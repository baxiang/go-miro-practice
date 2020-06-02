package main

import (
	"flag"
	"fmt"
)

var configFile = flag.String("f", "details.yml", "set config file which viper will loading.")

func main() {
	flag.Parse()
	app, err := CreateApp(*configFile)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if err := app.Start(); err != nil {
		panic(err)
	}

	app.AwaitSignal()
}