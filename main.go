package main

import (
	_ "streamApi/routers"
	"github.com/astaxie/beego"
)

import (
	"github.com/astaxie/beego/orm"
	"github.com/tkanos/gonfig"
	_ "github.com/go-sql-driver/mysql"
)

type Configuration struct {
	Dbuser string
	Dbpass string
	Dbname string
}

func init() {

	configuration := Configuration{}
	err := gonfig.GetConf("secrets.yml", &configuration)

	if err != nil {
		panic(err)
	}

	orm.Debug = true

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", configuration.Dbuser+":"+configuration.Dbpass+"@tcp(172.17.0.2:3306)/"+configuration.Dbname+"?charset=utf8")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

