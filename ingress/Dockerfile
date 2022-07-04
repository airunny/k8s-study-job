FROM golang:1.17
WORKDIR /app
COPY app app
ENV VERSION=v0.0.1
ENV TZ=Asia/Shanghai
EXPOSE 4000
ENTRYPOINT [ "./app" ]