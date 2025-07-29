package hash_object

import (
	"encoding/hex"
	"fmt"
	"gel/internal/core/repository"
	"gel/internal/porcelain/add"
	"os"
)

func HashObject(repo *repository.Repository, args []string) error {
	if len(args) < 2 {
		fmt.Errorf("usage: gel hash-object <option> <file>")
	}

	switch args[0] {
	case "-w":
		return Write(repo, args[1])
	default:
		return fmt.Errorf("unknown option: %s", args[0])
	}
}

func Write(repo *repository.Repository, filepath string) error {
	hash, err := add.File(repo, filepath)
	if err != nil {
		return fmt.Errorf("error saving file: %v", err)
	}
	fmt.Fprintf(os.Stdout, "%v\n", hex.EncodeToString(hash))
	return nil
}
