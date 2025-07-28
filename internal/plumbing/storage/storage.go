package storage

import "gel/internal/core/object"

type ObjectStorage interface {
	Get(hash []byte) (object.Object, error)
	Save(object object.Object) ([]byte, error)
	Exists(hash []byte) bool
}
