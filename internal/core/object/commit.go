package object

import (
	"encoding/hex"
	"fmt"
	"gel/pkg/constant"
	"strconv"
	"strings"
)

type Author struct {
	Name      string
	Email     string
	Timestamp string
	Timezone  string
}

type Committer struct {
	Name      string
	Email     string
	Timestamp string
	Timezone  string
}
type Commit struct {
	Author       Author
	Committer    Committer
	HexTreeHash  string
	ParentHashes [][]byte
	Message      string
}

func NewCommit(author Author, committer Committer, hexTreeHash string, parentHashes [][]byte, message string) *Commit {
	return &Commit{
		Author:       author,
		Committer:    committer,
		HexTreeHash:  hexTreeHash,
		ParentHashes: parentHashes,
		Message:      message,
	}
}

func (commit *Commit) Type() string {
	return constant.GEL_OBJECT_TYPE_COMMIT
}
func (commit *Commit) Size() int {
	//TODO implement me
	panic("implement me")
}

func (commit *Commit) Serialize() []byte {
	content := constant.GEL_COMMIT_FIELD_TREE + commit.HexTreeHash + constant.GEL_COMMIT_MESSAGE_SEPARATOR

	for _, parentHash := range commit.ParentHashes {
		content += constant.GEL_COMMIT_FIELD_PARENT + hex.EncodeToString(parentHash) + constant.GEL_COMMIT_MESSAGE_SEPARATOR
	}

	content += constant.GEL_COMMIT_FIELD_AUTHOR + commit.Author.Name + constant.GEL_FIELD_DELIMITER_STRING +
		"<" + commit.Author.Email + ">" + constant.GEL_FIELD_DELIMITER_STRING +
		commit.Author.Timestamp + constant.GEL_FIELD_DELIMITER_STRING +
		commit.Author.Timezone + constant.GEL_COMMIT_MESSAGE_SEPARATOR

	content += constant.GEL_COMMIT_FIELD_COMMITTER + commit.Committer.Name + constant.GEL_FIELD_DELIMITER_STRING +
		"<" + commit.Committer.Email + ">" + constant.GEL_FIELD_DELIMITER_STRING +
		commit.Committer.Timestamp + constant.GEL_FIELD_DELIMITER_STRING +
		commit.Committer.Timezone + constant.GEL_COMMIT_MESSAGE_SEPARATOR

	content += constant.GEL_COMMIT_MESSAGE_SEPARATOR + commit.Message

	contentBytes := []byte(content)
	header := constant.GEL_OBJECT_HEADER_COMMIT + strconv.Itoa(len(contentBytes)) + string(constant.GEL_OBJECT_DELIMITER)

	return append([]byte(header), contentBytes...)
}

func (commit *Commit) Deserialize(data []byte) error {
	dataStr := string(data)
	lines := strings.Split(dataStr, constant.GEL_COMMIT_MESSAGE_SEPARATOR)

	var messageStartIndex int

	for i, line := range lines {
		if line == "" {
			messageStartIndex = i + 1
			break
		}

		if strings.HasPrefix(line, constant.GEL_COMMIT_FIELD_TREE) {
			hexTreeHash := strings.TrimPrefix(line, constant.GEL_COMMIT_FIELD_TREE)
			commit.HexTreeHash = hexTreeHash

		} else if strings.HasPrefix(line, constant.GEL_COMMIT_FIELD_PARENT) {
			parentHashStr := strings.TrimPrefix(line, constant.GEL_COMMIT_FIELD_PARENT)
			parentHash, err := hex.DecodeString(parentHashStr)
			if err != nil {
				return fmt.Errorf("invalid parent hash: %v", err)
			}
			commit.ParentHashes = append(commit.ParentHashes, parentHash)

		} else if strings.HasPrefix(line, constant.GEL_COMMIT_FIELD_AUTHOR) {
			authorInfo := strings.TrimPrefix(line, constant.GEL_COMMIT_FIELD_AUTHOR)
			err := parsePersonInfo(authorInfo, &commit.Author.Name, &commit.Author.Email, &commit.Author.Timestamp, &commit.Author.Timezone)
			if err != nil {
				return fmt.Errorf("invalid author info: %v", err)
			}

		} else if strings.HasPrefix(line, constant.GEL_COMMIT_FIELD_COMMITTER) {
			committerInfo := strings.TrimPrefix(line, constant.GEL_COMMIT_FIELD_COMMITTER)
			err := parsePersonInfo(committerInfo, &commit.Committer.Name, &commit.Committer.Email, &commit.Committer.Timestamp, &commit.Committer.Timezone)
			if err != nil {
				return fmt.Errorf("invalid committer info: %v", err)
			}
		}
	}

	if messageStartIndex < len(lines) {
		commit.Message = strings.Join(lines[messageStartIndex:], constant.GEL_COMMIT_MESSAGE_SEPARATOR)
	}

	return nil
}

func parsePersonInfo(info string, name, email, timestamp, timezone *string) error {
	emailStart := strings.Index(info, "<")
	emailEnd := strings.Index(info, ">")

	if emailStart == -1 || emailEnd == -1 || emailEnd <= emailStart {
		return fmt.Errorf("invalid person info format")
	}

	*name = strings.TrimSpace(info[:emailStart])
	*email = info[emailStart+1 : emailEnd]

	remaining := strings.TrimSpace(info[emailEnd+1:])
	parts := strings.Fields(remaining)

	if len(parts) != 2 {
		return fmt.Errorf("invalid timestamp/timezone format")
	}

	*timestamp = parts[0]
	*timezone = parts[1]

	return nil
}
