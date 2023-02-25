package rc

import "fmt"

type Release struct {
	Suffix *string `yaml:"suffix,omitempty"`
}

// HasSuffix returns true if release name has a suffix
func (receiver Release) HasSuffix() bool {
	return receiver.Suffix != nil
}

func (receiver Release) AppendSuffix(nextVersion *string) {
	*nextVersion = fmt.Sprintf("%s-%s", *nextVersion, *receiver.Suffix)
}
