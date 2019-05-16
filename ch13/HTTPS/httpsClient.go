package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		return
	}
	URL := os.Args[1]

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{},
	}
	client := &http.Client{Transport: tr}
	response, err := client.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	s := strings.TrimSpace(string(content))

	fmt.Println(s)
}
