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
	// [map[string][]StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiff].
	AssetsDiffs interface{} `json:"assets_diffs,required"`
	// This field can have the runtime type of
	// [map[string][]StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposure].
	Exposures interface{} `json:"exposures,required"`
	// This field can have the runtime type of
	// [[]StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAddressDetail].
	AddressDetails interface{} `json:"address_details,required"`
	// This field can have the runtime type of
	// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummary].
	AccountSummary interface{} `json:"account_summary,required"`
	// Error message
	Error string                                        `json:"error"`
	JSON  starknetTransactionScanResponseSimulationJSON `json:"-"`
	union StarknetTransactionScanResponseSimulationUnion
}

// starknetTransactionScanResponseSimulationJSON contains the JSON metadata for the
// struct [StarknetTransactionScanResponseSimulation]
type starknetTransactionScanResponseSimulationJSON struct {
	Status         apijson.Field
	AssetsDiffs    apijson.Field
	Exposures      apijson.Field
	AddressDetails apijson.Field
	AccountSummary apijson.Field
	Error          apijson.Field
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
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchema],
// [StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchema].
func (r StarknetTransactionScanResponseSimulation) AsUnion() StarknetTransactionScanResponseSimulationUnion {
	return r.union
}

// Simulation result; Only present if simulation option is included in the request
//
// Union satisfied by
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchema] or
// [StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchema].
type StarknetTransactionScanResponseSimulationUnion interface {
	implementsStarknetTransactionScanResponseSimulation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StarknetTransactionScanResponseSimulationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetSimulationResultSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchema{}),
		},
	)
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchema struct {
	// Summary of the actions and asset transfers that were made by the requested
	// account address
	AccountSummary StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummary `json:"account_summary,required"`
	Status         StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaStatus         `json:"status,required"`
	// Details of addresses involved in the transaction
	AddressDetails []StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAddressDetail `json:"address_details"`
	// Mapping between the address of an account to the assets diff during the
	// transaction
	AssetsDiffs map[string][]StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiff `json:"assets_diffs"`
	// Mapping between the address of an account to the exposure of the assets during
	// the transaction
	Exposures map[string][]StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposure `json:"exposures"`
	JSON      starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaJSON                  `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchema]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaJSON struct {
	AccountSummary apijson.Field
	Status         apijson.Field
	AddressDetails apijson.Field
	AssetsDiffs    apijson.Field
	Exposures      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchema) implementsStarknetTransactionScanResponseSimulation() {
}

// Summary of the actions and asset transfers that were made by the requested
// account address
type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummary struct {
	// Exposures made by the requested account address
	Exposures []StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposure `json:"exposures,required"`
	// Total USD diff for the requested account address
	TotalUsdDiff StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryTotalUsdDiff `json:"total_usd_diff,required"`
	// Assets diffs of the requested account address
	AssetsDiffs []StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiff `json:"assets_diffs"`
	// Total USD exposure for each of the spender addresses during the transaction
	TotalUsdExposure map[string]float64                                                                        `json:"total_usd_exposure"`
	JSON             starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummary]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryJSON struct {
	Exposures        apijson.Field
	TotalUsdDiff     apijson.Field
	AssetsDiffs      apijson.Field
	TotalUsdExposure apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposure struct {
	// This field can have the runtime type of [StarknetErc20Details],
	// [StarknetErc721Details], [StarknetErc1155Details].
	Asset interface{} `json:"asset"`
	// This field can have the runtime type of
	// [map[string]StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaSpender],
	// [map[string]StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaSpender],
	// [map[string]StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaSpender].
	Spenders interface{}                                                                                       `json:"spenders,required"`
	JSON     starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposureJSON `json:"-"`
	union    StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresUnion
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposure]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposureJSON) RawJSON() string {
	return r.raw
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposure) UnmarshalJSON(data []byte) (err error) {
	*r = StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchema],
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchema],
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchema].
func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposure) AsUnion() StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresUnion {
	return r.union
}

// Union satisfied by
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchema],
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchema]
// or
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchema].
type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresUnion interface {
	implementsStarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchema{}),
		},
	)
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchema struct {
	Asset StarknetErc20Details `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaSpender `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaJSON               `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchema]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchema) implementsStarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposure() {
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaSpender struct {
	// Approval value of the ERC20 token
	Approval int64               `json:"approval,required"`
	Exposure []StarknetErc20Diff `json:"exposure,required"`
	// Expiration date of the approval
	Expiration time.Time `json:"expiration,nullable" format:"date-time"`
	// Summarized description of the exposure
	Summary string                                                                                                                                                                           `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaSpenderJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaSpenderJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaSpender]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaSpenderJSON struct {
	Approval    apijson.Field
	Exposure    apijson.Field
	Expiration  apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaSpenderJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchema struct {
	Asset StarknetErc721Details `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaSpender `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaJSON               `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchema]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchema) implementsStarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposure() {
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaSpender struct {
	Exposure []StarknetErc721Diff `json:"exposure,required"`
	// Whether `setApprovalForAll` was invoked
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// Summarized description of the exposure
	Summary string                                                                                                                                                                             `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaSpenderJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaSpenderJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaSpender]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaSpenderJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchema struct {
	Asset StarknetErc1155Details `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaSpender `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaJSON               `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchema]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchema) implementsStarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposure() {
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaSpender struct {
	Exposure []StarknetErc1155Diff `json:"exposure,required"`
	// Whether `setApprovalForAll` was invoked
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// Summarized description of the exposure
	Summary string                                                                                                                                                                               `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaSpenderJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaSpenderJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaSpender]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaSpenderJSON) RawJSON() string {
	return r.raw
}

// Total USD diff for the requested account address
type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryTotalUsdDiff struct {
	// Total incoming USD transfers
	In float64 `json:"in,required"`
	// Total outgoing USD transfers
	Out float64 `json:"out,required"`
	// Total USD transfers
	Total float64                                                                                               `json:"total"`
	JSON  starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryTotalUsdDiffJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryTotalUsdDiffJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryTotalUsdDiff]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiff struct {
	// This field can have the runtime type of
	// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAsset],
	// [StarknetErc20Details], [StarknetErc721Details], [StarknetErc1155Details].
	Asset interface{} `json:"asset"`
	// This field can have the runtime type of [StarknetNativeDiff],
	// [StarknetErc20Diff], [StarknetErc721Diff], [StarknetErc1155Diff].
	In interface{} `json:"in,required"`
	// This field can have the runtime type of [StarknetNativeDiff],
	// [StarknetErc20Diff], [StarknetErc721Diff], [StarknetErc1155Diff].
	Out   interface{}                                                                                         `json:"out,required"`
	JSON  starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffJSON `json:"-"`
	union StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsUnion
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiff]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchema],
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaErc20DiffSchema],
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaErc721DiffSchema],
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaErc1155DiffSchema].
func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiff) AsUnion() StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchema],
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaErc20DiffSchema],
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaErc721DiffSchema]
// or
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaErc1155DiffSchema].
type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsUnion interface {
	implementsStarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaErc20DiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaErc721DiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaErc1155DiffSchema{}),
		},
	)
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchema struct {
	Asset StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAsset `json:"asset,required"`
	// Details of the incoming transfer
	In StarknetNativeDiff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetNativeDiff                                                                                                                                                                   `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchema]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchema) implementsStarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAsset struct {
	// Decimals of the asset
	Decimals StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetDecimals `json:"decimals"`
	// Name of the asset
	Name StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetName `json:"name"`
	// Symbol of the asset
	Symbol StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetSymbol `json:"symbol"`
	// Type of the asset (`NATIVE`)
	Type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAsset]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetJSON struct {
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Decimals of the asset
type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetDecimals int64

const (
	StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetDecimals18 StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetDecimals = 18
)

func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetDecimals) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetDecimals18:
		return true
	}
	return false
}

// Name of the asset
type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetName string

const (
	StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetNameStrk StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetName = "STRK"
)

func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetName) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetNameStrk:
		return true
	}
	return false
}

// Symbol of the asset
type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetSymbol string

const (
	StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetSymbolStrk StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetSymbol = "STRK"
)

func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetSymbol) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetSymbolStrk:
		return true
	}
	return false
}

// Type of the asset (`NATIVE`)
type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetTypeNative StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetType = "NATIVE"
)

func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetTypeNative:
		return true
	}
	return false
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaErc20DiffSchema struct {
	Asset StarknetErc20Details `json:"asset,required"`
	// Details of the incoming transfer
	In StarknetErc20Diff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetErc20Diff                                                                                                                                                             `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaErc20DiffSchemaJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaErc20DiffSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaErc20DiffSchema]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaErc20DiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaErc20DiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaErc20DiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaErc20DiffSchema) implementsStarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaErc721DiffSchema struct {
	Asset StarknetErc721Details `json:"asset,required"`
	// Details of the incoming transfer
	In StarknetErc721Diff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetErc721Diff                                                                                                                                                              `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaErc721DiffSchemaJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaErc721DiffSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaErc721DiffSchema]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaErc721DiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaErc721DiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaErc721DiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaErc721DiffSchema) implementsStarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaErc1155DiffSchema struct {
	Asset StarknetErc1155Details `json:"asset,required"`
	// Details of the incoming transfer
	In StarknetErc1155Diff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetErc1155Diff                                                                                                                                                               `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaErc1155DiffSchemaJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaErc1155DiffSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaErc1155DiffSchema]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaErc1155DiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaErc1155DiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaErc1155DiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaErc1155DiffSchema) implementsStarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAccountSummaryAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaStatus string

const (
	StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaStatusSuccess StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaStatus = "Success"
)

func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaStatus) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaStatusSuccess:
		return true
	}
	return false
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAddressDetail struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// Description of the account
	Description string                                                                                   `json:"description,nullable"`
	JSON        starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAddressDetailJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAddressDetailJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAddressDetail]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAddressDetailJSON struct {
	AccountAddress apijson.Field
	Description    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAddressDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAddressDetailJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiff struct {
	// This field can have the runtime type of
	// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAsset],
	// [StarknetErc20Details], [StarknetErc721Details], [StarknetErc1155Details].
	Asset interface{} `json:"asset"`
	// This field can have the runtime type of [StarknetNativeDiff],
	// [StarknetErc20Diff], [StarknetErc721Diff], [StarknetErc1155Diff].
	In interface{} `json:"in,required"`
	// This field can have the runtime type of [StarknetNativeDiff],
	// [StarknetErc20Diff], [StarknetErc721Diff], [StarknetErc1155Diff].
	Out   interface{}                                                                           `json:"out,required"`
	JSON  starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffJSON `json:"-"`
	union StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsUnion
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiff]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchema],
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaErc20DiffSchema],
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaErc721DiffSchema],
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaErc1155DiffSchema].
func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiff) AsUnion() StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchema],
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaErc20DiffSchema],
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaErc721DiffSchema]
// or
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaErc1155DiffSchema].
type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsUnion interface {
	implementsStarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaErc20DiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaErc721DiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaErc1155DiffSchema{}),
		},
	)
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchema struct {
	Asset StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAsset `json:"asset,required"`
	// Details of the incoming transfer
	In StarknetNativeDiff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetNativeDiff                                                                                                                                                     `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchema]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchema) implementsStarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAsset struct {
	// Decimals of the asset
	Decimals StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetDecimals `json:"decimals"`
	// Name of the asset
	Name StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetName `json:"name"`
	// Symbol of the asset
	Symbol StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetSymbol `json:"symbol"`
	// Type of the asset (`NATIVE`)
	Type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAsset]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetJSON struct {
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Decimals of the asset
type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetDecimals int64

const (
	StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetDecimals18 StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetDecimals = 18
)

func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetDecimals) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetDecimals18:
		return true
	}
	return false
}

// Name of the asset
type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetName string

const (
	StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetNameStrk StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetName = "STRK"
)

func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetName) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetNameStrk:
		return true
	}
	return false
}

// Symbol of the asset
type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetSymbol string

const (
	StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetSymbolStrk StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetSymbol = "STRK"
)

func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetSymbol) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetSymbolStrk:
		return true
	}
	return false
}

// Type of the asset (`NATIVE`)
type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetTypeNative StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetType = "NATIVE"
)

func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetTypeNative:
		return true
	}
	return false
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaErc20DiffSchema struct {
	Asset StarknetErc20Details `json:"asset,required"`
	// Details of the incoming transfer
	In StarknetErc20Diff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetErc20Diff                                                                                                                                               `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaErc20DiffSchemaJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaErc20DiffSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaErc20DiffSchema]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaErc20DiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaErc20DiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaErc20DiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaErc20DiffSchema) implementsStarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaErc721DiffSchema struct {
	Asset StarknetErc721Details `json:"asset,required"`
	// Details of the incoming transfer
	In StarknetErc721Diff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetErc721Diff                                                                                                                                                `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaErc721DiffSchemaJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaErc721DiffSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaErc721DiffSchema]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaErc721DiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaErc721DiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaErc721DiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaErc721DiffSchema) implementsStarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaErc1155DiffSchema struct {
	Asset StarknetErc1155Details `json:"asset,required"`
	// Details of the incoming transfer
	In StarknetErc1155Diff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetErc1155Diff                                                                                                                                                 `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaErc1155DiffSchemaJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaErc1155DiffSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaErc1155DiffSchema]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaErc1155DiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaErc1155DiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaErc1155DiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaErc1155DiffSchema) implementsStarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposure struct {
	// This field can have the runtime type of [StarknetErc20Details],
	// [StarknetErc721Details], [StarknetErc1155Details].
	Asset interface{} `json:"asset"`
	// This field can have the runtime type of
	// [map[string]StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaSpender],
	// [map[string]StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaSpender],
	// [map[string]StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaSpender].
	Spenders interface{}                                                                         `json:"spenders,required"`
	JSON     starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposureJSON `json:"-"`
	union    StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresUnion
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposure]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposureJSON) RawJSON() string {
	return r.raw
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposure) UnmarshalJSON(data []byte) (err error) {
	*r = StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchema],
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchema],
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchema].
func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposure) AsUnion() StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresUnion {
	return r.union
}

// Union satisfied by
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchema],
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchema]
// or
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchema].
type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresUnion interface {
	implementsStarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchema{}),
		},
	)
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchema struct {
	Asset StarknetErc20Details `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaSpender `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaJSON               `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchema]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchema) implementsStarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposure() {
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaSpender struct {
	// Approval value of the ERC20 token
	Approval int64               `json:"approval,required"`
	Exposure []StarknetErc20Diff `json:"exposure,required"`
	// Expiration date of the approval
	Expiration time.Time `json:"expiration,nullable" format:"date-time"`
	// Summarized description of the exposure
	Summary string                                                                                                                                                             `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaSpenderJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaSpenderJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaSpender]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaSpenderJSON struct {
	Approval    apijson.Field
	Exposure    apijson.Field
	Expiration  apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc20DetailsSchemaErc20ExposureSchemaSpenderJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchema struct {
	Asset StarknetErc721Details `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaSpender `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaJSON               `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchema]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchema) implementsStarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposure() {
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaSpender struct {
	Exposure []StarknetErc721Diff `json:"exposure,required"`
	// Whether `setApprovalForAll` was invoked
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// Summarized description of the exposure
	Summary string                                                                                                                                                               `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaSpenderJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaSpenderJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaSpender]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc721DetailsSchemaErc721ExposureSchemaSpenderJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchema struct {
	Asset StarknetErc1155Details `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaSpender `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaJSON               `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchema]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchema) implementsStarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposure() {
}

type StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaSpender struct {
	Exposure []StarknetErc1155Diff `json:"exposure,required"`
	// Whether `setApprovalForAll` was invoked
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// Summarized description of the exposure
	Summary string                                                                                                                                                                 `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaSpenderJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaSpenderJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaSpender]
type starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaErc1155DetailsSchemaErc1155ExposureSchemaSpenderJSON) RawJSON() string {
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
	// Verdict of the validation
	ResultType StarknetTransactionScanResponseValidationResultType `json:"result_type"`
	// A textual description about the validation result
	Description string `json:"description"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason string `json:"reason"`
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification"`
	// This field can have the runtime type of
	// [[]StarknetTransactionScanResponseValidationStarknetValidationResultSchemaFeature].
	Features interface{} `json:"features,required"`
	// Error message
	Error string                                        `json:"error"`
	JSON  starknetTransactionScanResponseValidationJSON `json:"-"`
	union StarknetTransactionScanResponseValidationUnion
}

// starknetTransactionScanResponseValidationJSON contains the JSON metadata for the
// struct [StarknetTransactionScanResponseValidation]
type starknetTransactionScanResponseValidationJSON struct {
	Status         apijson.Field
	ResultType     apijson.Field
	Description    apijson.Field
	Reason         apijson.Field
	Classification apijson.Field
	Features       apijson.Field
	Error          apijson.Field
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
// [StarknetTransactionScanResponseValidationStarknetValidationResultSchema],
// [StarknetTransactionScanResponseValidationStarknetValidationErrorSchema].
func (r StarknetTransactionScanResponseValidation) AsUnion() StarknetTransactionScanResponseValidationUnion {
	return r.union
}

// Validation result; Only present if validation option is included in the request
//
// Union satisfied by
// [StarknetTransactionScanResponseValidationStarknetValidationResultSchema] or
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
			Type:       reflect.TypeOf(StarknetTransactionScanResponseValidationStarknetValidationResultSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseValidationStarknetValidationErrorSchema{}),
		},
	)
}

type StarknetTransactionScanResponseValidationStarknetValidationResultSchema struct {
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification,required"`
	// A textual description about the validation result
	Description string `json:"description,required"`
	// A list of features about this transaction explaining the validation
	Features []StarknetTransactionScanResponseValidationStarknetValidationResultSchemaFeature `json:"features,required"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason string `json:"reason,required"`
	// Verdict of the validation
	ResultType StarknetTransactionScanResponseValidationStarknetValidationResultSchemaResultType `json:"result_type,required"`
	Status     StarknetTransactionScanResponseValidationStarknetValidationResultSchemaStatus     `json:"status,required"`
	JSON       starknetTransactionScanResponseValidationStarknetValidationResultSchemaJSON       `json:"-"`
}

// starknetTransactionScanResponseValidationStarknetValidationResultSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseValidationStarknetValidationResultSchema]
type starknetTransactionScanResponseValidationStarknetValidationResultSchemaJSON struct {
	Classification apijson.Field
	Description    apijson.Field
	Features       apijson.Field
	Reason         apijson.Field
	ResultType     apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseValidationStarknetValidationResultSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseValidationStarknetValidationResultSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseValidationStarknetValidationResultSchema) implementsStarknetTransactionScanResponseValidation() {
}

type StarknetTransactionScanResponseValidationStarknetValidationResultSchemaFeature struct {
	// Address the feature refers to
	Address string `json:"address,required"`
	// Textual description
	Description string `json:"description,required"`
	FeatureID string `json:"feature_id,required"`
	// Feature Classification
	Type StarknetTransactionScanResponseValidationStarknetValidationResultSchemaFeaturesType `json:"type,required"`
	JSON starknetTransactionScanResponseValidationStarknetValidationResultSchemaFeatureJSON  `json:"-"`
}

// starknetTransactionScanResponseValidationStarknetValidationResultSchemaFeatureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseValidationStarknetValidationResultSchemaFeature]
type starknetTransactionScanResponseValidationStarknetValidationResultSchemaFeatureJSON struct {
	Address     apijson.Field
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseValidationStarknetValidationResultSchemaFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseValidationStarknetValidationResultSchemaFeatureJSON) RawJSON() string {
	return r.raw
}

// Feature Classification
type StarknetTransactionScanResponseValidationStarknetValidationResultSchemaFeaturesType string

const (
	StarknetTransactionScanResponseValidationStarknetValidationResultSchemaFeaturesTypeBenign    StarknetTransactionScanResponseValidationStarknetValidationResultSchemaFeaturesType = "Benign"
	StarknetTransactionScanResponseValidationStarknetValidationResultSchemaFeaturesTypeWarning   StarknetTransactionScanResponseValidationStarknetValidationResultSchemaFeaturesType = "Warning"
	StarknetTransactionScanResponseValidationStarknetValidationResultSchemaFeaturesTypeMalicious StarknetTransactionScanResponseValidationStarknetValidationResultSchemaFeaturesType = "Malicious"
	StarknetTransactionScanResponseValidationStarknetValidationResultSchemaFeaturesTypeInfo      StarknetTransactionScanResponseValidationStarknetValidationResultSchemaFeaturesType = "Info"
)

func (r StarknetTransactionScanResponseValidationStarknetValidationResultSchemaFeaturesType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseValidationStarknetValidationResultSchemaFeaturesTypeBenign, StarknetTransactionScanResponseValidationStarknetValidationResultSchemaFeaturesTypeWarning, StarknetTransactionScanResponseValidationStarknetValidationResultSchemaFeaturesTypeMalicious, StarknetTransactionScanResponseValidationStarknetValidationResultSchemaFeaturesTypeInfo:
		return true
	}
	return false
}

// Verdict of the validation
type StarknetTransactionScanResponseValidationStarknetValidationResultSchemaResultType string

const (
	StarknetTransactionScanResponseValidationStarknetValidationResultSchemaResultTypeBenign    StarknetTransactionScanResponseValidationStarknetValidationResultSchemaResultType = "Benign"
	StarknetTransactionScanResponseValidationStarknetValidationResultSchemaResultTypeWarning   StarknetTransactionScanResponseValidationStarknetValidationResultSchemaResultType = "Warning"
	StarknetTransactionScanResponseValidationStarknetValidationResultSchemaResultTypeMalicious StarknetTransactionScanResponseValidationStarknetValidationResultSchemaResultType = "Malicious"
)

func (r StarknetTransactionScanResponseValidationStarknetValidationResultSchemaResultType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseValidationStarknetValidationResultSchemaResultTypeBenign, StarknetTransactionScanResponseValidationStarknetValidationResultSchemaResultTypeWarning, StarknetTransactionScanResponseValidationStarknetValidationResultSchemaResultTypeMalicious:
		return true
	}
	return false
}

type StarknetTransactionScanResponseValidationStarknetValidationResultSchemaStatus string

const (
	StarknetTransactionScanResponseValidationStarknetValidationResultSchemaStatusSuccess StarknetTransactionScanResponseValidationStarknetValidationResultSchemaStatus = "Success"
)

func (r StarknetTransactionScanResponseValidationStarknetValidationResultSchemaStatus) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseValidationStarknetValidationResultSchemaStatusSuccess:
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

type StarknetTransactionScanParams struct {
	AccountAddress param.Field[string] `json:"account_address,required"`
	// The chain name or chain ID
	Chain param.Field[StarknetTransactionScanParamsChain] `json:"chain,required"`
	// Metadata
	Metadata    param.Field[StarknetTransactionScanParamsMetadataUnion]    `json:"metadata,required"`
	Transaction param.Field[StarknetTransactionScanParamsTransactionUnion] `json:"transaction,required"`
	// List of options to include in the response
	//
	// - `simulation`: Include simulation output in the response
	// - `validation`: Include security validation of the transaction in the response
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
	Type param.Field[StarknetTransactionScanParamsMetadataType] `json:"type,required"`
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
	Type param.Field[StarknetTransactionScanParamsMetadataStarknetInAppRequestMetadataType] `json:"type,required"`
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
	// The version of the transaction.
	Version param.Field[StarknetTransactionScanParamsTransactionVersion] `json:"version,required"`
	// The nonce of the transaction.
	Nonce param.Field[string] `json:"nonce,required"`
	// The address of the sender.
	SenderAddress param.Field[string]      `json:"sender_address"`
	Calldata      param.Field[interface{}] `json:"calldata,required"`
	// The maximum fee that the sender is willing to pay.
	MaxFee param.Field[string] `json:"max_fee"`
	// The id of the chain to which the transaction is sent.
	ChainID param.Field[string] `json:"chain_id"`
	// The nonce data availability mode.
	NonceDataAvailabilityMode param.Field[StarknetTransactionScanParamsTransactionNonceDataAvailabilityMode] `json:"nonce_data_availability_mode"`
	PaymasterData             param.Field[interface{}]                                                       `json:"paymaster_data,required"`
	AccountDeploymentData     param.Field[interface{}]                                                       `json:"account_deployment_data,required"`
	// The hash of the contract class.
	ClassHash param.Field[string] `json:"class_hash"`
	// The salt of the contract address.
	ContractAddressSalt param.Field[string]      `json:"contract_address_salt"`
	ConstructorCalldata param.Field[interface{}] `json:"constructor_calldata,required"`
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
