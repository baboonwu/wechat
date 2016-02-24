package wechat

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
获取用户个人信息（UnionID机制）
接口说明
此接口用于获取用户个人信息。开发者可通过OpenID来获取用户基本信息。
特别需要注意的是，如果开发者拥有多个移动应用、网站应用和公众帐号，可通过获取用户基本信息中的unionid来区分用户的唯一性，
因为只要是同一个微信开放平台帐号下的移动应用、网站应用和公众帐号，用户的unionid是唯一的。
换句话说，同一用户，对同一个微信开放平台下的不同应用，unionid是相同的。

请注意，在用户修改微信头像后，旧的微信头像URL将会失效，因此开发者应该自己在获取用户信息后，将头像图片保存下来，避免微信头像URL失效后的异常情况。

请求说明

http请求方式: GET
https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID

参数说明
参数	是否必须	说明
	access_token	是	调用凭证
	openid	        是	普通用户的标识，对当前开发者帐号唯一
	lang	        否	国家地区语言版本，zh_CN 简体，zh_TW 繁体，en 英语，默认为zh-CN

返回说明
正确的Json返回结果：
{
	"openid":"OPENID",
	"nickname":"NICKNAME",
	"sex":1,
	"province":"PROVINCE",
	"city":"CITY",
	"country":"COUNTRY",
	"headimgurl": "http://wx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/0",
	"privilege":[
		"PRIVILEGE1",
		"PRIVILEGE2"
	],
	"unionid": " o6_bmasdasdsad6_2sgVt7hMZOPfL"
}
参数	说明
openid	普通用户的标识，对当前开发者帐号唯一
nickname	普通用户昵称
sex	普通用户性别，1为男性，2为女性
province	普通用户个人资料填写的省份
city	普通用户个人资料填写的城市
country	国家，如中国为CN
headimgurl	用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空
privilege	用户特权信息，json数组，如微信沃卡用户为（chinaunicom）
unionid	用户统一标识。针对一个微信开放平台帐号下的应用，同一用户的unionid是唯一的。
错误的Json返回示例:
{
"errcode":40003,"errmsg":"invalid openid"
}

*/

type RespUserInfo struct {
	Openid     string   `json:"openid"`
	NickName   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}

func (w *Wechat) GetUserInfo(accesstoken, openid string) (*RespUserInfo, *Error, error) {

	url := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s", accesstoken, openid)

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

	result := &RespUserInfo{}
	if e := json.Unmarshal(resp, result); e != nil {
		log.Println(e, url)
		return nil, nil, e
	}

	return result, nil, nil
}
