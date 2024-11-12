// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"net/http"
	"reflect"
	"time"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/param"
	"github.com/blockaid-official/blockaid-client-go/internal/requestconfig"
	"github.com/blockaid-official/blockaid-client-go/option"
	"github.com/tidwall/gjson"
)

// StarknetTransactionService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStarknetTransactionService] method instead.
type StarknetTransactionService struct {
	Options []option.RequestOption
}

// NewStarknetTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStarknetTransactionService(opts ...option.RequestOption) (r *StarknetTransactionService) {
	r = &StarknetTransactionService{}
	r.Options = opts
	return
}

// Report Transaction
func (r *StarknetTransactionService) Report(ctx context.Context, body StarknetTransactionReportParams, opts ...option.RequestOption) (res *int64, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/starknet/transaction/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Scan Transactions
func (r *StarknetTransactionService) Scan(ctx context.Context, body StarknetTransactionScanParams, opts ...option.RequestOption) (res *StarknetTransactionScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/starknet/transaction/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type StarknetTransactionScanResponse struct {
	// Simulation result; Only present if simulation option is included in the request
	Simulation StarknetTransactionScanResponseSimulation `json:"simulation,nullable"`
	// Validation result; Only present if validation option is included in the request
	Validation StarknetTransactionScanResponseValidation `json:"validation,nullable"`
	JSON       starknetTransactionScanResponseJSON       `json:"-"`
}

// starknetTransactionScanResponseJSON contains the JSON metadata for the struct
// [StarknetTransactionScanResponse]
type starknetTransactionScanResponseJSON struct {
	Simulation  apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseJSON) RawJSON() string {
	return r.raw
}

// Simulation result; Only present if simulation option is included in the request
type StarknetTransactionScanResponseSimulation struct {
	Status StarknetTransactionScanResponseSimulationStatus `json:"status,required"`
	// This field can have the runtime type of
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummary].
	AccountSummary interface{} `json:"account_summary"`
	// This field can have the runtime type of
	// [[]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAddressDetail].
	AddressDetails interface{} `json:"address_details"`
	// This field can have the runtime type of
	// [map[string][]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff].
	AssetsDiffs interface{} `json:"assets_diffs"`
	// Optional block number or tag context for the simulation
	BlockNumber string `json:"block_number,nullable"`
	// Error message
	Error string `json:"error"`
	// This field can have the runtime type of
	// [map[string][]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure].
	Exposures interface{}                                   `json:"exposures"`
	JSON      starknetTransactionScanResponseSimulationJSON `json:"-"`
	union     StarknetTransactionScanResponseSimulationUnion
}

// starknetTransactionScanResponseSimulationJSON contains the JSON metadata for the
// struct [StarknetTransactionScanResponseSimulation]
type starknetTransactionScanResponseSimulationJSON struct {
	Status         apijson.Field
	AccountSummary apijson.Field
	AddressDetails apijson.Field
	AssetsDiffs    apijson.Field
	BlockNumber    apijson.Field
	Error          apijson.Field
	Exposures      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r starknetTransactionScanResponseSimulationJSON) RawJSON() string {
	return r.raw
}

func (r *StarknetTransactionScanResponseSimulation) UnmarshalJSON(data []byte) (err error) {
	*r = StarknetTransactionScanResponseSimulation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [StarknetTransactionScanResponseSimulationUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchema],
// [StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchema].
func (r StarknetTransactionScanResponseSimulation) AsUnion() StarknetTransactionScanResponseSimulationUnion {
	return r.union
}

// Simulation result; Only present if simulation option is included in the request
//
// Union satisfied by
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchema]
// or [StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchema].
type StarknetTransactionScanResponseSimulationUnion interface {
	implementsStarknetTransactionScanResponseSimulation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StarknetTransactionScanResponseSimulationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchema{}),
		},
	)
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchema struct {
	// Summary of the actions and asset transfers that were made by the requested
	// account address
	AccountSummary StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummary `json:"account_summary,required"`
	Status         StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaStatus         `json:"status,required"`
	// Details of addresses involved in the transaction
	AddressDetails []StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAddressDetail `json:"address_details"`
	// Mapping between the address of an account to the assets diff during the
	// transaction
	AssetsDiffs map[string][]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff `json:"assets_diffs"`
	// Optional block number or tag context for the simulation
	BlockNumber string `json:"block_number,nullable"`
	// Mapping between the address of an account to the exposure of the assets during
	// the transaction
	Exposures map[string][]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure `json:"exposures"`
	JSON      starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaJSON                  `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchema]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaJSON struct {
	AccountSummary apijson.Field
	Status         apijson.Field
	AddressDetails apijson.Field
	AssetsDiffs    apijson.Field
	BlockNumber    apijson.Field
	Exposures      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchema) implementsStarknetTransactionScanResponseSimulation() {
}

// Summary of the actions and asset transfers that were made by the requested
// account address
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummary struct {
	// Exposures made by the requested account address
	AccountExposures []StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure `json:"account_exposures,required"`
	// Total USD diff for the requested account address
	TotalUsdDiff StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryTotalUsdDiff `json:"total_usd_diff,required"`
	// Assets diffs of the requested account address
	AccountAssetsDiffs []StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiff `json:"account_assets_diffs"`
	// Total USD exposure for each of the spender addresses during the transaction
	TotalUsdExposure map[string]float64                                                                                `json:"total_usd_exposure"`
	JSON             starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummary]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryJSON struct {
	AccountExposures   apijson.Field
	TotalUsdDiff       apijson.Field
	AccountAssetsDiffs apijson.Field
	TotalUsdExposure   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure struct {
	// This field can have the runtime type of [StarknetAccountErc20Exposure],
	// [StarknetAccountErc721Exposure], [StarknetAccountErc1155Exposure].
	Asset interface{} `json:"asset,required"`
	// This field can have the runtime type of
	// [map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpender],
	// [map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpender],
	// [map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpender].
	Spenders interface{}                                                                                                      `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposureJSON `json:"-"`
	union    StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresUnion
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposureJSON) RawJSON() string {
	return r.raw
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure) UnmarshalJSON(data []byte) (err error) {
	*r = StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20Exposure],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721Exposure],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155Exposure].
func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure) AsUnion() StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresUnion {
	return r.union
}

// Union satisfied by
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20Exposure],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721Exposure]
// or
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155Exposure].
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresUnion interface {
	implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20Exposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721Exposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155Exposure{}),
		},
	)
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20Exposure struct {
	Asset StarknetAccountErc20Exposure `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpender `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureJSON               `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20Exposure]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20Exposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20Exposure) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpender struct {
	// Approval value of the ERC20 token
	Approval string              `json:"approval,required"`
	Exposure []StarknetErc20Diff `json:"exposure,required"`
	// Expiration date of the approval
	Expiration time.Time `json:"expiration,nullable" format:"date-time"`
	// Summarized description of the exposure
	Summary string                                                                                                                                        `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpenderJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpenderJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpender]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpenderJSON struct {
	Approval    apijson.Field
	Exposure    apijson.Field
	Expiration  apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721Exposure struct {
	Asset StarknetAccountErc721Exposure `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpender `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureJSON               `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721Exposure]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721Exposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721Exposure) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpender struct {
	Exposure []StarknetErc721Diff `json:"exposure,required"`
	// Whether `setApprovalForAll` was invoked
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// Summarized description of the exposure
	Summary string                                                                                                                                         `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpenderJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpenderJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpender]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155Exposure struct {
	Asset StarknetAccountErc1155Exposure `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpender `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureJSON               `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155Exposure]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155Exposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155Exposure) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpender struct {
	Exposure []StarknetErc1155Diff `json:"exposure,required"`
	// Whether `setApprovalForAll` was invoked
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// Summarized description of the exposure
	Summary string                                                                                                                                          `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpenderJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpenderJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpender]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpenderJSON) RawJSON() string {
	return r.raw
}

// Total USD diff for the requested account address
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryTotalUsdDiff struct {
	// Total incoming USD transfers
	In float64 `json:"in,required"`
	// Total outgoing USD transfers
	Out float64 `json:"out,required"`
	// Total USD transfers
	Total float64                                                                                                       `json:"total"`
	JSON  starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryTotalUsdDiffJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryTotalUsdDiffJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryTotalUsdDiff]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiff struct {
	// This field can have the runtime type of [StarknetAccountErc20Exposure],
	// [StarknetAccountErc721Exposure], [StarknetAccountErc1155Exposure].
	Asset interface{} `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of [StarknetErc20Diff],
	// [StarknetErc721Diff], [StarknetErc1155Diff].
	In interface{} `json:"in"`
	// This field can have the runtime type of [StarknetErc20Diff],
	// [StarknetErc721Diff], [StarknetErc1155Diff].
	Out   interface{}                                                                                                        `json:"out"`
	JSON  starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffJSON `json:"-"`
	union StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsUnion
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiff]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiff],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiff],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiff].
func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiff) AsUnion() StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiff],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiff]
// or
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiff].
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsUnion interface {
	implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiff{}),
		},
	)
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiff struct {
	Asset StarknetAccountErc20Exposure `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StarknetErc20Diff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetErc20Diff                                                                                                                         `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiff]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiff) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiff struct {
	Asset StarknetAccountErc721Exposure `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StarknetErc721Diff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetErc721Diff                                                                                                                         `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiff]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiff) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiff struct {
	Asset StarknetAccountErc1155Exposure `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StarknetErc1155Diff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetErc1155Diff                                                                                                                         `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiff]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiff) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaStatus string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaStatusSuccess StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaStatus = "Success"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaStatus) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaStatusSuccess:
		return true
	}
	return false
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAddressDetail struct {
	// Encoded public key of the account
	AccountAddress interface{} `json:"account_address,required"`
	// Description of the account
	Description string                                                                                           `json:"description,nullable"`
	JSON        starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAddressDetailJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAddressDetailJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAddressDetail]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAddressDetailJSON struct {
	AccountAddress apijson.Field
	Description    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAddressDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAddressDetailJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff struct {
	// This field can have the runtime type of [StarknetAccountErc20Exposure],
	// [StarknetAccountErc721Exposure], [StarknetAccountErc1155Exposure].
	Asset interface{} `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of [StarknetErc20Diff],
	// [StarknetErc721Diff], [StarknetErc1155Diff].
	In interface{} `json:"in"`
	// This field can have the runtime type of [StarknetErc20Diff],
	// [StarknetErc721Diff], [StarknetErc1155Diff].
	Out   interface{}                                                                                   `json:"out"`
	JSON  starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffJSON `json:"-"`
	union StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsUnion
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiff],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiff],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiff].
func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff) AsUnion() StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiff],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiff]
// or
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiff].
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsUnion interface {
	implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiff{}),
		},
	)
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiff struct {
	Asset StarknetAccountErc20Exposure `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StarknetErc20Diff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetErc20Diff                                                                                                    `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiff]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiff) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiff struct {
	Asset StarknetAccountErc721Exposure `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StarknetErc721Diff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetErc721Diff                                                                                                    `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiff]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiff) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiff struct {
	Asset StarknetAccountErc1155Exposure `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StarknetErc1155Diff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetErc1155Diff                                                                                                    `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiff]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiff) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure struct {
	// This field can have the runtime type of [StarknetAccountErc20Exposure],
	// [StarknetAccountErc721Exposure], [StarknetAccountErc1155Exposure].
	Asset interface{} `json:"asset,required"`
	// This field can have the runtime type of
	// [map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpender],
	// [map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpender],
	// [map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpender].
	Spenders interface{}                                                                                 `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposureJSON `json:"-"`
	union    StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresUnion
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposureJSON) RawJSON() string {
	return r.raw
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure) UnmarshalJSON(data []byte) (err error) {
	*r = StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20Exposure],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721Exposure],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155Exposure].
func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure) AsUnion() StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresUnion {
	return r.union
}

// Union satisfied by
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20Exposure],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721Exposure]
// or
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155Exposure].
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresUnion interface {
	implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20Exposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721Exposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155Exposure{}),
		},
	)
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20Exposure struct {
	Asset StarknetAccountErc20Exposure `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpender `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureJSON               `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20Exposure]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20Exposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20Exposure) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpender struct {
	// Approval value of the ERC20 token
	Approval string              `json:"approval,required"`
	Exposure []StarknetErc20Diff `json:"exposure,required"`
	// Expiration date of the approval
	Expiration time.Time `json:"expiration,nullable" format:"date-time"`
	// Summarized description of the exposure
	Summary string                                                                                                                   `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpenderJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpenderJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpender]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpenderJSON struct {
	Approval    apijson.Field
	Exposure    apijson.Field
	Expiration  apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721Exposure struct {
	Asset StarknetAccountErc721Exposure `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpender `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureJSON               `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721Exposure]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721Exposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721Exposure) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpender struct {
	Exposure []StarknetErc721Diff `json:"exposure,required"`
	// Whether `setApprovalForAll` was invoked
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// Summarized description of the exposure
	Summary string                                                                                                                    `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpenderJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpenderJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpender]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155Exposure struct {
	Asset StarknetAccountErc1155Exposure `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpender `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureJSON               `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155Exposure]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155Exposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155Exposure) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpender struct {
	Exposure []StarknetErc1155Diff `json:"exposure,required"`
	// Whether `setApprovalForAll` was invoked
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// Summarized description of the exposure
	Summary string                                                                                                                     `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpenderJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpenderJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpender]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchema struct {
	// Error message
	Error  string                                                                       `json:"error,required"`
	Status StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchemaStatus `json:"status,required"`
	JSON   starknetTransactionScanResponseSimulationStarknetSimulationErrorSchemaJSON   `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationErrorSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchema]
type starknetTransactionScanResponseSimulationStarknetSimulationErrorSchemaJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationErrorSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchema) implementsStarknetTransactionScanResponseSimulation() {
}

type StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchemaStatus string

const (
	StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchemaStatusError StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchemaStatus = "Error"
)

func (r StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchemaStatus) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchemaStatusError:
		return true
	}
	return false
}

type StarknetTransactionScanResponseSimulationStatus string

const (
	StarknetTransactionScanResponseSimulationStatusSuccess StarknetTransactionScanResponseSimulationStatus = "Success"
	StarknetTransactionScanResponseSimulationStatusError   StarknetTransactionScanResponseSimulationStatus = "Error"
)

func (r StarknetTransactionScanResponseSimulationStatus) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStatusSuccess, StarknetTransactionScanResponseSimulationStatusError:
		return true
	}
	return false
}

// Validation result; Only present if validation option is included in the request
type StarknetTransactionScanResponseValidation struct {
	Status StarknetTransactionScanResponseValidationStatus `json:"status,required"`
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification"`
	// A textual description about the validation result
	Description string `json:"description"`
	// Error message
	Error string `json:"error"`
	// This field can have the runtime type of
	// [[]StarknetTransactionScanResponseValidationStarknetValidationResultFeature].
	Features interface{} `json:"features"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason string `json:"reason"`
	// Verdict of the validation
	ResultType StarknetTransactionScanResponseValidationResultType `json:"result_type"`
	JSON       starknetTransactionScanResponseValidationJSON       `json:"-"`
	union      StarknetTransactionScanResponseValidationUnion
}

// starknetTransactionScanResponseValidationJSON contains the JSON metadata for the
// struct [StarknetTransactionScanResponseValidation]
type starknetTransactionScanResponseValidationJSON struct {
	Status         apijson.Field
	Classification apijson.Field
	Description    apijson.Field
	Error          apijson.Field
	Features       apijson.Field
	Reason         apijson.Field
	ResultType     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r starknetTransactionScanResponseValidationJSON) RawJSON() string {
	return r.raw
}

func (r *StarknetTransactionScanResponseValidation) UnmarshalJSON(data []byte) (err error) {
	*r = StarknetTransactionScanResponseValidation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [StarknetTransactionScanResponseValidationUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StarknetTransactionScanResponseValidationStarknetValidationResult],
// [StarknetTransactionScanResponseValidationStarknetValidationErrorSchema].
func (r StarknetTransactionScanResponseValidation) AsUnion() StarknetTransactionScanResponseValidationUnion {
	return r.union
}

// Validation result; Only present if validation option is included in the request
//
// Union satisfied by
// [StarknetTransactionScanResponseValidationStarknetValidationResult] or
// [StarknetTransactionScanResponseValidationStarknetValidationErrorSchema].
type StarknetTransactionScanResponseValidationUnion interface {
	implementsStarknetTransactionScanResponseValidation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StarknetTransactionScanResponseValidationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseValidationStarknetValidationResult{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseValidationStarknetValidationErrorSchema{}),
		},
	)
}

type StarknetTransactionScanResponseValidationStarknetValidationResult struct {
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification,required"`
	// A textual description about the validation result
	Description string                                                                     `json:"description,required"`
	Features    []StarknetTransactionScanResponseValidationStarknetValidationResultFeature `json:"features,required"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason string `json:"reason,required"`
	// Verdict of the validation
	ResultType StarknetTransactionScanResponseValidationStarknetValidationResultResultType `json:"result_type,required"`
	Status     StarknetTransactionScanResponseValidationStarknetValidationResultStatus     `json:"status,required"`
	JSON       starknetTransactionScanResponseValidationStarknetValidationResultJSON       `json:"-"`
}

// starknetTransactionScanResponseValidationStarknetValidationResultJSON contains
// the JSON metadata for the struct
// [StarknetTransactionScanResponseValidationStarknetValidationResult]
type starknetTransactionScanResponseValidationStarknetValidationResultJSON struct {
	Classification apijson.Field
	Description    apijson.Field
	Features       apijson.Field
	Reason         apijson.Field
	ResultType     apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseValidationStarknetValidationResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseValidationStarknetValidationResultJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseValidationStarknetValidationResult) implementsStarknetTransactionScanResponseValidation() {
}

type StarknetTransactionScanResponseValidationStarknetValidationResultFeature struct {
	// Address the feature refers to
	Address string `json:"address,required"`
	// Textual description
	Description string `json:"description,required"`
	FeatureID string `json:"feature_id,required"`
	// Feature Classification
	Type StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesType `json:"type,required"`
	JSON starknetTransactionScanResponseValidationStarknetValidationResultFeatureJSON  `json:"-"`
}

// starknetTransactionScanResponseValidationStarknetValidationResultFeatureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseValidationStarknetValidationResultFeature]
type starknetTransactionScanResponseValidationStarknetValidationResultFeatureJSON struct {
	Address     apijson.Field
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseValidationStarknetValidationResultFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseValidationStarknetValidationResultFeatureJSON) RawJSON() string {
	return r.raw
}

// Feature Classification
type StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesType string

const (
	StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesTypeBenign    StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesType = "Benign"
	StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesTypeWarning   StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesType = "Warning"
	StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesTypeMalicious StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesType = "Malicious"
	StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesTypeInfo      StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesType = "Info"
)

func (r StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesTypeBenign, StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesTypeWarning, StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesTypeMalicious, StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesTypeInfo:
		return true
	}
	return false
}

// Verdict of the validation
type StarknetTransactionScanResponseValidationStarknetValidationResultResultType string

const (
	StarknetTransactionScanResponseValidationStarknetValidationResultResultTypeBenign    StarknetTransactionScanResponseValidationStarknetValidationResultResultType = "Benign"
	StarknetTransactionScanResponseValidationStarknetValidationResultResultTypeWarning   StarknetTransactionScanResponseValidationStarknetValidationResultResultType = "Warning"
	StarknetTransactionScanResponseValidationStarknetValidationResultResultTypeMalicious StarknetTransactionScanResponseValidationStarknetValidationResultResultType = "Malicious"
)

func (r StarknetTransactionScanResponseValidationStarknetValidationResultResultType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseValidationStarknetValidationResultResultTypeBenign, StarknetTransactionScanResponseValidationStarknetValidationResultResultTypeWarning, StarknetTransactionScanResponseValidationStarknetValidationResultResultTypeMalicious:
		return true
	}
	return false
}

type StarknetTransactionScanResponseValidationStarknetValidationResultStatus string

const (
	StarknetTransactionScanResponseValidationStarknetValidationResultStatusSuccess StarknetTransactionScanResponseValidationStarknetValidationResultStatus = "Success"
)

func (r StarknetTransactionScanResponseValidationStarknetValidationResultStatus) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseValidationStarknetValidationResultStatusSuccess:
		return true
	}
	return false
}

type StarknetTransactionScanResponseValidationStarknetValidationErrorSchema struct {
	// Error message
	Error  string                                                                       `json:"error,required"`
	Status StarknetTransactionScanResponseValidationStarknetValidationErrorSchemaStatus `json:"status,required"`
	JSON   starknetTransactionScanResponseValidationStarknetValidationErrorSchemaJSON   `json:"-"`
}

// starknetTransactionScanResponseValidationStarknetValidationErrorSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseValidationStarknetValidationErrorSchema]
type starknetTransactionScanResponseValidationStarknetValidationErrorSchemaJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseValidationStarknetValidationErrorSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseValidationStarknetValidationErrorSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseValidationStarknetValidationErrorSchema) implementsStarknetTransactionScanResponseValidation() {
}

type StarknetTransactionScanResponseValidationStarknetValidationErrorSchemaStatus string

const (
	StarknetTransactionScanResponseValidationStarknetValidationErrorSchemaStatusError StarknetTransactionScanResponseValidationStarknetValidationErrorSchemaStatus = "Error"
)

func (r StarknetTransactionScanResponseValidationStarknetValidationErrorSchemaStatus) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseValidationStarknetValidationErrorSchemaStatusError:
		return true
	}
	return false
}

type StarknetTransactionScanResponseValidationStatus string

const (
	StarknetTransactionScanResponseValidationStatusSuccess StarknetTransactionScanResponseValidationStatus = "Success"
	StarknetTransactionScanResponseValidationStatusError   StarknetTransactionScanResponseValidationStatus = "Error"
)

func (r StarknetTransactionScanResponseValidationStatus) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseValidationStatusSuccess, StarknetTransactionScanResponseValidationStatusError:
		return true
	}
	return false
}

// Verdict of the validation
type StarknetTransactionScanResponseValidationResultType string

const (
	StarknetTransactionScanResponseValidationResultTypeBenign    StarknetTransactionScanResponseValidationResultType = "Benign"
	StarknetTransactionScanResponseValidationResultTypeWarning   StarknetTransactionScanResponseValidationResultType = "Warning"
	StarknetTransactionScanResponseValidationResultTypeMalicious StarknetTransactionScanResponseValidationResultType = "Malicious"
)

func (r StarknetTransactionScanResponseValidationResultType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseValidationResultTypeBenign, StarknetTransactionScanResponseValidationResultTypeWarning, StarknetTransactionScanResponseValidationResultTypeMalicious:
		return true
	}
	return false
}

type StarknetTransactionReportParams struct {
	Details param.Field[string]                                     `json:"details,required"`
	Event   param.Field[StarknetTransactionReportParamsEvent]       `json:"event,required"`
	Report  param.Field[StarknetTransactionReportParamsReportUnion] `json:"report,required"`
}

func (r StarknetTransactionReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StarknetTransactionReportParamsEvent string

const (
	StarknetTransactionReportParamsEventShouldBeMalicious StarknetTransactionReportParamsEvent = "should_be_malicious"
	StarknetTransactionReportParamsEventShouldBeBenign    StarknetTransactionReportParamsEvent = "should_be_benign"
)

func (r StarknetTransactionReportParamsEvent) IsKnown() bool {
	switch r {
	case StarknetTransactionReportParamsEventShouldBeMalicious, StarknetTransactionReportParamsEventShouldBeBenign:
		return true
	}
	return false
}

type StarknetTransactionReportParamsReport struct {
	ID     param.Field[string]                                    `json:"id"`
	Params param.Field[interface{}]                               `json:"params"`
	Type   param.Field[StarknetTransactionReportParamsReportType] `json:"type"`
}

func (r StarknetTransactionReportParamsReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionReportParamsReport) implementsStarknetTransactionReportParamsReportUnion() {
}

// Satisfied by [StarknetTransactionReportParamsReportStarknetAppealRequestID],
// [StarknetTransactionReportParamsReportStarknetAppealTransactionDataReport],
// [StarknetTransactionReportParamsReport].
type StarknetTransactionReportParamsReportUnion interface {
	implementsStarknetTransactionReportParamsReportUnion()
}

type StarknetTransactionReportParamsReportStarknetAppealRequestID struct {
	ID   param.Field[string]                                                           `json:"id,required"`
	Type param.Field[StarknetTransactionReportParamsReportStarknetAppealRequestIDType] `json:"type"`
}

func (r StarknetTransactionReportParamsReportStarknetAppealRequestID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionReportParamsReportStarknetAppealRequestID) implementsStarknetTransactionReportParamsReportUnion() {
}

type StarknetTransactionReportParamsReportStarknetAppealRequestIDType string

const (
	StarknetTransactionReportParamsReportStarknetAppealRequestIDTypeRequestID StarknetTransactionReportParamsReportStarknetAppealRequestIDType = "request_id"
)

func (r StarknetTransactionReportParamsReportStarknetAppealRequestIDType) IsKnown() bool {
	switch r {
	case StarknetTransactionReportParamsReportStarknetAppealRequestIDTypeRequestID:
		return true
	}
	return false
}

type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReport struct {
	Params param.Field[StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParams] `json:"params,required"`
	Type   param.Field[StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportType]   `json:"type"`
}

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReport) implementsStarknetTransactionReportParamsReportUnion() {
}

type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParams struct {
	AccountAddress param.Field[string] `json:"account_address,required"`
	// The chain name or chain ID
	Chain param.Field[StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsChain] `json:"chain,required"`
	// Metadata
	Metadata    param.Field[StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataUnion]    `json:"metadata,required"`
	Transaction param.Field[StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionUnion] `json:"transaction,required"`
	// Optional block number or tag context for the simulation
	BlockNumber param.Field[string] `json:"block_number"`
	// List of options to include in the response
	//
	// - `Options.validation`: Include Options.validation output in the response
	//
	// - `Options.simulation`: Include Options.simulation output in the response
	Options param.Field[[]StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsOption] `json:"options"`
}

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A CAIP-2 or a Starknet network name or a Starknet network name
type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsChain string

const (
	StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsChainMainnet            StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsChain = "mainnet"
	StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsChainSepolia            StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsChain = "sepolia"
	StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsChainSepoliaIntegration StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsChain = "sepolia_integration"
)

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsChain) IsKnown() bool {
	switch r {
	case StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsChainMainnet, StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsChainSepolia, StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsChainSepoliaIntegration:
		return true
	}
	return false
}

// Metadata
type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadata struct {
	// Metadata for wallet requests
	Type param.Field[StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataType] `json:"type"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url"`
}

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadata) implementsStarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataUnion() {
}

// Metadata
//
// Satisfied by
// [StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataStarknetWalletRequestMetadata],
// [StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataStarknetInAppRequestMetadata],
// [StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadata].
type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataUnion interface {
	implementsStarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataUnion()
}

type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataStarknetWalletRequestMetadata struct {
	// Metadata for wallet requests
	Type param.Field[StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataStarknetWalletRequestMetadataType] `json:"type,required"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url,required"`
}

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataStarknetWalletRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataStarknetWalletRequestMetadata) implementsStarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataUnion() {
}

// Metadata for wallet requests
type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataStarknetWalletRequestMetadataType string

const (
	StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataStarknetWalletRequestMetadataTypeWallet StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataStarknetWalletRequestMetadataType = "wallet"
)

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataStarknetWalletRequestMetadataType) IsKnown() bool {
	switch r {
	case StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataStarknetWalletRequestMetadataTypeWallet:
		return true
	}
	return false
}

type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataStarknetInAppRequestMetadata struct {
	// Metadata for in-app requests
	Type param.Field[StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataStarknetInAppRequestMetadataType] `json:"type"`
}

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataStarknetInAppRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataStarknetInAppRequestMetadata) implementsStarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataUnion() {
}

// Metadata for in-app requests
type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataStarknetInAppRequestMetadataType string

const (
	StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataStarknetInAppRequestMetadataTypeInApp StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataStarknetInAppRequestMetadataType = "in_app"
)

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataStarknetInAppRequestMetadataType) IsKnown() bool {
	switch r {
	case StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataStarknetInAppRequestMetadataTypeInApp:
		return true
	}
	return false
}

// Metadata for wallet requests
type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataType string

const (
	StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataTypeWallet StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataType = "wallet"
	StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataTypeInApp  StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataType = "in_app"
)

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataType) IsKnown() bool {
	switch r {
	case StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataTypeWallet, StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsMetadataTypeInApp:
		return true
	}
	return false
}

type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransaction struct {
	// The nonce of the transaction.
	Nonce param.Field[string] `json:"nonce,required"`
	// The version of the transaction.
	Version               param.Field[StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionVersion] `json:"version,required"`
	AccountDeploymentData param.Field[interface{}]                                                                                      `json:"account_deployment_data"`
	Calldata              param.Field[interface{}]                                                                                      `json:"calldata"`
	// The id of the chain to which the transaction is sent.
	ChainID param.Field[string] `json:"chain_id"`
	// The hash of the contract class.
	ClassHash           param.Field[string]      `json:"class_hash"`
	ConstructorCalldata param.Field[interface{}] `json:"constructor_calldata"`
	// The salt of the contract address.
	ContractAddressSalt param.Field[string] `json:"contract_address_salt"`
	// The maximum fee that the sender is willing to pay.
	MaxFee param.Field[string] `json:"max_fee"`
	// The nonce data availability mode.
	NonceDataAvailabilityMode param.Field[StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionNonceDataAvailabilityMode] `json:"nonce_data_availability_mode"`
	PaymasterData             param.Field[interface{}]                                                                                                        `json:"paymaster_data"`
	// The address of the sender.
	SenderAddress param.Field[string] `json:"sender_address"`
}

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransaction) implementsStarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionUnion() {
}

// Satisfied by
// [StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV1TransactionSchema],
// [StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV3TransactionSchema],
// [StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetDeployAccountV1TransactionSchema],
// [StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetDeployAccountV3TransactionSchema],
// [StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransaction].
type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionUnion interface {
	implementsStarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionUnion()
}

type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV1TransactionSchema struct {
	// The maximum fee that the sender is willing to pay.
	MaxFee param.Field[string] `json:"max_fee,required"`
	// The nonce of the transaction.
	Nonce param.Field[string] `json:"nonce,required"`
	// The address of the sender.
	SenderAddress param.Field[string] `json:"sender_address,required"`
	// The version of the transaction.
	Version param.Field[StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV1TransactionSchemaVersion] `json:"version,required"`
	// The arguments that are passed to the validate and execute functions.
	Calldata param.Field[[]string] `json:"calldata"`
}

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV1TransactionSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV1TransactionSchema) implementsStarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionUnion() {
}

// The version of the transaction.
type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV1TransactionSchemaVersion int64

const (
	StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV1TransactionSchemaVersion1 StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV1TransactionSchemaVersion = 1
)

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV1TransactionSchemaVersion) IsKnown() bool {
	switch r {
	case StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV1TransactionSchemaVersion1:
		return true
	}
	return false
}

type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV3TransactionSchema struct {
	// The arguments that are passed to the validate and execute functions.
	Calldata param.Field[[]string] `json:"calldata,required"`
	// The id of the chain to which the transaction is sent.
	ChainID param.Field[string] `json:"chain_id,required"`
	// The nonce of the transaction.
	Nonce param.Field[string] `json:"nonce,required"`
	// The address of the sender.
	SenderAddress param.Field[string] `json:"sender_address,required"`
	// The version of the transaction.
	Version param.Field[StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV3TransactionSchemaVersion] `json:"version,required"`
	// For future use. Currently this value is always empty.
	AccountDeploymentData param.Field[[]string] `json:"account_deployment_data"`
	// The nonce data availability mode.
	NonceDataAvailabilityMode param.Field[StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV3TransactionSchemaNonceDataAvailabilityMode] `json:"nonce_data_availability_mode"`
	// For future use. Currently this value is always empty.
	PaymasterData param.Field[[]string] `json:"paymaster_data"`
}

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV3TransactionSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV3TransactionSchema) implementsStarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionUnion() {
}

// The version of the transaction.
type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV3TransactionSchemaVersion int64

const (
	StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV3TransactionSchemaVersion3 StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV3TransactionSchemaVersion = 3
)

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV3TransactionSchemaVersion) IsKnown() bool {
	switch r {
	case StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV3TransactionSchemaVersion3:
		return true
	}
	return false
}

// The nonce data availability mode.
type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV3TransactionSchemaNonceDataAvailabilityMode int64

const (
	StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV3TransactionSchemaNonceDataAvailabilityMode0 StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV3TransactionSchemaNonceDataAvailabilityMode = 0
)

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV3TransactionSchemaNonceDataAvailabilityMode) IsKnown() bool {
	switch r {
	case StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetInvokeV3TransactionSchemaNonceDataAvailabilityMode0:
		return true
	}
	return false
}

type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetDeployAccountV1TransactionSchema struct {
	// The hash of the contract class.
	ClassHash param.Field[string] `json:"class_hash,required"`
	// The arguments that are passed to the constructor function.
	ConstructorCalldata param.Field[[]string] `json:"constructor_calldata,required"`
	// The salt of the contract address.
	ContractAddressSalt param.Field[string] `json:"contract_address_salt,required"`
	// The maximum fee that the sender is willing to pay.
	MaxFee param.Field[string] `json:"max_fee,required"`
	// The nonce of the transaction.
	Nonce param.Field[string] `json:"nonce,required"`
	// The version of the transaction.
	Version param.Field[StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetDeployAccountV1TransactionSchemaVersion] `json:"version,required"`
}

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetDeployAccountV1TransactionSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetDeployAccountV1TransactionSchema) implementsStarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionUnion() {
}

// The version of the transaction.
type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetDeployAccountV1TransactionSchemaVersion int64

const (
	StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetDeployAccountV1TransactionSchemaVersion1 StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetDeployAccountV1TransactionSchemaVersion = 1
)

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetDeployAccountV1TransactionSchemaVersion) IsKnown() bool {
	switch r {
	case StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetDeployAccountV1TransactionSchemaVersion1:
		return true
	}
	return false
}

type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetDeployAccountV3TransactionSchema struct {
	// The hash of the contract class.
	ClassHash param.Field[string] `json:"class_hash,required"`
	// The arguments that are passed to the constructor function.
	ConstructorCalldata param.Field[[]string] `json:"constructor_calldata,required"`
	// The salt of the contract address.
	ContractAddressSalt param.Field[string] `json:"contract_address_salt,required"`
	// The maximum fee that the sender is willing to pay.
	MaxFee param.Field[string] `json:"max_fee,required"`
	// The nonce of the transaction.
	Nonce param.Field[string] `json:"nonce,required"`
	// The version of the transaction.
	Version param.Field[StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetDeployAccountV3TransactionSchemaVersion] `json:"version,required"`
}

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetDeployAccountV3TransactionSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetDeployAccountV3TransactionSchema) implementsStarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionUnion() {
}

// The version of the transaction.
type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetDeployAccountV3TransactionSchemaVersion int64

const (
	StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetDeployAccountV3TransactionSchemaVersion3 StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetDeployAccountV3TransactionSchemaVersion = 3
)

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetDeployAccountV3TransactionSchemaVersion) IsKnown() bool {
	switch r {
	case StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionStarknetDeployAccountV3TransactionSchemaVersion3:
		return true
	}
	return false
}

// The version of the transaction.
type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionVersion int64

const (
	StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionVersion1 StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionVersion = 1
	StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionVersion3 StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionVersion = 3
)

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionVersion) IsKnown() bool {
	switch r {
	case StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionVersion1, StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionVersion3:
		return true
	}
	return false
}

// The nonce data availability mode.
type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionNonceDataAvailabilityMode int64

const (
	StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionNonceDataAvailabilityMode0 StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionNonceDataAvailabilityMode = 0
)

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionNonceDataAvailabilityMode) IsKnown() bool {
	switch r {
	case StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionNonceDataAvailabilityMode0:
		return true
	}
	return false
}

type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsOption string

const (
	StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsOptionValidation StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsOption = "validation"
	StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsOptionSimulation StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsOption = "simulation"
)

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsOption) IsKnown() bool {
	switch r {
	case StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsOptionValidation, StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsOptionSimulation:
		return true
	}
	return false
}

type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportType string

const (
	StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportTypeParams StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportType = "params"
)

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportType) IsKnown() bool {
	switch r {
	case StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportTypeParams:
		return true
	}
	return false
}

type StarknetTransactionReportParamsReportType string

const (
	StarknetTransactionReportParamsReportTypeRequestID StarknetTransactionReportParamsReportType = "request_id"
	StarknetTransactionReportParamsReportTypeParams    StarknetTransactionReportParamsReportType = "params"
)

func (r StarknetTransactionReportParamsReportType) IsKnown() bool {
	switch r {
	case StarknetTransactionReportParamsReportTypeRequestID, StarknetTransactionReportParamsReportTypeParams:
		return true
	}
	return false
}

type StarknetTransactionScanParams struct {
	AccountAddress param.Field[string] `json:"account_address,required"`
	// The chain name or chain ID
	Chain param.Field[StarknetTransactionScanParamsChain] `json:"chain,required"`
	// Metadata
	Metadata    param.Field[StarknetTransactionScanParamsMetadataUnion]    `json:"metadata,required"`
	Transaction param.Field[StarknetTransactionScanParamsTransactionUnion] `json:"transaction,required"`
	// Optional block number or tag context for the simulation
	BlockNumber param.Field[string] `json:"block_number"`
	// List of options to include in the response
	//
	// - `Options.validation`: Include Options.validation output in the response
	//
	// - `Options.simulation`: Include Options.simulation output in the response
	Options param.Field[[]StarknetTransactionScanParamsOption] `json:"options"`
}

func (r StarknetTransactionScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A CAIP-2 or a Starknet network name or a Starknet network name
type StarknetTransactionScanParamsChain string

const (
	StarknetTransactionScanParamsChainMainnet            StarknetTransactionScanParamsChain = "mainnet"
	StarknetTransactionScanParamsChainSepolia            StarknetTransactionScanParamsChain = "sepolia"
	StarknetTransactionScanParamsChainSepoliaIntegration StarknetTransactionScanParamsChain = "sepolia_integration"
)

func (r StarknetTransactionScanParamsChain) IsKnown() bool {
	switch r {
	case StarknetTransactionScanParamsChainMainnet, StarknetTransactionScanParamsChainSepolia, StarknetTransactionScanParamsChainSepoliaIntegration:
		return true
	}
	return false
}

// Metadata
type StarknetTransactionScanParamsMetadata struct {
	// Metadata for wallet requests
	Type param.Field[StarknetTransactionScanParamsMetadataType] `json:"type"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url"`
}

func (r StarknetTransactionScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionScanParamsMetadata) implementsStarknetTransactionScanParamsMetadataUnion() {
}

// Metadata
//
// Satisfied by
// [StarknetTransactionScanParamsMetadataStarknetWalletRequestMetadata],
// [StarknetTransactionScanParamsMetadataStarknetInAppRequestMetadata],
// [StarknetTransactionScanParamsMetadata].
type StarknetTransactionScanParamsMetadataUnion interface {
	implementsStarknetTransactionScanParamsMetadataUnion()
}

type StarknetTransactionScanParamsMetadataStarknetWalletRequestMetadata struct {
	// Metadata for wallet requests
	Type param.Field[StarknetTransactionScanParamsMetadataStarknetWalletRequestMetadataType] `json:"type,required"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url,required"`
}

func (r StarknetTransactionScanParamsMetadataStarknetWalletRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionScanParamsMetadataStarknetWalletRequestMetadata) implementsStarknetTransactionScanParamsMetadataUnion() {
}

// Metadata for wallet requests
type StarknetTransactionScanParamsMetadataStarknetWalletRequestMetadataType string

const (
	StarknetTransactionScanParamsMetadataStarknetWalletRequestMetadataTypeWallet StarknetTransactionScanParamsMetadataStarknetWalletRequestMetadataType = "wallet"
)

func (r StarknetTransactionScanParamsMetadataStarknetWalletRequestMetadataType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanParamsMetadataStarknetWalletRequestMetadataTypeWallet:
		return true
	}
	return false
}

type StarknetTransactionScanParamsMetadataStarknetInAppRequestMetadata struct {
	// Metadata for in-app requests
	Type param.Field[StarknetTransactionScanParamsMetadataStarknetInAppRequestMetadataType] `json:"type"`
}

func (r StarknetTransactionScanParamsMetadataStarknetInAppRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionScanParamsMetadataStarknetInAppRequestMetadata) implementsStarknetTransactionScanParamsMetadataUnion() {
}

// Metadata for in-app requests
type StarknetTransactionScanParamsMetadataStarknetInAppRequestMetadataType string

const (
	StarknetTransactionScanParamsMetadataStarknetInAppRequestMetadataTypeInApp StarknetTransactionScanParamsMetadataStarknetInAppRequestMetadataType = "in_app"
)

func (r StarknetTransactionScanParamsMetadataStarknetInAppRequestMetadataType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanParamsMetadataStarknetInAppRequestMetadataTypeInApp:
		return true
	}
	return false
}

// Metadata for wallet requests
type StarknetTransactionScanParamsMetadataType string

const (
	StarknetTransactionScanParamsMetadataTypeWallet StarknetTransactionScanParamsMetadataType = "wallet"
	StarknetTransactionScanParamsMetadataTypeInApp  StarknetTransactionScanParamsMetadataType = "in_app"
)

func (r StarknetTransactionScanParamsMetadataType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanParamsMetadataTypeWallet, StarknetTransactionScanParamsMetadataTypeInApp:
		return true
	}
	return false
}

type StarknetTransactionScanParamsTransaction struct {
	// The nonce of the transaction.
	Nonce param.Field[string] `json:"nonce,required"`
	// The version of the transaction.
	Version               param.Field[StarknetTransactionScanParamsTransactionVersion] `json:"version,required"`
	AccountDeploymentData param.Field[interface{}]                                     `json:"account_deployment_data"`
	Calldata              param.Field[interface{}]                                     `json:"calldata"`
	// The id of the chain to which the transaction is sent.
	ChainID param.Field[string] `json:"chain_id"`
	// The hash of the contract class.
	ClassHash           param.Field[string]      `json:"class_hash"`
	ConstructorCalldata param.Field[interface{}] `json:"constructor_calldata"`
	// The salt of the contract address.
	ContractAddressSalt param.Field[string] `json:"contract_address_salt"`
	// The maximum fee that the sender is willing to pay.
	MaxFee param.Field[string] `json:"max_fee"`
	// The nonce data availability mode.
	NonceDataAvailabilityMode param.Field[StarknetTransactionScanParamsTransactionNonceDataAvailabilityMode] `json:"nonce_data_availability_mode"`
	PaymasterData             param.Field[interface{}]                                                       `json:"paymaster_data"`
	// The address of the sender.
	SenderAddress param.Field[string] `json:"sender_address"`
}

func (r StarknetTransactionScanParamsTransaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionScanParamsTransaction) implementsStarknetTransactionScanParamsTransactionUnion() {
}

// Satisfied by
// [StarknetTransactionScanParamsTransactionStarknetInvokeV1TransactionSchema],
// [StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchema],
// [StarknetTransactionScanParamsTransactionStarknetDeployAccountV1TransactionSchema],
// [StarknetTransactionScanParamsTransactionStarknetDeployAccountV3TransactionSchema],
// [StarknetTransactionScanParamsTransaction].
type StarknetTransactionScanParamsTransactionUnion interface {
	implementsStarknetTransactionScanParamsTransactionUnion()
}

type StarknetTransactionScanParamsTransactionStarknetInvokeV1TransactionSchema struct {
	// The maximum fee that the sender is willing to pay.
	MaxFee param.Field[string] `json:"max_fee,required"`
	// The nonce of the transaction.
	Nonce param.Field[string] `json:"nonce,required"`
	// The address of the sender.
	SenderAddress param.Field[string] `json:"sender_address,required"`
	// The version of the transaction.
	Version param.Field[StarknetTransactionScanParamsTransactionStarknetInvokeV1TransactionSchemaVersion] `json:"version,required"`
	// The arguments that are passed to the validate and execute functions.
	Calldata param.Field[[]string] `json:"calldata"`
}

func (r StarknetTransactionScanParamsTransactionStarknetInvokeV1TransactionSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionScanParamsTransactionStarknetInvokeV1TransactionSchema) implementsStarknetTransactionScanParamsTransactionUnion() {
}

// The version of the transaction.
type StarknetTransactionScanParamsTransactionStarknetInvokeV1TransactionSchemaVersion int64

const (
	StarknetTransactionScanParamsTransactionStarknetInvokeV1TransactionSchemaVersion1 StarknetTransactionScanParamsTransactionStarknetInvokeV1TransactionSchemaVersion = 1
)

func (r StarknetTransactionScanParamsTransactionStarknetInvokeV1TransactionSchemaVersion) IsKnown() bool {
	switch r {
	case StarknetTransactionScanParamsTransactionStarknetInvokeV1TransactionSchemaVersion1:
		return true
	}
	return false
}

type StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchema struct {
	// The arguments that are passed to the validate and execute functions.
	Calldata param.Field[[]string] `json:"calldata,required"`
	// The id of the chain to which the transaction is sent.
	ChainID param.Field[string] `json:"chain_id,required"`
	// The nonce of the transaction.
	Nonce param.Field[string] `json:"nonce,required"`
	// The address of the sender.
	SenderAddress param.Field[string] `json:"sender_address,required"`
	// The version of the transaction.
	Version param.Field[StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaVersion] `json:"version,required"`
	// For future use. Currently this value is always empty.
	AccountDeploymentData param.Field[[]string] `json:"account_deployment_data"`
	// The nonce data availability mode.
	NonceDataAvailabilityMode param.Field[StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaNonceDataAvailabilityMode] `json:"nonce_data_availability_mode"`
	// For future use. Currently this value is always empty.
	PaymasterData param.Field[[]string] `json:"paymaster_data"`
}

func (r StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchema) implementsStarknetTransactionScanParamsTransactionUnion() {
}

// The version of the transaction.
type StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaVersion int64

const (
	StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaVersion3 StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaVersion = 3
)

func (r StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaVersion) IsKnown() bool {
	switch r {
	case StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaVersion3:
		return true
	}
	return false
}

// The nonce data availability mode.
type StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaNonceDataAvailabilityMode int64

const (
	StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaNonceDataAvailabilityMode0 StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaNonceDataAvailabilityMode = 0
)

func (r StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaNonceDataAvailabilityMode) IsKnown() bool {
	switch r {
	case StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaNonceDataAvailabilityMode0:
		return true
	}
	return false
}

type StarknetTransactionScanParamsTransactionStarknetDeployAccountV1TransactionSchema struct {
	// The hash of the contract class.
	ClassHash param.Field[string] `json:"class_hash,required"`
	// The arguments that are passed to the constructor function.
	ConstructorCalldata param.Field[[]string] `json:"constructor_calldata,required"`
	// The salt of the contract address.
	ContractAddressSalt param.Field[string] `json:"contract_address_salt,required"`
	// The maximum fee that the sender is willing to pay.
	MaxFee param.Field[string] `json:"max_fee,required"`
	// The nonce of the transaction.
	Nonce param.Field[string] `json:"nonce,required"`
	// The version of the transaction.
	Version param.Field[StarknetTransactionScanParamsTransactionStarknetDeployAccountV1TransactionSchemaVersion] `json:"version,required"`
}

func (r StarknetTransactionScanParamsTransactionStarknetDeployAccountV1TransactionSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionScanParamsTransactionStarknetDeployAccountV1TransactionSchema) implementsStarknetTransactionScanParamsTransactionUnion() {
}

// The version of the transaction.
type StarknetTransactionScanParamsTransactionStarknetDeployAccountV1TransactionSchemaVersion int64

const (
	StarknetTransactionScanParamsTransactionStarknetDeployAccountV1TransactionSchemaVersion1 StarknetTransactionScanParamsTransactionStarknetDeployAccountV1TransactionSchemaVersion = 1
)

func (r StarknetTransactionScanParamsTransactionStarknetDeployAccountV1TransactionSchemaVersion) IsKnown() bool {
	switch r {
	case StarknetTransactionScanParamsTransactionStarknetDeployAccountV1TransactionSchemaVersion1:
		return true
	}
	return false
}

type StarknetTransactionScanParamsTransactionStarknetDeployAccountV3TransactionSchema struct {
	// The hash of the contract class.
	ClassHash param.Field[string] `json:"class_hash,required"`
	// The arguments that are passed to the constructor function.
	ConstructorCalldata param.Field[[]string] `json:"constructor_calldata,required"`
	// The salt of the contract address.
	ContractAddressSalt param.Field[string] `json:"contract_address_salt,required"`
	// The maximum fee that the sender is willing to pay.
	MaxFee param.Field[string] `json:"max_fee,required"`
	// The nonce of the transaction.
	Nonce param.Field[string] `json:"nonce,required"`
	// The version of the transaction.
	Version param.Field[StarknetTransactionScanParamsTransactionStarknetDeployAccountV3TransactionSchemaVersion] `json:"version,required"`
}

func (r StarknetTransactionScanParamsTransactionStarknetDeployAccountV3TransactionSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionScanParamsTransactionStarknetDeployAccountV3TransactionSchema) implementsStarknetTransactionScanParamsTransactionUnion() {
}

// The version of the transaction.
type StarknetTransactionScanParamsTransactionStarknetDeployAccountV3TransactionSchemaVersion int64

const (
	StarknetTransactionScanParamsTransactionStarknetDeployAccountV3TransactionSchemaVersion3 StarknetTransactionScanParamsTransactionStarknetDeployAccountV3TransactionSchemaVersion = 3
)

func (r StarknetTransactionScanParamsTransactionStarknetDeployAccountV3TransactionSchemaVersion) IsKnown() bool {
	switch r {
	case StarknetTransactionScanParamsTransactionStarknetDeployAccountV3TransactionSchemaVersion3:
		return true
	}
	return false
}

// The version of the transaction.
type StarknetTransactionScanParamsTransactionVersion int64

const (
	StarknetTransactionScanParamsTransactionVersion1 StarknetTransactionScanParamsTransactionVersion = 1
	StarknetTransactionScanParamsTransactionVersion3 StarknetTransactionScanParamsTransactionVersion = 3
)

func (r StarknetTransactionScanParamsTransactionVersion) IsKnown() bool {
	switch r {
	case StarknetTransactionScanParamsTransactionVersion1, StarknetTransactionScanParamsTransactionVersion3:
		return true
	}
	return false
}

// The nonce data availability mode.
type StarknetTransactionScanParamsTransactionNonceDataAvailabilityMode int64

const (
	StarknetTransactionScanParamsTransactionNonceDataAvailabilityMode0 StarknetTransactionScanParamsTransactionNonceDataAvailabilityMode = 0
)

func (r StarknetTransactionScanParamsTransactionNonceDataAvailabilityMode) IsKnown() bool {
	switch r {
	case StarknetTransactionScanParamsTransactionNonceDataAvailabilityMode0:
		return true
	}
	return false
}

type StarknetTransactionScanParamsOption string

const (
	StarknetTransactionScanParamsOptionValidation StarknetTransactionScanParamsOption = "validation"
	StarknetTransactionScanParamsOptionSimulation StarknetTransactionScanParamsOption = "simulation"
)

func (r StarknetTransactionScanParamsOption) IsKnown() bool {
	switch r {
	case StarknetTransactionScanParamsOptionValidation, StarknetTransactionScanParamsOptionSimulation:
		return true
	}
	return false
}
