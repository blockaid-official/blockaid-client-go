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

// Scan a transaction that was already executed on chain, returns validation with
// features indicating address poisoning entities and malicious operators.
func (r *SuiPostTransactionService) Scan(ctx context.Context, body SuiPostTransactionScanParams, opts ...option.RequestOption) (res *SuiPostTransactionScanResponse, err error) {
	opts = slices.Concat(r.Options, opts)
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
	// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummary].
	AccountSummary interface{} `json:"account_summary"`
	// This field can have the runtime type of
	// [[]SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAddressDetail].
	AddressDetails interface{} `json:"address_details"`
	// This field can have the runtime type of
	// [map[string][]SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiff].
	AssetsDiffs interface{} `json:"assets_diffs"`
	// Error message
	Error string `json:"error"`
	// This field can have the runtime type of
	// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultParams].
	Params interface{}                                  `json:"params"`
	JSON   suiPostTransactionScanResponseSimulationJSON `json:"-"`
	union  SuiPostTransactionScanResponseSimulationUnion
}

// suiPostTransactionScanResponseSimulationJSON contains the JSON metadata for the
// struct [SuiPostTransactionScanResponseSimulation]
type suiPostTransactionScanResponseSimulationJSON struct {
	Status         apijson.Field
	AccountSummary apijson.Field
	AddressDetails apijson.Field
	AssetsDiffs    apijson.Field
	Error          apijson.Field
	Params         apijson.Field
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
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResult],
// [SuiPostTransactionScanResponseSimulationSuiSimulationErrorSchema].
func (r SuiPostTransactionScanResponseSimulation) AsUnion() SuiPostTransactionScanResponseSimulationUnion {
	return r.union
}

// Simulation result; Only present if simulation option is included in the request
//
// Union satisfied by
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResult] or
// [SuiPostTransactionScanResponseSimulationSuiSimulationErrorSchema].
type SuiPostTransactionScanResponseSimulationUnion interface {
	implementsSuiPostTransactionScanResponseSimulation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SuiPostTransactionScanResponseSimulationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResult{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiPostTransactionScanResponseSimulationSuiSimulationErrorSchema{}),
		},
	)
}

type SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResult struct {
	// Summary of the actions and asset transfers that were made by the requested
	// account address
	AccountSummary SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummary `json:"account_summary,required"`
	Params         SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultParams         `json:"params,required"`
	Status         SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultStatus         `json:"status,required"`
	// Details of addresses involved in the transaction
	AddressDetails []SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAddressDetail `json:"address_details"`
	// Mapping between the address of an account to the assets diff during the
	// transaction
	AssetsDiffs map[string][]SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiff `json:"assets_diffs"`
	JSON        suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultJSON                    `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResult]
type suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultJSON struct {
	AccountSummary apijson.Field
	Params         apijson.Field
	Status         apijson.Field
	AddressDetails apijson.Field
	AssetsDiffs    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultJSON) RawJSON() string {
	return r.raw
}

func (r SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResult) implementsSuiPostTransactionScanResponseSimulation() {
}

// Summary of the actions and asset transfers that were made by the requested
// account address
type SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummary struct {
	// Total USD diff for the requested account address
	TotalUsdDiff SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryTotalUsdDiff `json:"total_usd_diff,required"`
	// Assets diffs of the requested account address
	AccountAssetsDiffs []SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiff `json:"account_assets_diffs"`
	JSON               suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryJSON                `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummary]
type suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryJSON struct {
	TotalUsdDiff       apijson.Field
	AccountAssetsDiffs apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryJSON) RawJSON() string {
	return r.raw
}

// Total USD diff for the requested account address
type SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryTotalUsdDiff struct {
	// Total incoming USD transfers
	In float64 `json:"in,required"`
	// Total outgoing USD transfers
	Out float64 `json:"out,required"`
	// Total USD transfers
	Total float64                                                                                                  `json:"total"`
	JSON  suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryTotalUsdDiffJSON `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryTotalUsdDiffJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryTotalUsdDiff]
type suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

type SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiff struct {
	// This field can have the runtime type of [SuiNativeAssetDetailsSchema],
	// [SuiNFTDetailsSchema],
	// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAsset].
	Asset interface{} `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of [SuiAssetTransferDetailsSchema],
	// [SuiNFTDiffSchema].
	In interface{} `json:"in"`
	// This field can have the runtime type of [SuiAssetTransferDetailsSchema],
	// [SuiNFTDiffSchema].
	Out   interface{}                                                                                                   `json:"out"`
	JSON  suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffJSON `json:"-"`
	union SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsUnion
}

// suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiff]
type suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiff],
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiff],
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiff].
func (r SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiff) AsUnion() SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiff],
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiff]
// or
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiff].
type SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsUnion interface {
	implementsSuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiff{}),
		},
	)
}

type SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiff struct {
	Asset SuiNativeAssetDetailsSchema `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SuiAssetTransferDetailsSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SuiAssetTransferDetailsSchema                                                                                                    `json:"out,nullable"`
	JSON suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiffJSON `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiffJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiff]
type suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiNativeAssetDiff) implementsSuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiff() {
}

type SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiff struct {
	Asset SuiNFTDetailsSchema `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SuiNFTDiffSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SuiNFTDiffSchema                                                                                                              `json:"out,nullable"`
	JSON suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiffJSON `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiffJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiff]
type suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiNFTAssetDiff) implementsSuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiff() {
}

type SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiff struct {
	Asset SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SuiAssetTransferDetailsSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SuiAssetTransferDetailsSchema                                                                                                   `json:"out,nullable"`
	JSON suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffJSON `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiff]
type suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiff) implementsSuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiff() {
}

type SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAsset struct {
	// Token's package address
	Address string `json:"address,required"`
	// Token's decimal precision
	Decimals int64 `json:"decimals,required"`
	// Token's name
	Name string `json:"name,required"`
	// Token's symbol (abbreviation)
	Symbol            string    `json:"symbol,required"`
	CreationTimestamp time.Time `json:"creation_timestamp,nullable" format:"date-time"`
	// Address of the token's deployer
	Deployer string `json:"deployer,nullable"`
	// URL of the token's logo
	LogoURL     string `json:"logo_url,nullable"`
	Scam        bool   `json:"scam,nullable"`
	TotalSupply int64  `json:"total_supply,nullable"`
	// Type of the asset (`Coin`)
	Type     SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetType `json:"type"`
	Verified bool                                                                                                                                 `json:"verified,nullable"`
	JSON     suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetJSON `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAsset]
type suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetJSON struct {
	Address           apijson.Field
	Decimals          apijson.Field
	Name              apijson.Field
	Symbol            apijson.Field
	CreationTimestamp apijson.Field
	Deployer          apijson.Field
	LogoURL           apijson.Field
	Scam              apijson.Field
	TotalSupply       apijson.Field
	Type              apijson.Field
	Verified          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`Coin`)
type SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetType string

const (
	SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetTypeFungible SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetType = "fungible"
)

func (r SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetType) IsKnown() bool {
	switch r {
	case SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAccountSummaryAccountAssetsDiffsSuiCoinsAssetDiffAssetTypeFungible:
		return true
	}
	return false
}

type SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultParams struct {
	BlockTag    string                                                                               `json:"block_tag,required"`
	Chain       string                                                                               `json:"chain,required"`
	From        string                                                                               `json:"from,required"`
	Value       string                                                                               `json:"value,required"`
	ExtraFields map[string]interface{}                                                               `json:"-,extras"`
	JSON        suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultParamsJSON `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultParamsJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultParams]
type suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultParamsJSON struct {
	BlockTag    apijson.Field
	Chain       apijson.Field
	From        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultParamsJSON) RawJSON() string {
	return r.raw
}

type SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultStatus string

const (
	SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultStatusSuccess SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultStatus = "Success"
)

func (r SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultStatus) IsKnown() bool {
	switch r {
	case SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultStatusSuccess:
		return true
	}
	return false
}

type SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAddressDetail struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// Description of the account
	Description string                                                                                      `json:"description,nullable"`
	JSON        suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAddressDetailJSON `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAddressDetailJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAddressDetail]
type suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAddressDetailJSON struct {
	AccountAddress apijson.Field
	Description    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAddressDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAddressDetailJSON) RawJSON() string {
	return r.raw
}

type SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiff struct {
	// This field can have the runtime type of [SuiNativeAssetDetailsSchema],
	// [SuiNFTDetailsSchema],
	// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiffAsset].
	Asset interface{} `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of [SuiAssetTransferDetailsSchema],
	// [SuiNFTDiffSchema].
	In interface{} `json:"in"`
	// This field can have the runtime type of [SuiAssetTransferDetailsSchema],
	// [SuiNFTDiffSchema].
	Out   interface{}                                                                              `json:"out"`
	JSON  suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffJSON `json:"-"`
	union SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsUnion
}

// suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiff]
type suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiNativeAssetDiff],
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiNFTAssetDiff],
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiff].
func (r SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiff) AsUnion() SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiNativeAssetDiff],
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiNFTAssetDiff]
// or
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiff].
type SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsUnion interface {
	implementsSuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiNativeAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiNFTAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiff{}),
		},
	)
}

type SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiNativeAssetDiff struct {
	Asset SuiNativeAssetDetailsSchema `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SuiAssetTransferDetailsSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SuiAssetTransferDetailsSchema                                                                               `json:"out,nullable"`
	JSON suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiNativeAssetDiffJSON `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiNativeAssetDiffJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiNativeAssetDiff]
type suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiNativeAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiNativeAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiNativeAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiNativeAssetDiff) implementsSuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiff() {
}

type SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiNFTAssetDiff struct {
	Asset SuiNFTDetailsSchema `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SuiNFTDiffSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SuiNFTDiffSchema                                                                                         `json:"out,nullable"`
	JSON suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiNFTAssetDiffJSON `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiNFTAssetDiffJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiNFTAssetDiff]
type suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiNFTAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiNFTAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiNFTAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiNFTAssetDiff) implementsSuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiff() {
}

type SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiff struct {
	Asset SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SuiAssetTransferDetailsSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SuiAssetTransferDetailsSchema                                                                              `json:"out,nullable"`
	JSON suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiffJSON `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiffJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiff]
type suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiff) implementsSuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiff() {
}

type SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiffAsset struct {
	// Token's package address
	Address string `json:"address,required"`
	// Token's decimal precision
	Decimals int64 `json:"decimals,required"`
	// Token's name
	Name string `json:"name,required"`
	// Token's symbol (abbreviation)
	Symbol            string    `json:"symbol,required"`
	CreationTimestamp time.Time `json:"creation_timestamp,nullable" format:"date-time"`
	// Address of the token's deployer
	Deployer string `json:"deployer,nullable"`
	// URL of the token's logo
	LogoURL     string `json:"logo_url,nullable"`
	Scam        bool   `json:"scam,nullable"`
	TotalSupply int64  `json:"total_supply,nullable"`
	// Type of the asset (`Coin`)
	Type     SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetType `json:"type"`
	Verified bool                                                                                                            `json:"verified,nullable"`
	JSON     suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetJSON `json:"-"`
}

// suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiffAsset]
type suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetJSON struct {
	Address           apijson.Field
	Decimals          apijson.Field
	Name              apijson.Field
	Symbol            apijson.Field
	CreationTimestamp apijson.Field
	Deployer          apijson.Field
	LogoURL           apijson.Field
	Scam              apijson.Field
	TotalSupply       apijson.Field
	Type              apijson.Field
	Verified          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`Coin`)
type SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetType string

const (
	SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetTypeFungible SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetType = "fungible"
)

func (r SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetType) IsKnown() bool {
	switch r {
	case SuiPostTransactionScanResponseSimulationSuiPostTransactionSimulationResultAssetsDiffsSuiCoinsAssetDiffAssetTypeFungible:
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
	SuiPostTransactionScanResponseValidationSuiValidationResultResultTypeError     SuiPostTransactionScanResponseValidationSuiValidationResultResultType = "Error"
)

func (r SuiPostTransactionScanResponseValidationSuiValidationResultResultType) IsKnown() bool {
	switch r {
	case SuiPostTransactionScanResponseValidationSuiValidationResultResultTypeBenign, SuiPostTransactionScanResponseValidationSuiValidationResultResultTypeWarning, SuiPostTransactionScanResponseValidationSuiValidationResultResultTypeMalicious, SuiPostTransactionScanResponseValidationSuiValidationResultResultTypeError:
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
	SuiPostTransactionScanResponseValidationResultTypeError     SuiPostTransactionScanResponseValidationResultType = "Error"
)

func (r SuiPostTransactionScanResponseValidationResultType) IsKnown() bool {
	switch r {
	case SuiPostTransactionScanResponseValidationResultTypeBenign, SuiPostTransactionScanResponseValidationResultTypeWarning, SuiPostTransactionScanResponseValidationResultTypeMalicious, SuiPostTransactionScanResponseValidationResultTypeError:
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
