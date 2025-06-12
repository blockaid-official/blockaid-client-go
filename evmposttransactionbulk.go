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

// EvmPostTransactionBulkService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvmPostTransactionBulkService] method instead.
type EvmPostTransactionBulkService struct {
	Options []option.RequestOption
}

// NewEvmPostTransactionBulkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEvmPostTransactionBulkService(opts ...option.RequestOption) (r *EvmPostTransactionBulkService) {
	r = &EvmPostTransactionBulkService{}
	r.Options = opts
	return
}

// Scan transactions that were already executed on chain, returns validation with
// features indicating address poisoning entites and malicious operators.
func (r *EvmPostTransactionBulkService) Scan(ctx context.Context, body EvmPostTransactionBulkScanParams, opts ...option.RequestOption) (res *[]TransactionScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/evm/post-transaction-bulk/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EvmPostTransactionBulkScanParams struct {
	// The chain name or chain ID
	Chain param.Field[TransactionScanSupportedChain] `json:"chain,required"`
	// Transaction hashes to scan
	Data param.Field[[]string] `json:"data,required"`
	// Object of additional information to validate against.
	Metadata param.Field[EvmPostTransactionBulkScanParamsMetadata] `json:"metadata,required"`
	// The relative block for the block validation. Can be "latest" or a block number.
	Block param.Field[EvmPostTransactionBulkScanParamsBlockUnion] `json:"block"`
	// List of one or both of options for the desired output. "simulation" - include
	// simulation output in your response. "validation" - include security validation
	// of the transaction in your response. Default is ["validation"]
	Options param.Field[[]EvmPostTransactionBulkScanParamsOption] `json:"options"`
	// Override the state of the chain. This is useful for testing purposes.
	StateOverride param.Field[map[string]EvmPostTransactionBulkScanParamsStateOverride] `json:"state_override"`
}

func (r EvmPostTransactionBulkScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Object of additional information to validate against.
type EvmPostTransactionBulkScanParamsMetadata struct {
	// cross reference transaction against the domain.
	Domain param.Field[string] `json:"domain,required"`
}

func (r EvmPostTransactionBulkScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The relative block for the block validation. Can be "latest" or a block number.
//
// Satisfied by [shared.UnionInt], [shared.UnionString].
type EvmPostTransactionBulkScanParamsBlockUnion interface {
	ImplementsEvmPostTransactionBulkScanParamsBlockUnion()
}

// An enumeration.
type EvmPostTransactionBulkScanParamsOption string

const (
	EvmPostTransactionBulkScanParamsOptionValidation    EvmPostTransactionBulkScanParamsOption = "validation"
	EvmPostTransactionBulkScanParamsOptionSimulation    EvmPostTransactionBulkScanParamsOption = "simulation"
	EvmPostTransactionBulkScanParamsOptionGasEstimation EvmPostTransactionBulkScanParamsOption = "gas_estimation"
	EvmPostTransactionBulkScanParamsOptionEvents        EvmPostTransactionBulkScanParamsOption = "events"
)

func (r EvmPostTransactionBulkScanParamsOption) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanParamsOptionValidation, EvmPostTransactionBulkScanParamsOptionSimulation, EvmPostTransactionBulkScanParamsOptionGasEstimation, EvmPostTransactionBulkScanParamsOptionEvents:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanParamsStateOverride struct {
	// Fake balance to set for the account before executing the call.
	Balance param.Field[string] `json:"balance"`
	// Fake EVM bytecode to inject into the account before executing the call.
	Code param.Field[string] `json:"code"`
	// Moves precompile to given address
	MovePrecompileToAddress param.Field[string] `json:"movePrecompileToAddress"`
	// Fake nonce to set for the account before executing the call.
	Nonce param.Field[string] `json:"nonce"`
	// Fake key-value mapping to override all slots in the account storage before
	// executing the call.
	State param.Field[map[string]string] `json:"state"`
	// Fake key-value mapping to override individual slots in the account storage
	// before executing the call.
	StateDiff param.Field[map[string]string] `json:"stateDiff"`
}

func (r EvmPostTransactionBulkScanParamsStateOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
