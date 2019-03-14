package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("json")
	viper.SetConfigFile("./myJSONConfig.json")
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())
	viper.ReadInConfig()

	if viper.IsSet("item1.key1") {
		fmt.Println("item1.key1:", viper.Get("item1.key1"))
	} else {
		fmt.Println("item1.key1 not set!")
	}

	if viper.IsSet("item2.key3") {
		fmt.Println("item2.key3:", viper.Get("item2.key3"))
	} else {
		fmt.Println("item2.key3 is not set!")
	}

	if !viper.IsSet("item3.key1") {
		fmt.Println("item3.key1 is not set!")
	}
}
