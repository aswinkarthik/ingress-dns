package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetServices() ServiceList {
	log.Println("Preparing to Get all services")
	response := request("api/v1/services", http.MethodGet, nil)
	defer response.Body.Close()
	var serviceList ServiceList
	if err := json.NewDecoder(response.Body).Decode(&serviceList); err != nil {
		log.Fatal(err)
	}
	return serviceList
}

func request(path string, method string, data io.Reader) *http.Response {
	url := fmt.Sprintf("%s://%s:%s/%s", appConfig.Protocol, appConfig.Host, appConfig.Port, path)
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: appConfig.SkipTlsVerify},
		},
	}
	request, err := http.NewRequest(method, url, data)
	if err != nil {
		log.Fatal(err)
	}
	authorizationToken := fmt.Sprintf("Bearer %s", appConfig.APIKey)
	request.Header.Set("Authorization", authorizationToken)
	response, respErr := client.Do(request)
	if respErr != nil {
		log.Fatal(respErr)
	}
	return response
}
