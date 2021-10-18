FROM golang:1.17 as builder

ENV GOPROXY=https://goproxy.cn,direct 

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

FROM scratch

WORKDIR /app

COPY --from=builder /app/app .

ENV VERSION=lyb

EXPOSE 1024

ENTRYPOINT [ "/app/app" ]


