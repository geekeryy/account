// Package yunpian @Description  TODO
// @Author  	 jiangyang
// @Created  	 2021/11/27 9:17 下午
package yunpian

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	singleSendUrl = "https://sms.yunpian.com/v2/sms/single_send.json"
	msgTemplate   = "【江杨】您的验证码是%d"
)

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}

func (y Client) SendCode(mobile string, code int) error {

	httpCli := http.DefaultClient

	data := url.Values{}
	data.Set("apikey", y.apiKey)
	data.Set("mobile", mobile)
	data.Set("text", fmt.Sprintf(msgTemplate, code))

	request, err := http.NewRequest(http.MethodPost, singleSendUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	request.Header.Set("Accept", "application/json;charset=utf-8")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := httpCli.Do(request)
	if err != nil {
		return err
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	res := SingleSendResponse{}
	if err := json.Unmarshal(all, &res); err != nil {
		return err
	}
	if res.Code != 0 {
		return errors.New("single msg err: " + mobile + res.Msg)
	}
	return nil

}

type SingleSendResponse struct {
	Code   int     `json:"code"`
	Msg    string  `json:"msg"`
	Count  int     `json:"count"`
	Fee    float64 `json:"fee"`
	Unit   string  `json:"unit"`
	Mobile string  `json:"mobile"`
	Sid    int64   `json:"sid"`
}
