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

For realtime ini read
```golang
package db

import (
	"github.com/gwaylib/conf"
	"github.com/gwaylib/conf/ini"
)

func main() {
    etcRoot := filepath.Join(conf.RootDir(), "etc")
    etc := ini.NewIni(etcRoot).GetFile("etc.ini") // read disk file every times.
    str := etc.String("test", "str")
    if str != "abc" {
        panic("expect abc, but : " + str)
    }
}
```

For memory cache ini read
```
package db

import (
	"github.com/gwaylib/conf"
	"github.com/gwaylib/conf/ini"
)

func main() {
    etcRoot := filepath.Join(conf.RootDir(), "etc")
    etcCache := ini.NewCacheIni(etcRoot).GetFile("etc.ini") // the cache should reload when loaded after 5*time.Minute
    str := etc.String("test", "str")
    if str != "abc" {
        panic("expect abc, but : " + str)
    }
}
```

