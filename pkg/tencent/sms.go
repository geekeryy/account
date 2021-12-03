// Package tencent @Description  TODO
// @Author  	 jiangyang
// @Created  	 2021/10/16 10:51 上午
package tencent

import (
	"encoding/json"
	"log"

	"github.com/comeonjy/account/configs"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

type TenSms struct {
	client          *sms.Client
	request         *sms.SendSmsRequest
	LoginTemplateId string
}

type smsConf struct {
	LoginTemplateId string `json:"login_template_id"`
	SmsSdkAppId     string `json:"sms_sdk_app_id"`
	SignName        string `json:"sign_name"`
}

func NewTenSms(cfg configs.Interface) *TenSms {
	conf := smsConf{}
	if err := json.Unmarshal([]byte(cfg.Get().TenSmsConf), &conf); err != nil {
		log.Fatalln("NewTenSms err:", err)
	}
	credential := common.NewCredential(cfg.Get().TenSecretId, cfg.Get().TenSecretKey)
	cpf := profile.NewClientProfile()
	client, _ := sms.NewClient(credential, "ap-guangzhou", cpf)
	request := sms.NewSendSmsRequest()
	request.SmsSdkAppId = common.StringPtr(conf.SmsSdkAppId)
	request.SignName = common.StringPtr(conf.SignName)
	request.SenderId = common.StringPtr("")
	return &TenSms{
		client:          client,
		request:         request,
		LoginTemplateId: conf.LoginTemplateId,
	}
}

// SendLoginCode 短信登录验证
// TODO 短信发送记录
func (s *TenSms) SendLoginCode(mobile, code string) error {

	s.request.TemplateId = common.StringPtr(s.LoginTemplateId)
	s.request.SessionContext = common.StringPtr("user1")
	s.request.TemplateParamSet = common.StringPtrs([]string{code})
	s.request.PhoneNumberSet = common.StringPtrs([]string{"+86" + mobile})
	_, err := s.client.SendSms(s.request)
	if err != nil {
		return err
	}
	return nil
}
