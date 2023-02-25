package git

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/semankit/semankit/pkg/git"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Git struct {
	path string
}

func New(repoPath *string) (git Git) {
	if repoPath != nil {
		git.path = *repoPath
	} else {
		execPath, _ := os.Getwd()
		git.path = execPath
	}

	return
}

func (receiver Git) IsInstalled() (isInstalled bool, err error) {
	cmd := exec.Command("which", "git")
	_, cmdErr := cmd.Output()
	if cmdErr != nil {
		if cmdErr.(*exec.ExitError).ExitCode() == 1 {
			err = errors.New("error, git is not installed")
		} else {
			err = errors.New("error, could not determine if git is installed")
		}

		return
	}

	isInstalled = true

	return
}

// Version returns the Git versioning
func (receiver Git) Version() (version string, err error) {
	version = "undetermined"
	cmd := exec.Command("git", "--versioning")
	cmd.Dir = receiver.path
	output, cmdErr := cmd.Output()
	if cmdErr != nil {
		err = errors.New("could not determine git versioning")
		return
	}

	buffer := strings.Split(string(output), "\n")
	if len(buffer) == 1 {
		version = buffer[0]
	}

	if len(buffer) == 1 || len(buffer) == 2 {
		version = buffer[0]
	}

	return
}

// ListCommitsFromTag list all commits since last tag
func (receiver Git) ListCommitsFromTag(tag string) []git.Commit {
	cmd := exec.Command("git", "log", fmt.Sprintf("%s...HEAD", tag), "--pretty=oneline")
	cmd.Dir = receiver.path
	out, err := cmd.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err)
	}
	// otherwise, print the output from running the command

	commits := make([]git.Commit, 0)
	buffer := strings.Split(string(out), "\n")

	for _, cursor := range buffer {
		totalChar := len(cursor)
		if totalChar == 0 {
			continue
		}

		commit := git.Commit{}
		commit.Hash = cursor[0:40]
		commit.Message = cursor[40:totalChar]
		commits = append(commits, commit)
	}

	return commits
}

// ListCommits list all commits, can be limited by depth
func (receiver Git) ListCommits(depth uint8) []git.Commit {
	arg := []string{
		"log",
		"--pretty=oneline",
	}

	if depth != 0 {
		arg = append(arg, "-n", strconv.Itoa(int(depth)))
	}

	cmd := exec.Command("git", arg...)
	cmd.Dir = receiver.path
	out, err := cmd.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err)
	}
	// otherwise, print the output from running the command

	commits := make([]git.Commit, 0)
	buffer := strings.Split(string(out), "\n")

	for _, cursor := range buffer {
		totalChar := len(cursor)
		if totalChar == 0 {
			continue
		}

		commit := git.Commit{}
		commit.Hash = cursor[0:40]
		commit.Message = cursor[40:totalChar]
		commits = append(commits, commit)
	}

	return commits
}

// ListTags list all repository tags
func (receiver Git) ListTags() []string {
	cmd := exec.Command("git", "tag")
	cmd.Dir = receiver.path
	out, err := cmd.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err)
	}
	// otherwise, print the output from running the command

	tags := strings.Split(string(out), "\n")
	if len(tags[len(tags)-1]) == 0 {
		tags = tags[:len(tags)-1]
	}

	return tags
}

// CurrentBranch Return checked-out branch
func (receiver Git) CurrentBranch() string {
	cmd := exec.Command("git", "branch", "--show-current")
	cmd.Dir = receiver.path
	out, err := cmd.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err)
	}
	// otherwise, print the output from running the command

	tags := strings.Split(string(out), "\n")
	if len(tags[len(tags)-1]) == 0 {
		tags = tags[:len(tags)-1]
	}

	return tags[0]
}

// DoesBranchExist check if a given branch exist
func (receiver Git) DoesBranchExist(branchName string) (exist bool, err error) {
	cmd := exec.Command("git", "branch", "-l")
	cmd.Dir = receiver.path
	buffer, err := cmd.Output()
	buffer = bytes.TrimSuffix(buffer, []byte{10})
	buffer = bytes.TrimPrefix(buffer, []byte{42})
	buffer = bytes.TrimPrefix(buffer, []byte{32})

	for _, cursor := range strings.Split(string(buffer), "\n") {
		if cursor != branchName {
			continue
		}

		exist = true
		break
	}

	return
}

func (receiver Git) IsCommitBreakingChange(commit git.Commit) (is bool, err error) {
	cmd := exec.Command("git", "show", "-s", commit.Hash)
	cmd.Dir = receiver.path
	output, cmdErr := cmd.Output()
	if cmdErr != nil {
		err = errors.New("could not determine if breaking")
		return
	}

	buffer := string(output)
	if strings.Contains(buffer, "BREAKING CHANGE") {
		is = true
	}

	return
}
