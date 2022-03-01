package main

import (
	"github.com/aurorazl/go_module_version"
	"github.com/aurorazl/go_module_version/DirectoryA"
)

func main() {
	go_module_version.Test()
	DirectoryA.MethodA()
}
