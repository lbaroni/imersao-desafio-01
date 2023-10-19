package usecase

import (
	"fmt"

	"github.com/lbaroni/imersao-desafio-01/grpcserver/domain/model"
)

type ProductUseCase struct {
	ProductRepository model.ProductRepositoryInterface
}

func (p *ProductUseCase) AddProduct(name string, description string, price float64) (*model.Product,error){
	product, err := model.NewProduct(name,description,price)

	if err != nil {
		return nil, err
	}

	productadded, err := p.ProductRepository.AddProduct(product)

	if err != nil {
		return nil, fmt.Errorf("error adding product (%d) - %s", product.ID, product.Name)
	}    

	return productadded, nil
}

func (p *ProductUseCase) FindAllProducts() ([]*model.Product, error){
	products, err := p.ProductRepository.FindAllProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}