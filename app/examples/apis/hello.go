// description:
// @author renshiwei
// Date: 2022/10/5 21:01

package apis

import (
	"github.com/duktig666/go-common/app/examples/service"
	"github.com/duktig666/go-common/common/api"
	"github.com/duktig666/go-common/common/logger"
	"github.com/duktig666/go-common/common/res"
	"github.com/gin-gonic/gin"
)

type Hello struct {
	api.Api
}

func (h *Hello) SayHelloHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "duktig666")

	hello := service.Hello{}
	helloRes, err := hello.SayHello(name)
	if err != nil {
		logger.Errorf("", err)
	}
	res.Success(c, helloRes)
}
