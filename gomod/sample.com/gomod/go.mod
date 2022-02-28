module sample.com/gomod

go 1.13

//replace github.com/qiniu/qmgo077 => github.com/labstack/echo v3.3.10+incompatible

replace (
	github.com/BurntSushi/toml v0.5.0 => github.com/BurntSushi/toml v0.3.0
	sample.com/testMod v0.5.0 => D:\workstate\GoProJect\src\sample.com\testMod
)

require github.com/BurntSushi/toml v0.5.0

require (
	github.com/aurorazl/go_module_version v1.1.0
	sample.com/testMod v0.5.0
)
