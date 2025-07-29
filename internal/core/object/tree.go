package object

import (
	"errors"
	"gel/pkg/constant"
	"strconv"
)

type EntryType int

const (
	BLOB_ENTRY EntryType = iota
	TREE_ENTRY
	COMMIT_ENTRY
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
		for data[i] != constant.GEL_FIELD_DELIMITER {
			i++
		}
		mode := string(data[modeStart:i])
		i++

		nameStart := i
		for data[i] != constant.GEL_OBJECT_DELIMITER {
			i++
		}
		name := string(data[nameStart:i])
		i++

		hash := data[i : i+20]
		i += 20

		var entryType EntryType
		switch mode {
		case constant.GEL_OBJECT_MODE_BLOB:
			entryType = BLOB_ENTRY
		case constant.GEL_OBJECT_MODE_TREE:
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
	return constant.GEL_OBJECT_TYPE_TREE
}

func (tree *Tree) Size() int {
	serialized := tree.serializeEntries()
	return len(serialized)
}

func (tree *Tree) Serialize() []byte {

	serialized := tree.serializeEntries()
	header := constant.GEL_OBJECT_HEADER_TREE + strconv.Itoa(len(serialized)) + string(constant.GEL_OBJECT_DELIMITER)
	return append([]byte(header), serialized...)
}

func (tree *Tree) Deserialize(data []byte) error {

	i := 0
	for i < len(data) {
		modeStart := i
		for i != len(data) && data[i] != constant.GEL_FIELD_DELIMITER {
			i++
		}
		if i == len(data) {
			return errors.New("error parsing tree entry mode")
		}
		mode := string(data[modeStart:i])
		i++

		nameStart := i
		for i != len(data) && data[i] != constant.GEL_OBJECT_DELIMITER {
			i++
		}

		if i == len(data) {
			return errors.New("error parsing tree entry name")
		}
		name := string(data[nameStart:i])
		i++

		hash := data[i : i+20]
		i += 20

		var entryType EntryType
		switch mode {
		case constant.GEL_OBJECT_MODE_BLOB:
			entryType = BLOB_ENTRY
		case constant.GEL_OBJECT_MODE_TREE:
			entryType = TREE_ENTRY
		case constant.GEL_OBJECT_MODE_COMMIT:
			entryType = COMMIT_ENTRY
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

	return nil

}

func (tree *Tree) serializeEntries() []byte {
	var serialized []byte
	for _, entry := range tree.Entries {
		line := []byte(entry.Mode)
		line = append(line, constant.GEL_FIELD_DELIMITER)
		line = append(line, []byte(entry.Name)...)
		line = append(line, constant.GEL_OBJECT_DELIMITER)
		line = append(line, entry.Hash...)
		serialized = append(serialized, line...)
	}
	return serialized
}
