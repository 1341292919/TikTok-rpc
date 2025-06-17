package mq

import (
	"TikTok-rpc/app/interact/domain/model"
	"TikTok-rpc/pkg/constants"
	"TikTok-rpc/pkg/errno"
	"TikTok-rpc/pkg/kafka"
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"strconv"
	"strings"
)

func (ka *KafkaAdapter) sendLikeMessage(ctx context.Context, msg []*kafka.Message) error {
	var err error
	if !ka.done.Load() {
		err = ka.mq.SetWriter(constants.KafkaInteractLikeTopic, true)
		if err != nil {
			return err
		}
		ka.done.Swap(true)
	}
	errs := ka.mq.Send(ctx, constants.KafkaInteractLikeTopic, msg)
	if len(errs) != 0 {
		var errMsg string
		for _, e := range errs {
			errMsg = strings.Join([]string{errMsg, e.Error(), ";"}, "")
		}
		err = fmt.Errorf("mq.Send: send msg failed, errs: %v", errMsg)
		return err
	}
	return nil
}

func (ka *KafkaAdapter) SendLikeMessage(ctx context.Context, like *model.UserLike) error {
	v, err := sonic.Marshal(like)
	if err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "json turn userlike error"+err.Error())
	}
	msg := []*kafka.Message{
		{
			K: []byte(strconv.FormatInt(like.Uid%constants.KafkaInteractPartitionNum, 10)),
			V: v,
		},
	}
	if err = ka.sendLikeMessage(ctx, msg); err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "kafka send message error"+err.Error())
	}
	return nil
}

func (ka *KafkaAdapter) SendCommentMessage(ctx context.Context, comment *model.CommentMessage) error {
	v, err := sonic.Marshal(comment)
	if err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "json turn commentmessage error"+err.Error())
	}
	msg := []*kafka.Message{
		{
			K: []byte(strconv.FormatInt(comment.UId%constants.KafkaInteractPartitionNum, 10)),
			V: v,
		},
	}
	if err = ka.sendCommentMessage(ctx, msg); err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "kafka send message error"+err.Error())
	}
	return nil
}

func (ka *KafkaAdapter) sendCommentMessage(ctx context.Context, msg []*kafka.Message) error {
	var err error
	if !ka.done.Load() {
		err = ka.mq.SetWriter(constants.KafkaInteractCommentTopic, true)
		if err != nil {
			return err
		}
		ka.done.Swap(true)
	}
	errs := ka.mq.Send(ctx, constants.KafkaInteractCommentTopic, msg)
	if len(errs) != 0 {
		var errMsg string
		for _, e := range errs {
			errMsg = strings.Join([]string{errMsg, e.Error(), ";"}, "")
		}
		err = fmt.Errorf("mq.Send: send msg failed, errs: %v", errMsg)
		return err
	}
	return nil
}
