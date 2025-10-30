// cache ini to memory
package ini

import (
	"context"
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
	timeout    time.Duration
	readSignal chan string
	loadLk     sync.Mutex
}

// rootPath -- point out the etc directory
// timeout -- reread init file afeter timeout
func NewTimeoutIniCache(rootPath string, timeout time.Duration) *IniCache {
	i := &IniCache{
		rootPath:   rootPath,
		timeout:    timeout,
		readSignal: make(chan string, 200), // buffer for 200 concurrency
	}
	go func() {
		ctx := context.Background()
		for {
			select {
			case filePath := <-i.readSignal:
				i.load(filePath)
			case <-ctx.Done():
			}
		}
	}()
	return i
}

func NewIniCache(rootPath string) *IniCache {
	return NewTimeoutIniCache(rootPath, 0)
}

func (ini *IniCache) load(filePath string) (*File, error) {
	ini.loadLk.Lock()
	defer ini.loadLk.Unlock()

	storeFile, ok := ini.cache.Load(filePath)
	if ok {
		cache := storeFile.(*cacheFile)
		if !cache.IsTimeout(time.Now()) {
			return cache.File, cache.Error
		}
	}

	file, err := GetFile(filePath)
	if err != nil {
		err = errors.As(err, filePath)
	}
	endedAt := time.Now().Add(ini.timeout)
	if ini.timeout == 0 {
		endedAt = time.Date(9999, 12, 31, 24, 0, 0, 0, time.Local)
	}
	ini.cache.Store(filePath, &cacheFile{File: file, Error: err, EndedAt: endedAt})
	return file, err

}

func (ini *IniCache) Reload(subFileName string) {
	filePath := filepath.Join(ini.rootPath, subFileName)
	storeFile, ok := ini.cache.Load(filePath)
	if ok {
		storeFile.(*cacheFile).EndedAt = time.Now()
		ini.cache.Store(filePath, storeFile)
	}
	ini.readSignal <- filePath
}

func (ini *IniCache) DelCache(subFileName string) {
	filePath := filepath.Join(ini.rootPath, subFileName)
	ini.cache.Delete(filePath)
}

func (ini *IniCache) ClearCache() {
	ini.cache.Clear()
}

func (ini *IniCache) getFile(subFileName string) (*File, error) {
	filePath := filepath.Join(ini.rootPath, subFileName)

	storeFile, ok := ini.cache.Load(filePath)
	if ok {
		cache := storeFile.(*cacheFile)
		if cache.IsTimeout(time.Now()) {
			ini.readSignal <- filePath
		}
		return cache.File, cache.Error
	}

	// read first time
	return ini.load(filePath)
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
