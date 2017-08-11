package url

const (
	LOGIN         = "https://login.weixin.qq.com"
	UUID          = LOGIN + "/jslogin"
	QR            = LOGIN + "/qrcode/"
	LOGIN_STATE   = LOGIN + "/cgi-bin/mmwebwx-bin/login"
	INFO          = "https://wx.qq.com/cgi-bin/mmwebwx-bin"
	INIT          = INFO + "/webwxinit" //注意次地址,网上的教程不对
	CONTACT_LIST  = INFO + "/webwxgetcontact"
	STATUS_NOTIFY = INFO + "/webwxstatusnotify"
	SYNCCHECK     = "https://webpush.wx.qq.com/cgi-bin/mmwebwx-bin/synccheck"
	SEND_MSG      = INFO + "/webwxsendmsg"
	GET_MSG       = INFO + "/webwxsync"
)
