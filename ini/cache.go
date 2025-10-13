// cache ini to memory
package ini

import (
	"path/filepath"
	"sync"

	"github.com/gwaylib/errors"
)

type IniCache struct {
	rootPath string
	cache    sync.Map
}

func NewIniCache(rootPath string) *IniCache {
	return &IniCache{rootPath: rootPath}
}

func (ini *IniCache) GetFile(subFileName string) *File {
	filePath := filepath.Join(ini.rootPath, subFileName)

	cacheFile, ok := ini.cache.Load(filePath)
	if ok {
		return cacheFile.(*File)
	}

	file, err := GetFile(filePath)
	if err != nil {
		panic(errors.As(err, filePath))
	}

	ini.cache.Store(filePath, file)

	return file
}

func (ini *IniCache) GetDefaultFile(subFileName, subDefaultFileName string) *File {
	if len(subFileName) == 0 {
		return ini.GetFile(subDefaultFileName)
	}

	filePath := filepath.Join(ini.rootPath, subFileName)
	cacheFile, ok := ini.cache.Load(filePath)
	if ok {
		return cacheFile.(*File)
	}

	file, err := GetFile(filePath)
	if err != nil {
		if !errors.ErrNoData.Equal(err) {
			panic(errors.As(err, filePath))
		}

		// cache the default file to current key when the target not exists.
		file = ini.GetFile(subDefaultFileName)
	}

	ini.cache.Store(filePath, file)
	return file
}
