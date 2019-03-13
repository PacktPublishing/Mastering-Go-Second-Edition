package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.BindEnv("GOMAXPROCS")
	val := viper.Get("GOMAXPROCS")
	fmt.Println("GOMAXPROCS:", val)
	viper.Set("GOMAXPROCS", 10)
	val = viper.Get("GOMAXPROCS")
	fmt.Println("GOMAXPROCS:", val)

	viper.BindEnv("NEW_VARIABLE")
	val = viper.Get("NEW_VARIABLE")
	if val == nil {
		fmt.Println("NEW_VARIABLE not defined.")
		return
	}
	fmt.Println(val)
}
