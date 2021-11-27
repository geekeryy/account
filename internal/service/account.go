package service

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	v1 "account/api/v1"
	"account/internal/data"
	"account/pkg/errcode"
	"account/pkg/redis"
	"account/pkg/util"
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
		return nil, xerror.NewError(errcode.SystemErr, "数据查询失败，请稍后再试")
	}

	respUser := v1.UserModel{}
	if err := copier.Copy(&respUser, user); err != nil {
		return nil, xerror.NewError(errcode.CopierErr, "数据查询失败，请稍后再试", err.Error())
	}

	return &v1.GetByIDResp{User: &respUser}, nil
}

func (svc *AccountService) MiniLogin(ctx context.Context, in *v1.MiniLoginReq) (*v1.MiniLoginResp, error) {
	session, err := svc.mini.JsCode2session(in.Code)
	if err != nil {
		return nil, xerror.NewError(errcode.WechatErr, "登录失败，请稍后重试！", err.Error())
	}

	user := data.UserModel{
		WechatOpenid: session.Openid,
	}
	if err := svc.accountRepo.GetByWechatOpenid(ctx, &user); err != nil {
		return nil, xerror.NewError(errcode.SQLErr, "登录失败，请稍后重试！", err.Error())
	}
	if user.Id == 0 {
		user.UUID = uuid.NewString()
		if err := svc.accountRepo.Create(ctx, &data.UserModel{
			UUID:         user.UUID,
			WechatOpenid: session.Openid,
		}); err != nil {
			return nil, xerror.NewError(errcode.SQLErr, "登录失败，请稍后重试！", err.Error())
		}
	}

	bus := xjwt.Business{
		UUID: user.UUID,
		Role: user.Role,
	}
	marshal, err := json.Marshal(bus)
	if err != nil {
		return nil, xerror.NewError(errcode.MarshalErr, "登录失败，请稍后重试！", err.Error())
	}
	token, err := xjwt.CreateToken(string(marshal), time.Hour*24)
	if err != nil {
		return nil, xerror.NewError(errcode.JwtErr, "登录失败，请稍后重试！", err.Error())
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
		return nil, xerror.NewError(errcode.AuthErr, "", err.Error())
	}
	if err := svc.accountRepo.Updates(ctx, &data.UserModel{
		UUID:      bus.UUID,
		NickName:  in.NickName,
		AvatarUrl: in.AvatarUrl,
	}); err != nil {
		return nil, xerror.NewError(errcode.SQLErr, "登录失败，请稍后重试！", err.Error())
	}
	return &v1.Empty{}, nil
}

func (svc *AccountService) SendMsgCode(ctx context.Context, in *v1.SendMsgCodeReq) (*v1.Empty, error) {
	rand.Seed(time.Now().Unix())
	code := rand.Intn(10000)
	if err := svc.redis.Set(ctx, fmt.Sprintf(redis.SmsLoginCode, util.Md5(in.GetMobile())), code, 5*time.Minute).Err(); err != nil {
		return nil, xerror.NewError(errcode.RedisErr, "发送失败，请重新发送", err.Error())
	}
	if err := svc.sms.SendCode(in.Mobile, code); err != nil {
		return nil, xerror.NewError(errcode.YunPianErr, "发送失败，请重新发送", err.Error())
	}
	return nil, nil
}

func (svc *AccountService) SmsLogin(ctx context.Context, in *v1.SmsLoginReq) (*v1.SmsLoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SmsLogin not implemented")
}
