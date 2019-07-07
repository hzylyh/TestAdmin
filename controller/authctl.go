package controller

import (
	"github.com/gin-gonic/gin"
	"salotto/service"
	"salotto/utils"
)

func Login(c *gin.Context) {
	var (
		token string
		err   error
	)
	if token, err = service.Auth.VerifyUser(c); err != nil {
		utils.ResponseOk(c, err.Error())
	} else {
		utils.ResponseOk(c, token)
	}
}
