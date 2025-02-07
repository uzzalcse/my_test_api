package main

import (
	_ "my_test_api/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

