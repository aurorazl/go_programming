package main

import (
	//sm "github.com/BurntSushi/toml"
	"gomod/moduleWithoutGoMod"
	"sample.com/testMod"
)

func main() {
	//fmt.Println(sm.Key{})
	moduleWithoutGoMod.Test()
	testMod.Test()
}
