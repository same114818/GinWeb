package main

import (
	"GinWeb/configs"
	"GinWeb/db"
	"GinWeb/global"
	"GinWeb/routers"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	err := setupConfigs()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

// @title GinWeb Demo
// @version 1.0
// @description 这里写描述信息
// @termsOfService http://swagger.io/terms/
// @contact.name 沈洁
// @contact.url https://www.baidu.com
// @contact.email same114818@hotmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	r := gin.Default()
	routers.SetupRouters(r)
	//fmt.Println(global.DbConnectionSetting.ConnectionString)
	time.LoadLocation("Asia/Shanghai”")

	err := db.InitDatabase(global.DbConnectionSetting)
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	r.Run()
}

func setupConfigs() error {
	setting, err := configs.InitSetting()
	if err != nil {
		return err
	}
	err = setting.ReadConfigs("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadConfigs("Sqlite3Connection", &global.DbConnectionSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}
