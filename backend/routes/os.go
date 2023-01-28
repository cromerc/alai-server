package routes

import (
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/controllers"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/middlewares"

	"github.com/julienschmidt/httprouter"
)

func OSRoutes(router *httprouter.Router) {
	router.GET("/os", controllers.ListOS)
	router.GET("/os/:id", controllers.GetOS)
	router.POST("/os", middlewares.Authenticate(controllers.CreateOS))
	router.PATCH("/os/:id", middlewares.Authenticate(controllers.UpdateOS))
	router.DELETE("/os/:id", middlewares.Authenticate(controllers.DeleteOS))
}
