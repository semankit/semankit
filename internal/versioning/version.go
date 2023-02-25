package versioning

import (
	gitClient "github.com/bitbreakr/semankit/internal/git"
	"github.com/bitbreakr/semankit/pkg/git"
	"github.com/bitbreakr/semankit/pkg/version"
	"github.com/charmbracelet/log"
	"strings"
)

type Bump struct {
	IsDirty    bool
	repository gitClient.Git
}

func New() Bump {
	return Bump{}
}

func (receiver *Bump) calculateWeight(commit git.Commit, strategy version.Strategy) {
	commit.Message = strings.TrimLeft(commit.Message, " ")
	commit.Message = strings.TrimRight(commit.Message, " ")

	if strings.HasPrefix(commit.Message, "feat:") ||
		strings.HasPrefix(commit.Message, "feat(") {
		if is, err := receiver.repository.IsCommitBreakingChange(commit); err != nil {
			log.Error(err)
		} else {
			if is {
				strategy.UpdateVersion(git.Major)
			} else {
				strategy.UpdateVersion(git.Minor)
			}
		}
	}

	if strings.HasPrefix(commit.Message, "fix:") ||
		strings.HasPrefix(commit.Message, "fix(") {
		if is, err := receiver.repository.IsCommitBreakingChange(commit); err != nil {
			log.Error(err)
		} else {
			if is {
				strategy.UpdateVersion(git.Minor)
			} else {
				strategy.UpdateVersion(git.Patch)
			}
		}
	}
}

func (receiver *Bump) Bump(
	currentVersion *string,
	commits []git.Commit,
	strategy version.Strategy,
) (next string, err error) {
	if currentVersion != nil {
		if currVerErr := strategy.SetCurrentVersion(*currentVersion); currVerErr != nil {
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