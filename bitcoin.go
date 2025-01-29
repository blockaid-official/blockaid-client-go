// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"reflect"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/param"
	"github.com/blockaid-official/blockaid-client-go/option"
	"github.com/tidwall/gjson"
)

// BitcoinService contains methods and other services that help with interacting
// with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBitcoinService] method instead.
type BitcoinService struct {
	Options        []option.RequestOption
	TransactionRaw *BitcoinTransactionRawService
}

// NewBitcoinService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBitcoinService(opts ...option.RequestOption) (r *BitcoinService) {
	r = &BitcoinService{}
	r.Options = opts
	r.TransactionRaw = NewBitcoinTransactionRawService(opts...)
	return
}

type BitcoinTransactionScanRequestParam struct {
	AccountAddress param.Field[string]                             `json:"account_address,required"`
	Chain          param.Field[BitcoinTransactionScanRequestChain] `json:"chain,required"`
	// Metadata
	Metadata    param.Field[BitcoinTransactionScanRequestMetadataUnionParam] `json:"metadata,required"`
	Transaction param.Field[string]                                          `json:"transaction,required"`
	// List of options to include in the response
	//
	// - `Options.validation`: Include Options.validation output in the response
	//
	// - `Options.simulation`: Include Options.simulation output in the response
	Options param.Field[[]BitcoinTransactionScanRequestOption] `json:"options"`
}

func (r BitcoinTransactionScanRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BitcoinTransactionScanRequestChain string

const (
	BitcoinTransactionScanRequestChainBitcoin BitcoinTransactionScanRequestChain = "bitcoin"
)

func (r BitcoinTransactionScanRequestChain) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanRequestChainBitcoin:
		return true
	}
	return false
}

// Metadata
type BitcoinTransactionScanRequestMetadataParam struct {
	// Metadata for wallet requests
	Type param.Field[BitcoinTransactionScanRequestMetadataType] `json:"type"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url"`
}

func (r BitcoinTransactionScanRequestMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionScanRequestMetadataParam) implementsBitcoinTransactionScanRequestMetadataUnionParam() {
}

// Metadata
//
// Satisfied by
// [BitcoinTransactionScanRequestMetadataBitcoinWalletRequestMetadataParam],
// [BitcoinTransactionScanRequestMetadataBitcoinInAppRequestMetadataParam],
// [BitcoinTransactionScanRequestMetadataParam].
type BitcoinTransactionScanRequestMetadataUnionParam interface {
	implementsBitcoinTransactionScanRequestMetadataUnionParam()
}

type BitcoinTransactionScanRequestMetadataBitcoinWalletRequestMetadataParam struct {
	// Metadata for wallet requests
	Type param.Field[BitcoinTransactionScanRequestMetadataBitcoinWalletRequestMetadataType] `json:"type,required"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url,required"`
}

func (r BitcoinTransactionScanRequestMetadataBitcoinWalletRequestMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionScanRequestMetadataBitcoinWalletRequestMetadataParam) implementsBitcoinTransactionScanRequestMetadataUnionParam() {
}

// Metadata for wallet requests
type BitcoinTransactionScanRequestMetadataBitcoinWalletRequestMetadataType string

const (
	BitcoinTransactionScanRequestMetadataBitcoinWalletRequestMetadataTypeWallet BitcoinTransactionScanRequestMetadataBitcoinWalletRequestMetadataType = "wallet"
)

func (r BitcoinTransactionScanRequestMetadataBitcoinWalletRequestMetadataType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanRequestMetadataBitcoinWalletRequestMetadataTypeWallet:
		return true
	}
	return false
}

type BitcoinTransactionScanRequestMetadataBitcoinInAppRequestMetadataParam struct {
	// Metadata for in-app requests
	Type param.Field[BitcoinTransactionScanRequestMetadataBitcoinInAppRequestMetadataType] `json:"type"`
}

func (r BitcoinTransactionScanRequestMetadataBitcoinInAppRequestMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionScanRequestMetadataBitcoinInAppRequestMetadataParam) implementsBitcoinTransactionScanRequestMetadataUnionParam() {
}

// Metadata for in-app requests
type BitcoinTransactionScanRequestMetadataBitcoinInAppRequestMetadataType string

const (
	BitcoinTransactionScanRequestMetadataBitcoinInAppRequestMetadataTypeInApp BitcoinTransactionScanRequestMetadataBitcoinInAppRequestMetadataType = "in_app"
)

func (r BitcoinTransactionScanRequestMetadataBitcoinInAppRequestMetadataType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanRequestMetadataBitcoinInAppRequestMetadataTypeInApp:
		return true
	}
	return false
}

// Metadata for wallet requests
type BitcoinTransactionScanRequestMetadataType string

const (
	BitcoinTransactionScanRequestMetadataTypeWallet BitcoinTransactionScanRequestMetadataType = "wallet"
	BitcoinTransactionScanRequestMetadataTypeInApp  BitcoinTransactionScanRequestMetadataType = "in_app"
)

func (r BitcoinTransactionScanRequestMetadataType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanRequestMetadataTypeWallet, BitcoinTransactionScanRequestMetadataTypeInApp:
		return true
	}
	return false
}

type BitcoinTransactionScanRequestOption string

const (
	BitcoinTransactionScanRequestOptionValidation BitcoinTransactionScanRequestOption = "validation"
	BitcoinTransactionScanRequestOptionSimulation BitcoinTransactionScanRequestOption = "simulation"
)

func (r BitcoinTransactionScanRequestOption) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanRequestOptionValidation, BitcoinTransactionScanRequestOptionSimulation:
		return true
	}
	return false
}

type BitcoinTransactionScanResponse struct {
	// Simulation result; Only present if simulation option is included in the request
	Simulation BitcoinTransactionScanResponseSimulation `json:"simulation,nullable"`
	// Validation result; Only present if validation option is included in the request
	Validation BitcoinTransactionScanResponseValidation `json:"validation,nullable"`
	JSON       bitcoinTransactionScanResponseJSON       `json:"-"`
}

// bitcoinTransactionScanResponseJSON contains the JSON metadata for the struct
// [BitcoinTransactionScanResponse]
type bitcoinTransactionScanResponseJSON struct {
	Simulation  apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseJSON) RawJSON() string {
	return r.raw
}

// Simulation result; Only present if simulation option is included in the request
type BitcoinTransactionScanResponseSimulation struct {
	Status BitcoinTransactionScanResponseSimulationStatus `json:"status,required"`
	// This field can have the runtime type of
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummary].
	AccountSummary interface{} `json:"account_summary"`
	// This field can have the runtime type of
	// [[]BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAddressDetail].
	AddressDetails interface{} `json:"address_details"`
	// This field can have the runtime type of
	// [map[string][]BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiff].
	AssetsDiffs interface{} `json:"assets_diffs"`
	// Error message
	Error string                                       `json:"error"`
	JSON  bitcoinTransactionScanResponseSimulationJSON `json:"-"`
	union BitcoinTransactionScanResponseSimulationUnion
}

// bitcoinTransactionScanResponseSimulationJSON contains the JSON metadata for the
// struct [BitcoinTransactionScanResponseSimulation]
type bitcoinTransactionScanResponseSimulationJSON struct {
	Status         apijson.Field
	AccountSummary apijson.Field
	AddressDetails apijson.Field
	AssetsDiffs    apijson.Field
	Error          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r bitcoinTransactionScanResponseSimulationJSON) RawJSON() string {
	return r.raw
}

func (r *BitcoinTransactionScanResponseSimulation) UnmarshalJSON(data []byte) (err error) {
	*r = BitcoinTransactionScanResponseSimulation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BitcoinTransactionScanResponseSimulationUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResult],
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchema].
func (r BitcoinTransactionScanResponseSimulation) AsUnion() BitcoinTransactionScanResponseSimulationUnion {
	return r.union
}

// Simulation result; Only present if simulation option is included in the request
//
// Union satisfied by
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResult] or
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchema].
type BitcoinTransactionScanResponseSimulationUnion interface {
	implementsBitcoinTransactionScanResponseSimulation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BitcoinTransactionScanResponseSimulationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BitcoinTransactionScanResponseSimulationBitcoinSimulationResult{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchema{}),
		},
	)
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResult struct {
	// Summary of the actions and asset transfers that were made by the requested
	// account address
	AccountSummary BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummary `json:"account_summary,required"`
	Status         BitcoinTransactionScanResponseSimulationBitcoinSimulationResultStatus         `json:"status,required"`
	// Details of addresses involved in the transaction
	AddressDetails []BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAddressDetail `json:"address_details"`
	// Mapping between the address of an account to the assets diff during the
	// transaction
	AssetsDiffs map[string][]BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiff `json:"assets_diffs"`
	JSON        bitcoinTransactionScanResponseSimulationBitcoinSimulationResultJSON                    `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultJSON contains the
// JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResult]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultJSON struct {
	AccountSummary apijson.Field
	Status         apijson.Field
	AddressDetails apijson.Field
	AssetsDiffs    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultJSON) RawJSON() string {
	return r.raw
}

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResult) implementsBitcoinTransactionScanResponseSimulation() {
}

// Summary of the actions and asset transfers that were made by the requested
// account address
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummary struct {
	// Total USD diff for the requested account address
	TotalUsdDiff BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryTotalUsdDiff `json:"total_usd_diff,required"`
	// Assets diffs of the requested account address
	AccountAssetsDiffs []BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiff `json:"account_assets_diffs"`
	JSON               bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryJSON                `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummary]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryJSON struct {
	TotalUsdDiff       apijson.Field
	AccountAssetsDiffs apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryJSON) RawJSON() string {
	return r.raw
}

// Total USD diff for the requested account address
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryTotalUsdDiff struct {
	// Total incoming USD transfers
	In float64 `json:"in,required"`
	// Total outgoing USD transfers
	Out float64 `json:"out,required"`
	// Total USD transfers
	Total float64                                                                                       `json:"total"`
	JSON  bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryTotalUsdDiffJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryTotalUsdDiffJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryTotalUsdDiff]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiff struct {
	// This field can have the runtime type of
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAsset],
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffAsset],
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffAsset].
	Asset interface{} `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffIn],
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffIn],
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffIn].
	In interface{} `json:"in"`
	// This field can have the runtime type of
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffOut],
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffOut],
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffOut].
	Out   interface{}                                                                                        `json:"out"`
	JSON  bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffJSON `json:"-"`
	union BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsUnion
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiff]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiff],
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiff],
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiff].
func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiff) AsUnion() BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiff],
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiff]
// or
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiff].
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsUnion interface {
	implementsBitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiff{}),
		},
	)
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiff struct {
	Asset BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffOut  `json:"out,nullable"`
	JSON bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiff]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiff) implementsBitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiff() {
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAsset struct {
	// URL of the asset's logo
	LogoURL string `json:"logo_url,required,nullable"`
	// Decimals of the asset
	Decimals BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetDecimals `json:"decimals"`
	// Name of the asset
	Name BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetName `json:"name"`
	// Symbol of the asset
	Symbol BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetSymbol `json:"symbol"`
	// Type of the asset (`NATIVE`)
	Type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetType `json:"type"`
	JSON bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAsset]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetJSON struct {
	LogoURL     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Decimals of the asset
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetDecimals int64

const (
	BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetDecimals8 BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetDecimals = 8
)

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetDecimals) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetDecimals8:
		return true
	}
	return false
}

// Name of the asset
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetName string

const (
	BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetNameBitcoin BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetName = "Bitcoin"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetName) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetNameBitcoin:
		return true
	}
	return false
}

// Symbol of the asset
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetSymbol string

const (
	BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetSymbolBtc BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetSymbol = "BTC"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetSymbol) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetSymbolBtc:
		return true
	}
	return false
}

// Type of the asset (`NATIVE`)
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetType string

const (
	BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetTypeNative BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetType = "NATIVE"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffAssetTypeNative:
		return true
	}
	return false
}

// Details of the incoming transfer
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice string                                                                                                                      `json:"usd_price,nullable"`
	JSON     bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffInJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffInJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffIn]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice string                                                                                                                       `json:"usd_price,nullable"`
	JSON     bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffOutJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffOutJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffOut]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinNativeAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiff struct {
	Asset BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffOut  `json:"out,nullable"`
	JSON bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiff]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiff) implementsBitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiff() {
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffAsset struct {
	// The Inscription ID
	ID string `json:"id,required"`
	// Inscription's display name
	Name string `json:"name,required"`
	// The Inscription sat
	Sat int64 `json:"sat,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`INSCRIPTION`)
	Type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffAssetType `json:"type"`
	JSON bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffAssetJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffAsset]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffAssetJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Sat         apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`INSCRIPTION`)
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffAssetType string

const (
	BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffAssetTypeInscription BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffAssetType = "INSCRIPTION"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffAssetType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffAssetTypeInscription:
		return true
	}
	return false
}

// Details of the incoming transfer
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffIn struct {
	// Inscription ID of the transfer
	InscriptionID string `json:"inscription_id,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice string                                                                                                                           `json:"usd_price,nullable"`
	JSON     bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffInJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffInJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffIn]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffInJSON struct {
	InscriptionID apijson.Field
	Summary       apijson.Field
	UsdPrice      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffOut struct {
	// Inscription ID of the transfer
	InscriptionID string `json:"inscription_id,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice string                                                                                                                            `json:"usd_price,nullable"`
	JSON     bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffOutJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffOutJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffOut]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffOutJSON struct {
	InscriptionID apijson.Field
	Summary       apijson.Field
	UsdPrice      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinInscriptionAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiff struct {
	Asset BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffOut  `json:"out,nullable"`
	JSON bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiff]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiff) implementsBitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiff() {
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffAsset struct {
	// The Rune ID
	ID string `json:"id,required"`
	// Decimals of the asset
	Decimals int64 `json:"decimals,required"`
	// The Rune name
	Name string `json:"name,required"`
	// The Rune spaced name
	SpacedName string `json:"spaced_name,required"`
	// The Rune's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`RUNE`)
	Type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffAssetType `json:"type"`
	JSON bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffAssetJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffAsset]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffAssetJSON struct {
	ID          apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	SpacedName  apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`RUNE`)
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffAssetType string

const (
	BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffAssetTypeRune BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffAssetType = "RUNE"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffAssetType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffAssetTypeRune:
		return true
	}
	return false
}

// Details of the incoming transfer
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice string                                                                                                                     `json:"usd_price,nullable"`
	JSON     bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffInJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffInJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffIn]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice string                                                                                                                      `json:"usd_price,nullable"`
	JSON     bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffOutJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffOutJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffOut]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAccountSummaryAccountAssetsDiffsBitcoinRunesAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultStatus string

const (
	BitcoinTransactionScanResponseSimulationBitcoinSimulationResultStatusSuccess BitcoinTransactionScanResponseSimulationBitcoinSimulationResultStatus = "Success"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResultStatus) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinSimulationResultStatusSuccess:
		return true
	}
	return false
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAddressDetail struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// Description of the account
	Description string                                                                           `json:"description,nullable"`
	JSON        bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAddressDetailJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAddressDetailJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAddressDetail]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAddressDetailJSON struct {
	AccountAddress apijson.Field
	Description    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAddressDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAddressDetailJSON) RawJSON() string {
	return r.raw
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiff struct {
	// This field can have the runtime type of
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAsset],
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffAsset],
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffAsset].
	Asset interface{} `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffIn],
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffIn],
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffIn].
	In interface{} `json:"in"`
	// This field can have the runtime type of
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffOut],
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffOut],
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffOut].
	Out   interface{}                                                                   `json:"out"`
	JSON  bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffJSON `json:"-"`
	union BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsUnion
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiff]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiff],
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiff],
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiff].
func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiff) AsUnion() BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiff],
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiff]
// or
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiff].
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsUnion interface {
	implementsBitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiff{}),
		},
	)
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiff struct {
	Asset BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffOut  `json:"out,nullable"`
	JSON bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiff]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiff) implementsBitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiff() {
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAsset struct {
	// URL of the asset's logo
	LogoURL string `json:"logo_url,required,nullable"`
	// Decimals of the asset
	Decimals BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetDecimals `json:"decimals"`
	// Name of the asset
	Name BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetName `json:"name"`
	// Symbol of the asset
	Symbol BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetSymbol `json:"symbol"`
	// Type of the asset (`NATIVE`)
	Type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetType `json:"type"`
	JSON bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAsset]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetJSON struct {
	LogoURL     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Decimals of the asset
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetDecimals int64

const (
	BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetDecimals8 BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetDecimals = 8
)

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetDecimals) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetDecimals8:
		return true
	}
	return false
}

// Name of the asset
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetName string

const (
	BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetNameBitcoin BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetName = "Bitcoin"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetName) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetNameBitcoin:
		return true
	}
	return false
}

// Symbol of the asset
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetSymbol string

const (
	BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetSymbolBtc BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetSymbol = "BTC"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetSymbol) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetSymbolBtc:
		return true
	}
	return false
}

// Type of the asset (`NATIVE`)
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetType string

const (
	BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetTypeNative BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetType = "NATIVE"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffAssetTypeNative:
		return true
	}
	return false
}

// Details of the incoming transfer
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice string                                                                                                 `json:"usd_price,nullable"`
	JSON     bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffInJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffInJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffIn]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice string                                                                                                  `json:"usd_price,nullable"`
	JSON     bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffOutJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffOutJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffOut]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinNativeAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiff struct {
	Asset BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffOut  `json:"out,nullable"`
	JSON bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiff]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiff) implementsBitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiff() {
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffAsset struct {
	// The Inscription ID
	ID string `json:"id,required"`
	// Inscription's display name
	Name string `json:"name,required"`
	// The Inscription sat
	Sat int64 `json:"sat,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`INSCRIPTION`)
	Type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffAssetType `json:"type"`
	JSON bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffAssetJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffAsset]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffAssetJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Sat         apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`INSCRIPTION`)
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffAssetType string

const (
	BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffAssetTypeInscription BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffAssetType = "INSCRIPTION"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffAssetType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffAssetTypeInscription:
		return true
	}
	return false
}

// Details of the incoming transfer
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffIn struct {
	// Inscription ID of the transfer
	InscriptionID string `json:"inscription_id,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice string                                                                                                      `json:"usd_price,nullable"`
	JSON     bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffInJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffInJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffIn]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffInJSON struct {
	InscriptionID apijson.Field
	Summary       apijson.Field
	UsdPrice      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffOut struct {
	// Inscription ID of the transfer
	InscriptionID string `json:"inscription_id,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice string                                                                                                       `json:"usd_price,nullable"`
	JSON     bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffOutJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffOutJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffOut]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffOutJSON struct {
	InscriptionID apijson.Field
	Summary       apijson.Field
	UsdPrice      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinInscriptionAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiff struct {
	Asset BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffOut  `json:"out,nullable"`
	JSON bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiff]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiff) implementsBitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiff() {
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffAsset struct {
	// The Rune ID
	ID string `json:"id,required"`
	// Decimals of the asset
	Decimals int64 `json:"decimals,required"`
	// The Rune name
	Name string `json:"name,required"`
	// The Rune spaced name
	SpacedName string `json:"spaced_name,required"`
	// The Rune's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`RUNE`)
	Type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffAssetType `json:"type"`
	JSON bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffAssetJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffAsset]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffAssetJSON struct {
	ID          apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	SpacedName  apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`RUNE`)
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffAssetType string

const (
	BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffAssetTypeRune BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffAssetType = "RUNE"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffAssetType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffAssetTypeRune:
		return true
	}
	return false
}

// Details of the incoming transfer
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice string                                                                                                `json:"usd_price,nullable"`
	JSON     bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffInJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffInJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffIn]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice string                                                                                                 `json:"usd_price,nullable"`
	JSON     bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffOutJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffOutJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffOut]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResultAssetsDiffsBitcoinRunesAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchema struct {
	// Error message
	Error  string                                                                     `json:"error,required"`
	Status BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchemaStatus `json:"status,required"`
	JSON   bitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchemaJSON   `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchemaJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchema]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchemaJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchemaJSON) RawJSON() string {
	return r.raw
}

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchema) implementsBitcoinTransactionScanResponseSimulation() {
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchemaStatus string

const (
	BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchemaStatusError BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchemaStatus = "Error"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchemaStatus) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchemaStatusError:
		return true
	}
	return false
}

type BitcoinTransactionScanResponseSimulationStatus string

const (
	BitcoinTransactionScanResponseSimulationStatusSuccess BitcoinTransactionScanResponseSimulationStatus = "Success"
	BitcoinTransactionScanResponseSimulationStatusError   BitcoinTransactionScanResponseSimulationStatus = "Error"
)

func (r BitcoinTransactionScanResponseSimulationStatus) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationStatusSuccess, BitcoinTransactionScanResponseSimulationStatusError:
		return true
	}
	return false
}

// Validation result; Only present if validation option is included in the request
type BitcoinTransactionScanResponseValidation struct {
	Status BitcoinTransactionScanResponseValidationStatus `json:"status,required"`
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification"`
	// A textual description about the validation result
	Description string `json:"description"`
	// Error message
	Error string `json:"error"`
	// This field can have the runtime type of
	// [[]BitcoinTransactionScanResponseValidationBitcoinValidationResultFeature].
	Features interface{} `json:"features"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason string `json:"reason"`
	// Verdict of the validation
	ResultType BitcoinTransactionScanResponseValidationResultType `json:"result_type"`
	JSON       bitcoinTransactionScanResponseValidationJSON       `json:"-"`
	union      BitcoinTransactionScanResponseValidationUnion
}

// bitcoinTransactionScanResponseValidationJSON contains the JSON metadata for the
// struct [BitcoinTransactionScanResponseValidation]
type bitcoinTransactionScanResponseValidationJSON struct {
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

func (r bitcoinTransactionScanResponseValidationJSON) RawJSON() string {
	return r.raw
}

func (r *BitcoinTransactionScanResponseValidation) UnmarshalJSON(data []byte) (err error) {
	*r = BitcoinTransactionScanResponseValidation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BitcoinTransactionScanResponseValidationUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [BitcoinTransactionScanResponseValidationBitcoinValidationResult],
// [BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchema].
func (r BitcoinTransactionScanResponseValidation) AsUnion() BitcoinTransactionScanResponseValidationUnion {
	return r.union
}

// Validation result; Only present if validation option is included in the request
//
// Union satisfied by
// [BitcoinTransactionScanResponseValidationBitcoinValidationResult] or
// [BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchema].
type BitcoinTransactionScanResponseValidationUnion interface {
	implementsBitcoinTransactionScanResponseValidation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BitcoinTransactionScanResponseValidationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BitcoinTransactionScanResponseValidationBitcoinValidationResult{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchema{}),
		},
	)
}

type BitcoinTransactionScanResponseValidationBitcoinValidationResult struct {
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification,required"`
	// A textual description about the validation result
	Description string                                                                   `json:"description,required"`
	Features    []BitcoinTransactionScanResponseValidationBitcoinValidationResultFeature `json:"features,required"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason string `json:"reason,required"`
	// Verdict of the validation
	ResultType BitcoinTransactionScanResponseValidationBitcoinValidationResultResultType `json:"result_type,required"`
	Status     BitcoinTransactionScanResponseValidationBitcoinValidationResultStatus     `json:"status,required"`
	JSON       bitcoinTransactionScanResponseValidationBitcoinValidationResultJSON       `json:"-"`
}

// bitcoinTransactionScanResponseValidationBitcoinValidationResultJSON contains the
// JSON metadata for the struct
// [BitcoinTransactionScanResponseValidationBitcoinValidationResult]
type bitcoinTransactionScanResponseValidationBitcoinValidationResultJSON struct {
	Classification apijson.Field
	Description    apijson.Field
	Features       apijson.Field
	Reason         apijson.Field
	ResultType     apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseValidationBitcoinValidationResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseValidationBitcoinValidationResultJSON) RawJSON() string {
	return r.raw
}

func (r BitcoinTransactionScanResponseValidationBitcoinValidationResult) implementsBitcoinTransactionScanResponseValidation() {
}

type BitcoinTransactionScanResponseValidationBitcoinValidationResultFeature struct {
	// Address the feature refers to
	Address string `json:"address,required"`
	// Textual description
	Description string `json:"description,required"`
	FeatureID   string `json:"feature_id,required"`
	// Feature Classification
	Type BitcoinTransactionScanResponseValidationBitcoinValidationResultFeaturesType `json:"type,required"`
	JSON bitcoinTransactionScanResponseValidationBitcoinValidationResultFeatureJSON  `json:"-"`
}

// bitcoinTransactionScanResponseValidationBitcoinValidationResultFeatureJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseValidationBitcoinValidationResultFeature]
type bitcoinTransactionScanResponseValidationBitcoinValidationResultFeatureJSON struct {
	Address     apijson.Field
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseValidationBitcoinValidationResultFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseValidationBitcoinValidationResultFeatureJSON) RawJSON() string {
	return r.raw
}

// Feature Classification
type BitcoinTransactionScanResponseValidationBitcoinValidationResultFeaturesType string

const (
	BitcoinTransactionScanResponseValidationBitcoinValidationResultFeaturesTypeBenign    BitcoinTransactionScanResponseValidationBitcoinValidationResultFeaturesType = "Benign"
	BitcoinTransactionScanResponseValidationBitcoinValidationResultFeaturesTypeWarning   BitcoinTransactionScanResponseValidationBitcoinValidationResultFeaturesType = "Warning"
	BitcoinTransactionScanResponseValidationBitcoinValidationResultFeaturesTypeMalicious BitcoinTransactionScanResponseValidationBitcoinValidationResultFeaturesType = "Malicious"
	BitcoinTransactionScanResponseValidationBitcoinValidationResultFeaturesTypeInfo      BitcoinTransactionScanResponseValidationBitcoinValidationResultFeaturesType = "Info"
)

func (r BitcoinTransactionScanResponseValidationBitcoinValidationResultFeaturesType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseValidationBitcoinValidationResultFeaturesTypeBenign, BitcoinTransactionScanResponseValidationBitcoinValidationResultFeaturesTypeWarning, BitcoinTransactionScanResponseValidationBitcoinValidationResultFeaturesTypeMalicious, BitcoinTransactionScanResponseValidationBitcoinValidationResultFeaturesTypeInfo:
		return true
	}
	return false
}

// Verdict of the validation
type BitcoinTransactionScanResponseValidationBitcoinValidationResultResultType string

const (
	BitcoinTransactionScanResponseValidationBitcoinValidationResultResultTypeBenign    BitcoinTransactionScanResponseValidationBitcoinValidationResultResultType = "Benign"
	BitcoinTransactionScanResponseValidationBitcoinValidationResultResultTypeWarning   BitcoinTransactionScanResponseValidationBitcoinValidationResultResultType = "Warning"
	BitcoinTransactionScanResponseValidationBitcoinValidationResultResultTypeMalicious BitcoinTransactionScanResponseValidationBitcoinValidationResultResultType = "Malicious"
)

func (r BitcoinTransactionScanResponseValidationBitcoinValidationResultResultType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseValidationBitcoinValidationResultResultTypeBenign, BitcoinTransactionScanResponseValidationBitcoinValidationResultResultTypeWarning, BitcoinTransactionScanResponseValidationBitcoinValidationResultResultTypeMalicious:
		return true
	}
	return false
}

type BitcoinTransactionScanResponseValidationBitcoinValidationResultStatus string

const (
	BitcoinTransactionScanResponseValidationBitcoinValidationResultStatusSuccess BitcoinTransactionScanResponseValidationBitcoinValidationResultStatus = "Success"
)

func (r BitcoinTransactionScanResponseValidationBitcoinValidationResultStatus) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseValidationBitcoinValidationResultStatusSuccess:
		return true
	}
	return false
}

type BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchema struct {
	// Error message
	Error  string                                                                     `json:"error,required"`
	Status BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchemaStatus `json:"status,required"`
	JSON   bitcoinTransactionScanResponseValidationBitcoinValidationErrorSchemaJSON   `json:"-"`
}

// bitcoinTransactionScanResponseValidationBitcoinValidationErrorSchemaJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchema]
type bitcoinTransactionScanResponseValidationBitcoinValidationErrorSchemaJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseValidationBitcoinValidationErrorSchemaJSON) RawJSON() string {
	return r.raw
}

func (r BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchema) implementsBitcoinTransactionScanResponseValidation() {
}

type BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchemaStatus string

const (
	BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchemaStatusError BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchemaStatus = "Error"
)

func (r BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchemaStatus) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchemaStatusError:
		return true
	}
	return false
}

type BitcoinTransactionScanResponseValidationStatus string

const (
	BitcoinTransactionScanResponseValidationStatusSuccess BitcoinTransactionScanResponseValidationStatus = "Success"
	BitcoinTransactionScanResponseValidationStatusError   BitcoinTransactionScanResponseValidationStatus = "Error"
)

func (r BitcoinTransactionScanResponseValidationStatus) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseValidationStatusSuccess, BitcoinTransactionScanResponseValidationStatusError:
		return true
	}
	return false
}

// Verdict of the validation
type BitcoinTransactionScanResponseValidationResultType string

const (
	BitcoinTransactionScanResponseValidationResultTypeBenign    BitcoinTransactionScanResponseValidationResultType = "Benign"
	BitcoinTransactionScanResponseValidationResultTypeWarning   BitcoinTransactionScanResponseValidationResultType = "Warning"
	BitcoinTransactionScanResponseValidationResultTypeMalicious BitcoinTransactionScanResponseValidationResultType = "Malicious"
)

func (r BitcoinTransactionScanResponseValidationResultType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseValidationResultTypeBenign, BitcoinTransactionScanResponseValidationResultTypeWarning, BitcoinTransactionScanResponseValidationResultTypeMalicious:
		return true
	}
	return false
}
