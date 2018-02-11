package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	APIKey         string
	KubeHost       string
	KubePort       string
	ConsulHost     string
	ConsulPort     string
	ConsulDomain   string
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
		KubeHost:       viper.GetString("KUBERNETES_SERVICE_HOST"),
		KubePort:       viper.GetString("KUBERNETES_SERVICE_PORT"),
		ConsulHost:     viper.GetString("CONSUL_HOST"),
		ConsulPort:     viper.GetString("CONSUL_PORT"),
		ConsulDomain:   viper.GetString("CONSUL_DOMAIN"),
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

func getConsulDomain() string {
	consulDomain := viper.GetString("CONSUL_DOMAIN")
	if consulDomain == "" {
		return "service.consul"
	}
	return consulDomain
}
