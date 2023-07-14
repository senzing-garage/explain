/*
 */
package cmd

import (
	"context"
	"os"

	"github.com/senzing/explain/explainer"
	"github.com/senzing/senzing-tools/cmdhelper"
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

var (
	envarExplainError  = "SENZING_TOOLS_EXPLAIN_ERROR"
	helpExplainError   = "Give an explanation of a specific Senzing error [%s]"
	optionExplainError = "error"
)

// ----------------------------------------------------------------------------
// Context variables
// ----------------------------------------------------------------------------

var ContextStrings = []cmdhelper.ContextString{
	{
		Default: cmdhelper.OsLookupEnvString(envarExplainError, ""),
		Envar:   envarExplainError,
		Help:    helpExplainError,
		Option:  optionExplainError,
	},
}

var ContextVariables = &cmdhelper.ContextVariables{
	Strings: append(ContextStrings, ContextStringsForOsArch...),
}

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

// Since init() is always invoked, define command line parameters.
func init() {
	cmdhelper.Init(RootCmd, *ContextVariables)
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

// Used in construction of cobra.Command
func PreRun(cobraCommand *cobra.Command, args []string) {
	cmdhelper.PreRun(cobraCommand, args, Use, *ContextVariables)
}

// Used in construction of cobra.Command
func RunE(_ *cobra.Command, _ []string) error {
	ctx := context.Background()
	var anExplainer explainer.Explainer

	// Choose explainer.

	if len(viper.GetString(optionExplainError)) > 0 {
		anExplainer = &explainer.ExplainerError{
			ErrorMessage: viper.GetString(optionExplainError),
		}
	} else {
		anExplainer = &explainer.ExplainerNull{}

	}

	return anExplainer.Explain(ctx)
}

// Used in construction of cobra.Command
func Version() string {
	return cmdhelper.Version(githubVersion, githubIteration)
}

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
