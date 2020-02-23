package main

import (
	"github.com/kataras/iris"
	"secondskillforgo/imooc_iris/repositories"
	"secondskillforgo/imooc_iris/services"
	"github.com/kataras/iris/mvc"
	"context"
	"secondskillforgo/imooc_iris/web/controllers"
	
	
)


func main(){
	//1. 创建iris例子 
	app := iris.New()
	//2. 设置错误等级
	//3. 注册模版
	template := iris.HTML("./immoc_iris/web/views",".html").Layout("ddd.html").Reload(true)
	app.RegisterView(template)
	//4. 设置模版目标
	//5. 出现异常跳到指定界面
	app.OnAnyErrorCode(func(ctx iris.Context){
		ctx.ViewData("message",ctx.Values().GetStringDefault("message","bad"))
		ctx.ViewLayout("")
		ctx.View("ddd.html")
	})
	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()



	//6 注册控制器
	productRepository := repositories.NewProcductManager("")
	productService := services.NewProductService(productRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx,productService)
	product.Handle(new(controllers.ProductController))

	//7 启动
	app.Run(
		iris.Addr("localhost:8080"),
	)




	
	
	


}