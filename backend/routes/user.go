package routes

import (
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/controllers"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/middlewares"

	"github.com/julienschmidt/httprouter"
)

func UserRoutes(router *httprouter.Router) {
	router.POST("/login", controllers.Login)
	router.GET("/auth", middlewares.Authenticate(controllers.AuthenticateUser))

	router.GET("/user", middlewares.Authenticate(controllers.ListUser))
	router.GET("/user/:id", middlewares.Authenticate(controllers.GetUser))
	router.POST("/user", middlewares.Authenticate(controllers.CreateUser))
	router.PATCH("/user/:id", middlewares.Authenticate(controllers.UpdateUser))
	router.DELETE("/user/:id", middlewares.Authenticate(controllers.DeleteUser))
}
