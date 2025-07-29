package cat_file

import (
	"encoding/hex"
	"fmt"
	"gel/internal/core/object"
	"gel/internal/core/repository"
	"gel/pkg/constant"
	"os"
)

func CatFile(repo *repository.Repository, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: gel cat-file <option> <object>")
	}

	switch args[0] {
	case "-p":
		return catFilePrettyPrint(repo, args[1])
	case "-t":
		return catFileType(repo, args[1])
	case "-s":
		return catFileSize(repo, args[1])
	default:
		return fmt.Errorf("unknown option: %s", args[0])
	}
}

func catFilePrettyPrint(repo *repository.Repository, hexHash string) error {

	obj, err := repo.GetObject(hexHash)
	if err != nil {
		return fmt.Errorf("error retrieving object: %v", err)
	}

	switch obj.Type() {
	case constant.GEL_OBJECT_TYPE_BLOB:
		blob := obj.(*object.Blob)
		fmt.Fprintf(os.Stdout, "%s", blob.Data)
	case constant.GEL_OBJECT_TYPE_TREE:
		tree := obj.(*object.Tree)
		for _, entry := range tree.Entries {
			entryType := "blob"
			if entry.Type == object.TREE_ENTRY {
				entryType = "tree"
			}
			fmt.Fprintf(os.Stdout, "%s %s %s\t%s\n", entry.Mode, entryType, hex.EncodeToString(entry.Hash), entry.Name)
		}
	case constant.GEL_OBJECT_TYPE_COMMIT:
		commit := obj.(*object.Commit)
		fmt.Fprintf(os.Stdout, "tree %v\n", commit.HexTreeHash)
		for _, parentHash := range commit.ParentHashes {
			fmt.Fprintf(os.Stdout, "parent %s\n", hex.EncodeToString(parentHash))
		}
		fmt.Fprintf(os.Stdout, "author %s <%s> %s %s\n", commit.Author.Name, commit.Author.Email, commit.Author.Timestamp, commit.Author.Timezone)
		fmt.Fprintf(os.Stdout, "committer %s <%s> %s %s\n", commit.Committer.Name, commit.Committer.Email, commit.Committer.Timestamp, commit.Committer.Timezone)
		fmt.Fprintf(os.Stdout, "\n%s\n", commit.Message)
	default:
		return fmt.Errorf("unknown object type: %s", obj.Type())
	}

	return nil
}

func catFileType(repo *repository.Repository, hexHash string) error {

	obj, err := repo.GetObject(hexHash)
	if err != nil {
		return fmt.Errorf("error retrieving object: %v", err)
	}

	fmt.Println(obj.Type())
	return nil
}

func catFileSize(repo *repository.Repository, hexHash string) error {

	obj, err := repo.GetObject(hexHash)
	if err != nil {
		return fmt.Errorf("error retrieving object: %v", err)
	}

	fmt.Println(obj.Size())
	return nil
}
