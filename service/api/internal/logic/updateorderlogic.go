package logic

import (
	"context"

	"menul-service/service/api/internal/svc"
	"menul-service/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderLogic {
	return &UpdateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOrderLogic) UpdateOrder(req *types.OrderUpdateReq) (resp *types.OrderReply, err error) {
	// todo: add your logic here and delete this line

	return
}
