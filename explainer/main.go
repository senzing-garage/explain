package explainer

import (
	"context"
	"errors"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The Explainer interface is used as a "Command Pattern".
// See https://en.wikipedia.org/wiki/Command_pattern
type Explainer interface {
	Explain(ctx context.Context) error
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// An example constant.
// const ExampleConstant = 1

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var ComponentID2WebPage = map[int]string{
	6001: page("sz-sdk-go-core"),
	6002: page("sz-sdk-go-core"),
	6003: page("sz-sdk-go-core"),
	6004: page("sz-sdk-go-core"),
	6005: page("sz-sdk-go-core"),
	6006: page("sz-sdk-go-core"),
	6007: page("sz-sdk-go-core"),
	6008: page("sz-sdk-go-core"),
	6009: page("sz-sdk-go-core"),
	6011: page("serve-grpc"),
	6012: page("serve-grpc"),
	6013: page("serve-grpc"),
	6014: page("serve-grpc"),
	6015: page("serve-grpc"),
	6016: page("serve-grpc"),
	6017: page("serve-grpc"),
	6021: page("sz-sdk-go-grpc"),
	6022: page("sz-sdk-go-grpc"),
	6023: page("sz-sdk-go-grpc"),
	6024: page("sz-sdk-go-grpc"),
	6025: page("sz-sdk-go-grpc"),
	6026: page("sz-sdk-go-grpc"),
	6027: page("sz-sdk-go-grpc"),
	6031: page("sz-sdk-go-mock"),
	6032: page("sz-sdk-go-mock"),
	6033: page("sz-sdk-go-mock"),
	6034: page("sz-sdk-go-mock"),
	6035: page("sz-sdk-go-mock"),
	6036: page("sz-sdk-go-mock"),
	6037: page("sz-sdk-go-mock"),
	6038: page("sz-sdk-go-mock"),
	6039: page("sz-sdk-go-mock"),
	6041: page("go-sdk-abstract-factory"),
	6100: page("senzing-tools"),
	6201: page("load"),
	6202: page("move"),
	6203: page("validate"),
	6204: page("serve-grpc"),
	6205: page("init-database"),
	6206: page("shuffle"),
	6207: page("observe"),
	6401: page("go-helpers"),
	6402: page("go-helpers"),
	6403: page("go-helpers"),
	6404: page("go-helpers"),
	6421: page("go-databasing"),
	6422: page("go-databasing"),
	6423: page("go-databasing"),
	6441: page("go-logging"),
	6461: page("go-observing"),
	6462: page("go-observing"),
	6463: page("go-observing"),
	6464: page("go-observing"),
	6481: page("go-queuing"),
	6501: page("init-database"),
	6502: page("init-database"),
	6503: page("init-database"),
	6601: page("load"),
	6602: page("move"),
	6620: page("serve-chat"),
}

var errPackage = errors.New("explainer")
