package rest

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

type server struct {
	App *fiber.App
}

func NewHttpServer() *server {
	//set concurrency
	concurrency := viper.GetInt("rest.maxConcurrency")
	if concurrency == 0 {
		concurrency = 10000
	}

	//server configuration setting
	cfg := fiber.Config{
		Prefork:       false,
		ServerHeader:  "sbit",
		StrictRouting: true,
		CaseSensitive: false,
		Immutable:     false,
		UnescapePath:  true,
		Concurrency:   concurrency,
	}

	//crete new instance server
	app := fiber.New(cfg)

	//use default CORST setting
	app.Use(cors.New(cors.Config{
		AllowMethods: "GET, POST",
		AllowOrigins: "*",
	}))

	return &server{
		App: app,
	}
}

// Run start the HTTP server and performs a graceful shutdown
func (s *server) Run() {
	port := viper.GetString("rest.port")
	if port == "" {
		port = "8008"
	}

	//liten from different go routine
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		_ = s.App.Shutdown()
	}()

	err := s.App.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}
