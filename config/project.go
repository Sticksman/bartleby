package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Project is the top level object
type Project struct {
	Title        string   `json:"title"`
	RootFilepath string   `json:"root_filepath"`
	IgnoredFiles []string `json:"ignored_files"`
	FileOrder    []string `json:"file_order"`
	// TODO: Add in per file/directory stuff
}

// Block is the most basic unit of a story
type Block struct {
	Name     string
	Contents string
}

// Compile takes all the files in the file ordering and combines it into a single file
func (p *Project) Compile() (string, error) {
	blockCh := make(chan *Block, len(p.FileOrder))
	processCh := make(chan bool, 4)
	blockContents := map[string]string{}
	output := ""

	for _, filename := range p.FileOrder {
		go func(name string) {
			processCh <- true
			p.compileFile(name, blockCh)
			<-processCh
		}(filename)
	}

	for i := 0; i < cap(processCh); i++ {
		processCh <- true
	}

	close(processCh)

	for block := range blockCh {
		blockContents[block.Name] = block.Contents
	}
	close(blockCh)

	for _, filename := range p.FileOrder {
		content, ok := blockContents[filename]
		if !ok {
			continue
		}
		output = fmt.Sprintf("%s\n", content)
	}

	if output == "" {
		return "", errors.New("Could not find any files to join")
	}

	output = fmt.Sprintf("#%s\n\n%s", p.Title, output)

	return output, nil
}

func (p *Project) compileFile(name string, ch chan *Block) error {
	blockPath := filepath.Join(p.RootFilepath, name)

	fi, err := os.Stat(blockPath)
	if err != nil {
		return err
	}

	mode := fi.Mode()
	if mode.IsDir() {
		return errors.New(blockPath + " is not a file")
	}

	dat, err := ioutil.ReadFile(blockPath)
	if err != nil {
		return err
	}

	body := string(dat)
	body = strings.Trim(body, " \n")

	b := Block{
		Name:     name,
		Contents: body,
	}
	ch <- &b

	return nil
}
