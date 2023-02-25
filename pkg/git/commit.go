package git

import "fmt"

type Commit struct {
	Hash    string
	Message string
}

func (receiver Commit) String() string {
	return fmt.Sprintf("%s %s", receiver.Hash, receiver.Message)
}
