package explainer

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/browser"
	"github.com/senzing-garage/go-helpers/wraperror"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// MessageExplainer is a Explainer for messages.
type MessageExplainer struct {
	MessageID string
	TtyOnly   bool
}

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

func ParseMessage(message string) (int, int, error) {
	const expectedLength = 8

	var err error

	componentID := 0
	messageID := 0
	intString := strings.TrimPrefix(message, "SZSDK")

	if len(intString) != expectedLength {
		return componentID, messageID, wraperror.Errorf(errPackage, "could not parse message: %s", message)
	}

	componentIDString := intString[0:4]

	componentID, err = strconv.Atoi(componentIDString)
	if err != nil {
		return componentID, messageID, wraperror.Errorf(err, "componentIDString strconv.Atoi(%s)", componentIDString)
	}

	messageID, err = strconv.Atoi(intString)
	if err != nil {
		return componentID, messageID, wraperror.Errorf(err, "intString strconv.Atoi(%s)", intString)
	}

	return componentID, messageID, wraperror.Errorf(err, wraperror.NoMessage)
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
func (explainer *MessageExplainer) Explain(ctx context.Context) error {
	_ = ctx

	componentID, messageNumber, err := ParseMessage(explainer.MessageID)
	if err != nil {
		return err
	}

	webpage, ok := ComponentID2WebPage[componentID]
	if !ok {
		return wraperror.Errorf(errPackage, "no information for --message-id %s", explainer.MessageID)
	}

	explainURL := fmt.Sprintf("https://%s#SZSDK%d", webpage, messageNumber)

	fmt.Printf("For information on that message, visit %s\n", explainURL) //nolint

	// Reference: https://gist.github.com/hyg/9c4afcd91fe24316cbf0

	if !explainer.TtyOnly {
		err = browser.OpenURL(explainURL)
	}

	return wraperror.Errorf(err, wraperror.NoMessage)
}
