// Package wechat @Description  TODO
// @Author  	 jiangyang
// @Created  	 2021/10/16 6:16 下午
package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Mini struct {
	appid  string
	secret string
}

func NewMini(appid, secret string) *Mini {
	return &Mini{
		appid:  appid,
		secret: secret,
	}
}

type JsCode2sessionResp struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrorMsg   string `json:"errormsg"`
}

func (m *Mini) JsCode2session(code string) (*JsCode2sessionResp, error) {
	urlStr := "https://api.weixin.qq.com/sns/jscode2session?appid=" + m.appid + "&secret=" + m.secret + "&js_code=" + code + "&grant_type=authorization_code"
	resp, err := http.Get(urlStr)
	if err != nil {
		return nil, err
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	res := JsCode2sessionResp{}
	if err := json.Unmarshal(all, &res); err != nil {
		return nil, err
	}
	if res.ErrCode != 0 {
		return nil, errors.New(fmt.Sprintf("%d:%s", res.ErrCode, res.ErrorMsg))
	}
	if len(res.Openid) == 0 {
		return nil, errors.New("jscode2session err resp")
	}
	return &res, nil
}
