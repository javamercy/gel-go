package add

import (
	"gel/constant"
	"gel/internal/core/object"
	"gel/internal/core/repository"
	"os"
	"path/filepath"
)

func SaveFile(repository *repository.Repository, path string) ([]byte, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	blob := object.NewBlob(data)
	return repository.SaveObject(blob)
}

func SaveDirectory(repository *repository.Repository, path string) ([]byte, error) {

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	tree := object.NewTree()
	for _, entry := range entries {
		entryPath := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			hash, err := SaveDirectory(repository, entryPath)
			if err != nil {
				return nil, err
			}
			treeEntry := object.TreeEntry{
				Mode: constant.GEL_OBJECT_MODE_TREE,
				Name: entry.Name(),
				Hash: hash,
				Type: object.TREE_ENTRY,
			}
			tree.AddEntry(treeEntry)
		} else {
			hash, err := SaveFile(repository, entryPath)
			if err != nil {
				return nil, err
			}

			treeEntry := object.TreeEntry{
				Mode: constant.GEL_OBJECT_MODE_BLOB,
				Name: entry.Name(),
				Hash: hash,
				Type: object.BLOB_ENTRY,
			}
			tree.AddEntry(treeEntry)
		}
	}

	return repository.SaveObject(tree)

}
