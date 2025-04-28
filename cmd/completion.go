/*
 */
package cmd

import (
	"io"
	"os"

	"github.com/senzing-garage/go-helpers/wraperror"
	"github.com/spf13/cobra"
)

// completionCmd represents the completion command.
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate bash completion for the command",
	Long: `To load completions, run:
source < (explain completion)

To load completions automaticallon on login, add this line to your .bashrc file:
source < (explain completion)
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		_ = cmd
		_ = args

		return completionAction(os.Stdout)
	},
}

func init() {
	RootCmd.AddCommand(completionCmd)
}

func completionAction(out io.Writer) error {
	err := RootCmd.GenBashCompletion(out)

	return wraperror.Errorf(err, "cmd.completionAction error: %w", err)
}
