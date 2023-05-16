package version

import (
	"github.com/semankit/semankit/pkg/commit"
)

type Strategy interface {
	Next() string
	SetCurrentVersion(version string) error
	InitVersion()
	UpdateVersion(commitType commit.Type)
}
