package main

import (
	"GormSSP_Example/db"
	_ "GormSSP_Example/routers"

	"github.com/astaxie/beego"
)

func main() {
	db.InitORM()
	beego.Run()
}
