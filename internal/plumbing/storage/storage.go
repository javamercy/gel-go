package storage

import "gel/internal/core/object"

type ObjectStorage interface {
	Get(hexHash string) (object.Object, error)
	Save(object object.Object) ([]byte, error)
	Exists(hexHash string) bool
}
