package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	v1 "github.com/comeonjy/account/api/v1"
	"github.com/comeonjy/account/internal/data"
	"github.com/comeonjy/account/pkg/consts"
	"github.com/comeonjy/account/pkg/redis"
	"github.com/comeonjy/account/pkg/util"
	"github.com/comeonjy/go-kit/pkg/xerror"
	"github.com/comeonjy/go-kit/pkg/xjwt"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (svc *AccountService) GetByID(ctx context.Context, in *v1.GetByIDReq) (*v1.GetByIDResp, error) {
	user, err := svc.accountRepo.Get(ctx, in.GetId())
	if err != nil {
		return nil, xerror.NewError(xerror.SystemErr, "数据查询失败，请稍后再试")
	}

	respUser := v1.UserModel{}
	if err := copier.Copy(&respUser, user); err != nil {
		return nil, xerror.NewError(xerror.CopierErr, "数据查询失败，请稍后再试", err.Error())
	}

	return &v1.GetByIDResp{User: &respUser}, nil
}

func (svc *AccountService) MiniLogin(ctx context.Context, in *v1.MiniLoginReq) (*v1.MiniLoginResp, error) {
	session, err := svc.mini.JsCode2session(in.Code)
	if err != nil {
		return nil, xerror.NewError(xerror.WechatErr, "登录失败，请稍后重试！", err.Error())
	}

	user := data.UserModel{
		WechatOpenid: session.Openid,
	}
	if err := svc.accountRepo.GetByWechatOpenid(ctx, &user); err != nil {
		return nil, xerror.NewError(xerror.SQLErr, "登录失败，请稍后重试！", err.Error())
	}
	if user.Id == 0 {
		user.UUID = uuid.NewString()
		if err := svc.accountRepo.Create(ctx, &data.UserModel{
			UUID:         user.UUID,
			WechatOpenid: session.Openid,
		}); err != nil {
			return nil, xerror.NewError(xerror.SQLErr, "登录失败，请稍后重试！", err.Error())
		}
	}

	bus := xjwt.Business{
		UUID: user.UUID,
		Role: user.Role,
	}
	marshal, err := json.Marshal(bus)
	if err != nil {
		return nil, xerror.NewError(xerror.MarshalErr, "登录失败，请稍后重试！", err.Error())
	}
	token, err := xjwt.CreateToken(string(marshal), time.Hour*24)
	if err != nil {
		return nil, xerror.NewError(xerror.JwtErr, "登录失败，请稍后重试！", err.Error())
	}

	return &v1.MiniLoginResp{
		Token: token,
		UserInfo: &v1.UserInfo{
			NickName:  user.NickName,
			AvatarUrl: user.AvatarUrl,
		}}, nil
}
func (svc *AccountService) UpdatesUser(ctx context.Context, in *v1.UpdatesUserReq) (*v1.Empty, error) {
	bus, err := svc.getCurrentUser(ctx)
	if err != nil {
		return nil, xerror.NewError(xerror.AuthErr, "", err.Error())
	}
	if err := svc.accountRepo.Updates(ctx, &data.UserModel{
		UUID:      bus.UUID,
		NickName:  in.NickName,
		AvatarUrl: in.AvatarUrl,
	}); err != nil {
		return nil, xerror.NewError(xerror.SQLErr, "登录失败，请稍后重试！", err.Error())
	}
	return &v1.Empty{}, nil
}

func (svc *AccountService) SendVerificationCode(ctx context.Context, in *v1.SendVerificationCodeReq) (*v1.Empty, error) {
	rand.Seed(time.Now().Unix())
	code := rand.Intn(9000) + 1000
	if err := svc.redis.Set(ctx, fmt.Sprintf(redis.SmsLoginCode, util.Md5(in.GetAccount())), code, 5*time.Minute).Err(); err != nil {
		return nil, xerror.NewError(xerror.RedisErr, "发送失败，请重新发送", err.Error())
	}
	switch in.GetType() {
	case "mobile":
		if err := svc.sms.SendCode(in.GetAccount(), code); err != nil {
			return nil, xerror.NewError(xerror.YunPianErr, "发送失败，请重新发送", err.Error())
		}
	case "email":
		if err := svc.email.SendMail([]string{in.GetAccount()}, "验证码", strings.Replace(consts.VerificationCodeTpl, "{{code}}", strconv.Itoa(code), 1)); err != nil {
			return nil, xerror.NewError(xerror.EmailErr, "发送失败，请重新发送", err.Error())
		}
	}

	return &v1.Empty{}, nil
}

func (svc *AccountService) Login(ctx context.Context, in *v1.LoginReq) (*v1.LoginResp, error) {
	var code string
	var err error

	if in.GetType() != "password" {
		code, err = svc.redis.Get(ctx, fmt.Sprintf(redis.SmsLoginCode, util.Md5(in.GetAccount()))).Result()
		if err != nil {
			return nil, xerror.NewError(xerror.RedisErr, "验证失败", err.Error())
		}
		if len(code) == 0 || code != in.GetCode() {
			return nil, xerror.NewError(xerror.ParamErr, "验证失败", errors.New(fmt.Sprintf("err code get:%s shoud:%s", in.GetCode(), code)))
		}
	}

	user := &data.UserModel{}

	switch in.GetType() {
	case "password":
		user, err = svc.accountRepo.GetByAccount(ctx, in.GetAccount())
		if err != nil {
			return nil, xerror.NewError(xerror.SQLErr, "", err.Error())
		}
		if user.Id == 0 {
			return nil, xerror.NewError(xerror.Invalid, "账号不存在")
		}
		if user.Password != util.Md5(in.GetPassword()) {
			return nil, xerror.NewError(xerror.Invalid, "密码错误")
		}
	case "email":
		user.Email = in.GetAccount()
		if err := svc.accountRepo.GetByEmail(ctx, user); err != nil {
			return nil, xerror.NewError(xerror.SQLErr, "", err.Error())
		}
		if user.Id == 0 {
			return nil, xerror.NewError(xerror.Invalid, "账号不存在")
		}
	case "mobile":
		user.Mobile = in.GetAccount()
		if err := svc.accountRepo.GetByMobile(ctx, user); err != nil {
			return nil, xerror.NewError(xerror.SQLErr, "", err.Error())
		}
		if user.Id == 0 {
			if err := svc.accountRepo.Create(ctx, &data.UserModel{UUID: uuid.NewString(), Mobile: in.GetAccount()}); err != nil {
				return nil, xerror.NewError(xerror.SQLErr, "用户注册失败，请重试", err.Error())
			}
		}
	}

	bus := xjwt.Business{
		UUID: user.UUID,
		Role: user.Role,
	}
	marshal, err := json.Marshal(bus)
	if err != nil {
		return nil, xerror.NewError(xerror.MarshalErr, "登录失败，请稍后重试！", err.Error())
	}
	token, err := xjwt.CreateToken(string(marshal), time.Hour*24)
	if err != nil {
		return nil, xerror.NewError(xerror.JwtErr, "登录失败，请稍后重试！", err.Error())
	}

	return &v1.LoginResp{Token: token}, nil
}

func (svc *AccountService) GetMiniQRCode(ctx context.Context, in *v1.GetMiniQRCodeReq) (*v1.GetMiniQRCodeResp, error) {
	// redis 存储随机数
	// 返回小程序码+随机数
	// 手机扫码调起小程序，登录后同步给随机数设置token
	// 前端轮询随机数，成功后获取到token
	return nil, status.Errorf(codes.Unimplemented, "method GetMiniQRCode not implemented")
}
