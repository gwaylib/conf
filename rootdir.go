package conf

import (
	"os"
	"path/filepath"
)

func envRoot() string {
	dir := os.Getenv("PRJ_ROOT")
	if len(dir) == 0 {
		// Compatible the old version
		dir = os.Getenv("PJ_ROOT")
	}

	if len(dir) == 0 {
		panic("Need PRJ_ROOT environment for project directory")
	}

	p, err := filepath.Abs(dir)
	if err != nil {
		panic(err)
	}
	return p
}

var (
	rootDir = ""
)

func InitRootDir(path string) {
	rootDir = path
}

func RootDir() string {
	if len(rootDir) == 0 {
		rootDir = envRoot()
	}
	return rootDir
}
