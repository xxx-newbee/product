package logic

import (
	"context"

	"github.com/xxx-newbee/product/internal/svc"
	"github.com/xxx-newbee/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSkuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSkuLogic {
	return &DeleteSkuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteSkuLogic) DeleteSku(in *product.DeleteSkuRequest) (*product.DeleteSkuResponse, error) {
	s, err := l.svcCtx.ProductSkuModel.FindById(uint(in.SkuId))
	if err != nil {
		return &product.DeleteSkuResponse{Success: false, Msg: err.Error()}, nil
	}
	if s == nil {
		return &product.DeleteSkuResponse{Success: false, Msg: "SKU不存在"}, nil
	}

	if err := l.svcCtx.ProductSkuModel.Delete(uint(in.SkuId)); err != nil {
		return &product.DeleteSkuResponse{Success: false, Msg: err.Error()}, nil
	}

	return &product.DeleteSkuResponse{Success: true, Msg: "删除成功"}, nil
}
