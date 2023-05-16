package main

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/semankit/karmic"
	"github.com/semankit/semankit/internal/rc"
	"github.com/semankit/semankit/internal/versioning"
	"github.com/semankit/semankit/internal/versioning/strategy"
	"os"
)

func main() {
	var err error

	git := karmic.New(nil)
	// Guard to check if git is installed
	if isGitInstalled, err := git.IsInstalled(); err != nil || !isGitInstalled {
		if err != nil {
			log.Error(err)
		}
		if !isGitInstalled {
			log.Error("error, git is not installed... Exiting!")
		}
		os.Exit(1)
	}

	rules, err := rc.New(nil)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	currentBranch := git.CurrentBranch()
	if branchCfn := rules.FindConfOfBranch(string(currentBranch)); branchCfn == nil {
		log.Info(fmt.Sprintf("no conf found for %s, skipping current execution", currentBranch))
		os.Exit(0)
	} else {
		var nextVersion string
		version := versioning.New()
		tags := git.Tags()
		commits := git.List(0)

		log.Info(fmt.Sprintf("found %d tag(s)", len(tags)))

		if len(tags) == 0 && len(commits) > 0 {
			log.Info(fmt.Sprintf("found %d commit(s)", len(commits)))
			nextVersion, _ = version.Bump(nil, commits, strategy.Default())
		}

		if len(tags) > 0 {
			log.Info(fmt.Sprintf("last tag found is %s", tags[0]))
			commits = tags[0].Commits()
			log.Info(fmt.Sprintf("found %d new commit(s) since last release", len(commits)))
			nextVersion, _ = version.Bump(&tags[0], commits, strategy.Default())
		}

		if !version.IsDirty {
			log.Info("no commit meets the prerequisites to bump the current version")
			os.Exit(0)
		}

		if branchCfn.Release.HasSuffix() {
			branchCfn.Release.AppendSuffix(&nextVersion)
		}

		fmt.Print(nextVersion)
	}

	os.Exit(0)
}
