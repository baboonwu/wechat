package wechat

import (
	"log"
	"testing"
)

var We *Wechat

func init() {

	log.SetFlags(log.Lshortfile)

	We = New("wx1ac20fc50222a85e", "9690c087a0d508bd8fe3faab63fb39a8")
}

func Test_AccessToken(t *testing.T) {

	result, errResult, e := We.AccessToken("code")
	log.Println(result, errResult, e)

	if e != nil {
		t.Error(e)
	}
}

func Test_GetUserInfo(t *testing.T) {

	result, errResult, e := We.AccessToken("code")
	log.Println(result, errResult, e)

	if e != nil {
		t.Error(e)
	}

	if result != nil {
		result, errResult, e := We.GetUserInfo(result.AccessToken, result.Openid)
		log.Println(result, errResult, e)
	}

}

func Test_RefreshToken(t *testing.T) {

}
