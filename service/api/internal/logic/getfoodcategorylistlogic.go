package logic

import (
	"context"

	"menul-service/service/api/internal/svc"
	"menul-service/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFoodCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFoodCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFoodCategoryListLogic {
	return &GetFoodCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFoodCategoryListLogic) GetFoodCategoryList() (resp *types.GetFoodCategoryListReply, err error) {
	resp = &types.GetFoodCategoryListReply{}

	// 查询去重的 category 字段（即 group by category）
	list, err := l.svcCtx.FoodModel.Food.GetFoodCategoryList()

	if err != nil {
		return nil, err
	}

	for _, item := range list {
		resp.Item = append(resp.Item, types.GetFoodCategoryListReplyItem{
			Name: item,
		})
	}

	return resp, nil
}
