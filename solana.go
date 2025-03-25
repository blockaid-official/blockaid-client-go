// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"github.com/blockaid-official/blockaid-client-go/option"
)

// SolanaService contains methods and other services that help with interacting
// with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSolanaService] method instead.
type SolanaService struct {
	Options []option.RequestOption
	Message *SolanaMessageService
	Address *SolanaAddressService
}

// NewSolanaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSolanaService(opts ...option.RequestOption) (r *SolanaService) {
	r = &SolanaService{}
	r.Options = opts
	r.Message = NewSolanaMessageService(opts...)
	r.Address = NewSolanaAddressService(opts...)
	return
}
