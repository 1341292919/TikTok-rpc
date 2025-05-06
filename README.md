# West2-online(learn-go考核5)

## 考核内容
[考核内容参考](https://github.com/west2-online/learn-go/blob/main/docs/5(2025)-微服务.md)

### 程序部署
请参考[docker部署](docs/deploy)

## 飞书
[文档](https://vcn9ra8gf7nh.feishu.cn/wiki/L8x3w7M2BioRRNkecYYcTnWJnRX)

## 接口实现

| 已实现接口                   | 备注         | 已实现接口                   | 备注         |
|-------------------------|------------|-------------------------|------------|
| POST /user/register     | 用户注册       | POST /user/login        | 用户登录       |
| GET /user/info          | 用户信息       | PUT /user/avatar/upload | 上传头像       |
| GET /auth/mfa/qrcode    | 获取MFA绑定二维码 | POST /auth/mfa/bind     | 绑定MFA      |
| POST /auth/mfa/status   | 启用（关闭）MFA  | GET /video/feed/        | 获取视频流      |
| POST /video/publish     | 投稿         | GET /video/list         | 发布列表       |
| GET /video/popular      | 热门视频排行榜    | POST /video/search      | 搜索视频       |
| POST /like/action       | 点赞         | GET /like/list          | 点赞列表       |
| POST /comment/publish   | 评论         | DEL /comment/delete     | 删除评论       |
| GET /comment/list       | 评论列表       | WS /ws                  | Websock聊天  |




