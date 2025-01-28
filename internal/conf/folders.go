package conf

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"pkg.mattglei.ch/timber"
)

func CreateClassFolders(classes []Class) error {
	for _, class := range classes {
		_, err := os.Stat(class.Folder)
		if errors.Is(err, fs.ErrNotExist) {
			err = os.Mkdir(class.Folder, os.ModeDir)
			if err != nil {
				return fmt.Errorf("%v failed to create directory: %s", err, class.Folder)
			}
			timber.Done("created", class.Folder, "directory")
		}
	}
	return nil
}
