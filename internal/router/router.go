package router

import (
	"phonebook/internal/handler"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter builds a gin.Engine with routes + swagger endpoint.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API Routes
	r.GET("/contacts", handler.GetContactsHandler)
	return r
}
