package main

import (
	"log"
)

func runWorker() {
	log.Println("Worker running...")

	serviceList := client.GetServices()
	ingressList := client.GetIngresses()
	userConfigs := appConfig.UserConfigs

	bindings := NewBindings(serviceList, ingressList, userConfigs)

	if debugEnabled {
		prettyPrint(bindings)
	}
}