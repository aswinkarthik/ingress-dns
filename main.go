package main

import (
	"github.com/spf13/cobra"
)

func main() {
	loadConfig()
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start Ingress-dns in non-daemon mode",
		Run: func(cmd *cobra.Command, args []string) {
			start(args)
		},
	}

	rootCmd := &cobra.Command{
		Use: "ingress-dns",
	}
	rootCmd.AddCommand(startCmd)
	rootCmd.Execute()
}

func start(args []string) {
	SendToConsul(getBindings())
	blockForever()
}

func blockForever() {
	select {}
}

func getBindings() []Binding {
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
