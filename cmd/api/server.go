// description:
// @author renshiwei
// Date: 2022/10/5 22:00

package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiaoshurui/couples-subtotal/app/examples/router"
	"github.com/qiaoshurui/couples-subtotal/common/global"
	"github.com/qiaoshurui/couples-subtotal/common/logger"
	"net/http"
)

func Run() {
	initRouter()
}

func initRouter() {
	global.Router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	router.InitHelloRouter(global.Router)

	err := global.Router.Run(fmt.Sprintf("%s:%d", global.Config.Server.Http.Host, global.Config.Server.Http.Port))
	if err != nil {
		logger.Errorf("Failed to start http server: %+v", err)
	}
}
