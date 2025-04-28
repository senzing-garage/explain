/*
 */
package cmd

import (
	"context"
	"os"

	"github.com/senzing-garage/explain/explainer"
	"github.com/senzing-garage/go-cmdhelping/cmdhelper"
	"github.com/senzing-garage/go-cmdhelping/option"
	"github.com/senzing-garage/go-helpers/wraperror"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	Short string = "provide an explanation"
	Use   string = "explain"
	Long  string = `
Explain an aspect of Senzing.
    `
)

// ----------------------------------------------------------------------------
// Context variables
// ----------------------------------------------------------------------------

var ContextVariablesForMultiPlatform = []option.ContextVariable{
	option.MessageID,
	option.TtyOnly,
}

var ContextVariables = append(ContextVariablesForMultiPlatform, ContextVariablesForOsArch...)

// ----------------------------------------------------------------------------
// Command
// ----------------------------------------------------------------------------

// RootCmd represents the command.
var RootCmd = &cobra.Command{
	Use:     Use,
	Short:   Short,
	Long:    Long,
	PreRun:  PreRun,
	RunE:    RunE,
	Version: Version(),
}

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Used in construction of cobra.Command.
func PreRun(cobraCommand *cobra.Command, args []string) {
	cmdhelper.PreRun(cobraCommand, args, Use, ContextVariables)
}

// Used in construction of cobra.Command.
func RunE(_ *cobra.Command, _ []string) error {
	ctx := context.Background()

	var anExplainer explainer.Explainer

	// Choose and run explainer.

	if len(viper.GetString(option.MessageID.Arg)) > 0 {
		anExplainer = &explainer.MessageExplainer{
			MessageID: viper.GetString(option.MessageID.Arg),
			TtyOnly:   viper.GetBool(option.TtyOnly.Arg),
		}
	} else {
		anExplainer = &explainer.NullExplainer{}
	}

	err := anExplainer.Explain(ctx)

	return wraperror.Errorf(err, "cmd.RunE error: %w", err)
}

// Used in construction of cobra.Command.
func Version() string {
	return cmdhelper.Version(githubVersion, githubIteration)
}

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

// Since init() is always invoked, define command line parameters.
func init() {
	cmdhelper.Init(RootCmd, ContextVariables)
}
