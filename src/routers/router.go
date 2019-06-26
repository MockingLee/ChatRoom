package routers

import (
	"github.com/astaxie/beego"
	"../controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/test" , &controllers.ApiController{})
}
