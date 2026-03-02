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

// Get a risk recommendation with plain-language reasons for a bulk of
// transactions.
func (r *EvmTransactionBulkService) Scan(ctx context.Context, body EvmTransactionBulkScanParams, opts ...option.RequestOption) (res *[]EvmTransactionBulkScanResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/evm/transaction-bulk/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EvmTransactionBulkScanResponse struct {
	Block                      string                                                   `json:"block" api:"required"`
	Chain                      string                                                   `json:"chain" api:"required"`
	AccountAddress             string                                                   `json:"account_address"`
	Events                     []EvmTransactionBulkScanResponseEvent                    `json:"events"`
	Features                   interface{}                                              `json:"features"`
	GasEstimation              EvmTransactionBulkScanResponseGasEstimation              `json:"gas_estimation"`
	Simulation                 EvmTransactionBulkScanResponseSimulation                 `json:"simulation"`
	UserOperationGasEstimation EvmTransactionBulkScanResponseUserOperationGasEstimation `json:"user_operation_gas_estimation"`
	Validation                 EvmTransactionBulkScanResponseValidation                 `json:"validation"`
	JSON                       evmTransactionBulkScanResponseJSON                       `json:"-"`
}

// evmTransactionBulkScanResponseJSON contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponse]
type evmTransactionBulkScanResponseJSON struct {
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

func (r *EvmTransactionBulkScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseEvent struct {
	Data           string                                      `json:"data" api:"required"`
	EmitterAddress string                                      `json:"emitter_address" api:"required"`
	Topics         []string                                    `json:"topics" api:"required"`
	EmitterName    string                                      `json:"emitter_name"`
	Name           string                                      `json:"name"`
	Params         []EvmTransactionBulkScanResponseEventsParam `json:"params"`
	JSON           evmTransactionBulkScanResponseEventJSON     `json:"-"`
}

// evmTransactionBulkScanResponseEventJSON contains the JSON metadata for the
// struct [EvmTransactionBulkScanResponseEvent]
type evmTransactionBulkScanResponseEventJSON struct {
	Data           apijson.Field
	EmitterAddress apijson.Field
	Topics         apijson.Field
	EmitterName    apijson.Field
	Name           apijson.Field
	Params         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseEventJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseEventsParam struct {
	Type         string                                               `json:"type" api:"required"`
	Value        EvmTransactionBulkScanResponseEventsParamsValueUnion `json:"value" api:"required"`
	InternalType string                                               `json:"internalType"`
	Name         string                                               `json:"name"`
	JSON         evmTransactionBulkScanResponseEventsParamJSON        `json:"-"`
}

// evmTransactionBulkScanResponseEventsParamJSON contains the JSON metadata for the
// struct [EvmTransactionBulkScanResponseEventsParam]
type evmTransactionBulkScanResponseEventsParamJSON struct {
	Type         apijson.Field
	Value        apijson.Field
	InternalType apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseEventsParam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseEventsParamJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or
// [EvmTransactionBulkScanResponseEventsParamsValueArray].
type EvmTransactionBulkScanResponseEventsParamsValueUnion interface {
	ImplementsEvmTransactionBulkScanResponseEventsParamsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseEventsParamsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseEventsParamsValueArray{}),
		},
	)
}

type EvmTransactionBulkScanResponseEventsParamsValueArray []interface{}

func (r EvmTransactionBulkScanResponseEventsParamsValueArray) ImplementsEvmTransactionBulkScanResponseEventsParamsValueUnion() {
}

type EvmTransactionBulkScanResponseGasEstimation struct {
	Status   EvmTransactionBulkScanResponseGasEstimationStatus `json:"status" api:"required"`
	Error    string                                            `json:"error"`
	Estimate string                                            `json:"estimate"`
	Used     string                                            `json:"used"`
	JSON     evmTransactionBulkScanResponseGasEstimationJSON   `json:"-"`
	union    EvmTransactionBulkScanResponseGasEstimationUnion
}

// evmTransactionBulkScanResponseGasEstimationJSON contains the JSON metadata for
// the struct [EvmTransactionBulkScanResponseGasEstimation]
type evmTransactionBulkScanResponseGasEstimationJSON struct {
	Status      apijson.Field
	Error       apijson.Field
	Estimate    apijson.Field
	Used        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseGasEstimationJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseGasEstimation) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseGasEstimation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EvmTransactionBulkScanResponseGasEstimationUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation],
// [EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError].
func (r EvmTransactionBulkScanResponseGasEstimation) AsUnion() EvmTransactionBulkScanResponseGasEstimationUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation]
// or
// [EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError].
type EvmTransactionBulkScanResponseGasEstimationUnion interface {
	implementsEvmTransactionBulkScanResponseGasEstimation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseGasEstimationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError{}),
		},
	)
}

type EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation struct {
	Estimate string                                                                                        `json:"estimate" api:"required"`
	Status   EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatus `json:"status" api:"required"`
	Used     string                                                                                        `json:"used" api:"required"`
	JSON     evmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationJSON   `json:"-"`
}

// evmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation]
type evmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationJSON struct {
	Estimate    apijson.Field
	Status      apijson.Field
	Used        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation) implementsEvmTransactionBulkScanResponseGasEstimation() {
}

type EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatus string

const (
	EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatusSuccess EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatus = "Success"
)

func (r EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatus) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatusSuccess:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError struct {
	Error  string                                                                                             `json:"error" api:"required"`
	Status EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus `json:"status" api:"required"`
	JSON   evmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON   `json:"-"`
}

// evmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError]
type evmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError) implementsEvmTransactionBulkScanResponseGasEstimation() {
}

type EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus string

const (
	EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatusError EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus = "Error"
)

func (r EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatusError:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseGasEstimationStatus string

const (
	EvmTransactionBulkScanResponseGasEstimationStatusSuccess EvmTransactionBulkScanResponseGasEstimationStatus = "Success"
	EvmTransactionBulkScanResponseGasEstimationStatusError   EvmTransactionBulkScanResponseGasEstimationStatus = "Error"
)

func (r EvmTransactionBulkScanResponseGasEstimationStatus) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseGasEstimationStatusSuccess, EvmTransactionBulkScanResponseGasEstimationStatusError:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulation struct {
	// A string indicating if the simulation was successful or not.
	Status EvmTransactionBulkScanResponseSimulationStatus `json:"status" api:"required"`
	// This field can have the runtime type of
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummary].
	AccountSummary interface{} `json:"account_summary"`
	// This field can have the runtime type of
	// [map[string]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetail].
	AddressDetails interface{} `json:"address_details"`
	// This field can have the runtime type of
	// [map[string][]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff].
	AssetsDiffs interface{} `json:"assets_diffs"`
	// This field can have the runtime type of
	// [map[string][]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement].
	ContractManagement interface{} `json:"contract_management"`
	// A string explaining why the transaction failed
	Description string `json:"description"`
	// An error message if the simulation failed.
	Error string `json:"error"`
	// This field can have the runtime type of
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails].
	ErrorDetails interface{} `json:"error_details"`
	// This field can have the runtime type of
	// [map[string][]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure].
	Exposures interface{} `json:"exposures"`
	// This field can have the runtime type of
	// [[]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalance].
	MissingBalances interface{} `json:"missing_balances"`
	// This field can have the runtime type of
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParams].
	Params interface{} `json:"params"`
	// This field can have the runtime type of
	// [map[string][]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKey].
	SessionKey interface{} `json:"session_key"`
	// The number of times the simulation ran until success
	SimulationRunCount int64 `json:"simulation_run_count"`
	// This field can have the runtime type of
	// [map[string]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiff].
	TotalUsdDiff interface{} `json:"total_usd_diff"`
	// This field can have the runtime type of [map[string]map[string]string].
	TotalUsdExposure interface{} `json:"total_usd_exposure"`
	// This field can have the runtime type of
	// [[]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions].
	TransactionActions interface{}                                  `json:"transaction_actions"`
	JSON               evmTransactionBulkScanResponseSimulationJSON `json:"-"`
	union              EvmTransactionBulkScanResponseSimulationUnion
}

// evmTransactionBulkScanResponseSimulationJSON contains the JSON metadata for the
// struct [EvmTransactionBulkScanResponseSimulation]
type evmTransactionBulkScanResponseSimulationJSON struct {
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

func (r evmTransactionBulkScanResponseSimulationJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulation) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EvmTransactionBulkScanResponseSimulationUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulation],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationError].
func (r EvmTransactionBulkScanResponseSimulation) AsUnion() EvmTransactionBulkScanResponseSimulationUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulation]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationError].
type EvmTransactionBulkScanResponseSimulationUnion interface {
	implementsEvmTransactionBulkScanResponseSimulation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationError{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulation struct {
	// Account summary for the account address. account address is determined implicit
	// by the `from` field in the transaction request, or explicit by the
	// account_address field in the request.
	AccountSummary EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummary `json:"account_summary" api:"required"`
	// a dictionary including additional information about each relevant address in the
	// transaction.
	AddressDetails map[string]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetail `json:"address_details" api:"required"`
	// dictionary describes the assets differences as a result of this transaction for
	// every involved address
	AssetsDiffs map[string][]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff `json:"assets_diffs" api:"required"`
	// dictionary describes the exposure differences as a result of this transaction
	// for every involved address (as a result of any approval / setApproval / permit
	// function)
	Exposures map[string][]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure `json:"exposures" api:"required"`
	// Session keys created in this transaction per address
	SessionKey map[string][]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKey `json:"session_key" api:"required"`
	// A string indicating if the simulation was successful or not.
	Status EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationStatus `json:"status" api:"required"`
	// dictionary represents the usd value each address gained / lost during this
	// transaction
	TotalUsdDiff map[string]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiff `json:"total_usd_diff" api:"required"`
	// a dictionary representing the usd value each address is exposed to, split by
	// spenders
	TotalUsdExposure map[string]map[string]string `json:"total_usd_exposure" api:"required"`
	// Describes the nature of the transaction and what happened as part of it
	TransactionActions []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions `json:"transaction_actions" api:"required"`
	// Describes the state differences as a result of this transaction for every
	// involved address
	ContractManagement map[string][]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement `json:"contract_management"`
	// Missing balances in the transaction
	MissingBalances []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalance `json:"missing_balances"`
	// The parameters of the transaction that was simulated.
	Params EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParams `json:"params"`
	// The number of times the simulation ran until success
	SimulationRunCount int64                                                                               `json:"simulation_run_count"`
	JSON               evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulation]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationJSON struct {
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

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulation) implementsEvmTransactionBulkScanResponseSimulation() {
}

// Account summary for the account address. account address is determined implicit
// by the `from` field in the transaction request, or explicit by the
// account_address field in the request.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummary struct {
	// All assets diffs related to the account address
	AssetsDiffs []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff `json:"assets_diffs" api:"required"`
	// All assets exposures related to the account address
	Exposures []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure `json:"exposures" api:"required"`
	// Total usd diff related to the account address
	TotalUsdDiff EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiff `json:"total_usd_diff" api:"required"`
	// Total usd exposure related to the account address
	TotalUsdExposure map[string]string `json:"total_usd_exposure" api:"required"`
	// All assets traces related to the account address
	Traces []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace `json:"traces" api:"required"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryJSON    `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummary]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryJSON struct {
	AssetsDiffs      apijson.Field
	Exposures        apijson.Field
	TotalUsdDiff     apijson.Field
	TotalUsdExposure apijson.Field
	Traces           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff struct {
	// This field can have the runtime type of
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAsset].
	Asset interface{} `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType `json:"asset_type" api:"required"`
	// This field can have the runtime type of
	// [[]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffIn],
	// [[]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffIn],
	// [[]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffIn],
	// [[]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffIn].
	In interface{} `json:"in" api:"required"`
	// This field can have the runtime type of
	// [[]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOut],
	// [[]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOut],
	// [[]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOut],
	// [[]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOut].
	Out interface{} `json:"out" api:"required"`
	// This field can have the runtime type of
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChanges],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChanges],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChanges],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChanges].
	BalanceChanges interface{}                                                                                                 `json:"balance_changes"`
	JSON           evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffJSON `json:"-"`
	union          EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff struct {
	// description of the asset for the current diff
	Asset EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType `json:"asset_type" api:"required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffIn `json:"in" api:"required"`
	// amount of the asset that was transferred from the address in this transaction
	Out []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOut `json:"out" api:"required"`
	// shows the balance before making the transaction and after
	BalanceChanges EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChanges `json:"balance_changes"`
	JSON           evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffJSON           `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff() {
}

// description of the asset for the current diff
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType `json:"type" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetJSON `json:"-"`
	union  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                       `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                        `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetTypeErc20  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType = "ERC20"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetTypeErc20, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffIn struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                             `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffInJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffInJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffIn]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffInJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOut struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                              `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOutJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOutJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOut]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOutJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOutJSON) RawJSON() string {
	return r.raw
}

// shows the balance before making the transaction and after
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChanges struct {
	// balance of the account after making the transaction
	After EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfter `json:"after" api:"required"`
	// balance of the account before making the transaction
	Before EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBefore `json:"before" api:"required"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesJSON   `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChanges]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChanges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesJSON) RawJSON() string {
	return r.raw
}

// balance of the account after making the transaction
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfter struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                              `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfterJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfterJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfter]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfterJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfterJSON) RawJSON() string {
	return r.raw
}

// balance of the account before making the transaction
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBefore struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                               `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBefore]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff struct {
	// description of the asset for the current diff
	Asset EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType `json:"asset_type" api:"required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffIn `json:"in" api:"required"`
	// amount of the asset that was transferred from the address in this transaction
	Out []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOut `json:"out" api:"required"`
	// shows the balance before making the transaction and after
	BalanceChanges EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChanges `json:"balance_changes"`
	JSON           evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffJSON           `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff() {
}

// description of the asset for the current diff
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                 `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetJSON `json:"-"`
	union  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                         `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                         `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetTypeErc721 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType = "ERC721"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetTypeErc721, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffIn struct {
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
	UsdPrice string                                                                                                                                                              `json:"usd_price"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffInJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffInJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffIn]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffInJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOut struct {
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
	UsdPrice string                                                                                                                                                               `json:"usd_price"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOutJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOutJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOut]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOutJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOutJSON) RawJSON() string {
	return r.raw
}

// shows the balance before making the transaction and after
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChanges struct {
	// balance of the account after making the transaction
	After EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfter `json:"after" api:"required"`
	// balance of the account before making the transaction
	Before EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBefore `json:"before" api:"required"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesJSON   `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChanges]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChanges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesJSON) RawJSON() string {
	return r.raw
}

// balance of the account after making the transaction
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfter struct {
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
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfterJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfterJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfter]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfterJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfterJSON) RawJSON() string {
	return r.raw
}

// balance of the account before making the transaction
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBefore struct {
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
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBefore]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff struct {
	// description of the asset for the current diff
	Asset EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType `json:"asset_type" api:"required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffIn `json:"in" api:"required"`
	// amount of the asset that was transferred from the address in this transaction
	Out []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOut `json:"out" api:"required"`
	// shows the balance before making the transaction and after
	BalanceChanges EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChanges `json:"balance_changes"`
	JSON           evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffJSON           `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff() {
}

// description of the asset for the current diff
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                  `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetJSON `json:"-"`
	union  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                           `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                          `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetTypeErc1155 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType = "ERC1155"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetTypeNonerc  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetTypeErc1155, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffIn struct {
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
	UsdPrice string                                                                                                                                                               `json:"usd_price"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffInJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffInJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffIn]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffInJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOut struct {
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
	UsdPrice string                                                                                                                                                                `json:"usd_price"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOutJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOutJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOut]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOutJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOutJSON) RawJSON() string {
	return r.raw
}

// shows the balance before making the transaction and after
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChanges struct {
	// balance of the account after making the transaction
	After EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfter `json:"after" api:"required"`
	// balance of the account before making the transaction
	Before EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBefore `json:"before" api:"required"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesJSON   `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChanges]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChanges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesJSON) RawJSON() string {
	return r.raw
}

// balance of the account after making the transaction
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfter struct {
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
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfterJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfterJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfter]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfterJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfterJSON) RawJSON() string {
	return r.raw
}

// balance of the account before making the transaction
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBefore struct {
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
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBefore]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff struct {
	// description of the asset for the current diff
	Asset EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetType `json:"asset_type" api:"required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffIn `json:"in" api:"required"`
	// amount of the asset that was transferred from the address in this transaction
	Out []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOut `json:"out" api:"required"`
	// shows the balance before making the transaction and after
	BalanceChanges EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChanges `json:"balance_changes"`
	JSON           evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffJSON           `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff() {
}

// description of the asset for the current diff
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAsset struct {
	ChainID   int64  `json:"chain_id" api:"required"`
	ChainName string `json:"chain_name" api:"required"`
	Decimals  int64  `json:"decimals" api:"required"`
	LogoURL   string `json:"logo_url" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetType `json:"type" api:"required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                 `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAsset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetJSON struct {
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

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetTypeNative EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetType = "NATIVE"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetTypeNative:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffIn struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                              `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffInJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffInJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffIn]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffInJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOut struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                               `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOutJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOutJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOut]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOutJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOutJSON) RawJSON() string {
	return r.raw
}

// shows the balance before making the transaction and after
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChanges struct {
	// balance of the account after making the transaction
	After EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfter `json:"after" api:"required"`
	// balance of the account before making the transaction
	Before EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBefore `json:"before" api:"required"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesJSON   `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChanges]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChanges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesJSON) RawJSON() string {
	return r.raw
}

// balance of the account after making the transaction
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfter struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                               `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfterJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfterJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfter]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfterJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfterJSON) RawJSON() string {
	return r.raw
}

// balance of the account before making the transaction
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBefore struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBeforeJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBeforeJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBefore]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBeforeJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBeforeJSON) RawJSON() string {
	return r.raw
}

// type of the asset for the current diff
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc20   EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType = "ERC20"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc721  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType = "ERC721"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc1155 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType = "ERC1155"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeNative  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType = "NATIVE"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc20, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc721, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc1155, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeNative:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure struct {
	// This field can have the runtime type of
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset].
	Asset interface{} `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType `json:"asset_type" api:"required"`
	// This field can have the runtime type of
	// [map[string]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpender],
	// [map[string]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpender],
	// [map[string]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpender].
	Spenders interface{}                                                                                               `json:"spenders" api:"required"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposureJSON `json:"-"`
	union    EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType `json:"asset_type" api:"required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpender `json:"spenders" api:"required"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureJSON               `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure() {
}

// description of the asset for the current diff
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType `json:"type" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetJSON `json:"-"`
	union  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                       `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                        `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetTypeErc20  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType = "ERC20"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetTypeErc20, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpender struct {
	// the amount that was asked in the approval request for this spender from the
	// current address and asset
	Approval string                                                                                                                                                         `json:"approval" api:"required"`
	Exposure []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure `json:"exposure" api:"required"`
	// the usd price of the approval amount
	ApprovalUsdPrice string `json:"approval_usd_price"`
	// the expiration time of the permit2 protocol
	Expiration time.Time `json:"expiration" format:"date-time"`
	// user friendly description of the approval
	Summary string                                                                                                                                                  `json:"summary"`
	JSON    evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpender]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON struct {
	Approval         apijson.Field
	Exposure         apijson.Field
	ApprovalUsdPrice apijson.Field
	Expiration       apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure struct {
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
	Value string                                                                                                                                                           `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON `json:"-"`
	union EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON struct {
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

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
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
	UsdPrice string                                                                                                                                                                                        `json:"usd_price"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
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
	UsdPrice string                                                                                                                                                                                       `json:"usd_price"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                      `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                       `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType `json:"asset_type" api:"required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpender `json:"spenders" api:"required"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureJSON               `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure() {
}

// description of the asset for the current diff
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                 `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetJSON `json:"-"`
	union  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                         `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                         `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetTypeErc721 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType = "ERC721"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetTypeErc721, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpender struct {
	Exposure []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure `json:"exposure" api:"required"`
	// boolean indicates whether an is_approved_for_all function was used (missing in
	// case of ERC20 / ERC1155)
	IsApprovedForAll bool `json:"is_approved_for_all" api:"required"`
	// user friendly description of the approval
	Summary string                                                                                                                                                   `json:"summary"`
	JSON    evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpender]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure struct {
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
	Value string                                                                                                                                                            `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON `json:"-"`
	union EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON struct {
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

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
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
	UsdPrice string                                                                                                                                                                                         `json:"usd_price"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
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
	UsdPrice string                                                                                                                                                                                        `json:"usd_price"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                       `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                        `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType `json:"asset_type" api:"required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpender `json:"spenders" api:"required"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureJSON               `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure() {
}

// description of the asset for the current diff
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                  `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON `json:"-"`
	union  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                           `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                          `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeErc1155 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType = "ERC1155"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeNonerc  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeErc1155, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpender struct {
	Exposure []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure `json:"exposure" api:"required"`
	// boolean indicates whether an is_approved_for_all function was used (missing in
	// case of ERC20 / ERC1155)
	IsApprovedForAll bool `json:"is_approved_for_all" api:"required"`
	// user friendly description of the approval
	Summary string                                                                                                                                                    `json:"summary"`
	JSON    evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpender]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure struct {
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
	Value string                                                                                                                                                             `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON `json:"-"`
	union EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON struct {
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

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
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
	UsdPrice string                                                                                                                                                                                          `json:"usd_price"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
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
	UsdPrice string                                                                                                                                                                                         `json:"usd_price"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                        `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                         `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

// type of the asset for the current diff
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc20   EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType = "ERC20"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc721  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType = "ERC721"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc1155 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType = "ERC1155"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc20, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc721, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc1155:
		return true
	}
	return false
}

// Total usd diff related to the account address
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiff struct {
	In    string                                                                                                        `json:"in" api:"required"`
	Out   string                                                                                                        `json:"out" api:"required"`
	Total string                                                                                                        `json:"total" api:"required"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace struct {
	// This field can have the runtime type of
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAsset],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset].
	Asset interface{} `json:"asset" api:"required"`
	// type of the trace
	TraceType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceType `json:"trace_type" api:"required"`
	// The type of the model
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType `json:"type" api:"required"`
	// This field can have the runtime type of
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiff],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiff],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiff],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiff].
	Diff interface{} `json:"diff"`
	// This field can have the runtime type of
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposed],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposed].
	Exposed interface{} `json:"exposed"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address"`
	// This field can have the runtime type of
	// [[]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabels],
	// [[]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabels],
	// [[]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabels],
	// [[]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabels].
	Labels interface{} `json:"labels"`
	// The owner of the assets
	Owner string `json:"owner"`
	// The spender of the assets
	Spender string `json:"spender"`
	// The address where the assets are moved to
	ToAddress string                                                                                                 `json:"to_address"`
	JSON      evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTraceJSON `json:"-"`
	union     EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTraceJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTraceJSON struct {
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

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTraceJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace struct {
	// Description of the asset in the trace
	Asset EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset `json:"asset" api:"required"`
	// The difference in value for the asset in the trace
	Diff EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiff `json:"diff" api:"required"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address" api:"required"`
	// The address where the assets are moved to
	ToAddress string `json:"to_address" api:"required"`
	// type of the trace
	TraceType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceType `json:"trace_type" api:"required"`
	// The type of the model
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceType `json:"type" api:"required"`
	// List of labels that describe the trace
	Labels []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabels `json:"labels"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceJSON     `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceJSON struct {
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

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetType `json:"type" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                        `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetJSON `json:"-"`
	union  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                               `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetTypeErc20  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetType = "ERC20"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetTypeErc20, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetTypeNonerc:
		return true
	}
	return false
}

// The difference in value for the asset in the trace
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                       `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiffJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceTypeAssetTrace EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceType = "AssetTrace"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceTypeAssetTrace:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTypeErc20AssetTrace EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceType = "ERC20AssetTrace"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTypeErc20AssetTrace:
		return true
	}
	return false
}

// Label assigned to a trace entry.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabels string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabelsGasFee EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabels = "GAS_FEE"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabels) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabelsGasFee:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace struct {
	// Description of the asset in the trace
	Asset EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset `json:"asset" api:"required"`
	// The difference in value for the asset in the trace
	Diff EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiff `json:"diff" api:"required"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address" api:"required"`
	// The address where the assets are moved to
	ToAddress string `json:"to_address" api:"required"`
	// type of the trace
	TraceType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceType `json:"trace_type" api:"required"`
	// The type of the model
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceType `json:"type" api:"required"`
	// List of labels that describe the trace
	Labels []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabels `json:"labels"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceJSON     `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceJSON struct {
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

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                         `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetJSON `json:"-"`
	union  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                 `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                 `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetTypeErc721 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetType = "ERC721"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetTypeErc721, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetTypeNonerc:
		return true
	}
	return false
}

// The difference in value for the asset in the trace
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiff struct {
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
	UsdPrice string                                                                                                                                        `json:"usd_price"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiffJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceTypeAssetTrace EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceType = "AssetTrace"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceTypeAssetTrace:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTypeErc721AssetTrace EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceType = "ERC721AssetTrace"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTypeErc721AssetTrace:
		return true
	}
	return false
}

// Label assigned to a trace entry.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabels string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabelsGasFee EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabels = "GAS_FEE"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabels) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabelsGasFee:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace struct {
	// Description of the asset in the trace
	Asset EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset `json:"asset" api:"required"`
	// The difference in value for the asset in the trace
	Diff EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiff `json:"diff" api:"required"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address" api:"required"`
	// The address where the assets are moved to
	ToAddress string `json:"to_address" api:"required"`
	// type of the trace
	TraceType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceType `json:"trace_type" api:"required"`
	// The type of the model
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceType `json:"type" api:"required"`
	// List of labels that describe the trace
	Labels []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabels `json:"labels"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceJSON     `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceJSON struct {
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

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                          `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetJSON `json:"-"`
	union  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                   `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                  `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetTypeErc1155 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetType = "ERC1155"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetTypeNonerc  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetTypeErc1155, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetTypeNonerc:
		return true
	}
	return false
}

// The difference in value for the asset in the trace
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiff struct {
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
	UsdPrice string                                                                                                                                         `json:"usd_price"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiffJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceTypeAssetTrace EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceType = "AssetTrace"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceTypeAssetTrace:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTypeErc1155AssetTrace EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceType = "ERC1155AssetTrace"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTypeErc1155AssetTrace:
		return true
	}
	return false
}

// Label assigned to a trace entry.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabels string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabelsGasFee EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabels = "GAS_FEE"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabels) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabelsGasFee:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace struct {
	// Description of the asset in the trace
	Asset EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAsset `json:"asset" api:"required"`
	// The difference in value for the asset in the trace
	Diff EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiff `json:"diff" api:"required"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address" api:"required"`
	// The address where the assets are moved to
	ToAddress string `json:"to_address" api:"required"`
	// type of the trace
	TraceType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceType `json:"trace_type" api:"required"`
	// The type of the model
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceType `json:"type" api:"required"`
	// List of labels that describe the trace
	Labels []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabels `json:"labels"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceJSON     `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceJSON struct {
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

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAsset struct {
	ChainID   int64  `json:"chain_id" api:"required"`
	ChainName string `json:"chain_name" api:"required"`
	Decimals  int64  `json:"decimals" api:"required"`
	LogoURL   string `json:"logo_url" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetType `json:"type" api:"required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                         `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAsset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetJSON struct {
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

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetTypeNative EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetType = "NATIVE"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetTypeNative:
		return true
	}
	return false
}

// The difference in value for the asset in the trace
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                        `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiffJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceTypeAssetTrace EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceType = "AssetTrace"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceTypeAssetTrace:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTypeNativeAssetTrace EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceType = "NativeAssetTrace"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTypeNativeAssetTrace:
		return true
	}
	return false
}

// Label assigned to a trace entry.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabels string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabelsGasFee EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabels = "GAS_FEE"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabels) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabelsGasFee:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace struct {
	// Description of the asset in the trace
	Asset   EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset   `json:"asset" api:"required"`
	Exposed EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposed `json:"exposed" api:"required"`
	// The owner of the assets
	Owner string `json:"owner" api:"required"`
	// The spender of the assets
	Spender string `json:"spender" api:"required"`
	// type of the trace
	TraceType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceType `json:"trace_type" api:"required"`
	// The type of the model
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceType `json:"type" api:"required"`
	JSON evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceJSON struct {
	Asset       apijson.Field
	Exposed     apijson.Field
	Owner       apijson.Field
	Spender     apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetType `json:"type" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                           `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetJSON `json:"-"`
	union  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                  `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                   `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetTypeErc20  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetType = "ERC20"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetTypeErc20, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposed struct {
	RawValue string                                                                                                                                             `json:"raw_value" api:"required"`
	UsdPrice float64                                                                                                                                            `json:"usd_price"`
	Value    float64                                                                                                                                            `json:"value"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposedJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposedJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposed]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposedJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposedJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceTypeExposureTrace EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceType = "ExposureTrace"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceTypeExposureTrace:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTypeErc20ExposureTrace EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceType = "ERC20ExposureTrace"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTypeErc20ExposureTrace:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace struct {
	// Description of the asset in the trace
	Asset EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset `json:"asset" api:"required"`
	// The owner of the assets
	Owner string `json:"owner" api:"required"`
	// The spender of the assets
	Spender string `json:"spender" api:"required"`
	// type of the trace
	TraceType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceType `json:"trace_type" api:"required"`
	// The type of the model
	Type    EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceType    `json:"type" api:"required"`
	Exposed EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposed `json:"exposed"`
	JSON    evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceJSON    `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceJSON struct {
	Asset       apijson.Field
	Owner       apijson.Field
	Spender     apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	Exposed     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                            `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetJSON `json:"-"`
	union  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                    `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                    `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetTypeErc721 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetType = "ERC721"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetTypeErc721, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetTypeNonerc:
		return true
	}
	return false
}

// type of the trace
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceTypeExposureTrace EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceType = "ExposureTrace"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceTypeExposureTrace:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTypeErc721ExposureTrace EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceType = "ERC721ExposureTrace"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTypeErc721ExposureTrace:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposed struct {
	Amount   int64                                                                                                                                               `json:"amount" api:"required"`
	TokenID  string                                                                                                                                              `json:"token_id" api:"required"`
	IsMint   bool                                                                                                                                                `json:"is_mint"`
	LogoURL  string                                                                                                                                              `json:"logo_url"`
	UsdPrice float64                                                                                                                                             `json:"usd_price"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposedJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposedJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposed]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposedJSON struct {
	Amount      apijson.Field
	TokenID     apijson.Field
	IsMint      apijson.Field
	LogoURL     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposedJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace struct {
	// Description of the asset in the trace
	Asset EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset `json:"asset" api:"required"`
	// The owner of the assets
	Owner string `json:"owner" api:"required"`
	// The spender of the assets
	Spender string `json:"spender" api:"required"`
	// type of the trace
	TraceType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceType `json:"trace_type" api:"required"`
	// The type of the model
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceType `json:"type" api:"required"`
	JSON evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceJSON struct {
	Asset       apijson.Field
	Owner       apijson.Field
	Spender     apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                             `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetJSON `json:"-"`
	union  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                      `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                     `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetTypeErc1155 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetType = "ERC1155"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetTypeNonerc  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetTypeErc1155, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetTypeNonerc:
		return true
	}
	return false
}

// type of the trace
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceTypeExposureTrace EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceType = "ExposureTrace"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceTypeExposureTrace:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTypeErc1155ExposureTrace EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceType = "ERC1155ExposureTrace"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTypeErc1155ExposureTrace:
		return true
	}
	return false
}

// type of the trace
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceTypeAssetTrace    EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceType = "AssetTrace"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceTypeExposureTrace EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceType = "ExposureTrace"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceTypeAssetTrace, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceTypeExposureTrace:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc20AssetTrace      EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC20AssetTrace"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc721AssetTrace     EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC721AssetTrace"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc1155AssetTrace    EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC1155AssetTrace"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeNativeAssetTrace     EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "NativeAssetTrace"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc20ExposureTrace   EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC20ExposureTrace"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc721ExposureTrace  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC721ExposureTrace"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc1155ExposureTrace EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC1155ExposureTrace"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc20AssetTrace, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc721AssetTrace, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc1155AssetTrace, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeNativeAssetTrace, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc20ExposureTrace, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc721ExposureTrace, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc1155ExposureTrace:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetail struct {
	// contains the contract's name if the address is a verified contract
	ContractName string `json:"contract_name"`
	// Whether the address is an eoa or a contract
	IsEoa bool `json:"is_eoa"`
	// known name tag for the address
	NameTag string                                                                                           `json:"name_tag"`
	JSON    evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetailJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetailJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetail]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetailJSON struct {
	ContractName apijson.Field
	IsEoa        apijson.Field
	NameTag      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetailJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff struct {
	// This field can have the runtime type of
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAsset].
	Asset interface{} `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType `json:"asset_type" api:"required"`
	// This field can have the runtime type of
	// [[]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffIn],
	// [[]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffIn],
	// [[]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffIn],
	// [[]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffIn].
	In interface{} `json:"in" api:"required"`
	// This field can have the runtime type of
	// [[]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOut],
	// [[]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOut],
	// [[]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOut],
	// [[]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOut].
	Out   interface{}                                                                                   `json:"out" api:"required"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffJSON `json:"-"`
	union EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff struct {
	// description of the asset for the current diff
	Asset EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType `json:"asset_type" api:"required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffIn `json:"in" api:"required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOut `json:"out" api:"required"`
	JSON evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffJSON  `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff() {
}

// description of the asset for the current diff
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType `json:"type" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                     `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetJSON `json:"-"`
	union  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                            `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                             `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetTypeErc20  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType = "ERC20"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetTypeErc20, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffIn struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                  `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffInJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffInJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffIn]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffInJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOut struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                   `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOutJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOutJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOut]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff struct {
	// description of the asset for the current diff
	Asset EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType `json:"asset_type" api:"required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffIn `json:"in" api:"required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOut `json:"out" api:"required"`
	JSON evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffJSON  `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff() {
}

// description of the asset for the current diff
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                      `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetJSON `json:"-"`
	union  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                              `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                              `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetTypeErc721 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType = "ERC721"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetTypeErc721, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffIn struct {
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
	UsdPrice string                                                                                                                                   `json:"usd_price"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffInJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffInJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffIn]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffInJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOut struct {
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
	UsdPrice string                                                                                                                                    `json:"usd_price"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOutJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOutJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOut]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOutJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff struct {
	// description of the asset for the current diff
	Asset EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType `json:"asset_type" api:"required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffIn `json:"in" api:"required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOut `json:"out" api:"required"`
	JSON evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffJSON  `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff() {
}

// description of the asset for the current diff
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                       `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetJSON `json:"-"`
	union  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                               `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetTypeErc1155 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType = "ERC1155"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetTypeNonerc  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetTypeErc1155, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffIn struct {
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
	UsdPrice string                                                                                                                                    `json:"usd_price"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffInJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffInJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffIn]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffInJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOut struct {
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
	UsdPrice string                                                                                                                                     `json:"usd_price"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOutJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOutJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOut]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOutJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff struct {
	// description of the asset for the current diff
	Asset EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetType `json:"asset_type" api:"required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffIn `json:"in" api:"required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOut `json:"out" api:"required"`
	JSON evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffJSON  `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff() {
}

// description of the asset for the current diff
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAsset struct {
	ChainID   int64  `json:"chain_id" api:"required"`
	ChainName string `json:"chain_name" api:"required"`
	Decimals  int64  `json:"decimals" api:"required"`
	LogoURL   string `json:"logo_url" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetType `json:"type" api:"required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                      `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAsset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetJSON struct {
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

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetTypeNative EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetType = "NATIVE"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetTypeNative:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffIn struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                   `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffInJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffInJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffIn]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffInJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOut struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                    `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOutJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOutJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOut]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

// type of the asset for the current diff
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc20   EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType = "ERC20"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc721  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType = "ERC721"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc1155 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType = "ERC1155"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeNative  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType = "NATIVE"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc20, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc721, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc1155, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeNative:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure struct {
	// This field can have the runtime type of
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset].
	Asset interface{} `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType `json:"asset_type" api:"required"`
	// This field can have the runtime type of
	// [map[string]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpender],
	// [map[string]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpender],
	// [map[string]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpender].
	Spenders interface{}                                                                                 `json:"spenders" api:"required"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposureJSON `json:"-"`
	union    EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType `json:"asset_type" api:"required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpender `json:"spenders" api:"required"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureJSON               `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure() {
}

// description of the asset for the current diff
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType `json:"type" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                  `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetJSON `json:"-"`
	union  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                         `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                          `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetTypeErc20  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType = "ERC20"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetTypeErc20, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpender struct {
	// the amount that was asked in the approval request for this spender from the
	// current address and asset
	Approval string                                                                                                                                           `json:"approval" api:"required"`
	Exposure []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure `json:"exposure" api:"required"`
	// the usd price of the approval amount
	ApprovalUsdPrice string `json:"approval_usd_price"`
	// the expiration time of the permit2 protocol
	Expiration time.Time `json:"expiration" format:"date-time"`
	// user friendly description of the approval
	Summary string                                                                                                                                    `json:"summary"`
	JSON    evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpender]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON struct {
	Approval         apijson.Field
	Exposure         apijson.Field
	ApprovalUsdPrice apijson.Field
	Expiration       apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure struct {
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
	Value string                                                                                                                                             `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON `json:"-"`
	union EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON struct {
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

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
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
	UsdPrice string                                                                                                                                                                          `json:"usd_price"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
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
	UsdPrice string                                                                                                                                                                         `json:"usd_price"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                        `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                         `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType `json:"asset_type" api:"required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpender `json:"spenders" api:"required"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureJSON               `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure() {
}

// description of the asset for the current diff
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                   `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetJSON `json:"-"`
	union  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                           `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                           `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetTypeErc721 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType = "ERC721"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetTypeErc721, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpender struct {
	Exposure []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure `json:"exposure" api:"required"`
	// boolean indicates whether an is_approved_for_all function was used (missing in
	// case of ERC20 / ERC1155)
	IsApprovedForAll bool `json:"is_approved_for_all" api:"required"`
	// user friendly description of the approval
	Summary string                                                                                                                                     `json:"summary"`
	JSON    evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpender]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure struct {
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
	Value string                                                                                                                                              `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON `json:"-"`
	union EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON struct {
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

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
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
	UsdPrice string                                                                                                                                                                           `json:"usd_price"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
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
	UsdPrice string                                                                                                                                                                          `json:"usd_price"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                         `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                          `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset `json:"asset" api:"required"`
	// type of the asset for the current diff
	AssetType EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType `json:"asset_type" api:"required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpender `json:"spenders" api:"required"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureJSON               `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure() {
}

// description of the asset for the current diff
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                    `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON `json:"-"`
	union  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                             `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                            `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeErc1155 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType = "ERC1155"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeNonerc  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType = "NONERC"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeErc1155, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpender struct {
	Exposure []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure `json:"exposure" api:"required"`
	// boolean indicates whether an is_approved_for_all function was used (missing in
	// case of ERC20 / ERC1155)
	IsApprovedForAll bool `json:"is_approved_for_all" api:"required"`
	// user friendly description of the approval
	Summary string                                                                                                                                      `json:"summary"`
	JSON    evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpender]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure struct {
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
	Value string                                                                                                                                               `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON `json:"-"`
	union EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON struct {
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

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
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
	UsdPrice string                                                                                                                                                                            `json:"usd_price"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
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
	UsdPrice string                                                                                                                                                                           `json:"usd_price"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                          `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value" api:"required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                           `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

// type of the asset for the current diff
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc20   EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType = "ERC20"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc721  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType = "ERC721"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc1155 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType = "ERC1155"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc20, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc721, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc1155:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKey struct {
	Key      string                                                                                            `json:"key" api:"required"`
	Policies []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy `json:"policies" api:"required"`
	Signer   string                                                                                            `json:"signer" api:"required"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyJSON     `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKey]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyJSON struct {
	Key         apijson.Field
	Policies    apijson.Field
	Signer      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy struct {
	// This field can have the runtime type of
	// [[]EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArg].
	Args interface{} `json:"args"`
	// This field can have the runtime type of
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails].
	AssetDetails interface{}                                                                                           `json:"asset_details"`
	EndTime      time.Time                                                                                             `json:"end_time" format:"date-time"`
	Method       string                                                                                                `json:"method"`
	Recipient    string                                                                                                `json:"recipient"`
	StartTime    time.Time                                                                                             `json:"start_time" format:"date-time"`
	ToAddress    string                                                                                                `json:"to_address"`
	Type         EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType `json:"type"`
	ValueLimit   string                                                                                                `json:"value_limit"`
	JSON         evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicyJSON   `json:"-"`
	union        EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicyJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicyJSON struct {
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

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicyJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy struct {
	AssetDetails EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails `json:"asset_details" api:"required"`
	Recipient    string                                                                                                                                           `json:"recipient"`
	Type         EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyType         `json:"type"`
	ValueLimit   string                                                                                                                                           `json:"value_limit"`
	JSON         evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyJSON         `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyJSON struct {
	AssetDetails apijson.Field
	Recipient    apijson.Field
	Type         apijson.Field
	ValueLimit   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails struct {
	Decimals int64 `json:"decimals" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsType `json:"type" api:"required"`
	// address of the token
	Address   string `json:"address"`
	ChainID   int64  `json:"chain_id"`
	ChainName string `json:"chain_name"`
	LogoURL   string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                               `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsJSON `json:"-"`
	union  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsJSON struct {
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

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails struct {
	ChainID   int64  `json:"chain_id" api:"required"`
	ChainName string `json:"chain_name" api:"required"`
	Decimals  int64  `json:"decimals" api:"required"`
	LogoURL   string `json:"logo_url" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsType `json:"type" api:"required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                       `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsJSON struct {
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

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsTypeNative EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsType = "NATIVE"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsTypeNative:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                      `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsTypeNative EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsType = "NATIVE"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsTypeErc20  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsType = "ERC20"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsTypeNative, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsTypeErc20:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyTypeTransferPolicy EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyType = "TRANSFER_POLICY"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyTypeTransferPolicy:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy struct {
	ToAddress  string                                                                                                                                `json:"to_address" api:"required"`
	Args       []EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArg `json:"args"`
	Method     string                                                                                                                                `json:"method"`
	Type       EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyType  `json:"type"`
	ValueLimit string                                                                                                                                `json:"value_limit"`
	JSON       evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyJSON  `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyJSON struct {
	ToAddress   apijson.Field
	Args        apijson.Field
	Method      apijson.Field
	Type        apijson.Field
	ValueLimit  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArg struct {
	// Comparison operator used to evaluate an argument/value against a policy
	// constraint.
	Condition EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition `json:"condition" api:"required"`
	Index     int64                                                                                                                                         `json:"index" api:"required"`
	Value     string                                                                                                                                        `json:"value" api:"required"`
	JSON      evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgJSON       `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArg]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgJSON struct {
	Condition   apijson.Field
	Index       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgJSON) RawJSON() string {
	return r.raw
}

// Comparison operator used to evaluate an argument/value against a policy
// constraint.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionUnconstrained  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "UNCONSTRAINED"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionEqual          EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "EQUAL"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionGreater        EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "GREATER"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionLess           EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "LESS"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionGreaterOrEqual EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "GREATER_OR_EQUAL"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionLessOrEqual    EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "LESS_OR_EQUAL"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionNotEqual       EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "NOT_EQUAL"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionUnconstrained, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionEqual, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionGreater, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionLess, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionGreaterOrEqual, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionLessOrEqual, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionNotEqual:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyTypeCallPolicy EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyType = "CALL_POLICY"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyTypeCallPolicy:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy struct {
	ValueLimit string                                                                                                                              `json:"value_limit" api:"required"`
	Type       EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyType `json:"type"`
	JSON       evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyJSON struct {
	ValueLimit  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyTypeGasPolicy EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyType = "GAS_POLICY"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyTypeGasPolicy:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy struct {
	EndTime   time.Time                                                                                                                                  `json:"end_time" format:"date-time"`
	StartTime time.Time                                                                                                                                  `json:"start_time" format:"date-time"`
	Type      EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyType `json:"type"`
	JSON      evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyTypeExpirationPolicy EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyType = "EXPIRATION_POLICY"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyTypeExpirationPolicy:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeTransferPolicy   EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType = "TRANSFER_POLICY"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeCallPolicy       EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType = "CALL_POLICY"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeGasPolicy        EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType = "GAS_POLICY"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeExpirationPolicy EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType = "EXPIRATION_POLICY"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeTransferPolicy, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeCallPolicy, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeGasPolicy, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeExpirationPolicy:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationStatus string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationStatusSuccess EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationStatus = "Success"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationStatus) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationStatusSuccess:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiff struct {
	In    string                                                                                          `json:"in" api:"required"`
	Out   string                                                                                          `json:"out" api:"required"`
	Total string                                                                                          `json:"total" api:"required"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiffJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiff]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

// Type of action performed in the transaction.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsMint            EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "mint"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsStake           EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "stake"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsSwap            EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "swap"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsNativeTransfer  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "native_transfer"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsTokenTransfer   EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "token_transfer"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsApproval        EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "approval"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsSetCodeAccount  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "set_code_account"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsProxyUpgrade    EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "proxy_upgrade"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsOwnershipChange EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "ownership_change"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsMint, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsStake, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsSwap, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsNativeTransfer, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsTokenTransfer, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsApproval, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsSetCodeAccount, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsProxyUpgrade, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsOwnershipChange:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement struct {
	// The type of the state change
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType `json:"type" api:"required"`
	// This field can have the runtime type of
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfter],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfter],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfter].
	After interface{} `json:"after"`
	// This field can have the runtime type of
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBefore],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBefore],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBefore].
	Before interface{} `json:"before"`
	// The delegated address
	DelegatedAddress string `json:"delegated_address"`
	// The direct creator address of the new contract
	Deployer string                                                                                                `json:"deployer"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementJSON `json:"-"`
	union    EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementJSON struct {
	Type             apijson.Field
	After            apijson.Field
	Before           apijson.Field
	DelegatedAddress apijson.Field
	Deployer         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement struct {
	// The state after the transaction
	After EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfter `json:"after" api:"required"`
	// The state before the transaction
	Before EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBefore `json:"before" api:"required"`
	// The type of the state change
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementType `json:"type" api:"required"`
	JSON evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementJSON struct {
	After       apijson.Field
	Before      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement() {
}

// The state after the transaction
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfter struct {
	Address string                                                                                                                                             `json:"address" api:"required"`
	JSON    evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfterJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfterJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfter]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfterJSON struct {
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfterJSON) RawJSON() string {
	return r.raw
}

// The state before the transaction
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBefore struct {
	Address string                                                                                                                                              `json:"address" api:"required"`
	JSON    evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBeforeJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBeforeJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBefore]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBeforeJSON struct {
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBeforeJSON) RawJSON() string {
	return r.raw
}

// The type of the state change
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementTypeProxyUpgrade EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementType = "PROXY_UPGRADE"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementTypeProxyUpgrade:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement struct {
	// The state after the transaction
	After EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfter `json:"after" api:"required"`
	// The state before the transaction
	Before EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBefore `json:"before" api:"required"`
	// The type of the state change
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementType `json:"type" api:"required"`
	JSON evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementJSON struct {
	After       apijson.Field
	Before      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement() {
}

// The state after the transaction
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfter struct {
	Owners []string                                                                                                                                              `json:"owners" api:"required"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfterJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfterJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfter]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfterJSON struct {
	Owners      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfterJSON) RawJSON() string {
	return r.raw
}

// The state before the transaction
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBefore struct {
	Owners []string                                                                                                                                               `json:"owners" api:"required"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBeforeJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBeforeJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBefore]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBeforeJSON struct {
	Owners      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBeforeJSON) RawJSON() string {
	return r.raw
}

// The type of the state change
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementTypeOwnershipChange EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementType = "OWNERSHIP_CHANGE"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementTypeOwnershipChange:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement struct {
	// The state after the transaction
	After EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfter `json:"after" api:"required"`
	// The state before the transaction
	Before EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBefore `json:"before" api:"required"`
	// The type of the state change
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementType `json:"type" api:"required"`
	JSON evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementJSON struct {
	After       apijson.Field
	Before      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement() {
}

// The state after the transaction
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfter struct {
	Modules []string                                                                                                                                            `json:"modules" api:"required"`
	JSON    evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfterJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfterJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfter]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfterJSON struct {
	Modules     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfterJSON) RawJSON() string {
	return r.raw
}

// The state before the transaction
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBefore struct {
	Modules []string                                                                                                                                             `json:"modules" api:"required"`
	JSON    evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBeforeJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBeforeJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBefore]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBeforeJSON struct {
	Modules     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBeforeJSON) RawJSON() string {
	return r.raw
}

// The type of the state change
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementTypeModuleChange EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementType = "MODULE_CHANGE"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementTypeModuleChange:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement struct {
	// The delegated address
	DelegatedAddress string `json:"delegated_address" api:"required"`
	// The type of the state change
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementType `json:"type" api:"required"`
	JSON evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementJSON struct {
	DelegatedAddress apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement() {
}

// The type of the state change
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementTypeSetCodeAccount EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementType = "SET_CODE_ACCOUNT"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementTypeSetCodeAccount:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation struct {
	// The direct creator address of the new contract
	Deployer string `json:"deployer" api:"required"`
	// The type of the state change
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationType `json:"type" api:"required"`
	JSON evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationJSON struct {
	Deployer    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement() {
}

// The type of the state change
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationTypeContractCreation EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationType = "CONTRACT_CREATION"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationTypeContractCreation:
		return true
	}
	return false
}

// The type of the state change
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeProxyUpgrade     EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType = "PROXY_UPGRADE"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeOwnershipChange  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType = "OWNERSHIP_CHANGE"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeModuleChange     EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType = "MODULE_CHANGE"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeSetCodeAccount   EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType = "SET_CODE_ACCOUNT"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeContractCreation EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType = "CONTRACT_CREATION"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeProxyUpgrade, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeOwnershipChange, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeModuleChange, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeSetCodeAccount, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeContractCreation:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalance struct {
	// The asset that is missing balance
	Asset EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset `json:"asset" api:"required"`
	// The account address's current balance of the asset
	CurrentBalance EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalance `json:"current_balance" api:"required"`
	// The account address's missing balance of the asset
	MissingBalance EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalance `json:"missing_balance" api:"required"`
	// The required balance of the asset for this action
	RequiredBalance EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalance `json:"required_balance" api:"required"`
	JSON            evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalanceJSON             `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalanceJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalance]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalanceJSON struct {
	Asset           apijson.Field
	CurrentBalance  apijson.Field
	MissingBalance  apijson.Field
	RequiredBalance apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalanceJSON) RawJSON() string {
	return r.raw
}

// The asset that is missing balance
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset struct {
	// asset's decimals
	Decimals int64 `json:"decimals" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetType `json:"type" api:"required"`
	// address of the token
	Address   string `json:"address"`
	ChainID   int64  `json:"chain_id"`
	ChainName string `json:"chain_name"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                  `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetJSON `json:"-"`
	union  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetJSON struct {
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

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetUnion {
	return r.union
}

// The asset that is missing balance
//
// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                         `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails struct {
	ChainID   int64  `json:"chain_id" api:"required"`
	ChainName string `json:"chain_name" api:"required"`
	Decimals  int64  `json:"decimals" api:"required"`
	LogoURL   string `json:"logo_url" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsType `json:"type" api:"required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                          `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsJSON struct {
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

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset() {
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsTypeNative EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsType = "NATIVE"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsTypeNative:
		return true
	}
	return false
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetTypeErc20  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetType = "ERC20"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetTypeNative EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetType = "NATIVE"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetTypeErc20, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetTypeNative:
		return true
	}
	return false
}

// The account address's current balance of the asset
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalance struct {
	// The raw value of the balance in hex string format
	RawValue string `json:"raw_value" api:"required"`
	// The value of the balance in decimal string format
	Value string                                                                                                           `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalanceJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalanceJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalance]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalanceJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalanceJSON) RawJSON() string {
	return r.raw
}

// The account address's missing balance of the asset
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalance struct {
	// The raw value of the balance in hex string format
	RawValue string `json:"raw_value" api:"required"`
	// The value of the balance in decimal string format
	Value string                                                                                                           `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalanceJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalanceJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalance]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalanceJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalanceJSON) RawJSON() string {
	return r.raw
}

// The required balance of the asset for this action
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalance struct {
	// The raw value of the balance in hex string format
	RawValue string `json:"raw_value" api:"required"`
	// The value of the balance in decimal string format
	Value string                                                                                                            `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalanceJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalanceJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalance]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalanceJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalanceJSON) RawJSON() string {
	return r.raw
}

// The parameters of the transaction that was simulated.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParams struct {
	// The block tag to be sent.
	BlockTag string `json:"block_tag"`
	// The calldata to be sent.
	Calldata EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldata `json:"calldata"`
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
	UserOperationCalldata EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldata `json:"user_operation_calldata"`
	// The value to be sent.
	Value string                                                                                    `json:"value"`
	JSON  evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParams]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsJSON struct {
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

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsJSON) RawJSON() string {
	return r.raw
}

// The calldata to be sent.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldata struct {
	// The function selector of the function called in the transaction
	FunctionSelector string `json:"function_selector" api:"required"`
	// The function declaration of the function called in the transaction
	FunctionDeclaration string `json:"function_declaration"`
	// The function signature of the function called in the transaction
	FunctionSignature string                                                                                            `json:"function_signature"`
	JSON              evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldataJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldataJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldata]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldataJSON struct {
	FunctionSelector    apijson.Field
	FunctionDeclaration apijson.Field
	FunctionSignature   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldataJSON) RawJSON() string {
	return r.raw
}

// The user operation call data to be sent.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldata struct {
	// The function selector of the function called in the transaction
	FunctionSelector string `json:"function_selector" api:"required"`
	// The function declaration of the function called in the transaction
	FunctionDeclaration string `json:"function_declaration"`
	// The function signature of the function called in the transaction
	FunctionSignature string                                                                                                         `json:"function_signature"`
	JSON              evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldataJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldataJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldata]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldataJSON struct {
	FunctionSelector    apijson.Field
	FunctionDeclaration apijson.Field
	FunctionSignature   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldataJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationError struct {
	// A string explaining why the transaction failed
	Description string `json:"description" api:"required"`
	// An error message if the simulation failed.
	Error string `json:"error" api:"required"`
	// A string indicating if the simulation was successful or not.
	Status EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatus `json:"status" api:"required"`
	// Error details if the simulation failed.
	ErrorDetails EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails `json:"error_details"`
	JSON         evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorJSON         `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationError]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorJSON struct {
	Description  apijson.Field
	Error        apijson.Field
	Status       apijson.Field
	ErrorDetails apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationError) implementsEvmTransactionBulkScanResponseSimulation() {
}

// A string indicating if the simulation was successful or not.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatus string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatusError EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatus = "Error"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatus) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatusError:
		return true
	}
	return false
}

// Error details if the simulation failed.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails struct {
	// The type of the model
	Code string `json:"code" api:"required"`
	// The address of the account
	AccountAddress string `json:"account_address"`
	// The address that is invalid
	Address string `json:"address"`
	// This field can have the runtime type of
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset].
	Asset    interface{} `json:"asset"`
	Category string      `json:"category"`
	// The current balance of the account
	CurrentBalance string `json:"current_balance"`
	// The domain name that is unsupported
	DomainName string `json:"domain_name"`
	// The message type that is unsupported
	MessageType string `json:"message_type"`
	// The required balance of the account
	RequiredBalance string                                                                                               `json:"required_balance"`
	JSON            evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsJSON `json:"-"`
	union           EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsJSON struct {
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

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsUnion {
	return r.union
}

// Error details if the simulation failed.
//
// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails struct {
	// The address of the account
	AccountAddress string `json:"account_address" api:"required"`
	// The asset that the account does not have enough balance for
	Asset EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset `json:"asset" api:"required"`
	// The type of the model
	Code     EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCode     `json:"code" api:"required"`
	Category EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCategory `json:"category"`
	// The current balance of the account
	CurrentBalance string `json:"current_balance"`
	// The required balance of the account
	RequiredBalance string                                                                                                                                                     `json:"required_balance"`
	JSON            evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsJSON struct {
	AccountAddress  apijson.Field
	Asset           apijson.Field
	Code            apijson.Field
	Category        apijson.Field
	CurrentBalance  apijson.Field
	RequiredBalance apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails() {
}

// The asset that the account does not have enough balance for
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset struct {
	// This field can have the runtime type of
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetails],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetails],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetails],
	// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetails].
	Details interface{} `json:"details" api:"required"`
	// The type of the model
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType `json:"type" api:"required"`
	// Token Id
	TokenID int64                                                                                                                                                           `json:"token_id"`
	JSON    evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetJSON `json:"-"`
	union   EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetUnion
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetJSON struct {
	Details     apijson.Field
	Type        apijson.Field
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset].
func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset) AsUnion() EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetUnion {
	return r.union
}

// The asset that the account does not have enough balance for
//
// Union satisfied by
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset],
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset]
// or
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset].
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetUnion interface {
	implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset{}),
		},
	)
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset struct {
	// Details
	Details EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetails `json:"details" api:"required"`
	// The type of the model
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetType `json:"type" api:"required"`
	JSON evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetJSON struct {
	Details     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset() {
}

// Details
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetails struct {
	ChainID   int64  `json:"chain_id" api:"required"`
	ChainName string `json:"chain_name" api:"required"`
	Decimals  int64  `json:"decimals" api:"required"`
	LogoURL   string `json:"logo_url" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsType `json:"type" api:"required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                              `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsJSON struct {
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

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsTypeNative EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsType = "NATIVE"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsTypeNative:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetTypeNative EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetType = "NATIVE"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetTypeNative:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset struct {
	// Details
	Details EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetails `json:"details" api:"required"`
	// The type of the model
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetType `json:"type" api:"required"`
	JSON evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetJSON struct {
	Details     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset() {
}

// Details
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset's decimals
	Decimals int64 `json:"decimals" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                             `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsTypeErc20 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsType = "ERC20"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsTypeErc20:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetTypeErc20 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetType = "ERC20"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetTypeErc20:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset struct {
	// Details
	Details EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetails `json:"details" api:"required"`
	// Token Id
	TokenID int64 `json:"token_id" api:"required"`
	// The type of the model
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetType `json:"type" api:"required"`
	JSON evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetJSON struct {
	Details     apijson.Field
	TokenID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset() {
}

// Details
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                              `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsTypeErc1155 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsType = "ERC1155"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsTypeErc1155:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetTypeErc721 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetType = "ERC721"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetTypeErc721:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset struct {
	// Details
	Details EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetails `json:"details" api:"required"`
	// Token Id
	TokenID int64 `json:"token_id" api:"required"`
	// The type of the model
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetType `json:"type" api:"required"`
	JSON evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetJSON struct {
	Details     apijson.Field
	TokenID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset() {
}

// Details
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetails struct {
	// address of the token
	Address string `json:"address" api:"required"`
	// asset type.
	Type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsType `json:"type" api:"required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                               `json:"symbol"`
	JSON   evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsTypeErc1155 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsType = "ERC1155"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsTypeErc1155:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetTypeErc1155 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetType = "ERC1155"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetTypeErc1155:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeNative  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType = "NATIVE"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc20   EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType = "ERC20"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc721  EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType = "ERC721"
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc1155 EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType = "ERC1155"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeNative, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc20, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc721, EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc1155:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCode string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCodeGeneralInsufficientFunds EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCode = "GENERAL_INSUFFICIENT_FUNDS"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCode) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCodeGeneralInsufficientFunds:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCategory string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCategoryRevert EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCategory = "REVERT"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCategory) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCategoryRevert:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails struct {
	// The address that is invalid
	Address string `json:"address" api:"required"`
	// The type of the model
	Code     EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCode     `json:"code" api:"required"`
	Category EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCategory `json:"category"`
	JSON     evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsJSON     `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsJSON struct {
	Address     apijson.Field
	Code        apijson.Field
	Category    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails() {
}

// The type of the model
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCode string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCodeGeneralInvalidAddress EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCode = "GENERAL_INVALID_ADDRESS"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCode) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCodeGeneralInvalidAddress:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCategory string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCategoryInvalidInput EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCategory = "INVALID_INPUT"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCategory) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCategoryInvalidInput:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails struct {
	Category string `json:"category" api:"required"`
	// The error code
	Code string                                                                                                                                    `json:"code" api:"required"`
	JSON evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetailsJSON `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetailsJSON struct {
	Category    apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails() {
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails struct {
	// The type of the model
	Code EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCode `json:"code" api:"required"`
	// The domain name that is unsupported
	DomainName string `json:"domain_name" api:"required"`
	// The message type that is unsupported
	MessageType string                                                                                                                                                         `json:"message_type" api:"required"`
	Category    EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCategory `json:"category"`
	JSON        evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsJSON     `json:"-"`
}

// evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails]
type evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsJSON struct {
	Code        apijson.Field
	DomainName  apijson.Field
	MessageType apijson.Field
	Category    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails) implementsEvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails() {
}

// The type of the model
type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCode string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCodeUnsupportedEip712Message EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCode = "UNSUPPORTED_EIP712_MESSAGE"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCode) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCodeUnsupportedEip712Message:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCategory string

const (
	EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCategoryInvalidInput EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCategory = "INVALID_INPUT"
)

func (r EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCategory) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCategoryInvalidInput:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type EvmTransactionBulkScanResponseSimulationStatus string

const (
	EvmTransactionBulkScanResponseSimulationStatusSuccess EvmTransactionBulkScanResponseSimulationStatus = "Success"
	EvmTransactionBulkScanResponseSimulationStatusError   EvmTransactionBulkScanResponseSimulationStatus = "Error"
)

func (r EvmTransactionBulkScanResponseSimulationStatus) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseSimulationStatusSuccess, EvmTransactionBulkScanResponseSimulationStatusError:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseUserOperationGasEstimation struct {
	Status                           EvmTransactionBulkScanResponseUserOperationGasEstimationStatus `json:"status" api:"required"`
	CallGasEstimate                  string                                                         `json:"call_gas_estimate"`
	Error                            string                                                         `json:"error"`
	PaymasterVerificationGasEstimate string                                                         `json:"paymaster_verification_gas_estimate"`
	PreVerificationGasEstimate       string                                                         `json:"pre_verification_gas_estimate"`
	VerificationGasEstimate          string                                                         `json:"verification_gas_estimate"`
	JSON                             evmTransactionBulkScanResponseUserOperationGasEstimationJSON   `json:"-"`
	union                            EvmTransactionBulkScanResponseUserOperationGasEstimationUnion
}

// evmTransactionBulkScanResponseUserOperationGasEstimationJSON contains the JSON
// metadata for the struct
// [EvmTransactionBulkScanResponseUserOperationGasEstimation]
type evmTransactionBulkScanResponseUserOperationGasEstimationJSON struct {
	Status                           apijson.Field
	CallGasEstimate                  apijson.Field
	Error                            apijson.Field
	PaymasterVerificationGasEstimate apijson.Field
	PreVerificationGasEstimate       apijson.Field
	VerificationGasEstimate          apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r evmTransactionBulkScanResponseUserOperationGasEstimationJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseUserOperationGasEstimation) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseUserOperationGasEstimation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionBulkScanResponseUserOperationGasEstimationUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [UserOperationV6GasEstimation],
// [UserOperationV7GasEstimation],
// [EvmTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError].
func (r EvmTransactionBulkScanResponseUserOperationGasEstimation) AsUnion() EvmTransactionBulkScanResponseUserOperationGasEstimationUnion {
	return r.union
}

// Union satisfied by [UserOperationV6GasEstimation],
// [UserOperationV7GasEstimation] or
// [EvmTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError].
type EvmTransactionBulkScanResponseUserOperationGasEstimationUnion interface {
	implementsEvmTransactionBulkScanResponseUserOperationGasEstimation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseUserOperationGasEstimationUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError{}),
		},
	)
}

type EvmTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError struct {
	Error  string                                                                                                          `json:"error" api:"required"`
	Status EvmTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus `json:"status" api:"required"`
	JSON   evmTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON   `json:"-"`
}

// evmTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError]
type evmTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError) implementsEvmTransactionBulkScanResponseUserOperationGasEstimation() {
}

type EvmTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus string

const (
	EvmTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatusError EvmTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus = "Error"
)

func (r EvmTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatusError:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseUserOperationGasEstimationStatus string

const (
	EvmTransactionBulkScanResponseUserOperationGasEstimationStatusSuccess EvmTransactionBulkScanResponseUserOperationGasEstimationStatus = "Success"
	EvmTransactionBulkScanResponseUserOperationGasEstimationStatusError   EvmTransactionBulkScanResponseUserOperationGasEstimationStatus = "Error"
)

func (r EvmTransactionBulkScanResponseUserOperationGasEstimationStatus) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseUserOperationGasEstimationStatusSuccess, EvmTransactionBulkScanResponseUserOperationGasEstimationStatusError:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseValidation struct {
	// This field can have the runtime type of
	// [[]EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeature],
	// [[]EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeature].
	Features interface{} `json:"features" api:"required"`
	// Result type returned when validation succeeds.
	ResultType EvmTransactionBulkScanResponseValidationResultType `json:"result_type" api:"required"`
	// A string indicating if the simulation was successful or not.
	Status EvmTransactionBulkScanResponseValidationStatus `json:"status" api:"required"`
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
	Reason string                                       `json:"reason"`
	JSON   evmTransactionBulkScanResponseValidationJSON `json:"-"`
	union  EvmTransactionBulkScanResponseValidationUnion
}

// evmTransactionBulkScanResponseValidationJSON contains the JSON metadata for the
// struct [EvmTransactionBulkScanResponseValidation]
type evmTransactionBulkScanResponseValidationJSON struct {
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

func (r evmTransactionBulkScanResponseValidationJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionBulkScanResponseValidation) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionBulkScanResponseValidation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EvmTransactionBulkScanResponseValidationUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidation],
// [EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationError].
func (r EvmTransactionBulkScanResponseValidation) AsUnion() EvmTransactionBulkScanResponseValidationUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidation]
// or
// [EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationError].
type EvmTransactionBulkScanResponseValidationUnion interface {
	implementsEvmTransactionBulkScanResponseValidation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionBulkScanResponseValidationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationError{}),
		},
	)
}

type EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidation struct {
	// A list of features about this transaction explaining the validation.
	Features []EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeature `json:"features" api:"required"`
	// Result type returned when validation succeeds.
	ResultType EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationResultType `json:"result_type" api:"required"`
	// A string indicating if the simulation was successful or not.
	Status EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationStatus `json:"status" api:"required"`
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification"`
	// A textual description that can be presented to the user about what this
	// transaction is doing.
	Description string `json:"description"`
	// A textual description about the reasons the transaction was flagged with
	// result_type.
	Reason string                                                                              `json:"reason"`
	JSON   evmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationJSON `json:"-"`
}

// evmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidation]
type evmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationJSON struct {
	Features       apijson.Field
	ResultType     apijson.Field
	Status         apijson.Field
	Classification apijson.Field
	Description    apijson.Field
	Reason         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidation) implementsEvmTransactionBulkScanResponseValidation() {
}

type EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeature struct {
	// Textual description
	Description string `json:"description" api:"required"`
	// Feature name
	FeatureID string `json:"feature_id" api:"required"`
	// Security result of a transaction scan feature.
	Type EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType `json:"type" api:"required"`
	// Address the feature refers to
	Address string `json:"address"`
	// Metadata related to the feature
	Metadata interface{}                                                                                `json:"metadata"`
	JSON     evmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeatureJSON `json:"-"`
}

// evmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeatureJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeature]
type evmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeatureJSON struct {
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	Address     apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeatureJSON) RawJSON() string {
	return r.raw
}

// Security result of a transaction scan feature.
type EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType string

const (
	EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeMalicious EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType = "Malicious"
	EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeWarning   EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType = "Warning"
	EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeBenign    EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType = "Benign"
	EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeInfo      EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType = "Info"
)

func (r EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeMalicious, EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeWarning, EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeBenign, EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeInfo:
		return true
	}
	return false
}

// Result type returned when validation succeeds.
type EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationResultType string

const (
	EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeBenign    EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationResultType = "Benign"
	EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeWarning   EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationResultType = "Warning"
	EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeMalicious EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationResultType = "Malicious"
)

func (r EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationResultType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeBenign, EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeWarning, EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeMalicious:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationStatus string

const (
	EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationStatusSuccess EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationStatus = "Success"
)

func (r EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationStatus) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationStatusSuccess:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationError struct {
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassification `json:"classification" api:"required"`
	// A textual description that can be presented to the user about what this
	// transaction is doing.
	Description EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescription `json:"description" api:"required"`
	// An error message if the validation failed.
	Error string `json:"error" api:"required"`
	// A list of features about this transaction explaining the validation.
	Features []EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeature `json:"features" api:"required"`
	// A textual description about the reasons the transaction was flagged with
	// result_type.
	Reason EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorReason `json:"reason" api:"required"`
	// A string indicating if the transaction is safe to sign or not.
	ResultType EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultType `json:"result_type" api:"required"`
	// A string indicating if the simulation was successful or not.
	Status EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatus `json:"status" api:"required"`
	JSON   evmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorJSON   `json:"-"`
}

// evmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationError]
type evmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorJSON struct {
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

func (r *EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationError) implementsEvmTransactionBulkScanResponseValidation() {
}

// A textual classification that can be presented to the user explaining the
// reason.
type EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassification string

const (
	EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassificationEmpty EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassification = ""
)

func (r EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassification) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassificationEmpty:
		return true
	}
	return false
}

// A textual description that can be presented to the user about what this
// transaction is doing.
type EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescription string

const (
	EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescriptionEmpty EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescription = ""
)

func (r EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescription) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescriptionEmpty:
		return true
	}
	return false
}

type EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeature struct {
	// Textual description
	Description string `json:"description" api:"required"`
	// Feature name
	FeatureID string `json:"feature_id" api:"required"`
	// Security result of a transaction scan feature.
	Type EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType `json:"type" api:"required"`
	// Address the feature refers to
	Address string `json:"address"`
	// Metadata related to the feature
	Metadata interface{}                                                                                     `json:"metadata"`
	JSON     evmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeatureJSON `json:"-"`
}

// evmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeatureJSON
// contains the JSON metadata for the struct
// [EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeature]
type evmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeatureJSON struct {
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	Address     apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeatureJSON) RawJSON() string {
	return r.raw
}

// Security result of a transaction scan feature.
type EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType string

const (
	EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeMalicious EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType = "Malicious"
	EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeWarning   EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType = "Warning"
	EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeBenign    EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType = "Benign"
	EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeInfo      EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType = "Info"
)

func (r EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeMalicious, EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeWarning, EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeBenign, EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeInfo:
		return true
	}
	return false
}

// A textual description about the reasons the transaction was flagged with
// result_type.
type EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorReason string

const (
	EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorReasonEmpty EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorReason = ""
)

func (r EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorReason) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorReasonEmpty:
		return true
	}
	return false
}

// A string indicating if the transaction is safe to sign or not.
type EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultType string

const (
	EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultTypeError EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultType = "Error"
)

func (r EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultTypeError:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatus string

const (
	EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatusSuccess EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatus = "Success"
)

func (r EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatus) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatusSuccess:
		return true
	}
	return false
}

// Result type returned when validation succeeds.
type EvmTransactionBulkScanResponseValidationResultType string

const (
	EvmTransactionBulkScanResponseValidationResultTypeBenign    EvmTransactionBulkScanResponseValidationResultType = "Benign"
	EvmTransactionBulkScanResponseValidationResultTypeWarning   EvmTransactionBulkScanResponseValidationResultType = "Warning"
	EvmTransactionBulkScanResponseValidationResultTypeMalicious EvmTransactionBulkScanResponseValidationResultType = "Malicious"
	EvmTransactionBulkScanResponseValidationResultTypeError     EvmTransactionBulkScanResponseValidationResultType = "Error"
)

func (r EvmTransactionBulkScanResponseValidationResultType) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseValidationResultTypeBenign, EvmTransactionBulkScanResponseValidationResultTypeWarning, EvmTransactionBulkScanResponseValidationResultTypeMalicious, EvmTransactionBulkScanResponseValidationResultTypeError:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type EvmTransactionBulkScanResponseValidationStatus string

const (
	EvmTransactionBulkScanResponseValidationStatusSuccess EvmTransactionBulkScanResponseValidationStatus = "Success"
)

func (r EvmTransactionBulkScanResponseValidationStatus) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanResponseValidationStatusSuccess:
		return true
	}
	return false
}

type EvmTransactionBulkScanParams struct {
	// The chain name or chain ID
	Chain param.Field[TransactionScanSupportedChain] `json:"chain" api:"required"`
	// Transaction bulk parameters
	Data param.Field[[]EvmTransactionBulkScanParamsData] `json:"data" api:"required"`
	// Additional context for the scan (e.g., dapp URL/domain, integration source).
	// Used to enrich results and reduce false positives/negatives.
	Metadata param.Field[EvmTransactionBulkScanParamsMetadataUnion] `json:"metadata" api:"required"`
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
	From param.Field[string] `json:"from" api:"required"`
	// The authorization list
	AuthorizationList param.Field[[]AuthorizationParam] `json:"authorization_list"`
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

// Additional context for the scan (e.g., dapp URL/domain, integration source).
// Used to enrich results and reduce false positives/negatives.
type EvmTransactionBulkScanParamsMetadata struct {
	// The full URL of the DApp or website that initiated the transaction, for
	// cross-reference. Must use the https or http scheme and contain a valid hostname.
	// Cannot contain JSON, braces, or other embedded data structures.
	Domain param.Field[string] `json:"domain"`
	// Indicates that the transaction was not initiated by a dapp.
	NonDapp param.Field[bool] `json:"non_dapp"`
}

func (r EvmTransactionBulkScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmTransactionBulkScanParamsMetadata) implementsEvmTransactionBulkScanParamsMetadataUnion() {}

// Additional context for the scan (e.g., dapp URL/domain, integration source).
// Used to enrich results and reduce false positives/negatives.
//
// Satisfied by
// [EvmTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataNonDapp],
// [EvmTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataDapp],
// [EvmTransactionBulkScanParamsMetadata].
type EvmTransactionBulkScanParamsMetadataUnion interface {
	implementsEvmTransactionBulkScanParamsMetadataUnion()
}

type EvmTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataNonDapp struct {
	// Indicates that the transaction was not initiated by a dapp.
	NonDapp param.Field[EvmTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataNonDappNonDapp] `json:"non_dapp"`
}

func (r EvmTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataNonDapp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataNonDapp) implementsEvmTransactionBulkScanParamsMetadataUnion() {
}

// Indicates that the transaction was not initiated by a dapp.
type EvmTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataNonDappNonDapp bool

const (
	EvmTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataNonDappNonDappTrue EvmTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataNonDappNonDapp = true
)

func (r EvmTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataNonDappNonDapp) IsKnown() bool {
	switch r {
	case EvmTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataNonDappNonDappTrue:
		return true
	}
	return false
}

type EvmTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataDapp struct {
	// The full URL of the DApp or website that initiated the transaction, for
	// cross-reference. Must use the https or http scheme and contain a valid hostname.
	// Cannot contain JSON, braces, or other embedded data structures.
	Domain param.Field[string] `json:"domain" api:"required"`
	// Indicates that the transaction was not initiated by a dapp. Use false when the
	// transaction is from a dapp.
	NonDapp param.Field[bool] `json:"non_dapp"`
}

func (r EvmTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataDapp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataDapp) implementsEvmTransactionBulkScanParamsMetadataUnion() {
}

// The relative block for the block validation. Can be "latest" or a block number.
//
// Satisfied by [shared.UnionInt], [shared.UnionString].
type EvmTransactionBulkScanParamsBlockUnion interface {
	ImplementsEvmTransactionBulkScanParamsBlockUnion()
}

// Response sections to include (e.g., validation, simulation, gas estimation,
// events).
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
