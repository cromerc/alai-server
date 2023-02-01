package routes

import (
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/controllers"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/middlewares"

	"github.com/julienschmidt/httprouter"
)

func ObjectRoutes(router *httprouter.Router) {
	router.GET("/object", controllers.ListObject)
	router.GET("/object/:id", controllers.GetObject)
	router.POST("/object", middlewares.Authenticate(controllers.CreateObject))
	router.PATCH("/object/:id", middlewares.Authenticate(controllers.UpdateObject))
	router.DELETE("/object/:id", middlewares.Authenticate(controllers.DeleteObject))
}
