package routes

import (
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/controllers"

	"github.com/julienschmidt/httprouter"
)

func GameRoutes(router *httprouter.Router) {
	router.GET("/game", controllers.ListGame)
	router.GET("/game/:id", controllers.GetGame)
	router.POST("/game", controllers.CreateGame)
	router.PATCH("/game/:id", controllers.UpdateGame)
	router.DELETE("/game/:id", controllers.DeleteGame)
}
