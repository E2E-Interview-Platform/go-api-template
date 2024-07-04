package db

import (
	"os"
	"strconv"
	"strings"

	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/helpers"
)

func createFile(filename string) (err error) {
	f, err := os.Create(filename)
	if err != nil {
		return
	}

	err = f.Close()
	return
}

// getMigrationVersionsFromDir reads the migration files from the local directory,
// parses the version from file name and returns them in slice sorted in descending order
func getMigrationVersionsFromDir(directoryName string) ([]int, error) {
	keys := make(map[int]bool)
	fileVersions := make([]int, 0)

	files, err := os.ReadDir(directoryName)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		version := strings.Split(file.Name(), "_")[0]

		v, err := strconv.Atoi(version)
		if err != nil {
			return nil, err
		}

		if _, ok := keys[v]; !ok {
			keys[v] = true
			fileVersions = append(fileVersions, v)
		}
	}

	helpers.SortIntDescending(fileVersions)

	return fileVersions, nil
}
