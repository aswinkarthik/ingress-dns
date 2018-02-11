package main

func GetProcessableBindings() []Binding {
	bindings := make([]Binding, len(appConfig.IngressConfigs))

	counter := 0

	serviceMap := GetServices().GetServiceMap()

	for _, config := range appConfig.IngressConfigs {
		if service, present := serviceMap[config.ControllerService]; present {
			bindings[counter] = Binding{config, service, Ingress{}}
			counter += 1
		}
	}
	bindings = bindings[:counter]
	for _, ingress := range GetIngresses().Items {
		for i, config := range bindings {
			if ingress.Metadata.ContainsAnnotations(config.Annotation) {
				bindings[i].Ingress = ingress
			}
		}
	}

	return bindings
}
