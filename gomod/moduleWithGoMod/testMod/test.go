package testMod

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

func Test() {
	fmt.Println("module with go mod, Test")
	fmt.Println(toml.Key{})
}
