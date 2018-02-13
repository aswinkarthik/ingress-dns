package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var hostRegexPattern *regexp.Regexp

func init() {
	hostRegexPattern, _ = regexp.Compile("^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\-]*[A-Za-z0-9])$")
}

// ConsulService represents a single object to register a given service to Consul
type ConsulService struct {
	ID      string   `json:"ID"`
	Name    string   `json:"Name"`
	Tags    []string `json:"Tags"`
	Address string   `json:"Address"`
}

// SendToConsul Converts the list of Bindings into ConsulService objects and sends them to consul
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
		address := binding.GetIPAddress()
		hosts := binding.GetHosts()
		tags := getTags(hosts, serviceDomain)
		if len(tags) > 0 && address != "" && id != "" && name != "" {
			consulServices[counter] = ConsulService{
				ID:      id,
				Name:    name,
				Address: address,
				Tags:    tags,
			}
			counter++
		}
	}
	return consulServices[:counter]
}

func groupBindings(bindings []Binding) map[string][]Binding {
	groupedBindings := make(map[string][]Binding, len(bindings))

	for _, binding := range bindings {
		existingBindings := groupedBindings[binding.IngressConfig.Name]
		updatedBindings := append(existingBindings, binding)
		groupedBindings[binding.IngressConfig.Name] = updatedBindings
	}

	return groupedBindings
}

// getTags returns all the valid hosts as Tags after removing the parent domain
func getTags(hosts []string, parentDomain string) []string {
	tags := make([]string, len(hosts))
	counter := 0
	for _, host := range hosts {
		validHost := hostRegexPattern.MatchString(host)
		if validHost && strings.HasSuffix(host, parentDomain) {
			tags[counter] = strings.Replace(host, parentDomain, "", -1)
			counter++
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
