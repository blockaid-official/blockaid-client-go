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

// EvmTransactionBulkService contains methods and other services that help with
// interacting with the blockaid API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewEvmTransactionBulkService] method
// instead.
type EvmTransactionBulkService struct {
	Options []option.RequestOption
}

// NewEvmTransactionBulkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEvmTransactionBulkService(opts ...option.RequestOption) (r *EvmTransactionBulkService) {
	r = &EvmTransactionBulkService{}
	r.Options = opts
	return
}

// Gets a bulk of transactions and returns a simulation showcasing the outcome
// after executing the transactions synchronously, along with a suggested course of
// action and textual explanations highlighting the reasons for flagging the bulk
// in that manner.
func (r *EvmTransactionBulkService) Scan(ctx context.Context, body EvmTransactionBulkScanParams, opts ...option.RequestOption) (res *[]TransactionBulkResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "evm/transaction-bulk/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EvmTransactionBulkScanParams struct {
	// The chain name
	Chain param.Field[Chain] `json:"chain,required"`
	// Transaction bulk parameters
	Data param.Field[[]EvmTransactionBulkScanParamsData] `json:"data,required"`
	// Object of additional information to validate against.
	Metadata param.Field[MetadataParam] `json:"metadata,required"`
	// List of one or both of options for the desired output. "simulation" - include
	// simulation output in your response. "validation" - include security validation
	// of the transaction in your response. Default is ["validation"]
	Options param.Field[[]EvmTransactionBulkScanParamsOption] `json:"options"`
}

func (r EvmTransactionBulkScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EvmTransactionBulkScanParamsData struct {
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

func (r EvmTransactionBulkScanParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An enumeration.
type EvmTransactionBulkScanParamsOption string

const (
	EvmTransactionBulkScanParamsOptionValidation    EvmTransactionBulkScanParamsOption = "validation"
	EvmTransactionBulkScanParamsOptionSimulation    EvmTransactionBulkScanParamsOption = "simulation"
	EvmTransactionBulkScanParamsOptionGasEstimation EvmTransactionBulkScanParamsOption = "gas_estimation"
	EvmTransactionBulkScanParamsOptionEvents        EvmTransactionBulkScanParamsOption = "events"
)

func (r EvmTransactionBulkScanParamsOption) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanParamsOptionValidation, EvmTransactionBulkScanParamsOptionSimulation, EvmTransactionBulkScanParamsOptionGasEstimation, EvmTransactionBulkScanParamsOptionEvents:
		return true
	}
	return false
}
