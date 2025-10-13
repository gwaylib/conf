package ini

import (
	"github.com/gwaylib/errors"
)

// 用于省略前缀长路径写法
// 例如,以下可用于多语言处理：
// ini := NewIni(conf.RootDir()+"/app.default)
// lang := ".zh_cn"
// cfg := ini.GetFile(lang)
// cfg.String("msg", "1001")
type Ini struct {
	rootPath string
}

func NewIni(rootPath string) *Ini {
	return &Ini{rootPath}
}

func (ini *Ini) GetFile(subFileName string) *File {
	file, err := GetFile(ini.rootPath + subFileName)
	if err != nil {
		panic(errors.As(err, ini.rootPath+subFileName))
	}
	return file
}

func (ini *Ini) GetDefaultFile(subFileName, subDefaultFileName string) *File {
	f, err := GetFile(ini.rootPath + subFileName)
	if err != nil {
		if !errors.ErrNoData.Equal(err) {
			panic(errors.As(err, ini.rootPath+subFileName))
		}
		return ini.GetFile(subDefaultFileName)
	}
	return f
}
