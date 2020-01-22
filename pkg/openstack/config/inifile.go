package config

import (
	"fmt"
	"strings"
)

// IniFile abstraction
type IniFile map[string]map[string]string

// GetSection returns articular section of the config file
func (i IniFile) GetSection(name string) (map[string]string, error) {
	section, found := i[name]
	if !found {
		return nil, fmt.Errorf("Section '%s' not found", name)
	}
	return section, nil
}

// GetValue returns value for particular key in particular section
func (i IniFile) GetValue(section, key string) (string, error) {
	opts, err := i.GetSection(section)
	if err != nil {
		return "", nil
	}

	val, found := opts[key]
	if !found {
		return "", fmt.Errorf("Value for key '%s' in section '%s' not found", key, section)
	}
	return val, nil
}

func (i IniFile) HasSection(name string) bool {
	_, found := i[name]
	return found
}

func (i IniFile) HasKey(section, key string) bool {
	opts, found := i[section]

	if found {
		_, found = opts[key]
	}

	return found
}

// AddSection extends IniFile with new sections. If section exists method
// does nothing
func (i IniFile) AddSection(name string) {
	if _, found := i[name]; !found {
		i[name] = make(map[string]string)
	}
}

// SetValue sets value of the key of a section
func (i IniFile) SetValue(section, key, value string, create bool) error {
	if !i.HasSection(section) {
		if create {
			i[section] = make(map[string]string)
		} else {
			return fmt.Errorf("Section '%s' not found", section)
		}

	}
	i[section][key] = value
	return nil
}

// Merge another IniFile content
func (i IniFile) Merge(src IniFile) {
	for sec, opts := range src {
		if !i.HasSection(sec) {
			i[sec] = opts
		} else {
			for key, val := range opts {
				i[sec][key] = val
			}
		}
	}
}

// ToString convert Ini config to string
func (i IniFile) ToString() string {
	var lines []string
	for sec, opts := range i {
		lines = append(lines, fmt.Sprintf("[%s]", sec))
		for key, val := range opts {
			lines = append(lines, fmt.Sprintf("%s = %s", key, val))
		}
	}
	return strings.Join(lines, "\n")
}
