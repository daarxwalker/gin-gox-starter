package facade

import (
	"github.com/gin-gonic/gin"
	
	"common/pkg/service/gox_service"
)

func Gox(c *gin.Context) *gox_service.GoxService {
	return gox_service.New(c)
}
