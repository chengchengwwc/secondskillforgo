package datamodels

import (
	
)

type Product struct {
	ID int64 `json:"id" imooc:"id" sql:"ID`
	ProductName string `json:"productName" immoc:"ProductName" sql:"productName"`
	ProductNum int64 `json:"productNum" sql:"productNum" immoc:"ProductNum"`
	ProductImages string `json:"productImages" sql:"productImages" immoc:"ProductImages"`
	ProductUrl string `json:"productUrl" sql:"productUrl" immoc:"ProductUrl"`
}