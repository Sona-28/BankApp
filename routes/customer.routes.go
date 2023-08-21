package routes

import (
	"bankDemo/controllers"

	"github.com/gin-gonic/gin"
)

func CustRoute(router *gin.Engine, controller controllers.TransactionController) {
	router.POST("/api/profile/create", controller.CreateTransaction)

}