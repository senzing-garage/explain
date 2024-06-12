package explainer

import (
	"context"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleMessageExplainer_Explain() {
	// For more information, visit https://github.com/senzing-garage/explain/blob/main/explainer/explainer_message_test.go
	ctx := context.TODO()
	explainer := &MessageExplainer{
		MessageID: "SZSDK60010000",
		TtyOnly:   true,
	}
	err := explainer.Explain(ctx)
	if err != nil {
		panic(err)
	}
	// Output: For information on that message, visit https://senzing-garage.github.io/sz-sdk-go-core/errors#SZSDK60010000
}
