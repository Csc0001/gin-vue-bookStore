package middleware

import (
	"gin-vue-bookStore/common"
	"gin-vue-bookStore/model"
	"gin-vue-bookStore/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware()gin.HandlerFunc  {
	return func(ctx *gin.Context){
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString,"Bearer "){
			response.Response(ctx,http.StatusUnauthorized,401,nil,"错误的token格式")
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		token,claims,err := common.ParseToken(tokenString)
		if err != nil||!token.Valid{ // token失效或出错
			response.Response(ctx,http.StatusUnauthorized,401,nil,"token失效")
			ctx.Abort()
			return
		}
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user,userId)
		if user.ID == 0{
			response.Response(ctx,http.StatusUnauthorized,401,nil,"查无此人")
			ctx.Abort()
			return
		}
		//用户存在，写入用户信息到上下文
		ctx.Set("user",user)
		ctx.Next()
	}
}