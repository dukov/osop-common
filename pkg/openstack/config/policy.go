package config

import (
	"sigs.k8s.io/yaml"
)

// Policy abstraction for service policy.yaml
type Policy map[string]string

// Merge configuration to object
func (p Policy) Merge(src Policy) {
	for k, v := range src {
		p[k] = v
	}
}

// ToString returns string representation of the object
func (p Policy) ToString() string {
	content, err := yaml.Marshal(p)
	if err != nil {
		return ""
	}
	return string(content)
}
