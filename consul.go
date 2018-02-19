package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

var hostRegexPattern *regexp.Regexp

// ConsulDto represents a single object to register a given service to Consul
type ConsulDto struct {
	ID      string   `json:"ID"`
	Name    string   `json:"Name"`
	Tags    []string `json:"Tags"`
	Address string   `json:"Address"`
}

func consulRequest(service ConsulDto) {
	url := fmt.Sprintf("http://%s:%s/v1/agent/service/register", appConfig.ConsulHost, appConfig.ConsulPort)
	client := http.Client{}
	body, bodyErr := json.Marshal(service)
	catch(bodyErr)

	req, reqErr := http.NewRequest(http.MethodPut, url, bytes.NewReader(body))
	catch(reqErr)

	res, resErr := client.Do(req)
	catch(resErr)

	if res.StatusCode != 200 {
		log.Print("Failed to update service in Consul ")
		log.Println(res.StatusCode)
	}
	log.Println(fmt.Sprintf("Updated nginx-controller %s ", service.Name))
}
