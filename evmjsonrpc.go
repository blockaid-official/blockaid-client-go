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
	Chain param.Field[TransactionScanSupportedChain] `json:"chain,required"`
	// JSON-RPC request that was received by the wallet.
	Data param.Field[EvmJsonRpcScanParamsData] `json:"data,required"`
	// Object of additional information to validate against.
	Metadata param.Field[EvmJsonRpcScanParamsMetadata] `json:"metadata,required"`
	// The address of the account (wallet) received the request in hex string format
	AccountAddress param.Field[string] `json:"account_address"`
	// The relative block for the block validation. Can be "latest" or a block number.
	Block param.Field[EvmJsonRpcScanParamsBlockUnion] `json:"block"`
	// List of one or more of options for the desired output. "simulation" - include
	// simulation output in your response. "validation" - include security validation
	// of the transaction in your response. "gas_estimation" - include gas estimation
	// result in your response. Default is ["validation"]
	Options param.Field[[]EvmJsonRpcScanParamsOption] `json:"options"`
	// Simulate transactions using gas estimation result. This requires
	// "gas_estimation" option to be enabled.
	SimulateWithEstimatedGas param.Field[bool] `json:"simulate_with_estimated_gas"`
	// Override the state of the chain. This is useful for testing purposes.
	StateOverride param.Field[map[string]EvmJsonRpcScanParamsStateOverride] `json:"state_override"`
}

func (r EvmJsonRpcScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

// Object of additional information to validate against.
type EvmJsonRpcScanParamsMetadata struct {
	// cross reference transaction against the domain.
	Domain param.Field[string] `json:"domain,required"`
}

func (r EvmJsonRpcScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The relative block for the block validation. Can be "latest" or a block number.
//
// Satisfied by [shared.UnionInt], [shared.UnionString].
type EvmJsonRpcScanParamsBlockUnion interface {
	ImplementsEvmJsonRpcScanParamsBlockUnion()
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

type EvmJsonRpcScanParamsStateOverride struct {
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

func (r EvmJsonRpcScanParamsStateOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
