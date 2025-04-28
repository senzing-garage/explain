package explainer_test

import (
	"context"

	"github.com/senzing-garage/explain/explainer"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleNullExplainer() {
	// For more information, visit https://github.com/senzing-garage/explain/blob/main/explainer/explainer_message_test.go
	ctx := context.TODO()
	explainer := &explainer.NullExplainer{}

	err := explainer.Explain(ctx)
	if err != nil {
		panic(err)
	}
	// Output:
}
