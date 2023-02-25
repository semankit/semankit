package strategy

import (
	"errors"
	"fmt"
	"github.com/semankit/semankit/pkg/git"
	"strconv"
	"strings"
	"unicode"
)

type Semantic struct {
	major int
	minor int
	patch int
}

func Default() *Semantic {
	return &Semantic{}
}

func (s *Semantic) InitVersion() {
	s.major = 0
	s.minor = 0
	s.patch = 0
}

func (s *Semantic) SetCurrentVersion(tag string) error {
	var err error
	var major, minor, patch int

	if unicode.IsLetter(rune(tag[0])) {
		tag = strings.TrimPrefix(tag, tag[0:1])
	}

	version := strings.Split(tag, ".")
	if len(version) != 3 {
		return errors.New("error, tag format is invalid")
	}

	major, err = strconv.Atoi(version[0])
	minor, err = strconv.Atoi(version[1])
	patch, err = strconv.Atoi(version[2])

	if err != nil {
		return errors.New("error, could not parse last tag")
	}

	s.major = major
	s.minor = minor
	s.patch = patch

	return nil
}

func (s *Semantic) UpdateVersion(commitType git.CommitType) {
	switch commitType {
	case git.Major:
		{
			s.major++
		}
	case git.Minor:
		{
			s.minor++
		}
	case git.Patch:
		{
			s.patch++
		}
	}
}

func (s Semantic) Next() string {
	return fmt.Sprintf("v%d.%d.%d", s.major, s.minor, s.patch)
}
