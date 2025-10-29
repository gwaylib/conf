// cache ini to memory
package ini

import (
	"log"
	"path/filepath"
	"sync"
	"time"

	"github.com/gwaylib/errors"
)

type cacheFile struct {
	File    *File
	Error   error
	EndedAt time.Time
}

func (c *cacheFile) IsTimeout(now time.Time) bool {
	return c.EndedAt.Before(now)
}

type IniCache struct {
	rootPath   string
	cache      sync.Map
	cacheOut   time.Duration
	readSignal chan string
}

// rootPath -- point out the etc directory
// timeout -- reread init file afeter timeout
func NewTimeoutIniCache(rootPath string, timeout time.Duration) *IniCache {
	i := &IniCache{
		rootPath:   rootPath,
		cacheOut:   timeout,
		readSignal: make(chan string, 200), // buffer for 200 concurrency
	}
	go i.read()
	return i
}

func NewIniCache(rootPath string) *IniCache {
	return NewTimeoutIniCache(rootPath, 5*time.Minute)
}

func (ini *IniCache) read() {
	for {
		filePath := <-ini.readSignal
		storeFile, ok := ini.cache.Load(filePath)
		if ok {
			cache := storeFile.(*cacheFile)
			if !cache.IsTimeout(time.Now()) {
				return
			}
		}

		file, err := GetFile(filePath)
		if err != nil {
			log.Println(errors.As(err, filePath))
			return
		}
		ini.cache.Store(filePath, &cacheFile{File: file, Error: err, EndedAt: time.Now().Add(ini.cacheOut)})
	}
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
		cache := storeFile.(*cacheFile)
		if cache.IsTimeout(time.Now()) {
			ini.readSignal <- filePath
		}
		return cache.File, cache.Error
	}

	// first time for read
	file, err = GetFile(filePath)
	if err != nil {
		err = errors.As(err, filePath)
	}
	ini.cache.Store(filePath, &cacheFile{File: file, Error: err, EndedAt: time.Now().Add(ini.cacheOut)})
	return file, err
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
