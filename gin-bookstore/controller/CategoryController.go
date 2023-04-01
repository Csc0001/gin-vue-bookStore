package controller

import (
	"gin-vue-bookStore/common"
	"gin-vue-bookStore/model"
	"gin-vue-bookStore/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

type ICategoryController interface {
	RestController
}
type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController()ICategoryController{
	db := common.GetDB()
	db.AutoMigrate(model.Category{})
	return CategoryController{DB:db}
}
func (c CategoryController) Creat(ctx *gin.Context) {
	var requestCategory model.Category
	ctx.Bind(&requestCategory)
	if requestCategory.Name == ""{
		response.Fail(ctx,nil,"数据验证错误，分类名称必填")
		return
	}
	c.DB.Create(&requestCategory)
	response.Success(ctx,gin.H{"category":requestCategory},"创建成功")

}

func (c CategoryController) Update(ctx *gin.Context) {
	// 绑定body中的参数
	var requestCategory model.Category
	ctx.Bind(&requestCategory)
	if requestCategory.Name == ""{
		response.Fail(ctx,nil,"数据验证错误，分类名称必填")
		return
	}
	// 获取path中的参数
	categoryId,_ := strconv.Atoi(ctx.Params.ByName("id"))

	var updateCategory model.Category
	if c.DB.First(&updateCategory,categoryId).RecordNotFound(){
		response.Fail(ctx,nil,"此分类不存在")
		return
	}

	//更新
	c.DB.Model(&updateCategory).Update("name",requestCategory.Name)
	response.Success(ctx,nil,"修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	// 获取path中的参数
	categoryId,_ := strconv.Atoi(ctx.Params.ByName("id"))

	var findCategory model.Category
	if c.DB.First(&findCategory,categoryId).RecordNotFound(){
		response.Fail(ctx,nil,"此分类不存在")
		return
	}
	response.Success(ctx,gin.H{"category":findCategory},"成功查询")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取path中的参数
	categoryId,_ := strconv.Atoi(ctx.Params.ByName("id"))

	if err := c.DB.Delete(model.Category{},categoryId).Error;err !=nil{
		response.Fail(ctx,nil,"删除失败")
		return
	}
	response.Success(ctx,nil,"成功删除")
}





