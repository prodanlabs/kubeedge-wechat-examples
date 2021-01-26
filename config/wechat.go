package config

import (
	"github.com/yaotian/gowechat"
	"github.com/yaotian/gowechat/wxcontext"

	"github.com/prodanlabs/kubeedge-wechat-examples/utils"
)

func InitWechat() *gowechat.Wechat {
	//配置微信参数
	config := wxcontext.Config{
		AppID:          utils.GetEnv("WECHAT_APP_ID", "xxx"),
		AppSecret:      utils.GetEnv("WECHAT_APP_SECRET", "xxx"),
		Token:          utils.GetEnv("WECHAT_TOKEN", "xxx"),
		EncodingAESKey: utils.GetEnv("WECHAT_ENCODING_AES_KEY", "xxx"),
	}
	return gowechat.NewWechat(config)
}