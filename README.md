# GoChat   
wechat lib for Go

## usage
```GO

    package main

    import (
        "github.com/luopotaotao/gochat/cons/msg"
        "github.com/luopotaotao/gochat/core"
        "log"
    )

    type WeChat struct {
	    core.Core
    }

    func main() {

        chat := core.New()
        f := func(chat *core.Core, m core.AddMsg) {
            switch m.Type() {
            case msg.MSG_BLANK:
                log.Println("空消息,不做处理")
            default:
                log.Printf("new msg from %s to %s ,content:%s\n", m.FromUserName, m.ToUserName, m.Content)
                chat.SendMsg(chat.User.UserName, "filehelper", "测试", chat.Ticket)
            }
    
        }
        chat.RegHandlers(map[int]func(chat *core.Core, msg core.AddMsg){
            msg.MSG_TEXT:           f,
            msg.MSG_POSITION:       f,
            msg.MSG_SHARING_MUSIC:  f,
            msg.MSG_SHARING_COMMON: f,
            msg.MSG_SYSTEM:         f,
            msg.MSG_IMG:            f,
            msg.MSG_RECORDING:      f,
            msg.MSG_EMOJI:          f,
            msg.MSG_VEDIO:          f,
            msg.MSG_CARD:           f,
            msg.MSG_UNKNOW:         f,
            msg.MSG_BLANK:          f,
        })
        chat.Start(core.Config{})
    }
```
# Thanks
refrences and to my sincerely thanks:
    
1. [单眼皮的老虎的专栏 微信通信协议，用自己的程序收发微信，微信网页web版分析](http://blog.csdn.net/avsuper/article/details/63678827)  
2. [itchat](https://github.com/littlecodersh/ItChat) 
3. [itchat4j](https://github.com/yaphone/itchat4j) 
# Note:
much more work needed,welcome to help!
***

# GoChat
微信的Go语言库
# 致谢
此库参考了以下内容,向这些先行者表示诚挚的感谢:
    
1. [单眼皮的老虎的专栏 微信通信协议，用自己的程序收发微信，微信网页web版分析](http://blog.csdn.net/avsuper/article/details/63678827)  
2. [itchat](https://github.com/littlecodersh/ItChat) 
3. [itchat4j](https://github.com/yaphone/itchat4j) 

# 说明:
更多功能待完善中,欢迎指点!
`
