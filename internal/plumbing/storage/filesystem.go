package storage

import (
	"encoding/hex"
	"errors"
	"gel/internal/core/object"
	"gel/internal/plumbing/gel-path"
	"gel/pkg/compression"
	"gel/pkg/constant"
	"gel/pkg/hashing"
	"os"
	fp "path/filepath"
	"strings"
)

type Filesystem struct {
	objectsPath string
}

func NewFilesystem() *Filesystem {
	path, err := gel_path.GetObjectsPath()
	if err != nil {
		panic(err)
	}
	return &Filesystem{
		objectsPath: path,
	}
}

func (filesystem *Filesystem) Get(hexHash string) (object.Object, error) {
	fullPath := fp.Join(filesystem.objectsPath, hexHash[:2], hexHash[2:])

	data, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}

	decompressedData, err := compression.DecompressZlib(data)
	if err != nil {
		return nil, err
	}

	return parseObject(decompressedData)
}

func (filesystem *Filesystem) Save(object object.Object) ([]byte, error) {
	data := object.Serialize()
	hash := hashing.ComputeSha1Hash(data)
	compressedData, err := compression.CompressZlib(data)
	if err != nil {
		return nil, err
	}

	hexHash := hex.EncodeToString(hash)
	dirPath := fp.Join(filesystem.objectsPath, hexHash[:2])

	if err := os.MkdirAll(dirPath, constant.GEL_DIRECTORY_PERMISSIONS); err != nil {
		return nil, err
	}

	filePath := fp.Join(dirPath, hexHash[2:])
	if err := os.WriteFile(filePath, compressedData, constant.GEL_REGULAR_FILE_PERMISSIONS); err != nil {
		return nil, err
	}

	return hash, nil
}

func (filesystem *Filesystem) Exists(hexHash string) bool {
	filePath := fp.Join(filesystem.objectsPath, hexHash[:2], hexHash[2:])
	_, err := os.Stat(filePath)
	return err == nil
}

func parseObject(decompressedData []byte) (object.Object, error) {
	objectDelimiterIndex := -1
	for i, b := range decompressedData {
		if b == constant.GEL_OBJECT_DELIMITER {
			objectDelimiterIndex = i
			break
		}
	}

	if objectDelimiterIndex == -1 {
		return nil, errors.New("invalid git object")
	}

	header := string(decompressedData[:objectDelimiterIndex])
	content := decompressedData[objectDelimiterIndex+1:]

	headerParts := strings.SplitN(header, " ", 2)
	if len(headerParts) != 2 {
		return nil, errors.New("invalid git object header")
	}

	objectType := headerParts[0]

	switch objectType {
	case constant.GEL_OBJECT_TYPE_BLOB:
		blob := &object.Blob{}
		err := blob.Deserialize(content)
		if err != nil {
			return nil, err
		}
		return blob, nil
	case constant.GEL_OBJECT_TYPE_TREE:
		tree := &object.Tree{}
		err := tree.Deserialize(content)
		if err != nil {
			return nil, err
		}
		return tree, nil
	case constant.GEL_OBJECT_TYPE_COMMIT:
		commit := &object.Commit{}
		err := commit.Deserialize(content)
		if err != nil {
			return nil, err
		}
		return commit, nil
	default:
		return nil, errors.New("invalid git object type")
	}
}
