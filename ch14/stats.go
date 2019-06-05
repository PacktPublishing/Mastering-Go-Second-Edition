package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
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
	sum := float64(0)
	for _, v := range x {
		sum = sum + v
	}
	return sum / float64(len(x))
}

func medianValue(x []float64) float64 {
	length := len(x)
	if length%2 == 1 {
		// Odd
		return x[(length-1)/2]
	} else {
		// Even
		return (x[length/2] + x[(length/2)-1]) / 2
	}
	return 0
}

func variance(x []float64) float64 {
	mean := meanValue(x)
	sum := float64(0)
	for _, v := range x {
		sum = sum + (v-mean)*(v-mean)
	}
	return sum / float64(len(x))
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: stats filename\n")
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

	sort.Float64s(data)

	fmt.Println("Min:", min(data))
	fmt.Println("Max:", max(data))
	fmt.Println("Mean:", meanValue(data))
	fmt.Println("Median:", medianValue(data))
	fmt.Println("Variance:", variance(data))
	fmt.Println("Standard Deviation:", math.Sqrt(variance(data)))
}
