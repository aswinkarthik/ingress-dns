package main

import (
	"encoding/json"
	"log"
)

func catch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func prettyPrint(v interface{}) {
	data, err := json.MarshalIndent(v, "", "  ")
	catch(err)
	log.Println(string(data))
}
