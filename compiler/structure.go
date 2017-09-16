package compiler

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Sticksman/bartleby/config"
)

// MapDirectoryTree walks the project directory and tries to create blocks from what it finds
func MapDirectoryTree(path string, ignoreFiles []string) ([]*config.Block, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	mode := fi.Mode()
	if !mode.IsDir() {
		return nil, errors.New(path + " is not a directory")
	}

	blocks := []*config.Block{}
	fileList, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, fileInfo := range fileList {
		if !shouldReadFile(fileInfo.Name(), ignoreFiles) {
			continue
		}
		// If not a dir, we want to create a block from the filename
		block := config.NewBlock(fileInfo.Name())
		if fileInfo.IsDir() {
			fp := filepath.Join(path, fileInfo.Name())
			subblocks, err := MapDirectoryTree(fp, []string{})
			if err != nil {
				fmt.Printf("Could not map directory %v: %v\n", fp, err)
				continue
			}
			block.Substructure = subblocks
		}
		blocks = append(blocks, &block)
	}

	return blocks, nil
}

func shouldReadFile(filename string, ignoreFiles []string) bool {
	for _, ignoreFileName := range ignoreFiles {
		if ignoreFileName == filename {
			return false
		}
	}
	return true
}
