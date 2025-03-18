namespace go model

struct User{
    1: required string id,           //用户id
    2: required string username,  //用户名
    3: required string avatar_url
    4: required string created_at
    5: required string updated_at
    6: required string deleted_at
}

struct SimpleUser{
    1: required string id,           //用户id
    2: required string username,  //用户名
    3: required string avatar_url
}

struct Video{
    1: required string id,            //视频id
    2: required string user_id,
    3: required string video_url,     //视频文件链接
    4: required string cover_url ,    //封面链接
    5: required string title ,//视频标题
    6: required string description ,
    7: required i64 visit_count,
    8: required i64 like_count,
    9: required i64 comment_count,
    10: required string created_at
    11: required string updated_at
    12: required string deleted_at
}
struct Comment{
    1: required string id,//评论id
    2: required string user_id,
    3: required string video_id,
    4: required string content,
    5: required string created_at
    6: required string updated_at
    7: required string deleted_at
}

struct SimpleUserList{
     1: required list<SimpleUser> items,
     2: required i64 total,          //总数
}

struct VideoList{
    1: required list<Video> items,   //视频列表
    2: required i64 total,          //总数
}
struct CommentList{
    1: required list<Comment> items,   //评论列表
    2: required i64 total,          //总数
}

struct UserList{
    1: required list<User> items,
    2: required i64 total,
}
struct MFAMessage{
    1:required string secret,
    2:required string qrcode,
}
struct BaseResp {
    1: required i64 code,          //请求返回的状态码
    2: required string msg,        //返回的消息
}
