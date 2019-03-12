package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.BindEnv("GOMAXPROCS")
	val := viper.Get("GOMAXPROCS")
	fmt.Println(val)
	viper.Set("GOMAXPROCS", 10)
	val = viper.Get("GOMAXPROCS")
	fmt.Println(val)
}
