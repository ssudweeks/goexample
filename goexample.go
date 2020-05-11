package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/builder"
	"github.com/docker/docker/client"
	"io/ioutil"
	"strings"
)

func main() {
	hello()
	imageList()
	imageBuild()
}

func hello() {
	fmt.Println("Hello world")
}

func imageList(){
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		for _, title := range image.RepoTags {
			fmt.Println(title)
		}
	}
}

func imageBuild() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	buildContextDirectory, _, _ := builder.GetContextFromLocalDir(".","Dockerfile")
	buildContext := strings.NewReader(buildContextDirectory)
	build, err := cli.ImageBuild(context.Background(), buildContext, types.ImageBuildOptions{})

	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(build.Body)
	build.Body.Close()
	println(body)
}




