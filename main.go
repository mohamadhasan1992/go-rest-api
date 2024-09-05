package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mohamadhasan1992/go-rest-api.git/db"
	"github.com/mohamadhasan1992/go-rest-api.git/routes"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func router01() http.Handler {
	server := gin.Default()
	routes.RegisterAuthRoutes(server)
	return server
}

func router02() http.Handler {
	server := gin.Default()
	routes.RegisterEventRoutes(server)
	return server
}

func router03() http.Handler {
	server := gin.Default()
	routes.RegisterTestRoutes(server)
	return server
}

func main() {
	db.InitDB()
	server01 := &http.Server{
		Addr:         ":8081",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":8082",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server03 := &http.Server{
		Addr:         ":8083",
		Handler:      router03(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	g.Go(func() error {
		return server03.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

	// gin.ForceConsoleColor()
	// server := gin.Default()
	// routes.RegisterRoutes(server)
	// server.Run(":8080")
}
