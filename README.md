### gin_note

### docker部署腾讯云服务器

### 执行docker build
docker build  -t go-docker:latest .

### 执行doker run
docker run -d -p 8080:8080 go-docker:latest

### 查看docker镜像
docker images

### 查看docker 是否在运行
docker ps
