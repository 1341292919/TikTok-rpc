package mysql

import (
	"TikTok-rpc/app/websocket/domain/model"
	"TikTok-rpc/app/websocket/domain/repository"
	"TikTok-rpc/pkg/constants"
	"TikTok-rpc/pkg/errno"
	"context"
	"errors"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
)

type websocketDB struct {
	client *gorm.DB
}

func NewWebsocketDB(client *gorm.DB) repository.WebsocketDB {
	return &websocketDB{client: client}
}

func (db *websocketDB) CreateNewMessage(ctx context.Context, m *model.Message) error {
	var messageResp *Message
	messageResp = &Message{
		UserId:   m.UserId,
		Type:     m.Type,
		TargetId: m.TargetId,
		Content:  m.Content,
		Status:   m.Status,
	}
	err := db.client.WithContext(ctx).
		Table(constants.TableMessage).
		Create(&messageResp).
		Error
	if err != nil {
		return errno.NewErrNo(errno.InternalDatabaseErrorCode, "CreateNewMessage:"+err.Error())
	}
	return nil
}

func (db *websocketDB) QueryTargetMessage(ctx context.Context, targetid int64) ([]*model.Message, error) {
	var messageResp []*Message
	var count int64
	err := db.client.Transaction(func(tx *gorm.DB) error {
		err := db.client.WithContext(ctx).
			Table(constants.TableMessage).
			Where("target_id=?", targetid).
			Count(&count).
			Find(&messageResp).
			Error
		if err != nil {
			return nil
		}
		updateId := make([]int64, 0)
		for _, message := range messageResp {
			hlog.Info(message.Status)
			if message.Status == 0 {
				updateId = append(updateId, message.Id)
			}
		}
		err = updateMessageStatus(tx, updateId, 1)
		if err != nil {
			return errno.NewErrNo(errno.InternalDatabaseErrorCode, "QueryTargetMessage:"+err.Error())
		}
		return nil
	})
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, "QueryTargetMessage:"+err.Error())
	}
	return buildMessageList(messageResp), nil
}

func (db *websocketDB) QueryPrivateMessage(ctx context.Context, req *model.ChatReq) ([]*model.Message, error) {
	var messageResp []*Message
	var count int64
	err := db.client.WithContext(ctx).
		Table(constants.TableMessage).
		Where("target_id=? and user_id = ?", req.TargetId, req.UserId).
		Where("type = ?", 0).
		Limit(int(req.PageSize)).
		Offset(int((req.PageNum - 1) * req.PageSize)).
		Count(&count).
		Find(&messageResp).
		Error
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, "QueryPrivateMessage:"+err.Error())
	}
	return buildMessageList(messageResp), nil
}

func (db *websocketDB) QueryGroupMessage(ctx context.Context, req *model.ChatReq) ([]*model.Message, error) {
	var messageResp []*Message
	var count int64
	err := db.client.WithContext(ctx).
		Table(constants.TableMessage).
		Where("target_id=?", req.TargetId).
		Where("type = ?", 1).
		Limit(int(req.PageSize)).
		Offset(int((req.PageNum - 1) * req.PageSize)).
		Count(&count).
		Find(&messageResp).
		Error
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, "QueryPrivateMessage:"+err.Error())
	}
	return buildMessageList(messageResp), nil
}
func (db *websocketDB) UpdateMessageList(ctx context.Context, messages []*model.Message) error {
	// 1. 获取数据库最新消息时间
	var latestMsg Message
	err := db.client.WithContext(ctx).
		Table(constants.TableMessage).
		Order("created_at DESC").
		First(&latestMsg).
		Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errno.NewErrNo(errno.InternalDatabaseErrorCode, "UpdateMessageList: "+err.Error())
	}

	latestTime := latestMsg.CreatedAt.Unix()

	// 2. 筛选需要插入的消息
	var toInsert []*Message
	for _, m := range messages {
		if m.CreatedAT > latestTime { // 只处理比数据库更新的消息
			toInsert = append(toInsert, &Message{
				UserId:    m.UserId,
				Type:      m.Type,
				TargetId:  m.TargetId,
				Content:   m.Content,
				Status:    m.Status,
				CreatedAt: TimestampToTime(m.CreatedAT),
			})
		}
	}

	// 3. 批量插入（使用事务）
	if len(toInsert) == 0 {
		return nil
	}

	return db.client.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(constants.TableMessage).Create(&toInsert).Error; err != nil {
			return errno.NewErrNo(errno.InternalDatabaseErrorCode, "UpdateMessageList: "+err.Error())
		}
		return nil
	})
}

func TimestampToTime(timestamp int64) time.Time {
	// 根据时间戳长度判断单位
	switch {
	case timestamp > 1e18: // 纳秒级(19位)
		return time.Unix(0, timestamp)
	case timestamp > 1e15: // 微秒级(16位)
		return time.Unix(0, timestamp*1000)
	case timestamp > 1e12: // 毫秒级(13位)
		return time.Unix(0, timestamp*1e6)
	default: // 秒级(10位)
		return time.Unix(timestamp, 0)
	}
}
func buildMessage(data *Message) *model.Message {
	return &model.Message{
		UserId:    data.UserId,
		Type:      data.Type,
		TargetId:  data.TargetId,
		Content:   data.Content,
		Id:        data.Id,
		Status:    data.Status,
		CreatedAT: data.CreatedAt.Unix(),
	}
}
func buildMessageList(data []*Message) []*model.Message {
	messageList := make([]*model.Message, 0)
	for _, m := range data {
		messageList = append(messageList, buildMessage(m))
	}
	return messageList
}
func updateMessageStatus(tx *gorm.DB, messageIDs []int64, status int) error {
	return tx.Table(constants.TableMessage).
		Where("id IN (?)", messageIDs).
		Update("status", status).
		Error
}
