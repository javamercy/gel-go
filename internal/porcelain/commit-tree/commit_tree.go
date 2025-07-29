package commit_tree

import (
	"encoding/hex"
	"fmt"
	"gel/internal/core/object"
	"gel/internal/core/repository"
	"gel/pkg/constant"
	"os"
	"strings"
	"time"
)

func CommitTree(repo *repository.Repository, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: gel commit-tree <tree-sha> [-p <parent>]... -m <message>")
	}

	hexTreeHash := args[0]
	tree, err := repo.GetObject(hexTreeHash)
	if err != nil {
		return fmt.Errorf("error retrieving tree: %v", err)
	}

	if tree.Type() != constant.GEL_OBJECT_TYPE_TREE {
		return fmt.Errorf("object is not a tree")
	}

	var message string
	var parentHashes [][]byte

	for i := 1; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "-p", "--parent":
			if i+1 >= len(args) {
				return fmt.Errorf("option %s requires an argument", arg)
			}
			parentHash, err := hex.DecodeString(args[i+1])
			if err != nil {
				return fmt.Errorf("error decoding parent hash: %v", err)
			}
			parentHashes = append(parentHashes, parentHash)
			i++
		case "-m", "--message":
			if i+1 >= len(args) {
				return fmt.Errorf("option %s requires an argument", arg)
			}
			message = args[i+1]
			i++
		default:
			if strings.HasPrefix(arg, "--message=") {
				message = strings.TrimPrefix(arg, "--message=")
			} else if strings.HasPrefix(arg, "--parent=") {
				hashStr := strings.TrimPrefix(arg, "--parent=")
				parentHash, err := hex.DecodeString(hashStr)
				if err != nil {
					return fmt.Errorf("error decoding parent hash: %v", err)
				}
				parentHashes = append(parentHashes, parentHash)
			} else {
				return fmt.Errorf("unknown argument: %s", arg)
			}
		}
	}

	if message == "" {
		return fmt.Errorf("commit message is required (-m or --message)")
	}

	authorName := os.Getenv("GEL_AUTHOR_NAME")
	if authorName == "" {
		authorName = "Gel User"
	}
	authorEmail := os.Getenv("GEL_AUTHOR_EMAIL")
	if authorEmail == "" {
		authorEmail = "gel@example.com"
	}

	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	timezone := "+0000"

	author := object.Author{
		Name:      authorName,
		Email:     authorEmail,
		Timestamp: timestamp,
		Timezone:  timezone,
	}

	committer := object.Committer{
		Name:      authorName,
		Email:     authorEmail,
		Timestamp: timestamp,
		Timezone:  timezone,
	}

	commit := object.NewCommit(author, committer, hexTreeHash, parentHashes, message)

	commitHash, err := repo.SaveObject(commit)
	if err != nil {
		return fmt.Errorf("error saving commit: %v", err)
	}

	fmt.Fprintf(os.Stdout, "%s\n", hex.EncodeToString(commitHash))
	return nil
}
