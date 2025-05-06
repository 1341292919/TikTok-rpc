# 🛠️ 部署

程序使用以下端口，请先确保这些端口没有被占用：
10000、10001、10002、10003、10004、10005、10006
或者你可以通过修改 `config/config.yaml` 文件来更改端口。
#### 配置文件
请通过参考[config](docs/config.md) 在程序部署前完善配置文件`config/config.yaml`

#### Docker 容器

1. **安装 Docker：**
   首先[下载并安装 Docker](https://docs.docker.com/get-docker/)。

2. **下载仓库：**
   使用以下命令克隆仓库：
```sh
git clone https://github.com/1341292919/TikTok-rpc
```
3. **必要的运行环境环境：**
   可以通过以下命令搭建必要的环境：
```sh
    make env-up
```
  也可以根据docker/docker-compose.yaml搭建，并修改相应配置文件

4. **部署与构建：**
   使用以下命令构建镜像并启动相应服务：
```sh
cd Tiktok-rpc
make builid -% # 请在%处补充相应的服务名称，可在makefile中查看服务名
```

