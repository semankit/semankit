package commit

type Type uint

const (
	Major Type = iota + 1
	Minor
	Patch
)
