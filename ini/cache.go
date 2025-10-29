// cache ini to memory
package ini

import (
	"path/filepath"
	"sync"
	"time"

	"github.com/gwaylib/errors"
)

type CacheFile struct {
	File    *File
	EndedAt time.Time
}

func (c *CacheFile) IsTimeout(now time.Time) bool {
	return c.EndedAt.After(now)
}

type IniCache struct {
	rootPath string
	cache    sync.Map
	cacheOut time.Duration
}

// rootPath -- point out the etc directory
// timeout -- reread init file afeter timeout
func NewTimeoutIniCache(rootPath string, timeout time.Duration) *IniCache {
	return &IniCache{rootPath: rootPath, cacheOut: timeout}
}

func NewIniCache(rootPath string) *IniCache {
	return NewTimeoutIniCache(rootPath, 5*time.Minute)
}

func (ini *IniCache) DelCache(subFileName string) {
	filePath := filepath.Join(ini.rootPath, subFileName)
	ini.cache.Delete(filePath)
}

func (ini *IniCache) getFile(subFileName string) (*File, error) {
	filePath := filepath.Join(ini.rootPath, subFileName)

	var file *File
	var err error

	storeFile, ok := ini.cache.Load(filePath)
	if ok {
		cacheFile := storeFile.(*CacheFile)
		if !cacheFile.IsTimeout(time.Now()) {
			return cacheFile.File, nil
		}
		file, err = GetFile(filePath)
		if err != nil {
			println(errors.As(err, filePath))
			return cacheFile.File, nil
		}
	} else {
		file, err = GetFile(filePath)
		if err != nil {
			return nil, errors.As(err, filePath)
		}
	}

	ini.cache.Store(filePath, &CacheFile{File: file, EndedAt: time.Now().Add(ini.cacheOut)})
	return file, nil

}

func (ini *IniCache) GetFile(subFileName string) *File {
	file, err := ini.getFile(subFileName)
	if err != nil {
		panic(errors.As(err))
	}
	return file
}

func (ini *IniCache) GetDefaultFile(subFileName, subDefaultFileName string) *File {
	if len(subFileName) == 0 {
		return ini.GetFile(subDefaultFileName)
	}

	file, err := ini.getFile(subFileName)
	if err != nil {
		return ini.GetFile(subDefaultFileName)
	}
	return file
}
