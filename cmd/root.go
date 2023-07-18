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
	envarErrorId  = "SENZING_TOOLS_ERROR_ID"
	helpErrorId   = "Give an explanation of a specific Senzing error [%s]"
	optionErrorId = "error-id"
	envarTtyOnly  = "SENZING_TOOLS_TTY_ONLY"
	helpTtyOnly   = "Output confined to terminal (TTY) [%s]"
	optionTtyOnly = "tty-only"
)

// ----------------------------------------------------------------------------
// Context variables
// ----------------------------------------------------------------------------

var ContextStrings = []cmdhelper.ContextString{
	{
		Default: cmdhelper.OsLookupEnvString(envarErrorId, ""),
		Envar:   envarErrorId,
		Help:    helpErrorId,
		Option:  optionErrorId,
	},
}

var ContextBools = []cmdhelper.ContextBool{
	{
		Default: cmdhelper.OsLookupEnvBool(envarTtyOnly, false),
		Envar:   envarTtyOnly,
		Help:    helpTtyOnly,
		Option:  optionTtyOnly,
	},
}

var ContextVariables = &cmdhelper.ContextVariables{
	Strings: append(ContextStrings, ContextStringsForOsArch...),
	Bools:   append(ContextBools, ContextBoolsForOsArch...),
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

	if len(viper.GetString(optionErrorId)) > 0 {
		anExplainer = &explainer.ExplainerError{
			ErrorId: viper.GetString(optionErrorId),
			TtyOnly: viper.GetBool(optionTtyOnly),
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
