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
