FROM alpine

# 工作目录路径
WORKDIR /app

COPY main ./

# 端口号
EXPOSE 8080

# run
ENTRYPOINT ["/app/main"]
