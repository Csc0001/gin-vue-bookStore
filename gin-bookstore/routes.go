package main

import (
	"gin-vue-bookStore/controller"
	"gin-vue-bookStore/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine)*gin.Engine{
	r.Use(middleware.CORSMiddleware(),middleware.RecoveryMiddleware())
	r.POST("/api/user/register",controller.Register)
	r.POST("/api/user/login",controller.Login)
	r.GET("/api/user/info",middleware.AuthMiddleware(),controller.Info)

	categoryRoutes := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("",categoryController.Create)
	categoryRoutes.PUT("/:id",categoryController.Update)
	categoryRoutes.GET("/:id",categoryController.Show)
	categoryRoutes.DELETE(":id",categoryController.Delete)

	postRoutes := r.Group("/posts")
	postController := controller.NewPostController()
	postRoutes.POST("",middleware.AuthMiddleware(),postController.Create)
	postRoutes.PUT("/:id",middleware.AuthMiddleware(),postController.Update)
	postRoutes.GET("/:id",middleware.AuthMiddleware(),postController.Show)
	postRoutes.DELETE(":id",middleware.AuthMiddleware(),postController.Delete)
	postRoutes.POST("page/list",postController.PageList)
	return r
}