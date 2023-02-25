package rc

import (
	"errors"
	"fmt"
	"github.com/charmbracelet/log"
	"os"
	"path/filepath"
	"strings"
)

type RC struct {
	path  string
	Files []File
}

func New(rcFolderPath *string) (rc RC, err error) {
	if rcFolderPath != nil {
		rc.path = *rcFolderPath
	} else {
		execPath, execErr := os.Getwd()
		if execErr != nil {
			err = errors.New("error getting execution path")
			return
		}

		rc.path = execPath
	}

	rc.path = fmt.Sprintf("%s/.semankit", rc.path)
	if fileInfo, statErr := os.Stat(rc.path); statErr != nil {
		if os.IsExist(statErr) {
			err = errors.New(".semankit folder is missing")
		}

		if os.IsPermission(statErr) {
			err = errors.New("insufficient permissions to access .semankit folder")
		}

		return
	} else {
		if !fileInfo.IsDir() {
			err = errors.New(".semankit must be folder")

			return
		}
	}

	if err != nil {
		return RC{}, err
	}

	files, err := os.ReadDir(rc.path)
	if err != nil {
		log.Fatal(err)
	}

	for _, cursor := range files {
		if cursor.IsDir() {
			continue
		}

		filePath := fmt.Sprintf("%s/%s", rc.path, cursor.Name())
		if content, readErr := os.ReadFile(filePath); readErr == nil {
			if file, newFileErr := NewFile(content); newFileErr == nil {
				file.Branch = strings.TrimSuffix(cursor.Name(), filepath.Ext(cursor.Name()))
				rc.Files = append(rc.Files, file)
			}
		}
	}

	return
}

func (receiver RC) FindConfOfBranch(branchName string) (file *File) {
	for _, cursor := range receiver.Files {
		if cursor.Branch != branchName {
			continue
		}

		file = &cursor
		break
	}

	return
}
