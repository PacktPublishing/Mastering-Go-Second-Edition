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

func variance(x []float64) float64 {
	mean := meanValue(x)
	sum := float64(0)
	for _, v := range x {
		sum = sum + (v-mean)*(v-mean)
	}
	return sum / float64(len(x))
}

func meanValue(x []float64) float64 {
	sum := float64(0)
	for _, v := range x {
		sum = sum + v
	}
	return sum / float64(len(x))
}

func outliers(x []float64, limit float64) []float64 {
	deviation := math.Sqrt(variance(x))
	mean := meanValue(x)
	anomaly := deviation * limit
	lower_limit := mean - anomaly
	upper_limit := mean + anomaly
	fmt.Println(lower_limit, upper_limit)

	y := make([]float64, 0)
	for _, val := range x {
		if val < lower_limit || val > upper_limit {
			y = append(y, val)
		}
	}

	return y
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 2 {
		fmt.Printf("usage: stats filename limit\n")
		return
	}

	file := flag.Args()[0]
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	limit, err := strconv.ParseFloat(flag.Args()[1], 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := make([]float64, 0)
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
	out := outliers(data, limit)
	fmt.Println(out)
}
