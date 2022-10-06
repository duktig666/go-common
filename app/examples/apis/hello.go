// description:
// @author renshiwei
// Date: 2022/10/5 21:01

package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/qiaoshurui/couples-subtotal/app/examples/service"
	"github.com/qiaoshurui/couples-subtotal/common/api"
	"github.com/qiaoshurui/couples-subtotal/common/logger"
	"github.com/qiaoshurui/couples-subtotal/common/res"
)

type Hello struct {
	api.Api
}

func (h *Hello) SayHelloHandler(c *gin.Context) {
	name := c.Query("name")

	hello := service.Hello{}
	helloRes, err := hello.SayHello(name)
	if err != nil {
		logger.Errorf("", err)
	}
	res.Success(c, helloRes)
}
