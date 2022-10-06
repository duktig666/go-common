// description:
// @author renshiwei
// Date: 2022/10/5 17:06

package global

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

const VERSION = "v0.0.1"

var (
	Gorm   *gorm.DB
	Viper  *viper.Viper
	Config ConfigYaml
	Router *gin.Engine
)

type ConfigYaml struct {
	Server struct {
		Name    string
		Version string
		Http    struct {
			Host string
			Port int
		}
	}

	Db struct {
		Dsn          string
		MaxOpenConns int
		MaxIdleConns int
	}

	Log struct {
		Level struct {
			Server string
			Gorm   string
		}
	}
}
