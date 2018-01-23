package cmd

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Sticksman/bartleby/config"
	"github.com/spf13/cobra"
)

var metadataPath string

var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compiles all files in the file ordering field",
	Long:  "Takes all files listed in the config command and joins them together.",
}

func init() {
	compileCmd.Flags().StringVarP(&metadataPath, "metadata-path", "m", ".metadata", "path to metadata")
}

func compileProject(cmd *cobra.Command, args []string) error {
	configPath := filepath.Join(filepath.FromSlash(metadataPath), "config.json")

	fi, err := os.Stat(configPath)
	if err != nil {
		return err
	}

	mode := fi.Mode()
	if mode.IsDir() {
		return errors.New(configPath + " is not a file")
	}

	dat, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	project := config.Project{}
	json.Unmarshal(dat, &project)
	output, err := project.Compile()
	if err != nil {
		return err
	}

	ioutil.WriteFile("output.md", []byte(output), 0755)

	return nil
}
