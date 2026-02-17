// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"net/http"
	"reflect"
	"slices"
	"time"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/requestconfig"
	"github.com/blockaid-official/blockaid-client-go/option"
	"github.com/blockaid-official/blockaid-client-go/shared"
	"github.com/tidwall/gjson"
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

// Get a risk recommendation with plain-language reasons for a user operation.
func (r *EvmUserOperationService) Scan(ctx context.Context, body EvmUserOperationScanParams, opts ...option.RequestOption) (res *EvmUserOperationScanResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/evm/user-operation/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EvmUserOperationScanResponse struct {
	Block                      string                                                 `json:"block,required"`
	Chain                      string                                                 `json:"chain,required"`
	AccountAddress             string                                                 `json:"account_address"`
	Events                     []EvmUserOperationScanResponseEvent                    `json:"events"`
	Features                   interface{}                                            `json:"features"`
	GasEstimation              EvmUserOperationScanResponseGasEstimation              `json:"gas_estimation"`
	Simulation                 EvmUserOperationScanResponseSimulation                 `json:"simulation"`
	UserOperationGasEstimation EvmUserOperationScanResponseUserOperationGasEstimation `json:"user_operation_gas_estimation"`
	Validation                 EvmUserOperationScanResponseValidation                 `json:"validation"`
	JSON                       evmUserOperationScanResponseJSON                       `json:"-"`
}

// evmUserOperationScanResponseJSON contains the JSON metadata for the struct
// [EvmUserOperationScanResponse]
type evmUserOperationScanResponseJSON struct {
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

func (r *EvmUserOperationScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseEvent struct {
	Data           string                                    `json:"data,required"`
	EmitterAddress string                                    `json:"emitter_address,required"`
	Topics         []string                                  `json:"topics,required"`
	EmitterName    string                                    `json:"emitter_name"`
	Name           string                                    `json:"name"`
	Params         []EvmUserOperationScanResponseEventsParam `json:"params"`
	JSON           evmUserOperationScanResponseEventJSON     `json:"-"`
}

// evmUserOperationScanResponseEventJSON contains the JSON metadata for the struct
// [EvmUserOperationScanResponseEvent]
type evmUserOperationScanResponseEventJSON struct {
	Data           apijson.Field
	EmitterAddress apijson.Field
	Topics         apijson.Field
	EmitterName    apijson.Field
	Name           apijson.Field
	Params         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseEventJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseEventsParam struct {
	Type         string                                             `json:"type,required"`
	Value        EvmUserOperationScanResponseEventsParamsValueUnion `json:"value,required"`
	InternalType string                                             `json:"internalType"`
	Name         string                                             `json:"name"`
	JSON         evmUserOperationScanResponseEventsParamJSON        `json:"-"`
}

// evmUserOperationScanResponseEventsParamJSON contains the JSON metadata for the
// struct [EvmUserOperationScanResponseEventsParam]
type evmUserOperationScanResponseEventsParamJSON struct {
	Type         apijson.Field
	Value        apijson.Field
	InternalType apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseEventsParam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseEventsParamJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or
// [EvmUserOperationScanResponseEventsParamsValueArray].
type EvmUserOperationScanResponseEventsParamsValueUnion interface {
	ImplementsEvmUserOperationScanResponseEventsParamsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseEventsParamsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseEventsParamsValueArray{}),
		},
	)
}

type EvmUserOperationScanResponseEventsParamsValueArray []interface{}

func (r EvmUserOperationScanResponseEventsParamsValueArray) ImplementsEvmUserOperationScanResponseEventsParamsValueUnion() {
}

type EvmUserOperationScanResponseGasEstimation struct {
	Status   EvmUserOperationScanResponseGasEstimationStatus `json:"status,required"`
	Error    string                                          `json:"error"`
	Estimate string                                          `json:"estimate"`
	Used     string                                          `json:"used"`
	JSON     evmUserOperationScanResponseGasEstimationJSON   `json:"-"`
	union    EvmUserOperationScanResponseGasEstimationUnion
}

// evmUserOperationScanResponseGasEstimationJSON contains the JSON metadata for the
// struct [EvmUserOperationScanResponseGasEstimation]
type evmUserOperationScanResponseGasEstimationJSON struct {
	Status      apijson.Field
	Error       apijson.Field
	Estimate    apijson.Field
	Used        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmUserOperationScanResponseGasEstimationJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseGasEstimation) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseGasEstimation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EvmUserOperationScanResponseGasEstimationUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation],
// [EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError].
func (r EvmUserOperationScanResponseGasEstimation) AsUnion() EvmUserOperationScanResponseGasEstimationUnion {
	return r.union
}

// Union satisfied by
// [EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation]
// or
// [EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError].
type EvmUserOperationScanResponseGasEstimationUnion interface {
	implementsEvmUserOperationScanResponseGasEstimation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseGasEstimationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError{}),
		},
	)
}

type EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation struct {
	Estimate string                                                                                      `json:"estimate,required"`
	Status   EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatus `json:"status,required"`
	Used     string                                                                                      `json:"used,required"`
	JSON     evmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationJSON   `json:"-"`
}

// evmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation]
type evmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationJSON struct {
	Estimate    apijson.Field
	Status      apijson.Field
	Used        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation) implementsEvmUserOperationScanResponseGasEstimation() {
}

type EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatus string

const (
	EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatusSuccess EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatus = "Success"
)

func (r EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatus) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatusSuccess:
		return true
	}
	return false
}

type EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError struct {
	Error  string                                                                                           `json:"error,required"`
	Status EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus `json:"status,required"`
	JSON   evmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON   `json:"-"`
}

// evmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError]
type evmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError) implementsEvmUserOperationScanResponseGasEstimation() {
}

type EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus string

const (
	EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatusError EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus = "Error"
)

func (r EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatusError:
		return true
	}
	return false
}

type EvmUserOperationScanResponseGasEstimationStatus string

const (
	EvmUserOperationScanResponseGasEstimationStatusSuccess EvmUserOperationScanResponseGasEstimationStatus = "Success"
	EvmUserOperationScanResponseGasEstimationStatusError   EvmUserOperationScanResponseGasEstimationStatus = "Error"
)

func (r EvmUserOperationScanResponseGasEstimationStatus) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseGasEstimationStatusSuccess, EvmUserOperationScanResponseGasEstimationStatusError:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulation struct {
	// A string indicating if the simulation was successful or not.
	Status EvmUserOperationScanResponseSimulationStatus `json:"status,required"`
	// This field can have the runtime type of
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummary].
	AccountSummary interface{} `json:"account_summary"`
	// This field can have the runtime type of
	// [map[string]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetail].
	AddressDetails interface{} `json:"address_details"`
	// This field can have the runtime type of
	// [map[string][]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff].
	AssetsDiffs interface{} `json:"assets_diffs"`
	// This field can have the runtime type of
	// [map[string][]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement].
	ContractManagement interface{} `json:"contract_management"`
	// A string explaining why the transaction failed
	Description string `json:"description"`
	// An error message if the simulation failed.
	Error string `json:"error"`
	// This field can have the runtime type of
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails].
	ErrorDetails interface{} `json:"error_details"`
	// This field can have the runtime type of
	// [map[string][]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure].
	Exposures interface{} `json:"exposures"`
	// This field can have the runtime type of
	// [[]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalance].
	MissingBalances interface{} `json:"missing_balances"`
	// This field can have the runtime type of
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationParams].
	Params interface{} `json:"params"`
	// This field can have the runtime type of
	// [map[string][]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKey].
	SessionKey interface{} `json:"session_key"`
	// The number of times the simulation ran until success
	SimulationRunCount int64 `json:"simulation_run_count"`
	// This field can have the runtime type of
	// [map[string]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiff].
	TotalUsdDiff interface{} `json:"total_usd_diff"`
	// This field can have the runtime type of [map[string]map[string]string].
	TotalUsdExposure interface{} `json:"total_usd_exposure"`
	// This field can have the runtime type of
	// [[]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions].
	TransactionActions interface{}                                `json:"transaction_actions"`
	JSON               evmUserOperationScanResponseSimulationJSON `json:"-"`
	union              EvmUserOperationScanResponseSimulationUnion
}

// evmUserOperationScanResponseSimulationJSON contains the JSON metadata for the
// struct [EvmUserOperationScanResponseSimulation]
type evmUserOperationScanResponseSimulationJSON struct {
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

func (r evmUserOperationScanResponseSimulationJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulation) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EvmUserOperationScanResponseSimulationUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulation],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationError].
func (r EvmUserOperationScanResponseSimulation) AsUnion() EvmUserOperationScanResponseSimulationUnion {
	return r.union
}

// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulation]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationError].
type EvmUserOperationScanResponseSimulationUnion interface {
	implementsEvmUserOperationScanResponseSimulation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationError{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulation struct {
	// Account summary for the account address. account address is determined implicit
	// by the `from` field in the transaction request, or explicit by the
	// account_address field in the request.
	AccountSummary EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummary `json:"account_summary,required"`
	// a dictionary including additional information about each relevant address in the
	// transaction.
	AddressDetails map[string]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetail `json:"address_details,required"`
	// dictionary describes the assets differences as a result of this transaction for
	// every involved address
	AssetsDiffs map[string][]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff `json:"assets_diffs,required"`
	// dictionary describes the exposure differences as a result of this transaction
	// for every involved address (as a result of any approval / setApproval / permit
	// function)
	Exposures map[string][]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure `json:"exposures,required"`
	// Session keys created in this transaction per address
	SessionKey map[string][]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKey `json:"session_key,required"`
	// A string indicating if the simulation was successful or not.
	Status EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationStatus `json:"status,required"`
	// dictionary represents the usd value each address gained / lost during this
	// transaction
	TotalUsdDiff map[string]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiff `json:"total_usd_diff,required"`
	// a dictionary representing the usd value each address is exposed to, split by
	// spenders
	TotalUsdExposure map[string]map[string]string `json:"total_usd_exposure,required"`
	// Describes the nature of the transaction and what happened as part of it
	TransactionActions []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions `json:"transaction_actions,required"`
	// Describes the state differences as a result of this transaction for every
	// involved address
	ContractManagement map[string][]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement `json:"contract_management"`
	// Missing balances in the transaction
	MissingBalances []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalance `json:"missing_balances"`
	// The parameters of the transaction that was simulated.
	Params EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationParams `json:"params"`
	// The number of times the simulation ran until success
	SimulationRunCount int64                                                                             `json:"simulation_run_count"`
	JSON               evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulation]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationJSON struct {
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

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulation) implementsEvmUserOperationScanResponseSimulation() {
}

// Account summary for the account address. account address is determined implicit
// by the `from` field in the transaction request, or explicit by the
// account_address field in the request.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummary struct {
	// All assets diffs related to the account address
	AssetsDiffs []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff `json:"assets_diffs,required"`
	// All assets exposures related to the account address
	Exposures []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure `json:"exposures,required"`
	// Total usd diff related to the account address
	TotalUsdDiff EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiff `json:"total_usd_diff,required"`
	// Total usd exposure related to the account address
	TotalUsdExposure map[string]string `json:"total_usd_exposure,required"`
	// All assets traces related to the account address
	Traces []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace `json:"traces,required"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryJSON    `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummary]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryJSON struct {
	AssetsDiffs      apijson.Field
	Exposures        apijson.Field
	TotalUsdDiff     apijson.Field
	TotalUsdExposure apijson.Field
	Traces           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff struct {
	// This field can have the runtime type of
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAsset].
	Asset interface{} `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType `json:"asset_type,required"`
	// This field can have the runtime type of
	// [[]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffIn],
	// [[]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffIn],
	// [[]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffIn],
	// [[]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffIn].
	In interface{} `json:"in,required"`
	// This field can have the runtime type of
	// [[]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOut],
	// [[]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOut],
	// [[]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOut],
	// [[]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOut].
	Out interface{} `json:"out,required"`
	// This field can have the runtime type of
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChanges],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChanges],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChanges],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChanges].
	BalanceChanges interface{}                                                                                               `json:"balance_changes"`
	JSON           evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffJSON `json:"-"`
	union          EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff struct {
	// description of the asset for the current diff
	Asset EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOut `json:"out,required"`
	// shows the balance before making the transaction and after
	BalanceChanges EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChanges `json:"balance_changes"`
	JSON           evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffJSON           `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff() {
}

// description of the asset for the current diff
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType `json:"type,required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                              `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetJSON `json:"-"`
	union  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                     `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                      `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetTypeErc20  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType = "ERC20"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetTypeErc20, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffIn struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                           `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffInJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffInJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffIn]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffInJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOut struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                            `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOutJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOutJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOut]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOutJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOutJSON) RawJSON() string {
	return r.raw
}

// shows the balance before making the transaction and after
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChanges struct {
	// balance of the account after making the transaction
	After EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfter `json:"after,required"`
	// balance of the account before making the transaction
	Before EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBefore `json:"before,required"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesJSON   `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChanges]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChanges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesJSON) RawJSON() string {
	return r.raw
}

// balance of the account after making the transaction
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfter struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                            `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfterJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfterJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfter]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfterJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfterJSON) RawJSON() string {
	return r.raw
}

// balance of the account before making the transaction
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBefore struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                             `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBefore]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff struct {
	// description of the asset for the current diff
	Asset EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOut `json:"out,required"`
	// shows the balance before making the transaction and after
	BalanceChanges EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChanges `json:"balance_changes"`
	JSON           evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffJSON           `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff() {
}

// description of the asset for the current diff
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                               `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetJSON `json:"-"`
	union  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                       `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                       `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetTypeErc721 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType = "ERC721"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetTypeErc721, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffIn struct {
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
	UsdPrice string                                                                                                                                                            `json:"usd_price"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffInJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffInJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffIn]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffInJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOut struct {
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
	UsdPrice string                                                                                                                                                             `json:"usd_price"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOutJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOutJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOut]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOutJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOutJSON) RawJSON() string {
	return r.raw
}

// shows the balance before making the transaction and after
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChanges struct {
	// balance of the account after making the transaction
	After EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfter `json:"after,required"`
	// balance of the account before making the transaction
	Before EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBefore `json:"before,required"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesJSON   `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChanges]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChanges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesJSON) RawJSON() string {
	return r.raw
}

// balance of the account after making the transaction
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfter struct {
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
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfterJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfterJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfter]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfterJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfterJSON) RawJSON() string {
	return r.raw
}

// balance of the account before making the transaction
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBefore struct {
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
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBefore]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff struct {
	// description of the asset for the current diff
	Asset EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOut `json:"out,required"`
	// shows the balance before making the transaction and after
	BalanceChanges EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChanges `json:"balance_changes"`
	JSON           evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffJSON           `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff() {
}

// description of the asset for the current diff
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetJSON `json:"-"`
	union  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                         `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                        `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetTypeErc1155 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType = "ERC1155"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetTypeNonerc  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetTypeErc1155, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffIn struct {
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
	UsdPrice string                                                                                                                                                             `json:"usd_price"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffInJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffInJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffIn]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffInJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOut struct {
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
	UsdPrice string                                                                                                                                                              `json:"usd_price"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOutJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOutJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOut]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOutJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOutJSON) RawJSON() string {
	return r.raw
}

// shows the balance before making the transaction and after
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChanges struct {
	// balance of the account after making the transaction
	After EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfter `json:"after,required"`
	// balance of the account before making the transaction
	Before EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBefore `json:"before,required"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesJSON   `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChanges]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChanges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesJSON) RawJSON() string {
	return r.raw
}

// balance of the account after making the transaction
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfter struct {
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
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfterJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfterJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfter]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfterJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfterJSON) RawJSON() string {
	return r.raw
}

// balance of the account before making the transaction
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBefore struct {
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
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBefore]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff struct {
	// description of the asset for the current diff
	Asset EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOut `json:"out,required"`
	// shows the balance before making the transaction and after
	BalanceChanges EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChanges `json:"balance_changes"`
	JSON           evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffJSON           `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff() {
}

// description of the asset for the current diff
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAsset struct {
	ChainID   int64  `json:"chain_id,required"`
	ChainName string `json:"chain_name,required"`
	Decimals  int64  `json:"decimals,required"`
	LogoURL   string `json:"logo_url,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetType `json:"type,required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                               `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAsset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetJSON struct {
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

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetTypeNative EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetType = "NATIVE"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetTypeNative:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffIn struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                            `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffInJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffInJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffIn]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffInJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOut struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                             `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOutJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOutJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOut]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOutJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOutJSON) RawJSON() string {
	return r.raw
}

// shows the balance before making the transaction and after
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChanges struct {
	// balance of the account after making the transaction
	After EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfter `json:"after,required"`
	// balance of the account before making the transaction
	Before EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBefore `json:"before,required"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesJSON   `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChanges]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChanges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesJSON) RawJSON() string {
	return r.raw
}

// balance of the account after making the transaction
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfter struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                             `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfterJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfterJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfter]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfterJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfterJSON) RawJSON() string {
	return r.raw
}

// balance of the account before making the transaction
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBefore struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                              `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBeforeJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBeforeJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBefore]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBeforeJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBeforeJSON) RawJSON() string {
	return r.raw
}

// type of the asset for the current diff
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc20   EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType = "ERC20"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc721  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType = "ERC721"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc1155 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType = "ERC1155"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeNative  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType = "NATIVE"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc20, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc721, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc1155, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeNative:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure struct {
	// This field can have the runtime type of
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset].
	Asset interface{} `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType `json:"asset_type,required"`
	// This field can have the runtime type of
	// [map[string]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpender],
	// [map[string]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpender],
	// [map[string]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpender].
	Spenders interface{}                                                                                             `json:"spenders,required"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposureJSON `json:"-"`
	union    EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposureJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresUnion {
	return r.union
}

// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType `json:"asset_type,required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpender `json:"spenders,required"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureJSON               `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure() {
}

// description of the asset for the current diff
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType `json:"type,required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                              `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetJSON `json:"-"`
	union  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                     `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                      `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetTypeErc20  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType = "ERC20"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetTypeErc20, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpender struct {
	// the amount that was asked in the approval request for this spender from the
	// current address and asset
	Approval string                                                                                                                                                       `json:"approval,required"`
	Exposure []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure `json:"exposure,required"`
	// the usd price of the approval amount
	ApprovalUsdPrice string `json:"approval_usd_price"`
	// the expiration time of the permit2 protocol
	Expiration time.Time `json:"expiration" format:"date-time"`
	// user friendly description of the approval
	Summary string                                                                                                                                                `json:"summary"`
	JSON    evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpender]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON struct {
	Approval         apijson.Field
	Exposure         apijson.Field
	ApprovalUsdPrice apijson.Field
	Expiration       apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure struct {
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
	Value string                                                                                                                                                         `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON `json:"-"`
	union EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON struct {
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

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
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
	UsdPrice string                                                                                                                                                                                      `json:"usd_price"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
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
	UsdPrice string                                                                                                                                                                                     `json:"usd_price"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                    `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                     `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType `json:"asset_type,required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpender `json:"spenders,required"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureJSON               `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure() {
}

// description of the asset for the current diff
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                               `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetJSON `json:"-"`
	union  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                       `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                       `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetTypeErc721 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType = "ERC721"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetTypeErc721, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpender struct {
	Exposure []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure `json:"exposure,required"`
	// boolean indicates whether an is_approved_for_all function was used (missing in
	// case of ERC20 / ERC1155)
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// user friendly description of the approval
	Summary string                                                                                                                                                 `json:"summary"`
	JSON    evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpender]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure struct {
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
	Value string                                                                                                                                                          `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON `json:"-"`
	union EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON struct {
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

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
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
	UsdPrice string                                                                                                                                                                                       `json:"usd_price"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
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
	UsdPrice string                                                                                                                                                                                      `json:"usd_price"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                     `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                      `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType `json:"asset_type,required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpender `json:"spenders,required"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureJSON               `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure() {
}

// description of the asset for the current diff
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON `json:"-"`
	union  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                         `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                        `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeErc1155 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType = "ERC1155"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeNonerc  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeErc1155, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpender struct {
	Exposure []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure `json:"exposure,required"`
	// boolean indicates whether an is_approved_for_all function was used (missing in
	// case of ERC20 / ERC1155)
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// user friendly description of the approval
	Summary string                                                                                                                                                  `json:"summary"`
	JSON    evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpender]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure struct {
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
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON `json:"-"`
	union EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON struct {
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

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
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
	UsdPrice string                                                                                                                                                                                        `json:"usd_price"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
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
	UsdPrice string                                                                                                                                                                                       `json:"usd_price"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                      `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                       `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

// type of the asset for the current diff
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc20   EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType = "ERC20"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc721  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType = "ERC721"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc1155 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType = "ERC1155"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc20, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc721, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc1155:
		return true
	}
	return false
}

// Total usd diff related to the account address
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiff struct {
	In    string                                                                                                      `json:"in,required"`
	Out   string                                                                                                      `json:"out,required"`
	Total string                                                                                                      `json:"total,required"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace struct {
	// This field can have the runtime type of
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAsset],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset].
	Asset interface{} `json:"asset,required"`
	// type of the trace
	TraceType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceType `json:"trace_type,required"`
	// The type of the model
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType `json:"type,required"`
	// This field can have the runtime type of
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiff],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiff],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiff],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiff].
	Diff interface{} `json:"diff"`
	// This field can have the runtime type of
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposed],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposed].
	Exposed interface{} `json:"exposed"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address"`
	// This field can have the runtime type of
	// [[]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabels],
	// [[]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabels],
	// [[]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabels],
	// [[]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabels].
	Labels interface{} `json:"labels"`
	// The owner of the assets
	Owner string `json:"owner"`
	// The spender of the assets
	Spender string `json:"spender"`
	// The address where the assets are moved to
	ToAddress string                                                                                               `json:"to_address"`
	JSON      evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTraceJSON `json:"-"`
	union     EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTraceJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTraceJSON struct {
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

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTraceJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesUnion {
	return r.union
}

// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace struct {
	// Description of the asset in the trace
	Asset EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset `json:"asset,required"`
	// The difference in value for the asset in the trace
	Diff EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiff `json:"diff,required"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address,required"`
	// The address where the assets are moved to
	ToAddress string `json:"to_address,required"`
	// type of the trace
	TraceType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceType `json:"type,required"`
	// List of labels that describe the trace
	Labels []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabels `json:"labels"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceJSON     `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceJSON struct {
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

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetType `json:"type,required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                      `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetJSON `json:"-"`
	union  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                             `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                              `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetTypeErc20  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetType = "ERC20"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetTypeErc20, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetTypeNonerc:
		return true
	}
	return false
}

// The difference in value for the asset in the trace
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                     `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiffJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceTypeAssetTrace EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceType = "AssetTrace"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceTypeAssetTrace:
		return true
	}
	return false
}

// The type of the model
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTypeErc20AssetTrace EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceType = "ERC20AssetTrace"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTypeErc20AssetTrace:
		return true
	}
	return false
}

// Label assigned to a trace entry.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabels string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabelsGasFee EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabels = "GAS_FEE"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabels) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabelsGasFee:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace struct {
	// Description of the asset in the trace
	Asset EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset `json:"asset,required"`
	// The difference in value for the asset in the trace
	Diff EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiff `json:"diff,required"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address,required"`
	// The address where the assets are moved to
	ToAddress string `json:"to_address,required"`
	// type of the trace
	TraceType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceType `json:"type,required"`
	// List of labels that describe the trace
	Labels []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabels `json:"labels"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceJSON     `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceJSON struct {
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

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                       `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetJSON `json:"-"`
	union  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                               `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                               `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetTypeErc721 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetType = "ERC721"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetTypeErc721, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetTypeNonerc:
		return true
	}
	return false
}

// The difference in value for the asset in the trace
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiff struct {
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
	UsdPrice string                                                                                                                                      `json:"usd_price"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiffJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceTypeAssetTrace EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceType = "AssetTrace"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceTypeAssetTrace:
		return true
	}
	return false
}

// The type of the model
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTypeErc721AssetTrace EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceType = "ERC721AssetTrace"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTypeErc721AssetTrace:
		return true
	}
	return false
}

// Label assigned to a trace entry.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabels string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabelsGasFee EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabels = "GAS_FEE"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabels) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabelsGasFee:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace struct {
	// Description of the asset in the trace
	Asset EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset `json:"asset,required"`
	// The difference in value for the asset in the trace
	Diff EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiff `json:"diff,required"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address,required"`
	// The address where the assets are moved to
	ToAddress string `json:"to_address,required"`
	// type of the trace
	TraceType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceType `json:"type,required"`
	// List of labels that describe the trace
	Labels []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabels `json:"labels"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceJSON     `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceJSON struct {
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

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                        `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetJSON `json:"-"`
	union  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                 `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetTypeErc1155 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetType = "ERC1155"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetTypeNonerc  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetTypeErc1155, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetTypeNonerc:
		return true
	}
	return false
}

// The difference in value for the asset in the trace
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiff struct {
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
	UsdPrice string                                                                                                                                       `json:"usd_price"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiffJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceTypeAssetTrace EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceType = "AssetTrace"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceTypeAssetTrace:
		return true
	}
	return false
}

// The type of the model
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTypeErc1155AssetTrace EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceType = "ERC1155AssetTrace"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTypeErc1155AssetTrace:
		return true
	}
	return false
}

// Label assigned to a trace entry.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabels string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabelsGasFee EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabels = "GAS_FEE"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabels) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabelsGasFee:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace struct {
	// Description of the asset in the trace
	Asset EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAsset `json:"asset,required"`
	// The difference in value for the asset in the trace
	Diff EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiff `json:"diff,required"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address,required"`
	// The address where the assets are moved to
	ToAddress string `json:"to_address,required"`
	// type of the trace
	TraceType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceType `json:"type,required"`
	// List of labels that describe the trace
	Labels []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabels `json:"labels"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceJSON     `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceJSON struct {
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

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAsset struct {
	ChainID   int64  `json:"chain_id,required"`
	ChainName string `json:"chain_name,required"`
	Decimals  int64  `json:"decimals,required"`
	LogoURL   string `json:"logo_url,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetType `json:"type,required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                       `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAsset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetJSON struct {
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

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetTypeNative EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetType = "NATIVE"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetTypeNative:
		return true
	}
	return false
}

// The difference in value for the asset in the trace
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                      `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiffJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceTypeAssetTrace EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceType = "AssetTrace"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceTypeAssetTrace:
		return true
	}
	return false
}

// The type of the model
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTypeNativeAssetTrace EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceType = "NativeAssetTrace"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTypeNativeAssetTrace:
		return true
	}
	return false
}

// Label assigned to a trace entry.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabels string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabelsGasFee EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabels = "GAS_FEE"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabels) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabelsGasFee:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace struct {
	// Description of the asset in the trace
	Asset   EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset   `json:"asset,required"`
	Exposed EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposed `json:"exposed,required"`
	// The owner of the assets
	Owner string `json:"owner,required"`
	// The spender of the assets
	Spender string `json:"spender,required"`
	// type of the trace
	TraceType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceType `json:"type,required"`
	JSON evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceJSON struct {
	Asset       apijson.Field
	Exposed     apijson.Field
	Owner       apijson.Field
	Spender     apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetType `json:"type,required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                         `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetJSON `json:"-"`
	union  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                 `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetTypeErc20  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetType = "ERC20"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetTypeErc20, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetTypeNonerc:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposed struct {
	RawValue string                                                                                                                                           `json:"raw_value,required"`
	UsdPrice float64                                                                                                                                          `json:"usd_price"`
	Value    float64                                                                                                                                          `json:"value"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposedJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposedJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposed]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposedJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposedJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceTypeExposureTrace EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceType = "ExposureTrace"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceTypeExposureTrace:
		return true
	}
	return false
}

// The type of the model
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTypeErc20ExposureTrace EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceType = "ERC20ExposureTrace"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTypeErc20ExposureTrace:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace struct {
	// Description of the asset in the trace
	Asset EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset `json:"asset,required"`
	// The owner of the assets
	Owner string `json:"owner,required"`
	// The spender of the assets
	Spender string `json:"spender,required"`
	// type of the trace
	TraceType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type    EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceType    `json:"type,required"`
	Exposed EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposed `json:"exposed"`
	JSON    evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceJSON    `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceJSON struct {
	Asset       apijson.Field
	Owner       apijson.Field
	Spender     apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	Exposed     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                          `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetJSON `json:"-"`
	union  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                  `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                  `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetTypeErc721 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetType = "ERC721"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetTypeErc721, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetTypeNonerc:
		return true
	}
	return false
}

// type of the trace
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceTypeExposureTrace EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceType = "ExposureTrace"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceTypeExposureTrace:
		return true
	}
	return false
}

// The type of the model
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTypeErc721ExposureTrace EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceType = "ERC721ExposureTrace"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTypeErc721ExposureTrace:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposed struct {
	Amount   int64                                                                                                                                             `json:"amount,required"`
	TokenID  string                                                                                                                                            `json:"token_id,required"`
	IsMint   bool                                                                                                                                              `json:"is_mint"`
	LogoURL  string                                                                                                                                            `json:"logo_url"`
	UsdPrice float64                                                                                                                                           `json:"usd_price"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposedJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposedJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposed]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposedJSON struct {
	Amount      apijson.Field
	TokenID     apijson.Field
	IsMint      apijson.Field
	LogoURL     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposedJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace struct {
	// Description of the asset in the trace
	Asset EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset `json:"asset,required"`
	// The owner of the assets
	Owner string `json:"owner,required"`
	// The spender of the assets
	Spender string `json:"spender,required"`
	// type of the trace
	TraceType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceType `json:"type,required"`
	JSON evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceJSON struct {
	Asset       apijson.Field
	Owner       apijson.Field
	Spender     apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                           `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetJSON `json:"-"`
	union  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                    `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                   `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetTypeErc1155 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetType = "ERC1155"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetTypeNonerc  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetTypeErc1155, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetTypeNonerc:
		return true
	}
	return false
}

// type of the trace
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceTypeExposureTrace EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceType = "ExposureTrace"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceTypeExposureTrace:
		return true
	}
	return false
}

// The type of the model
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTypeErc1155ExposureTrace EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceType = "ERC1155ExposureTrace"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTypeErc1155ExposureTrace:
		return true
	}
	return false
}

// type of the trace
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceTypeAssetTrace    EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceType = "AssetTrace"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceTypeExposureTrace EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceType = "ExposureTrace"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceTypeAssetTrace, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceTypeExposureTrace:
		return true
	}
	return false
}

// The type of the model
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc20AssetTrace      EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC20AssetTrace"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc721AssetTrace     EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC721AssetTrace"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc1155AssetTrace    EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC1155AssetTrace"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeNativeAssetTrace     EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "NativeAssetTrace"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc20ExposureTrace   EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC20ExposureTrace"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc721ExposureTrace  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC721ExposureTrace"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc1155ExposureTrace EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC1155ExposureTrace"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc20AssetTrace, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc721AssetTrace, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc1155AssetTrace, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeNativeAssetTrace, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc20ExposureTrace, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc721ExposureTrace, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc1155ExposureTrace:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetail struct {
	// contains the contract's name if the address is a verified contract
	ContractName string `json:"contract_name"`
	// Whether the address is an eoa or a contract
	IsEoa bool `json:"is_eoa"`
	// known name tag for the address
	NameTag string                                                                                         `json:"name_tag"`
	JSON    evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetailJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetailJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetail]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetailJSON struct {
	ContractName apijson.Field
	IsEoa        apijson.Field
	NameTag      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetailJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff struct {
	// This field can have the runtime type of
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAsset].
	Asset interface{} `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType `json:"asset_type,required"`
	// This field can have the runtime type of
	// [[]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffIn],
	// [[]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffIn],
	// [[]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffIn],
	// [[]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffIn].
	In interface{} `json:"in,required"`
	// This field can have the runtime type of
	// [[]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOut],
	// [[]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOut],
	// [[]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOut],
	// [[]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOut].
	Out   interface{}                                                                                 `json:"out,required"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffJSON `json:"-"`
	union EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff struct {
	// description of the asset for the current diff
	Asset EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOut `json:"out,required"`
	JSON evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffJSON  `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff() {
}

// description of the asset for the current diff
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType `json:"type,required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                   `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetJSON `json:"-"`
	union  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                          `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                           `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetTypeErc20  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType = "ERC20"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetTypeErc20, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffIn struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffInJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffInJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffIn]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffInJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOut struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                 `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOutJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOutJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOut]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff struct {
	// description of the asset for the current diff
	Asset EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOut `json:"out,required"`
	JSON evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffJSON  `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff() {
}

// description of the asset for the current diff
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                    `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetJSON `json:"-"`
	union  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                            `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                            `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetTypeErc721 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType = "ERC721"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetTypeErc721, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffIn struct {
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
	UsdPrice string                                                                                                                                 `json:"usd_price"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffInJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffInJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffIn]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffInJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOut struct {
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
	UsdPrice string                                                                                                                                  `json:"usd_price"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOutJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOutJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOut]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOutJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff struct {
	// description of the asset for the current diff
	Asset EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOut `json:"out,required"`
	JSON evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffJSON  `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff() {
}

// description of the asset for the current diff
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                     `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetJSON `json:"-"`
	union  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                              `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                             `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetTypeErc1155 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType = "ERC1155"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetTypeNonerc  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetTypeErc1155, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffIn struct {
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
	UsdPrice string                                                                                                                                  `json:"usd_price"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffInJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffInJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffIn]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffInJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOut struct {
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
	UsdPrice string                                                                                                                                   `json:"usd_price"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOutJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOutJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOut]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOutJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff struct {
	// description of the asset for the current diff
	Asset EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOut `json:"out,required"`
	JSON evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffJSON  `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff() {
}

// description of the asset for the current diff
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAsset struct {
	ChainID   int64  `json:"chain_id,required"`
	ChainName string `json:"chain_name,required"`
	Decimals  int64  `json:"decimals,required"`
	LogoURL   string `json:"logo_url,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetType `json:"type,required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                    `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAsset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetJSON struct {
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

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetTypeNative EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetType = "NATIVE"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetTypeNative:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffIn struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                 `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffInJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffInJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffIn]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffInJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOut struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                  `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOutJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOutJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOut]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

// type of the asset for the current diff
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc20   EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType = "ERC20"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc721  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType = "ERC721"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc1155 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType = "ERC1155"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeNative  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType = "NATIVE"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc20, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc721, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc1155, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeNative:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure struct {
	// This field can have the runtime type of
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset].
	Asset interface{} `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType `json:"asset_type,required"`
	// This field can have the runtime type of
	// [map[string]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpender],
	// [map[string]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpender],
	// [map[string]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpender].
	Spenders interface{}                                                                               `json:"spenders,required"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposureJSON `json:"-"`
	union    EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposureJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresUnion {
	return r.union
}

// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType `json:"asset_type,required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpender `json:"spenders,required"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureJSON               `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure() {
}

// description of the asset for the current diff
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType `json:"type,required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetJSON `json:"-"`
	union  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                       `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                        `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetTypeErc20  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType = "ERC20"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetTypeErc20, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpender struct {
	// the amount that was asked in the approval request for this spender from the
	// current address and asset
	Approval string                                                                                                                                         `json:"approval,required"`
	Exposure []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure `json:"exposure,required"`
	// the usd price of the approval amount
	ApprovalUsdPrice string `json:"approval_usd_price"`
	// the expiration time of the permit2 protocol
	Expiration time.Time `json:"expiration" format:"date-time"`
	// user friendly description of the approval
	Summary string                                                                                                                                  `json:"summary"`
	JSON    evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpender]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON struct {
	Approval         apijson.Field
	Exposure         apijson.Field
	ApprovalUsdPrice apijson.Field
	Expiration       apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure struct {
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
	Value string                                                                                                                                           `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON `json:"-"`
	union EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON struct {
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

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
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
	UsdPrice string                                                                                                                                                                        `json:"usd_price"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
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
	UsdPrice string                                                                                                                                                                       `json:"usd_price"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                      `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                       `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType `json:"asset_type,required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpender `json:"spenders,required"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureJSON               `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure() {
}

// description of the asset for the current diff
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                 `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetJSON `json:"-"`
	union  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                         `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                         `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetTypeErc721 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType = "ERC721"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetTypeErc721, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpender struct {
	Exposure []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure `json:"exposure,required"`
	// boolean indicates whether an is_approved_for_all function was used (missing in
	// case of ERC20 / ERC1155)
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// user friendly description of the approval
	Summary string                                                                                                                                   `json:"summary"`
	JSON    evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpender]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure struct {
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
	Value string                                                                                                                                            `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON `json:"-"`
	union EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON struct {
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

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
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
	UsdPrice string                                                                                                                                                                         `json:"usd_price"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
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
	UsdPrice string                                                                                                                                                                        `json:"usd_price"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                       `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                        `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType `json:"asset_type,required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpender `json:"spenders,required"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureJSON               `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure() {
}

// description of the asset for the current diff
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                  `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON `json:"-"`
	union  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                           `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                          `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeErc1155 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType = "ERC1155"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeNonerc  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType = "NONERC"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeErc1155, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpender struct {
	Exposure []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure `json:"exposure,required"`
	// boolean indicates whether an is_approved_for_all function was used (missing in
	// case of ERC20 / ERC1155)
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// user friendly description of the approval
	Summary string                                                                                                                                    `json:"summary"`
	JSON    evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpender]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure struct {
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
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON `json:"-"`
	union EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON struct {
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

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
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
	UsdPrice string                                                                                                                                                                          `json:"usd_price"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
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
	UsdPrice string                                                                                                                                                                         `json:"usd_price"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                        `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                         `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

// type of the asset for the current diff
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc20   EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType = "ERC20"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc721  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType = "ERC721"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc1155 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType = "ERC1155"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc20, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc721, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc1155:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKey struct {
	Key      string                                                                                          `json:"key,required"`
	Policies []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy `json:"policies,required"`
	Signer   string                                                                                          `json:"signer,required"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyJSON     `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKey]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyJSON struct {
	Key         apijson.Field
	Policies    apijson.Field
	Signer      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy struct {
	// This field can have the runtime type of
	// [[]EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArg].
	Args interface{} `json:"args"`
	// This field can have the runtime type of
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails].
	AssetDetails interface{}                                                                                         `json:"asset_details"`
	EndTime      time.Time                                                                                           `json:"end_time" format:"date-time"`
	Method       string                                                                                              `json:"method"`
	Recipient    string                                                                                              `json:"recipient"`
	StartTime    time.Time                                                                                           `json:"start_time" format:"date-time"`
	ToAddress    string                                                                                              `json:"to_address"`
	Type         EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType `json:"type"`
	ValueLimit   string                                                                                              `json:"value_limit"`
	JSON         evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicyJSON   `json:"-"`
	union        EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicyJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicyJSON struct {
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

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicyJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesUnion {
	return r.union
}

// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy struct {
	AssetDetails EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails `json:"asset_details,required"`
	Recipient    string                                                                                                                                         `json:"recipient"`
	Type         EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyType         `json:"type"`
	ValueLimit   string                                                                                                                                         `json:"value_limit"`
	JSON         evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyJSON         `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyJSON struct {
	AssetDetails apijson.Field
	Recipient    apijson.Field
	Type         apijson.Field
	ValueLimit   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails struct {
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsType `json:"type,required"`
	// address of the token
	Address   string `json:"address"`
	ChainID   int64  `json:"chain_id"`
	ChainName string `json:"chain_name"`
	LogoURL   string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                             `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsJSON `json:"-"`
	union  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsJSON struct {
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

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsUnion {
	return r.union
}

// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails struct {
	ChainID   int64  `json:"chain_id,required"`
	ChainName string `json:"chain_name,required"`
	Decimals  int64  `json:"decimals,required"`
	LogoURL   string `json:"logo_url,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsType `json:"type,required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                     `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsJSON struct {
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

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsTypeNative EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsType = "NATIVE"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsTypeNative:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                    `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsTypeNative EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsType = "NATIVE"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsTypeErc20  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsType = "ERC20"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsTypeNative, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsTypeErc20:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyTypeTransferPolicy EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyType = "TRANSFER_POLICY"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyTypeTransferPolicy:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy struct {
	ToAddress  string                                                                                                                              `json:"to_address,required"`
	Args       []EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArg `json:"args"`
	Method     string                                                                                                                              `json:"method"`
	Type       EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyType  `json:"type"`
	ValueLimit string                                                                                                                              `json:"value_limit"`
	JSON       evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyJSON  `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyJSON struct {
	ToAddress   apijson.Field
	Args        apijson.Field
	Method      apijson.Field
	Type        apijson.Field
	ValueLimit  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArg struct {
	// Comparison operator used to evaluate an argument/value against a policy
	// constraint.
	Condition EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition `json:"condition,required"`
	Index     int64                                                                                                                                       `json:"index,required"`
	Value     string                                                                                                                                      `json:"value,required"`
	JSON      evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgJSON       `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArg]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgJSON struct {
	Condition   apijson.Field
	Index       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgJSON) RawJSON() string {
	return r.raw
}

// Comparison operator used to evaluate an argument/value against a policy
// constraint.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionUnconstrained  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "UNCONSTRAINED"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionEqual          EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "EQUAL"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionGreater        EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "GREATER"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionLess           EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "LESS"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionGreaterOrEqual EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "GREATER_OR_EQUAL"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionLessOrEqual    EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "LESS_OR_EQUAL"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionNotEqual       EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "NOT_EQUAL"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionUnconstrained, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionEqual, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionGreater, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionLess, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionGreaterOrEqual, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionLessOrEqual, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionNotEqual:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyTypeCallPolicy EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyType = "CALL_POLICY"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyTypeCallPolicy:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy struct {
	ValueLimit string                                                                                                                            `json:"value_limit,required"`
	Type       EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyType `json:"type"`
	JSON       evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyJSON struct {
	ValueLimit  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyTypeGasPolicy EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyType = "GAS_POLICY"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyTypeGasPolicy:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy struct {
	EndTime   time.Time                                                                                                                                `json:"end_time" format:"date-time"`
	StartTime time.Time                                                                                                                                `json:"start_time" format:"date-time"`
	Type      EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyType `json:"type"`
	JSON      evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyTypeExpirationPolicy EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyType = "EXPIRATION_POLICY"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyTypeExpirationPolicy:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeTransferPolicy   EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType = "TRANSFER_POLICY"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeCallPolicy       EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType = "CALL_POLICY"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeGasPolicy        EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType = "GAS_POLICY"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeExpirationPolicy EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType = "EXPIRATION_POLICY"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeTransferPolicy, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeCallPolicy, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeGasPolicy, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeExpirationPolicy:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationStatus string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationStatusSuccess EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationStatus = "Success"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationStatus) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationStatusSuccess:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiff struct {
	In    string                                                                                        `json:"in,required"`
	Out   string                                                                                        `json:"out,required"`
	Total string                                                                                        `json:"total,required"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiffJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiffJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiff]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

// Type of action performed in the transaction.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsMint            EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "mint"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsStake           EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "stake"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsSwap            EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "swap"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsNativeTransfer  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "native_transfer"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsTokenTransfer   EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "token_transfer"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsApproval        EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "approval"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsSetCodeAccount  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "set_code_account"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsProxyUpgrade    EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "proxy_upgrade"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsOwnershipChange EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "ownership_change"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsMint, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsStake, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsSwap, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsNativeTransfer, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsTokenTransfer, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsApproval, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsSetCodeAccount, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsProxyUpgrade, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsOwnershipChange:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement struct {
	// The type of the state change
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType `json:"type,required"`
	// This field can have the runtime type of
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfter],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfter],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfter].
	After interface{} `json:"after"`
	// This field can have the runtime type of
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBefore],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBefore],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBefore].
	Before interface{} `json:"before"`
	// The delegated address
	DelegatedAddress string `json:"delegated_address"`
	// The direct creator address of the new contract
	Deployer string                                                                                              `json:"deployer"`
	JSON     evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementJSON `json:"-"`
	union    EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementJSON struct {
	Type             apijson.Field
	After            apijson.Field
	Before           apijson.Field
	DelegatedAddress apijson.Field
	Deployer         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementUnion {
	return r.union
}

// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement struct {
	// The state after the transaction
	After EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfter `json:"after,required"`
	// The state before the transaction
	Before EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBefore `json:"before,required"`
	// The type of the state change
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementType `json:"type,required"`
	JSON evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementJSON struct {
	After       apijson.Field
	Before      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement() {
}

// The state after the transaction
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfter struct {
	Address string                                                                                                                                           `json:"address,required"`
	JSON    evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfterJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfterJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfter]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfterJSON struct {
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfterJSON) RawJSON() string {
	return r.raw
}

// The state before the transaction
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBefore struct {
	Address string                                                                                                                                            `json:"address,required"`
	JSON    evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBeforeJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBeforeJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBefore]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBeforeJSON struct {
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBeforeJSON) RawJSON() string {
	return r.raw
}

// The type of the state change
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementTypeProxyUpgrade EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementType = "PROXY_UPGRADE"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementTypeProxyUpgrade:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement struct {
	// The state after the transaction
	After EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfter `json:"after,required"`
	// The state before the transaction
	Before EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBefore `json:"before,required"`
	// The type of the state change
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementType `json:"type,required"`
	JSON evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementJSON struct {
	After       apijson.Field
	Before      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement() {
}

// The state after the transaction
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfter struct {
	Owners []string                                                                                                                                            `json:"owners,required"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfterJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfterJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfter]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfterJSON struct {
	Owners      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfterJSON) RawJSON() string {
	return r.raw
}

// The state before the transaction
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBefore struct {
	Owners []string                                                                                                                                             `json:"owners,required"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBeforeJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBeforeJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBefore]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBeforeJSON struct {
	Owners      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBeforeJSON) RawJSON() string {
	return r.raw
}

// The type of the state change
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementTypeOwnershipChange EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementType = "OWNERSHIP_CHANGE"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementTypeOwnershipChange:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement struct {
	// The state after the transaction
	After EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfter `json:"after,required"`
	// The state before the transaction
	Before EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBefore `json:"before,required"`
	// The type of the state change
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementType `json:"type,required"`
	JSON evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementJSON struct {
	After       apijson.Field
	Before      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement() {
}

// The state after the transaction
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfter struct {
	Modules []string                                                                                                                                          `json:"modules,required"`
	JSON    evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfterJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfterJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfter]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfterJSON struct {
	Modules     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfterJSON) RawJSON() string {
	return r.raw
}

// The state before the transaction
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBefore struct {
	Modules []string                                                                                                                                           `json:"modules,required"`
	JSON    evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBeforeJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBeforeJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBefore]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBeforeJSON struct {
	Modules     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBeforeJSON) RawJSON() string {
	return r.raw
}

// The type of the state change
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementTypeModuleChange EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementType = "MODULE_CHANGE"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementTypeModuleChange:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement struct {
	// The delegated address
	DelegatedAddress string `json:"delegated_address,required"`
	// The type of the state change
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementType `json:"type,required"`
	JSON evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementJSON struct {
	DelegatedAddress apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement() {
}

// The type of the state change
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementTypeSetCodeAccount EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementType = "SET_CODE_ACCOUNT"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementTypeSetCodeAccount:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation struct {
	// The direct creator address of the new contract
	Deployer string `json:"deployer,required"`
	// The type of the state change
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationType `json:"type,required"`
	JSON evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationJSON struct {
	Deployer    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement() {
}

// The type of the state change
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationTypeContractCreation EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationType = "CONTRACT_CREATION"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationTypeContractCreation:
		return true
	}
	return false
}

// The type of the state change
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeProxyUpgrade     EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType = "PROXY_UPGRADE"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeOwnershipChange  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType = "OWNERSHIP_CHANGE"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeModuleChange     EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType = "MODULE_CHANGE"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeSetCodeAccount   EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType = "SET_CODE_ACCOUNT"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeContractCreation EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType = "CONTRACT_CREATION"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeProxyUpgrade, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeOwnershipChange, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeModuleChange, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeSetCodeAccount, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeContractCreation:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalance struct {
	// The asset that is missing balance
	Asset EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset `json:"asset,required"`
	// The account address's current balance of the asset
	CurrentBalance EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalance `json:"current_balance,required"`
	// The account address's missing balance of the asset
	MissingBalance EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalance `json:"missing_balance,required"`
	// The required balance of the asset for this action
	RequiredBalance EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalance `json:"required_balance,required"`
	JSON            evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalanceJSON             `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalanceJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalance]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalanceJSON struct {
	Asset           apijson.Field
	CurrentBalance  apijson.Field
	MissingBalance  apijson.Field
	RequiredBalance apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalanceJSON) RawJSON() string {
	return r.raw
}

// The asset that is missing balance
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset struct {
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetType `json:"type,required"`
	// address of the token
	Address   string `json:"address"`
	ChainID   int64  `json:"chain_id"`
	ChainName string `json:"chain_name"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetJSON `json:"-"`
	union  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetJSON struct {
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

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetUnion {
	return r.union
}

// The asset that is missing balance
//
// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                       `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails struct {
	ChainID   int64  `json:"chain_id,required"`
	ChainName string `json:"chain_name,required"`
	Decimals  int64  `json:"decimals,required"`
	LogoURL   string `json:"logo_url,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsType `json:"type,required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                        `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsJSON struct {
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

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset() {
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsTypeNative EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsType = "NATIVE"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsTypeNative:
		return true
	}
	return false
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetTypeErc20  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetType = "ERC20"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetTypeNative EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetType = "NATIVE"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetTypeErc20, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetTypeNative:
		return true
	}
	return false
}

// The account address's current balance of the asset
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalance struct {
	// The raw value of the balance in hex string format
	RawValue string `json:"raw_value,required"`
	// The value of the balance in decimal string format
	Value string                                                                                                         `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalanceJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalanceJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalance]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalanceJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalanceJSON) RawJSON() string {
	return r.raw
}

// The account address's missing balance of the asset
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalance struct {
	// The raw value of the balance in hex string format
	RawValue string `json:"raw_value,required"`
	// The value of the balance in decimal string format
	Value string                                                                                                         `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalanceJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalanceJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalance]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalanceJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalanceJSON) RawJSON() string {
	return r.raw
}

// The required balance of the asset for this action
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalance struct {
	// The raw value of the balance in hex string format
	RawValue string `json:"raw_value,required"`
	// The value of the balance in decimal string format
	Value string                                                                                                          `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalanceJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalanceJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalance]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalanceJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalanceJSON) RawJSON() string {
	return r.raw
}

// The parameters of the transaction that was simulated.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationParams struct {
	// The block tag to be sent.
	BlockTag string `json:"block_tag"`
	// The calldata to be sent.
	Calldata EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldata `json:"calldata"`
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
	UserOperationCalldata EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldata `json:"user_operation_calldata"`
	// The value to be sent.
	Value string                                                                                  `json:"value"`
	JSON  evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationParams]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsJSON struct {
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

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsJSON) RawJSON() string {
	return r.raw
}

// The calldata to be sent.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldata struct {
	// The function selector of the function called in the transaction
	FunctionSelector string `json:"function_selector,required"`
	// The function declaration of the function called in the transaction
	FunctionDeclaration string `json:"function_declaration"`
	// The function signature of the function called in the transaction
	FunctionSignature string                                                                                          `json:"function_signature"`
	JSON              evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldataJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldataJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldata]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldataJSON struct {
	FunctionSelector    apijson.Field
	FunctionDeclaration apijson.Field
	FunctionSignature   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldataJSON) RawJSON() string {
	return r.raw
}

// The user operation call data to be sent.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldata struct {
	// The function selector of the function called in the transaction
	FunctionSelector string `json:"function_selector,required"`
	// The function declaration of the function called in the transaction
	FunctionDeclaration string `json:"function_declaration"`
	// The function signature of the function called in the transaction
	FunctionSignature string                                                                                                       `json:"function_signature"`
	JSON              evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldataJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldataJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldata]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldataJSON struct {
	FunctionSelector    apijson.Field
	FunctionDeclaration apijson.Field
	FunctionSignature   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldataJSON) RawJSON() string {
	return r.raw
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationError struct {
	// A string explaining why the transaction failed
	Description string `json:"description,required"`
	// An error message if the simulation failed.
	Error string `json:"error,required"`
	// A string indicating if the simulation was successful or not.
	Status EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatus `json:"status,required"`
	// Error details if the simulation failed.
	ErrorDetails EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails `json:"error_details"`
	JSON         evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorJSON         `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationError]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorJSON struct {
	Description  apijson.Field
	Error        apijson.Field
	Status       apijson.Field
	ErrorDetails apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationError) implementsEvmUserOperationScanResponseSimulation() {
}

// A string indicating if the simulation was successful or not.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatus string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatusError EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatus = "Error"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatus) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatusError:
		return true
	}
	return false
}

// Error details if the simulation failed.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails struct {
	// The type of the model
	Code string `json:"code,required"`
	// The address of the account
	AccountAddress string `json:"account_address"`
	// The address that is invalid
	Address string `json:"address"`
	// This field can have the runtime type of
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset].
	Asset interface{} `json:"asset"`
	// The current balance of the account
	CurrentBalance string `json:"current_balance"`
	// The domain name that is unsupported
	DomainName string `json:"domain_name"`
	// The message type that is unsupported
	MessageType string `json:"message_type"`
	// The required balance of the account
	RequiredBalance string                                                                                             `json:"required_balance"`
	JSON            evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsJSON `json:"-"`
	union           EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsJSON struct {
	Code            apijson.Field
	AccountAddress  apijson.Field
	Address         apijson.Field
	Asset           apijson.Field
	CurrentBalance  apijson.Field
	DomainName      apijson.Field
	MessageType     apijson.Field
	RequiredBalance apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsUnion {
	return r.union
}

// Error details if the simulation failed.
//
// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails struct {
	// The address of the account
	AccountAddress string `json:"account_address,required"`
	// The asset that the account does not have enough balance for
	Asset EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset `json:"asset,required"`
	// The type of the model
	Code EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCode `json:"code,required"`
	// The current balance of the account
	CurrentBalance string `json:"current_balance"`
	// The required balance of the account
	RequiredBalance string                                                                                                                                                   `json:"required_balance"`
	JSON            evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsJSON struct {
	AccountAddress  apijson.Field
	Asset           apijson.Field
	Code            apijson.Field
	CurrentBalance  apijson.Field
	RequiredBalance apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails() {
}

// The asset that the account does not have enough balance for
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset struct {
	// This field can have the runtime type of
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetails],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetails],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetails],
	// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetails].
	Details interface{} `json:"details,required"`
	// The type of the model
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType `json:"type,required"`
	// Token Id
	TokenID int64                                                                                                                                                         `json:"token_id"`
	JSON    evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetJSON `json:"-"`
	union   EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetUnion
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetJSON struct {
	Details     apijson.Field
	Type        apijson.Field
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset].
func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset) AsUnion() EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetUnion {
	return r.union
}

// The asset that the account does not have enough balance for
//
// Union satisfied by
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset],
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset]
// or
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset].
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetUnion interface {
	implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset{}),
		},
	)
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset struct {
	// Details
	Details EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetails `json:"details,required"`
	// The type of the model
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetType `json:"type,required"`
	JSON evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetJSON struct {
	Details     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset() {
}

// Details
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetails struct {
	ChainID   int64  `json:"chain_id,required"`
	ChainName string `json:"chain_name,required"`
	Decimals  int64  `json:"decimals,required"`
	LogoURL   string `json:"logo_url,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsType `json:"type,required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                            `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsJSON struct {
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

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsTypeNative EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsType = "NATIVE"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsTypeNative:
		return true
	}
	return false
}

// The type of the model
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetTypeNative EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetType = "NATIVE"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetTypeNative:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset struct {
	// Details
	Details EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetails `json:"details,required"`
	// The type of the model
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetType `json:"type,required"`
	JSON evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetJSON struct {
	Details     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset() {
}

// Details
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                           `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsTypeErc20 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsType = "ERC20"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsTypeErc20:
		return true
	}
	return false
}

// The type of the model
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetTypeErc20 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetType = "ERC20"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetTypeErc20:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset struct {
	// Details
	Details EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetails `json:"details,required"`
	// Token Id
	TokenID int64 `json:"token_id,required"`
	// The type of the model
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetType `json:"type,required"`
	JSON evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetJSON struct {
	Details     apijson.Field
	TokenID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset() {
}

// Details
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                            `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsTypeErc1155 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsType = "ERC1155"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsTypeErc1155:
		return true
	}
	return false
}

// The type of the model
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetTypeErc721 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetType = "ERC721"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetTypeErc721:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset struct {
	// Details
	Details EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetails `json:"details,required"`
	// Token Id
	TokenID int64 `json:"token_id,required"`
	// The type of the model
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetType `json:"type,required"`
	JSON evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetJSON struct {
	Details     apijson.Field
	TokenID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset() {
}

// Details
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                             `json:"symbol"`
	JSON   evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsTypeErc1155 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsType = "ERC1155"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsTypeErc1155:
		return true
	}
	return false
}

// The type of the model
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetTypeErc1155 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetType = "ERC1155"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetTypeErc1155:
		return true
	}
	return false
}

// The type of the model
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeNative  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType = "NATIVE"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc20   EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType = "ERC20"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc721  EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType = "ERC721"
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc1155 EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType = "ERC1155"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeNative, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc20, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc721, EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc1155:
		return true
	}
	return false
}

// The type of the model
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCode string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCodeGeneralInsufficientFunds EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCode = "GENERAL_INSUFFICIENT_FUNDS"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCode) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCodeGeneralInsufficientFunds:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails struct {
	// The address that is invalid
	Address string `json:"address,required"`
	// The type of the model
	Code EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCode `json:"code,required"`
	JSON evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsJSON struct {
	Address     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails() {
}

// The type of the model
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCode string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCodeGeneralInvalidAddress EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCode = "GENERAL_INVALID_ADDRESS"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCode) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCodeGeneralInvalidAddress:
		return true
	}
	return false
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails struct {
	// The error code
	Code string                                                                                                                                  `json:"code,required"`
	JSON evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetailsJSON struct {
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails() {
}

type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails struct {
	// The type of the model
	Code EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCode `json:"code,required"`
	// The domain name that is unsupported
	DomainName string `json:"domain_name,required"`
	// The message type that is unsupported
	MessageType string                                                                                                                                                   `json:"message_type,required"`
	JSON        evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsJSON `json:"-"`
}

// evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails]
type evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsJSON struct {
	Code        apijson.Field
	DomainName  apijson.Field
	MessageType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails) implementsEvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails() {
}

// The type of the model
type EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCode string

const (
	EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCodeUnsupportedEip712Message EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCode = "UNSUPPORTED_EIP712_MESSAGE"
)

func (r EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCode) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCodeUnsupportedEip712Message:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type EvmUserOperationScanResponseSimulationStatus string

const (
	EvmUserOperationScanResponseSimulationStatusSuccess EvmUserOperationScanResponseSimulationStatus = "Success"
	EvmUserOperationScanResponseSimulationStatusError   EvmUserOperationScanResponseSimulationStatus = "Error"
)

func (r EvmUserOperationScanResponseSimulationStatus) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseSimulationStatusSuccess, EvmUserOperationScanResponseSimulationStatusError:
		return true
	}
	return false
}

type EvmUserOperationScanResponseUserOperationGasEstimation struct {
	Status                           EvmUserOperationScanResponseUserOperationGasEstimationStatus `json:"status,required"`
	CallGasEstimate                  string                                                       `json:"call_gas_estimate"`
	Error                            string                                                       `json:"error"`
	PaymasterVerificationGasEstimate string                                                       `json:"paymaster_verification_gas_estimate"`
	PreVerificationGasEstimate       string                                                       `json:"pre_verification_gas_estimate"`
	VerificationGasEstimate          string                                                       `json:"verification_gas_estimate"`
	JSON                             evmUserOperationScanResponseUserOperationGasEstimationJSON   `json:"-"`
	union                            EvmUserOperationScanResponseUserOperationGasEstimationUnion
}

// evmUserOperationScanResponseUserOperationGasEstimationJSON contains the JSON
// metadata for the struct [EvmUserOperationScanResponseUserOperationGasEstimation]
type evmUserOperationScanResponseUserOperationGasEstimationJSON struct {
	Status                           apijson.Field
	CallGasEstimate                  apijson.Field
	Error                            apijson.Field
	PaymasterVerificationGasEstimate apijson.Field
	PreVerificationGasEstimate       apijson.Field
	VerificationGasEstimate          apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r evmUserOperationScanResponseUserOperationGasEstimationJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseUserOperationGasEstimation) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseUserOperationGasEstimation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EvmUserOperationScanResponseUserOperationGasEstimationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [UserOperationV6GasEstimation],
// [UserOperationV7GasEstimation],
// [EvmUserOperationScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError].
func (r EvmUserOperationScanResponseUserOperationGasEstimation) AsUnion() EvmUserOperationScanResponseUserOperationGasEstimationUnion {
	return r.union
}

// Union satisfied by [UserOperationV6GasEstimation],
// [UserOperationV7GasEstimation] or
// [EvmUserOperationScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError].
type EvmUserOperationScanResponseUserOperationGasEstimationUnion interface {
	implementsEvmUserOperationScanResponseUserOperationGasEstimation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseUserOperationGasEstimationUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(EvmUserOperationScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError{}),
		},
	)
}

type EvmUserOperationScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError struct {
	Error  string                                                                                                        `json:"error,required"`
	Status EvmUserOperationScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus `json:"status,required"`
	JSON   evmUserOperationScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON   `json:"-"`
}

// evmUserOperationScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError]
type evmUserOperationScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError) implementsEvmUserOperationScanResponseUserOperationGasEstimation() {
}

type EvmUserOperationScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus string

const (
	EvmUserOperationScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatusError EvmUserOperationScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus = "Error"
)

func (r EvmUserOperationScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatusError:
		return true
	}
	return false
}

type EvmUserOperationScanResponseUserOperationGasEstimationStatus string

const (
	EvmUserOperationScanResponseUserOperationGasEstimationStatusSuccess EvmUserOperationScanResponseUserOperationGasEstimationStatus = "Success"
	EvmUserOperationScanResponseUserOperationGasEstimationStatusError   EvmUserOperationScanResponseUserOperationGasEstimationStatus = "Error"
)

func (r EvmUserOperationScanResponseUserOperationGasEstimationStatus) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseUserOperationGasEstimationStatusSuccess, EvmUserOperationScanResponseUserOperationGasEstimationStatusError:
		return true
	}
	return false
}

type EvmUserOperationScanResponseValidation struct {
	// This field can have the runtime type of
	// [[]EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationFeature],
	// [[]EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeature].
	Features interface{} `json:"features,required"`
	// Result type returned when validation succeeds.
	ResultType EvmUserOperationScanResponseValidationResultType `json:"result_type,required"`
	// A string indicating if the simulation was successful or not.
	Status EvmUserOperationScanResponseValidationStatus `json:"status,required"`
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
	Reason string                                     `json:"reason"`
	JSON   evmUserOperationScanResponseValidationJSON `json:"-"`
	union  EvmUserOperationScanResponseValidationUnion
}

// evmUserOperationScanResponseValidationJSON contains the JSON metadata for the
// struct [EvmUserOperationScanResponseValidation]
type evmUserOperationScanResponseValidationJSON struct {
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

func (r evmUserOperationScanResponseValidationJSON) RawJSON() string {
	return r.raw
}

func (r *EvmUserOperationScanResponseValidation) UnmarshalJSON(data []byte) (err error) {
	*r = EvmUserOperationScanResponseValidation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EvmUserOperationScanResponseValidationUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidation],
// [EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationError].
func (r EvmUserOperationScanResponseValidation) AsUnion() EvmUserOperationScanResponseValidationUnion {
	return r.union
}

// Union satisfied by
// [EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidation]
// or
// [EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationError].
type EvmUserOperationScanResponseValidationUnion interface {
	implementsEvmUserOperationScanResponseValidation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmUserOperationScanResponseValidationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationError{}),
		},
	)
}

type EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidation struct {
	// A list of features about this transaction explaining the validation.
	Features []EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationFeature `json:"features,required"`
	// Result type returned when validation succeeds.
	ResultType EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationResultType `json:"result_type,required"`
	// A string indicating if the simulation was successful or not.
	Status EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationStatus `json:"status,required"`
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification"`
	// A textual description that can be presented to the user about what this
	// transaction is doing.
	Description string `json:"description"`
	// A textual description about the reasons the transaction was flagged with
	// result_type.
	Reason string                                                                            `json:"reason"`
	JSON   evmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationJSON `json:"-"`
}

// evmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidation]
type evmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationJSON struct {
	Features       apijson.Field
	ResultType     apijson.Field
	Status         apijson.Field
	Classification apijson.Field
	Description    apijson.Field
	Reason         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidation) implementsEvmUserOperationScanResponseValidation() {
}

type EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationFeature struct {
	// Textual description
	Description string `json:"description,required"`
	// Feature name
	FeatureID string `json:"feature_id,required"`
	// Security result of a transaction scan feature.
	Type EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType `json:"type,required"`
	// Address the feature refers to
	Address string `json:"address"`
	// Metadata related to the feature
	Metadata interface{}                                                                              `json:"metadata"`
	JSON     evmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationFeatureJSON `json:"-"`
}

// evmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationFeatureJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationFeature]
type evmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationFeatureJSON struct {
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	Address     apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationFeatureJSON) RawJSON() string {
	return r.raw
}

// Security result of a transaction scan feature.
type EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType string

const (
	EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeMalicious EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType = "Malicious"
	EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeWarning   EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType = "Warning"
	EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeBenign    EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType = "Benign"
	EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeInfo      EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType = "Info"
)

func (r EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeMalicious, EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeWarning, EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeBenign, EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeInfo:
		return true
	}
	return false
}

// Result type returned when validation succeeds.
type EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationResultType string

const (
	EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeBenign    EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationResultType = "Benign"
	EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeWarning   EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationResultType = "Warning"
	EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeMalicious EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationResultType = "Malicious"
)

func (r EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationResultType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeBenign, EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeWarning, EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeMalicious:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationStatus string

const (
	EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationStatusSuccess EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationStatus = "Success"
)

func (r EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationStatus) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationStatusSuccess:
		return true
	}
	return false
}

type EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationError struct {
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassification `json:"classification,required"`
	// A textual description that can be presented to the user about what this
	// transaction is doing.
	Description EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescription `json:"description,required"`
	// An error message if the validation failed.
	Error string `json:"error,required"`
	// A list of features about this transaction explaining the validation.
	Features []EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeature `json:"features,required"`
	// A textual description about the reasons the transaction was flagged with
	// result_type.
	Reason EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorReason `json:"reason,required"`
	// A string indicating if the transaction is safe to sign or not.
	ResultType EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultType `json:"result_type,required"`
	// A string indicating if the simulation was successful or not.
	Status EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatus `json:"status,required"`
	JSON   evmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorJSON   `json:"-"`
}

// evmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationError]
type evmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorJSON struct {
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

func (r *EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorJSON) RawJSON() string {
	return r.raw
}

func (r EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationError) implementsEvmUserOperationScanResponseValidation() {
}

// A textual classification that can be presented to the user explaining the
// reason.
type EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassification string

const (
	EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassificationEmpty EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassification = ""
)

func (r EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassification) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassificationEmpty:
		return true
	}
	return false
}

// A textual description that can be presented to the user about what this
// transaction is doing.
type EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescription string

const (
	EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescriptionEmpty EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescription = ""
)

func (r EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescription) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescriptionEmpty:
		return true
	}
	return false
}

type EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeature struct {
	// Textual description
	Description string `json:"description,required"`
	// Feature name
	FeatureID string `json:"feature_id,required"`
	// Security result of a transaction scan feature.
	Type EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType `json:"type,required"`
	// Address the feature refers to
	Address string `json:"address"`
	// Metadata related to the feature
	Metadata interface{}                                                                                   `json:"metadata"`
	JSON     evmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeatureJSON `json:"-"`
}

// evmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeatureJSON
// contains the JSON metadata for the struct
// [EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeature]
type evmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeatureJSON struct {
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	Address     apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeatureJSON) RawJSON() string {
	return r.raw
}

// Security result of a transaction scan feature.
type EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType string

const (
	EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeMalicious EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType = "Malicious"
	EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeWarning   EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType = "Warning"
	EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeBenign    EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType = "Benign"
	EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeInfo      EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType = "Info"
)

func (r EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeMalicious, EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeWarning, EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeBenign, EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeInfo:
		return true
	}
	return false
}

// A textual description about the reasons the transaction was flagged with
// result_type.
type EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorReason string

const (
	EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorReasonEmpty EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorReason = ""
)

func (r EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorReason) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorReasonEmpty:
		return true
	}
	return false
}

// A string indicating if the transaction is safe to sign or not.
type EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultType string

const (
	EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultTypeError EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultType = "Error"
)

func (r EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultTypeError:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatus string

const (
	EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatusSuccess EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatus = "Success"
)

func (r EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatus) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatusSuccess:
		return true
	}
	return false
}

// Result type returned when validation succeeds.
type EvmUserOperationScanResponseValidationResultType string

const (
	EvmUserOperationScanResponseValidationResultTypeBenign    EvmUserOperationScanResponseValidationResultType = "Benign"
	EvmUserOperationScanResponseValidationResultTypeWarning   EvmUserOperationScanResponseValidationResultType = "Warning"
	EvmUserOperationScanResponseValidationResultTypeMalicious EvmUserOperationScanResponseValidationResultType = "Malicious"
	EvmUserOperationScanResponseValidationResultTypeError     EvmUserOperationScanResponseValidationResultType = "Error"
)

func (r EvmUserOperationScanResponseValidationResultType) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseValidationResultTypeBenign, EvmUserOperationScanResponseValidationResultTypeWarning, EvmUserOperationScanResponseValidationResultTypeMalicious, EvmUserOperationScanResponseValidationResultTypeError:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type EvmUserOperationScanResponseValidationStatus string

const (
	EvmUserOperationScanResponseValidationStatusSuccess EvmUserOperationScanResponseValidationStatus = "Success"
)

func (r EvmUserOperationScanResponseValidationStatus) IsKnown() bool {
	switch r {
	case EvmUserOperationScanResponseValidationStatusSuccess:
		return true
	}
	return false
}

type EvmUserOperationScanParams struct {
	UserOperationRequest UserOperationRequestParam `json:"UserOperationRequest,required"`
}

func (r EvmUserOperationScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.UserOperationRequest)
}
