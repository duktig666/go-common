// description:
// @author renshiwei
// Date: 2022/10/5 20:11

package initialize

import (
	"github.com/duktig666/go-common/common/global"
	"github.com/gin-gonic/gin"
)

func InitGin() {
	r := gin.New()
	// 添加默认的中间件
	r.Use(gin.Recovery())
	global.Router = r
}
