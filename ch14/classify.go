package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 2 {
		fmt.Printf("usage: classify filename k\n")
		return
	}

	dataset := flag.Args()[0]
	rawData, err := base.ParseCSVToInstances(dataset, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	k, err := strconv.Atoi(flag.Args()[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	cls := knn.NewKnnClassifier("euclidean", "linear", k)
	train, test := base.InstancesTrainTestSplit(rawData, 0.50)
	cls.Fit(train)

	p, err := cls.Predict(test)
	if err != nil {
		fmt.Println(err)
		return
	}

	confusionMat, err := evaluation.GetConfusionMatrix(test, p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(evaluation.GetSummary(confusionMat))
}
