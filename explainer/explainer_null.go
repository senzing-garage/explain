package explainer

import (
	"context"
	"errors"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// ExplainError is an example type-struct.
type ExplainerNull struct {
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
func (explainer *ExplainerNull) Explain(ctx context.Context) error {
	return errors.New("need to specify an option (e.g. '--error')")
}
