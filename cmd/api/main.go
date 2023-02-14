package main

import (
	"context"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/api/handlers"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/api/rpc"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/user"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/constants"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/errno"
	JWT "github.com/dzc1997/DouyinSimplifyEdition/pkg/jwt"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/tracer"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Init() {
	tracer.InitJaeger(constants.ApiServiceName)
	rpc.InitRPC()
}

func main() {
	Init()
	r := gin.New()
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentiryKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVar handlers.UserLogin
			if err := c.ShouldBind(&loginVar); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if len(loginVar.Username) == 0 || len(loginVar.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			return rpc.UserLogin(context.Background(), &user.UserLoginRequest{Username: loginVar.Username, Password: loginVar.Password})
		},
		LoginResponse: func(c *gin.Context, code int, message string, time time.Time) {
			Jwt := JWT.NewJWT([]byte(constants.SecretKey))
			userId, err := Jwt.CheckToken(message)
			if err != nil {
				panic(err)
			}
			handlers.SendUserResponse(c, errno.Success, userId, message)
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		klog.Fatal("JWT Error:" + err.Error())
	}

	douyin := r.Group("/douyin")
	user_ := douyin.Group("/user")
	user_.GET("/", handlers.UserInfo)
	user_.POST("/register/", handlers.UserRegister)
	user_.POST("/login/", authMiddleware.LoginHandler)

	if err := http.ListenAndServe(constants.ApiAddress, r); err != nil {
		klog.Fatal(err)
	}
}
