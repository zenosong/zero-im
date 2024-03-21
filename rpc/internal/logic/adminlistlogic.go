package logic

import (
	"context"
	"zero/rpc/internal/svc"
	"zero/rpc/zero"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminListLogic {
	return &AdminListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Admin service
func (l *AdminListLogic) AdminList(in *zero.AdminListReq) (*zero.AdminListResp, error) {
	// todo: add your logic here and delete this line

	list, total, err := l.svcCtx.Query.User.WithContext(l.ctx).FindByPage(int(*in.Page), int(*in.Size))
	if err != nil {
		return nil, err
	}
	var data []*zero.AdminInfo
	for _, user := range list {
		data = append(data, &zero.AdminInfo{
			Id:       user.ID,
			Username: &user.Username,
			Password: &user.Password,
		})
	}
	i := int32(total)
	return &zero.AdminListResp{
		Data:  data,
		Total: &i,
	}, nil
}
