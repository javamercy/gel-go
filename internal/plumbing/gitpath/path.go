package gitpath

import (
	"encoding/hex"
	"errors"
	"gel/constant"
	"os"
	"path/filepath"
	"sync"
)

var (
	gelPath     string
	gelPathOnce sync.Once
	pathError   error
)

func ensureGelPath() error {
	gelPathOnce.Do(func() {
		currentPath, err := os.Getwd()
		if err != nil {
			pathError = err
			return
		}
		gelPath, pathError = FindGelPath(currentPath)
	})
	return pathError
}
func FindGelPath(startPath string) (string, error) {
	currentPath := startPath
	for {
		gelDir := filepath.Join(currentPath, ".gel")
		info, err := os.Stat(gelDir)
		if err == nil && info.IsDir() {
			return filepath.Clean(gelDir), nil

		}
		parentPath := filepath.Dir(currentPath)
		if parentPath == currentPath {
			return "", errors.New(constant.ERR_GIT_NOT_REPOSITORY)
		}
		currentPath = parentPath
	}
}

func GetObjectPath(hash []byte) (string, error) {
	if err := ensureGelPath(); err != nil {
		return "", err
	}
	hexHash := hex.EncodeToString(hash)
	return filepath.Join(gelPath, "objects", hexHash[:2], hexHash[2:]), nil
}

func GetObjectsPath() (string, error) {
	if err := ensureGelPath(); err != nil {
		return "", err
	}
	return filepath.Join(gelPath, "objects"), nil
}
