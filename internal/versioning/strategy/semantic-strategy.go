package strategy

import (
	"errors"
	"fmt"
	"github.com/semankit/semankit/pkg/commit"
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

func (s *Semantic) hasReachCap(commitType commit.Type) bool {
	capLimit := 255
	switch commitType {
	case commit.Minor:
		return s.minor < capLimit
	case commit.Patch:
		return s.patch < capLimit
	}

	return false
}

func (s *Semantic) UpdateVersion(commitType commit.Type) {
	switch commitType {
	case commit.Major:
		{
			s.major++
		}
	case commit.Minor:
		{
			if s.hasReachCap(commit.Minor) {
				s.minor++
			} else {
				s.minor = 0
				s.patch++
			}
		}
	case commit.Patch:
		{
			if s.hasReachCap(commit.Patch) {
				s.patch++
			} else {
				s.patch = 0
				s.minor++
			}
		}
	}
}

func (s Semantic) Next() string {
	return fmt.Sprintf("v%d.%d.%d", s.major, s.minor, s.patch)
}
