// description:
// @author renshiwei
// Date: 2022/10/5 21:10

package router

import (
	"github.com/duktig666/go-common/app/examples/apis"
	"github.com/gin-gonic/gin"
)

func RegisterHelloRouter(v1 *gin.RouterGroup) {
	hello := &apis.Hello{}
	{
		v1.GET("/hello/say", hello.SayHelloHandler)
	}

}
