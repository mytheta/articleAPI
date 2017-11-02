package controllers

import (
	"github.com/revel/revel"
  _ "github.com/go-sql-driver/mysql"
	 "github.com/jinzhu/gorm"
	 "strings"
	 "fmt"
)

type App struct {
	*revel.Controller
}


func (c App) Index() revel.Result {

     return c.RenderJSON(ret)
}
