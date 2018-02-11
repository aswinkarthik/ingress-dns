package main

type ServiceList struct {
	Kind    string    `json:"kind"`
	Message string    `json:"message"`
	Items   []Service `json:"items"`
}

type Service struct {
	Metadata `json:"metadata"`
	Spec     `json:"spec"`
}

type Metadata struct {
	Name string `json:"name"`
}

type Spec struct {
	ClusterIP  string `json:"clusterIP"`
	ExternalIP string `json:"externalIP"`
}
