package conf

import (
	"os"
	"path/filepath"
)

func RootDir() string {
	dir := os.Getenv("GOSPACE")
	if len(dir) == 0 {
		// Compatible the old version
		dir = os.Getenv("PJ_ROOT")
	}

	if len(dir) == 0 {
		panic("Need GOSPACE environment for project directory")
	}

	p, err := filepath.Abs(dir)
	if err != nil {
		panic(err)
	}
	return p
}
