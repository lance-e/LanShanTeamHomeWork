# 第七次课后作业

### 架构：

采用 微服务架构 ，在http服务端 远程调用(实际上是在同一个项目目录中) rpc服务端的代码，
http服务端结构：
rpc服务端结构：

### 不足：

1.读取配置文件步骤未解耦，主要是因为仅使用了mysql数据库
2.代码冗余，
3.rpc的客户端和服务端同时写在一个项目中，重名率非常高，代码可读性不好，不好维护

### 接口文档

1.注册：
请求:
postForm :username
postForm :password
响应：
出错：
code：1000
status:1000
error:string(error的信息)
失败：
code：1001
status：1001
message：register failed
成功：
code：1002
status:1002
message:register success
