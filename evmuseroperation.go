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

// EvmUserOperationService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvmUserOperationService] method instead.
type EvmUserOperationService struct {
	Options []option.RequestOption
}

// NewEvmUserOperationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEvmUserOperationService(opts ...option.RequestOption) (r *EvmUserOperationService) {
	r = &EvmUserOperationService{}
	r.Options = opts
	return
}

// Gets a user operation request and returns a full simulation indicating what will
// happen in the transaction together with a recommended action and some textual
// reasons of why the transaction was flagged that way.
func (r *EvmUserOperationService) Scan(ctx context.Context, body EvmUserOperationScanParams, opts ...option.RequestOption) (res *TransactionScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/evm/user-operation/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EvmUserOperationScanParams struct {
	// The chain name
	Chain param.Field[string] `json:"chain,required"`
	// The user operation request that was received by the wallet
	Data param.Field[EvmUserOperationScanParamsData] `json:"data,required"`
	// Object of additional information to validate against.
	Metadata param.Field[MetadataParam] `json:"metadata,required"`
	// The address of the account (wallet) sending the request in hex string format
	AccountAddress param.Field[string] `json:"account_address"`
	// list of one or both of options for the desired output. "simulation" - include
	// simulation output in your response. "validation" - include security validation
	// of the transaction in your response. Default is ["validation"]
	Options param.Field[[]EvmUserOperationScanParamsOption] `json:"options"`
}

func (r EvmUserOperationScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The user operation request that was received by the wallet
type EvmUserOperationScanParamsData struct {
	// The operation parameters of the user operation request
	Operation param.Field[EvmUserOperationScanParamsDataOperation] `json:"operation,required"`
	// The address of the entrypoint receiving the request in hex string format
	Entrypoint param.Field[string] `json:"entrypoint"`
}

func (r EvmUserOperationScanParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The operation parameters of the user operation request
type EvmUserOperationScanParamsDataOperation struct {
	// The call data value in hex string format.
	CallData param.Field[string] `json:"call_data"`
	// The call gas limit value in hex string format.
	CallGasLimit param.Field[string] `json:"call_gas_limit"`
	// The init code value in hex string format.
	InitCode param.Field[string] `json:"init_code"`
	// The max fee per gas value in hex string format.
	MaxFeePerGas param.Field[string] `json:"max_fee_per_gas"`
	// The max priority fee per gas value in hex string format.
	MaxPriorityFeePerGas param.Field[string] `json:"max_priority_fee_per_gas"`
	// The nonce value in hex string format.
	Nonce param.Field[string] `json:"nonce"`
	// The paymaster and data value in hex string format.
	PaymasterAndData param.Field[string] `json:"paymaster_and_data"`
	// The pre verification gas value in hex string format.
	PreVerificationGas param.Field[string] `json:"pre_verification_gas"`
	// The sender address of the operation in hex string format
	Sender param.Field[string] `json:"sender"`
	// The signature value in hex string format.
	Signature param.Field[string] `json:"signature"`
	// The verification gas limit value in hex string format.
	VerificationGasLimit param.Field[string] `json:"verification_gas_limit"`
}

func (r EvmUserOperationScanParamsDataOperation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An enumeration.
type EvmUserOperationScanParamsOption string

const (
	EvmUserOperationScanParamsOptionValidation    EvmUserOperationScanParamsOption = "validation"
	EvmUserOperationScanParamsOptionSimulation    EvmUserOperationScanParamsOption = "simulation"
	EvmUserOperationScanParamsOptionGasEstimation EvmUserOperationScanParamsOption = "gas_estimation"
	EvmUserOperationScanParamsOptionEvents        EvmUserOperationScanParamsOption = "events"
)

func (r EvmUserOperationScanParamsOption) IsKnown() bool {
	switch r {
	case EvmUserOperationScanParamsOptionValidation, EvmUserOperationScanParamsOptionSimulation, EvmUserOperationScanParamsOptionGasEstimation, EvmUserOperationScanParamsOptionEvents:
		return true
	}
	return false
}
