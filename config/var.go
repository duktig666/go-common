// description:
// @author renshiwei
// Date: 2022/10/5 17:06

//go:generate go-bindata -pkg=config -nocompress -o=default_conf.go ../conf/config-default.yaml

package config

var (
	GlobalConfig Config
)

type Config struct {
	Cli struct {
		Name string
	}

	Log struct {
		Level  string
		Format string
	}
}
