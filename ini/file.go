package ini

import (
	"strings"
	"sync"
	"time"

	"github.com/go-ini/ini"
	"github.com/gwaylib/errors"
)

var cache = sync.Map{}

func GetFile(fileName string) (*File, error) {
	f, ok := cache.Load(fileName)
	if ok {
		return f.(*File), nil
	}
	file, err := ini.Load(fileName)
	if err != nil {
		if strings.Index(err.Error(), "no such file or directory") > -1 {
			return nil, errors.ErrNoData.As(err, fileName)
		}
		return nil, err
	}
	ff := &File{file}
	cache.Store(fileName, ff)
	return ff, nil
}

type File struct {
	*ini.File
}

func (f *File) String(section, key string) string {
	s := f.Section(section)
	result := s.Key(key).String()
	if len(result) == 0 && !s.HasKey(key) {
		panic(errors.ErrNoData.As(section, key))
	}
	return result
}
func (f *File) Bool(section, key string) bool {
	result, err := f.Section(section).Key(key).Bool()
	if err != nil {
		panic(errors.As(err, section, key))
	}
	return result
}
func (f *File) Float64(section, key string) float64 {
	result, err := f.Section(section).Key(key).Float64()
	if err != nil {
		panic(errors.As(err, section, key))
	}
	return result
}
func (f *File) Int(section, key string) int {
	result, err := f.Section(section).Key(key).Int()
	if err != nil {
		panic(errors.As(err, section, key))
	}
	return result
}
func (f *File) Int64(section, key string) int64 {
	result, err := f.Section(section).Key(key).Int64()
	if err != nil {
		panic(errors.As(err, section, key))
	}
	return result
}
func (f *File) Uint(section, key string) uint {
	result, err := f.Section(section).Key(key).Uint()
	if err != nil {
		panic(errors.As(err, section, key))
	}
	return result
}
func (f *File) Uint64(section, key string) uint64 {
	result, err := f.Section(section).Key(key).Uint64()
	if err != nil {
		panic(errors.As(err, section, key))
	}
	return result
}
func (f *File) Duration(section, key string) time.Duration {
	result, err := f.Section(section).Key(key).Duration()
	if err != nil {
		panic(errors.As(err, section, key))
	}
	return result
}
func (f *File) TimeFormat(section, key, format string) time.Time {
	result, err := f.Section(section).Key(key).TimeFormat(format)
	if err != nil {
		panic(errors.As(err, section, key))
	}
	return result
}
func (f *File) Time(section, key string) time.Time {
	result, err := f.Section(section).Key(key).Time()
	if err != nil {
		panic(errors.As(err, section, key))
	}
	return result
}
