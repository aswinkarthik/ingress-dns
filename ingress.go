package main

type IngressConfig struct {
	Annotation        map[string]string `json:"ingress-annotation"`
	ControllerService string            `json:"ingress-controller-service"`
	IPType            string            `json:"ingress-ip"`
	Name              string            `json:"name"`
}

// Binding struct binds IngressConfig (user-defined) with an Ingress (Kube resource) and Service (Kube resource)
type Binding struct {
	IngressConfig
	Service
	Ingress
}

type IngressList struct {
	Kind    string    `json:"kind"`
	Message string    `json:"message"`
	Items   []Ingress `json:"items"`
}

type Ingress struct {
	Metadata `json:"metadata"`
	Spec     ingressSpec `json:"spec"`
}

type ingressSpec struct {
	Rules []ingressRule `json:"rules"`
}

type ingressRule struct {
	Host string `json:"host"`
}

// GetIPAddress returns the IP address of the service based on the IPType
// A service can either have a clusterIP or externalIP. The method returns approprately
// If the IngressConfig.IPType is anything other than clusterIP or externalIP empty string is returned
func (b Binding) GetIPAddress() string {
	switch b.IngressConfig.IPType {
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
	return b.IngressConfig.Name
}

// GetId returns the name of the user provided ingress configurations
func (b Binding) GetName() string {
	return b.IngressConfig.Name
}
