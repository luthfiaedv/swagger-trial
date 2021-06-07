package docs

import (
	"github.com/luthfiaedv/swagger-trial/spec"
)

// swagger:route GET / response-tag response-get
// Response does some amazing stuff.
// responses:
//	200: responseSuccess

// This text will appear as description of your response body.
// swagger:response responseSuccess
type trialResponseWrapper struct {
	// in:body
	Body spec.Response
}
