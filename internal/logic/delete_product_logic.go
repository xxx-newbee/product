package logic

import (
	"context"

	"github.com/xxx-newbee/product/internal/svc"
	"github.com/xxx-newbee/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteProductLogic) DeleteProduct(in *product.DeleteProductRequest) (*product.DeleteProductResponse, error) {
	p, err := l.svcCtx.ProductModel.FindById(uint(in.ProductId))
	if err != nil {
		return &product.DeleteProductResponse{Success: false, Msg: err.Error()}, nil
	}
	if p == nil {
		return &product.DeleteProductResponse{Success: false, Msg: "商品不存在"}, nil
	}

	if err := l.svcCtx.ProductModel.Delete(uint(in.ProductId)); err != nil {
		return &product.DeleteProductResponse{Success: false, Msg: err.Error()}, nil
	}

	return &product.DeleteProductResponse{Success: true, Msg: "删除成功"}, nil
}
