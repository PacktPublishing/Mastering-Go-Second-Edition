package main

import (
	"flag"
	"fmt"
	"github.com/mash/gokmeans"
	"strconv"
)

var observations []gokmeans.Node = []gokmeans.Node{
	gokmeans.Node{20.0, 20.0, 20.0, 20.0},
	gokmeans.Node{21.0, 21.0, 21.0, 21.0},
	gokmeans.Node{100.5, 100.5, 100.5, 100.5},
	gokmeans.Node{50.1, 50.1, 50.1, 50.1},
	gokmeans.Node{64.2, 64.2, 64.2, 64.2},
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: cluster k\n")
		return
	}

	k, err := strconv.Atoi(flag.Args()[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	if success, centroids := gokmeans.Train(observations, k, 50); success {
		fmt.Println("The centroids are")
		for _, centroid := range centroids {
			fmt.Println(centroid)
		}

		fmt.Println("The clusters are the following:")
		for _, observation := range observations {
			index := gokmeans.Nearest(observation, centroids)
			fmt.Println(observation, "belongs in cluster", index+1, ".")
		}
	}
}
