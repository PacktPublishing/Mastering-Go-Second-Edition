package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var minusD bool = false
var minusF bool = false

func walk(path string, info os.FileInfo, err error) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	mode := fileInfo.Mode()
	if mode.IsRegular() && minusF {
		fmt.Println("+", path)
		return nil
	}

	if mode.IsDir() && minusD {
		fmt.Println("*", path)
		return nil
	}

	fmt.Println(path)
	return nil
}

func main() {
	starD := flag.Bool("d", false, "Signify directories")
	plusF := flag.Bool("f", false, "Signify regular files")
	flag.Parse()
	flags := flag.Args()

	Path := "."
	if len(flags) == 1 {
		Path = flags[0]
	}

	minusD = *starD
	minusF = *plusF

	err := filepath.Walk(Path, walk)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
