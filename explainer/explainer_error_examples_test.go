package explainer_test

import (
	"context"

	"github.com/senzing-garage/explain/explainer"
)

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
	// Output: For information on that error, visit https://hub.senzing.com/sz-sdk-go-core/errors#senzing-60010000
}
