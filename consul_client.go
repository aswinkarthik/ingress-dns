package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ConsulClient struct {
	Host string
	Port string
}

var consulClient *ConsulClient

func (c *ConsulClient) Register(service ConsulDto) {
	url := fmt.Sprintf("http://%s:%s/v1/agent/service/register", c.Host, c.Port)
	client := http.Client{}
	body, bodyErr := json.Marshal(service)
	catch(bodyErr)

	if debugEnabled {
		log.Println("Sending the following consul request:")
		prettyPrint(service)
	}

	req, reqErr := http.NewRequest(http.MethodPut, url, bytes.NewReader(body))
	catch(reqErr)

	res, resErr := client.Do(req)
	catch(resErr)

	if res.StatusCode != 200 {
		log.Fatal(fmt.Sprintf("Failed to update service in Consul. Status Code: %d", res.StatusCode))
	}
	log.Println(fmt.Sprintf("Updated nginx-controller %s ", service.Name))
}
