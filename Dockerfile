FROM alpine

# 工作目录路径
WORKDIR /app

ADD ./main ./main

EXPOSE 8080

# run
ENTRYPOINT ["./main"]
