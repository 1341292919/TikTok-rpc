namespace go api.video

include"model.thrift"
//上传视频
struct PublishRequest{
    1: required string title(api.form="title"),
    2: required string description(api.form="description"),
    3: binary data (api.form="data"),//视频file
}

struct PublishResponse{
     1: model.BaseResp base,
}
//发布列表
struct QueryPublishListRequest{
     1: required i64 user_id(api.form="user_id"),
     2: required i64 page_size(api.form="page_size"),  //每一页的数量
     3: required i64 page_num(api.form="page_num"),   //页码
}

struct QueryPublishListResponse{
    1:model.BaseResp base,
    2:model.VideoList data,
}
//搜索视频
struct SearchVideoByKeywordRequest{
     1: required i64 page_size(api.form="page_size"),  //每一页的数量
     2: required i64 page_num(api.form="page_num"),   //页码
     3: required string keyword(api.form="keyword"), //关键词
     4: optional i64 from_date(api.form="from_date"),         //起始日期
     5: optional i64 to_date(api.form="to_date"),            //终止日期
     6:string username(api.form="username"),        //对应用户的视频
}

struct SearchVideoByKeywordResponse{
     1:model.BaseResp base,
     2:model.VideoList data,
}
//热门排行榜
struct GetPopularListRequest{
      1: required i64 page_size(api.form="page_size"),  //每一页的数量
      2: required i64 page_num(api.form="page_num"),   //页码
}
struct GetPopularListResponse{
     1:model.BaseResp base,
     2:model.VideoList data,
}
//视频流
struct VideoStreamRequest{
    1:optional i64 latest_time(api.query="latest_time"),
    2:required i64 page_num (api.query="page_num"),
    3:required i64 page_size (api.query="page_size")
}

struct VideoStreamResponse{
    1:model.BaseResp base,
    2:model.VideoList data,
}
service VideoService{
    PublishResponse PublishVideo(1:PublishRequest req) (api.post="/video/publish"),
    QueryPublishListResponse QueryList(1:QueryPublishListRequest req)(api.get="/video/list"),
    SearchVideoByKeywordResponse SearchVideo(1:SearchVideoByKeywordRequest req)(api.post="/video/search"),
    GetPopularListResponse GetPopularVideo(1:GetPopularListRequest req)(api.get="/video/popular"),
    VideoStreamResponse GetVideoStream(1:VideoStreamRequest req)(api.get="/video/feed")
}