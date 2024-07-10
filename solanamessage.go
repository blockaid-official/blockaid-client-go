// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"net/http"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/requestconfig"
	"github.com/blockaid-official/blockaid-client-go/option"
)

// SolanaMessageService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSolanaMessageService] method instead.
type SolanaMessageService struct {
	Options []option.RequestOption
}

// NewSolanaMessageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSolanaMessageService(opts ...option.RequestOption) (r *SolanaMessageService) {
	r = &SolanaMessageService{}
	r.Options = opts
	return
}

// Scan Messages
func (r *SolanaMessageService) Scan(ctx context.Context, body SolanaMessageScanParams, opts ...option.RequestOption) (res *ResponseSchema, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/solana/message/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SolanaMessageScanParams struct {
	TxScanRequestSchema TxScanRequestSchemaParam `json:"TxScanRequestSchema,required"`
}

func (r SolanaMessageScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.TxScanRequestSchema)
}
