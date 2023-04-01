package controller

import (

	"gin-vue-bookStore/model"
	"gin-vue-bookStore/repository"
	"gin-vue-bookStore/response"
	"gin-vue-bookStore/vo"
	"github.com/gin-gonic/gin"

	"strconv"
)

type ICategoryController interface {
	RestController
}
type CategoryController struct {
	Repository repository.CategoryRepository
}

func NewCategoryController()ICategoryController{
	repository := repository.NewCategoryRepository()

	repository.DB.AutoMigrate(model.Category{})
	return CategoryController{Repository: repository}
}
func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory);err !=nil{
		response.Fail(ctx,nil,"数据验证错误，分类名称缺省")
		return
	}
	category,err := c.Repository.Create(requestCategory.Name)
	if err!=nil{
		panic(err)
		return
	}
	response.Success(ctx,gin.H{"category":category},"创建成功")

}

func (c CategoryController) Update(ctx *gin.Context) {
	// 绑定body中的参数
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory);err !=nil{
		response.Fail(ctx,nil,"数据验证错误，分类名称缺省")
		return
	}
	// 获取path中的参数
	categoryId,_ := strconv.Atoi(ctx.Params.ByName("id"))

	updateCategory,err := c.Repository.SelectById(categoryId)
	if err!=nil{
		response.Fail(ctx,nil,"分类不存在")
	}
	//更新
	if category,err := c.Repository.Update(*updateCategory,requestCategory.Name);err!=nil{
		response.Fail(ctx,nil,"更新失败")
	}else {
		response.Success(ctx, gin.H{"category": category}, "修改成功")
	}
}

func (c CategoryController) Show(ctx *gin.Context) {
	// 获取path中的参数
	categoryId,_ := strconv.Atoi(ctx.Params.ByName("id"))

	if findCategory,err := c.Repository.SelectById(categoryId);err!=nil{
		response.Fail(ctx,nil,"分类不存在")
	}else{
		response.Success(ctx,gin.H{"category":findCategory},"成功查询")
	}

}

func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取path中的参数
	categoryId,_ := strconv.Atoi(ctx.Params.ByName("id"))

	if err := c.Repository.DeleteById(categoryId);err !=nil{
		response.Fail(ctx,nil,"删除失败")
	}else{
		response.Success(ctx,nil,"成功删除")
	}
}





