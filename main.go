package main

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var ControllerString = `package #MODULE_NAME

import (
	"github.com/hive-go/hive"
  "github.com/gofiber/fiber/v2"
)

var #CAPITALIZED_NAMEController = hive.CreateController()

func init() {

	#CAPITALIZED_NAMEController.SetConfig(hive.ControllerConfig{
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

	#CAPITALIZED_NAMEController.Get("", func(c *fiber.Ctx) (any, error) {
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

var ServiceString = `package #MODULE_NAME

import (
  "github.com/gofiber/fiber/v2"
)

type #CAPITALIZED_NAMEServiceStruct struct{}

var #CAPITALIZED_NAMEService = #CAPITALIZED_NAMEServiceStruct{}

func (u *#CAPITALIZED_NAMEServiceStruct) Create#CAPITALIZED_NAME(data *Create#CAPITALIZED_NAMEDto) (any, error) {
	return fiber.Map{
		"message": "#CAPITALIZED_NAME created successfully !!!",
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

func (u *#CAPITALIZED_NAMEServiceStruct) Update#CAPITALIZED_NAME(id string, data *Update#CAPITALIZED_NAMEDto) (any, error) {
  return fiber.Map{
    "message": "#CAPITALIZED_NAME updated successfully for id " + id,
  }, nil
}

func (u *#CAPITALIZED_NAMEServiceStruct) Delete#CAPITALIZED_NAME(data *Delete#CAPITALIZED_NAMEDto) (any, error) {
  return fiber.Map{
    "message": "#CAPITALIZED_NAME deleted successfully",
  }, nil
}
`

var ModuleString = `package #MODULE_NAME

import (
	"github.com/hive-go/hive"
)

var #CAPITALIZED_NAMEModule = hive.CreateModule()

func init() {
	#CAPITALIZED_NAMEModule.AddController(#CAPITALIZED_NAMEController)
}
`

var DTOString = `package #MODULE_NAME

type Create#CAPITALIZED_NAMEDto struct {
}

type Update#CAPITALIZED_NAMEDto struct {
}
`

func main() {
	var rootCmd = &cobra.Command{Use: "hive"}

	var generateCmd = &cobra.Command{
		Use:   "generate_resource [resource name]",
		Short: "Generate a new resource",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			moduleName := args[0]
			capitalizedModuleName := capitalize(moduleName)
			os.Mkdir("src", 0755)
			os.Mkdir("src/modules", 0755)
			os.Mkdir("src/modules/"+moduleName, 0755)

			//prepare service
			//read service.txt in root and replace all occurrences of #MODULE_NAME with moduleName and all occurrences of #CAPITALIZED_NAME with capitalizedModuleName
			//after it write the content to src/modules/moduleName/module_name.service.go
			//start
			serviceContent := strings.ReplaceAll(ServiceString, "#MODULE_NAME", moduleName)
			serviceContent = strings.ReplaceAll(serviceContent, "#CAPITALIZED_NAME", capitalizedModuleName)
			os.WriteFile("src/modules/"+moduleName+"/"+moduleName+".service.go", []byte(serviceContent), 0644)

			dtoContent := strings.ReplaceAll(DTOString, "#MODULE_NAME", moduleName)
			dtoContent = strings.ReplaceAll(dtoContent, "#CAPITALIZED_NAME", capitalizedModuleName)
			os.WriteFile("src/modules/"+moduleName+"/"+moduleName+".dto.go", []byte(dtoContent), 0644)

			//prepare controller
			controllerContent := strings.ReplaceAll(ControllerString, "#MODULE_NAME", moduleName)
			controllerContent = strings.ReplaceAll(controllerContent, "#CAPITALIZED_NAME", capitalizedModuleName)
			os.WriteFile("src/modules/"+moduleName+"/"+moduleName+".controller.go", []byte(controllerContent), 0644)

			moduleContent := strings.ReplaceAll(ModuleString, "#MODULE_NAME", moduleName)
			moduleContent = strings.ReplaceAll(moduleContent, "#CAPITALIZED_NAME", capitalizedModuleName)
			os.WriteFile("src/modules/"+moduleName+"/"+moduleName+".module.go", []byte(moduleContent), 0644)

		},
	}

	rootCmd.AddCommand(generateCmd)
	rootCmd.Execute()
}

func capitalize(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}
