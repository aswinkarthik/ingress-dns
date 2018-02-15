package main

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
func (b Binding) GetIPAddress() string {
	switch b.UserConfig.IPType {
	case "clusterIP":
		return b.Service.Spec.ClusterIP
	case "externalIP":
		return b.Service.Spec.ExternalIP
	default:
		return ""
	}
}

// GetHosts returns the list of hosts configured in the binding's Ingress resource's rules
func (b Binding) GetHosts() []string {
	rules := b.Ingress.Spec.Rules
	hosts := make([]string, len(rules))
	for i, rule := range rules {
		hosts[i] = rule.Host
	}
	return hosts
}

// GetId returns the name of the user provided ingress configurations
func (b Binding) GetId() string {
	return b.UserConfig.Name
}

// GetId returns the name of the user provided ingress configurations
func (b Binding) GetName() string {
	return b.UserConfig.Name
}
