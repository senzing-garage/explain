package explainer

import (
	"context"
	"errors"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// NullExplainer is a "null object" Explainer.
type NullExplainer struct {
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// const exampleConstant = "examplePackage"

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The Explain method simply returns an error to drive the "Usage" output by the caller.

Input
  - ctx: A context to control lifecycle.
*/
func (explainer *NullExplainer) Explain(ctx context.Context) error {
	_ = ctx
	return errors.New("need to specify an option (e.g. '--error-id')")
}
