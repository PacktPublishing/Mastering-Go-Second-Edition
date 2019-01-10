package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	r := strings.NewReader("test")
	fmt.Println("r length:", r.Len())

	b := make([]byte, 1)
	for {
		n, err := r.Read(b)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Read %s Bytes: %d\n", b, n)
	}

	s := strings.NewReader("This is an error!\n")
	fmt.Println("r length:", s.Len())
	n, err := s.WriteTo(os.Stderr)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Wrote %d bytes to os.Stderr\n", n)
}
