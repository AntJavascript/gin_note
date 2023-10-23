### gin_note

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

// int到string
str1 := strconv.Itoa(int)

// int64到string
str2 := strconv.FormatInt(int64,10）//10为进制

// string到int
intStr, err := strconv.Atoi(string)

// string到int64  
int64Str, err := strconv.ParseInt(string, 10, 64)//10为进制，64为bit位
