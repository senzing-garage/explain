package explainer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test public functions
// ----------------------------------------------------------------------------

func TestParseMessage(test *testing.T) {
	messageID := "SZSDK60010000"
	_, _, err := ParseMessage(messageID)
	require.NoError(test, err)
}

func TestParseMessage_badPrefix(test *testing.T) {
	badMessagID := "123456789"
	_, _, err := ParseMessage(badMessagID)
	require.Error(test, err)
}

func TestParseMessage_badComponentID(test *testing.T) {
	badMessagID := "SZSDK6A010000"
	_, _, err := ParseMessage(badMessagID)
	require.Error(test, err)
}

func TestParseMessage_badInstanceID(test *testing.T) {
	badMessagID := "SZSDK60010A00"
	_, _, err := ParseMessage(badMessagID)
	require.Error(test, err)
}
