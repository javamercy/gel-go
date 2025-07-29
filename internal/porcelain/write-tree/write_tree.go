package write_tree

import (
	"encoding/hex"
	"fmt"
	"gel/internal/core/repository"
	"gel/internal/porcelain/add"
	"os"
)

func WriteTree(repo *repository.Repository, directory string) error {
	hash, err := add.Directory(repo, directory)
	if err != nil {
		return fmt.Errorf("error writing tree: %v", err)
	}
	fmt.Fprintf(os.Stdout, "%v\n", hex.EncodeToString(hash))
	return nil
}
