package versioning

import (
	"github.com/charmbracelet/log"
	"github.com/semankit/karmic"
	commitType "github.com/semankit/semankit/pkg/commit"
	"github.com/semankit/semankit/pkg/version"
	"strings"
)

type Bump struct {
	IsDirty bool
}

func New() Bump {
	return Bump{}
}

func (receiver *Bump) calculateWeight(commit karmic.Commit, strategy version.Strategy) {
	commit.Message = strings.TrimLeft(commit.Message, " ")
	commit.Message = strings.TrimRight(commit.Message, " ")

	if strings.HasPrefix(commit.Message, "feat:") ||
		strings.HasPrefix(commit.Message, "feat(") {
		if is, err := commit.IsBreakingChange(); err != nil {
			log.Error(err)
		} else {
			if is {
				strategy.UpdateVersion(commitType.Major)
			} else {
				strategy.UpdateVersion(commitType.Minor)
			}
		}

		receiver.IsDirty = true
	}

	if strings.HasPrefix(commit.Message, "fix:") ||
		strings.HasPrefix(commit.Message, "fix(") {
		if is, err := commit.IsBreakingChange(); err != nil {
			log.Error(err)
		} else {
			if is {
				strategy.UpdateVersion(commitType.Minor)
			} else {
				strategy.UpdateVersion(commitType.Patch)
			}
		}

		receiver.IsDirty = true
	}
}

func (receiver *Bump) Bump(
	currentVersion *karmic.Tag,
	commits []karmic.Commit,
	strategy version.Strategy,
) (next string, err error) {
	if currentVersion != nil {
		if currVerErr := strategy.SetCurrentVersion(currentVersion.String()); currVerErr != nil {
			log.Error(currVerErr)
			strategy.InitVersion()
		}
	} else {
		strategy.InitVersion()
	}

	for _, cursor := range commits {
		receiver.calculateWeight(cursor, strategy)
	}

	next = strategy.Next()

	return
}
