package version

import (
	"github.com/bitbreakr/semankit/pkg/git"
)

type Strategy interface {
	Next() string
	SetCurrentVersion(version string) error
	InitVersion()
	UpdateVersion(commitType git.CommitType)
}
