package config

import (
	"fmt"
	"testing"
)

func TestConfigMerge(t *testing.T) {
	//InitConfig("../conf/config-default.dev.yaml")
	var s string
	InitConfig(s)
	fmt.Println("cli.name", GlobalConfig.Cli.Name)
	fmt.Println("log.level", GlobalConfig.Log.Level)
}
