package main

import (
	"github.com/ermos/annotation"
	"github.com/ermos/annotation/parser"
	"log"
)

func main() {
	var api []parser.API
	err := annotation.Fetch("module/auth", &api, parser.ToAPI)
	if err != nil {
		log.Fatal(err)
	}
	err = annotation.Save(api, "./api.json")
	if err != nil {
		log.Fatal(err)
	}
}
