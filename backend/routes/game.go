package routes

import (
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/controllers"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/middlewares"

	"github.com/julienschmidt/httprouter"
)

func GameRoutes(router *httprouter.Router) {
	router.GET("/game", middlewares.Authenticate(controllers.ListGames))
	router.GET("/game/:id", middlewares.Authenticate(controllers.GetGame))
	router.POST("/game", controllers.CreateGame)
}
