package controllers

import (
	"bankDemo/interfaces"
	"bankDemo/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerController struct{
    CustomerService  interfaces.Icustomer
}


func InitCustController(customerService interfaces.Icustomer) CustomerController {
    return CustomerController{customerService}
}

func (t *CustomerController)CreateCustomer(ctx *gin.Context){
    var trans *models.Customer  
    if err := ctx.ShouldBindJSON(&trans); err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }
    newtrans, err := t.CustomerService.CreateCustomer(trans)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newtrans})

}



func (t *CustomerController)GetCustomerById(ctx *gin.Context){
    id:= ctx.Param("id")
    id1,err := strconv.ParseInt(id,10,64)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    val, err := t.CustomerService.GetCustomerById(id1)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": val})
}

func (t *CustomerController)UpdateCustomerById(ctx *gin.Context){
    id:= ctx.Param("id")
    fv := &models.UpdateModel{}
    if err := ctx.ShouldBindJSON(&fv); err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }
    fmt.Println(fv)
    id1,err := strconv.ParseInt(id,10,64)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    res,err := t.CustomerService.UpdateCustomerById(id1,fv)
    if err!=nil{
        fmt.Println("error")
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}

func (t *CustomerController)DeleteCustomerById(ctx *gin.Context){
    id:= ctx.Param("id")
    id1,err := strconv.ParseInt(id,10,64)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    res,err := t.CustomerService.DeleteCustomerById(id1)
    if err!=nil{
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}

func (t *CustomerController)GetAllCustomerTransaction(ctx *gin.Context){
    id:= ctx.Param("id")
    id1,err := strconv.ParseInt(id,10,64)
    if(err!=nil){
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    res,err := t.CustomerService.GetAllCustomerTransaction(id1)                                                                
    if err!=nil{
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
    }
    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}

func (t *CustomerController)GetAllTransactionSum(ctx *gin.Context){
    id := ctx.Param("id")
    id1,_ := strconv.ParseInt(id,10,64)
    var date *Date
	if err := ctx.ShouldBindJSON(&date); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
    res,err := t.CustomerService.GetAllTransactionSum(id1, date.From, date.To)
    if err!=nil{
        ctx.JSON(http.StatusBadGateway, gin.H{"status":"fail","message":err.Error()})
    }
    ctx.JSON(http.StatusAccepted, gin.H{"status":"success", "sum":res})
}