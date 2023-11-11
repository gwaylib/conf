# README

Export the env to the shell
```shell
export PRJ_ROOT=$HOME/ws/go_project
```

Example for `RootDir`
```golang
package db

import (
	"github.com/gwaylib/conf"
	"github.com//gwaylib/qsql"
	_ "github.com/go-sql-driver/mysql"
)

var dbFile = conf.RootDir() + "/etc/db.cfg"

func DB(section string) *qsql.DB {
	return qsql.CacheDB(dbFile, section)
}

func HasDB(section string) (*qsql.DB, error) {
	return qsql.HasDB(dbFile, section)
}
```

Example for ini read
```golang
package db

import (
	"github.com/gwaylib/conf"
	"github.com/gwaylib/conf/ini"
)

func main() {
    //
    // etc.cfg example
    // [test]
    // str: abc
    //
    path := conf.RootDir() + "/etc/etc.cfg"
    cfg := ini.NewIni(path)
    str := cfg.String("test", "str")
    if str != "abc" {
        panic("expect abc, but : " + str)
    }
}
```
