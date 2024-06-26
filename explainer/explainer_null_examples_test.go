package explainer

import (
	"context"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleNullExplainer() {
	// For more information, visit https://github.com/senzing-garage/explain/blob/main/explainer/explainer_message_test.go
	ctx := context.TODO()
	explainer := &NullExplainer{}
	err := explainer.Explain(ctx)
	if err != nil {
		panic(err)
	}
	// Output:
}
