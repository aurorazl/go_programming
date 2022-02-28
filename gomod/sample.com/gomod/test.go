package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/aurorazl/go_module_version/v2"
	sample "sample.com/testMod"
)

func main(){
	fmt.Println(toml.Key{})
	sample.Test()
	go_module_version.Test()
}