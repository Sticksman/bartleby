package cmd

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var projectName string
var projectPath string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates a bartleby project",
	Long: `Creates a bartleby project with the specified path
or in the current directory if no path is specified.
Takes a name or uses the name of the directory as the project name.
This command creates a .config.json file in the directory.
You can then edit the config file to finish configuring the project.`,
	RunE: initProject,
}

func init() {
	initCmd.Flags().StringVarP(&projectName, "name", "n", "", "name of the project")
	initCmd.Flags().StringVarP(&projectPath, "path", "p", ".", "path of the project")
}

func initProject(cmd *cobra.Command, args []string) error {
	path := filepath.FromSlash(projectPath)

	// Check if we're pointing to a directory
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}

	mode := fi.Mode()
	if !mode.IsDir() {
		return errors.New(path + " already exists but is not a directory")
	}

	name := projectName
	if name == "" {
		name = filepath.Base(path)
	}

	err = createConfig(name, path)
	if err != nil {
		return err
	}

	return nil
}

func createConfig(name string, path string) error {
	return nil
}
