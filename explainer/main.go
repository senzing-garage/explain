package explainer

import (
	"context"
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

var ComponentId2WebPage = map[int]string{
	6001: page("g2-sdk-go-base"),
	6002: page("g2-sdk-go-base"),
	6003: page("g2-sdk-go-base"),
	6004: page("g2-sdk-go-base"),
	6005: page("g2-sdk-go-base"),
	6006: page("g2-sdk-go-base"),
	6007: page("g2-sdk-go-base"),
	6008: page("g2-sdk-go-base"),
	6009: page("g2-sdk-go-base"),
	6011: page("servegrpc"),
	6012: page("servegrpc"),
	6013: page("servegrpc"),
	6014: page("servegrpc"),
	6015: page("servegrpc"),
	6016: page("servegrpc"),
	6017: page("servegrpc"),
	6021: page("g2-sdk-go-grpc"),
	6022: page("g2-sdk-go-grpc"),
	6023: page("g2-sdk-go-grpc"),
	6024: page("g2-sdk-go-grpc"),
	6025: page("g2-sdk-go-grpc"),
	6026: page("g2-sdk-go-grpc"),
	6027: page("g2-sdk-go-grpc"),
	6031: page("g2-sdk-go-mock"),
	6032: page("g2-sdk-go-mock"),
	6033: page("g2-sdk-go-mock"),
	6034: page("g2-sdk-go-mock"),
	6035: page("g2-sdk-go-mock"),
	6036: page("g2-sdk-go-mock"),
	6037: page("g2-sdk-go-mock"),
	6038: page("g2-sdk-go-mock"),
	6039: page("g2-sdk-go-mock"),
	6041: page("go-sdk-abstract-factory"),
	6100: page("senzing-tools"),
	6201: page("load"),
	6202: page("move"),
	6203: page("validate"),
	6204: page("serve-grpc"),
	6205: page("init-database"),
	6206: page("shuffle"),
	6207: page("observe"),
	6401: page("go-common"),
	6402: page("go-common"),
	6403: page("go-common"),
	6404: page("go-common"),
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
