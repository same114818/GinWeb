package db

import (
	"GinWeb/configs"
	"GinWeb/global"
	models "GinWeb/models"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

//var engine *xorm.Engine

func InitDatabase(DbConnectionSetting *configs.Sqlite3Connection) error {
	var err error
	global.DBEngine, err = xorm.NewEngine(DbConnectionSetting.DbType, DbConnectionSetting.ConnectionString)
	if err != nil {
		return err
	}
	global.DBEngine.SetMapper(names.SameMapper{})
	global.DBEngine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	global.DBEngine.DatabaseTZ, _ = time.LoadLocation("Asia/Shanghai")
	err1 := global.DBEngine.Sync2(new(models.KnowledgeMain))
	if err1 != nil {
		return err1
	}
	return nil
}
