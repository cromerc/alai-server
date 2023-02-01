package routes

import (
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/controllers"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/middlewares"

	"github.com/julienschmidt/httprouter"
)

func FrameRoutes(router *httprouter.Router) {
	router.GET("/frame", controllers.ListFrame)
	router.GET("/frame/:id", controllers.GetFrame)
	router.POST("/frame", middlewares.Authenticate(controllers.CreateFrame))
	router.PATCH("/frame/:id", middlewares.Authenticate(controllers.UpdateFrame))
	router.DELETE("/frame/:id", middlewares.Authenticate(controllers.DeleteFrame))
}
