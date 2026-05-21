package logic

import (
	"context"

	"github.com/xxx-newbee/product/internal/svc"
	"github.com/xxx-newbee/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSkuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSkuLogic {
	return &UpdateSkuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSkuLogic) UpdateSku(in *product.UpdateSkuRequest) (*product.UpdateSkuResponse, error) {
	s, err := l.svcCtx.ProductSkuModel.FindById(uint(in.SkuId))
	if err != nil {
		return &product.UpdateSkuResponse{Success: false, Msg: err.Error()}, nil
	}
	if s == nil {
		return &product.UpdateSkuResponse{Success: false, Msg: "SKU不存在"}, nil
	}

	if in.SkuCode != "" {
		s.SkuCode = in.SkuCode
	}
	if in.Specs != "" {
		s.Specs = in.Specs
	}
	s.Price = in.Price
	s.Stock = int(in.Stock)
	s.Status = int(in.Status)

	if err := l.svcCtx.ProductSkuModel.Update(s); err != nil {
		return &product.UpdateSkuResponse{Success: false, Msg: err.Error()}, nil
	}

	return &product.UpdateSkuResponse{Success: true, Msg: "更新成功"}, nil
}
