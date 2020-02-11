package service

import (
	"github.com/gin-gonic/gin"
	"salotto/model"
)

var Auth = &authService{}

type authService struct {
}

func (srv *authService) VerifyUser(c *gin.Context) (string, error) {
	var (
		user model.TSysUser
		err  error
	)
	if err = c.ShouldBind(&user); err != nil {
		return "", err
	}
	return "token", err
}
