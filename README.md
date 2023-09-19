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


### 部署云服务器
https://blog.csdn.net/weixin_51261234/article/details/125906988?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522169501917316800180628651%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fall.%2522%257D&request_id=169501917316800180628651&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~first_rank_ecpm_v1~times_rank-17-125906988-null-null.142^v94^insert_down28v1&utm_term=golang%E9%A1%B9%E7%9B%AE%E6%89%93%E5%8C%85%20docker%E9%83%A8%E7%BD%B2&spm=1018.2226.3001.4187
