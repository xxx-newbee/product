package model

import (
	"errors"

	"gorm.io/gorm"
)

type (
	Category struct {
		gorm.Model
		Name      string `db:"name"`
		ParentId  int64  `db:"parent_id"`
		SortOrder int    `db:"sort_order"`
		Status    int    `db:"status"` // 0禁用 1启用
	}

	CategoryModel interface {
		Insert(data *Category) (uint, error)
		Update(data *Category) error
		Delete(id uint) error
		FindById(id uint) (*Category, error)
		FindByParentId(parentId int64) ([]*Category, error)
		FindAll() ([]*Category, error)
		CountByParentId(parentId int64) (int64, error)
	}

	defaultCategoryModel struct {
		db    *gorm.DB
		table string
	}
)

func (Category) TableName() string {
	return "categories"
}

func NewCategoryModel(db *gorm.DB) CategoryModel {
	return &defaultCategoryModel{
		db:    db,
		table: "categories",
	}
}

func (m *defaultCategoryModel) Insert(data *Category) (uint, error) {
	if err := m.db.Create(data).Error; err != nil {
		return 0, err
	}
	return data.ID, nil
}

func (m *defaultCategoryModel) Update(data *Category) error {
	return m.db.Table(m.table).Where("id = ?", data.ID).Updates(data).Error
}

func (m *defaultCategoryModel) Delete(id uint) error {
	return m.db.Table(m.table).Where("id = ?", id).Delete(&Category{}).Error
}

func (m *defaultCategoryModel) FindById(id uint) (*Category, error) {
	var c Category
	res := m.db.Table(m.table).Where("id = ?", id).First(&c)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	return &c, nil
}

func (m *defaultCategoryModel) FindByParentId(parentId int64) ([]*Category, error) {
	var list []*Category
	res := m.db.Table(m.table).Where("parent_id = ?", parentId).Order("sort_order ASC").Find(&list)
	return list, res.Error
}

func (m *defaultCategoryModel) FindAll() ([]*Category, error) {
	var list []*Category
	res := m.db.Table(m.table).Order("sort_order ASC").Find(&list)
	return list, res.Error
}

func (m *defaultCategoryModel) CountByParentId(parentId int64) (int64, error) {
	var count int64
	res := m.db.Table(m.table).Where("parent_id = ?", parentId).Count(&count)
	return count, res.Error
}
