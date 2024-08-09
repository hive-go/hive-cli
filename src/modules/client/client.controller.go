package client

import (
	"github.com/hive-go/hive"
  "github.com/gofiber/fiber/v2"
)

var ClientController = hive.CreateController()

func init() {

	DomainController.SetConfig(hive.ControllerConfig{
		Prefix: "/client",
	})

	ClientController.Post("", func(c *fiber.Ctx) (any, error) {

		payload := new(CreateClientDto)

		if err := c.BodyParser(&payload); err != nil {
			return nil, fiber.NewError(fiber.StatusNotFound, "BODY_PARSING_ERROR")
		}

		result, err := ClientService.CreateClient(payload)

		if err != nil {
			return nil, err
		}

		return result, nil
	})

	ClientController.Get("", func(c *hive.Ctx) (any, error) {
		result, err := ClientService.GetClients()

		if err != nil {
			return nil, err
		}

		return result, nil
	})

  ClientController.Get("/:id", func(c *fiber.Ctx) (any, error) {
		id := c.Params("id")

		body, err := ClientService.GetClientById(id)

		if err != nil {
			return nil, err
		}

		return body, nil
	})

  ClientController.Patch("/:id", func(c *fiber.Ctx) (any, error) {
    id := c.Params("id")

    payload := new(UpdateClientDto)

    if err := c.BodyParser(&payload); err != nil {
      return nil, fiber.NewError(fiber.StatusNotFound, "BODY_PARSING_ERROR")
    }

    body, err := ClientService.UpdateClient(id, payload)

    if err != nil {
      return nil, err
    }

    return body, nil
  })

	ClientController.Delete("/:id", func(c *fiber.Ctx) (any, error) {
		id := c.Params("id")

		body, err := ClientService.DeleteClient(id)

		if err != nil {
			return nil, err
		}

		return body, nil
	})
}
