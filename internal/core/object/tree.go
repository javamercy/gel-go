package object

import (
	"gel/constant"
	"strconv"
)

type EntryType int

const (
	BLOB_ENTRY EntryType = iota
	TREE_ENTRY
	UNKNOWN_ENTRY
)

type TreeEntry struct {
	Mode string
	Name string
	Hash []byte
	Type EntryType
}

type Tree struct {
	Entries []TreeEntry
}

func NewTree() *Tree {
	return &Tree{
		Entries: []TreeEntry{},
	}
}

func NewTreeFromEntriesData(data []byte) *Tree {
	tree := NewTree()
	i := 0
	for i < len(data) {

		modeStart := i
		for data[i] != constant.GIT_FIELD_DELIMITER {
			i++
		}
		mode := string(data[modeStart:i])
		i++

		nameStart := i
		for data[i] != constant.GIT_OBJECT_DELIMITER {
			i++
		}
		name := string(data[nameStart:i])
		i++

		hash := data[i : i+20]
		i += 20

		var entryType EntryType
		switch mode {
		case constant.GIT_OBJECT_MODE_BLOB:
			entryType = BLOB_ENTRY
		case constant.GIT_OBJECT_MODE_TREE:
			entryType = TREE_ENTRY
		default:
			entryType = UNKNOWN_ENTRY
		}
		tree.Entries = append(tree.Entries, TreeEntry{
			Name: name,
			Mode: mode,
			Hash: hash,
			Type: entryType,
		})
	}
	return tree
}

func (tree *Tree) AddEntry(entry TreeEntry) {
	tree.Entries = append(tree.Entries, entry)
}

func (tree *Tree) Type() string {
	return constant.GIT_OBJECT_TYPE_TREE
}

func (tree *Tree) Size() int {
	serialized := tree.serializeEntries()
	return len(serialized)
}

func (tree *Tree) Serialize() []byte {

	serialized := tree.serializeEntries()
	header := constant.GIT_OBJECT_HEADER_TREE + strconv.Itoa(len(serialized)) + string(constant.GIT_OBJECT_DELIMITER)
	return append([]byte(header), serialized...)
}

func (tree *Tree) Deserialize(hash []byte) ([]byte, error) {

	return nil, nil
}

func (tree *Tree) serializeEntries() []byte {
	var serialized []byte
	for _, entry := range tree.Entries {
		line := []byte(entry.Mode)
		line = append(line, constant.GIT_FIELD_DELIMITER)
		line = append(line, []byte(entry.Name)...)
		line = append(line, constant.GIT_OBJECT_DELIMITER)
		line = append(line, entry.Hash...)
		serialized = append(serialized, line...)
	}
	return serialized
}
