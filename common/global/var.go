// description:
// @author renshiwei
// Date: 2022/10/5 17:06

package global

import (
	"github.com/gin-gonic/gin"
)

var (
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

	Cli struct {
		Name string
	}

	Log struct {
		Level struct {
			Server string
			Gorm   string
		}
	}
}
