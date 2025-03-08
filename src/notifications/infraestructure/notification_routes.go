package infraestructure
import (
	"mi-notificacion/src/notifications/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)


func NotificationRoutes(router *gin.Engine) {

	router.POST("/notifications", controllers.NotificationController)
}