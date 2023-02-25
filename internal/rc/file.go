package rc

import (
	"gopkg.in/yaml.v3"
)

type File struct {
	Branch     string `yaml:"-"`
	Prerelease bool   `yaml:"prerelease"`
}

func NewFile(content []byte) (File, error) {
	var file File
	if err := yaml.Unmarshal(content, &file); err != nil {
		return file, err
	}

	return file, nil
}
