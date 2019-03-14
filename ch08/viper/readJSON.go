package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("json")
	viper.SetConfigFile("./myConfig.json")
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())
	viper.ReadInConfig()

	if viper.IsSet("item1.key1") {
		fmt.Println("item1.val1:", viper.Get("item1.key1"))
	} else {
		fmt.Println("item1.key1 not set!")
	}

	if !viper.IsSet("item3.key1") {
		fmt.Println("item3.key1 not set!")
	}
}
