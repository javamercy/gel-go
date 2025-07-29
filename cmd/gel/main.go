package main

import (
	"fmt"
	"gel/internal/core/repository"
	"gel/internal/plumbing/storage"
	"gel/internal/porcelain/add"
	"gel/internal/porcelain/cat-file"
	"gel/internal/porcelain/commit-tree"
	"gel/internal/porcelain/hash-object"
	"gel/internal/porcelain/init"
	"gel/internal/porcelain/ls-tree"
	"gel/internal/porcelain/write-tree"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: gel <command> [<args>...]\n")
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	if command == "init" {
		err := init_gel.Init()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error initializing repository: %s\n", err)
			os.Exit(1)
		}
		return
	}

	repo := repository.NewRepository(storage.NewFilesystem())

	var err error
	switch command {
	case "add":
		err = add.Add(repo, args)
	case "cat-file":
		err = cat_file.CatFile(repo, args)
	case "hash-object":
		err = hash_object.HashObject(repo, args)
	case "ls-tree":
		err = ls_tree.LsTree(repo, args)
	case "write-tree":
		directory, _ := os.Getwd()
		err = write_tree.WriteTree(repo, directory)
	case "commit-tree":
		err = commit_tree.CommitTree(repo, args)
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
