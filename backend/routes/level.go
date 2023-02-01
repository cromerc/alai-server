package routes

import (
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/controllers"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/middlewares"

	"github.com/julienschmidt/httprouter"
)

func LevelRoutes(router *httprouter.Router) {
	router.GET("/level", controllers.ListLevel)
	router.GET("/level/:id", controllers.GetLevel)
	router.POST("/level", middlewares.Authenticate(controllers.CreateLevel))
	router.PATCH("/level/:id", middlewares.Authenticate(controllers.UpdateLevel))
	router.DELETE("/level/:id", middlewares.Authenticate(controllers.DeleteLevel))
}
