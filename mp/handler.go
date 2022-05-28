package mp

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yaotian/gowechat"
	"github.com/yaotian/gowechat/mp/message"
	"github.com/yaotian/gowechat/wxcontext"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Handler(c *gin.Context) {
	mpCfg := wxcontext.Config{
		AppID:          Config.Mp.AppID,
		AppSecret:      Config.Mp.AppSecret,
		Token:          Config.Mp.Token,
		EncodingAESKey: Config.Mp.EncodingAesKey,
	}
	wechat := gowechat.NewWechat(mpCfg)
	mp, err := wechat.MpMgr()
	if err != nil {
		return
	}
	msgHandler := mp.GetMsgHandler(c.Request, c.Writer)
	msgHandler.SetHandleMessageFunc(replayMsg)
	err = msgHandler.Handle()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func replayMsg(msg message.MixMessage) *message.Reply {
	text := message.NewText(msg.Content)
	receive := text.Content
	text.Content = qingyueke(receive)
	go RecordSendMsg(MongoClient(), msg.FromUserName, receive, text.Content)
	return &message.Reply{message.MsgTypeText, text}
}

func qingyueke(query string) string {
	resp, err := http.Get("https://vqdata.chenshaowen.com/api/v1/bot?keyword=" + url.QueryEscape(query))
	if err != nil {
		fmt.Println(err)
		return "能在说一遍么，刚才走神了"
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	res := struct {
		Data string `json:"data"`
	}{}
	err = json.Unmarshal(s, &res)
	return res.Data
}
