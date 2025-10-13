// read ini from file
package ini

import (
	"path/filepath"

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
	filePath := filepath.Join(ini.rootPath, subFileName)
	file, err := GetFile(filePath)
	if err != nil {
		panic(errors.As(err, filePath))
	}
	return file
}

func (ini *Ini) GetDefaultFile(subFileName, subDefaultFileName string) *File {
	if len(subFileName) == 0 {
		return ini.GetFile(subDefaultFileName)
	}

	filePath := filepath.Join(ini.rootPath, subFileName)
	f, err := GetFile(filePath)
	if err != nil {
		if !errors.ErrNoData.Equal(err) {
			panic(errors.As(err, filePath))
		}
		return ini.GetFile(subDefaultFileName)
	}
	return f
}
