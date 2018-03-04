package main

// ConsulDto represents a single object to register a given service to Consul
type ConsulDto struct {
	ID      string   `json:"ID"`
	Name    string   `json:"Name"`
	Tags    []string `json:"Tags"`
	Address string   `json:"Address"`
}
