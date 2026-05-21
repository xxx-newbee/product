package logic

import (
	"context"

	"github.com/xxx-newbee/product/internal/svc"
	"github.com/xxx-newbee/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(in *product.UpdateCategoryRequest) (*product.UpdateCategoryResponse, error) {
	category, err := l.svcCtx.CategoryModel.FindById(uint(in.CategoryId))
	if err != nil {
		return &product.UpdateCategoryResponse{Success: false, Msg: err.Error()}, nil
	}
	if category == nil {
		return &product.UpdateCategoryResponse{Success: false, Msg: "分类不存在"}, nil
	}

	if in.Name != "" {
		category.Name = in.Name
	}
	category.ParentId = in.ParentId
	category.SortOrder = int(in.SortOrder)
	category.Status = int(in.Status)

	if err := l.svcCtx.CategoryModel.Update(category); err != nil {
		return &product.UpdateCategoryResponse{Success: false, Msg: err.Error()}, nil
	}

	return &product.UpdateCategoryResponse{Success: true, Msg: "更新成功"}, nil
}
