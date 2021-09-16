package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type SportContainer []types.Port

func (listPorts SportContainer) ListPort() string {
	if listPorts != nil {
		resultPort := []string{}
		for _, items := range listPorts {
			if items.PublicPort != 0 {
				resultPort = append(resultPort, fmt.Sprintf("%v:%v->%v/%v ", items.IP, items.PublicPort, items.PrivatePort, items.Type))
			} else {
				resultPort = append(resultPort, fmt.Sprintf("%v/%v ", items.PrivatePort, items.Type))
			}
		}
		return strings.Join(resultPort, "")
	}
	return ""
}

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		listPorts := SportContainer(container.Ports)
		fmt.Printf("%s %s %s %s %s\n", container.ID[:10], container.Names[0][1:], container.Image, listPorts.ListPort(), container.Status)
	}
}
