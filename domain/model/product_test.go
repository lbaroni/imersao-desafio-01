package model

import (
	"testing"
)

func TestNewProduct(t *testing.T){
	name := "Product 01"
	description := "Product 01 description"
	price := 100.00

	_, err := NewProduct(name, description, price)

	if err != nil {
		t.Fatalf("NewProduct(%s,%s,%f) failed with error: %v",name,description,price,err)
	}
}

func TestNewProductFail(t *testing.T){
	name := ""
	description := "Product 01 description"
	price := 100.00

	_, err := NewProduct(name, description, price)

	if err == nil {
		t.Fatalf("NewProduct(%s,%s,%f) should fail: invalid name",name,description,price)
	}

	name = "Product 01"
	description = ""
	price = 100.00

	_, err = NewProduct(name, description, price)

	if err == nil {
		t.Fatalf("NewProduct(%s,%s,%f) should fail: invalid description",name,description,price)
	}

	name = "Product 01"
	description = "Product 01 description"
	price = -100.00

	_, err = NewProduct(name, description, price)

	if err == nil {
		t.Fatalf("NewProduct(%s,%s,%f) should fail: invalid price",name,description,price)
	}
}