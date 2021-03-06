package controllers

import (
	"github.com/astaxie/beego"

	"GormSSP_Example/models"

	SSP "github.com/juaismar/gormssp/v2"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
}

func (c *MainController) SimplePaginate() {

	var err error

	columns := []SSP.Data{
		SSP.Data{Db: "id", Dt: "id", Formatter: nil},
		SSP.Data{Db: "name", Dt: "name", Formatter: nil},
		SSP.Data{Db: "email", Dt: "email", Formatter: nil},
		SSP.Data{Db: "age", Dt: "age", Formatter: nil},
	}

	c.Data["json"], err = SSP.Simple(c, models.Conn, "users", columns)
	if err != nil {
		beego.Info(err)
	}

	c.ServeJSON()
}
