package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"secondskillforgo/imooc_iris/services"
)

type ProductController struct {
	Ctx            iris.Context
	ProductService services.IProductService
}

func (p *ProductController) GetAll() mvc.View {
	productArray, _ := p.ProductService.GetAllProduct()
	return mvc.View{
		Name: "product/view.html",
		Data: iris.Map{
			"productArray": productArray,
		},
	}
}

func (p *ProductController) PostUpdate() {
	p.Ctx.Request().ParseForm()

}

func (p *ProductController) GetOrder() mvc.View {

}
