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

type AddressAssetExposure struct {
	// description of the asset for the current diff
	Asset AddressAssetExposureAsset `json:"asset,required"`
	// dictionary of spender addresses where the exposure has changed during this
	// transaction for the current address and asset
	Spenders map[string]AddressAssetExposureSpender `json:"spenders,required"`
	JSON     addressAssetExposureJSON               `json:"-"`
}

// addressAssetExposureJSON contains the JSON metadata for the struct
// [AddressAssetExposure]
type addressAssetExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAssetExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressAssetExposureJSON) RawJSON() string {
	return r.raw
}

// description of the asset for the current diff
type AddressAssetExposureAsset struct {
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string `json:"symbol"`
	// address of the token
	Address string `json:"address,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// asset type.
	Type AddressAssetExposureAssetType `json:"type,required"`
	// asset's decimals
	Decimals int64                         `json:"decimals"`
	JSON     addressAssetExposureAssetJSON `json:"-"`
	union    AddressAssetExposureAssetUnion
}

// addressAssetExposureAssetJSON contains the JSON metadata for the struct
// [AddressAssetExposureAsset]
type addressAssetExposureAssetJSON struct {
	Name        apijson.Field
	Symbol      apijson.Field
	Address     apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r addressAssetExposureAssetJSON) RawJSON() string {
	return r.raw
}

func (r *AddressAssetExposureAsset) UnmarshalJSON(data []byte) (err error) {
	*r = AddressAssetExposureAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AddressAssetExposureAssetUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [Erc20TokenDetails],
// [Erc1155TokenDetails], [Erc721TokenDetails], [NonercTokenDetails].
func (r AddressAssetExposureAsset) AsUnion() AddressAssetExposureAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by [Erc20TokenDetails], [Erc1155TokenDetails],
// [Erc721TokenDetails] or [NonercTokenDetails].
type AddressAssetExposureAssetUnion interface {
	implementsAddressAssetExposureAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressAssetExposureAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc1155TokenDetails{}),
		},
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
type AddressAssetExposureAssetType string

const (
	AddressAssetExposureAssetTypeErc20   AddressAssetExposureAssetType = "ERC20"
	AddressAssetExposureAssetTypeErc1155 AddressAssetExposureAssetType = "ERC1155"
	AddressAssetExposureAssetTypeErc721  AddressAssetExposureAssetType = "ERC721"
	AddressAssetExposureAssetTypeNonerc  AddressAssetExposureAssetType = "NONERC"
)

func (r AddressAssetExposureAssetType) IsKnown() bool {
	switch r {
	case AddressAssetExposureAssetTypeErc20, AddressAssetExposureAssetTypeErc1155, AddressAssetExposureAssetTypeErc721, AddressAssetExposureAssetTypeNonerc:
		return true
	}
	return false
}

type AddressAssetExposureSpender struct {
	// This field can have the runtime type of [[]Erc20ExposureExposure],
	// [[]Erc721ExposureExposure], [[]Erc1155ExposureExposure].
	Exposure interface{} `json:"exposure"`
	// user friendly description of the approval
	Summary string `json:"summary"`
	// the amount that was asked in the approval request for this spender from the
	// current address and asset
	Approval string `json:"approval"`
	// the expiration time of the permit2 protocol
	Expiration time.Time `json:"expiration" format:"date-time"`
	// boolean indicates whether an is_approved_for_all function was used (missing in
	// case of ERC20 / ERC1155)
	IsApprovedForAll bool                            `json:"is_approved_for_all"`
	JSON             addressAssetExposureSpenderJSON `json:"-"`
	union            AddressAssetExposureSpendersUnion
}

// addressAssetExposureSpenderJSON contains the JSON metadata for the struct
// [AddressAssetExposureSpender]
type addressAssetExposureSpenderJSON struct {
	Exposure         apijson.Field
	Summary          apijson.Field
	Approval         apijson.Field
	Expiration       apijson.Field
	IsApprovedForAll apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r addressAssetExposureSpenderJSON) RawJSON() string {
	return r.raw
}

func (r *AddressAssetExposureSpender) UnmarshalJSON(data []byte) (err error) {
	*r = AddressAssetExposureSpender{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AddressAssetExposureSpendersUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [Erc20Exposure], [Erc721Exposure],
// [Erc1155Exposure].
func (r AddressAssetExposureSpender) AsUnion() AddressAssetExposureSpendersUnion {
	return r.union
}

// Union satisfied by [Erc20Exposure], [Erc721Exposure] or [Erc1155Exposure].
type AddressAssetExposureSpendersUnion interface {
	implementsAddressAssetExposureSpender()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressAssetExposureSpendersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc20Exposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc721Exposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc1155Exposure{}),
		},
	)
}

type AssetDiff struct {
	// description of the asset for the current diff
	Asset AssetDiffAsset `json:"asset,required"`
	// An enumeration.
	AssetType AssetDiffAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []AssetDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []AssetDiffOut `json:"out,required"`
	JSON assetDiffJSON  `json:"-"`
}

// assetDiffJSON contains the JSON metadata for the struct [AssetDiff]
type assetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetDiffJSON) RawJSON() string {
	return r.raw
}

// description of the asset for the current diff
type AssetDiffAsset struct {
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string `json:"symbol"`
	// address of the token
	Address string `json:"address"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// asset type.
	Type AssetDiffAssetType `json:"type,required"`
	// asset's decimals
	Decimals  int64              `json:"decimals"`
	ChainName string             `json:"chain_name"`
	ChainID   int64              `json:"chain_id"`
	JSON      assetDiffAssetJSON `json:"-"`
	union     AssetDiffAssetUnion
}

// assetDiffAssetJSON contains the JSON metadata for the struct [AssetDiffAsset]
type assetDiffAssetJSON struct {
	Name        apijson.Field
	Symbol      apijson.Field
	Address     apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	ChainName   apijson.Field
	ChainID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r assetDiffAssetJSON) RawJSON() string {
	return r.raw
}

func (r *AssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	*r = AssetDiffAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AssetDiffAssetUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [Erc20TokenDetails],
// [Erc1155TokenDetails], [Erc721TokenDetails], [NonercTokenDetails],
// [NativeAssetDetails].
func (r AssetDiffAsset) AsUnion() AssetDiffAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by [Erc20TokenDetails], [Erc1155TokenDetails],
// [Erc721TokenDetails], [NonercTokenDetails] or [NativeAssetDetails].
type AssetDiffAssetUnion interface {
	implementsAssetDiffAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AssetDiffAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NonercTokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NativeAssetDetails{}),
		},
	)
}

// asset type.
type AssetDiffAssetType string

const (
	AssetDiffAssetTypeErc20   AssetDiffAssetType = "ERC20"
	AssetDiffAssetTypeErc1155 AssetDiffAssetType = "ERC1155"
	AssetDiffAssetTypeErc721  AssetDiffAssetType = "ERC721"
	AssetDiffAssetTypeNonerc  AssetDiffAssetType = "NONERC"
	AssetDiffAssetTypeNative  AssetDiffAssetType = "NATIVE"
)

func (r AssetDiffAssetType) IsKnown() bool {
	switch r {
	case AssetDiffAssetTypeErc20, AssetDiffAssetTypeErc1155, AssetDiffAssetTypeErc721, AssetDiffAssetTypeNonerc, AssetDiffAssetTypeNative:
		return true
	}
	return false
}

type AssetDiffIn struct {
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// id of the token
	TokenID string `json:"token_id"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value"`
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// value before divided by decimal, that was transferred from this address
	RawValue string          `json:"raw_value"`
	JSON     assetDiffInJSON `json:"-"`
	union    AssetDiffInUnion
}

// assetDiffInJSON contains the JSON metadata for the struct [AssetDiffIn]
type assetDiffInJSON struct {
	UsdPrice                 apijson.Field
	Summary                  apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	ArbitraryCollectionToken apijson.Field
	LogoURL                  apijson.Field
	RawValue                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r assetDiffInJSON) RawJSON() string {
	return r.raw
}

func (r *AssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	*r = AssetDiffIn{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AssetDiffInUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [Erc1155Diff], [Erc721Diff],
// [Erc20Diff], [NativeDiff].
func (r AssetDiffIn) AsUnion() AssetDiffInUnion {
	return r.union
}

// Union satisfied by [Erc1155Diff], [Erc721Diff], [Erc20Diff] or [NativeDiff].
type AssetDiffInUnion interface {
	implementsAssetDiffIn()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AssetDiffInUnion)(nil)).Elem(),
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

type AssetDiffOut struct {
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// id of the token
	TokenID string `json:"token_id"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value"`
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// value before divided by decimal, that was transferred from this address
	RawValue string           `json:"raw_value"`
	JSON     assetDiffOutJSON `json:"-"`
	union    AssetDiffOutUnion
}

// assetDiffOutJSON contains the JSON metadata for the struct [AssetDiffOut]
type assetDiffOutJSON struct {
	UsdPrice                 apijson.Field
	Summary                  apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	ArbitraryCollectionToken apijson.Field
	LogoURL                  apijson.Field
	RawValue                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r assetDiffOutJSON) RawJSON() string {
	return r.raw
}

func (r *AssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	*r = AssetDiffOut{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AssetDiffOutUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [Erc1155Diff], [Erc721Diff],
// [Erc20Diff], [NativeDiff].
func (r AssetDiffOut) AsUnion() AssetDiffOutUnion {
	return r.union
}

// Union satisfied by [Erc1155Diff], [Erc721Diff], [Erc20Diff] or [NativeDiff].
type AssetDiffOutUnion interface {
	implementsAssetDiffOut()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AssetDiffOutUnion)(nil)).Elem(),
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

func (r Erc1155Diff) implementsErc20ExposureExposure() {}

func (r Erc1155Diff) implementsErc721ExposureExposure() {}

func (r Erc1155Diff) implementsErc1155ExposureExposure() {}

func (r Erc1155Diff) implementsAssetDiffIn() {}

func (r Erc1155Diff) implementsAssetDiffOut() {}

func (r Erc1155Diff) implementsTransactionSimulationAccountSummaryAssetsDiffsIn() {}

func (r Erc1155Diff) implementsTransactionSimulationAccountSummaryAssetsDiffsOut() {}

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

func (r Erc1155Exposure) implementsAddressAssetExposureSpender() {}

type Erc1155ExposureExposure struct {
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// id of the token
	TokenID string `json:"token_id"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value"`
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// value before divided by decimal, that was transferred from this address
	RawValue string                      `json:"raw_value"`
	JSON     erc1155ExposureExposureJSON `json:"-"`
	union    Erc1155ExposureExposureUnion
}

// erc1155ExposureExposureJSON contains the JSON metadata for the struct
// [Erc1155ExposureExposure]
type erc1155ExposureExposureJSON struct {
	UsdPrice                 apijson.Field
	Summary                  apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	ArbitraryCollectionToken apijson.Field
	LogoURL                  apijson.Field
	RawValue                 apijson.Field
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

func (r Erc1155TokenDetails) implementsAddressAssetExposureAsset() {}

func (r Erc1155TokenDetails) implementsAssetDiffAsset() {}

func (r Erc1155TokenDetails) implementsTransactionSimulationAccountSummaryAssetsDiffsAsset() {}

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

func (r Erc20Diff) implementsErc20ExposureExposure() {}

func (r Erc20Diff) implementsErc721ExposureExposure() {}

func (r Erc20Diff) implementsErc1155ExposureExposure() {}

func (r Erc20Diff) implementsAssetDiffIn() {}

func (r Erc20Diff) implementsAssetDiffOut() {}

func (r Erc20Diff) implementsTransactionSimulationAccountSummaryAssetsDiffsIn() {}

func (r Erc20Diff) implementsTransactionSimulationAccountSummaryAssetsDiffsOut() {}

type Erc20Exposure struct {
	// the amount that was asked in the approval request for this spender from the
	// current address and asset
	Approval string                  `json:"approval,required"`
	Exposure []Erc20ExposureExposure `json:"exposure,required"`
	// the expiration time of the permit2 protocol
	Expiration time.Time `json:"expiration" format:"date-time"`
	// user friendly description of the approval
	Summary string            `json:"summary"`
	JSON    erc20ExposureJSON `json:"-"`
}

// erc20ExposureJSON contains the JSON metadata for the struct [Erc20Exposure]
type erc20ExposureJSON struct {
	Approval    apijson.Field
	Exposure    apijson.Field
	Expiration  apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Erc20Exposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r erc20ExposureJSON) RawJSON() string {
	return r.raw
}

func (r Erc20Exposure) implementsAddressAssetExposureSpender() {}

type Erc20ExposureExposure struct {
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// id of the token
	TokenID string `json:"token_id"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value"`
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// value before divided by decimal, that was transferred from this address
	RawValue string                    `json:"raw_value"`
	JSON     erc20ExposureExposureJSON `json:"-"`
	union    Erc20ExposureExposureUnion
}

// erc20ExposureExposureJSON contains the JSON metadata for the struct
// [Erc20ExposureExposure]
type erc20ExposureExposureJSON struct {
	UsdPrice                 apijson.Field
	Summary                  apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	ArbitraryCollectionToken apijson.Field
	LogoURL                  apijson.Field
	RawValue                 apijson.Field
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

func (r Erc20TokenDetails) implementsAddressAssetExposureAsset() {}

func (r Erc20TokenDetails) implementsAssetDiffAsset() {}

func (r Erc20TokenDetails) implementsTransactionSimulationAccountSummaryAssetsDiffsAsset() {}

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

func (r Erc721Diff) implementsErc20ExposureExposure() {}

func (r Erc721Diff) implementsErc721ExposureExposure() {}

func (r Erc721Diff) implementsErc1155ExposureExposure() {}

func (r Erc721Diff) implementsAssetDiffIn() {}

func (r Erc721Diff) implementsAssetDiffOut() {}

func (r Erc721Diff) implementsTransactionSimulationAccountSummaryAssetsDiffsIn() {}

func (r Erc721Diff) implementsTransactionSimulationAccountSummaryAssetsDiffsOut() {}

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

func (r Erc721Exposure) implementsAddressAssetExposureSpender() {}

type Erc721ExposureExposure struct {
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// id of the token
	TokenID string `json:"token_id"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value"`
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// value before divided by decimal, that was transferred from this address
	RawValue string                     `json:"raw_value"`
	JSON     erc721ExposureExposureJSON `json:"-"`
	union    Erc721ExposureExposureUnion
}

// erc721ExposureExposureJSON contains the JSON metadata for the struct
// [Erc721ExposureExposure]
type erc721ExposureExposureJSON struct {
	UsdPrice                 apijson.Field
	Summary                  apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	ArbitraryCollectionToken apijson.Field
	LogoURL                  apijson.Field
	RawValue                 apijson.Field
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

func (r Erc721TokenDetails) implementsAddressAssetExposureAsset() {}

func (r Erc721TokenDetails) implementsAssetDiffAsset() {}

func (r Erc721TokenDetails) implementsTransactionSimulationAccountSummaryAssetsDiffsAsset() {}

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

func (r NativeAssetDetails) implementsAssetDiffAsset() {}

func (r NativeAssetDetails) implementsTransactionSimulationAccountSummaryAssetsDiffsAsset() {}

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

func (r NativeDiff) implementsErc20ExposureExposure() {}

func (r NativeDiff) implementsErc721ExposureExposure() {}

func (r NativeDiff) implementsErc1155ExposureExposure() {}

func (r NativeDiff) implementsAssetDiffIn() {}

func (r NativeDiff) implementsAssetDiffOut() {}

func (r NativeDiff) implementsTransactionSimulationAccountSummaryAssetsDiffsIn() {}

func (r NativeDiff) implementsTransactionSimulationAccountSummaryAssetsDiffsOut() {}

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

func (r NonercTokenDetails) implementsAddressAssetExposureAsset() {}

func (r NonercTokenDetails) implementsAssetDiffAsset() {}

func (r NonercTokenDetails) implementsTransactionSimulationAccountSummaryAssetsDiffsAsset() {}

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
	TokenScanSupportedChainArbitrum  TokenScanSupportedChain = "arbitrum"
	TokenScanSupportedChainAvalanche TokenScanSupportedChain = "avalanche"
	TokenScanSupportedChainBase      TokenScanSupportedChain = "base"
	TokenScanSupportedChainBsc       TokenScanSupportedChain = "bsc"
	TokenScanSupportedChainEthereum  TokenScanSupportedChain = "ethereum"
	TokenScanSupportedChainOptimism  TokenScanSupportedChain = "optimism"
	TokenScanSupportedChainPolygon   TokenScanSupportedChain = "polygon"
	TokenScanSupportedChainZora      TokenScanSupportedChain = "zora"
	TokenScanSupportedChainSolana    TokenScanSupportedChain = "solana"
	TokenScanSupportedChainStarknet  TokenScanSupportedChain = "starknet"
	TokenScanSupportedChainStellar   TokenScanSupportedChain = "stellar"
	TokenScanSupportedChainLinea     TokenScanSupportedChain = "linea"
	TokenScanSupportedChainBlast     TokenScanSupportedChain = "blast"
	TokenScanSupportedChainZksync    TokenScanSupportedChain = "zksync"
	TokenScanSupportedChainScroll    TokenScanSupportedChain = "scroll"
	TokenScanSupportedChainDegen     TokenScanSupportedChain = "degen"
	TokenScanSupportedChainBitcoin   TokenScanSupportedChain = "bitcoin"
)

func (r TokenScanSupportedChain) IsKnown() bool {
	switch r {
	case TokenScanSupportedChainArbitrum, TokenScanSupportedChainAvalanche, TokenScanSupportedChainBase, TokenScanSupportedChainBsc, TokenScanSupportedChainEthereum, TokenScanSupportedChainOptimism, TokenScanSupportedChainPolygon, TokenScanSupportedChainZora, TokenScanSupportedChainSolana, TokenScanSupportedChainStarknet, TokenScanSupportedChainStellar, TokenScanSupportedChainLinea, TokenScanSupportedChainBlast, TokenScanSupportedChainZksync, TokenScanSupportedChainScroll, TokenScanSupportedChainDegen, TokenScanSupportedChainBitcoin:
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
	Used     int64                                      `json:"used"`
	Estimate int64                                      `json:"estimate"`
	Error    string                                     `json:"error"`
	JSON     transactionScanResponseGasEstimationJSON   `json:"-"`
	union    TransactionScanResponseGasEstimationUnion
}

// transactionScanResponseGasEstimationJSON contains the JSON metadata for the
// struct [TransactionScanResponseGasEstimation]
type transactionScanResponseGasEstimationJSON struct {
	Status      apijson.Field
	Used        apijson.Field
	Estimate    apijson.Field
	Error       apijson.Field
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
	// This field can have the runtime type of [map[string][]AssetDiff].
	AssetsDiffs interface{} `json:"assets_diffs,required"`
	// This field can have the runtime type of [map[string]UsdDiff].
	TotalUsdDiff interface{} `json:"total_usd_diff,required"`
	// This field can have the runtime type of [map[string][]AddressAssetExposure].
	Exposures interface{} `json:"exposures,required"`
	// This field can have the runtime type of [map[string]map[string]string].
	TotalUsdExposure interface{} `json:"total_usd_exposure,required"`
	// This field can have the runtime type of
	// [map[string]TransactionSimulationAddressDetail].
	AddressDetails interface{} `json:"address_details,required"`
	// This field can have the runtime type of [TransactionSimulationAccountSummary].
	AccountSummary interface{} `json:"account_summary,required"`
	// This field can have the runtime type of [TransactionSimulationParams].
	Params interface{} `json:"params,required"`
	// This field can have the runtime type of
	// [map[string][]TransactionSimulationContractManagement].
	ContractManagement interface{} `json:"contract_management,required"`
	// An error message if the simulation failed.
	Error string                                `json:"error"`
	JSON  transactionScanResponseSimulationJSON `json:"-"`
	union TransactionScanResponseSimulationUnion
}

// transactionScanResponseSimulationJSON contains the JSON metadata for the struct
// [TransactionScanResponseSimulation]
type transactionScanResponseSimulationJSON struct {
	Status             apijson.Field
	AssetsDiffs        apijson.Field
	TotalUsdDiff       apijson.Field
	Exposures          apijson.Field
	TotalUsdExposure   apijson.Field
	AddressDetails     apijson.Field
	AccountSummary     apijson.Field
	Params             apijson.Field
	ContractManagement apijson.Field
	Error              apijson.Field
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
	// A string indicating if the simulation was successful or not.
	Status TransactionScanResponseValidationStatus `json:"status,required"`
	// An enumeration.
	ResultType TransactionScanResponseValidationResultType `json:"result_type,required"`
	// A textual description that can be presented to the user about what this
	// transaction is doing.
	Description string `json:"description"`
	// A textual description about the reasons the transaction was flagged with
	// result_type.
	Reason string `json:"reason"`
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification"`
	// This field can have the runtime type of [[]TransactionScanFeature].
	Features interface{} `json:"features"`
	// An error message if the validation failed.
	Error string                                `json:"error"`
	JSON  transactionScanResponseValidationJSON `json:"-"`
	union TransactionScanResponseValidationUnion
}

// transactionScanResponseValidationJSON contains the JSON metadata for the struct
// [TransactionScanResponseValidation]
type transactionScanResponseValidationJSON struct {
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
)

func (r TransactionScanSupportedChain) IsKnown() bool {
	switch r {
	case TransactionScanSupportedChainArbitrum, TransactionScanSupportedChainAvalanche, TransactionScanSupportedChainBase, TransactionScanSupportedChainBaseSepolia, TransactionScanSupportedChainBsc, TransactionScanSupportedChainEthereum, TransactionScanSupportedChainOptimism, TransactionScanSupportedChainPolygon, TransactionScanSupportedChainZksync, TransactionScanSupportedChainZksyncSepolia, TransactionScanSupportedChainZora, TransactionScanSupportedChainLinea, TransactionScanSupportedChainBlast, TransactionScanSupportedChainScroll, TransactionScanSupportedChainEthereumSepolia, TransactionScanSupportedChainDegen, TransactionScanSupportedChainAvalancheFuji, TransactionScanSupportedChainImmutableZkevm, TransactionScanSupportedChainImmutableZkevmTestnet, TransactionScanSupportedChainGnosis, TransactionScanSupportedChainWorldchain:
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
	AssetsDiffs map[string][]AssetDiff `json:"assets_diffs,required"`
	// dictionary describes the exposure differences as a result of this transaction
	// for every involved address (as a result of any approval / setApproval / permit
	// function)
	Exposures map[string][]AddressAssetExposure `json:"exposures,required"`
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
	Exposures []AddressAssetExposure `json:"exposures,required"`
	// Total usd diff related to the account address
	TotalUsdDiff UsdDiff `json:"total_usd_diff,required"`
	// Total usd exposure related to the account address
	TotalUsdExposure map[string]string                       `json:"total_usd_exposure,required"`
	JSON             transactionSimulationAccountSummaryJSON `json:"-"`
}

// transactionSimulationAccountSummaryJSON contains the JSON metadata for the
// struct [TransactionSimulationAccountSummary]
type transactionSimulationAccountSummaryJSON struct {
	AssetsDiffs      apijson.Field
	Exposures        apijson.Field
	TotalUsdDiff     apijson.Field
	TotalUsdExposure apijson.Field
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
	// description of the asset for the current diff
	Asset TransactionSimulationAccountSummaryAssetsDiffsAsset `json:"asset,required"`
	// An enumeration.
	AssetType TransactionSimulationAccountSummaryAssetsDiffsAssetType `json:"asset_type,required"`
	// amount of the asset that was transferred to the address in this transaction
	In []TransactionSimulationAccountSummaryAssetsDiffsIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out []TransactionSimulationAccountSummaryAssetsDiffsOut `json:"out,required"`
	// shows the balance before making the transaction and after
	BalanceChanges TransactionSimulationAccountSummaryAssetsDiffsBalanceChanges `json:"balance_changes"`
	JSON           transactionSimulationAccountSummaryAssetsDiffJSON            `json:"-"`
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

func (r *TransactionSimulationAccountSummaryAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAccountSummaryAssetsDiffJSON) RawJSON() string {
	return r.raw
}

// description of the asset for the current diff
type TransactionSimulationAccountSummaryAssetsDiffsAsset struct {
	// string represents the name of the asset
	Name string `json:"name"`
	// asset's symbol name
	Symbol string `json:"symbol"`
	// address of the token
	Address string `json:"address"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// asset type.
	Type TransactionSimulationAccountSummaryAssetsDiffsAssetType `json:"type,required"`
	// asset's decimals
	Decimals  int64                                                   `json:"decimals"`
	ChainName string                                                  `json:"chain_name"`
	ChainID   int64                                                   `json:"chain_id"`
	JSON      transactionSimulationAccountSummaryAssetsDiffsAssetJSON `json:"-"`
	union     TransactionSimulationAccountSummaryAssetsDiffsAssetUnion
}

// transactionSimulationAccountSummaryAssetsDiffsAssetJSON contains the JSON
// metadata for the struct [TransactionSimulationAccountSummaryAssetsDiffsAsset]
type transactionSimulationAccountSummaryAssetsDiffsAssetJSON struct {
	Name        apijson.Field
	Symbol      apijson.Field
	Address     apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	ChainName   apijson.Field
	ChainID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionSimulationAccountSummaryAssetsDiffsAssetJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationAccountSummaryAssetsDiffsAsset) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationAccountSummaryAssetsDiffsAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TransactionSimulationAccountSummaryAssetsDiffsAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [Erc20TokenDetails],
// [Erc1155TokenDetails], [Erc721TokenDetails], [NonercTokenDetails],
// [NativeAssetDetails].
func (r TransactionSimulationAccountSummaryAssetsDiffsAsset) AsUnion() TransactionSimulationAccountSummaryAssetsDiffsAssetUnion {
	return r.union
}

// description of the asset for the current diff
//
// Union satisfied by [Erc20TokenDetails], [Erc1155TokenDetails],
// [Erc721TokenDetails], [NonercTokenDetails] or [NativeAssetDetails].
type TransactionSimulationAccountSummaryAssetsDiffsAssetUnion interface {
	implementsTransactionSimulationAccountSummaryAssetsDiffsAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationAccountSummaryAssetsDiffsAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc20TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc1155TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Erc721TokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NonercTokenDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NativeAssetDetails{}),
		},
	)
}

// asset type.
type TransactionSimulationAccountSummaryAssetsDiffsAssetType string

const (
	TransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc20   TransactionSimulationAccountSummaryAssetsDiffsAssetType = "ERC20"
	TransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc1155 TransactionSimulationAccountSummaryAssetsDiffsAssetType = "ERC1155"
	TransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc721  TransactionSimulationAccountSummaryAssetsDiffsAssetType = "ERC721"
	TransactionSimulationAccountSummaryAssetsDiffsAssetTypeNonerc  TransactionSimulationAccountSummaryAssetsDiffsAssetType = "NONERC"
	TransactionSimulationAccountSummaryAssetsDiffsAssetTypeNative  TransactionSimulationAccountSummaryAssetsDiffsAssetType = "NATIVE"
)

func (r TransactionSimulationAccountSummaryAssetsDiffsAssetType) IsKnown() bool {
	switch r {
	case TransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc20, TransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc1155, TransactionSimulationAccountSummaryAssetsDiffsAssetTypeErc721, TransactionSimulationAccountSummaryAssetsDiffsAssetTypeNonerc, TransactionSimulationAccountSummaryAssetsDiffsAssetTypeNative:
		return true
	}
	return false
}

type TransactionSimulationAccountSummaryAssetsDiffsIn struct {
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// id of the token
	TokenID string `json:"token_id"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value"`
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// value before divided by decimal, that was transferred from this address
	RawValue string                                               `json:"raw_value"`
	JSON     transactionSimulationAccountSummaryAssetsDiffsInJSON `json:"-"`
	union    TransactionSimulationAccountSummaryAssetsDiffsInUnion
}

// transactionSimulationAccountSummaryAssetsDiffsInJSON contains the JSON metadata
// for the struct [TransactionSimulationAccountSummaryAssetsDiffsIn]
type transactionSimulationAccountSummaryAssetsDiffsInJSON struct {
	UsdPrice                 apijson.Field
	Summary                  apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	ArbitraryCollectionToken apijson.Field
	LogoURL                  apijson.Field
	RawValue                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r transactionSimulationAccountSummaryAssetsDiffsInJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationAccountSummaryAssetsDiffsIn) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationAccountSummaryAssetsDiffsIn{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TransactionSimulationAccountSummaryAssetsDiffsInUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [Erc1155Diff], [Erc721Diff],
// [Erc20Diff], [NativeDiff].
func (r TransactionSimulationAccountSummaryAssetsDiffsIn) AsUnion() TransactionSimulationAccountSummaryAssetsDiffsInUnion {
	return r.union
}

// Union satisfied by [Erc1155Diff], [Erc721Diff], [Erc20Diff] or [NativeDiff].
type TransactionSimulationAccountSummaryAssetsDiffsInUnion interface {
	implementsTransactionSimulationAccountSummaryAssetsDiffsIn()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationAccountSummaryAssetsDiffsInUnion)(nil)).Elem(),
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

type TransactionSimulationAccountSummaryAssetsDiffsOut struct {
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// id of the token
	TokenID string `json:"token_id"`
	// value before divided by decimal, that was transferred from this address
	Value string `json:"value"`
	// Indicates whether the token ID represents an arbitrary token from a collection,
	// unpredictable while running the simulation
	ArbitraryCollectionToken bool `json:"arbitrary_collection_token"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// value before divided by decimal, that was transferred from this address
	RawValue string                                                `json:"raw_value"`
	JSON     transactionSimulationAccountSummaryAssetsDiffsOutJSON `json:"-"`
	union    TransactionSimulationAccountSummaryAssetsDiffsOutUnion
}

// transactionSimulationAccountSummaryAssetsDiffsOutJSON contains the JSON metadata
// for the struct [TransactionSimulationAccountSummaryAssetsDiffsOut]
type transactionSimulationAccountSummaryAssetsDiffsOutJSON struct {
	UsdPrice                 apijson.Field
	Summary                  apijson.Field
	TokenID                  apijson.Field
	Value                    apijson.Field
	ArbitraryCollectionToken apijson.Field
	LogoURL                  apijson.Field
	RawValue                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r transactionSimulationAccountSummaryAssetsDiffsOutJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationAccountSummaryAssetsDiffsOut) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationAccountSummaryAssetsDiffsOut{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TransactionSimulationAccountSummaryAssetsDiffsOutUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [Erc1155Diff], [Erc721Diff],
// [Erc20Diff], [NativeDiff].
func (r TransactionSimulationAccountSummaryAssetsDiffsOut) AsUnion() TransactionSimulationAccountSummaryAssetsDiffsOutUnion {
	return r.union
}

// Union satisfied by [Erc1155Diff], [Erc721Diff], [Erc20Diff] or [NativeDiff].
type TransactionSimulationAccountSummaryAssetsDiffsOutUnion interface {
	implementsTransactionSimulationAccountSummaryAssetsDiffsOut()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationAccountSummaryAssetsDiffsOutUnion)(nil)).Elem(),
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

// shows the balance before making the transaction and after
type TransactionSimulationAccountSummaryAssetsDiffsBalanceChanges struct {
	// balance of the account after making the transaction
	After TransactionSimulationAccountSummaryAssetsDiffsBalanceChangesAfter `json:"after,required"`
	// balance of the account before making the transaction
	Before TransactionSimulationAccountSummaryAssetsDiffsBalanceChangesBefore `json:"before,required"`
	JSON   transactionSimulationAccountSummaryAssetsDiffsBalanceChangesJSON   `json:"-"`
}

// transactionSimulationAccountSummaryAssetsDiffsBalanceChangesJSON contains the
// JSON metadata for the struct
// [TransactionSimulationAccountSummaryAssetsDiffsBalanceChanges]
type transactionSimulationAccountSummaryAssetsDiffsBalanceChangesJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationAccountSummaryAssetsDiffsBalanceChanges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAccountSummaryAssetsDiffsBalanceChangesJSON) RawJSON() string {
	return r.raw
}

// balance of the account after making the transaction
type TransactionSimulationAccountSummaryAssetsDiffsBalanceChangesAfter struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                `json:"value"`
	JSON  transactionSimulationAccountSummaryAssetsDiffsBalanceChangesAfterJSON `json:"-"`
}

// transactionSimulationAccountSummaryAssetsDiffsBalanceChangesAfterJSON contains
// the JSON metadata for the struct
// [TransactionSimulationAccountSummaryAssetsDiffsBalanceChangesAfter]
type transactionSimulationAccountSummaryAssetsDiffsBalanceChangesAfterJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationAccountSummaryAssetsDiffsBalanceChangesAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAccountSummaryAssetsDiffsBalanceChangesAfterJSON) RawJSON() string {
	return r.raw
}

// balance of the account before making the transaction
type TransactionSimulationAccountSummaryAssetsDiffsBalanceChangesBefore struct {
	// value before divided by decimal, that was transferred from this address
	RawValue string `json:"raw_value,required"`
	// usd equal of the asset that was transferred from this address
	UsdPrice string `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value string                                                                 `json:"value"`
	JSON  transactionSimulationAccountSummaryAssetsDiffsBalanceChangesBeforeJSON `json:"-"`
}

// transactionSimulationAccountSummaryAssetsDiffsBalanceChangesBeforeJSON contains
// the JSON metadata for the struct
// [TransactionSimulationAccountSummaryAssetsDiffsBalanceChangesBefore]
type transactionSimulationAccountSummaryAssetsDiffsBalanceChangesBeforeJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationAccountSummaryAssetsDiffsBalanceChangesBefore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationAccountSummaryAssetsDiffsBalanceChangesBeforeJSON) RawJSON() string {
	return r.raw
}

type TransactionSimulationAddressDetail struct {
	// contains the contract's name if the address is a verified contract
	ContractName string `json:"contract_name"`
	// known name tag for the address
	NameTag string                                 `json:"name_tag"`
	JSON    transactionSimulationAddressDetailJSON `json:"-"`
}

// transactionSimulationAddressDetailJSON contains the JSON metadata for the struct
// [TransactionSimulationAddressDetail]
type transactionSimulationAddressDetailJSON struct {
	ContractName apijson.Field
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
	// The state after the transaction
	After TransactionSimulationContractManagementAfter `json:"after,required"`
	// The state before the transaction
	Before TransactionSimulationContractManagementBefore `json:"before,required"`
	// An enumeration.
	Type TransactionSimulationContractManagementType `json:"type,required"`
	JSON transactionSimulationContractManagementJSON `json:"-"`
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

func (r *TransactionSimulationContractManagement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationContractManagementJSON) RawJSON() string {
	return r.raw
}

// The state after the transaction
type TransactionSimulationContractManagementAfter struct {
	Address string `json:"address"`
	// This field can have the runtime type of [[]string].
	Owners interface{} `json:"owners,required"`
	// This field can have the runtime type of [[]string].
	Modules interface{}                                      `json:"modules,required"`
	JSON    transactionSimulationContractManagementAfterJSON `json:"-"`
	union   TransactionSimulationContractManagementAfterUnion
}

// transactionSimulationContractManagementAfterJSON contains the JSON metadata for
// the struct [TransactionSimulationContractManagementAfter]
type transactionSimulationContractManagementAfterJSON struct {
	Address     apijson.Field
	Owners      apijson.Field
	Modules     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionSimulationContractManagementAfterJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationContractManagementAfter) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationContractManagementAfter{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TransactionSimulationContractManagementAfterUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TransactionSimulationContractManagementAfterAddressChange],
// [TransactionSimulationContractManagementAfterOwnershipChange],
// [TransactionSimulationContractManagementAfterModulesChange].
func (r TransactionSimulationContractManagementAfter) AsUnion() TransactionSimulationContractManagementAfterUnion {
	return r.union
}

// The state after the transaction
//
// Union satisfied by [TransactionSimulationContractManagementAfterAddressChange],
// [TransactionSimulationContractManagementAfterOwnershipChange] or
// [TransactionSimulationContractManagementAfterModulesChange].
type TransactionSimulationContractManagementAfterUnion interface {
	implementsTransactionSimulationContractManagementAfter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationContractManagementAfterUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationContractManagementAfterAddressChange{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationContractManagementAfterOwnershipChange{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationContractManagementAfterModulesChange{}),
		},
	)
}

type TransactionSimulationContractManagementAfterAddressChange struct {
	Address string                                                        `json:"address,required"`
	JSON    transactionSimulationContractManagementAfterAddressChangeJSON `json:"-"`
}

// transactionSimulationContractManagementAfterAddressChangeJSON contains the JSON
// metadata for the struct
// [TransactionSimulationContractManagementAfterAddressChange]
type transactionSimulationContractManagementAfterAddressChangeJSON struct {
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationContractManagementAfterAddressChange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationContractManagementAfterAddressChangeJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationContractManagementAfterAddressChange) implementsTransactionSimulationContractManagementAfter() {
}

type TransactionSimulationContractManagementAfterOwnershipChange struct {
	Owners []string                                                        `json:"owners,required"`
	JSON   transactionSimulationContractManagementAfterOwnershipChangeJSON `json:"-"`
}

// transactionSimulationContractManagementAfterOwnershipChangeJSON contains the
// JSON metadata for the struct
// [TransactionSimulationContractManagementAfterOwnershipChange]
type transactionSimulationContractManagementAfterOwnershipChangeJSON struct {
	Owners      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationContractManagementAfterOwnershipChange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationContractManagementAfterOwnershipChangeJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationContractManagementAfterOwnershipChange) implementsTransactionSimulationContractManagementAfter() {
}

type TransactionSimulationContractManagementAfterModulesChange struct {
	Modules []string                                                      `json:"modules,required"`
	JSON    transactionSimulationContractManagementAfterModulesChangeJSON `json:"-"`
}

// transactionSimulationContractManagementAfterModulesChangeJSON contains the JSON
// metadata for the struct
// [TransactionSimulationContractManagementAfterModulesChange]
type transactionSimulationContractManagementAfterModulesChangeJSON struct {
	Modules     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationContractManagementAfterModulesChange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationContractManagementAfterModulesChangeJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationContractManagementAfterModulesChange) implementsTransactionSimulationContractManagementAfter() {
}

// The state before the transaction
type TransactionSimulationContractManagementBefore struct {
	Address string `json:"address"`
	// This field can have the runtime type of [[]string].
	Owners interface{} `json:"owners,required"`
	// This field can have the runtime type of [[]string].
	Modules interface{}                                       `json:"modules,required"`
	JSON    transactionSimulationContractManagementBeforeJSON `json:"-"`
	union   TransactionSimulationContractManagementBeforeUnion
}

// transactionSimulationContractManagementBeforeJSON contains the JSON metadata for
// the struct [TransactionSimulationContractManagementBefore]
type transactionSimulationContractManagementBeforeJSON struct {
	Address     apijson.Field
	Owners      apijson.Field
	Modules     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionSimulationContractManagementBeforeJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionSimulationContractManagementBefore) UnmarshalJSON(data []byte) (err error) {
	*r = TransactionSimulationContractManagementBefore{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TransactionSimulationContractManagementBeforeUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TransactionSimulationContractManagementBeforeAddressChange],
// [TransactionSimulationContractManagementBeforeOwnershipChange],
// [TransactionSimulationContractManagementBeforeModulesChange].
func (r TransactionSimulationContractManagementBefore) AsUnion() TransactionSimulationContractManagementBeforeUnion {
	return r.union
}

// The state before the transaction
//
// Union satisfied by [TransactionSimulationContractManagementBeforeAddressChange],
// [TransactionSimulationContractManagementBeforeOwnershipChange] or
// [TransactionSimulationContractManagementBeforeModulesChange].
type TransactionSimulationContractManagementBeforeUnion interface {
	implementsTransactionSimulationContractManagementBefore()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionSimulationContractManagementBeforeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationContractManagementBeforeAddressChange{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationContractManagementBeforeOwnershipChange{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionSimulationContractManagementBeforeModulesChange{}),
		},
	)
}

type TransactionSimulationContractManagementBeforeAddressChange struct {
	Address string                                                         `json:"address,required"`
	JSON    transactionSimulationContractManagementBeforeAddressChangeJSON `json:"-"`
}

// transactionSimulationContractManagementBeforeAddressChangeJSON contains the JSON
// metadata for the struct
// [TransactionSimulationContractManagementBeforeAddressChange]
type transactionSimulationContractManagementBeforeAddressChangeJSON struct {
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationContractManagementBeforeAddressChange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationContractManagementBeforeAddressChangeJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationContractManagementBeforeAddressChange) implementsTransactionSimulationContractManagementBefore() {
}

type TransactionSimulationContractManagementBeforeOwnershipChange struct {
	Owners []string                                                         `json:"owners,required"`
	JSON   transactionSimulationContractManagementBeforeOwnershipChangeJSON `json:"-"`
}

// transactionSimulationContractManagementBeforeOwnershipChangeJSON contains the
// JSON metadata for the struct
// [TransactionSimulationContractManagementBeforeOwnershipChange]
type transactionSimulationContractManagementBeforeOwnershipChangeJSON struct {
	Owners      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationContractManagementBeforeOwnershipChange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationContractManagementBeforeOwnershipChangeJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationContractManagementBeforeOwnershipChange) implementsTransactionSimulationContractManagementBefore() {
}

type TransactionSimulationContractManagementBeforeModulesChange struct {
	Modules []string                                                       `json:"modules,required"`
	JSON    transactionSimulationContractManagementBeforeModulesChangeJSON `json:"-"`
}

// transactionSimulationContractManagementBeforeModulesChangeJSON contains the JSON
// metadata for the struct
// [TransactionSimulationContractManagementBeforeModulesChange]
type transactionSimulationContractManagementBeforeModulesChangeJSON struct {
	Modules     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionSimulationContractManagementBeforeModulesChange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionSimulationContractManagementBeforeModulesChangeJSON) RawJSON() string {
	return r.raw
}

func (r TransactionSimulationContractManagementBeforeModulesChange) implementsTransactionSimulationContractManagementBefore() {
}

// An enumeration.
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
	// The value to be sent.
	Value string                          `json:"value"`
	JSON  transactionSimulationParamsJSON `json:"-"`
}

// transactionSimulationParamsJSON contains the JSON metadata for the struct
// [TransactionSimulationParams]
type transactionSimulationParamsJSON struct {
	BlockTag    apijson.Field
	Calldata    apijson.Field
	Chain       apijson.Field
	Data        apijson.Field
	From        apijson.Field
	Gas         apijson.Field
	GasPrice    apijson.Field
	To          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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

type TransactionSimulationError struct {
	// An error message if the simulation failed.
	Error string `json:"error,required"`
	// A string indicating if the simulation was successful or not.
	Status TransactionSimulationErrorStatus `json:"status,required"`
	JSON   transactionSimulationErrorJSON   `json:"-"`
}

// transactionSimulationErrorJSON contains the JSON metadata for the struct
// [TransactionSimulationError]
type transactionSimulationErrorJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
