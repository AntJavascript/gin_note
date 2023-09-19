### gin_note

### docker部署腾讯云服务器

### 执行docker build
docker build  -t go-docker:v1 .

### 执行doker run
docker run -d -p 8080:8080 go-docker:v1

### 查看docker镜像
docker images

### 查看docker 是否在运行
docker ps

### 访问测试接口
127.0.0.1:8080/test

Error response from daemon: failed to create task for container: failed to create shim task: OCI runtime create failed: runc create failed: unable to start container process: exec: "./build/main-linux": permission denied: unknown.
