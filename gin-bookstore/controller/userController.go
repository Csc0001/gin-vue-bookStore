package controller

import (
	"gin-vue-bookStore/common"
	"gin-vue-bookStore/dto"
	"gin-vue-bookStore/model"
	"gin-vue-bookStore/response"
	"gin-vue-bookStore/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register(c *gin.Context){
	DB := common.GetDB()
	name := c.PostForm("name")
	tel := c.PostForm("tel")
	password := c.PostForm("password")
	//手机号长度
	if len(tel) != 11{
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"手机号长度必须为11位")
		return
	}
	//密码不少于6位
	if len(password) < 6{
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"密码长度不少于6位")
	}
	//姓名缺省
	if len(name) == 0{
		name = util.RandomString(10)
	}
	//手机号已存在不允许注册
	if isTelephoneExit(DB,tel){
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"用户已存在")

	}
	//创建用户
	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil{
		response.Response(c,http.StatusInternalServerError,500,nil,"密码加密错误")
		return
	}
	newUser :=  model.User{
		Name: name,
		Tel: tel,
		Password: string(hashedPassword),
	}
	DB.Create(&newUser)
	log.Println(name,tel,password)
	response.Success(c,nil,"注册成功")
}



func Login(c *gin.Context){
	DB := common.GetDB()
	tel := c.PostForm("tel")
	password := c.PostForm("password")
	//手机号长度
	if len(tel) != 11{
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"手机号长度必须为11位")
		return
	}
	//密码不少于6位
	if len(password) < 6{
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"密码长度不少于6位")
		return
	}
	//判断手机号是否存在
	var user model.User
	DB.Where("tel=?",tel).First(&user)
	if user.ID == 0{
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"用户不存在")
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));err != nil{
		response.Response(c,http.StatusBadRequest,400,nil,"密码错误")
		return
	}
	//发放token
	token,err := common.ReleaseTocken(user)
	if err != nil {
		response.Response(c,http.StatusInternalServerError,500,nil,"系统异常")
		log.Printf("token gennerate error: %v",err)
		return
	}
	response.Success(c,gin.H{"token":token},"登录成功")
	return
}
func Info(ctx *gin.Context){
	user,_ := ctx.Get("user")//对应中间件的Set，从上下文获取用户信息
	response.Success(ctx,gin.H{"user":dto.ToUserDto(user.(model.User))},"获取成功")
}

func isTelephoneExit(db *gorm.DB,tel string)bool{
	var user model.User
	db.Where("tel=?",tel).First(&user)
	if user.ID != 0{
		return true
	}
	return false
}
