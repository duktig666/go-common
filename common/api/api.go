// description:
// @author renshiwei
// Date: 2022/10/5 21:03

package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DefaultLanguage = "zh-CN"

type Api struct {
	Context *gin.Context
	Gorm    *gorm.DB
	Errors  error
}
