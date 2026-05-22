package model

import (
	"errors"

	"gorm.io/gorm"
)

type (
	ProductSku struct {
		gorm.Model
		ProductId uint    `db:"product_id"`
		SkuCode   string  `db:"sku_code"`
		Specs     string  `db:"specs"` // json: {"颜色":"红","尺寸":"L"}
		Price     float64 `db:"price"`
		Stock     int     `db:"stock"`
		Status    int     `db:"status"` // 0禁用 1启用
	}

	ProductSkuModel interface {
		Insert(data *ProductSku) (uint, error)
		Update(data *ProductSku) error
		Delete(id uint) error
		FindById(id uint) (*ProductSku, error)
		FindByProductId(productId uint) ([]*ProductSku, error)
	}

	defaultProductSkuModel struct {
		db    *gorm.DB
		table string
	}
)

func (ProductSku) TableName() string {
	return "product_skus"
}

func NewProductSkuModel(db *gorm.DB) ProductSkuModel {
	return &defaultProductSkuModel{
		db:    db,
		table: "product_skus",
	}
}

func (m *defaultProductSkuModel) Insert(data *ProductSku) (uint, error) {
	if err := m.db.Create(data).Error; err != nil {
		return 0, err
	}
	return data.ID, nil
}

func (m *defaultProductSkuModel) Update(data *ProductSku) error {
	return m.db.Table(m.table).Where("id = ?", data.ID).Select("*").Updates(data).Error
}

func (m *defaultProductSkuModel) Delete(id uint) error {
	return m.db.Table(m.table).Where("id = ?", id).Delete(&ProductSku{}).Error
}

func (m *defaultProductSkuModel) FindById(id uint) (*ProductSku, error) {
	var s ProductSku
	res := m.db.Table(m.table).Where("id = ?", id).First(&s)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	return &s, nil
}

func (m *defaultProductSkuModel) FindByProductId(productId uint) ([]*ProductSku, error) {
	var list []*ProductSku
	res := m.db.Table(m.table).Where("product_id = ?", productId).Find(&list)
	return list, res.Error
}
