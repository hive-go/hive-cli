package main

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

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
			serviceFile, err := os.ReadFile("service.txt")
			if err != nil {
				panic(err)
			}
			serviceContent := strings.ReplaceAll(string(serviceFile), "#MODULE_NAME", moduleName)
			serviceContent = strings.ReplaceAll(serviceContent, "#CAPITALIZED_NAME", capitalizedModuleName)
			os.WriteFile("src/modules/"+moduleName+"/"+moduleName+".service.go", []byte(serviceContent), 0644)

			//prepare DTO
			dtoFile, err := os.ReadFile("dto.txt")
			if err != nil {
				panic(err)
			}
			dtoContent := strings.ReplaceAll(string(dtoFile), "#MODULE_NAME", moduleName)
			dtoContent = strings.ReplaceAll(dtoContent, "#CAPITALIZED_NAME", capitalizedModuleName)
			os.WriteFile("src/modules/"+moduleName+"/"+moduleName+".dto.go", []byte(dtoContent), 0644)

			//prepare controller
			controllerFile, err := os.ReadFile("controller.txt")
			if err != nil {
				panic(err)
			}
			controllerContent := strings.ReplaceAll(string(controllerFile), "#MODULE_NAME", moduleName)
			controllerContent = strings.ReplaceAll(controllerContent, "#CAPITALIZED_NAME", capitalizedModuleName)
			os.WriteFile("src/modules/"+moduleName+"/"+moduleName+".controller.go", []byte(controllerContent), 0644)

			//prepare module
			moduleFile, err := os.ReadFile("module.txt")
			if err != nil {
				panic(err)
			}
			moduleContent := strings.ReplaceAll(string(moduleFile), "#MODULE_NAME", moduleName)
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
