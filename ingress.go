package main

type UserConfig struct {
	Name              string            `json:"name"`
	Annotation        map[string]string `json:"ingress-annotation"`
	ControllerService string            `json:"ingress-controller-service"`
	IPType            string            `json:"ingress-ip"`
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
