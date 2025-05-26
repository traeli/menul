package logic

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"menul-service/service/model"
	"time"

	"menul-service/service/api/internal/svc"
	"menul-service/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req *types.OrderReq) (resp *types.OrderReply, err error) {
	resp = &types.OrderReply{}

	if len(req.Foods) == 0 {
		return nil, errors.New("foods is empty")
	}

	table := l.svcCtx.FoodModel.Food
	foodMap := make(map[string]*model.Food, len(req.Foods))
	foodsList, err := table.WithContext(l.ctx).Where(table.ID.In(req.Foods...)).Find()
	for _, food := range foodsList {
		foodMap[food.ID] = food
	}

	db := l.svcCtx.DBEngin
	db.WithContext(l.ctx)
	err = db.WithContext(l.ctx).Transaction(func(tx *gorm.DB) error {
		order := model.Order{
			CreateAt: time.Now(),
			ID:       uuid.New().String(),
			//UserID:   , // 假设 OrderReq 中有 UserID
			Status: "pending",
		}
		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		var orderItems []model.OrderItem
		for _, foodID := range req.Foods {
			item := model.OrderItem{
				ID:      uuid.New().String(),
				OrderID: order.ID,
				FoodID:  foodID,
				Name:    foodMap[foodID].Name,
				Price:   foodMap[foodID].Price,
			}
			orderItems = append(orderItems, item)
		}

		if err := tx.Create(&orderItems).Error; err != nil {
			return err
		}

		resp.OrderID = order.ID
		return nil
	})

	if err != nil {
		return nil, err
	}

	return resp, nil
}
