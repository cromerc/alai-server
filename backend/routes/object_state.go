package routes

import (
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/controllers"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/middlewares"

	"github.com/julienschmidt/httprouter"
)

func ObjectStateRoutes(router *httprouter.Router) {
	router.GET("/object-state", controllers.ListObjectState)
	router.GET("/object-state/:id", controllers.GetObjectState)
	router.POST("/object-state", middlewares.Authenticate(controllers.CreateObjectState))
	router.PATCH("/object-state/:id", middlewares.Authenticate(controllers.UpdateObjectState))
	router.DELETE("/object-state/:id", middlewares.Authenticate(controllers.DeleteObjectState))
}
