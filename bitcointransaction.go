// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"github.com/blockaid-official/blockaid-client-go/option"
)

// BitcoinTransactionService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBitcoinTransactionService] method instead.
type BitcoinTransactionService struct {
	Options []option.RequestOption
}

// NewBitcoinTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBitcoinTransactionService(opts ...option.RequestOption) (r *BitcoinTransactionService) {
	r = &BitcoinTransactionService{}
	r.Options = opts
	return
}
