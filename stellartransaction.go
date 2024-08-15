// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"net/http"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/param"
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
	AccountAddress param.Field[string] `json:"account_address,required"`
	// A CAIP-2 chain ID or a Stellar network name
	Chain param.Field[StellarTransactionScanParamsChain] `json:"chain,required"`
	// Metadata
	Metadata    param.Field[StellarTransactionScanParamsMetadataUnion] `json:"metadata,required"`
	Transaction param.Field[string]                                    `json:"transaction,required"`
	// List of options to include in the response
	//
	// - `simulation`: Include simulation output in the response
	// - `validation`: Include security validation of the transaction in the response
	Options param.Field[[]StellarTransactionScanParamsOption] `json:"options"`
}

func (r StellarTransactionScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A CAIP-2 chain ID or a Stellar network name
type StellarTransactionScanParamsChain string

const (
	StellarTransactionScanParamsChainPubnet    StellarTransactionScanParamsChain = "pubnet"
	StellarTransactionScanParamsChainFuturenet StellarTransactionScanParamsChain = "futurenet"
	StellarTransactionScanParamsChainTestnet   StellarTransactionScanParamsChain = "testnet"
)

func (r StellarTransactionScanParamsChain) IsKnown() bool {
	switch r {
	case StellarTransactionScanParamsChainPubnet, StellarTransactionScanParamsChainFuturenet, StellarTransactionScanParamsChainTestnet:
		return true
	}
	return false
}

// Metadata
type StellarTransactionScanParamsMetadata struct {
	// Metadata for wallet requests
	Type param.Field[StellarTransactionScanParamsMetadataType] `json:"type,required"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url"`
}

func (r StellarTransactionScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionScanParamsMetadata) implementsStellarTransactionScanParamsMetadataUnion() {}

// Metadata
//
// Satisfied by [StellarTransactionScanParamsMetadataStellarWalletRequestMetadata],
// [StellarTransactionScanParamsMetadataStellarInAppRequestMetadata],
// [StellarTransactionScanParamsMetadata].
type StellarTransactionScanParamsMetadataUnion interface {
	implementsStellarTransactionScanParamsMetadataUnion()
}

type StellarTransactionScanParamsMetadataStellarWalletRequestMetadata struct {
	// Metadata for wallet requests
	Type param.Field[StellarTransactionScanParamsMetadataStellarWalletRequestMetadataType] `json:"type,required"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url,required"`
}

func (r StellarTransactionScanParamsMetadataStellarWalletRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionScanParamsMetadataStellarWalletRequestMetadata) implementsStellarTransactionScanParamsMetadataUnion() {
}

// Metadata for wallet requests
type StellarTransactionScanParamsMetadataStellarWalletRequestMetadataType string

const (
	StellarTransactionScanParamsMetadataStellarWalletRequestMetadataTypeWallet StellarTransactionScanParamsMetadataStellarWalletRequestMetadataType = "wallet"
)

func (r StellarTransactionScanParamsMetadataStellarWalletRequestMetadataType) IsKnown() bool {
	switch r {
	case StellarTransactionScanParamsMetadataStellarWalletRequestMetadataTypeWallet:
		return true
	}
	return false
}

type StellarTransactionScanParamsMetadataStellarInAppRequestMetadata struct {
	// Metadata for in-app requests
	Type param.Field[StellarTransactionScanParamsMetadataStellarInAppRequestMetadataType] `json:"type,required"`
}

func (r StellarTransactionScanParamsMetadataStellarInAppRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionScanParamsMetadataStellarInAppRequestMetadata) implementsStellarTransactionScanParamsMetadataUnion() {
}

// Metadata for in-app requests
type StellarTransactionScanParamsMetadataStellarInAppRequestMetadataType string

const (
	StellarTransactionScanParamsMetadataStellarInAppRequestMetadataTypeInApp StellarTransactionScanParamsMetadataStellarInAppRequestMetadataType = "in_app"
)

func (r StellarTransactionScanParamsMetadataStellarInAppRequestMetadataType) IsKnown() bool {
	switch r {
	case StellarTransactionScanParamsMetadataStellarInAppRequestMetadataTypeInApp:
		return true
	}
	return false
}

// Metadata for wallet requests
type StellarTransactionScanParamsMetadataType string

const (
	StellarTransactionScanParamsMetadataTypeWallet StellarTransactionScanParamsMetadataType = "wallet"
	StellarTransactionScanParamsMetadataTypeInApp  StellarTransactionScanParamsMetadataType = "in_app"
)

func (r StellarTransactionScanParamsMetadataType) IsKnown() bool {
	switch r {
	case StellarTransactionScanParamsMetadataTypeWallet, StellarTransactionScanParamsMetadataTypeInApp:
		return true
	}
	return false
}

type StellarTransactionScanParamsOption string

const (
	StellarTransactionScanParamsOptionValidation StellarTransactionScanParamsOption = "validation"
	StellarTransactionScanParamsOptionSimulation StellarTransactionScanParamsOption = "simulation"
)

func (r StellarTransactionScanParamsOption) IsKnown() bool {
	switch r {
	case StellarTransactionScanParamsOptionValidation, StellarTransactionScanParamsOptionSimulation:
		return true
	}
	return false
}
