package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type Father struct {
	Id int
	Name string

	Father_order []*Father_order `orm:"reverse(many)"`

}

type Father_order struct {
	Id int
	Order_data string

	Father *Father `orm:"rel(fk)"`
	

}

func init(){
	orm.RegisterDataBase("default","mysql","root:1234@tcp(127.0.0.1:3306)/logininfo?charset=utf8")
	orm.RegisterModel(new(Father),new(Father_order))
	orm.RunSyncdb("default",false,true)
	fmt.Println("hello")
}