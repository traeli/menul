package logic

import (
	"context"

	"menul-service/service/api/internal/svc"
	"menul-service/service/api/internal/types"

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
	resp = &types.GetCurrentFoodReqReply{}
	resp.Food = "test食物"
	resp.Image = "image"
	resp.Desc = "很好吃"
	resp.NearbyPrice = 23
	return resp, nil
}
