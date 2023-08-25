package routes

import (
	"bankDemo/controllers"

	"github.com/gin-gonic/gin"
)

func TransRoute(r *gin.Engine, controller controllers.TransactionController){
	r.POST("/transfer", controller.TransferMoney)
}