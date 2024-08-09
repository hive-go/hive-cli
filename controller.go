package main

var ControllerString = `package #MODULE_NAME

import (
	"github.com/hive-go/hive"
  "github.com/gofiber/fiber/v2"
)

var #CAPITALIZED_NAMEController = hive.CreateController()

func init() {

	DomainController.SetConfig(hive.ControllerConfig{
		Prefix: "/#MODULE_NAME",
	})

	#CAPITALIZED_NAMEController.Post("", func(c *fiber.Ctx) (any, error) {

		payload := new(Create#CAPITALIZED_NAMEDto)

		if err := c.BodyParser(&payload); err != nil {
			return nil, fiber.NewError(fiber.StatusNotFound, "BODY_PARSING_ERROR")
		}

		result, err := #CAPITALIZED_NAMEService.Create#CAPITALIZED_NAME(payload)

		if err != nil {
			return nil, err
		}

		return result, nil
	})

	#CAPITALIZED_NAMEController.Get("", func(c *hive.Ctx) (any, error) {
		result, err := #CAPITALIZED_NAMEService.Get#CAPITALIZED_NAMEs()

		if err != nil {
			return nil, err
		}

		return result, nil
	})

  #CAPITALIZED_NAMEController.Get("/:id", func(c *fiber.Ctx) (any, error) {
		id := c.Params("id")

		body, err := #CAPITALIZED_NAMEService.Get#CAPITALIZED_NAMEById(id)

		if err != nil {
			return nil, err
		}

		return body, nil
	})

  #CAPITALIZED_NAMEController.Patch("/:id", func(c *fiber.Ctx) (any, error) {
    id := c.Params("id")

    payload := new(Update#CAPITALIZED_NAMEDto)

    if err := c.BodyParser(&payload); err != nil {
      return nil, fiber.NewError(fiber.StatusNotFound, "BODY_PARSING_ERROR")
    }

    body, err := #CAPITALIZED_NAMEService.Update#CAPITALIZED_NAME(id, payload)

    if err != nil {
      return nil, err
    }

    return body, nil
  })

	#CAPITALIZED_NAMEController.Delete("/:id", func(c *fiber.Ctx) (any, error) {
		id := c.Params("id")

		body, err := #CAPITALIZED_NAMEService.Delete#CAPITALIZED_NAME(id)

		if err != nil {
			return nil, err
		}

		return body, nil
	})
}
`