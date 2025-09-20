// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/param"
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

// Gets a transaction and returns a full simulation indicating what will happen in
// the transaction together with a recommended action and some textual reasons of
// why the transaction was flagged that way.
func (r *SuiTransactionService) Scan(ctx context.Context, body SuiTransactionScanParams, opts ...option.RequestOption) (res *SuiTransactionScanResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/sui/transaction/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SuiTransactionScanParams struct {
	AccountAddress param.Field[interface{}]                   `json:"account_address,required"`
	Chain          param.Field[SuiTransactionScanParamsChain] `json:"chain,required"`
	// Metadata
	Metadata    param.Field[SuiTransactionScanParamsMetadataUnion] `json:"metadata,required"`
	Transaction param.Field[string]                                `json:"transaction,required"`
	// List of options to include in the response
	//
	// - `Options.validation`: Include Options.validation output in the response
	//
	// - `Options.simulation`: Include Options.simulation output in the response
	Options param.Field[[]SuiTransactionScanParamsOption] `json:"options"`
}

func (r SuiTransactionScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuiTransactionScanParamsChain string

const (
	SuiTransactionScanParamsChainMainnet SuiTransactionScanParamsChain = "mainnet"
	SuiTransactionScanParamsChainTestnet SuiTransactionScanParamsChain = "testnet"
	SuiTransactionScanParamsChainDevnet  SuiTransactionScanParamsChain = "devnet"
)

func (r SuiTransactionScanParamsChain) IsKnown() bool {
	switch r {
	case SuiTransactionScanParamsChainMainnet, SuiTransactionScanParamsChainTestnet, SuiTransactionScanParamsChainDevnet:
		return true
	}
	return false
}

// Metadata
type SuiTransactionScanParamsMetadata struct {
	// Metadata for wallet requests
	Type param.Field[SuiTransactionScanParamsMetadataType] `json:"type"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url"`
}

func (r SuiTransactionScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuiTransactionScanParamsMetadata) implementsSuiTransactionScanParamsMetadataUnion() {}

// Metadata
//
// Satisfied by [SuiTransactionScanParamsMetadataSuiWalletRequestMetadata],
// [SuiTransactionScanParamsMetadataSuiInAppRequestMetadata],
// [SuiTransactionScanParamsMetadata].
type SuiTransactionScanParamsMetadataUnion interface {
	implementsSuiTransactionScanParamsMetadataUnion()
}

type SuiTransactionScanParamsMetadataSuiWalletRequestMetadata struct {
	// Metadata for wallet requests
	Type param.Field[SuiTransactionScanParamsMetadataSuiWalletRequestMetadataType] `json:"type,required"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url,required"`
}

func (r SuiTransactionScanParamsMetadataSuiWalletRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuiTransactionScanParamsMetadataSuiWalletRequestMetadata) implementsSuiTransactionScanParamsMetadataUnion() {
}

// Metadata for wallet requests
type SuiTransactionScanParamsMetadataSuiWalletRequestMetadataType string

const (
	SuiTransactionScanParamsMetadataSuiWalletRequestMetadataTypeWallet SuiTransactionScanParamsMetadataSuiWalletRequestMetadataType = "wallet"
)

func (r SuiTransactionScanParamsMetadataSuiWalletRequestMetadataType) IsKnown() bool {
	switch r {
	case SuiTransactionScanParamsMetadataSuiWalletRequestMetadataTypeWallet:
		return true
	}
	return false
}

type SuiTransactionScanParamsMetadataSuiInAppRequestMetadata struct {
	// Metadata for in-app requests
	Type param.Field[SuiTransactionScanParamsMetadataSuiInAppRequestMetadataType] `json:"type"`
}

func (r SuiTransactionScanParamsMetadataSuiInAppRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuiTransactionScanParamsMetadataSuiInAppRequestMetadata) implementsSuiTransactionScanParamsMetadataUnion() {
}

// Metadata for in-app requests
type SuiTransactionScanParamsMetadataSuiInAppRequestMetadataType string

const (
	SuiTransactionScanParamsMetadataSuiInAppRequestMetadataTypeInApp SuiTransactionScanParamsMetadataSuiInAppRequestMetadataType = "in_app"
)

func (r SuiTransactionScanParamsMetadataSuiInAppRequestMetadataType) IsKnown() bool {
	switch r {
	case SuiTransactionScanParamsMetadataSuiInAppRequestMetadataTypeInApp:
		return true
	}
	return false
}

// Metadata for wallet requests
type SuiTransactionScanParamsMetadataType string

const (
	SuiTransactionScanParamsMetadataTypeWallet SuiTransactionScanParamsMetadataType = "wallet"
	SuiTransactionScanParamsMetadataTypeInApp  SuiTransactionScanParamsMetadataType = "in_app"
)

func (r SuiTransactionScanParamsMetadataType) IsKnown() bool {
	switch r {
	case SuiTransactionScanParamsMetadataTypeWallet, SuiTransactionScanParamsMetadataTypeInApp:
		return true
	}
	return false
}

type SuiTransactionScanParamsOption string

const (
	SuiTransactionScanParamsOptionValidation SuiTransactionScanParamsOption = "validation"
	SuiTransactionScanParamsOptionSimulation SuiTransactionScanParamsOption = "simulation"
)

func (r SuiTransactionScanParamsOption) IsKnown() bool {
	switch r {
	case SuiTransactionScanParamsOptionValidation, SuiTransactionScanParamsOptionSimulation:
		return true
	}
	return false
}
