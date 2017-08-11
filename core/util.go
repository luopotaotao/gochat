package core

import (
	"bytes"
	"strings"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"time"
	"regexp"
	"math/rand"
)

func(self *Core) Get(url string, p map[string]string) ([]byte, error) {

	size := len(p)

	if size > 0 {
		buff := new(bytes.Buffer)
		buff.WriteString(url)
		buff.WriteString("?")
		for key, val := range p {
			buff.WriteString(key)
			buff.WriteString("=")
			buff.WriteString(val)
			buff.WriteString("&")
		}
		url = strings.Trim(buff.String(), "&")
	}
	res, err := self.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (self *Core) Post(url string, t Ticket, params map[string]interface{}) (ret []byte, err error) {
	//log.Println(url)
	br := BaseRequest{
		t.Wxuin,
		t.Wxsid,
		t.Skey,
		self.DeviceId,
	}
	body := map[string]interface{}{"BaseRequest": br}
	if len(params) > 0 {
		for key, val := range params {
			body[key] = val
		}
	}
	data, _ := json.Marshal(body)
	res, err := self.client.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func Timestamp() string {
	return fmt.Sprintf("%d", time.Now().UTC().Unix())
}

var chars []rune = []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func Rand() string {
	buff := new(bytes.Buffer)
	buff.WriteRune('e')
	for i := 0; i < 15; i++ {
		buff.WriteRune(chars[rand.Intn(len(chars))])
	}
	return buff.String()
}

//n除code外还要匹配几个
func ParseResReg(reg string, txt []byte, n int) (code string, strs []string, err error) {
	res_text := string(txt)
	re := regexp.MustCompile(reg)
	re.MatchString(res_text)
	matches := re.FindAllStringSubmatch(res_text, n+2)
	if len(matches) < 1 || len(matches[0]) != n+2 {
		return "", nil, ErrUnexpectedResponse
	}
	//全句匹配,code,其他
	code = matches[0][1]
	strs = matches[0][2:]
	return code, strs, nil
}

