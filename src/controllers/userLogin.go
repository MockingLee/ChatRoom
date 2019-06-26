package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type UserLoginControllers struct {
	beego.Controller
}

func (c *UserLoginControllers) Get(){
	fmt.Println(c.Data)
	c.ServeJSON()
}

type user struct {
	Username string `json:"username"`
}

func (this  *UserLoginControllers) Post(){
	var user user
	data := this.Ctx.Input.RequestBody
	fmt.Println(data)
	//json数据封装到user对象中
	err := json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	fmt.Println(user)
	this.Data["json"] = "{" +
		"\"success\":\"" + "true" + "\"}"
	this.ServeJSON()
}