package logic

import (
	"context"

	"github.com/xxx-newbee/product/internal/model"
	"github.com/xxx-newbee/product/internal/svc"
	"github.com/xxx-newbee/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCategoryLogic {
	return &ListCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListCategoryLogic) ListCategory(in *product.ListCategoryRequest) (*product.ListCategoryResponse, error) {
	all, err := l.svcCtx.CategoryModel.FindAll()
	if err != nil {
		return &product.ListCategoryResponse{Success: false, Msg: err.Error()}, nil
	}

	tree := buildCategoryTree(all, 0)

	return &product.ListCategoryResponse{
		Categories: tree,
		Success:    true,
		Msg:        "ok",
	}, nil
}

func buildCategoryTree(all []*model.Category, parentId int64) []*product.CategoryInfo {
	var result []*product.CategoryInfo
	for _, c := range all {
		if c.ParentId == parentId {
			info := &product.CategoryInfo{
				CategoryId: uint32(c.ID),
				Name:       c.Name,
				ParentId:   c.ParentId,
				SortOrder:  int32(c.SortOrder),
				Status:     int32(c.Status),
				CreatedAt:  c.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:  c.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
			info.Children = buildCategoryTree(all, int64(c.ID))
			result = append(result, info)
		}
	}
	return result
}
