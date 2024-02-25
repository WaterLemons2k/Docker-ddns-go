FROM scratch
LABEL maintainer="waterlemons2k <docker@waterlemons2k.com>"

COPY ca-certificates.crt /etc/ssl/certs/
COPY ddns-go /app/
COPY Shanghai /usr/share/zoneinfo/Asia/

ENV TZ=Asia/Shanghai
EXPOSE 9876

ENTRYPOINT ["/app/ddns-go"]
CMD ["-l", ":9876", "-f", "300"]
