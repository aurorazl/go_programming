module gomod

go 1.13

//replace github.com/qiniu/qmgo077 => github.com/labstack/echo v3.3.10+incompatible

//replace github.com/BurntSushi/toml v0.5.0 => github.com/BurntSushi/toml v0.3.0
//
//require github.com/BurntSushi/toml v0.5.0
//

replace sample.com/gomod/testMod v0.5.0 => /Users/kelvin/Documents/workstate/GoProJect/src/go_programming/gomod/moduleWithGoMod/testMod

require sample.com/gomod/testMod v0.5.0
