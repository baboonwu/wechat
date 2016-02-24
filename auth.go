package wechat

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
获取第一步的code后，请求以下链接获取access_token：
https://api.weixin.qq.com/sns/e/access_token?appid=APPID&secret=SECRET&code=CODE&grant_type=authorization_code

参数说明
参数	是否必须	说明
	appid	是	应用唯一标识，在微信开放平台提交应用审核通过后获得
	secret	是	应用密钥AppSecret，在微信开放平台提交应用审核通过后获得
	code	是	填写第一步获取的code参数
	grant_type	是	填authorization_code

返回说明
正确的返回：
{
	"access_token":"ACCESS_TOKEN",
	"expires_in":7200,
	"refresh_token":"REFRESH_TOKEN",
	"openid":"OPENID",
	"scope":"SCOPE",
	"unionid":"o6_bmasdasdsad6_2sgVt7hMZOPfL"
}

参数	说明
	access_token	接口调用凭证
	expires_in	    access_token接口调用凭证超时时间，单位（秒）
	refresh_token	用户刷新access_token
	openid	        授权用户唯一标识
	scope	        用户授权的作用域，使用逗号（,）分隔
	unionid	        当且仅当该移动应用已获得该用户的userinfo授权时，才会出现该字段

错误返回样例：
{"errcode":40029,"errmsg":"invalid code"}

*/

type RespAccessToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	Unionid      string `json:"unionid"`
}

func (w *Wechat) AccessToken(code string) (*RespAccessToken, *Error, error) {

	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", w.AppID, w.AppSecret, code)

	// 1. http request
	resp, err := doHTTP("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	// 2. process response
	wechatError := &Error{}
	if e := json.Unmarshal(resp, wechatError); e != nil {
		log.Println(e, url)
		return nil, nil, e
	}

	if wechatError.ErrCode != 0 {
		log.Println(wechatError)
		return nil, wechatError, nil
	}

	result := &RespAccessToken{}
	if e := json.Unmarshal(resp, result); e != nil {
		log.Println(e, url)
		return nil, nil, e
	}

	return result, nil, nil
}
