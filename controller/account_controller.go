package controller

import (
	"github.com/gin-gonic/gin"
	"golangmicro/AccountService"
	"golangmicro/models"
)

type UserController struct {
	Service AccountService.ServiceAccount
}

func NewAccountController(service AccountService.ServiceAccount) *UserController {
	return &UserController{Service: service}
}

func (nac *UserController) CreateBalance(c *gin.Context) {
	var account models.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Request"})
		return
	}

	if err := nac.Service.CreateBalance()
}
