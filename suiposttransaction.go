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

// SuiPostTransactionService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSuiPostTransactionService] method instead.
type SuiPostTransactionService struct {
	Options []option.RequestOption
}

// NewSuiPostTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSuiPostTransactionService(opts ...option.RequestOption) (r *SuiPostTransactionService) {
	r = &SuiPostTransactionService{}
	r.Options = opts
	return
}

// Scan Post Transaction
func (r *SuiPostTransactionService) Scan(ctx context.Context, body SuiPostTransactionScanParams, opts ...option.RequestOption) (res *SuiPostTransactionScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/sui/post-transaction/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SuiPostTransactionScanResponse struct {
	AccountAddress string `json:"account_address,required"`
	// Simulation result; Only present if simulation option is included in the request
	Simulation SuiPostTransactionScanResponseSimulation `json:"simulation,nullable"`
	// Validation result; Only present if validation option is included in the request
	Validation SuiPostTransactionScanResponseValidation `json:"validation,nullable"`
	JSON       suiPostTransactionScanResponseJSON       `json:"-"`
}

// suiPostTransactionScanResponseJSON contains the JSON metadata for the struct
// [SuiPostTransactionScanResponse]
type suiPostTransactionScanResponseJSON struct {
	AccountAddress apijson.Field
	Simulation     apijson.Field
	Validation     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseJSON) RawJSON() string {
	return r.raw
}

// Simulation result; Only present if simulation option is included in the request
type SuiPostTransactionScanResponseSimulation struct {
	Status SuiPostTransactionScanResponseSimulationStatus `json:"status,required"`
	// This field can have the runtime type of
	// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummary].
	AccountSummary interface{} `json:"account_summary"`
	// This field can have the runtime type of
	// [[]SuiPostTransactionScanResponseSimulationSuiSimulationResultAddressDetail].
	AddressDetails interface{} `json:"address_details"`
	// This field can have the runtime type of
	// [map[string][]SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiff].
	AssetsDiffs interface{} `json:"assets_diffs"`
	// Error message
	Error string                                       `json:"error"`
	JSON  suiPostTransactionScanResponseSimulationJSON `json:"-"`
	union SuiPostTransactionScanResponseSimulationUnion
}

// suiPostTransactionScanResponseSimulationJSON contains the JSON metadata for the
// struct [SuiPostTransactionScanResponseSimulation]
type suiPostTransactionScanResponseSimulationJSON struct {
	Status         apijson.Field
	AccountSummary apijson.Field
	AddressDetails apijson.Field
	AssetsDiffs    apijson.Field
	Error          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r suiPostTransactionScanResponseSimulationJSON) RawJSON() string {
	return r.raw
}

func (r *SuiPostTransactionScanResponseSimulation) UnmarshalJSON(data []byte) (err error) {
	*r = SuiPostTransactionScanResponseSimulation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SuiPostTransactionScanResponseSimulationUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SuiPostTransactionScanResponseSimulationSuiSimulationResult],
// [SuiPostTransactionScanResponseSimulationSuiSimulationErrorSchema].
func (r SuiPostTransactionScanResponseSimulation) AsUnion() SuiPostTransactionScanResponseSimulationUnion {
	return r.union
}

// Simulation result; Only present if simulation option is included in the request
//
// Union satisfied by [SuiPostTransactionScanResponseSimulationSuiSimulationResult]
// or [SuiPostTransactionScanResponseSimulationSuiSimulationErrorSchema].
type SuiPostTransactionScanResponseSimulationUnion interface {
	implementsSuiPostTransactionScanResponseSimulation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SuiPostTransactionScanResponseSimulationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiPostTransactionScanResponseSimulationSuiSimulationResult{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiPostTransactionScanResponseSimulationSuiSimulationErrorSchema{}),
		},
	)
}

type SuiPostTransactionScanResponseSimulationSuiSimulationResult struct {
	// Summary of the actions and asset transfers that were made by the requested
	// account address
	AccountSummary SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummary `json:"account_summary,required"`
	Status         SuiPostTransactionScanResponseSimulationSuiSimulationResultStatus         `json:"status,required"`
	// Details of addresses involved in the transaction
	AddressDetails []SuiPostTransactionScanResponseSimulationSuiSimulationResultAddressDetail `json:"address_details"`
	// Mapping between the address of an account to the assets diff during the
	// transaction
	AssetsDiffs map[string][]SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiff `json:"assets_diffs"`
	JSON        suiPostTransactionScanResponseSimulationSuiSimulationResultJSON                    `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiSimulationResultJSON contains the
// JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiSimulationResult]
type suiPostTransactionScanResponseSimulationSuiSimulationResultJSON struct {
	AccountSummary apijson.Field
	Status         apijson.Field
	AddressDetails apijson.Field
	AssetsDiffs    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseSimulationSuiSimulationResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiSimulationResultJSON) RawJSON() string {
	return r.raw
}

func (r SuiPostTransactionScanResponseSimulationSuiSimulationResult) implementsSuiPostTransactionScanResponseSimulation() {
}

// Summary of the actions and asset transfers that were made by the requested
// account address
type SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummary struct {
	// Total USD diff for the requested account address
	TotalUsdDiff SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryTotalUsdDiff `json:"total_usd_diff,required"`
	// Assets diffs of the requested account address
	AccountAssetsDiffs []SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiff `json:"account_assets_diffs"`
	JSON               suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryJSON                `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummary]
type suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryJSON struct {
	TotalUsdDiff       apijson.Field
	AccountAssetsDiffs apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryJSON) RawJSON() string {
	return r.raw
}

// Total USD diff for the requested account address
type SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryTotalUsdDiff struct {
	// Total incoming USD transfers
	In float64 `json:"in,required"`
	// Total outgoing USD transfers
	Out float64 `json:"out,required"`
	// Total USD transfers
	Total float64                                                                                   `json:"total"`
	JSON  suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryTotalUsdDiffJSON `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryTotalUsdDiffJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryTotalUsdDiff]
type suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

type SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiff struct {
	// This field can have the runtime type of [SuiNativeAssetDetailsSchema],
	// [SuiNFTDetailsSchema],
	// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAsset].
	Asset interface{} `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of [SuiAssetTransferDetailsSchema],
	// [SuiNFTDiffSchema].
	In interface{} `json:"in"`
	// This field can have the runtime type of [SuiAssetTransferDetailsSchema],
	// [SuiNFTDiffSchema].
	Out   interface{}                                                                                    `json:"out"`
	JSON  suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffJSON `json:"-"`
	union SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsUnion
}

// suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiff]
type suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiff],
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiff],
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiff].
func (r SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiff) AsUnion() SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiff],
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiff]
// or
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiff].
type SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsUnion interface {
	implementsSuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiff{}),
		},
	)
}

type SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiff struct {
	Asset SuiNativeAssetDetailsSchema `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SuiAssetTransferDetailsSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SuiAssetTransferDetailsSchema                                                                                     `json:"out,nullable"`
	JSON suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiffJSON `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiffJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiff]
type suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiff) implementsSuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiff() {
}

type SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiff struct {
	Asset SuiNFTDetailsSchema `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SuiNFTDiffSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SuiNFTDiffSchema                                                                                               `json:"out,nullable"`
	JSON suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiffJSON `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiffJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiff]
type suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiff) implementsSuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiff() {
}

type SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiff struct {
	Asset SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SuiAssetTransferDetailsSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SuiAssetTransferDetailsSchema                                                                                    `json:"out,nullable"`
	JSON suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffJSON `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiff]
type suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiff) implementsSuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiff() {
}

type SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAsset struct {
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
	Type     SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetType `json:"type"`
	Verified bool                                                                                                                  `json:"verified,nullable"`
	JSON     suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetJSON `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAsset]
type suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetJSON struct {
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

func (r *SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`Coin`)
type SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetType string

const (
	SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetTypeFungible SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetType = "fungible"
)

func (r SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetType) IsKnown() bool {
	switch r {
	case SuiPostTransactionScanResponseSimulationSuiSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetTypeFungible:
		return true
	}
	return false
}

type SuiPostTransactionScanResponseSimulationSuiSimulationResultStatus string

const (
	SuiPostTransactionScanResponseSimulationSuiSimulationResultStatusSuccess SuiPostTransactionScanResponseSimulationSuiSimulationResultStatus = "Success"
)

func (r SuiPostTransactionScanResponseSimulationSuiSimulationResultStatus) IsKnown() bool {
	switch r {
	case SuiPostTransactionScanResponseSimulationSuiSimulationResultStatusSuccess:
		return true
	}
	return false
}

type SuiPostTransactionScanResponseSimulationSuiSimulationResultAddressDetail struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// Description of the account
	Description string                                                                       `json:"description,nullable"`
	JSON        suiPostTransactionScanResponseSimulationSuiSimulationResultAddressDetailJSON `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiSimulationResultAddressDetailJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAddressDetail]
type suiPostTransactionScanResponseSimulationSuiSimulationResultAddressDetailJSON struct {
	AccountAddress apijson.Field
	Description    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseSimulationSuiSimulationResultAddressDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiSimulationResultAddressDetailJSON) RawJSON() string {
	return r.raw
}

type SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiff struct {
	// This field can have the runtime type of [SuiNativeAssetDetailsSchema],
	// [SuiNFTDetailsSchema],
	// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAsset].
	Asset interface{} `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of [SuiAssetTransferDetailsSchema],
	// [SuiNFTDiffSchema].
	In interface{} `json:"in"`
	// This field can have the runtime type of [SuiAssetTransferDetailsSchema],
	// [SuiNFTDiffSchema].
	Out   interface{}                                                               `json:"out"`
	JSON  suiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffJSON `json:"-"`
	union SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsUnion
}

// suiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiff]
type suiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r suiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNativeAssetDiff],
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNFTAssetDiff],
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiff].
func (r SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiff) AsUnion() SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNativeAssetDiff],
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNFTAssetDiff]
// or
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiff].
type SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsUnion interface {
	implementsSuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNativeAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNFTAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiff{}),
		},
	)
}

type SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNativeAssetDiff struct {
	Asset SuiNativeAssetDetailsSchema `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SuiAssetTransferDetailsSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SuiAssetTransferDetailsSchema                                                                `json:"out,nullable"`
	JSON suiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNativeAssetDiffJSON `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNativeAssetDiffJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNativeAssetDiff]
type suiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNativeAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNativeAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNativeAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNativeAssetDiff) implementsSuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiff() {
}

type SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNFTAssetDiff struct {
	Asset SuiNFTDetailsSchema `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SuiNFTDiffSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SuiNFTDiffSchema                                                                          `json:"out,nullable"`
	JSON suiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNFTAssetDiffJSON `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNFTAssetDiffJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNFTAssetDiff]
type suiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNFTAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNFTAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNFTAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiNFTAssetDiff) implementsSuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiff() {
}

type SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiff struct {
	Asset SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SuiAssetTransferDetailsSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SuiAssetTransferDetailsSchema                                                               `json:"out,nullable"`
	JSON suiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffJSON `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiff]
type suiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiff) implementsSuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiff() {
}

type SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAsset struct {
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
	Type     SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetType `json:"type"`
	Verified bool                                                                                             `json:"verified,nullable"`
	JSON     suiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetJSON `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAsset]
type suiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetJSON struct {
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

func (r *SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`Coin`)
type SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetType string

const (
	SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetTypeFungible SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetType = "fungible"
)

func (r SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetType) IsKnown() bool {
	switch r {
	case SuiPostTransactionScanResponseSimulationSuiSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetTypeFungible:
		return true
	}
	return false
}

type SuiPostTransactionScanResponseSimulationSuiSimulationErrorSchema struct {
	// Error message
	Error  string                                                                 `json:"error,required"`
	Status SuiPostTransactionScanResponseSimulationSuiSimulationErrorSchemaStatus `json:"status,required"`
	JSON   suiPostTransactionScanResponseSimulationSuiSimulationErrorSchemaJSON   `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiSimulationErrorSchemaJSON contains
// the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiSimulationErrorSchema]
type suiPostTransactionScanResponseSimulationSuiSimulationErrorSchemaJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseSimulationSuiSimulationErrorSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiSimulationErrorSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SuiPostTransactionScanResponseSimulationSuiSimulationErrorSchema) implementsSuiPostTransactionScanResponseSimulation() {
}

type SuiPostTransactionScanResponseSimulationSuiSimulationErrorSchemaStatus string

const (
	SuiPostTransactionScanResponseSimulationSuiSimulationErrorSchemaStatusError SuiPostTransactionScanResponseSimulationSuiSimulationErrorSchemaStatus = "Error"
)

func (r SuiPostTransactionScanResponseSimulationSuiSimulationErrorSchemaStatus) IsKnown() bool {
	switch r {
	case SuiPostTransactionScanResponseSimulationSuiSimulationErrorSchemaStatusError:
		return true
	}
	return false
}

type SuiPostTransactionScanResponseSimulationStatus string

const (
	SuiPostTransactionScanResponseSimulationStatusSuccess SuiPostTransactionScanResponseSimulationStatus = "Success"
	SuiPostTransactionScanResponseSimulationStatusError   SuiPostTransactionScanResponseSimulationStatus = "Error"
)

func (r SuiPostTransactionScanResponseSimulationStatus) IsKnown() bool {
	switch r {
	case SuiPostTransactionScanResponseSimulationStatusSuccess, SuiPostTransactionScanResponseSimulationStatusError:
		return true
	}
	return false
}

// Validation result; Only present if validation option is included in the request
type SuiPostTransactionScanResponseValidation struct {
	Status SuiPostTransactionScanResponseValidationStatus `json:"status,required"`
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification"`
	// A textual description about the validation result
	Description string `json:"description"`
	// Error message
	Error string `json:"error"`
	// This field can have the runtime type of
	// [[]SuiPostTransactionScanResponseValidationSuiValidationResultFeature].
	Features interface{} `json:"features"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason string `json:"reason"`
	// Verdict of the validation
	ResultType SuiPostTransactionScanResponseValidationResultType `json:"result_type"`
	JSON       suiPostTransactionScanResponseValidationJSON       `json:"-"`
	union      SuiPostTransactionScanResponseValidationUnion
}

// suiPostTransactionScanResponseValidationJSON contains the JSON metadata for the
// struct [SuiPostTransactionScanResponseValidation]
type suiPostTransactionScanResponseValidationJSON struct {
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

func (r suiPostTransactionScanResponseValidationJSON) RawJSON() string {
	return r.raw
}

func (r *SuiPostTransactionScanResponseValidation) UnmarshalJSON(data []byte) (err error) {
	*r = SuiPostTransactionScanResponseValidation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SuiPostTransactionScanResponseValidationUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SuiPostTransactionScanResponseValidationSuiValidationResult],
// [SuiPostTransactionScanResponseValidationSuiValidationErrorSchema].
func (r SuiPostTransactionScanResponseValidation) AsUnion() SuiPostTransactionScanResponseValidationUnion {
	return r.union
}

// Validation result; Only present if validation option is included in the request
//
// Union satisfied by [SuiPostTransactionScanResponseValidationSuiValidationResult]
// or [SuiPostTransactionScanResponseValidationSuiValidationErrorSchema].
type SuiPostTransactionScanResponseValidationUnion interface {
	implementsSuiPostTransactionScanResponseValidation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SuiPostTransactionScanResponseValidationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiPostTransactionScanResponseValidationSuiValidationResult{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiPostTransactionScanResponseValidationSuiValidationErrorSchema{}),
		},
	)
}

type SuiPostTransactionScanResponseValidationSuiValidationResult struct {
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification,required"`
	// A textual description about the validation result
	Description string                                                               `json:"description,required"`
	Features    []SuiPostTransactionScanResponseValidationSuiValidationResultFeature `json:"features,required"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason string `json:"reason,required"`
	// Verdict of the validation
	ResultType SuiPostTransactionScanResponseValidationSuiValidationResultResultType `json:"result_type,required"`
	Status     SuiPostTransactionScanResponseValidationSuiValidationResultStatus     `json:"status,required"`
	JSON       suiPostTransactionScanResponseValidationSuiValidationResultJSON       `json:"-"`
}

// suiPostTransactionScanResponseValidationSuiValidationResultJSON contains the
// JSON metadata for the struct
// [SuiPostTransactionScanResponseValidationSuiValidationResult]
type suiPostTransactionScanResponseValidationSuiValidationResultJSON struct {
	Classification apijson.Field
	Description    apijson.Field
	Features       apijson.Field
	Reason         apijson.Field
	ResultType     apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseValidationSuiValidationResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseValidationSuiValidationResultJSON) RawJSON() string {
	return r.raw
}

func (r SuiPostTransactionScanResponseValidationSuiValidationResult) implementsSuiPostTransactionScanResponseValidation() {
}

type SuiPostTransactionScanResponseValidationSuiValidationResultFeature struct {
	// Address the feature refers to
	Address string `json:"address,required"`
	// Textual description
	Description string `json:"description,required"`
	FeatureID   string `json:"feature_id,required"`
	// Feature Classification
	Type SuiPostTransactionScanResponseValidationSuiValidationResultFeaturesType `json:"type,required"`
	JSON suiPostTransactionScanResponseValidationSuiValidationResultFeatureJSON  `json:"-"`
}

// suiPostTransactionScanResponseValidationSuiValidationResultFeatureJSON contains
// the JSON metadata for the struct
// [SuiPostTransactionScanResponseValidationSuiValidationResultFeature]
type suiPostTransactionScanResponseValidationSuiValidationResultFeatureJSON struct {
	Address     apijson.Field
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseValidationSuiValidationResultFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseValidationSuiValidationResultFeatureJSON) RawJSON() string {
	return r.raw
}

// Feature Classification
type SuiPostTransactionScanResponseValidationSuiValidationResultFeaturesType string

const (
	SuiPostTransactionScanResponseValidationSuiValidationResultFeaturesTypeBenign    SuiPostTransactionScanResponseValidationSuiValidationResultFeaturesType = "Benign"
	SuiPostTransactionScanResponseValidationSuiValidationResultFeaturesTypeWarning   SuiPostTransactionScanResponseValidationSuiValidationResultFeaturesType = "Warning"
	SuiPostTransactionScanResponseValidationSuiValidationResultFeaturesTypeMalicious SuiPostTransactionScanResponseValidationSuiValidationResultFeaturesType = "Malicious"
	SuiPostTransactionScanResponseValidationSuiValidationResultFeaturesTypeInfo      SuiPostTransactionScanResponseValidationSuiValidationResultFeaturesType = "Info"
)

func (r SuiPostTransactionScanResponseValidationSuiValidationResultFeaturesType) IsKnown() bool {
	switch r {
	case SuiPostTransactionScanResponseValidationSuiValidationResultFeaturesTypeBenign, SuiPostTransactionScanResponseValidationSuiValidationResultFeaturesTypeWarning, SuiPostTransactionScanResponseValidationSuiValidationResultFeaturesTypeMalicious, SuiPostTransactionScanResponseValidationSuiValidationResultFeaturesTypeInfo:
		return true
	}
	return false
}

// Verdict of the validation
type SuiPostTransactionScanResponseValidationSuiValidationResultResultType string

const (
	SuiPostTransactionScanResponseValidationSuiValidationResultResultTypeBenign    SuiPostTransactionScanResponseValidationSuiValidationResultResultType = "Benign"
	SuiPostTransactionScanResponseValidationSuiValidationResultResultTypeWarning   SuiPostTransactionScanResponseValidationSuiValidationResultResultType = "Warning"
	SuiPostTransactionScanResponseValidationSuiValidationResultResultTypeMalicious SuiPostTransactionScanResponseValidationSuiValidationResultResultType = "Malicious"
)

func (r SuiPostTransactionScanResponseValidationSuiValidationResultResultType) IsKnown() bool {
	switch r {
	case SuiPostTransactionScanResponseValidationSuiValidationResultResultTypeBenign, SuiPostTransactionScanResponseValidationSuiValidationResultResultTypeWarning, SuiPostTransactionScanResponseValidationSuiValidationResultResultTypeMalicious:
		return true
	}
	return false
}

type SuiPostTransactionScanResponseValidationSuiValidationResultStatus string

const (
	SuiPostTransactionScanResponseValidationSuiValidationResultStatusSuccess SuiPostTransactionScanResponseValidationSuiValidationResultStatus = "Success"
)

func (r SuiPostTransactionScanResponseValidationSuiValidationResultStatus) IsKnown() bool {
	switch r {
	case SuiPostTransactionScanResponseValidationSuiValidationResultStatusSuccess:
		return true
	}
	return false
}

type SuiPostTransactionScanResponseValidationSuiValidationErrorSchema struct {
	// Error message
	Error  string                                                                 `json:"error,required"`
	Status SuiPostTransactionScanResponseValidationSuiValidationErrorSchemaStatus `json:"status,required"`
	JSON   suiPostTransactionScanResponseValidationSuiValidationErrorSchemaJSON   `json:"-"`
}

// suiPostTransactionScanResponseValidationSuiValidationErrorSchemaJSON contains
// the JSON metadata for the struct
// [SuiPostTransactionScanResponseValidationSuiValidationErrorSchema]
type suiPostTransactionScanResponseValidationSuiValidationErrorSchemaJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseValidationSuiValidationErrorSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseValidationSuiValidationErrorSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SuiPostTransactionScanResponseValidationSuiValidationErrorSchema) implementsSuiPostTransactionScanResponseValidation() {
}

type SuiPostTransactionScanResponseValidationSuiValidationErrorSchemaStatus string

const (
	SuiPostTransactionScanResponseValidationSuiValidationErrorSchemaStatusError SuiPostTransactionScanResponseValidationSuiValidationErrorSchemaStatus = "Error"
)

func (r SuiPostTransactionScanResponseValidationSuiValidationErrorSchemaStatus) IsKnown() bool {
	switch r {
	case SuiPostTransactionScanResponseValidationSuiValidationErrorSchemaStatusError:
		return true
	}
	return false
}

type SuiPostTransactionScanResponseValidationStatus string

const (
	SuiPostTransactionScanResponseValidationStatusSuccess SuiPostTransactionScanResponseValidationStatus = "Success"
	SuiPostTransactionScanResponseValidationStatusError   SuiPostTransactionScanResponseValidationStatus = "Error"
)

func (r SuiPostTransactionScanResponseValidationStatus) IsKnown() bool {
	switch r {
	case SuiPostTransactionScanResponseValidationStatusSuccess, SuiPostTransactionScanResponseValidationStatusError:
		return true
	}
	return false
}

// Verdict of the validation
type SuiPostTransactionScanResponseValidationResultType string

const (
	SuiPostTransactionScanResponseValidationResultTypeBenign    SuiPostTransactionScanResponseValidationResultType = "Benign"
	SuiPostTransactionScanResponseValidationResultTypeWarning   SuiPostTransactionScanResponseValidationResultType = "Warning"
	SuiPostTransactionScanResponseValidationResultTypeMalicious SuiPostTransactionScanResponseValidationResultType = "Malicious"
)

func (r SuiPostTransactionScanResponseValidationResultType) IsKnown() bool {
	switch r {
	case SuiPostTransactionScanResponseValidationResultTypeBenign, SuiPostTransactionScanResponseValidationResultTypeWarning, SuiPostTransactionScanResponseValidationResultTypeMalicious:
		return true
	}
	return false
}

type SuiPostTransactionScanParams struct {
	Chain    param.Field[SuiPostTransactionScanParamsChain]    `json:"chain,required"`
	Data     param.Field[SuiPostTransactionScanParamsData]     `json:"data,required"`
	Metadata param.Field[SuiPostTransactionScanParamsMetadata] `json:"metadata,required"`
	// List of options to include in the response
	//
	// - `Options.validation`: Include Options.validation output in the response
	//
	// - `Options.simulation`: Include Options.simulation output in the response
	Options param.Field[[]SuiPostTransactionScanParamsOption] `json:"options"`
}

func (r SuiPostTransactionScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuiPostTransactionScanParamsChain string

const (
	SuiPostTransactionScanParamsChainMainnet SuiPostTransactionScanParamsChain = "mainnet"
	SuiPostTransactionScanParamsChainTestnet SuiPostTransactionScanParamsChain = "testnet"
	SuiPostTransactionScanParamsChainDevnet  SuiPostTransactionScanParamsChain = "devnet"
)

func (r SuiPostTransactionScanParamsChain) IsKnown() bool {
	switch r {
	case SuiPostTransactionScanParamsChainMainnet, SuiPostTransactionScanParamsChainTestnet, SuiPostTransactionScanParamsChainDevnet:
		return true
	}
	return false
}

type SuiPostTransactionScanParamsData struct {
	TxHash param.Field[string] `json:"tx_hash,required"`
}

func (r SuiPostTransactionScanParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuiPostTransactionScanParamsMetadata struct {
	// cross reference transaction against the domain.
	Domain param.Field[string] `json:"domain"`
	// whether the transaction is initiated by a dapp. Default is false.
	NonDapp param.Field[bool] `json:"non_dapp"`
}

func (r SuiPostTransactionScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuiPostTransactionScanParamsOption string

const (
	SuiPostTransactionScanParamsOptionValidation SuiPostTransactionScanParamsOption = "validation"
	SuiPostTransactionScanParamsOptionSimulation SuiPostTransactionScanParamsOption = "simulation"
)

func (r SuiPostTransactionScanParamsOption) IsKnown() bool {
	switch r {
	case SuiPostTransactionScanParamsOptionValidation, SuiPostTransactionScanParamsOptionSimulation:
		return true
	}
	return false
}
