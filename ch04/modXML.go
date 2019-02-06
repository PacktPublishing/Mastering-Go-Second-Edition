package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func main() {
	type Address struct {
		City, Country string
	}
	type Employee struct {
		XMLName   xml.Name `xml:"employee"`
		Id        int      `xml:"id,attr"`
		FirstName string   `xml:"name>first"`
		LastName  string   `xml:"name>last"`
		Initials  string   `xml:"name>initials"`
		Height    float32  `xml:"height,omitempty"`
		Address
		Comment string `xml:",comment"`
	}

	r := &Employee{Id: 7, FirstName: "Mihalis", LastName: "Tsoukalos", Initials: "MIT"}
	r.Comment = "Technical Writer + DevOps"
	r.Address = Address{"SomeWhere 12", "12312, Greece"}

	output, err := xml.MarshalIndent(r, "  ", "    ")
	if err != nil {
		fmt.Println("Error:", err)
	}
	output = []byte(xml.Header + string(output))
	os.Stdout.Write(output)
	os.Stdout.Write([]byte("\n"))
}
