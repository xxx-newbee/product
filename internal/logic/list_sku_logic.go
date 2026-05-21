package logic

import (
	"context"

	"github.com/xxx-newbee/product/internal/svc"
	"github.com/xxx-newbee/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSkuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSkuLogic {
	return &ListSkuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListSkuLogic) ListSku(in *product.ListSkuRequest) (*product.ListSkuResponse, error) {
	list, err := l.svcCtx.ProductSkuModel.FindByProductId(uint(in.ProductId))
	if err != nil {
		return &product.ListSkuResponse{Success: false, Msg: err.Error()}, nil
	}

	var skus []*product.SkuInfo
	for _, s := range list {
		skus = append(skus, &product.SkuInfo{
			SkuId:     uint32(s.ID),
			ProductId: uint32(s.ProductId),
			SkuCode:   s.SkuCode,
			Specs:     s.Specs,
			Price:     s.Price,
			Stock:     int32(s.Stock),
			Status:    int32(s.Status),
			CreatedAt: s.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &product.ListSkuResponse{
		Skus:    skus,
		Success: true,
		Msg:     "ok",
	}, nil
}
