package core

import (
	"errors"
	"gochat/cons/msg"
	"strings"
	"regexp"
	"log"
)

type (
	Ticket struct {
		Ret         int
		Message     string `xml:"message"`
		Skey        string `xml:"skey"`
		Wxsid       string `xml:"wxsid"`
		Wxuin       string `xml:"wxuin"`
		PassTicket  string `xml:"pass_ticket"`
		Isgrayscale int `xml:"isgrayscale"`
	}
	BaseRequest struct {
		Uin      string `json:"Uin"`
		Sid      string `json:"Sid"`
		Skey     string `json:"Skey"`
		DeviceID string `json:"DeviceID"`
	}

	BaseResponse struct {
		Ret    int `json:"Ret"`
		ErrMsg string `json:"ErrMsg"`
	}
	Contact struct {
		Uin              int `json:"Uin"`
		UserName         string `json:"UserName"`
		NickName         string `json:"NickName"`
		HeadImgUrl       string `json:"HeadImgUrl"`
		ContactFlag      int `json:"ContactFlag"`
		MemberCount      int `json:"MemberCount"`
		MemberList       []*Member `json:"MemberList"`
		RemarkName       string `json:"RemarkName"`
		HideInputBarFlag int `json:"HideInputBarFlag"`
		Sex              int `json:"Sex"`
		Signature        string `json:"Signature"`
		VerifyFlag       int `json:"VerifyFlag"`
		OwnerUin         int `json:"OwnerUin"`
		PYInitial        string `json:"PYInitial"`
		PYQuanPin        string `json:"PYQuanPin"`
		RemarkPYInitial  string `json:"RemarkPYInitial"`
		RemarkPYQuanPin  string `json:"RemarkPYQuanPin"`
		StarFriend       int `json:"StarFriend"`
		AppAccountFlag   int `json:"AppAccountFlag"`
		Statues          int `json:"Statues"`
		AttrStatus       int `json:"AttrStatus"`
		Province         string `json:"Province"`
		City             string `json:"City"`
		Alias            string `json:"Alias"`
		SnsFlag          int `json:"SnsFlag"`
		UniFriend        int `json:"UniFriend"`
		DisplayName      string `json:"DisplayName"`
		ChatRoomId       int `json:"ChatRoomId"`
		KeyWord          string `json:"KeyWord"`
		EncryChatRoomId  string `json:"EncryChatRoomId"`
		IsOwner          int `json:"IsOwner"`
	}
	Member struct {
		Uin             int `json:"Uin"`
		UserName        string `json:"UserName"`
		NickName        string `json:"NickName"`
		AttrStatus      int `json:"AttrStatus"`
		PYInitial       string `json:"PYInitial"`
		PYQuanPin       string `json:"PYQuanPin"`
		RemarkPYInitial string `json:"RemarkPYInitial"`
		RemarkPYQuanPin string `json:"RemarkPYQuanPin"`
		MemberStatus    int `json:"MemberStatus"`
		DisplayName     string `json:"DisplayName"`
		KeyWord         string `json:"KeyWord"`
	}
	SyncKeyItem struct {
		Key int `json:"Key"`
		Val int `json:"Val"`
	}
	SyncKey struct {
		Count int `json:"Count"`
		List  []*SyncKeyItem `json:"List"`
	}
	User struct {
		Uin               int `json:"Uin"`
		UserName          string `json:"UserName"`
		NickName          string `json:"NickName"`
		HeadImgUrl        string `json:"HeadImgUrl"`
		RemarkName        string `json:"RemarkName"`
		PYInitial         string `json:"PYInitial"`
		PYQuanPin         string `json:"PYQuanPin"`
		RemarkPYInitial   string `json:"RemarkPYInitial"`
		RemarkPYQuanPin   string `json:"RemarkPYQuanPin"`
		HideInputBarFlag  int `json:"HideInputBarFlag"`
		StarFriend        int `json:"StarFriend"`
		Sex               int `json:"Sex"`
		Signature         string `json:"Signature"`
		AppAccountFlag    int `json:"AppAccountFlag"`
		VerifyFlag        int `json:"VerifyFlag"`
		ContactFlag       int `json:"ContactFlag"`
		WebWxPluginSwitch int `json:"WebWxPluginSwitch"`
		HeadImgFlag       int `json:"HeadImgFlag"`
		SnsFlag           int `json:"SnsFlag"`
	}
	MPArticle struct {
		Title  string `json:"Title"`
		Digest string `json:"Digest"`
		Cover  string `json:"Cover"`
		Url    string `json:"Url"`
	}
	MPSubscribeMsg struct {
		UserName       string `json:"UserName"`
		MPArticleCount int `json:"MPArticleCount"`
		MPArticleList  []*MPArticle `json:"MPArticleList"`
		Time           int `json:"Time"`
		NickName       string `json:"NickName"`
	}
	ResponseInit struct {
		BaseResponse        BaseResponse `json:"BaseResponse"`
		Count               int `json:"Count"`
		ContactList         []*Contact `json:"ContactList"`
		SyncKey             SyncKey `json:"SyncKey"`
		User                User `json:"User"`
		ChatSet             string `json:"ChatSet"`
		SKey                string `json:"SKey"`
		ClientVersion       int `json:"ClientVersion"`
		SystemTime          int `json:"SystemTime"`
		GrayScale           int `json:"GrayScale"`
		InviteStartCount    int `json:"InviteStartCount"`
		MPSubscribeMsgCount int `json:"MPSubscribeMsgCount"`
		MPSubscribeMsgList  []*MPSubscribeMsg `json:"MPSubscribeMsgList"`
		ClickReportInterval int `json:"ClickReportInterval"`
	}
	ResponseContactList struct {
		BaseResponse BaseResponse `json:"BaseResponse"`
		MemberCount  int `json:"MemberCount"`
		MemberList   []*Contact `json:"MemberList"`
		Seq          int `json:"Seq"`
	}
	ResponseSync struct {
		RetCode  string `json:"retcode"`
		Selector string `json:"selector"`
	}
	Msg struct {
		Type         int `json:"Type"`
		Content      string `json:"Content"`
		FromUserName string `json:"FromUserName"`
		ToUserName   string `json:"ToUserName"`
		LocalID      string `json:"LocalID"`
		ClientMsgId  string `json:"ClientMsgId"`
	}
	RecommendInfo struct {
		UserName   string `json:"UserName"`
		NickName   string `json:"NickName"`
		QQNum      int `json:"QQNum"`
		Province   string `json:"Province"`
		City       string `json:"City"`
		Content    string `json:"Content"`
		Signature  string `json:"Signature"`
		Alias      string `json:"Alias"`
		Scene      int `json:"Scene"`
		VerifyFlag int `json:"VerifyFlag"`
		AttrStatus int `json:"AttrStatus"`
		Sex        int `json:"Sex"`
		Ticket     string `json:"Ticket"`
		OpCode     int `json:"OpCode"`
	}
	AppInfo struct {
		AppID string `json:"AppID"`
		Type  int `json:"Type"`
	}
	AddMsg struct {
		MsgId                string `json:"MsgId"`
		FromUserName         string `json:"FromUserName"`
		ToUserName           string `json:"ToUserName"`
		MsgType              int `json:"MsgType"`
		Content              string `json:"Content"`
		Status               int `json:"Status"`
		ImgStatus            int `json:"ImgStatus"`
		CreateTime           int `json:"ImgStatus"`
		VoiceLength          int `json:"VoiceLength"`
		PlayLength           int `json:"PlayLength"`
		FileName             string `json:"FileName"`
		FileSize             string `json:"FileSize"`
		MediaId              string `json:"MediaId"`
		Url                  string `json:"Url"`
		AppMsgType           int `json:"AppMsgType"`
		StatusNotifyCode     int `json:"StatusNotifyCode"`
		StatusNotifyUserName string `json:"StatusNotifyUserName"`
		RecommendInfo        RecommendInfo `json:"RecommendInfo"`
		ForwardFlag          int `json:"ForwardFlag"`
		AppInfo              AppInfo `json:"AppInfo"`
		HasProductId         int `json:"HasProductId"`
		Ticket               string `json:"Ticket"`
		ImgHeight            int `json:"ImgHeight"`
		ImgWidth             int `json:"ImgWidth"`
		SubMsgType           int `json:"SubMsgType"`
		NewMsgId             int `json:"NewMsgId"`
		OriContent           string `json:"OriContent"`
	}
	ProfileUnit struct {
		Buff string `json:"Buff"`
	}
	Profile struct {
		BitFlag           int `json:"BitFlag"`
		UserName          ProfileUnit `json:"UserName"`
		NickName          ProfileUnit `json:"NickName"`
		BindUin           int `json:"BindUin"`
		BindEmail         ProfileUnit `json:"BindEmail"`
		BindMobile        ProfileUnit `json:"BindMobile"`
		Status            int `json:"Status"`
		Sex               int `json:"Sex"`
		PersonalCard      int `json:"PersonalCard"`
		Alias             string `json:"Alias"`
		HeadImgUpdateFlag int `json:"HeadImgUpdateFlag"`
		HeadImgUrl        string `json:"HeadImgUrl"`
		Signature         string `json:"Signature"`
	}
	ResponseAddMsg struct {
		BaseResponse           BaseResponse `json:"BaseResponse"`
		AddMsgCount            int `json:"AddMsgCount"`
		AddMsgList             []AddMsg `json:"AddMsgList"`
		ModContactCount        int `json:"ModContactCount"`
		ModContactList         [] interface{} `json:"ModContactList"`
		DelContactCount        int `json:"DelContactCount"`
		DelContactList         [] interface{} `json:"DelContactList"`
		ModChatRoomMemberCount int `json:"ModChatRoomMemberCount"`
		ModChatRoomMemberList  [] interface{} `json:"ModChatRoomMemberList"`
		Profile                Profile `json:"Profile"`
		ContinueFlag           int `json:"ContinueFlag"`
		SyncKey                SyncKey `json:"SyncKey"`
		SKey                   string `json:"SKey"`
		SyncCheckKey           SyncKey `json:"SyncCheckKey"`
	}
)

var ErrUnexpectedResponse = errors.New("Unexpected Response")

func (self *AddMsg) Type() (ret int) {
	content := self.Content
	if content == "" {
		return msg.MSG_BLANK
	}
	switch self.MsgType {
	case msg.TYPE_TEXT:
		log.Printf("text msg:%+v",self)
		ok, _ := regexp.MatchString("/cgi-bin/mmwebwx-bin/webwxgetpubliclinkimg\\?url=xxx&msgid=\\d+&pictype=location", content)
		switch {
		case ok:
			return msg.MSG_POSITION
		case strings.HasPrefix(self.FromUserName, "@@"):
			return msg.MSG_GROUP
		}
		return msg.MSG_TEXT
	case msg.TYPE_PICTURE:
		return msg.MSG_IMG
	case msg.TYPE_VEDIO:
		return msg.MSG_VEDIO
	case msg.TYPE_NAME_CARD:
		return msg.MSG_CARD
	case msg.TYPE_RECORDING:
		return msg.MSG_RECORDING
	case msg.TYPE_EMOJI:
		return msg.MSG_EMOJI
	case msg.TYPE_SHARING:
		switch self.AppMsgType {
		case msg.TYPE_SHARING_MUSIC:
			return msg.MSG_SHARING_MUSIC
		case msg.TYPE_SHARING_COMMON:
			return msg.MSG_SHARING_COMMON
		case msg.TYPE_SHARING_RED_PACKET:
			return msg.MSG_RED_PACKET
		}
	case msg.TYPE_SYSTEM:
		return msg.MSG_SYSTEM
	}

	return msg.MSG_UNKNOW
}
