FROM golang:1.19

# 工作目录路径
WORKDIR /app

COPY linux-main ./

# 端口号
EXPOSE 8080

# run
CMD ["/build/main"]
