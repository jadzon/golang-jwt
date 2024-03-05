package main

import (
	"jwt_najnowszy/controllers"
	"jwt_najnowszy/initializers"
	"jwt_najnowszy/middleware"
	"jwt_najnowszy/models"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}
func main() {
	userDB := models.CreateEmptyUserDB()
	router := gin.Default()
	router.POST("/signup", func(c *gin.Context) {
		controllers.Signup(c, userDB)
	})
	router.POST("/login", func(c *gin.Context) {
		controllers.Login(c, userDB)
	})
	router.POST("/logout", func(c *gin.Context) {
		middleware.RequireAuth(c, userDB)
	}, controllers.Logout)
	router.GET("/validate", func(c *gin.Context) {
		middleware.RequireAuth(c, userDB)
	}, controllers.Validate)
	router.Run(os.Getenv("PORT_NUM"))
}
