package main

import (
	"fmt"
	"io/ioutil"
	"os"

	router "./router"
)

func main() {

	//read config
	jsonFile, err := os.Open("config.json")
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	//initiate router
	router.Route(&byteValue)
}
