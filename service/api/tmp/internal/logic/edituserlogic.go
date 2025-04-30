package logic

import (
	"context"

	"menul-service/service/api/tmp/internal/svc"
	"menul-service/service/api/tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditUserLogic {
	return &EditUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditUserLogic) EditUser(req *types.EditUserReq) (resp *types.EditUserReply, err error) {
	// todo: add your logic here and delete this line

	return
}
