package model

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type (
	Product struct {
		gorm.Model
		Name        string  `db:"name"`
		Description string  `db:"description"`
		CategoryId  int64   `db:"category_id"`
		MainImage   string  `db:"main_image"`
		Images      string  `db:"images"` // json数组
		Price       float64 `db:"price"`
		Stock       int     `db:"stock"`
		Status      int     `db:"status"` // 0下架 1上架
		SalesCount  int     `db:"sales_count"`
	}

	ProductModel interface {
		Insert(data *Product) (uint, error)
		Update(data *Product) error
		Delete(id uint) error
		FindById(id uint) (*Product, error)
		List(page, pageSize int, categoryId int64, status int, keyword string) ([]*Product, int64, error)
		UpdateStatus(id uint, status int) error
	}

	defaultProductModel struct {
		db    *gorm.DB
		table string
	}
)

func (Product) TableName() string {
	return "products"
}

func NewProductModel(db *gorm.DB) ProductModel {
	return &defaultProductModel{
		db:    db,
		table: "products",
	}
}

func (m *defaultProductModel) Insert(data *Product) (uint, error) {
	if err := m.db.Create(data).Error; err != nil {
		return 0, err
	}
	return data.ID, nil
}

func (m *defaultProductModel) Update(data *Product) error {
	return m.db.Table(m.table).Where("id = ?", data.ID).Select("*").Updates(data).Error
}

func (m *defaultProductModel) Delete(id uint) error {
	return m.db.Table(m.table).Where("id = ?", id).Delete(&Product{}).Error
}

func (m *defaultProductModel) FindById(id uint) (*Product, error) {
	var p Product
	res := m.db.Table(m.table).Where("id = ?", id).First(&p)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	return &p, nil
}

func (m *defaultProductModel) List(page, pageSize int, categoryId int64, status int, keyword string) ([]*Product, int64, error) {
	var list []*Product
	var total int64

	query := m.db.Table(m.table)
	if categoryId > 0 {
		query = query.Where("category_id = ?", categoryId)
	}
	if status >= 0 {
		query = query.Where("status = ?", status)
	}
	if keyword != "" {
		query = query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", keyword))
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	res := query.Order("id DESC").Offset(offset).Limit(pageSize).Find(&list)
	return list, total, res.Error
}

func (m *defaultProductModel) UpdateStatus(id uint, status int) error {
	return m.db.Table(m.table).Where("id = ?", id).Update("status", status).Error
}
