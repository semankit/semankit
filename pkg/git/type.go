package git

type CommitType uint

const (
	Major CommitType = iota + 1
	Minor
	Patch
)
