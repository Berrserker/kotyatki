// Package classification Account API.
//
// this is to show how to write RESTful APIs in golang.
// that is to provide a detailed overview of the language specs
//
// Terms Of Service:
//
//     Schemes: http, https
//     Host: localhost:4900
//     Version: 3.0.0
//     Contact: YagoSevatarion<Berrserk.s@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//      -JWT: []
//
//     securityDefinitions:
//		  JWT:
//		    type: apiKey
//		    in: header
//		    name: access_token
//
// swagger:meta
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
