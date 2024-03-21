package logic

import (
	"context"

	"zero/rpc/internal/svc"
	"zero/rpc/zero"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminDetailLogic {
	return &AdminDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminDetailLogic) AdminDetail(in *zero.IDReq) (*zero.AdminInfo, error) {
	// todo: add your logic here and delete this line

	return &zero.AdminInfo{}, nil
}
