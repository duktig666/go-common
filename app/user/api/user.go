package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/qiaoshurui/couples-subtotal/common/api"
)

type User struct{
	api.Api
}

func (u *User) SignUpHandler(c *gin.Context) {
	//获取参数校验
	var s models.
}
