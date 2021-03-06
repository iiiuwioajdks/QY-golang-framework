package super_admin

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"jiaoshoujia/web/api"
	"jiaoshoujia/web/forms"
	"jiaoshoujia/web/handler/super_admin_handler"
)

func Create(c *gin.Context) {
	var csa forms.CreateSAdmin
	if err := c.ShouldBindJSON(&csa); err != nil {
		zap.S().Info(err)
		api.InvalidParam(c, err)
		return
	}
	err := super_admin_handler.Create(&csa)
	if err != nil {
		zap.S().Info(err)
		api.HandlerErr(c)
		return
	}
	api.Success(c, nil)
}
