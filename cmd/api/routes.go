package main

import (
	"github.com/Fabimhso/go-categories-msvc/cmd/api/controllers"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes (router *gin.Engine) {
	CategoryRoutes := router.Group("/categories")

	CategoryRoutes.POST("/", controllers.CreateCategory)
}