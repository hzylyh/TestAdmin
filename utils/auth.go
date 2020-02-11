package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"salotto/conf"
	"salotto/model"
	"salotto/utils/redisUtil"
	"time"
)

func GenerateToken(user *model.TSysUser) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.UserAccount,
		//"exp":      time.Now().Add(time.Hour * 2).Unix(),
		"iat": time.Now(),
	})

	mytoken, err := token.SignedString([]byte("secret"))
	redisUtil.SetKeyWithExp(user.UserAccount, mytoken, "1200")
	return mytoken, err

}

func RefreshToken(username, token string) {
	redisUtil.SetKeyWithExp(username, token, "1200")
}

//func TokenMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		tokenStr := r.Header.Get("authorization")
//		if tokenStr == "" {
//			ResponseFailOther(w, conf.UnAuth, "not authorized")
//		} else {
//			token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
//				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//					ResponseFailOther(w, conf.UnAuth, "not authorized")
//					return nil, fmt.Errorf("not authorization")
//				}
//				return []byte("secret"), nil
//			})
//			if !token.Valid {
//				ResponseFailOther(w, conf.UnAuth, "not authorized")
//			} else {
//				claim := token.Claims.(jwt.MapClaims)
//				username := claim["username"].(string)
//				redisToken := redisUtil.GetStringValue(username)
//				if redisToken == tokenStr {
//					RefreshToken(username, tokenStr)
//					next.ServeHTTP(w, r)
//				} else {
//					ResponseFailOther(w, conf.UnAuth, "not authorized")
//				}
//
//			}
//		}
//	})
//}

func TokenAuthMid() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("开始认证")
		//tokenStr := r.Header.Get("authorization")
		a := c.Request.Header
		fmt.Println(a)
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			ResponseFailOther(c, conf.UnAuth, conf.UnAuthMessage)
			fmt.Println("kong")
			c.Abort()
			return // return也是可以省略的，执行了abort操作，会内置在中间件defer前，return，写出来也只是解答为什么Abort()之后，还能执行返回JSON数据
		} else {
			token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					ResponseFailOther(c, conf.UnAuth, conf.UnAuthMessage)
					fmt.Println("else")
					return nil, fmt.Errorf(conf.UnAuthMessage)
				}
				return []byte("secret"), nil
			})
			if !token.Valid {
				ResponseFailOther(c, conf.UnAuth, conf.UnAuthMessage)
				c.Abort()
				return
			} else {
				claim := token.Claims.(jwt.MapClaims)
				username := claim["username"].(string)
				redisToken := redisUtil.GetStringValue(username)
				if redisToken == tokenStr {
					RefreshToken(username, tokenStr)
					fmt.Println("guo")
					c.Next()
				} else {
					ResponseFailOther(c, conf.UnAuth, conf.UnAuthMessage)
					c.Abort()
					return
				}

			}
		}
		fmt.Println("结束认证")
	}
}
