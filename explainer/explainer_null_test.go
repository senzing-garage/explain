package explainer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Test public functions
// ----------------------------------------------------------------------------

func TestNullExplainer(test *testing.T) {
	ctx := context.TODO()
	testObject := &NullExplainer{}
	// os.Args = []string{"command-name", "--message-id", "SZSDK60010000"}
	err := testObject.Explain(ctx)
	assert.NoError(test, err)
}
