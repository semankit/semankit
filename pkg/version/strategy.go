package version

import (
	"github.com/semankit/semankit/pkg/git"
)

type Strategy interface {
	Next() string
	SetCurrentVersion(version string) error
	InitVersion()
	UpdateVersion(commitType git.CommitType)
}
