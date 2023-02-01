package routes

import (
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/controllers"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/middlewares"

	"github.com/julienschmidt/httprouter"
)

func PlayerRoutes(router *httprouter.Router) {
	router.GET("/player", controllers.ListPlayer)
	router.GET("/player/:id", controllers.GetPlayer)
	router.POST("/player", middlewares.Authenticate(controllers.CreatePlayer))
	router.PATCH("/player/:id", middlewares.Authenticate(controllers.UpdatePlayer))
	router.DELETE("/player/:id", middlewares.Authenticate(controllers.DeletePlayer))
}
