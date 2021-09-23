package container

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type iDockerClient interface {
	ListAll() ([]string, error)
	StopContainer(a string) error
	StartContainer(a string) error
}

type listContainer []types.Container

func (c listContainer) ListAll() ([]string, error) {
	if c != nil {
		result := []string{}
		for _, container := range c {
			resultPort := []string{}
			for _, items := range container.Ports {
				if items.PublicPort != 0 {
					resultPort = append(resultPort, fmt.Sprintf("%v:%v->%v/%v ", items.IP, items.PublicPort, items.PrivatePort, items.Type))
				} else {
					resultPort = append(resultPort, fmt.Sprintf("%v/%v ", items.PrivatePort, items.Type))
				}
			}
			result = append(result, fmt.Sprintf("%s %s %s %s %s\n", container.ID[:10], container.Names[0][1:], container.Image, strings.Join(resultPort, ""), container.State))
		}
		return result, nil
	}
	return []string{}, errors.New("empty container list")
}

func (c listContainer) StopContainer(a string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	err = cli.ContainerStop(ctx, a, nil)
	if err != nil {
		panic(err)
	}
	return nil
}

func (c listContainer) StartContainer(a string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	err = cli.ContainerStart(ctx, a, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}
	return nil
}

func getContainer() ([]types.Container, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}
	return containers, nil
}

func ActiontListContainer() {
	var actions iDockerClient
	if result, err := getContainer(); err != nil {
		fmt.Println(err)
	} else {
		actions = listContainer(result)
	}
	if result, err := actions.ListAll(); err != nil {
		fmt.Println(err)
	} else {
		for _, items := range result {
			fmt.Println(items)
		}
	}
}

func ActionStopContainer(c string) error {
	var actions iDockerClient
	actions = listContainer(make([]types.Container, 0))
	if err := actions.StopContainer(c); err != nil {
		return err
	} else {
		return nil
	}
}

func ActionStartContainer(c string) error {
	var actions iDockerClient
	actions = listContainer(make([]types.Container, 0))
	if err := actions.StartContainer(c); err != nil {
		return err
	} else {
		return nil
	}
}
