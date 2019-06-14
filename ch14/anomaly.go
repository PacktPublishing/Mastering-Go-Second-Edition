package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/lytics/anomalyzer"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: anomaly MAX\n")
		return
	}

	MAX, err := strconv.Atoi(flag.Args()[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	conf := &anomalyzer.AnomalyzerConf{
		Sensitivity: 0.1,
		UpperBound:  5,
		LowerBound:  anomalyzer.NA,
		ActiveSize:  1,
		NSeasons:    4,
		Methods:     []string{"diff", "fence", "magnitude", "ks"},
	}

	data := []float64{}
	SEED := time.Now().Unix()
	rand.Seed(SEED)

	for i := 0; i < MAX; i++ {
		data = append(data, float64(random(0, MAX)))
	}
	fmt.Println("data:", data)

	anom, _ := anomalyzer.NewAnomalyzer(conf, data)
	prob := anom.Push(8.0)
	fmt.Println("Anomalous Probability:", prob)
}
