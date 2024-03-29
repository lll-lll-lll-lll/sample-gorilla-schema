package main

import (
	"log"
	"net/http"
	"github.com/gorilla/schema"
	
)
var decoder = schema.NewDecoder()

type D struct {
	S string   `schema:"s,default:test1"`
}


func main() {
	data := map[string][]string{}

	d := D{}


	if err := decoder.Decode(&d, data); err != nil {
		log.Fatal("Error while decoding:", err)
	}

}