package init_gel

import (
	"gel/constant"
	"os"
)

func Init() error {
	for _, dir := range []string{".gel", ".gel/objects", ".gel/refs"} {
		if err := os.MkdirAll(dir, constant.GIT_EXECUTABLE_FILE_PERMISSIONS); err != nil {
			return err
		}
	}
	headFileContents := []byte("ref: refs/heads/main\n")
	if err := os.WriteFile(".gel/HEAD", headFileContents, constant.GIT_REGULAR_FILE_PERMISSIONS); err != nil {
		return err
	}
	return nil
}
