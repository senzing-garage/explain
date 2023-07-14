package explainer

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Test harness
// ----------------------------------------------------------------------------

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	code := m.Run()
	err = teardown()
	if err != nil {
		fmt.Print(err)
	}
	os.Exit(code)
}

func setup() error {
	var err error = nil
	return err
}

func teardown() error {
	var err error = nil
	return err
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestExplainerError(test *testing.T) {
	ctx := context.TODO()
	testObject := &ExplainerError{
		ErrorMessage: "senzing-12345678",
	}
	err := testObject.Explain(ctx)
	assert.Nil(test, err)
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleExplainerError() {
	// For more information, visit https://github.com/Senzing/explain/blob/main/examplepackage/examplepackage_test.go
	ctx := context.TODO()
	explainer := &ExplainerError{
		ErrorMessage: "senzing-12345678",
	}
	explainer.Explain(ctx)
	//Output:
	//examplePackage: I'm here
}
