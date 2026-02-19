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

// EvmPostTransactionBulkService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvmPostTransactionBulkService] method instead.
type EvmPostTransactionBulkService struct {
	Options []option.RequestOption
}

// NewEvmPostTransactionBulkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEvmPostTransactionBulkService(opts ...option.RequestOption) (r *EvmPostTransactionBulkService) {
	r = &EvmPostTransactionBulkService{}
	r.Options = opts
	return
}

// Scan transactions that were already executed on chain, returns validation with
// features indicating address poisoning entites and malicious operators.
func (r *EvmPostTransactionBulkService) Scan(ctx context.Context, body EvmPostTransactionBulkScanParams, opts ...option.RequestOption) (res *[]EvmPostTransactionBulkScanResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/evm/post-transaction-bulk/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EvmPostTransactionBulkScanResponse struct {
	Block                      string                                                       `json:"block,required"`
	Chain                      string                                                       `json:"chain,required"`
	AccountAddress             string                                                       `json:"account_address"`
	Events                     []EvmPostTransactionBulkScanResponseEvent                    `json:"events"`
	Features                   interface{}                                                  `json:"features"`
	GasEstimation              EvmPostTransactionBulkScanResponseGasEstimation              `json:"gas_estimation"`
	Simulation                 EvmPostTransactionBulkScanResponseSimulation                 `json:"simulation"`
	UserOperationGasEstimation EvmPostTransactionBulkScanResponseUserOperationGasEstimation `json:"user_operation_gas_estimation"`
	Validation                 EvmPostTransactionBulkScanResponseValidation                 `json:"validation"`
	JSON                       evmPostTransactionBulkScanResponseJSON                       `json:"-"`
}

// evmPostTransactionBulkScanResponseJSON contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponse]
type evmPostTransactionBulkScanResponseJSON struct {
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

func (r *EvmPostTransactionBulkScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseEvent struct {
	Data           string                                          `json:"data,required"`
	EmitterAddress string                                          `json:"emitter_address,required"`
	Topics         []string                                        `json:"topics,required"`
	EmitterName    string                                          `json:"emitter_name"`
	Name           string                                          `json:"name"`
	Params         []EvmPostTransactionBulkScanResponseEventsParam `json:"params"`
	JSON           evmPostTransactionBulkScanResponseEventJSON     `json:"-"`
}

// evmPostTransactionBulkScanResponseEventJSON contains the JSON metadata for the
// struct [EvmPostTransactionBulkScanResponseEvent]
type evmPostTransactionBulkScanResponseEventJSON struct {
	Data           apijson.Field
	EmitterAddress apijson.Field
	Topics         apijson.Field
	EmitterName    apijson.Field
	Name           apijson.Field
	Params         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseEventJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseEventsParam struct {
	Type         string                                                   `json:"type,required"`
	Value        EvmPostTransactionBulkScanResponseEventsParamsValueUnion `json:"value,required"`
	InternalType string                                                   `json:"internalType"`
	Name         string                                                   `json:"name"`
	JSON         evmPostTransactionBulkScanResponseEventsParamJSON        `json:"-"`
}

// evmPostTransactionBulkScanResponseEventsParamJSON contains the JSON metadata for
// the struct [EvmPostTransactionBulkScanResponseEventsParam]
type evmPostTransactionBulkScanResponseEventsParamJSON struct {
	Type         apijson.Field
	Value        apijson.Field
	InternalType apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseEventsParam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseEventsParamJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or
// [EvmPostTransactionBulkScanResponseEventsParamsValueArray].
type EvmPostTransactionBulkScanResponseEventsParamsValueUnion interface {
	ImplementsEvmPostTransactionBulkScanResponseEventsParamsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseEventsParamsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseEventsParamsValueArray{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseEventsParamsValueArray []interface{}

func (r EvmPostTransactionBulkScanResponseEventsParamsValueArray) ImplementsEvmPostTransactionBulkScanResponseEventsParamsValueUnion() {
}

type EvmPostTransactionBulkScanResponseGasEstimation struct {
	Status   EvmPostTransactionBulkScanResponseGasEstimationStatus `json:"status,required"`
	Error    string                                                `json:"error"`
	Estimate string                                                `json:"estimate"`
	Used     string                                                `json:"used"`
	JSON     evmPostTransactionBulkScanResponseGasEstimationJSON   `json:"-"`
	union    EvmPostTransactionBulkScanResponseGasEstimationUnion
}

// evmPostTransactionBulkScanResponseGasEstimationJSON contains the JSON metadata
// for the struct [EvmPostTransactionBulkScanResponseGasEstimation]
type evmPostTransactionBulkScanResponseGasEstimationJSON struct {
	Status      apijson.Field
	Error       apijson.Field
	Estimate    apijson.Field
	Used        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseGasEstimationJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseGasEstimation) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseGasEstimation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EvmPostTransactionBulkScanResponseGasEstimationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation],
// [EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError].
func (r EvmPostTransactionBulkScanResponseGasEstimation) AsUnion() EvmPostTransactionBulkScanResponseGasEstimationUnion {
	return r.union
}

// Union satisfied by
// [EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation]
// or
// [EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError].
type EvmPostTransactionBulkScanResponseGasEstimationUnion interface {
	implementsEvmPostTransactionBulkScanResponseGasEstimation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseGasEstimationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation struct {
	Estimate string                                                                                            `json:"estimate,required"`
	Status   EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatus `json:"status,required"`
	Used     string                                                                                            `json:"used,required"`
	JSON     evmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationJSON   `json:"-"`
}

// evmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation]
type evmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationJSON struct {
	Estimate    apijson.Field
	Status      apijson.Field
	Used        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation) implementsEvmPostTransactionBulkScanResponseGasEstimation() {
}

type EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatus string

const (
	EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatusSuccess EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatus = "Success"
)

func (r EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatus) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatusSuccess:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError struct {
	Error  string                                                                                                 `json:"error,required"`
	Status EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus `json:"status,required"`
	JSON   evmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON   `json:"-"`
}

// evmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError]
type evmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError) implementsEvmPostTransactionBulkScanResponseGasEstimation() {
}

type EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus string

const (
	EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatusError EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus = "Error"
)

func (r EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatusError:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseGasEstimationStatus string

const (
	EvmPostTransactionBulkScanResponseGasEstimationStatusSuccess EvmPostTransactionBulkScanResponseGasEstimationStatus = "Success"
	EvmPostTransactionBulkScanResponseGasEstimationStatusError   EvmPostTransactionBulkScanResponseGasEstimationStatus = "Error"
)

func (r EvmPostTransactionBulkScanResponseGasEstimationStatus) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseGasEstimationStatusSuccess, EvmPostTransactionBulkScanResponseGasEstimationStatusError:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulation struct {
	// A string indicating if the simulation was successful or not.
	Status EvmPostTransactionBulkScanResponseSimulationStatus `json:"status,required"`
	// This field can have the runtime type of
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummary].
	AccountSummary interface{} `json:"account_summary"`
	// This field can have the runtime type of
	// [map[string]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetail].
	AddressDetails interface{} `json:"address_details"`
	// This field can have the runtime type of
	// [map[string][]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff].
	AssetsDiffs interface{} `json:"assets_diffs"`
	// This field can have the runtime type of
	// [map[string][]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement].
	ContractManagement interface{} `json:"contract_management"`
	// A string explaining why the transaction failed
	Description string `json:"description"`
	// An error message if the simulation failed.
	Error string `json:"error"`
	// This field can have the runtime type of
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails].
	ErrorDetails interface{} `json:"error_details"`
	// This field can have the runtime type of
	// [map[string][]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure].
	Exposures interface{} `json:"exposures"`
	// This field can have the runtime type of
	// [[]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalance].
	MissingBalances interface{} `json:"missing_balances"`
	// This field can have the runtime type of
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParams].
	Params interface{} `json:"params"`
	// This field can have the runtime type of
	// [map[string][]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKey].
	SessionKey interface{} `json:"session_key"`
	// The number of times the simulation ran until success
	SimulationRunCount int64 `json:"simulation_run_count"`
	// This field can have the runtime type of
	// [map[string]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiff].
	TotalUsdDiff interface{} `json:"total_usd_diff"`
	// This field can have the runtime type of [map[string]map[string]string].
	TotalUsdExposure interface{} `json:"total_usd_exposure"`
	// This field can have the runtime type of
	// [[]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions].
	TransactionActions interface{}                                      `json:"transaction_actions"`
	JSON               evmPostTransactionBulkScanResponseSimulationJSON `json:"-"`
	union              EvmPostTransactionBulkScanResponseSimulationUnion
}

// evmPostTransactionBulkScanResponseSimulationJSON contains the JSON metadata for
// the struct [EvmPostTransactionBulkScanResponseSimulation]
type evmPostTransactionBulkScanResponseSimulationJSON struct {
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

func (r evmPostTransactionBulkScanResponseSimulationJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulation) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EvmPostTransactionBulkScanResponseSimulationUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulation],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationError].
func (r EvmPostTransactionBulkScanResponseSimulation) AsUnion() EvmPostTransactionBulkScanResponseSimulationUnion {
	return r.union
}

// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulation]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationError].
type EvmPostTransactionBulkScanResponseSimulationUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationError{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulation struct {
	// Account summary for the account address. account address is determined implicit
	// by the `from` field in the transaction request, or explicit by the
	// account_address field in the request.
	AccountSummary EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummary `json:"account_summary,required"`
	// a dictionary including additional information about each relevant address in the
	// transaction.
	AddressDetails map[string]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetail `json:"address_details,required"`
	// dictionary describes the assets differences as a result of this transaction for
	// every involved address
	AssetsDiffs map[string][]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff `json:"assets_diffs,required"`
	// dictionary describes the exposure differences as a result of this transaction
	// for every involved address (as a result of any approval / setApproval / permit
	// function)
	Exposures map[string][]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure `json:"exposures,required"`
	// Session keys created in this transaction per address
	SessionKey map[string][]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKey `json:"session_key,required"`
	// A string indicating if the simulation was successful or not.
	Status EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationStatus `json:"status,required"`
	// dictionary represents the usd value each address gained / lost during this
	// transaction
	TotalUsdDiff map[string]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiff `json:"total_usd_diff,required"`
	// a dictionary representing the usd value each address is exposed to, split by
	// spenders
	TotalUsdExposure map[string]map[string]string `json:"total_usd_exposure,required"`
	// Describes the nature of the transaction and what happened as part of it
	TransactionActions []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions `json:"transaction_actions,required"`
	// Describes the state differences as a result of this transaction for every
	// involved address
	ContractManagement map[string][]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement `json:"contract_management"`
	// Missing balances in the transaction
	MissingBalances []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalance `json:"missing_balances"`
	// The parameters of the transaction that was simulated.
	Params EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParams `json:"params"`
	// The number of times the simulation ran until success
	SimulationRunCount int64                                                                                   `json:"simulation_run_count"`
	JSON               evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulation]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationJSON struct {
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

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulation) implementsEvmPostTransactionBulkScanResponseSimulation() {
}

// Account summary for the account address. account address is determined implicit
// by the `from` field in the transaction request, or explicit by the
// account_address field in the request.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummary struct {
	// All assets diffs related to the account address
	AssetsDiffs []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff `json:"assets_diffs,required"`
	// All assets exposures related to the account address
	Exposures []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure `json:"exposures,required"`
	// Total usd diff related to the account address
	TotalUsdDiff EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiff `json:"total_usd_diff,required"`
	// Total usd exposure related to the account address
	TotalUsdExposure map[string]string `json:"total_usd_exposure,required"`
	// All assets traces related to the account address
	Traces []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace `json:"traces,required"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryJSON    `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummary]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryJSON struct {
	AssetsDiffs      apijson.Field
	Exposures        apijson.Field
	TotalUsdDiff     apijson.Field
	TotalUsdExposure apijson.Field
	Traces           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff struct {
	// This field can have the runtime type of
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAsset].
	Asset interface{} `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType `json:"asset_type,required"`
	// This field can have the runtime type of
	// [[]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffIn],
	// [[]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffIn],
	// [[]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffIn],
	// [[]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffIn].
	In interface{} `json:"in,required"`
	// This field can have the runtime type of
	// [[]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOut],
	// [[]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOut],
	// [[]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOut],
	// [[]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOut].
	Out interface{} `json:"out,required"`
	// This field can have the runtime type of
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChanges],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChanges],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChanges],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChanges].
	BalanceChanges interface{}                                                                                                     `json:"balance_changes"`
	JSON           evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffJSON `json:"-"`
	union          EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff struct {
	// description of the asset for the current diff
	Asset EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOut `json:"out,required"`
	// shows the balance before making the transaction and after
	BalanceChanges EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChanges `json:"balance_changes"`
	JSON           evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffJSON           `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff() {
}

// description of the asset for the current diff
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType `json:"type,required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                    `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetJSON `json:"-"`
	union  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                           `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                            `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetTypeErc20  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType = "ERC20"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetTypeErc20, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffIn struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                 `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffInJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffInJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffIn]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffInJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOut struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                  `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOutJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOutJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOut]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOutJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOutJSON) RawJSON() string {
	return r.raw
}

// shows the balance before making the transaction and after
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChanges struct {
	// balance of the account after making the transaction
	After EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfter `json:"after,required"`
	// balance of the account before making the transaction
	Before EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBefore `json:"before,required"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesJSON   `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChanges]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChanges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesJSON) RawJSON() string {
	return r.raw
}

// balance of the account after making the transaction
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfter struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                  `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfterJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfterJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfter]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfterJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfterJSON) RawJSON() string {
	return r.raw
}

// balance of the account before making the transaction
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBefore struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                   `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBefore]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff struct {
	// description of the asset for the current diff
	Asset EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOut `json:"out,required"`
	// shows the balance before making the transaction and after
	BalanceChanges EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChanges `json:"balance_changes"`
	JSON           evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffJSON           `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff() {
}

// description of the asset for the current diff
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                     `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetJSON `json:"-"`
	union  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                             `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                             `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetTypeErc721 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType = "ERC721"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetTypeErc721, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffIn struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                  `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffInJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffInJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffIn]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffInJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOut struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                   `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOutJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOutJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOut]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOutJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOutJSON) RawJSON() string {
	return r.raw
}

// shows the balance before making the transaction and after
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChanges struct {
	// balance of the account after making the transaction
	After EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfter `json:"after,required"`
	// balance of the account before making the transaction
	Before EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBefore `json:"before,required"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesJSON   `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChanges]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChanges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesJSON) RawJSON() string {
	return r.raw
}

// balance of the account after making the transaction
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfter struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                                   `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfterJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfterJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfter]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfterJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfterJSON) RawJSON() string {
	return r.raw
}

// balance of the account before making the transaction
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBefore struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                                    `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBefore]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff struct {
	// description of the asset for the current diff
	Asset EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOut `json:"out,required"`
	// shows the balance before making the transaction and after
	BalanceChanges EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChanges `json:"balance_changes"`
	JSON           evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffJSON           `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff() {
}

// description of the asset for the current diff
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                      `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetJSON `json:"-"`
	union  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                               `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                              `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetTypeErc1155 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType = "ERC1155"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetTypeNonerc  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetTypeErc1155, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffIn struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                   `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffInJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffInJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffIn]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffInJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOut struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                    `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOutJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOutJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOut]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOutJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOutJSON) RawJSON() string {
	return r.raw
}

// shows the balance before making the transaction and after
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChanges struct {
	// balance of the account after making the transaction
	After EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfter `json:"after,required"`
	// balance of the account before making the transaction
	Before EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBefore `json:"before,required"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesJSON   `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChanges]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChanges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesJSON) RawJSON() string {
	return r.raw
}

// balance of the account after making the transaction
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfter struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                                    `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfterJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfterJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfter]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfterJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfterJSON) RawJSON() string {
	return r.raw
}

// balance of the account before making the transaction
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBefore struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                                     `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBefore]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff struct {
	// description of the asset for the current diff
	Asset EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOut `json:"out,required"`
	// shows the balance before making the transaction and after
	BalanceChanges EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChanges `json:"balance_changes"`
	JSON           evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffJSON           `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff() {
}

// description of the asset for the current diff
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAsset struct {
	ChainID   int64  `json:"chain_id,required"`
	ChainName string `json:"chain_name,required"`
	Decimals  int64  `json:"decimals,required"`
	LogoURL   string `json:"logo_url,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetType `json:"type,required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                     `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAsset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetJSON struct {
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

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetTypeNative EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetType = "NATIVE"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetTypeNative:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffIn struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                  `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffInJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffInJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffIn]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffInJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOut struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                   `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOutJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOutJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOut]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOutJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOutJSON) RawJSON() string {
	return r.raw
}

// shows the balance before making the transaction and after
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChanges struct {
	// balance of the account after making the transaction
	After EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfter `json:"after,required"`
	// balance of the account before making the transaction
	Before EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBefore `json:"before,required"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesJSON   `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChanges]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChanges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesJSON) RawJSON() string {
	return r.raw
}

// balance of the account after making the transaction
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfter struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                   `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfterJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfterJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfter]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfterJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfterJSON) RawJSON() string {
	return r.raw
}

// balance of the account before making the transaction
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBefore struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                    `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBeforeJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBeforeJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBefore]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBeforeJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBeforeJSON) RawJSON() string {
	return r.raw
}

// type of the asset for the current diff
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc20   EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType = "ERC20"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc721  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType = "ERC721"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc1155 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType = "ERC1155"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeNative  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType = "NATIVE"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc20, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc721, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc1155, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeNative:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure struct {
	// This field can have the runtime type of
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset].
	Asset interface{} `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType `json:"asset_type,required"`
	// This field can have the runtime type of
	// [map[string]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpender],
	// [map[string]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpender],
	// [map[string]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpender].
	Spenders interface{}                                                                                                   `json:"spenders,required"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposureJSON `json:"-"`
	union    EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposureJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresUnion {
	return r.union
}

// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType `json:"asset_type,required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpender `json:"spenders,required"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureJSON               `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure() {
}

// description of the asset for the current diff
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType `json:"type,required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                    `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetJSON `json:"-"`
	union  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                           `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                            `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetTypeErc20  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType = "ERC20"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetTypeErc20, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpender struct {
	// the amount that was asked in the approval request for this spender from the
	// current address and asset
	Approval string                                                                                                                                                             `json:"approval,required"`
	Exposure []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure `json:"exposure,required"`
	// the usd price of the approval amount
	ApprovalUsdPrice string `json:"approval_usd_price"`
	// the expiration time of the permit2 protocol
	Expiration time.Time `json:"expiration" format:"date-time"`
	// user friendly description of the approval
	Summary string                                                                                                                                                      `json:"summary"`
	JSON    evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpender]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON struct {
	Approval         apijson.Field
	Exposure         apijson.Field
	ApprovalUsdPrice apijson.Field
	Expiration       apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure struct {
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
	Value string                                                                                                                                                               `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON `json:"-"`
	union EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON struct {
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

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                                            `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                                           `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                          `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                           `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType `json:"asset_type,required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpender `json:"spenders,required"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureJSON               `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure() {
}

// description of the asset for the current diff
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                     `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetJSON `json:"-"`
	union  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                             `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                             `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetTypeErc721 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType = "ERC721"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetTypeErc721, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpender struct {
	Exposure []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure `json:"exposure,required"`
	// boolean indicates whether an is_approved_for_all function was used (missing in
	// case of ERC20 / ERC1155)
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// user friendly description of the approval
	Summary string                                                                                                                                                       `json:"summary"`
	JSON    evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpender]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure struct {
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
	Value string                                                                                                                                                                `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON `json:"-"`
	union EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON struct {
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

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                                             `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                                            `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                           `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                            `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType `json:"asset_type,required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpender `json:"spenders,required"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureJSON               `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure() {
}

// description of the asset for the current diff
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                      `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON `json:"-"`
	union  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                               `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                              `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeErc1155 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType = "ERC1155"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeNonerc  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeErc1155, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpender struct {
	Exposure []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure `json:"exposure,required"`
	// boolean indicates whether an is_approved_for_all function was used (missing in
	// case of ERC20 / ERC1155)
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// user friendly description of the approval
	Summary string                                                                                                                                                        `json:"summary"`
	JSON    evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpender]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure struct {
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
	Value string                                                                                                                                                                 `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON `json:"-"`
	union EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON struct {
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

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                                              `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                                             `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                            `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                             `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

// type of the asset for the current diff
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc20   EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType = "ERC20"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc721  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType = "ERC721"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc1155 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType = "ERC1155"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc20, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc721, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc1155:
		return true
	}
	return false
}

// Total usd diff related to the account address
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiff struct {
	In    string                                                                                                            `json:"in,required"`
	Out   string                                                                                                            `json:"out,required"`
	Total string                                                                                                            `json:"total,required"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace struct {
	// This field can have the runtime type of
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAsset],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset].
	Asset interface{} `json:"asset,required"`
	// type of the trace
	TraceType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceType `json:"trace_type,required"`
	// The type of the model
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType `json:"type,required"`
	// This field can have the runtime type of
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiff],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiff],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiff],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiff].
	Diff interface{} `json:"diff"`
	// This field can have the runtime type of
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposed],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposed].
	Exposed interface{} `json:"exposed"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address"`
	// This field can have the runtime type of
	// [[]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabels],
	// [[]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabels],
	// [[]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabels],
	// [[]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabels].
	Labels interface{} `json:"labels"`
	// The owner of the assets
	Owner string `json:"owner"`
	// The spender of the assets
	Spender string `json:"spender"`
	// The address where the assets are moved to
	ToAddress string                                                                                                     `json:"to_address"`
	JSON      evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTraceJSON `json:"-"`
	union     EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTraceJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTraceJSON struct {
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

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTraceJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesUnion {
	return r.union
}

// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace struct {
	// Description of the asset in the trace
	Asset EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset `json:"asset,required"`
	// The difference in value for the asset in the trace
	Diff EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiff `json:"diff,required"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address,required"`
	// The address where the assets are moved to
	ToAddress string `json:"to_address,required"`
	// type of the trace
	TraceType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceType `json:"type,required"`
	// List of labels that describe the trace
	Labels []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabels `json:"labels"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceJSON     `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceJSON struct {
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

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetType `json:"type,required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                            `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetJSON `json:"-"`
	union  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                   `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                    `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetTypeErc20  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetType = "ERC20"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetTypeErc20, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetTypeNonerc:
		return true
	}
	return false
}

// The difference in value for the asset in the trace
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                           `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiffJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceTypeAssetTrace EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceType = "AssetTrace"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceTypeAssetTrace:
		return true
	}
	return false
}

// The type of the model
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTypeErc20AssetTrace EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceType = "ERC20AssetTrace"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTypeErc20AssetTrace:
		return true
	}
	return false
}

// Label assigned to a trace entry.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabels string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabelsGasFee EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabels = "GAS_FEE"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabels) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabelsGasFee:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace struct {
	// Description of the asset in the trace
	Asset EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset `json:"asset,required"`
	// The difference in value for the asset in the trace
	Diff EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiff `json:"diff,required"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address,required"`
	// The address where the assets are moved to
	ToAddress string `json:"to_address,required"`
	// type of the trace
	TraceType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceType `json:"type,required"`
	// List of labels that describe the trace
	Labels []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabels `json:"labels"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceJSON     `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceJSON struct {
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

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                             `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetJSON `json:"-"`
	union  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                     `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                     `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetTypeErc721 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetType = "ERC721"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetTypeErc721, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetTypeNonerc:
		return true
	}
	return false
}

// The difference in value for the asset in the trace
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                            `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiffJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceTypeAssetTrace EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceType = "AssetTrace"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceTypeAssetTrace:
		return true
	}
	return false
}

// The type of the model
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTypeErc721AssetTrace EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceType = "ERC721AssetTrace"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTypeErc721AssetTrace:
		return true
	}
	return false
}

// Label assigned to a trace entry.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabels string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabelsGasFee EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabels = "GAS_FEE"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabels) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabelsGasFee:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace struct {
	// Description of the asset in the trace
	Asset EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset `json:"asset,required"`
	// The difference in value for the asset in the trace
	Diff EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiff `json:"diff,required"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address,required"`
	// The address where the assets are moved to
	ToAddress string `json:"to_address,required"`
	// type of the trace
	TraceType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceType `json:"type,required"`
	// List of labels that describe the trace
	Labels []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabels `json:"labels"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceJSON     `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceJSON struct {
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

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                              `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetJSON `json:"-"`
	union  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                       `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                      `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetTypeErc1155 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetType = "ERC1155"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetTypeNonerc  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetTypeErc1155, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetTypeNonerc:
		return true
	}
	return false
}

// The difference in value for the asset in the trace
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                             `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiffJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceTypeAssetTrace EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceType = "AssetTrace"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceTypeAssetTrace:
		return true
	}
	return false
}

// The type of the model
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTypeErc1155AssetTrace EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceType = "ERC1155AssetTrace"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTypeErc1155AssetTrace:
		return true
	}
	return false
}

// Label assigned to a trace entry.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabels string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabelsGasFee EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabels = "GAS_FEE"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabels) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabelsGasFee:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace struct {
	// Description of the asset in the trace
	Asset EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAsset `json:"asset,required"`
	// The difference in value for the asset in the trace
	Diff EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiff `json:"diff,required"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address,required"`
	// The address where the assets are moved to
	ToAddress string `json:"to_address,required"`
	// type of the trace
	TraceType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceType `json:"type,required"`
	// List of labels that describe the trace
	Labels []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabels `json:"labels"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceJSON     `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceJSON struct {
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

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAsset struct {
	ChainID   int64  `json:"chain_id,required"`
	ChainName string `json:"chain_name,required"`
	Decimals  int64  `json:"decimals,required"`
	LogoURL   string `json:"logo_url,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetType `json:"type,required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                             `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAsset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetJSON struct {
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

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetTypeNative EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetType = "NATIVE"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetTypeNative:
		return true
	}
	return false
}

// The difference in value for the asset in the trace
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                            `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiffJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceTypeAssetTrace EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceType = "AssetTrace"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceTypeAssetTrace:
		return true
	}
	return false
}

// The type of the model
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTypeNativeAssetTrace EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceType = "NativeAssetTrace"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTypeNativeAssetTrace:
		return true
	}
	return false
}

// Label assigned to a trace entry.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabels string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabelsGasFee EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabels = "GAS_FEE"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabels) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabelsGasFee:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace struct {
	// Description of the asset in the trace
	Asset   EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset   `json:"asset,required"`
	Exposed EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposed `json:"exposed,required"`
	// The owner of the assets
	Owner string `json:"owner,required"`
	// The spender of the assets
	Spender string `json:"spender,required"`
	// type of the trace
	TraceType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceType `json:"type,required"`
	JSON evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceJSON struct {
	Asset       apijson.Field
	Exposed     apijson.Field
	Owner       apijson.Field
	Spender     apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetType `json:"type,required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                               `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetJSON `json:"-"`
	union  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                      `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                       `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetTypeErc20  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetType = "ERC20"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetTypeErc20, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetTypeNonerc:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposed struct {
	RawValue string                                                                                                                                                 `json:"raw_value,required"`
	UsdPrice float64                                                                                                                                                `json:"usd_price"`
	Value    float64                                                                                                                                                `json:"value"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposedJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposedJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposed]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposedJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposedJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceTypeExposureTrace EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceType = "ExposureTrace"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceTypeExposureTrace:
		return true
	}
	return false
}

// The type of the model
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTypeErc20ExposureTrace EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceType = "ERC20ExposureTrace"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTypeErc20ExposureTrace:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace struct {
	// Description of the asset in the trace
	Asset EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset `json:"asset,required"`
	// The owner of the assets
	Owner string `json:"owner,required"`
	// The spender of the assets
	Spender string `json:"spender,required"`
	// type of the trace
	TraceType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type    EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceType    `json:"type,required"`
	Exposed EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposed `json:"exposed"`
	JSON    evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceJSON    `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceJSON struct {
	Asset       apijson.Field
	Owner       apijson.Field
	Spender     apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	Exposed     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetJSON `json:"-"`
	union  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                        `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                        `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetTypeErc721 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetType = "ERC721"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetTypeErc721, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetTypeNonerc:
		return true
	}
	return false
}

// type of the trace
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceTypeExposureTrace EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceType = "ExposureTrace"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceTypeExposureTrace:
		return true
	}
	return false
}

// The type of the model
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTypeErc721ExposureTrace EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceType = "ERC721ExposureTrace"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTypeErc721ExposureTrace:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposed struct {
	Amount   int64                                                                                                                                                   `json:"amount,required"`
	TokenID  string                                                                                                                                                  `json:"token_id,required"`
	IsMint   bool                                                                                                                                                    `json:"is_mint"`
	LogoURL  string                                                                                                                                                  `json:"logo_url"`
	UsdPrice float64                                                                                                                                                 `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposedJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposedJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposed]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposedJSON struct {
	Amount      apijson.Field
	TokenID     apijson.Field
	IsMint      apijson.Field
	LogoURL     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposedJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace struct {
	// Description of the asset in the trace
	Asset EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset `json:"asset,required"`
	// The owner of the assets
	Owner string `json:"owner,required"`
	// The spender of the assets
	Spender string `json:"spender,required"`
	// type of the trace
	TraceType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceType `json:"type,required"`
	JSON evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceJSON struct {
	Asset       apijson.Field
	Owner       apijson.Field
	Spender     apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                 `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetJSON `json:"-"`
	union  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                          `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                         `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetTypeErc1155 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetType = "ERC1155"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetTypeNonerc  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetTypeErc1155, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetTypeNonerc:
		return true
	}
	return false
}

// type of the trace
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceTypeExposureTrace EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceType = "ExposureTrace"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceTypeExposureTrace:
		return true
	}
	return false
}

// The type of the model
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTypeErc1155ExposureTrace EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceType = "ERC1155ExposureTrace"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTypeErc1155ExposureTrace:
		return true
	}
	return false
}

// type of the trace
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceTypeAssetTrace    EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceType = "AssetTrace"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceTypeExposureTrace EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceType = "ExposureTrace"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceTypeAssetTrace, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceTypeExposureTrace:
		return true
	}
	return false
}

// The type of the model
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc20AssetTrace      EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC20AssetTrace"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc721AssetTrace     EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC721AssetTrace"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc1155AssetTrace    EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC1155AssetTrace"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeNativeAssetTrace     EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "NativeAssetTrace"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc20ExposureTrace   EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC20ExposureTrace"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc721ExposureTrace  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC721ExposureTrace"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc1155ExposureTrace EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC1155ExposureTrace"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc20AssetTrace, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc721AssetTrace, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc1155AssetTrace, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeNativeAssetTrace, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc20ExposureTrace, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc721ExposureTrace, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc1155ExposureTrace:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetail struct {
	// contains the contract's name if the address is a verified contract
	ContractName string `json:"contract_name"`
	// Whether the address is an eoa or a contract
	IsEoa bool `json:"is_eoa"`
	// known name tag for the address
	NameTag string                                                                                               `json:"name_tag"`
	JSON    evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetailJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetailJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetail]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetailJSON struct {
	ContractName apijson.Field
	IsEoa        apijson.Field
	NameTag      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetailJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff struct {
	// This field can have the runtime type of
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAsset].
	Asset interface{} `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType `json:"asset_type,required"`
	// This field can have the runtime type of
	// [[]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffIn],
	// [[]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffIn],
	// [[]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffIn],
	// [[]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffIn].
	In interface{} `json:"in,required"`
	// This field can have the runtime type of
	// [[]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOut],
	// [[]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOut],
	// [[]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOut],
	// [[]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOut].
	Out   interface{}                                                                                       `json:"out,required"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffJSON `json:"-"`
	union EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff struct {
	// description of the asset for the current diff
	Asset EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOut `json:"out,required"`
	JSON evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffJSON  `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff() {
}

// description of the asset for the current diff
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType `json:"type,required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                         `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetJSON `json:"-"`
	union  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                 `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetTypeErc20  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType = "ERC20"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetTypeErc20, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffIn struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                      `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffInJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffInJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffIn]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffInJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOut struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                       `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOutJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOutJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOut]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff struct {
	// description of the asset for the current diff
	Asset EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOut `json:"out,required"`
	JSON evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffJSON  `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff() {
}

// description of the asset for the current diff
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                          `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetJSON `json:"-"`
	union  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                  `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                  `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetTypeErc721 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType = "ERC721"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetTypeErc721, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffIn struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                       `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffInJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffInJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffIn]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffInJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOut struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                        `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOutJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOutJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOut]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOutJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff struct {
	// description of the asset for the current diff
	Asset EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOut `json:"out,required"`
	JSON evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffJSON  `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff() {
}

// description of the asset for the current diff
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                           `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetJSON `json:"-"`
	union  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                    `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                   `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetTypeErc1155 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType = "ERC1155"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetTypeNonerc  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetTypeErc1155, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffIn struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                        `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffInJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffInJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffIn]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffInJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOut struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                         `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOutJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOutJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOut]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOutJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff struct {
	// description of the asset for the current diff
	Asset EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOut `json:"out,required"`
	JSON evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffJSON  `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff() {
}

// description of the asset for the current diff
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAsset struct {
	ChainID   int64  `json:"chain_id,required"`
	ChainName string `json:"chain_name,required"`
	Decimals  int64  `json:"decimals,required"`
	LogoURL   string `json:"logo_url,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetType `json:"type,required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                          `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAsset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetJSON struct {
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

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetTypeNative EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetType = "NATIVE"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetTypeNative:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffIn struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                       `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffInJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffInJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffIn]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffInJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOut struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                        `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOutJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOutJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOut]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

// type of the asset for the current diff
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc20   EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType = "ERC20"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc721  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType = "ERC721"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc1155 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType = "ERC1155"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeNative  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType = "NATIVE"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc20, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc721, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc1155, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeNative:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure struct {
	// This field can have the runtime type of
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset].
	Asset interface{} `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType `json:"asset_type,required"`
	// This field can have the runtime type of
	// [map[string]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpender],
	// [map[string]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpender],
	// [map[string]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpender].
	Spenders interface{}                                                                                     `json:"spenders,required"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposureJSON `json:"-"`
	union    EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposureJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresUnion {
	return r.union
}

// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType `json:"asset_type,required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpender `json:"spenders,required"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureJSON               `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure() {
}

// description of the asset for the current diff
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType `json:"type,required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                      `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetJSON `json:"-"`
	union  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                             `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                              `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetTypeErc20  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType = "ERC20"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetTypeErc20, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpender struct {
	// the amount that was asked in the approval request for this spender from the
	// current address and asset
	Approval string                                                                                                                                               `json:"approval,required"`
	Exposure []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure `json:"exposure,required"`
	// the usd price of the approval amount
	ApprovalUsdPrice string `json:"approval_usd_price"`
	// the expiration time of the permit2 protocol
	Expiration time.Time `json:"expiration" format:"date-time"`
	// user friendly description of the approval
	Summary string                                                                                                                                        `json:"summary"`
	JSON    evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpender]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON struct {
	Approval         apijson.Field
	Exposure         apijson.Field
	ApprovalUsdPrice apijson.Field
	Expiration       apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure struct {
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
	Value string                                                                                                                                                 `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON `json:"-"`
	union EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON struct {
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

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                              `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                             `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                            `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                             `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType `json:"asset_type,required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpender `json:"spenders,required"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureJSON               `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure() {
}

// description of the asset for the current diff
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                       `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetJSON `json:"-"`
	union  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                               `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                               `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetTypeErc721 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType = "ERC721"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetTypeErc721, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpender struct {
	Exposure []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure `json:"exposure,required"`
	// boolean indicates whether an is_approved_for_all function was used (missing in
	// case of ERC20 / ERC1155)
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// user friendly description of the approval
	Summary string                                                                                                                                         `json:"summary"`
	JSON    evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpender]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure struct {
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
	Value string                                                                                                                                                  `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON `json:"-"`
	union EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON struct {
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

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                               `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                              `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                             `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                              `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType `json:"asset_type,required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpender `json:"spenders,required"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureJSON               `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure() {
}

// description of the asset for the current diff
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                        `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON `json:"-"`
	union  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                 `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeErc1155 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType = "ERC1155"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeNonerc  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType = "NONERC"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeErc1155, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpender struct {
	Exposure []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure `json:"exposure,required"`
	// boolean indicates whether an is_approved_for_all function was used (missing in
	// case of ERC20 / ERC1155)
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// user friendly description of the approval
	Summary string                                                                                                                                          `json:"summary"`
	JSON    evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpender]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure struct {
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
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON `json:"-"`
	union EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON struct {
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

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                                `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string                                                                                                                                                                               `json:"usd_price"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                              `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                               `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

// type of the asset for the current diff
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc20   EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType = "ERC20"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc721  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType = "ERC721"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc1155 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType = "ERC1155"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc20, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc721, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc1155:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKey struct {
	Key      string                                                                                                `json:"key,required"`
	Policies []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy `json:"policies,required"`
	Signer   string                                                                                                `json:"signer,required"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyJSON     `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKey]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyJSON struct {
	Key         apijson.Field
	Policies    apijson.Field
	Signer      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy struct {
	// This field can have the runtime type of
	// [[]EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArg].
	Args interface{} `json:"args"`
	// This field can have the runtime type of
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails].
	AssetDetails interface{}                                                                                               `json:"asset_details"`
	EndTime      time.Time                                                                                                 `json:"end_time" format:"date-time"`
	Method       string                                                                                                    `json:"method"`
	Recipient    string                                                                                                    `json:"recipient"`
	StartTime    time.Time                                                                                                 `json:"start_time" format:"date-time"`
	ToAddress    string                                                                                                    `json:"to_address"`
	Type         EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType `json:"type"`
	ValueLimit   string                                                                                                    `json:"value_limit"`
	JSON         evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicyJSON   `json:"-"`
	union        EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicyJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicyJSON struct {
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

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicyJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesUnion {
	return r.union
}

// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy struct {
	AssetDetails EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails `json:"asset_details,required"`
	Recipient    string                                                                                                                                               `json:"recipient"`
	Type         EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyType         `json:"type"`
	ValueLimit   string                                                                                                                                               `json:"value_limit"`
	JSON         evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyJSON         `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyJSON struct {
	AssetDetails apijson.Field
	Recipient    apijson.Field
	Type         apijson.Field
	ValueLimit   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails struct {
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsType `json:"type,required"`
	// address of the token
	Address   string `json:"address"`
	ChainID   int64  `json:"chain_id"`
	ChainName string `json:"chain_name"`
	LogoURL   string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                   `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsJSON `json:"-"`
	union  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsJSON struct {
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

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsUnion {
	return r.union
}

// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails struct {
	ChainID   int64  `json:"chain_id,required"`
	ChainName string `json:"chain_name,required"`
	Decimals  int64  `json:"decimals,required"`
	LogoURL   string `json:"logo_url,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsType `json:"type,required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                           `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsJSON struct {
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

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsTypeNative EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsType = "NATIVE"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsTypeNative:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                          `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsTypeNative EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsType = "NATIVE"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsTypeErc20  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsType = "ERC20"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsTypeNative, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsTypeErc20:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyTypeTransferPolicy EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyType = "TRANSFER_POLICY"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyTypeTransferPolicy:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy struct {
	ToAddress  string                                                                                                                                    `json:"to_address,required"`
	Args       []EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArg `json:"args"`
	Method     string                                                                                                                                    `json:"method"`
	Type       EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyType  `json:"type"`
	ValueLimit string                                                                                                                                    `json:"value_limit"`
	JSON       evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyJSON  `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyJSON struct {
	ToAddress   apijson.Field
	Args        apijson.Field
	Method      apijson.Field
	Type        apijson.Field
	ValueLimit  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArg struct {
	// Comparison operator used to evaluate an argument/value against a policy
	// constraint.
	Condition EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition `json:"condition,required"`
	Index     int64                                                                                                                                             `json:"index,required"`
	Value     string                                                                                                                                            `json:"value,required"`
	JSON      evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgJSON       `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArg]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgJSON struct {
	Condition   apijson.Field
	Index       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgJSON) RawJSON() string {
	return r.raw
}

// Comparison operator used to evaluate an argument/value against a policy
// constraint.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionUnconstrained  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "UNCONSTRAINED"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionEqual          EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "EQUAL"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionGreater        EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "GREATER"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionLess           EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "LESS"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionGreaterOrEqual EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "GREATER_OR_EQUAL"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionLessOrEqual    EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "LESS_OR_EQUAL"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionNotEqual       EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "NOT_EQUAL"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionUnconstrained, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionEqual, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionGreater, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionLess, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionGreaterOrEqual, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionLessOrEqual, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionNotEqual:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyTypeCallPolicy EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyType = "CALL_POLICY"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyTypeCallPolicy:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy struct {
	ValueLimit string                                                                                                                                  `json:"value_limit,required"`
	Type       EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyType `json:"type"`
	JSON       evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyJSON struct {
	ValueLimit  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyTypeGasPolicy EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyType = "GAS_POLICY"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyTypeGasPolicy:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy struct {
	EndTime   time.Time                                                                                                                                      `json:"end_time" format:"date-time"`
	StartTime time.Time                                                                                                                                      `json:"start_time" format:"date-time"`
	Type      EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyType `json:"type"`
	JSON      evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyTypeExpirationPolicy EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyType = "EXPIRATION_POLICY"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyTypeExpirationPolicy:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeTransferPolicy   EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType = "TRANSFER_POLICY"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeCallPolicy       EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType = "CALL_POLICY"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeGasPolicy        EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType = "GAS_POLICY"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeExpirationPolicy EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType = "EXPIRATION_POLICY"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeTransferPolicy, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeCallPolicy, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeGasPolicy, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeExpirationPolicy:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationStatus string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationStatusSuccess EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationStatus = "Success"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationStatus) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationStatusSuccess:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiff struct {
	In    string                                                                                              `json:"in,required"`
	Out   string                                                                                              `json:"out,required"`
	Total string                                                                                              `json:"total,required"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiffJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiffJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiff]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

// Type of action performed in the transaction.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsMint            EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "mint"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsStake           EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "stake"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsSwap            EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "swap"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsNativeTransfer  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "native_transfer"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsTokenTransfer   EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "token_transfer"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsApproval        EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "approval"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsSetCodeAccount  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "set_code_account"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsProxyUpgrade    EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "proxy_upgrade"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsOwnershipChange EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "ownership_change"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsMint, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsStake, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsSwap, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsNativeTransfer, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsTokenTransfer, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsApproval, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsSetCodeAccount, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsProxyUpgrade, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsOwnershipChange:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement struct {
	// The type of the state change
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType `json:"type,required"`
	// This field can have the runtime type of
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfter],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfter],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfter].
	After interface{} `json:"after"`
	// This field can have the runtime type of
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBefore],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBefore],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBefore].
	Before interface{} `json:"before"`
	// The delegated address
	DelegatedAddress string `json:"delegated_address"`
	// The direct creator address of the new contract
	Deployer string                                                                                                    `json:"deployer"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementJSON `json:"-"`
	union    EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementJSON struct {
	Type             apijson.Field
	After            apijson.Field
	Before           apijson.Field
	DelegatedAddress apijson.Field
	Deployer         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementUnion {
	return r.union
}

// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement struct {
	// The state after the transaction
	After EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfter `json:"after,required"`
	// The state before the transaction
	Before EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBefore `json:"before,required"`
	// The type of the state change
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementType `json:"type,required"`
	JSON evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementJSON struct {
	After       apijson.Field
	Before      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement() {
}

// The state after the transaction
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfter struct {
	Address string                                                                                                                                                 `json:"address,required"`
	JSON    evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfterJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfterJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfter]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfterJSON struct {
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfterJSON) RawJSON() string {
	return r.raw
}

// The state before the transaction
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBefore struct {
	Address string                                                                                                                                                  `json:"address,required"`
	JSON    evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBeforeJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBeforeJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBefore]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBeforeJSON struct {
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBeforeJSON) RawJSON() string {
	return r.raw
}

// The type of the state change
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementTypeProxyUpgrade EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementType = "PROXY_UPGRADE"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementTypeProxyUpgrade:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement struct {
	// The state after the transaction
	After EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfter `json:"after,required"`
	// The state before the transaction
	Before EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBefore `json:"before,required"`
	// The type of the state change
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementType `json:"type,required"`
	JSON evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementJSON struct {
	After       apijson.Field
	Before      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement() {
}

// The state after the transaction
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfter struct {
	Owners []string                                                                                                                                                  `json:"owners,required"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfterJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfterJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfter]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfterJSON struct {
	Owners      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfterJSON) RawJSON() string {
	return r.raw
}

// The state before the transaction
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBefore struct {
	Owners []string                                                                                                                                                   `json:"owners,required"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBeforeJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBeforeJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBefore]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBeforeJSON struct {
	Owners      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBeforeJSON) RawJSON() string {
	return r.raw
}

// The type of the state change
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementTypeOwnershipChange EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementType = "OWNERSHIP_CHANGE"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementTypeOwnershipChange:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement struct {
	// The state after the transaction
	After EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfter `json:"after,required"`
	// The state before the transaction
	Before EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBefore `json:"before,required"`
	// The type of the state change
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementType `json:"type,required"`
	JSON evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementJSON struct {
	After       apijson.Field
	Before      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement() {
}

// The state after the transaction
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfter struct {
	Modules []string                                                                                                                                                `json:"modules,required"`
	JSON    evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfterJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfterJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfter]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfterJSON struct {
	Modules     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfterJSON) RawJSON() string {
	return r.raw
}

// The state before the transaction
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBefore struct {
	Modules []string                                                                                                                                                 `json:"modules,required"`
	JSON    evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBeforeJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBeforeJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBefore]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBeforeJSON struct {
	Modules     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBeforeJSON) RawJSON() string {
	return r.raw
}

// The type of the state change
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementTypeModuleChange EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementType = "MODULE_CHANGE"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementTypeModuleChange:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement struct {
	// The delegated address
	DelegatedAddress string `json:"delegated_address,required"`
	// The type of the state change
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementType `json:"type,required"`
	JSON evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementJSON struct {
	DelegatedAddress apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement() {
}

// The type of the state change
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementTypeSetCodeAccount EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementType = "SET_CODE_ACCOUNT"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementTypeSetCodeAccount:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation struct {
	// The direct creator address of the new contract
	Deployer string `json:"deployer,required"`
	// The type of the state change
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationType `json:"type,required"`
	JSON evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationJSON struct {
	Deployer    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement() {
}

// The type of the state change
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationTypeContractCreation EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationType = "CONTRACT_CREATION"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationTypeContractCreation:
		return true
	}
	return false
}

// The type of the state change
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeProxyUpgrade     EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType = "PROXY_UPGRADE"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeOwnershipChange  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType = "OWNERSHIP_CHANGE"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeModuleChange     EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType = "MODULE_CHANGE"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeSetCodeAccount   EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType = "SET_CODE_ACCOUNT"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeContractCreation EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType = "CONTRACT_CREATION"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeProxyUpgrade, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeOwnershipChange, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeModuleChange, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeSetCodeAccount, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeContractCreation:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalance struct {
	// The asset that is missing balance
	Asset EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset `json:"asset,required"`
	// The account address's current balance of the asset
	CurrentBalance EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalance `json:"current_balance,required"`
	// The account address's missing balance of the asset
	MissingBalance EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalance `json:"missing_balance,required"`
	// The required balance of the asset for this action
	RequiredBalance EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalance `json:"required_balance,required"`
	JSON            evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalanceJSON             `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalanceJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalance]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalanceJSON struct {
	Asset           apijson.Field
	CurrentBalance  apijson.Field
	MissingBalance  apijson.Field
	RequiredBalance apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalanceJSON) RawJSON() string {
	return r.raw
}

// The asset that is missing balance
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset struct {
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetType `json:"type,required"`
	// address of the token
	Address   string `json:"address"`
	ChainID   int64  `json:"chain_id"`
	ChainName string `json:"chain_name"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                      `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetJSON `json:"-"`
	union  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetJSON struct {
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

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetUnion {
	return r.union
}

// The asset that is missing balance
//
// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                             `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails struct {
	ChainID   int64  `json:"chain_id,required"`
	ChainName string `json:"chain_name,required"`
	Decimals  int64  `json:"decimals,required"`
	LogoURL   string `json:"logo_url,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsType `json:"type,required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                              `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsJSON struct {
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

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset() {
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsTypeNative EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsType = "NATIVE"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsTypeNative:
		return true
	}
	return false
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetTypeErc20  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetType = "ERC20"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetTypeNative EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetType = "NATIVE"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetTypeErc20, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetTypeNative:
		return true
	}
	return false
}

// The account address's current balance of the asset
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalance struct {
	// The raw value of the balance in hex string format
	RawValue string `json:"raw_value,required"`
	// The value of the balance in decimal string format
	Value string                                                                                                               `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalanceJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalanceJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalance]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalanceJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalanceJSON) RawJSON() string {
	return r.raw
}

// The account address's missing balance of the asset
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalance struct {
	// The raw value of the balance in hex string format
	RawValue string `json:"raw_value,required"`
	// The value of the balance in decimal string format
	Value string                                                                                                               `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalanceJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalanceJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalance]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalanceJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalanceJSON) RawJSON() string {
	return r.raw
}

// The required balance of the asset for this action
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalance struct {
	// The raw value of the balance in hex string format
	RawValue string `json:"raw_value,required"`
	// The value of the balance in decimal string format
	Value string                                                                                                                `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalanceJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalanceJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalance]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalanceJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalanceJSON) RawJSON() string {
	return r.raw
}

// The parameters of the transaction that was simulated.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParams struct {
	// The block tag to be sent.
	BlockTag string `json:"block_tag"`
	// The calldata to be sent.
	Calldata EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldata `json:"calldata"`
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
	UserOperationCalldata EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldata `json:"user_operation_calldata"`
	// The value to be sent.
	Value string                                                                                        `json:"value"`
	JSON  evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParams]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsJSON struct {
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

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsJSON) RawJSON() string {
	return r.raw
}

// The calldata to be sent.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldata struct {
	// The function selector of the function called in the transaction
	FunctionSelector string `json:"function_selector,required"`
	// The function declaration of the function called in the transaction
	FunctionDeclaration string `json:"function_declaration"`
	// The function signature of the function called in the transaction
	FunctionSignature string                                                                                                `json:"function_signature"`
	JSON              evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldataJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldataJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldata]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldataJSON struct {
	FunctionSelector    apijson.Field
	FunctionDeclaration apijson.Field
	FunctionSignature   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldataJSON) RawJSON() string {
	return r.raw
}

// The user operation call data to be sent.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldata struct {
	// The function selector of the function called in the transaction
	FunctionSelector string `json:"function_selector,required"`
	// The function declaration of the function called in the transaction
	FunctionDeclaration string `json:"function_declaration"`
	// The function signature of the function called in the transaction
	FunctionSignature string                                                                                                             `json:"function_signature"`
	JSON              evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldataJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldataJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldata]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldataJSON struct {
	FunctionSelector    apijson.Field
	FunctionDeclaration apijson.Field
	FunctionSignature   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldataJSON) RawJSON() string {
	return r.raw
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationError struct {
	// A string explaining why the transaction failed
	Description string `json:"description,required"`
	// An error message if the simulation failed.
	Error string `json:"error,required"`
	// A string indicating if the simulation was successful or not.
	Status EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatus `json:"status,required"`
	// Error details if the simulation failed.
	ErrorDetails EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails `json:"error_details"`
	JSON         evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorJSON         `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationError]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorJSON struct {
	Description  apijson.Field
	Error        apijson.Field
	Status       apijson.Field
	ErrorDetails apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationError) implementsEvmPostTransactionBulkScanResponseSimulation() {
}

// A string indicating if the simulation was successful or not.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatus string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatusError EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatus = "Error"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatus) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatusError:
		return true
	}
	return false
}

// Error details if the simulation failed.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails struct {
	// The type of the model
	Code string `json:"code,required"`
	// The address of the account
	AccountAddress string `json:"account_address"`
	// The address that is invalid
	Address string `json:"address"`
	// This field can have the runtime type of
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset].
	Asset    interface{} `json:"asset"`
	Category string      `json:"category"`
	// The current balance of the account
	CurrentBalance string `json:"current_balance"`
	// The domain name that is unsupported
	DomainName string `json:"domain_name"`
	// The message type that is unsupported
	MessageType string `json:"message_type"`
	// The required balance of the account
	RequiredBalance string                                                                                                   `json:"required_balance"`
	JSON            evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsJSON `json:"-"`
	union           EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsJSON struct {
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

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsUnion {
	return r.union
}

// Error details if the simulation failed.
//
// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails struct {
	// The address of the account
	AccountAddress string `json:"account_address,required"`
	// The asset that the account does not have enough balance for
	Asset EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset `json:"asset,required"`
	// The type of the model
	Code     EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCode     `json:"code,required"`
	Category EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCategory `json:"category"`
	// The current balance of the account
	CurrentBalance string `json:"current_balance"`
	// The required balance of the account
	RequiredBalance string                                                                                                                                                         `json:"required_balance"`
	JSON            evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsJSON struct {
	AccountAddress  apijson.Field
	Asset           apijson.Field
	Code            apijson.Field
	Category        apijson.Field
	CurrentBalance  apijson.Field
	RequiredBalance apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails() {
}

// The asset that the account does not have enough balance for
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset struct {
	// This field can have the runtime type of
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetails],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetails],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetails],
	// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetails].
	Details interface{} `json:"details,required"`
	// The type of the model
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType `json:"type,required"`
	// Token Id
	TokenID int64                                                                                                                                                               `json:"token_id"`
	JSON    evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetJSON `json:"-"`
	union   EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetUnion
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetJSON struct {
	Details     apijson.Field
	Type        apijson.Field
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset].
func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset) AsUnion() EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetUnion {
	return r.union
}

// The asset that the account does not have enough balance for
//
// Union satisfied by
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset],
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset]
// or
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset].
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetUnion interface {
	implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset struct {
	// Details
	Details EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetails `json:"details,required"`
	// The type of the model
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetType `json:"type,required"`
	JSON evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetJSON struct {
	Details     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset() {
}

// Details
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetails struct {
	ChainID   int64  `json:"chain_id,required"`
	ChainName string `json:"chain_name,required"`
	Decimals  int64  `json:"decimals,required"`
	LogoURL   string `json:"logo_url,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsType `json:"type,required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                  `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsJSON struct {
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

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsTypeNative EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsType = "NATIVE"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsTypeNative:
		return true
	}
	return false
}

// The type of the model
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetTypeNative EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetType = "NATIVE"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetTypeNative:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset struct {
	// Details
	Details EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetails `json:"details,required"`
	// The type of the model
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetType `json:"type,required"`
	JSON evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetJSON struct {
	Details     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset() {
}

// Details
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                 `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsTypeErc20 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsType = "ERC20"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsTypeErc20:
		return true
	}
	return false
}

// The type of the model
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetTypeErc20 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetType = "ERC20"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetTypeErc20:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset struct {
	// Details
	Details EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetails `json:"details,required"`
	// Token Id
	TokenID int64 `json:"token_id,required"`
	// The type of the model
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetType `json:"type,required"`
	JSON evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetJSON struct {
	Details     apijson.Field
	TokenID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset() {
}

// Details
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                  `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsTypeErc1155 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsType = "ERC1155"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsTypeErc1155:
		return true
	}
	return false
}

// The type of the model
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetTypeErc721 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetType = "ERC721"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetTypeErc721:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset struct {
	// Details
	Details EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetails `json:"details,required"`
	// Token Id
	TokenID int64 `json:"token_id,required"`
	// The type of the model
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetType `json:"type,required"`
	JSON evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetJSON struct {
	Details     apijson.Field
	TokenID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset() {
}

// Details
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                   `json:"symbol"`
	JSON   evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsTypeErc1155 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsType = "ERC1155"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsTypeErc1155:
		return true
	}
	return false
}

// The type of the model
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetTypeErc1155 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetType = "ERC1155"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetTypeErc1155:
		return true
	}
	return false
}

// The type of the model
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeNative  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType = "NATIVE"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc20   EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType = "ERC20"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc721  EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType = "ERC721"
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc1155 EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType = "ERC1155"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeNative, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc20, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc721, EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc1155:
		return true
	}
	return false
}

// The type of the model
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCode string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCodeGeneralInsufficientFunds EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCode = "GENERAL_INSUFFICIENT_FUNDS"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCode) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCodeGeneralInsufficientFunds:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCategory string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCategoryRevert EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCategory = "REVERT"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCategory) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCategoryRevert:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails struct {
	// The address that is invalid
	Address string `json:"address,required"`
	// The type of the model
	Code     EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCode     `json:"code,required"`
	Category EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCategory `json:"category"`
	JSON     evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsJSON     `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsJSON struct {
	Address     apijson.Field
	Code        apijson.Field
	Category    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails() {
}

// The type of the model
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCode string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCodeGeneralInvalidAddress EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCode = "GENERAL_INVALID_ADDRESS"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCode) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCodeGeneralInvalidAddress:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCategory string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCategoryInvalidInput EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCategory = "INVALID_INPUT"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCategory) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCategoryInvalidInput:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails struct {
	Category string `json:"category,required"`
	// The error code
	Code string                                                                                                                                        `json:"code,required"`
	JSON evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetailsJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetailsJSON struct {
	Category    apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails() {
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails struct {
	// The type of the model
	Code EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCode `json:"code,required"`
	// The domain name that is unsupported
	DomainName string `json:"domain_name,required"`
	// The message type that is unsupported
	MessageType string                                                                                                                                                             `json:"message_type,required"`
	Category    EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCategory `json:"category"`
	JSON        evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsJSON     `json:"-"`
}

// evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails]
type evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsJSON struct {
	Code        apijson.Field
	DomainName  apijson.Field
	MessageType apijson.Field
	Category    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails) implementsEvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails() {
}

// The type of the model
type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCode string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCodeUnsupportedEip712Message EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCode = "UNSUPPORTED_EIP712_MESSAGE"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCode) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCodeUnsupportedEip712Message:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCategory string

const (
	EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCategoryInvalidInput EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCategory = "INVALID_INPUT"
)

func (r EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCategory) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCategoryInvalidInput:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type EvmPostTransactionBulkScanResponseSimulationStatus string

const (
	EvmPostTransactionBulkScanResponseSimulationStatusSuccess EvmPostTransactionBulkScanResponseSimulationStatus = "Success"
	EvmPostTransactionBulkScanResponseSimulationStatusError   EvmPostTransactionBulkScanResponseSimulationStatus = "Error"
)

func (r EvmPostTransactionBulkScanResponseSimulationStatus) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseSimulationStatusSuccess, EvmPostTransactionBulkScanResponseSimulationStatusError:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseUserOperationGasEstimation struct {
	Status                           EvmPostTransactionBulkScanResponseUserOperationGasEstimationStatus `json:"status,required"`
	CallGasEstimate                  string                                                             `json:"call_gas_estimate"`
	Error                            string                                                             `json:"error"`
	PaymasterVerificationGasEstimate string                                                             `json:"paymaster_verification_gas_estimate"`
	PreVerificationGasEstimate       string                                                             `json:"pre_verification_gas_estimate"`
	VerificationGasEstimate          string                                                             `json:"verification_gas_estimate"`
	JSON                             evmPostTransactionBulkScanResponseUserOperationGasEstimationJSON   `json:"-"`
	union                            EvmPostTransactionBulkScanResponseUserOperationGasEstimationUnion
}

// evmPostTransactionBulkScanResponseUserOperationGasEstimationJSON contains the
// JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseUserOperationGasEstimation]
type evmPostTransactionBulkScanResponseUserOperationGasEstimationJSON struct {
	Status                           apijson.Field
	CallGasEstimate                  apijson.Field
	Error                            apijson.Field
	PaymasterVerificationGasEstimate apijson.Field
	PreVerificationGasEstimate       apijson.Field
	VerificationGasEstimate          apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r evmPostTransactionBulkScanResponseUserOperationGasEstimationJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseUserOperationGasEstimation) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseUserOperationGasEstimation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmPostTransactionBulkScanResponseUserOperationGasEstimationUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [UserOperationV6GasEstimation],
// [UserOperationV7GasEstimation],
// [EvmPostTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError].
func (r EvmPostTransactionBulkScanResponseUserOperationGasEstimation) AsUnion() EvmPostTransactionBulkScanResponseUserOperationGasEstimationUnion {
	return r.union
}

// Union satisfied by [UserOperationV6GasEstimation],
// [UserOperationV7GasEstimation] or
// [EvmPostTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError].
type EvmPostTransactionBulkScanResponseUserOperationGasEstimationUnion interface {
	implementsEvmPostTransactionBulkScanResponseUserOperationGasEstimation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseUserOperationGasEstimationUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError struct {
	Error  string                                                                                                              `json:"error,required"`
	Status EvmPostTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus `json:"status,required"`
	JSON   evmPostTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON   `json:"-"`
}

// evmPostTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError]
type evmPostTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError) implementsEvmPostTransactionBulkScanResponseUserOperationGasEstimation() {
}

type EvmPostTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus string

const (
	EvmPostTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatusError EvmPostTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus = "Error"
)

func (r EvmPostTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatusError:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseUserOperationGasEstimationStatus string

const (
	EvmPostTransactionBulkScanResponseUserOperationGasEstimationStatusSuccess EvmPostTransactionBulkScanResponseUserOperationGasEstimationStatus = "Success"
	EvmPostTransactionBulkScanResponseUserOperationGasEstimationStatusError   EvmPostTransactionBulkScanResponseUserOperationGasEstimationStatus = "Error"
)

func (r EvmPostTransactionBulkScanResponseUserOperationGasEstimationStatus) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseUserOperationGasEstimationStatusSuccess, EvmPostTransactionBulkScanResponseUserOperationGasEstimationStatusError:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseValidation struct {
	// This field can have the runtime type of
	// [[]EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeature],
	// [[]EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeature].
	Features interface{} `json:"features,required"`
	// Result type returned when validation succeeds.
	ResultType EvmPostTransactionBulkScanResponseValidationResultType `json:"result_type,required"`
	// A string indicating if the simulation was successful or not.
	Status EvmPostTransactionBulkScanResponseValidationStatus `json:"status,required"`
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
	Reason string                                           `json:"reason"`
	JSON   evmPostTransactionBulkScanResponseValidationJSON `json:"-"`
	union  EvmPostTransactionBulkScanResponseValidationUnion
}

// evmPostTransactionBulkScanResponseValidationJSON contains the JSON metadata for
// the struct [EvmPostTransactionBulkScanResponseValidation]
type evmPostTransactionBulkScanResponseValidationJSON struct {
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

func (r evmPostTransactionBulkScanResponseValidationJSON) RawJSON() string {
	return r.raw
}

func (r *EvmPostTransactionBulkScanResponseValidation) UnmarshalJSON(data []byte) (err error) {
	*r = EvmPostTransactionBulkScanResponseValidation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EvmPostTransactionBulkScanResponseValidationUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidation],
// [EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationError].
func (r EvmPostTransactionBulkScanResponseValidation) AsUnion() EvmPostTransactionBulkScanResponseValidationUnion {
	return r.union
}

// Union satisfied by
// [EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidation]
// or
// [EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationError].
type EvmPostTransactionBulkScanResponseValidationUnion interface {
	implementsEvmPostTransactionBulkScanResponseValidation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmPostTransactionBulkScanResponseValidationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationError{}),
		},
	)
}

type EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidation struct {
	// A list of features about this transaction explaining the validation.
	Features []EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeature `json:"features,required"`
	// Result type returned when validation succeeds.
	ResultType EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationResultType `json:"result_type,required"`
	// A string indicating if the simulation was successful or not.
	Status EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationStatus `json:"status,required"`
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification"`
	// A textual description that can be presented to the user about what this
	// transaction is doing.
	Description string `json:"description"`
	// A textual description about the reasons the transaction was flagged with
	// result_type.
	Reason string                                                                                  `json:"reason"`
	JSON   evmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidation]
type evmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationJSON struct {
	Features       apijson.Field
	ResultType     apijson.Field
	Status         apijson.Field
	Classification apijson.Field
	Description    apijson.Field
	Reason         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidation) implementsEvmPostTransactionBulkScanResponseValidation() {
}

type EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeature struct {
	// Textual description
	Description string `json:"description,required"`
	// Feature name
	FeatureID string `json:"feature_id,required"`
	// Security result of a transaction scan feature.
	Type EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType `json:"type,required"`
	// Address the feature refers to
	Address string `json:"address"`
	// Metadata related to the feature
	Metadata interface{}                                                                                    `json:"metadata"`
	JSON     evmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeatureJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeatureJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeature]
type evmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeatureJSON struct {
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	Address     apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeatureJSON) RawJSON() string {
	return r.raw
}

// Security result of a transaction scan feature.
type EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType string

const (
	EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeMalicious EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType = "Malicious"
	EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeWarning   EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType = "Warning"
	EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeBenign    EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType = "Benign"
	EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeInfo      EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType = "Info"
)

func (r EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeMalicious, EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeWarning, EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeBenign, EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeInfo:
		return true
	}
	return false
}

// Result type returned when validation succeeds.
type EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationResultType string

const (
	EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeBenign    EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationResultType = "Benign"
	EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeWarning   EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationResultType = "Warning"
	EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeMalicious EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationResultType = "Malicious"
)

func (r EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationResultType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeBenign, EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeWarning, EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeMalicious:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationStatus string

const (
	EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationStatusSuccess EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationStatus = "Success"
)

func (r EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationStatus) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationStatusSuccess:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationError struct {
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassification `json:"classification,required"`
	// A textual description that can be presented to the user about what this
	// transaction is doing.
	Description EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescription `json:"description,required"`
	// An error message if the validation failed.
	Error string `json:"error,required"`
	// A list of features about this transaction explaining the validation.
	Features []EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeature `json:"features,required"`
	// A textual description about the reasons the transaction was flagged with
	// result_type.
	Reason EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorReason `json:"reason,required"`
	// A string indicating if the transaction is safe to sign or not.
	ResultType EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultType `json:"result_type,required"`
	// A string indicating if the simulation was successful or not.
	Status EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatus `json:"status,required"`
	JSON   evmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorJSON   `json:"-"`
}

// evmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationError]
type evmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorJSON struct {
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

func (r *EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorJSON) RawJSON() string {
	return r.raw
}

func (r EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationError) implementsEvmPostTransactionBulkScanResponseValidation() {
}

// A textual classification that can be presented to the user explaining the
// reason.
type EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassification string

const (
	EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassificationEmpty EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassification = ""
)

func (r EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassification) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassificationEmpty:
		return true
	}
	return false
}

// A textual description that can be presented to the user about what this
// transaction is doing.
type EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescription string

const (
	EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescriptionEmpty EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescription = ""
)

func (r EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescription) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescriptionEmpty:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeature struct {
	// Textual description
	Description string `json:"description,required"`
	// Feature name
	FeatureID string `json:"feature_id,required"`
	// Security result of a transaction scan feature.
	Type EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType `json:"type,required"`
	// Address the feature refers to
	Address string `json:"address"`
	// Metadata related to the feature
	Metadata interface{}                                                                                         `json:"metadata"`
	JSON     evmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeatureJSON `json:"-"`
}

// evmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeatureJSON
// contains the JSON metadata for the struct
// [EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeature]
type evmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeatureJSON struct {
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	Address     apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeatureJSON) RawJSON() string {
	return r.raw
}

// Security result of a transaction scan feature.
type EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType string

const (
	EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeMalicious EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType = "Malicious"
	EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeWarning   EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType = "Warning"
	EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeBenign    EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType = "Benign"
	EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeInfo      EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType = "Info"
)

func (r EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeMalicious, EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeWarning, EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeBenign, EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeInfo:
		return true
	}
	return false
}

// A textual description about the reasons the transaction was flagged with
// result_type.
type EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorReason string

const (
	EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorReasonEmpty EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorReason = ""
)

func (r EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorReason) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorReasonEmpty:
		return true
	}
	return false
}

// A string indicating if the transaction is safe to sign or not.
type EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultType string

const (
	EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultTypeError EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultType = "Error"
)

func (r EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultTypeError:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatus string

const (
	EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatusSuccess EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatus = "Success"
)

func (r EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatus) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatusSuccess:
		return true
	}
	return false
}

// Result type returned when validation succeeds.
type EvmPostTransactionBulkScanResponseValidationResultType string

const (
	EvmPostTransactionBulkScanResponseValidationResultTypeBenign    EvmPostTransactionBulkScanResponseValidationResultType = "Benign"
	EvmPostTransactionBulkScanResponseValidationResultTypeWarning   EvmPostTransactionBulkScanResponseValidationResultType = "Warning"
	EvmPostTransactionBulkScanResponseValidationResultTypeMalicious EvmPostTransactionBulkScanResponseValidationResultType = "Malicious"
	EvmPostTransactionBulkScanResponseValidationResultTypeError     EvmPostTransactionBulkScanResponseValidationResultType = "Error"
)

func (r EvmPostTransactionBulkScanResponseValidationResultType) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseValidationResultTypeBenign, EvmPostTransactionBulkScanResponseValidationResultTypeWarning, EvmPostTransactionBulkScanResponseValidationResultTypeMalicious, EvmPostTransactionBulkScanResponseValidationResultTypeError:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type EvmPostTransactionBulkScanResponseValidationStatus string

const (
	EvmPostTransactionBulkScanResponseValidationStatusSuccess EvmPostTransactionBulkScanResponseValidationStatus = "Success"
)

func (r EvmPostTransactionBulkScanResponseValidationStatus) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanResponseValidationStatusSuccess:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanParams struct {
	// The chain name or chain ID
	Chain param.Field[TransactionScanSupportedChain] `json:"chain,required"`
	// Transaction hashes to scan
	Data param.Field[[]string] `json:"data,required"`
	// Additional context for the scan (e.g., dapp URL/domain, integration source).
	// Used to enrich results and reduce false positives/negatives.
	Metadata param.Field[EvmPostTransactionBulkScanParamsMetadataUnion] `json:"metadata,required"`
	// The relative block for the block validation. Can be "latest" or a block number.
	Block param.Field[EvmPostTransactionBulkScanParamsBlockUnion] `json:"block"`
	// List of one or both of options for the desired output. "simulation" - include
	// simulation output in your response. "validation" - include security validation
	// of the transaction in your response. Default is ["validation"]
	Options param.Field[[]EvmPostTransactionBulkScanParamsOption] `json:"options"`
	// Simulate transactions using gas estimation result. This requires
	// "gas_estimation" option to be enabled.
	SimulateWithEstimatedGas param.Field[bool] `json:"simulate_with_estimated_gas"`
	// Override the state of the chain. This is useful for testing purposes.
	StateOverride param.Field[map[string]EvmPostTransactionBulkScanParamsStateOverride] `json:"state_override"`
}

func (r EvmPostTransactionBulkScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Additional context for the scan (e.g., dapp URL/domain, integration source).
// Used to enrich results and reduce false positives/negatives.
type EvmPostTransactionBulkScanParamsMetadata struct {
	// The full URL of the DApp or website that initiated the transaction, for
	// cross-reference. Must use the https or http scheme and contain a valid hostname.
	// Cannot contain JSON, braces, or other embedded data structures.
	Domain param.Field[string] `json:"domain"`
	// Indicates that the transaction was not initiated by a dapp.
	NonDapp param.Field[bool] `json:"non_dapp"`
}

func (r EvmPostTransactionBulkScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmPostTransactionBulkScanParamsMetadata) implementsEvmPostTransactionBulkScanParamsMetadataUnion() {
}

// Additional context for the scan (e.g., dapp URL/domain, integration source).
// Used to enrich results and reduce false positives/negatives.
//
// Satisfied by
// [EvmPostTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataNonDapp],
// [EvmPostTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataDapp],
// [EvmPostTransactionBulkScanParamsMetadata].
type EvmPostTransactionBulkScanParamsMetadataUnion interface {
	implementsEvmPostTransactionBulkScanParamsMetadataUnion()
}

type EvmPostTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataNonDapp struct {
	// Indicates that the transaction was not initiated by a dapp.
	NonDapp param.Field[EvmPostTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataNonDappNonDapp] `json:"non_dapp"`
}

func (r EvmPostTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataNonDapp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmPostTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataNonDapp) implementsEvmPostTransactionBulkScanParamsMetadataUnion() {
}

// Indicates that the transaction was not initiated by a dapp.
type EvmPostTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataNonDappNonDapp bool

const (
	EvmPostTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataNonDappNonDappTrue EvmPostTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataNonDappNonDapp = true
)

func (r EvmPostTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataNonDappNonDapp) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataNonDappNonDappTrue:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataDapp struct {
	// The full URL of the DApp or website that initiated the transaction, for
	// cross-reference. Must use the https or http scheme and contain a valid hostname.
	// Cannot contain JSON, braces, or other embedded data structures.
	Domain param.Field[string] `json:"domain,required"`
	// Indicates that the transaction was not initiated by a dapp. Use false when the
	// transaction is from a dapp.
	NonDapp param.Field[bool] `json:"non_dapp"`
}

func (r EvmPostTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataDapp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmPostTransactionBulkScanParamsMetadataRoutersEvmModelsMetadataDapp) implementsEvmPostTransactionBulkScanParamsMetadataUnion() {
}

// The relative block for the block validation. Can be "latest" or a block number.
//
// Satisfied by [shared.UnionInt], [shared.UnionString].
type EvmPostTransactionBulkScanParamsBlockUnion interface {
	ImplementsEvmPostTransactionBulkScanParamsBlockUnion()
}

// Response sections to include (e.g., validation, simulation, gas estimation,
// events).
type EvmPostTransactionBulkScanParamsOption string

const (
	EvmPostTransactionBulkScanParamsOptionValidation    EvmPostTransactionBulkScanParamsOption = "validation"
	EvmPostTransactionBulkScanParamsOptionSimulation    EvmPostTransactionBulkScanParamsOption = "simulation"
	EvmPostTransactionBulkScanParamsOptionGasEstimation EvmPostTransactionBulkScanParamsOption = "gas_estimation"
	EvmPostTransactionBulkScanParamsOptionEvents        EvmPostTransactionBulkScanParamsOption = "events"
)

func (r EvmPostTransactionBulkScanParamsOption) IsKnown() bool {
	switch r {
	case EvmPostTransactionBulkScanParamsOptionValidation, EvmPostTransactionBulkScanParamsOptionSimulation, EvmPostTransactionBulkScanParamsOptionGasEstimation, EvmPostTransactionBulkScanParamsOptionEvents:
		return true
	}
	return false
}

type EvmPostTransactionBulkScanParamsStateOverride struct {
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

func (r EvmPostTransactionBulkScanParamsStateOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
