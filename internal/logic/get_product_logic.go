package logic

import (
	"context"

	"github.com/xxx-newbee/product/internal/svc"
	"github.com/xxx-newbee/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductLogic {
	return &GetProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductLogic) GetProduct(in *product.GetProductRequest) (*product.GetProductResponse, error) {
	p, err := l.svcCtx.ProductModel.FindById(uint(in.ProductId))
	if err != nil {
		return &product.GetProductResponse{Success: false, Msg: err.Error()}, nil
	}
	if p == nil {
		return &product.GetProductResponse{Success: false, Msg: "商品不存在"}, nil
	}

	return &product.GetProductResponse{
		ProductId:   uint32(p.ID),
		Name:        p.Name,
		Description: p.Description,
		CategoryId:  p.CategoryId,
		MainImage:   p.MainImage,
		Images:      p.Images,
		Price:       p.Price,
		Stock:       int32(p.Stock),
		Status:      int32(p.Status),
		SalesCount:  int32(p.SalesCount),
		CreatedAt:   p.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   p.UpdatedAt.Format("2006-01-02 15:04:05"),
		Success:     true,
		Msg:         "ok",
	}, nil
}
