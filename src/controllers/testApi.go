package controllers

import "github.com/astaxie/beego"

type ApiController struct {
	beego.Controller
}


type testStrut struct {
Code int `json:"code"`
Desc string `json:"desc"`
}


func (c * ApiController) Get() {
	mystruct := testStrut{
		Code:1,
		Desc:"hh",
	}
	c.Data["json"] = &mystruct
	c.ServeJSON()
}
