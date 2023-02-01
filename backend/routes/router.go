package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"git.cromer.cl/Proyecto-Titulo/alai-server/backend/middlewares"
	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"
)

func Initialize() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", index)
	FrameRoutes(router)
	GameRoutes(router)
	GodotVersionRoutes(router)
	LevelRoutes(router)
	OSRoutes(router)
	PlayerRoutes(router)
	UserRoutes(router)
	return router
}

func Serve(router *httprouter.Router) {
	newRouter := handlers.CombinedLoggingHandler(os.Stdout, router)
	newRouter = handlers.CompressHandler(newRouter)
	newRouter = middlewares.Cors(newRouter)

	idleConnsClosed := make(chan struct{})

	server := &http.Server{Addr: ":3001", Handler: newRouter}

	// Listen for CTRL-C(SIGTERM)
	sigterm := make(chan os.Signal)
	signal.Notify(sigterm, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigterm
		// When CTRL-C is pressed shutdown the server
		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
		os.Exit(0)
	}()

	// Run the server
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}

func index(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprintf(writer, "This is the Alai API server!")
}
