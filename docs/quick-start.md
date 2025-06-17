# quickly start
## 依赖要求
### 环境要求
    Docker 20.10+
    Docker Compose v2.0+
    Go 1.21+
    Make

### 占用端口
请保证 [参考配置文件](../config/config.example.yaml) 中addr所指向的端口未被占用

或者请在需要改变处更改对应的addr

### 快速开始
1. 克隆项目
```bash
git clone https://github.com/1341292919/TikTok-rpc
cd TikTok-rpc
```

2. 配置环境变量
```bash
cp config/config.yaml.example config/config.yaml
```
除此之外请参考 [配置内容说明](config.md) 进行oss部分配置（用于文件对象存储）

3. 启动基础服务
```bash
make env-up 
# 启动 MySQL、Redis、Kafka、etcd
```
**必须安装 FFmpeg**（用于视频处理）：
  ```bash
  # Linux (Debian/Ubuntu)
  sudo apt install ffmpeg
  ```

4. 编译并启动服务
```
make run-user        # 启动用户服务
make run-video       # 启动视频服务
make run-interact    # 启动互动服务
make run-websocket      # 启动websocket服务
make run-gateway     # 启动网关服务

```  