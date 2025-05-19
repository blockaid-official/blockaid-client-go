// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"reflect"
	"time"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/option"
	"github.com/tidwall/gjson"
)

// SuiService contains methods and other services that help with interacting with
// the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSuiService] method instead.
type SuiService struct {
	Options         []option.RequestOption
	Transaction     *SuiTransactionService
	PostTransaction *SuiPostTransactionService
}

// NewSuiService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSuiService(opts ...option.RequestOption) (r *SuiService) {
	r = &SuiService{}
	r.Options = opts
	r.Transaction = NewSuiTransactionService(opts...)
	r.PostTransaction = NewSuiPostTransactionService(opts...)
	return
}

type SuiAssetTransferDetailsSchema struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                           `json:"usd_price,nullable"`
	JSON     suiAssetTransferDetailsSchemaJSON `json:"-"`
}

// suiAssetTransferDetailsSchemaJSON contains the JSON metadata for the struct
// [SuiAssetTransferDetailsSchema]
type suiAssetTransferDetailsSchemaJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiAssetTransferDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiAssetTransferDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

type SuiNativeAssetDetailsSchema struct {
	// Decimals of the asset
	Decimals SuiNativeAssetDetailsSchemaDecimals `json:"decimals"`
	// URL of the asset's logo
	LogoURL SuiNativeAssetDetailsSchemaLogoURL `json:"logo_url"`
	// Name of the asset
	Name SuiNativeAssetDetailsSchemaName `json:"name"`
	// Symbol of the asset
	Symbol SuiNativeAssetDetailsSchemaSymbol `json:"symbol"`
	// Type of the asset (`NATIVE`)
	Type SuiNativeAssetDetailsSchemaType `json:"type"`
	JSON suiNativeAssetDetailsSchemaJSON `json:"-"`
}

// suiNativeAssetDetailsSchemaJSON contains the JSON metadata for the struct
// [SuiNativeAssetDetailsSchema]
type suiNativeAssetDetailsSchemaJSON struct {
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiNativeAssetDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiNativeAssetDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

// Decimals of the asset
type SuiNativeAssetDetailsSchemaDecimals int64

const (
	SuiNativeAssetDetailsSchemaDecimals9 SuiNativeAssetDetailsSchemaDecimals = 9
)

func (r SuiNativeAssetDetailsSchemaDecimals) IsKnown() bool {
	switch r {
	case SuiNativeAssetDetailsSchemaDecimals9:
		return true
	}
	return false
}

// URL of the asset's logo
type SuiNativeAssetDetailsSchemaLogoURL string

const (
	SuiNativeAssetDetailsSchemaLogoURLHTTPSImagedeliveryNetCBndGgkrsEaBIxIp9SkQSuiSvgPublic SuiNativeAssetDetailsSchemaLogoURL = "https://imagedelivery.net/cBNDGgkrsEA-b_ixIp9SkQ/sui.svg/public"
)

func (r SuiNativeAssetDetailsSchemaLogoURL) IsKnown() bool {
	switch r {
	case SuiNativeAssetDetailsSchemaLogoURLHTTPSImagedeliveryNetCBndGgkrsEaBIxIp9SkQSuiSvgPublic:
		return true
	}
	return false
}

// Name of the asset
type SuiNativeAssetDetailsSchemaName string

const (
	SuiNativeAssetDetailsSchemaNameSui SuiNativeAssetDetailsSchemaName = "Sui"
)

func (r SuiNativeAssetDetailsSchemaName) IsKnown() bool {
	switch r {
	case SuiNativeAssetDetailsSchemaNameSui:
		return true
	}
	return false
}

// Symbol of the asset
type SuiNativeAssetDetailsSchemaSymbol string

const (
	SuiNativeAssetDetailsSchemaSymbolSui SuiNativeAssetDetailsSchemaSymbol = "SUI"
)

func (r SuiNativeAssetDetailsSchemaSymbol) IsKnown() bool {
	switch r {
	case SuiNativeAssetDetailsSchemaSymbolSui:
		return true
	}
	return false
}

// Type of the asset (`NATIVE`)
type SuiNativeAssetDetailsSchemaType string

const (
	SuiNativeAssetDetailsSchemaTypeNative SuiNativeAssetDetailsSchemaType = "NATIVE"
)

func (r SuiNativeAssetDetailsSchemaType) IsKnown() bool {
	switch r {
	case SuiNativeAssetDetailsSchemaTypeNative:
		return true
	}
	return false
}

type SuiNFTDetailsSchema struct {
	// The NFT ID
	ID string `json:"id,required"`
	// The NFT's description
	Description string `json:"description,required"`
	// NFT's display name
	Name string `json:"name,required"`
	// The NFT's collection ID
	NFTType string `json:"nft_type,required"`
	// Type of the asset (`NFT`)
	Type SuiNFTDetailsSchemaType `json:"type"`
	// URL of the nft's image
	URL  string                  `json:"url,nullable"`
	JSON suiNFTDetailsSchemaJSON `json:"-"`
}

// suiNFTDetailsSchemaJSON contains the JSON metadata for the struct
// [SuiNFTDetailsSchema]
type suiNFTDetailsSchemaJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	NFTType     apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiNFTDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiNFTDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`NFT`)
type SuiNFTDetailsSchemaType string

const (
	SuiNFTDetailsSchemaTypeNFT SuiNFTDetailsSchemaType = "NFT"
)

func (r SuiNFTDetailsSchemaType) IsKnown() bool {
	switch r {
	case SuiNFTDetailsSchemaTypeNFT:
		return true
	}
	return false
}

type SuiNFTDiffSchema struct {
	// NFT ID of the transfer
	ID string `json:"id,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64              `json:"usd_price,nullable"`
	JSON     suiNFTDiffSchemaJSON `json:"-"`
}

// suiNFTDiffSchemaJSON contains the JSON metadata for the struct
// [SuiNFTDiffSchema]
type suiNFTDiffSchemaJSON struct {
	ID          apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiNFTDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiNFTDiffSchemaJSON) RawJSON() string {
	return r.raw
}

type SuiTransactionScanResponse struct {
	// Simulation result; Only present if simulation option is included in the request
	Simulation SuiTransactionScanResponseSimulation `json:"simulation,nullable"`
	// Validation result; Only present if validation option is included in the request
	Validation SuiTransactionScanResponseValidation `json:"validation,nullable"`
	JSON       suiTransactionScanResponseJSON       `json:"-"`
}

// suiTransactionScanResponseJSON contains the JSON metadata for the struct
// [SuiTransactionScanResponse]
type suiTransactionScanResponseJSON struct {
	Simulation  apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiTransactionScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiTransactionScanResponseJSON) RawJSON() string {
	return r.raw
}

// Simulation result; Only present if simulation option is included in the request
type SuiTransactionScanResponseSimulation struct {
	Status SuiTransactionScanResponseSimulationStatus `json:"status,required"`
	// This field can have the runtime type of
	// [SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummary].
	AccountSummary interface{} `json:"account_summary"`
	// This field can have the runtime type of [[]interface{}].
	AddressDetails interface{} `json:"address_details"`
	// This field can have the runtime type of
	// [map[string][]SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiff].
	AssetsDiffs interface{} `json:"assets_diffs"`
	// Error message
	Error string                                   `json:"error"`
	JSON  suiTransactionScanResponseSimulationJSON `json:"-"`
	union SuiTransactionScanResponseSimulationUnion
}

// suiTransactionScanResponseSimulationJSON contains the JSON metadata for the
// struct [SuiTransactionScanResponseSimulation]
type suiTransactionScanResponseSimulationJSON struct {
	Status         apijson.Field
	AccountSummary apijson.Field
	AddressDetails apijson.Field
	AssetsDiffs    apijson.Field
	Error          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r suiTransactionScanResponseSimulationJSON) RawJSON() string {
	return r.raw
}

func (r *SuiTransactionScanResponseSimulation) UnmarshalJSON(data []byte) (err error) {
	*r = SuiTransactionScanResponseSimulation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SuiTransactionScanResponseSimulationUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SuiTransactionScanResponseSimulationSuiSimulationResult],
// [SuiTransactionScanResponseSimulationSuiSimulationErrorSchema].
func (r SuiTransactionScanResponseSimulation) AsUnion() SuiTransactionScanResponseSimulationUnion {
	return r.union
}

// Simulation result; Only present if simulation option is included in the request
//
// Union satisfied by [SuiTransactionScanResponseSimulationSuiSimulationResult] or
// [SuiTransactionScanResponseSimulationSuiSimulationErrorSchema].
type SuiTransactionScanResponseSimulationUnion interface {
	implementsSuiTransactionScanResponseSimulation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SuiTransactionScanResponseSimulationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiTransactionScanResponseSimulationSuiSimulationResult{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiTransactionScanResponseSimulationSuiSimulationErrorSchema{}),
		},
	)
}

type SuiTransactionScanResponseSimulationSuiSimulationResult struct {
	// Summary of the actions and asset transfers that were made by the requested
	// account address
	AccountSummary SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummary `json:"account_summary,required"`
	Status         SuiTransactionScanResponseSimulationSuiSimulationResultStatus         `json:"status,required"`
	// Details of addresses involved in the transaction
	AddressDetails []interface{} `json:"address_details"`
	// Mapping between the address of an account to the assets diff during the
	// transaction
	AssetsDiffs map[string][]SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiff `json:"assets_diffs"`
	JSON        suiTransactionScanResponseSimulationSuiSimulationResultJSON                    `json:"-"`
}

// suiTransactionScanResponseSimulationSuiSimulationResultJSON contains the JSON
// metadata for the struct
// [SuiTransactionScanResponseSimulationSuiSimulationResult]
type suiTransactionScanResponseSimulationSuiSimulationResultJSON struct {
	AccountSummary apijson.Field
	Status         apijson.Field
	AddressDetails apijson.Field
	AssetsDiffs    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SuiTransactionScanResponseSimulationSuiSimulationResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiTransactionScanResponseSimulationSuiSimulationResultJSON) RawJSON() string {
	return r.raw
}

func (r SuiTransactionScanResponseSimulationSuiSimulationResult) implementsSuiTransactionScanResponseSimulation() {
}

// Summary of the actions and asset transfers that were made by the requested
// account address
type SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummary struct {
	// Total USD diff for the requested account address
	TotalUsdDiff SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryTotalUsdDiff `json:"total_usd_diff,required"`
	// Assets diffs of the requested account address
	AccountAssetsDiffs []SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiff `json:"account_assets_diffs"`
	JSON               suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryJSON                `json:"-"`
}

// suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryJSON
// contains the JSON metadata for the struct
// [SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummary]
type suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryJSON struct {
	TotalUsdDiff       apijson.Field
	AccountAssetsDiffs apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryJSON) RawJSON() string {
	return r.raw
}

// Total USD diff for the requested account address
type SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryTotalUsdDiff struct {
	// Total incoming USD transfers
	In float64 `json:"in,required"`
	// Total outgoing USD transfers
	Out float64 `json:"out,required"`
	// Total USD transfers
	Total float64                                                                               `json:"total"`
	JSON  suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryTotalUsdDiffJSON `json:"-"`
}

// suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryTotalUsdDiffJSON
// contains the JSON metadata for the struct
// [SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryTotalUsdDiff]
type suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

type SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiff struct {
	// This field can have the runtime type of [SuiNativeAssetDetailsSchema],
	// [SuiNFTDetailsSchema],
	// [SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAsset].
	Asset interface{} `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of [SuiAssetTransferDetailsSchema],
	// [SuiNFTDiffSchema].
	In interface{} `json:"in"`
	// This field can have the runtime type of [SuiAssetTransferDetailsSchema],
	// [SuiNFTDiffSchema].
	Out   interface{}                                                                                `json:"out"`
	JSON  suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffJSON `json:"-"`
	union SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsUnion
}

// suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffJSON
// contains the JSON metadata for the struct
// [SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiff]
type suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiff],
// [SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiff],
// [SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiff].
func (r SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiff) AsUnion() SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiff],
// [SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiff]
// or
// [SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiff].
type SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsUnion interface {
	implementsSuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiff{}),
		},
	)
}

type SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiff struct {
	Asset SuiNativeAssetDetailsSchema `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SuiAssetTransferDetailsSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SuiAssetTransferDetailsSchema                                                                                 `json:"out,nullable"`
	JSON suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiffJSON `json:"-"`
}

// suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiffJSON
// contains the JSON metadata for the struct
// [SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiff]
type suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiff) implementsSuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiff() {
}

type SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiff struct {
	Asset SuiNFTDetailsSchema `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SuiNFTDiffSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SuiNFTDiffSchema                                                                                           `json:"out,nullable"`
	JSON suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiffJSON `json:"-"`
}

// suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiffJSON
// contains the JSON metadata for the struct
// [SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiff]
type suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiff) implementsSuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiff() {
}

type SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiff struct {
	Asset SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SuiAssetTransferDetailsSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SuiAssetTransferDetailsSchema                                                                                `json:"out,nullable"`
	JSON suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffJSON `json:"-"`
}

// suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffJSON
// contains the JSON metadata for the struct
// [SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiff]
type suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiff) implementsSuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiff() {
}

type SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAsset struct {
	// Token's package address
	Address string `json:"address,required"`
	// Token's decimal precision
	Decimals int64 `json:"decimals,required"`
	// Token's name
	Name string `json:"name,required"`
	// Token's symbol (abbreviation)
	Symbol            string    `json:"symbol,required"`
	CreationTimestamp time.Time `json:"creation_timestamp,nullable" format:"date-time"`
	// URL of the token's logo
	LogoURL string `json:"logo_url,nullable"`
	Scam    bool   `json:"scam,nullable"`
	// Type of the asset (`Coin`)
	Type     SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetType `json:"type"`
	Verified bool                                                                                                              `json:"verified,nullable"`
	JSON     suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetJSON `json:"-"`
}

// suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAsset]
type suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetJSON struct {
	Address           apijson.Field
	Decimals          apijson.Field
	Name              apijson.Field
	Symbol            apijson.Field
	CreationTimestamp apijson.Field
	LogoURL           apijson.Field
	Scam              apijson.Field
	Type              apijson.Field
	Verified          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`Coin`)
type SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetType string

const (
	SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetTypeFungible SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetType = "fungible"
)

func (r SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetType) IsKnown() bool {
	switch r {
	case SuiTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetTypeFungible:
		return true
	}
	return false
}

type SuiTransactionScanResponseSimulationSuiSimulationResultStatus string

const (
	SuiTransactionScanResponseSimulationSuiSimulationResultStatusSuccess SuiTransactionScanResponseSimulationSuiSimulationResultStatus = "Success"
)

func (r SuiTransactionScanResponseSimulationSuiSimulationResultStatus) IsKnown() bool {
	switch r {
	case SuiTransactionScanResponseSimulationSuiSimulationResultStatusSuccess:
		return true
	}
	return false
}

type SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiff struct {
	// This field can have the runtime type of [SuiNativeAssetDetailsSchema],
	// [SuiNFTDetailsSchema],
	// [SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAsset].
	Asset interface{} `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of [SuiAssetTransferDetailsSchema],
	// [SuiNFTDiffSchema].
	In interface{} `json:"in"`
	// This field can have the runtime type of [SuiAssetTransferDetailsSchema],
	// [SuiNFTDiffSchema].
	Out   interface{}                                                           `json:"out"`
	JSON  suiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffJSON `json:"-"`
	union SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsUnion
}

// suiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffJSON contains
// the JSON metadata for the struct
// [SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiff]
type suiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r suiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNativeAssetDiff],
// [SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNFTAssetDiff],
// [SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiff].
func (r SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiff) AsUnion() SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNativeAssetDiff],
// [SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNFTAssetDiff]
// or
// [SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiff].
type SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsUnion interface {
	implementsSuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNativeAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNFTAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiff{}),
		},
	)
}

type SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNativeAssetDiff struct {
	Asset SuiNativeAssetDetailsSchema `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SuiAssetTransferDetailsSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SuiAssetTransferDetailsSchema                                                            `json:"out,nullable"`
	JSON suiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNativeAssetDiffJSON `json:"-"`
}

// suiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNativeAssetDiffJSON
// contains the JSON metadata for the struct
// [SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNativeAssetDiff]
type suiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNativeAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNativeAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNativeAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNativeAssetDiff) implementsSuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiff() {
}

type SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNFTAssetDiff struct {
	Asset SuiNFTDetailsSchema `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SuiNFTDiffSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SuiNFTDiffSchema                                                                      `json:"out,nullable"`
	JSON suiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNFTAssetDiffJSON `json:"-"`
}

// suiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNFTAssetDiffJSON
// contains the JSON metadata for the struct
// [SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNFTAssetDiff]
type suiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNFTAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNFTAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNFTAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNFTAssetDiff) implementsSuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiff() {
}

type SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiff struct {
	Asset SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SuiAssetTransferDetailsSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SuiAssetTransferDetailsSchema                                                           `json:"out,nullable"`
	JSON suiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffJSON `json:"-"`
}

// suiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffJSON
// contains the JSON metadata for the struct
// [SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiff]
type suiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiff) implementsSuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiff() {
}

type SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAsset struct {
	// Token's package address
	Address string `json:"address,required"`
	// Token's decimal precision
	Decimals int64 `json:"decimals,required"`
	// Token's name
	Name string `json:"name,required"`
	// Token's symbol (abbreviation)
	Symbol            string    `json:"symbol,required"`
	CreationTimestamp time.Time `json:"creation_timestamp,nullable" format:"date-time"`
	// URL of the token's logo
	LogoURL string `json:"logo_url,nullable"`
	Scam    bool   `json:"scam,nullable"`
	// Type of the asset (`Coin`)
	Type     SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetType `json:"type"`
	Verified bool                                                                                         `json:"verified,nullable"`
	JSON     suiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetJSON `json:"-"`
}

// suiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAsset]
type suiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetJSON struct {
	Address           apijson.Field
	Decimals          apijson.Field
	Name              apijson.Field
	Symbol            apijson.Field
	CreationTimestamp apijson.Field
	LogoURL           apijson.Field
	Scam              apijson.Field
	Type              apijson.Field
	Verified          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`Coin`)
type SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetType string

const (
	SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetTypeFungible SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetType = "fungible"
)

func (r SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetType) IsKnown() bool {
	switch r {
	case SuiTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetTypeFungible:
		return true
	}
	return false
}

type SuiTransactionScanResponseSimulationSuiSimulationErrorSchema struct {
	// Error message
	Error  string                                                             `json:"error,required"`
	Status SuiTransactionScanResponseSimulationSuiSimulationErrorSchemaStatus `json:"status,required"`
	JSON   suiTransactionScanResponseSimulationSuiSimulationErrorSchemaJSON   `json:"-"`
}

// suiTransactionScanResponseSimulationSuiSimulationErrorSchemaJSON contains the
// JSON metadata for the struct
// [SuiTransactionScanResponseSimulationSuiSimulationErrorSchema]
type suiTransactionScanResponseSimulationSuiSimulationErrorSchemaJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiTransactionScanResponseSimulationSuiSimulationErrorSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiTransactionScanResponseSimulationSuiSimulationErrorSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SuiTransactionScanResponseSimulationSuiSimulationErrorSchema) implementsSuiTransactionScanResponseSimulation() {
}

type SuiTransactionScanResponseSimulationSuiSimulationErrorSchemaStatus string

const (
	SuiTransactionScanResponseSimulationSuiSimulationErrorSchemaStatusError SuiTransactionScanResponseSimulationSuiSimulationErrorSchemaStatus = "Error"
)

func (r SuiTransactionScanResponseSimulationSuiSimulationErrorSchemaStatus) IsKnown() bool {
	switch r {
	case SuiTransactionScanResponseSimulationSuiSimulationErrorSchemaStatusError:
		return true
	}
	return false
}

type SuiTransactionScanResponseSimulationStatus string

const (
	SuiTransactionScanResponseSimulationStatusSuccess SuiTransactionScanResponseSimulationStatus = "Success"
	SuiTransactionScanResponseSimulationStatusError   SuiTransactionScanResponseSimulationStatus = "Error"
)

func (r SuiTransactionScanResponseSimulationStatus) IsKnown() bool {
	switch r {
	case SuiTransactionScanResponseSimulationStatusSuccess, SuiTransactionScanResponseSimulationStatusError:
		return true
	}
	return false
}

// Validation result; Only present if validation option is included in the request
type SuiTransactionScanResponseValidation struct {
	Status SuiTransactionScanResponseValidationStatus `json:"status,required"`
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification"`
	// A textual description about the validation result
	Description string `json:"description"`
	// Error message
	Error string `json:"error"`
	// This field can have the runtime type of
	// [[]SuiTransactionScanResponseValidationSuiValidationResultFeature].
	Features interface{} `json:"features"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason string `json:"reason"`
	// Verdict of the validation
	ResultType SuiTransactionScanResponseValidationResultType `json:"result_type"`
	JSON       suiTransactionScanResponseValidationJSON       `json:"-"`
	union      SuiTransactionScanResponseValidationUnion
}

// suiTransactionScanResponseValidationJSON contains the JSON metadata for the
// struct [SuiTransactionScanResponseValidation]
type suiTransactionScanResponseValidationJSON struct {
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

func (r suiTransactionScanResponseValidationJSON) RawJSON() string {
	return r.raw
}

func (r *SuiTransactionScanResponseValidation) UnmarshalJSON(data []byte) (err error) {
	*r = SuiTransactionScanResponseValidation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SuiTransactionScanResponseValidationUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SuiTransactionScanResponseValidationSuiValidationResult],
// [SuiTransactionScanResponseValidationSuiValidationErrorSchema].
func (r SuiTransactionScanResponseValidation) AsUnion() SuiTransactionScanResponseValidationUnion {
	return r.union
}

// Validation result; Only present if validation option is included in the request
//
// Union satisfied by [SuiTransactionScanResponseValidationSuiValidationResult] or
// [SuiTransactionScanResponseValidationSuiValidationErrorSchema].
type SuiTransactionScanResponseValidationUnion interface {
	implementsSuiTransactionScanResponseValidation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SuiTransactionScanResponseValidationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiTransactionScanResponseValidationSuiValidationResult{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiTransactionScanResponseValidationSuiValidationErrorSchema{}),
		},
	)
}

type SuiTransactionScanResponseValidationSuiValidationResult struct {
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification,required"`
	// A textual description about the validation result
	Description string                                                           `json:"description,required"`
	Features    []SuiTransactionScanResponseValidationSuiValidationResultFeature `json:"features,required"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason string `json:"reason,required"`
	// Verdict of the validation
	ResultType SuiTransactionScanResponseValidationSuiValidationResultResultType `json:"result_type,required"`
	Status     SuiTransactionScanResponseValidationSuiValidationResultStatus     `json:"status,required"`
	JSON       suiTransactionScanResponseValidationSuiValidationResultJSON       `json:"-"`
}

// suiTransactionScanResponseValidationSuiValidationResultJSON contains the JSON
// metadata for the struct
// [SuiTransactionScanResponseValidationSuiValidationResult]
type suiTransactionScanResponseValidationSuiValidationResultJSON struct {
	Classification apijson.Field
	Description    apijson.Field
	Features       apijson.Field
	Reason         apijson.Field
	ResultType     apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SuiTransactionScanResponseValidationSuiValidationResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiTransactionScanResponseValidationSuiValidationResultJSON) RawJSON() string {
	return r.raw
}

func (r SuiTransactionScanResponseValidationSuiValidationResult) implementsSuiTransactionScanResponseValidation() {
}

type SuiTransactionScanResponseValidationSuiValidationResultFeature struct {
	// Address the feature refers to
	Address string `json:"address,required"`
	// Textual description
	Description string `json:"description,required"`
	FeatureID   string `json:"feature_id,required"`
	// Feature Classification
	Type SuiTransactionScanResponseValidationSuiValidationResultFeaturesType `json:"type,required"`
	JSON suiTransactionScanResponseValidationSuiValidationResultFeatureJSON  `json:"-"`
}

// suiTransactionScanResponseValidationSuiValidationResultFeatureJSON contains the
// JSON metadata for the struct
// [SuiTransactionScanResponseValidationSuiValidationResultFeature]
type suiTransactionScanResponseValidationSuiValidationResultFeatureJSON struct {
	Address     apijson.Field
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiTransactionScanResponseValidationSuiValidationResultFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiTransactionScanResponseValidationSuiValidationResultFeatureJSON) RawJSON() string {
	return r.raw
}

// Feature Classification
type SuiTransactionScanResponseValidationSuiValidationResultFeaturesType string

const (
	SuiTransactionScanResponseValidationSuiValidationResultFeaturesTypeBenign    SuiTransactionScanResponseValidationSuiValidationResultFeaturesType = "Benign"
	SuiTransactionScanResponseValidationSuiValidationResultFeaturesTypeWarning   SuiTransactionScanResponseValidationSuiValidationResultFeaturesType = "Warning"
	SuiTransactionScanResponseValidationSuiValidationResultFeaturesTypeMalicious SuiTransactionScanResponseValidationSuiValidationResultFeaturesType = "Malicious"
	SuiTransactionScanResponseValidationSuiValidationResultFeaturesTypeInfo      SuiTransactionScanResponseValidationSuiValidationResultFeaturesType = "Info"
)

func (r SuiTransactionScanResponseValidationSuiValidationResultFeaturesType) IsKnown() bool {
	switch r {
	case SuiTransactionScanResponseValidationSuiValidationResultFeaturesTypeBenign, SuiTransactionScanResponseValidationSuiValidationResultFeaturesTypeWarning, SuiTransactionScanResponseValidationSuiValidationResultFeaturesTypeMalicious, SuiTransactionScanResponseValidationSuiValidationResultFeaturesTypeInfo:
		return true
	}
	return false
}

// Verdict of the validation
type SuiTransactionScanResponseValidationSuiValidationResultResultType string

const (
	SuiTransactionScanResponseValidationSuiValidationResultResultTypeBenign    SuiTransactionScanResponseValidationSuiValidationResultResultType = "Benign"
	SuiTransactionScanResponseValidationSuiValidationResultResultTypeWarning   SuiTransactionScanResponseValidationSuiValidationResultResultType = "Warning"
	SuiTransactionScanResponseValidationSuiValidationResultResultTypeMalicious SuiTransactionScanResponseValidationSuiValidationResultResultType = "Malicious"
)

func (r SuiTransactionScanResponseValidationSuiValidationResultResultType) IsKnown() bool {
	switch r {
	case SuiTransactionScanResponseValidationSuiValidationResultResultTypeBenign, SuiTransactionScanResponseValidationSuiValidationResultResultTypeWarning, SuiTransactionScanResponseValidationSuiValidationResultResultTypeMalicious:
		return true
	}
	return false
}

type SuiTransactionScanResponseValidationSuiValidationResultStatus string

const (
	SuiTransactionScanResponseValidationSuiValidationResultStatusSuccess SuiTransactionScanResponseValidationSuiValidationResultStatus = "Success"
)

func (r SuiTransactionScanResponseValidationSuiValidationResultStatus) IsKnown() bool {
	switch r {
	case SuiTransactionScanResponseValidationSuiValidationResultStatusSuccess:
		return true
	}
	return false
}

type SuiTransactionScanResponseValidationSuiValidationErrorSchema struct {
	// Error message
	Error  string                                                             `json:"error,required"`
	Status SuiTransactionScanResponseValidationSuiValidationErrorSchemaStatus `json:"status,required"`
	JSON   suiTransactionScanResponseValidationSuiValidationErrorSchemaJSON   `json:"-"`
}

// suiTransactionScanResponseValidationSuiValidationErrorSchemaJSON contains the
// JSON metadata for the struct
// [SuiTransactionScanResponseValidationSuiValidationErrorSchema]
type suiTransactionScanResponseValidationSuiValidationErrorSchemaJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiTransactionScanResponseValidationSuiValidationErrorSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiTransactionScanResponseValidationSuiValidationErrorSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SuiTransactionScanResponseValidationSuiValidationErrorSchema) implementsSuiTransactionScanResponseValidation() {
}

type SuiTransactionScanResponseValidationSuiValidationErrorSchemaStatus string

const (
	SuiTransactionScanResponseValidationSuiValidationErrorSchemaStatusError SuiTransactionScanResponseValidationSuiValidationErrorSchemaStatus = "Error"
)

func (r SuiTransactionScanResponseValidationSuiValidationErrorSchemaStatus) IsKnown() bool {
	switch r {
	case SuiTransactionScanResponseValidationSuiValidationErrorSchemaStatusError:
		return true
	}
	return false
}

type SuiTransactionScanResponseValidationStatus string

const (
	SuiTransactionScanResponseValidationStatusSuccess SuiTransactionScanResponseValidationStatus = "Success"
	SuiTransactionScanResponseValidationStatusError   SuiTransactionScanResponseValidationStatus = "Error"
)

func (r SuiTransactionScanResponseValidationStatus) IsKnown() bool {
	switch r {
	case SuiTransactionScanResponseValidationStatusSuccess, SuiTransactionScanResponseValidationStatusError:
		return true
	}
	return false
}

// Verdict of the validation
type SuiTransactionScanResponseValidationResultType string

const (
	SuiTransactionScanResponseValidationResultTypeBenign    SuiTransactionScanResponseValidationResultType = "Benign"
	SuiTransactionScanResponseValidationResultTypeWarning   SuiTransactionScanResponseValidationResultType = "Warning"
	SuiTransactionScanResponseValidationResultTypeMalicious SuiTransactionScanResponseValidationResultType = "Malicious"
)

func (r SuiTransactionScanResponseValidationResultType) IsKnown() bool {
	switch r {
	case SuiTransactionScanResponseValidationResultTypeBenign, SuiTransactionScanResponseValidationResultTypeWarning, SuiTransactionScanResponseValidationResultTypeMalicious:
		return true
	}
	return false
}
