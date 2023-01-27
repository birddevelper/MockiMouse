package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	cfg "github.com/birddevelper/mockimouse/config"
	"github.com/birddevelper/mockimouse/utils"
	"github.com/gofiber/fiber/v2"
)

type EndpointHandler struct {
	Endpoint *cfg.EndPoint
}

func (endpointHandler *EndpointHandler) handler(c *fiber.Ctx) error {

	scenario := endpointHandler.getMatchSenario(c)
	c.Set("content-type", endpointHandler.Endpoint.ContentType)
	c.Status(scenario.Status)
	response, err := scenario.GetResponse()
	if err != nil {
		fmt.Println("Error :" + err.Error())
	}
	time.Sleep(time.Duration(endpointHandler.Endpoint.Delay) * time.Millisecond)
	return c.SendString(response)

}

func (endpointHandler *EndpointHandler) getMatchSenario(c *fiber.Ctx) cfg.Scenario {

	for _, scenario := range endpointHandler.Endpoint.Scenarios {

		senarioMatch := true
	paramChecking:
		for _, param := range scenario.Condition.Params {
			var parameterValue string
			switch param.Type {
			case "query":
				parameterValue = c.Query(param.Name)

			case "form":
				parameterValue = c.FormValue(param.Name)

			case "header":
				parameterValue = c.Get(param.Name)

			case "path":
				parameterValue = c.Params(param.Name)

			case "body":
				parameterValue, _ = utils.GetParamFromJson(c.Body(), param.Name)

			default:
				senarioMatch = false
				break paramChecking
			}

			switch param.Operand {
			case "equal":
				if senarioMatch = string(parameterValue) == param.Value; !senarioMatch {
					break paramChecking
				}

			case "notEqual":
				if senarioMatch = string(parameterValue) != param.Value; !senarioMatch {
					break paramChecking
				}

			case "greaterThan":
				numericParam, _ := strconv.ParseFloat(parameterValue, 64)
				numericValue, _ := strconv.ParseFloat(param.Value, 64)
				if senarioMatch = numericParam > numericValue; !senarioMatch {
					break paramChecking
				}

			case "lessThan":
				numericParam, _ := strconv.ParseFloat(parameterValue, 64)
				numericValue, _ := strconv.ParseFloat(param.Value, 64)
				if senarioMatch = numericParam < numericValue; !senarioMatch {
					break paramChecking
				}

			case "greaterEqual":
				numericParam, _ := strconv.ParseFloat(parameterValue, 64)
				numericValue, _ := strconv.ParseFloat(param.Value, 64)
				if senarioMatch = numericParam >= numericValue; !senarioMatch {
					break paramChecking
				}

			case "lessEqual":
				numericParam, _ := strconv.ParseFloat(parameterValue, 64)
				numericValue, _ := strconv.ParseFloat(param.Value, 64)
				if senarioMatch = numericParam <= numericValue; !senarioMatch {
					break paramChecking
				}

			case "contain":
				if senarioMatch = strings.Contains(string(parameterValue), param.Value); !senarioMatch {
					break paramChecking
				}

			default:
				senarioMatch = false
				break paramChecking
			}

		}

		if senarioMatch {
			return scenario
		}
	}

	return cfg.Scenario{Response: "None of Secnarios matched with given parameter(s)", Status: 404}
}
