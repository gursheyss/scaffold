package docker

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

type progressDetail struct {
	Current int `json:"current"`
	Total   int `json:"total"`
}

type progressMessage struct {
	Status         string         `json:"status"`
	ProgressDetail progressDetail `json:"progressDetail"`
	ID             string         `json:"id"`
}

func CheckDockerRunning() error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	_, err = cli.Ping(context.Background())
	if err != nil {
		return fmt.Errorf("docker daemon is not running, either start it up or install it from https://docs.docker.com/engine/install/")
	}

	return nil
}

func StartMySQLContainer(username string, password string, dbname string, port int) (string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", err
	}

	ctx := context.Background()

	out, err := cli.ImagePull(ctx, "docker.io/library/mysql:latest", types.ImagePullOptions{})
	if err != nil {
		return "", err
	}
	defer out.Close()

	dec := json.NewDecoder(out)
	for {
		var msg progressMessage
		if err := dec.Decode(&msg); err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}

		if msg.ProgressDetail.Total > 0 {
			progress := float64(msg.ProgressDetail.Current) / float64(msg.ProgressDetail.Total) * 100
			bar := strings.Repeat("=", int(progress)/2) + strings.Repeat(" ", 50-int(progress)/2)
			fmt.Printf("\r%s: [%-50s] %.2f%%", msg.Status, bar, progress)
		}
	}
	fmt.Println()

	exposedPort, err := nat.NewPort("tcp", fmt.Sprintf("%d", port))
	if err != nil {
		return "", err
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "mysql:latest",
		Env: []string{
			fmt.Sprintf("MYSQL_ROOT_PASSWORD=%s", password),
		},
		ExposedPorts: nat.PortSet{
			exposedPort: struct{}{},
		},
	}, nil, nil, nil, "")
	if err != nil {
		return "", err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", err
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, "localhost", port, dbname)

	return connectionString, nil
}
