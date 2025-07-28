package object

import (
	"gel/constant"
	"strconv"
)

type Blob struct {
	Data []byte
}

func NewBlob(data []byte) *Blob {
	return &Blob{
		Data: data,
	}
}

func (blob *Blob) Type() string {
	return constant.GIT_OBJECT_TYPE_BLOB
}

func (blob *Blob) Size() int {
	return len(blob.Data)
}

func (blob *Blob) Serialize() []byte {
	header := constant.GIT_OBJECT_HEADER_BLOB + strconv.Itoa(blob.Size()) + string(constant.GIT_OBJECT_DELIMITER)

	return append([]byte(header), blob.Data...)
}

func (blob *Blob) Deserialize(hash []byte) ([]byte, error) {
	return nil, nil
}
