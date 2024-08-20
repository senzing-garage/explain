package main

import (
	"os"
	"testing"
)

func TestMain(test *testing.T) {
	_ = test
	os.Args = []string{"command-name", "--message-id", "SZSDK60010000", "--tty-only"}
	main()
}
