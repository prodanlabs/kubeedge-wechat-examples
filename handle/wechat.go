package handle

import (
	"log"
	"regexp"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/yaotian/gowechat/mp/message"
	"github.com/yaotian/gowechat/mp/user"

	"github.com/prodanlabs/kubeedge-wechat-examples/config"
	"github.com/prodanlabs/kubeedge-wechat-examples/utils"
)

// A regular expression which evaluates to true if the text contains `play`.
var rp = regexp.MustCompile(`\w*play\w*`)

// A regular expression which evaluates to true if the text contains `stop`.
var rs = regexp.MustCompile(`\w*stop\w*`)

// TextHandler
func TextHandler(ctx iris.Context) {

	wc := config.InitWechat()
	mp, err := wc.MpMgr()
	if err != nil {
		return
	}

	// 传入 request 和 responseWriter
	msgHandler := mp.GetMsgHandler(ctx.Request(), ctx.ResponseWriter())
	//设置接收消息的处理方法
	msgHandler.SetHandleMessageFunc(func(msg message.MixMessage) *message.Reply {
		log.Printf("Got text %s\n", msg.Content)
		requestText := strings.ToLower(strings.TrimSpace(msg.Content))
		replyText := ""
		if requestText == "list" {
			replyText += "raspberrypi的音乐\n"
			replyText += S.Songs
			replyText += "\n"
			replyText += "回复 play+数字 播放音乐(如: play 1)\n"
			replyText += "回复 stop 停止播放音乐"
			log.Println("requestText stop")

		} else if requestText == "kubeedge" {
			utils.UpdateDeviceTwinWithDesiredTrack(requestText)
			S.Songs = ""
			go func(s *StreamServer) {
				for {
					msg := <-s.Ch
					if msg == "EOF" {
						log.Println("Channel receive message complete")
						break
					} else {
						log.Printf("The value obtained from Channel is %s\n", msg)
						s.Songs += msg + "\n"
					}
				}
			}(S)
			replyText += "云边音乐信息正在同步中，请稍等..."
		} else if rp.MatchString(requestText) {
			items := rp.Split(requestText, -1)
			songTrack := strings.TrimSpace(items[1])
			log.Printf("Music Track: %s\n", songTrack)
			// Update device twin
			utils.UpdateDeviceTwinWithDesiredTrack(songTrack)
			replyText += "正在开始播放音乐，请稍等..."
		} else if rs.MatchString(requestText) {
			songTrack := "stop"
			// Update device twin
			utils.UpdateDeviceTwinWithDesiredTrack(songTrack)
			replyText += "正在停止播放音乐，请稍等..."
		} else {
			log.Println("Could not parse the message")
			replyText += "不能识别您的输入"
		}
		//回复消息：演示回复用户发送的消息
		return &message.Reply{message.MsgTypeText, message.NewText(replyText)}
	})

	//处理消息接收以及回复
	err = msgHandler.Handle()
	if err != nil {
		log.Printf("msgHandler error: %s", err)
		return
	}
}

// WxOAuth 微信公众平台，网页授权
func WxOAuth(ctx iris.Context) {
	wc := config.InitWechat()
	mp, err := wc.MpMgr()
	if err != nil {
		return
	}

	addr := utils.GetEnv("SERVER_ADDR_PORT", "foo.junengcloud.com:443")
	oauthHandler := mp.GetPageOAuthHandler(ctx.Request(), ctx.ResponseWriter(), "https://"+addr+"/oauth")

	oauthHandler.SetFuncCheckOpenIDExisting(func(openID string) (existing bool, stopNow bool) {
		//看自己的系统中是否已经存在此openID的用户
		//如果已经存在， 调用自己的Login 方法，设置cookie等，return true
		//如果还不存在，return false, handler会自动去取用户信息
		return false, true
	})

	oauthHandler.SetFuncAfterGetUserInfo(func(user user.Info) bool {
		//已获得用户信息，这里用信息做注册使用
		//调用自己的Login方法，设置cookie等
		return false
	})
	oauthHandler.Handle()
}
