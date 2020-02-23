package repositories


import (
	"secondskillforgo/imooc_iris/datamodels"
	"database/sql"
)

type IOrderRepository interface {
	Conn() error
	Insert(*datamodels.Order)(int64,error)
	Delete(int64) bool
	Update(*datamodels.Order) error
	SelectByKey(int64)(*datamodels.Order,error)
	SelectAll()([]*datamodels.Order,error)
}

type OrderManagerRepository struct {
	table string
	mysqlConn *sql.DB
}

func NewOrderManagerRepository(table string,ddd *sql.DB) IOrderRepository {
	return &OrderManagerRepository{
		table: table,
		mysqlConn:ddd,
	}
}


func (p *OrderManagerRepository) Conn() error{
	return nil
}

func (p *OrderManagerRepository) Insert(*datamodels.Order)(int64,error){
	return 0,nil
}

func (p *OrderManagerRepository) Delete(int64) bool{
	return true
}

func(p *OrderManagerRepository) Update(*datamodels.Order) error{
	return nil
}

func(p *OrderManagerRepository)SelectByKey(int64)(*datamodels.Order,error){
	return nil,nil
}

func(p *OrderManagerRepository) SelectAll()([]*datamodels.Order,error){
	return nil,nil
}















