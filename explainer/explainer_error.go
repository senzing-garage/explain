package explainer

import (
	"context"
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// ExplainError is an example type-struct.
type ExplainerError struct {
	ErrorMessage string
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// const exampleConstant = "examplePackage"

// ----------------------------------------------------------------------------
// private functions
// ----------------------------------------------------------------------------

func openBrowser(url string) error {
	var err error = nil
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return err
}

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

func ParseErrorMessage(errorMessage string) (int, int, error) {
	var err error = nil
	componentId := 0
	messageId := 0
	intString := strings.TrimPrefix(errorMessage, "senzing-")
	if len(intString) != 8 {
		return componentId, messageId, fmt.Errorf("could not parse error message: %s", errorMessage)
	}
	componentIdString := intString[0:4]
	componentId, err = strconv.Atoi(componentIdString)
	if err != nil {
		return componentId, messageId, err
	}
	messageId, err = strconv.Atoi(intString)
	if err != nil {
		return componentId, messageId, err
	}
	return componentId, messageId, err
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The Explain method...

Input
  - ctx: A context to control lifecycle.

Output
  - Nothing is returned, except for an error.  However, something is printed.
    See the example output.
*/
func (explainer *ExplainerError) Explain(ctx context.Context) error {

	componentId, errorNumber, err := ParseErrorMessage(explainer.ErrorMessage)
	if err != nil {
		return err
	}

	webpage, ok := ComponentId2WebPage[componentId]
	if !ok {
		return fmt.Errorf("no information for --error-message %s", explainer.ErrorMessage)
	}

	explainUrl := fmt.Sprintf("https://%s#senzing-%d", webpage, errorNumber)

	fmt.Printf("For information on that error, visit %s\n", explainUrl)
	err = openBrowser(explainUrl)
	return err
}
