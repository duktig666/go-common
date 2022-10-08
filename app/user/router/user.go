package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qiaoshurui/couples-subtotal/app/examples/apis"
)

func RegisterUserRouter(v1 *gin.RouterGroup) {
	user := &apis.User{}
	{
		v1.POST("/sign", user.SignUpHandler)
	}

}
