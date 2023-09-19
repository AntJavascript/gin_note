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


### docker部署linux云服务器（文章参考）
https://blog.csdn.net/weixin_51261234/article/details/125906988?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522169501917316800180628651%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fall.%2522%257D&request_id=169501917316800180628651&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~first_rank_ecpm_v1~times_rank-17-125906988-null-null.142^v94^insert_down28v1&utm_term=golang%E9%A1%B9%E7%9B%AE%E6%89%93%E5%8C%85%20docker%E9%83%A8%E7%BD%B2&spm=1018.2226.3001.4187

https://blog.csdn.net/qq_46028694/article/details/131033905?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522169510447216800226546202%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fall.%2522%257D&request_id=169510447216800226546202&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~first_rank_ecpm_v1~times_rank-1-131033905-null-null.142^v94^insert_down28v1&utm_term=docker%E9%83%A8%E7%BD%B2go%E4%BA%8C%E8%BF%9B%E5%88%B6%E6%96%87%E4%BB%B6&spm=1018.2226.3001.4187

https://blog.csdn.net/qq_28478281/article/details/84137836?utm_medium=distribute.pc_relevant.none-task-blog-2~default~baidujs_utm_term~default-0-84137836-blog-131033905.235^v38^pc_relevant_sort_base3&spm=1001.2101.3001.4242.1&utm_relevant_index=3
