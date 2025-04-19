namespace go interact

include "model.thrift"
//点赞
struct LikeRequest{
    //两者必须存在其一
    1:optional i64 video_id
    2:optional i64 comment_id,
    3:required i64 action_type,//1-点赞，2-取消点赞
    4:required i64 user_id
}
struct LikeResponse{
    1:model.BaseResp base,
}
//用户的点赞列表
struct QueryLikeListRequest{
    1:required i64 user_id,
    2: required i64 page_size,  //每一页的数量
    3: required i64 page_num   //页码
}

struct QueryLikeListResponse{
    1:model.BaseResp base,
    2:optional model.VideoList data,
}
//评论
struct CommentRequest{
    //两者必须存在其一
    1:optional i64 video_id,
    2:optional i64 comment_id,
    3:required string content,
    4:required i64 user_id
}

struct CommentResponse{
    1:model.BaseResp base,
    2:optional i64 comment_id,
}
//查看评论列表
struct QueryCommentListRequest{
    1:optional i64 video_id,
    2:optional i64 comment_id,
    3: required i64 page_size,  //每一页的数量
    4: required i64 page_num,   //页码
}

struct QueryCommentListResponse{
    1:model.BaseResp base,
    2:optional model.CommentList data,
}
//删除评论
struct DeleteCommentRequest{
    1:optional i64 video_id,
    2:optional i64 comment_id,
    3:required i64 user_id
}

struct DeleteCommentResponse{
    1:model.BaseResp base,
}

service InteractService{
   LikeResponse Like(1:LikeRequest req),
   QueryLikeListResponse QueryLikeList(1:QueryLikeListRequest req),
   CommentResponse Comment(1:CommentRequest req),
   QueryCommentListResponse QueryCommentList(1:QueryCommentListRequest req),
   DeleteCommentResponse DeleteComment(1:DeleteCommentRequest req),
}