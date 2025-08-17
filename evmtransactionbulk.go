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
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvmTransactionBulkService] method instead.
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
func (r *EvmTransactionBulkService) Scan(ctx context.Context, body EvmTransactionBulkScanParams, opts ...option.RequestOption) (res *[]TransactionScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/evm/transaction-bulk/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EvmTransactionBulkScanParams struct {
	// The chain name or chain ID
	Chain param.Field[TransactionScanSupportedChain] `json:"chain,required"`
	// Transaction bulk parameters
	Data param.Field[[]EvmTransactionBulkScanParamsData] `json:"data,required"`
	// Object of additional information to validate against.
	Metadata param.Field[EvmTransactionBulkScanParamsMetadata] `json:"metadata,required"`
	// Should aggregate the results to one result
	Aggregated param.Field[bool] `json:"aggregated"`
	// The relative block for the block validation. Can be "latest" or a block number.
	Block param.Field[EvmTransactionBulkScanParamsBlockUnion] `json:"block"`
	// List of one or both of options for the desired output. "simulation" - include
	// simulation output in your response. "validation" - include security validation
	// of the transaction in your response. Default is ["validation"]
	Options param.Field[[]EvmTransactionBulkScanParamsOption] `json:"options"`
	// Simulate transactions using gas estimation result. This requires
	// "gas_estimation" option to be enabled.
	SimulateWithEstimatedGas param.Field[bool] `json:"simulate_with_estimated_gas"`
	// Override the state of the chain. This is useful for testing purposes.
	StateOverride param.Field[map[string]EvmTransactionBulkScanParamsStateOverride] `json:"state_override"`
}

func (r EvmTransactionBulkScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EvmTransactionBulkScanParamsData struct {
	// The source address of the transaction in hex string format
	From param.Field[string] `json:"from,required"`
	// The authorization list
	AuthorizationList param.Field[[]EvmTransactionBulkScanParamsDataAuthorizationList] `json:"authorization_list"`
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

type EvmTransactionBulkScanParamsDataAuthorizationList struct {
	// The delegation designation address
	Address param.Field[string] `json:"address,required"`
	// The chain ID as hex string
	ChainID param.Field[string] `json:"chainId"`
	// The authority address of the delegation, should be provided when the signature
	// (r,s,yParity) is not provided in order to simulate the transaction with the
	// correct delegation
	Eoa param.Field[string] `json:"eoa"`
	// The nonce value as hex string
	Nonce param.Field[string] `json:"nonce"`
	// The r value as hex string (excluding leading 0 digits)
	R param.Field[string] `json:"r"`
	// The s value as hex string (excluding leading 0 digits)
	S param.Field[string] `json:"s"`
	// The yParity value as hex string
	YParity param.Field[string] `json:"yParity"`
}

func (r EvmTransactionBulkScanParamsDataAuthorizationList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Object of additional information to validate against.
type EvmTransactionBulkScanParamsMetadata struct {
	// cross reference transaction against the domain.
	Domain param.Field[string] `json:"domain,required"`
}

func (r EvmTransactionBulkScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The relative block for the block validation. Can be "latest" or a block number.
//
// Satisfied by [shared.UnionInt], [shared.UnionString].
type EvmTransactionBulkScanParamsBlockUnion interface {
	ImplementsEvmTransactionBulkScanParamsBlockUnion()
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

type EvmTransactionBulkScanParamsStateOverride struct {
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

func (r EvmTransactionBulkScanParamsStateOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
