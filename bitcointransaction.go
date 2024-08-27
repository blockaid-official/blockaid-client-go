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
	JSON       bitcoinTransactionScanResponseJSON       `json:"-"`
}

// bitcoinTransactionScanResponseJSON contains the JSON metadata for the struct
// [BitcoinTransactionScanResponse]
type bitcoinTransactionScanResponseJSON struct {
	Simulation  apijson.Field
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
	Status BitcoinTransactionScanResponseSimulationStatus `json:"status"`
	// This field can have the runtime type of
	// [map[string][]BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiff].
	AssetsDiffs interface{} `json:"assets_diffs,required"`
	// Error message
	Error string                                       `json:"error"`
	JSON  bitcoinTransactionScanResponseSimulationJSON `json:"-"`
	union BitcoinTransactionScanResponseSimulationUnion
}

// bitcoinTransactionScanResponseSimulationJSON contains the JSON metadata for the
// struct [BitcoinTransactionScanResponseSimulation]
type bitcoinTransactionScanResponseSimulationJSON struct {
	Status      apijson.Field
	AssetsDiffs apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationSchema],
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchema].
func (r BitcoinTransactionScanResponseSimulation) AsUnion() BitcoinTransactionScanResponseSimulationUnion {
	return r.union
}

// Simulation result; Only present if simulation option is included in the request
//
// Union satisfied by
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationSchema] or
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
			Type:       reflect.TypeOf(BitcoinTransactionScanResponseSimulationBitcoinSimulationSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchema{}),
		},
	)
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationSchema struct {
	// A dictionary describing the assets differences as a result of this transaction
	// for every involved address
	AssetsDiffs map[string][]BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiff `json:"assets_diffs,required"`
	Status      BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaStatus                  `json:"status"`
	JSON        bitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaJSON                    `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaJSON contains the
// JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationSchema]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaJSON struct {
	AssetsDiffs apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaJSON) RawJSON() string {
	return r.raw
}

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationSchema) implementsBitcoinTransactionScanResponseSimulation() {
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiff struct {
	// Description of the asset for the current diff
	Asset BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsAsset `json:"asset,required"`
	// The assets received by the address
	In []BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsIn `json:"in"`
	// The assets sent by the address
	Out  []BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsOut `json:"out"`
	JSON bitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffJSON   `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiff]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffJSON) RawJSON() string {
	return r.raw
}

// Description of the asset for the current diff
type BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsAsset struct {
	ChainName  string                                                                              `json:"chain_name,required"`
	Decimals   int64                                                                               `json:"decimals,required"`
	LogoURL    string                                                                              `json:"logo_url,required"`
	Name       string                                                                              `json:"name,required"`
	Type       BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsAssetType `json:"type,required"`
	ID         string                                                                              `json:"id"`
	SpacedName string                                                                              `json:"spaced_name"`
	Symbol     string                                                                              `json:"symbol"`
	JSON       bitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsAssetJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsAssetJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsAsset]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsAssetJSON struct {
	ChainName   apijson.Field
	Decimals    apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SpacedName  apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsAssetJSON) RawJSON() string {
	return r.raw
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsAssetType string

const (
	BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsAssetTypeNative  BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsAssetType = "NATIVE"
	BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsAssetTypeRune    BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsAssetType = "RUNE"
	BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsAssetTypeOrdinal BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsAssetType = "ORDINAL"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsAssetType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsAssetTypeNative, BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsAssetTypeRune, BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsAssetTypeOrdinal:
		return true
	}
	return false
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsIn struct {
	Name     string                                                                           `json:"name"`
	RawValue string                                                                           `json:"raw_value"`
	Summary  string                                                                           `json:"summary"`
	TokenID  string                                                                           `json:"token_id"`
	UsdPrice string                                                                           `json:"usd_price"`
	Value    string                                                                           `json:"value"`
	JSON     bitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsInJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsInJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsIn]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsInJSON struct {
	Name        apijson.Field
	RawValue    apijson.Field
	Summary     apijson.Field
	TokenID     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsInJSON) RawJSON() string {
	return r.raw
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsOut struct {
	Name     string                                                                            `json:"name"`
	RawValue string                                                                            `json:"raw_value"`
	Summary  string                                                                            `json:"summary"`
	TokenID  string                                                                            `json:"token_id"`
	UsdPrice string                                                                            `json:"usd_price"`
	Value    string                                                                            `json:"value"`
	JSON     bitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsOutJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsOutJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsOut]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsOutJSON struct {
	Name        apijson.Field
	RawValue    apijson.Field
	Summary     apijson.Field
	TokenID     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaAssetsDiffsOutJSON) RawJSON() string {
	return r.raw
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaStatus string

const (
	BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaStatusSuccess BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaStatus = "Success"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaStatus) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinSimulationSchemaStatusSuccess:
		return true
	}
	return false
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchema struct {
	// Error message
	Error  string                                                                     `json:"error,required"`
	Status BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchemaStatus `json:"status"`
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

type BitcoinTransactionScanParams struct {
	// List of options to include in the response
	//
	// - `simulation`: Include simulation output in the response
	Options param.Field[[]BitcoinTransactionScanParamsOption] `json:"options,required"`
	// The transaction encoded as a hex string
	Transaction param.Field[string]                            `json:"transaction,required"`
	Chain       param.Field[BitcoinTransactionScanParamsChain] `json:"chain"`
	// Metadata
	Metadata param.Field[BitcoinTransactionScanParamsMetadataUnion] `json:"metadata"`
}

func (r BitcoinTransactionScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BitcoinTransactionScanParamsOption string

const (
	BitcoinTransactionScanParamsOptionSimulation BitcoinTransactionScanParamsOption = "simulation"
)

func (r BitcoinTransactionScanParamsOption) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanParamsOptionSimulation:
		return true
	}
	return false
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
	// Metadata for in-app requests
	Type param.Field[BitcoinTransactionScanParamsMetadataType] `json:"type,required"`
	// URL of the dApp that originated the transaction
	URL param.Field[string] `json:"url"`
}

func (r BitcoinTransactionScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionScanParamsMetadata) implementsBitcoinTransactionScanParamsMetadataUnion() {}

// Metadata
//
// Satisfied by [BitcoinTransactionScanParamsMetadataBitcoinInAppRequestMetadata],
// [BitcoinTransactionScanParamsMetadataBitcoinWalletRequestMetadata],
// [BitcoinTransactionScanParamsMetadata].
type BitcoinTransactionScanParamsMetadataUnion interface {
	implementsBitcoinTransactionScanParamsMetadataUnion()
}

type BitcoinTransactionScanParamsMetadataBitcoinInAppRequestMetadata struct {
	// Metadata for in-app requests
	Type param.Field[BitcoinTransactionScanParamsMetadataBitcoinInAppRequestMetadataType] `json:"type,required"`
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

type BitcoinTransactionScanParamsMetadataBitcoinWalletRequestMetadata struct {
	// Metadata for wallet requests
	Type param.Field[BitcoinTransactionScanParamsMetadataBitcoinWalletRequestMetadataType] `json:"type,required"`
	// URL of the dApp that originated the transaction
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

// Metadata for in-app requests
type BitcoinTransactionScanParamsMetadataType string

const (
	BitcoinTransactionScanParamsMetadataTypeInApp  BitcoinTransactionScanParamsMetadataType = "in_app"
	BitcoinTransactionScanParamsMetadataTypeWallet BitcoinTransactionScanParamsMetadataType = "wallet"
)

func (r BitcoinTransactionScanParamsMetadataType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanParamsMetadataTypeInApp, BitcoinTransactionScanParamsMetadataTypeWallet:
		return true
	}
	return false
}
