package main

func GetProcessableBindings() []Binding {
	bindings := make([]Binding, len(appConfig.IngressConfigs))

	counter := 0
	for _, service := range GetServices().Items {
		for _, config := range appConfig.IngressConfigs {
			if config.ControllerService == service.Metadata.Name {
				bindings[counter] = Binding{config, service}
				counter += 1
			}
		}
	}

	return bindings[:counter]
}
