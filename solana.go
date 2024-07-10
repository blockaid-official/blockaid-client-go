// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"reflect"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/param"
	"github.com/blockaid-official/blockaid-client-go/option"
	"github.com/tidwall/gjson"
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

type AccountSummarySchema struct {
	// Total USD diff for the requested account address
	TotalUsdDiff AccountSummarySchemaTotalUsdDiff `json:"total_usd_diff,required"`
	// Assets diff of the requested account address
	AccountAssetsDiff []AccountSummarySchemaAccountAssetsDiff `json:"account_assets_diff"`
	// Delegated assets of the requested account address
	AccountDelegations []DelegatedAssetDetailsSchema `json:"account_delegations"`
	// Account ownerships diff of the requested account address
	AccountOwnershipsDiff []AccountSummarySchemaAccountOwnershipsDiff `json:"account_ownerships_diff"`
	JSON                  accountSummarySchemaJSON                    `json:"-"`
}

// accountSummarySchemaJSON contains the JSON metadata for the struct
// [AccountSummarySchema]
type accountSummarySchemaJSON struct {
	TotalUsdDiff          apijson.Field
	AccountAssetsDiff     apijson.Field
	AccountDelegations    apijson.Field
	AccountOwnershipsDiff apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccountSummarySchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaJSON) RawJSON() string {
	return r.raw
}

// Total USD diff for the requested account address
type AccountSummarySchemaTotalUsdDiff struct {
	// Total incoming USD transfers
	In float64 `json:"in,required"`
	// Total outgoing USD transfers
	Out float64 `json:"out,required"`
	// Total USD transfers
	Total float64                              `json:"total,required"`
	JSON  accountSummarySchemaTotalUsdDiffJSON `json:"-"`
}

// accountSummarySchemaTotalUsdDiffJSON contains the JSON metadata for the struct
// [AccountSummarySchemaTotalUsdDiff]
type accountSummarySchemaTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

type AccountSummarySchemaAccountAssetsDiff struct {
	// This field can have the runtime type of
	// [AccountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaAsset],
	// [AccountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaAsset],
	// [AccountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaAsset],
	// [AccountSummarySchemaAccountAssetsDiffCnftDiffSchemaAsset].
	Asset interface{} `json:"asset"`
	// This field can have the runtime type of
	// [AccountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaIn],
	// [AccountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaIn],
	// [AccountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaIn],
	// [AccountSummarySchemaAccountAssetsDiffCnftDiffSchemaIn].
	In interface{} `json:"in,required"`
	// This field can have the runtime type of
	// [AccountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaOut],
	// [AccountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaOut],
	// [AccountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaOut],
	// [AccountSummarySchemaAccountAssetsDiffCnftDiffSchemaOut].
	Out   interface{}                               `json:"out,required"`
	JSON  accountSummarySchemaAccountAssetsDiffJSON `json:"-"`
	union AccountSummarySchemaAccountAssetsDiffUnion
}

// accountSummarySchemaAccountAssetsDiffJSON contains the JSON metadata for the
// struct [AccountSummarySchemaAccountAssetsDiff]
type accountSummarySchemaAccountAssetsDiffJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r accountSummarySchemaAccountAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *AccountSummarySchemaAccountAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccountSummarySchemaAccountAssetsDiffUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AccountSummarySchemaAccountAssetsDiffNativeSolDiffSchema],
// [AccountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchema],
// [AccountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchema],
// [AccountSummarySchemaAccountAssetsDiffCnftDiffSchema].
func (r AccountSummarySchemaAccountAssetsDiff) AsUnion() AccountSummarySchemaAccountAssetsDiffUnion {
	return r.union
}

// Union satisfied by [AccountSummarySchemaAccountAssetsDiffNativeSolDiffSchema],
// [AccountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchema],
// [AccountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchema] or
// [AccountSummarySchemaAccountAssetsDiffCnftDiffSchema].
type AccountSummarySchemaAccountAssetsDiffUnion interface {
	implementsAccountSummarySchemaAccountAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountSummarySchemaAccountAssetsDiffUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountSummarySchemaAccountAssetsDiffNativeSolDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountSummarySchemaAccountAssetsDiffCnftDiffSchema{}),
		},
	)
}

type AccountSummarySchemaAccountAssetsDiffNativeSolDiffSchema struct {
	Asset AccountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaAsset `json:"asset,required"`
	// Incoming transfers of the asset
	In   AccountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaIn   `json:"in,nullable"`
	Out  AccountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaOut  `json:"out,nullable"`
	JSON accountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaJSON `json:"-"`
}

// accountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaJSON contains the JSON
// metadata for the struct
// [AccountSummarySchemaAccountAssetsDiffNativeSolDiffSchema]
type accountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountAssetsDiffNativeSolDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r AccountSummarySchemaAccountAssetsDiffNativeSolDiffSchema) implementsAccountSummarySchemaAccountAssetsDiff() {
}

type AccountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaAsset struct {
	Decimals int64 `json:"decimals"`
	// Type of the asset (`"SOL"`)
	Type string                                                            `json:"type"`
	JSON accountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaAssetJSON `json:"-"`
}

// accountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaAssetJSON contains the
// JSON metadata for the struct
// [AccountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaAsset]
type accountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Incoming transfers of the asset
type AccountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                        `json:"usd_price,nullable"`
	JSON     accountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaInJSON `json:"-"`
}

// accountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaInJSON contains the JSON
// metadata for the struct
// [AccountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaIn]
type accountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

type AccountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                         `json:"usd_price,nullable"`
	JSON     accountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaOutJSON `json:"-"`
}

// accountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaOutJSON contains the
// JSON metadata for the struct
// [AccountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaOut]
type accountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountAssetsDiffNativeSolDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type AccountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchema struct {
	Asset AccountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaAsset `json:"asset,required"`
	// Incoming transfers of the asset
	In   AccountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaIn   `json:"in,nullable"`
	Out  AccountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaOut  `json:"out,nullable"`
	JSON accountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaJSON `json:"-"`
}

// accountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaJSON contains the
// JSON metadata for the struct
// [AccountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchema]
type accountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r AccountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchema) implementsAccountSummarySchemaAccountAssetsDiff() {
}

type AccountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaAsset struct {
	Address  string `json:"address,required"`
	Decimals int64  `json:"decimals,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Logo     string `json:"logo"`
	// Type of the asset (`"TOKEN"`)
	Type string                                                                   `json:"type"`
	JSON accountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaAssetJSON `json:"-"`
}

// accountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [AccountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaAsset]
type accountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Incoming transfers of the asset
type AccountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                               `json:"usd_price,nullable"`
	JSON     accountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaInJSON `json:"-"`
}

// accountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaInJSON contains
// the JSON metadata for the struct
// [AccountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaIn]
type accountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

type AccountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                `json:"usd_price,nullable"`
	JSON     accountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaOutJSON `json:"-"`
}

// accountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaOutJSON contains
// the JSON metadata for the struct
// [AccountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaOut]
type accountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountAssetsDiffSplFungibleTokenDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type AccountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchema struct {
	Asset AccountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaAsset `json:"asset,required"`
	// Incoming transfers of the asset
	In   AccountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaIn   `json:"in,nullable"`
	Out  AccountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaOut  `json:"out,nullable"`
	JSON accountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaJSON `json:"-"`
}

// accountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaJSON contains
// the JSON metadata for the struct
// [AccountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchema]
type accountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r AccountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchema) implementsAccountSummarySchemaAccountAssetsDiff() {
}

type AccountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaAsset struct {
	Address  string `json:"address,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	Logo     string `json:"logo"`
	// Type of the asset (`"NFT"`)
	Type string                                                                      `json:"type"`
	JSON accountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaAssetJSON `json:"-"`
}

// accountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [AccountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaAsset]
type accountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Incoming transfers of the asset
type AccountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                  `json:"usd_price,nullable"`
	JSON     accountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaInJSON `json:"-"`
}

// accountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaInJSON
// contains the JSON metadata for the struct
// [AccountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaIn]
type accountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

type AccountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                   `json:"usd_price,nullable"`
	JSON     accountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaOutJSON `json:"-"`
}

// accountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaOutJSON
// contains the JSON metadata for the struct
// [AccountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaOut]
type accountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountAssetsDiffSplNonFungibleTokenDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type AccountSummarySchemaAccountAssetsDiffCnftDiffSchema struct {
	Asset AccountSummarySchemaAccountAssetsDiffCnftDiffSchemaAsset `json:"asset,required"`
	// Incoming transfers of the asset
	In   AccountSummarySchemaAccountAssetsDiffCnftDiffSchemaIn   `json:"in,nullable"`
	Out  AccountSummarySchemaAccountAssetsDiffCnftDiffSchemaOut  `json:"out,nullable"`
	JSON accountSummarySchemaAccountAssetsDiffCnftDiffSchemaJSON `json:"-"`
}

// accountSummarySchemaAccountAssetsDiffCnftDiffSchemaJSON contains the JSON
// metadata for the struct [AccountSummarySchemaAccountAssetsDiffCnftDiffSchema]
type accountSummarySchemaAccountAssetsDiffCnftDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountAssetsDiffCnftDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountAssetsDiffCnftDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r AccountSummarySchemaAccountAssetsDiffCnftDiffSchema) implementsAccountSummarySchemaAccountAssetsDiff() {
}

type AccountSummarySchemaAccountAssetsDiffCnftDiffSchemaAsset struct {
	Address  string `json:"address,required"`
	Decimals int64  `json:"decimals,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Logo     string `json:"logo"`
	// Type of the asset (`"CNFT"`)
	Type string                                                       `json:"type"`
	JSON accountSummarySchemaAccountAssetsDiffCnftDiffSchemaAssetJSON `json:"-"`
}

// accountSummarySchemaAccountAssetsDiffCnftDiffSchemaAssetJSON contains the JSON
// metadata for the struct
// [AccountSummarySchemaAccountAssetsDiffCnftDiffSchemaAsset]
type accountSummarySchemaAccountAssetsDiffCnftDiffSchemaAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountAssetsDiffCnftDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountAssetsDiffCnftDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Incoming transfers of the asset
type AccountSummarySchemaAccountAssetsDiffCnftDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                   `json:"usd_price,nullable"`
	JSON     accountSummarySchemaAccountAssetsDiffCnftDiffSchemaInJSON `json:"-"`
}

// accountSummarySchemaAccountAssetsDiffCnftDiffSchemaInJSON contains the JSON
// metadata for the struct [AccountSummarySchemaAccountAssetsDiffCnftDiffSchemaIn]
type accountSummarySchemaAccountAssetsDiffCnftDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountAssetsDiffCnftDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountAssetsDiffCnftDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

type AccountSummarySchemaAccountAssetsDiffCnftDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                    `json:"usd_price,nullable"`
	JSON     accountSummarySchemaAccountAssetsDiffCnftDiffSchemaOutJSON `json:"-"`
}

// accountSummarySchemaAccountAssetsDiffCnftDiffSchemaOutJSON contains the JSON
// metadata for the struct [AccountSummarySchemaAccountAssetsDiffCnftDiffSchemaOut]
type accountSummarySchemaAccountAssetsDiffCnftDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountAssetsDiffCnftDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountAssetsDiffCnftDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type AccountSummarySchemaAccountOwnershipsDiff struct {
	// This field can have the runtime type of
	// [AccountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaAsset],
	// [AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset],
	// [AccountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAsset].
	Asset interface{} `json:"asset,required"`
	// This field can have the runtime type of
	// [AccountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaIn],
	// [AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaIn],
	// [AccountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaIn].
	In interface{} `json:"in_,required"`
	// This field can have the runtime type of
	// [AccountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaOut],
	// [AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaOut],
	// [AccountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaOut].
	Out interface{} `json:"out,required"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,nullable"`
	// The owner post the transaction
	PostOwner string                                        `json:"post_owner,required"`
	JSON      accountSummarySchemaAccountOwnershipsDiffJSON `json:"-"`
	union     AccountSummarySchemaAccountOwnershipsDiffUnion
}

// accountSummarySchemaAccountOwnershipsDiffJSON contains the JSON metadata for the
// struct [AccountSummarySchemaAccountOwnershipsDiff]
type accountSummarySchemaAccountOwnershipsDiffJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	PostOwner   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r accountSummarySchemaAccountOwnershipsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *AccountSummarySchemaAccountOwnershipsDiff) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccountSummarySchemaAccountOwnershipsDiffUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AccountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchema],
// [AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchema],
// [AccountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchema].
func (r AccountSummarySchemaAccountOwnershipsDiff) AsUnion() AccountSummarySchemaAccountOwnershipsDiffUnion {
	return r.union
}

// Union satisfied by
// [AccountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchema],
// [AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchema] or
// [AccountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchema].
type AccountSummarySchemaAccountOwnershipsDiffUnion interface {
	implementsAccountSummarySchemaAccountOwnershipsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountSummarySchemaAccountOwnershipsDiffUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchema{}),
		},
	)
}

type AccountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchema struct {
	Asset AccountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaAsset `json:"asset,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required"`
	// Incoming transfers of the asset
	In AccountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaIn `json:"in_,nullable"`
	// Details of the moved value
	Out AccountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaOut `json:"out,nullable"`
	// The owner prior to the transaction
	PreOwner string                                                                    `json:"pre_owner,nullable"`
	JSON     accountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaJSON `json:"-"`
}

// accountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaJSON
// contains the JSON metadata for the struct
// [AccountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchema]
type accountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaJSON struct {
	Asset       apijson.Field
	PostOwner   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r AccountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchema) implementsAccountSummarySchemaAccountOwnershipsDiff() {
}

type AccountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaAsset struct {
	Decimals int64 `json:"decimals"`
	// Type of the asset (`"SOL"`)
	Type string                                                                         `json:"type"`
	JSON accountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaAssetJSON `json:"-"`
}

// accountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [AccountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaAsset]
type accountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Incoming transfers of the asset
type AccountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                     `json:"usd_price,nullable"`
	JSON     accountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaInJSON `json:"-"`
}

// accountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaInJSON
// contains the JSON metadata for the struct
// [AccountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaIn]
type accountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

// Details of the moved value
type AccountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                      `json:"usd_price,nullable"`
	JSON     accountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaOutJSON `json:"-"`
}

// accountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaOutJSON
// contains the JSON metadata for the struct
// [AccountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaOut]
type accountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountOwnershipsDiffNativeSolOwnershipDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchema struct {
	Asset AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset `json:"asset,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required"`
	// Incoming transfers of the asset
	In AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaIn `json:"in_,nullable"`
	// Details of the moved value
	Out AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaOut `json:"out,nullable"`
	// The owner prior to the transaction
	PreOwner string                                                                   `json:"pre_owner,nullable"`
	JSON     accountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaJSON `json:"-"`
}

// accountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaJSON
// contains the JSON metadata for the struct
// [AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchema]
type accountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaJSON struct {
	Asset       apijson.Field
	PostOwner   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchema) implementsAccountSummarySchemaAccountOwnershipsDiff() {
}

type AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset struct {
	Address string `json:"address,required"`
	Symbol  string `json:"symbol,required"`
	Name    string `json:"name,required"`
	Logo    string `json:"logo"`
	// Type of the asset (`"TOKEN"`)
	Type     string                                                                        `json:"type"`
	Decimals int64                                                                         `json:"decimals"`
	JSON     accountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetJSON `json:"-"`
	union    AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetUnion
}

// accountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset]
type accountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetJSON struct {
	Address     apijson.Field
	Symbol      apijson.Field
	Name        apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r accountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

func (r *AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema],
// [AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema].
func (r AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset) AsUnion() AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetUnion {
	return r.union
}

// Union satisfied by
// [AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema]
// or
// [AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema].
type AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetUnion interface {
	implementsAccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema{}),
		},
	)
}

type AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema struct {
	Address  string `json:"address,required"`
	Decimals int64  `json:"decimals,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Logo     string `json:"logo"`
	// Type of the asset (`"TOKEN"`)
	Type string                                                                                                     `json:"type"`
	JSON accountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchemaJSON `json:"-"`
}

// accountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchemaJSON
// contains the JSON metadata for the struct
// [AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema]
type accountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchemaJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema) implementsAccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset() {
}

type AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema struct {
	Address  string `json:"address,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	Logo     string `json:"logo"`
	// Type of the asset (`"NFT"`)
	Type string                                                                                                        `json:"type"`
	JSON accountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchemaJSON `json:"-"`
}

// accountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchemaJSON
// contains the JSON metadata for the struct
// [AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema]
type accountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchemaJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema) implementsAccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset() {
}

// Incoming transfers of the asset
type AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                    `json:"usd_price,nullable"`
	JSON     accountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaInJSON `json:"-"`
}

// accountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaInJSON
// contains the JSON metadata for the struct
// [AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaIn]
type accountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

// Details of the moved value
type AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                     `json:"usd_price,nullable"`
	JSON     accountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaOutJSON `json:"-"`
}

// accountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaOutJSON
// contains the JSON metadata for the struct
// [AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaOut]
type accountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountOwnershipsDiffSplTokenOwnershipDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type AccountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchema struct {
	// The owner post the transaction
	PostOwner string                                                                             `json:"post_owner,required"`
	Asset     AccountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAsset `json:"asset"`
	// Incoming transfers of the asset
	In AccountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaIn `json:"in_,nullable"`
	// Details of the moved value
	Out AccountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaOut `json:"out,nullable"`
	// The owner prior to the transaction
	PreOwner string                                                                            `json:"pre_owner,nullable"`
	JSON     accountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaJSON `json:"-"`
}

// accountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaJSON
// contains the JSON metadata for the struct
// [AccountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchema]
type accountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaJSON struct {
	PostOwner   apijson.Field
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r AccountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchema) implementsAccountSummarySchemaAccountOwnershipsDiff() {
}

type AccountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAsset struct {
	Decimals int64 `json:"decimals"`
	// Type of the asset (`"STAKED_SOL"`)
	Type string                                                                                 `json:"type"`
	JSON accountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAssetJSON `json:"-"`
}

// accountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [AccountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAsset]
type accountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Incoming transfers of the asset
type AccountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                             `json:"usd_price,nullable"`
	JSON     accountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaInJSON `json:"-"`
}

// accountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaInJSON
// contains the JSON metadata for the struct
// [AccountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaIn]
type accountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

// Details of the moved value
type AccountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                              `json:"usd_price,nullable"`
	JSON     accountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaOutJSON `json:"-"`
}

// accountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaOutJSON
// contains the JSON metadata for the struct
// [AccountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaOut]
type accountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSummarySchemaAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type AddressScanRequestSchemaParam struct {
	// Encoded public key
	Address  param.Field[string]                                `json:"address,required"`
	Metadata param.Field[AddressScanRequestSchemaMetadataParam] `json:"metadata,required"`
	// Chain to scan the transaction on
	Chain param.Field[string] `json:"chain"`
}

func (r AddressScanRequestSchemaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressScanRequestSchemaMetadataParam struct {
	// URL of the dApp related to the address
	URL param.Field[string] `json:"url,required"`
}

func (r AddressScanRequestSchemaMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressScanResponseSchema struct {
	// Features about the result
	Features []AddressScanResponseSchemaFeature `json:"features,required"`
	// An enumeration.
	ResultType AddressScanResponseSchemaResultType `json:"result_type,required"`
	JSON       addressScanResponseSchemaJSON       `json:"-"`
}

// addressScanResponseSchemaJSON contains the JSON metadata for the struct
// [AddressScanResponseSchema]
type addressScanResponseSchemaJSON struct {
	Features    apijson.Field
	ResultType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressScanResponseSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressScanResponseSchemaJSON) RawJSON() string {
	return r.raw
}

type AddressScanResponseSchemaFeature struct {
	// Description of the feature
	Description string `json:"description,required"`
	// ID of the feature
	FeatureID string `json:"feature_id,required"`
	// An enumeration.
	Type AddressScanResponseSchemaFeaturesType `json:"type,required"`
	JSON addressScanResponseSchemaFeatureJSON  `json:"-"`
}

// addressScanResponseSchemaFeatureJSON contains the JSON metadata for the struct
// [AddressScanResponseSchemaFeature]
type addressScanResponseSchemaFeatureJSON struct {
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressScanResponseSchemaFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressScanResponseSchemaFeatureJSON) RawJSON() string {
	return r.raw
}

// An enumeration.
type AddressScanResponseSchemaFeaturesType string

const (
	AddressScanResponseSchemaFeaturesTypeMalicious AddressScanResponseSchemaFeaturesType = "Malicious"
	AddressScanResponseSchemaFeaturesTypeWarning   AddressScanResponseSchemaFeaturesType = "Warning"
	AddressScanResponseSchemaFeaturesTypeBenign    AddressScanResponseSchemaFeaturesType = "Benign"
	AddressScanResponseSchemaFeaturesTypeInfo      AddressScanResponseSchemaFeaturesType = "Info"
)

func (r AddressScanResponseSchemaFeaturesType) IsKnown() bool {
	switch r {
	case AddressScanResponseSchemaFeaturesTypeMalicious, AddressScanResponseSchemaFeaturesTypeWarning, AddressScanResponseSchemaFeaturesTypeBenign, AddressScanResponseSchemaFeaturesTypeInfo:
		return true
	}
	return false
}

// An enumeration.
type AddressScanResponseSchemaResultType string

const (
	AddressScanResponseSchemaResultTypeMalicious AddressScanResponseSchemaResultType = "Malicious"
	AddressScanResponseSchemaResultTypeWarning   AddressScanResponseSchemaResultType = "Warning"
	AddressScanResponseSchemaResultTypeBenign    AddressScanResponseSchemaResultType = "Benign"
)

func (r AddressScanResponseSchemaResultType) IsKnown() bool {
	switch r {
	case AddressScanResponseSchemaResultTypeMalicious, AddressScanResponseSchemaResultTypeWarning, AddressScanResponseSchemaResultTypeBenign:
		return true
	}
	return false
}

type CombinedValidationResult struct {
	// Transaction validation result
	Validation CombinedValidationResultValidation `json:"validation,required"`
	// Transaction simulation result
	Simulation SuccessfulSimulationResultSchema `json:"simulation,nullable"`
	JSON       combinedValidationResultJSON     `json:"-"`
}

// combinedValidationResultJSON contains the JSON metadata for the struct
// [CombinedValidationResult]
type combinedValidationResultJSON struct {
	Validation  apijson.Field
	Simulation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultJSON) RawJSON() string {
	return r.raw
}

// Transaction validation result
type CombinedValidationResultValidation struct {
	// A list of features about this transaction explaining the validation
	Features []string `json:"features,required"`
	// An enumeration.
	Reason CombinedValidationResultValidationReason `json:"reason,required"`
	// An enumeration.
	Verdict CombinedValidationResultValidationVerdict `json:"verdict,required"`
	JSON    combinedValidationResultValidationJSON    `json:"-"`
}

// combinedValidationResultValidationJSON contains the JSON metadata for the struct
// [CombinedValidationResultValidation]
type combinedValidationResultValidationJSON struct {
	Features    apijson.Field
	Reason      apijson.Field
	Verdict     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultValidationJSON) RawJSON() string {
	return r.raw
}

// An enumeration.
type CombinedValidationResultValidationReason string

const (
	CombinedValidationResultValidationReasonEmpty                   CombinedValidationResultValidationReason = ""
	CombinedValidationResultValidationReasonSharedStateInBulk       CombinedValidationResultValidationReason = "shared_state_in_bulk"
	CombinedValidationResultValidationReasonUnknownProfiter         CombinedValidationResultValidationReason = "unknown_profiter"
	CombinedValidationResultValidationReasonUnfairTrade             CombinedValidationResultValidationReason = "unfair_trade"
	CombinedValidationResultValidationReasonTransferFarming         CombinedValidationResultValidationReason = "transfer_farming"
	CombinedValidationResultValidationReasonWritableAccountsFarming CombinedValidationResultValidationReason = "writable_accounts_farming"
	CombinedValidationResultValidationReasonNativeOwnershipChange   CombinedValidationResultValidationReason = "native_ownership_change"
	CombinedValidationResultValidationReasonSplTokenOwnershipChange CombinedValidationResultValidationReason = "spl_token_ownership_change"
	CombinedValidationResultValidationReasonExposureFarming         CombinedValidationResultValidationReason = "exposure_farming"
	CombinedValidationResultValidationReasonKnownAttacker           CombinedValidationResultValidationReason = "known_attacker"
	CombinedValidationResultValidationReasonInvalidSignature        CombinedValidationResultValidationReason = "invalid_signature"
	CombinedValidationResultValidationReasonOther                   CombinedValidationResultValidationReason = "other"
)

func (r CombinedValidationResultValidationReason) IsKnown() bool {
	switch r {
	case CombinedValidationResultValidationReasonEmpty, CombinedValidationResultValidationReasonSharedStateInBulk, CombinedValidationResultValidationReasonUnknownProfiter, CombinedValidationResultValidationReasonUnfairTrade, CombinedValidationResultValidationReasonTransferFarming, CombinedValidationResultValidationReasonWritableAccountsFarming, CombinedValidationResultValidationReasonNativeOwnershipChange, CombinedValidationResultValidationReasonSplTokenOwnershipChange, CombinedValidationResultValidationReasonExposureFarming, CombinedValidationResultValidationReasonKnownAttacker, CombinedValidationResultValidationReasonInvalidSignature, CombinedValidationResultValidationReasonOther:
		return true
	}
	return false
}

// An enumeration.
type CombinedValidationResultValidationVerdict string

const (
	CombinedValidationResultValidationVerdictBenign    CombinedValidationResultValidationVerdict = "Benign"
	CombinedValidationResultValidationVerdictWarning   CombinedValidationResultValidationVerdict = "Warning"
	CombinedValidationResultValidationVerdictMalicious CombinedValidationResultValidationVerdict = "Malicious"
)

func (r CombinedValidationResultValidationVerdict) IsKnown() bool {
	switch r {
	case CombinedValidationResultValidationVerdictBenign, CombinedValidationResultValidationVerdictWarning, CombinedValidationResultValidationVerdictMalicious:
		return true
	}
	return false
}

type DelegatedAssetDetailsSchema struct {
	Asset DelegatedAssetDetailsSchemaAsset `json:"asset,required"`
	// The delegate's address
	Delegate string `json:"delegate,required"`
	// Details of the delegation
	Delegation DelegatedAssetDetailsSchemaDelegation `json:"delegation,required"`
	JSON       delegatedAssetDetailsSchemaJSON       `json:"-"`
}

// delegatedAssetDetailsSchemaJSON contains the JSON metadata for the struct
// [DelegatedAssetDetailsSchema]
type delegatedAssetDetailsSchemaJSON struct {
	Asset       apijson.Field
	Delegate    apijson.Field
	Delegation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DelegatedAssetDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r delegatedAssetDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

type DelegatedAssetDetailsSchemaAsset struct {
	Address string `json:"address,required"`
	Symbol  string `json:"symbol,required"`
	Name    string `json:"name,required"`
	Logo    string `json:"logo"`
	// Type of the asset (`"TOKEN"`)
	Type     string                               `json:"type"`
	Decimals int64                                `json:"decimals"`
	JSON     delegatedAssetDetailsSchemaAssetJSON `json:"-"`
	union    DelegatedAssetDetailsSchemaAssetUnion
}

// delegatedAssetDetailsSchemaAssetJSON contains the JSON metadata for the struct
// [DelegatedAssetDetailsSchemaAsset]
type delegatedAssetDetailsSchemaAssetJSON struct {
	Address     apijson.Field
	Symbol      apijson.Field
	Name        apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r delegatedAssetDetailsSchemaAssetJSON) RawJSON() string {
	return r.raw
}

func (r *DelegatedAssetDetailsSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DelegatedAssetDetailsSchemaAssetUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DelegatedAssetDetailsSchemaAssetSplFungibleTokenDetailsSchema],
// [DelegatedAssetDetailsSchemaAssetSplNonFungibleTokenDetailsSchema],
// [DelegatedAssetDetailsSchemaAssetCnftDetailsSchema].
func (r DelegatedAssetDetailsSchemaAsset) AsUnion() DelegatedAssetDetailsSchemaAssetUnion {
	return r.union
}

// Union satisfied by
// [DelegatedAssetDetailsSchemaAssetSplFungibleTokenDetailsSchema],
// [DelegatedAssetDetailsSchemaAssetSplNonFungibleTokenDetailsSchema] or
// [DelegatedAssetDetailsSchemaAssetCnftDetailsSchema].
type DelegatedAssetDetailsSchemaAssetUnion interface {
	implementsDelegatedAssetDetailsSchemaAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DelegatedAssetDetailsSchemaAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DelegatedAssetDetailsSchemaAssetSplFungibleTokenDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DelegatedAssetDetailsSchemaAssetSplNonFungibleTokenDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DelegatedAssetDetailsSchemaAssetCnftDetailsSchema{}),
		},
	)
}

type DelegatedAssetDetailsSchemaAssetSplFungibleTokenDetailsSchema struct {
	Address  string `json:"address,required"`
	Decimals int64  `json:"decimals,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Logo     string `json:"logo"`
	// Type of the asset (`"TOKEN"`)
	Type string                                                            `json:"type"`
	JSON delegatedAssetDetailsSchemaAssetSplFungibleTokenDetailsSchemaJSON `json:"-"`
}

// delegatedAssetDetailsSchemaAssetSplFungibleTokenDetailsSchemaJSON contains the
// JSON metadata for the struct
// [DelegatedAssetDetailsSchemaAssetSplFungibleTokenDetailsSchema]
type delegatedAssetDetailsSchemaAssetSplFungibleTokenDetailsSchemaJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DelegatedAssetDetailsSchemaAssetSplFungibleTokenDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r delegatedAssetDetailsSchemaAssetSplFungibleTokenDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r DelegatedAssetDetailsSchemaAssetSplFungibleTokenDetailsSchema) implementsDelegatedAssetDetailsSchemaAsset() {
}

type DelegatedAssetDetailsSchemaAssetSplNonFungibleTokenDetailsSchema struct {
	Address  string `json:"address,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	Logo     string `json:"logo"`
	// Type of the asset (`"NFT"`)
	Type string                                                               `json:"type"`
	JSON delegatedAssetDetailsSchemaAssetSplNonFungibleTokenDetailsSchemaJSON `json:"-"`
}

// delegatedAssetDetailsSchemaAssetSplNonFungibleTokenDetailsSchemaJSON contains
// the JSON metadata for the struct
// [DelegatedAssetDetailsSchemaAssetSplNonFungibleTokenDetailsSchema]
type delegatedAssetDetailsSchemaAssetSplNonFungibleTokenDetailsSchemaJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DelegatedAssetDetailsSchemaAssetSplNonFungibleTokenDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r delegatedAssetDetailsSchemaAssetSplNonFungibleTokenDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r DelegatedAssetDetailsSchemaAssetSplNonFungibleTokenDetailsSchema) implementsDelegatedAssetDetailsSchemaAsset() {
}

type DelegatedAssetDetailsSchemaAssetCnftDetailsSchema struct {
	Address  string `json:"address,required"`
	Decimals int64  `json:"decimals,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Logo     string `json:"logo"`
	// Type of the asset (`"CNFT"`)
	Type string                                                `json:"type"`
	JSON delegatedAssetDetailsSchemaAssetCnftDetailsSchemaJSON `json:"-"`
}

// delegatedAssetDetailsSchemaAssetCnftDetailsSchemaJSON contains the JSON metadata
// for the struct [DelegatedAssetDetailsSchemaAssetCnftDetailsSchema]
type delegatedAssetDetailsSchemaAssetCnftDetailsSchemaJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DelegatedAssetDetailsSchemaAssetCnftDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r delegatedAssetDetailsSchemaAssetCnftDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r DelegatedAssetDetailsSchemaAssetCnftDetailsSchema) implementsDelegatedAssetDetailsSchemaAsset() {
}

// Details of the delegation
type DelegatedAssetDetailsSchemaDelegation struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                   `json:"usd_price,nullable"`
	JSON     delegatedAssetDetailsSchemaDelegationJSON `json:"-"`
}

// delegatedAssetDetailsSchemaDelegationJSON contains the JSON metadata for the
// struct [DelegatedAssetDetailsSchemaDelegation]
type delegatedAssetDetailsSchemaDelegationJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DelegatedAssetDetailsSchemaDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r delegatedAssetDetailsSchemaDelegationJSON) RawJSON() string {
	return r.raw
}

type ResponseSchema struct {
	// Encoding of the public keys
	Encoding string `json:"encoding"`
	// Error message if the simulation failed
	Error string `json:"error,nullable"`
	// Summary of the result
	Result CombinedValidationResult `json:"result,nullable"`
	JSON   responseSchemaJSON       `json:"-"`
}

// responseSchemaJSON contains the JSON metadata for the struct [ResponseSchema]
type responseSchemaJSON struct {
	Encoding    apijson.Field
	Error       apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseSchemaJSON) RawJSON() string {
	return r.raw
}

type SuccessfulSimulationResultSchema struct {
	// Summary of the requested account address
	AccountSummary AccountSummarySchema `json:"account_summary,required"`
	// Summary of the accounts involved in the transaction simulation
	AccountsDetails []SuccessfulSimulationResultSchemaAccountsDetail `json:"accounts_details,required"`
	// Summary of the assets involved in the transaction simulation
	AssetsDiff map[string][]SuccessfulSimulationResultSchemaAssetsDiff `json:"assets_diff,required"`
	// Summary of ownership changes; By account address
	AssetsOwnershipDiff map[string][]SuccessfulSimulationResultSchemaAssetsOwnershipDiff `json:"assets_ownership_diff,required"`
	// Summary of the delegations, by account address
	Delegations map[string][]DelegatedAssetDetailsSchema `json:"delegations,required"`
	JSON        successfulSimulationResultSchemaJSON     `json:"-"`
}

// successfulSimulationResultSchemaJSON contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchema]
type successfulSimulationResultSchemaJSON struct {
	AccountSummary      apijson.Field
	AccountsDetails     apijson.Field
	AssetsDiff          apijson.Field
	AssetsOwnershipDiff apijson.Field
	Delegations         apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaJSON) RawJSON() string {
	return r.raw
}

type SuccessfulSimulationResultSchemaAccountsDetail struct {
	Type SuccessfulSimulationResultSchemaAccountsDetailsType `json:"type"`
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// Description of the account
	Description string `json:"description,nullable"`
	// Whether the account had been written to during the simulation
	WasWrittenTo bool `json:"was_written_to,required"`
	// Encoded public key of the mint
	MintAddress string `json:"mint_address"`
	// Encoded public key of the owner
	OwnerAddress string `json:"owner_address"`
	// Name of the mint
	Name string `json:"name"`
	// Symbol of the mint
	Symbol string `json:"symbol"`
	// Logo of the mint
	Logo string `json:"logo"`
	// URI of the mint
	Uri string `json:"uri"`
	// The address of the owning program
	Owner string                                             `json:"owner"`
	JSON  successfulSimulationResultSchemaAccountsDetailJSON `json:"-"`
	union SuccessfulSimulationResultSchemaAccountsDetailsUnion
}

// successfulSimulationResultSchemaAccountsDetailJSON contains the JSON metadata
// for the struct [SuccessfulSimulationResultSchemaAccountsDetail]
type successfulSimulationResultSchemaAccountsDetailJSON struct {
	Type           apijson.Field
	AccountAddress apijson.Field
	Description    apijson.Field
	WasWrittenTo   apijson.Field
	MintAddress    apijson.Field
	OwnerAddress   apijson.Field
	Name           apijson.Field
	Symbol         apijson.Field
	Logo           apijson.Field
	Uri            apijson.Field
	Owner          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r successfulSimulationResultSchemaAccountsDetailJSON) RawJSON() string {
	return r.raw
}

func (r *SuccessfulSimulationResultSchemaAccountsDetail) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SuccessfulSimulationResultSchemaAccountsDetailsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SuccessfulSimulationResultSchemaAccountsDetailsSystemAccountDetailsSchema],
// [SuccessfulSimulationResultSchemaAccountsDetailsTokenAccountDetailsSchema],
// [SuccessfulSimulationResultSchemaAccountsDetailsFungibleMintAccountDetailsSchema],
// [SuccessfulSimulationResultSchemaAccountsDetailsNonFungibleMintAccountDetailsSchema],
// [SuccessfulSimulationResultSchemaAccountsDetailsProgramAccountDetailsSchema],
// [SuccessfulSimulationResultSchemaAccountsDetailsPdaAccountSchema],
// [SuccessfulSimulationResultSchemaAccountsDetailsCnftMintAccountDetailsSchema].
func (r SuccessfulSimulationResultSchemaAccountsDetail) AsUnion() SuccessfulSimulationResultSchemaAccountsDetailsUnion {
	return r.union
}

// Union satisfied by
// [SuccessfulSimulationResultSchemaAccountsDetailsSystemAccountDetailsSchema],
// [SuccessfulSimulationResultSchemaAccountsDetailsTokenAccountDetailsSchema],
// [SuccessfulSimulationResultSchemaAccountsDetailsFungibleMintAccountDetailsSchema],
// [SuccessfulSimulationResultSchemaAccountsDetailsNonFungibleMintAccountDetailsSchema],
// [SuccessfulSimulationResultSchemaAccountsDetailsProgramAccountDetailsSchema],
// [SuccessfulSimulationResultSchemaAccountsDetailsPdaAccountSchema] or
// [SuccessfulSimulationResultSchemaAccountsDetailsCnftMintAccountDetailsSchema].
type SuccessfulSimulationResultSchemaAccountsDetailsUnion interface {
	implementsSuccessfulSimulationResultSchemaAccountsDetail()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SuccessfulSimulationResultSchemaAccountsDetailsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuccessfulSimulationResultSchemaAccountsDetailsSystemAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuccessfulSimulationResultSchemaAccountsDetailsTokenAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuccessfulSimulationResultSchemaAccountsDetailsFungibleMintAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuccessfulSimulationResultSchemaAccountsDetailsNonFungibleMintAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuccessfulSimulationResultSchemaAccountsDetailsProgramAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuccessfulSimulationResultSchemaAccountsDetailsPdaAccountSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuccessfulSimulationResultSchemaAccountsDetailsCnftMintAccountDetailsSchema{}),
		},
	)
}

type SuccessfulSimulationResultSchemaAccountsDetailsSystemAccountDetailsSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// Whether the account had been written to during the simulation
	WasWrittenTo bool `json:"was_written_to,required"`
	// Description of the account
	Description string                                                                        `json:"description,nullable"`
	Type        SuccessfulSimulationResultSchemaAccountsDetailsSystemAccountDetailsSchemaType `json:"type"`
	JSON        successfulSimulationResultSchemaAccountsDetailsSystemAccountDetailsSchemaJSON `json:"-"`
}

// successfulSimulationResultSchemaAccountsDetailsSystemAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAccountsDetailsSystemAccountDetailsSchema]
type successfulSimulationResultSchemaAccountsDetailsSystemAccountDetailsSchemaJSON struct {
	AccountAddress apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAccountsDetailsSystemAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAccountsDetailsSystemAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SuccessfulSimulationResultSchemaAccountsDetailsSystemAccountDetailsSchema) implementsSuccessfulSimulationResultSchemaAccountsDetail() {
}

type SuccessfulSimulationResultSchemaAccountsDetailsSystemAccountDetailsSchemaType string

const (
	SuccessfulSimulationResultSchemaAccountsDetailsSystemAccountDetailsSchemaTypeSystemAccount SuccessfulSimulationResultSchemaAccountsDetailsSystemAccountDetailsSchemaType = "SYSTEM_ACCOUNT"
)

func (r SuccessfulSimulationResultSchemaAccountsDetailsSystemAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case SuccessfulSimulationResultSchemaAccountsDetailsSystemAccountDetailsSchemaTypeSystemAccount:
		return true
	}
	return false
}

type SuccessfulSimulationResultSchemaAccountsDetailsTokenAccountDetailsSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// Encoded public key of the mint
	MintAddress string `json:"mint_address,required"`
	// Encoded public key of the owner
	OwnerAddress string `json:"owner_address,required"`
	// Whether the account had been written to during the simulation
	WasWrittenTo bool `json:"was_written_to,required"`
	// Description of the account
	Description string                                                                       `json:"description,nullable"`
	Type        SuccessfulSimulationResultSchemaAccountsDetailsTokenAccountDetailsSchemaType `json:"type"`
	JSON        successfulSimulationResultSchemaAccountsDetailsTokenAccountDetailsSchemaJSON `json:"-"`
}

// successfulSimulationResultSchemaAccountsDetailsTokenAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAccountsDetailsTokenAccountDetailsSchema]
type successfulSimulationResultSchemaAccountsDetailsTokenAccountDetailsSchemaJSON struct {
	AccountAddress apijson.Field
	MintAddress    apijson.Field
	OwnerAddress   apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAccountsDetailsTokenAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAccountsDetailsTokenAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SuccessfulSimulationResultSchemaAccountsDetailsTokenAccountDetailsSchema) implementsSuccessfulSimulationResultSchemaAccountsDetail() {
}

type SuccessfulSimulationResultSchemaAccountsDetailsTokenAccountDetailsSchemaType string

const (
	SuccessfulSimulationResultSchemaAccountsDetailsTokenAccountDetailsSchemaTypeTokenAccount SuccessfulSimulationResultSchemaAccountsDetailsTokenAccountDetailsSchemaType = "TOKEN_ACCOUNT"
)

func (r SuccessfulSimulationResultSchemaAccountsDetailsTokenAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case SuccessfulSimulationResultSchemaAccountsDetailsTokenAccountDetailsSchemaTypeTokenAccount:
		return true
	}
	return false
}

type SuccessfulSimulationResultSchemaAccountsDetailsFungibleMintAccountDetailsSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// Name of the mint
	Name string `json:"name,required"`
	// Symbol of the mint
	Symbol string `json:"symbol,required"`
	// Whether the account had been written to during the simulation
	WasWrittenTo bool `json:"was_written_to,required"`
	// Description of the account
	Description string `json:"description,nullable"`
	// Logo of the mint
	Logo string                                                                              `json:"logo"`
	Type SuccessfulSimulationResultSchemaAccountsDetailsFungibleMintAccountDetailsSchemaType `json:"type"`
	JSON successfulSimulationResultSchemaAccountsDetailsFungibleMintAccountDetailsSchemaJSON `json:"-"`
}

// successfulSimulationResultSchemaAccountsDetailsFungibleMintAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAccountsDetailsFungibleMintAccountDetailsSchema]
type successfulSimulationResultSchemaAccountsDetailsFungibleMintAccountDetailsSchemaJSON struct {
	AccountAddress apijson.Field
	Name           apijson.Field
	Symbol         apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Logo           apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAccountsDetailsFungibleMintAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAccountsDetailsFungibleMintAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SuccessfulSimulationResultSchemaAccountsDetailsFungibleMintAccountDetailsSchema) implementsSuccessfulSimulationResultSchemaAccountsDetail() {
}

type SuccessfulSimulationResultSchemaAccountsDetailsFungibleMintAccountDetailsSchemaType string

const (
	SuccessfulSimulationResultSchemaAccountsDetailsFungibleMintAccountDetailsSchemaTypeFungibleMintAccount SuccessfulSimulationResultSchemaAccountsDetailsFungibleMintAccountDetailsSchemaType = "FUNGIBLE_MINT_ACCOUNT"
)

func (r SuccessfulSimulationResultSchemaAccountsDetailsFungibleMintAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case SuccessfulSimulationResultSchemaAccountsDetailsFungibleMintAccountDetailsSchemaTypeFungibleMintAccount:
		return true
	}
	return false
}

type SuccessfulSimulationResultSchemaAccountsDetailsNonFungibleMintAccountDetailsSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// Name of the mint
	Name string `json:"name,required"`
	// Symbol of the mint
	Symbol string `json:"symbol,required"`
	// URI of the mint
	Uri string `json:"uri,required"`
	// Whether the account had been written to during the simulation
	WasWrittenTo bool `json:"was_written_to,required"`
	// Description of the account
	Description string `json:"description,nullable"`
	// Logo of the mint
	Logo string                                                                                 `json:"logo"`
	Type SuccessfulSimulationResultSchemaAccountsDetailsNonFungibleMintAccountDetailsSchemaType `json:"type"`
	JSON successfulSimulationResultSchemaAccountsDetailsNonFungibleMintAccountDetailsSchemaJSON `json:"-"`
}

// successfulSimulationResultSchemaAccountsDetailsNonFungibleMintAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAccountsDetailsNonFungibleMintAccountDetailsSchema]
type successfulSimulationResultSchemaAccountsDetailsNonFungibleMintAccountDetailsSchemaJSON struct {
	AccountAddress apijson.Field
	Name           apijson.Field
	Symbol         apijson.Field
	Uri            apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Logo           apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAccountsDetailsNonFungibleMintAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAccountsDetailsNonFungibleMintAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SuccessfulSimulationResultSchemaAccountsDetailsNonFungibleMintAccountDetailsSchema) implementsSuccessfulSimulationResultSchemaAccountsDetail() {
}

type SuccessfulSimulationResultSchemaAccountsDetailsNonFungibleMintAccountDetailsSchemaType string

const (
	SuccessfulSimulationResultSchemaAccountsDetailsNonFungibleMintAccountDetailsSchemaTypeNonFungibleMintAccount SuccessfulSimulationResultSchemaAccountsDetailsNonFungibleMintAccountDetailsSchemaType = "NON_FUNGIBLE_MINT_ACCOUNT"
)

func (r SuccessfulSimulationResultSchemaAccountsDetailsNonFungibleMintAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case SuccessfulSimulationResultSchemaAccountsDetailsNonFungibleMintAccountDetailsSchemaTypeNonFungibleMintAccount:
		return true
	}
	return false
}

type SuccessfulSimulationResultSchemaAccountsDetailsProgramAccountDetailsSchema struct {
	// Encoded public key of the account
	AccountAddress string                                                                         `json:"account_address,required"`
	Type           SuccessfulSimulationResultSchemaAccountsDetailsProgramAccountDetailsSchemaType `json:"type,required"`
	// Whether the account had been written to during the simulation
	WasWrittenTo bool `json:"was_written_to,required"`
	// Description of the account
	Description string                                                                         `json:"description,nullable"`
	JSON        successfulSimulationResultSchemaAccountsDetailsProgramAccountDetailsSchemaJSON `json:"-"`
}

// successfulSimulationResultSchemaAccountsDetailsProgramAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAccountsDetailsProgramAccountDetailsSchema]
type successfulSimulationResultSchemaAccountsDetailsProgramAccountDetailsSchemaJSON struct {
	AccountAddress apijson.Field
	Type           apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAccountsDetailsProgramAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAccountsDetailsProgramAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SuccessfulSimulationResultSchemaAccountsDetailsProgramAccountDetailsSchema) implementsSuccessfulSimulationResultSchemaAccountsDetail() {
}

type SuccessfulSimulationResultSchemaAccountsDetailsProgramAccountDetailsSchemaType string

const (
	SuccessfulSimulationResultSchemaAccountsDetailsProgramAccountDetailsSchemaTypeProgram       SuccessfulSimulationResultSchemaAccountsDetailsProgramAccountDetailsSchemaType = "PROGRAM"
	SuccessfulSimulationResultSchemaAccountsDetailsProgramAccountDetailsSchemaTypeNativeProgram SuccessfulSimulationResultSchemaAccountsDetailsProgramAccountDetailsSchemaType = "NATIVE_PROGRAM"
)

func (r SuccessfulSimulationResultSchemaAccountsDetailsProgramAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case SuccessfulSimulationResultSchemaAccountsDetailsProgramAccountDetailsSchemaTypeProgram, SuccessfulSimulationResultSchemaAccountsDetailsProgramAccountDetailsSchemaTypeNativeProgram:
		return true
	}
	return false
}

type SuccessfulSimulationResultSchemaAccountsDetailsPdaAccountSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// The address of the owning program
	Owner string `json:"owner,required"`
	// Whether the account had been written to during the simulation
	WasWrittenTo bool `json:"was_written_to,required"`
	// Description of the account
	Description string                                                              `json:"description,nullable"`
	Type        SuccessfulSimulationResultSchemaAccountsDetailsPdaAccountSchemaType `json:"type"`
	JSON        successfulSimulationResultSchemaAccountsDetailsPdaAccountSchemaJSON `json:"-"`
}

// successfulSimulationResultSchemaAccountsDetailsPdaAccountSchemaJSON contains the
// JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAccountsDetailsPdaAccountSchema]
type successfulSimulationResultSchemaAccountsDetailsPdaAccountSchemaJSON struct {
	AccountAddress apijson.Field
	Owner          apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAccountsDetailsPdaAccountSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAccountsDetailsPdaAccountSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SuccessfulSimulationResultSchemaAccountsDetailsPdaAccountSchema) implementsSuccessfulSimulationResultSchemaAccountsDetail() {
}

type SuccessfulSimulationResultSchemaAccountsDetailsPdaAccountSchemaType string

const (
	SuccessfulSimulationResultSchemaAccountsDetailsPdaAccountSchemaTypePda SuccessfulSimulationResultSchemaAccountsDetailsPdaAccountSchemaType = "PDA"
)

func (r SuccessfulSimulationResultSchemaAccountsDetailsPdaAccountSchemaType) IsKnown() bool {
	switch r {
	case SuccessfulSimulationResultSchemaAccountsDetailsPdaAccountSchemaTypePda:
		return true
	}
	return false
}

type SuccessfulSimulationResultSchemaAccountsDetailsCnftMintAccountDetailsSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// Name of the mint
	Name string `json:"name,required"`
	// Symbol of the mint
	Symbol string `json:"symbol,required"`
	// URI of the mint
	Uri string `json:"uri,required"`
	// Whether the account had been written to during the simulation
	WasWrittenTo bool `json:"was_written_to,required"`
	// Description of the account
	Description string `json:"description,nullable"`
	// Logo of the mint
	Logo string                                                                          `json:"logo"`
	Type SuccessfulSimulationResultSchemaAccountsDetailsCnftMintAccountDetailsSchemaType `json:"type"`
	JSON successfulSimulationResultSchemaAccountsDetailsCnftMintAccountDetailsSchemaJSON `json:"-"`
}

// successfulSimulationResultSchemaAccountsDetailsCnftMintAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAccountsDetailsCnftMintAccountDetailsSchema]
type successfulSimulationResultSchemaAccountsDetailsCnftMintAccountDetailsSchemaJSON struct {
	AccountAddress apijson.Field
	Name           apijson.Field
	Symbol         apijson.Field
	Uri            apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Logo           apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAccountsDetailsCnftMintAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAccountsDetailsCnftMintAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SuccessfulSimulationResultSchemaAccountsDetailsCnftMintAccountDetailsSchema) implementsSuccessfulSimulationResultSchemaAccountsDetail() {
}

type SuccessfulSimulationResultSchemaAccountsDetailsCnftMintAccountDetailsSchemaType string

const (
	SuccessfulSimulationResultSchemaAccountsDetailsCnftMintAccountDetailsSchemaTypeCnftMintAccount SuccessfulSimulationResultSchemaAccountsDetailsCnftMintAccountDetailsSchemaType = "CNFT_MINT_ACCOUNT"
)

func (r SuccessfulSimulationResultSchemaAccountsDetailsCnftMintAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case SuccessfulSimulationResultSchemaAccountsDetailsCnftMintAccountDetailsSchemaTypeCnftMintAccount:
		return true
	}
	return false
}

type SuccessfulSimulationResultSchemaAccountsDetailsType string

const (
	SuccessfulSimulationResultSchemaAccountsDetailsTypeSystemAccount          SuccessfulSimulationResultSchemaAccountsDetailsType = "SYSTEM_ACCOUNT"
	SuccessfulSimulationResultSchemaAccountsDetailsTypeTokenAccount           SuccessfulSimulationResultSchemaAccountsDetailsType = "TOKEN_ACCOUNT"
	SuccessfulSimulationResultSchemaAccountsDetailsTypeFungibleMintAccount    SuccessfulSimulationResultSchemaAccountsDetailsType = "FUNGIBLE_MINT_ACCOUNT"
	SuccessfulSimulationResultSchemaAccountsDetailsTypeNonFungibleMintAccount SuccessfulSimulationResultSchemaAccountsDetailsType = "NON_FUNGIBLE_MINT_ACCOUNT"
	SuccessfulSimulationResultSchemaAccountsDetailsTypeProgram                SuccessfulSimulationResultSchemaAccountsDetailsType = "PROGRAM"
	SuccessfulSimulationResultSchemaAccountsDetailsTypeNativeProgram          SuccessfulSimulationResultSchemaAccountsDetailsType = "NATIVE_PROGRAM"
	SuccessfulSimulationResultSchemaAccountsDetailsTypePda                    SuccessfulSimulationResultSchemaAccountsDetailsType = "PDA"
	SuccessfulSimulationResultSchemaAccountsDetailsTypeCnftMintAccount        SuccessfulSimulationResultSchemaAccountsDetailsType = "CNFT_MINT_ACCOUNT"
)

func (r SuccessfulSimulationResultSchemaAccountsDetailsType) IsKnown() bool {
	switch r {
	case SuccessfulSimulationResultSchemaAccountsDetailsTypeSystemAccount, SuccessfulSimulationResultSchemaAccountsDetailsTypeTokenAccount, SuccessfulSimulationResultSchemaAccountsDetailsTypeFungibleMintAccount, SuccessfulSimulationResultSchemaAccountsDetailsTypeNonFungibleMintAccount, SuccessfulSimulationResultSchemaAccountsDetailsTypeProgram, SuccessfulSimulationResultSchemaAccountsDetailsTypeNativeProgram, SuccessfulSimulationResultSchemaAccountsDetailsTypePda, SuccessfulSimulationResultSchemaAccountsDetailsTypeCnftMintAccount:
		return true
	}
	return false
}

type SuccessfulSimulationResultSchemaAssetsDiff struct {
	// This field can have the runtime type of
	// [SuccessfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaAsset],
	// [SuccessfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaAsset],
	// [SuccessfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaAsset],
	// [SuccessfulSimulationResultSchemaAssetsDiffCnftDiffSchemaAsset].
	Asset interface{} `json:"asset"`
	// This field can have the runtime type of
	// [SuccessfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaIn],
	// [SuccessfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaIn],
	// [SuccessfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaIn],
	// [SuccessfulSimulationResultSchemaAssetsDiffCnftDiffSchemaIn].
	In interface{} `json:"in,required"`
	// This field can have the runtime type of
	// [SuccessfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaOut],
	// [SuccessfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaOut],
	// [SuccessfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaOut],
	// [SuccessfulSimulationResultSchemaAssetsDiffCnftDiffSchemaOut].
	Out   interface{}                                    `json:"out,required"`
	JSON  successfulSimulationResultSchemaAssetsDiffJSON `json:"-"`
	union SuccessfulSimulationResultSchemaAssetsDiffUnion
}

// successfulSimulationResultSchemaAssetsDiffJSON contains the JSON metadata for
// the struct [SuccessfulSimulationResultSchemaAssetsDiff]
type successfulSimulationResultSchemaAssetsDiffJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r successfulSimulationResultSchemaAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *SuccessfulSimulationResultSchemaAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SuccessfulSimulationResultSchemaAssetsDiffUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SuccessfulSimulationResultSchemaAssetsDiffNativeSolDiffSchema],
// [SuccessfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchema],
// [SuccessfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchema],
// [SuccessfulSimulationResultSchemaAssetsDiffCnftDiffSchema].
func (r SuccessfulSimulationResultSchemaAssetsDiff) AsUnion() SuccessfulSimulationResultSchemaAssetsDiffUnion {
	return r.union
}

// Union satisfied by
// [SuccessfulSimulationResultSchemaAssetsDiffNativeSolDiffSchema],
// [SuccessfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchema],
// [SuccessfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchema] or
// [SuccessfulSimulationResultSchemaAssetsDiffCnftDiffSchema].
type SuccessfulSimulationResultSchemaAssetsDiffUnion interface {
	implementsSuccessfulSimulationResultSchemaAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SuccessfulSimulationResultSchemaAssetsDiffUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuccessfulSimulationResultSchemaAssetsDiffNativeSolDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuccessfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuccessfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuccessfulSimulationResultSchemaAssetsDiffCnftDiffSchema{}),
		},
	)
}

type SuccessfulSimulationResultSchemaAssetsDiffNativeSolDiffSchema struct {
	Asset SuccessfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaAsset `json:"asset,required"`
	// Incoming transfers of the asset
	In   SuccessfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaIn   `json:"in,nullable"`
	Out  SuccessfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaOut  `json:"out,nullable"`
	JSON successfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaJSON contains the
// JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsDiffNativeSolDiffSchema]
type successfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsDiffNativeSolDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SuccessfulSimulationResultSchemaAssetsDiffNativeSolDiffSchema) implementsSuccessfulSimulationResultSchemaAssetsDiff() {
}

type SuccessfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaAsset struct {
	Decimals int64 `json:"decimals"`
	// Type of the asset (`"SOL"`)
	Type string                                                                 `json:"type"`
	JSON successfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaAssetJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaAssetJSON contains
// the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaAsset]
type successfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Incoming transfers of the asset
type SuccessfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                             `json:"usd_price,nullable"`
	JSON     successfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaInJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaInJSON contains the
// JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaIn]
type successfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

type SuccessfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                              `json:"usd_price,nullable"`
	JSON     successfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaOutJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaOutJSON contains
// the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaOut]
type successfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsDiffNativeSolDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type SuccessfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchema struct {
	Asset SuccessfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaAsset `json:"asset,required"`
	// Incoming transfers of the asset
	In   SuccessfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaIn   `json:"in,nullable"`
	Out  SuccessfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaOut  `json:"out,nullable"`
	JSON successfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchema]
type successfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SuccessfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchema) implementsSuccessfulSimulationResultSchemaAssetsDiff() {
}

type SuccessfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaAsset struct {
	Address  string `json:"address,required"`
	Decimals int64  `json:"decimals,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Logo     string `json:"logo"`
	// Type of the asset (`"TOKEN"`)
	Type string                                                                        `json:"type"`
	JSON successfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaAssetJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaAsset]
type successfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Incoming transfers of the asset
type SuccessfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                    `json:"usd_price,nullable"`
	JSON     successfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaInJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaInJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaIn]
type successfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

type SuccessfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                     `json:"usd_price,nullable"`
	JSON     successfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaOutJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaOutJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaOut]
type successfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsDiffSplFungibleTokenDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type SuccessfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchema struct {
	Asset SuccessfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaAsset `json:"asset,required"`
	// Incoming transfers of the asset
	In   SuccessfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaIn   `json:"in,nullable"`
	Out  SuccessfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaOut  `json:"out,nullable"`
	JSON successfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchema]
type successfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SuccessfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchema) implementsSuccessfulSimulationResultSchemaAssetsDiff() {
}

type SuccessfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaAsset struct {
	Address  string `json:"address,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	Logo     string `json:"logo"`
	// Type of the asset (`"NFT"`)
	Type string                                                                           `json:"type"`
	JSON successfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaAssetJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaAsset]
type successfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Incoming transfers of the asset
type SuccessfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                       `json:"usd_price,nullable"`
	JSON     successfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaInJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaInJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaIn]
type successfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

type SuccessfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                        `json:"usd_price,nullable"`
	JSON     successfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaOutJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaOutJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaOut]
type successfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsDiffSplNonFungibleTokenDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type SuccessfulSimulationResultSchemaAssetsDiffCnftDiffSchema struct {
	Asset SuccessfulSimulationResultSchemaAssetsDiffCnftDiffSchemaAsset `json:"asset,required"`
	// Incoming transfers of the asset
	In   SuccessfulSimulationResultSchemaAssetsDiffCnftDiffSchemaIn   `json:"in,nullable"`
	Out  SuccessfulSimulationResultSchemaAssetsDiffCnftDiffSchemaOut  `json:"out,nullable"`
	JSON successfulSimulationResultSchemaAssetsDiffCnftDiffSchemaJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsDiffCnftDiffSchemaJSON contains the JSON
// metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsDiffCnftDiffSchema]
type successfulSimulationResultSchemaAssetsDiffCnftDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsDiffCnftDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsDiffCnftDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SuccessfulSimulationResultSchemaAssetsDiffCnftDiffSchema) implementsSuccessfulSimulationResultSchemaAssetsDiff() {
}

type SuccessfulSimulationResultSchemaAssetsDiffCnftDiffSchemaAsset struct {
	Address  string `json:"address,required"`
	Decimals int64  `json:"decimals,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Logo     string `json:"logo"`
	// Type of the asset (`"CNFT"`)
	Type string                                                            `json:"type"`
	JSON successfulSimulationResultSchemaAssetsDiffCnftDiffSchemaAssetJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsDiffCnftDiffSchemaAssetJSON contains the
// JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsDiffCnftDiffSchemaAsset]
type successfulSimulationResultSchemaAssetsDiffCnftDiffSchemaAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsDiffCnftDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsDiffCnftDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Incoming transfers of the asset
type SuccessfulSimulationResultSchemaAssetsDiffCnftDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                        `json:"usd_price,nullable"`
	JSON     successfulSimulationResultSchemaAssetsDiffCnftDiffSchemaInJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsDiffCnftDiffSchemaInJSON contains the JSON
// metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsDiffCnftDiffSchemaIn]
type successfulSimulationResultSchemaAssetsDiffCnftDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsDiffCnftDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsDiffCnftDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

type SuccessfulSimulationResultSchemaAssetsDiffCnftDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                         `json:"usd_price,nullable"`
	JSON     successfulSimulationResultSchemaAssetsDiffCnftDiffSchemaOutJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsDiffCnftDiffSchemaOutJSON contains the
// JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsDiffCnftDiffSchemaOut]
type successfulSimulationResultSchemaAssetsDiffCnftDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsDiffCnftDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsDiffCnftDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type SuccessfulSimulationResultSchemaAssetsOwnershipDiff struct {
	// This field can have the runtime type of
	// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaAsset],
	// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset],
	// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAsset].
	Asset interface{} `json:"asset,required"`
	// This field can have the runtime type of
	// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaIn],
	// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaIn],
	// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaIn].
	In interface{} `json:"in_,required"`
	// This field can have the runtime type of
	// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaOut],
	// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaOut],
	// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaOut].
	Out interface{} `json:"out,required"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,nullable"`
	// The owner post the transaction
	PostOwner string                                                  `json:"post_owner,required"`
	JSON      successfulSimulationResultSchemaAssetsOwnershipDiffJSON `json:"-"`
	union     SuccessfulSimulationResultSchemaAssetsOwnershipDiffUnion
}

// successfulSimulationResultSchemaAssetsOwnershipDiffJSON contains the JSON
// metadata for the struct [SuccessfulSimulationResultSchemaAssetsOwnershipDiff]
type successfulSimulationResultSchemaAssetsOwnershipDiffJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	PostOwner   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r successfulSimulationResultSchemaAssetsOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

func (r *SuccessfulSimulationResultSchemaAssetsOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SuccessfulSimulationResultSchemaAssetsOwnershipDiffUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchema],
// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchema],
// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchema].
func (r SuccessfulSimulationResultSchemaAssetsOwnershipDiff) AsUnion() SuccessfulSimulationResultSchemaAssetsOwnershipDiffUnion {
	return r.union
}

// Union satisfied by
// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchema],
// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchema]
// or
// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchema].
type SuccessfulSimulationResultSchemaAssetsOwnershipDiffUnion interface {
	implementsSuccessfulSimulationResultSchemaAssetsOwnershipDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SuccessfulSimulationResultSchemaAssetsOwnershipDiffUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuccessfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuccessfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchema{}),
		},
	)
}

type SuccessfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchema struct {
	Asset SuccessfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaAsset `json:"asset,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required"`
	// Incoming transfers of the asset
	In SuccessfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaIn `json:"in_,nullable"`
	// Details of the moved value
	Out SuccessfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaOut `json:"out,nullable"`
	// The owner prior to the transaction
	PreOwner string                                                                              `json:"pre_owner,nullable"`
	JSON     successfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchema]
type successfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaJSON struct {
	Asset       apijson.Field
	PostOwner   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SuccessfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchema) implementsSuccessfulSimulationResultSchemaAssetsOwnershipDiff() {
}

type SuccessfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaAsset struct {
	Decimals int64 `json:"decimals"`
	// Type of the asset (`"SOL"`)
	Type string                                                                                   `json:"type"`
	JSON successfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaAssetJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaAsset]
type successfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Incoming transfers of the asset
type SuccessfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                               `json:"usd_price,nullable"`
	JSON     successfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaInJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaInJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaIn]
type successfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

// Details of the moved value
type SuccessfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                `json:"usd_price,nullable"`
	JSON     successfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaOutJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaOutJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaOut]
type successfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsOwnershipDiffNativeSolOwnershipDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchema struct {
	Asset SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset `json:"asset,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required"`
	// Incoming transfers of the asset
	In SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaIn `json:"in_,nullable"`
	// Details of the moved value
	Out SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaOut `json:"out,nullable"`
	// The owner prior to the transaction
	PreOwner string                                                                             `json:"pre_owner,nullable"`
	JSON     successfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchema]
type successfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaJSON struct {
	Asset       apijson.Field
	PostOwner   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchema) implementsSuccessfulSimulationResultSchemaAssetsOwnershipDiff() {
}

type SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset struct {
	Address string `json:"address,required"`
	Symbol  string `json:"symbol,required"`
	Name    string `json:"name,required"`
	Logo    string `json:"logo"`
	// Type of the asset (`"TOKEN"`)
	Type     string                                                                                  `json:"type"`
	Decimals int64                                                                                   `json:"decimals"`
	JSON     successfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetJSON `json:"-"`
	union    SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetUnion
}

// successfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset]
type successfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetJSON struct {
	Address     apijson.Field
	Symbol      apijson.Field
	Name        apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r successfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

func (r *SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema],
// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema].
func (r SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset) AsUnion() SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetUnion {
	return r.union
}

// Union satisfied by
// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema]
// or
// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema].
type SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetUnion interface {
	implementsSuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema{}),
		},
	)
}

type SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema struct {
	Address  string `json:"address,required"`
	Decimals int64  `json:"decimals,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Logo     string `json:"logo"`
	// Type of the asset (`"TOKEN"`)
	Type string                                                                                                               `json:"type"`
	JSON successfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchemaJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema]
type successfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchemaJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema) implementsSuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset() {
}

type SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema struct {
	Address  string `json:"address,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	Logo     string `json:"logo"`
	// Type of the asset (`"NFT"`)
	Type string                                                                                                                  `json:"type"`
	JSON successfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchemaJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema]
type successfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchemaJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema) implementsSuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset() {
}

// Incoming transfers of the asset
type SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                              `json:"usd_price,nullable"`
	JSON     successfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaInJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaInJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaIn]
type successfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

// Details of the moved value
type SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                               `json:"usd_price,nullable"`
	JSON     successfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaOutJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaOutJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaOut]
type successfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsOwnershipDiffSplTokenOwnershipDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type SuccessfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchema struct {
	// The owner post the transaction
	PostOwner string                                                                                       `json:"post_owner,required"`
	Asset     SuccessfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAsset `json:"asset"`
	// Incoming transfers of the asset
	In SuccessfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaIn `json:"in_,nullable"`
	// Details of the moved value
	Out SuccessfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaOut `json:"out,nullable"`
	// The owner prior to the transaction
	PreOwner string                                                                                      `json:"pre_owner,nullable"`
	JSON     successfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchema]
type successfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaJSON struct {
	PostOwner   apijson.Field
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SuccessfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchema) implementsSuccessfulSimulationResultSchemaAssetsOwnershipDiff() {
}

type SuccessfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAsset struct {
	Decimals int64 `json:"decimals"`
	// Type of the asset (`"STAKED_SOL"`)
	Type string                                                                                           `json:"type"`
	JSON successfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAssetJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAsset]
type successfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Incoming transfers of the asset
type SuccessfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                       `json:"usd_price,nullable"`
	JSON     successfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaInJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaInJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaIn]
type successfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

// Details of the moved value
type SuccessfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                        `json:"usd_price,nullable"`
	JSON     successfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaOutJSON `json:"-"`
}

// successfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaOutJSON
// contains the JSON metadata for the struct
// [SuccessfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaOut]
type successfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuccessfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r successfulSimulationResultSchemaAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type TxScanRequestSchemaParam struct {
	// Encoded public key of the account to simulate the transaction on
	AccountAddress param.Field[string]                           `json:"account_address,required"`
	Metadata       param.Field[TxScanRequestSchemaMetadataParam] `json:"metadata,required"`
	// Transactions to scan
	Transactions param.Field[[]string] `json:"transactions,required"`
	// Chain to scan the transaction on
	Chain param.Field[string] `json:"chain"`
	// Encoding of the transaction and public keys
	Encoding param.Field[string] `json:"encoding"`
	// The RPC method used by dApp to propose the transaction
	Method param.Field[string] `json:"method"`
}

func (r TxScanRequestSchemaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TxScanRequestSchemaMetadataParam struct {
	// URL of the dApp that originated the transaction
	URL param.Field[string] `json:"url"`
}

func (r TxScanRequestSchemaMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
