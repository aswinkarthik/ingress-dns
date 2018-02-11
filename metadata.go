package main

type Metadata struct {
	Name        string            `json:"name"`
	Annotations map[string]string `json:"annotations"`
}

func (m Metadata) ContainsAnnotations(annotations map[string]string) bool {
	for k, v := range annotations {
		if value, present := m.Annotations[k]; !present || v != value {
			return false
		}
	}
	return true
}
