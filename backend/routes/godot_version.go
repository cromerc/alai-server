package routes

import (
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/controllers"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/middlewares"

	"github.com/julienschmidt/httprouter"
)

func GodotVersionRoutes(router *httprouter.Router) {
	router.GET("/godot-version", controllers.ListGodotVersion)
	router.GET("/godot-version/:id", controllers.GetGodotVersion)
	router.POST("/godot-version", middlewares.Authenticate(controllers.CreateGodotVersion))
	router.PATCH("/godot-version/:id", middlewares.Authenticate(controllers.UpdateGodotVersion))
	router.DELETE("/godot-version/:id", middlewares.Authenticate(controllers.DeleteGodotVersion))
}
