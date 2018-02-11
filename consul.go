package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type ConsulService struct {
	ID      string   `json:"ID"`
	Name    string   `json:"Name"`
	Tags    []string `json:"Tags"`
	Address string   `json:"Address"`
}

func SendToConsul(bindings []Binding) {
	consulServices := convertToConsulService(bindings)
	for _, service := range consulServices {
		consulRequest(service)
	}
}

func convertToConsulService(bindings []Binding) []ConsulService {
	consulServices := make([]ConsulService, len(bindings))
	counter := 0
	for _, binding := range bindings {
		id := binding.GetId()
		name := binding.GetName()
		serviceDomain := fmt.Sprintf(".%s.%s", id, appConfig.ConsulDomain)
		address := binding.GetIpAddress()
		hosts := binding.GetHosts()
		tags := getTags(hosts, serviceDomain)
		if len(tags) > 0 && address != "" && id != "" && name != "" {
			consulServices[counter] = ConsulService{
				ID:      id,
				Name:    name,
				Address: address,
				Tags:    tags,
			}
		}
	}
	return consulServices
}

func getTags(hosts []string, serviceDomain string) []string {
	tags := make([]string, len(hosts))
	counter := 0
	for _, host := range hosts {
		if strings.HasSuffix(host, serviceDomain) {
			tags[counter] = strings.Replace(host, serviceDomain, "", -1)
			counter += 1
		}
	}
	return tags[:counter]
}

func consulRequest(service ConsulService) {
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
