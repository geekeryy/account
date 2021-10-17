package service

import (
	"context"
	"encoding/json"
	"time"

	v1 "account/api/v1"
	"account/internal/data"
	"account/pkg/errcode"
	"github.com/comeonjy/go-kit/pkg/xerror"
	"github.com/comeonjy/go-kit/pkg/xjwt"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
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
		if err:=svc.accountRepo.Create(ctx, &data.UserModel{
			UUID:         user.UUID,
			WechatOpenid: session.Openid,
		});err!=nil{
			return nil, xerror.NewError(errcode.SQLErr, "登录失败，请稍后重试！", err.Error())
		}
	}

	bus := xjwt.Business{
		UUID: user.UUID,
		Role: uint(user.Role),
	}
	marshal, err := json.Marshal(bus)
	if err != nil {
		return nil, xerror.NewError(errcode.MarshalErr, "登录失败，请稍后重试！", err.Error())
	}
	token, err := xjwt.CreateToken(string(marshal), time.Hour*24)
	if err != nil {
		return nil, xerror.NewError(errcode.JwtErr, "登录失败，请稍后重试！", err.Error())
	}

	return &v1.MiniLoginResp{Token: token}, nil
}
