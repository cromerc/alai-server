package routes

import (
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/controllers"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/middlewares"

	"github.com/julienschmidt/httprouter"
)

func ObjectNameRoutes(router *httprouter.Router) {
	router.GET("/object-name", controllers.ListObjectName)
	router.GET("/object-name/:id", controllers.GetObjectName)
	router.POST("/object-name", middlewares.Authenticate(controllers.CreateObjectName))
	router.PATCH("/object-name/:id", middlewares.Authenticate(controllers.UpdateObjectName))
	router.DELETE("/object-name/:id", middlewares.Authenticate(controllers.DeleteObjectName))
}
