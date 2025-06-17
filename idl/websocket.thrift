namespace go websocket

include "model.thrift"

struct AddMessageRequest{
    1:required string content,
    2:required i64 id,
    3:required i64 target_id,
    4:required i64 status,
    5:required i64 type,
}

struct AddMessageResponse{
    1:model.BaseResp base,
}

struct QueryOfflineMessageRequest{
    1:required i64 id,
}
struct QueryOfflineMessageResponse{
    1:model.BaseResp base,
    2: optional model.ChatMessageList data,
}
struct QueryPrivateHistoryMessageRequest{
    1:required i64 user_id,
    2:required i64 target_id,
    3:required i64 page_num,
    4:required i64 page_size,
}
struct QueryPrivateHistoryMessageResponse{
    1:model.BaseResp base,
    2:optional model.ChatMessageList data,
}
struct QueryGroupHistoryMessageRequest{
    1:required i64 user_id,
    2:required i64 target_id,
    3:required i64 page_num,
    4:required i64 page_size,
}
struct QueryGroupHistoryMessageResponse{
    1:model.BaseResp base,
    2:optional model.ChatMessageList data,
}
service WebsocketService{
    AddMessageResponse AddMessage (1:AddMessageRequest req),
    QueryOfflineMessageResponse QueryOfflineMessage (1:QueryOfflineMessageRequest req),
    QueryPrivateHistoryMessageResponse QueryPrivateHistoryMessage(1:QueryPrivateHistoryMessageRequest req),
    QueryGroupHistoryMessageResponse QueryGroupHistoryMessage(1:QueryGroupHistoryMessageRequest req),
}
