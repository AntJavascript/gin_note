FROM goalng:1.2-app

# 工作目录路径
WORKDIR /build

# 端口号
EXPOSE 8080

# 入口文件
ENTRYPOINT ["/build/main"]
