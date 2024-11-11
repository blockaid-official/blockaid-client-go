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
	// This field can have the runtime type of
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummary].
	AccountSummary interface{} `json:"account_summary,required"`
	// This field can have the runtime type of
	// [[]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAddressDetail].
	AddressDetails interface{} `json:"address_details,required"`
	// This field can have the runtime type of
	// [map[string][]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff].
	AssetsDiffs interface{} `json:"assets_diffs,required"`
	// This field can have the runtime type of
	// [map[string][]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure].
	Exposures interface{}                                     `json:"exposures,required"`
	Status    StarknetTransactionScanResponseSimulationStatus `json:"status,required"`
	// Optional block number or tag context for the simulation
	BlockNumber string `json:"block_number,nullable"`
	// Error message
	Error string                                        `json:"error"`
	JSON  starknetTransactionScanResponseSimulationJSON `json:"-"`
	union StarknetTransactionScanResponseSimulationUnion
}

// starknetTransactionScanResponseSimulationJSON contains the JSON metadata for the
// struct [StarknetTransactionScanResponseSimulation]
type starknetTransactionScanResponseSimulationJSON struct {
	AccountSummary apijson.Field
	AddressDetails apijson.Field
	AssetsDiffs    apijson.Field
	Exposures      apijson.Field
	Status         apijson.Field
	BlockNumber    apijson.Field
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
	// This field can have the runtime type of
	// [map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpender],
	// [map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpender],
	// [map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpender].
	Spenders interface{} `json:"spenders,required"`
	// This field can have the runtime type of
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureAsset],
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureAsset],
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureAsset].
	Asset interface{}                                                                                                      `json:"asset"`
	JSON  starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposureJSON `json:"-"`
	union StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresUnion
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposureJSON struct {
	Spenders    apijson.Field
	Asset       apijson.Field
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
	Asset StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureAsset `json:"asset,required"`
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

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's decimals
	Decimals int64 `json:"decimals,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC20`)
	Type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureAsset]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC20`)
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureAssetTypeErc20 StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureAssetType = "ERC20"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureAssetTypeErc20:
		return true
	}
	return false
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpender struct {
	// Approval value of the ERC20 token
	Approval string                                                                                                                                               `json:"approval,required"`
	Exposure []StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpendersExposure `json:"exposure,required"`
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

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpendersExposure struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                                                 `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpendersExposureJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpendersExposure]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpendersExposureJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721Exposure struct {
	Asset StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureAsset `json:"asset,required"`
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

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC721`)
	Type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureAsset]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC721`)
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureAssetTypeErc721 StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureAssetType = "ERC721"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureAssetTypeErc721:
		return true
	}
	return false
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpender struct {
	Exposure []StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpendersExposure `json:"exposure,required"`
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

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpendersExposure struct {
	// Token ID of the transfer
	TokenID string `json:"token_id,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                                                  `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpendersExposureJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpendersExposure]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpendersExposureJSON struct {
	TokenID     apijson.Field
	UsdPrice    apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155Exposure struct {
	Asset StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureAsset `json:"asset,required"`
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

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC1155`)
	Type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureAsset]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC1155`)
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureAssetTypeErc1155 StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureAssetType = "ERC1155"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureAssetTypeErc1155:
		return true
	}
	return false
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpender struct {
	Exposure []StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpendersExposure `json:"exposure,required"`
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

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpendersExposure struct {
	// Token ID of the transfer
	TokenID string `json:"token_id,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value int64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                                                   `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpendersExposureJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpendersExposure]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpendersExposureJSON struct {
	TokenID     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpendersExposureJSON) RawJSON() string {
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
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffIn],
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffIn],
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffIn].
	In interface{} `json:"in,required"`
	// This field can have the runtime type of
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffOut],
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffOut],
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffOut].
	Out interface{} `json:"out,required"`
	// This field can have the runtime type of
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffAsset],
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffAsset],
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffAsset].
	Asset interface{}                                                                                                        `json:"asset"`
	JSON  starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffJSON `json:"-"`
	union StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsUnion
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiff]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffJSON struct {
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	Asset       apijson.Field
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
	Asset StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffOut  `json:"out,nullable"`
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

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's decimals
	Decimals int64 `json:"decimals,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC20`)
	Type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffAsset]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC20`)
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffAssetTypeErc20 StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffAssetType = "ERC20"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffAssetTypeErc20:
		return true
	}
	return false
}

// Details of the incoming transfer
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                                      `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffInJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffInJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffIn]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffInJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                                       `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffOutJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffOutJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffOut]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffOutJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiff struct {
	Asset StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffOut  `json:"out,nullable"`
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

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC721`)
	Type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffAsset]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC721`)
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffAssetTypeErc721 StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffAssetType = "ERC721"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffAssetTypeErc721:
		return true
	}
	return false
}

// Details of the incoming transfer
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffIn struct {
	// Token ID of the transfer
	TokenID string `json:"token_id,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                                       `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffInJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffInJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffIn]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffInJSON struct {
	TokenID     apijson.Field
	UsdPrice    apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffOut struct {
	// Token ID of the transfer
	TokenID string `json:"token_id,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                                        `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffOutJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffOutJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffOut]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffOutJSON struct {
	TokenID     apijson.Field
	UsdPrice    apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiff struct {
	Asset StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffOut  `json:"out,nullable"`
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

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC1155`)
	Type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffAsset]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC1155`)
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffAssetTypeErc1155 StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffAssetType = "ERC1155"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffAssetTypeErc1155:
		return true
	}
	return false
}

// Details of the incoming transfer
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffIn struct {
	// Token ID of the transfer
	TokenID string `json:"token_id,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value int64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                                        `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffInJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffInJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffIn]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffInJSON struct {
	TokenID     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffOut struct {
	// Token ID of the transfer
	TokenID string `json:"token_id,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value int64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                                         `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffOutJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffOutJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffOut]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffOutJSON struct {
	TokenID     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffOutJSON) RawJSON() string {
	return r.raw
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
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffIn],
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffIn],
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffIn].
	In interface{} `json:"in,required"`
	// This field can have the runtime type of
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffOut],
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffOut],
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffOut].
	Out interface{} `json:"out,required"`
	// This field can have the runtime type of
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffAsset],
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffAsset],
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffAsset].
	Asset interface{}                                                                                   `json:"asset"`
	JSON  starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffJSON `json:"-"`
	union StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsUnion
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffJSON struct {
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	Asset       apijson.Field
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
	Asset StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffOut  `json:"out,nullable"`
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

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's decimals
	Decimals int64 `json:"decimals,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC20`)
	Type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffAsset]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC20`)
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffAssetTypeErc20 StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffAssetType = "ERC20"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffAssetTypeErc20:
		return true
	}
	return false
}

// Details of the incoming transfer
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                 `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffInJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffInJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffIn]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffInJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                  `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffOutJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffOutJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffOut]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffOutJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiff struct {
	Asset StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffOut  `json:"out,nullable"`
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

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC721`)
	Type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffAsset]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC721`)
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffAssetTypeErc721 StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffAssetType = "ERC721"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffAssetTypeErc721:
		return true
	}
	return false
}

// Details of the incoming transfer
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffIn struct {
	// Token ID of the transfer
	TokenID string `json:"token_id,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                  `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffInJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffInJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffIn]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffInJSON struct {
	TokenID     apijson.Field
	UsdPrice    apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffOut struct {
	// Token ID of the transfer
	TokenID string `json:"token_id,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                   `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffOutJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffOutJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffOut]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffOutJSON struct {
	TokenID     apijson.Field
	UsdPrice    apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiff struct {
	Asset StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffOut  `json:"out,nullable"`
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

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC1155`)
	Type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffAsset]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC1155`)
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffAssetTypeErc1155 StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffAssetType = "ERC1155"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffAssetTypeErc1155:
		return true
	}
	return false
}

// Details of the incoming transfer
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffIn struct {
	// Token ID of the transfer
	TokenID string `json:"token_id,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value int64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                   `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffInJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffInJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffIn]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffInJSON struct {
	TokenID     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffOut struct {
	// Token ID of the transfer
	TokenID string `json:"token_id,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value int64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                    `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffOutJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffOutJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffOut]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffOutJSON struct {
	TokenID     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure struct {
	// This field can have the runtime type of
	// [map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpender],
	// [map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpender],
	// [map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpender].
	Spenders interface{} `json:"spenders,required"`
	// This field can have the runtime type of
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureAsset],
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureAsset],
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureAsset].
	Asset interface{}                                                                                 `json:"asset"`
	JSON  starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposureJSON `json:"-"`
	union StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresUnion
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposureJSON struct {
	Spenders    apijson.Field
	Asset       apijson.Field
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
	Asset StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureAsset `json:"asset,required"`
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

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's decimals
	Decimals int64 `json:"decimals,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC20`)
	Type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureAsset]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC20`)
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureAssetTypeErc20 StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureAssetType = "ERC20"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureAssetTypeErc20:
		return true
	}
	return false
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpender struct {
	// Approval value of the ERC20 token
	Approval string                                                                                                                          `json:"approval,required"`
	Exposure []StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpendersExposure `json:"exposure,required"`
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

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpendersExposure struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                            `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpendersExposureJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpendersExposure]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpendersExposureJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721Exposure struct {
	Asset StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureAsset `json:"asset,required"`
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

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC721`)
	Type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureAsset]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC721`)
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureAssetTypeErc721 StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureAssetType = "ERC721"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureAssetTypeErc721:
		return true
	}
	return false
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpender struct {
	Exposure []StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpendersExposure `json:"exposure,required"`
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

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpendersExposure struct {
	// Token ID of the transfer
	TokenID string `json:"token_id,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                             `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpendersExposureJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpendersExposure]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpendersExposureJSON struct {
	TokenID     apijson.Field
	UsdPrice    apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155Exposure struct {
	Asset StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureAsset `json:"asset,required"`
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

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC1155`)
	Type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureAsset]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC1155`)
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureAssetTypeErc1155 StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureAssetType = "ERC1155"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureAssetTypeErc1155:
		return true
	}
	return false
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpender struct {
	Exposure []StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpendersExposure `json:"exposure,required"`
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

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpendersExposure struct {
	// Token ID of the transfer
	TokenID string `json:"token_id,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value int64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                              `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpendersExposureJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpendersExposure]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpendersExposureJSON struct {
	TokenID     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpendersExposureJSON) RawJSON() string {
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
	// This field can have the runtime type of
	// [[]StarknetTransactionScanResponseValidationStarknetValidationResultFeature].
	Features interface{}                                     `json:"features,required"`
	Status   StarknetTransactionScanResponseValidationStatus `json:"status,required"`
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification"`
	// A textual description about the validation result
	Description string `json:"description"`
	// Error message
	Error string `json:"error"`
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
	Features       apijson.Field
	Status         apijson.Field
	Classification apijson.Field
	Description    apijson.Field
	Error          apijson.Field
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
	Params param.Field[interface{}]                               `json:"params,required"`
	ID     param.Field[string]                                    `json:"id"`
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
	AccountDeploymentData param.Field[interface{}] `json:"account_deployment_data,required"`
	Calldata              param.Field[interface{}] `json:"calldata,required"`
	ConstructorCalldata   param.Field[interface{}] `json:"constructor_calldata,required"`
	// The nonce of the transaction.
	Nonce         param.Field[string]      `json:"nonce,required"`
	PaymasterData param.Field[interface{}] `json:"paymaster_data,required"`
	// The version of the transaction.
	Version param.Field[StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionVersion] `json:"version,required"`
	// The id of the chain to which the transaction is sent.
	ChainID param.Field[string] `json:"chain_id"`
	// The hash of the contract class.
	ClassHash param.Field[string] `json:"class_hash"`
	// The salt of the contract address.
	ContractAddressSalt param.Field[string] `json:"contract_address_salt"`
	// The maximum fee that the sender is willing to pay.
	MaxFee param.Field[string] `json:"max_fee"`
	// The nonce data availability mode.
	NonceDataAvailabilityMode param.Field[StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportParamsTransactionNonceDataAvailabilityMode] `json:"nonce_data_availability_mode"`
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
	AccountDeploymentData param.Field[interface{}] `json:"account_deployment_data,required"`
	Calldata              param.Field[interface{}] `json:"calldata,required"`
	ConstructorCalldata   param.Field[interface{}] `json:"constructor_calldata,required"`
	// The nonce of the transaction.
	Nonce         param.Field[string]      `json:"nonce,required"`
	PaymasterData param.Field[interface{}] `json:"paymaster_data,required"`
	// The version of the transaction.
	Version param.Field[StarknetTransactionScanParamsTransactionVersion] `json:"version,required"`
	// The id of the chain to which the transaction is sent.
	ChainID param.Field[string] `json:"chain_id"`
	// The hash of the contract class.
	ClassHash param.Field[string] `json:"class_hash"`
	// The salt of the contract address.
	ContractAddressSalt param.Field[string] `json:"contract_address_salt"`
	// The maximum fee that the sender is willing to pay.
	MaxFee param.Field[string] `json:"max_fee"`
	// The nonce data availability mode.
	NonceDataAvailabilityMode param.Field[StarknetTransactionScanParamsTransactionNonceDataAvailabilityMode] `json:"nonce_data_availability_mode"`
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
