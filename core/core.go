package core

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"fmt"
	"os"
	"gochat/cons/url"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"
	"encoding/json"
	"strconv"
	"time"
	"errors"
	"encoding/xml"
	"sync"
)

func New() *Core {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
		Transport: &http.Transport{
			DisableCompression: true,
			DisableKeepAlives:  false,
		},
		Timeout:15*time.Second,
	}
	return &Core{
		client:   client,
		Handlers: map[int]func(chat *Core,msg AddMsg){},
		DeviceId:Rand(),
	}
}

type Core struct {
	ResponseInit
	client   *http.Client
	Uuid     string
	Ticket   Ticket
	DeviceId string
	Handlers map[int]func(chat *Core,msg AddMsg)
}
type Config struct {
}

func (self *Core) Start(config Config) {
	uuid, err := self.step1LodadUuid()
	if err != nil {
		log.Print("获取uuid失败")
	}
	if err = self.step2GetQrAndShow(uuid); err != nil {
		log.Println("获取二维码失败")
	}
	var ticket Ticket

	tip := "1"
LoopQr:
	for i := 0; i < 9; i++ {
		code, uri, err := self.step3QrState(uuid, tip)
		if err != nil {
			log.Println("查询二维码状态失败")
		}
		switch code {
		case "200":
			log.Println("200")
			ticket, err = self.step4GetTicket(uri)
			self.Ticket = ticket
			break LoopQr
		case "201":
			tip = "0"
		case "408":
			tip = "1"
		}
	}
	if err != nil {
		log.Println("获取ticket失败",err)
	}
	info, err := self.step5WxInit(self.Ticket)
	self.ResponseInit = info
	if err != nil {
		log.Println("获取初始化信息失败")
	}
	contacts, err := self.step6LoadContactList(ticket, 0)
	if err != nil {
		log.Println("获取联系人列表失败")
	}
	log.Println(contacts.BaseResponse, contacts.MemberCount)

	_, err = self.step7StatusNotify(info.User.UserName, ticket)
	if err != nil {
		log.Println("StatusNotify失败")
	}
	var s sync.WaitGroup
	s.Add(1)
	go func() {
		counter := 0
		max := 10  //为非正整数时表示永不退出
		for {
			_, err := self.step8SyncCheck(url.SYNCCHECK, self.Ticket, self.SyncKey)
			if err != nil {
				log.Println("同步心跳失败!",counter,err)
				counter++
				if max>0&&counter >= max {
					log.Println("同步心跳失败次数达到上限,退出 ")
					s.Done()
				}
			} else {
				counter = 0
			}
			time.Sleep(15 * time.Second)
		}
	}()
	s.Wait()
}

//----------------------------------消息处理部分 start---------------------------------//
func (self *Core) RegHandlers(handlers map[int]func(chat *Core,msg AddMsg)) {
	for key, f := range handlers {
		self.Handlers[key] = f
	}
}
func (self *Core) OnMsg(msgs ResponseAddMsg) {
	//更新同步Key
	if msgs.BaseResponse.Ret==0&&msgs.SyncKey.Count > 0 {
		self.SyncKey = msgs.SyncKey
	}

	//处理消息
	for _, msg := range msgs.AddMsgList {
		fun, ok := self.Handlers[msg.Type()]
		if ok {
			log.Println("获取到处理函数,开始处理!")
			fun(self,msg)
		} else {
			log.Printf("on unkonwn msg type : %d %d ,content: %s",msg.MsgType, msg.AppMsgType, msg.Content)
		}
	}

}

//----------------------------------消息处理部分 end---------------------------------//
//----------------------------------微信登陆流程 begin---------------------------------//

func (self *Core) step1LodadUuid() (string, error) {
	log.Println("第1步,加载UUID")
	res_bytes, err := self.Get(url.UUID, map[string]string{
		"appid": "wx782c26e4c19acffb",
		"fun":   "new",
		"lang":  "zh_CN",
		"_":     Timestamp(),
	})
	if err != nil {
		log.Printf("获取UUID失败!%v", err)
		return "", os.ErrInvalid
	}
	reg := "window.QRLogin.code = (?P<code>\\d+); window.QRLogin.uuid = \"(?P<uuid>\\S+?)\";"
	code, strs, err := ParseResReg(reg, res_bytes, 1)
	if err != nil {
		return "", ErrUnexpectedResponse
	}
	uuid := strs[0]
	switch code {
	case "200":
		self.Uuid = uuid
		return uuid, nil
	case "201":
		log.Println("TODO") //TODO
	case "401":
		log.Println("TODO") //TODO
	}
	return "", ErrUnexpectedResponse
}

//第2步,获取二维码并打开供扫描
func (self *Core) step2GetQrAndShow(uuid string) error {
	log.Println("第1步,加载二维码")
	qr, err := self.Get(url.QR+uuid, nil)
	if err != nil {
		log.Println("获取二维码失败!")
		return err
	}
	err = ioutil.WriteFile("D:\\QR.jpg", qr, 777)
	if err != nil {
		log.Println("写入文件失败!")
		return err
	}
	return exec.Command("cmd.exe", "/c", "start D:\\QR.jpg").Run()
}

//第3步,查询是否扫描二维码
func (self *Core) step3QrState(uuid, tip string) (code, uri string, err error) {
	log.Println("第3步,查询是否扫描二维码")
	res, err := self.Get(url.LOGIN_STATE, map[string]string{"uuid": uuid, "tip": tip, "_": Timestamp()})
	res_txt := string(res)
	//reg := "window.code=(?P<code>\\d+); window.redirect_uri=\"(?P<uri>\\S+)\";"
	code = regexp.MustCompile("\\d+").FindString(res_txt)
	uri = regexp.MustCompile("\"(\\S+)\"").FindString(res_txt)
	uri = strings.Trim(uri, "\"")
	if err != nil {
		return "", "", err
	}
	switch code {
	case "200":
		return code, uri, nil
	case "201":
		log.Println("已扫描二维码,请在手机微信确认登陆") //TODO
	case "408":
		log.Println("等待超时") //TODO
	}
	return code, uri, nil

}

//第4步,获取Ticket
func (self *Core) step4GetTicket(uri string) (t Ticket, err error) {
	log.Println("第4步,获取Ticket")
	ret_bytes, err := self.Get(fmt.Sprintf("%s&fun=new", uri), nil)
	if err != nil {
		return t, ErrUnexpectedResponse
	}
	err = xml.Unmarshal(ret_bytes, &t)
	return t, err
}

//第5步,初始化
func (self *Core) step5WxInit(t Ticket) (ret ResponseInit, err error) {
	log.Println("第5步,初始化")
	url := fmt.Sprintf("%s?r=%s", url.INIT, Timestamp())
	res_bytes, err := self.Post(url, t, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(res_bytes, &ret)
	return
}

//第6步,加载联系人列表
func (self *Core) step6LoadContactList(ticket Ticket, seq int) (ret ResponseContactList, err error) {
	log.Println("第6步,加载联系人列表")
	res_bytes, err := self.Get(url.CONTACT_LIST, map[string]string{
		"lang":        "zh_CN",
		"pass_ticket": ticket.PassTicket,
		"r":           Timestamp(),
		"seq":         strconv.Itoa(seq),
		"skey":        ticket.Skey,
	})
	if err != nil {
		return
	}
	err = json.Unmarshal(res_bytes, &ret)
	return
}

//第7步,开启微信状态通知 //TODO这步有啥用?
func (self *Core) step7StatusNotify(uname string, t Ticket) (ret []byte, err error) {
	log.Println("第7步,开启微信状态通知")
	url := fmt.Sprintf("%s?lang=zh_CN&pass_ticket=%s", url.STATUS_NOTIFY, t.PassTicket)
	params := map[string]interface{}{
		"Code":         3,
		"FromUserName": uname,
		"ToUserName":   uname,
		"ClientMsgId":  time.Now().UTC().Unix(),
	}
	return self.Post(url, t, params)
}

const (
	CodeNewMsg2   = "2" //新消息
	CodeNewMsg6   = "6" //新消息
	CodePublicMsg = "4" //公众号消息
)

//第8步,心跳同步
func (self *Core) step8SyncCheck(url string, t Ticket, keys SyncKey) (ret ResponseSync, err error) {
	log.Println("第8步,心跳同步")
	timestamp := Timestamp()
	key := make([]string, keys.Count, keys.Count)
	for i, item := range keys.List {
		key[i] = strings.Join([]string{
			strconv.Itoa(item.Key),
			strconv.Itoa(item.Val),
		}, "_")
	}
	res_bytes, err := self.Get(url, map[string]string{
		"r":        timestamp,
		"skey":     t.Skey,
		"sid":      t.Wxsid,
		"uin":      t.Wxuin,
		"deviceid": self.DeviceId,
		"synckey":  strings.Join(key, "|"),
		"_":        timestamp,
	})
	re_json := regexp.MustCompile("\\{\\S+,\\S+\\}")
	re_num := regexp.MustCompile("\\d+")
	str := re_json.FindString(string(res_bytes))
	if str != "" {
		arr := strings.Split(str, ",")
		ret = ResponseSync{
			re_num.FindString(arr[0]),
			re_num.FindString(arr[1]),
		}
	}
	if ret.RetCode != "0" {
		err = errors.New("心跳同步失败!"+string(res_bytes))
		return
	}
	switch ret.Selector {
	case CodeNewMsg2, CodeNewMsg6:
		msgs, err1 := self.loadMsgs()
		if err1 != nil {
			log.Println("获取新消息失败!")
			return ret, err1
		}
		self.OnMsg(msgs)
	case CodePublicMsg:
		log.Println("公众号消息")
	default:
		log.Println("同步完成")
	}
	return
}
func (self *Core) loadMsgs() (msgs ResponseAddMsg, err error) {
	log.Println("第9步,加载消息")
	url := fmt.Sprintf("%s?sid=%s&skey=%s", url.GET_MSG, self.Ticket.Wxsid, self.Ticket.Skey)
	params := map[string]interface{}{
		"SyncKey": self.SyncKey,
		"rr":      time.Now().UTC().Unix(),
	}
	res_bytes, err := self.Post(url, self.Ticket, params)

	if err != nil {
		log.Println("获取新消息失败", err)
		return
	}
	err = json.Unmarshal(res_bytes, &msgs)
	return
}

func (self *Core) SendMsg(from, to, msg string, ticket Ticket) error {
	timestamp := Timestamp()
	m := map[string]interface{}{
		"Msg": Msg{
			1,
			msg,
			from,
			to,
			timestamp,
			timestamp,
		},
	}
	_, err := self.Post(url.SEND_MSG, ticket, m)
	return err
}

//----------------------------------微信登陆流程 end---------------------------------//
