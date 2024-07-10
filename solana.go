// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/option"
)

// SolanaService contains methods and other services that help with interacting
// with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSolanaService] method instead.
type SolanaService struct {
	Options []option.RequestOption
	Message *SolanaMessageService
	Address *SolanaAddressService
}

// NewSolanaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSolanaService(opts ...option.RequestOption) (r *SolanaService) {
	r = &SolanaService{}
	r.Options = opts
	r.Message = NewSolanaMessageService(opts...)
	r.Address = NewSolanaAddressService(opts...)
	return
}

type AssetTransferDetailsSchema struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                        `json:"usd_price,nullable"`
	JSON     assetTransferDetailsSchemaJSON `json:"-"`
}

// assetTransferDetailsSchemaJSON contains the JSON metadata for the struct
// [AssetTransferDetailsSchema]
type assetTransferDetailsSchemaJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetTransferDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetTransferDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

type CnftDetailsSchema struct {
	Address  string `json:"address,required"`
	Decimals int64  `json:"decimals,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Logo     string `json:"logo"`
	// Type of the asset (`"CNFT"`)
	Type string                `json:"type"`
	JSON cnftDetailsSchemaJSON `json:"-"`
}

// cnftDetailsSchemaJSON contains the JSON metadata for the struct
// [CnftDetailsSchema]
type cnftDetailsSchemaJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CnftDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cnftDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CnftDetailsSchema) implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsAsset() {
}

func (r CnftDetailsSchema) implementsSolanaMessageScanResponseResultSimulationDelegationsAsset() {}

type NativeSolDetailsSchema struct {
	Decimals int64 `json:"decimals"`
	// Type of the asset (`"SOL"`)
	Type string                     `json:"type"`
	JSON nativeSolDetailsSchemaJSON `json:"-"`
}

// nativeSolDetailsSchemaJSON contains the JSON metadata for the struct
// [NativeSolDetailsSchema]
type nativeSolDetailsSchemaJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NativeSolDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nativeSolDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

type SplFungibleTokenDetailsSchema struct {
	Address  string `json:"address,required"`
	Decimals int64  `json:"decimals,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Logo     string `json:"logo"`
	// Type of the asset (`"TOKEN"`)
	Type string                            `json:"type"`
	JSON splFungibleTokenDetailsSchemaJSON `json:"-"`
}

// splFungibleTokenDetailsSchemaJSON contains the JSON metadata for the struct
// [SplFungibleTokenDetailsSchema]
type splFungibleTokenDetailsSchemaJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplFungibleTokenDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splFungibleTokenDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SplFungibleTokenDetailsSchema) implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsAsset() {
}

func (r SplFungibleTokenDetailsSchema) implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset() {
}

func (r SplFungibleTokenDetailsSchema) implementsSolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset() {
}

func (r SplFungibleTokenDetailsSchema) implementsSolanaMessageScanResponseResultSimulationDelegationsAsset() {
}

type SplNonFungibleTokenDetailsSchema struct {
	Address  string `json:"address,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	Logo     string `json:"logo"`
	// Type of the asset (`"NFT"`)
	Type string                               `json:"type"`
	JSON splNonFungibleTokenDetailsSchemaJSON `json:"-"`
}

// splNonFungibleTokenDetailsSchemaJSON contains the JSON metadata for the struct
// [SplNonFungibleTokenDetailsSchema]
type splNonFungibleTokenDetailsSchemaJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplNonFungibleTokenDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splNonFungibleTokenDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SplNonFungibleTokenDetailsSchema) implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsAsset() {
}

func (r SplNonFungibleTokenDetailsSchema) implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset() {
}

func (r SplNonFungibleTokenDetailsSchema) implementsSolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset() {
}

func (r SplNonFungibleTokenDetailsSchema) implementsSolanaMessageScanResponseResultSimulationDelegationsAsset() {
}
