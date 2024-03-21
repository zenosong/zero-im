package logic

import (
	"context"

	"zero/rpc/internal/svc"
	"zero/rpc/zero"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminUpdateLogic {
	return &AdminUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminUpdateLogic) AdminUpdate(in *zero.AdminInfo) (*zero.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &zero.BaseResp{}, nil
}
