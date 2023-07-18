package explainer

import "fmt"

func page(uriSegment string) string {
	return fmt.Sprintf("hub.senzing.com/%s/errors", uriSegment)
}
