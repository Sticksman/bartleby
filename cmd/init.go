package cmd

import (
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
	return nil
}
