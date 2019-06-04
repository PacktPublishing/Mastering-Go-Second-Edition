package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func min(x []float64) float64 {
	return x[0]
}

func max(x []float64) float64 {
	return x[len(x)-1]
}

func meanValue(x []float64) float64 {

	return 0
}

func medianValue(x []float64) float64 {

	// Odd

	// Even

	return 0
}

func variance(x []float64) float64 {

	return 0
}

func covariance(x []float64) float64 {

	return 0
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: byLine filename\n")
		return
	}

	data := make([]float64, 0)

	file := flag.Args()[0]
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
		line = strings.TrimRight(line, "\r\n")
		value, err := strconv.ParseFloat(line, 64)
		if err == nil {
			data = append(data, value)
		}
	}

	// Sort slice
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})

	fmt.Println("min:", min(data))
	fmt.Println("min:", max(data))

}
