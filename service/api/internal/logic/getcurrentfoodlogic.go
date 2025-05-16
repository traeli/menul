package logic

import (
	"context"
	"math/rand"
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

	timePeriod := middleware.GetTimePeriod(time.Now())

	// 查询对象初始化
	query := table.WithContext(l.ctx).Where(table.TimePeriod.Eq(timePeriod))

	// 如果有关键词，则按名称模糊查找
	if req.Food != "" {
		query = query.Where(table.Name.Like(req.Food))

		// 尝试查找一个匹配项
		food, selectErr := query.First()
		if selectErr != nil {
			resp.Food = "查询结果为空"
			return resp, nil
		}

		resp.Food = food.Name
		resp.Image = food.Image
		resp.Desc = food.Desc
		resp.NearbyPrice = float64(food.Price)
		return resp, nil
	}

	// 如果关键词为空，则随机查找
	offset := rand.Intn(10)
	food, selectErr := query.Offset(offset).Limit(1).First()
	if selectErr != nil {
		return nil, selectErr
	}

	resp.Food = food.Name
	resp.Image = food.Image
	resp.Desc = food.Desc
	resp.NearbyPrice = float64(food.Price)

	return resp, nil
}
