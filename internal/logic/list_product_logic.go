package logic

import (
	"context"

	"github.com/xxx-newbee/product/internal/svc"
	"github.com/xxx-newbee/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProductLogic {
	return &ListProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListProductLogic) ListProduct(in *product.ListProductRequest) (*product.ListProductResponse, error) {
	page := int(in.Page)
	pageSize := int(in.PageSize)
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	list, total, err := l.svcCtx.ProductModel.List(page, pageSize, in.CategoryId, int(in.Status), in.Keyword)
	if err != nil {
		return &product.ListProductResponse{Success: false, Msg: err.Error()}, nil
	}

	var products []*product.ProductInfo
	for _, p := range list {
		products = append(products, &product.ProductInfo{
			ProductId:   uint32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			CategoryId:  p.CategoryId,
			MainImage:   p.MainImage,
			Price:       p.Price,
			Stock:       int32(p.Stock),
			Status:      int32(p.Status),
			SalesCount:  int32(p.SalesCount),
			CreatedAt:   p.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   p.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &product.ListProductResponse{
		Products: products,
		Total:    total,
		Success:  true,
		Msg:      "ok",
	}, nil
}
