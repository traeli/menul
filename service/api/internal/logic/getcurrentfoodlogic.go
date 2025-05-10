package logic

import (
	"context"
	"menul-service/service/api/internal/middleware"
	"menul-service/service/api/internal/svc"
	"menul-service/service/api/internal/types"
	"time"

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
	resp = &types.GetCurrentFoodReqReply{}
	table := l.svcCtx.FoodModel.Food

	// 获取当前时间段
	timePeriod := middleware.GetTimePeriod(time.Now())

	food, selectErr := table.WithContext(l.ctx).
		Where(table.TimePeriod.Eq(timePeriod)).
		//Order(field.RawExpr("RANDOM()")).
		Limit(1).
		First()

	if selectErr != nil {
		return nil, selectErr
	}

	resp.Food = food.Name
	resp.Image = food.Image
	resp.Desc = food.Desc
	resp.NearbyPrice = float64(food.Price)

	return resp, nil
}
