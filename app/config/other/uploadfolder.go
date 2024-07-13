package other

import (
	"log"
	"os"
	"path/filepath"
)

func InitUploadFolder() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	newFolderPath := filepath.Join(cwd, "upload", "images")

	if _, err := os.Stat(newFolderPath); os.IsNotExist(err) {
		err = os.MkdirAll(newFolderPath, 0755)
		if err != nil {
			panic(err)
		}
		log.Println("Upload Folder Initialized")
	}
}
