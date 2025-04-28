package explainer_test

import (
	"testing"

	"github.com/senzing-garage/explain/explainer"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test public functions
// ----------------------------------------------------------------------------

func TestParseMessage(test *testing.T) {
	test.Parallel()

	messageID := "SZSDK60010000"
	_, _, err := explainer.ParseMessage(messageID)
	require.NoError(test, err)
}

func TestParseMessage_badPrefix(test *testing.T) {
	test.Parallel()

	badMessagID := "123456789"
	_, _, err := explainer.ParseMessage(badMessagID)
	require.Error(test, err)
}

func TestParseMessage_badComponentID(test *testing.T) {
	test.Parallel()

	badMessagID := "SZSDK6A010000"
	_, _, err := explainer.ParseMessage(badMessagID)
	require.Error(test, err)
}

func TestParseMessage_badInstanceID(test *testing.T) {
	test.Parallel()

	badMessagID := "SZSDK60010A00"
	_, _, err := explainer.ParseMessage(badMessagID)
	require.Error(test, err)
}
