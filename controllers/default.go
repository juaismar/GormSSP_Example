package controllers

import (
	"github.com/astaxie/beego"

	"GormSSP_Example/models"

	SSP "github.com/juaismar/gormssp"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
}

func (c *MainController) SimplePaginate() {

	columns := make(map[int]SSP.Data)
	columns[0] = SSP.Data{Db: "id", Dt: "id", Formatter: nil}
	columns[1] = SSP.Data{Db: "name", Dt: "name", Formatter: nil}
	columns[2] = SSP.Data{Db: "email", Dt: "email", Formatter: nil}
	columns[3] = SSP.Data{Db: "age", Dt: "age", Formatter: nil}

	c.Data["json"] = SSP.Simple(c, models.Conn, "users", columns)
	c.ServeJSON()
}
