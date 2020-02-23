package services

import (
	"secondskillforgo/imooc_iris/datamodels"
	"secondskillforgo/imooc_iris/repositories"
)



type IProductService interface {
	GetProductById(int64)(*datamodels.Product,error)
	GetAllProduct()([]*datamodels.Product,error)
	DeleteProductById(int64)bool
	InsertProduct(*datamodels.Product)(int64,error)
	UpdateProduct(datamodels.Product)(error)
}

type ProductService struct {
	productRepository repositories.IPrdouct

}


func NewProductService(repository repositories.IPrdouct)IProductService{
	return &ProductService{productRepository:repository}
}

func(p *ProductService)GetProductById(int64)(*datamodels.Product,error){
	return nil,nil
}

func(p *ProductService)GetAllProduct()([]*datamodels.Product,error){
	return nil,nil
}

func (p *ProductService)DeleteProductById(int64)bool{
	return true
}

func(p *ProductService)InsertProduct(*datamodels.Product)(int64,error){
	return 0,nil
}

func(p *ProductService)UpdateProduct(datamodels.Product)(error){
	return nil
}











