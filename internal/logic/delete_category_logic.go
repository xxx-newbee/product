package logic

import (
	"context"

	"github.com/xxx-newbee/product/internal/svc"
	"github.com/xxx-newbee/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryLogic {
	return &DeleteCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCategoryLogic) DeleteCategory(in *product.DeleteCategoryRequest) (*product.DeleteCategoryResponse, error) {
	id := uint(in.CategoryId)

	category, err := l.svcCtx.CategoryModel.FindById(id)
	if err != nil {
		return &product.DeleteCategoryResponse{Success: false, Msg: err.Error()}, nil
	}
	if category == nil {
		return &product.DeleteCategoryResponse{Success: false, Msg: "分类不存在"}, nil
	}

	// 检查是否有子分类
	count, err := l.svcCtx.CategoryModel.CountByParentId(int64(id))
	if err != nil {
		return &product.DeleteCategoryResponse{Success: false, Msg: err.Error()}, nil
	}
	if count > 0 {
		return &product.DeleteCategoryResponse{Success: false, Msg: "该分类下存在子分类，无法删除"}, nil
	}

	if err := l.svcCtx.CategoryModel.Delete(id); err != nil {
		return &product.DeleteCategoryResponse{Success: false, Msg: err.Error()}, nil
	}

	return &product.DeleteCategoryResponse{Success: true, Msg: "删除成功"}, nil
}
