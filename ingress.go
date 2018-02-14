package main

type UserConfig struct {
	Annotation        map[string]string `json:"ingress-annotation"`
	ControllerService string            `json:"ingress-controller-service"`
	IPType            string            `json:"ingress-ip"`
	Name              string            `json:"name"`
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
