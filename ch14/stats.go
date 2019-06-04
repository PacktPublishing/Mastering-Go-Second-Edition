package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func min() float64 {

	return 0
}

func max() float64 {

	return 0
}

func meanValue() float64 {

	return 0
}

func medianValue() float64 {

	// Odd

	// Even

	return 0
}

func variance() float64 {

	return 0
}

func covariance() float64 {

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

		value, err := strconv.ParseFloat(line, 64)
		if err == nil {
			data = append(data, value)
		}
	}

	// Sort slice

	fmt.Println(data)

}
