package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/miltlima/goact/pkg"
	"github.com/spf13/cobra"
)

var stack string

var rootCmd = &cobra.Command{
	Use:   "goact",
	Short: "Tool for creating default file for Github Actions ",
}

var createDockerfile bool

var createStackCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Github action configuration for stack specified",
	Run: func(cmd *cobra.Command, args []string) {
		if stack == "" {
			fmt.Println("Error: you need pass the flag -s or --stack.")
			return
		}

		actionsConfig, err := pkg.GenerateGithubActionsConfig(stack)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = pkg.CreateGithubActionsFiles(stack, string(actionsConfig))
		if err != nil {
			fmt.Println("Error to create pipeline.yaml", err)
		} else {
			fmt.Println("File created successfully")
		}

		if createDockerfile {
			fmt.Println("Creating Dockerfile...")
			dockerfileContent, err := pkg.GenerateDockerfile(stack)
			if err != nil {
				fmt.Println("Error generating Dockerfile:", err)
				return
			}

			err = ioutil.WriteFile("Dockerfile", []byte(dockerfileContent), 0644)
			if err != nil {
				fmt.Println("Erro creating Dockerfile", err)
				return
			}
			fmt.Println("Dockerfile created sucessfully")
		}
	},
}

func init() {
	createStackCmd.Flags().StringVarP(&stack, "stack", "s", "", "Define the stack (node.js, python and others)")
	createStackCmd.Flags().BoolVarP(&createDockerfile, "dockerfile", "d", false, "Create Dockerfile for the specified Stack")
	rootCmd.AddCommand(createStackCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
