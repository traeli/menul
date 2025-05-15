package logic

import (
	"context"
	"menul-service/service/api/internal/svc"
	"menul-service/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFoodListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFoodListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFoodListLogic {
	return &GetFoodListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFoodListLogic) GetFoodList(req *types.GetFoodListReq) (resp *types.GetFoodListReqReply, err error) {
	resp = &types.GetFoodListReqReply{}
	table := l.svcCtx.FoodModel.Food

	foods, selectErr := table.WithContext(l.ctx).
		Where(table.Category.Eq(req.Category)).
		Offset(int(req.Page * req.PageSize)).
		Limit(int(req.PageSize)).
		Find()

	if selectErr != nil {
		return nil, selectErr
	}

	for _, f := range foods {
		resp.List = append(resp.List, types.GetFoodListReqReplyItem{
			Food:        f.Name,
			Desc:        f.Desc,
			NearbyPrice: float64(f.Price),
			Image:       f.Image,
		})
	}

	return
}
