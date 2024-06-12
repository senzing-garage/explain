package explainer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Test public functions
// ----------------------------------------------------------------------------

func TestErrorExplainer(test *testing.T) {
	ctx := context.TODO()
	testObject := &ErrorExplainer{
		ErrorID: "SZSDK60010000",
		TtyOnly: true,
	}
	err := testObject.Explain(ctx)
	assert.NoError(test, err)
}
