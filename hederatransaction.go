// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"net/http"
	"reflect"
	"slices"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/param"
	"github.com/blockaid-official/blockaid-client-go/internal/requestconfig"
	"github.com/blockaid-official/blockaid-client-go/option"
	"github.com/tidwall/gjson"
)

// HederaTransactionService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHederaTransactionService] method instead.
type HederaTransactionService struct {
	Options []option.RequestOption
}

// NewHederaTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHederaTransactionService(opts ...option.RequestOption) (r *HederaTransactionService) {
	r = &HederaTransactionService{}
	r.Options = opts
	return
}

// Get a risk recommendation with plain-language reasons for a Hedera transaction.
func (r *HederaTransactionService) Scan(ctx context.Context, body HederaTransactionScanParams, opts ...option.RequestOption) (res *HederaTransactionScanResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/hedera/transaction/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Transaction scan response schema.
type HederaTransactionScanResponse struct {
	// Simulation result; Only present if simulation option is included in the request
	Simulation HederaTransactionScanResponseSimulation `json:"simulation,nullable"`
	// Validation result; Only present if validation option is included in the request
	Validation HederaTransactionScanResponseValidation `json:"validation,nullable"`
	JSON       hederaTransactionScanResponseJSON       `json:"-"`
}

// hederaTransactionScanResponseJSON contains the JSON metadata for the struct
// [HederaTransactionScanResponse]
type hederaTransactionScanResponseJSON struct {
	Simulation  apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseJSON) RawJSON() string {
	return r.raw
}

// Simulation result; Only present if simulation option is included in the request
type HederaTransactionScanResponseSimulation struct {
	Status HederaTransactionScanResponseSimulationStatus `json:"status,required"`
	// This field can have the runtime type of
	// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummary].
	AccountSummary interface{} `json:"account_summary"`
	// This field can have the runtime type of
	// [[]HederaTransactionScanResponseSimulationHederaSimulationResponseAddressDetail].
	AddressDetails interface{} `json:"address_details"`
	// This field can have the runtime type of
	// [map[string][]HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiff].
	AssetsDiffs interface{} `json:"assets_diffs"`
	// Error message
	Error string `json:"error"`
	// This field can have the runtime type of
	// [map[string][]HederaTransactionScanResponseSimulationHederaSimulationResponseExposure].
	Exposures interface{} `json:"exposures"`
	// This field can have the runtime type of
	// [[]HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActions].
	TransactionActions interface{}                                 `json:"transaction_actions"`
	JSON               hederaTransactionScanResponseSimulationJSON `json:"-"`
	union              HederaTransactionScanResponseSimulationUnion
}

// hederaTransactionScanResponseSimulationJSON contains the JSON metadata for the
// struct [HederaTransactionScanResponseSimulation]
type hederaTransactionScanResponseSimulationJSON struct {
	Status             apijson.Field
	AccountSummary     apijson.Field
	AddressDetails     apijson.Field
	AssetsDiffs        apijson.Field
	Error              apijson.Field
	Exposures          apijson.Field
	TransactionActions apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r hederaTransactionScanResponseSimulationJSON) RawJSON() string {
	return r.raw
}

func (r *HederaTransactionScanResponseSimulation) UnmarshalJSON(data []byte) (err error) {
	*r = HederaTransactionScanResponseSimulation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [HederaTransactionScanResponseSimulationUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [HederaTransactionScanResponseSimulationHederaSimulationResponse],
// [HederaTransactionScanResponseSimulationHederaSimulationErrorSchema].
func (r HederaTransactionScanResponseSimulation) AsUnion() HederaTransactionScanResponseSimulationUnion {
	return r.union
}

// Simulation result; Only present if simulation option is included in the request
//
// Union satisfied by
// [HederaTransactionScanResponseSimulationHederaSimulationResponse] or
// [HederaTransactionScanResponseSimulationHederaSimulationErrorSchema].
type HederaTransactionScanResponseSimulationUnion interface {
	implementsHederaTransactionScanResponseSimulation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*HederaTransactionScanResponseSimulationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HederaTransactionScanResponseSimulationHederaSimulationResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HederaTransactionScanResponseSimulationHederaSimulationErrorSchema{}),
		},
	)
}

type HederaTransactionScanResponseSimulationHederaSimulationResponse struct {
	// Summary of the actions and asset transfers that were made by the requested
	// account address
	AccountSummary HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummary `json:"account_summary,required"`
	Status         HederaTransactionScanResponseSimulationHederaSimulationResponseStatus         `json:"status,required"`
	// Details of addresses involved in the transaction
	AddressDetails []HederaTransactionScanResponseSimulationHederaSimulationResponseAddressDetail `json:"address_details"`
	// Mapping between the address of an account to the assets diff during the
	// transaction
	AssetsDiffs map[string][]HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiff `json:"assets_diffs"`
	// Mapping between the address of an account to the exposure of the assets during
	// the transaction
	Exposures          map[string][]HederaTransactionScanResponseSimulationHederaSimulationResponseExposure `json:"exposures"`
	TransactionActions []HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActions  `json:"transaction_actions,nullable"`
	JSON               hederaTransactionScanResponseSimulationHederaSimulationResponseJSON                  `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseJSON contains the
// JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponse]
type hederaTransactionScanResponseSimulationHederaSimulationResponseJSON struct {
	AccountSummary     apijson.Field
	Status             apijson.Field
	AddressDetails     apijson.Field
	AssetsDiffs        apijson.Field
	Exposures          apijson.Field
	TransactionActions apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseJSON) RawJSON() string {
	return r.raw
}

func (r HederaTransactionScanResponseSimulationHederaSimulationResponse) implementsHederaTransactionScanResponseSimulation() {
}

// Summary of the actions and asset transfers that were made by the requested
// account address
type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummary struct {
	// Exposures made by the requested account address
	AccountExposures []HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposure `json:"account_exposures,required"`
	// Total USD diff for the requested account address
	TotalUsdDiff HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryTotalUsdDiff `json:"total_usd_diff,required"`
	// Assets diffs of the requested account address
	AccountAssetsDiffs []HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiff `json:"account_assets_diffs"`
	// Total USD exposure for each of the spender addresses during the transaction
	TotalUsdExposure map[string]float64                                                                `json:"total_usd_exposure"`
	JSON             hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummary]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryJSON struct {
	AccountExposures   apijson.Field
	TotalUsdDiff       apijson.Field
	AccountAssetsDiffs apijson.Field
	TotalUsdExposure   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryJSON) RawJSON() string {
	return r.raw
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposure struct {
	Asset HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAsset `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresSpender `json:"spenders"`
	JSON     hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposureJSON                `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposureJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposure]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposureJSON) RawJSON() string {
	return r.raw
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAsset struct {
	// The token ID
	ID string `json:"id"`
	// Decimals of the asset
	Decimals int64 `json:"decimals"`
	// The token description
	Description string `json:"description"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// The token name
	Name string `json:"name"`
	// The NFT's collection ID (token ID)
	NFTType string `json:"nft_type"`
	// The token's symbol
	Symbol string `json:"symbol"`
	// Type of the asset (`NATIVE`)
	Type  HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetType `json:"type"`
	JSON  hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetJSON `json:"-"`
	union HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetUnion
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAsset]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetJSON struct {
	ID          apijson.Field
	Decimals    apijson.Field
	Description apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	NFTType     apijson.Field
	Symbol      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetJSON) RawJSON() string {
	return r.raw
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAsset) UnmarshalJSON(data []byte) (err error) {
	*r = HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNativeAssetDetailsSchema],
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaTokenDetailsSchema],
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNFTDetailsSchema].
func (r HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAsset) AsUnion() HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetUnion {
	return r.union
}

// Union satisfied by
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNativeAssetDetailsSchema],
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaTokenDetailsSchema]
// or
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNFTDetailsSchema].
type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetUnion interface {
	implementsHederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNativeAssetDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaTokenDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNFTDetailsSchema{}),
		},
	)
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNativeAssetDetailsSchema struct {
	// Type of the asset (`NATIVE`)
	Type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNativeAssetDetailsSchemaType `json:"type"`
	JSON hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNativeAssetDetailsSchemaJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNativeAssetDetailsSchemaJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNativeAssetDetailsSchema]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNativeAssetDetailsSchemaJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNativeAssetDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNativeAssetDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNativeAssetDetailsSchema) implementsHederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAsset() {
}

// Type of the asset (`NATIVE`)
type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNativeAssetDetailsSchemaType string

const (
	HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNativeAssetDetailsSchemaTypeNative HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNativeAssetDetailsSchemaType = "NATIVE"
)

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNativeAssetDetailsSchemaType) IsKnown() bool {
	switch r {
	case HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNativeAssetDetailsSchemaTypeNative:
		return true
	}
	return false
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaTokenDetailsSchema struct {
	// The token ID
	ID string `json:"id,required"`
	// Decimals of the asset
	Decimals int64 `json:"decimals,required"`
	// The token description
	Description string `json:"description,required"`
	// The token name
	Name string `json:"name,required"`
	// The token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`TOKEN`)
	Type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaTokenDetailsSchemaType `json:"type"`
	JSON hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaTokenDetailsSchemaJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaTokenDetailsSchemaJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaTokenDetailsSchema]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaTokenDetailsSchemaJSON struct {
	ID          apijson.Field
	Decimals    apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaTokenDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaTokenDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaTokenDetailsSchema) implementsHederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAsset() {
}

// Type of the asset (`TOKEN`)
type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaTokenDetailsSchemaType string

const (
	HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaTokenDetailsSchemaTypeToken HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaTokenDetailsSchemaType = "TOKEN"
)

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaTokenDetailsSchemaType) IsKnown() bool {
	switch r {
	case HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaTokenDetailsSchemaTypeToken:
		return true
	}
	return false
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNFTDetailsSchema struct {
	// The NFT ID
	ID string `json:"id,required"`
	// The NFT's description
	Description string `json:"description,required"`
	// NFT's display name
	Name string `json:"name,required"`
	// The NFT's collection ID (token ID)
	NFTType string `json:"nft_type,required"`
	// URL of the NFT's image
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`NFT`)
	Type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNFTDetailsSchemaType `json:"type"`
	JSON hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNFTDetailsSchemaJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNFTDetailsSchemaJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNFTDetailsSchema]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNFTDetailsSchemaJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	NFTType     apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNFTDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNFTDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNFTDetailsSchema) implementsHederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAsset() {
}

// Type of the asset (`NFT`)
type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNFTDetailsSchemaType string

const (
	HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNFTDetailsSchemaTypeNFT HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNFTDetailsSchemaType = "NFT"
)

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNFTDetailsSchemaType) IsKnown() bool {
	switch r {
	case HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetHederaNFTDetailsSchemaTypeNFT:
		return true
	}
	return false
}

// Type of the asset (`NATIVE`)
type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetType string

const (
	HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetTypeNative HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetType = "NATIVE"
	HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetTypeToken  HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetType = "TOKEN"
	HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetTypeNFT    HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetType = "NFT"
)

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetType) IsKnown() bool {
	switch r {
	case HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetTypeNative, HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetTypeToken, HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresAssetTypeNFT:
		return true
	}
	return false
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresSpender struct {
	Exposure []HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresSpendersExposure `json:"exposure,required"`
	// Summarized description of the exposure
	Summary string                                                                                                   `json:"summary,nullable"`
	JSON    hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresSpenderJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresSpenderJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresSpender]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresSpenderJSON struct {
	Exposure    apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresSpenderJSON) RawJSON() string {
	return r.raw
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresSpendersExposure struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                           `json:"usd_price,nullable"`
	JSON     hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresSpendersExposureJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresSpendersExposureJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresSpendersExposure]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresSpendersExposureJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountExposuresSpendersExposureJSON) RawJSON() string {
	return r.raw
}

// Total USD diff for the requested account address
type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryTotalUsdDiff struct {
	// Total incoming USD transfers
	In float64 `json:"in,required"`
	// Total outgoing USD transfers
	Out float64 `json:"out,required"`
	// Total USD transfers
	Total float64                                                                                       `json:"total"`
	JSON  hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryTotalUsdDiffJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryTotalUsdDiffJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryTotalUsdDiff]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiff struct {
	// This field can have the runtime type of
	// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAsset],
	// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAsset],
	// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAsset].
	Asset interface{} `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of
	// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaIn],
	// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaIn],
	// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaIn].
	In interface{} `json:"in"`
	// This field can have the runtime type of
	// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaOut],
	// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaOut],
	// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaOut].
	Out   interface{}                                                                                        `json:"out"`
	JSON  hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffJSON `json:"-"`
	union HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsUnion
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiff]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchema],
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchema],
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchema].
func (r HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiff) AsUnion() HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchema],
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchema]
// or
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchema].
type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsUnion interface {
	implementsHederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchema{}),
		},
	)
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchema struct {
	Asset HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaOut  `json:"out,nullable"`
	JSON hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchema]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchema) implementsHederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiff() {
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAsset struct {
	// Type of the asset (`NATIVE`)
	Type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAssetType `json:"type"`
	JSON hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAssetJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAssetJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAsset]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAssetJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`NATIVE`)
type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAssetType string

const (
	HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAssetTypeNative HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAssetType = "NATIVE"
)

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAssetType) IsKnown() bool {
	switch r {
	case HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAssetTypeNative:
		return true
	}
	return false
}

// Details of the incoming transfer
type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                                                `json:"usd_price,nullable"`
	JSON     hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaInJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaInJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaIn]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                                                 `json:"usd_price,nullable"`
	JSON     hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaOutJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaOutJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaOut]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaOutJSON) RawJSON() string {
	return r.raw
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchema struct {
	Asset HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaOut  `json:"out,nullable"`
	JSON hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchema]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchema) implementsHederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiff() {
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAsset struct {
	// The token ID
	ID string `json:"id,required"`
	// Decimals of the asset
	Decimals int64 `json:"decimals,required"`
	// The token description
	Description string `json:"description,required"`
	// The token name
	Name string `json:"name,required"`
	// The token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`TOKEN`)
	Type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAssetType `json:"type"`
	JSON hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAssetJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAssetJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAsset]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAssetJSON struct {
	ID          apijson.Field
	Decimals    apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`TOKEN`)
type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAssetType string

const (
	HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAssetTypeToken HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAssetType = "TOKEN"
)

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAssetType) IsKnown() bool {
	switch r {
	case HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAssetTypeToken:
		return true
	}
	return false
}

// Details of the incoming transfer
type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                                          `json:"usd_price,nullable"`
	JSON     hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaInJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaInJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaIn]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                                           `json:"usd_price,nullable"`
	JSON     hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaOutJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaOutJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaOut]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaOutJSON) RawJSON() string {
	return r.raw
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchema struct {
	Asset HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaOut  `json:"out,nullable"`
	JSON hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchema]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchema) implementsHederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiff() {
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAsset struct {
	// The NFT ID
	ID string `json:"id,required"`
	// The NFT's description
	Description string `json:"description,required"`
	// NFT's display name
	Name string `json:"name,required"`
	// The NFT's collection ID (token ID)
	NFTType string `json:"nft_type,required"`
	// URL of the NFT's image
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`NFT`)
	Type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAssetType `json:"type"`
	JSON hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAssetJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAsset]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAssetJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	NFTType     apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`NFT`)
type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAssetType string

const (
	HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAssetTypeNFT HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAssetType = "NFT"
)

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAssetType) IsKnown() bool {
	switch r {
	case HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAssetTypeNFT:
		return true
	}
	return false
}

// Details of the incoming transfer
type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaIn struct {
	// Token ID of the transfer
	TokenID string `json:"token_id,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                               `json:"usd_price,nullable"`
	JSON     hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaInJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaInJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaIn]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaInJSON struct {
	TokenID     apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaOut struct {
	// Token ID of the transfer
	TokenID string `json:"token_id,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                                `json:"usd_price,nullable"`
	JSON     hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaOutJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaOutJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaOut]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaOutJSON struct {
	TokenID     apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAccountSummaryAccountAssetsDiffsNFTDetailsSchemaErc721DiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseStatus string

const (
	HederaTransactionScanResponseSimulationHederaSimulationResponseStatusSuccess HederaTransactionScanResponseSimulationHederaSimulationResponseStatus = "Success"
)

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseStatus) IsKnown() bool {
	switch r {
	case HederaTransactionScanResponseSimulationHederaSimulationResponseStatusSuccess:
		return true
	}
	return false
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseAddressDetail struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// Description of the account
	Description string                                                                           `json:"description,nullable"`
	JSON        hederaTransactionScanResponseSimulationHederaSimulationResponseAddressDetailJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAddressDetailJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAddressDetail]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAddressDetailJSON struct {
	AccountAddress apijson.Field
	Description    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAddressDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAddressDetailJSON) RawJSON() string {
	return r.raw
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiff struct {
	// This field can have the runtime type of
	// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAsset],
	// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAsset],
	// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAsset].
	Asset interface{} `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of
	// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaIn],
	// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaIn],
	// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaIn].
	In interface{} `json:"in"`
	// This field can have the runtime type of
	// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaOut],
	// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaOut],
	// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaOut].
	Out   interface{}                                                                   `json:"out"`
	JSON  hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffJSON `json:"-"`
	union HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsUnion
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiff]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchema],
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchema],
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchema].
func (r HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiff) AsUnion() HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchema],
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchema]
// or
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchema].
type HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsUnion interface {
	implementsHederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchema{}),
		},
	)
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchema struct {
	Asset HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaOut  `json:"out,nullable"`
	JSON hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchema]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchema) implementsHederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiff() {
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAsset struct {
	// Type of the asset (`NATIVE`)
	Type HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAssetType `json:"type"`
	JSON hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAssetJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAssetJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAsset]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAssetJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`NATIVE`)
type HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAssetType string

const (
	HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAssetTypeNative HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAssetType = "NATIVE"
)

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAssetType) IsKnown() bool {
	switch r {
	case HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaAssetTypeNative:
		return true
	}
	return false
}

// Details of the incoming transfer
type HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                           `json:"usd_price,nullable"`
	JSON     hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaInJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaInJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaIn]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                            `json:"usd_price,nullable"`
	JSON     hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaOutJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaOutJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaOut]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountNativeAssetTransferDiffDetailsSchemaOutJSON) RawJSON() string {
	return r.raw
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchema struct {
	Asset HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaOut  `json:"out,nullable"`
	JSON hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchema]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchema) implementsHederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiff() {
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAsset struct {
	// The token ID
	ID string `json:"id,required"`
	// Decimals of the asset
	Decimals int64 `json:"decimals,required"`
	// The token description
	Description string `json:"description,required"`
	// The token name
	Name string `json:"name,required"`
	// The token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`TOKEN`)
	Type HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAssetType `json:"type"`
	JSON hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAssetJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAssetJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAsset]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAssetJSON struct {
	ID          apijson.Field
	Decimals    apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`TOKEN`)
type HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAssetType string

const (
	HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAssetTypeToken HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAssetType = "TOKEN"
)

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAssetType) IsKnown() bool {
	switch r {
	case HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaAssetTypeToken:
		return true
	}
	return false
}

// Details of the incoming transfer
type HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                     `json:"usd_price,nullable"`
	JSON     hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaInJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaInJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaIn]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                      `json:"usd_price,nullable"`
	JSON     hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaOutJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaOutJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaOut]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsHederaAccountTokenTransferDiffDetailsSchemaOutJSON) RawJSON() string {
	return r.raw
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchema struct {
	Asset HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaOut  `json:"out,nullable"`
	JSON hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchema]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchema) implementsHederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiff() {
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAsset struct {
	// The NFT ID
	ID string `json:"id,required"`
	// The NFT's description
	Description string `json:"description,required"`
	// NFT's display name
	Name string `json:"name,required"`
	// The NFT's collection ID (token ID)
	NFTType string `json:"nft_type,required"`
	// URL of the NFT's image
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`NFT`)
	Type HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAssetType `json:"type"`
	JSON hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAssetJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAsset]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAssetJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	NFTType     apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`NFT`)
type HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAssetType string

const (
	HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAssetTypeNFT HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAssetType = "NFT"
)

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAssetType) IsKnown() bool {
	switch r {
	case HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaAssetTypeNFT:
		return true
	}
	return false
}

// Details of the incoming transfer
type HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaIn struct {
	// Token ID of the transfer
	TokenID string `json:"token_id,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                          `json:"usd_price,nullable"`
	JSON     hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaInJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaInJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaIn]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaInJSON struct {
	TokenID     apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaOut struct {
	// Token ID of the transfer
	TokenID string `json:"token_id,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                           `json:"usd_price,nullable"`
	JSON     hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaOutJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaOutJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaOut]
type hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaOutJSON struct {
	TokenID     apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseAssetsDiffsNFTDetailsSchemaErc721DiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseExposure struct {
	Asset HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAsset `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresSpender `json:"spenders"`
	JSON     hederaTransactionScanResponseSimulationHederaSimulationResponseExposureJSON                `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseExposureJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseExposure]
type hederaTransactionScanResponseSimulationHederaSimulationResponseExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseExposureJSON) RawJSON() string {
	return r.raw
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAsset struct {
	// The token ID
	ID string `json:"id"`
	// Decimals of the asset
	Decimals int64 `json:"decimals"`
	// The token description
	Description string `json:"description"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// The token name
	Name string `json:"name"`
	// The NFT's collection ID (token ID)
	NFTType string `json:"nft_type"`
	// The token's symbol
	Symbol string `json:"symbol"`
	// Type of the asset (`NATIVE`)
	Type  HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetType `json:"type"`
	JSON  hederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetJSON `json:"-"`
	union HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetUnion
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAsset]
type hederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetJSON struct {
	ID          apijson.Field
	Decimals    apijson.Field
	Description apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	NFTType     apijson.Field
	Symbol      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetJSON) RawJSON() string {
	return r.raw
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAsset) UnmarshalJSON(data []byte) (err error) {
	*r = HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNativeAssetDetailsSchema],
// [HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaTokenDetailsSchema],
// [HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNFTDetailsSchema].
func (r HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAsset) AsUnion() HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetUnion {
	return r.union
}

// Union satisfied by
// [HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNativeAssetDetailsSchema],
// [HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaTokenDetailsSchema]
// or
// [HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNFTDetailsSchema].
type HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetUnion interface {
	implementsHederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNativeAssetDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaTokenDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNFTDetailsSchema{}),
		},
	)
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNativeAssetDetailsSchema struct {
	// Type of the asset (`NATIVE`)
	Type HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNativeAssetDetailsSchemaType `json:"type"`
	JSON hederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNativeAssetDetailsSchemaJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNativeAssetDetailsSchemaJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNativeAssetDetailsSchema]
type hederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNativeAssetDetailsSchemaJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNativeAssetDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNativeAssetDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNativeAssetDetailsSchema) implementsHederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAsset() {
}

// Type of the asset (`NATIVE`)
type HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNativeAssetDetailsSchemaType string

const (
	HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNativeAssetDetailsSchemaTypeNative HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNativeAssetDetailsSchemaType = "NATIVE"
)

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNativeAssetDetailsSchemaType) IsKnown() bool {
	switch r {
	case HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNativeAssetDetailsSchemaTypeNative:
		return true
	}
	return false
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaTokenDetailsSchema struct {
	// The token ID
	ID string `json:"id,required"`
	// Decimals of the asset
	Decimals int64 `json:"decimals,required"`
	// The token description
	Description string `json:"description,required"`
	// The token name
	Name string `json:"name,required"`
	// The token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`TOKEN`)
	Type HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaTokenDetailsSchemaType `json:"type"`
	JSON hederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaTokenDetailsSchemaJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaTokenDetailsSchemaJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaTokenDetailsSchema]
type hederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaTokenDetailsSchemaJSON struct {
	ID          apijson.Field
	Decimals    apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaTokenDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaTokenDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaTokenDetailsSchema) implementsHederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAsset() {
}

// Type of the asset (`TOKEN`)
type HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaTokenDetailsSchemaType string

const (
	HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaTokenDetailsSchemaTypeToken HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaTokenDetailsSchemaType = "TOKEN"
)

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaTokenDetailsSchemaType) IsKnown() bool {
	switch r {
	case HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaTokenDetailsSchemaTypeToken:
		return true
	}
	return false
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNFTDetailsSchema struct {
	// The NFT ID
	ID string `json:"id,required"`
	// The NFT's description
	Description string `json:"description,required"`
	// NFT's display name
	Name string `json:"name,required"`
	// The NFT's collection ID (token ID)
	NFTType string `json:"nft_type,required"`
	// URL of the NFT's image
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`NFT`)
	Type HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNFTDetailsSchemaType `json:"type"`
	JSON hederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNFTDetailsSchemaJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNFTDetailsSchemaJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNFTDetailsSchema]
type hederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNFTDetailsSchemaJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	NFTType     apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNFTDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNFTDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNFTDetailsSchema) implementsHederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAsset() {
}

// Type of the asset (`NFT`)
type HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNFTDetailsSchemaType string

const (
	HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNFTDetailsSchemaTypeNFT HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNFTDetailsSchemaType = "NFT"
)

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNFTDetailsSchemaType) IsKnown() bool {
	switch r {
	case HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetHederaNFTDetailsSchemaTypeNFT:
		return true
	}
	return false
}

// Type of the asset (`NATIVE`)
type HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetType string

const (
	HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetTypeNative HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetType = "NATIVE"
	HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetTypeToken  HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetType = "TOKEN"
	HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetTypeNFT    HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetType = "NFT"
)

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetType) IsKnown() bool {
	switch r {
	case HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetTypeNative, HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetTypeToken, HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresAssetTypeNFT:
		return true
	}
	return false
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresSpender struct {
	Exposure []HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresSpendersExposure `json:"exposure,required"`
	// Summarized description of the exposure
	Summary string                                                                              `json:"summary,nullable"`
	JSON    hederaTransactionScanResponseSimulationHederaSimulationResponseExposuresSpenderJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseExposuresSpenderJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresSpender]
type hederaTransactionScanResponseSimulationHederaSimulationResponseExposuresSpenderJSON struct {
	Exposure    apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseExposuresSpenderJSON) RawJSON() string {
	return r.raw
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresSpendersExposure struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                      `json:"usd_price,nullable"`
	JSON     hederaTransactionScanResponseSimulationHederaSimulationResponseExposuresSpendersExposureJSON `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationResponseExposuresSpendersExposureJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresSpendersExposure]
type hederaTransactionScanResponseSimulationHederaSimulationResponseExposuresSpendersExposureJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationResponseExposuresSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationResponseExposuresSpendersExposureJSON) RawJSON() string {
	return r.raw
}

type HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActions string

const (
	HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActionsNativeWrap      HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActions = "native_wrap"
	HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActionsNativeTransfer  HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActions = "native_transfer"
	HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActionsTokenTransfer   HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActions = "token_transfer"
	HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActionsSwap            HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActions = "swap"
	HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActionsMint            HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActions = "mint"
	HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActionsStake           HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActions = "stake"
	HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActionsApproval        HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActions = "approval"
	HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActionsProxyUpgrade    HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActions = "proxy_upgrade"
	HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActionsOwnershipChange HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActions = "ownership_change"
	HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActionsSetCodeAccount  HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActions = "set_code_account"
)

func (r HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActions) IsKnown() bool {
	switch r {
	case HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActionsNativeWrap, HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActionsNativeTransfer, HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActionsTokenTransfer, HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActionsSwap, HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActionsMint, HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActionsStake, HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActionsApproval, HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActionsProxyUpgrade, HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActionsOwnershipChange, HederaTransactionScanResponseSimulationHederaSimulationResponseTransactionActionsSetCodeAccount:
		return true
	}
	return false
}

type HederaTransactionScanResponseSimulationHederaSimulationErrorSchema struct {
	// Error message
	Error  string                                                                   `json:"error,required"`
	Status HederaTransactionScanResponseSimulationHederaSimulationErrorSchemaStatus `json:"status,required"`
	JSON   hederaTransactionScanResponseSimulationHederaSimulationErrorSchemaJSON   `json:"-"`
}

// hederaTransactionScanResponseSimulationHederaSimulationErrorSchemaJSON contains
// the JSON metadata for the struct
// [HederaTransactionScanResponseSimulationHederaSimulationErrorSchema]
type hederaTransactionScanResponseSimulationHederaSimulationErrorSchemaJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseSimulationHederaSimulationErrorSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseSimulationHederaSimulationErrorSchemaJSON) RawJSON() string {
	return r.raw
}

func (r HederaTransactionScanResponseSimulationHederaSimulationErrorSchema) implementsHederaTransactionScanResponseSimulation() {
}

type HederaTransactionScanResponseSimulationHederaSimulationErrorSchemaStatus string

const (
	HederaTransactionScanResponseSimulationHederaSimulationErrorSchemaStatusError HederaTransactionScanResponseSimulationHederaSimulationErrorSchemaStatus = "Error"
)

func (r HederaTransactionScanResponseSimulationHederaSimulationErrorSchemaStatus) IsKnown() bool {
	switch r {
	case HederaTransactionScanResponseSimulationHederaSimulationErrorSchemaStatusError:
		return true
	}
	return false
}

type HederaTransactionScanResponseSimulationStatus string

const (
	HederaTransactionScanResponseSimulationStatusSuccess HederaTransactionScanResponseSimulationStatus = "Success"
	HederaTransactionScanResponseSimulationStatusError   HederaTransactionScanResponseSimulationStatus = "Error"
)

func (r HederaTransactionScanResponseSimulationStatus) IsKnown() bool {
	switch r {
	case HederaTransactionScanResponseSimulationStatusSuccess, HederaTransactionScanResponseSimulationStatusError:
		return true
	}
	return false
}

// Validation result; Only present if validation option is included in the request
type HederaTransactionScanResponseValidation struct {
	Status HederaTransactionScanResponseValidationStatus `json:"status,required"`
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification"`
	// A textual description about the validation result
	Description string `json:"description"`
	// Error message
	Error string `json:"error"`
	// This field can have the runtime type of
	// [[]HederaTransactionScanResponseValidationHederaValidationResultFeature].
	Features interface{} `json:"features"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason string `json:"reason"`
	// Verdict of the validation
	ResultType HederaTransactionScanResponseValidationResultType `json:"result_type"`
	JSON       hederaTransactionScanResponseValidationJSON       `json:"-"`
	union      HederaTransactionScanResponseValidationUnion
}

// hederaTransactionScanResponseValidationJSON contains the JSON metadata for the
// struct [HederaTransactionScanResponseValidation]
type hederaTransactionScanResponseValidationJSON struct {
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

func (r hederaTransactionScanResponseValidationJSON) RawJSON() string {
	return r.raw
}

func (r *HederaTransactionScanResponseValidation) UnmarshalJSON(data []byte) (err error) {
	*r = HederaTransactionScanResponseValidation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [HederaTransactionScanResponseValidationUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [HederaTransactionScanResponseValidationHederaValidationResult],
// [HederaTransactionScanResponseValidationHederaValidationErrorSchema].
func (r HederaTransactionScanResponseValidation) AsUnion() HederaTransactionScanResponseValidationUnion {
	return r.union
}

// Validation result; Only present if validation option is included in the request
//
// Union satisfied by
// [HederaTransactionScanResponseValidationHederaValidationResult] or
// [HederaTransactionScanResponseValidationHederaValidationErrorSchema].
type HederaTransactionScanResponseValidationUnion interface {
	implementsHederaTransactionScanResponseValidation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*HederaTransactionScanResponseValidationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HederaTransactionScanResponseValidationHederaValidationResult{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HederaTransactionScanResponseValidationHederaValidationErrorSchema{}),
		},
	)
}

type HederaTransactionScanResponseValidationHederaValidationResult struct {
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification,required"`
	// A textual description about the validation result
	Description string                                                                 `json:"description,required"`
	Features    []HederaTransactionScanResponseValidationHederaValidationResultFeature `json:"features,required"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason string `json:"reason,required"`
	// Verdict of the validation
	ResultType HederaTransactionScanResponseValidationHederaValidationResultResultType `json:"result_type,required"`
	Status     HederaTransactionScanResponseValidationHederaValidationResultStatus     `json:"status,required"`
	JSON       hederaTransactionScanResponseValidationHederaValidationResultJSON       `json:"-"`
}

// hederaTransactionScanResponseValidationHederaValidationResultJSON contains the
// JSON metadata for the struct
// [HederaTransactionScanResponseValidationHederaValidationResult]
type hederaTransactionScanResponseValidationHederaValidationResultJSON struct {
	Classification apijson.Field
	Description    apijson.Field
	Features       apijson.Field
	Reason         apijson.Field
	ResultType     apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HederaTransactionScanResponseValidationHederaValidationResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseValidationHederaValidationResultJSON) RawJSON() string {
	return r.raw
}

func (r HederaTransactionScanResponseValidationHederaValidationResult) implementsHederaTransactionScanResponseValidation() {
}

type HederaTransactionScanResponseValidationHederaValidationResultFeature struct {
	// Address the feature relates to.
	Address string `json:"address,required,nullable"`
	// Textual description of the feature.
	Description string `json:"description,required"`
	// The ID of the feature associated with this transaction.
	FeatureID string `json:"feature_id,required"`
	// The security classification of the feature (Benign, Warning, Malicious or Info).
	Type HederaTransactionScanResponseValidationHederaValidationResultFeaturesType `json:"type,required"`
	JSON hederaTransactionScanResponseValidationHederaValidationResultFeatureJSON  `json:"-"`
}

// hederaTransactionScanResponseValidationHederaValidationResultFeatureJSON
// contains the JSON metadata for the struct
// [HederaTransactionScanResponseValidationHederaValidationResultFeature]
type hederaTransactionScanResponseValidationHederaValidationResultFeatureJSON struct {
	Address     apijson.Field
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseValidationHederaValidationResultFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseValidationHederaValidationResultFeatureJSON) RawJSON() string {
	return r.raw
}

// The security classification of the feature (Benign, Warning, Malicious or Info).
type HederaTransactionScanResponseValidationHederaValidationResultFeaturesType string

const (
	HederaTransactionScanResponseValidationHederaValidationResultFeaturesTypeBenign    HederaTransactionScanResponseValidationHederaValidationResultFeaturesType = "Benign"
	HederaTransactionScanResponseValidationHederaValidationResultFeaturesTypeWarning   HederaTransactionScanResponseValidationHederaValidationResultFeaturesType = "Warning"
	HederaTransactionScanResponseValidationHederaValidationResultFeaturesTypeMalicious HederaTransactionScanResponseValidationHederaValidationResultFeaturesType = "Malicious"
	HederaTransactionScanResponseValidationHederaValidationResultFeaturesTypeInfo      HederaTransactionScanResponseValidationHederaValidationResultFeaturesType = "Info"
)

func (r HederaTransactionScanResponseValidationHederaValidationResultFeaturesType) IsKnown() bool {
	switch r {
	case HederaTransactionScanResponseValidationHederaValidationResultFeaturesTypeBenign, HederaTransactionScanResponseValidationHederaValidationResultFeaturesTypeWarning, HederaTransactionScanResponseValidationHederaValidationResultFeaturesTypeMalicious, HederaTransactionScanResponseValidationHederaValidationResultFeaturesTypeInfo:
		return true
	}
	return false
}

// Verdict of the validation
type HederaTransactionScanResponseValidationHederaValidationResultResultType string

const (
	HederaTransactionScanResponseValidationHederaValidationResultResultTypeBenign    HederaTransactionScanResponseValidationHederaValidationResultResultType = "Benign"
	HederaTransactionScanResponseValidationHederaValidationResultResultTypeWarning   HederaTransactionScanResponseValidationHederaValidationResultResultType = "Warning"
	HederaTransactionScanResponseValidationHederaValidationResultResultTypeMalicious HederaTransactionScanResponseValidationHederaValidationResultResultType = "Malicious"
	HederaTransactionScanResponseValidationHederaValidationResultResultTypeError     HederaTransactionScanResponseValidationHederaValidationResultResultType = "Error"
)

func (r HederaTransactionScanResponseValidationHederaValidationResultResultType) IsKnown() bool {
	switch r {
	case HederaTransactionScanResponseValidationHederaValidationResultResultTypeBenign, HederaTransactionScanResponseValidationHederaValidationResultResultTypeWarning, HederaTransactionScanResponseValidationHederaValidationResultResultTypeMalicious, HederaTransactionScanResponseValidationHederaValidationResultResultTypeError:
		return true
	}
	return false
}

type HederaTransactionScanResponseValidationHederaValidationResultStatus string

const (
	HederaTransactionScanResponseValidationHederaValidationResultStatusSuccess HederaTransactionScanResponseValidationHederaValidationResultStatus = "Success"
)

func (r HederaTransactionScanResponseValidationHederaValidationResultStatus) IsKnown() bool {
	switch r {
	case HederaTransactionScanResponseValidationHederaValidationResultStatusSuccess:
		return true
	}
	return false
}

type HederaTransactionScanResponseValidationHederaValidationErrorSchema struct {
	// Error message
	Error  string                                                                   `json:"error,required"`
	Status HederaTransactionScanResponseValidationHederaValidationErrorSchemaStatus `json:"status,required"`
	JSON   hederaTransactionScanResponseValidationHederaValidationErrorSchemaJSON   `json:"-"`
}

// hederaTransactionScanResponseValidationHederaValidationErrorSchemaJSON contains
// the JSON metadata for the struct
// [HederaTransactionScanResponseValidationHederaValidationErrorSchema]
type hederaTransactionScanResponseValidationHederaValidationErrorSchemaJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaTransactionScanResponseValidationHederaValidationErrorSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaTransactionScanResponseValidationHederaValidationErrorSchemaJSON) RawJSON() string {
	return r.raw
}

func (r HederaTransactionScanResponseValidationHederaValidationErrorSchema) implementsHederaTransactionScanResponseValidation() {
}

type HederaTransactionScanResponseValidationHederaValidationErrorSchemaStatus string

const (
	HederaTransactionScanResponseValidationHederaValidationErrorSchemaStatusError HederaTransactionScanResponseValidationHederaValidationErrorSchemaStatus = "Error"
)

func (r HederaTransactionScanResponseValidationHederaValidationErrorSchemaStatus) IsKnown() bool {
	switch r {
	case HederaTransactionScanResponseValidationHederaValidationErrorSchemaStatusError:
		return true
	}
	return false
}

type HederaTransactionScanResponseValidationStatus string

const (
	HederaTransactionScanResponseValidationStatusSuccess HederaTransactionScanResponseValidationStatus = "Success"
	HederaTransactionScanResponseValidationStatusError   HederaTransactionScanResponseValidationStatus = "Error"
)

func (r HederaTransactionScanResponseValidationStatus) IsKnown() bool {
	switch r {
	case HederaTransactionScanResponseValidationStatusSuccess, HederaTransactionScanResponseValidationStatusError:
		return true
	}
	return false
}

// Verdict of the validation
type HederaTransactionScanResponseValidationResultType string

const (
	HederaTransactionScanResponseValidationResultTypeBenign    HederaTransactionScanResponseValidationResultType = "Benign"
	HederaTransactionScanResponseValidationResultTypeWarning   HederaTransactionScanResponseValidationResultType = "Warning"
	HederaTransactionScanResponseValidationResultTypeMalicious HederaTransactionScanResponseValidationResultType = "Malicious"
	HederaTransactionScanResponseValidationResultTypeError     HederaTransactionScanResponseValidationResultType = "Error"
)

func (r HederaTransactionScanResponseValidationResultType) IsKnown() bool {
	switch r {
	case HederaTransactionScanResponseValidationResultTypeBenign, HederaTransactionScanResponseValidationResultTypeWarning, HederaTransactionScanResponseValidationResultTypeMalicious, HederaTransactionScanResponseValidationResultTypeError:
		return true
	}
	return false
}

type HederaTransactionScanParams struct {
	// The address to relate the transaction to. Account address determines in which
	// perspective the transaction is simulated and validated.
	AccountAddress param.Field[interface{}] `json:"account_address,required"`
	// The chain the transaction runs on.
	Chain param.Field[HederaTransactionScanParamsChain] `json:"chain,required"`
	// Additional information regarding the wallet involved in the transaction.
	Metadata    param.Field[HederaTransactionScanParamsMetadataUnion] `json:"metadata,required"`
	Transaction param.Field[string]                                   `json:"transaction,required"`
	// Select which component will be included in the response.
	//
	//   - `simulation` - Include the results of the transaction simulation in your
	//     response.
	//   - `validation` - Include a security validation of the transaction in your
	//     response.
	Options param.Field[[]HederaTransactionScanParamsOption] `json:"options"`
}

func (r HederaTransactionScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The chain the transaction runs on.
type HederaTransactionScanParamsChain string

const (
	HederaTransactionScanParamsChainMainnet HederaTransactionScanParamsChain = "mainnet"
)

func (r HederaTransactionScanParamsChain) IsKnown() bool {
	switch r {
	case HederaTransactionScanParamsChainMainnet:
		return true
	}
	return false
}

// Additional information regarding the wallet involved in the transaction.
type HederaTransactionScanParamsMetadata struct {
	// Metadata for wallet requests
	Type param.Field[HederaTransactionScanParamsMetadataType] `json:"type"`
	// URL of the dApp the transaction originated from.
	URL param.Field[string] `json:"url"`
}

func (r HederaTransactionScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HederaTransactionScanParamsMetadata) implementsHederaTransactionScanParamsMetadataUnion() {}

// Additional information regarding the wallet involved in the transaction.
//
// Satisfied by [HederaTransactionScanParamsMetadataHederaWalletRequestMetadata],
// [HederaTransactionScanParamsMetadataHederaInAppRequestMetadata],
// [HederaTransactionScanParamsMetadata].
type HederaTransactionScanParamsMetadataUnion interface {
	implementsHederaTransactionScanParamsMetadataUnion()
}

type HederaTransactionScanParamsMetadataHederaWalletRequestMetadata struct {
	// Metadata for wallet requests
	Type param.Field[HederaTransactionScanParamsMetadataHederaWalletRequestMetadataType] `json:"type,required"`
	// URL of the dApp the transaction originated from.
	URL param.Field[string] `json:"url,required"`
}

func (r HederaTransactionScanParamsMetadataHederaWalletRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HederaTransactionScanParamsMetadataHederaWalletRequestMetadata) implementsHederaTransactionScanParamsMetadataUnion() {
}

// Metadata for wallet requests
type HederaTransactionScanParamsMetadataHederaWalletRequestMetadataType string

const (
	HederaTransactionScanParamsMetadataHederaWalletRequestMetadataTypeWallet HederaTransactionScanParamsMetadataHederaWalletRequestMetadataType = "wallet"
)

func (r HederaTransactionScanParamsMetadataHederaWalletRequestMetadataType) IsKnown() bool {
	switch r {
	case HederaTransactionScanParamsMetadataHederaWalletRequestMetadataTypeWallet:
		return true
	}
	return false
}

type HederaTransactionScanParamsMetadataHederaInAppRequestMetadata struct {
	// Metadata for in-app requests
	Type param.Field[HederaTransactionScanParamsMetadataHederaInAppRequestMetadataType] `json:"type"`
}

func (r HederaTransactionScanParamsMetadataHederaInAppRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HederaTransactionScanParamsMetadataHederaInAppRequestMetadata) implementsHederaTransactionScanParamsMetadataUnion() {
}

// Metadata for in-app requests
type HederaTransactionScanParamsMetadataHederaInAppRequestMetadataType string

const (
	HederaTransactionScanParamsMetadataHederaInAppRequestMetadataTypeInApp HederaTransactionScanParamsMetadataHederaInAppRequestMetadataType = "in_app"
)

func (r HederaTransactionScanParamsMetadataHederaInAppRequestMetadataType) IsKnown() bool {
	switch r {
	case HederaTransactionScanParamsMetadataHederaInAppRequestMetadataTypeInApp:
		return true
	}
	return false
}

// Metadata for wallet requests
type HederaTransactionScanParamsMetadataType string

const (
	HederaTransactionScanParamsMetadataTypeWallet HederaTransactionScanParamsMetadataType = "wallet"
	HederaTransactionScanParamsMetadataTypeInApp  HederaTransactionScanParamsMetadataType = "in_app"
)

func (r HederaTransactionScanParamsMetadataType) IsKnown() bool {
	switch r {
	case HederaTransactionScanParamsMetadataTypeWallet, HederaTransactionScanParamsMetadataTypeInApp:
		return true
	}
	return false
}

type HederaTransactionScanParamsOption string

const (
	HederaTransactionScanParamsOptionValidation HederaTransactionScanParamsOption = "validation"
	HederaTransactionScanParamsOptionSimulation HederaTransactionScanParamsOption = "simulation"
)

func (r HederaTransactionScanParamsOption) IsKnown() bool {
	switch r {
	case HederaTransactionScanParamsOptionValidation, HederaTransactionScanParamsOptionSimulation:
		return true
	}
	return false
}
