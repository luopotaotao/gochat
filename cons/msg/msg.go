package msg

const (
	TYPE_PHONE_INIT         = 51
	TYPE_TEXT               = 1
	TYPE_PICTURE            = 3
	TYPE_VEDIO              = 43
	TYPE_NAME_CARD          = 42
	TYPE_RECORDING          = 34
	TYPE_EMOJI              = 47
	TYPE_SHARING            = 49
	TYPE_SHARING_COMMON     = 5
	TYPE_SHARING_MUSIC      = 3
	TYPE_SHARING_RED_PACKET = 2001
	TYPE_SYSTEM             = 10000
)
const (
	MSG_UNKNOW = iota
	MSG_BLANK
	MSG_TEXT
	MSG_IMG
	MSG_VEDIO
	MSG_POSITION
	MSG_CARD
	MSG_RECORDING
	MSG_EMOJI
	MSG_SHARING_COMMON
	MSG_SHARING_MUSIC
	MSG_GROUP
	MSG_RED_PACKET
	MSG_SYSTEM
)
