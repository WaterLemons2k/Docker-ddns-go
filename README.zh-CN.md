# Docker-ddns-go

[ddns-go](https://github.com/jeessy2/ddns-go) 的最小 Docker 镜像  
版本：v6.12.2

[English](https://github.com/WaterLemons2k/Docker-ddns-go/blob/main/README.md) | 简体中文

## 为什么？

[官方镜像](https://hub.docker.com/r/jeessy/ddns-go)基于 [alpine](https://hub.docker.com/_/alpine)，并且由于调试、通过命令获取 IP 等需求，不会改变（有关更多信息，请参阅 [jeessy2/ddns-go#340](https://github.com/jeessy2/ddns-go/pull/340)）。

因此，创建了基于 [scratch](https://hub.docker.com/_/scratch) 的最小镜像，其中只包含 ddns-go、时区和证书。

如果你需要最小的镜像体积且没有上述需求，可以尝试使用此镜像。

## 使用

- 挂载主机目录, 使用 docker host 模式。可把 `/opt/ddns-go` 替换为你主机任意目录, 配置文件为隐藏文件

  ```bash
  docker run -d --name ddns-go --restart=always --net=host -v /opt/ddns-go:/root waterlemons2k/ddns-go
  ```

- 在浏览器中打开`http://主机IP:9876`，并修改你的配置

- [可选] 支持启动带参数 `-l`监听地址 `-f`间隔时间(秒)

  ```bash
  docker run -d --name ddns-go --restart=always --net=host -v /opt/ddns-go:/root waterlemons2k/ddns-go -l :9877 -f 600
  ```

- [可选] 不使用 docker host 模式

  ```bash
  docker run -d --name ddns-go --restart=always -p 9876:9876 -v /opt/ddns-go:/root waterlemons2k/ddns-go
  ```
