FROM golang:1.16-alpine AS build
RUN apk add --no-cache git
WORKDIR /project
RUN git clone https://github.com/airunny/k8s-study-job && \
    cd k8s-study-job/one && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

FROM alpine:3.13.6
WORKDIR /app
COPY --from=build /project/k8s-study-job/one/app app
ENV VERSION=v0.0.1
ENV TZ=Asia/Shanghai
EXPOSE 1024
ENTRYPOINT [ "./app" ]

