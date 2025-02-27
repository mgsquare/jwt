package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/mgsquare/jwt/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.SignUp())
	incomingRoutes.POST("users/login", controller.Login())
}
