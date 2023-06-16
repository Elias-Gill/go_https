package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/elias-gill/go_pokemon/authentication"
	_ "github.com/elias-gill/go_pokemon/docs"
	"github.com/elias-gill/go_pokemon/routers"
	"github.com/elias-gill/go_pokemon/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title			go_https ft: pokeapi
// @version		1.0
// @description	pokeapi "wrapper" made using golang
// @BasePath		/
func main() {
	// instantiate chi router
	port := ":3000"
	r := chi.NewRouter()

	// middlewares
	r.Use(cors.AllowAll().Handler)
	r.Use(middleware.Logger)
	r.Use(authentication.JwtMidleware)
	r.Mount("/swagger", httpSwagger.WrapHandler)

	// routes
	r.Route("/user", routers.UserHandlers)
	r.Route("/teams", routers.TeamsHandlers)

	// start server
	println("Starting server in port" + port)
	go http.ListenAndServe(port, r)

	// Wait for an in interrupt and attempt a graceful shutdown.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// cerrar conexion con mongo cuando se cierre el programa
	if a := <-c; a != nil {
		server.C.CloseMongo()
		return
	}
}
