package cmd

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// InitCmd represents the initialization of a new to be compiled directory
var InitCmd = &cobra.Command{
	Use:   "init [path?]",
	Short: "Creates a bartleby project",
	Long: `Creates a bartleby project at the specified path
or in the current directory if no path is specified.
It does so by creating a .config.json file in the directory.
You can then edit the config file to finish configuring the project.`,
	RunE: initProject,
}

func initProject(cmd *cobra.Command, args []string) error {
	path := "."
	if len(args) >= 1 {
		path = args[0]
	}
	path = filepath.FromSlash(path)

	// Check if we're pointing to a directory
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}

	mode := fi.Mode()
	if !mode.IsDir() {
		return errors.New(path + " already exists but is not a directory")
	}

	err = createConfig(path)
	if err != nil {
		return err
	}

	return nil
}

func createConfig(path string) error {
	return nil
}
