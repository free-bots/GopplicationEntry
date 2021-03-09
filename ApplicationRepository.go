package GoppilcationEntry

import (
	"fmt"
	"os"
	"path/filepath"
)

func FindAllEntries() []*ApplicationEntry {
	applications := make([]*ApplicationEntry, 0)
	for _, path := range GetApplicationPaths() {

		path = filepath.Join(path, ApplicationDirName)

		fileInfo, err := os.Stat(path)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if !fileInfo.IsDir() {
			continue
		}

		file, err := os.Open(path)

		if err != nil {
			return nil
		}

		fileNames, err := file.Readdirnames(0)

		_ = file.Close()

		if err != nil {
			return nil
		}

		for _, name := range fileNames {
			application := Parse(filepath.Join(path, name), true)
			if application == nil {
				continue
			}

			applications = append(applications, application)
		}
	}

	return applications
}
