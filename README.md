# kubeedge-wechat-examples

> 参考的源项目：https://github.com/kubeedge/examples

##### 与原项目不同之处：
- mqtt 库使用的是`eclipse/paho.mqtt.golang` ;
- web 框架使用 `Iris + gRPC` ;
- 边缘端服务增加上传数据到云端服务的功能，模拟数据协同;

<img src="https://github.com/prodanlabs/kubeedge-wechat-examples/blob/main/image/2021-01-24_14-32.png" width="960">

##### 先决条件：
- `kubernetes KubeEdge` 环境；
- 树莓派；
- 微信公众号；
- 扬声器；
- 因为 `gRPC` 需要 `TLS` 证书认证，且微信回调接口识别不了自签证书，故需要 `SSL` 证书颁发机构的证书；



##### 云端服务

| 变量名                    |      默认值      | 说明                                                         |
| :------------------------ | :--------------: | :----------------------------------------------------------- |
| `SERVER_ADDR_PORT`        |  `0.0.0.0:443`   | 服务监听的`IP+端口`                                          |
| `WECHAT_APP_ID`           |      `xxx`       | 微信`AppID`                                                  |
| `WECHAT_APP_SECRET`       |      `xxx`       | 微信`AppSecret`                                              |
| `WECHAT_TOKEN`            |      `xxx`       | 微信`Token`                                                  |
| `WECHAT_ENCODING_AES_KEY` |      `xxx`       | 微信`EncodingAESKey`                                         |
| `IN_CLUSTER`              | `out-of-cluster` | 使用默认值时，通过 `~/.kube/config` 连接`k8s`，主要用于开发调试 |
| `DEVICE_ID`               |   `speaker-01`   | `kubeedge-speaker.yaml` 文件定义的`name` ,非默认值，边缘端服务订阅的`topic` 需要对应 |
| `DEVICE_NAMESPACE`        |    `default`     | `kubeedge-speaker.yaml` 文件定义的`namespace` ,非默认值，边缘端服务订阅的`topic` 需要对应` |

##### 边缘端服务

1. 音乐目录：`/home/pi/music/*.mp3`
2. 需要 `omxplayer` 播放器
3. `gRPC` 服务端地址需修改 `edgecore-client-pi/utils/grpc_client.go`

##### 编译

```sh
# 生成 grpc 代码
make gen
# 编译云端服务
make build
# 交叉编译边缘端服务
make build-client-arm
```
> 其他的看Makefile文件

感兴趣的朋友，可以去公众号看看效果~~~

<img src="https://github.com/prodanlabs/kubeedge-wechat-examples/blob/main/image/weixin.png" width="460">