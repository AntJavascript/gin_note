FROM alpine

# 工作目录路径
WORKDIR /app

ADD ./main ./main

# 端口号
EXPOSE 8080

# run
ENTRYPOINT ["./main"]
