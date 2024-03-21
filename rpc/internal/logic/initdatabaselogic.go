package logic

import (
	"context"

	"zero/rpc/internal/svc"
	"zero/rpc/zero"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitDatabaseLogic) InitDatabase(in *zero.Empty) (*zero.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &zero.BaseResp{}, nil
}
