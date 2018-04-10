# 使用说明

在shell中导入环境变量
export PJ_ROOT=$HOME/ws/go_project

使用ini包读取配置文件数据
``` text
package db

import (
	"github.com/CodeInEverest/golib/conf"
	"github.com/CodeInEverest/golib/database"
	_ "github.com/go-sql-driver/mysql"
)

var dbFile = conf.RootDir() + "/etc/db.cfg"

func DB(section string) *database.DB {
	return database.CacheDB(dbFile, section)
}

func HasDB(section string) (*database.DB, error) {
	return database.HasDB(dbFile, section)
}
```
