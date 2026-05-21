package logic

import (
	"context"

	"github.com/xxx-newbee/product/internal/svc"
	"github.com/xxx-newbee/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductStatusLogic {
	return &UpdateProductStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductStatusLogic) UpdateProductStatus(in *product.UpdateProductStatusRequest) (*product.UpdateProductStatusResponse, error) {
	p, err := l.svcCtx.ProductModel.FindById(uint(in.ProductId))
	if err != nil {
		return &product.UpdateProductStatusResponse{Success: false, Msg: err.Error()}, nil
	}
	if p == nil {
		return &product.UpdateProductStatusResponse{Success: false, Msg: "商品不存在"}, nil
	}

	if in.Status != 0 && in.Status != 1 {
		return &product.UpdateProductStatusResponse{Success: false, Msg: "状态值无效，只能为0(下架)或1(上架)"}, nil
	}

	if err := l.svcCtx.ProductModel.UpdateStatus(uint(in.ProductId), int(in.Status)); err != nil {
		return &product.UpdateProductStatusResponse{Success: false, Msg: err.Error()}, nil
	}

	statusText := "下架"
	if in.Status == 1 {
		statusText = "上架"
	}
	return &product.UpdateProductStatusResponse{Success: true, Msg: statusText + "成功"}, nil
}
