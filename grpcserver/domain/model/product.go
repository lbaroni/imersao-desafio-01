package model

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

type ProductRepositoryInterface interface {
	AddProduct(product *Product) (*Product, error)
	FindAllProducts()([]*Product, error)
}

func init(){
 	govalidator.SetFieldsRequiredByDefault(true)
}

type Product struct {
	ID uint `json:"id" gorm:"primarykey; notnull" valid:"-"`
	Name string `json:"name" gorm:"type:varchar(100)" valid:"notnull"`
	Description string `json:"description" gorm:"type:varchar(255)" valid:"notnull"`
	Price float64 `json:"price" gorm:"type:float" valid:"float,notnull"`
}

func (p *Product) isValid() error {
	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return err
	}

	if p.Price <= 0 {
		return errors.New("invalid price")
	}

	return nil
}

func NewProduct(name string, description string, price float64) (*Product, error) {
	product := Product{
		Name: name,
		Description: description,
		Price: price,
	}

	err := product.isValid()
	if err != nil {
		return nil, err
	}

	return &product, nil
}


