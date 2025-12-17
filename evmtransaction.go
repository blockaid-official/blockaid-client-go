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

// Report a misclassification of a transaction.
func (r *EvmTransactionService) Report(ctx context.Context, body EvmTransactionReportParams, opts ...option.RequestOption) (res *EvmTransactionReportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/evm/transaction/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets a transaction and returns a full simulation indicating what will happen in
// the transaction together with a recommended action and some textual reasons of
// why the transaction was flagged that way.
func (r *EvmTransactionService) Scan(ctx context.Context, body EvmTransactionScanParams, opts ...option.RequestOption) (res *EvmTransactionScanResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/evm/transaction/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EvmTransactionReportResponse = interface{}

type EvmTransactionScanResponse struct {
	Block                      string                                               `json:"block,required"`
	Chain                      string                                               `json:"chain,required"`
	AccountAddress             string                                               `json:"account_address"`
	Events                     []EvmTransactionScanResponseEvent                    `json:"events"`
	Features                   interface{}                                          `json:"features"`
	GasEstimation              EvmTransactionScanResponseGasEstimation              `json:"gas_estimation"`
	Simulation                 EvmTransactionScanResponseSimulation                 `json:"simulation"`
	UserOperationGasEstimation EvmTransactionScanResponseUserOperationGasEstimation `json:"user_operation_gas_estimation"`
	Validation                 EvmTransactionScanResponseValidation                 `json:"validation"`
	JSON                       evmTransactionScanResponseJSON                       `json:"-"`
}

// evmTransactionScanResponseJSON contains the JSON metadata for the struct
// [EvmTransactionScanResponse]
type evmTransactionScanResponseJSON struct {
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

func (r *EvmTransactionScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseEvent struct {
	Data           string                                  `json:"data,required"`
	EmitterAddress string                                  `json:"emitter_address,required"`
	Topics         []string                                `json:"topics,required"`
	EmitterName    string                                  `json:"emitter_name"`
	Name           string                                  `json:"name"`
	Params         []EvmTransactionScanResponseEventsParam `json:"params"`
	JSON           evmTransactionScanResponseEventJSON     `json:"-"`
}

// evmTransactionScanResponseEventJSON contains the JSON metadata for the struct
// [EvmTransactionScanResponseEvent]
type evmTransactionScanResponseEventJSON struct {
	Data           apijson.Field
	EmitterAddress apijson.Field
	Topics         apijson.Field
	EmitterName    apijson.Field
	Name           apijson.Field
	Params         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmTransactionScanResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseEventJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseEventsParam struct {
	Type         string                                           `json:"type,required"`
	Value        EvmTransactionScanResponseEventsParamsValueUnion `json:"value,required"`
	InternalType string                                           `json:"internalType"`
	Name         string                                           `json:"name"`
	JSON         evmTransactionScanResponseEventsParamJSON        `json:"-"`
}

// evmTransactionScanResponseEventsParamJSON contains the JSON metadata for the
// struct [EvmTransactionScanResponseEventsParam]
type evmTransactionScanResponseEventsParamJSON struct {
	Type         apijson.Field
	Value        apijson.Field
	InternalType apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EvmTransactionScanResponseEventsParam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseEventsParamJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or
// [EvmTransactionScanResponseEventsParamsValueArray].
type EvmTransactionScanResponseEventsParamsValueUnion interface {
	ImplementsEvmTransactionScanResponseEventsParamsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseEventsParamsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseEventsParamsValueArray{}),
		},
	)
}

type EvmTransactionScanResponseEventsParamsValueArray []interface{}

func (r EvmTransactionScanResponseEventsParamsValueArray) ImplementsEvmTransactionScanResponseEventsParamsValueUnion() {
}

type EvmTransactionScanResponseGasEstimation struct {
	Status   EvmTransactionScanResponseGasEstimationStatus `json:"status,required"`
	Error    string                                        `json:"error"`
	Estimate string                                        `json:"estimate"`
	Used     string                                        `json:"used"`
	JSON     evmTransactionScanResponseGasEstimationJSON   `json:"-"`
	union    EvmTransactionScanResponseGasEstimationUnion
}

// evmTransactionScanResponseGasEstimationJSON contains the JSON metadata for the
// struct [EvmTransactionScanResponseGasEstimation]
type evmTransactionScanResponseGasEstimationJSON struct {
	Status      apijson.Field
	Error       apijson.Field
	Estimate    apijson.Field
	Used        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionScanResponseGasEstimationJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseGasEstimation) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseGasEstimation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EvmTransactionScanResponseGasEstimationUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation],
// [EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError].
func (r EvmTransactionScanResponseGasEstimation) AsUnion() EvmTransactionScanResponseGasEstimationUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation]
// or
// [EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError].
type EvmTransactionScanResponseGasEstimationUnion interface {
	implementsEvmTransactionScanResponseGasEstimation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseGasEstimationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError{}),
		},
	)
}

type EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation struct {
	Estimate string                                                                                    `json:"estimate,required"`
	Status   EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatus `json:"status,required"`
	Used     string                                                                                    `json:"used,required"`
	JSON     evmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationJSON   `json:"-"`
}

// evmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation]
type evmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationJSON struct {
	Estimate    apijson.Field
	Status      apijson.Field
	Used        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimation) implementsEvmTransactionScanResponseGasEstimation() {
}

type EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatus string

const (
	EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatusSuccess EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatus = "Success"
)

func (r EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatus) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationStatusSuccess:
		return true
	}
	return false
}

type EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError struct {
	Error  string                                                                                         `json:"error,required"`
	Status EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus `json:"status,required"`
	JSON   evmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON   `json:"-"`
}

// evmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError]
type evmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationError) implementsEvmTransactionScanResponseGasEstimation() {
}

type EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus string

const (
	EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatusError EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus = "Error"
)

func (r EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatusError:
		return true
	}
	return false
}

type EvmTransactionScanResponseGasEstimationStatus string

const (
	EvmTransactionScanResponseGasEstimationStatusSuccess EvmTransactionScanResponseGasEstimationStatus = "Success"
	EvmTransactionScanResponseGasEstimationStatusError   EvmTransactionScanResponseGasEstimationStatus = "Error"
)

func (r EvmTransactionScanResponseGasEstimationStatus) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseGasEstimationStatusSuccess, EvmTransactionScanResponseGasEstimationStatusError:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulation struct {
	// A string indicating if the simulation was successful or not.
	Status EvmTransactionScanResponseSimulationStatus `json:"status,required"`
	// This field can have the runtime type of
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummary].
	AccountSummary interface{} `json:"account_summary"`
	// This field can have the runtime type of
	// [map[string]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetail].
	AddressDetails interface{} `json:"address_details"`
	// This field can have the runtime type of
	// [map[string][]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff].
	AssetsDiffs interface{} `json:"assets_diffs"`
	// This field can have the runtime type of
	// [map[string][]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement].
	ContractManagement interface{} `json:"contract_management"`
	// A string explaining why the transaction failed
	Description string `json:"description"`
	// An error message if the simulation failed.
	Error string `json:"error"`
	// This field can have the runtime type of
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails].
	ErrorDetails interface{} `json:"error_details"`
	// This field can have the runtime type of
	// [map[string][]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure].
	Exposures interface{} `json:"exposures"`
	// This field can have the runtime type of
	// [[]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalance].
	MissingBalances interface{} `json:"missing_balances"`
	// This field can have the runtime type of
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationParams].
	Params interface{} `json:"params"`
	// This field can have the runtime type of
	// [map[string][]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKey].
	SessionKey interface{} `json:"session_key"`
	// The number of times the simulation ran until success
	SimulationRunCount int64 `json:"simulation_run_count"`
	// This field can have the runtime type of
	// [map[string]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiff].
	TotalUsdDiff interface{} `json:"total_usd_diff"`
	// This field can have the runtime type of [map[string]map[string]string].
	TotalUsdExposure interface{} `json:"total_usd_exposure"`
	// This field can have the runtime type of
	// [[]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions].
	TransactionActions interface{}                              `json:"transaction_actions"`
	JSON               evmTransactionScanResponseSimulationJSON `json:"-"`
	union              EvmTransactionScanResponseSimulationUnion
}

// evmTransactionScanResponseSimulationJSON contains the JSON metadata for the
// struct [EvmTransactionScanResponseSimulation]
type evmTransactionScanResponseSimulationJSON struct {
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

func (r evmTransactionScanResponseSimulationJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulation) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EvmTransactionScanResponseSimulationUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulation],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationError].
func (r EvmTransactionScanResponseSimulation) AsUnion() EvmTransactionScanResponseSimulationUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulation] or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationError].
type EvmTransactionScanResponseSimulationUnion interface {
	implementsEvmTransactionScanResponseSimulation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationError{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulation struct {
	// Account summary for the account address. account address is determined implicit
	// by the `from` field in the transaction request, or explicit by the
	// account_address field in the request.
	AccountSummary EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummary `json:"account_summary,required"`
	// a dictionary including additional information about each relevant address in the
	// transaction.
	AddressDetails map[string]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetail `json:"address_details,required"`
	// dictionary describes the assets differences as a result of this transaction for
	// every involved address
	AssetsDiffs map[string][]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff `json:"assets_diffs,required"`
	// dictionary describes the exposure differences as a result of this transaction
	// for every involved address (as a result of any approval / setApproval / permit
	// function)
	Exposures map[string][]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure `json:"exposures,required"`
	// Session keys created in this transaction per address
	SessionKey map[string][]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKey `json:"session_key,required"`
	// A string indicating if the simulation was successful or not.
	Status EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationStatus `json:"status,required"`
	// dictionary represents the usd value each address gained / lost during this
	// transaction
	TotalUsdDiff map[string]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiff `json:"total_usd_diff,required"`
	// a dictionary representing the usd value each address is exposed to, split by
	// spenders
	TotalUsdExposure map[string]map[string]string `json:"total_usd_exposure,required"`
	// Describes the nature of the transaction and what happened as part of it
	TransactionActions []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions `json:"transaction_actions,required"`
	// Describes the state differences as a result of this transaction for every
	// involved address
	ContractManagement map[string][]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement `json:"contract_management"`
	// Missing balances in the transaction
	MissingBalances []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalance `json:"missing_balances"`
	// The parameters of the transaction that was simulated.
	Params EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationParams `json:"params"`
	// The number of times the simulation ran until success
	SimulationRunCount int64                                                                           `json:"simulation_run_count"`
	JSON               evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulation]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationJSON struct {
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

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulation) implementsEvmTransactionScanResponseSimulation() {
}

// Account summary for the account address. account address is determined implicit
// by the `from` field in the transaction request, or explicit by the
// account_address field in the request.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummary struct {
	// All assets diffs related to the account address
	AssetsDiffs []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff `json:"assets_diffs,required"`
	// All assets exposures related to the account address
	Exposures []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure `json:"exposures,required"`
	// Total usd diff related to the account address
	TotalUsdDiff EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiff `json:"total_usd_diff,required"`
	// Total usd exposure related to the account address
	TotalUsdExposure map[string]string `json:"total_usd_exposure,required"`
	// All assets traces related to the account address
	Traces []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace `json:"traces,required"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryJSON    `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummary]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryJSON struct {
	AssetsDiffs      apijson.Field
	Exposures        apijson.Field
	TotalUsdDiff     apijson.Field
	TotalUsdExposure apijson.Field
	Traces           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff struct {
	// This field can have the runtime type of
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAsset].
	Asset interface{} `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType `json:"asset_type,required"`
	// This field can have the runtime type of
	// [[]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffIn],
	// [[]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffIn],
	// [[]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffIn],
	// [[]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffIn].
	In interface{} `json:"in,required"`
	// This field can have the runtime type of
	// [[]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOut],
	// [[]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOut],
	// [[]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOut],
	// [[]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOut].
	Out interface{} `json:"out,required"`
	// This field can have the runtime type of
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChanges],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChanges],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChanges],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChanges].
	BalanceChanges interface{}                                                                                             `json:"balance_changes"`
	JSON           evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffJSON `json:"-"`
	union          EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff struct {
	// description of the asset for the current diff
	Asset EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOut `json:"out,required"`
	// shows the balance before making the transaction and after
	BalanceChanges EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChanges `json:"balance_changes"`
	JSON           evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffJSON           `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff() {
}

// description of the asset for the current diff
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType `json:"type,required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                            `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetJSON `json:"-"`
	union  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                   `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                    `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetTypeErc20  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType = "ERC20"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetTypeErc20, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffIn struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                         `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffInJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffInJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffIn]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffInJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOut struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                          `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOutJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOutJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOut]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOutJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffOutJSON) RawJSON() string {
	return r.raw
}

// shows the balance before making the transaction and after
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChanges struct {
	// balance of the account after making the transaction
	After EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfter `json:"after,required"`
	// balance of the account before making the transaction
	Before EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBefore `json:"before,required"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesJSON   `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChanges]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChanges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesJSON) RawJSON() string {
	return r.raw
}

// balance of the account after making the transaction
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfter struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                          `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfterJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfterJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfter]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfterJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesAfterJSON) RawJSON() string {
	return r.raw
}

// balance of the account before making the transaction
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBefore struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                           `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBefore]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc20AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff struct {
	// description of the asset for the current diff
	Asset EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOut `json:"out,required"`
	// shows the balance before making the transaction and after
	BalanceChanges EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChanges `json:"balance_changes"`
	JSON           evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffJSON           `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff() {
}

// description of the asset for the current diff
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                             `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetJSON `json:"-"`
	union  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                     `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                     `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetTypeErc721 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType = "ERC721"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetTypeErc721, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffIn struct {
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
	UsdPrice string                                                                                                                                                          `json:"usd_price"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffInJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffInJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffIn]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffInJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOut struct {
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
	UsdPrice string                                                                                                                                                           `json:"usd_price"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOutJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOutJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOut]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOutJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffOutJSON) RawJSON() string {
	return r.raw
}

// shows the balance before making the transaction and after
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChanges struct {
	// balance of the account after making the transaction
	After EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfter `json:"after,required"`
	// balance of the account before making the transaction
	Before EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBefore `json:"before,required"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesJSON   `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChanges]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChanges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesJSON) RawJSON() string {
	return r.raw
}

// balance of the account after making the transaction
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfter struct {
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
	UsdPrice string                                                                                                                                                                           `json:"usd_price"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfterJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfterJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfter]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfterJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesAfterJSON) RawJSON() string {
	return r.raw
}

// balance of the account before making the transaction
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBefore struct {
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
	UsdPrice string                                                                                                                                                                            `json:"usd_price"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBefore]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc721AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff struct {
	// description of the asset for the current diff
	Asset EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOut `json:"out,required"`
	// shows the balance before making the transaction and after
	BalanceChanges EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChanges `json:"balance_changes"`
	JSON           evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffJSON           `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff() {
}

// description of the asset for the current diff
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                              `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetJSON `json:"-"`
	union  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                       `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                                      `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetTypeErc1155 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType = "ERC1155"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetTypeNonerc  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetTypeErc1155, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffIn struct {
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
	UsdPrice string                                                                                                                                                           `json:"usd_price"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffInJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffInJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffIn]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffInJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOut struct {
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
	UsdPrice string                                                                                                                                                            `json:"usd_price"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOutJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOutJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOut]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOutJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffOutJSON) RawJSON() string {
	return r.raw
}

// shows the balance before making the transaction and after
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChanges struct {
	// balance of the account after making the transaction
	After EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfter `json:"after,required"`
	// balance of the account before making the transaction
	Before EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBefore `json:"before,required"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesJSON   `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChanges]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChanges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesJSON) RawJSON() string {
	return r.raw
}

// balance of the account after making the transaction
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfter struct {
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
	UsdPrice string                                                                                                                                                                            `json:"usd_price"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfterJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfterJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfter]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfterJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesAfterJSON) RawJSON() string {
	return r.raw
}

// balance of the account before making the transaction
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBefore struct {
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
	UsdPrice string                                                                                                                                                                             `json:"usd_price"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBefore]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseErc1155AddressAssetBalanceChangeDiffBalanceChangesBeforeJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff struct {
	// description of the asset for the current diff
	Asset EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOut `json:"out,required"`
	// shows the balance before making the transaction and after
	BalanceChanges EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChanges `json:"balance_changes"`
	JSON           evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffJSON           `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiff() {
}

// description of the asset for the current diff
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAsset struct {
	ChainID   int64  `json:"chain_id,required"`
	ChainName string `json:"chain_name,required"`
	Decimals  int64  `json:"decimals,required"`
	LogoURL   string `json:"logo_url,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetType `json:"type,required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                             `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAsset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetJSON struct {
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

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetTypeNative EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetType = "NATIVE"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffAssetTypeNative:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffIn struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                          `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffInJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffInJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffIn]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffInJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOut struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                           `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOutJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOutJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOut]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOutJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffOutJSON) RawJSON() string {
	return r.raw
}

// shows the balance before making the transaction and after
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChanges struct {
	// balance of the account after making the transaction
	After EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfter `json:"after,required"`
	// balance of the account before making the transaction
	Before EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBefore `json:"before,required"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesJSON   `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChanges]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChanges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesJSON) RawJSON() string {
	return r.raw
}

// balance of the account after making the transaction
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfter struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                           `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfterJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfterJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfter]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfterJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesAfterJSON) RawJSON() string {
	return r.raw
}

// balance of the account before making the transaction
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBefore struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                            `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBeforeJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBeforeJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBefore]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBeforeJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsRoutersEvmResponseNativeAddressAssetBalanceChangeDiffBalanceChangesBeforeJSON) RawJSON() string {
	return r.raw
}

// type of the asset for the current diff
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc20   EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType = "ERC20"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc721  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType = "ERC721"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc1155 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType = "ERC1155"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeNative  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType = "NATIVE"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc20, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc721, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc1155, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryAssetsDiffsAssetTypeNative:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure struct {
	// This field can have the runtime type of
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset].
	Asset interface{} `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType `json:"asset_type,required"`
	// This field can have the runtime type of
	// [map[string]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpender],
	// [map[string]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpender],
	// [map[string]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpender].
	Spenders interface{}                                                                                           `json:"spenders,required"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposureJSON `json:"-"`
	union    EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType `json:"asset_type,required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpender `json:"spenders,required"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureJSON               `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposure) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure() {
}

// description of the asset for the current diff
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType `json:"type,required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                            `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetJSON `json:"-"`
	union  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                   `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                    `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetTypeErc20  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType = "ERC20"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetTypeErc20, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpender struct {
	// the amount that was asked in the approval request for this spender from the
	// current address and asset
	Approval string                                                                                                                                                     `json:"approval,required"`
	Exposure []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure `json:"exposure,required"`
	// the usd price of the approval amount
	ApprovalUsdPrice string `json:"approval_usd_price"`
	// the expiration time of the permit2 protocol
	Expiration time.Time `json:"expiration" format:"date-time"`
	// user friendly description of the approval
	Summary string                                                                                                                                              `json:"summary"`
	JSON    evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpender]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON struct {
	Approval         apijson.Field
	Exposure         apijson.Field
	ApprovalUsdPrice apijson.Field
	Expiration       apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure struct {
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
	Value string                                                                                                                                                       `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON `json:"-"`
	union EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON struct {
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

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
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
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
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
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                  `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                   `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType `json:"asset_type,required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpender `json:"spenders,required"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureJSON               `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposure) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure() {
}

// description of the asset for the current diff
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                             `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetJSON `json:"-"`
	union  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                     `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                     `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetTypeErc721 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType = "ERC721"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetTypeErc721, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpender struct {
	Exposure []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure `json:"exposure,required"`
	// boolean indicates whether an is_approved_for_all function was used (missing in
	// case of ERC20 / ERC1155)
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// user friendly description of the approval
	Summary string                                                                                                                                               `json:"summary"`
	JSON    evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpender]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure struct {
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
	Value string                                                                                                                                                        `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON `json:"-"`
	union EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON struct {
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

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
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
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
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
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                   `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                    `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType `json:"asset_type,required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpender `json:"spenders,required"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureJSON               `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposure) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposure() {
}

// description of the asset for the current diff
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                              `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON `json:"-"`
	union  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                       `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                      `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeErc1155 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType = "ERC1155"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeNonerc  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeErc1155, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpender struct {
	Exposure []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure `json:"exposure,required"`
	// boolean indicates whether an is_approved_for_all function was used (missing in
	// case of ERC20 / ERC1155)
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// user friendly description of the approval
	Summary string                                                                                                                                                `json:"summary"`
	JSON    evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpender]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure struct {
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
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON `json:"-"`
	union EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON struct {
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

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
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
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
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
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                    `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                                     `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

// type of the asset for the current diff
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc20   EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType = "ERC20"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc721  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType = "ERC721"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc1155 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType = "ERC1155"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc20, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc721, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryExposuresAssetTypeErc1155:
		return true
	}
	return false
}

// Total usd diff related to the account address
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiff struct {
	In    string                                                                                                    `json:"in,required"`
	Out   string                                                                                                    `json:"out,required"`
	Total string                                                                                                    `json:"total,required"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace struct {
	// This field can have the runtime type of
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAsset],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset].
	Asset interface{} `json:"asset,required"`
	// type of the trace
	TraceType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceType `json:"trace_type,required"`
	// The type of the model
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType `json:"type,required"`
	// This field can have the runtime type of
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiff],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiff],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiff],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiff].
	Diff interface{} `json:"diff"`
	// This field can have the runtime type of
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposed],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposed].
	Exposed interface{} `json:"exposed"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address"`
	// This field can have the runtime type of
	// [[]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabels],
	// [[]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabels],
	// [[]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabels],
	// [[]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabels].
	Labels interface{} `json:"labels"`
	// The owner of the assets
	Owner string `json:"owner"`
	// The spender of the assets
	Spender string `json:"spender"`
	// The address where the assets are moved to
	ToAddress string                                                                                             `json:"to_address"`
	JSON      evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTraceJSON `json:"-"`
	union     EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTraceJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTraceJSON struct {
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

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTraceJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace struct {
	// Description of the asset in the trace
	Asset EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset `json:"asset,required"`
	// The difference in value for the asset in the trace
	Diff EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiff `json:"diff,required"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address,required"`
	// The address where the assets are moved to
	ToAddress string `json:"to_address,required"`
	// type of the trace
	TraceType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceType `json:"type,required"`
	// List of labels that describe the trace
	Labels []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabels `json:"labels"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceJSON     `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceJSON struct {
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

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTrace) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetType `json:"type,required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                    `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetJSON `json:"-"`
	union  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                           `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                            `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetTypeErc20  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetType = "ERC20"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetTypeErc20, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceAssetTypeNonerc:
		return true
	}
	return false
}

// The difference in value for the asset in the trace
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                   `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceDiffJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceTypeAssetTrace EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceType = "AssetTrace"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTraceTypeAssetTrace:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTypeErc20AssetTrace EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceType = "ERC20AssetTrace"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceTypeErc20AssetTrace:
		return true
	}
	return false
}

// An enumeration.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabels string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabelsGasFee EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabels = "GAS_FEE"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabels) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20AssetTraceLabelsGasFee:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace struct {
	// Description of the asset in the trace
	Asset EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset `json:"asset,required"`
	// The difference in value for the asset in the trace
	Diff EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiff `json:"diff,required"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address,required"`
	// The address where the assets are moved to
	ToAddress string `json:"to_address,required"`
	// type of the trace
	TraceType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceType `json:"type,required"`
	// List of labels that describe the trace
	Labels []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabels `json:"labels"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceJSON     `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceJSON struct {
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

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTrace) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                     `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetJSON `json:"-"`
	union  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                             `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                             `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetTypeErc721 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetType = "ERC721"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetTypeErc721, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceAssetTypeNonerc:
		return true
	}
	return false
}

// The difference in value for the asset in the trace
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiff struct {
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
	UsdPrice string                                                                                                                                    `json:"usd_price"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceDiffJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceTypeAssetTrace EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceType = "AssetTrace"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTraceTypeAssetTrace:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTypeErc721AssetTrace EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceType = "ERC721AssetTrace"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceTypeErc721AssetTrace:
		return true
	}
	return false
}

// An enumeration.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabels string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabelsGasFee EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabels = "GAS_FEE"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabels) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721AssetTraceLabelsGasFee:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace struct {
	// Description of the asset in the trace
	Asset EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset `json:"asset,required"`
	// The difference in value for the asset in the trace
	Diff EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiff `json:"diff,required"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address,required"`
	// The address where the assets are moved to
	ToAddress string `json:"to_address,required"`
	// type of the trace
	TraceType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceType `json:"type,required"`
	// List of labels that describe the trace
	Labels []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabels `json:"labels"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceJSON     `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceJSON struct {
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

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTrace) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                      `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetJSON `json:"-"`
	union  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                               `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                              `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetTypeErc1155 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetType = "ERC1155"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetTypeNonerc  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetTypeErc1155, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceAssetTypeNonerc:
		return true
	}
	return false
}

// The difference in value for the asset in the trace
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiff struct {
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
	UsdPrice string                                                                                                                                     `json:"usd_price"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceDiffJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceTypeAssetTrace EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceType = "AssetTrace"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTraceTypeAssetTrace:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTypeErc1155AssetTrace EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceType = "ERC1155AssetTrace"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceTypeErc1155AssetTrace:
		return true
	}
	return false
}

// An enumeration.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabels string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabelsGasFee EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabels = "GAS_FEE"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabels) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155AssetTraceLabelsGasFee:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace struct {
	// Description of the asset in the trace
	Asset EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAsset `json:"asset,required"`
	// The difference in value for the asset in the trace
	Diff EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiff `json:"diff,required"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address,required"`
	// The address where the assets are moved to
	ToAddress string `json:"to_address,required"`
	// type of the trace
	TraceType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceType `json:"type,required"`
	// List of labels that describe the trace
	Labels []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabels `json:"labels"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceJSON     `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceJSON struct {
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

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTrace) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAsset struct {
	ChainID   int64  `json:"chain_id,required"`
	ChainName string `json:"chain_name,required"`
	Decimals  int64  `json:"decimals,required"`
	LogoURL   string `json:"logo_url,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetType `json:"type,required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                     `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAsset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetJSON struct {
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

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetTypeNative EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetType = "NATIVE"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceAssetTypeNative:
		return true
	}
	return false
}

// The difference in value for the asset in the trace
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                    `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceDiffJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceTypeAssetTrace EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceType = "AssetTrace"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTraceTypeAssetTrace:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTypeNativeAssetTrace EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceType = "NativeAssetTrace"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceTypeNativeAssetTrace:
		return true
	}
	return false
}

// An enumeration.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabels string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabelsGasFee EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabels = "GAS_FEE"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabels) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseNativeAssetTraceLabelsGasFee:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace struct {
	// Description of the asset in the trace
	Asset   EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset   `json:"asset,required"`
	Exposed EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposed `json:"exposed,required"`
	// The owner of the assets
	Owner string `json:"owner,required"`
	// The spender of the assets
	Spender string `json:"spender,required"`
	// type of the trace
	TraceType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceType `json:"type,required"`
	JSON evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceJSON struct {
	Asset       apijson.Field
	Exposed     apijson.Field
	Owner       apijson.Field
	Spender     apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTrace) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetType `json:"type,required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                       `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetJSON `json:"-"`
	union  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                              `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                               `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetTypeErc20  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetType = "ERC20"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetTypeErc20, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposed struct {
	RawValue string                                                                                                                                         `json:"raw_value,required"`
	UsdPrice float64                                                                                                                                        `json:"usd_price"`
	Value    float64                                                                                                                                        `json:"value"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposedJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposedJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposed]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposedJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceExposedJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceTypeExposureTrace EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceType = "ExposureTrace"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTraceTypeExposureTrace:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTypeErc20ExposureTrace EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceType = "ERC20ExposureTrace"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc20ExposureTraceTypeErc20ExposureTrace:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace struct {
	// Description of the asset in the trace
	Asset EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset `json:"asset,required"`
	// The owner of the assets
	Owner string `json:"owner,required"`
	// The spender of the assets
	Spender string `json:"spender,required"`
	// type of the trace
	TraceType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type    EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceType    `json:"type,required"`
	Exposed EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposed `json:"exposed"`
	JSON    evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceJSON    `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceJSON struct {
	Asset       apijson.Field
	Owner       apijson.Field
	Spender     apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	Exposed     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTrace) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                        `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetJSON `json:"-"`
	union  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetTypeErc721 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetType = "ERC721"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetTypeErc721, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceAssetTypeNonerc:
		return true
	}
	return false
}

// type of the trace
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceTypeExposureTrace EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceType = "ExposureTrace"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTraceTypeExposureTrace:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTypeErc721ExposureTrace EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceType = "ERC721ExposureTrace"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceTypeErc721ExposureTrace:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposed struct {
	Amount   int64                                                                                                                                           `json:"amount,required"`
	TokenID  string                                                                                                                                          `json:"token_id,required"`
	IsMint   bool                                                                                                                                            `json:"is_mint"`
	LogoURL  string                                                                                                                                          `json:"logo_url"`
	UsdPrice float64                                                                                                                                         `json:"usd_price"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposedJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposedJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposed]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposedJSON struct {
	Amount      apijson.Field
	TokenID     apijson.Field
	IsMint      apijson.Field
	LogoURL     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc721ExposureTraceExposedJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace struct {
	// Description of the asset in the trace
	Asset EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset `json:"asset,required"`
	// The owner of the assets
	Owner string `json:"owner,required"`
	// The spender of the assets
	Spender string `json:"spender,required"`
	// type of the trace
	TraceType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceType `json:"type,required"`
	JSON evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceJSON struct {
	Asset       apijson.Field
	Owner       apijson.Field
	Spender     apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTrace) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                         `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetJSON `json:"-"`
	union  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                  `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                 `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetTypeErc1155 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetType = "ERC1155"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetTypeNonerc  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetTypeErc1155, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceAssetTypeNonerc:
		return true
	}
	return false
}

// type of the trace
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceTypeExposureTrace EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceType = "ExposureTrace"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTraceTypeExposureTrace:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTypeErc1155ExposureTrace EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceType = "ERC1155ExposureTrace"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesRoutersEvmResponseErc1155ExposureTraceTypeErc1155ExposureTrace:
		return true
	}
	return false
}

// type of the trace
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceTypeAssetTrace    EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceType = "AssetTrace"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceTypeExposureTrace EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceType = "ExposureTrace"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceTypeAssetTrace, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTraceTypeExposureTrace:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc20AssetTrace      EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC20AssetTrace"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc721AssetTrace     EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC721AssetTrace"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc1155AssetTrace    EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC1155AssetTrace"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeNativeAssetTrace     EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "NativeAssetTrace"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc20ExposureTrace   EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC20ExposureTrace"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc721ExposureTrace  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC721ExposureTrace"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc1155ExposureTrace EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType = "ERC1155ExposureTrace"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc20AssetTrace, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc721AssetTrace, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc1155AssetTrace, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeNativeAssetTrace, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc20ExposureTrace, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc721ExposureTrace, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAccountSummaryTracesTypeErc1155ExposureTrace:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetail struct {
	// contains the contract's name if the address is a verified contract
	ContractName string `json:"contract_name"`
	// Whether the address is an eoa or a contract
	IsEoa bool `json:"is_eoa"`
	// known name tag for the address
	NameTag string                                                                                       `json:"name_tag"`
	JSON    evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetailJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetailJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetail]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetailJSON struct {
	ContractName apijson.Field
	IsEoa        apijson.Field
	NameTag      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAddressDetailJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff struct {
	// This field can have the runtime type of
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAsset].
	Asset interface{} `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType `json:"asset_type,required"`
	// This field can have the runtime type of
	// [[]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffIn],
	// [[]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffIn],
	// [[]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffIn],
	// [[]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffIn].
	In interface{} `json:"in,required"`
	// This field can have the runtime type of
	// [[]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOut],
	// [[]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOut],
	// [[]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOut],
	// [[]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOut].
	Out   interface{}                                                                               `json:"out,required"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffJSON `json:"-"`
	union EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff struct {
	// description of the asset for the current diff
	Asset EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOut `json:"out,required"`
	JSON evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffJSON  `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff() {
}

// description of the asset for the current diff
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType `json:"type,required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                 `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetJSON `json:"-"`
	union  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                        `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                         `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetTypeErc20  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType = "ERC20"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetTypeErc20, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffIn struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                              `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffInJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffInJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffIn]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffInJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOut struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                               `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOutJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOutJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOut]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc20AddressAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff struct {
	// description of the asset for the current diff
	Asset EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOut `json:"out,required"`
	JSON evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffJSON  `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff() {
}

// description of the asset for the current diff
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                  `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetJSON `json:"-"`
	union  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                          `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                          `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetTypeErc721 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType = "ERC721"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetTypeErc721, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffIn struct {
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
	UsdPrice string                                                                                                                               `json:"usd_price"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffInJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffInJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffIn]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffInJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOut struct {
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
	UsdPrice string                                                                                                                                `json:"usd_price"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOutJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOutJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOut]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOutJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc721AddressAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff struct {
	// description of the asset for the current diff
	Asset EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOut `json:"out,required"`
	JSON evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffJSON  `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff() {
}

// description of the asset for the current diff
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                   `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetJSON `json:"-"`
	union  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                            `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                           `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetTypeErc1155 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType = "ERC1155"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetTypeNonerc  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetTypeErc1155, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffIn struct {
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
	UsdPrice string                                                                                                                                `json:"usd_price"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffInJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffInJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffIn]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffInJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOut struct {
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
	UsdPrice string                                                                                                                                 `json:"usd_price"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOutJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOutJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOut]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOutJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseErc1155AddressAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff struct {
	// description of the asset for the current diff
	Asset EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOut `json:"out,required"`
	JSON evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffJSON  `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiff() {
}

// description of the asset for the current diff
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAsset struct {
	ChainID   int64  `json:"chain_id,required"`
	ChainName string `json:"chain_name,required"`
	Decimals  int64  `json:"decimals,required"`
	LogoURL   string `json:"logo_url,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetType `json:"type,required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                  `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAsset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetJSON struct {
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

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetTypeNative EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetType = "NATIVE"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffAssetTypeNative:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffIn struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                               `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffInJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffInJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffIn]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffInJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffInJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOut struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOutJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOutJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOut]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsRoutersEvmResponseNativeAddressAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

// type of the asset for the current diff
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc20   EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType = "ERC20"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc721  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType = "ERC721"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc1155 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType = "ERC1155"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeNative  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType = "NATIVE"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc20, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc721, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeErc1155, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationAssetsDiffsAssetTypeNative:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure struct {
	// This field can have the runtime type of
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset].
	Asset interface{} `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType `json:"asset_type,required"`
	// This field can have the runtime type of
	// [map[string]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpender],
	// [map[string]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpender],
	// [map[string]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpender].
	Spenders interface{}                                                                             `json:"spenders,required"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposureJSON `json:"-"`
	union    EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType `json:"asset_type,required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpender `json:"spenders,required"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureJSON               `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposure) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure() {
}

// description of the asset for the current diff
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType `json:"type,required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                              `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetJSON `json:"-"`
	union  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                     `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                      `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetTypeErc20  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType = "ERC20"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetTypeErc20, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpender struct {
	// the amount that was asked in the approval request for this spender from the
	// current address and asset
	Approval string                                                                                                                                       `json:"approval,required"`
	Exposure []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure `json:"exposure,required"`
	// the usd price of the approval amount
	ApprovalUsdPrice string `json:"approval_usd_price"`
	// the expiration time of the permit2 protocol
	Expiration time.Time `json:"expiration" format:"date-time"`
	// user friendly description of the approval
	Summary string                                                                                                                                `json:"summary"`
	JSON    evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpender]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON struct {
	Approval         apijson.Field
	Exposure         apijson.Field
	ApprovalUsdPrice apijson.Field
	Expiration       apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure struct {
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
	Value string                                                                                                                                         `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON `json:"-"`
	union EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON struct {
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

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
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
	UsdPrice string                                                                                                                                                                      `json:"usd_price"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
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
	UsdPrice string                                                                                                                                                                     `json:"usd_price"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                    `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                     `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc20AddressExposureSpendersExposure() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType `json:"asset_type,required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpender `json:"spenders,required"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureJSON               `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposure) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure() {
}

// description of the asset for the current diff
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                               `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetJSON `json:"-"`
	union  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                       `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType = "ERC721"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsErc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                       `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetTypeErc721 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType = "ERC721"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetTypeErc721, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpender struct {
	Exposure []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure `json:"exposure,required"`
	// boolean indicates whether an is_approved_for_all function was used (missing in
	// case of ERC20 / ERC1155)
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// user friendly description of the approval
	Summary string                                                                                                                                 `json:"summary"`
	JSON    evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpender]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure struct {
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
	Value string                                                                                                                                          `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON `json:"-"`
	union EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON struct {
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

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
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
	UsdPrice string                                                                                                                                                                       `json:"usd_price"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
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
	UsdPrice string                                                                                                                                                                      `json:"usd_price"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                     `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                      `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc721AddressExposureSpendersExposure() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure struct {
	// description of the asset for the current diff
	Asset EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType `json:"asset_type,required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpender `json:"spenders,required"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureJSON               `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposure) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposure() {
}

// description of the asset for the current diff
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON `json:"-"`
	union  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                         `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType = "ERC1155"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsErc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                        `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetRoutersEvmTokenDetailsNonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeErc1155 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType = "ERC1155"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeNonerc  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType = "NONERC"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeErc1155, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpender struct {
	Exposure []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure `json:"exposure,required"`
	// boolean indicates whether an is_approved_for_all function was used (missing in
	// case of ERC20 / ERC1155)
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// user friendly description of the approval
	Summary string                                                                                                                                  `json:"summary"`
	JSON    evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpender]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure struct {
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
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON `json:"-"`
	union EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON struct {
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

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff struct {
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
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc1155Diff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff struct {
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
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc721Diff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                      `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseErc20Diff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                                                                                                                       `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposureRoutersEvmResponseNativeDiff) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresRoutersEvmResponseErc1155AddressExposureSpendersExposure() {
}

// type of the asset for the current diff
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc20   EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType = "ERC20"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc721  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType = "ERC721"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc1155 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType = "ERC1155"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc20, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc721, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationExposuresAssetTypeErc1155:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKey struct {
	Key      string                                                                                        `json:"key,required"`
	Policies []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy `json:"policies,required"`
	Signer   string                                                                                        `json:"signer,required"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyJSON     `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKey]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyJSON struct {
	Key         apijson.Field
	Policies    apijson.Field
	Signer      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy struct {
	// This field can have the runtime type of
	// [[]EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArg].
	Args interface{} `json:"args"`
	// This field can have the runtime type of
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails].
	AssetDetails interface{}                                                                                       `json:"asset_details"`
	EndTime      time.Time                                                                                         `json:"end_time" format:"date-time"`
	Method       string                                                                                            `json:"method"`
	Recipient    string                                                                                            `json:"recipient"`
	StartTime    time.Time                                                                                         `json:"start_time" format:"date-time"`
	ToAddress    string                                                                                            `json:"to_address"`
	Type         EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType `json:"type"`
	ValueLimit   string                                                                                            `json:"value_limit"`
	JSON         evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicyJSON   `json:"-"`
	union        EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicyJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicyJSON struct {
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

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicyJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy struct {
	AssetDetails EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails `json:"asset_details,required"`
	Recipient    string                                                                                                                                       `json:"recipient"`
	Type         EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyType         `json:"type"`
	ValueLimit   string                                                                                                                                       `json:"value_limit"`
	JSON         evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyJSON         `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyJSON struct {
	AssetDetails apijson.Field
	Recipient    apijson.Field
	Type         apijson.Field
	ValueLimit   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicy) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails struct {
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsType `json:"type,required"`
	// address of the token
	Address   string `json:"address"`
	ChainID   int64  `json:"chain_id"`
	ChainName string `json:"chain_name"`
	LogoURL   string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                           `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsJSON `json:"-"`
	union  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsJSON struct {
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

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails struct {
	ChainID   int64  `json:"chain_id,required"`
	ChainName string `json:"chain_name,required"`
	Decimals  int64  `json:"decimals,required"`
	LogoURL   string `json:"logo_url,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsType `json:"type,required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                   `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsJSON struct {
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

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsTypeNative EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsType = "NATIVE"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsNativeAssetDetailsTypeNative:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                  `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetails() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsTypeNative EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsType = "NATIVE"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsTypeErc20  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsType = "ERC20"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsTypeNative, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyAssetDetailsTypeErc20:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyTypeTransferPolicy EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyType = "TRANSFER_POLICY"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysTransferPolicyTypeTransferPolicy:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy struct {
	ToAddress  string                                                                                                                            `json:"to_address,required"`
	Args       []EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArg `json:"args"`
	Method     string                                                                                                                            `json:"method"`
	Type       EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyType  `json:"type"`
	ValueLimit string                                                                                                                            `json:"value_limit"`
	JSON       evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyJSON  `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyJSON struct {
	ToAddress   apijson.Field
	Args        apijson.Field
	Method      apijson.Field
	Type        apijson.Field
	ValueLimit  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicy) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArg struct {
	// An enumeration.
	Condition EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition `json:"condition,required"`
	Index     int64                                                                                                                                     `json:"index,required"`
	Value     string                                                                                                                                    `json:"value,required"`
	JSON      evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgJSON       `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArg]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgJSON struct {
	Condition   apijson.Field
	Index       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgJSON) RawJSON() string {
	return r.raw
}

// An enumeration.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionUnconstrained  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "UNCONSTRAINED"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionEqual          EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "EQUAL"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionGreater        EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "GREATER"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionLess           EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "LESS"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionGreaterOrEqual EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "GREATER_OR_EQUAL"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionLessOrEqual    EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "LESS_OR_EQUAL"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionNotEqual       EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition = "NOT_EQUAL"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsCondition) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionUnconstrained, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionEqual, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionGreater, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionLess, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionGreaterOrEqual, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionLessOrEqual, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyArgsConditionNotEqual:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyTypeCallPolicy EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyType = "CALL_POLICY"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysCallPolicyTypeCallPolicy:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy struct {
	ValueLimit string                                                                                                                          `json:"value_limit,required"`
	Type       EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyType `json:"type"`
	JSON       evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyJSON struct {
	ValueLimit  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicy) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyTypeGasPolicy EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyType = "GAS_POLICY"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysGasPolicyTypeGasPolicy:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy struct {
	EndTime   time.Time                                                                                                                              `json:"end_time" format:"date-time"`
	StartTime time.Time                                                                                                                              `json:"start_time" format:"date-time"`
	Type      EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyType `json:"type"`
	JSON      evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicy) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPolicy() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyTypeExpirationPolicy EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyType = "EXPIRATION_POLICY"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesRoutersEvmSessionKeysExpirationPolicyTypeExpirationPolicy:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeTransferPolicy   EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType = "TRANSFER_POLICY"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeCallPolicy       EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType = "CALL_POLICY"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeGasPolicy        EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType = "GAS_POLICY"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeExpirationPolicy EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType = "EXPIRATION_POLICY"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeTransferPolicy, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeCallPolicy, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeGasPolicy, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationSessionKeyPoliciesTypeExpirationPolicy:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationStatus string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationStatusSuccess EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationStatus = "Success"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationStatus) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationStatusSuccess:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiff struct {
	In    string                                                                                      `json:"in,required"`
	Out   string                                                                                      `json:"out,required"`
	Total string                                                                                      `json:"total,required"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiffJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiffJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiff]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

// An enumeration.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsMint            EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "mint"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsStake           EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "stake"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsSwap            EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "swap"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsNativeTransfer  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "native_transfer"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsTokenTransfer   EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "token_transfer"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsApproval        EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "approval"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsSetCodeAccount  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "set_code_account"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsProxyUpgrade    EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "proxy_upgrade"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsOwnershipChange EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions = "ownership_change"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActions) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsMint, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsStake, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsSwap, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsNativeTransfer, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsTokenTransfer, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsApproval, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsSetCodeAccount, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsProxyUpgrade, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationTransactionActionsOwnershipChange:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement struct {
	// The type of the state change
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType `json:"type,required"`
	// This field can have the runtime type of
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfter],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfter],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfter].
	After interface{} `json:"after"`
	// This field can have the runtime type of
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBefore],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBefore],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBefore].
	Before interface{} `json:"before"`
	// The delegated address
	DelegatedAddress string `json:"delegated_address"`
	// The direct creator address of the new contract
	Deployer string                                                                                            `json:"deployer"`
	JSON     evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementJSON `json:"-"`
	union    EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementJSON struct {
	Type             apijson.Field
	After            apijson.Field
	Before           apijson.Field
	DelegatedAddress apijson.Field
	Deployer         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement struct {
	// The state after the transaction
	After EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfter `json:"after,required"`
	// The state before the transaction
	Before EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBefore `json:"before,required"`
	// The type of the state change
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementType `json:"type,required"`
	JSON evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementJSON struct {
	After       apijson.Field
	Before      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagement) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement() {
}

// The state after the transaction
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfter struct {
	Address string                                                                                                                                         `json:"address,required"`
	JSON    evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfterJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfterJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfter]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfterJSON struct {
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementAfterJSON) RawJSON() string {
	return r.raw
}

// The state before the transaction
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBefore struct {
	Address string                                                                                                                                          `json:"address,required"`
	JSON    evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBeforeJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBeforeJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBefore]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBeforeJSON struct {
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementBeforeJSON) RawJSON() string {
	return r.raw
}

// The type of the state change
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementTypeProxyUpgrade EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementType = "PROXY_UPGRADE"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseProxyUpgradeManagementTypeProxyUpgrade:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement struct {
	// The state after the transaction
	After EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfter `json:"after,required"`
	// The state before the transaction
	Before EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBefore `json:"before,required"`
	// The type of the state change
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementType `json:"type,required"`
	JSON evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementJSON struct {
	After       apijson.Field
	Before      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagement) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement() {
}

// The state after the transaction
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfter struct {
	Owners []string                                                                                                                                          `json:"owners,required"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfterJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfterJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfter]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfterJSON struct {
	Owners      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementAfterJSON) RawJSON() string {
	return r.raw
}

// The state before the transaction
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBefore struct {
	Owners []string                                                                                                                                           `json:"owners,required"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBeforeJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBeforeJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBefore]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBeforeJSON struct {
	Owners      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementBeforeJSON) RawJSON() string {
	return r.raw
}

// The type of the state change
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementTypeOwnershipChange EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementType = "OWNERSHIP_CHANGE"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseOwnershipChangeManagementTypeOwnershipChange:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement struct {
	// The state after the transaction
	After EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfter `json:"after,required"`
	// The state before the transaction
	Before EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBefore `json:"before,required"`
	// The type of the state change
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementType `json:"type,required"`
	JSON evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementJSON struct {
	After       apijson.Field
	Before      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagement) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement() {
}

// The state after the transaction
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfter struct {
	Modules []string                                                                                                                                        `json:"modules,required"`
	JSON    evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfterJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfterJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfter]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfterJSON struct {
	Modules     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementAfterJSON) RawJSON() string {
	return r.raw
}

// The state before the transaction
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBefore struct {
	Modules []string                                                                                                                                         `json:"modules,required"`
	JSON    evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBeforeJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBeforeJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBefore]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBeforeJSON struct {
	Modules     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementBeforeJSON) RawJSON() string {
	return r.raw
}

// The type of the state change
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementTypeModuleChange EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementType = "MODULE_CHANGE"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseModulesChangeManagementTypeModuleChange:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement struct {
	// The delegated address
	DelegatedAddress string `json:"delegated_address,required"`
	// The type of the state change
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementType `json:"type,required"`
	JSON evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementJSON struct {
	DelegatedAddress apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagement) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement() {
}

// The type of the state change
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementTypeSetCodeAccount EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementType = "SET_CODE_ACCOUNT"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseSetCodeAccountManagementTypeSetCodeAccount:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation struct {
	// The direct creator address of the new contract
	Deployer string `json:"deployer,required"`
	// The type of the state change
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationType `json:"type,required"`
	JSON evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationJSON struct {
	Deployer    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreation) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagement() {
}

// The type of the state change
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationTypeContractCreation EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationType = "CONTRACT_CREATION"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementRoutersEvmResponseContractCreationTypeContractCreation:
		return true
	}
	return false
}

// The type of the state change
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeProxyUpgrade     EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType = "PROXY_UPGRADE"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeOwnershipChange  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType = "OWNERSHIP_CHANGE"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeModuleChange     EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType = "MODULE_CHANGE"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeSetCodeAccount   EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType = "SET_CODE_ACCOUNT"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeContractCreation EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType = "CONTRACT_CREATION"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeProxyUpgrade, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeOwnershipChange, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeModuleChange, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeSetCodeAccount, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationContractManagementTypeContractCreation:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalance struct {
	// The asset that is missing balance
	Asset EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset `json:"asset,required"`
	// The account address's current balance of the asset
	CurrentBalance EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalance `json:"current_balance,required"`
	// The account address's missing balance of the asset
	MissingBalance EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalance `json:"missing_balance,required"`
	// The required balance of the asset for this action
	RequiredBalance EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalance `json:"required_balance,required"`
	JSON            evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalanceJSON             `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalanceJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalance]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalanceJSON struct {
	Asset           apijson.Field
	CurrentBalance  apijson.Field
	MissingBalance  apijson.Field
	RequiredBalance apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalanceJSON) RawJSON() string {
	return r.raw
}

// The asset that is missing balance
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset struct {
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetType `json:"type,required"`
	// address of the token
	Address   string `json:"address"`
	ChainID   int64  `json:"chain_id"`
	ChainName string `json:"chain_name"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                              `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetJSON `json:"-"`
	union  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetJSON struct {
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

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetUnion {
	return r.union
}

// The asset that is missing balance
//
// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                     `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsType = "ERC20"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsErc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails struct {
	ChainID   int64  `json:"chain_id,required"`
	ChainName string `json:"chain_name,required"`
	Decimals  int64  `json:"decimals,required"`
	LogoURL   string `json:"logo_url,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsType `json:"type,required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                      `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsJSON struct {
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

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAsset() {
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsTypeNative EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsType = "NATIVE"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetRoutersEvmTokenDetailsNativeAssetDetailsTypeNative:
		return true
	}
	return false
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetTypeErc20  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetType = "ERC20"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetTypeNative EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetType = "NATIVE"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetTypeErc20, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesAssetTypeNative:
		return true
	}
	return false
}

// The account address's current balance of the asset
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalance struct {
	// The raw value of the balance in hex string format
	RawValue string `json:"raw_value,required"`
	// The value of the balance in decimal string format
	Value string                                                                                                       `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalanceJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalanceJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalance]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalanceJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesCurrentBalanceJSON) RawJSON() string {
	return r.raw
}

// The account address's missing balance of the asset
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalance struct {
	// The raw value of the balance in hex string format
	RawValue string `json:"raw_value,required"`
	// The value of the balance in decimal string format
	Value string                                                                                                       `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalanceJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalanceJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalance]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalanceJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesMissingBalanceJSON) RawJSON() string {
	return r.raw
}

// The required balance of the asset for this action
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalance struct {
	// The raw value of the balance in hex string format
	RawValue string `json:"raw_value,required"`
	// The value of the balance in decimal string format
	Value string                                                                                                        `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalanceJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalanceJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalance]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalanceJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationMissingBalancesRequiredBalanceJSON) RawJSON() string {
	return r.raw
}

// The parameters of the transaction that was simulated.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationParams struct {
	// The block tag to be sent.
	BlockTag string `json:"block_tag"`
	// The calldata to be sent.
	Calldata EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldata `json:"calldata"`
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
	UserOperationCalldata EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldata `json:"user_operation_calldata"`
	// The value to be sent.
	Value string                                                                                `json:"value"`
	JSON  evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationParams]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsJSON struct {
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

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsJSON) RawJSON() string {
	return r.raw
}

// The calldata to be sent.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldata struct {
	// The function selector of the function called in the transaction
	FunctionSelector string `json:"function_selector,required"`
	// The function declaration of the function called in the transaction
	FunctionDeclaration string `json:"function_declaration"`
	// The function signature of the function called in the transaction
	FunctionSignature string                                                                                        `json:"function_signature"`
	JSON              evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldataJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldataJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldata]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldataJSON struct {
	FunctionSelector    apijson.Field
	FunctionDeclaration apijson.Field
	FunctionSignature   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsCalldataJSON) RawJSON() string {
	return r.raw
}

// The user operation call data to be sent.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldata struct {
	// The function selector of the function called in the transaction
	FunctionSelector string `json:"function_selector,required"`
	// The function declaration of the function called in the transaction
	FunctionDeclaration string `json:"function_declaration"`
	// The function signature of the function called in the transaction
	FunctionSignature string                                                                                                     `json:"function_signature"`
	JSON              evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldataJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldataJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldata]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldataJSON struct {
	FunctionSelector    apijson.Field
	FunctionDeclaration apijson.Field
	FunctionSignature   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationParamsUserOperationCalldataJSON) RawJSON() string {
	return r.raw
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationError struct {
	// A string explaining why the transaction failed
	Description string `json:"description,required"`
	// An error message if the simulation failed.
	Error string `json:"error,required"`
	// A string indicating if the simulation was successful or not.
	Status EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatus `json:"status,required"`
	// Error details if the simulation failed.
	ErrorDetails EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails `json:"error_details"`
	JSON         evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorJSON         `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationError]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorJSON struct {
	Description  apijson.Field
	Error        apijson.Field
	Status       apijson.Field
	ErrorDetails apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationError) implementsEvmTransactionScanResponseSimulation() {
}

// A string indicating if the simulation was successful or not.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatus string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatusError EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatus = "Error"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatus) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorStatusError:
		return true
	}
	return false
}

// Error details if the simulation failed.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails struct {
	// The type of the model
	Code string `json:"code,required"`
	// The address of the account
	AccountAddress string `json:"account_address"`
	// The address that is invalid
	Address string `json:"address"`
	// This field can have the runtime type of
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset].
	Asset interface{} `json:"asset"`
	// The current balance of the account
	CurrentBalance string `json:"current_balance"`
	// The domain name that is unsupported
	DomainName string `json:"domain_name"`
	// The message type that is unsupported
	MessageType string `json:"message_type"`
	// The required balance of the account
	RequiredBalance string                                                                                           `json:"required_balance"`
	JSON            evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsJSON `json:"-"`
	union           EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsJSON struct {
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

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsUnion {
	return r.union
}

// Error details if the simulation failed.
//
// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails struct {
	// The address of the account
	AccountAddress string `json:"account_address,required"`
	// The asset that the account does not have enough balance for
	Asset EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset `json:"asset,required"`
	// The type of the model
	Code EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCode `json:"code,required"`
	// The current balance of the account
	CurrentBalance string `json:"current_balance"`
	// The required balance of the account
	RequiredBalance string                                                                                                                                                 `json:"required_balance"`
	JSON            evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsJSON struct {
	AccountAddress  apijson.Field
	Asset           apijson.Field
	Code            apijson.Field
	CurrentBalance  apijson.Field
	RequiredBalance apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails() {
}

// The asset that the account does not have enough balance for
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset struct {
	// This field can have the runtime type of
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetails],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetails],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetails],
	// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetails].
	Details interface{} `json:"details,required"`
	// The type of the model
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType `json:"type,required"`
	// Token Id
	TokenID int64                                                                                                                                                       `json:"token_id"`
	JSON    evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetJSON `json:"-"`
	union   EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetUnion
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetJSON struct {
	Details     apijson.Field
	Type        apijson.Field
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset].
func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset) AsUnion() EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetUnion {
	return r.union
}

// The asset that the account does not have enough balance for
//
// Union satisfied by
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset],
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset]
// or
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset].
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetUnion interface {
	implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset{}),
		},
	)
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset struct {
	// Details
	Details EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetails `json:"details,required"`
	// The type of the model
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetType `json:"type,required"`
	JSON evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetJSON struct {
	Details     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAsset) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset() {
}

// Details
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetails struct {
	ChainID   int64  `json:"chain_id,required"`
	ChainName string `json:"chain_name,required"`
	Decimals  int64  `json:"decimals,required"`
	LogoURL   string `json:"logo_url,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsType `json:"type,required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                          `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsJSON struct {
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

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsTypeNative EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsType = "NATIVE"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetDetailsTypeNative:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetTypeNative EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetType = "NATIVE"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseNativeAssetTypeNative:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset struct {
	// Details
	Details EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetails `json:"details,required"`
	// The type of the model
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetType `json:"type,required"`
	JSON evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetJSON struct {
	Details     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20Asset) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset() {
}

// Details
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                         `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsTypeErc20 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsType = "ERC20"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetDetailsTypeErc20:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetTypeErc20 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetType = "ERC20"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc20AssetTypeErc20:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset struct {
	// Details
	Details EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetails `json:"details,required"`
	// Token Id
	TokenID int64 `json:"token_id,required"`
	// The type of the model
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetType `json:"type,required"`
	JSON evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetJSON struct {
	Details     apijson.Field
	TokenID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721Asset) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset() {
}

// Details
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                          `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsTypeErc1155 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsType = "ERC1155"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetDetailsTypeErc1155:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetTypeErc721 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetType = "ERC721"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc721AssetTypeErc721:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset struct {
	// Details
	Details EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetails `json:"details,required"`
	// Token Id
	TokenID int64 `json:"token_id,required"`
	// The type of the model
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetType `json:"type,required"`
	JSON evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetJSON struct {
	Details     apijson.Field
	TokenID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155Asset) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAsset() {
}

// Details
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                                                                                                                           `json:"symbol"`
	JSON   evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsJSON) RawJSON() string {
	return r.raw
}

// asset type.
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsTypeErc1155 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsType = "ERC1155"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetDetailsTypeErc1155:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetTypeErc1155 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetType = "ERC1155"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetRoutersEvmResponseErc1155AssetTypeErc1155:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeNative  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType = "NATIVE"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc20   EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType = "ERC20"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc721  EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType = "ERC721"
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc1155 EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType = "ERC1155"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeNative, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc20, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc721, EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsAssetTypeErc1155:
		return true
	}
	return false
}

// The type of the model
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCode string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCodeGeneralInsufficientFunds EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCode = "GENERAL_INSUFFICIENT_FUNDS"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCode) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInsufficientFundsErrorDetailsCodeGeneralInsufficientFunds:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails struct {
	// The address that is invalid
	Address string `json:"address,required"`
	// The type of the model
	Code EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCode `json:"code,required"`
	JSON evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsJSON struct {
	Address     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails() {
}

// The type of the model
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCode string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCodeGeneralInvalidAddress EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCode = "GENERAL_INVALID_ADDRESS"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCode) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGeneralInvalidAddressErrorDetailsCodeGeneralInvalidAddress:
		return true
	}
	return false
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails struct {
	// The error code
	Code string                                                                                                                                `json:"code,required"`
	JSON evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetailsJSON struct {
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseGenericErrorDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails() {
}

type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails struct {
	// The type of the model
	Code EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCode `json:"code,required"`
	// The domain name that is unsupported
	DomainName string `json:"domain_name,required"`
	// The message type that is unsupported
	MessageType string                                                                                                                                                 `json:"message_type,required"`
	JSON        evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsJSON `json:"-"`
}

// evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails]
type evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsJSON struct {
	Code        apijson.Field
	DomainName  apijson.Field
	MessageType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetails) implementsEvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetails() {
}

// The type of the model
type EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCode string

const (
	EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCodeUnsupportedEip712Message EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCode = "UNSUPPORTED_EIP712_MESSAGE"
)

func (r EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCode) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationRoutersEvmResponseTransactionSimulationErrorErrorDetailsRoutersEvmResponseUnsupportedEip712MessageErrorDetailsCodeUnsupportedEip712Message:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type EvmTransactionScanResponseSimulationStatus string

const (
	EvmTransactionScanResponseSimulationStatusSuccess EvmTransactionScanResponseSimulationStatus = "Success"
	EvmTransactionScanResponseSimulationStatusError   EvmTransactionScanResponseSimulationStatus = "Error"
)

func (r EvmTransactionScanResponseSimulationStatus) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseSimulationStatusSuccess, EvmTransactionScanResponseSimulationStatusError:
		return true
	}
	return false
}

type EvmTransactionScanResponseUserOperationGasEstimation struct {
	Status                           EvmTransactionScanResponseUserOperationGasEstimationStatus `json:"status,required"`
	CallGasEstimate                  string                                                     `json:"call_gas_estimate"`
	Error                            string                                                     `json:"error"`
	PaymasterVerificationGasEstimate string                                                     `json:"paymaster_verification_gas_estimate"`
	PreVerificationGasEstimate       string                                                     `json:"pre_verification_gas_estimate"`
	VerificationGasEstimate          string                                                     `json:"verification_gas_estimate"`
	JSON                             evmTransactionScanResponseUserOperationGasEstimationJSON   `json:"-"`
	union                            EvmTransactionScanResponseUserOperationGasEstimationUnion
}

// evmTransactionScanResponseUserOperationGasEstimationJSON contains the JSON
// metadata for the struct [EvmTransactionScanResponseUserOperationGasEstimation]
type evmTransactionScanResponseUserOperationGasEstimationJSON struct {
	Status                           apijson.Field
	CallGasEstimate                  apijson.Field
	Error                            apijson.Field
	PaymasterVerificationGasEstimate apijson.Field
	PreVerificationGasEstimate       apijson.Field
	VerificationGasEstimate          apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r evmTransactionScanResponseUserOperationGasEstimationJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseUserOperationGasEstimation) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseUserOperationGasEstimation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EvmTransactionScanResponseUserOperationGasEstimationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV6GasEstimation],
// [EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV7GasEstimation],
// [EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError].
func (r EvmTransactionScanResponseUserOperationGasEstimation) AsUnion() EvmTransactionScanResponseUserOperationGasEstimationUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV6GasEstimation],
// [EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV7GasEstimation]
// or
// [EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError].
type EvmTransactionScanResponseUserOperationGasEstimationUnion interface {
	implementsEvmTransactionScanResponseUserOperationGasEstimation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseUserOperationGasEstimationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV6GasEstimation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV7GasEstimation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError{}),
		},
	)
}

type EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV6GasEstimation struct {
	CallGasEstimate            string                                                                                                 `json:"call_gas_estimate,required"`
	PreVerificationGasEstimate string                                                                                                 `json:"pre_verification_gas_estimate,required"`
	Status                     EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV6GasEstimationStatus `json:"status,required"`
	VerificationGasEstimate    string                                                                                                 `json:"verification_gas_estimate,required"`
	JSON                       evmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV6GasEstimationJSON   `json:"-"`
}

// evmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV6GasEstimationJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV6GasEstimation]
type evmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV6GasEstimationJSON struct {
	CallGasEstimate            apijson.Field
	PreVerificationGasEstimate apijson.Field
	Status                     apijson.Field
	VerificationGasEstimate    apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV6GasEstimation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV6GasEstimationJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV6GasEstimation) implementsEvmTransactionScanResponseUserOperationGasEstimation() {
}

type EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV6GasEstimationStatus string

const (
	EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV6GasEstimationStatusSuccess EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV6GasEstimationStatus = "Success"
)

func (r EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV6GasEstimationStatus) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV6GasEstimationStatusSuccess:
		return true
	}
	return false
}

type EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV7GasEstimation struct {
	CallGasEstimate                  string                                                                                                 `json:"call_gas_estimate,required"`
	PaymasterVerificationGasEstimate string                                                                                                 `json:"paymaster_verification_gas_estimate,required"`
	PreVerificationGasEstimate       string                                                                                                 `json:"pre_verification_gas_estimate,required"`
	Status                           EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV7GasEstimationStatus `json:"status,required"`
	VerificationGasEstimate          string                                                                                                 `json:"verification_gas_estimate,required"`
	JSON                             evmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV7GasEstimationJSON   `json:"-"`
}

// evmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV7GasEstimationJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV7GasEstimation]
type evmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV7GasEstimationJSON struct {
	CallGasEstimate                  apijson.Field
	PaymasterVerificationGasEstimate apijson.Field
	PreVerificationGasEstimate       apijson.Field
	Status                           apijson.Field
	VerificationGasEstimate          apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r *EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV7GasEstimation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV7GasEstimationJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV7GasEstimation) implementsEvmTransactionScanResponseUserOperationGasEstimation() {
}

type EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV7GasEstimationStatus string

const (
	EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV7GasEstimationStatusSuccess EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV7GasEstimationStatus = "Success"
)

func (r EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV7GasEstimationStatus) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsUserOperationV7GasEstimationStatusSuccess:
		return true
	}
	return false
}

type EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError struct {
	Error  string                                                                                                      `json:"error,required"`
	Status EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus `json:"status,required"`
	JSON   evmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON   `json:"-"`
}

// evmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError]
type evmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationError) implementsEvmTransactionScanResponseUserOperationGasEstimation() {
}

type EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus string

const (
	EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatusError EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus = "Error"
)

func (r EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatus) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseUserOperationGasEstimationRoutersEvmModelsTransactionScanGasEstimationErrorStatusError:
		return true
	}
	return false
}

type EvmTransactionScanResponseUserOperationGasEstimationStatus string

const (
	EvmTransactionScanResponseUserOperationGasEstimationStatusSuccess EvmTransactionScanResponseUserOperationGasEstimationStatus = "Success"
	EvmTransactionScanResponseUserOperationGasEstimationStatusError   EvmTransactionScanResponseUserOperationGasEstimationStatus = "Error"
)

func (r EvmTransactionScanResponseUserOperationGasEstimationStatus) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseUserOperationGasEstimationStatusSuccess, EvmTransactionScanResponseUserOperationGasEstimationStatusError:
		return true
	}
	return false
}

type EvmTransactionScanResponseValidation struct {
	// This field can have the runtime type of
	// [[]EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationFeature],
	// [[]EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeature].
	Features interface{} `json:"features,required"`
	// An enumeration.
	ResultType EvmTransactionScanResponseValidationResultType `json:"result_type,required"`
	// A string indicating if the simulation was successful or not.
	Status EvmTransactionScanResponseValidationStatus `json:"status,required"`
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
	Reason string                                   `json:"reason"`
	JSON   evmTransactionScanResponseValidationJSON `json:"-"`
	union  EvmTransactionScanResponseValidationUnion
}

// evmTransactionScanResponseValidationJSON contains the JSON metadata for the
// struct [EvmTransactionScanResponseValidation]
type evmTransactionScanResponseValidationJSON struct {
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

func (r evmTransactionScanResponseValidationJSON) RawJSON() string {
	return r.raw
}

func (r *EvmTransactionScanResponseValidation) UnmarshalJSON(data []byte) (err error) {
	*r = EvmTransactionScanResponseValidation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EvmTransactionScanResponseValidationUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidation],
// [EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationError].
func (r EvmTransactionScanResponseValidation) AsUnion() EvmTransactionScanResponseValidationUnion {
	return r.union
}

// Union satisfied by
// [EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidation] or
// [EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationError].
type EvmTransactionScanResponseValidationUnion interface {
	implementsEvmTransactionScanResponseValidation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvmTransactionScanResponseValidationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationError{}),
		},
	)
}

type EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidation struct {
	// A list of features about this transaction explaining the validation.
	Features []EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationFeature `json:"features,required"`
	// An enumeration.
	ResultType EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationResultType `json:"result_type,required"`
	// A string indicating if the simulation was successful or not.
	Status EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationStatus `json:"status,required"`
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification"`
	// A textual description that can be presented to the user about what this
	// transaction is doing.
	Description string `json:"description"`
	// A textual description about the reasons the transaction was flagged with
	// result_type.
	Reason string                                                                          `json:"reason"`
	JSON   evmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationJSON `json:"-"`
}

// evmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidation]
type evmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationJSON struct {
	Features       apijson.Field
	ResultType     apijson.Field
	Status         apijson.Field
	Classification apijson.Field
	Description    apijson.Field
	Reason         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidation) implementsEvmTransactionScanResponseValidation() {
}

type EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationFeature struct {
	// Textual description
	Description string `json:"description,required"`
	// Feature name
	FeatureID string `json:"feature_id,required"`
	// An enumeration.
	Type EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType `json:"type,required"`
	// Address the feature refers to
	Address string `json:"address"`
	// Metadata related to the feature
	Metadata interface{}                                                                            `json:"metadata"`
	JSON     evmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationFeatureJSON `json:"-"`
}

// evmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationFeatureJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationFeature]
type evmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationFeatureJSON struct {
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	Address     apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationFeatureJSON) RawJSON() string {
	return r.raw
}

// An enumeration.
type EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType string

const (
	EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeMalicious EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType = "Malicious"
	EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeWarning   EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType = "Warning"
	EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeBenign    EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType = "Benign"
	EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeInfo      EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType = "Info"
)

func (r EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeMalicious, EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeWarning, EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeBenign, EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationFeaturesTypeInfo:
		return true
	}
	return false
}

// An enumeration.
type EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationResultType string

const (
	EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeBenign    EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationResultType = "Benign"
	EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeWarning   EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationResultType = "Warning"
	EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeMalicious EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationResultType = "Malicious"
)

func (r EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationResultType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeBenign, EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeWarning, EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationResultTypeMalicious:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationStatus string

const (
	EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationStatusSuccess EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationStatus = "Success"
)

func (r EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationStatus) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationStatusSuccess:
		return true
	}
	return false
}

type EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationError struct {
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassification `json:"classification,required"`
	// A textual description that can be presented to the user about what this
	// transaction is doing.
	Description EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescription `json:"description,required"`
	// An error message if the validation failed.
	Error string `json:"error,required"`
	// A list of features about this transaction explaining the validation.
	Features []EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeature `json:"features,required"`
	// A textual description about the reasons the transaction was flagged with
	// result_type.
	Reason EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorReason `json:"reason,required"`
	// A string indicating if the transaction is safe to sign or not.
	ResultType EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultType `json:"result_type,required"`
	// A string indicating if the simulation was successful or not.
	Status EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatus `json:"status,required"`
	JSON   evmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorJSON   `json:"-"`
}

// evmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationError]
type evmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorJSON struct {
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

func (r *EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorJSON) RawJSON() string {
	return r.raw
}

func (r EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationError) implementsEvmTransactionScanResponseValidation() {
}

// A textual classification that can be presented to the user explaining the
// reason.
type EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassification string

const (
	EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassificationEmpty EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassification = ""
)

func (r EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassification) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorClassificationEmpty:
		return true
	}
	return false
}

// A textual description that can be presented to the user about what this
// transaction is doing.
type EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescription string

const (
	EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescriptionEmpty EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescription = ""
)

func (r EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescription) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorDescriptionEmpty:
		return true
	}
	return false
}

type EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeature struct {
	// Textual description
	Description string `json:"description,required"`
	// Feature name
	FeatureID string `json:"feature_id,required"`
	// An enumeration.
	Type EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType `json:"type,required"`
	// Address the feature refers to
	Address string `json:"address"`
	// Metadata related to the feature
	Metadata interface{}                                                                                 `json:"metadata"`
	JSON     evmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeatureJSON `json:"-"`
}

// evmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeatureJSON
// contains the JSON metadata for the struct
// [EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeature]
type evmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeatureJSON struct {
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	Address     apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeatureJSON) RawJSON() string {
	return r.raw
}

// An enumeration.
type EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType string

const (
	EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeMalicious EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType = "Malicious"
	EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeWarning   EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType = "Warning"
	EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeBenign    EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType = "Benign"
	EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeInfo      EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType = "Info"
)

func (r EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeMalicious, EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeWarning, EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeBenign, EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorFeaturesTypeInfo:
		return true
	}
	return false
}

// A textual description about the reasons the transaction was flagged with
// result_type.
type EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorReason string

const (
	EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorReasonEmpty EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorReason = ""
)

func (r EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorReason) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorReasonEmpty:
		return true
	}
	return false
}

// A string indicating if the transaction is safe to sign or not.
type EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultType string

const (
	EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultTypeError EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultType = "Error"
)

func (r EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorResultTypeError:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatus string

const (
	EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatusSuccess EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatus = "Success"
)

func (r EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatus) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseValidationRoutersEvmResponseTransactionValidationErrorStatusSuccess:
		return true
	}
	return false
}

// An enumeration.
type EvmTransactionScanResponseValidationResultType string

const (
	EvmTransactionScanResponseValidationResultTypeBenign    EvmTransactionScanResponseValidationResultType = "Benign"
	EvmTransactionScanResponseValidationResultTypeWarning   EvmTransactionScanResponseValidationResultType = "Warning"
	EvmTransactionScanResponseValidationResultTypeMalicious EvmTransactionScanResponseValidationResultType = "Malicious"
	EvmTransactionScanResponseValidationResultTypeError     EvmTransactionScanResponseValidationResultType = "Error"
)

func (r EvmTransactionScanResponseValidationResultType) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseValidationResultTypeBenign, EvmTransactionScanResponseValidationResultTypeWarning, EvmTransactionScanResponseValidationResultTypeMalicious, EvmTransactionScanResponseValidationResultTypeError:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type EvmTransactionScanResponseValidationStatus string

const (
	EvmTransactionScanResponseValidationStatusSuccess EvmTransactionScanResponseValidationStatus = "Success"
)

func (r EvmTransactionScanResponseValidationStatus) IsKnown() bool {
	switch r {
	case EvmTransactionScanResponseValidationStatusSuccess:
		return true
	}
	return false
}

type EvmTransactionReportParams struct {
	// Details about the report.
	Details param.Field[string] `json:"details,required"`
	// The event type of the report. Could be `FALSE_POSITIVE` or `FALSE_NEGATIVE`.
	Event param.Field[EvmTransactionReportParamsEvent] `json:"event,required"`
	// Parameters identifying the transaction to report, provided either as transaction
	// details (chain, account address, and transaction data) or as a request ID from a
	// previous scan.
	Report param.Field[EvmTransactionReportParamsReportUnion] `json:"report,required"`
}

func (r EvmTransactionReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The event type of the report. Could be `FALSE_POSITIVE` or `FALSE_NEGATIVE`.
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

// Parameters identifying the transaction to report, provided either as transaction
// details (chain, account address, and transaction data) or as a request ID from a
// previous scan.
type EvmTransactionReportParamsReport struct {
	Type   param.Field[EvmTransactionReportParamsReportType] `json:"type,required"`
	Params param.Field[interface{}]                          `json:"params"`
	// The request ID of a previous request. This can be found in the value of the
	// `x-request-id` field in the headers of the response of the previous request. For
	// instance: `6c3cf6c1-a80d-4927-91b9-03d841ea61fe`.
	RequestID param.Field[string] `json:"request_id"`
}

func (r EvmTransactionReportParamsReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmTransactionReportParamsReport) implementsEvmTransactionReportParamsReportUnion() {}

// Parameters identifying the transaction to report, provided either as transaction
// details (chain, account address, and transaction data) or as a request ID from a
// previous scan.
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
	Metadata param.Field[EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataUnion] `json:"metadata,required"`
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
	// An enumeration.
	Method param.Field[EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethod] `json:"method"`
	Params param.Field[interface{}]                                                                        `json:"params"`
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
// [EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataRoutersEvmTransactionModelsTransaction],
// [EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpc],
// [EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsData].
type EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataUnion interface {
	implementsEvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataUnion()
}

type EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataRoutersEvmTransactionModelsTransaction struct {
	// The source address of the transaction in hex string format
	From param.Field[string] `json:"from,required"`
	// The authorization list
	AuthorizationList param.Field[[]EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataRoutersEvmTransactionModelsTransactionAuthorizationList] `json:"authorization_list"`
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

func (r EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataRoutersEvmTransactionModelsTransaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataRoutersEvmTransactionModelsTransaction) implementsEvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataUnion() {
}

type EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataRoutersEvmTransactionModelsTransactionAuthorizationList struct {
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
	// The r value as hex string
	R param.Field[string] `json:"r"`
	// The s value as hex string
	S param.Field[string] `json:"s"`
	// The yParity value as hex string
	YParity param.Field[string] `json:"yParity"`
}

func (r EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataRoutersEvmTransactionModelsTransactionAuthorizationList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpc struct {
	// An enumeration.
	Method param.Field[EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethod] `json:"method,required"`
	// The parameters of the JSON-RPC request in JSON format
	Params param.Field[[]interface{}] `json:"params,required"`
}

func (r EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpc) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpc) implementsEvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataUnion() {
}

// An enumeration.
type EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethod string

const (
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethodEthSendTransaction    EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethod = "eth_sendTransaction"
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethodEthSendRawTransaction EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethod = "eth_sendRawTransaction"
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethodEthSignTransaction    EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethod = "eth_signTransaction"
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethodEthSignTypedData      EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethod = "eth_signTypedData"
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethodEthSignTypedDataV1    EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethod = "eth_signTypedData_v1"
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethodEthSignTypedDataV2    EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethod = "eth_signTypedData_v2"
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethodEthSignTypedDataV3    EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethod = "eth_signTypedData_v3"
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethodEthSignTypedDataV4    EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethod = "eth_signTypedData_v4"
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethodEthSendUserOperation  EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethod = "eth_sendUserOperation"
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethodPersonalSign          EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethod = "personal_sign"
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethodEthSign               EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethod = "eth_sign"
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethodWalletSendCalls       EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethod = "wallet_sendCalls"
)

func (r EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethod) IsKnown() bool {
	switch r {
	case EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethodEthSendTransaction, EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethodEthSendRawTransaction, EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethodEthSignTransaction, EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethodEthSignTypedData, EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethodEthSignTypedDataV1, EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethodEthSignTypedDataV2, EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethodEthSignTypedDataV3, EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethodEthSignTypedDataV4, EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethodEthSendUserOperation, EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethodPersonalSign, EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethodEthSign, EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataJsonRpcMethodWalletSendCalls:
		return true
	}
	return false
}

// An enumeration.
type EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethod string

const (
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethodEthSendTransaction    EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethod = "eth_sendTransaction"
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethodEthSendRawTransaction EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethod = "eth_sendRawTransaction"
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethodEthSignTransaction    EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethod = "eth_signTransaction"
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethodEthSignTypedData      EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethod = "eth_signTypedData"
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethodEthSignTypedDataV1    EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethod = "eth_signTypedData_v1"
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethodEthSignTypedDataV2    EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethod = "eth_signTypedData_v2"
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethodEthSignTypedDataV3    EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethod = "eth_signTypedData_v3"
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethodEthSignTypedDataV4    EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethod = "eth_signTypedData_v4"
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethodEthSendUserOperation  EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethod = "eth_sendUserOperation"
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethodPersonalSign          EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethod = "personal_sign"
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethodEthSign               EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethod = "eth_sign"
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethodWalletSendCalls       EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethod = "wallet_sendCalls"
)

func (r EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethod) IsKnown() bool {
	switch r {
	case EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethodEthSendTransaction, EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethodEthSendRawTransaction, EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethodEthSignTransaction, EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethodEthSignTypedData, EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethodEthSignTypedDataV1, EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethodEthSignTypedDataV2, EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethodEthSignTypedDataV3, EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethodEthSignTypedDataV4, EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethodEthSendUserOperation, EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethodPersonalSign, EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethodEthSign, EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsDataMethodWalletSendCalls:
		return true
	}
	return false
}

// Object of additional information to validate against.
type EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadata struct {
	// The full URL of the DApp or website that initiated the transaction, for
	// cross-reference. Must use the https or http scheme and contain a valid hostname.
	// Cannot contain JSON, braces, or other embedded data structures.
	Domain param.Field[string] `json:"domain"`
	// Indicates that the transaction was not initiated by a dapp.
	NonDapp param.Field[EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataNonDapp] `json:"non_dapp"`
}

func (r EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadata) implementsEvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataUnion() {
}

// Object of additional information to validate against.
//
// Satisfied by
// [EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataRoutersEvmModelsMetadataNonDapp],
// [EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataRoutersEvmModelsMetadataDapp],
// [EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadata].
type EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataUnion interface {
	implementsEvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataUnion()
}

type EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataRoutersEvmModelsMetadataNonDapp struct {
	// Indicates that the transaction was not initiated by a dapp.
	NonDapp param.Field[EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataRoutersEvmModelsMetadataNonDappNonDapp] `json:"non_dapp"`
}

func (r EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataRoutersEvmModelsMetadataNonDapp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataRoutersEvmModelsMetadataNonDapp) implementsEvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataUnion() {
}

// Indicates that the transaction was not initiated by a dapp.
type EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataRoutersEvmModelsMetadataNonDappNonDapp bool

const (
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataRoutersEvmModelsMetadataNonDappNonDappTrue EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataRoutersEvmModelsMetadataNonDappNonDapp = true
)

func (r EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataRoutersEvmModelsMetadataNonDappNonDapp) IsKnown() bool {
	switch r {
	case EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataRoutersEvmModelsMetadataNonDappNonDappTrue:
		return true
	}
	return false
}

type EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataRoutersEvmModelsMetadataDapp struct {
	// The full URL of the DApp or website that initiated the transaction, for
	// cross-reference. Must use the https or http scheme and contain a valid hostname.
	// Cannot contain JSON, braces, or other embedded data structures.
	Domain param.Field[string] `json:"domain,required"`
}

func (r EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataRoutersEvmModelsMetadataDapp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataRoutersEvmModelsMetadataDapp) implementsEvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataUnion() {
}

// Indicates that the transaction was not initiated by a dapp.
type EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataNonDapp bool

const (
	EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataNonDappTrue EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataNonDapp = true
)

func (r EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataNonDapp) IsKnown() bool {
	switch r {
	case EvmTransactionReportParamsReportParamReportTransactionReportParamsParamsMetadataNonDappTrue:
		return true
	}
	return false
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
	// The request ID of a previous request. This can be found in the value of the
	// `x-request-id` field in the headers of the response of the previous request. For
	// instance: `6c3cf6c1-a80d-4927-91b9-03d841ea61fe`.
	RequestID param.Field[string] `json:"request_id,required"`
	// The type identifier indicating that a request ID from a previous scan is being
	// used.
	Type param.Field[EvmTransactionReportParamsReportRequestIDReportType] `json:"type,required"`
}

func (r EvmTransactionReportParamsReportRequestIDReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmTransactionReportParamsReportRequestIDReport) implementsEvmTransactionReportParamsReportUnion() {
}

// The type identifier indicating that a request ID from a previous scan is being
// used.
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
	Metadata param.Field[EvmTransactionScanParamsMetadataUnion] `json:"metadata,required"`
	// The relative block for the block validation. Can be "latest" or a block number.
	Block param.Field[EvmTransactionScanParamsBlockUnion] `json:"block"`
	// List of one or more of options for the desired output. "simulation" - include
	// simulation output in your response. "validation" - include security validation
	// of the transaction in your response. "gas_estimation" - include gas estimation
	// result in your response. Default is ["validation"]
	Options param.Field[[]EvmTransactionScanParamsOption] `json:"options"`
	// Simulate transactions using gas estimation result. This requires
	// "gas_estimation" option to be enabled.
	SimulateWithEstimatedGas param.Field[bool] `json:"simulate_with_estimated_gas"`
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
	// The chain ID as hex string
	ChainID param.Field[string] `json:"chainId"`
	// The authority address of the delegation, should be provided when the signature
	// (r,s,yParity) is not provided in order to simulate the transaction with the
	// correct delegation
	Eoa param.Field[string] `json:"eoa"`
	// The nonce value as hex string
	Nonce param.Field[string] `json:"nonce"`
	// The r value as hex string
	R param.Field[string] `json:"r"`
	// The s value as hex string
	S param.Field[string] `json:"s"`
	// The yParity value as hex string
	YParity param.Field[string] `json:"yParity"`
}

func (r EvmTransactionScanParamsDataAuthorizationList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Object of additional information to validate against.
type EvmTransactionScanParamsMetadata struct {
	// The full URL of the DApp or website that initiated the transaction, for
	// cross-reference. Must use the https or http scheme and contain a valid hostname.
	// Cannot contain JSON, braces, or other embedded data structures.
	Domain param.Field[string] `json:"domain"`
	// Indicates that the transaction was not initiated by a dapp.
	NonDapp param.Field[EvmTransactionScanParamsMetadataNonDapp] `json:"non_dapp"`
}

func (r EvmTransactionScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmTransactionScanParamsMetadata) implementsEvmTransactionScanParamsMetadataUnion() {}

// Object of additional information to validate against.
//
// Satisfied by [EvmTransactionScanParamsMetadataRoutersEvmModelsMetadataNonDapp],
// [EvmTransactionScanParamsMetadataRoutersEvmModelsMetadataDapp],
// [EvmTransactionScanParamsMetadata].
type EvmTransactionScanParamsMetadataUnion interface {
	implementsEvmTransactionScanParamsMetadataUnion()
}

type EvmTransactionScanParamsMetadataRoutersEvmModelsMetadataNonDapp struct {
	// Indicates that the transaction was not initiated by a dapp.
	NonDapp param.Field[EvmTransactionScanParamsMetadataRoutersEvmModelsMetadataNonDappNonDapp] `json:"non_dapp"`
}

func (r EvmTransactionScanParamsMetadataRoutersEvmModelsMetadataNonDapp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmTransactionScanParamsMetadataRoutersEvmModelsMetadataNonDapp) implementsEvmTransactionScanParamsMetadataUnion() {
}

// Indicates that the transaction was not initiated by a dapp.
type EvmTransactionScanParamsMetadataRoutersEvmModelsMetadataNonDappNonDapp bool

const (
	EvmTransactionScanParamsMetadataRoutersEvmModelsMetadataNonDappNonDappTrue EvmTransactionScanParamsMetadataRoutersEvmModelsMetadataNonDappNonDapp = true
)

func (r EvmTransactionScanParamsMetadataRoutersEvmModelsMetadataNonDappNonDapp) IsKnown() bool {
	switch r {
	case EvmTransactionScanParamsMetadataRoutersEvmModelsMetadataNonDappNonDappTrue:
		return true
	}
	return false
}

type EvmTransactionScanParamsMetadataRoutersEvmModelsMetadataDapp struct {
	// The full URL of the DApp or website that initiated the transaction, for
	// cross-reference. Must use the https or http scheme and contain a valid hostname.
	// Cannot contain JSON, braces, or other embedded data structures.
	Domain param.Field[string] `json:"domain,required"`
}

func (r EvmTransactionScanParamsMetadataRoutersEvmModelsMetadataDapp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmTransactionScanParamsMetadataRoutersEvmModelsMetadataDapp) implementsEvmTransactionScanParamsMetadataUnion() {
}

// Indicates that the transaction was not initiated by a dapp.
type EvmTransactionScanParamsMetadataNonDapp bool

const (
	EvmTransactionScanParamsMetadataNonDappTrue EvmTransactionScanParamsMetadataNonDapp = true
)

func (r EvmTransactionScanParamsMetadataNonDapp) IsKnown() bool {
	switch r {
	case EvmTransactionScanParamsMetadataNonDappTrue:
		return true
	}
	return false
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
