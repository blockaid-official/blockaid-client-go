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

// EvmPostTransactionService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvmPostTransactionService] method instead.
type EvmPostTransactionService struct {
	Options []option.RequestOption
}

// NewEvmPostTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEvmPostTransactionService(opts ...option.RequestOption) (r *EvmPostTransactionService) {
	r = &EvmPostTransactionService{}
	r.Options = opts
	return
}

// Scan a transaction that was already executed on chain, returns validation with
// features indicating address poisoning entites and malicious operators.
func (r *EvmPostTransactionService) Scan(ctx context.Context, body EvmPostTransactionScanParams, opts ...option.RequestOption) (res *TransactionScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/evm/post-transaction/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EvmPostTransactionScanParams struct {
	// The chain name or chain ID
	Chain param.Field[TransactionScanSupportedChain]    `json:"chain,required"`
	Data  param.Field[EvmPostTransactionScanParamsData] `json:"data,required"`
	// Object of additional information to validate against.
	Metadata param.Field[MetadataParam] `json:"metadata,required"`
	// The relative block for the block validation. Can be "latest", "earliest",
	// "pending" or a block number.
	Block param.Field[EvmPostTransactionScanParamsBlockUnion] `json:"block"`
	// list of one or both of options for the desired output. "simulation" - include
	// simulation output in your response. "validation" - include security validation
	// of the transaction in your response. Default is ["validation"]
	Options param.Field[[]EvmPostTransactionScanParamsOption] `json:"options"`
}

func (r EvmPostTransactionScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EvmPostTransactionScanParamsData struct {
	// The transaction hash to scan
	TxHash param.Field[string] `json:"tx_hash,required"`
}

func (r EvmPostTransactionScanParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The relative block for the block validation. Can be "latest", "earliest",
// "pending" or a block number.
//
// Satisfied by [shared.UnionInt], [shared.UnionString].
type EvmPostTransactionScanParamsBlockUnion interface {
	ImplementsEvmPostTransactionScanParamsBlockUnion()
}

// An enumeration.
type EvmPostTransactionScanParamsOption string

const (
	EvmPostTransactionScanParamsOptionValidation    EvmPostTransactionScanParamsOption = "validation"
	EvmPostTransactionScanParamsOptionSimulation    EvmPostTransactionScanParamsOption = "simulation"
	EvmPostTransactionScanParamsOptionGasEstimation EvmPostTransactionScanParamsOption = "gas_estimation"
	EvmPostTransactionScanParamsOptionEvents        EvmPostTransactionScanParamsOption = "events"
)

func (r EvmPostTransactionScanParamsOption) IsKnown() bool {
	switch r {
	case EvmPostTransactionScanParamsOptionValidation, EvmPostTransactionScanParamsOptionSimulation, EvmPostTransactionScanParamsOptionGasEstimation, EvmPostTransactionScanParamsOptionEvents:
		return true
	}
	return false
}
