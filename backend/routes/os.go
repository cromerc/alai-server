package routes

import (
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/controllers"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/middlewares"

	"github.com/julienschmidt/httprouter"
)

func OSRoutes(router *httprouter.Router) {
	router.GET("/os", controllers.ListUsers)
	router.GET("/os/:id", controllers.GetUser)
	router.POST("/os", middlewares.Authenticate(controllers.CreateUser))
	router.PATCH("/os/:id", middlewares.Authenticate(controllers.UpdateUser))
	router.DELETE("/os/:id", middlewares.Authenticate(controllers.DeleteUser))
}
