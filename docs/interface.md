1. UserService

| 实现接口         | 路径             | 说明           |
|--------------|------------------|----------------|
| Register     |  POST /user/register        |         用户注册       |
| Login        | POST /user/login          |       用户登录          |
| GetUserInfo  | GET /user/info        |        用户信息            |
| UploadAvatar | PUT /user/avatar/upload         |          上传头像       |
| GetQrcode    |GET /auth/mfa/qrcode |        获取MFA绑定二维码        |
| BindQrcode   | POST /auth/mfa/bind         |        绑定MFA            |

2. VideoService

| 实现接口          | 路径             | 说明           |
|---------------|------------------|----------------|
| VideoStream   |   GET /video/feed/        | 获取视频流 |
| Publish       |POST /video/publish     | 投稿         | 
| PublishedList | GET /video/list         | 发布列表       |
| PopularVideo  | GET /video/popular      | 热门视频排行榜 |
| Search        |POST /video/search      | 搜索视频       |

3. InteractService

| 实现接口          | 路径             | 说明     |
|---------------|------------------|--------|
| Like          |  POST /like/action       | 点赞     | 
| GetLikes      | GET /like/list          | 查看点赞列表 |
| Comment       | POST /comment/publish   | 评论     |
| DeleteComment |DEL /comment/delete     | 删除评论   |
| CommentList   |GET /comment/list       | 评论列表   |

4. WebsocketService
   借助Hertz 框架内置 WebSocket 的实现方法
   参考 websocket 的聊天功能，基于聊天的实时性，请使用 Redis + MySQL 方式