package repository

import (
	"gel/internal/core/object"
	"gel/internal/plumbing/storage"
)

type Repository struct {
	storage storage.ObjectStorage
}

func NewRepository(storage storage.ObjectStorage) *Repository {
	return &Repository{
		storage,
	}
}

func (repository *Repository) GetObject(hash []byte) (object.Object, error) {
	return repository.storage.Get(hash)
}

func (repository *Repository) SaveObject(object object.Object) ([]byte, error) {
	return repository.storage.Save(object)
}
