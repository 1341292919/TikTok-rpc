# TikTok
基于golang实现的分布式架构的视频平台，支持点赞评论发布聊天等视频功能

## 系统架构
基于微服务架构，用hertz框架搭建gateway收取http请求，划分为user、video、interact和web四个核心服务，通过Gateway统一入口处理请求。服务间通过Kitex RPC框架通信，使用etcd进行服务发现。

## 技术栈

### 核心框架
- **编程语言**: Go 1.21+
- **RPC框架**: Kitex
- **HTTP框架**: Hertz
- **接口定义**: Thrift

### 数据存储
- **关系型数据库**: MySQL 9.0.1
- **缓存数据库**: Redis
- **时序数据库**: VictoriaMetrics

### 消息队列
- **流处理平台**: Kafka
- **可视化管理**: Kafka-Ui

### 服务治理
- **服务发现与配置**: etcd
- **分布式追踪**: Jaeger + OpenTelemetry

### 监控体系
- **指标收集**: Prometheus
- **数据可视化**: Grafana
- **数据采集器**: Node Exporter, Redis Exporter, MySQL Exporter

### 基础设施
- **容器化**: Docker + Docker Compose
- **容器监控**: cAdvisor


## 服务接口
[接口文档参考](https://apifox.com/apidoc/shared/d0864798-bfd5-4288-bc37-7d802a9f52e3)

实现的接口、请参考[接口说明](docs/interface.md)

## 核心功能实现
### 缓存与持久化
- 点赞信息、访问量采取redis缓存，mysql存储异步进行
- 热门排行榜从维护的redis Sorted Set快速访问
- 视频与图片等数据对象存储到云端
- web聊天通信在线信息优先存储进redis

### 异步处理
- 点赞与评论接口添加消息队列kafka异步处理请求
- 另启线程消费点赞评论这些高频信息
- 懒加载处理，有效节省算力与存储成本

### 访问令牌
- hertz中间件颁发Access-Token、 Refresh-Token双token
- Refresh-Token可在规定时间内刷新Access-Token

### 基础设施服务
- etcd作为分布式键值存储，用于服务发现和配置管理
- OpenTelemetry可观测性框架，统一收集追踪、指标和日志数据 
- Jaeger参与构建分布式追踪系统，用于性能监控和故障排查
- Prometheus监控系统，收集和存储时间序列数据
- Grafana搭建可视化平台，展示监控指标和仪表盘
- 多种采集器Exporter：Node Exporter、Redis Exporter、MySQL Exporter负责采集各类指标

## 文档说明-飞书
关于业务处理逻辑、redis、kafka的运用详细、监测与链路追踪体系可见文档
[文档](https://vcn9ra8gf7nh.feishu.cn/wiki/TVjswJ5uLi0yLekhLvDcG27ineZ) 

## 开发与部署
项目使用Makefile自动化构建，支持单独编译和部署各个服务。环境依赖通过Docker Compose管理，包括MySQL、Redis、etcd、Elasticsearch和Kafka等。

### 环境要求
- Docker 20.10+
- Docker Compose v2.0+
- Go 1.21+
- Make
### 快速开始
请参考 [quick start](docs/quick-start.md)

### 程序部署
请参考[docker部署](docs/deploy.md)
该项目也成功在kubernetes上完成部署，详细可见上述飞书文档。




