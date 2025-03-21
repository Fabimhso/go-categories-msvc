package main

import (
	"github.com/Fabimhso/go-categories-msvc/cmd/api/controllers"
	"github.com/Fabimhso/go-categories-msvc/internal/repositories"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes (router *gin.Engine) {
	CategoryRoutes := router.Group("/categories")

	inMemoryCategoryRepository := repositories.NewInMemoryCategoryRepository()
	
	CategoryRoutes.POST("/", func (ctx *gin.Context) {
		controllers.CreateCategory(ctx, inMemoryCategoryRepository)
	})

	CategoryRoutes.GET("/", func (ctx *gin.Context) {
		controllers.ListCategories(ctx, inMemoryCategoryRepository)
	})
}