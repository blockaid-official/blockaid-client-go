// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"reflect"
	"time"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/param"
	"github.com/blockaid-official/blockaid-client-go/option"
	"github.com/blockaid-official/blockaid-client-go/shared"
	"github.com/tidwall/gjson"
)

// EvmService contains methods and other services that help with interacting with
// the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvmService] method instead.
type EvmService struct {
	Options             []option.RequestOption
	JsonRpc             *EvmJsonRpcService
	Transaction         *EvmTransactionService
	TransactionBulk     *EvmTransactionBulkService
	TransactionRaw      *EvmTransactionRawService
	UserOperation       *EvmUserOperationService
	PostTransaction     *EvmPostTransactionService
	PostTransactionBulk *EvmPostTransactionBulkService
}

// NewEvmService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEvmService(opts ...option.RequestOption) (r *EvmService) {
	r = &EvmService{}
	r.Options = opts
	r.JsonRpc = NewEvmJsonRpcService(opts...)
	r.Transaction = NewEvmTransactionService(opts...)
	r.TransactionBulk = NewEvmTransactionBulkService(opts...)
	r.TransactionRaw = NewEvmTransactionRawService(opts...)
	r.UserOperation = NewEvmUserOperationService(opts...)
	r.PostTransaction = NewEvmPostTransactionService(opts...)
	r.PostTransactionBulk = NewEvmPostTransactionBulkService(opts...)
	return
}

type Erc1155Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string          `json:"usd_price"`
	JSON     erc1155DiffJSON `json:"-"`
}

// erc1155DiffJSON contains the JSON metadata for the struct [Erc1155Diff]
type erc1155DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *Erc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r erc1155DiffJSON) RawJSON() string {
	return r.raw
}

func (r Erc1155Diff) implementsErc1155ExposureExposure() {}

func (r Erc1155Diff) implementsErc20ExposureExposure() {}

func (r Erc1155Diff) implementsErc721ExposureExposure() {}

type Erc1155Exposure struct {
	Exposure []Erc1155ExposureExposure `json:"exposure,required"`
	// boolean indicates whether an is_approved_for_all function was used (missing in
	// case of ERC20 / ERC1155)
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// user friendly description of the approval
	Summary string              `json:"summary"`
	JSON    erc1155ExposureJSON `json:"-"`
}

// erc1155ExposureJSON contains the JSON metadata for the struct [Erc1155Exposure]
type erc1155ExposureJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Erc1155Exposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r erc1155ExposureJSON) RawJSON() string {
	return r.raw
}

type Erc1155ExposureExposure struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// id of the token
	TokenID string `json:"token_id"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value before divided by decimal, that was transferred from this address
	Value string                      `json:"value"`
	JSON  erc1155ExposureExposureJSON `json:"-"`
	union Erc1155ExposureExposureUnion
}

// erc1155ExposureExposureJSON contains the JSON metadata for the struct
// [Erc1155ExposureExposure]
type erc1155ExposureExposureJSON struct {
	ArbitraryCollectionToken apijson.Field
	LogoURL                  apijson.Field
	RawValue                 apijson.Field
	Summary                  apijson.Field
	TokenID                  apijson.Field
	UsdPrice                 apijson.Field
	Value                    apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r erc1155ExposureExposureJSON) RawJSON() string {
	return r.raw
}

func (r *Erc1155ExposureExposure) UnmarshalJSON(data []byte) (err error) {
	*r = Erc1155ExposureExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [Erc1155ExposureExposureUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [Erc1155Diff], [Erc721Diff],
// [Erc20Diff], [NativeDiff].
func (r Erc1155ExposureExposure) AsUnion() Erc1155ExposureExposureUnion {
	return r.union
}

// Union satisfied by [Erc1155Diff], [Erc721Diff], [Erc20Diff] or [NativeDiff].
type Erc1155ExposureExposureUnion interface {
	implementsErc1155ExposureExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*Erc1155ExposureExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NativeDiff{}),
		},
	)
}

type Erc1155TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type Erc1155TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                  `json:"symbol"`
	JSON   erc1155TokenDetailsJSON `json:"-"`
}

// erc1155TokenDetailsJSON contains the JSON metadata for the struct
// [Erc1155TokenDetails]
type erc1155TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Erc1155TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r erc1155TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r Erc1155TokenDetails) implementsTransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAsset() {
}

func (r Erc1155TokenDetails) implementsTransactionSimulationAccountSummaryExposuresErc1155AddressExposureAsset() {
}

func (r Erc1155TokenDetails) implementsTransactionSimulationAccountSummaryTracesErc1155AssetTraceAsset() {
}

func (r Erc1155TokenDetails) implementsTransactionSimulationAssetsDiffsErc1155AddressAssetDiffAsset() {
}

func (r Erc1155TokenDetails) implementsTransactionSimulationExposuresErc1155AddressExposureAsset() {}

// asset type.
type Erc1155TokenDetailsType string

const (
	Erc1155TokenDetailsTypeErc1155 Erc1155TokenDetailsType = "ERC1155"
)

func (r Erc1155TokenDetailsType) IsKnown() bool {
	switch r {
	case Erc1155TokenDetailsTypeErc1155:
		return true
	}
	return false
}

type Erc20Diff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string        `json:"value"`
	JSON  erc20DiffJSON `json:"-"`
}

// erc20DiffJSON contains the JSON metadata for the struct [Erc20Diff]
type erc20DiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Erc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r erc20DiffJSON) RawJSON() string {
	return r.raw
}

func (r Erc20Diff) implementsErc1155ExposureExposure() {}

func (r Erc20Diff) implementsErc20ExposureExposure() {}

func (r Erc20Diff) implementsErc721ExposureExposure() {}

type Erc20Exposure struct {
	// the amount that was asked in the approval request for this spender from the
	// current address and asset
	Approval string                  `json:"approval,required"`
	Exposure []Erc20ExposureExposure `json:"exposure,required"`
	// the usd price of the approval amount
	ApprovalUsdPrice string `json:"approval_usd_price"`
	// the expiration time of the permit2 protocol
	Expiration time.Time `json:"expiration" format:"date-time"`
	// user friendly description of the approval
	Summary string            `json:"summary"`
	JSON    erc20ExposureJSON `json:"-"`
}

// erc20ExposureJSON contains the JSON metadata for the struct [Erc20Exposure]
type erc20ExposureJSON struct {
	Approval         apijson.Field
	Exposure         apijson.Field
	ApprovalUsdPrice apijson.Field
	Expiration       apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Erc20Exposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r erc20ExposureJSON) RawJSON() string {
	return r.raw
}

type Erc20ExposureExposure struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// id of the token
	TokenID string `json:"token_id"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value before divided by decimal, that was transferred from this address
	Value string                    `json:"value"`
	JSON  erc20ExposureExposureJSON `json:"-"`
	union Erc20ExposureExposureUnion
}

// erc20ExposureExposureJSON contains the JSON metadata for the struct
// [Erc20ExposureExposure]
type erc20ExposureExposureJSON struct {
	ArbitraryCollectionToken apijson.Field
	LogoURL                  apijson.Field
	RawValue                 apijson.Field
	Summary                  apijson.Field
	TokenID                  apijson.Field
	UsdPrice                 apijson.Field
	Value                    apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r erc20ExposureExposureJSON) RawJSON() string {
	return r.raw
}

func (r *Erc20ExposureExposure) UnmarshalJSON(data []byte) (err error) {
	*r = Erc20ExposureExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [Erc20ExposureExposureUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [Erc1155Diff], [Erc721Diff],
// [Erc20Diff], [NativeDiff].
func (r Erc20ExposureExposure) AsUnion() Erc20ExposureExposureUnion {
	return r.union
}

// Union satisfied by [Erc1155Diff], [Erc721Diff], [Erc20Diff] or [NativeDiff].
type Erc20ExposureExposureUnion interface {
	implementsErc20ExposureExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*Erc20ExposureExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NativeDiff{}),
		},
	)
}

type Erc20TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset's decimals
	Decimals int64 `json:"decimals,required"`
	// asset type.
	Type Erc20TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                `json:"symbol"`
	JSON   erc20TokenDetailsJSON `json:"-"`
}

// erc20TokenDetailsJSON contains the JSON metadata for the struct
// [Erc20TokenDetails]
type erc20TokenDetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Erc20TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r erc20TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r Erc20TokenDetails) implementsTransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAsset() {
}

func (r Erc20TokenDetails) implementsTransactionSimulationAccountSummaryExposuresErc20AddressExposureAsset() {
}

func (r Erc20TokenDetails) implementsTransactionSimulationAccountSummaryTracesErc20AssetTraceAsset() {
}

func (r Erc20TokenDetails) implementsTransactionSimulationAssetsDiffsErc20AddressAssetDiffAsset() {}

func (r Erc20TokenDetails) implementsTransactionSimulationExposuresErc20AddressExposureAsset() {}

// asset type.
type Erc20TokenDetailsType string

const (
	Erc20TokenDetailsTypeErc20 Erc20TokenDetailsType = "ERC20"
)

func (r Erc20TokenDetailsType) IsKnown() bool {
	switch r {
	case Erc20TokenDetailsTypeErc20:
		return true
	}
	return false
}

type Erc721Diff struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token,required"`
	// id of the token
	TokenID string `json:"token_id,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string         `json:"usd_price"`
	JSON     erc721DiffJSON `json:"-"`
}

// erc721DiffJSON contains the JSON metadata for the struct [Erc721Diff]
type erc721DiffJSON struct {
	ArbitraryCollectionToken apijson.Field
	TokenID                  apijson.Field
	LogoURL                  apijson.Field
	Summary                  apijson.Field
	UsdPrice                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *Erc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r erc721DiffJSON) RawJSON() string {
	return r.raw
}

func (r Erc721Diff) implementsErc1155ExposureExposure() {}

func (r Erc721Diff) implementsErc20ExposureExposure() {}

func (r Erc721Diff) implementsErc721ExposureExposure() {}

type Erc721Exposure struct {
	Exposure []Erc721ExposureExposure `json:"exposure,required"`
	// boolean indicates whether an is_approved_for_all function was used (missing in
	// case of ERC20 / ERC1155)
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// user friendly description of the approval
	Summary string             `json:"summary"`
	JSON    erc721ExposureJSON `json:"-"`
}

// erc721ExposureJSON contains the JSON metadata for the struct [Erc721Exposure]
type erc721ExposureJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Erc721Exposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r erc721ExposureJSON) RawJSON() string {
	return r.raw
}

type Erc721ExposureExposure struct {
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// id of the token
	TokenID string `json:"token_id"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value before divided by decimal, that was transferred from this address
	Value string                     `json:"value"`
	JSON  erc721ExposureExposureJSON `json:"-"`
	union Erc721ExposureExposureUnion
}

// erc721ExposureExposureJSON contains the JSON metadata for the struct
// [Erc721ExposureExposure]
type erc721ExposureExposureJSON struct {
	ArbitraryCollectionToken apijson.Field
	LogoURL                  apijson.Field
	RawValue                 apijson.Field
	Summary                  apijson.Field
	TokenID                  apijson.Field
	UsdPrice                 apijson.Field
	Value                    apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r erc721ExposureExposureJSON) RawJSON() string {
	return r.raw
}

func (r *Erc721ExposureExposure) UnmarshalJSON(data []byte) (err error) {
	*r = Erc721ExposureExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [Erc721ExposureExposureUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [Erc1155Diff], [Erc721Diff],
// [Erc20Diff], [NativeDiff].
func (r Erc721ExposureExposure) AsUnion() Erc721ExposureExposureUnion {
	return r.union
}

// Union satisfied by [Erc1155Diff], [Erc721Diff], [Erc20Diff] or [NativeDiff].
type Erc721ExposureExposureUnion interface {
	implementsErc721ExposureExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*Erc721ExposureExposureUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc1155Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc721Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc20Diff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NativeDiff{}),
		},
	)
}

type Erc721TokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type Erc721TokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                 `json:"symbol"`
	JSON   erc721TokenDetailsJSON `json:"-"`
}

// erc721TokenDetailsJSON contains the JSON metadata for the struct
// [Erc721TokenDetails]
type erc721TokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Erc721TokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r erc721TokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r Erc721TokenDetails) implementsTransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAsset() {
}

func (r Erc721TokenDetails) implementsTransactionSimulationAccountSummaryExposuresErc721AddressExposureAsset() {
}

func (r Erc721TokenDetails) implementsTransactionSimulationAccountSummaryTracesErc721AssetTraceAsset() {
}

func (r Erc721TokenDetails) implementsTransactionSimulationAssetsDiffsErc721AddressAssetDiffAsset() {}

func (r Erc721TokenDetails) implementsTransactionSimulationExposuresErc721AddressExposureAsset() {}

// asset type.
type Erc721TokenDetailsType string

const (
	Erc721TokenDetailsTypeErc721 Erc721TokenDetailsType = "ERC721"
)

func (r Erc721TokenDetailsType) IsKnown() bool {
	switch r {
	case Erc721TokenDetailsTypeErc721:
		return true
	}
	return false
}

type MetadataParam struct {
	// cross reference transaction against the domain.
	Domain param.Field[string] `json:"domain,required"`
}

func (r MetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NativeAssetDetails struct {
	ChainID   int64  `json:"chain_id,required"`
	ChainName string `json:"chain_name,required"`
	Decimals  int64  `json:"decimals,required"`
	LogoURL   string `json:"logo_url,required"`
	// asset type.
	Type NativeAssetDetailsType `json:"type,required"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                 `json:"symbol"`
	JSON   nativeAssetDetailsJSON `json:"-"`
}

// nativeAssetDetailsJSON contains the JSON metadata for the struct
// [NativeAssetDetails]
type nativeAssetDetailsJSON struct {
	ChainID     apijson.Field
	ChainName   apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NativeAssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nativeAssetDetailsJSON) RawJSON() string {
	return r.raw
}

// asset type.
type NativeAssetDetailsType string

const (
	NativeAssetDetailsTypeNative NativeAssetDetailsType = "NATIVE"
)

func (r NativeAssetDetailsType) IsKnown() bool {
	switch r {
	case NativeAssetDetailsTypeNative:
		return true
	}
	return false
}

type NativeDiff struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string         `json:"value"`
	JSON  nativeDiffJSON `json:"-"`
}

// nativeDiffJSON contains the JSON metadata for the struct [NativeDiff]
type nativeDiffJSON struct {
	RawValue    apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NativeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nativeDiffJSON) RawJSON() string {
	return r.raw
}

func (r NativeDiff) implementsErc1155ExposureExposure() {}

func (r NativeDiff) implementsErc20ExposureExposure() {}

func (r NativeDiff) implementsErc721ExposureExposure() {}

type NonercTokenDetails struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type NonercTokenDetailsType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                 `json:"symbol"`
	JSON   nonercTokenDetailsJSON `json:"-"`
}

// nonercTokenDetailsJSON contains the JSON metadata for the struct
// [NonercTokenDetails]
type nonercTokenDetailsJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NonercTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nonercTokenDetailsJSON) RawJSON() string {
	return r.raw
}

func (r NonercTokenDetails) implementsTransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAsset() {
}

func (r NonercTokenDetails) implementsTransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAsset() {
}

func (r NonercTokenDetails) implementsTransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAsset() {
}

func (r NonercTokenDetails) implementsTransactionSimulationAccountSummaryExposuresErc20AddressExposureAsset() {
}

func (r NonercTokenDetails) implementsTransactionSimulationAccountSummaryExposuresErc721AddressExposureAsset() {
}

func (r NonercTokenDetails) implementsTransactionSimulationAccountSummaryExposuresErc1155AddressExposureAsset() {
}

func (r NonercTokenDetails) implementsTransactionSimulationAccountSummaryTracesErc20AssetTraceAsset() {
}

func (r NonercTokenDetails) implementsTransactionSimulationAccountSummaryTracesErc721AssetTraceAsset() {
}

func (r NonercTokenDetails) implementsTransactionSimulationAccountSummaryTracesErc1155AssetTraceAsset() {
}

func (r NonercTokenDetails) implementsTransactionSimulationAssetsDiffsErc20AddressAssetDiffAsset() {}

func (r NonercTokenDetails) implementsTransactionSimulationAssetsDiffsErc721AddressAssetDiffAsset() {}

func (r NonercTokenDetails) implementsTransactionSimulationAssetsDiffsErc1155AddressAssetDiffAsset() {
}

func (r NonercTokenDetails) implementsTransactionSimulationExposuresErc20AddressExposureAsset() {}

func (r NonercTokenDetails) implementsTransactionSimulationExposuresErc721AddressExposureAsset() {}

func (r NonercTokenDetails) implementsTransactionSimulationExposuresErc1155AddressExposureAsset() {}

// asset type.
type NonercTokenDetailsType string

const (
	NonercTokenDetailsTypeNonerc NonercTokenDetailsType = "NONERC"
)

func (r NonercTokenDetailsType) IsKnown() bool {
	switch r {
	case NonercTokenDetailsTypeNonerc:
		return true
	}
	return false
}

// The chain name
type TokenScanSupportedChain string

const (
	TokenScanSupportedChainArbitrum    TokenScanSupportedChain = "arbitrum"
	TokenScanSupportedChainAvalanche   TokenScanSupportedChain = "avalanche"
	TokenScanSupportedChainBase        TokenScanSupportedChain = "base"
	TokenScanSupportedChainBsc         TokenScanSupportedChain = "bsc"
	TokenScanSupportedChainEthereum    TokenScanSupportedChain = "ethereum"
	TokenScanSupportedChainOptimism    TokenScanSupportedChain = "optimism"
	TokenScanSupportedChainPolygon     TokenScanSupportedChain = "polygon"
	TokenScanSupportedChainZora        TokenScanSupportedChain = "zora"
	TokenScanSupportedChainSolana      TokenScanSupportedChain = "solana"
	TokenScanSupportedChainStarknet    TokenScanSupportedChain = "starknet"
	TokenScanSupportedChainStellar     TokenScanSupportedChain = "stellar"
	TokenScanSupportedChainLinea       TokenScanSupportedChain = "linea"
	TokenScanSupportedChainBlast       TokenScanSupportedChain = "blast"
	TokenScanSupportedChainZksync      TokenScanSupportedChain = "zksync"
	TokenScanSupportedChainScroll      TokenScanSupportedChain = "scroll"
	TokenScanSupportedChainDegen       TokenScanSupportedChain = "degen"
	TokenScanSupportedChainAbstract    TokenScanSupportedChain = "abstract"
	TokenScanSupportedChainSoneium     TokenScanSupportedChain = "soneium"
	TokenScanSupportedChainInk         TokenScanSupportedChain = "ink"
	TokenScanSupportedChainZeroNetwork TokenScanSupportedChain = "zero-network"
	TokenScanSupportedChainBerachain   TokenScanSupportedChain = "berachain"
	TokenScanSupportedChainUnichain    TokenScanSupportedChain = "unichain"
)

func (r TokenScanSupportedChain) IsKnown() bool {
	switch r {
	case TokenScanSupportedChainArbitrum, TokenScanSupportedChainAvalanche, TokenScanSupportedChainBase, TokenScanSupportedChainBsc, TokenScanSupportedChainEthereum, TokenScanSupportedChainOptimism, TokenScanSupportedChainPolygon, TokenScanSupportedChainZora, TokenScanSupportedChainSolana, TokenScanSupportedChainStarknet, TokenScanSupportedChainStellar, TokenScanSupportedChainLinea, TokenScanSupportedChainBlast, TokenScanSupportedChainZksync, TokenScanSupportedChainScroll, TokenScanSupportedChainDegen, TokenScanSupportedChainAbstract, TokenScanSupportedChainSoneium, TokenScanSupportedChainInk, TokenScanSupportedChainZeroNetwork, TokenScanSupportedChainBerachain, TokenScanSupportedChainUnichain:
		return true
	}
	return false
}

type TransactionScanFeature struct {
	// Textual description
	Description string `json:"description,required"`
	// Feature name
	FeatureID string `json:"feature_id,required"`
	// An enumeration.
	Type TransactionScanFeatureType `json:"type,required"`
	// Address the feature refers to
	Address string                     `json:"address"`
	JSON    transactionScanFeatureJSON `json:"-"`
}

// transactionScanFeatureJSON contains the JSON metadata for the struct
// [TransactionScanFeature]
type transactionScanFeatureJSON struct {
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionScanFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanFeatureJSON) RawJSON() string {
	return r.raw
}

// An enumeration.
type TransactionScanFeatureType string

const (
	TransactionScanFeatureTypeMalicious TransactionScanFeatureType = "Malicious"
	TransactionScanFeatureTypeWarning   TransactionScanFeatureType = "Warning"
	TransactionScanFeatureTypeBenign    TransactionScanFeatureType = "Benign"
	TransactionScanFeatureTypeInfo      TransactionScanFeatureType = "Info"
)

func (r TransactionScanFeatureType) IsKnown() bool {
	switch r {
	case TransactionScanFeatureTypeMalicious, TransactionScanFeatureTypeWarning, TransactionScanFeatureTypeBenign, TransactionScanFeatureTypeInfo:
		return true
	}
	return false
}

type TransactionScanResponse struct {
	Block          string                               `json:"block,required"`
	Chain          string                               `json:"chain,required"`
	AccountAddress string                               `json:"account_address"`
	Events         []TransactionScanResponseEvent       `json:"events"`
	Features       interface{}                          `json:"features"`
	GasEstimation  TransactionScanResponseGasEstimation `json:"gas_estimation"`
	Simulation     TransactionScanResponseSimulation    `json:"simulation"`
	Validation     TransactionScanResponseValidation    `json:"validation"`
	JSON           transactionScanResponseJSON          `json:"-"`
}

// transactionScanResponseJSON contains the JSON metadata for the struct
// [TransactionScanResponse]
type transactionScanResponseJSON struct {
	Block          apijson.Field
	Chain          apijson.Field
	AccountAddress apijson.Field
	Events         apijson.Field
	Features       apijson.Field
	GasEstimation  apijson.Field
	Simulation     apijson.Field
	Validation     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionScanResponseEvent struct {
	Data           string                               `json:"data,required"`
	EmitterAddress string                               `json:"emitter_address,required"`
	Topics         []string                             `json:"topics,required"`
	EmitterName    string                               `json:"emitter_name"`
	Name           string                               `json:"name"`
	Params         []TransactionScanResponseEventsParam `json:"params"`
	JSON           transactionScanResponseEventJSON     `json:"-"`
}

// transactionScanResponseEventJSON contains the JSON metadata for the struct
// [TransactionScanResponseEvent]
type transactionScanResponseEventJSON struct {
	Data           apijson.Field
	EmitterAddress apijson.Field
	Topics         apijson.Field
	EmitterName    apijson.Field
	Name           apijson.Field
	Params         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionScanResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseEventJSON) RawJSON() string {
	return r.raw
}

type TransactionScanResponseEventsParam struct {
	Type         string                                        `json:"type,required"`
	Value        TransactionScanResponseEventsParamsValueUnion `json:"value,required"`
	InternalType string                                        `json:"internalType"`
	Name         string                                        `json:"name"`
	JSON         transactionScanResponseEventsParamJSON        `json:"-"`
}

// transactionScanResponseEventsParamJSON contains the JSON metadata for the struct
// [TransactionScanResponseEventsParam]
type transactionScanResponseEventsParamJSON struct {
	Type         apijson.Field
	Value        apijson.Field
	InternalType apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TransactionScanResponseEventsParam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseEventsParamJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or
// [TransactionScanResponseEventsParamsValueArray].
type TransactionScanResponseEventsParamsValueUnion interface {
	ImplementsTransactionScanResponseEventsParamsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionScanResponseEventsParamsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionScanResponseEventsParamsValueArray{}),
		},
	)
}

type TransactionScanResponseEventsParamsValueArray []interface{}

func (r TransactionScanResponseEventsParamsValueArray) ImplementsTransactionScanResponseEventsParamsValueUnion() {
}

type TransactionScanResponseGasEstimation struct {
	Status   TransactionScanResponseGasEstimationStatus `json:"status,required"`
	Error    string                                     `json:"error"`
	Estimate int64                                      `json:"estimate"`
	Used     int64                                      `json:"used"`
	JSON     transactionScanResponseGasEstimationJSON   `json:"-"`
	union    TransactionScanResponseGasEstimationUnion
}

// transactionScanResponseGasEstimationJSON contains the JSON metadata for the
// struct [TransactionScanResponseGasEstimation]
type transactionScanResponseGasEstimationJSON struct {
	Status      apijson.Field
	Error       apijson.Field
	Estimate    apijson.Field
	Used        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionScanResponseGasEstimationJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionScanResponseGasEstimation) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionScanResponseGasEstimation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TransactionScanResponseGasEstimationUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TransactionScanResponseGasEstimationTransactionScanGasEstimation],
// [TransactionScanResponseGasEstimationTransactionScanGasEstimationError].
func (r TransactionScanResponseGasEstimation) AsUnion() TransactionScanResponseGasEstimationUnion {
	return r.union
}

// Union satisfied by
// [TransactionScanResponseGasEstimationTransactionScanGasEstimation] or
// [TransactionScanResponseGasEstimationTransactionScanGasEstimationError].
type TransactionScanResponseGasEstimationUnion interface {
	implementsTransactionScanResponseGasEstimation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionScanResponseGasEstimationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionScanResponseGasEstimationTransactionScanGasEstimation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionScanResponseGasEstimationTransactionScanGasEstimationError{}),
		},
	)
}

type TransactionScanResponseGasEstimationTransactionScanGasEstimation struct {
	Estimate int64                                                                  `json:"estimate,required"`
	Status   TransactionScanResponseGasEstimationTransactionScanGasEstimationStatus `json:"status,required"`
	Used     int64                                                                  `json:"used,required"`
	JSON     transactionScanResponseGasEstimationTransactionScanGasEstimationJSON   `json:"-"`
}

// transactionScanResponseGasEstimationTransactionScanGasEstimationJSON contains
// the JSON metadata for the struct
// [TransactionScanResponseGasEstimationTransactionScanGasEstimation]
type transactionScanResponseGasEstimationTransactionScanGasEstimationJSON struct {
	Estimate    apijson.Field
	Status      apijson.Field
	Used        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionScanResponseGasEstimationTransactionScanGasEstimation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseGasEstimationTransactionScanGasEstimationJSON) RawJSON() string {
	return r.raw
}

func (r TransactionScanResponseGasEstimationTransactionScanGasEstimation) implementsTransactionScanResponseGasEstimation() {
}

type TransactionScanResponseGasEstimationTransactionScanGasEstimationStatus string

const (
	TransactionScanResponseGasEstimationTransactionScanGasEstimationStatusSuccess TransactionScanResponseGasEstimationTransactionScanGasEstimationStatus = "Success"
)

func (r TransactionScanResponseGasEstimationTransactionScanGasEstimationStatus) IsKnown() bool {
	switch r {
	case TransactionScanResponseGasEstimationTransactionScanGasEstimationStatusSuccess:
		return true
	}
	return false
}

type TransactionScanResponseGasEstimationTransactionScanGasEstimationError struct {
	Error  string                                                                      `json:"error,required"`
	Status TransactionScanResponseGasEstimationTransactionScanGasEstimationErrorStatus `json:"status,required"`
	JSON   transactionScanResponseGasEstimationTransactionScanGasEstimationErrorJSON   `json:"-"`
}

// transactionScanResponseGasEstimationTransactionScanGasEstimationErrorJSON
// contains the JSON metadata for the struct
// [TransactionScanResponseGasEstimationTransactionScanGasEstimationError]
type transactionScanResponseGasEstimationTransactionScanGasEstimationErrorJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionScanResponseGasEstimationTransactionScanGasEstimationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseGasEstimationTransactionScanGasEstimationErrorJSON) RawJSON() string {
	return r.raw
}

func (r TransactionScanResponseGasEstimationTransactionScanGasEstimationError) implementsTransactionScanResponseGasEstimation() {
}

type TransactionScanResponseGasEstimationTransactionScanGasEstimationErrorStatus string

const (
	TransactionScanResponseGasEstimationTransactionScanGasEstimationErrorStatusError TransactionScanResponseGasEstimationTransactionScanGasEstimationErrorStatus = "Error"
)

func (r TransactionScanResponseGasEstimationTransactionScanGasEstimationErrorStatus) IsKnown() bool {
	switch r {
	case TransactionScanResponseGasEstimationTransactionScanGasEstimationErrorStatusError:
		return true
	}
	return false
}

type TransactionScanResponseGasEstimationStatus string

const (
	TransactionScanResponseGasEstimationStatusSuccess TransactionScanResponseGasEstimationStatus = "Success"
	TransactionScanResponseGasEstimationStatusError   TransactionScanResponseGasEstimationStatus = "Error"
)

func (r TransactionScanResponseGasEstimationStatus) IsKnown() bool {
	switch r {
	case TransactionScanResponseGasEstimationStatusSuccess, TransactionScanResponseGasEstimationStatusError:
		return true
	}
	return false
}

type TransactionScanResponseSimulation struct {
	// A string indicating if the simulation was successful or not.
	Status TransactionScanResponseSimulationStatus `json:"status,required"`
	// This field can have the runtime type of [TransactionSimulationAccountSummary].
	AccountSummary interface{} `json:"account_summary"`
	// This field can have the runtime type of
	// [map[string]TransactionSimulationAddressDetail].
	AddressDetails interface{} `json:"address_details"`
	// This field can have the runtime type of
	// [map[string][]TransactionSimulationAssetsDiff].
	AssetsDiffs interface{} `json:"assets_diffs"`
	// This field can have the runtime type of
	// [map[string][]TransactionSimulationContractManagement].
	ContractManagement interface{} `json:"contract_management"`
	// An error message if the simulation failed.
	Error string `json:"error"`
	// This field can have the runtime type of
	// [TransactionSimulationErrorErrorDetails].
	ErrorDetails interface{} `json:"error_details"`
	// This field can have the runtime type of
	// [map[string][]TransactionSimulationExposure].
	Exposures interface{} `json:"exposures"`
	// This field can have the runtime type of [TransactionSimulationParams].
	Params interface{} `json:"params"`
	// This field can have the runtime type of [map[string]UsdDiff].
	TotalUsdDiff interface{} `json:"total_usd_diff"`
	// This field can have the runtime type of [map[string]map[string]string].
	TotalUsdExposure interface{}                           `json:"total_usd_exposure"`
	JSON             transactionScanResponseSimulationJSON `json:"-"`
	union            TransactionScanResponseSimulationUnion
}

// transactionScanResponseSimulationJSON contains the JSON metadata for the struct
// [TransactionScanResponseSimulation]
type transactionScanResponseSimulationJSON struct {
	Status             apijson.Field
	AccountSummary     apijson.Field
	AddressDetails     apijson.Field
	AssetsDiffs        apijson.Field
	ContractManagement apijson.Field
	Error              apijson.Field
	ErrorDetails       apijson.Field
	Exposures          apijson.Field
	Params             apijson.Field
	TotalUsdDiff       apijson.Field
	TotalUsdExposure   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r transactionScanResponseSimulationJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionScanResponseSimulation) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionScanResponseSimulation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TransactionScanResponseSimulationUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [TransactionSimulation],
// [TransactionSimulationError].
func (r TransactionScanResponseSimulation) AsUnion() TransactionScanResponseSimulationUnion {
	return r.union
}

// Union satisfied by [TransactionSimulation] or [TransactionSimulationError].
type TransactionScanResponseSimulationUnion interface {
	implementsTransactionScanResponseSimulation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionScanResponseSimulationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationError{}),
		},
	)
}

// A string indicating if the simulation was successful or not.
type TransactionScanResponseSimulationStatus string

const (
	TransactionScanResponseSimulationStatusSuccess TransactionScanResponseSimulationStatus = "Success"
	TransactionScanResponseSimulationStatusError   TransactionScanResponseSimulationStatus = "Error"
)

func (r TransactionScanResponseSimulationStatus) IsKnown() bool {
	switch r {
	case TransactionScanResponseSimulationStatusSuccess, TransactionScanResponseSimulationStatusError:
		return true
	}
	return false
}

type TransactionScanResponseValidation struct {
	// This field can have the runtime type of [[]TransactionScanFeature].
	Features interface{} `json:"features,required"`
	// An enumeration.
	ResultType TransactionScanResponseValidationResultType `json:"result_type,required"`
	// A string indicating if the simulation was successful or not.
	Status TransactionScanResponseValidationStatus `json:"status,required"`
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification"`
	// A textual description that can be presented to the user about what this
	// transaction is doing.
	Description string `json:"description"`
	// An error message if the validation failed.
	Error string `json:"error"`
	// A textual description about the reasons the transaction was flagged with
	// result_type.
	Reason string                                `json:"reason"`
	JSON   transactionScanResponseValidationJSON `json:"-"`
	union  TransactionScanResponseValidationUnion
}

// transactionScanResponseValidationJSON contains the JSON metadata for the struct
// [TransactionScanResponseValidation]
type transactionScanResponseValidationJSON struct {
	Features       apijson.Field
	ResultType     apijson.Field
	Status         apijson.Field
	Classification apijson.Field
	Description    apijson.Field
	Error          apijson.Field
	Reason         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r transactionScanResponseValidationJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionScanResponseValidation) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionScanResponseValidation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TransactionScanResponseValidationUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [TransactionValidation],
// [TransactionValidationError].
func (r TransactionScanResponseValidation) AsUnion() TransactionScanResponseValidationUnion {
	return r.union
}

// Union satisfied by [TransactionValidation] or [TransactionValidationError].
type TransactionScanResponseValidationUnion interface {
	implementsTransactionScanResponseValidation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionScanResponseValidationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionValidation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionValidationError{}),
		},
	)
}

// An enumeration.
type TransactionScanResponseValidationResultType string

const (
	TransactionScanResponseValidationResultTypeBenign    TransactionScanResponseValidationResultType = "Benign"
	TransactionScanResponseValidationResultTypeWarning   TransactionScanResponseValidationResultType = "Warning"
	TransactionScanResponseValidationResultTypeMalicious TransactionScanResponseValidationResultType = "Malicious"
	TransactionScanResponseValidationResultTypeError     TransactionScanResponseValidationResultType = "Error"
)

func (r TransactionScanResponseValidationResultType) IsKnown() bool {
	switch r {
	case TransactionScanResponseValidationResultTypeBenign, TransactionScanResponseValidationResultTypeWarning, TransactionScanResponseValidationResultTypeMalicious, TransactionScanResponseValidationResultTypeError:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type TransactionScanResponseValidationStatus string

const (
	TransactionScanResponseValidationStatusSuccess TransactionScanResponseValidationStatus = "Success"
)

func (r TransactionScanResponseValidationStatus) IsKnown() bool {
	switch r {
	case TransactionScanResponseValidationStatusSuccess:
		return true
	}
	return false
}

// The chain name
type TransactionScanSupportedChain string

const (
	TransactionScanSupportedChainArbitrum              TransactionScanSupportedChain = "arbitrum"
	TransactionScanSupportedChainAvalanche             TransactionScanSupportedChain = "avalanche"
	TransactionScanSupportedChainBase                  TransactionScanSupportedChain = "base"
	TransactionScanSupportedChainBaseSepolia           TransactionScanSupportedChain = "base-sepolia"
	TransactionScanSupportedChainBsc                   TransactionScanSupportedChain = "bsc"
	TransactionScanSupportedChainEthereum              TransactionScanSupportedChain = "ethereum"
	TransactionScanSupportedChainOptimism              TransactionScanSupportedChain = "optimism"
	TransactionScanSupportedChainPolygon               TransactionScanSupportedChain = "polygon"
	TransactionScanSupportedChainZksync                TransactionScanSupportedChain = "zksync"
	TransactionScanSupportedChainZksyncSepolia         TransactionScanSupportedChain = "zksync-sepolia"
	TransactionScanSupportedChainZora                  TransactionScanSupportedChain = "zora"
	TransactionScanSupportedChainLinea                 TransactionScanSupportedChain = "linea"
	TransactionScanSupportedChainBlast                 TransactionScanSupportedChain = "blast"
	TransactionScanSupportedChainScroll                TransactionScanSupportedChain = "scroll"
	TransactionScanSupportedChainEthereumSepolia       TransactionScanSupportedChain = "ethereum-sepolia"
	TransactionScanSupportedChainDegen                 TransactionScanSupportedChain = "degen"
	TransactionScanSupportedChainAvalancheFuji         TransactionScanSupportedChain = "avalanche-fuji"
	TransactionScanSupportedChainImmutableZkevm        TransactionScanSupportedChain = "immutable-zkevm"
	TransactionScanSupportedChainImmutableZkevmTestnet TransactionScanSupportedChain = "immutable-zkevm-testnet"
	TransactionScanSupportedChainGnosis                TransactionScanSupportedChain = "gnosis"
	TransactionScanSupportedChainWorldchain            TransactionScanSupportedChain = "worldchain"
	TransactionScanSupportedChainSoneiumMinato         TransactionScanSupportedChain = "soneium-minato"
	TransactionScanSupportedChainRonin                 TransactionScanSupportedChain = "ronin"
	TransactionScanSupportedChainApechain              TransactionScanSupportedChain = "apechain"
	TransactionScanSupportedChainZeroNetwork           TransactionScanSupportedChain = "zero-network"
	TransactionScanSupportedChainBerachain             TransactionScanSupportedChain = "berachain"
	TransactionScanSupportedChainBerachainBartio       TransactionScanSupportedChain = "berachain-bartio"
	TransactionScanSupportedChainInk                   TransactionScanSupportedChain = "ink"
	TransactionScanSupportedChainInkSepolia            TransactionScanSupportedChain = "ink-sepolia"
	TransactionScanSupportedChainAbstract              TransactionScanSupportedChain = "abstract"
	TransactionScanSupportedChainSoneium               TransactionScanSupportedChain = "soneium"
	TransactionScanSupportedChainUnichain              TransactionScanSupportedChain = "unichain"
)

func (r TransactionScanSupportedChain) IsKnown() bool {
	switch r {
	case TransactionScanSupportedChainArbitrum, TransactionScanSupportedChainAvalanche, TransactionScanSupportedChainBase, TransactionScanSupportedChainBaseSepolia, TransactionScanSupportedChainBsc, TransactionScanSupportedChainEthereum, TransactionScanSupportedChainOptimism, TransactionScanSupportedChainPolygon, TransactionScanSupportedChainZksync, TransactionScanSupportedChainZksyncSepolia, TransactionScanSupportedChainZora, TransactionScanSupportedChainLinea, TransactionScanSupportedChainBlast, TransactionScanSupportedChainScroll, TransactionScanSupportedChainEthereumSepolia, TransactionScanSupportedChainDegen, TransactionScanSupportedChainAvalancheFuji, TransactionScanSupportedChainImmutableZkevm, TransactionScanSupportedChainImmutableZkevmTestnet, TransactionScanSupportedChainGnosis, TransactionScanSupportedChainWorldchain, TransactionScanSupportedChainSoneiumMinato, TransactionScanSupportedChainRonin, TransactionScanSupportedChainApechain, TransactionScanSupportedChainZeroNetwork, TransactionScanSupportedChainBerachain, TransactionScanSupportedChainBerachainBartio, TransactionScanSupportedChainInk, TransactionScanSupportedChainInkSepolia, TransactionScanSupportedChainAbstract, TransactionScanSupportedChainSoneium, TransactionScanSupportedChainUnichain:
		return true
	}
	return false
}

type TransactionSimulation struct {
	// Account summary for the account address. account address is determined implicit
	// by the `from` field in the transaction request, or explicit by the
	// account_address field in the request.
	AccountSummary TransactionSimulationAccountSummary `json:"account_summary,required"`
	// a dictionary including additional information about each relevant address in the
	// transaction.
	AddressDetails map[string]TransactionSimulationAddressDetail `json:"address_details,required"`
	// dictionary describes the assets differences as a result of this transaction for
	// every involved address
	AssetsDiffs map[string][]TransactionSimulationAssetsDiff `json:"assets_diffs,required"`
	// dictionary describes the exposure differences as a result of this transaction
	// for every involved address (as a result of any approval / setApproval / permit
	// function)
	Exposures map[string][]TransactionSimulationExposure `json:"exposures,required"`
	// A string indicating if the simulation was successful or not.
	Status TransactionSimulationStatus `json:"status,required"`
	// dictionary represents the usd value each address gained / lost during this
	// transaction
	TotalUsdDiff map[string]UsdDiff `json:"total_usd_diff,required"`
	// a dictionary representing the usd value each address is exposed to, split by
	// spenders
	TotalUsdExposure map[string]map[string]string `json:"total_usd_exposure,required"`
	// Describes the state differences as a result of this transaction for every
	// involved address
	ContractManagement map[string][]TransactionSimulationContractManagement `json:"contract_management"`
	// The parameters of the transaction that was simulated.
	Params TransactionSimulationParams `json:"params"`
	JSON   transactionSimulationJSON   `json:"-"`
}

// transactionSimulationJSON contains the JSON metadata for the struct
// [TransactionSimulation]
type transactionSimulationJSON struct {
	AccountSummary     apijson.Field
	AddressDetails     apijson.Field
	AssetsDiffs        apijson.Field
	Exposures          apijson.Field
	Status             apijson.Field
	TotalUsdDiff       apijson.Field
	TotalUsdExposure   apijson.Field
	ContractManagement apijson.Field
	Params             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TransactionSimulation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulation) implementsTransactionScanResponseSimulation() {}

// Account summary for the account address. account address is determined implicit
// by the `from` field in the transaction request, or explicit by the
// account_address field in the request.
type TransactionSimulationAccountSummary struct {
	// All assets diffs related to the account address
	AssetsDiffs []TransactionSimulationAccountSummaryAssetsDiff `json:"assets_diffs,required"`
	// All assets exposures related to the account address
	Exposures []TransactionSimulationAccountSummaryExposure `json:"exposures,required"`
	// Total usd diff related to the account address
	TotalUsdDiff UsdDiff `json:"total_usd_diff,required"`
	// Total usd exposure related to the account address
	TotalUsdExposure map[string]string `json:"total_usd_exposure,required"`
	// All assets traces related to the account address
	Traces []TransactionSimulationAccountSummaryTrace `json:"traces,required"`
	JSON   transactionSimulationAccountSummaryJSON    `json:"-"`
}

// transactionSimulationAccountSummaryJSON contains the JSON metadata for the
// struct [TransactionSimulationAccountSummary]
type transactionSimulationAccountSummaryJSON struct {
	AssetsDiffs      apijson.Field
	Exposures        apijson.Field
	TotalUsdDiff     apijson.Field
	TotalUsdExposure apijson.Field
	Traces           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransactionSimulationAccountSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAccountSummaryJSON) RawJSON() string {
	return r.raw
}

type TransactionSimulationAccountSummaryAssetsDiff struct {
	// This field can have the runtime type of
	// [TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAsset],
	// [TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAsset],
	// [TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAsset],
	// [NativeAssetDetails].
	Asset interface{} `json:"asset,required"`
	// type of the asset for the current diff
	AssetType TransactionSimulationAccountSummaryAssetsDiffsAssetType `json:"asset_type,required"`
	// This field can have the runtime type of [[]Erc20Diff], [[]Erc721Diff],
	// [[]Erc1155Diff], [[]NativeDiff].
	In interface{} `json:"in,required"`
	// This field can have the runtime type of [[]Erc20Diff], [[]Erc721Diff],
	// [[]Erc1155Diff], [[]NativeDiff].
	Out interface{} `json:"out,required"`
	// This field can have the runtime type of
	// [TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffBalanceChanges],
	// [TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffBalanceChanges],
	// [TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffBalanceChanges],
	// [TransactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiffBalanceChanges].
	BalanceChanges interface{}                                       `json:"balance_changes"`
	JSON           transactionSimulationAccountSummaryAssetsDiffJSON `json:"-"`
	union          TransactionSimulationAccountSummaryAssetsDiffsUnion
}

// transactionSimulationAccountSummaryAssetsDiffJSON contains the JSON metadata for
// the struct [TransactionSimulationAccountSummaryAssetsDiff]
type transactionSimulationAccountSummaryAssetsDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r transactionSimulationAccountSummaryAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationAccountSummaryAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationAccountSummaryAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TransactionSimulationAccountSummaryAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiff],
// [TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiff],
// [TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiff],
// [TransactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiff].
func (r TransactionSimulationAccountSummaryAssetsDiff) AsUnion() TransactionSimulationAccountSummaryAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiff],
// [TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiff],
// [TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiff]
// or
// [TransactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiff].
type TransactionSimulationAccountSummaryAssetsDiffsUnion interface {
	implementsTransactionSimulationAccountSummaryAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationAccountSummaryAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiff{}),
		},
	)
}

type TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiff struct {
	// description of the asset for the current diff
	Asset TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []Erc20Diff `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out []Erc20Diff `json:"out,required"`
	// shows the balance before making the transaction and after
	BalanceChanges TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffBalanceChanges `json:"balance_changes"`
	JSON           transactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffJSON           `json:"-"`
}

// transactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffJSON
// contains the JSON metadata for the struct
// [TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiff]
type transactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiff) implementsTransactionSimulationAccountSummaryAssetsDiff() {
}

// description of the asset for the current diff
type TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAssetType `json:"type,required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                    `json:"symbol"`
	JSON   transactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAssetJSON `json:"-"`
	union  TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAssetUnion
}

// transactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAssetJSON
// contains the JSON metadata for the struct
// [TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAsset]
type transactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [Erc20TokenDetails],
// [NonercTokenDetails].
func (r TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAsset) AsUnion() TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by [Erc20TokenDetails] or [NonercTokenDetails].
type TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAssetUnion interface {
	implementsTransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NonercTokenDetails{}),
		},
	)
}

// asset type.
type TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAssetType string

const (
	TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAssetTypeErc20  TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAssetType = "ERC20"
	TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAssetTypeNonerc TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAssetType = "NONERC"
)

func (r TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAssetTypeErc20, TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffAssetTypeNonerc:
		return true
	}
	return false
}

// shows the balance before making the transaction and after
type TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffBalanceChanges struct {
	// balance of the account after making the transaction
	After Erc20Diff `json:"after,required"`
	// balance of the account before making the transaction
	Before Erc20Diff                                                                                          `json:"before,required"`
	JSON   transactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffBalanceChangesJSON `json:"-"`
}

// transactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffBalanceChangesJSON
// contains the JSON metadata for the struct
// [TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffBalanceChanges]
type transactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffBalanceChangesJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffBalanceChanges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAccountSummaryAssetsDiffsErc20AddressAssetBalanceChangeDiffBalanceChangesJSON) RawJSON() string {
	return r.raw
}

type TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiff struct {
	// description of the asset for the current diff
	Asset TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []Erc721Diff `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out []Erc721Diff `json:"out,required"`
	// shows the balance before making the transaction and after
	BalanceChanges TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffBalanceChanges `json:"balance_changes"`
	JSON           transactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffJSON           `json:"-"`
}

// transactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffJSON
// contains the JSON metadata for the struct
// [TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiff]
type transactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiff) implementsTransactionSimulationAccountSummaryAssetsDiff() {
}

// description of the asset for the current diff
type TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                     `json:"symbol"`
	JSON   transactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAssetJSON `json:"-"`
	union  TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAssetUnion
}

// transactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAssetJSON
// contains the JSON metadata for the struct
// [TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAsset]
type transactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [Erc721TokenDetails],
// [NonercTokenDetails].
func (r TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAsset) AsUnion() TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by [Erc721TokenDetails] or [NonercTokenDetails].
type TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAssetUnion interface {
	implementsTransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NonercTokenDetails{}),
		},
	)
}

// asset type.
type TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAssetType string

const (
	TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAssetTypeErc721 TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAssetType = "ERC721"
	TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAssetTypeNonerc TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAssetType = "NONERC"
)

func (r TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAssetTypeErc721, TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffAssetTypeNonerc:
		return true
	}
	return false
}

// shows the balance before making the transaction and after
type TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffBalanceChanges struct {
	// balance of the account after making the transaction
	After Erc721Diff `json:"after,required"`
	// balance of the account before making the transaction
	Before Erc721Diff                                                                                          `json:"before,required"`
	JSON   transactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffBalanceChangesJSON `json:"-"`
}

// transactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffBalanceChangesJSON
// contains the JSON metadata for the struct
// [TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffBalanceChanges]
type transactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffBalanceChangesJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffBalanceChanges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAccountSummaryAssetsDiffsErc721AddressAssetBalanceChangeDiffBalanceChangesJSON) RawJSON() string {
	return r.raw
}

type TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiff struct {
	// description of the asset for the current diff
	Asset TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []Erc1155Diff `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out []Erc1155Diff `json:"out,required"`
	// shows the balance before making the transaction and after
	BalanceChanges TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffBalanceChanges `json:"balance_changes"`
	JSON           transactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffJSON           `json:"-"`
}

// transactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffJSON
// contains the JSON metadata for the struct
// [TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiff]
type transactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiff) implementsTransactionSimulationAccountSummaryAssetsDiff() {
}

// description of the asset for the current diff
type TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                                      `json:"symbol"`
	JSON   transactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAssetJSON `json:"-"`
	union  TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAssetUnion
}

// transactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAssetJSON
// contains the JSON metadata for the struct
// [TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAsset]
type transactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [Erc1155TokenDetails],
// [NonercTokenDetails].
func (r TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAsset) AsUnion() TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by [Erc1155TokenDetails] or [NonercTokenDetails].
type TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAssetUnion interface {
	implementsTransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NonercTokenDetails{}),
		},
	)
}

// asset type.
type TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAssetType string

const (
	TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAssetTypeErc1155 TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAssetType = "ERC1155"
	TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAssetTypeNonerc  TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAssetType = "NONERC"
)

func (r TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAssetTypeErc1155, TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffAssetTypeNonerc:
		return true
	}
	return false
}

// shows the balance before making the transaction and after
type TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffBalanceChanges struct {
	// balance of the account after making the transaction
	After Erc1155Diff `json:"after,required"`
	// balance of the account before making the transaction
	Before Erc1155Diff                                                                                          `json:"before,required"`
	JSON   transactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffBalanceChangesJSON `json:"-"`
}

// transactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffBalanceChangesJSON
// contains the JSON metadata for the struct
// [TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffBalanceChanges]
type transactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffBalanceChangesJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffBalanceChanges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAccountSummaryAssetsDiffsErc1155AddressAssetBalanceChangeDiffBalanceChangesJSON) RawJSON() string {
	return r.raw
}

type TransactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiff struct {
	// description of the asset for the current diff
	Asset NativeAssetDetails `json:"asset,required"`
	// type of the asset for the current diff
	AssetType TransactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []NativeDiff `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out []NativeDiff `json:"out,required"`
	// shows the balance before making the transaction and after
	BalanceChanges TransactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiffBalanceChanges `json:"balance_changes"`
	JSON           transactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiffJSON           `json:"-"`
}

// transactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiffJSON
// contains the JSON metadata for the struct
// [TransactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiff]
type transactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiffJSON struct {
	Asset          apijson.Field
	AssetType      apijson.Field
	In             apijson.Field
	Out            apijson.Field
	BalanceChanges apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiffJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiff) implementsTransactionSimulationAccountSummaryAssetsDiff() {
}

// type of the asset for the current diff
type TransactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiffAssetType string

const (
	TransactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiffAssetTypeNative TransactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiffAssetType = "NATIVE"
)

func (r TransactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiffAssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiffAssetTypeNative:
		return true
	}
	return false
}

// shows the balance before making the transaction and after
type TransactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiffBalanceChanges struct {
	// balance of the account after making the transaction
	After NativeDiff `json:"after,required"`
	// balance of the account before making the transaction
	Before NativeDiff                                                                                          `json:"before,required"`
	JSON   transactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiffBalanceChangesJSON `json:"-"`
}

// transactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiffBalanceChangesJSON
// contains the JSON metadata for the struct
// [TransactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiffBalanceChanges]
type transactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiffBalanceChangesJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiffBalanceChanges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAccountSummaryAssetsDiffsNativeAddressAssetBalanceChangeDiffBalanceChangesJSON) RawJSON() string {
	return r.raw
}

// type of the asset for the current diff
type TransactionSimulationAccountSummaryAssetsDiffsAssetType string

const (
	TransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc20   TransactionSimulationAccountSummaryAssetsDiffsAssetType = "ERC20"
	TransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc721  TransactionSimulationAccountSummaryAssetsDiffsAssetType = "ERC721"
	TransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc1155 TransactionSimulationAccountSummaryAssetsDiffsAssetType = "ERC1155"
	TransactionSimulationAccountSummaryAssetsDiffsAssetTypeNative  TransactionSimulationAccountSummaryAssetsDiffsAssetType = "NATIVE"
)

func (r TransactionSimulationAccountSummaryAssetsDiffsAssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc20, TransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc721, TransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc1155, TransactionSimulationAccountSummaryAssetsDiffsAssetTypeNative:
		return true
	}
	return false
}

type TransactionSimulationAccountSummaryExposure struct {
	// This field can have the runtime type of
	// [TransactionSimulationAccountSummaryExposuresErc20AddressExposureAsset],
	// [TransactionSimulationAccountSummaryExposuresErc721AddressExposureAsset],
	// [TransactionSimulationAccountSummaryExposuresErc1155AddressExposureAsset].
	Asset interface{} `json:"asset,required"`
	// type of the asset for the current diff
	AssetType TransactionSimulationAccountSummaryExposuresAssetType `json:"asset_type,required"`
	// This field can have the runtime type of [map[string]Erc20Exposure],
	// [map[string]Erc721Exposure], [map[string]Erc1155Exposure].
	Spenders interface{}                                     `json:"spenders,required"`
	JSON     transactionSimulationAccountSummaryExposureJSON `json:"-"`
	union    TransactionSimulationAccountSummaryExposuresUnion
}

// transactionSimulationAccountSummaryExposureJSON contains the JSON metadata for
// the struct [TransactionSimulationAccountSummaryExposure]
type transactionSimulationAccountSummaryExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionSimulationAccountSummaryExposureJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationAccountSummaryExposure) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationAccountSummaryExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TransactionSimulationAccountSummaryExposuresUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TransactionSimulationAccountSummaryExposuresErc20AddressExposure],
// [TransactionSimulationAccountSummaryExposuresErc721AddressExposure],
// [TransactionSimulationAccountSummaryExposuresErc1155AddressExposure].
func (r TransactionSimulationAccountSummaryExposure) AsUnion() TransactionSimulationAccountSummaryExposuresUnion {
	return r.union
}

// Union satisfied by
// [TransactionSimulationAccountSummaryExposuresErc20AddressExposure],
// [TransactionSimulationAccountSummaryExposuresErc721AddressExposure] or
// [TransactionSimulationAccountSummaryExposuresErc1155AddressExposure].
type TransactionSimulationAccountSummaryExposuresUnion interface {
	implementsTransactionSimulationAccountSummaryExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationAccountSummaryExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationAccountSummaryExposuresErc20AddressExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationAccountSummaryExposuresErc721AddressExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationAccountSummaryExposuresErc1155AddressExposure{}),
		},
	)
}

type TransactionSimulationAccountSummaryExposuresErc20AddressExposure struct {
	// description of the asset for the current diff
	Asset TransactionSimulationAccountSummaryExposuresErc20AddressExposureAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType TransactionSimulationAccountSummaryExposuresErc20AddressExposureAssetType `json:"asset_type,required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]Erc20Exposure                                             `json:"spenders,required"`
	JSON     transactionSimulationAccountSummaryExposuresErc20AddressExposureJSON `json:"-"`
}

// transactionSimulationAccountSummaryExposuresErc20AddressExposureJSON contains
// the JSON metadata for the struct
// [TransactionSimulationAccountSummaryExposuresErc20AddressExposure]
type transactionSimulationAccountSummaryExposuresErc20AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationAccountSummaryExposuresErc20AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAccountSummaryExposuresErc20AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationAccountSummaryExposuresErc20AddressExposure) implementsTransactionSimulationAccountSummaryExposure() {
}

// description of the asset for the current diff
type TransactionSimulationAccountSummaryExposuresErc20AddressExposureAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type TransactionSimulationAccountSummaryExposuresErc20AddressExposureAssetType `json:"type,required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                    `json:"symbol"`
	JSON   transactionSimulationAccountSummaryExposuresErc20AddressExposureAssetJSON `json:"-"`
	union  TransactionSimulationAccountSummaryExposuresErc20AddressExposureAssetUnion
}

// transactionSimulationAccountSummaryExposuresErc20AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [TransactionSimulationAccountSummaryExposuresErc20AddressExposureAsset]
type transactionSimulationAccountSummaryExposuresErc20AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionSimulationAccountSummaryExposuresErc20AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationAccountSummaryExposuresErc20AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationAccountSummaryExposuresErc20AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [TransactionSimulationAccountSummaryExposuresErc20AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [Erc20TokenDetails],
// [NonercTokenDetails].
func (r TransactionSimulationAccountSummaryExposuresErc20AddressExposureAsset) AsUnion() TransactionSimulationAccountSummaryExposuresErc20AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by [Erc20TokenDetails] or [NonercTokenDetails].
type TransactionSimulationAccountSummaryExposuresErc20AddressExposureAssetUnion interface {
	implementsTransactionSimulationAccountSummaryExposuresErc20AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationAccountSummaryExposuresErc20AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NonercTokenDetails{}),
		},
	)
}

// asset type.
type TransactionSimulationAccountSummaryExposuresErc20AddressExposureAssetType string

const (
	TransactionSimulationAccountSummaryExposuresErc20AddressExposureAssetTypeErc20  TransactionSimulationAccountSummaryExposuresErc20AddressExposureAssetType = "ERC20"
	TransactionSimulationAccountSummaryExposuresErc20AddressExposureAssetTypeNonerc TransactionSimulationAccountSummaryExposuresErc20AddressExposureAssetType = "NONERC"
)

func (r TransactionSimulationAccountSummaryExposuresErc20AddressExposureAssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryExposuresErc20AddressExposureAssetTypeErc20, TransactionSimulationAccountSummaryExposuresErc20AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type TransactionSimulationAccountSummaryExposuresErc721AddressExposure struct {
	// description of the asset for the current diff
	Asset TransactionSimulationAccountSummaryExposuresErc721AddressExposureAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType TransactionSimulationAccountSummaryExposuresErc721AddressExposureAssetType `json:"asset_type,required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]Erc721Exposure                                             `json:"spenders,required"`
	JSON     transactionSimulationAccountSummaryExposuresErc721AddressExposureJSON `json:"-"`
}

// transactionSimulationAccountSummaryExposuresErc721AddressExposureJSON contains
// the JSON metadata for the struct
// [TransactionSimulationAccountSummaryExposuresErc721AddressExposure]
type transactionSimulationAccountSummaryExposuresErc721AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationAccountSummaryExposuresErc721AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAccountSummaryExposuresErc721AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationAccountSummaryExposuresErc721AddressExposure) implementsTransactionSimulationAccountSummaryExposure() {
}

// description of the asset for the current diff
type TransactionSimulationAccountSummaryExposuresErc721AddressExposureAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type TransactionSimulationAccountSummaryExposuresErc721AddressExposureAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                     `json:"symbol"`
	JSON   transactionSimulationAccountSummaryExposuresErc721AddressExposureAssetJSON `json:"-"`
	union  TransactionSimulationAccountSummaryExposuresErc721AddressExposureAssetUnion
}

// transactionSimulationAccountSummaryExposuresErc721AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [TransactionSimulationAccountSummaryExposuresErc721AddressExposureAsset]
type transactionSimulationAccountSummaryExposuresErc721AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionSimulationAccountSummaryExposuresErc721AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationAccountSummaryExposuresErc721AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationAccountSummaryExposuresErc721AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [TransactionSimulationAccountSummaryExposuresErc721AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [Erc721TokenDetails],
// [NonercTokenDetails].
func (r TransactionSimulationAccountSummaryExposuresErc721AddressExposureAsset) AsUnion() TransactionSimulationAccountSummaryExposuresErc721AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by [Erc721TokenDetails] or [NonercTokenDetails].
type TransactionSimulationAccountSummaryExposuresErc721AddressExposureAssetUnion interface {
	implementsTransactionSimulationAccountSummaryExposuresErc721AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationAccountSummaryExposuresErc721AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NonercTokenDetails{}),
		},
	)
}

// asset type.
type TransactionSimulationAccountSummaryExposuresErc721AddressExposureAssetType string

const (
	TransactionSimulationAccountSummaryExposuresErc721AddressExposureAssetTypeErc721 TransactionSimulationAccountSummaryExposuresErc721AddressExposureAssetType = "ERC721"
	TransactionSimulationAccountSummaryExposuresErc721AddressExposureAssetTypeNonerc TransactionSimulationAccountSummaryExposuresErc721AddressExposureAssetType = "NONERC"
)

func (r TransactionSimulationAccountSummaryExposuresErc721AddressExposureAssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryExposuresErc721AddressExposureAssetTypeErc721, TransactionSimulationAccountSummaryExposuresErc721AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type TransactionSimulationAccountSummaryExposuresErc1155AddressExposure struct {
	// description of the asset for the current diff
	Asset TransactionSimulationAccountSummaryExposuresErc1155AddressExposureAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType TransactionSimulationAccountSummaryExposuresErc1155AddressExposureAssetType `json:"asset_type,required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]Erc1155Exposure                                             `json:"spenders,required"`
	JSON     transactionSimulationAccountSummaryExposuresErc1155AddressExposureJSON `json:"-"`
}

// transactionSimulationAccountSummaryExposuresErc1155AddressExposureJSON contains
// the JSON metadata for the struct
// [TransactionSimulationAccountSummaryExposuresErc1155AddressExposure]
type transactionSimulationAccountSummaryExposuresErc1155AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationAccountSummaryExposuresErc1155AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAccountSummaryExposuresErc1155AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationAccountSummaryExposuresErc1155AddressExposure) implementsTransactionSimulationAccountSummaryExposure() {
}

// description of the asset for the current diff
type TransactionSimulationAccountSummaryExposuresErc1155AddressExposureAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type TransactionSimulationAccountSummaryExposuresErc1155AddressExposureAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                                      `json:"symbol"`
	JSON   transactionSimulationAccountSummaryExposuresErc1155AddressExposureAssetJSON `json:"-"`
	union  TransactionSimulationAccountSummaryExposuresErc1155AddressExposureAssetUnion
}

// transactionSimulationAccountSummaryExposuresErc1155AddressExposureAssetJSON
// contains the JSON metadata for the struct
// [TransactionSimulationAccountSummaryExposuresErc1155AddressExposureAsset]
type transactionSimulationAccountSummaryExposuresErc1155AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionSimulationAccountSummaryExposuresErc1155AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationAccountSummaryExposuresErc1155AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationAccountSummaryExposuresErc1155AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [TransactionSimulationAccountSummaryExposuresErc1155AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [Erc1155TokenDetails],
// [NonercTokenDetails].
func (r TransactionSimulationAccountSummaryExposuresErc1155AddressExposureAsset) AsUnion() TransactionSimulationAccountSummaryExposuresErc1155AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by [Erc1155TokenDetails] or [NonercTokenDetails].
type TransactionSimulationAccountSummaryExposuresErc1155AddressExposureAssetUnion interface {
	implementsTransactionSimulationAccountSummaryExposuresErc1155AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationAccountSummaryExposuresErc1155AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NonercTokenDetails{}),
		},
	)
}

// asset type.
type TransactionSimulationAccountSummaryExposuresErc1155AddressExposureAssetType string

const (
	TransactionSimulationAccountSummaryExposuresErc1155AddressExposureAssetTypeErc1155 TransactionSimulationAccountSummaryExposuresErc1155AddressExposureAssetType = "ERC1155"
	TransactionSimulationAccountSummaryExposuresErc1155AddressExposureAssetTypeNonerc  TransactionSimulationAccountSummaryExposuresErc1155AddressExposureAssetType = "NONERC"
)

func (r TransactionSimulationAccountSummaryExposuresErc1155AddressExposureAssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryExposuresErc1155AddressExposureAssetTypeErc1155, TransactionSimulationAccountSummaryExposuresErc1155AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

// type of the asset for the current diff
type TransactionSimulationAccountSummaryExposuresAssetType string

const (
	TransactionSimulationAccountSummaryExposuresAssetTypeErc20   TransactionSimulationAccountSummaryExposuresAssetType = "ERC20"
	TransactionSimulationAccountSummaryExposuresAssetTypeErc721  TransactionSimulationAccountSummaryExposuresAssetType = "ERC721"
	TransactionSimulationAccountSummaryExposuresAssetTypeErc1155 TransactionSimulationAccountSummaryExposuresAssetType = "ERC1155"
)

func (r TransactionSimulationAccountSummaryExposuresAssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryExposuresAssetTypeErc20, TransactionSimulationAccountSummaryExposuresAssetTypeErc721, TransactionSimulationAccountSummaryExposuresAssetTypeErc1155:
		return true
	}
	return false
}

type TransactionSimulationAccountSummaryTrace struct {
	// type of the trace
	TraceType TransactionSimulationAccountSummaryTracesTraceType `json:"trace_type,required"`
	// The type of the model
	Type TransactionSimulationAccountSummaryTracesType `json:"type,required"`
	// This field can have the runtime type of
	// [TransactionSimulationAccountSummaryTracesErc20AssetTraceAsset],
	// [TransactionSimulationAccountSummaryTracesErc721AssetTraceAsset],
	// [TransactionSimulationAccountSummaryTracesErc1155AssetTraceAsset],
	// [NativeAssetDetails].
	Asset interface{} `json:"asset"`
	// This field can have the runtime type of [Erc20Diff], [Erc721Diff],
	// [Erc1155Diff], [NativeDiff].
	Diff interface{} `json:"diff"`
	// This field can have the runtime type of
	// [TransactionSimulationAccountSummaryTracesErc20ExposureTraceExposed],
	// [TransactionSimulationAccountSummaryTracesErc721ExposureTraceExposed].
	Exposed interface{} `json:"exposed"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address"`
	// This field can have the runtime type of
	// [[]TransactionSimulationAccountSummaryTracesErc20AssetTraceLabels],
	// [[]TransactionSimulationAccountSummaryTracesErc721AssetTraceLabels],
	// [[]TransactionSimulationAccountSummaryTracesErc1155AssetTraceLabels],
	// [[]TransactionSimulationAccountSummaryTracesNativeAssetTraceLabels].
	Labels interface{} `json:"labels"`
	// The owner of the assets
	Owner string `json:"owner"`
	// The spender of the assets
	Spender string `json:"spender"`
	// The address where the assets are moved to
	ToAddress string                                       `json:"to_address"`
	JSON      transactionSimulationAccountSummaryTraceJSON `json:"-"`
	union     TransactionSimulationAccountSummaryTracesUnion
}

// transactionSimulationAccountSummaryTraceJSON contains the JSON metadata for the
// struct [TransactionSimulationAccountSummaryTrace]
type transactionSimulationAccountSummaryTraceJSON struct {
	TraceType   apijson.Field
	Type        apijson.Field
	Asset       apijson.Field
	Diff        apijson.Field
	Exposed     apijson.Field
	FromAddress apijson.Field
	Labels      apijson.Field
	Owner       apijson.Field
	Spender     apijson.Field
	ToAddress   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionSimulationAccountSummaryTraceJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationAccountSummaryTrace) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationAccountSummaryTrace{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TransactionSimulationAccountSummaryTracesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TransactionSimulationAccountSummaryTracesErc20AssetTrace],
// [TransactionSimulationAccountSummaryTracesErc721AssetTrace],
// [TransactionSimulationAccountSummaryTracesErc1155AssetTrace],
// [TransactionSimulationAccountSummaryTracesNativeAssetTrace],
// [TransactionSimulationAccountSummaryTracesErc20ExposureTrace],
// [TransactionSimulationAccountSummaryTracesErc721ExposureTrace],
// [TransactionSimulationAccountSummaryTracesErc1155ExposureTrace].
func (r TransactionSimulationAccountSummaryTrace) AsUnion() TransactionSimulationAccountSummaryTracesUnion {
	return r.union
}

// Union satisfied by [TransactionSimulationAccountSummaryTracesErc20AssetTrace],
// [TransactionSimulationAccountSummaryTracesErc721AssetTrace],
// [TransactionSimulationAccountSummaryTracesErc1155AssetTrace],
// [TransactionSimulationAccountSummaryTracesNativeAssetTrace],
// [TransactionSimulationAccountSummaryTracesErc20ExposureTrace],
// [TransactionSimulationAccountSummaryTracesErc721ExposureTrace] or
// [TransactionSimulationAccountSummaryTracesErc1155ExposureTrace].
type TransactionSimulationAccountSummaryTracesUnion interface {
	implementsTransactionSimulationAccountSummaryTrace()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationAccountSummaryTracesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationAccountSummaryTracesErc20AssetTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationAccountSummaryTracesErc721AssetTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationAccountSummaryTracesErc1155AssetTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationAccountSummaryTracesNativeAssetTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationAccountSummaryTracesErc20ExposureTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationAccountSummaryTracesErc721ExposureTrace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationAccountSummaryTracesErc1155ExposureTrace{}),
		},
	)
}

type TransactionSimulationAccountSummaryTracesErc20AssetTrace struct {
	// Description of the asset in the trace
	Asset TransactionSimulationAccountSummaryTracesErc20AssetTraceAsset `json:"asset,required"`
	// The difference in value for the asset in the trace
	Diff Erc20Diff `json:"diff,required"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address,required"`
	// The address where the assets are moved to
	ToAddress string `json:"to_address,required"`
	// type of the trace
	TraceType TransactionSimulationAccountSummaryTracesErc20AssetTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type TransactionSimulationAccountSummaryTracesErc20AssetTraceType `json:"type,required"`
	// List of labels that describe the trace
	Labels []TransactionSimulationAccountSummaryTracesErc20AssetTraceLabels `json:"labels"`
	JSON   transactionSimulationAccountSummaryTracesErc20AssetTraceJSON     `json:"-"`
}

// transactionSimulationAccountSummaryTracesErc20AssetTraceJSON contains the JSON
// metadata for the struct
// [TransactionSimulationAccountSummaryTracesErc20AssetTrace]
type transactionSimulationAccountSummaryTracesErc20AssetTraceJSON struct {
	Asset       apijson.Field
	Diff        apijson.Field
	FromAddress apijson.Field
	ToAddress   apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	Labels      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationAccountSummaryTracesErc20AssetTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAccountSummaryTracesErc20AssetTraceJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationAccountSummaryTracesErc20AssetTrace) implementsTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type TransactionSimulationAccountSummaryTracesErc20AssetTraceAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type TransactionSimulationAccountSummaryTracesErc20AssetTraceAssetType `json:"type,required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                            `json:"symbol"`
	JSON   transactionSimulationAccountSummaryTracesErc20AssetTraceAssetJSON `json:"-"`
	union  TransactionSimulationAccountSummaryTracesErc20AssetTraceAssetUnion
}

// transactionSimulationAccountSummaryTracesErc20AssetTraceAssetJSON contains the
// JSON metadata for the struct
// [TransactionSimulationAccountSummaryTracesErc20AssetTraceAsset]
type transactionSimulationAccountSummaryTracesErc20AssetTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionSimulationAccountSummaryTracesErc20AssetTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationAccountSummaryTracesErc20AssetTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationAccountSummaryTracesErc20AssetTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [TransactionSimulationAccountSummaryTracesErc20AssetTraceAssetUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [Erc20TokenDetails],
// [NonercTokenDetails].
func (r TransactionSimulationAccountSummaryTracesErc20AssetTraceAsset) AsUnion() TransactionSimulationAccountSummaryTracesErc20AssetTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by [Erc20TokenDetails] or [NonercTokenDetails].
type TransactionSimulationAccountSummaryTracesErc20AssetTraceAssetUnion interface {
	implementsTransactionSimulationAccountSummaryTracesErc20AssetTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationAccountSummaryTracesErc20AssetTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NonercTokenDetails{}),
		},
	)
}

// asset type.
type TransactionSimulationAccountSummaryTracesErc20AssetTraceAssetType string

const (
	TransactionSimulationAccountSummaryTracesErc20AssetTraceAssetTypeErc20  TransactionSimulationAccountSummaryTracesErc20AssetTraceAssetType = "ERC20"
	TransactionSimulationAccountSummaryTracesErc20AssetTraceAssetTypeNonerc TransactionSimulationAccountSummaryTracesErc20AssetTraceAssetType = "NONERC"
)

func (r TransactionSimulationAccountSummaryTracesErc20AssetTraceAssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryTracesErc20AssetTraceAssetTypeErc20, TransactionSimulationAccountSummaryTracesErc20AssetTraceAssetTypeNonerc:
		return true
	}
	return false
}

// type of the trace
type TransactionSimulationAccountSummaryTracesErc20AssetTraceTraceType string

const (
	TransactionSimulationAccountSummaryTracesErc20AssetTraceTraceTypeAssetTrace TransactionSimulationAccountSummaryTracesErc20AssetTraceTraceType = "AssetTrace"
)

func (r TransactionSimulationAccountSummaryTracesErc20AssetTraceTraceType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryTracesErc20AssetTraceTraceTypeAssetTrace:
		return true
	}
	return false
}

// The type of the model
type TransactionSimulationAccountSummaryTracesErc20AssetTraceType string

const (
	TransactionSimulationAccountSummaryTracesErc20AssetTraceTypeErc20AssetTrace TransactionSimulationAccountSummaryTracesErc20AssetTraceType = "ERC20AssetTrace"
)

func (r TransactionSimulationAccountSummaryTracesErc20AssetTraceType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryTracesErc20AssetTraceTypeErc20AssetTrace:
		return true
	}
	return false
}

// An enumeration.
type TransactionSimulationAccountSummaryTracesErc20AssetTraceLabels string

const (
	TransactionSimulationAccountSummaryTracesErc20AssetTraceLabelsGasFee TransactionSimulationAccountSummaryTracesErc20AssetTraceLabels = "GAS_FEE"
)

func (r TransactionSimulationAccountSummaryTracesErc20AssetTraceLabels) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryTracesErc20AssetTraceLabelsGasFee:
		return true
	}
	return false
}

type TransactionSimulationAccountSummaryTracesErc721AssetTrace struct {
	// Description of the asset in the trace
	Asset TransactionSimulationAccountSummaryTracesErc721AssetTraceAsset `json:"asset,required"`
	// The difference in value for the asset in the trace
	Diff Erc721Diff `json:"diff,required"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address,required"`
	// The address where the assets are moved to
	ToAddress string `json:"to_address,required"`
	// type of the trace
	TraceType TransactionSimulationAccountSummaryTracesErc721AssetTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type TransactionSimulationAccountSummaryTracesErc721AssetTraceType `json:"type,required"`
	// List of labels that describe the trace
	Labels []TransactionSimulationAccountSummaryTracesErc721AssetTraceLabels `json:"labels"`
	JSON   transactionSimulationAccountSummaryTracesErc721AssetTraceJSON     `json:"-"`
}

// transactionSimulationAccountSummaryTracesErc721AssetTraceJSON contains the JSON
// metadata for the struct
// [TransactionSimulationAccountSummaryTracesErc721AssetTrace]
type transactionSimulationAccountSummaryTracesErc721AssetTraceJSON struct {
	Asset       apijson.Field
	Diff        apijson.Field
	FromAddress apijson.Field
	ToAddress   apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	Labels      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationAccountSummaryTracesErc721AssetTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAccountSummaryTracesErc721AssetTraceJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationAccountSummaryTracesErc721AssetTrace) implementsTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type TransactionSimulationAccountSummaryTracesErc721AssetTraceAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type TransactionSimulationAccountSummaryTracesErc721AssetTraceAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                             `json:"symbol"`
	JSON   transactionSimulationAccountSummaryTracesErc721AssetTraceAssetJSON `json:"-"`
	union  TransactionSimulationAccountSummaryTracesErc721AssetTraceAssetUnion
}

// transactionSimulationAccountSummaryTracesErc721AssetTraceAssetJSON contains the
// JSON metadata for the struct
// [TransactionSimulationAccountSummaryTracesErc721AssetTraceAsset]
type transactionSimulationAccountSummaryTracesErc721AssetTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionSimulationAccountSummaryTracesErc721AssetTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationAccountSummaryTracesErc721AssetTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationAccountSummaryTracesErc721AssetTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [TransactionSimulationAccountSummaryTracesErc721AssetTraceAssetUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [Erc721TokenDetails],
// [NonercTokenDetails].
func (r TransactionSimulationAccountSummaryTracesErc721AssetTraceAsset) AsUnion() TransactionSimulationAccountSummaryTracesErc721AssetTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by [Erc721TokenDetails] or [NonercTokenDetails].
type TransactionSimulationAccountSummaryTracesErc721AssetTraceAssetUnion interface {
	implementsTransactionSimulationAccountSummaryTracesErc721AssetTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationAccountSummaryTracesErc721AssetTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NonercTokenDetails{}),
		},
	)
}

// asset type.
type TransactionSimulationAccountSummaryTracesErc721AssetTraceAssetType string

const (
	TransactionSimulationAccountSummaryTracesErc721AssetTraceAssetTypeErc721 TransactionSimulationAccountSummaryTracesErc721AssetTraceAssetType = "ERC721"
	TransactionSimulationAccountSummaryTracesErc721AssetTraceAssetTypeNonerc TransactionSimulationAccountSummaryTracesErc721AssetTraceAssetType = "NONERC"
)

func (r TransactionSimulationAccountSummaryTracesErc721AssetTraceAssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryTracesErc721AssetTraceAssetTypeErc721, TransactionSimulationAccountSummaryTracesErc721AssetTraceAssetTypeNonerc:
		return true
	}
	return false
}

// type of the trace
type TransactionSimulationAccountSummaryTracesErc721AssetTraceTraceType string

const (
	TransactionSimulationAccountSummaryTracesErc721AssetTraceTraceTypeAssetTrace TransactionSimulationAccountSummaryTracesErc721AssetTraceTraceType = "AssetTrace"
)

func (r TransactionSimulationAccountSummaryTracesErc721AssetTraceTraceType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryTracesErc721AssetTraceTraceTypeAssetTrace:
		return true
	}
	return false
}

// The type of the model
type TransactionSimulationAccountSummaryTracesErc721AssetTraceType string

const (
	TransactionSimulationAccountSummaryTracesErc721AssetTraceTypeErc721AssetTrace TransactionSimulationAccountSummaryTracesErc721AssetTraceType = "ERC721AssetTrace"
)

func (r TransactionSimulationAccountSummaryTracesErc721AssetTraceType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryTracesErc721AssetTraceTypeErc721AssetTrace:
		return true
	}
	return false
}

// An enumeration.
type TransactionSimulationAccountSummaryTracesErc721AssetTraceLabels string

const (
	TransactionSimulationAccountSummaryTracesErc721AssetTraceLabelsGasFee TransactionSimulationAccountSummaryTracesErc721AssetTraceLabels = "GAS_FEE"
)

func (r TransactionSimulationAccountSummaryTracesErc721AssetTraceLabels) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryTracesErc721AssetTraceLabelsGasFee:
		return true
	}
	return false
}

type TransactionSimulationAccountSummaryTracesErc1155AssetTrace struct {
	// Description of the asset in the trace
	Asset TransactionSimulationAccountSummaryTracesErc1155AssetTraceAsset `json:"asset,required"`
	// The difference in value for the asset in the trace
	Diff Erc1155Diff `json:"diff,required"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address,required"`
	// The address where the assets are moved to
	ToAddress string `json:"to_address,required"`
	// type of the trace
	TraceType TransactionSimulationAccountSummaryTracesErc1155AssetTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type TransactionSimulationAccountSummaryTracesErc1155AssetTraceType `json:"type,required"`
	// List of labels that describe the trace
	Labels []TransactionSimulationAccountSummaryTracesErc1155AssetTraceLabels `json:"labels"`
	JSON   transactionSimulationAccountSummaryTracesErc1155AssetTraceJSON     `json:"-"`
}

// transactionSimulationAccountSummaryTracesErc1155AssetTraceJSON contains the JSON
// metadata for the struct
// [TransactionSimulationAccountSummaryTracesErc1155AssetTrace]
type transactionSimulationAccountSummaryTracesErc1155AssetTraceJSON struct {
	Asset       apijson.Field
	Diff        apijson.Field
	FromAddress apijson.Field
	ToAddress   apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	Labels      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationAccountSummaryTracesErc1155AssetTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAccountSummaryTracesErc1155AssetTraceJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationAccountSummaryTracesErc1155AssetTrace) implementsTransactionSimulationAccountSummaryTrace() {
}

// Description of the asset in the trace
type TransactionSimulationAccountSummaryTracesErc1155AssetTraceAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type TransactionSimulationAccountSummaryTracesErc1155AssetTraceAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                              `json:"symbol"`
	JSON   transactionSimulationAccountSummaryTracesErc1155AssetTraceAssetJSON `json:"-"`
	union  TransactionSimulationAccountSummaryTracesErc1155AssetTraceAssetUnion
}

// transactionSimulationAccountSummaryTracesErc1155AssetTraceAssetJSON contains the
// JSON metadata for the struct
// [TransactionSimulationAccountSummaryTracesErc1155AssetTraceAsset]
type transactionSimulationAccountSummaryTracesErc1155AssetTraceAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionSimulationAccountSummaryTracesErc1155AssetTraceAssetJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationAccountSummaryTracesErc1155AssetTraceAsset) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationAccountSummaryTracesErc1155AssetTraceAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [TransactionSimulationAccountSummaryTracesErc1155AssetTraceAssetUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [Erc1155TokenDetails],
// [NonercTokenDetails].
func (r TransactionSimulationAccountSummaryTracesErc1155AssetTraceAsset) AsUnion() TransactionSimulationAccountSummaryTracesErc1155AssetTraceAssetUnion {
	return r.union
}

// Description of the asset in the trace
//
// Union satisfied by [Erc1155TokenDetails] or [NonercTokenDetails].
type TransactionSimulationAccountSummaryTracesErc1155AssetTraceAssetUnion interface {
	implementsTransactionSimulationAccountSummaryTracesErc1155AssetTraceAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationAccountSummaryTracesErc1155AssetTraceAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NonercTokenDetails{}),
		},
	)
}

// asset type.
type TransactionSimulationAccountSummaryTracesErc1155AssetTraceAssetType string

const (
	TransactionSimulationAccountSummaryTracesErc1155AssetTraceAssetTypeErc1155 TransactionSimulationAccountSummaryTracesErc1155AssetTraceAssetType = "ERC1155"
	TransactionSimulationAccountSummaryTracesErc1155AssetTraceAssetTypeNonerc  TransactionSimulationAccountSummaryTracesErc1155AssetTraceAssetType = "NONERC"
)

func (r TransactionSimulationAccountSummaryTracesErc1155AssetTraceAssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryTracesErc1155AssetTraceAssetTypeErc1155, TransactionSimulationAccountSummaryTracesErc1155AssetTraceAssetTypeNonerc:
		return true
	}
	return false
}

// type of the trace
type TransactionSimulationAccountSummaryTracesErc1155AssetTraceTraceType string

const (
	TransactionSimulationAccountSummaryTracesErc1155AssetTraceTraceTypeAssetTrace TransactionSimulationAccountSummaryTracesErc1155AssetTraceTraceType = "AssetTrace"
)

func (r TransactionSimulationAccountSummaryTracesErc1155AssetTraceTraceType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryTracesErc1155AssetTraceTraceTypeAssetTrace:
		return true
	}
	return false
}

// The type of the model
type TransactionSimulationAccountSummaryTracesErc1155AssetTraceType string

const (
	TransactionSimulationAccountSummaryTracesErc1155AssetTraceTypeErc1155AssetTrace TransactionSimulationAccountSummaryTracesErc1155AssetTraceType = "ERC1155AssetTrace"
)

func (r TransactionSimulationAccountSummaryTracesErc1155AssetTraceType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryTracesErc1155AssetTraceTypeErc1155AssetTrace:
		return true
	}
	return false
}

// An enumeration.
type TransactionSimulationAccountSummaryTracesErc1155AssetTraceLabels string

const (
	TransactionSimulationAccountSummaryTracesErc1155AssetTraceLabelsGasFee TransactionSimulationAccountSummaryTracesErc1155AssetTraceLabels = "GAS_FEE"
)

func (r TransactionSimulationAccountSummaryTracesErc1155AssetTraceLabels) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryTracesErc1155AssetTraceLabelsGasFee:
		return true
	}
	return false
}

type TransactionSimulationAccountSummaryTracesNativeAssetTrace struct {
	// Description of the asset in the trace
	Asset NativeAssetDetails `json:"asset,required"`
	// The difference in value for the asset in the trace
	Diff NativeDiff `json:"diff,required"`
	// The address where the assets are moved from
	FromAddress string `json:"from_address,required"`
	// The address where the assets are moved to
	ToAddress string `json:"to_address,required"`
	// type of the trace
	TraceType TransactionSimulationAccountSummaryTracesNativeAssetTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type TransactionSimulationAccountSummaryTracesNativeAssetTraceType `json:"type,required"`
	// List of labels that describe the trace
	Labels []TransactionSimulationAccountSummaryTracesNativeAssetTraceLabels `json:"labels"`
	JSON   transactionSimulationAccountSummaryTracesNativeAssetTraceJSON     `json:"-"`
}

// transactionSimulationAccountSummaryTracesNativeAssetTraceJSON contains the JSON
// metadata for the struct
// [TransactionSimulationAccountSummaryTracesNativeAssetTrace]
type transactionSimulationAccountSummaryTracesNativeAssetTraceJSON struct {
	Asset       apijson.Field
	Diff        apijson.Field
	FromAddress apijson.Field
	ToAddress   apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	Labels      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationAccountSummaryTracesNativeAssetTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAccountSummaryTracesNativeAssetTraceJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationAccountSummaryTracesNativeAssetTrace) implementsTransactionSimulationAccountSummaryTrace() {
}

// type of the trace
type TransactionSimulationAccountSummaryTracesNativeAssetTraceTraceType string

const (
	TransactionSimulationAccountSummaryTracesNativeAssetTraceTraceTypeAssetTrace TransactionSimulationAccountSummaryTracesNativeAssetTraceTraceType = "AssetTrace"
)

func (r TransactionSimulationAccountSummaryTracesNativeAssetTraceTraceType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryTracesNativeAssetTraceTraceTypeAssetTrace:
		return true
	}
	return false
}

// The type of the model
type TransactionSimulationAccountSummaryTracesNativeAssetTraceType string

const (
	TransactionSimulationAccountSummaryTracesNativeAssetTraceTypeNativeAssetTrace TransactionSimulationAccountSummaryTracesNativeAssetTraceType = "NativeAssetTrace"
)

func (r TransactionSimulationAccountSummaryTracesNativeAssetTraceType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryTracesNativeAssetTraceTypeNativeAssetTrace:
		return true
	}
	return false
}

// An enumeration.
type TransactionSimulationAccountSummaryTracesNativeAssetTraceLabels string

const (
	TransactionSimulationAccountSummaryTracesNativeAssetTraceLabelsGasFee TransactionSimulationAccountSummaryTracesNativeAssetTraceLabels = "GAS_FEE"
)

func (r TransactionSimulationAccountSummaryTracesNativeAssetTraceLabels) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryTracesNativeAssetTraceLabelsGasFee:
		return true
	}
	return false
}

type TransactionSimulationAccountSummaryTracesErc20ExposureTrace struct {
	Exposed TransactionSimulationAccountSummaryTracesErc20ExposureTraceExposed `json:"exposed,required"`
	// The owner of the assets
	Owner string `json:"owner,required"`
	// The spender of the assets
	Spender string `json:"spender,required"`
	// type of the trace
	TraceType TransactionSimulationAccountSummaryTracesErc20ExposureTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type TransactionSimulationAccountSummaryTracesErc20ExposureTraceType `json:"type,required"`
	JSON transactionSimulationAccountSummaryTracesErc20ExposureTraceJSON `json:"-"`
}

// transactionSimulationAccountSummaryTracesErc20ExposureTraceJSON contains the
// JSON metadata for the struct
// [TransactionSimulationAccountSummaryTracesErc20ExposureTrace]
type transactionSimulationAccountSummaryTracesErc20ExposureTraceJSON struct {
	Exposed     apijson.Field
	Owner       apijson.Field
	Spender     apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationAccountSummaryTracesErc20ExposureTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAccountSummaryTracesErc20ExposureTraceJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationAccountSummaryTracesErc20ExposureTrace) implementsTransactionSimulationAccountSummaryTrace() {
}

type TransactionSimulationAccountSummaryTracesErc20ExposureTraceExposed struct {
	RawValue string                                                                 `json:"raw_value,required"`
	UsdPrice float64                                                                `json:"usd_price"`
	Value    float64                                                                `json:"value"`
	JSON     transactionSimulationAccountSummaryTracesErc20ExposureTraceExposedJSON `json:"-"`
}

// transactionSimulationAccountSummaryTracesErc20ExposureTraceExposedJSON contains
// the JSON metadata for the struct
// [TransactionSimulationAccountSummaryTracesErc20ExposureTraceExposed]
type transactionSimulationAccountSummaryTracesErc20ExposureTraceExposedJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationAccountSummaryTracesErc20ExposureTraceExposed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAccountSummaryTracesErc20ExposureTraceExposedJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type TransactionSimulationAccountSummaryTracesErc20ExposureTraceTraceType string

const (
	TransactionSimulationAccountSummaryTracesErc20ExposureTraceTraceTypeExposureTrace TransactionSimulationAccountSummaryTracesErc20ExposureTraceTraceType = "ExposureTrace"
)

func (r TransactionSimulationAccountSummaryTracesErc20ExposureTraceTraceType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryTracesErc20ExposureTraceTraceTypeExposureTrace:
		return true
	}
	return false
}

// The type of the model
type TransactionSimulationAccountSummaryTracesErc20ExposureTraceType string

const (
	TransactionSimulationAccountSummaryTracesErc20ExposureTraceTypeErc20ExposureTrace TransactionSimulationAccountSummaryTracesErc20ExposureTraceType = "ERC20ExposureTrace"
)

func (r TransactionSimulationAccountSummaryTracesErc20ExposureTraceType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryTracesErc20ExposureTraceTypeErc20ExposureTrace:
		return true
	}
	return false
}

type TransactionSimulationAccountSummaryTracesErc721ExposureTrace struct {
	Exposed TransactionSimulationAccountSummaryTracesErc721ExposureTraceExposed `json:"exposed,required"`
	// The owner of the assets
	Owner string `json:"owner,required"`
	// The spender of the assets
	Spender string `json:"spender,required"`
	// type of the trace
	TraceType TransactionSimulationAccountSummaryTracesErc721ExposureTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type TransactionSimulationAccountSummaryTracesErc721ExposureTraceType `json:"type,required"`
	JSON transactionSimulationAccountSummaryTracesErc721ExposureTraceJSON `json:"-"`
}

// transactionSimulationAccountSummaryTracesErc721ExposureTraceJSON contains the
// JSON metadata for the struct
// [TransactionSimulationAccountSummaryTracesErc721ExposureTrace]
type transactionSimulationAccountSummaryTracesErc721ExposureTraceJSON struct {
	Exposed     apijson.Field
	Owner       apijson.Field
	Spender     apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationAccountSummaryTracesErc721ExposureTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAccountSummaryTracesErc721ExposureTraceJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationAccountSummaryTracesErc721ExposureTrace) implementsTransactionSimulationAccountSummaryTrace() {
}

type TransactionSimulationAccountSummaryTracesErc721ExposureTraceExposed struct {
	Amount   int64                                                                   `json:"amount,required"`
	TokenID  string                                                                  `json:"token_id,required"`
	IsMint   bool                                                                    `json:"is_mint"`
	LogoURL  string                                                                  `json:"logo_url"`
	UsdPrice float64                                                                 `json:"usd_price"`
	JSON     transactionSimulationAccountSummaryTracesErc721ExposureTraceExposedJSON `json:"-"`
}

// transactionSimulationAccountSummaryTracesErc721ExposureTraceExposedJSON contains
// the JSON metadata for the struct
// [TransactionSimulationAccountSummaryTracesErc721ExposureTraceExposed]
type transactionSimulationAccountSummaryTracesErc721ExposureTraceExposedJSON struct {
	Amount      apijson.Field
	TokenID     apijson.Field
	IsMint      apijson.Field
	LogoURL     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationAccountSummaryTracesErc721ExposureTraceExposed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAccountSummaryTracesErc721ExposureTraceExposedJSON) RawJSON() string {
	return r.raw
}

// type of the trace
type TransactionSimulationAccountSummaryTracesErc721ExposureTraceTraceType string

const (
	TransactionSimulationAccountSummaryTracesErc721ExposureTraceTraceTypeExposureTrace TransactionSimulationAccountSummaryTracesErc721ExposureTraceTraceType = "ExposureTrace"
)

func (r TransactionSimulationAccountSummaryTracesErc721ExposureTraceTraceType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryTracesErc721ExposureTraceTraceTypeExposureTrace:
		return true
	}
	return false
}

// The type of the model
type TransactionSimulationAccountSummaryTracesErc721ExposureTraceType string

const (
	TransactionSimulationAccountSummaryTracesErc721ExposureTraceTypeErc721ExposureTrace TransactionSimulationAccountSummaryTracesErc721ExposureTraceType = "ERC721ExposureTrace"
)

func (r TransactionSimulationAccountSummaryTracesErc721ExposureTraceType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryTracesErc721ExposureTraceTypeErc721ExposureTrace:
		return true
	}
	return false
}

type TransactionSimulationAccountSummaryTracesErc1155ExposureTrace struct {
	// The owner of the assets
	Owner string `json:"owner,required"`
	// The spender of the assets
	Spender string `json:"spender,required"`
	// type of the trace
	TraceType TransactionSimulationAccountSummaryTracesErc1155ExposureTraceTraceType `json:"trace_type,required"`
	// The type of the model
	Type TransactionSimulationAccountSummaryTracesErc1155ExposureTraceType `json:"type,required"`
	JSON transactionSimulationAccountSummaryTracesErc1155ExposureTraceJSON `json:"-"`
}

// transactionSimulationAccountSummaryTracesErc1155ExposureTraceJSON contains the
// JSON metadata for the struct
// [TransactionSimulationAccountSummaryTracesErc1155ExposureTrace]
type transactionSimulationAccountSummaryTracesErc1155ExposureTraceJSON struct {
	Owner       apijson.Field
	Spender     apijson.Field
	TraceType   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationAccountSummaryTracesErc1155ExposureTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAccountSummaryTracesErc1155ExposureTraceJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationAccountSummaryTracesErc1155ExposureTrace) implementsTransactionSimulationAccountSummaryTrace() {
}

// type of the trace
type TransactionSimulationAccountSummaryTracesErc1155ExposureTraceTraceType string

const (
	TransactionSimulationAccountSummaryTracesErc1155ExposureTraceTraceTypeExposureTrace TransactionSimulationAccountSummaryTracesErc1155ExposureTraceTraceType = "ExposureTrace"
)

func (r TransactionSimulationAccountSummaryTracesErc1155ExposureTraceTraceType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryTracesErc1155ExposureTraceTraceTypeExposureTrace:
		return true
	}
	return false
}

// The type of the model
type TransactionSimulationAccountSummaryTracesErc1155ExposureTraceType string

const (
	TransactionSimulationAccountSummaryTracesErc1155ExposureTraceTypeErc1155ExposureTrace TransactionSimulationAccountSummaryTracesErc1155ExposureTraceType = "ERC1155ExposureTrace"
)

func (r TransactionSimulationAccountSummaryTracesErc1155ExposureTraceType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryTracesErc1155ExposureTraceTypeErc1155ExposureTrace:
		return true
	}
	return false
}

// type of the trace
type TransactionSimulationAccountSummaryTracesTraceType string

const (
	TransactionSimulationAccountSummaryTracesTraceTypeAssetTrace    TransactionSimulationAccountSummaryTracesTraceType = "AssetTrace"
	TransactionSimulationAccountSummaryTracesTraceTypeExposureTrace TransactionSimulationAccountSummaryTracesTraceType = "ExposureTrace"
)

func (r TransactionSimulationAccountSummaryTracesTraceType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryTracesTraceTypeAssetTrace, TransactionSimulationAccountSummaryTracesTraceTypeExposureTrace:
		return true
	}
	return false
}

// The type of the model
type TransactionSimulationAccountSummaryTracesType string

const (
	TransactionSimulationAccountSummaryTracesTypeErc20AssetTrace      TransactionSimulationAccountSummaryTracesType = "ERC20AssetTrace"
	TransactionSimulationAccountSummaryTracesTypeErc721AssetTrace     TransactionSimulationAccountSummaryTracesType = "ERC721AssetTrace"
	TransactionSimulationAccountSummaryTracesTypeErc1155AssetTrace    TransactionSimulationAccountSummaryTracesType = "ERC1155AssetTrace"
	TransactionSimulationAccountSummaryTracesTypeNativeAssetTrace     TransactionSimulationAccountSummaryTracesType = "NativeAssetTrace"
	TransactionSimulationAccountSummaryTracesTypeErc20ExposureTrace   TransactionSimulationAccountSummaryTracesType = "ERC20ExposureTrace"
	TransactionSimulationAccountSummaryTracesTypeErc721ExposureTrace  TransactionSimulationAccountSummaryTracesType = "ERC721ExposureTrace"
	TransactionSimulationAccountSummaryTracesTypeErc1155ExposureTrace TransactionSimulationAccountSummaryTracesType = "ERC1155ExposureTrace"
)

func (r TransactionSimulationAccountSummaryTracesType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryTracesTypeErc20AssetTrace, TransactionSimulationAccountSummaryTracesTypeErc721AssetTrace, TransactionSimulationAccountSummaryTracesTypeErc1155AssetTrace, TransactionSimulationAccountSummaryTracesTypeNativeAssetTrace, TransactionSimulationAccountSummaryTracesTypeErc20ExposureTrace, TransactionSimulationAccountSummaryTracesTypeErc721ExposureTrace, TransactionSimulationAccountSummaryTracesTypeErc1155ExposureTrace:
		return true
	}
	return false
}

type TransactionSimulationAddressDetail struct {
	// contains the contract's name if the address is a verified contract
	ContractName string `json:"contract_name"`
	// Whether the address is an eoa or a contract
	IsEoa bool `json:"is_eoa"`
	// known name tag for the address
	NameTag string                                 `json:"name_tag"`
	JSON    transactionSimulationAddressDetailJSON `json:"-"`
}

// transactionSimulationAddressDetailJSON contains the JSON metadata for the struct
// [TransactionSimulationAddressDetail]
type transactionSimulationAddressDetailJSON struct {
	ContractName apijson.Field
	IsEoa        apijson.Field
	NameTag      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TransactionSimulationAddressDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAddressDetailJSON) RawJSON() string {
	return r.raw
}

type TransactionSimulationAssetsDiff struct {
	// This field can have the runtime type of
	// [TransactionSimulationAssetsDiffsErc20AddressAssetDiffAsset],
	// [TransactionSimulationAssetsDiffsErc721AddressAssetDiffAsset],
	// [TransactionSimulationAssetsDiffsErc1155AddressAssetDiffAsset],
	// [NativeAssetDetails].
	Asset interface{} `json:"asset,required"`
	// type of the asset for the current diff
	AssetType TransactionSimulationAssetsDiffsAssetType `json:"asset_type,required"`
	// This field can have the runtime type of [[]Erc20Diff], [[]Erc721Diff],
	// [[]Erc1155Diff], [[]NativeDiff].
	In interface{} `json:"in,required"`
	// This field can have the runtime type of [[]Erc20Diff], [[]Erc721Diff],
	// [[]Erc1155Diff], [[]NativeDiff].
	Out   interface{}                         `json:"out,required"`
	JSON  transactionSimulationAssetsDiffJSON `json:"-"`
	union TransactionSimulationAssetsDiffsUnion
}

// transactionSimulationAssetsDiffJSON contains the JSON metadata for the struct
// [TransactionSimulationAssetsDiff]
type transactionSimulationAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionSimulationAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TransactionSimulationAssetsDiffsUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TransactionSimulationAssetsDiffsErc20AddressAssetDiff],
// [TransactionSimulationAssetsDiffsErc721AddressAssetDiff],
// [TransactionSimulationAssetsDiffsErc1155AddressAssetDiff],
// [TransactionSimulationAssetsDiffsNativeAddressAssetDiff].
func (r TransactionSimulationAssetsDiff) AsUnion() TransactionSimulationAssetsDiffsUnion {
	return r.union
}

// Union satisfied by [TransactionSimulationAssetsDiffsErc20AddressAssetDiff],
// [TransactionSimulationAssetsDiffsErc721AddressAssetDiff],
// [TransactionSimulationAssetsDiffsErc1155AddressAssetDiff] or
// [TransactionSimulationAssetsDiffsNativeAddressAssetDiff].
type TransactionSimulationAssetsDiffsUnion interface {
	implementsTransactionSimulationAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationAssetsDiffsErc20AddressAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationAssetsDiffsErc721AddressAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationAssetsDiffsErc1155AddressAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationAssetsDiffsNativeAddressAssetDiff{}),
		},
	)
}

type TransactionSimulationAssetsDiffsErc20AddressAssetDiff struct {
	// description of the asset for the current diff
	Asset TransactionSimulationAssetsDiffsErc20AddressAssetDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType TransactionSimulationAssetsDiffsErc20AddressAssetDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []Erc20Diff `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []Erc20Diff                                               `json:"out,required"`
	JSON transactionSimulationAssetsDiffsErc20AddressAssetDiffJSON `json:"-"`
}

// transactionSimulationAssetsDiffsErc20AddressAssetDiffJSON contains the JSON
// metadata for the struct [TransactionSimulationAssetsDiffsErc20AddressAssetDiff]
type transactionSimulationAssetsDiffsErc20AddressAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationAssetsDiffsErc20AddressAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAssetsDiffsErc20AddressAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationAssetsDiffsErc20AddressAssetDiff) implementsTransactionSimulationAssetsDiff() {
}

// description of the asset for the current diff
type TransactionSimulationAssetsDiffsErc20AddressAssetDiffAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type TransactionSimulationAssetsDiffsErc20AddressAssetDiffAssetType `json:"type,required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                         `json:"symbol"`
	JSON   transactionSimulationAssetsDiffsErc20AddressAssetDiffAssetJSON `json:"-"`
	union  TransactionSimulationAssetsDiffsErc20AddressAssetDiffAssetUnion
}

// transactionSimulationAssetsDiffsErc20AddressAssetDiffAssetJSON contains the JSON
// metadata for the struct
// [TransactionSimulationAssetsDiffsErc20AddressAssetDiffAsset]
type transactionSimulationAssetsDiffsErc20AddressAssetDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionSimulationAssetsDiffsErc20AddressAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationAssetsDiffsErc20AddressAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationAssetsDiffsErc20AddressAssetDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [TransactionSimulationAssetsDiffsErc20AddressAssetDiffAssetUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [Erc20TokenDetails],
// [NonercTokenDetails].
func (r TransactionSimulationAssetsDiffsErc20AddressAssetDiffAsset) AsUnion() TransactionSimulationAssetsDiffsErc20AddressAssetDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by [Erc20TokenDetails] or [NonercTokenDetails].
type TransactionSimulationAssetsDiffsErc20AddressAssetDiffAssetUnion interface {
	implementsTransactionSimulationAssetsDiffsErc20AddressAssetDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationAssetsDiffsErc20AddressAssetDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NonercTokenDetails{}),
		},
	)
}

// asset type.
type TransactionSimulationAssetsDiffsErc20AddressAssetDiffAssetType string

const (
	TransactionSimulationAssetsDiffsErc20AddressAssetDiffAssetTypeErc20  TransactionSimulationAssetsDiffsErc20AddressAssetDiffAssetType = "ERC20"
	TransactionSimulationAssetsDiffsErc20AddressAssetDiffAssetTypeNonerc TransactionSimulationAssetsDiffsErc20AddressAssetDiffAssetType = "NONERC"
)

func (r TransactionSimulationAssetsDiffsErc20AddressAssetDiffAssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationAssetsDiffsErc20AddressAssetDiffAssetTypeErc20, TransactionSimulationAssetsDiffsErc20AddressAssetDiffAssetTypeNonerc:
		return true
	}
	return false
}

type TransactionSimulationAssetsDiffsErc721AddressAssetDiff struct {
	// description of the asset for the current diff
	Asset TransactionSimulationAssetsDiffsErc721AddressAssetDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType TransactionSimulationAssetsDiffsErc721AddressAssetDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []Erc721Diff `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []Erc721Diff                                               `json:"out,required"`
	JSON transactionSimulationAssetsDiffsErc721AddressAssetDiffJSON `json:"-"`
}

// transactionSimulationAssetsDiffsErc721AddressAssetDiffJSON contains the JSON
// metadata for the struct [TransactionSimulationAssetsDiffsErc721AddressAssetDiff]
type transactionSimulationAssetsDiffsErc721AddressAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationAssetsDiffsErc721AddressAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAssetsDiffsErc721AddressAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationAssetsDiffsErc721AddressAssetDiff) implementsTransactionSimulationAssetsDiff() {
}

// description of the asset for the current diff
type TransactionSimulationAssetsDiffsErc721AddressAssetDiffAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type TransactionSimulationAssetsDiffsErc721AddressAssetDiffAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                          `json:"symbol"`
	JSON   transactionSimulationAssetsDiffsErc721AddressAssetDiffAssetJSON `json:"-"`
	union  TransactionSimulationAssetsDiffsErc721AddressAssetDiffAssetUnion
}

// transactionSimulationAssetsDiffsErc721AddressAssetDiffAssetJSON contains the
// JSON metadata for the struct
// [TransactionSimulationAssetsDiffsErc721AddressAssetDiffAsset]
type transactionSimulationAssetsDiffsErc721AddressAssetDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionSimulationAssetsDiffsErc721AddressAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationAssetsDiffsErc721AddressAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationAssetsDiffsErc721AddressAssetDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [TransactionSimulationAssetsDiffsErc721AddressAssetDiffAssetUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [Erc721TokenDetails],
// [NonercTokenDetails].
func (r TransactionSimulationAssetsDiffsErc721AddressAssetDiffAsset) AsUnion() TransactionSimulationAssetsDiffsErc721AddressAssetDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by [Erc721TokenDetails] or [NonercTokenDetails].
type TransactionSimulationAssetsDiffsErc721AddressAssetDiffAssetUnion interface {
	implementsTransactionSimulationAssetsDiffsErc721AddressAssetDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationAssetsDiffsErc721AddressAssetDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NonercTokenDetails{}),
		},
	)
}

// asset type.
type TransactionSimulationAssetsDiffsErc721AddressAssetDiffAssetType string

const (
	TransactionSimulationAssetsDiffsErc721AddressAssetDiffAssetTypeErc721 TransactionSimulationAssetsDiffsErc721AddressAssetDiffAssetType = "ERC721"
	TransactionSimulationAssetsDiffsErc721AddressAssetDiffAssetTypeNonerc TransactionSimulationAssetsDiffsErc721AddressAssetDiffAssetType = "NONERC"
)

func (r TransactionSimulationAssetsDiffsErc721AddressAssetDiffAssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationAssetsDiffsErc721AddressAssetDiffAssetTypeErc721, TransactionSimulationAssetsDiffsErc721AddressAssetDiffAssetTypeNonerc:
		return true
	}
	return false
}

type TransactionSimulationAssetsDiffsErc1155AddressAssetDiff struct {
	// description of the asset for the current diff
	Asset TransactionSimulationAssetsDiffsErc1155AddressAssetDiffAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType TransactionSimulationAssetsDiffsErc1155AddressAssetDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []Erc1155Diff `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []Erc1155Diff                                               `json:"out,required"`
	JSON transactionSimulationAssetsDiffsErc1155AddressAssetDiffJSON `json:"-"`
}

// transactionSimulationAssetsDiffsErc1155AddressAssetDiffJSON contains the JSON
// metadata for the struct
// [TransactionSimulationAssetsDiffsErc1155AddressAssetDiff]
type transactionSimulationAssetsDiffsErc1155AddressAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationAssetsDiffsErc1155AddressAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAssetsDiffsErc1155AddressAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationAssetsDiffsErc1155AddressAssetDiff) implementsTransactionSimulationAssetsDiff() {
}

// description of the asset for the current diff
type TransactionSimulationAssetsDiffsErc1155AddressAssetDiffAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type TransactionSimulationAssetsDiffsErc1155AddressAssetDiffAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                           `json:"symbol"`
	JSON   transactionSimulationAssetsDiffsErc1155AddressAssetDiffAssetJSON `json:"-"`
	union  TransactionSimulationAssetsDiffsErc1155AddressAssetDiffAssetUnion
}

// transactionSimulationAssetsDiffsErc1155AddressAssetDiffAssetJSON contains the
// JSON metadata for the struct
// [TransactionSimulationAssetsDiffsErc1155AddressAssetDiffAsset]
type transactionSimulationAssetsDiffsErc1155AddressAssetDiffAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionSimulationAssetsDiffsErc1155AddressAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationAssetsDiffsErc1155AddressAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationAssetsDiffsErc1155AddressAssetDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [TransactionSimulationAssetsDiffsErc1155AddressAssetDiffAssetUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [Erc1155TokenDetails],
// [NonercTokenDetails].
func (r TransactionSimulationAssetsDiffsErc1155AddressAssetDiffAsset) AsUnion() TransactionSimulationAssetsDiffsErc1155AddressAssetDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by [Erc1155TokenDetails] or [NonercTokenDetails].
type TransactionSimulationAssetsDiffsErc1155AddressAssetDiffAssetUnion interface {
	implementsTransactionSimulationAssetsDiffsErc1155AddressAssetDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationAssetsDiffsErc1155AddressAssetDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NonercTokenDetails{}),
		},
	)
}

// asset type.
type TransactionSimulationAssetsDiffsErc1155AddressAssetDiffAssetType string

const (
	TransactionSimulationAssetsDiffsErc1155AddressAssetDiffAssetTypeErc1155 TransactionSimulationAssetsDiffsErc1155AddressAssetDiffAssetType = "ERC1155"
	TransactionSimulationAssetsDiffsErc1155AddressAssetDiffAssetTypeNonerc  TransactionSimulationAssetsDiffsErc1155AddressAssetDiffAssetType = "NONERC"
)

func (r TransactionSimulationAssetsDiffsErc1155AddressAssetDiffAssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationAssetsDiffsErc1155AddressAssetDiffAssetTypeErc1155, TransactionSimulationAssetsDiffsErc1155AddressAssetDiffAssetTypeNonerc:
		return true
	}
	return false
}

type TransactionSimulationAssetsDiffsNativeAddressAssetDiff struct {
	// description of the asset for the current diff
	Asset NativeAssetDetails `json:"asset,required"`
	// type of the asset for the current diff
	AssetType TransactionSimulationAssetsDiffsNativeAddressAssetDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []NativeDiff `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []NativeDiff                                               `json:"out,required"`
	JSON transactionSimulationAssetsDiffsNativeAddressAssetDiffJSON `json:"-"`
}

// transactionSimulationAssetsDiffsNativeAddressAssetDiffJSON contains the JSON
// metadata for the struct [TransactionSimulationAssetsDiffsNativeAddressAssetDiff]
type transactionSimulationAssetsDiffsNativeAddressAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationAssetsDiffsNativeAddressAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAssetsDiffsNativeAddressAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationAssetsDiffsNativeAddressAssetDiff) implementsTransactionSimulationAssetsDiff() {
}

// type of the asset for the current diff
type TransactionSimulationAssetsDiffsNativeAddressAssetDiffAssetType string

const (
	TransactionSimulationAssetsDiffsNativeAddressAssetDiffAssetTypeNative TransactionSimulationAssetsDiffsNativeAddressAssetDiffAssetType = "NATIVE"
)

func (r TransactionSimulationAssetsDiffsNativeAddressAssetDiffAssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationAssetsDiffsNativeAddressAssetDiffAssetTypeNative:
		return true
	}
	return false
}

// type of the asset for the current diff
type TransactionSimulationAssetsDiffsAssetType string

const (
	TransactionSimulationAssetsDiffsAssetTypeErc20   TransactionSimulationAssetsDiffsAssetType = "ERC20"
	TransactionSimulationAssetsDiffsAssetTypeErc721  TransactionSimulationAssetsDiffsAssetType = "ERC721"
	TransactionSimulationAssetsDiffsAssetTypeErc1155 TransactionSimulationAssetsDiffsAssetType = "ERC1155"
	TransactionSimulationAssetsDiffsAssetTypeNative  TransactionSimulationAssetsDiffsAssetType = "NATIVE"
)

func (r TransactionSimulationAssetsDiffsAssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationAssetsDiffsAssetTypeErc20, TransactionSimulationAssetsDiffsAssetTypeErc721, TransactionSimulationAssetsDiffsAssetTypeErc1155, TransactionSimulationAssetsDiffsAssetTypeNative:
		return true
	}
	return false
}

type TransactionSimulationExposure struct {
	// This field can have the runtime type of
	// [TransactionSimulationExposuresErc20AddressExposureAsset],
	// [TransactionSimulationExposuresErc721AddressExposureAsset],
	// [TransactionSimulationExposuresErc1155AddressExposureAsset].
	Asset interface{} `json:"asset,required"`
	// type of the asset for the current diff
	AssetType TransactionSimulationExposuresAssetType `json:"asset_type,required"`
	// This field can have the runtime type of [map[string]Erc20Exposure],
	// [map[string]Erc721Exposure], [map[string]Erc1155Exposure].
	Spenders interface{}                       `json:"spenders,required"`
	JSON     transactionSimulationExposureJSON `json:"-"`
	union    TransactionSimulationExposuresUnion
}

// transactionSimulationExposureJSON contains the JSON metadata for the struct
// [TransactionSimulationExposure]
type transactionSimulationExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionSimulationExposureJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationExposure) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TransactionSimulationExposuresUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TransactionSimulationExposuresErc20AddressExposure],
// [TransactionSimulationExposuresErc721AddressExposure],
// [TransactionSimulationExposuresErc1155AddressExposure].
func (r TransactionSimulationExposure) AsUnion() TransactionSimulationExposuresUnion {
	return r.union
}

// Union satisfied by [TransactionSimulationExposuresErc20AddressExposure],
// [TransactionSimulationExposuresErc721AddressExposure] or
// [TransactionSimulationExposuresErc1155AddressExposure].
type TransactionSimulationExposuresUnion interface {
	implementsTransactionSimulationExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationExposuresErc20AddressExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationExposuresErc721AddressExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationExposuresErc1155AddressExposure{}),
		},
	)
}

type TransactionSimulationExposuresErc20AddressExposure struct {
	// description of the asset for the current diff
	Asset TransactionSimulationExposuresErc20AddressExposureAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType TransactionSimulationExposuresErc20AddressExposureAssetType `json:"asset_type,required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]Erc20Exposure                               `json:"spenders,required"`
	JSON     transactionSimulationExposuresErc20AddressExposureJSON `json:"-"`
}

// transactionSimulationExposuresErc20AddressExposureJSON contains the JSON
// metadata for the struct [TransactionSimulationExposuresErc20AddressExposure]
type transactionSimulationExposuresErc20AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationExposuresErc20AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationExposuresErc20AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationExposuresErc20AddressExposure) implementsTransactionSimulationExposure() {
}

// description of the asset for the current diff
type TransactionSimulationExposuresErc20AddressExposureAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type TransactionSimulationExposuresErc20AddressExposureAssetType `json:"type,required"`
	// asset's decimals
	Decimals int64 `json:"decimals"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                      `json:"symbol"`
	JSON   transactionSimulationExposuresErc20AddressExposureAssetJSON `json:"-"`
	union  TransactionSimulationExposuresErc20AddressExposureAssetUnion
}

// transactionSimulationExposuresErc20AddressExposureAssetJSON contains the JSON
// metadata for the struct
// [TransactionSimulationExposuresErc20AddressExposureAsset]
type transactionSimulationExposuresErc20AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionSimulationExposuresErc20AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationExposuresErc20AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationExposuresErc20AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TransactionSimulationExposuresErc20AddressExposureAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [Erc20TokenDetails],
// [NonercTokenDetails].
func (r TransactionSimulationExposuresErc20AddressExposureAsset) AsUnion() TransactionSimulationExposuresErc20AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by [Erc20TokenDetails] or [NonercTokenDetails].
type TransactionSimulationExposuresErc20AddressExposureAssetUnion interface {
	implementsTransactionSimulationExposuresErc20AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationExposuresErc20AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NonercTokenDetails{}),
		},
	)
}

// asset type.
type TransactionSimulationExposuresErc20AddressExposureAssetType string

const (
	TransactionSimulationExposuresErc20AddressExposureAssetTypeErc20  TransactionSimulationExposuresErc20AddressExposureAssetType = "ERC20"
	TransactionSimulationExposuresErc20AddressExposureAssetTypeNonerc TransactionSimulationExposuresErc20AddressExposureAssetType = "NONERC"
)

func (r TransactionSimulationExposuresErc20AddressExposureAssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationExposuresErc20AddressExposureAssetTypeErc20, TransactionSimulationExposuresErc20AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type TransactionSimulationExposuresErc721AddressExposure struct {
	// description of the asset for the current diff
	Asset TransactionSimulationExposuresErc721AddressExposureAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType TransactionSimulationExposuresErc721AddressExposureAssetType `json:"asset_type,required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]Erc721Exposure                               `json:"spenders,required"`
	JSON     transactionSimulationExposuresErc721AddressExposureJSON `json:"-"`
}

// transactionSimulationExposuresErc721AddressExposureJSON contains the JSON
// metadata for the struct [TransactionSimulationExposuresErc721AddressExposure]
type transactionSimulationExposuresErc721AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationExposuresErc721AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationExposuresErc721AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationExposuresErc721AddressExposure) implementsTransactionSimulationExposure() {
}

// description of the asset for the current diff
type TransactionSimulationExposuresErc721AddressExposureAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type TransactionSimulationExposuresErc721AddressExposureAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                       `json:"symbol"`
	JSON   transactionSimulationExposuresErc721AddressExposureAssetJSON `json:"-"`
	union  TransactionSimulationExposuresErc721AddressExposureAssetUnion
}

// transactionSimulationExposuresErc721AddressExposureAssetJSON contains the JSON
// metadata for the struct
// [TransactionSimulationExposuresErc721AddressExposureAsset]
type transactionSimulationExposuresErc721AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionSimulationExposuresErc721AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationExposuresErc721AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationExposuresErc721AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [TransactionSimulationExposuresErc721AddressExposureAssetUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [Erc721TokenDetails],
// [NonercTokenDetails].
func (r TransactionSimulationExposuresErc721AddressExposureAsset) AsUnion() TransactionSimulationExposuresErc721AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by [Erc721TokenDetails] or [NonercTokenDetails].
type TransactionSimulationExposuresErc721AddressExposureAssetUnion interface {
	implementsTransactionSimulationExposuresErc721AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationExposuresErc721AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NonercTokenDetails{}),
		},
	)
}

// asset type.
type TransactionSimulationExposuresErc721AddressExposureAssetType string

const (
	TransactionSimulationExposuresErc721AddressExposureAssetTypeErc721 TransactionSimulationExposuresErc721AddressExposureAssetType = "ERC721"
	TransactionSimulationExposuresErc721AddressExposureAssetTypeNonerc TransactionSimulationExposuresErc721AddressExposureAssetType = "NONERC"
)

func (r TransactionSimulationExposuresErc721AddressExposureAssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationExposuresErc721AddressExposureAssetTypeErc721, TransactionSimulationExposuresErc721AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

type TransactionSimulationExposuresErc1155AddressExposure struct {
	// description of the asset for the current diff
	Asset TransactionSimulationExposuresErc1155AddressExposureAsset `json:"asset,required"`
	// type of the asset for the current diff
	AssetType TransactionSimulationExposuresErc1155AddressExposureAssetType `json:"asset_type,required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]Erc1155Exposure                               `json:"spenders,required"`
	JSON     transactionSimulationExposuresErc1155AddressExposureJSON `json:"-"`
}

// transactionSimulationExposuresErc1155AddressExposureJSON contains the JSON
// metadata for the struct [TransactionSimulationExposuresErc1155AddressExposure]
type transactionSimulationExposuresErc1155AddressExposureJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationExposuresErc1155AddressExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationExposuresErc1155AddressExposureJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationExposuresErc1155AddressExposure) implementsTransactionSimulationExposure() {
}

// description of the asset for the current diff
type TransactionSimulationExposuresErc1155AddressExposureAsset struct {
	// address of the token
	Address string `json:"address,required"`
	// asset type.
	Type TransactionSimulationExposuresErc1155AddressExposureAssetType `json:"type,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string                                                        `json:"symbol"`
	JSON   transactionSimulationExposuresErc1155AddressExposureAssetJSON `json:"-"`
	union  TransactionSimulationExposuresErc1155AddressExposureAssetUnion
}

// transactionSimulationExposuresErc1155AddressExposureAssetJSON contains the JSON
// metadata for the struct
// [TransactionSimulationExposuresErc1155AddressExposureAsset]
type transactionSimulationExposuresErc1155AddressExposureAssetJSON struct {
	Address     apijson.Field
	Type        apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionSimulationExposuresErc1155AddressExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationExposuresErc1155AddressExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationExposuresErc1155AddressExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [TransactionSimulationExposuresErc1155AddressExposureAssetUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [Erc1155TokenDetails],
// [NonercTokenDetails].
func (r TransactionSimulationExposuresErc1155AddressExposureAsset) AsUnion() TransactionSimulationExposuresErc1155AddressExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by [Erc1155TokenDetails] or [NonercTokenDetails].
type TransactionSimulationExposuresErc1155AddressExposureAssetUnion interface {
	implementsTransactionSimulationExposuresErc1155AddressExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationExposuresErc1155AddressExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NonercTokenDetails{}),
		},
	)
}

// asset type.
type TransactionSimulationExposuresErc1155AddressExposureAssetType string

const (
	TransactionSimulationExposuresErc1155AddressExposureAssetTypeErc1155 TransactionSimulationExposuresErc1155AddressExposureAssetType = "ERC1155"
	TransactionSimulationExposuresErc1155AddressExposureAssetTypeNonerc  TransactionSimulationExposuresErc1155AddressExposureAssetType = "NONERC"
)

func (r TransactionSimulationExposuresErc1155AddressExposureAssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationExposuresErc1155AddressExposureAssetTypeErc1155, TransactionSimulationExposuresErc1155AddressExposureAssetTypeNonerc:
		return true
	}
	return false
}

// type of the asset for the current diff
type TransactionSimulationExposuresAssetType string

const (
	TransactionSimulationExposuresAssetTypeErc20   TransactionSimulationExposuresAssetType = "ERC20"
	TransactionSimulationExposuresAssetTypeErc721  TransactionSimulationExposuresAssetType = "ERC721"
	TransactionSimulationExposuresAssetTypeErc1155 TransactionSimulationExposuresAssetType = "ERC1155"
)

func (r TransactionSimulationExposuresAssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationExposuresAssetTypeErc20, TransactionSimulationExposuresAssetTypeErc721, TransactionSimulationExposuresAssetTypeErc1155:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type TransactionSimulationStatus string

const (
	TransactionSimulationStatusSuccess TransactionSimulationStatus = "Success"
)

func (r TransactionSimulationStatus) IsKnown() bool {
	switch r {
	case TransactionSimulationStatusSuccess:
		return true
	}
	return false
}

type TransactionSimulationContractManagement struct {
	// This field can have the runtime type of
	// [TransactionSimulationContractManagementProxyUpgradeManagementAfter],
	// [TransactionSimulationContractManagementOwnershipChangeManagementAfter],
	// [TransactionSimulationContractManagementModulesChangeManagementAfter].
	After interface{} `json:"after,required"`
	// This field can have the runtime type of
	// [TransactionSimulationContractManagementProxyUpgradeManagementBefore],
	// [TransactionSimulationContractManagementOwnershipChangeManagementBefore],
	// [TransactionSimulationContractManagementModulesChangeManagementBefore].
	Before interface{} `json:"before,required"`
	// The type of the state change
	Type  TransactionSimulationContractManagementType `json:"type,required"`
	JSON  transactionSimulationContractManagementJSON `json:"-"`
	union TransactionSimulationContractManagementUnion
}

// transactionSimulationContractManagementJSON contains the JSON metadata for the
// struct [TransactionSimulationContractManagement]
type transactionSimulationContractManagementJSON struct {
	After       apijson.Field
	Before      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionSimulationContractManagementJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationContractManagement) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationContractManagement{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TransactionSimulationContractManagementUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TransactionSimulationContractManagementProxyUpgradeManagement],
// [TransactionSimulationContractManagementOwnershipChangeManagement],
// [TransactionSimulationContractManagementModulesChangeManagement].
func (r TransactionSimulationContractManagement) AsUnion() TransactionSimulationContractManagementUnion {
	return r.union
}

// Union satisfied by
// [TransactionSimulationContractManagementProxyUpgradeManagement],
// [TransactionSimulationContractManagementOwnershipChangeManagement] or
// [TransactionSimulationContractManagementModulesChangeManagement].
type TransactionSimulationContractManagementUnion interface {
	implementsTransactionSimulationContractManagement()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationContractManagementUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationContractManagementProxyUpgradeManagement{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationContractManagementOwnershipChangeManagement{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationContractManagementModulesChangeManagement{}),
		},
	)
}

type TransactionSimulationContractManagementProxyUpgradeManagement struct {
	// The state after the transaction
	After TransactionSimulationContractManagementProxyUpgradeManagementAfter `json:"after,required"`
	// The state before the transaction
	Before TransactionSimulationContractManagementProxyUpgradeManagementBefore `json:"before,required"`
	// The type of the state change
	Type TransactionSimulationContractManagementProxyUpgradeManagementType `json:"type,required"`
	JSON transactionSimulationContractManagementProxyUpgradeManagementJSON `json:"-"`
}

// transactionSimulationContractManagementProxyUpgradeManagementJSON contains the
// JSON metadata for the struct
// [TransactionSimulationContractManagementProxyUpgradeManagement]
type transactionSimulationContractManagementProxyUpgradeManagementJSON struct {
	After       apijson.Field
	Before      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationContractManagementProxyUpgradeManagement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationContractManagementProxyUpgradeManagementJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationContractManagementProxyUpgradeManagement) implementsTransactionSimulationContractManagement() {
}

// The state after the transaction
type TransactionSimulationContractManagementProxyUpgradeManagementAfter struct {
	Address string                                                                 `json:"address,required"`
	JSON    transactionSimulationContractManagementProxyUpgradeManagementAfterJSON `json:"-"`
}

// transactionSimulationContractManagementProxyUpgradeManagementAfterJSON contains
// the JSON metadata for the struct
// [TransactionSimulationContractManagementProxyUpgradeManagementAfter]
type transactionSimulationContractManagementProxyUpgradeManagementAfterJSON struct {
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationContractManagementProxyUpgradeManagementAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationContractManagementProxyUpgradeManagementAfterJSON) RawJSON() string {
	return r.raw
}

// The state before the transaction
type TransactionSimulationContractManagementProxyUpgradeManagementBefore struct {
	Address string                                                                  `json:"address,required"`
	JSON    transactionSimulationContractManagementProxyUpgradeManagementBeforeJSON `json:"-"`
}

// transactionSimulationContractManagementProxyUpgradeManagementBeforeJSON contains
// the JSON metadata for the struct
// [TransactionSimulationContractManagementProxyUpgradeManagementBefore]
type transactionSimulationContractManagementProxyUpgradeManagementBeforeJSON struct {
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationContractManagementProxyUpgradeManagementBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationContractManagementProxyUpgradeManagementBeforeJSON) RawJSON() string {
	return r.raw
}

// The type of the state change
type TransactionSimulationContractManagementProxyUpgradeManagementType string

const (
	TransactionSimulationContractManagementProxyUpgradeManagementTypeProxyUpgrade TransactionSimulationContractManagementProxyUpgradeManagementType = "PROXY_UPGRADE"
)

func (r TransactionSimulationContractManagementProxyUpgradeManagementType) IsKnown() bool {
	switch r {
	case TransactionSimulationContractManagementProxyUpgradeManagementTypeProxyUpgrade:
		return true
	}
	return false
}

type TransactionSimulationContractManagementOwnershipChangeManagement struct {
	// The state after the transaction
	After TransactionSimulationContractManagementOwnershipChangeManagementAfter `json:"after,required"`
	// The state before the transaction
	Before TransactionSimulationContractManagementOwnershipChangeManagementBefore `json:"before,required"`
	// The type of the state change
	Type TransactionSimulationContractManagementOwnershipChangeManagementType `json:"type,required"`
	JSON transactionSimulationContractManagementOwnershipChangeManagementJSON `json:"-"`
}

// transactionSimulationContractManagementOwnershipChangeManagementJSON contains
// the JSON metadata for the struct
// [TransactionSimulationContractManagementOwnershipChangeManagement]
type transactionSimulationContractManagementOwnershipChangeManagementJSON struct {
	After       apijson.Field
	Before      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationContractManagementOwnershipChangeManagement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationContractManagementOwnershipChangeManagementJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationContractManagementOwnershipChangeManagement) implementsTransactionSimulationContractManagement() {
}

// The state after the transaction
type TransactionSimulationContractManagementOwnershipChangeManagementAfter struct {
	Owners []string                                                                  `json:"owners,required"`
	JSON   transactionSimulationContractManagementOwnershipChangeManagementAfterJSON `json:"-"`
}

// transactionSimulationContractManagementOwnershipChangeManagementAfterJSON
// contains the JSON metadata for the struct
// [TransactionSimulationContractManagementOwnershipChangeManagementAfter]
type transactionSimulationContractManagementOwnershipChangeManagementAfterJSON struct {
	Owners      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationContractManagementOwnershipChangeManagementAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationContractManagementOwnershipChangeManagementAfterJSON) RawJSON() string {
	return r.raw
}

// The state before the transaction
type TransactionSimulationContractManagementOwnershipChangeManagementBefore struct {
	Owners []string                                                                   `json:"owners,required"`
	JSON   transactionSimulationContractManagementOwnershipChangeManagementBeforeJSON `json:"-"`
}

// transactionSimulationContractManagementOwnershipChangeManagementBeforeJSON
// contains the JSON metadata for the struct
// [TransactionSimulationContractManagementOwnershipChangeManagementBefore]
type transactionSimulationContractManagementOwnershipChangeManagementBeforeJSON struct {
	Owners      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationContractManagementOwnershipChangeManagementBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationContractManagementOwnershipChangeManagementBeforeJSON) RawJSON() string {
	return r.raw
}

// The type of the state change
type TransactionSimulationContractManagementOwnershipChangeManagementType string

const (
	TransactionSimulationContractManagementOwnershipChangeManagementTypeOwnershipChange TransactionSimulationContractManagementOwnershipChangeManagementType = "OWNERSHIP_CHANGE"
)

func (r TransactionSimulationContractManagementOwnershipChangeManagementType) IsKnown() bool {
	switch r {
	case TransactionSimulationContractManagementOwnershipChangeManagementTypeOwnershipChange:
		return true
	}
	return false
}

type TransactionSimulationContractManagementModulesChangeManagement struct {
	// The state after the transaction
	After TransactionSimulationContractManagementModulesChangeManagementAfter `json:"after,required"`
	// The state before the transaction
	Before TransactionSimulationContractManagementModulesChangeManagementBefore `json:"before,required"`
	// The type of the state change
	Type TransactionSimulationContractManagementModulesChangeManagementType `json:"type,required"`
	JSON transactionSimulationContractManagementModulesChangeManagementJSON `json:"-"`
}

// transactionSimulationContractManagementModulesChangeManagementJSON contains the
// JSON metadata for the struct
// [TransactionSimulationContractManagementModulesChangeManagement]
type transactionSimulationContractManagementModulesChangeManagementJSON struct {
	After       apijson.Field
	Before      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationContractManagementModulesChangeManagement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationContractManagementModulesChangeManagementJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationContractManagementModulesChangeManagement) implementsTransactionSimulationContractManagement() {
}

// The state after the transaction
type TransactionSimulationContractManagementModulesChangeManagementAfter struct {
	Modules []string                                                                `json:"modules,required"`
	JSON    transactionSimulationContractManagementModulesChangeManagementAfterJSON `json:"-"`
}

// transactionSimulationContractManagementModulesChangeManagementAfterJSON contains
// the JSON metadata for the struct
// [TransactionSimulationContractManagementModulesChangeManagementAfter]
type transactionSimulationContractManagementModulesChangeManagementAfterJSON struct {
	Modules     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationContractManagementModulesChangeManagementAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationContractManagementModulesChangeManagementAfterJSON) RawJSON() string {
	return r.raw
}

// The state before the transaction
type TransactionSimulationContractManagementModulesChangeManagementBefore struct {
	Modules []string                                                                 `json:"modules,required"`
	JSON    transactionSimulationContractManagementModulesChangeManagementBeforeJSON `json:"-"`
}

// transactionSimulationContractManagementModulesChangeManagementBeforeJSON
// contains the JSON metadata for the struct
// [TransactionSimulationContractManagementModulesChangeManagementBefore]
type transactionSimulationContractManagementModulesChangeManagementBeforeJSON struct {
	Modules     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationContractManagementModulesChangeManagementBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationContractManagementModulesChangeManagementBeforeJSON) RawJSON() string {
	return r.raw
}

// The type of the state change
type TransactionSimulationContractManagementModulesChangeManagementType string

const (
	TransactionSimulationContractManagementModulesChangeManagementTypeModuleChange TransactionSimulationContractManagementModulesChangeManagementType = "MODULE_CHANGE"
)

func (r TransactionSimulationContractManagementModulesChangeManagementType) IsKnown() bool {
	switch r {
	case TransactionSimulationContractManagementModulesChangeManagementTypeModuleChange:
		return true
	}
	return false
}

// The type of the state change
type TransactionSimulationContractManagementType string

const (
	TransactionSimulationContractManagementTypeProxyUpgrade    TransactionSimulationContractManagementType = "PROXY_UPGRADE"
	TransactionSimulationContractManagementTypeOwnershipChange TransactionSimulationContractManagementType = "OWNERSHIP_CHANGE"
	TransactionSimulationContractManagementTypeModuleChange    TransactionSimulationContractManagementType = "MODULE_CHANGE"
)

func (r TransactionSimulationContractManagementType) IsKnown() bool {
	switch r {
	case TransactionSimulationContractManagementTypeProxyUpgrade, TransactionSimulationContractManagementTypeOwnershipChange, TransactionSimulationContractManagementTypeModuleChange:
		return true
	}
	return false
}

// The parameters of the transaction that was simulated.
type TransactionSimulationParams struct {
	// The block tag to be sent.
	BlockTag string `json:"block_tag"`
	// The calldata to be sent.
	Calldata TransactionSimulationParamsCalldata `json:"calldata"`
	// The chain to be sent.
	Chain string `json:"chain"`
	// The data to be sent.
	Data string `json:"data"`
	// The address the transaction is sent from.
	From string `json:"from"`
	// The gas to be sent.
	Gas string `json:"gas"`
	// The gas price to be sent.
	GasPrice string `json:"gas_price"`
	// The address the transaction is directed to.
	To string `json:"to"`
	// The user operation call data to be sent.
	UserOperationCalldata TransactionSimulationParamsUserOperationCalldata `json:"user_operation_calldata"`
	// The value to be sent.
	Value string                          `json:"value"`
	JSON  transactionSimulationParamsJSON `json:"-"`
}

// transactionSimulationParamsJSON contains the JSON metadata for the struct
// [TransactionSimulationParams]
type transactionSimulationParamsJSON struct {
	BlockTag              apijson.Field
	Calldata              apijson.Field
	Chain                 apijson.Field
	Data                  apijson.Field
	From                  apijson.Field
	Gas                   apijson.Field
	GasPrice              apijson.Field
	To                    apijson.Field
	UserOperationCalldata apijson.Field
	Value                 apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *TransactionSimulationParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationParamsJSON) RawJSON() string {
	return r.raw
}

// The calldata to be sent.
type TransactionSimulationParamsCalldata struct {
	// The function selector of the function called in the transaction
	FunctionSelector string `json:"function_selector,required"`
	// The function declaration of the function called in the transaction
	FunctionDeclaration string `json:"function_declaration"`
	// The function signature of the function called in the transaction
	FunctionSignature string                                  `json:"function_signature"`
	JSON              transactionSimulationParamsCalldataJSON `json:"-"`
}

// transactionSimulationParamsCalldataJSON contains the JSON metadata for the
// struct [TransactionSimulationParamsCalldata]
type transactionSimulationParamsCalldataJSON struct {
	FunctionSelector    apijson.Field
	FunctionDeclaration apijson.Field
	FunctionSignature   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *TransactionSimulationParamsCalldata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationParamsCalldataJSON) RawJSON() string {
	return r.raw
}

// The user operation call data to be sent.
type TransactionSimulationParamsUserOperationCalldata struct {
	// The function selector of the function called in the transaction
	FunctionSelector string `json:"function_selector,required"`
	// The function declaration of the function called in the transaction
	FunctionDeclaration string `json:"function_declaration"`
	// The function signature of the function called in the transaction
	FunctionSignature string                                               `json:"function_signature"`
	JSON              transactionSimulationParamsUserOperationCalldataJSON `json:"-"`
}

// transactionSimulationParamsUserOperationCalldataJSON contains the JSON metadata
// for the struct [TransactionSimulationParamsUserOperationCalldata]
type transactionSimulationParamsUserOperationCalldataJSON struct {
	FunctionSelector    apijson.Field
	FunctionDeclaration apijson.Field
	FunctionSignature   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *TransactionSimulationParamsUserOperationCalldata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationParamsUserOperationCalldataJSON) RawJSON() string {
	return r.raw
}

type TransactionSimulationError struct {
	// An error message if the simulation failed.
	Error string `json:"error,required"`
	// A string indicating if the simulation was successful or not.
	Status TransactionSimulationErrorStatus `json:"status,required"`
	// Error details if the simulation failed.
	ErrorDetails TransactionSimulationErrorErrorDetails `json:"error_details"`
	JSON         transactionSimulationErrorJSON         `json:"-"`
}

// transactionSimulationErrorJSON contains the JSON metadata for the struct
// [TransactionSimulationError]
type transactionSimulationErrorJSON struct {
	Error        apijson.Field
	Status       apijson.Field
	ErrorDetails apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TransactionSimulationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationErrorJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationError) implementsTransactionScanResponseSimulation() {}

// A string indicating if the simulation was successful or not.
type TransactionSimulationErrorStatus string

const (
	TransactionSimulationErrorStatusError TransactionSimulationErrorStatus = "Error"
)

func (r TransactionSimulationErrorStatus) IsKnown() bool {
	switch r {
	case TransactionSimulationErrorStatusError:
		return true
	}
	return false
}

// Error details if the simulation failed.
type TransactionSimulationErrorErrorDetails struct {
	// The type of the model
	Code string `json:"code,required"`
	// The address of the account
	AccountAddress string `json:"account_address"`
	// This field can have the runtime type of
	// [TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAsset].
	Asset interface{} `json:"asset"`
	// The current balance of the account
	CurrentBalance int64 `json:"current_balance"`
	// The required balance of the account
	RequiredBalance int64                                      `json:"required_balance"`
	JSON            transactionSimulationErrorErrorDetailsJSON `json:"-"`
	union           TransactionSimulationErrorErrorDetailsUnion
}

// transactionSimulationErrorErrorDetailsJSON contains the JSON metadata for the
// struct [TransactionSimulationErrorErrorDetails]
type transactionSimulationErrorErrorDetailsJSON struct {
	Code            apijson.Field
	AccountAddress  apijson.Field
	Asset           apijson.Field
	CurrentBalance  apijson.Field
	RequiredBalance apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r transactionSimulationErrorErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationErrorErrorDetails) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationErrorErrorDetails{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TransactionSimulationErrorErrorDetailsUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetails],
// [TransactionSimulationErrorErrorDetailsGenericErrorDetails].
func (r TransactionSimulationErrorErrorDetails) AsUnion() TransactionSimulationErrorErrorDetailsUnion {
	return r.union
}

// Error details if the simulation failed.
//
// Union satisfied by
// [TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetails] or
// [TransactionSimulationErrorErrorDetailsGenericErrorDetails].
type TransactionSimulationErrorErrorDetailsUnion interface {
	implementsTransactionSimulationErrorErrorDetails()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationErrorErrorDetailsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationErrorErrorDetailsGenericErrorDetails{}),
		},
	)
}

type TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetails struct {
	// The address of the account
	AccountAddress string `json:"account_address,required"`
	// The asset that the account does not have enough balance for
	Asset TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAsset `json:"asset,required"`
	// The type of the model
	Code TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsCode `json:"code,required"`
	// The current balance of the account
	CurrentBalance int64 `json:"current_balance"`
	// The required balance of the account
	RequiredBalance int64                                                                          `json:"required_balance"`
	JSON            transactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsJSON `json:"-"`
}

// transactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsJSON
// contains the JSON metadata for the struct
// [TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetails]
type transactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsJSON struct {
	AccountAddress  apijson.Field
	Asset           apijson.Field
	Code            apijson.Field
	CurrentBalance  apijson.Field
	RequiredBalance apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetails) implementsTransactionSimulationErrorErrorDetails() {
}

// The asset that the account does not have enough balance for
type TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAsset struct {
	// This field can have the runtime type of [NativeAssetDetails],
	// [Erc20TokenDetails], [Erc1155TokenDetails].
	Details interface{} `json:"details,required"`
	// The type of the model
	Type TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetType `json:"type,required"`
	// Token Id
	TokenID int64                                                                               `json:"token_id"`
	JSON    transactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetJSON `json:"-"`
	union   TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetUnion
}

// transactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetJSON
// contains the JSON metadata for the struct
// [TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAsset]
type transactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetJSON struct {
	Details     apijson.Field
	Type        apijson.Field
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAsset) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetNativeAsset],
// [TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc20Asset],
// [TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc721Asset],
// [TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc1155Asset].
func (r TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAsset) AsUnion() TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetUnion {
	return r.union
}

// The asset that the account does not have enough balance for
//
// Union satisfied by
// [TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetNativeAsset],
// [TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc20Asset],
// [TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc721Asset]
// or
// [TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc1155Asset].
type TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetUnion interface {
	implementsTransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetNativeAsset{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc20Asset{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc721Asset{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc1155Asset{}),
		},
	)
}

type TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetNativeAsset struct {
	// Details
	Details NativeAssetDetails `json:"details,required"`
	// The type of the model
	Type TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetNativeAssetType `json:"type,required"`
	JSON transactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetNativeAssetJSON `json:"-"`
}

// transactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetNativeAssetJSON
// contains the JSON metadata for the struct
// [TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetNativeAsset]
type transactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetNativeAssetJSON struct {
	Details     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetNativeAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetNativeAssetJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetNativeAsset) implementsTransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAsset() {
}

// The type of the model
type TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetNativeAssetType string

const (
	TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetNativeAssetTypeNative TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetNativeAssetType = "NATIVE"
)

func (r TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetNativeAssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetNativeAssetTypeNative:
		return true
	}
	return false
}

type TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc20Asset struct {
	// Details
	Details Erc20TokenDetails `json:"details,required"`
	// The type of the model
	Type TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc20AssetType `json:"type,required"`
	JSON transactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc20AssetJSON `json:"-"`
}

// transactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc20AssetJSON
// contains the JSON metadata for the struct
// [TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc20Asset]
type transactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc20AssetJSON struct {
	Details     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc20Asset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc20AssetJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc20Asset) implementsTransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAsset() {
}

// The type of the model
type TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc20AssetType string

const (
	TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc20AssetTypeErc20 TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc20AssetType = "ERC20"
)

func (r TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc20AssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc20AssetTypeErc20:
		return true
	}
	return false
}

type TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc721Asset struct {
	// Details
	Details Erc1155TokenDetails `json:"details,required"`
	// Token Id
	TokenID int64 `json:"token_id,required"`
	// The type of the model
	Type TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc721AssetType `json:"type,required"`
	JSON transactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc721AssetJSON `json:"-"`
}

// transactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc721AssetJSON
// contains the JSON metadata for the struct
// [TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc721Asset]
type transactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc721AssetJSON struct {
	Details     apijson.Field
	TokenID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc721Asset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc721AssetJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc721Asset) implementsTransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAsset() {
}

// The type of the model
type TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc721AssetType string

const (
	TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc721AssetTypeErc721 TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc721AssetType = "ERC721"
)

func (r TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc721AssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc721AssetTypeErc721:
		return true
	}
	return false
}

type TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc1155Asset struct {
	// Details
	Details Erc1155TokenDetails `json:"details,required"`
	// Token Id
	TokenID int64 `json:"token_id,required"`
	// The type of the model
	Type TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc1155AssetType `json:"type,required"`
	JSON transactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc1155AssetJSON `json:"-"`
}

// transactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc1155AssetJSON
// contains the JSON metadata for the struct
// [TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc1155Asset]
type transactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc1155AssetJSON struct {
	Details     apijson.Field
	TokenID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc1155Asset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc1155AssetJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc1155Asset) implementsTransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAsset() {
}

// The type of the model
type TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc1155AssetType string

const (
	TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc1155AssetTypeErc1155 TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc1155AssetType = "ERC1155"
)

func (r TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc1155AssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetErc1155AssetTypeErc1155:
		return true
	}
	return false
}

// The type of the model
type TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetType string

const (
	TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetTypeNative  TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetType = "NATIVE"
	TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetTypeErc20   TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetType = "ERC20"
	TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetTypeErc721  TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetType = "ERC721"
	TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetTypeErc1155 TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetType = "ERC1155"
)

func (r TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetTypeNative, TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetTypeErc20, TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetTypeErc721, TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsAssetTypeErc1155:
		return true
	}
	return false
}

// The type of the model
type TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsCode string

const (
	TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsCodeGeneralInsufficientFunds TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsCode = "GENERAL_INSUFFICIENT_FUNDS"
)

func (r TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsCode) IsKnown() bool {
	switch r {
	case TransactionSimulationErrorErrorDetailsGeneralInsufficientFundsErrorDetailsCodeGeneralInsufficientFunds:
		return true
	}
	return false
}

type TransactionSimulationErrorErrorDetailsGenericErrorDetails struct {
	// The error code
	Code string                                                        `json:"code,required"`
	JSON transactionSimulationErrorErrorDetailsGenericErrorDetailsJSON `json:"-"`
}

// transactionSimulationErrorErrorDetailsGenericErrorDetailsJSON contains the JSON
// metadata for the struct
// [TransactionSimulationErrorErrorDetailsGenericErrorDetails]
type transactionSimulationErrorErrorDetailsGenericErrorDetailsJSON struct {
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationErrorErrorDetailsGenericErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationErrorErrorDetailsGenericErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationErrorErrorDetailsGenericErrorDetails) implementsTransactionSimulationErrorErrorDetails() {
}

type TransactionValidation struct {
	// A list of features about this transaction explaining the validation.
	Features []TransactionScanFeature `json:"features,required"`
	// An enumeration.
	ResultType TransactionValidationResultType `json:"result_type,required"`
	// A string indicating if the simulation was successful or not.
	Status TransactionValidationStatus `json:"status,required"`
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification"`
	// A textual description that can be presented to the user about what this
	// transaction is doing.
	Description string `json:"description"`
	// A textual description about the reasons the transaction was flagged with
	// result_type.
	Reason string                    `json:"reason"`
	JSON   transactionValidationJSON `json:"-"`
}

// transactionValidationJSON contains the JSON metadata for the struct
// [TransactionValidation]
type transactionValidationJSON struct {
	Features       apijson.Field
	ResultType     apijson.Field
	Status         apijson.Field
	Classification apijson.Field
	Description    apijson.Field
	Reason         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionValidationJSON) RawJSON() string {
	return r.raw
}

func (r TransactionValidation) implementsTransactionScanResponseValidation() {}

// An enumeration.
type TransactionValidationResultType string

const (
	TransactionValidationResultTypeBenign    TransactionValidationResultType = "Benign"
	TransactionValidationResultTypeWarning   TransactionValidationResultType = "Warning"
	TransactionValidationResultTypeMalicious TransactionValidationResultType = "Malicious"
)

func (r TransactionValidationResultType) IsKnown() bool {
	switch r {
	case TransactionValidationResultTypeBenign, TransactionValidationResultTypeWarning, TransactionValidationResultTypeMalicious:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type TransactionValidationStatus string

const (
	TransactionValidationStatusSuccess TransactionValidationStatus = "Success"
)

func (r TransactionValidationStatus) IsKnown() bool {
	switch r {
	case TransactionValidationStatusSuccess:
		return true
	}
	return false
}

type TransactionValidationError struct {
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification TransactionValidationErrorClassification `json:"classification,required"`
	// A textual description that can be presented to the user about what this
	// transaction is doing.
	Description TransactionValidationErrorDescription `json:"description,required"`
	// An error message if the validation failed.
	Error string `json:"error,required"`
	// A list of features about this transaction explaining the validation.
	Features []TransactionScanFeature `json:"features,required"`
	// A textual description about the reasons the transaction was flagged with
	// result_type.
	Reason TransactionValidationErrorReason `json:"reason,required"`
	// A string indicating if the transaction is safe to sign or not.
	ResultType TransactionValidationErrorResultType `json:"result_type,required"`
	// A string indicating if the simulation was successful or not.
	Status TransactionValidationErrorStatus `json:"status,required"`
	JSON   transactionValidationErrorJSON   `json:"-"`
}

// transactionValidationErrorJSON contains the JSON metadata for the struct
// [TransactionValidationError]
type transactionValidationErrorJSON struct {
	Classification apijson.Field
	Description    apijson.Field
	Error          apijson.Field
	Features       apijson.Field
	Reason         apijson.Field
	ResultType     apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionValidationErrorJSON) RawJSON() string {
	return r.raw
}

func (r TransactionValidationError) implementsTransactionScanResponseValidation() {}

// A textual classification that can be presented to the user explaining the
// reason.
type TransactionValidationErrorClassification string

const (
	TransactionValidationErrorClassificationEmpty TransactionValidationErrorClassification = ""
)

func (r TransactionValidationErrorClassification) IsKnown() bool {
	switch r {
	case TransactionValidationErrorClassificationEmpty:
		return true
	}
	return false
}

// A textual description that can be presented to the user about what this
// transaction is doing.
type TransactionValidationErrorDescription string

const (
	TransactionValidationErrorDescriptionEmpty TransactionValidationErrorDescription = ""
)

func (r TransactionValidationErrorDescription) IsKnown() bool {
	switch r {
	case TransactionValidationErrorDescriptionEmpty:
		return true
	}
	return false
}

// A textual description about the reasons the transaction was flagged with
// result_type.
type TransactionValidationErrorReason string

const (
	TransactionValidationErrorReasonEmpty TransactionValidationErrorReason = ""
)

func (r TransactionValidationErrorReason) IsKnown() bool {
	switch r {
	case TransactionValidationErrorReasonEmpty:
		return true
	}
	return false
}

// A string indicating if the transaction is safe to sign or not.
type TransactionValidationErrorResultType string

const (
	TransactionValidationErrorResultTypeError TransactionValidationErrorResultType = "Error"
)

func (r TransactionValidationErrorResultType) IsKnown() bool {
	switch r {
	case TransactionValidationErrorResultTypeError:
		return true
	}
	return false
}

// A string indicating if the simulation was successful or not.
type TransactionValidationErrorStatus string

const (
	TransactionValidationErrorStatusSuccess TransactionValidationErrorStatus = "Success"
)

func (r TransactionValidationErrorStatus) IsKnown() bool {
	switch r {
	case TransactionValidationErrorStatusSuccess:
		return true
	}
	return false
}

type UsdDiff struct {
	In    string      `json:"in,required"`
	Out   string      `json:"out,required"`
	Total string      `json:"total,required"`
	JSON  usdDiffJSON `json:"-"`
}

// usdDiffJSON contains the JSON metadata for the struct [UsdDiff]
type usdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usdDiffJSON) RawJSON() string {
	return r.raw
}
