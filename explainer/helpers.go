package explainer

import "fmt"

func page(uriSegment string) string {
	return fmt.Sprintf("senzing-garage.github.io/%s/errors", uriSegment)
}
