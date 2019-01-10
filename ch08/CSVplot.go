package main

import (
	"encoding/csv"
	"fmt"
	"github.com/Arafatk/glot"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need a data file!")
		return
	}

	file := os.Args[1]
	_, err := os.Stat(file)
	if err != nil {
		fmt.Println("Cannot stat", file)
		return
	}

	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Cannot open", file)
		fmt.Println(err)
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	allRecords, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	xP := []float64{}
	yP := []float64{}
	for _, rec := range allRecords {
		x, _ := strconv.ParseFloat(rec[0], 64)
		y, _ := strconv.ParseFloat(rec[1], 64)
		xP = append(xP, x)
		yP = append(yP, y)
	}

	points := [][]float64{}
	points = append(points, xP)
	points = append(points, yP)
	fmt.Println(points)

	dimensions := 2
	persist := true
	debug := false
	plot, _ := glot.NewPlot(dimensions, persist, debug)

	plot.SetTitle("Using Glot with CSV data")
	plot.SetXLabel("X-Axis")
	plot.SetYLabel("Y-Axis")
	style := "circle"
	plot.AddPointGroup("Circle:", style, points)
	plot.SavePlot("output.png")
}
