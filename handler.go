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

	requestContentType := c.Get("content-type")
	requestContentTypeOK := false

	// Log the request
	fmt.Println(
		"\r\nRequest :\r\n",
		string(c.Request().Header.Header()),
		"\r\nParameters :\r\n",
		c.AllParams(),
		"\r\nBody :\r\n",
		string(c.Body()),
		"\r\n-----------------------")

	for _, accept := range strings.Split(endpointHandler.Endpoint.Accepts, " ") {
		if accept == strings.ToLower(requestContentType) {
			requestContentTypeOK = true
			break
		}
	}

	if !requestContentTypeOK && endpointHandler.Endpoint.Accepts != "" {
		c.Status(400)
		return c.SendString("400 Bad Request")
	}

	scenario := endpointHandler.getMatchScenario(c)

	// if ContentType is set in senario, set it to response header
	if scenario.ContentType != "" {
		c.Set("content-type", scenario.ContentType)
	}
	// if scenario status is not set, default would be 200
	if scenario.Status == 0 {
		scenario.Status = 200
	}

	c.Status(scenario.Status)
	response, err := scenario.GetResponse()
	if err != nil {
		fmt.Println("Error :" + err.Error())
	}
	time.Sleep(time.Duration(endpointHandler.Endpoint.Delay) * time.Millisecond)
	return c.SendString(response)

}

func (endpointHandler *EndpointHandler) getMatchScenario(c *fiber.Ctx) cfg.Scenario {

	// if there exist only one senario and it has no parametes, return it
	if len(endpointHandler.Endpoint.Scenarios) == 1 &&
		len(endpointHandler.Endpoint.Scenarios[0].Condition.Params) == 0 {
		return endpointHandler.Endpoint.Scenarios[0]
	}

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

	notFoundScenario := cfg.Scenario{Status: 404}
	notFoundScenario.Response = []string{"None of Secnarios matched with given parameter(s)"}

	return notFoundScenario
}
