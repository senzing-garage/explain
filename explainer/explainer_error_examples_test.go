package explainer

import (
	"context"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleErrorExplainer() {
	// For more information, visit https://github.com/senzing-garage/explain/blob/main/explainer/explainer_error_test.go
	ctx := context.TODO()
	explainer := &ErrorExplainer{
		ErrorID: "SZSDK60010000",
		TtyOnly: true,
	}
	err := explainer.Explain(ctx)
	if err != nil {
		panic(err)
	}
	// Output: For information on that error, visit https://hub.senzing.com/sz-sdk-go-core/errors#SZSDK60010000
}
