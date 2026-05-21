package logic

import (
	"context"

	"github.com/xxx-newbee/product/internal/model"
	"github.com/xxx-newbee/product/internal/svc"
	"github.com/xxx-newbee/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSkuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSkuLogic {
	return &CreateSkuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateSkuLogic) CreateSku(in *product.CreateSkuRequest) (*product.CreateSkuResponse, error) {
	p, err := l.svcCtx.ProductModel.FindById(uint(in.ProductId))
	if err != nil {
		return &product.CreateSkuResponse{Success: false, Msg: err.Error()}, nil
	}
	if p == nil {
		return &product.CreateSkuResponse{Success: false, Msg: "商品不存在"}, nil
	}

	id, err := l.svcCtx.ProductSkuModel.Insert(&model.ProductSku{
		ProductId: uint(in.ProductId),
		SkuCode:   in.SkuCode,
		Specs:     in.Specs,
		Price:     in.Price,
		Stock:     int(in.Stock),
		Status:    1,
	})
	if err != nil {
		return &product.CreateSkuResponse{Success: false, Msg: err.Error()}, nil
	}

	return &product.CreateSkuResponse{
		SkuId:   uint32(id),
		Success: true,
		Msg:     "创建成功",
	}, nil
}
