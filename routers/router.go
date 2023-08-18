package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"test-controller/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/users", &controllers.UserController{})
}
