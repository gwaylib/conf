# README

Export the env to the shell
```shell
export PRJ_ROOT=$HOME/ws/test
```

Example for `RootDir`
```golang
package db

import (
    "os/filepath"

	"github.com/gwaylib/conf"
	"github.com//gwaylib/qsql"
	_ "github.com/go-sql-driver/mysql"
)

var dbFile = filepath.Join(conf.RootDir(),  "etc/db.cfg")

func DB(section string) *qsql.DB {
	return qsql.CacheDB(dbFile, section)
}

func HasDB(section string) (*qsql.DB, error) {
	return qsql.HasDB(dbFile, section)
}
```

Example for ini read

example $HOME/ws/test/etc/etc.ini
```
[test]
str: abc
```

```shell

```golang
package db

import (
	"github.com/gwaylib/conf"
	"github.com/gwaylib/conf/ini"
)

func main() {
    path := conf.RootDir() + "/etc/"
    etc := ini.NewIni(path).GetFile("etc.ini")
    str := etc.String("test", "str")
    if str != "abc" {
        panic("expect abc, but : " + str)
    }
}
```
