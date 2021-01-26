package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/kubeedge/kubeedge/cloud/pkg/devicecontroller/types"

	"github.com/prodanlabs/kubeedge-wechat-examples/edgecore-client-pi/utils"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Received message: %s messageid: %d from topic: %s\n", msg.Payload(), msg.MessageID(), msg.Topic())

	messages := &types.DeviceTwinDocument{}
	if err := json.Unmarshal(msg.Payload(), messages); err != nil {
		log.Printf("error: %s", err)
	}

	current := *messages.Twin["track"].CurrentState.Expected.Value
	last := *messages.Twin["track"].LastState.Expected.Value
	log.Printf("LastState: %s CurrentState: %s ", last, current)
	if current == "kubeedge" {
		go utils.PUTMenu()
	} else if current == "stop" {
		utils.StopOmxplayer()
	} else {
		sfi, _ := strconv.Atoi(current)
		utils.StartOmxplayer(sfi)
	}
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Connection successfully")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Printf("The connection is broken\n: %v", err)
}

func main() {

	// create NewClientOptions
	opts := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1883")
	// SetClientID将设置客户机id，以便在连接到MQTT代理时由该客户机使用。根据MQTT规范，客户机id必须不超过23个字符
	opts.SetClientID("receive-client-ubuntu")
	opts.SetUsername("")
	opts.SetPassword("")

	// MQTT pub 消息处理
	//opts.SetDefaultPublishHandler(messagePubHandler)

	// 连接的回调，当客户机状态从未连接/断开连接更改为已连接时，将调用该回调函数，无论是在初始连接还是在重新连接时
	opts.OnConnect = connectHandler
	// 连接断开的回调，可以将其设置为在意外断开与MQTT代理的连接时执行，调用Disconnect或ForceDisconnect导致的断开不会导致执行OnConnectionLost回调
	opts.OnConnectionLost = connectLostHandler

	// create Client
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		log.Panic(token.Error())
	}
	// 退出 channel
	quit := make(chan os.Signal)
	// mqtt topic
	mqttTopic := `$hw/events/device/speaker-01/twin/update/document`
	// 断开连接
	defer c.Disconnect(250)
	// 取消订阅 topic
	defer unsubscribe(c, mqttTopic)

	// 订阅 topic
	subscribe(c, mqttTopic)

	//监听指定信号 ctrl+c kill
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	log.Printf("exit signal is : %s", <-quit)
}

func subscribe(client mqtt.Client, topic string) {
	// 在最大qos为零时，等待收据确认订阅
	if token := client.Subscribe(topic, 0, messagePubHandler); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		os.Exit(1)
	}
	log.Printf("Subscribed to topic: %s\n", topic)
}

func unsubscribe(client mqtt.Client, topic string) {
	if token := client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		os.Exit(1)
	}
}
