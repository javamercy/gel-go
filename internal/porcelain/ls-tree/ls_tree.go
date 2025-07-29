package ls_tree

import (
	"fmt"
	"gel/internal/core/object"
	"gel/internal/core/repository"
	"gel/pkg/constant"
)

func LsTree(repo *repository.Repository, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: gel ls-tree <option> <tree-sha>")
	}

	switch args[0] {
	case "--name-only":
		return lsTreeNameOnly(repo, args[1])
	default:
		return fmt.Errorf("unknown option: %s", args[0])
	}
}

func lsTreeNameOnly(repo *repository.Repository, hexHash string) error {

	obj, err := repo.GetObject(hexHash)
	if err != nil {
		return fmt.Errorf("error retrieving object: %v", err)
	}

	if obj.Type() != constant.GEL_OBJECT_TYPE_TREE {
		return fmt.Errorf("object is not a tree")
	}

	tree := obj.(*object.Tree)
	for _, entry := range tree.Entries {
		fmt.Println(entry.Name)
	}

	return nil
}
