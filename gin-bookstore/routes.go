package main

import (
	"gin-vue-bookStore/controller"
	"gin-vue-bookStore/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine)*gin.Engine{
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/user/register",controller.Register)
	r.POST("/api/user/login",controller.Login)
	r.GET("/api/user/info",middleware.AuthMiddleware(),controller.Info)

	categoryRoutes := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("",categoryController.Creat)
	categoryRoutes.PUT("/:id",categoryController.Update)
	categoryRoutes.GET("/:id",categoryController.Show)
	categoryRoutes.DELETE(":id",categoryController.Delete)
	return r
}