package cmd

import (
	"fmt"

	"github.com/miltlima/goact/pkg"
	"github.com/spf13/cobra"
)

var stack string

var rootCmd = &cobra.Command{
	Use:   "goact",
	Short: "Tool for creating default file for Github Actions ",
}

var createCmd = &cobra.Command{
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
	},
}

func init() {
	createCmd.Flags().StringVarP(&stack, "stack", "s", "", "Define the stack (node.js, python and others)")

	rootCmd.AddCommand(createCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
