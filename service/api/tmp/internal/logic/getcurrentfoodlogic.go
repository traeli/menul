package logic

import (
	"context"

	"menul-service/service/api/tmp/internal/svc"
	"menul-service/service/api/tmp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCurrentFoodLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCurrentFoodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCurrentFoodLogic {
	return &GetCurrentFoodLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCurrentFoodLogic) GetCurrentFood(req *types.GetCurrentFoodReq) (resp *types.GetCurrentFoodReqReply, err error) {
	// todo: add your logic here and delete this line

	return
}
