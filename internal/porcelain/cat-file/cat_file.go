package cat_file

import "gel/internal/core/repository"

type Option string

const (
	PrettyPrint = Option("-p")
)

type CatFile struct {
	repository repository.Repository
}

func NewCatFile(repository repository.Repository) *CatFile {
	return &CatFile{
		repository,
	}
}

func (catFile *CatFile) PrettyPrint(hash []byte) (string, error) {
	_, err := catFile.repository.GetObject(hash)

	if err != nil {
		return "", err
	}

	return "", nil
}
