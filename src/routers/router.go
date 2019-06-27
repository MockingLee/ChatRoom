package routers

import (
	"../controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/test" , &controllers.ApiController{})
    beego.Router("/api/userLogin" , &controllers.UserLoginControllers{})
    beego.Router("/api/userin" , &controllers.WebSocketController{})
    beego.Router("/api/ws",&controllers.WebSocketController{} , "get:Join")
}
