package main

import (
	"fmt"
	"strconv"

	cfg "github.com/birddevelper/mockimouse/config"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	api := app.Group(cfg.ConfigResolver.GetContextPath())
	endpoints := cfg.ConfigResolver.GetEndPoints()
	endpointHandlers := make([]EndpointHandler, len(endpoints))

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
