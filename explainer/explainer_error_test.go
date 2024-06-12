package explainer_test

import (
	"context"
	"testing"

	"github.com/senzing-garage/explain/explainer"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestExplainerError(test *testing.T) {
	ctx := context.TODO()
	testObject := &explainer.ErrorExplainer{
		ErrorID: "senzing-60010000",
		TtyOnly: true,
	}
	err := testObject.Explain(ctx)
	assert.NoError(test, err)
}
