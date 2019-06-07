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
		panic(err)
	}

	k, err := strconv.Atoi(flag.Args()[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print a pleasant summary of your data.
	fmt.Println(rawData)

	//Initialises a new KNN classifier
	cls := knn.NewKnnClassifier("euclidean", "linear", k)

	// Do a training-test split
	trainData, testData := base.InstancesTrainTestSplit(rawData, 0.50)
	cls.Fit(trainData)

	//Calculates the Euclidean distance and returns the most popular label
	predictions, err := cls.Predict(testData)
	if err != nil {
		panic(err)
	}

	// Prints precision/recall metrics
	confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
	}

	fmt.Println(evaluation.GetSummary(confusionMat))
}
