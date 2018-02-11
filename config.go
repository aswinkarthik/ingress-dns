package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	APIKey         string
	Host           string
	Port           string
	Protocol       string
	SkipTlsVerify  bool
	IngressConfigs []IngressConfig
}

const HTTP = "http"
const HTTPS = "https"

var appConfig *Config

func loadConfig() {
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	appConfig = &Config{
		APIKey:         getKubernetesAPIToken(),
		Host:           viper.GetString("KUBERNETES_SERVICE_HOST"),
		Port:           viper.GetString("KUBERNETES_SERVICE_PORT"),
		Protocol:       getProtocol(),
		SkipTlsVerify:  viper.GetBool("SKIP_TLS_VERIFY"),
		IngressConfigs: getIngressConfigs(),
	}
}

func getProtocol() string {
	if port := viper.GetInt("KUBERNETES_SERVICE_PORT"); port == 443 {
		return HTTPS
	} else if isHttps := viper.GetBool("USE_HTTPS"); isHttps {
		return HTTPS
	} else {
		return HTTP
	}
}

func getKubernetesAPIToken() string {
	tokenFile := "/var/run/secrets/kubernetes.io/serviceaccount/token"
	data, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		fmt.Println(err)
		return viper.GetString("KUBERNETES_API_TOKEN")
	}
	return string(data)
}

func getIngressConfigs() []IngressConfig {
	var configs []IngressConfig

	configAsString := viper.GetString("INGRESS_CONFIGS")
	if err := json.Unmarshal([]byte(configAsString), &configs); err != nil {
		log.Fatal(err)
	}
	return configs
}
