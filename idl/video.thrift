namespace go video

include"model.thrift"
//上传视频
struct PublishRequest{
    1: required string title,
    2: required string description,
    3: binary data ,//视频file
    4: required i64 user_id,
}

struct PublishResponse{
     1: model.BaseResp base,
}
//发布列表
struct QueryPublishListRequest{
     1: required i64 user_id,
     2: required i64 page_size,  //每一页的数量
     3: required i64 page_num,   //页码
}

struct QueryPublishListResponse{
    1:model.BaseResp base,
    2:model.VideoList data,
}
//搜索视频
struct SearchVideoByKeywordRequest{
     1: required i64 page_size  //每一页的数量
     2: required i64 page_num,   //页码
     3: required string keyword, //关键词
     4: optional i64 from_date,         //起始日期
     5: optional i64 to_date,            //终止日期
     6:string username,        //对应用户的视频
}

struct SearchVideoByKeywordResponse{
     1:model.BaseResp base,
     2:model.VideoList data,
}
//热门排行榜
struct GetPopularListRequest{
      1: required i64 page_size,  //每一页的数量
      2: required i64 page_num,   //页码
}
struct GetPopularListResponse{
     1:model.BaseResp base,
     2:model.VideoList data,
}
//视频流
struct VideoStreamRequest{
    1:optional i64 latest_time,
    2:required i64 page_num,
    3:required i64 page_size
}

struct VideoStreamResponse{
    1:model.BaseResp base,
    2:model.VideoList data,
}
service VideoService{
    PublishResponse PublishVideo(1:PublishRequest req),
    QueryPublishListResponse QueryList(1:QueryPublishListRequest req),
    SearchVideoByKeywordResponse SearchVideo(1:SearchVideoByKeywordRequest req),
    GetPopularListResponse GetPopularVideo(1:GetPopularListRequest req),
    VideoStreamResponse GetVideoStream(1:VideoStreamRequest req)
}