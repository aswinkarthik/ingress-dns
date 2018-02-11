package main

type IngressConfig struct {
	Annotation        map[string]string `json:"ingress-annotation"`
	ControllerService string            `json:"ingress-controller-service"`
	IPType            string            `json:"ingress-ip"`
	Name              string            `json:"name"`
}

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
	Spec     struct {
		Rules []struct {
			Host string `json:"host"`
		} `json:"rules"`
	} `json:"spec"`
}

func (b Binding) GetIpAddress() string {
	switch b.IngressConfig.IPType {
	case "clusterIP":
		return b.Service.Spec.ClusterIP
	case "externalIP":
		return b.Service.Spec.ExternalIP
	default:
		return ""
	}
}

func (b Binding) GetHosts() []string {
	rules := b.Ingress.Spec.Rules
	hosts := make([]string, len(rules))
	for i, rule := range rules {
		hosts[i] = rule.Host
	}
	return hosts
}

func (b Binding) GetId() string {
	return b.IngressConfig.Name
}

func (b Binding) GetName() string {
	return b.IngressConfig.Name
}
