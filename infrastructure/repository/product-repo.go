package repository

import (
	"fmt"

	"github.com/lbaroni/imersao-desafio-01/domain/model"
	"github.com/jinzhu/gorm"
)

// type ProductRepositoryInterface interface {
// 	AddProduct(product *Product) (*Product, error)
// 	FindAllProducts()([]*Product, error)
// }

type ProductRepositoryDB struct {
	Db *gorm.DB
}

func (r ProductRepositoryDB) AddProduct(product *model.Product) (*model.Product, error) {
	err := r.Db.Create(product).Error

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r ProductRepositoryDB) FindAllProducts() ([]*model.Product, error) {
	var products []model.Product

	result := r.Db.Find(&products)

	if result.RowsAffected <= 0 {
		return nil, fmt.Errorf("no products found")
	}
	
	productsSlice := make([]*model.Product, len(products))
	for i, v:= range products {
		vAux := v
		productsSlice[i] = &vAux
	}

	return productsSlice, nil
}