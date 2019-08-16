package controllers

import (
	"father/models"


	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type JSONStruct struct {
	Id int
	Name string

}

type ReController struct {
	beego.Controller
}

func (this *ReController)queryOrder(){
	var orders []models.Father_order
	o := orm.NewOrm()
	var qs orm.QuerySeter
	qs = o.QueryTable("Father_order")
	_,err := qs.Filter("id",1).All(orders)
	if err == nil{
		beego.Error("query erro")
		return

	}
	for _,order := range orders{
		beego.Info("queru order success,order =",order)
	}

//	var orders models.Father_order
//	err := orm.NewOrm().QueryTable("Father_order").Filter("id", 1).Limit(1).One(&orders)
//	if err == nil {
//		fmt.Println(orders)
//	}
//	this.Ctx.WriteString("查询成功")
//
//
}

//func(this *ReController)Search(){
//	name := this.Input().Get("name")
//	beego.Info(name)
//
//	o := orm.NewOrm()
//	goods := models.Father{}
//	goods.Name = name
//
//	err := o.Read(&goods,"Name")
//	if err!=nil{
//		beego.Info("查询失败",err)
//		this.Ctx.WriteString("查询失败")
//		return
//	}
//
//	mystruct := &JSONStruct{goods.Id,goods.Name}
//	this.Data["json"] = mystruct
//	this.ServeJSON()
//	beego.Info(goods.Id,goods.Name)
//	this.Ctx.WriteString("查询成功")
//}
