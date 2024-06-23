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
	// Transaction bulk parameters
	Data param.Field[[]EvmPostTransactionBulkScanParamsData] `json:"data,required"`
	// Object of additional information to validate against.
	Metadata param.Field[MetadataParam] `json:"metadata,required"`
	// List of one or both of options for the desired output. "simulation" - include
	// simulation output in your response. "validation" - include security validation
	// of the transaction in your response. Default is ["validation"]
	Options param.Field[[]EvmPostTransactionBulkScanParamsOption] `json:"options"`
}

func (r EvmPostTransactionBulkScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EvmPostTransactionBulkScanParamsData struct {
	// The source address of the transaction in hex string format
	From param.Field[string] `json:"from,required"`
	// The encoded call data of the transaction in hex string format
	Data param.Field[string] `json:"data"`
	// The gas required for the transaction in hex string format.
	Gas param.Field[string] `json:"gas"`
	// The gas price for the transaction in hex string format.
	GasPrice param.Field[string] `json:"gas_price"`
	// The destination address of the transaction in hex string format
	To param.Field[string] `json:"to"`
	// The value of the transaction in Wei in hex string format
	Value param.Field[string] `json:"value"`
}

func (r EvmPostTransactionBulkScanParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
