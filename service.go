package main

type ServiceList struct {
	Kind    string    `json:"kind"`
	Message string    `json:"message"`
	Items   []Service `json:"items"`
}

type Service struct {
	Metadata `json:"metadata"`
	Spec     serviceSpec `json:"spec"`
}

type serviceSpec struct {
	ClusterIP  string `json:"clusterIP"`
	ExternalIP string `json:"externalIP"`
}

func (s ServiceList) GetServiceMap() map[string]Service {
	serviceMap := make(map[string]Service, len(s.Items))

	for _, service := range s.Items {
		serviceMap[service.Metadata.Name] = service
	}

	return serviceMap
}
