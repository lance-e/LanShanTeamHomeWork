# 第七次课后作业

### 架构：

采用 微服务架构 ，在http服务端 远程调用(实际上是在同一个项目目录中) rpc服务端的代码，
~~在本次作业中使用体验下来，感觉代码量还多了~~

### 不足：

1.读取配置文件步骤未解耦，主要是因为仅使用了mysql数据库
2.代码冗余，
3.rpc的客户端和服务端同时写在一个项目中，重名率非常高，代码可读性不好，不好维护
4.越写越乱了，哈哈哈，光是登陆注册服务，就要分别启动服务，推测是protobuf文件写错了，代码可读性-1
~~5.为了简化操作，加好友的逻辑：在relationShip表中关系为双向（你follow我，我follow你）
6.用户查询出来的好友列表显示的是uuid
7.把搜索用户砍了，因为架构设计和表设计限制~~
8.好嘛，现在把好友模块全删了，架构和数据库设计缺陷，导致后续开发无法下手，可读性非常不好
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
