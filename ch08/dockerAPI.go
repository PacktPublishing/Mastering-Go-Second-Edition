package main

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func listContainers() error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return (err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return (err)
	}

	for _, container := range containers {
		fmt.Println("Images:", container.Image, "with ID:", container.ID)
	}
	return nil
}

func listImages() error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return (err)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return (err)
	}

	for _, image := range images {
		fmt.Printf("Images %s with size %d\n", image.RepoTags, image.Size)
	}
	return nil
}

func main() {
	fmt.Println("The available images are:")
	err := listImages()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("The running Containers are:")
	err = listContainers()
	if err != nil {
		fmt.Println(err)
	}
}
