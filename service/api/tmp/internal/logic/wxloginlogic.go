package logic

import (
	"context"

	"menul-service/service/api/tmp/internal/svc"
	"menul-service/service/api/tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WxLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxLoginLogic {
	return &WxLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxLoginLogic) WxLogin(req *types.WxLoginReq) (resp *types.WxLoginReply, err error) {
	// todo: add your logic here and delete this line

	return
}
