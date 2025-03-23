// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"net/http"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/requestconfig"
	"github.com/blockaid-official/blockaid-client-go/option"
)

// SolanaAddressService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSolanaAddressService] method instead.
type SolanaAddressService struct {
	Options []option.RequestOption
}

// NewSolanaAddressService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSolanaAddressService(opts ...option.RequestOption) (r *SolanaAddressService) {
	r = &SolanaAddressService{}
	r.Options = opts
	return
}

// Gets an address and returns a full security assessment indicating weather or not
// this address is malicious as well as textual reasons of why the address was
// flagged that way.
func (r *SolanaAddressService) Scan(ctx context.Context, body SolanaAddressScanParams, opts ...option.RequestOption) (res *AddressScanResponseSchema, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/solana/address/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SolanaAddressScanParams struct {
	AddressScanRequestSchema AddressScanRequestSchemaParam `json:"AddressScanRequestSchema,required"`
}

func (r SolanaAddressScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.AddressScanRequestSchema)
}
