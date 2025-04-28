package explainer_test

import (
	"testing"

	"github.com/senzing-garage/explain/explainer"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Test public functions
// ----------------------------------------------------------------------------

func TestNullExplainer(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	testObject := &explainer.NullExplainer{}
	err := testObject.Explain(ctx)
	assert.NoError(test, err)
}
