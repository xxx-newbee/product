package logic

import (
	"context"

	"github.com/xxx-newbee/product/internal/model"
	"github.com/xxx-newbee/product/internal/svc"
	"github.com/xxx-newbee/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCategoryLogic {
	return &CreateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCategoryLogic) CreateCategory(in *product.CreateCategoryRequest) (*product.CreateCategoryResponse, error) {
	if in.Name == "" {
		return &product.CreateCategoryResponse{Success: false, Msg: "分类名称不能为空"}, nil
	}

	id, err := l.svcCtx.CategoryModel.Insert(&model.Category{
		Name:      in.Name,
		ParentId:  in.ParentId,
		SortOrder: int(in.SortOrder),
		Status:    1,
	})
	if err != nil {
		return &product.CreateCategoryResponse{Success: false, Msg: err.Error()}, nil
	}

	return &product.CreateCategoryResponse{
		CategoryId: uint32(id),
		Success:    true,
		Msg:        "创建成功",
	}, nil
}
