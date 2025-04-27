namespace go api.socialize

include "model.thrift"

struct FollowRequest{
    1:required i64 target_user_id(api.form="to_user_id"),
    2:required i64 action_type(api.form="action_type"),   //0关注 1取关
}
struct FollowResponse{
    1:model.BaseResp base,
}

struct QueryFollowListRequest{  //查看对应id的关注
     1: required i64 user_id(api.form="user_id"),
    2: required i64 page_size(api.form="page_size"),  //每一页的数量
    3: required i64 page_num(api.form="page_num"),   //页码
}
struct QueryFollowListResponse{
      1:model.BaseResp base,
      2:model.SimpleUserList data,
}
struct QueryFansListRequest{ //查看指定id的粉丝
     1: required i64 user_id(api.form="user_id"),
    2: required i64 page_size(api.form="page_size"),  //每一页的数量
    3: required i64 page_num(api.form="page_num"),   //页码
}
struct QueryFansListResponse{
      1:model.BaseResp base,
      2:model.SimpleUserList data,
}
struct QueryFriendListRequest{
    1: required i64 page_size(api.form="page_size"),  //每一页的数量
    2: required i64 page_num(api.form="page_num"),   //页码
}
struct QueryFriendListResponse{
      1:model.BaseResp base,
      2:model.SimpleUserList data,
}

service SocializeService{
    FollowResponse Follow(1:FollowRequest req)(api.post="/relation/action"),
    QueryFollowListResponse QueryFollowList(1:QueryFollowListRequest req)(api.get="/following/list"),
    QueryFansListResponse QueryFansList(1:QueryFansListRequest req)(api.get="/follower/list"),
    QueryFriendListResponse QueryFriendList(1:QueryFriendListRequest req)(api.get="/friends/list"),
}
