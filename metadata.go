package main

type Metadata struct {
	Name        string            `json:"name"`
	Annotations map[string]string `json:"annotations"`
}
