package cmd

import (
	"encoding/json"
	"fmt"
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
	Run:   compileProject,
}

func init() {
	compileCmd.Flags().StringVarP(&metadataPath, "metadata-path", "m", ".metadata", "path to metadata")
}

func compileProject(cmd *cobra.Command, args []string) {
	configPath := filepath.Join(filepath.FromSlash(metadataPath), "config.json")

	fi, err := os.Stat(configPath)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	mode := fi.Mode()
	if mode.IsDir() {
		fmt.Println(configPath + " is not a file")
		return
	}

	dat, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	var project config.Project
	err = json.Unmarshal(dat, &project)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	output, err := project.Compile()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	ioutil.WriteFile("output.md", []byte(output), 0755)
}
