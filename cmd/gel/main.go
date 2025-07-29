package main

import (
	"encoding/hex"
	"fmt"
	"gel/constant"
	"gel/internal/core/object"
	"gel/internal/core/repository"
	"gel/internal/plumbing/storage"
	"gel/internal/porcelain/add"
	"gel/internal/porcelain/init"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: gel <command> [<args>...]\n")
		os.Exit(1)
	}

	command := os.Args[1]
	if command == "init" {
		err := init_gel.Init()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error initializing repository: %s\n", err)
			os.Exit(1)
		}
		return
	}
	repo := repository.NewRepository(storage.NewFilesystem())
	switch command {
	case "cat-file":
		switch argument := os.Args[2]; argument {
		case "-p":
			hash, err := hex.DecodeString(os.Args[3])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error decoding hash: %s\n", err)
				os.Exit(1)
			}
			obj, err := repo.GetObject(hash)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error retrieving object: %s\n", err)
				os.Exit(1)
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
			default:
				fmt.Printf("Unknown object type: %s\n", obj.Type())
			}
		}
	case "hash-object":
		switch argument := os.Args[2]; argument {
		case "-w":
			filepath := os.Args[3]
			hash, err := add.SaveFile(repo, filepath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error saving file: %s\n", err)
				os.Exit(1)
			}
			fmt.Fprintf(os.Stdout, "%s\n", hex.EncodeToString(hash))
		}
	case "ls-tree":
		switch argument := os.Args[2]; argument {
		case "--name-only":
			hash, err := hex.DecodeString(os.Args[3])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error decoding hash: %s\n", err)
				os.Exit(1)
			}
			obj, err := repo.GetObject(hash)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error retrieving object: %s\n", err)
			}

			if obj.Type() != constant.GEL_OBJECT_TYPE_TREE {
				fmt.Fprintf(os.Stderr, "Object is not a tree\n")
				os.Exit(1)
			}

			tree := obj.(*object.Tree)
			for _, entry := range tree.Entries {
				fmt.Println(entry.Name)
			}
		}
	case "write-tree":
		workingDirectory, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting current working directory: %s\n", err)
			os.Exit(1)
		}
		hash, err := add.SaveDirectory(repo, workingDirectory)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing tree: %s\n", err)
		}

		fmt.Fprintf(os.Stdout, "%s\n", hex.EncodeToString(hash))
	default:
		fmt.Fprintf(os.Stderr, "Unknown command %s\n", os.Args[2])
		os.Exit(1)
	}
}
