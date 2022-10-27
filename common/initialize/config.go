// description:
// @author renshiwei
// Date: 2022/10/5 17:33

package initialize

import (
	"fmt"
	"github.com/duktig666/go-common/common/global"
	"github.com/duktig666/go-common/common/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig(configDir string) {
	v := viper.New()
	v.SetConfigFile(configDir)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	if err := v.Unmarshal(&global.Config); err != nil {
		logger.Errorf("读取配置文件失败。")
	}

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})

	global.Viper = v
}
