package controller

import (
	"gin-vue-bookStore/common"
	"gin-vue-bookStore/model"
	"gin-vue-bookStore/response"
	"gin-vue-bookStore/vo"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
)

type IPostController interface {
	RestController
	PageList(ctx *gin.Context)
}
type PostController struct {
	DB *gorm.DB
}



func (p PostController) Create(ctx *gin.Context) {
	var requestPost vo.CreaterPostRequest
	if err := ctx.ShouldBind(&requestPost);err !=nil{
		response.Fail(ctx,nil,"数据验证错误，分类名称缺省")
		return
	}
	user,_ := ctx.Get("user")
	log.Println(user)
	post := model.Post{
		UserId: user.(model.User).ID,
		CategoryId: requestPost.CategoryId,
		Title: requestPost.Title,
		HeadImg: requestPost.HeadImg,
		Content: requestPost.Content,
	}

	if err := p.DB.Create(&post).Error;err!=nil{
		panic(err)
		return
	}
	response.Success(ctx,nil,"创建成功")
}


func (p PostController) Update(ctx *gin.Context) {
	var requestPost vo.CreaterPostRequest
	if err := ctx.ShouldBind(&requestPost);err !=nil{
		response.Fail(ctx,nil,"数据验证错误，分类名称缺省")
		return
	}
	user,_ := ctx.Get("user")

	postId := ctx.Params.ByName("id")

	var post model.Post
	if p.DB.Where("id=?",postId).First(&post).RecordNotFound(){
		response.Fail(ctx,nil,"文章不存在")
		return
	}
	userId := user.(model.User).ID
	if userId != post.UserId{
		response.Fail(ctx,nil,"文章不属于您，勿非法操作")
		return
	}
	if err := p.DB.Model(&post).Update(requestPost).Error;err != nil{
		response.Fail(ctx,nil,"更新失败")
		return
	}
	response.Success(ctx,gin.H{"post":post},"更新成功")
}

func (p PostController) Show(ctx *gin.Context) {
	postId := ctx.Params.ByName("id")

	var post model.Post
	if p.DB.Preload("Category").Where("id=?",postId).First(&post).RecordNotFound(){
		response.Fail(ctx,nil,"文章不存在")
		return
	}
	response.Success(ctx,gin.H{"post":post},"")
}

func (p PostController) Delete(ctx *gin.Context) {
	postId := ctx.Params.ByName("id")
	var post model.Post
	if p.DB.Where("id=?",postId).First(&post).RecordNotFound(){
		response.Fail(ctx,nil,"文章不存在")
		return
	}
	user,_ := ctx.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId{
		response.Fail(ctx,nil,"文章不属于您，勿非法操作")
		return
	}
	p.DB.Delete(&post)
	response.Success(ctx,nil,"删除成功")
}

func NewPostController() IPostController {
	db := common.GetDB()
	db.AutoMigrate(model.Post{})
	return PostController{DB: db}
}

// 分页
func (p PostController) PageList(ctx *gin.Context) {
	//获取分页参数
	pageNum,_ := strconv.Atoi( ctx.DefaultQuery("pageNum","1"))
	pageSize,_ := strconv.Atoi( ctx.DefaultQuery("pageSize","20"))
	//分页
	var posts []model.Post
	p.DB.Preload("Category").Order("created_at desc").Offset((pageNum-1)*pageSize).Limit(pageSize).Find(&posts)
	// 传递给前端的总数
	var total int
	p.DB.Model(model.Post{}).Count(&total)
	response.Success(ctx,gin.H{"data":posts,"count":total},"成功")
}