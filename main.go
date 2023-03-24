package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	createParents bool
	permissions   os.FileMode
)

var rootCmd = &cobra.Command{
	Use:   "mktouch <filename>",
	Short: "Create an empty file at the specified path",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		filename := args[0]

		if createParents {
			err := os.MkdirAll(filepath.Dir(filename), 0755)
			if err != nil {
				return fmt.Errorf("error creating directory: %w", err)
			}
		}

		file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, permissions)
		if err != nil {
			return fmt.Errorf("error creating file: %w", err)
		}

		file.Close()
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&createParents, "parents", "p", false, "Create parent directories if they don't exist")
	rootCmd.PersistentFlags().Uint32VarP((*uint32)(&permissions), "mode", "m", 0644, "File permissions (in octal)")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

