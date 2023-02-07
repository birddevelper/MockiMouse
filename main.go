package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2/middleware/cors"

	cfg "github.com/birddevelper/mockimouse/config"
	"github.com/birddevelper/mockimouse/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	api := app.Group(cfg.ConfigResolver.GetContextPath())

	// static files such as image or css files
	app.Static(cfg.ConfigResolver.GetStatics(), "./assets")
	app.Use(cors.New())
	endpoints := cfg.ConfigResolver.GetEndPoints()

	// print endpoint information
	utils.PrintEndpointsInfo(endpoints, cfg.ConfigResolver.GetPort())

	// initiate endpointHandlers
	endpointHandlers := make([]EndpointHandler, len(endpoints))

	// creating endpoints
	for i, endpoint := range endpoints {
		endpointHandlers[i] = EndpointHandler{Endpoint: &endpoints[i]}
		switch endpointHandlers[i].Endpoint.Method {
		case "GET":
			api.Get(endpoint.Path, endpointHandlers[i].handler)
		case "POST":
			api.Post(endpoint.Path, endpointHandlers[i].handler)
		case "PUT":
			api.Put(endpoint.Path, endpointHandlers[i].handler)
		case "PATCH":
			api.Patch(endpoint.Path, endpointHandlers[i].handler)
		case "DELETE":
			api.Delete(endpoint.Path, endpointHandlers[i].handler)
		default:
			fmt.Println(endpointHandlers[i].Endpoint.Name, " Endpoint method is not valid")
		}

	}

	app.Listen(":" + strconv.Itoa(cfg.ConfigResolver.GetPort()))
}
