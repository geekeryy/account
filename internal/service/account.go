package service

import (
	"context"

	v1 "account/api/v1"
	"account/pkg/errcode"
	"github.com/comeonjy/go-kit/pkg/xerror"
	"github.com/jinzhu/copier"
)

func (svc *AccountService) GetByID(ctx context.Context, in *v1.GetByIDReq) (*v1.GetByIDResp, error) {
	user, err := svc.accountRepo.Get(ctx,in.GetId())
	if err != nil {
		return nil, xerror.NewError(errcode.SystemErr,"数据查询失败，请稍后再试")
	}

	respUser := v1.UserModel{}
	if err := copier.Copy(&respUser, user); err != nil {
		return nil, xerror.NewError(errcode.CopierErr,"数据查询失败，请稍后再试",err)
	}

	return &v1.GetByIDResp{User: &respUser}, nil
}
