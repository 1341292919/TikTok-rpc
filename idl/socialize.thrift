namespace go socialize

include "model.thrift"

struct FollowRequest{
    1:required i64 target_user_id,
    2:required i64 action_type,   //0关注 1取关
    3:required i64 user_id,
}
struct FollowResponse{
    1:model.BaseResp base,
}

struct QueryFollowListRequest{  //查看对应id的关注
     1: required i64 user_id,
    2: required i64 page_size,  //每一页的数量
    3: required i64 page_num,   //页码
}
struct QueryFollowListResponse{
      1:model.BaseResp base,
      2:optional model.SimpleUserList data,
}
struct QueryFansListRequest{ //查看指定id的粉丝
     1: required i64 user_id,
    2: required i64 page_size,  //每一页的数量
    3: required i64 page_num,   //页码
}
struct QueryFansListResponse{
      1:model.BaseResp base,
      2:optional model.SimpleUserList data,
}
struct QueryFriendListRequest{
    1: required i64 page_size,  //每一页的数量
    2: required i64 page_num,   //页码
}
struct QueryFriendListResponse{
      1:model.BaseResp base,
      2: optional model.SimpleUserList data,
}

service SocializeService{
    FollowResponse Follow(1:FollowRequest req),
    QueryFollowListResponse QueryFollowList(1:QueryFollowListRequest req),
    QueryFansListResponse  QueryFansList(1:QueryFansListRequest req),
    QueryFriendListResponse QueryFriendList(1:QueryFriendListRequest req),
}
