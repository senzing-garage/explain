package explainer

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/browser"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// ExplainError is a Explainer for errors.
type ErrorExplainer struct {
	ErrorID string
	TtyOnly bool
}

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

func ParseErrorMessage(errorMessage string) (int, int, error) {
	var err error
	componentID := 0
	messageID := 0
	intString := strings.TrimPrefix(errorMessage, "senzing-")
	if len(intString) != 8 {
		return componentID, messageID, fmt.Errorf("could not parse error message: %s", errorMessage)
	}
	componentIDString := intString[0:4]
	componentID, err = strconv.Atoi(componentIDString)
	if err != nil {
		return componentID, messageID, err
	}
	messageID, err = strconv.Atoi(intString)
	if err != nil {
		return componentID, messageID, err
	}
	return componentID, messageID, err
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
func (explainer *ErrorExplainer) Explain(ctx context.Context) error {
	_ = ctx

	componentID, errorNumber, err := ParseErrorMessage(explainer.ErrorID)
	if err != nil {
		return err
	}

	webpage, ok := ComponentID2WebPage[componentID]
	if !ok {
		return fmt.Errorf("no information for --error-message %s", explainer.ErrorID)
	}

	explainURL := fmt.Sprintf("https://%s#senzing-%d", webpage, errorNumber)

	fmt.Printf("For information on that error, visit %s\n", explainURL)

	// Reference: https://gist.github.com/hyg/9c4afcd91fe24316cbf0

	if !explainer.TtyOnly {
		err = browser.OpenURL(explainURL)
	}
	return err
}
