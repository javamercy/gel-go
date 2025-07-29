package add

import (
	"fmt"
	"gel/internal/core/object"
	"gel/internal/core/repository"
	"gel/pkg/constant"
	"os"
	"path/filepath"
)

func Add(repo *repository.Repository, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: gel add <file(s)>")
	}

	for _, path := range args {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return fmt.Errorf("pathspec '%s' did not match any files", path)
		}

		info, err := os.Stat(path)
		if err != nil {
			return fmt.Errorf("error accessing '%s': %v", path, err)
		}

		if info.IsDir() {
			_, err = Directory(repo, path)
		} else {
			_, err = File(repo, path)
		}

		if err != nil {
			return fmt.Errorf("error adding '%s': %v", path, err)
		}
	}
	return nil
}

func File(repository *repository.Repository, path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	blob := object.NewBlob(data)
	return repository.SaveObject(blob)
}

func Directory(repository *repository.Repository, path string) ([]byte, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	tree := object.NewTree()
	for _, entry := range entries {
		entryPath := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			hash, err := Directory(repository, entryPath)
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
			hash, err := File(repository, entryPath)
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
