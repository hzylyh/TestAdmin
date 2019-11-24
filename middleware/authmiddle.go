package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthMiddle() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println()
	}
}
