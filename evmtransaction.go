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

// EvmTransactionService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvmTransactionService] method instead.
type EvmTransactionService struct {
	Options []option.RequestOption
}

// NewEvmTransactionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEvmTransactionService(opts ...option.RequestOption) (r *EvmTransactionService) {
	r = &EvmTransactionService{}
	r.Options = opts
	return
}

// Gets a transaction and returns a full simulation indicating what will happen in
// the transaction together with a recommended action and some textual reasons of
// why the transaction was flagged that way.
func (r *EvmTransactionService) Scan(ctx context.Context, body EvmTransactionScanParams, opts ...option.RequestOption) (res *TransactionScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/evm/transaction/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EvmTransactionScanParams struct {
	// The address to relate the transaction to. Account address determines in which
	// perspective the transaction is simulated and validated.
	AccountAddress param.Field[string] `json:"account_address,required"`
	// The chain name or chain ID
	Chain param.Field[EvmTransactionScanParamsChainUnion] `json:"chain,required"`
	// Transaction parameters
	Data param.Field[EvmTransactionScanParamsData] `json:"data,required"`
	// Object of additional information to validate against.
	Metadata param.Field[MetadataParam] `json:"metadata,required"`
	// list of one or both of options for the desired output. "simulation" - include
	// simulation output in your response. "validation" - include security validation
	// of the transaction in your response. Default is ["validation"]
	Options param.Field[[]EvmTransactionScanParamsOption] `json:"options"`
}

func (r EvmTransactionScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The chain name or chain ID
//
// Satisfied by [TransactionScanSupportedChain], [shared.UnionString].
type EvmTransactionScanParamsChainUnion interface {
	ImplementsEvmTransactionScanParamsChainUnion()
}

// Transaction parameters
type EvmTransactionScanParamsData struct {
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

func (r EvmTransactionScanParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An enumeration.
type EvmTransactionScanParamsOption string

const (
	EvmTransactionScanParamsOptionValidation    EvmTransactionScanParamsOption = "validation"
	EvmTransactionScanParamsOptionSimulation    EvmTransactionScanParamsOption = "simulation"
	EvmTransactionScanParamsOptionGasEstimation EvmTransactionScanParamsOption = "gas_estimation"
	EvmTransactionScanParamsOptionEvents        EvmTransactionScanParamsOption = "events"
)

func (r EvmTransactionScanParamsOption) IsKnown() bool {
	switch r {
	case EvmTransactionScanParamsOptionValidation, EvmTransactionScanParamsOptionSimulation, EvmTransactionScanParamsOptionGasEstimation, EvmTransactionScanParamsOptionEvents:
		return true
	}
	return false
}
