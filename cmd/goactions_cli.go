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
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stack = args[0]
		actionsConfig, err := pkg.GenerateGithubActionsConfig(stack)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(actionsConfig)

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
	createCmd.MarkFlagRequired("stack")

	rootCmd.AddCommand(createCmd)
}

func Execute() {
	if err := createCmd.Execute(); err != nil {

		fmt.Println(err)
	}
}