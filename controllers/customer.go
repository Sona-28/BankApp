package controllers

import (
	"bankDemo/interfaces"
	"bankDemo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct{
     TransactionService  interfaces.Icustomer
}

// func InitialiseTransactionController( transaction interfaces.Itransaction)(TransactionController){
//  return TransactionController{transactionService} 
// }

func InitTransController(transactionService interfaces.Icustomer) TransactionController {
    return TransactionController{transactionService}
}

func (t *TransactionController)CreateTransaction(ctx *gin.Context){
    var trans *models.Customer  
    if err := ctx.ShouldBindJSON(&trans); err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }
    newtrans, err := t.TransactionService.CreateCustomer(trans)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newtrans})

}

func (t *TransactionController)GetTransaction(ctx *gin.Context){
    val, err := t.TransactionService.GetCustomer()
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": val})

}