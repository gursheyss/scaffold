package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
)

func CheckDockerRunning() error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	_, err = cli.Ping(context.Background())
	if err != nil {
		return fmt.Errorf("Docker daemon is not running, either start it up or install it from https://docs.docker.com/engine/install/")
	}

	return nil
}
