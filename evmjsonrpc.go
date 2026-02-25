// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"net/http"
	"reflect"
	"slices"
	"time"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/param"
	"github.com/blockaid-official/blockaid-client-go/internal/requestconfig"
	"github.com/blockaid-official/blockaid-client-go/option"
	"github.com/blockaid-official/blockaid-client-go/shared"
	"github.com/tidwall/gjson"
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

// Get a risk recommendation with plain-language reasons for a JSON-RPC request.
func (r *EvmJsonRpcService) Scan(ctx context.Context, body EvmJsonRpcScanParams, opts ...option.RequestOption) (res *EvmJsonRpcScanResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/evm/json-rpc/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EvmJsonRpcScanResponse struct {
	Block                      string                                           `json:"block" api:"required"`
	Chain                      string                                           `json:"chain" api:"required"`
	AccountAddress             string                                           `json:"account_address"`
	Events                     []EvmJsonRpcScanResponseEvent                    `json:"events"`
	Features                   interface{}                                      `json:"features"`
	GasEstimation              EvmJsonRpcScanResponseGasEstimation              `json:"gas_estimation"`
	Simulation                 EvmJsonRpcScanResponseSimulation                 `json:"simulation"`
	UserOperationGasEstimation EvmJsonRpcScanResponseUserOperationGasEstimation `json:"user_operation_gas_estimation"`
	Validation                 EvmJsonRpcScanResponseValidation                 `json:"validation"`
	JSON                       evmJsonRpcScanResponseJSON                       `json:"-"`
}

// evmJsonRpcScanResponseJSON contains the JSON metadata for the struct
// [EvmJsonRpcScanResponse]
type evmJsonRpcScanResponseJSON struct {
	Block                      apijson.Field
	Chain                      apijson.Field
	AccountAddress             apijson.Field
	Events                     apijson.Field
	Features                   apijson.Field
	GasEstimation              apijson.Field
	Simulation                 apijson.Field
	UserOperationGasEstimation apijson.Field
	Validation                 apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseEvent struct {
	Data           string                              `json:"data" api:"required"`
	EmitterAddress string                              `json:"emitter_address" api:"required"`
	Topics         []string                            `json:"topics" api:"required"`
	EmitterName    string                              `json:"emitter_name"`
	Name           string                              `json:"name"`
	Params         []EvmJsonRpcScanResponseEventsParam `json:"params"`
	JSON           evmJsonRpcScanResponseEventJSON     `json:"-"`
}

// evmJsonRpcScanResponseEventJSON contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseEvent]
type evmJsonRpcScanResponseEventJSON struct {
	Data           apijson.Field
	EmitterAddress apijson.Field
	Topics         apijson.Field
	EmitterName    apijson.Field
	Name           apijson.Field
	Params         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseEventJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseEventsParam struct {
	Type         string                                       `json:"type" api:"required"`
	Value        EvmJsonRpcScanResponseEventsParamsValueUnion `json:"value" api:"required"`
	InternalType string                                       `json:"internalType"`
	Name         string                                       `json:"name"`
	JSON         evmJsonRpcScanResponseEventsParamJSON        `json:"-"`
}

// evmJsonRpcScanResponseEventsParamJSON contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseEventsParam]
type evmJsonRpcScanResponseEventsParamJSON struct {
	Type         apijson.Field
	Value        apijson.Field
	InternalType apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseEventsParam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseEventsParamJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or
// [EvmJsonRpcScanResponseEventsParamsValueArray].
type EvmJsonRpcScanResponseEventsParamsValueUnion interface {
	ImplementsEvmJsonRpcScanResponseEventsParamsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseEventsParamsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseEventsParamsValueArray{}),
		},
	)
}

type EvmJsonRpcScanResponseEventsParamsValueArray []interface{}

func (r EvmJsonRpcScanResponseEventsParamsValueArray) ImplementsEvmJsonRpcScanResponseEventsParamsValueUnion() {
}

type EvmJsonRpcScanResponseGasEstimation struct {
	Status   EvmJsonRpcScanResponseGasEstimationStatus `json:"status" api:"required"`
	Error    string                                    `json:"error"`
	Estimate string                                    `json:"estimate"`
	Used     string                                    `json:"used"`
	JSON     evmJsonRpcScanResponseGasEstimationJSON   `json:"-"`
	union    EvmJsonRpcScanResponseGasEstimationUnion
}

// evmJsonRpcScanResponseGasEstimationJSON contains the JSON metadata for the
// struct [EvmJsonRpcScanResponseGasEstimation]
type evmJsonRpcScanResponseGasEstimationJSON struct {
	Status      apijson.Field
	Error       apijson.Field
	Estimate    apijson.Field
	Used        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseGasEstimationJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseGasEstimation) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseGasEstimation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EvmJsonRpcScanResponseGasEstimationUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation],
// [EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError].
func (r EvmJsonRpcScanResponseGasEstimation) AsUnion() EvmJsonRpcScanResponseGasEstimationUnion {
	return r.union
}

// Union satisfied by
// [EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation]
// or
// [EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError].
type EvmJsonRpcScanResponseGasEstimationUnion interface {
	implementsEvmJsonRpcScanResponseGasEstimation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseGasEstimationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError{}),
		},
	)
}

type EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation struct {
	Estimate string                                                                                `json:"estimate" api:"required"`
	Status   EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatus `json:"status" api:"required"`
	Used     string                                                                                `json:"used" api:"required"`
	JSON     evmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationJSON   `json:"-"`
}

// evmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation]
type evmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationJSON struct {
	Estimate    apijson.Field
	Status      apijson.Field
	Used        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation) implementsEvmJsonRpcScanResponseGasEstimation() {
}

type EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatus string

const (
	EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatusSuccess EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatus = "Success"
)

func (r EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatus) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatusSuccess:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError struct {
	Error  string                                                                                     `json:"error" api:"required"`
	Status EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus `json:"status" api:"required"`
	JSON   evmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON   `json:"-"`
}

// evmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError]
type evmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError) implementsEvmJsonRpcScanResponseGasEstimation() {
}

type EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus string

const (
	EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatusError EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus = "Error"
)

func (r EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatusError:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseGasEstimationStatus string

const (
	EvmJsonRpcScanResponseGasEstimationStatusSuccess EvmJsonRpcScanResponseGasEstimationStatus = "Success"
	EvmJsonRpcScanResponseGasEstimationStatusError   EvmJsonRpcScanResponseGasEstimationStatus = "Error"
)

func (r EvmJsonRpcScanResponseGasEstimationStatus) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseGasEstimationStatusSuccess, EvmJsonRpcScanResponseGasEstimationStatusError:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulation struct {
	// A string indicating if the simulation was successful or not.
	Status EvmJsonRpcScanResponseSimulationStatus `json:"status" api:"required"`
	// This field can have the runtime type of
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummary].
	AccountSummary interface{} `json:"account_summary"`
	// This field can have the runtime type of
	// [map[string]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetail].
	AddressDetails interface{} `json:"address_details"`
	// This field can have the runtime type of
	// [map[string][]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff].
	AssetsDiffs interface{} `json:"assets_diffs"`
	// This field can have the runtime type of
	// [map[string][]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement].
	ContractManagement interface{} `json:"contract_management"`
	// A string explaining why the transaction failed
	Description string `json:"description"`
	// An error message if the simulation failed.
	Error string `json:"error"`
	// This field can have the runtime type of
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails].
	ErrorDetails interface{} `json:"error_details"`
	// This field can have the runtime type of
	// [map[string][]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure].
	Exposures interface{} `json:"exposures"`
	// This field can have the runtime type of
	// [[]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalance].
	MissingBalances interface{} `json:"missing_balances"`
	// This field can have the runtime type of
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationParams].
	Params interface{} `json:"params"`
	// This field can have the runtime type of
	// [map[string][]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKey].
	SessionKey interface{} `json:"session_key"`
	// The number of times the simulation ran until success
	SimulationRunCount int64 `json:"simulation_run_count"`
	// This field can have the runtime type of
	// [map[string]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiff].
	TotalUsdDiff interface{} `json:"total_usd_diff"`
	// This field can have the runtime type of [map[string]map[string]string].
	TotalUsdExposure interface{} `json:"total_usd_exposure"`
	// This field can have the runtime type of
	// [[]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions].
	TransactionActions interface{}                          `json:"transaction_actions"`
	JSON               evmJsonRpcScanResponseSimulationJSON `json:"-"`
	union              EvmJsonRpcScanResponseSimulationUnion
}

// evmJsonRpcScanResponseSimulationJSON contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulation]
type evmJsonRpcScanResponseSimulationJSON struct {
	Status             apijson.Field
	AccountSummary     apijson.Field
	AddressDetails     apijson.Field
	AssetsDiffs        apijson.Field
	ContractManagement apijson.Field
	Description        apijson.Field
	Error              apijson.Field
	ErrorDetails       apijson.Field
	Exposures          apijson.Field
	MissingBalances    apijson.Field
	Params             apijson.Field
	SessionKey         apijson.Field
	SimulationRunCount apijson.Field
	TotalUsdDiff       apijson.Field
	TotalUsdExposure   apijson.Field
	TransactionActions apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulation) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EvmJsonRpcScanResponseSimulationUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulation],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationError].
func (r EvmJsonRpcScanResponseSimulation) AsUnion() EvmJsonRpcScanResponseSimulationUnion {
	return r.union
}

// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulation] or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationError].
type EvmJsonRpcScanResponseSimulationUnion interface {
	implementsEvmJsonRpcScanResponseSimulation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationError{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulation struct {
	// Account summary for the account address. account address is determined implicit
	// by the `from` field in the transaction request, or explicit by the
	// account_address field in the request.
	AccountSummary EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummary `json:"account_summary" api:"required"`
	// a dictionary including additional information about each relevant address in the
	// transaction.
	AddressDetails map[string]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetail `json:"address_details" api:"required"`
	// dictionary describes the assets differences as a result of this transaction for
	// every involved address
	AssetsDiffs map[string][]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff `json:"assets_diffs" api:"required"`
	// dictionary describes the exposure differences as a result of this transaction
	// for every involved address (as a result of any approval / setApproval / permit
	// function)
	Exposures map[string][]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure `json:"exposures" api:"required"`
	// Session keys created in this transaction per address
	SessionKey map[string][]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKey `json:"session_key" api:"required"`
	// A string indicating if the simulation was successful or not.
	Status EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationStatus `json:"status" api:"required"`
	// dictionary represents the usd value each address gained / lost during this
	// transaction
	TotalUsdDiff map[string]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiff `json:"total_usd_diff" api:"required"`
	// a dictionary representing the usd value each address is exposed to, split by
	// spenders
	TotalUsdExposure map[string]map[string]string `json:"total_usd_exposure" api:"required"`
	// Describes the nature of the transaction and what happened as part of it
	TransactionActions []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions `json:"transaction_actions" api:"required"`
	// Describes the state differences as a result of this transaction for every
	// involved address
	ContractManagement map[string][]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement `json:"contract_management"`
	// Missing balances in the transaction
	MissingBalances []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalance `json:"missing_balances"`
	// The parameters of the transaction that was simulated.
	Params EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationParams `json:"params"`
	// The number of times the simulation ran until success
	SimulationRunCount int64                                                                       `json:"simulation_run_count"`
	JSON               evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulation]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationJSON struct {
	AccountSummary     apijson.Field
	AddressDetails     apijson.Field
	AssetsDiffs        apijson.Field
	Exposures          apijson.Field
	SessionKey         apijson.Field
	Status             apijson.Field
	TotalUsdDiff       apijson.Field
	TotalUsdExposure   apijson.Field
	TransactionActions apijson.Field
	ContractManagement apijson.Field
	MissingBalances    apijson.Field
	Params             apijson.Field
	SimulationRunCount apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulation) implementsEvmJsonRpcScanResponseSimulation() {
}

// Account summary for the account address. account address is determined implicit
// by the `from` field in the transaction request, or explicit by the
// account_address field in the request.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummary struct {
	// All assets diffs related to the account address
	AssetsDiffs []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff `json:"assets_diffs" api:"required"`
	// All assets exposures related to the account address
	Exposures []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure `json:"exposures" api:"required"`
	// Total usd diff related to the account address
	TotalUsdDiff EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiff `json:"total_usd_diff" api:"required"`
	// Total usd exposure related to the account address
	TotalUsdExposure map[string]string `json:"total_usd_exposure" api:"required"`
	// All assets traces related to the account address
	Traces []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace `json:"traces" api:"required"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryJSON    `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummary]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryJSON struct {
	AssetsDiffs      apijson.Field
	Exposures        apijson.Field
	TotalUsdDiff     apijson.Field
	TotalUsdExposure apijson.Field
	Traces           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff struct {
	// This field can have the runtime type of
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAsset].
	Asset interface{} `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType `json:"asset_type" api:"required"`
	// This field can have the runtime type of
	// [[]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffIn],
	// [[]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffIn],
	// [[]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffIn],
	// [[]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffIn].
	In interface{} `json:"in" api:"required"`
	// This field can have the runtime type of
	// [[]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOut],
	// [[]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOut],
	// [[]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOut],
	// [[]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOut].
	Out interface{} `json:"out" api:"required"`
	// This field can have the runtime type of
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChanges],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChanges],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChanges],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChanges].
	BalanceChanges interface{}                                                                                         `json:"balance_changes"`
	JSON           evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffJSON `json:"-"`
	union          EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff struct {
	// description of the asset for the current diff
	Asset EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType `json:"asset_type" api:"required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffIn `json:"in" api:"required"`
	// amount of the asset that was transferred from the address in this transaction
	Out []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOut `json:"out" api:"required"`
	// shows the balance before making the transaction and after
	BalanceChanges EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChanges `json:"balance_changes"`
	JSON           evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffJSON           `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff() {
}

// description of the asset for the current diff
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType `json:"type" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                        `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetJSON `json:"-"`
	union  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                               `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetTypeErc20  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType = "ERC20"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetTypeErc20, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffIn struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                     `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffInJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffInJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffIn]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffInJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOut struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                      `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOutJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOutJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOut]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOutJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOutJSON) RawJSON() string {
	return r.raw
}

// shows the balance before making the transaction and after
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChanges struct {
	// balance of the account after making the transaction
	After EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfter `json:"after" api:"required"`
	// balance of the account before making the transaction
	Before EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBefore `json:"before" api:"required"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesJSON   `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChanges]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChanges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesJSON) RawJSON() string {
	return r.raw
}

// balance of the account after making the transaction
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfter struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                      `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfterJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfterJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfter]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfterJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfterJSON) RawJSON() string {
	return r.raw
}

// balance of the account before making the transaction
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBefore struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                       `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBefore]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff struct {
	// description of the asset for the current diff
	Asset EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType `json:"asset_type" api:"required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffIn `json:"in" api:"required"`
	// amount of the asset that was transferred from the address in this transaction
	Out []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOut `json:"out" api:"required"`
	// shows the balance before making the transaction and after
	BalanceChanges EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChanges `json:"balance_changes"`
	JSON           evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffJSON           `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff() {
}

// description of the asset for the current diff
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                         `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetJSON `json:"-"`
	union  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                 `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                 `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetTypeErc721 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType = "ERC721"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetTypeErc721, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffIn struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                      `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffInJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffInJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffIn]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffInJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOut struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                       `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOutJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOutJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOut]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOutJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOutJSON) RawJSON() string {
	return r.raw
}

// shows the balance before making the transaction and after
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChanges struct {
	// balance of the account after making the transaction
	After EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfter `json:"after" api:"required"`
	// balance of the account before making the transaction
	Before EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBefore `json:"before" api:"required"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesJSON   `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChanges]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChanges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesJSON) RawJSON() string {
	return r.raw
}

// balance of the account after making the transaction
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfter struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                       `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfterJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfterJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfter]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfterJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfterJSON) RawJSON() string {
	return r.raw
}

// balance of the account before making the transaction
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBefore struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                        `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBefore]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff struct {
	// description of the asset for the current diff
	Asset EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType `json:"asset_type" api:"required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffIn `json:"in" api:"required"`
	// amount of the asset that was transferred from the address in this transaction
	Out []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOut `json:"out" api:"required"`
	// shows the balance before making the transaction and after
	BalanceChanges EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChanges `json:"balance_changes"`
	JSON           evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffJSON           `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff() {
}

// description of the asset for the current diff
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                          `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetJSON `json:"-"`
	union  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                   `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                  `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetTypeErc1155 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType = "ERC1155"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetTypeNonerc  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetTypeErc1155, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffIn struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                       `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffInJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffInJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffIn]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffInJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOut struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                        `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOutJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOutJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOut]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOutJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOutJSON) RawJSON() string {
	return r.raw
}

// shows the balance before making the transaction and after
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChanges struct {
	// balance of the account after making the transaction
	After EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfter `json:"after" api:"required"`
	// balance of the account before making the transaction
	Before EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBefore `json:"before" api:"required"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesJSON   `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChanges]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChanges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesJSON) RawJSON() string {
	return r.raw
}

// balance of the account after making the transaction
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfter struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                        `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfterJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfterJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfter]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfterJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfterJSON) RawJSON() string {
	return r.raw
}

// balance of the account before making the transaction
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBefore struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                         `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBefore]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff struct {
	// description of the asset for the current diff
	Asset EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetType `json:"asset_type" api:"required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffIn `json:"in" api:"required"`
	// amount of the asset that was transferred from the address in this transaction
	Out []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOut `json:"out" api:"required"`
	// shows the balance before making the transaction and after
	BalanceChanges EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChanges `json:"balance_changes"`
	JSON           evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffJSON           `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff() {
}

// description of the asset for the current diff
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAsset struct {
	ChainID   int64  `json:"chain_id" api:"required"`
	ChainName string `json:"chain_name" api:"required"`
	Decimals  int64  `json:"decimals" api:"required"`
	LogoURL   string `json:"logo_url" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetType `json:"type" api:"required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                         `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAsset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetJSON struct {
	ChainID     apijson.Field
	ChainName   apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetTypeNative EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetType = "NATIVE"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetTypeNative:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffIn struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                      `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffInJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffInJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffIn]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffInJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOut struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                       `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOutJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOutJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOut]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOutJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOutJSON) RawJSON() string {
	return r.raw
}

// shows the balance before making the transaction and after
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChanges struct {
	// balance of the account after making the transaction
	After EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfter `json:"after" api:"required"`
	// balance of the account before making the transaction
	Before EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBefore `json:"before" api:"required"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesJSON   `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChanges]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChanges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesJSON) RawJSON() string {
	return r.raw
}

// balance of the account after making the transaction
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfter struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                       `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfterJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfterJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfter]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfterJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfterJSON) RawJSON() string {
	return r.raw
}

// balance of the account before making the transaction
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBefore struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                        `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBeforeJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBeforeJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBefore]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBeforeJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBeforeJSON) RawJSON() string {
	return r.raw
}

// type of the asset for the current diff
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc20   EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType = "ERC20"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc721  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType = "ERC721"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc1155 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType = "ERC1155"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeNative  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType = "NATIVE"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc20, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc721, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc1155, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeNative:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure struct {
	// This field can have the runtime type of
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset].
	Asset interface{} `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType `json:"asset_type" api:"required"`
	// This field can have the runtime type of
	// [map[string]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpender],
	// [map[string]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpender],
	// [map[string]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpender].
	Spenders interface{}                                                                                       `json:"spenders" api:"required"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposureJSON `json:"-"`
	union    EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposureJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresUnion {
	return r.union
}

// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType `json:"asset_type" api:"required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpender `json:"spenders" api:"required"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureJSON               `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure() {
}

// description of the asset for the current diff
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType `json:"type" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                        `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetJSON `json:"-"`
	union  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                               `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetTypeErc20  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType = "ERC20"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetTypeErc20, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpender struct {
	// the amount that was asked in the approval request for this spender from the
	// current address and asset
	Approval string                                                                                                                                                 `json:"approval" api:"required"`
	Exposure []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure `json:"exposure" api:"required"`
	// the usd price of the approval amount
	ApprovalUsdPrice string `json:"approval_usd_price"`
	// the expiration time of the permit2 protocol
	Expiration time.Time `json:"expiration" format:"date-time"`
	// user friendly description of the approval
	Summary string                                                                                                                                          `json:"summary"`
	JSON    evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpender]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON struct {
	Approval         apijson.Field
	Exposure         apijson.Field
	ApprovalUsdPrice apijson.Field
	Expiration       apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// id of the token
	TokenID string `json:"token_id"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value before divided by decimal, that was transferred from this address
	Value string                                                                                                                                                   `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON `json:"-"`
	union EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON struct {
	ArbitraryCollectionToken apijson.Field
	LogoURL                  apijson.Field
	RawValue                 apijson.Field
	Summary                  apijson.Field
	TokenID                  apijson.Field
	UsdPrice                 apijson.Field
	Value                    apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                                `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                               `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                              `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                               `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType `json:"asset_type" api:"required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpender `json:"spenders" api:"required"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureJSON               `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure() {
}

// description of the asset for the current diff
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                         `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetJSON `json:"-"`
	union  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                 `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                 `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetTypeErc721 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType = "ERC721"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetTypeErc721, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpender struct {
	Exposure []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure `json:"exposure" api:"required"`
	// boolean indicates whether an is_approved_for_all function was used (missing in
	// case of ERC20 / ERC1155)
	IsApprovedForAll bool `json:"is_approved_for_all" api:"required"`
	// user friendly description of the approval
	Summary string                                                                                                                                           `json:"summary"`
	JSON    evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpender]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// id of the token
	TokenID string `json:"token_id"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value before divided by decimal, that was transferred from this address
	Value string                                                                                                                                                    `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON `json:"-"`
	union EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON struct {
	ArbitraryCollectionToken apijson.Field
	LogoURL                  apijson.Field
	RawValue                 apijson.Field
	Summary                  apijson.Field
	TokenID                  apijson.Field
	UsdPrice                 apijson.Field
	Value                    apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                                 `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                                `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                               `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType `json:"asset_type" api:"required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpender `json:"spenders" api:"required"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureJSON               `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure() {
}

// description of the asset for the current diff
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                          `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON `json:"-"`
	union  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                   `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                  `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeErc1155 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType = "ERC1155"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeNonerc  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeErc1155, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpender struct {
	Exposure []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure `json:"exposure" api:"required"`
	// boolean indicates whether an is_approved_for_all function was used (missing in
	// case of ERC20 / ERC1155)
	IsApprovedForAll bool `json:"is_approved_for_all" api:"required"`
	// user friendly description of the approval
	Summary string                                                                                                                                            `json:"summary"`
	JSON    evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpender]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// id of the token
	TokenID string `json:"token_id"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value before divided by decimal, that was transferred from this address
	Value string                                                                                                                                                     `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON `json:"-"`
	union EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON struct {
	ArbitraryCollectionToken apijson.Field
	LogoURL                  apijson.Field
	RawValue                 apijson.Field
	Summary                  apijson.Field
	TokenID                  apijson.Field
	UsdPrice                 apijson.Field
	Value                    apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                                  `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                                 `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                 `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

// type of the asset for the current diff
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc20   EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType = "ERC20"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc721  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType = "ERC721"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc1155 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType = "ERC1155"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc20, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc721, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc1155:
		return true
	}
	return false
}

// Total usd diff related to the account address
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiff struct {
	In    string                                                                                                `json:"in" api:"required"`
	Out   string                                                                                                `json:"out" api:"required"`
	Total string                                                                                                `json:"total" api:"required"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace struct {
	// This field can have the runtime type of
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAsset],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset].
	Asset interface{} `json:"asset" api:"required"`
	// type of the trace
	TraceType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceType `json:"trace_type" api:"required"`
	// The type of the model
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType `json:"type" api:"required"`
	// This field can have the runtime type of
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiff],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiff],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiff],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiff].
	Diff interface{} `json:"diff"`
	// This field can have the runtime type of
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposed],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposed].
	Exposed interface{} `json:"exposed"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address"`
	// This field can have the runtime type of
	// [[]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabels],
	// [[]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabels],
	// [[]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabels],
	// [[]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabels].
	Labels interface{} `json:"labels"`
	// The owner of the assets
	Owner string `json:"owner"`
	// The spender of the assets
	Spender string `json:"spender"`
	// The address where the assets are moved to
	ToAddress string                                                                                         `json:"to_address"`
	JSON      evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTraceJSON `json:"-"`
	union     EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTraceJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTraceJSON struct {
	Asset       apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	Diff        apijson.Field
	Exposed     apijson.Field
	FromAddress apijson.Field
	Labels      apijson.Field
	Owner       apijson.Field
	Spender     apijson.Field
	ToAddress   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTraceJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesUnion {
	return r.union
}

// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace struct {
	// Description of the asset in the trace
	Asset EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset `json:"asset" api:"required"`
	// The difference in value for the asset in the trace
	Diff EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiff `json:"diff" api:"required"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address" api:"required"`
	// The address where the assets are moved to
	ToAddress string `json:"to_address" api:"required"`
	// type of the trace
	TraceType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceType `json:"trace_type" api:"required"`
	// The type of the model
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceType `json:"type" api:"required"`
	// List of labels that describe the trace
	Labels []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabels `json:"labels"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceJSON     `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceJSON struct {
	Asset       apijson.Field
	Diff        apijson.Field
	FromAddress apijson.Field
	ToAddress   apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	Labels      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetType `json:"type" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetJSON `json:"-"`
	union  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                       `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                        `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetTypeErc20  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetType = "ERC20"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetTypeErc20, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetTypeNonerc:
		return true
	}
	return false
}

// The difference in value for the asset in the trace
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                               `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiffJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceTypeAssetTrace EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceType = "AssetTrace"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceTypeAssetTrace:
		return true
	}
	return false
}

// The type of the model
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTypeErc20AssetTrace EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceType = "ERC20AssetTrace"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTypeErc20AssetTrace:
		return true
	}
	return false
}

// Label assigned to a trace entry.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabels string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabelsGasFee EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabels = "GAS_FEE"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabels) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabelsGasFee:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace struct {
	// Description of the asset in the trace
	Asset EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset `json:"asset" api:"required"`
	// The difference in value for the asset in the trace
	Diff EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiff `json:"diff" api:"required"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address" api:"required"`
	// The address where the assets are moved to
	ToAddress string `json:"to_address" api:"required"`
	// type of the trace
	TraceType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceType `json:"trace_type" api:"required"`
	// The type of the model
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceType `json:"type" api:"required"`
	// List of labels that describe the trace
	Labels []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabels `json:"labels"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceJSON     `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceJSON struct {
	Asset       apijson.Field
	Diff        apijson.Field
	FromAddress apijson.Field
	ToAddress   apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	Labels      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                 `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetJSON `json:"-"`
	union  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                         `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                         `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetTypeErc721 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetType = "ERC721"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetTypeErc721, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetTypeNonerc:
		return true
	}
	return false
}

// The difference in value for the asset in the trace
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiffJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceTypeAssetTrace EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceType = "AssetTrace"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceTypeAssetTrace:
		return true
	}
	return false
}

// The type of the model
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTypeErc721AssetTrace EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceType = "ERC721AssetTrace"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTypeErc721AssetTrace:
		return true
	}
	return false
}

// Label assigned to a trace entry.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabels string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabelsGasFee EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabels = "GAS_FEE"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabels) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabelsGasFee:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace struct {
	// Description of the asset in the trace
	Asset EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset `json:"asset" api:"required"`
	// The difference in value for the asset in the trace
	Diff EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiff `json:"diff" api:"required"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address" api:"required"`
	// The address where the assets are moved to
	ToAddress string `json:"to_address" api:"required"`
	// type of the trace
	TraceType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceType `json:"trace_type" api:"required"`
	// The type of the model
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceType `json:"type" api:"required"`
	// List of labels that describe the trace
	Labels []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabels `json:"labels"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceJSON     `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceJSON struct {
	Asset       apijson.Field
	Diff        apijson.Field
	FromAddress apijson.Field
	ToAddress   apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	Labels      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                  `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetJSON `json:"-"`
	union  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                           `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                          `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetTypeErc1155 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetType = "ERC1155"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetTypeNonerc  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetTypeErc1155, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetTypeNonerc:
		return true
	}
	return false
}

// The difference in value for the asset in the trace
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                 `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiffJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceTypeAssetTrace EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceType = "AssetTrace"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceTypeAssetTrace:
		return true
	}
	return false
}

// The type of the model
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTypeErc1155AssetTrace EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceType = "ERC1155AssetTrace"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTypeErc1155AssetTrace:
		return true
	}
	return false
}

// Label assigned to a trace entry.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabels string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabelsGasFee EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabels = "GAS_FEE"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabels) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabelsGasFee:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace struct {
	// Description of the asset in the trace
	Asset EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAsset `json:"asset" api:"required"`
	// The difference in value for the asset in the trace
	Diff EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiff `json:"diff" api:"required"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address" api:"required"`
	// The address where the assets are moved to
	ToAddress string `json:"to_address" api:"required"`
	// type of the trace
	TraceType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceType `json:"trace_type" api:"required"`
	// The type of the model
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceType `json:"type" api:"required"`
	// List of labels that describe the trace
	Labels []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabels `json:"labels"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceJSON     `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceJSON struct {
	Asset       apijson.Field
	Diff        apijson.Field
	FromAddress apijson.Field
	ToAddress   apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	Labels      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAsset struct {
	ChainID   int64  `json:"chain_id" api:"required"`
	ChainName string `json:"chain_name" api:"required"`
	Decimals  int64  `json:"decimals" api:"required"`
	LogoURL   string `json:"logo_url" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetType `json:"type" api:"required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                 `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAsset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetJSON struct {
	ChainID     apijson.Field
	ChainName   apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetTypeNative EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetType = "NATIVE"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetTypeNative:
		return true
	}
	return false
}

// The difference in value for the asset in the trace
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiffJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceTypeAssetTrace EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceType = "AssetTrace"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceTypeAssetTrace:
		return true
	}
	return false
}

// The type of the model
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTypeNativeAssetTrace EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceType = "NativeAssetTrace"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTypeNativeAssetTrace:
		return true
	}
	return false
}

// Label assigned to a trace entry.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabels string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabelsGasFee EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabels = "GAS_FEE"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabels) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabelsGasFee:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace struct {
	// Description of the asset in the trace
	Asset   EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset   `json:"asset" api:"required"`
	Exposed EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposed `json:"exposed" api:"required"`
	// The owner of the assets
	Owner string `json:"owner" api:"required"`
	// The spender of the assets
	Spender string `json:"spender" api:"required"`
	// type of the trace
	TraceType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceType `json:"trace_type" api:"required"`
	// The type of the model
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceType `json:"type" api:"required"`
	JSON evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceJSON struct {
	Asset       apijson.Field
	Exposed     apijson.Field
	Owner       apijson.Field
	Spender     apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetType `json:"type" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                   `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetJSON `json:"-"`
	union  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                          `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                           `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetTypeErc20  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetType = "ERC20"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetTypeErc20, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetTypeNonerc:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposed struct {
	RawValue string                                                                                                                                     `json:"raw_value" api:"required"`
	UsdPrice float64                                                                                                                                    `json:"usd_price"`
	Value    float64                                                                                                                                    `json:"value"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposedJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposedJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposed]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposedJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposedJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceTypeExposureTrace EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceType = "ExposureTrace"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceTypeExposureTrace:
		return true
	}
	return false
}

// The type of the model
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTypeErc20ExposureTrace EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceType = "ERC20ExposureTrace"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTypeErc20ExposureTrace:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace struct {
	// Description of the asset in the trace
	Asset EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset `json:"asset" api:"required"`
	// The owner of the assets
	Owner string `json:"owner" api:"required"`
	// The spender of the assets
	Spender string `json:"spender" api:"required"`
	// type of the trace
	TraceType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceType `json:"trace_type" api:"required"`
	// The type of the model
	Type    EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceType    `json:"type" api:"required"`
	Exposed EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposed `json:"exposed"`
	JSON    evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceJSON    `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceJSON struct {
	Asset       apijson.Field
	Owner       apijson.Field
	Spender     apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	Exposed     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                    `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetJSON `json:"-"`
	union  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                            `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                            `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetTypeErc721 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetType = "ERC721"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetTypeErc721, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetTypeNonerc:
		return true
	}
	return false
}

// type of the trace
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceTypeExposureTrace EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceType = "ExposureTrace"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceTypeExposureTrace:
		return true
	}
	return false
}

// The type of the model
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTypeErc721ExposureTrace EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceType = "ERC721ExposureTrace"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTypeErc721ExposureTrace:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposed struct {
	Amount   int64                                                                                                                                       `json:"amount" api:"required"`
	TokenID  string                                                                                                                                      `json:"token_id" api:"required"`
	IsMint   bool                                                                                                                                        `json:"is_mint"`
	LogoURL  string                                                                                                                                      `json:"logo_url"`
	UsdPrice float64                                                                                                                                     `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposedJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposedJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposed]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposedJSON struct {
	Amount      apijson.Field
	TokenID     apijson.Field
	IsMint      apijson.Field
	LogoURL     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposedJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace struct {
	// Description of the asset in the trace
	Asset EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset `json:"asset" api:"required"`
	// The owner of the assets
	Owner string `json:"owner" api:"required"`
	// The spender of the assets
	Spender string `json:"spender" api:"required"`
	// type of the trace
	TraceType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceType `json:"trace_type" api:"required"`
	// The type of the model
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceType `json:"type" api:"required"`
	JSON evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceJSON struct {
	Asset       apijson.Field
	Owner       apijson.Field
	Spender     apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                     `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetJSON `json:"-"`
	union  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                              `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                             `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetTypeErc1155 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetType = "ERC1155"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetTypeNonerc  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetTypeErc1155, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetTypeNonerc:
		return true
	}
	return false
}

// type of the trace
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceTypeExposureTrace EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceType = "ExposureTrace"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceTypeExposureTrace:
		return true
	}
	return false
}

// The type of the model
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTypeErc1155ExposureTrace EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceType = "ERC1155ExposureTrace"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTypeErc1155ExposureTrace:
		return true
	}
	return false
}

// type of the trace
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceTypeAssetTrace    EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceType = "AssetTrace"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceTypeExposureTrace EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceType = "ExposureTrace"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceTypeAssetTrace, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceTypeExposureTrace:
		return true
	}
	return false
}

// The type of the model
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc20AssetTrace      EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC20AssetTrace"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc721AssetTrace     EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC721AssetTrace"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc1155AssetTrace    EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC1155AssetTrace"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeNativeAssetTrace     EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "NativeAssetTrace"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc20ExposureTrace   EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC20ExposureTrace"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc721ExposureTrace  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC721ExposureTrace"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc1155ExposureTrace EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC1155ExposureTrace"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc20AssetTrace, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc721AssetTrace, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc1155AssetTrace, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeNativeAssetTrace, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc20ExposureTrace, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc721ExposureTrace, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc1155ExposureTrace:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetail struct {
	// contains the contract's name if the address is a verified contract
	ContractName string `json:"contract_name"`
	// Whether the address is an eoa or a contract
	IsEoa bool `json:"is_eoa"`
	// known name tag for the address
	NameTag string                                                                                   `json:"name_tag"`
	JSON    evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetailJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetailJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetail]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetailJSON struct {
	ContractName apijson.Field
	IsEoa        apijson.Field
	NameTag      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetailJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff struct {
	// This field can have the runtime type of
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAsset].
	Asset interface{} `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType `json:"asset_type" api:"required"`
	// This field can have the runtime type of
	// [[]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffIn],
	// [[]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffIn],
	// [[]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffIn],
	// [[]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffIn].
	In interface{} `json:"in" api:"required"`
	// This field can have the runtime type of
	// [[]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOut],
	// [[]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOut],
	// [[]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOut],
	// [[]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOut].
	Out   interface{}                                                                           `json:"out" api:"required"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffJSON `json:"-"`
	union EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff struct {
	// description of the asset for the current diff
	Asset EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType `json:"asset_type" api:"required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffIn `json:"in" api:"required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOut `json:"out" api:"required"`
	JSON evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffJSON  `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff() {
}

// description of the asset for the current diff
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType `json:"type" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                             `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetJSON `json:"-"`
	union  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                    `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                     `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetTypeErc20  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType = "ERC20"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetTypeErc20, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffIn struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                          `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffInJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffInJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffIn]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffInJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOut struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                           `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOutJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOutJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOut]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff struct {
	// description of the asset for the current diff
	Asset EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType `json:"asset_type" api:"required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffIn `json:"in" api:"required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOut `json:"out" api:"required"`
	JSON evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffJSON  `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff() {
}

// description of the asset for the current diff
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                              `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetJSON `json:"-"`
	union  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                      `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                      `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetTypeErc721 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType = "ERC721"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetTypeErc721, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffIn struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                           `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffInJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffInJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffIn]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffInJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOut struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                            `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOutJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOutJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOut]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOutJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff struct {
	// description of the asset for the current diff
	Asset EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType `json:"asset_type" api:"required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffIn `json:"in" api:"required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOut `json:"out" api:"required"`
	JSON evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffJSON  `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff() {
}

// description of the asset for the current diff
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                               `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetJSON `json:"-"`
	union  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                        `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                       `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetTypeErc1155 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType = "ERC1155"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetTypeNonerc  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetTypeErc1155, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffIn struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                            `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffInJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffInJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffIn]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffInJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOut struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                             `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOutJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOutJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOut]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOutJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff struct {
	// description of the asset for the current diff
	Asset EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetType `json:"asset_type" api:"required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffIn `json:"in" api:"required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOut `json:"out" api:"required"`
	JSON evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffJSON  `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff() {
}

// description of the asset for the current diff
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAsset struct {
	ChainID   int64  `json:"chain_id" api:"required"`
	ChainName string `json:"chain_name" api:"required"`
	Decimals  int64  `json:"decimals" api:"required"`
	LogoURL   string `json:"logo_url" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetType `json:"type" api:"required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                              `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAsset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetJSON struct {
	ChainID     apijson.Field
	ChainName   apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetTypeNative EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetType = "NATIVE"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetTypeNative:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffIn struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                           `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffInJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffInJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffIn]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffInJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOut struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                            `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOutJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOutJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOut]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

// type of the asset for the current diff
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc20   EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType = "ERC20"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc721  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType = "ERC721"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc1155 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType = "ERC1155"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeNative  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType = "NATIVE"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc20, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc721, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc1155, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeNative:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure struct {
	// This field can have the runtime type of
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset].
	Asset interface{} `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType `json:"asset_type" api:"required"`
	// This field can have the runtime type of
	// [map[string]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpender],
	// [map[string]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpender],
	// [map[string]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpender].
	Spenders interface{}                                                                         `json:"spenders" api:"required"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposureJSON `json:"-"`
	union    EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposureJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresUnion {
	return r.union
}

// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType `json:"asset_type" api:"required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpender `json:"spenders" api:"required"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureJSON               `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure() {
}

// description of the asset for the current diff
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType `json:"type" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                          `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetJSON `json:"-"`
	union  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                 `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                  `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetTypeErc20  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType = "ERC20"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetTypeErc20, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpender struct {
	// the amount that was asked in the approval request for this spender from the
	// current address and asset
	Approval string                                                                                                                                   `json:"approval" api:"required"`
	Exposure []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure `json:"exposure" api:"required"`
	// the usd price of the approval amount
	ApprovalUsdPrice string `json:"approval_usd_price"`
	// the expiration time of the permit2 protocol
	Expiration time.Time `json:"expiration" format:"date-time"`
	// user friendly description of the approval
	Summary string                                                                                                                            `json:"summary"`
	JSON    evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpender]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON struct {
	Approval         apijson.Field
	Exposure         apijson.Field
	ApprovalUsdPrice apijson.Field
	Expiration       apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// id of the token
	TokenID string `json:"token_id"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value before divided by decimal, that was transferred from this address
	Value string                                                                                                                                     `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON `json:"-"`
	union EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON struct {
	ArbitraryCollectionToken apijson.Field
	LogoURL                  apijson.Field
	RawValue                 apijson.Field
	Summary                  apijson.Field
	TokenID                  apijson.Field
	UsdPrice                 apijson.Field
	Value                    apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                  `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                 `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                 `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType `json:"asset_type" api:"required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpender `json:"spenders" api:"required"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureJSON               `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure() {
}

// description of the asset for the current diff
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                           `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetJSON `json:"-"`
	union  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                   `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                   `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetTypeErc721 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType = "ERC721"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetTypeErc721, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpender struct {
	Exposure []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure `json:"exposure" api:"required"`
	// boolean indicates whether an is_approved_for_all function was used (missing in
	// case of ERC20 / ERC1155)
	IsApprovedForAll bool `json:"is_approved_for_all" api:"required"`
	// user friendly description of the approval
	Summary string                                                                                                                             `json:"summary"`
	JSON    evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpender]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// id of the token
	TokenID string `json:"token_id"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value before divided by decimal, that was transferred from this address
	Value string                                                                                                                                      `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON `json:"-"`
	union EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON struct {
	ArbitraryCollectionToken apijson.Field
	LogoURL                  apijson.Field
	RawValue                 apijson.Field
	Summary                  apijson.Field
	TokenID                  apijson.Field
	UsdPrice                 apijson.Field
	Value                    apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                   `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                  `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                 `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                  `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType `json:"asset_type" api:"required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpender `json:"spenders" api:"required"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureJSON               `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure() {
}

// description of the asset for the current diff
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                            `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON `json:"-"`
	union  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                     `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                    `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeErc1155 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType = "ERC1155"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeNonerc  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType = "NONERC"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeErc1155, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpender struct {
	Exposure []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure `json:"exposure" api:"required"`
	// boolean indicates whether an is_approved_for_all function was used (missing in
	// case of ERC20 / ERC1155)
	IsApprovedForAll bool `json:"is_approved_for_all" api:"required"`
	// user friendly description of the approval
	Summary string                                                                                                                              `json:"summary"`
	JSON    evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpender]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// id of the token
	TokenID string `json:"token_id"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value before divided by decimal, that was transferred from this address
	Value string                                                                                                                                       `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON `json:"-"`
	union EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON struct {
	ArbitraryCollectionToken apijson.Field
	LogoURL                  apijson.Field
	RawValue                 apijson.Field
	Summary                  apijson.Field
	TokenID                  apijson.Field
	UsdPrice                 apijson.Field
	Value                    apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                    `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token" api:"required"`
	// id of the token
	TokenID string `json:"token_id" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                   `json:"usd_price"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                  `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                   `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

// type of the asset for the current diff
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc20   EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType = "ERC20"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc721  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType = "ERC721"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc1155 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType = "ERC1155"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc20, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc721, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc1155:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKey struct {
	Key      string                                                                                    `json:"key" api:"required"`
	Policies []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy `json:"policies" api:"required"`
	Signer   string                                                                                    `json:"signer" api:"required"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyJSON     `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKey]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyJSON struct {
	Key         apijson.Field
	Policies    apijson.Field
	Signer      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy struct {
	// This field can have the runtime type of
	// [[]EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArg].
	Args interface{} `json:"args"`
	// This field can have the runtime type of
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails].
	AssetDetails interface{}                                                                                   `json:"asset_details"`
	EndTime      time.Time                                                                                     `json:"end_time" format:"date-time"`
	Method       string                                                                                        `json:"method"`
	Recipient    string                                                                                        `json:"recipient"`
	StartTime    time.Time                                                                                     `json:"start_time" format:"date-time"`
	ToAddress    string                                                                                        `json:"to_address"`
	Type         EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType `json:"type"`
	ValueLimit   string                                                                                        `json:"value_limit"`
	JSON         evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicyJSON   `json:"-"`
	union        EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicyJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicyJSON struct {
	Args         apijson.Field
	AssetDetails apijson.Field
	EndTime      apijson.Field
	Method       apijson.Field
	Recipient    apijson.Field
	StartTime    apijson.Field
	ToAddress    apijson.Field
	Type         apijson.Field
	ValueLimit   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicyJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesUnion {
	return r.union
}

// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy struct {
	AssetDetails EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails `json:"asset_details" api:"required"`
	Recipient    string                                                                                                                                   `json:"recipient"`
	Type         EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyType         `json:"type"`
	ValueLimit   string                                                                                                                                   `json:"value_limit"`
	JSON         evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyJSON         `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyJSON struct {
	AssetDetails apijson.Field
	Recipient    apijson.Field
	Type         apijson.Field
	ValueLimit   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails struct {
	Decimals int64 `json:"decimals" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsType `json:"type" api:"required"`
	// address of the token
	Address   string `json:"address"`
	ChainID   int64  `json:"chain_id"`
	ChainName string `json:"chain_name"`
	LogoURL   string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                       `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsJSON `json:"-"`
	union  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	Address     apijson.Field
	ChainID     apijson.Field
	ChainName   apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsUnion {
	return r.union
}

// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails struct {
	ChainID   int64  `json:"chain_id" api:"required"`
	ChainName string `json:"chain_name" api:"required"`
	Decimals  int64  `json:"decimals" api:"required"`
	LogoURL   string `json:"logo_url" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsType `json:"type" api:"required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                               `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsJSON struct {
	ChainID     apijson.Field
	ChainName   apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsTypeNative EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsType = "NATIVE"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsTypeNative:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                              `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsTypeNative EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsType = "NATIVE"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsTypeErc20  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsType = "ERC20"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsTypeNative, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsTypeErc20:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyTypeTransferPolicy EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyType = "TRANSFER_POLICY"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyTypeTransferPolicy:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy struct {
	ToAddress  string                                                                                                                        `json:"to_address" api:"required"`
	Args       []EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArg `json:"args"`
	Method     string                                                                                                                        `json:"method"`
	Type       EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyType  `json:"type"`
	ValueLimit string                                                                                                                        `json:"value_limit"`
	JSON       evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyJSON  `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyJSON struct {
	ToAddress   apijson.Field
	Args        apijson.Field
	Method      apijson.Field
	Type        apijson.Field
	ValueLimit  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArg struct {
	// Comparison operator used to evaluate an argument/value against a policy
	// constraint.
	Condition EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition `json:"condition" api:"required"`
	Index     int64                                                                                                                                 `json:"index" api:"required"`
	Value     string                                                                                                                                `json:"value" api:"required"`
	JSON      evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgJSON       `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArg]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgJSON struct {
	Condition   apijson.Field
	Index       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgJSON) RawJSON() string {
	return r.raw
}

// Comparison operator used to evaluate an argument/value against a policy
// constraint.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionUnconstrained  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "UNCONSTRAINED"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionEqual          EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "EQUAL"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionGreater        EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "GREATER"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionLess           EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "LESS"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionGreaterOrEqual EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "GREATER_OR_EQUAL"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionLessOrEqual    EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "LESS_OR_EQUAL"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionNotEqual       EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "NOT_EQUAL"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionUnconstrained, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionEqual, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionGreater, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionLess, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionGreaterOrEqual, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionLessOrEqual, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionNotEqual:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyTypeCallPolicy EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyType = "CALL_POLICY"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyTypeCallPolicy:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy struct {
	ValueLimit string                                                                                                                      `json:"value_limit" api:"required"`
	Type       EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyType `json:"type"`
	JSON       evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyJSON struct {
	ValueLimit  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyTypeGasPolicy EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyType = "GAS_POLICY"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyTypeGasPolicy:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy struct {
	EndTime   time.Time                                                                                                                          `json:"end_time" format:"date-time"`
	StartTime time.Time                                                                                                                          `json:"start_time" format:"date-time"`
	Type      EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyType `json:"type"`
	JSON      evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyTypeExpirationPolicy EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyType = "EXPIRATION_POLICY"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyTypeExpirationPolicy:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeTransferPolicy   EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType = "TRANSFER_POLICY"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeCallPolicy       EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType = "CALL_POLICY"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeGasPolicy        EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType = "GAS_POLICY"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeExpirationPolicy EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType = "EXPIRATION_POLICY"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeTransferPolicy, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeCallPolicy, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeGasPolicy, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeExpirationPolicy:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationStatus string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationStatusSuccess EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationStatus = "Success"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationStatus) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationStatusSuccess:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiff struct {
	In    string                                                                                  `json:"in" api:"required"`
	Out   string                                                                                  `json:"out" api:"required"`
	Total string                                                                                  `json:"total" api:"required"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiffJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiffJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiff]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

// Type of action performed in the transaction.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsMint            EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "mint"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsStake           EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "stake"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsSwap            EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "swap"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsNativeTransfer  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "native_transfer"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsTokenTransfer   EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "token_transfer"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsApproval        EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "approval"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsSetCodeAccount  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "set_code_account"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsProxyUpgrade    EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "proxy_upgrade"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsOwnershipChange EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "ownership_change"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsMint, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsStake, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsSwap, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsNativeTransfer, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsTokenTransfer, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsApproval, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsSetCodeAccount, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsProxyUpgrade, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsOwnershipChange:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement struct {
	// The type of the state change
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType `json:"type" api:"required"`
	// This field can have the runtime type of
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfter],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfter],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfter].
	After interface{} `json:"after"`
	// This field can have the runtime type of
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBefore],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBefore],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBefore].
	Before interface{} `json:"before"`
	// The delegated address
	DelegatedAddress string `json:"delegated_address"`
	// The direct creator address of the new contract
	Deployer string                                                                                        `json:"deployer"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementJSON `json:"-"`
	union    EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementJSON struct {
	Type             apijson.Field
	After            apijson.Field
	Before           apijson.Field
	DelegatedAddress apijson.Field
	Deployer         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementUnion {
	return r.union
}

// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement struct {
	// The state after the transaction
	After EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfter `json:"after" api:"required"`
	// The state before the transaction
	Before EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBefore `json:"before" api:"required"`
	// The type of the state change
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementType `json:"type" api:"required"`
	JSON evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementJSON struct {
	After       apijson.Field
	Before      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement() {
}

// The state after the transaction
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfter struct {
	Address string                                                                                                                                     `json:"address" api:"required"`
	JSON    evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfterJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfterJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfter]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfterJSON struct {
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfterJSON) RawJSON() string {
	return r.raw
}

// The state before the transaction
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBefore struct {
	Address string                                                                                                                                      `json:"address" api:"required"`
	JSON    evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBeforeJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBeforeJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBefore]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBeforeJSON struct {
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBeforeJSON) RawJSON() string {
	return r.raw
}

// The type of the state change
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementTypeProxyUpgrade EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementType = "PROXY_UPGRADE"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementTypeProxyUpgrade:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement struct {
	// The state after the transaction
	After EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfter `json:"after" api:"required"`
	// The state before the transaction
	Before EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBefore `json:"before" api:"required"`
	// The type of the state change
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementType `json:"type" api:"required"`
	JSON evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementJSON struct {
	After       apijson.Field
	Before      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement() {
}

// The state after the transaction
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfter struct {
	Owners []string                                                                                                                                      `json:"owners" api:"required"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfterJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfterJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfter]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfterJSON struct {
	Owners      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfterJSON) RawJSON() string {
	return r.raw
}

// The state before the transaction
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBefore struct {
	Owners []string                                                                                                                                       `json:"owners" api:"required"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBeforeJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBeforeJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBefore]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBeforeJSON struct {
	Owners      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBeforeJSON) RawJSON() string {
	return r.raw
}

// The type of the state change
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementTypeOwnershipChange EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementType = "OWNERSHIP_CHANGE"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementTypeOwnershipChange:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement struct {
	// The state after the transaction
	After EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfter `json:"after" api:"required"`
	// The state before the transaction
	Before EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBefore `json:"before" api:"required"`
	// The type of the state change
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementType `json:"type" api:"required"`
	JSON evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementJSON struct {
	After       apijson.Field
	Before      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement() {
}

// The state after the transaction
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfter struct {
	Modules []string                                                                                                                                    `json:"modules" api:"required"`
	JSON    evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfterJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfterJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfter]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfterJSON struct {
	Modules     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfterJSON) RawJSON() string {
	return r.raw
}

// The state before the transaction
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBefore struct {
	Modules []string                                                                                                                                     `json:"modules" api:"required"`
	JSON    evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBeforeJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBeforeJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBefore]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBeforeJSON struct {
	Modules     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBeforeJSON) RawJSON() string {
	return r.raw
}

// The type of the state change
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementTypeModuleChange EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementType = "MODULE_CHANGE"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementTypeModuleChange:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement struct {
	// The delegated address
	DelegatedAddress string `json:"delegated_address" api:"required"`
	// The type of the state change
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementType `json:"type" api:"required"`
	JSON evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementJSON struct {
	DelegatedAddress apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement() {
}

// The type of the state change
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementTypeSetCodeAccount EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementType = "SET_CODE_ACCOUNT"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementTypeSetCodeAccount:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation struct {
	// The direct creator address of the new contract
	Deployer string `json:"deployer" api:"required"`
	// The type of the state change
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationType `json:"type" api:"required"`
	JSON evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationJSON struct {
	Deployer    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement() {
}

// The type of the state change
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationTypeContractCreation EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationType = "CONTRACT_CREATION"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationTypeContractCreation:
		return true
	}
	return false
}

// The type of the state change
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeProxyUpgrade     EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType = "PROXY_UPGRADE"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeOwnershipChange  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType = "OWNERSHIP_CHANGE"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeModuleChange     EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType = "MODULE_CHANGE"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeSetCodeAccount   EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType = "SET_CODE_ACCOUNT"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeContractCreation EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType = "CONTRACT_CREATION"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeProxyUpgrade, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeOwnershipChange, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeModuleChange, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeSetCodeAccount, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeContractCreation:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalance struct {
	// The asset that is missing balance
	Asset EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset `json:"asset" api:"required"`
	// The account address's current balance of the asset
	CurrentBalance EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalance `json:"current_balance" api:"required"`
	// The account address's missing balance of the asset
	MissingBalance EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalance `json:"missing_balance" api:"required"`
	// The required balance of the asset for this action
	RequiredBalance EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalance `json:"required_balance" api:"required"`
	JSON            evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalanceJSON             `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalanceJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalance]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalanceJSON struct {
	Asset           apijson.Field
	CurrentBalance  apijson.Field
	MissingBalance  apijson.Field
	RequiredBalance apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalanceJSON) RawJSON() string {
	return r.raw
}

// The asset that is missing balance
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset struct {
	// asset's decimals
	Decimals int64 `json:"decimals" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetType `json:"type" api:"required"`
	// address of the token
	Address   string `json:"address"`
	ChainID   int64  `json:"chain_id"`
	ChainName string `json:"chain_name"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                          `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetJSON `json:"-"`
	union  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	Address     apijson.Field
	ChainID     apijson.Field
	ChainName   apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetUnion {
	return r.union
}

// The asset that is missing balance
//
// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                 `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails struct {
	ChainID   int64  `json:"chain_id" api:"required"`
	ChainName string `json:"chain_name" api:"required"`
	Decimals  int64  `json:"decimals" api:"required"`
	LogoURL   string `json:"logo_url" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsType `json:"type" api:"required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                  `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsJSON struct {
	ChainID     apijson.Field
	ChainName   apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset() {
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsTypeNative EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsType = "NATIVE"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsTypeNative:
		return true
	}
	return false
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetTypeErc20  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetType = "ERC20"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetTypeNative EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetType = "NATIVE"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetTypeErc20, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetTypeNative:
		return true
	}
	return false
}

// The account address's current balance of the asset
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalance struct {
	// The raw value of the balance in hex string format
	RawValue string `json:"raw_value" api:"required"`
	// The value of the balance in decimal string format
	Value string                                                                                                   `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalanceJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalanceJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalance]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalanceJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalanceJSON) RawJSON() string {
	return r.raw
}

// The account address's missing balance of the asset
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalance struct {
	// The raw value of the balance in hex string format
	RawValue string `json:"raw_value" api:"required"`
	// The value of the balance in decimal string format
	Value string                                                                                                   `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalanceJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalanceJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalance]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalanceJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalanceJSON) RawJSON() string {
	return r.raw
}

// The required balance of the asset for this action
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalance struct {
	// The raw value of the balance in hex string format
	RawValue string `json:"raw_value" api:"required"`
	// The value of the balance in decimal string format
	Value string                                                                                                    `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalanceJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalanceJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalance]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalanceJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalanceJSON) RawJSON() string {
	return r.raw
}

// The parameters of the transaction that was simulated.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationParams struct {
	// The block tag to be sent.
	BlockTag string `json:"block_tag"`
	// The calldata to be sent.
	Calldata EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldata `json:"calldata"`
	// The chain to be sent.
	Chain string `json:"chain"`
	// The data to be sent.
	Data string `json:"data"`
	// The address the transaction is sent from.
	From string `json:"from"`
	// The gas to be sent.
	Gas string `json:"gas"`
	// The gas price to be sent.
	GasPrice string `json:"gas_price"`
	// The address the transaction is directed to.
	To string `json:"to"`
	// The user operation call data to be sent.
	UserOperationCalldata EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldata `json:"user_operation_calldata"`
	// The value to be sent.
	Value string                                                                            `json:"value"`
	JSON  evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationParams]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsJSON struct {
	BlockTag              apijson.Field
	Calldata              apijson.Field
	Chain                 apijson.Field
	Data                  apijson.Field
	From                  apijson.Field
	Gas                   apijson.Field
	GasPrice              apijson.Field
	To                    apijson.Field
	UserOperationCalldata apijson.Field
	Value                 apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsJSON) RawJSON() string {
	return r.raw
}

// The calldata to be sent.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldata struct {
	// The function selector of the function called in the transaction
	FunctionSelector string `json:"function_selector" api:"required"`
	// The function declaration of the function called in the transaction
	FunctionDeclaration string `json:"function_declaration"`
	// The function signature of the function called in the transaction
	FunctionSignature string                                                                                    `json:"function_signature"`
	JSON              evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldataJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldataJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldata]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldataJSON struct {
	FunctionSelector    apijson.Field
	FunctionDeclaration apijson.Field
	FunctionSignature   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldataJSON) RawJSON() string {
	return r.raw
}

// The user operation call data to be sent.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldata struct {
	// The function selector of the function called in the transaction
	FunctionSelector string `json:"function_selector" api:"required"`
	// The function declaration of the function called in the transaction
	FunctionDeclaration string `json:"function_declaration"`
	// The function signature of the function called in the transaction
	FunctionSignature string                                                                                                 `json:"function_signature"`
	JSON              evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldataJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldataJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldata]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldataJSON struct {
	FunctionSelector    apijson.Field
	FunctionDeclaration apijson.Field
	FunctionSignature   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldataJSON) RawJSON() string {
	return r.raw
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationError struct {
	// A string explaining why the transaction failed
	Description string `json:"description" api:"required"`
	// An error message if the simulation failed.
	Error string `json:"error" api:"required"`
	// A string indicating if the simulation was successful or not.
	Status EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatus `json:"status" api:"required"`
	// Error details if the simulation failed.
	ErrorDetails EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails `json:"error_details"`
	JSON         evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorJSON         `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationError]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorJSON struct {
	Description  apijson.Field
	Error        apijson.Field
	Status       apijson.Field
	ErrorDetails apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationError) implementsEvmJsonRpcScanResponseSimulation() {
}

// A string indicating if the simulation was successful or not.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatus string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatusError EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatus = "Error"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatus) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatusError:
		return true
	}
	return false
}

// Error details if the simulation failed.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails struct {
	// The type of the model
	Code string `json:"code" api:"required"`
	// The address of the account
	AccountAddress string `json:"account_address"`
	// The address that is invalid
	Address string `json:"address"`
	// This field can have the runtime type of
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset].
	Asset    interface{} `json:"asset"`
	Category string      `json:"category"`
	// The current balance of the account
	CurrentBalance string `json:"current_balance"`
	// The domain name that is unsupported
	DomainName string `json:"domain_name"`
	// The message type that is unsupported
	MessageType string `json:"message_type"`
	// The required balance of the account
	RequiredBalance string                                                                                       `json:"required_balance"`
	JSON            evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsJSON `json:"-"`
	union           EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsJSON struct {
	Code            apijson.Field
	AccountAddress  apijson.Field
	Address         apijson.Field
	Asset           apijson.Field
	Category        apijson.Field
	CurrentBalance  apijson.Field
	DomainName      apijson.Field
	MessageType     apijson.Field
	RequiredBalance apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsUnion {
	return r.union
}

// Error details if the simulation failed.
//
// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails struct {
	// The address of the account
	AccountAddress string `json:"account_address" api:"required"`
	// The asset that the account does not have enough balance for
	Asset EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset `json:"asset" api:"required"`
	// The type of the model
	Code     EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCode     `json:"code" api:"required"`
	Category EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCategory `json:"category"`
	// The current balance of the account
	CurrentBalance string `json:"current_balance"`
	// The required balance of the account
	RequiredBalance string                                                                                                                                             `json:"required_balance"`
	JSON            evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsJSON struct {
	AccountAddress  apijson.Field
	Asset           apijson.Field
	Code            apijson.Field
	Category        apijson.Field
	CurrentBalance  apijson.Field
	RequiredBalance apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails() {
}

// The asset that the account does not have enough balance for
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset struct {
	// This field can have the runtime type of
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetails],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetails],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetails],
	// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetails].
	Details interface{} `json:"details" api:"required"`
	// The type of the model
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType `json:"type" api:"required"`
	// Token Id
	TokenID int64                                                                                                                                                   `json:"token_id"`
	JSON    evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetJSON `json:"-"`
	union   EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetUnion
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetJSON struct {
	Details     apijson.Field
	Type        apijson.Field
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset].
func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset) AsUnion() EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetUnion {
	return r.union
}

// The asset that the account does not have enough balance for
//
// Union satisfied by
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset],
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset]
// or
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset].
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetUnion interface {
	implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset{}),
		},
	)
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset struct {
	// Details
	Details EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetails `json:"details" api:"required"`
	// The type of the model
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetType `json:"type" api:"required"`
	JSON evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetJSON struct {
	Details     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset() {
}

// Details
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetails struct {
	ChainID   int64  `json:"chain_id" api:"required"`
	ChainName string `json:"chain_name" api:"required"`
	Decimals  int64  `json:"decimals" api:"required"`
	LogoURL   string `json:"logo_url" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsType `json:"type" api:"required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                      `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsJSON struct {
	ChainID     apijson.Field
	ChainName   apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsTypeNative EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsType = "NATIVE"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsTypeNative:
		return true
	}
	return false
}

// The type of the model
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetTypeNative EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetType = "NATIVE"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetTypeNative:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset struct {
	// Details
	Details EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetails `json:"details" api:"required"`
	// The type of the model
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetType `json:"type" api:"required"`
	JSON evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetJSON struct {
	Details     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset() {
}

// Details
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                     `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsTypeErc20 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsType = "ERC20"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsTypeErc20:
		return true
	}
	return false
}

// The type of the model
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetTypeErc20 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetType = "ERC20"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetTypeErc20:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset struct {
	// Details
	Details EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetails `json:"details" api:"required"`
	// Token Id
	TokenID int64 `json:"token_id" api:"required"`
	// The type of the model
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetType `json:"type" api:"required"`
	JSON evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetJSON struct {
	Details     apijson.Field
	TokenID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset() {
}

// Details
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                      `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsTypeErc1155 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsType = "ERC1155"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsTypeErc1155:
		return true
	}
	return false
}

// The type of the model
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetTypeErc721 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetType = "ERC721"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetTypeErc721:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset struct {
	// Details
	Details EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetails `json:"details" api:"required"`
	// Token Id
	TokenID int64 `json:"token_id" api:"required"`
	// The type of the model
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetType `json:"type" api:"required"`
	JSON evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetJSON struct {
	Details     apijson.Field
	TokenID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset() {
}

// Details
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                       `json:"symbol"`
	JSON   evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsTypeErc1155 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsType = "ERC1155"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsTypeErc1155:
		return true
	}
	return false
}

// The type of the model
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetTypeErc1155 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetType = "ERC1155"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetTypeErc1155:
		return true
	}
	return false
}

// The type of the model
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeNative  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType = "NATIVE"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc20   EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType = "ERC20"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc721  EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType = "ERC721"
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc1155 EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType = "ERC1155"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeNative, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc20, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc721, EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc1155:
		return true
	}
	return false
}

// The type of the model
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCode string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCodeGeneralInsufficientFunds EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCode = "GENERAL_INSUFFICIENT_FUNDS"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCode) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCodeGeneralInsufficientFunds:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCategory string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCategoryRevert EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCategory = "REVERT"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCategory) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCategoryRevert:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails struct {
	// The address that is invalid
	Address string `json:"address" api:"required"`
	// The type of the model
	Code     EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCode     `json:"code" api:"required"`
	Category EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCategory `json:"category"`
	JSON     evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsJSON     `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsJSON struct {
	Address     apijson.Field
	Code        apijson.Field
	Category    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails() {
}

// The type of the model
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCode string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCodeGeneralInvalidAddress EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCode = "GENERAL_INVALID_ADDRESS"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCode) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCodeGeneralInvalidAddress:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCategory string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCategoryInvalidInput EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCategory = "INVALID_INPUT"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCategory) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCategoryInvalidInput:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails struct {
	Category string `json:"category" api:"required"`
	// The error code
	Code string                                                                                                                            `json:"code" api:"required"`
	JSON evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetailsJSON `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetailsJSON struct {
	Category    apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails() {
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails struct {
	// The type of the model
	Code EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCode `json:"code" api:"required"`
	// The domain name that is unsupported
	DomainName string `json:"domain_name" api:"required"`
	// The message type that is unsupported
	MessageType string                                                                                                                                                 `json:"message_type" api:"required"`
	Category    EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCategory `json:"category"`
	JSON        evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsJSON     `json:"-"`
}

// evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails]
type evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsJSON struct {
	Code        apijson.Field
	DomainName  apijson.Field
	MessageType apijson.Field
	Category    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails) implementsEvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails() {
}

// The type of the model
type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCode string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCodeUnsupportedEip712Message EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCode = "UNSUPPORTED_EIP712_MESSAGE"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCode) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCodeUnsupportedEip712Message:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCategory string

const (
	EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCategoryInvalidInput EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCategory = "INVALID_INPUT"
)

func (r EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCategory) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCategoryInvalidInput:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type EvmJsonRpcScanResponseSimulationStatus string

const (
	EvmJsonRpcScanResponseSimulationStatusSuccess EvmJsonRpcScanResponseSimulationStatus = "Success"
	EvmJsonRpcScanResponseSimulationStatusError   EvmJsonRpcScanResponseSimulationStatus = "Error"
)

func (r EvmJsonRpcScanResponseSimulationStatus) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseSimulationStatusSuccess, EvmJsonRpcScanResponseSimulationStatusError:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseUserOperationGasEstimation struct {
	Status                           EvmJsonRpcScanResponseUserOperationGasEstimationStatus `json:"status" api:"required"`
	CallGasEstimate                  string                                                 `json:"call_gas_estimate"`
	Error                            string                                                 `json:"error"`
	PaymasterVerificationGasEstimate string                                                 `json:"paymaster_verification_gas_estimate"`
	PreVerificationGasEstimate       string                                                 `json:"pre_verification_gas_estimate"`
	VerificationGasEstimate          string                                                 `json:"verification_gas_estimate"`
	JSON                             evmJsonRpcScanResponseUserOperationGasEstimationJSON   `json:"-"`
	union                            EvmJsonRpcScanResponseUserOperationGasEstimationUnion
}

// evmJsonRpcScanResponseUserOperationGasEstimationJSON contains the JSON metadata
// for the struct [EvmJsonRpcScanResponseUserOperationGasEstimation]
type evmJsonRpcScanResponseUserOperationGasEstimationJSON struct {
	Status                           apijson.Field
	CallGasEstimate                  apijson.Field
	Error                            apijson.Field
	PaymasterVerificationGasEstimate apijson.Field
	PreVerificationGasEstimate       apijson.Field
	VerificationGasEstimate          apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r evmJsonRpcScanResponseUserOperationGasEstimationJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseUserOperationGasEstimation) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseUserOperationGasEstimation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EvmJsonRpcScanResponseUserOperationGasEstimationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [UserOperationV6GasEstimation],
// [UserOperationV7GasEstimation],
// [EvmJsonRpcScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError].
func (r EvmJsonRpcScanResponseUserOperationGasEstimation) AsUnion() EvmJsonRpcScanResponseUserOperationGasEstimationUnion {
	return r.union
}

// Union satisfied by [UserOperationV6GasEstimation],
// [UserOperationV7GasEstimation] or
// [EvmJsonRpcScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError].
type EvmJsonRpcScanResponseUserOperationGasEstimationUnion interface {
	implementsEvmJsonRpcScanResponseUserOperationGasEstimation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseUserOperationGasEstimationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserOperationV6GasEstimation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserOperationV7GasEstimation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError{}),
		},
	)
}

type EvmJsonRpcScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError struct {
	Error  string                                                                                                  `json:"error" api:"required"`
	Status EvmJsonRpcScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus `json:"status" api:"required"`
	JSON   evmJsonRpcScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON   `json:"-"`
}

// evmJsonRpcScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError]
type evmJsonRpcScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError) implementsEvmJsonRpcScanResponseUserOperationGasEstimation() {
}

type EvmJsonRpcScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus string

const (
	EvmJsonRpcScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatusError EvmJsonRpcScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus = "Error"
)

func (r EvmJsonRpcScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatusError:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseUserOperationGasEstimationStatus string

const (
	EvmJsonRpcScanResponseUserOperationGasEstimationStatusSuccess EvmJsonRpcScanResponseUserOperationGasEstimationStatus = "Success"
	EvmJsonRpcScanResponseUserOperationGasEstimationStatusError   EvmJsonRpcScanResponseUserOperationGasEstimationStatus = "Error"
)

func (r EvmJsonRpcScanResponseUserOperationGasEstimationStatus) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseUserOperationGasEstimationStatusSuccess, EvmJsonRpcScanResponseUserOperationGasEstimationStatusError:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseValidation struct {
	// This field can have the runtime type of
	// [[]EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationFeature],
	// [[]EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeature].
	Features interface{} `json:"features" api:"required"`
	// Result type returned when validation succeeds.
	ResultType EvmJsonRpcScanResponseValidationResultType `json:"result_type" api:"required"`
	// A string indicating if the simulation was successful or not.
	Status EvmJsonRpcScanResponseValidationStatus `json:"status" api:"required"`
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification"`
	// A textual description that can be presented to the user about what this
	// transaction is doing.
	Description string `json:"description"`
	// An error message if the validation failed.
	Error string `json:"error"`
	// A textual description about the reasons the transaction was flagged with
	// result_type.
	Reason string                               `json:"reason"`
	JSON   evmJsonRpcScanResponseValidationJSON `json:"-"`
	union  EvmJsonRpcScanResponseValidationUnion
}

// evmJsonRpcScanResponseValidationJSON contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseValidation]
type evmJsonRpcScanResponseValidationJSON struct {
	Features       apijson.Field
	ResultType     apijson.Field
	Status         apijson.Field
	Classification apijson.Field
	Description    apijson.Field
	Error          apijson.Field
	Reason         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r evmJsonRpcScanResponseValidationJSON) RawJSON() string {
	return r.raw
}

func (r *EvmJsonRpcScanResponseValidation) UnmarshalJSON(data []byte) (err error) {
	*r = EvmJsonRpcScanResponseValidation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EvmJsonRpcScanResponseValidationUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidation],
// [EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationError].
func (r EvmJsonRpcScanResponseValidation) AsUnion() EvmJsonRpcScanResponseValidationUnion {
	return r.union
}

// Union satisfied by
// [EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidation] or
// [EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationError].
type EvmJsonRpcScanResponseValidationUnion interface {
	implementsEvmJsonRpcScanResponseValidation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmJsonRpcScanResponseValidationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationError{}),
		},
	)
}

type EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidation struct {
	// A list of features about this transaction explaining the validation.
	Features []EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationFeature `json:"features" api:"required"`
	// Result type returned when validation succeeds.
	ResultType EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationResultType `json:"result_type" api:"required"`
	// A string indicating if the simulation was successful or not.
	Status EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationStatus `json:"status" api:"required"`
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification"`
	// A textual description that can be presented to the user about what this
	// transaction is doing.
	Description string `json:"description"`
	// A textual description about the reasons the transaction was flagged with
	// result_type.
	Reason string                                                                      `json:"reason"`
	JSON   evmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationJSON `json:"-"`
}

// evmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidation]
type evmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationJSON struct {
	Features       apijson.Field
	ResultType     apijson.Field
	Status         apijson.Field
	Classification apijson.Field
	Description    apijson.Field
	Reason         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidation) implementsEvmJsonRpcScanResponseValidation() {
}

type EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationFeature struct {
	// Textual description
	Description string `json:"description" api:"required"`
	// Feature name
	FeatureID string `json:"feature_id" api:"required"`
	// Security result of a transaction scan feature.
	Type EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType `json:"type" api:"required"`
	// Address the feature refers to
	Address string `json:"address"`
	// Metadata related to the feature
	Metadata interface{}                                                                        `json:"metadata"`
	JSON     evmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationFeatureJSON `json:"-"`
}

// evmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationFeatureJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationFeature]
type evmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationFeatureJSON struct {
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	Address     apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationFeatureJSON) RawJSON() string {
	return r.raw
}

// Security result of a transaction scan feature.
type EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType string

const (
	EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeMalicious EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType = "Malicious"
	EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeWarning   EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType = "Warning"
	EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeBenign    EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType = "Benign"
	EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeInfo      EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType = "Info"
)

func (r EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeMalicious, EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeWarning, EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeBenign, EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeInfo:
		return true
	}
	return false
}

// Result type returned when validation succeeds.
type EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationResultType string

const (
	EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeBenign    EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationResultType = "Benign"
	EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeWarning   EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationResultType = "Warning"
	EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeMalicious EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationResultType = "Malicious"
)

func (r EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationResultType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeBenign, EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeWarning, EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeMalicious:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationStatus string

const (
	EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationStatusSuccess EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationStatus = "Success"
)

func (r EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationStatus) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationStatusSuccess:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationError struct {
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassification `json:"classification" api:"required"`
	// A textual description that can be presented to the user about what this
	// transaction is doing.
	Description EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescription `json:"description" api:"required"`
	// An error message if the validation failed.
	Error string `json:"error" api:"required"`
	// A list of features about this transaction explaining the validation.
	Features []EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeature `json:"features" api:"required"`
	// A textual description about the reasons the transaction was flagged with
	// result_type.
	Reason EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorReason `json:"reason" api:"required"`
	// A string indicating if the transaction is safe to sign or not.
	ResultType EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultType `json:"result_type" api:"required"`
	// A string indicating if the simulation was successful or not.
	Status EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatus `json:"status" api:"required"`
	JSON   evmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorJSON   `json:"-"`
}

// evmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationError]
type evmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorJSON struct {
	Classification apijson.Field
	Description    apijson.Field
	Error          apijson.Field
	Features       apijson.Field
	Reason         apijson.Field
	ResultType     apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorJSON) RawJSON() string {
	return r.raw
}

func (r EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationError) implementsEvmJsonRpcScanResponseValidation() {
}

// A textual classification that can be presented to the user explaining the
// reason.
type EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassification string

const (
	EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassificationEmpty EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassification = ""
)

func (r EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassification) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassificationEmpty:
		return true
	}
	return false
}

// A textual description that can be presented to the user about what this
// transaction is doing.
type EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescription string

const (
	EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescriptionEmpty EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescription = ""
)

func (r EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescription) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescriptionEmpty:
		return true
	}
	return false
}

type EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeature struct {
	// Textual description
	Description string `json:"description" api:"required"`
	// Feature name
	FeatureID string `json:"feature_id" api:"required"`
	// Security result of a transaction scan feature.
	Type EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType `json:"type" api:"required"`
	// Address the feature refers to
	Address string `json:"address"`
	// Metadata related to the feature
	Metadata interface{}                                                                             `json:"metadata"`
	JSON     evmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeatureJSON `json:"-"`
}

// evmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeatureJSON
// contains the JSON metadata for the struct
// [EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeature]
type evmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeatureJSON struct {
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	Address     apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeatureJSON) RawJSON() string {
	return r.raw
}

// Security result of a transaction scan feature.
type EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType string

const (
	EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeMalicious EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType = "Malicious"
	EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeWarning   EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType = "Warning"
	EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeBenign    EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType = "Benign"
	EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeInfo      EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType = "Info"
)

func (r EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeMalicious, EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeWarning, EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeBenign, EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeInfo:
		return true
	}
	return false
}

// A textual description about the reasons the transaction was flagged with
// result_type.
type EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorReason string

const (
	EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorReasonEmpty EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorReason = ""
)

func (r EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorReason) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorReasonEmpty:
		return true
	}
	return false
}

// A string indicating if the transaction is safe to sign or not.
type EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultType string

const (
	EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultTypeError EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultType = "Error"
)

func (r EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultTypeError:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatus string

const (
	EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatusSuccess EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatus = "Success"
)

func (r EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatus) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatusSuccess:
		return true
	}
	return false
}

// Result type returned when validation succeeds.
type EvmJsonRpcScanResponseValidationResultType string

const (
	EvmJsonRpcScanResponseValidationResultTypeBenign    EvmJsonRpcScanResponseValidationResultType = "Benign"
	EvmJsonRpcScanResponseValidationResultTypeWarning   EvmJsonRpcScanResponseValidationResultType = "Warning"
	EvmJsonRpcScanResponseValidationResultTypeMalicious EvmJsonRpcScanResponseValidationResultType = "Malicious"
	EvmJsonRpcScanResponseValidationResultTypeError     EvmJsonRpcScanResponseValidationResultType = "Error"
)

func (r EvmJsonRpcScanResponseValidationResultType) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseValidationResultTypeBenign, EvmJsonRpcScanResponseValidationResultTypeWarning, EvmJsonRpcScanResponseValidationResultTypeMalicious, EvmJsonRpcScanResponseValidationResultTypeError:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type EvmJsonRpcScanResponseValidationStatus string

const (
	EvmJsonRpcScanResponseValidationStatusSuccess EvmJsonRpcScanResponseValidationStatus = "Success"
)

func (r EvmJsonRpcScanResponseValidationStatus) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanResponseValidationStatusSuccess:
		return true
	}
	return false
}

type EvmJsonRpcScanParams struct {
	// The chain name or chain ID
	Chain param.Field[TransactionScanSupportedChain] `json:"chain" api:"required"`
	// JSON-RPC request that was received by the wallet.
	Data param.Field[EvmJsonRpcScanParamsData] `json:"data" api:"required"`
	// Additional context for the scan (e.g., dapp URL/domain, integration source).
	// Used to enrich results and reduce false positives/negatives.
	Metadata param.Field[EvmJsonRpcScanParamsMetadataUnion] `json:"metadata" api:"required"`
	// The address of the account (wallet) that received the request, in hex string
	// format
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
	// Supported JSON-RPC methods that can be scanned.
	Method param.Field[EvmJsonRpcScanParamsDataMethod] `json:"method" api:"required"`
	// The parameters of the JSON-RPC request in JSON format
	Params param.Field[[]interface{}] `json:"params" api:"required"`
}

func (r EvmJsonRpcScanParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Supported JSON-RPC methods that can be scanned.
type EvmJsonRpcScanParamsDataMethod string

const (
	EvmJsonRpcScanParamsDataMethodEthSendTransaction    EvmJsonRpcScanParamsDataMethod = "eth_sendTransaction"
	EvmJsonRpcScanParamsDataMethodEthSendRawTransaction EvmJsonRpcScanParamsDataMethod = "eth_sendRawTransaction"
	EvmJsonRpcScanParamsDataMethodEthSignTransaction    EvmJsonRpcScanParamsDataMethod = "eth_signTransaction"
	EvmJsonRpcScanParamsDataMethodEthSignTypedData      EvmJsonRpcScanParamsDataMethod = "eth_signTypedData"
	EvmJsonRpcScanParamsDataMethodEthSignTypedDataV1    EvmJsonRpcScanParamsDataMethod = "eth_signTypedData_v1"
	EvmJsonRpcScanParamsDataMethodEthSignTypedDataV2    EvmJsonRpcScanParamsDataMethod = "eth_signTypedData_v2"
	EvmJsonRpcScanParamsDataMethodEthSignTypedDataV3    EvmJsonRpcScanParamsDataMethod = "eth_signTypedData_v3"
	EvmJsonRpcScanParamsDataMethodEthSignTypedDataV4    EvmJsonRpcScanParamsDataMethod = "eth_signTypedData_v4"
	EvmJsonRpcScanParamsDataMethodEthSendUserOperation  EvmJsonRpcScanParamsDataMethod = "eth_sendUserOperation"
	EvmJsonRpcScanParamsDataMethodPersonalSign          EvmJsonRpcScanParamsDataMethod = "personal_sign"
	EvmJsonRpcScanParamsDataMethodEthSign               EvmJsonRpcScanParamsDataMethod = "eth_sign"
	EvmJsonRpcScanParamsDataMethodWalletSendCalls       EvmJsonRpcScanParamsDataMethod = "wallet_sendCalls"
)

func (r EvmJsonRpcScanParamsDataMethod) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanParamsDataMethodEthSendTransaction, EvmJsonRpcScanParamsDataMethodEthSendRawTransaction, EvmJsonRpcScanParamsDataMethodEthSignTransaction, EvmJsonRpcScanParamsDataMethodEthSignTypedData, EvmJsonRpcScanParamsDataMethodEthSignTypedDataV1, EvmJsonRpcScanParamsDataMethodEthSignTypedDataV2, EvmJsonRpcScanParamsDataMethodEthSignTypedDataV3, EvmJsonRpcScanParamsDataMethodEthSignTypedDataV4, EvmJsonRpcScanParamsDataMethodEthSendUserOperation, EvmJsonRpcScanParamsDataMethodPersonalSign, EvmJsonRpcScanParamsDataMethodEthSign, EvmJsonRpcScanParamsDataMethodWalletSendCalls:
		return true
	}
	return false
}

// Additional context for the scan (e.g., dapp URL/domain, integration source).
// Used to enrich results and reduce false positives/negatives.
type EvmJsonRpcScanParamsMetadata struct {
	// The full URL of the DApp or website that initiated the transaction, for
	// cross-reference. Must use the https or http scheme and contain a valid hostname.
	// Cannot contain JSON, braces, or other embedded data structures.
	Domain param.Field[string] `json:"domain"`
	// Indicates that the transaction was not initiated by a dapp.
	NonDapp param.Field[bool] `json:"non_dapp"`
}

func (r EvmJsonRpcScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmJsonRpcScanParamsMetadata) implementsEvmJsonRpcScanParamsMetadataUnion() {}

// Additional context for the scan (e.g., dapp URL/domain, integration source).
// Used to enrich results and reduce false positives/negatives.
//
// Satisfied by [EvmJsonRpcScanParamsMetadataRoutersEvmModelsMetadataNonDapp],
// [EvmJsonRpcScanParamsMetadataRoutersEvmModelsMetadataDapp],
// [EvmJsonRpcScanParamsMetadata].
type EvmJsonRpcScanParamsMetadataUnion interface {
	implementsEvmJsonRpcScanParamsMetadataUnion()
}

type EvmJsonRpcScanParamsMetadataRoutersEvmModelsMetadataNonDapp struct {
	// Indicates that the transaction was not initiated by a dapp.
	NonDapp param.Field[EvmJsonRpcScanParamsMetadataRoutersEvmModelsMetadataNonDappNonDapp] `json:"non_dapp"`
}

func (r EvmJsonRpcScanParamsMetadataRoutersEvmModelsMetadataNonDapp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmJsonRpcScanParamsMetadataRoutersEvmModelsMetadataNonDapp) implementsEvmJsonRpcScanParamsMetadataUnion() {
}

// Indicates that the transaction was not initiated by a dapp.
type EvmJsonRpcScanParamsMetadataRoutersEvmModelsMetadataNonDappNonDapp bool

const (
	EvmJsonRpcScanParamsMetadataRoutersEvmModelsMetadataNonDappNonDappTrue EvmJsonRpcScanParamsMetadataRoutersEvmModelsMetadataNonDappNonDapp = true
)

func (r EvmJsonRpcScanParamsMetadataRoutersEvmModelsMetadataNonDappNonDapp) IsKnown() bool {
	switch r {
	case EvmJsonRpcScanParamsMetadataRoutersEvmModelsMetadataNonDappNonDappTrue:
		return true
	}
	return false
}

type EvmJsonRpcScanParamsMetadataRoutersEvmModelsMetadataDapp struct {
	// The full URL of the DApp or website that initiated the transaction, for
	// cross-reference. Must use the https or http scheme and contain a valid hostname.
	// Cannot contain JSON, braces, or other embedded data structures.
	Domain param.Field[string] `json:"domain" api:"required"`
	// Indicates that the transaction was not initiated by a dapp. Use false when the
	// transaction is from a dapp.
	NonDapp param.Field[bool] `json:"non_dapp"`
}

func (r EvmJsonRpcScanParamsMetadataRoutersEvmModelsMetadataDapp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmJsonRpcScanParamsMetadataRoutersEvmModelsMetadataDapp) implementsEvmJsonRpcScanParamsMetadataUnion() {
}

// The relative block for the block validation. Can be "latest" or a block number.
//
// Satisfied by [shared.UnionInt], [shared.UnionString].
type EvmJsonRpcScanParamsBlockUnion interface {
	ImplementsEvmJsonRpcScanParamsBlockUnion()
}

// Response sections to include (e.g., validation, simulation, gas estimation,
// events).
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
