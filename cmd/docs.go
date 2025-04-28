/*
 */
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/senzing-garage/go-helpers/wraperror"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// docsCmd represents the docs command.
var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Generate documentation for the command",
	RunE: func(cmd *cobra.Command, args []string) error {
		_ = args
		dir, err := cmd.Flags().GetString("dir")
		if err != nil {
			return wraperror.Errorf(err, "cmd.Flags.GetString error: %w", err)
		}
		if dir == "" {
			if dir, err = os.MkdirTemp("", "explain"); err != nil {
				return wraperror.Errorf(err, "cmd.os.MkdirTemp error: %w", err)
			}
		}

		return docsAction(os.Stdout, dir)
	},
}

func init() {
	RootCmd.AddCommand(docsCmd)
	docsCmd.Flags().StringP("dir", "d", "", "Destination directory for docs")
}

func docsAction(out io.Writer, dir string) error {
	if err := doc.GenMarkdownTree(RootCmd, dir); err != nil {
		return wraperror.Errorf(err, "cmd.docsAction error: %w", err)
	}

	_, err := fmt.Fprintf(out, "Documentation successfully created in %s\n", dir)

	return wraperror.Errorf(err, "cmd.docsAction error: %w", err)
}
