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

// EvmJsonRpcService contains methods and other services that help with interacting
// with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvmJsonRpcService] method instead.
type EvmJsonRpcService struct {
	Options []option.RequestOption
}

// NewEvmJsonRpcService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEvmJsonRpcService(opts ...option.RequestOption) (r *EvmJsonRpcService) {
	r = &EvmJsonRpcService{}
	r.Options = opts
	return
}

// Gets a json-rpc request and returns a full simulation indicating what will
// happen in the transaction together with a recommended action and some textual
// reasons of why the transaction was flagged that way.
func (r *EvmJsonRpcService) Scan(ctx context.Context, body EvmJsonRpcScanParams, opts ...option.RequestOption) (res *TransactionScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/evm/json-rpc/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EvmJsonRpcScanParams struct {
	// The chain name or chain ID
	Chain param.Field[EvmJsonRpcScanParamsChainUnion] `json:"chain,required"`
	// JSON-RPC request that was received by the wallet.
	Data param.Field[EvmJsonRpcScanParamsData] `json:"data,required"`
	// Object of additional information to validate against.
	Metadata param.Field[MetadataParam] `json:"metadata,required"`
	// The address of the account (wallet) received the request in hex string format
	AccountAddress param.Field[string] `json:"account_address"`
	// list of one or both of options for the desired output. "simulation" - include
	// simulation output in your response. "validation" - include security validation
	// of the transaction in your response. Default is ["validation"]
	Options param.Field[[]EvmJsonRpcScanParamsOption] `json:"options"`
}

func (r EvmJsonRpcScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The chain name or chain ID
//
// Satisfied by [TransactionScanSupportedChain], [shared.UnionString].
type EvmJsonRpcScanParamsChainUnion interface {
	ImplementsEvmJsonRpcScanParamsChainUnion()
}

// JSON-RPC request that was received by the wallet.
type EvmJsonRpcScanParamsData struct {
	// The method of the JSON-RPC request
	Method param.Field[string] `json:"method,required"`
	// The parameters of the JSON-RPC request in JSON format
	Params param.Field[[]interface{}] `json:"params,required"`
}

func (r EvmJsonRpcScanParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An enumeration.
type EvmJsonRpcScanParamsOption string

const (
	EvmJsonRpcScanParamsOptionValidation    EvmJsonRpcScanParamsOption = "validation"
	EvmJsonRpcScanParamsOptionSimulation    EvmJsonRpcScanParamsOption = "simulation"
	EvmJsonRpcScanParamsOptionGasEstimation EvmJsonRpcScanParamsOption = "gas_estimation"
	EvmJsonRpcScanParamsOptionEvents        EvmJsonRpcScanParamsOption = "events"
)

func (r EvmJsonRpcScanParamsOption) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanParamsOptionValidation, EvmJsonRpcScanParamsOptionSimulation, EvmJsonRpcScanParamsOptionGasEstimation, EvmJsonRpcScanParamsOptionEvents:
		return true
	}
	return false
}
