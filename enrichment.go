// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"github.com/blockaid-official/blockaid-client-go/option"
)

// EnrichmentService contains methods and other services that help with interacting
// with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnrichmentService] method instead.
type EnrichmentService struct {
	Options []option.RequestOption
}

// NewEnrichmentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEnrichmentService(opts ...option.RequestOption) (r *EnrichmentService) {
	r = &EnrichmentService{}
	r.Options = opts
	return
}
