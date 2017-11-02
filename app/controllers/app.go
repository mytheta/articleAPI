package controllers

import (
	"github.com/revel/revel"
  _ "github.com/go-sql-driver/mysql"
)

type App struct {
	*revel.Controller
}


func (c App) Index() revel.Result {

     return c.Render()
}
