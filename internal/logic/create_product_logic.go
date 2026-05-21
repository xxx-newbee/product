package logic

import (
	"context"

	"github.com/xxx-newbee/product/internal/model"
	"github.com/xxx-newbee/product/internal/svc"
	"github.com/xxx-newbee/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateProductLogic) CreateProduct(in *product.CreateProductRequest) (*product.CreateProductResponse, error) {
	if in.Name == "" {
		return &product.CreateProductResponse{Success: false, Msg: "商品名称不能为空"}, nil
	}
	if in.Price < 0 {
		return &product.CreateProductResponse{Success: false, Msg: "商品价格不能为负数"}, nil
	}

	id, err := l.svcCtx.ProductModel.Insert(&model.Product{
		Name:        in.Name,
		Description: in.Description,
		CategoryId:  in.CategoryId,
		MainImage:   in.MainImage,
		Images:      in.Images,
		Price:       in.Price,
		Stock:       int(in.Stock),
		Status:      1,
	})
	if err != nil {
		return &product.CreateProductResponse{Success: false, Msg: err.Error()}, nil
	}

	return &product.CreateProductResponse{
		ProductId: uint32(id),
		Success:   true,
		Msg:       "创建成功",
	}, nil
}
