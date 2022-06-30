package routes

import (
	"backend/controllers"
	"backend/middlewares"

	"github.com/julienschmidt/httprouter"
)

func UserRoutes(router *httprouter.Router) {
	router.POST("/login", controllers.Login)
	router.GET("/users", middlewares.Authenticate(controllers.ListUsers))
	router.POST("/user", middlewares.Authenticate(controllers.CreateUser))
	router.GET("/user/:id", middlewares.Authenticate(controllers.GetUser))
	router.PATCH("/user/:id", middlewares.Authenticate(controllers.UpdateUser))
	router.GET("/auth", middlewares.Authenticate(controllers.AuthenticateUser))
}
