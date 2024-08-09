package client

import (
  "github.com/gofiber/fiber/v2"
)

type ClientServiceStruct struct{}

var ClientService = ClientServiceStruct{}

func (u *ClientServiceStruct) CreateClient(data *CreateClientDto) (any, error) {
	return fiber.Map{
		"message": "Client created successfully",
	}, nil
}


func (u *ClientServiceStruct) GetClients(data *GetClientDto) (any, error) {
  return fiber.Map{
    "message": "Client retrieved successfully",
  }, nil
}

func (u *ClientServiceStruct) GetClientById(data *GetClientByIdDto) (any, error) {
  return fiber.Map{
    "message": "Client retrieved successfully",
  }, nil
}

func (u *ClientServiceStruct) UpdateClient(data *UpdateClientDto) (any, error) {
  return fiber.Map{
    "message": "Client updated successfully",
  }, nil
}

func (u *ClientServiceStruct) DeleteClient(data *DeleteClientDto) (any, error) {
  return fiber.Map{
    "message": "Client deleted successfully",
  }, nil
}
