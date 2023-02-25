package rc

import (
	"gopkg.in/yaml.v3"
)

type File struct {
	Branch  string  `yaml:"-"`
	Release Release `yaml:"release,omitempty"`
}

func NewFile(content []byte) (File, error) {
	var file File
	if err := yaml.Unmarshal(content, &file); err != nil {
		return file, err
	}

	return file, nil
}
