package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Binding struct binds IngressConfig (user-defined) with an Ingress (Kube resource) and Service (Kube resource)
type Binding struct {
	UserConfig
	Service
	Ingress
}

type BindingV2 struct {
	UserConfig
	Service
	Ingresses []Ingress
}

func init() {
	hostRegexPattern, _ = regexp.Compile("^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\-]*[A-Za-z0-9])$")
}

// NewBindings takes in ServiceList, IngressList, []UserConfig and creates a Binding
// Binding has all 3 grouped appropriately
func NewBindings(serviceList ServiceList, ingressList IngressList, userConfigs []UserConfig) []BindingV2 {
	serviceMap := serviceList.GetServiceMap()
	ingresses := ingressList.Items
	finalBindings := make([]BindingV2, len(userConfigs))

	for i, config := range userConfigs {
		controllerService := serviceMap[config.ControllerService]
		ingressesForConfig := make([]Ingress, len(ingresses))
		ingressCount := 0
		for _, ingress := range ingressList.Items {
			if ingress.ContainsAnnotations(config.Annotation) {
				ingressesForConfig[ingressCount] = ingress
				ingressCount++
			}
		}
		finalBindings[i] = BindingV2{config, controllerService, ingressesForConfig[:ingressCount]}
	}
	return finalBindings
}

// GetIPAddress returns the IP address of the service based on the IPType
// A service can either have a clusterIP or externalIP. The method returns approprately
// If the IngressConfig.IPType is anything other than clusterIP or externalIP empty string is returned
func (b *BindingV2) GetIPAddress() string {
	switch b.UserConfig.IPType {
	case "clusterIP":
		return b.Service.Spec.ClusterIP
	case "externalIP":
		return b.Service.Spec.ExternalIP
	default:
		return ""
	}
}

func (b *BindingV2) getHosts() []string {
	sum := 0
	for _, ingress := range b.Ingresses {
		sum += len(ingress.Spec.Rules)
	}

	hosts := make([]string, sum)
	counter := 0
	for _, ingress := range b.Ingresses {
		for _, rule := range ingress.Spec.Rules {
			hosts[counter] = rule.Host
			counter++
		}
	}
	return hosts[:counter]
}

// GetConsulDto converts a Binding object to a struct that can be PUT to consul HTTP Service
func (b *BindingV2) GetConsulDto() ConsulDto {
	domain := fmt.Sprintf(".%s.%s", b.Service.Metadata.Name, appConfig.ConsulDomain)
	hosts := b.getHosts()
	tags := make([]string, len(hosts))
	counter := 0
	for _, host := range hosts {
		if isValidHost(host, domain) {
			tags[counter] = getTag(host, domain)
			counter++
		}
	}
	return ConsulDto{
		ID:      b.Service.Metadata.Name,
		Name:    b.Service.Metadata.Name,
		Tags:    tags[:counter],
		Address: b.GetIPAddress(),
	}
}

func isValidHost(host string, domain string) bool {
	validHost := hostRegexPattern.MatchString(host)
	hasSameDomain := strings.HasSuffix(host, domain)
	return validHost && hasSameDomain
}

func getTag(host string, domain string) string {
	return strings.Replace(host, domain, "", -1)
}
