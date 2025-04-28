package explainer

import (
	"context"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// NullExplainer is a "null object" Explainer.
type NullExplainer struct{}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// const exampleConstant = "examplePackage"

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The Explain method simply returns a nil to drive the "Usage" output by the caller.

Input
  - ctx: A context to control lifecycle.
*/
func (explainer *NullExplainer) Explain(ctx context.Context) error {
	_ = ctx

	var err error

	return err
}
