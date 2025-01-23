// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"net/http"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/requestconfig"
	"github.com/blockaid-official/blockaid-client-go/option"
)

// SuiTransactionService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSuiTransactionService] method instead.
type SuiTransactionService struct {
	Options []option.RequestOption
}

// NewSuiTransactionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSuiTransactionService(opts ...option.RequestOption) (r *SuiTransactionService) {
	r = &SuiTransactionService{}
	r.Options = opts
	return
}

// Scan Transaction
func (r *SuiTransactionService) Scan(ctx context.Context, body SuiTransactionScanParams, opts ...option.RequestOption) (res *SuiTransactionScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/sui/transaction/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SuiTransactionScanParams struct {
	SuiTransactionScanRequest SuiTransactionScanRequestParam `json:"SuiTransactionScanRequest,required"`
}

func (r SuiTransactionScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SuiTransactionScanRequest)
}
