package conf

import (
	"fmt"
	"testing"
)

func TestRootDir(t *testing.T) {
	fmt.Println(RootDir())
}

func TestEtcDir(t *testing.T) {
	fmt.Println(EtcDir())
}
