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
	testObject := &explainer.ExplainerError{
		ErrorId: "senzing-60010000",
		TtyOnly: true,
	}
	err := testObject.Explain(ctx)
	assert.Nil(test, err)
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleExplainerError() {
	// For more information, visit https://github.com/senzing-garage/explain/blob/main/explainer/explainer_error_test.go
	ctx := context.TODO()
	explainer := &explainer.ExplainerError{
		ErrorId: "senzing-60010000",
		TtyOnly: true,
	}
	err := explainer.Explain(ctx)
	if err != nil {
		panic(err)
	}
	//Output: For information on that error, visit https://hub.senzing.com/g2-sdk-go-base/errors#senzing-60010000
}
