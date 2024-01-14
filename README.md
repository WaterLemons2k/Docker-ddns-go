# Docker-ddns-go

Minimal Docker Image for [ddns-go](https://github.com/jeessy2/ddns-go).  
Version: v6.0.2

English | [简体中文](https://github.com/WaterLemons2k/Docker-ddns-go/blob/main/README.zh-CN.md)

## Why?

The [Official Image](https://hub.docker.com/r/jeessy/ddns-go) is based on [alpine](https://hub.docker.com/_/alpine) and will not be changed due to the requirements for debugging, getting IP by command, etc. (for more infomation, see [jeessy2/ddns-go#340](https://github.com/jeessy2/ddns-go/pull/340))

Therefore, a minimal image based on [scratch](https://hub.docker.com/_/scratch) has been created, containing only ddns-go, the timezone and certificates.

If you need the minimum image size and don't have the above requirements, you can try this image.

## Usage

- Mount the host directory, use the docker host mode. You can replace `/opt/ddns-go` with any directory on your host, the configuration file is a hidden file

  ```bash
  docker run -d --name ddns-go --restart=always --net=host -v /opt/ddns-go:/root waterlemons2k/ddns-go
  ```

- Open `http://DOCKER_IP:9876` in the browser, modify your configuration

- [Optional] Support startup with parameters `-l`listen address `-f`Sync frequency(seconds)

  ```bash
  docker run -d --name ddns-go --restart=always --net=host -v /opt/ddns-go:/root waterlemons2k/ddns-go -l :9877 -f 600
  ```

- [Optional] Without using docker host mode

  ```bash
  docker run -d --name ddns-go --restart=always -p 9876:9876 -v /opt/ddns-go:/root waterlemons2k/ddns-go
  ```