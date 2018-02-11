package main

import "log"

func catch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
