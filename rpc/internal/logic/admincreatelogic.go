package logic

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"zero/rpc/dao/model"

	"zero/rpc/internal/svc"
	"zero/rpc/zero"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminCreateLogic {
	return &AdminCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminCreateLogic) AdminCreate(in *zero.AdminInfo) (*zero.BaseIDResp, error) {
	// todo: add your logic here and delete this line
	pass, _ := bcrypt.GenerateFromPassword([]byte(*in.Password), bcrypt.DefaultCost)

	admin := &model.Admin{
		Username: *in.Username,
		Password: string(pass),
		Avatar:   *in.Avatar,
		GroupID:  *in.GroupId,
		IsSuper:  *in.IsSuper,
	}
	err := l.svcCtx.Query.Admin.WithContext(l.ctx).Create(admin)
	if err != nil {
		l.Logger.Error("create admin error", err)
		return nil, err
	}
	return &zero.BaseIDResp{
		Id: admin.ID,
	}, nil
}
