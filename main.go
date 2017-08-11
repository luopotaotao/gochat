package main

import (
	"gochat/core"
	"gochat/cons/msg"
	"log"
)

type WeChat struct {
	core.Core
}

func main() {

	chat:=core.New()
	f:= func(chat *core.Core,m core.AddMsg) {
		switch m.Type(){
		case msg.MSG_BLANK:
			log.Println("空消息,不做处理")
		default:
			log.Printf("new msg from %s to %s ,content:%s\n",m.FromUserName,m.ToUserName,m.Content)
			chat.SendMsg(chat.User.UserName,"filehelper","测试",chat.Ticket)
		}

	}
	chat.RegHandlers(map[int]func(chat *core.Core,msg core.AddMsg){
		msg.MSG_TEXT :f,
		msg.MSG_POSITION :f,
		msg.MSG_SHARING_MUSIC :f,
		msg.MSG_SHARING_COMMON :f,
		msg.MSG_SYSTEM :f,
		msg.MSG_IMG :f,
		msg.MSG_RECORDING :f,
		msg.MSG_EMOJI :f,
		msg.MSG_VEDIO :f,
		msg.MSG_CARD :f,
		msg.MSG_UNKNOW :f,
		msg.MSG_BLANK :f,
	})
	chat.Start(core.Config{})
}

