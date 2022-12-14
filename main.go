package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/elias-gill/go_pokemon/authentication"
	"github.com/elias-gill/go_pokemon/routers"
	"github.com/elias-gill/go_pokemon/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	port := ":3000"
	// instantiate chi router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(authentication.JwtMidleware)
	r.Route("/user", routers.UserHandlers)
	r.Route("/teams", routers.TeamsHandlers)

	// start server
	println("Starting server in port" + port)
	go log.Fatal(http.ListenAndServe(port, r))

	// Wait for an in interrupt. Attempt a graceful shutdown.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
    // cerrar conexion con mongo cuando se cierre el programa
    server.C.CloseMongo()
	<-c
}
