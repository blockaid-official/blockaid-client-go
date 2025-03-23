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

// Report for misclassification of an EVM post transaction.
func (r *EvmPostTransactionService) Report(ctx context.Context, body EvmPostTransactionReportParams, opts ...option.RequestOption) (res *EvmPostTransactionReportResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/evm/post-transaction/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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

type EvmPostTransactionReportResponse = interface{}

type EvmPostTransactionReportParams struct {
	// Details about the report.
	Details param.Field[string] `json:"details,required"`
	// The event type of the report. Could be FALSE_POSITIVE or FALSE_NEGATIVE.
	Event param.Field[EvmPostTransactionReportParamsEvent] `json:"event,required"`
	// The report parameters.
	Report param.Field[EvmPostTransactionReportParamsReportUnion] `json:"report,required"`
}

func (r EvmPostTransactionReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The event type of the report. Could be FALSE_POSITIVE or FALSE_NEGATIVE.
type EvmPostTransactionReportParamsEvent string

const (
	EvmPostTransactionReportParamsEventFalsePositive EvmPostTransactionReportParamsEvent = "FALSE_POSITIVE"
	EvmPostTransactionReportParamsEventFalseNegative EvmPostTransactionReportParamsEvent = "FALSE_NEGATIVE"
)

func (r EvmPostTransactionReportParamsEvent) IsKnown() bool {
	switch r {
	case EvmPostTransactionReportParamsEventFalsePositive, EvmPostTransactionReportParamsEventFalseNegative:
		return true
	}
	return false
}

// The report parameters.
type EvmPostTransactionReportParamsReport struct {
	Type      param.Field[EvmPostTransactionReportParamsReportType] `json:"type,required"`
	Params    param.Field[interface{}]                              `json:"params"`
	RequestID param.Field[string]                                   `json:"request_id"`
}

func (r EvmPostTransactionReportParamsReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmPostTransactionReportParamsReport) implementsEvmPostTransactionReportParamsReportUnion() {}

// The report parameters.
//
// Satisfied by
// [EvmPostTransactionReportParamsReportParamReportChainTransactionHashParams],
// [EvmPostTransactionReportParamsReportRequestIDReport],
// [EvmPostTransactionReportParamsReport].
type EvmPostTransactionReportParamsReportUnion interface {
	implementsEvmPostTransactionReportParamsReportUnion()
}

type EvmPostTransactionReportParamsReportParamReportChainTransactionHashParams struct {
	Params param.Field[EvmPostTransactionReportParamsReportParamReportChainTransactionHashParamsParams] `json:"params,required"`
	Type   param.Field[EvmPostTransactionReportParamsReportParamReportChainTransactionHashParamsType]   `json:"type,required"`
}

func (r EvmPostTransactionReportParamsReportParamReportChainTransactionHashParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmPostTransactionReportParamsReportParamReportChainTransactionHashParams) implementsEvmPostTransactionReportParamsReportUnion() {
}

type EvmPostTransactionReportParamsReportParamReportChainTransactionHashParamsParams struct {
	// The chain name
	Chain param.Field[TransactionScanSupportedChain] `json:"chain,required"`
	// The transaction hash to report on.
	TxHash param.Field[string] `json:"tx_hash,required"`
}

func (r EvmPostTransactionReportParamsReportParamReportChainTransactionHashParamsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EvmPostTransactionReportParamsReportParamReportChainTransactionHashParamsType string

const (
	EvmPostTransactionReportParamsReportParamReportChainTransactionHashParamsTypeParams EvmPostTransactionReportParamsReportParamReportChainTransactionHashParamsType = "params"
)

func (r EvmPostTransactionReportParamsReportParamReportChainTransactionHashParamsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionReportParamsReportParamReportChainTransactionHashParamsTypeParams:
		return true
	}
	return false
}

type EvmPostTransactionReportParamsReportRequestIDReport struct {
	RequestID param.Field[string]                                                  `json:"request_id,required"`
	Type      param.Field[EvmPostTransactionReportParamsReportRequestIDReportType] `json:"type,required"`
}

func (r EvmPostTransactionReportParamsReportRequestIDReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmPostTransactionReportParamsReportRequestIDReport) implementsEvmPostTransactionReportParamsReportUnion() {
}

type EvmPostTransactionReportParamsReportRequestIDReportType string

const (
	EvmPostTransactionReportParamsReportRequestIDReportTypeRequestID EvmPostTransactionReportParamsReportRequestIDReportType = "request_id"
)

func (r EvmPostTransactionReportParamsReportRequestIDReportType) IsKnown() bool {
	switch r {
	case EvmPostTransactionReportParamsReportRequestIDReportTypeRequestID:
		return true
	}
	return false
}

type EvmPostTransactionReportParamsReportType string

const (
	EvmPostTransactionReportParamsReportTypeParams    EvmPostTransactionReportParamsReportType = "params"
	EvmPostTransactionReportParamsReportTypeRequestID EvmPostTransactionReportParamsReportType = "request_id"
)

func (r EvmPostTransactionReportParamsReportType) IsKnown() bool {
	switch r {
	case EvmPostTransactionReportParamsReportTypeParams, EvmPostTransactionReportParamsReportTypeRequestID:
		return true
	}
	return false
}

type EvmPostTransactionScanParams struct {
	// The chain name or chain ID
	Chain param.Field[TransactionScanSupportedChain]    `json:"chain,required"`
	Data  param.Field[EvmPostTransactionScanParamsData] `json:"data,required"`
	// Object of additional information to validate against.
	Metadata param.Field[MetadataParam] `json:"metadata,required"`
	// The relative block for the block validation. Can be "latest" or a block number.
	Block param.Field[EvmPostTransactionScanParamsBlockUnion] `json:"block"`
	// list of one or both of options for the desired output. "simulation" - include
	// simulation output in your response. "validation" - include security validation
	// of the transaction in your response. Default is ["validation"]
	Options param.Field[[]EvmPostTransactionScanParamsOption] `json:"options"`
	// Override the state of the chain. This is useful for testing purposes.
	StateOverride param.Field[map[string]EvmPostTransactionScanParamsStateOverride] `json:"state_override"`
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

// The relative block for the block validation. Can be "latest" or a block number.
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

type EvmPostTransactionScanParamsStateOverride struct {
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

func (r EvmPostTransactionScanParamsStateOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
