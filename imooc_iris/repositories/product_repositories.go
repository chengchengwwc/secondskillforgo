package repositories

// 先开发对一个接口
// 实现接口
import (
	"secondskillforgo/imooc_iris/datamodels"
	"database/sql"
)


type IPrdouct interface {
	Conn()(error)
	Insert(*datamodels.Product)(int64,error)
	Delete(int64)bool
	Update(*datamodels.Product) error
	SelectByKey(int64)(*datamodels.Product,error)
	SelectAll()([]*datamodels.Product,error)
}

type ProductManager struct {
	table string
	mysqlConn *sql.DB
}


func NewProcductManager(table string,db *sql.DB) IPrdouct {
	return &ProductManager{table: table,mysqlConn:db}
}

func (p *ProductManager) Conn()(error){
	return nil

}

func (p *ProductManager)Insert(*datamodels.Product)(int64,error){
	return 0,nil

}

func (p *ProductManager)Delete(id int64) bool {
	return true
}

func (p *ProductManager)Update(*datamodels.Product) error{
	return nil
}

func(p *ProductManager)SelectByKey(id int64)(*datamodels.Product,error){
	return nil,nil
}

func(p *ProductManager)SelectAll()([]*datamodels.Product,error){
	return nil,nil
}