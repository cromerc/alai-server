package routes

import (
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/controllers"
	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/middlewares"

	"github.com/julienschmidt/httprouter"
)

func AuthRoutes(router *httprouter.Router) {
	router.POST("/login", controllers.Login)
	router.GET("/auth", middlewares.Authenticate(controllers.AuthenticateUser))
}
