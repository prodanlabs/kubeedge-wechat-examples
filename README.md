# kubeedge-wechat-examples

> 参考的源项目：https://github.com/kubeedge/examples

##### 与原项目不同之处：
- mqtt 库使用的是`eclipse/paho.mqtt.golang` ;
- web 框架使用 `Iris + gRPC` ;
- 边缘端服务增加上传数据到云端服务的功能，模拟数据协同;


##### 先决条件：
- `kubernetes KubeEdge` 环境；
- 树莓派；
- 微信公众号；
- 扬声器；
- 因为 `gRPC` 需要 `TLS` 证书认证，且微信回调接口识别不了自签证书，故需要 `SSL` 证书颁发机构的证书；