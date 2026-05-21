package logic

import (
	"context"

	"github.com/xxx-newbee/product/internal/svc"
	"github.com/xxx-newbee/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductLogic) UpdateProduct(in *product.UpdateProductRequest) (*product.UpdateProductResponse, error) {
	p, err := l.svcCtx.ProductModel.FindById(uint(in.ProductId))
	if err != nil {
		return &product.UpdateProductResponse{Success: false, Msg: err.Error()}, nil
	}
	if p == nil {
		return &product.UpdateProductResponse{Success: false, Msg: "商品不存在"}, nil
	}

	if in.Name != "" {
		p.Name = in.Name
	}
	if in.Description != "" {
		p.Description = in.Description
	}
	p.CategoryId = in.CategoryId
	if in.MainImage != "" {
		p.MainImage = in.MainImage
	}
	if in.Images != "" {
		p.Images = in.Images
	}
	p.Price = in.Price
	p.Stock = int(in.Stock)

	if err := l.svcCtx.ProductModel.Update(p); err != nil {
		return &product.UpdateProductResponse{Success: false, Msg: err.Error()}, nil
	}

	return &product.UpdateProductResponse{Success: true, Msg: "更新成功"}, nil
}
