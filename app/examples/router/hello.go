// description:
// @author renshiwei
// Date: 2022/10/5 21:10

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qiaoshurui/couples-subtotal/app/examples/apis"
)

func RegisterHelloRouter(v1 *gin.RouterGroup) {
	hello := &apis.Hello{}
	{
		v1.GET("/hello/say", hello.SayHelloHandler)
	}

}
