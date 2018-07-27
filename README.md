## 微服务demo
---
### gateway   
    是一个基于HTTP协议的restful的API网关。可以作为统一的API接入层,请求通过grpc传到后端服务,对后端服务做负载均衡,服务发现
    

### service   
    后端服务,接收gateway过来的请求处理,后端服务器启动时注册服务到etcd   


