package routes

import (
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/controllers"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/middlewares"

	"github.com/julienschmidt/httprouter"
)

func PlayerRoutes(router *httprouter.Router) {
	router.GET("/os", controllers.ListPlayer)
	router.GET("/os/:id", controllers.GetPlayer)
	router.POST("/os", middlewares.Authenticate(controllers.CreatePlayer))
	router.PATCH("/os/:id", middlewares.Authenticate(controllers.UpdatePlayer))
	router.DELETE("/os/:id", middlewares.Authenticate(controllers.DeletePlayer))
}
