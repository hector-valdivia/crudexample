package routers

import (
	"crudexample/controllers"
	"crudexample/middlewares"
	"github.com/gin-gonic/gin"
)

func setAuthRoute(router *gin.Engine) {
	authController := new(controllers.AuthController)
	//non-auth required
	router.POST("/login", authController.Login)
	router.POST("/signup", authController.Signup)

	authGroup := router.Group("/")
	authGroup.Use(middlewares.Authentication())
	authGroup.GET("/profile", authController.Profile)
}

// InitRoute ..
func InitRoute() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	setAuthRoute(router)

	return router
}
