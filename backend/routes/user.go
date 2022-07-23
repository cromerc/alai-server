package routes

import (
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/controllers"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/middlewares"

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
