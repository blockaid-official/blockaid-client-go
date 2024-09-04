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
	// [map[string][]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff].
	AssetsDiffs interface{} `json:"assets_diffs,required"`
	// This field can have the runtime type of
	// [map[string][]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure].
	Exposures interface{} `json:"exposures,required"`
	// This field can have the runtime type of
	// [[]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAddressDetail].
	AddressDetails interface{} `json:"address_details,required"`
	// This field can have the runtime type of
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummary].
	AccountSummary interface{} `json:"account_summary,required"`
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
	Status         apijson.Field
	AssetsDiffs    apijson.Field
	Exposures      apijson.Field
	AddressDetails apijson.Field
	AccountSummary apijson.Field
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
	AssetsDiffs []StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiff `json:"assets_diffs"`
	// Total USD exposure for each of the spender addresses during the transaction
	TotalUsdExposure map[string]float64                                                                                `json:"total_usd_exposure"`
	JSON             starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummary]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryJSON struct {
	AccountExposures apijson.Field
	TotalUsdDiff     apijson.Field
	AssetsDiffs      apijson.Field
	TotalUsdExposure apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure struct {
	// This field can have the runtime type of
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAsset],
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAsset],
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAsset].
	Asset interface{} `json:"asset"`
	// This field can have the runtime type of
	// [map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaSpender],
	// [map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaSpender],
	// [map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaSpender].
	Spenders interface{}                                                                                                      `json:"spenders,required"`
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
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchema],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchema],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchema].
func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure) AsUnion() StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresUnion {
	return r.union
}

// Union satisfied by
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchema],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchema]
// or
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchema].
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresUnion interface {
	implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchema{}),
		},
	)
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchema struct {
	Asset StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAsset `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaSpender `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaJSON               `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchema]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchema) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAsset struct {
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
	Type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAsset]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC20`)
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAssetTypeErc20 StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAssetType = "ERC20"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAssetTypeErc20:
		return true
	}
	return false
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaSpender struct {
	// Approval value of the ERC20 token
	Approval string              `json:"approval,required"`
	Exposure []StarknetErc20Diff `json:"exposure,required"`
	// Expiration date of the approval
	Expiration time.Time `json:"expiration,nullable" format:"date-time"`
	// Summarized description of the exposure
	Summary string                                                                                                                                                                                                                                                                                                                  `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaSpenderJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaSpenderJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaSpender]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaSpenderJSON struct {
	Approval    apijson.Field
	Exposure    apijson.Field
	Expiration  apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaSpenderJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchema struct {
	Asset StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAsset `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaSpender `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaJSON               `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchema]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchema) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC721`)
	Type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAsset]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC721`)
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAssetTypeErc721 StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAssetType = "ERC721"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAssetTypeErc721:
		return true
	}
	return false
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaSpender struct {
	Exposure []StarknetErc721Diff `json:"exposure,required"`
	// Whether `setApprovalForAll` was invoked
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// Summarized description of the exposure
	Summary string                                                                                                                                                                                                                                                                                                                    `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaSpenderJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaSpenderJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaSpender]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaSpenderJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchema struct {
	Asset StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAsset `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaSpender `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaJSON               `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchema]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchema) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC1155`)
	Type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAsset]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC1155`)
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAssetTypeErc1155 StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAssetType = "ERC1155"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAssetTypeErc1155:
		return true
	}
	return false
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaSpender struct {
	Exposure []StarknetErc1155Diff `json:"exposure,required"`
	// Whether `setApprovalForAll` was invoked
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// Summarized description of the exposure
	Summary string                                                                                                                                                                                                                                                                                                                      `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaSpenderJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaSpenderJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaSpender]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaSpenderJSON) RawJSON() string {
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

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiff struct {
	// This field can have the runtime type of
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAsset],
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAsset],
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAsset].
	Asset interface{} `json:"asset"`
	// This field can have the runtime type of [StarknetErc20Diff],
	// [StarknetErc721Diff], [StarknetErc1155Diff].
	In interface{} `json:"in,required"`
	// This field can have the runtime type of [StarknetErc20Diff],
	// [StarknetErc721Diff], [StarknetErc1155Diff].
	Out   interface{}                                                                                                 `json:"out,required"`
	JSON  starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffJSON `json:"-"`
	union StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsUnion
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiff]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchema],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchema],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchema].
func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiff) AsUnion() StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchema],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchema]
// or
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchema].
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsUnion interface {
	implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchema{}),
		},
	)
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchema struct {
	Asset StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAsset `json:"asset,required"`
	// Details of the incoming transfer
	In StarknetErc20Diff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetErc20Diff                                                                                                                                                                                                                               `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchema]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchema) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAsset struct {
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
	Type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAsset]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC20`)
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAssetTypeErc20 StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAssetType = "ERC20"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAssetTypeErc20:
		return true
	}
	return false
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchema struct {
	Asset StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAsset `json:"asset,required"`
	// Details of the incoming transfer
	In StarknetErc721Diff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetErc721Diff                                                                                                                                                                                                                                `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchema]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchema) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC721`)
	Type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAsset]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC721`)
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAssetTypeErc721 StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAssetType = "ERC721"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAssetTypeErc721:
		return true
	}
	return false
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchema struct {
	Asset StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAsset `json:"asset,required"`
	// Details of the incoming transfer
	In StarknetErc1155Diff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetErc1155Diff                                                                                                                                                                                                                                 `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchema]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchema) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC1155`)
	Type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAsset]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC1155`)
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAssetTypeErc1155 StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAssetType = "ERC1155"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAssetTypeErc1155:
		return true
	}
	return false
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
	AccountAddress string `json:"account_address,required"`
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
	// This field can have the runtime type of
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAsset],
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAsset],
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAsset].
	Asset interface{} `json:"asset"`
	// This field can have the runtime type of [StarknetErc20Diff],
	// [StarknetErc721Diff], [StarknetErc1155Diff].
	In interface{} `json:"in,required"`
	// This field can have the runtime type of [StarknetErc20Diff],
	// [StarknetErc721Diff], [StarknetErc1155Diff].
	Out   interface{}                                                                                   `json:"out,required"`
	JSON  starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffJSON `json:"-"`
	union StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsUnion
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffJSON struct {
	Asset       apijson.Field
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
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchema],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchema],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchema].
func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff) AsUnion() StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchema],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchema]
// or
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchema].
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsUnion interface {
	implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchema{}),
		},
	)
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchema struct {
	Asset StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAsset `json:"asset,required"`
	// Details of the incoming transfer
	In StarknetErc20Diff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetErc20Diff                                                                                                                                                                                                                 `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchema]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchema) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAsset struct {
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
	Type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAsset]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC20`)
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAssetTypeErc20 StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAssetType = "ERC20"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DiffSchemaAssetTypeErc20:
		return true
	}
	return false
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchema struct {
	Asset StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAsset `json:"asset,required"`
	// Details of the incoming transfer
	In StarknetErc721Diff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetErc721Diff                                                                                                                                                                                                                  `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchema]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchema) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC721`)
	Type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAsset]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC721`)
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAssetTypeErc721 StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAssetType = "ERC721"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DiffSchemaAssetTypeErc721:
		return true
	}
	return false
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchema struct {
	Asset StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAsset `json:"asset,required"`
	// Details of the incoming transfer
	In StarknetErc1155Diff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetErc1155Diff                                                                                                                                                                                                                   `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchema]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchema) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC1155`)
	Type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAsset]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC1155`)
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAssetTypeErc1155 StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAssetType = "ERC1155"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetAccountSingleAssetDiffSchemaTypeErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DiffSchemaAssetTypeErc1155:
		return true
	}
	return false
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure struct {
	// This field can have the runtime type of
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAsset],
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAsset],
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAsset].
	Asset interface{} `json:"asset"`
	// This field can have the runtime type of
	// [map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaSpender],
	// [map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaSpender],
	// [map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaSpender].
	Spenders interface{}                                                                                 `json:"spenders,required"`
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
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchema],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchema],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchema].
func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure) AsUnion() StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresUnion {
	return r.union
}

// Union satisfied by
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchema],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchema]
// or
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchema].
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresUnion interface {
	implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchema{}),
		},
	)
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchema struct {
	Asset StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAsset `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaSpender `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaJSON               `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchema]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchema) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAsset struct {
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
	Type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAsset]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC20`)
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAssetTypeErc20 StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAssetType = "ERC20"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaAssetTypeErc20:
		return true
	}
	return false
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaSpender struct {
	// Approval value of the ERC20 token
	Approval string              `json:"approval,required"`
	Exposure []StarknetErc20Diff `json:"exposure,required"`
	// Expiration date of the approval
	Expiration time.Time `json:"expiration,nullable" format:"date-time"`
	// Summarized description of the exposure
	Summary string                                                                                                                                                                                                                                                                                             `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaSpenderJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaSpenderJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaSpender]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaSpenderJSON struct {
	Approval    apijson.Field
	Exposure    apijson.Field
	Expiration  apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc20ExposureSchemaSpenderJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchema struct {
	Asset StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAsset `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaSpender `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaJSON               `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchema]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchema) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC721`)
	Type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAsset]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC721`)
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAssetTypeErc721 StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAssetType = "ERC721"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaAssetTypeErc721:
		return true
	}
	return false
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaSpender struct {
	Exposure []StarknetErc721Diff `json:"exposure,required"`
	// Whether `setApprovalForAll` was invoked
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// Summarized description of the exposure
	Summary string                                                                                                                                                                                                                                                                                               `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaSpenderJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaSpenderJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaSpender]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc721ExposureSchemaSpenderJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchema struct {
	Asset StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAsset `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaSpender `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaJSON               `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchema]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchema) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC1155`)
	Type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAssetType `json:"type"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAssetJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAssetJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAsset]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC1155`)
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAssetType string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAssetTypeErc1155 StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAssetType = "ERC1155"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAssetType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaAssetTypeErc1155:
		return true
	}
	return false
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaSpender struct {
	Exposure []StarknetErc1155Diff `json:"exposure,required"`
	// Whether `setApprovalForAll` was invoked
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// Summarized description of the exposure
	Summary string                                                                                                                                                                                                                                                                                                 `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaSpenderJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaSpenderJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaSpender]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetAddressAssetExposureSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155DetailsSchemaAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaErc1155ExposureSchemaSpenderJSON) RawJSON() string {
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
	// [[]StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaFeature].
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
// [StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchema],
// [StarknetTransactionScanResponseValidationStarknetValidationErrorSchema].
func (r StarknetTransactionScanResponseValidation) AsUnion() StarknetTransactionScanResponseValidationUnion {
	return r.union
}

// Validation result; Only present if validation option is included in the request
//
// Union satisfied by
// [StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchema]
// or [StarknetTransactionScanResponseValidationStarknetValidationErrorSchema].
type StarknetTransactionScanResponseValidationUnion interface {
	implementsStarknetTransactionScanResponseValidation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StarknetTransactionScanResponseValidationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseValidationStarknetValidationErrorSchema{}),
		},
	)
}

type StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchema struct {
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification,required"`
	// A textual description about the validation result
	Description string `json:"description,required"`
	// A list of features about this transaction explaining the validation
	Features []StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaFeature `json:"features,required"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason string `json:"reason,required"`
	// Verdict of the validation
	ResultType StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaResultType `json:"result_type,required"`
	Status     StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaStatus     `json:"status,required"`
	JSON       starknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaJSON       `json:"-"`
}

// starknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchema]
type starknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaJSON struct {
	Classification apijson.Field
	Description    apijson.Field
	Features       apijson.Field
	Reason         apijson.Field
	ResultType     apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchema) implementsStarknetTransactionScanResponseValidation() {
}

type StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaFeature struct {
	// Address the feature refers to
	Address string `json:"address,required"`
	// Textual description
	Description string `json:"description,required"`
	FeatureID string `json:"feature_id,required"`
	// Feature Classification
	Type StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaFeaturesType `json:"type,required"`
	JSON starknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaFeatureJSON  `json:"-"`
}

// starknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaFeatureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaFeature]
type starknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaFeatureJSON struct {
	Address     apijson.Field
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaFeatureJSON) RawJSON() string {
	return r.raw
}

// Feature Classification
type StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaFeaturesType string

const (
	StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaFeaturesTypeBenign    StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaFeaturesType = "Benign"
	StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaFeaturesTypeWarning   StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaFeaturesType = "Warning"
	StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaFeaturesTypeMalicious StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaFeaturesType = "Malicious"
	StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaFeaturesTypeInfo      StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaFeaturesType = "Info"
)

func (r StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaFeaturesType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaFeaturesTypeBenign, StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaFeaturesTypeWarning, StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaFeaturesTypeMalicious, StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaFeaturesTypeInfo:
		return true
	}
	return false
}

// Verdict of the validation
type StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaResultType string

const (
	StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaResultTypeBenign    StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaResultType = "Benign"
	StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaResultTypeWarning   StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaResultType = "Warning"
	StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaResultTypeMalicious StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaResultType = "Malicious"
)

func (r StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaResultType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaResultTypeBenign, StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaResultTypeWarning, StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaResultTypeMalicious:
		return true
	}
	return false
}

type StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaStatus string

const (
	StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaStatusSuccess StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaStatus = "Success"
)

func (r StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaStatus) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseValidationStarknetValidationResultSchemaTypeAnnotatedIntSkipValidationPlainSerializerGetPydanticSchemaStatusSuccess:
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
