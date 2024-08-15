// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"net/http"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/requestconfig"
	"github.com/blockaid-official/blockaid-client-go/option"
)

// StellarTransactionService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStellarTransactionService] method instead.
type StellarTransactionService struct {
	Options []option.RequestOption
}

// NewStellarTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStellarTransactionService(opts ...option.RequestOption) (r *StellarTransactionService) {
	r = &StellarTransactionService{}
	r.Options = opts
	return
}

// Scan Transaction
func (r *StellarTransactionService) Scan(ctx context.Context, body StellarTransactionScanParams, opts ...option.RequestOption) (res *StellarTransactionScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/stellar/transaction/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type StellarTransactionScanParams struct {
	StellarTransactionScanRequest StellarTransactionScanRequestParam `json:"StellarTransactionScanRequest,required"`
}

func (r StellarTransactionScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.StellarTransactionScanRequest)
}
