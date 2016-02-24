package wechat

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const appID = ""
const appSecret = ""

type Wechat struct {
	AppID     string
	AppSecret string
}

func New(appid, secret string) *Wechat {
	return &Wechat{
		AppID:     appid,
		AppSecret: secret,
	}
}

func doHTTP(method, url string, body io.Reader) ([]byte, error) {

	log.Println(url)

	req, e := http.NewRequest(method, url, body)
	if e != nil {
		log.Println(e, url)
		return nil, e
	}

	c := &http.Client{}
	resp, e := c.Do(req)
	if e != nil {
		return nil, e
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("http code error", resp.StatusCode)
		return nil, fmt.Errorf("Expecting HTTP status code 200, but got %v", resp.StatusCode)
	}

	rbody, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return nil, e
	}

	return rbody, nil
}
