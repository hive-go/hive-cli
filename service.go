package main

var ServiceString = `package #MODULE_NAME

import (
  "github.com/gofiber/fiber/v2"
)

type #CAPITALIZED_NAMEServiceStruct struct{}

var #CAPITALIZED_NAMEService = #CAPITALIZED_NAMEServiceStruct{}

func (u *#CAPITALIZED_NAMEServiceStruct) Create#CAPITALIZED_NAME(data *Create#CAPITALIZED_NAMEDto) (any, error) {
	return fiber.Map{
		"message": "#CAPITALIZED_NAME created successfully",
	}, nil
}


func (u *#CAPITALIZED_NAMEServiceStruct) Get#CAPITALIZED_NAMEs(data *Get#CAPITALIZED_NAMEDto) (any, error) {
  return fiber.Map{
    "message": "#CAPITALIZED_NAME retrieved successfully",
  }, nil
}

func (u *#CAPITALIZED_NAMEServiceStruct) Get#CAPITALIZED_NAMEById(data *Get#CAPITALIZED_NAMEByIdDto) (any, error) {
  return fiber.Map{
    "message": "#CAPITALIZED_NAME retrieved successfully",
  }, nil
}

func (u *#CAPITALIZED_NAMEServiceStruct) Update#CAPITALIZED_NAME(data *Update#CAPITALIZED_NAMEDto) (any, error) {
  return fiber.Map{
    "message": "#CAPITALIZED_NAME updated successfully",
  }, nil
}

func (u *#CAPITALIZED_NAMEServiceStruct) Delete#CAPITALIZED_NAME(data *Delete#CAPITALIZED_NAMEDto) (any, error) {
  return fiber.Map{
    "message": "#CAPITALIZED_NAME deleted successfully",
  }, nil
}
`
