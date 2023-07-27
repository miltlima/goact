package pkg

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/miltlima/goact/pkg/stacks"
)

type GitHubActionsConfig string

func GenerateGithubActionsConfig(stack string) (GitHubActionsConfig, error) {
	switch stack {
	case "node.js":
		return GitHubActionsConfig(stacks.NodeJS), nil
	case "python":
		return GitHubActionsConfig(stacks.Python), nil
	case "ruby":
		return GitHubActionsConfig(stacks.Ruby), nil
	case "golang":
		return GitHubActionsConfig(stacks.Golang), nil
	case "java":
		return GitHubActionsConfig(stacks.Java), nil
	default:
		return "", fmt.Errorf("Stack '%s' is not supported.", stack)
	}
}

func CreateGithubActionsFiles(stack, config string) error {
	err := os.MkdirAll(".github/workflows", os.ModePerm)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(".github/workflows/pipeline.yaml", []byte(config), 0644)
	if err != nil {
		return err
	}

	return nil

}
