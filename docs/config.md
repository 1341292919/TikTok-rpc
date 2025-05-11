#环境变量

#

| 变量             | 环境变量名          | 描述          |
|----------------|----------------|-------------|
| MySQL.UserName | MYSQL_USERNAME | MySQL数据库的用户名 |
| MySQL.PassWord | MYSQL_PASSWORD | MySQL数据库的密码 |
| MySQL.Addr     | -              | MySQL数据库的地址 |
| MySQL.Database | MYSQL_DATABASE | MySQL数据库的名称 |
| Redis.UserName | REDIS_USERNAME | Redis数据库的用户名 |
| Redis.PassWord | REDIS_PASSWORD | Redis数据库的密码 |
| Redis.Addr     | -              | Redis数据库的地址 |
| Etcd.Addr      | ETCD_LISTEN_CLIENT_URLS            | Etcd地址      |
| Oss.Bucket     | -              | 七牛云存储的Bucket名 |
| Oss.AccessKey  | -              | 七牛云存储的访问密钥  |
| Oss.SecretKey  | -              | 七牛云存储的秘密密钥  |
| Oss.Domain     | -              | 七牛云存储的域名    |



请参照下文配置Oss部分
Oss使用七牛云对象存储
请根据下表配置 Oss.region：

|  storage.Zone  | region |
|---------------|--------|
| Zone_z0 - 华东（浙江） | z0     |
|Zone_z1 - 华北（北京） | z1     |
| Zone_z2 - 华南（广东）    | z2     | 
| Zone_na0 - 北美（洛杉矶）| na0    | 
|Zone_as0 - 东南亚（新加坡））| as0    | 

Oss.domain:
请补充七牛云对应桶的CDN 加速域名或 源站域名