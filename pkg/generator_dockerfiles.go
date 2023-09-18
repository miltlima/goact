package pkg

import (
	"fmt"

	"github.com/miltlima/goact/pkg/dockerfiles"
)

type DockerfileConfig string

func GenerateDockerfile(stack string) (DockerfileConfig, error) {
	switch stack {
	case "node.js", "nodejs":
		return DockerfileConfig(dockerfiles.NodeJS), nil
	case "python":
		return DockerfileConfig(dockerfiles.Python), nil
	case "java":
		return DockerfileConfig(dockerfiles.Java), nil
	case "golang":
		return DockerfileConfig(dockerfiles.Golang), nil
	default:
		return "", fmt.Errorf("Dockerfile '%s' is not supported")
	}
}
