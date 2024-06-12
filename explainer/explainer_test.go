package explainer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test interface methods
// ----------------------------------------------------------------------------

func TestErrorExplainer_Explain(test *testing.T) {
	ctx := context.TODO()
	testObject := &ErrorExplainer{
		ErrorID: "SZSDK60010000",
	}
	err := testObject.Explain(ctx)
	require.NoError(test, err)
}

func TestErrorExplainer_Explain_badErrorID(test *testing.T) {
	ctx := context.TODO()
	testObject := &ErrorExplainer{
		ErrorID: "SZSDK6A010000",
	}
	err := testObject.Explain(ctx)
	require.Error(test, err)
}

func TestErrorExplainer_Explain_badComponentID(test *testing.T) {
	ctx := context.TODO()
	testObject := &ErrorExplainer{
		ErrorID: "SZSDK99980000",
	}
	err := testObject.Explain(ctx)
	require.Error(test, err)
}

// ----------------------------------------------------------------------------
// Test public functions
// ----------------------------------------------------------------------------

func TestParseErrorMessage(test *testing.T) {
	messageID := "SZSDK60010000"
	_, _, err := ParseErrorMessage(messageID)
	require.NoError(test, err)
}

func TestParseErrorMessage_badPrefix(test *testing.T) {
	badMessagID := "123456789"
	_, _, err := ParseErrorMessage(badMessagID)
	require.Error(test, err)
}

func TestParseErrorMessage_badComponentID(test *testing.T) {
	badMessagID := "SZSDK6A010000"
	_, _, err := ParseErrorMessage(badMessagID)
	require.Error(test, err)
}

func TestParseErrorMessage_badInstanceID(test *testing.T) {
	badMessagID := "SZSDK60010A00"
	_, _, err := ParseErrorMessage(badMessagID)
	require.Error(test, err)
}
