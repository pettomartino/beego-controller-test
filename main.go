package main

import (
	_ "test-controller/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

