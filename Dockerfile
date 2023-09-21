FROM golang:latest

LABEL cjl "cjl_spy@163.com"

WORKDIR /root

ADD ./build/main-linux ./build/main-linux


EXPOSE 8080

ENTRYPOINT  ["./build/main-linux"]
