package explainer_test

import (
	"testing"

	"github.com/senzing-garage/explain/explainer"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test public functions
// ----------------------------------------------------------------------------

func TestMessageExplainer_Explain(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	testObject := &explainer.MessageExplainer{
		MessageID: "SZSDK60010000",
		TtyOnly:   true,
	}
	err := testObject.Explain(ctx)
	require.NoError(test, err)
}

func TestMessageExplainer_Explain_badMessageID(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	testObject := &explainer.MessageExplainer{
		MessageID: "SZSDK6A010000",
		TtyOnly:   true,
	}
	err := testObject.Explain(ctx)
	require.Error(test, err)
}

func TestMessageExplainer_Explain_badComponentID(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	testObject := &explainer.MessageExplainer{
		MessageID: "SZSDK99980000",
		TtyOnly:   true,
	}
	err := testObject.Explain(ctx)
	require.Error(test, err)
}
