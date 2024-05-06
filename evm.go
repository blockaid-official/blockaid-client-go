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
// the blockaid API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewEvmService] method instead.
type EvmService struct {
	Options         []option.RequestOption
	JsonRpc         *EvmJsonRpcService
	Transaction     *EvmTransactionService
	TransactionBulk *EvmTransactionBulkService
	TransactionRaw  *EvmTransactionRawService
	UserOperation   *EvmUserOperationService
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
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

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
	Exposure interface{} `json:"exposure"`
	// user friendly description of the approval
	Summary string `json:"summary"`
	// the amount that was asked in the approval request for this spender from the
	// current address and asset
	Approval int64 `json:"approval"`
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
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

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
	// amount of the asset that was transferred to the address in this transaction
	In []AssetDiffIn `json:"in,required"`
	// amount of the asset that was transferred from the address in this transaction
	Out  []AssetDiffOut `json:"out,required"`
	JSON assetDiffJSON  `json:"-"`
}

// assetDiffJSON contains the JSON metadata for the struct [AssetDiff]
type assetDiffJSON struct {
	Asset       apijson.Field
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
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

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
	UsdPrice float64 `json:"usd_price"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// id of the token
	TokenID int64       `json:"token_id"`
	Value   interface{} `json:"value,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// value before divided by decimal, that was transferred from this address
	RawValue int64           `json:"raw_value"`
	JSON     assetDiffInJSON `json:"-"`
	union    AssetDiffInUnion
}

// assetDiffInJSON contains the JSON metadata for the struct [AssetDiffIn]
type assetDiffInJSON struct {
	UsdPrice    apijson.Field
	Summary     apijson.Field
	TokenID     apijson.Field
	Value       apijson.Field
	LogoURL     apijson.Field
	RawValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r assetDiffInJSON) RawJSON() string {
	return r.raw
}

func (r *AssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

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
	UsdPrice float64 `json:"usd_price"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// id of the token
	TokenID int64       `json:"token_id"`
	Value   interface{} `json:"value,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// value before divided by decimal, that was transferred from this address
	RawValue int64            `json:"raw_value"`
	JSON     assetDiffOutJSON `json:"-"`
	union    AssetDiffOutUnion
}

// assetDiffOutJSON contains the JSON metadata for the struct [AssetDiffOut]
type assetDiffOutJSON struct {
	UsdPrice    apijson.Field
	Summary     apijson.Field
	TokenID     apijson.Field
	Value       apijson.Field
	LogoURL     apijson.Field
	RawValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r assetDiffOutJSON) RawJSON() string {
	return r.raw
}

func (r *AssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

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
	// id of the token
	TokenID int64 `json:"token_id,required"`
	// value before divided by decimal, that was transferred from this address
	Value int64 `json:"value,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice float64         `json:"usd_price"`
	JSON     erc1155DiffJSON `json:"-"`
}

// erc1155DiffJSON contains the JSON metadata for the struct [Erc1155Diff]
type erc1155DiffJSON struct {
	TokenID     apijson.Field
	Value       apijson.Field
	LogoURL     apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
	UsdPrice float64 `json:"usd_price"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// id of the token
	TokenID int64       `json:"token_id"`
	Value   interface{} `json:"value,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// value before divided by decimal, that was transferred from this address
	RawValue int64                       `json:"raw_value"`
	JSON     erc1155ExposureExposureJSON `json:"-"`
	union    Erc1155ExposureExposureUnion
}

// erc1155ExposureExposureJSON contains the JSON metadata for the struct
// [Erc1155ExposureExposure]
type erc1155ExposureExposureJSON struct {
	UsdPrice    apijson.Field
	Summary     apijson.Field
	TokenID     apijson.Field
	Value       apijson.Field
	LogoURL     apijson.Field
	RawValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r erc1155ExposureExposureJSON) RawJSON() string {
	return r.raw
}

func (r *Erc1155ExposureExposure) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

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
	RawValue int64 `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice float64 `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value float64       `json:"value"`
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

type Erc20Exposure struct {
	// the amount that was asked in the approval request for this spender from the
	// current address and asset
	Approval int64                   `json:"approval,required"`
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
	UsdPrice float64 `json:"usd_price"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// id of the token
	TokenID int64       `json:"token_id"`
	Value   interface{} `json:"value,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// value before divided by decimal, that was transferred from this address
	RawValue int64                     `json:"raw_value"`
	JSON     erc20ExposureExposureJSON `json:"-"`
	union    Erc20ExposureExposureUnion
}

// erc20ExposureExposureJSON contains the JSON metadata for the struct
// [Erc20ExposureExposure]
type erc20ExposureExposureJSON struct {
	UsdPrice    apijson.Field
	Summary     apijson.Field
	TokenID     apijson.Field
	Value       apijson.Field
	LogoURL     apijson.Field
	RawValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r erc20ExposureExposureJSON) RawJSON() string {
	return r.raw
}

func (r *Erc20ExposureExposure) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

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
	// id of the token
	TokenID int64 `json:"token_id,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice float64        `json:"usd_price"`
	JSON     erc721DiffJSON `json:"-"`
}

// erc721DiffJSON contains the JSON metadata for the struct [Erc721Diff]
type erc721DiffJSON struct {
	TokenID     apijson.Field
	LogoURL     apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
	UsdPrice float64 `json:"usd_price"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// id of the token
	TokenID int64       `json:"token_id"`
	Value   interface{} `json:"value,required"`
	// url of the token logo
	LogoURL string `json:"logo_url"`
	// value before divided by decimal, that was transferred from this address
	RawValue int64                      `json:"raw_value"`
	JSON     erc721ExposureExposureJSON `json:"-"`
	union    Erc721ExposureExposureUnion
}

// erc721ExposureExposureJSON contains the JSON metadata for the struct
// [Erc721ExposureExposure]
type erc721ExposureExposureJSON struct {
	UsdPrice    apijson.Field
	Summary     apijson.Field
	TokenID     apijson.Field
	Value       apijson.Field
	LogoURL     apijson.Field
	RawValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r erc721ExposureExposureJSON) RawJSON() string {
	return r.raw
}

func (r *Erc721ExposureExposure) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

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
	RawValue int64 `json:"raw_value,required"`
	// user friendly description of the asset transfer
	Summary string `json:"summary"`
	// usd equal of the asset that was transferred from this address
	UsdPrice float64 `json:"usd_price"`
	// value after divided by decimals, that was transferred from this address
	Value float64        `json:"value"`
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
	Block         string                               `json:"block,required"`
	Chain         string                               `json:"chain,required"`
	Events        []TransactionScanResponseEvent       `json:"events"`
	Features      interface{}                          `json:"features"`
	GasEstimation TransactionScanResponseGasEstimation `json:"gas_estimation"`
	Simulation    TransactionScanResponseSimulation    `json:"simulation"`
	Validation    TransactionScanResponseValidation    `json:"validation"`
	JSON          transactionScanResponseJSON          `json:"-"`
}

// transactionScanResponseJSON contains the JSON metadata for the struct
// [TransactionScanResponse]
type transactionScanResponseJSON struct {
	Block         apijson.Field
	Chain         apijson.Field
	Events        apijson.Field
	Features      apijson.Field
	GasEstimation apijson.Field
	Simulation    apijson.Field
	Validation    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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

// Union satisfied by [shared.UnionString],
// [TransactionScanResponseEventsParamsValueUnknown] or
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
	Used     int64                                    `json:"used"`
	Estimate int64                                    `json:"estimate"`
	Error    string                                   `json:"error"`
	JSON     transactionScanResponseGasEstimationJSON `json:"-"`
	union    TransactionScanResponseGasEstimationUnion
}

// transactionScanResponseGasEstimationJSON contains the JSON metadata for the
// struct [TransactionScanResponseGasEstimation]
type transactionScanResponseGasEstimationJSON struct {
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
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

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
	Estimate int64                                                                `json:"estimate,required"`
	Used     int64                                                                `json:"used,required"`
	JSON     transactionScanResponseGasEstimationTransactionScanGasEstimationJSON `json:"-"`
}

// transactionScanResponseGasEstimationTransactionScanGasEstimationJSON contains
// the JSON metadata for the struct
// [TransactionScanResponseGasEstimationTransactionScanGasEstimation]
type transactionScanResponseGasEstimationTransactionScanGasEstimationJSON struct {
	Estimate    apijson.Field
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

type TransactionScanResponseGasEstimationTransactionScanGasEstimationError struct {
	Error string                                                                    `json:"error,required"`
	JSON  transactionScanResponseGasEstimationTransactionScanGasEstimationErrorJSON `json:"-"`
}

// transactionScanResponseGasEstimationTransactionScanGasEstimationErrorJSON
// contains the JSON metadata for the struct
// [TransactionScanResponseGasEstimationTransactionScanGasEstimationError]
type transactionScanResponseGasEstimationTransactionScanGasEstimationErrorJSON struct {
	Error       apijson.Field
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

type TransactionScanResponseSimulation struct {
	AssetsDiffs      interface{} `json:"assets_diffs,required"`
	TotalUsdDiff     interface{} `json:"total_usd_diff,required"`
	Exposures        interface{} `json:"exposures,required"`
	TotalUsdExposure interface{} `json:"total_usd_exposure,required"`
	AddressDetails   interface{} `json:"address_details,required"`
	AccountSummary   interface{} `json:"account_summary,required"`
	// An error message if the simulation failed.
	Error string                                `json:"error"`
	JSON  transactionScanResponseSimulationJSON `json:"-"`
	union TransactionScanResponseSimulationUnion
}

// transactionScanResponseSimulationJSON contains the JSON metadata for the struct
// [TransactionScanResponseSimulation]
type transactionScanResponseSimulationJSON struct {
	AssetsDiffs      apijson.Field
	TotalUsdDiff     apijson.Field
	Exposures        apijson.Field
	TotalUsdExposure apijson.Field
	AddressDetails   apijson.Field
	AccountSummary   apijson.Field
	Error            apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r transactionScanResponseSimulationJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionScanResponseSimulation) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

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

type TransactionScanResponseValidation struct {
	// An enumeration.
	ResultType TransactionScanResponseValidationResultType `json:"result_type"`
	// A textual description that can be presented to the user about what this
	// transaction is doing.
	Description string `json:"description"`
	// A textual description about the reasons the transaction was flagged with
	// result_type.
	Reason string `json:"reason"`
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string                                `json:"classification"`
	Features       interface{}                           `json:"features,required"`
	Loc            interface{}                           `json:"loc,required"`
	Msg            string                                `json:"msg"`
	Type           string                                `json:"type"`
	JSON           transactionScanResponseValidationJSON `json:"-"`
	union          TransactionScanResponseValidationUnion
}

// transactionScanResponseValidationJSON contains the JSON metadata for the struct
// [TransactionScanResponseValidation]
type transactionScanResponseValidationJSON struct {
	ResultType     apijson.Field
	Description    apijson.Field
	Reason         apijson.Field
	Classification apijson.Field
	Features       apijson.Field
	Loc            apijson.Field
	Msg            apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r transactionScanResponseValidationJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionScanResponseValidation) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

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
)

func (r TransactionScanResponseValidationResultType) IsKnown() bool {
	switch r {
	case TransactionScanResponseValidationResultTypeBenign, TransactionScanResponseValidationResultTypeWarning, TransactionScanResponseValidationResultTypeMalicious:
		return true
	}
	return false
}

// The chain name
type TransactionScanSupportedChain string

const (
	TransactionScanSupportedChainArbitrum    TransactionScanSupportedChain = "arbitrum"
	TransactionScanSupportedChainAvalanche   TransactionScanSupportedChain = "avalanche"
	TransactionScanSupportedChainBase        TransactionScanSupportedChain = "base"
	TransactionScanSupportedChainBaseSepolia TransactionScanSupportedChain = "base-sepolia"
	TransactionScanSupportedChainBsc         TransactionScanSupportedChain = "bsc"
	TransactionScanSupportedChainEthereum    TransactionScanSupportedChain = "ethereum"
	TransactionScanSupportedChainOptimism    TransactionScanSupportedChain = "optimism"
	TransactionScanSupportedChainPolygon     TransactionScanSupportedChain = "polygon"
	TransactionScanSupportedChainZksync      TransactionScanSupportedChain = "zksync"
	TransactionScanSupportedChainZora        TransactionScanSupportedChain = "zora"
	TransactionScanSupportedChainLinea       TransactionScanSupportedChain = "linea"
	TransactionScanSupportedChainBlast       TransactionScanSupportedChain = "blast"
	TransactionScanSupportedChainUnknown     TransactionScanSupportedChain = "unknown"
)

func (r TransactionScanSupportedChain) IsKnown() bool {
	switch r {
	case TransactionScanSupportedChainArbitrum, TransactionScanSupportedChainAvalanche, TransactionScanSupportedChainBase, TransactionScanSupportedChainBaseSepolia, TransactionScanSupportedChainBsc, TransactionScanSupportedChainEthereum, TransactionScanSupportedChainOptimism, TransactionScanSupportedChainPolygon, TransactionScanSupportedChainZksync, TransactionScanSupportedChainZora, TransactionScanSupportedChainLinea, TransactionScanSupportedChainBlast, TransactionScanSupportedChainUnknown:
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
	// dictionary represents the usd value each address gained / lost during this
	// transaction
	TotalUsdDiff map[string]UsdDiff `json:"total_usd_diff,required"`
	// a dictionary representing the usd value each address is exposed to, split by
	// spenders
	TotalUsdExposure map[string]map[string]string `json:"total_usd_exposure,required"`
	JSON             transactionSimulationJSON    `json:"-"`
}

// transactionSimulationJSON contains the JSON metadata for the struct
// [TransactionSimulation]
type transactionSimulationJSON struct {
	AccountSummary   apijson.Field
	AddressDetails   apijson.Field
	AssetsDiffs      apijson.Field
	Exposures        apijson.Field
	TotalUsdDiff     apijson.Field
	TotalUsdExposure apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
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
	AssetsDiffs []AssetDiff `json:"assets_diffs,required"`
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

type TransactionSimulationError struct {
	// An error message if the simulation failed.
	Error string                         `json:"error,required"`
	JSON  transactionSimulationErrorJSON `json:"-"`
}

// transactionSimulationErrorJSON contains the JSON metadata for the struct
// [TransactionSimulationError]
type transactionSimulationErrorJSON struct {
	Error       apijson.Field
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

type TransactionValidation struct {
	// A list of features about this transaction explaining the validation.
	Features []TransactionScanFeature `json:"features,required"`
	// An enumeration.
	ResultType TransactionValidationResultType `json:"result_type,required"`
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

type TransactionValidationError struct {
	Loc  []TransactionValidationErrorLocUnion `json:"loc,required"`
	Msg  string                               `json:"msg,required"`
	Type string                               `json:"type,required"`
	JSON transactionValidationErrorJSON       `json:"-"`
}

// transactionValidationErrorJSON contains the JSON metadata for the struct
// [TransactionValidationError]
type transactionValidationErrorJSON struct {
	Loc         apijson.Field
	Msg         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionValidationErrorJSON) RawJSON() string {
	return r.raw
}

func (r TransactionValidationError) implementsTransactionScanResponseValidation() {}

// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type TransactionValidationErrorLocUnion interface {
	ImplementsTransactionValidationErrorLocUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionValidationErrorLocUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
	)
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
