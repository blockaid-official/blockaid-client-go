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

// Report for misclassification of a transaction.
func (r *EvmTransactionService) Report(ctx context.Context, body EvmTransactionReportParams, opts ...option.RequestOption) (res *EvmTransactionReportResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/evm/transaction/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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

type EvmTransactionReportResponse = interface{}

type EvmTransactionReportParams struct {
	// Details about the report.
	Details param.Field[string] `json:"details,required"`
	// The event type of the report. Could be FALSE_POSITIVE or FALSE_NEGATIVE.
	Event param.Field[EvmTransactionReportParamsEvent] `json:"event,required"`
	// The report parameters.
	Report param.Field[EvmTransactionReportParamsReportUnion] `json:"report,required"`
}

func (r EvmTransactionReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The event type of the report. Could be FALSE_POSITIVE or FALSE_NEGATIVE.
type EvmTransactionReportParamsEvent string

const (
	EvmTransactionReportParamsEventFalsePositive EvmTransactionReportParamsEvent = "FALSE_POSITIVE"
	EvmTransactionReportParamsEventFalseNegative EvmTransactionReportParamsEvent = "FALSE_NEGATIVE"
)

func (r EvmTransactionReportParamsEvent) IsKnown() bool {
	switch r {
	case EvmTransactionReportParamsEventFalsePositive, EvmTransactionReportParamsEventFalseNegative:
		return true
	}
	return false
}

// The report parameters.
type EvmTransactionReportParamsReport struct {
	Type      param.Field[EvmTransactionReportParamsReportType] `json:"type,required"`
	Params    param.Field[interface{}]                          `json:"params"`
	RequestID param.Field[string]                               `json:"request_id"`
}

func (r EvmTransactionReportParamsReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmTransactionReportParamsReport) implementsEvmTransactionReportParamsReportUnion() {}

// The report parameters.
//
// Satisfied by
// [EvmTransactionReportParamsReportParamReportTransactionReportParams],
// [EvmTransactionReportParamsReportRequestIDReport],
// [EvmTransactionReportParamsReport].
type EvmTransactionReportParamsReportUnion interface {
	implementsEvmTransactionReportParamsReportUnion()
}

type EvmTransactionReportParamsReportParamReportTransactionReportParams struct {
	Params param.Field[EvmTransactionReportParamsReportParamReportTransactionReportParamsParams] `json:"params,required"`
	Type   param.Field[EvmTransactionReportParamsReportParamReportTransactionReportParamsType]   `json:"type,required"`
}

func (r EvmTransactionReportParamsReportParamReportTransactionReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmTransactionReportParamsReportParamReportTransactionReportParams) implementsEvmTransactionReportParamsReportUnion() {
}

type EvmTransactionReportParamsReportParamReportTransactionReportParamsParams struct {
	// The address to relate the transaction to. Account address determines in which
	// perspective the transaction is simulated and validated.
	AccountAddress param.Field[string] `json:"account_address,required"`
	// The chain name
	Chain param.Field[TransactionScanSupportedChain] `json:"chain,required"`
	// Transaction parameters
	Data param.Field[EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataUnion] `json:"data,required"`
	// Object of additional information to validate against.
	Metadata param.Field[MetadataParam] `json:"metadata,required"`
}

func (r EvmTransactionReportParamsReportParamReportTransactionReportParamsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Transaction parameters
type EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsData struct {
	AuthorizationList param.Field[interface{}] `json:"authorization_list"`
	// The encoded call data of the transaction in hex string format
	Data param.Field[string] `json:"data"`
	// The source address of the transaction in hex string format
	From param.Field[string] `json:"from"`
	// The gas required for the transaction in hex string format.
	Gas param.Field[string] `json:"gas"`
	// The gas price for the transaction in hex string format.
	GasPrice param.Field[string] `json:"gas_price"`
	// The method of the JSON-RPC request
	Method param.Field[string]      `json:"method"`
	Params param.Field[interface{}] `json:"params"`
	// The destination address of the transaction in hex string format
	To param.Field[string] `json:"to"`
	// The value of the transaction in Wei in hex string format
	Value param.Field[string] `json:"value"`
}

func (r EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsData) implementsEvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataUnion() {
}

// Transaction parameters
//
// Satisfied by
// [EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataTransaction],
// [EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpc],
// [EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsData].
type EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataUnion interface {
	implementsEvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataUnion()
}

type EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataTransaction struct {
	// The source address of the transaction in hex string format
	From param.Field[string] `json:"from,required"`
	// The authorization list
	AuthorizationList param.Field[[]EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataTransactionAuthorizationList] `json:"authorization_list"`
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

func (r EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataTransaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataTransaction) implementsEvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataUnion() {
}

type EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataTransactionAuthorizationList struct {
	// The delegation designation address
	Address param.Field[string] `json:"address,required"`
	ChainID param.Field[string] `json:"chainId"`
	// The authority address of the delegation, should be provided when the signature
	// (r,s,yParity) is not provided in order to simulate the transaction with the
	// correct delegation
	Eoa     param.Field[string] `json:"eoa"`
	Nonce   param.Field[string] `json:"nonce"`
	R       param.Field[string] `json:"r"`
	S       param.Field[string] `json:"s"`
	YParity param.Field[string] `json:"yParity"`
}

func (r EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataTransactionAuthorizationList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpc struct {
	// The method of the JSON-RPC request
	Method param.Field[string] `json:"method,required"`
	// The parameters of the JSON-RPC request in JSON format
	Params param.Field[[]interface{}] `json:"params,required"`
}

func (r EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpc) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpc) implementsEvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataUnion() {
}

type EvmTransactionReportParamsReportParamReportTransactionReportParamsType string

const (
	EvmTransactionReportParamsReportParamReportTransactionReportParamsTypeParams EvmTransactionReportParamsReportParamReportTransactionReportParamsType = "params"
)

func (r EvmTransactionReportParamsReportParamReportTransactionReportParamsType) IsKnown() bool {
	switch r {
	case EvmTransactionReportParamsReportParamReportTransactionReportParamsTypeParams:
		return true
	}
	return false
}

type EvmTransactionReportParamsReportRequestIDReport struct {
	RequestID param.Field[string]                                              `json:"request_id,required"`
	Type      param.Field[EvmTransactionReportParamsReportRequestIDReportType] `json:"type,required"`
}

func (r EvmTransactionReportParamsReportRequestIDReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmTransactionReportParamsReportRequestIDReport) implementsEvmTransactionReportParamsReportUnion() {
}

type EvmTransactionReportParamsReportRequestIDReportType string

const (
	EvmTransactionReportParamsReportRequestIDReportTypeRequestID EvmTransactionReportParamsReportRequestIDReportType = "request_id"
)

func (r EvmTransactionReportParamsReportRequestIDReportType) IsKnown() bool {
	switch r {
	case EvmTransactionReportParamsReportRequestIDReportTypeRequestID:
		return true
	}
	return false
}

type EvmTransactionReportParamsReportType string

const (
	EvmTransactionReportParamsReportTypeParams    EvmTransactionReportParamsReportType = "params"
	EvmTransactionReportParamsReportTypeRequestID EvmTransactionReportParamsReportType = "request_id"
)

func (r EvmTransactionReportParamsReportType) IsKnown() bool {
	switch r {
	case EvmTransactionReportParamsReportTypeParams, EvmTransactionReportParamsReportTypeRequestID:
		return true
	}
	return false
}

type EvmTransactionScanParams struct {
	// The address to relate the transaction to. Account address determines in which
	// perspective the transaction is simulated and validated.
	AccountAddress param.Field[string] `json:"account_address,required"`
	// The chain name or chain ID
	Chain param.Field[TransactionScanSupportedChain] `json:"chain,required"`
	// Transaction parameters
	Data param.Field[EvmTransactionScanParamsData] `json:"data,required"`
	// Object of additional information to validate against.
	Metadata param.Field[MetadataParam] `json:"metadata,required"`
	// The relative block for the block validation. Can be "latest" or a block number.
	Block param.Field[EvmTransactionScanParamsBlockUnion] `json:"block"`
	// list of one or both of options for the desired output. "simulation" - include
	// simulation output in your response. "validation" - include security validation
	// of the transaction in your response. Default is ["validation"]
	Options param.Field[[]EvmTransactionScanParamsOption] `json:"options"`
	// Override the state of the chain. This is useful for testing purposes.
	StateOverride param.Field[map[string]EvmTransactionScanParamsStateOverride] `json:"state_override"`
}

func (r EvmTransactionScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Transaction parameters
type EvmTransactionScanParamsData struct {
	// The source address of the transaction in hex string format
	From param.Field[string] `json:"from,required"`
	// The authorization list
	AuthorizationList param.Field[[]EvmTransactionScanParamsDataAuthorizationList] `json:"authorization_list"`
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

type EvmTransactionScanParamsDataAuthorizationList struct {
	// The delegation designation address
	Address param.Field[string] `json:"address,required"`
	ChainID param.Field[string] `json:"chainId"`
	// The authority address of the delegation, should be provided when the signature
	// (r,s,yParity) is not provided in order to simulate the transaction with the
	// correct delegation
	Eoa     param.Field[string] `json:"eoa"`
	Nonce   param.Field[string] `json:"nonce"`
	R       param.Field[string] `json:"r"`
	S       param.Field[string] `json:"s"`
	YParity param.Field[string] `json:"yParity"`
}

func (r EvmTransactionScanParamsDataAuthorizationList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The relative block for the block validation. Can be "latest" or a block number.
//
// Satisfied by [shared.UnionInt], [shared.UnionString].
type EvmTransactionScanParamsBlockUnion interface {
	ImplementsEvmTransactionScanParamsBlockUnion()
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

type EvmTransactionScanParamsStateOverride struct {
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

func (r EvmTransactionScanParamsStateOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
