// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"net/http"
	"reflect"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/param"
	"github.com/blockaid-official/blockaid-client-go/internal/requestconfig"
	"github.com/blockaid-official/blockaid-client-go/option"
	"github.com/tidwall/gjson"
)

// BitcoinTransactionService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBitcoinTransactionService] method instead.
type BitcoinTransactionService struct {
	Options []option.RequestOption
}

// NewBitcoinTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBitcoinTransactionService(opts ...option.RequestOption) (r *BitcoinTransactionService) {
	r = &BitcoinTransactionService{}
	r.Options = opts
	return
}

// Report Transaction
func (r *BitcoinTransactionService) Report(ctx context.Context, body BitcoinTransactionReportParams, opts ...option.RequestOption) (res *int64, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/bitcoin/transaction/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Scan Transaction
func (r *BitcoinTransactionService) Scan(ctx context.Context, body BitcoinTransactionScanParams, opts ...option.RequestOption) (res *BitcoinTransactionScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/bitcoin/transaction/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
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
	Status         BitcoinTransactionScanResponseSimulationStatus `json:"status,required"`
	AccountSummary interface{}                                    `json:"account_summary"`
	// This field can have the runtime type of
	// [[]BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAddressDetail].
	AddressDetails interface{} `json:"address_details"`
	// This field can have the runtime type of
	// [map[string][]BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiff].
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
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponse],
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchema].
func (r BitcoinTransactionScanResponseSimulation) AsUnion() BitcoinTransactionScanResponseSimulationUnion {
	return r.union
}

// Simulation result; Only present if simulation option is included in the request
//
// Union satisfied by
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponse] or
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
			Type:       reflect.TypeOf(BitcoinTransactionScanResponseSimulationBitcoinSimulationResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchema{}),
		},
	)
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponse struct {
	Status         BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseStatus `json:"status,required"`
	AccountSummary interface{}                                                             `json:"account_summary"`
	// Details of addresses involved in the transaction
	AddressDetails []BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAddressDetail `json:"address_details"`
	// Mapping between the address of an account to the assets diff during the
	// transaction
	AssetsDiffs map[string][]BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiff `json:"assets_diffs"`
	JSON        bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseJSON                    `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseJSON contains
// the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponse]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseJSON struct {
	Status         apijson.Field
	AccountSummary apijson.Field
	AddressDetails apijson.Field
	AssetsDiffs    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseJSON) RawJSON() string {
	return r.raw
}

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResponse) implementsBitcoinTransactionScanResponseSimulation() {
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseStatus string

const (
	BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseStatusSuccess BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseStatus = "Success"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseStatus) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseStatusSuccess:
		return true
	}
	return false
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAddressDetail struct {
	// Encoded public key of the account
	AccountAddress interface{} `json:"account_address,required"`
	// Description of the account
	Description string                                                                             `json:"description,nullable"`
	JSON        bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAddressDetailJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAddressDetailJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAddressDetail]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAddressDetailJSON struct {
	AccountAddress apijson.Field
	Description    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAddressDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAddressDetailJSON) RawJSON() string {
	return r.raw
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiff struct {
	// This field can have the runtime type of
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAsset],
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffAsset],
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffAsset].
	Asset interface{} `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffIn],
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffIn],
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffIn].
	In interface{} `json:"in"`
	// This field can have the runtime type of
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffOut],
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffOut],
	// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffOut].
	Out   interface{}                                                                     `json:"out"`
	JSON  bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffJSON `json:"-"`
	union BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsUnion
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiff]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiff],
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiff],
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiff].
func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiff) AsUnion() BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiff],
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiff]
// or
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiff].
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsUnion interface {
	implementsBitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiff{}),
		},
	)
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiff struct {
	Asset BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffOut  `json:"out,nullable"`
	JSON bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiff]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiff) implementsBitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiff() {
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAsset struct {
	// URL of the asset's logo
	LogoURL string `json:"logo_url,required,nullable"`
	// Decimals of the asset
	Decimals BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetDecimals `json:"decimals"`
	// Name of the asset
	Name BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetName `json:"name"`
	// Symbol of the asset
	Symbol BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetSymbol `json:"symbol"`
	// Type of the asset (`NATIVE`)
	Type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetType `json:"type"`
	JSON bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAsset]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetJSON struct {
	LogoURL     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Decimals of the asset
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetDecimals int64

const (
	BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetDecimals8 BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetDecimals = 8
)

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetDecimals) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetDecimals8:
		return true
	}
	return false
}

// Name of the asset
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetName string

const (
	BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetNameBitcoin BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetName = "Bitcoin"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetName) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetNameBitcoin:
		return true
	}
	return false
}

// Symbol of the asset
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetSymbol string

const (
	BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetSymbolBtc BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetSymbol = "BTC"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetSymbol) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetSymbolBtc:
		return true
	}
	return false
}

// Type of the asset (`NATIVE`)
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetType string

const (
	BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetTypeNative BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetType = "NATIVE"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffAssetTypeNative:
		return true
	}
	return false
}

// Details of the incoming transfer
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffIn struct {
	// Raw value of the transfer
	RawValue string `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                   `json:"summary,nullable"`
	JSON    bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffInJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffInJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffIn]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffInJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffOut struct {
	// Raw value of the transfer
	RawValue string `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                    `json:"summary,nullable"`
	JSON    bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffOutJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffOutJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffOut]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffOutJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinNativeAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiff struct {
	Asset BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffOut  `json:"out,nullable"`
	JSON bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiff]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiff) implementsBitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiff() {
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffAsset struct {
	// token's name
	Name string `json:"name,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ORDINAL`)
	Type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffAssetType `json:"type"`
	JSON bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffAssetJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffAsset]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffAssetJSON struct {
	Name        apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ORDINAL`)
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffAssetType string

const (
	BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffAssetTypeOrdinal BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffAssetType = "ORDINAL"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffAssetType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffAssetTypeOrdinal:
		return true
	}
	return false
}

// Details of the incoming transfer
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffIn struct {
	// Id of the ordinal
	ID string `json:"id,required"`
	// Raw value of the transfer
	RawValue string `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                    `json:"summary,nullable"`
	JSON    bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffInJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffInJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffIn]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffInJSON struct {
	ID          apijson.Field
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffOut struct {
	// Id of the ordinal
	ID string `json:"id,required"`
	// Raw value of the transfer
	RawValue string `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                     `json:"summary,nullable"`
	JSON    bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffOutJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffOutJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffOut]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffOutJSON struct {
	ID          apijson.Field
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinOrdinalAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiff struct {
	Asset BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffOut  `json:"out,nullable"`
	JSON bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiff]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiff) implementsBitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiff() {
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffAsset struct {
	// The Rune ID
	ID string `json:"id,required"`
	// Decimals of the asset
	Decimals int64 `json:"decimals,required"`
	// The Rune name
	Name string `json:"name,required"`
	// The Rune spaced name
	SpacedName string `json:"spaced_name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`RUNE`)
	Type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffAssetType `json:"type"`
	JSON bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffAssetJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffAsset]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffAssetJSON struct {
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

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`RUNE`)
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffAssetType string

const (
	BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffAssetTypeRune BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffAssetType = "RUNE"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffAssetType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffAssetTypeRune:
		return true
	}
	return false
}

// Details of the incoming transfer
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffIn struct {
	// Raw value of the transfer
	RawValue string `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                  `json:"summary,nullable"`
	JSON    bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffInJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffInJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffIn]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffInJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffOut struct {
	// Raw value of the transfer
	RawValue string `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                   `json:"summary,nullable"`
	JSON    bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffOutJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffOutJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffOut]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffOutJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationResponseAssetsDiffsBitcoinRunesAssetDiffOutJSON) RawJSON() string {
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
	FeatureID string `json:"feature_id,required"`
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

type BitcoinTransactionReportParams struct {
	Details param.Field[string]                                    `json:"details,required"`
	Event   param.Field[BitcoinTransactionReportParamsEvent]       `json:"event,required"`
	Report  param.Field[BitcoinTransactionReportParamsReportUnion] `json:"report,required"`
}

func (r BitcoinTransactionReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BitcoinTransactionReportParamsEvent string

const (
	BitcoinTransactionReportParamsEventShouldBeMalicious BitcoinTransactionReportParamsEvent = "should_be_malicious"
	BitcoinTransactionReportParamsEventShouldBeBenign    BitcoinTransactionReportParamsEvent = "should_be_benign"
)

func (r BitcoinTransactionReportParamsEvent) IsKnown() bool {
	switch r {
	case BitcoinTransactionReportParamsEventShouldBeMalicious, BitcoinTransactionReportParamsEventShouldBeBenign:
		return true
	}
	return false
}

type BitcoinTransactionReportParamsReport struct {
	ID     param.Field[string]                                   `json:"id"`
	Params param.Field[interface{}]                              `json:"params"`
	Type   param.Field[BitcoinTransactionReportParamsReportType] `json:"type"`
}

func (r BitcoinTransactionReportParamsReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionReportParamsReport) implementsBitcoinTransactionReportParamsReportUnion() {}

// Satisfied by [BitcoinTransactionReportParamsReportBitcoinAppealRequestID],
// [BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReport],
// [BitcoinTransactionReportParamsReport].
type BitcoinTransactionReportParamsReportUnion interface {
	implementsBitcoinTransactionReportParamsReportUnion()
}

type BitcoinTransactionReportParamsReportBitcoinAppealRequestID struct {
	ID   param.Field[string]                                                         `json:"id,required"`
	Type param.Field[BitcoinTransactionReportParamsReportBitcoinAppealRequestIDType] `json:"type"`
}

func (r BitcoinTransactionReportParamsReportBitcoinAppealRequestID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionReportParamsReportBitcoinAppealRequestID) implementsBitcoinTransactionReportParamsReportUnion() {
}

type BitcoinTransactionReportParamsReportBitcoinAppealRequestIDType string

const (
	BitcoinTransactionReportParamsReportBitcoinAppealRequestIDTypeRequestID BitcoinTransactionReportParamsReportBitcoinAppealRequestIDType = "request_id"
)

func (r BitcoinTransactionReportParamsReportBitcoinAppealRequestIDType) IsKnown() bool {
	switch r {
	case BitcoinTransactionReportParamsReportBitcoinAppealRequestIDTypeRequestID:
		return true
	}
	return false
}

type BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReport struct {
	Params param.Field[BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParams] `json:"params,required"`
	Type   param.Field[BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportType]   `json:"type"`
}

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReport) implementsBitcoinTransactionReportParamsReportUnion() {
}

type BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParams struct {
	AccountAddress param.Field[string]                                                                            `json:"account_address,required"`
	Chain          param.Field[BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsChain] `json:"chain,required"`
	// Metadata
	Metadata    param.Field[BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataUnion] `json:"metadata,required"`
	Transaction param.Field[string]                                                                                    `json:"transaction,required"`
	// List of options to include in the response
	//
	// - `Options.validation`: Include Options.validation output in the response
	//
	// - `Options.simulation`: Include Options.simulation output in the response
	Options param.Field[[]BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsOption] `json:"options"`
}

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsChain string

const (
	BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsChainBitcoin BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsChain = "bitcoin"
)

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsChain) IsKnown() bool {
	switch r {
	case BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsChainBitcoin:
		return true
	}
	return false
}

// Metadata
type BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadata struct {
	// Metadata for wallet requests
	Type param.Field[BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataType] `json:"type"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url"`
}

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadata) implementsBitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataUnion() {
}

// Metadata
//
// Satisfied by
// [BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataBitcoinWalletRequestMetadata],
// [BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataBitcoinInAppRequestMetadata],
// [BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadata].
type BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataUnion interface {
	implementsBitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataUnion()
}

type BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataBitcoinWalletRequestMetadata struct {
	// Metadata for wallet requests
	Type param.Field[BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataBitcoinWalletRequestMetadataType] `json:"type,required"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url,required"`
}

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataBitcoinWalletRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataBitcoinWalletRequestMetadata) implementsBitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataUnion() {
}

// Metadata for wallet requests
type BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataBitcoinWalletRequestMetadataType string

const (
	BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataBitcoinWalletRequestMetadataTypeWallet BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataBitcoinWalletRequestMetadataType = "wallet"
)

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataBitcoinWalletRequestMetadataType) IsKnown() bool {
	switch r {
	case BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataBitcoinWalletRequestMetadataTypeWallet:
		return true
	}
	return false
}

type BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataBitcoinInAppRequestMetadata struct {
	// Metadata for in-app requests
	Type param.Field[BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataBitcoinInAppRequestMetadataType] `json:"type"`
}

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataBitcoinInAppRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataBitcoinInAppRequestMetadata) implementsBitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataUnion() {
}

// Metadata for in-app requests
type BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataBitcoinInAppRequestMetadataType string

const (
	BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataBitcoinInAppRequestMetadataTypeInApp BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataBitcoinInAppRequestMetadataType = "in_app"
)

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataBitcoinInAppRequestMetadataType) IsKnown() bool {
	switch r {
	case BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataBitcoinInAppRequestMetadataTypeInApp:
		return true
	}
	return false
}

// Metadata for wallet requests
type BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataType string

const (
	BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataTypeWallet BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataType = "wallet"
	BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataTypeInApp  BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataType = "in_app"
)

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataType) IsKnown() bool {
	switch r {
	case BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataTypeWallet, BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsMetadataTypeInApp:
		return true
	}
	return false
}

type BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsOption string

const (
	BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsOptionValidation BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsOption = "validation"
	BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsOptionSimulation BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsOption = "simulation"
)

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsOption) IsKnown() bool {
	switch r {
	case BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsOptionValidation, BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportParamsOptionSimulation:
		return true
	}
	return false
}

type BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportType string

const (
	BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportTypeParams BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportType = "params"
)

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportType) IsKnown() bool {
	switch r {
	case BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportTypeParams:
		return true
	}
	return false
}

type BitcoinTransactionReportParamsReportType string

const (
	BitcoinTransactionReportParamsReportTypeRequestID BitcoinTransactionReportParamsReportType = "request_id"
	BitcoinTransactionReportParamsReportTypeParams    BitcoinTransactionReportParamsReportType = "params"
)

func (r BitcoinTransactionReportParamsReportType) IsKnown() bool {
	switch r {
	case BitcoinTransactionReportParamsReportTypeRequestID, BitcoinTransactionReportParamsReportTypeParams:
		return true
	}
	return false
}

type BitcoinTransactionScanParams struct {
	AccountAddress param.Field[string]                            `json:"account_address,required"`
	Chain          param.Field[BitcoinTransactionScanParamsChain] `json:"chain,required"`
	// Metadata
	Metadata    param.Field[BitcoinTransactionScanParamsMetadataUnion] `json:"metadata,required"`
	Transaction param.Field[string]                                    `json:"transaction,required"`
	// List of options to include in the response
	//
	// - `Options.validation`: Include Options.validation output in the response
	//
	// - `Options.simulation`: Include Options.simulation output in the response
	Options param.Field[[]BitcoinTransactionScanParamsOption] `json:"options"`
}

func (r BitcoinTransactionScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BitcoinTransactionScanParamsChain string

const (
	BitcoinTransactionScanParamsChainBitcoin BitcoinTransactionScanParamsChain = "bitcoin"
)

func (r BitcoinTransactionScanParamsChain) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanParamsChainBitcoin:
		return true
	}
	return false
}

// Metadata
type BitcoinTransactionScanParamsMetadata struct {
	// Metadata for wallet requests
	Type param.Field[BitcoinTransactionScanParamsMetadataType] `json:"type"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url"`
}

func (r BitcoinTransactionScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionScanParamsMetadata) implementsBitcoinTransactionScanParamsMetadataUnion() {}

// Metadata
//
// Satisfied by [BitcoinTransactionScanParamsMetadataBitcoinWalletRequestMetadata],
// [BitcoinTransactionScanParamsMetadataBitcoinInAppRequestMetadata],
// [BitcoinTransactionScanParamsMetadata].
type BitcoinTransactionScanParamsMetadataUnion interface {
	implementsBitcoinTransactionScanParamsMetadataUnion()
}

type BitcoinTransactionScanParamsMetadataBitcoinWalletRequestMetadata struct {
	// Metadata for wallet requests
	Type param.Field[BitcoinTransactionScanParamsMetadataBitcoinWalletRequestMetadataType] `json:"type,required"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url,required"`
}

func (r BitcoinTransactionScanParamsMetadataBitcoinWalletRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionScanParamsMetadataBitcoinWalletRequestMetadata) implementsBitcoinTransactionScanParamsMetadataUnion() {
}

// Metadata for wallet requests
type BitcoinTransactionScanParamsMetadataBitcoinWalletRequestMetadataType string

const (
	BitcoinTransactionScanParamsMetadataBitcoinWalletRequestMetadataTypeWallet BitcoinTransactionScanParamsMetadataBitcoinWalletRequestMetadataType = "wallet"
)

func (r BitcoinTransactionScanParamsMetadataBitcoinWalletRequestMetadataType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanParamsMetadataBitcoinWalletRequestMetadataTypeWallet:
		return true
	}
	return false
}

type BitcoinTransactionScanParamsMetadataBitcoinInAppRequestMetadata struct {
	// Metadata for in-app requests
	Type param.Field[BitcoinTransactionScanParamsMetadataBitcoinInAppRequestMetadataType] `json:"type"`
}

func (r BitcoinTransactionScanParamsMetadataBitcoinInAppRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionScanParamsMetadataBitcoinInAppRequestMetadata) implementsBitcoinTransactionScanParamsMetadataUnion() {
}

// Metadata for in-app requests
type BitcoinTransactionScanParamsMetadataBitcoinInAppRequestMetadataType string

const (
	BitcoinTransactionScanParamsMetadataBitcoinInAppRequestMetadataTypeInApp BitcoinTransactionScanParamsMetadataBitcoinInAppRequestMetadataType = "in_app"
)

func (r BitcoinTransactionScanParamsMetadataBitcoinInAppRequestMetadataType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanParamsMetadataBitcoinInAppRequestMetadataTypeInApp:
		return true
	}
	return false
}

// Metadata for wallet requests
type BitcoinTransactionScanParamsMetadataType string

const (
	BitcoinTransactionScanParamsMetadataTypeWallet BitcoinTransactionScanParamsMetadataType = "wallet"
	BitcoinTransactionScanParamsMetadataTypeInApp  BitcoinTransactionScanParamsMetadataType = "in_app"
)

func (r BitcoinTransactionScanParamsMetadataType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanParamsMetadataTypeWallet, BitcoinTransactionScanParamsMetadataTypeInApp:
		return true
	}
	return false
}

type BitcoinTransactionScanParamsOption string

const (
	BitcoinTransactionScanParamsOptionValidation BitcoinTransactionScanParamsOption = "validation"
	BitcoinTransactionScanParamsOptionSimulation BitcoinTransactionScanParamsOption = "simulation"
)

func (r BitcoinTransactionScanParamsOption) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanParamsOptionValidation, BitcoinTransactionScanParamsOptionSimulation:
		return true
	}
	return false
}
