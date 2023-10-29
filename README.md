# Docker-ddns-go
[ddns-go](https://github.com/jeessy2/ddns-go) 的 Docker 镜像  
版本：v5.6.4
# 使用
- 挂载主机目录, 使用docker host模式。可把 `/opt/ddns-go` 替换为你主机任意目录, 配置文件为隐藏文件

  ```bash
  docker run -d --name ddns-go --restart=always --net=host -v /opt/ddns-go:/root waterlemons2k/ddns-go
  ```

- 在浏览器中打开`http://主机IP:9876`，修改你的配置，成功

- [可选] 支持启动带参数 `-l`监听地址 `-f`间隔时间(秒)

  ```bash
  docker run -d --name ddns-go --restart=always --net=host -v /opt/ddns-go:/root waterlemons2k/ddns-go -l :9877 -f 600
  ```

- [可选] 不使用docker host模式

  ```bash
  docker run -d --name ddns-go --restart=always -p 9876:9876 -v /opt/ddns-go:/root waterlemons2k/ddns-go
  ```