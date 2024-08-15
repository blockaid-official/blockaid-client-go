// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"reflect"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/param"
	"github.com/blockaid-official/blockaid-client-go/option"
	"github.com/tidwall/gjson"
)

// StellarService contains methods and other services that help with interacting
// with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStellarService] method instead.
type StellarService struct {
	Options     []option.RequestOption
	Transaction *StellarTransactionService
}

// NewStellarService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStellarService(opts ...option.RequestOption) (r *StellarService) {
	r = &StellarService{}
	r.Options = opts
	r.Transaction = NewStellarTransactionService(opts...)
	return
}

type StellarAssetContractDetailsSchema struct {
	// Address of the asset's contract
	Address string `json:"address,required"`
	// Asset code
	Name string `json:"name,required"`
	// Asset symbol
	Symbol string `json:"symbol,required"`
	// Type of the asset (`CONTRACT`)
	Type StellarAssetContractDetailsSchemaType `json:"type"`
	JSON stellarAssetContractDetailsSchemaJSON `json:"-"`
}

// stellarAssetContractDetailsSchemaJSON contains the JSON metadata for the struct
// [StellarAssetContractDetailsSchema]
type stellarAssetContractDetailsSchemaJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarAssetContractDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarAssetContractDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StellarAssetContractDetailsSchema) implementsStellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAsset() {
}

func (r StellarAssetContractDetailsSchema) implementsStellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAsset() {
}

// Type of the asset (`CONTRACT`)
type StellarAssetContractDetailsSchemaType string

const (
	StellarAssetContractDetailsSchemaTypeContract StellarAssetContractDetailsSchemaType = "CONTRACT"
)

func (r StellarAssetContractDetailsSchemaType) IsKnown() bool {
	switch r {
	case StellarAssetContractDetailsSchemaTypeContract:
		return true
	}
	return false
}

type StellarAssetTransferDetailsSchema struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                               `json:"usd_price"`
	JSON     stellarAssetTransferDetailsSchemaJSON `json:"-"`
}

// stellarAssetTransferDetailsSchemaJSON contains the JSON metadata for the struct
// [StellarAssetTransferDetailsSchema]
type stellarAssetTransferDetailsSchemaJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarAssetTransferDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarAssetTransferDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanRequestParam struct {
	AccountAddress param.Field[string] `json:"account_address,required"`
	// A CAIP-2 chain ID or a Stellar network name
	Chain param.Field[StellarTransactionScanRequestChain] `json:"chain,required"`
	// Metadata
	Metadata    param.Field[StellarTransactionScanRequestMetadataUnionParam] `json:"metadata,required"`
	Transaction param.Field[string]                                          `json:"transaction,required"`
	// List of options to include in the response
	//
	// - `simulation`: Include simulation output in the response
	// - `validation`: Include security validation of the transaction in the response
	Options param.Field[[]StellarTransactionScanRequestOption] `json:"options"`
}

func (r StellarTransactionScanRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A CAIP-2 chain ID or a Stellar network name
type StellarTransactionScanRequestChain string

const (
	StellarTransactionScanRequestChainPubnet    StellarTransactionScanRequestChain = "pubnet"
	StellarTransactionScanRequestChainFuturenet StellarTransactionScanRequestChain = "futurenet"
	StellarTransactionScanRequestChainTestnet   StellarTransactionScanRequestChain = "testnet"
)

func (r StellarTransactionScanRequestChain) IsKnown() bool {
	switch r {
	case StellarTransactionScanRequestChainPubnet, StellarTransactionScanRequestChainFuturenet, StellarTransactionScanRequestChainTestnet:
		return true
	}
	return false
}

// Metadata
type StellarTransactionScanRequestMetadataParam struct {
	// Metadata for wallet requests
	Type param.Field[StellarTransactionScanRequestMetadataType] `json:"type,required"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url"`
}

func (r StellarTransactionScanRequestMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionScanRequestMetadataParam) implementsStellarTransactionScanRequestMetadataUnionParam() {
}

// Metadata
//
// Satisfied by
// [StellarTransactionScanRequestMetadataStellarWalletRequestMetadataParam],
// [StellarTransactionScanRequestMetadataStellarInAppRequestMetadataParam],
// [StellarTransactionScanRequestMetadataParam].
type StellarTransactionScanRequestMetadataUnionParam interface {
	implementsStellarTransactionScanRequestMetadataUnionParam()
}

type StellarTransactionScanRequestMetadataStellarWalletRequestMetadataParam struct {
	// Metadata for wallet requests
	Type param.Field[StellarTransactionScanRequestMetadataStellarWalletRequestMetadataType] `json:"type,required"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url,required"`
}

func (r StellarTransactionScanRequestMetadataStellarWalletRequestMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionScanRequestMetadataStellarWalletRequestMetadataParam) implementsStellarTransactionScanRequestMetadataUnionParam() {
}

// Metadata for wallet requests
type StellarTransactionScanRequestMetadataStellarWalletRequestMetadataType string

const (
	StellarTransactionScanRequestMetadataStellarWalletRequestMetadataTypeWallet StellarTransactionScanRequestMetadataStellarWalletRequestMetadataType = "wallet"
)

func (r StellarTransactionScanRequestMetadataStellarWalletRequestMetadataType) IsKnown() bool {
	switch r {
	case StellarTransactionScanRequestMetadataStellarWalletRequestMetadataTypeWallet:
		return true
	}
	return false
}

type StellarTransactionScanRequestMetadataStellarInAppRequestMetadataParam struct {
	// Metadata for in-app requests
	Type param.Field[StellarTransactionScanRequestMetadataStellarInAppRequestMetadataType] `json:"type,required"`
}

func (r StellarTransactionScanRequestMetadataStellarInAppRequestMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionScanRequestMetadataStellarInAppRequestMetadataParam) implementsStellarTransactionScanRequestMetadataUnionParam() {
}

// Metadata for in-app requests
type StellarTransactionScanRequestMetadataStellarInAppRequestMetadataType string

const (
	StellarTransactionScanRequestMetadataStellarInAppRequestMetadataTypeInApp StellarTransactionScanRequestMetadataStellarInAppRequestMetadataType = "in_app"
)

func (r StellarTransactionScanRequestMetadataStellarInAppRequestMetadataType) IsKnown() bool {
	switch r {
	case StellarTransactionScanRequestMetadataStellarInAppRequestMetadataTypeInApp:
		return true
	}
	return false
}

// Metadata for wallet requests
type StellarTransactionScanRequestMetadataType string

const (
	StellarTransactionScanRequestMetadataTypeWallet StellarTransactionScanRequestMetadataType = "wallet"
	StellarTransactionScanRequestMetadataTypeInApp  StellarTransactionScanRequestMetadataType = "in_app"
)

func (r StellarTransactionScanRequestMetadataType) IsKnown() bool {
	switch r {
	case StellarTransactionScanRequestMetadataTypeWallet, StellarTransactionScanRequestMetadataTypeInApp:
		return true
	}
	return false
}

type StellarTransactionScanRequestOption string

const (
	StellarTransactionScanRequestOptionValidation StellarTransactionScanRequestOption = "validation"
	StellarTransactionScanRequestOptionSimulation StellarTransactionScanRequestOption = "simulation"
)

func (r StellarTransactionScanRequestOption) IsKnown() bool {
	switch r {
	case StellarTransactionScanRequestOptionValidation, StellarTransactionScanRequestOptionSimulation:
		return true
	}
	return false
}

type StellarTransactionScanResponse struct {
	// Simulation result; Only present if simulation option is included in the request
	Simulation StellarTransactionScanResponseSimulation `json:"simulation,nullable"`
	// Validation result; Only present if validation option is included in the request
	Validation StellarTransactionScanResponseValidation `json:"validation,nullable"`
	JSON       stellarTransactionScanResponseJSON       `json:"-"`
}

// stellarTransactionScanResponseJSON contains the JSON metadata for the struct
// [StellarTransactionScanResponse]
type stellarTransactionScanResponseJSON struct {
	Simulation  apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseJSON) RawJSON() string {
	return r.raw
}

// Simulation result; Only present if simulation option is included in the request
type StellarTransactionScanResponseSimulation struct {
	Status StellarTransactionScanResponseSimulationStatus `json:"status,required"`
	// This field can have the runtime type of
	// [map[string][]StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiff].
	AssetsDiffs interface{} `json:"assets_diffs,required"`
	// This field can have the runtime type of
	// [map[string][]StellarTransactionScanResponseSimulationStellarSimulationResultSchemaExposure].
	Exposures interface{} `json:"exposures,required"`
	// This field can have the runtime type of
	// [map[string][]StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsOwnershipDiff].
	AssetsOwnershipDiff interface{} `json:"assets_ownership_diff,required"`
	// This field can have the runtime type of
	// [[]StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAddressDetail].
	AddressDetails interface{} `json:"address_details,required"`
	// This field can have the runtime type of
	// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummary].
	AccountSummary interface{} `json:"account_summary,required"`
	// Error message
	Error string                                       `json:"error"`
	JSON  stellarTransactionScanResponseSimulationJSON `json:"-"`
	union StellarTransactionScanResponseSimulationUnion
}

// stellarTransactionScanResponseSimulationJSON contains the JSON metadata for the
// struct [StellarTransactionScanResponseSimulation]
type stellarTransactionScanResponseSimulationJSON struct {
	Status              apijson.Field
	AssetsDiffs         apijson.Field
	Exposures           apijson.Field
	AssetsOwnershipDiff apijson.Field
	AddressDetails      apijson.Field
	AccountSummary      apijson.Field
	Error               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r stellarTransactionScanResponseSimulationJSON) RawJSON() string {
	return r.raw
}

func (r *StellarTransactionScanResponseSimulation) UnmarshalJSON(data []byte) (err error) {
	*r = StellarTransactionScanResponseSimulation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [StellarTransactionScanResponseSimulationUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchema],
// [StellarTransactionScanResponseSimulationStellarSimulationErrorSchema].
func (r StellarTransactionScanResponseSimulation) AsUnion() StellarTransactionScanResponseSimulationUnion {
	return r.union
}

// Simulation result; Only present if simulation option is included in the request
//
// Union satisfied by
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchema] or
// [StellarTransactionScanResponseSimulationStellarSimulationErrorSchema].
type StellarTransactionScanResponseSimulationUnion interface {
	implementsStellarTransactionScanResponseSimulation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StellarTransactionScanResponseSimulationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationErrorSchema{}),
		},
	)
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchema struct {
	// Summary of the actions and asset transfers that were made by the requested
	// account address
	AccountSummary StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummary `json:"account_summary,required"`
	Status         StellarTransactionScanResponseSimulationStellarSimulationResultSchemaStatus         `json:"status,required"`
	// Details of addresses involved in the transaction
	AddressDetails []StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAddressDetail `json:"address_details"`
	// Mapping between the address of an account to the assets diff during the
	// transaction
	AssetsDiffs map[string][]StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiff `json:"assets_diffs"`
	// Mapping between the address of an account to the ownership diff of the account
	// during the transaction
	AssetsOwnershipDiff map[string][]StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsOwnershipDiff `json:"assets_ownership_diff"`
	// Mapping between the address of an account to the exposure of the assets during
	// the transaction
	Exposures map[string][]StellarTransactionScanResponseSimulationStellarSimulationResultSchemaExposure `json:"exposures"`
	JSON      stellarTransactionScanResponseSimulationStellarSimulationResultSchemaJSON                  `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchema]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaJSON struct {
	AccountSummary      apijson.Field
	Status              apijson.Field
	AddressDetails      apijson.Field
	AssetsDiffs         apijson.Field
	AssetsOwnershipDiff apijson.Field
	Exposures           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchema) implementsStellarTransactionScanResponseSimulation() {
}

// Summary of the actions and asset transfers that were made by the requested
// account address
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummary struct {
	// Exposures made by the requested account address
	AccountExposures []StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposure `json:"account_exposures,required"`
	// Account ownerships diff of the requested account address
	AccountOwnershipsDiff []StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountOwnershipsDiff `json:"account_ownerships_diff,required"`
	// Total USD diff for the requested account address
	TotalUsdDiff StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryTotalUsdDiff `json:"total_usd_diff,required"`
	// Assets diffs of the requested account address
	AssetsDiffs []StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiff `json:"assets_diffs"`
	// Total USD exposure for each of the spender addresses during the transaction
	TotalUsdExposure map[string]float64                                                                      `json:"total_usd_exposure"`
	JSON             stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummary]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryJSON struct {
	AccountExposures      apijson.Field
	AccountOwnershipsDiff apijson.Field
	TotalUsdDiff          apijson.Field
	AssetsDiffs           apijson.Field
	TotalUsdExposure      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposure struct {
	Asset StellarAssetContractDetailsSchema `json:"asset,required"`
	// Mapping between the address of a Spender to the exposure of the asset during the
	// transaction
	Spenders map[string]StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposuresSpender `json:"spenders"`
	JSON     stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposureJSON                `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposureJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposuresSpender struct {
	// Raw value of the exposure
	RawValue int64 `json:"raw_value,required"`
	// USD value of the exposure
	UsdPrice float64 `json:"usd_price,required"`
	// Value of the exposure
	Value float64 `json:"value,required"`
	// Summarized description of the exposure
	Summary string                                                                                                         `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposuresSpenderJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposuresSpenderJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposuresSpender]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposuresSpenderJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposuresSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposuresSpenderJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountOwnershipsDiff struct {
	// List of public keys that can sign on behalf of the account post-transaction
	PostSigners []string `json:"post_signers,required"`
	// List of public keys that can sign on behalf of the account pre-transaction
	PreSigners []string                                                                                                     `json:"pre_signers,required"`
	Type       StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountOwnershipsDiffType `json:"type,required"`
	JSON       stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountOwnershipsDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountOwnershipsDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountOwnershipsDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountOwnershipsDiffJSON struct {
	PostSigners apijson.Field
	PreSigners  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountOwnershipsDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountOwnershipsDiffJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountOwnershipsDiffType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountOwnershipsDiffTypeSetOptions StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountOwnershipsDiffType = "SET_OPTIONS"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountOwnershipsDiffType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountOwnershipsDiffTypeSetOptions:
		return true
	}
	return false
}

// Total USD diff for the requested account address
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryTotalUsdDiff struct {
	// Total incoming USD transfers
	In float64 `json:"in,required"`
	// Total outgoing USD transfers
	Out float64 `json:"out,required"`
	// Total USD transfers
	Total float64                                                                                             `json:"total"`
	JSON  stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryTotalUsdDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryTotalUsdDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryTotalUsdDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiff struct {
	// Asset involved in the transfer
	Asset StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAsset `json:"asset,required"`
	// Incoming transfers of the asset
	In StellarAssetTransferDetailsSchema `json:"in,nullable"`
	// Outgoing transfers of the asset
	Out  StellarAssetTransferDetailsSchema                                                                 `json:"out,nullable"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffJSON) RawJSON() string {
	return r.raw
}

// Asset involved in the transfer
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAsset struct {
	// Type of the asset (`ASSET`)
	Type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetType `json:"type"`
	// Asset code
	Code string `json:"code"`
	// Asset issuer address
	Issuer string `json:"issuer"`
	// Organization name
	OrgName string `json:"org_name"`
	// Organization URL
	OrgURL string `json:"org_url"`
	// Address of the asset's contract
	Address string `json:"address"`
	// Asset code
	Name string `json:"name"`
	// Asset symbol
	Symbol string                                                                                                  `json:"symbol"`
	JSON   stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetJSON `json:"-"`
	union  StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetUnion
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAsset]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetJSON struct {
	Type        apijson.Field
	Code        apijson.Field
	Issuer      apijson.Field
	OrgName     apijson.Field
	OrgURL      apijson.Field
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetJSON) RawJSON() string {
	return r.raw
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAsset) UnmarshalJSON(data []byte) (err error) {
	*r = StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchema],
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchema],
// [StellarAssetContractDetailsSchema].
func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAsset) AsUnion() StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetUnion {
	return r.union
}

// Asset involved in the transfer
//
// Union satisfied by
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchema],
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchema]
// or [StellarAssetContractDetailsSchema].
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetUnion interface {
	implementsStellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarAssetContractDetailsSchema{}),
		},
	)
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchema struct {
	// Asset code
	Code string `json:"code,required"`
	// Asset issuer address
	Issuer string `json:"issuer,required"`
	// Organization name
	OrgName string `json:"org_name,required"`
	// Organization URL
	OrgURL string `json:"org_url,required"`
	// Type of the asset (`ASSET`)
	Type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchemaType `json:"type"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchemaJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchemaJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchema]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchemaJSON struct {
	Code        apijson.Field
	Issuer      apijson.Field
	OrgName     apijson.Field
	OrgURL      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchema) implementsStellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAsset() {
}

// Type of the asset (`ASSET`)
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchemaType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchemaTypeAsset StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchemaType = "ASSET"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchemaType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchemaTypeAsset:
		return true
	}
	return false
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchema struct {
	// Asset code
	Code StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaCode `json:"code"`
	// Type of the asset (`NATIVE`)
	Type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaType `json:"type"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchema]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaJSON struct {
	Code        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchema) implementsStellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAsset() {
}

// Asset code
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaCode string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaCodeXlm StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaCode = "XLM"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaCode) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaCodeXlm:
		return true
	}
	return false
}

// Type of the asset (`NATIVE`)
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaTypeNative StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaType = "NATIVE"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaTypeNative:
		return true
	}
	return false
}

// Type of the asset (`ASSET`)
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetTypeAsset    StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetType = "ASSET"
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetTypeNative   StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetType = "NATIVE"
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetTypeContract StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetType = "CONTRACT"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetTypeAsset, StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetTypeNative, StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetTypeContract:
		return true
	}
	return false
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaStatus string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaStatusSuccess StellarTransactionScanResponseSimulationStellarSimulationResultSchemaStatus = "Success"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaStatus) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaStatusSuccess:
		return true
	}
	return false
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAddressDetail struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// Description of the account
	Description string                                                                                 `json:"description,nullable"`
	JSON        stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAddressDetailJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAddressDetailJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAddressDetail]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAddressDetailJSON struct {
	AccountAddress apijson.Field
	Description    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAddressDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAddressDetailJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiff struct {
	// Asset involved in the transfer
	Asset StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAsset `json:"asset,required"`
	// Incoming transfers of the asset
	In StellarAssetTransferDetailsSchema `json:"in,nullable"`
	// Outgoing transfers of the asset
	Out  StellarAssetTransferDetailsSchema                                                   `json:"out,nullable"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffJSON) RawJSON() string {
	return r.raw
}

// Asset involved in the transfer
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAsset struct {
	// Type of the asset (`ASSET`)
	Type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetType `json:"type"`
	// Asset code
	Code string `json:"code"`
	// Asset issuer address
	Issuer string `json:"issuer"`
	// Organization name
	OrgName string `json:"org_name"`
	// Organization URL
	OrgURL string `json:"org_url"`
	// Address of the asset's contract
	Address string `json:"address"`
	// Asset code
	Name string `json:"name"`
	// Asset symbol
	Symbol string                                                                                    `json:"symbol"`
	JSON   stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetJSON `json:"-"`
	union  StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetUnion
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAsset]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetJSON struct {
	Type        apijson.Field
	Code        apijson.Field
	Issuer      apijson.Field
	OrgName     apijson.Field
	OrgURL      apijson.Field
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetJSON) RawJSON() string {
	return r.raw
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAsset) UnmarshalJSON(data []byte) (err error) {
	*r = StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchema],
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchema],
// [StellarAssetContractDetailsSchema].
func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAsset) AsUnion() StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetUnion {
	return r.union
}

// Asset involved in the transfer
//
// Union satisfied by
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchema],
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchema]
// or [StellarAssetContractDetailsSchema].
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetUnion interface {
	implementsStellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarAssetContractDetailsSchema{}),
		},
	)
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchema struct {
	// Asset code
	Code string `json:"code,required"`
	// Asset issuer address
	Issuer string `json:"issuer,required"`
	// Organization name
	OrgName string `json:"org_name,required"`
	// Organization URL
	OrgURL string `json:"org_url,required"`
	// Type of the asset (`ASSET`)
	Type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchemaType `json:"type"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchemaJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchemaJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchema]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchemaJSON struct {
	Code        apijson.Field
	Issuer      apijson.Field
	OrgName     apijson.Field
	OrgURL      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchema) implementsStellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAsset() {
}

// Type of the asset (`ASSET`)
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchemaType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchemaTypeAsset StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchemaType = "ASSET"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchemaType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchemaTypeAsset:
		return true
	}
	return false
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchema struct {
	// Asset code
	Code StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaCode `json:"code"`
	// Type of the asset (`NATIVE`)
	Type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaType `json:"type"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchema]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaJSON struct {
	Code        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchema) implementsStellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAsset() {
}

// Asset code
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaCode string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaCodeXlm StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaCode = "XLM"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaCode) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaCodeXlm:
		return true
	}
	return false
}

// Type of the asset (`NATIVE`)
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaTypeNative StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaType = "NATIVE"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaTypeNative:
		return true
	}
	return false
}

// Type of the asset (`ASSET`)
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetTypeAsset    StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetType = "ASSET"
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetTypeNative   StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetType = "NATIVE"
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetTypeContract StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetType = "CONTRACT"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetTypeAsset, StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetTypeNative, StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetTypeContract:
		return true
	}
	return false
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsOwnershipDiff struct {
	// List of public keys that can sign on behalf of the account post-transaction
	PostSigners []string `json:"post_signers,required"`
	// List of public keys that can sign on behalf of the account pre-transaction
	PreSigners []string                                                                                     `json:"pre_signers,required"`
	Type       StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsOwnershipDiffType `json:"type,required"`
	JSON       stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsOwnershipDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsOwnershipDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsOwnershipDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsOwnershipDiffJSON struct {
	PostSigners apijson.Field
	PreSigners  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsOwnershipDiffType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsOwnershipDiffTypeSetOptions StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsOwnershipDiffType = "SET_OPTIONS"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsOwnershipDiffType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsOwnershipDiffTypeSetOptions:
		return true
	}
	return false
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaExposure struct {
	Asset StellarAssetContractDetailsSchema `json:"asset,required"`
	// Mapping between the address of a Spender to the exposure of the asset during the
	// transaction
	Spenders map[string]StellarTransactionScanResponseSimulationStellarSimulationResultSchemaExposuresSpender `json:"spenders"`
	JSON     stellarTransactionScanResponseSimulationStellarSimulationResultSchemaExposureJSON                `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaExposureJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaExposuresSpender struct {
	// Raw value of the exposure
	RawValue int64 `json:"raw_value,required"`
	// USD value of the exposure
	UsdPrice float64 `json:"usd_price,required"`
	// Value of the exposure
	Value float64 `json:"value,required"`
	// Summarized description of the exposure
	Summary string                                                                                    `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResultSchemaExposuresSpenderJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaExposuresSpenderJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaExposuresSpender]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaExposuresSpenderJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaExposuresSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaExposuresSpenderJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationErrorSchema struct {
	// Error message
	Error  string                                                                     `json:"error,required"`
	Status StellarTransactionScanResponseSimulationStellarSimulationErrorSchemaStatus `json:"status,required"`
	JSON   stellarTransactionScanResponseSimulationStellarSimulationErrorSchemaJSON   `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationErrorSchemaJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationErrorSchema]
type stellarTransactionScanResponseSimulationStellarSimulationErrorSchemaJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationErrorSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationErrorSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationErrorSchema) implementsStellarTransactionScanResponseSimulation() {
}

type StellarTransactionScanResponseSimulationStellarSimulationErrorSchemaStatus string

const (
	StellarTransactionScanResponseSimulationStellarSimulationErrorSchemaStatusError StellarTransactionScanResponseSimulationStellarSimulationErrorSchemaStatus = "Error"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationErrorSchemaStatus) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationErrorSchemaStatusError:
		return true
	}
	return false
}

type StellarTransactionScanResponseSimulationStatus string

const (
	StellarTransactionScanResponseSimulationStatusSuccess StellarTransactionScanResponseSimulationStatus = "Success"
	StellarTransactionScanResponseSimulationStatusError   StellarTransactionScanResponseSimulationStatus = "Error"
)

func (r StellarTransactionScanResponseSimulationStatus) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStatusSuccess, StellarTransactionScanResponseSimulationStatusError:
		return true
	}
	return false
}

// Validation result; Only present if validation option is included in the request
type StellarTransactionScanResponseValidation struct {
	Status StellarTransactionScanResponseValidationStatus `json:"status,required"`
	// Verdict of the validation
	ResultType StellarTransactionScanResponseValidationResultType `json:"result_type"`
	// A textual description about the validation result
	Description string `json:"description"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason string `json:"reason"`
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification"`
	// This field can have the runtime type of
	// [[]StellarTransactionScanResponseValidationStellarValidationResultSchemaFeature].
	Features interface{} `json:"features,required"`
	// Error message
	Error string                                       `json:"error"`
	JSON  stellarTransactionScanResponseValidationJSON `json:"-"`
	union StellarTransactionScanResponseValidationUnion
}

// stellarTransactionScanResponseValidationJSON contains the JSON metadata for the
// struct [StellarTransactionScanResponseValidation]
type stellarTransactionScanResponseValidationJSON struct {
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

func (r stellarTransactionScanResponseValidationJSON) RawJSON() string {
	return r.raw
}

func (r *StellarTransactionScanResponseValidation) UnmarshalJSON(data []byte) (err error) {
	*r = StellarTransactionScanResponseValidation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [StellarTransactionScanResponseValidationUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StellarTransactionScanResponseValidationStellarValidationResultSchema],
// [StellarTransactionScanResponseValidationStellarValidationErrorSchema].
func (r StellarTransactionScanResponseValidation) AsUnion() StellarTransactionScanResponseValidationUnion {
	return r.union
}

// Validation result; Only present if validation option is included in the request
//
// Union satisfied by
// [StellarTransactionScanResponseValidationStellarValidationResultSchema] or
// [StellarTransactionScanResponseValidationStellarValidationErrorSchema].
type StellarTransactionScanResponseValidationUnion interface {
	implementsStellarTransactionScanResponseValidation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StellarTransactionScanResponseValidationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseValidationStellarValidationResultSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseValidationStellarValidationErrorSchema{}),
		},
	)
}

type StellarTransactionScanResponseValidationStellarValidationResultSchema struct {
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification,required"`
	// A textual description about the validation result
	Description string `json:"description,required"`
	// A list of features about this transaction explaining the validation
	Features []StellarTransactionScanResponseValidationStellarValidationResultSchemaFeature `json:"features,required"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason string `json:"reason,required"`
	// Verdict of the validation
	ResultType StellarTransactionScanResponseValidationStellarValidationResultSchemaResultType `json:"result_type,required"`
	Status     StellarTransactionScanResponseValidationStellarValidationResultSchemaStatus     `json:"status,required"`
	JSON       stellarTransactionScanResponseValidationStellarValidationResultSchemaJSON       `json:"-"`
}

// stellarTransactionScanResponseValidationStellarValidationResultSchemaJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseValidationStellarValidationResultSchema]
type stellarTransactionScanResponseValidationStellarValidationResultSchemaJSON struct {
	Classification apijson.Field
	Description    apijson.Field
	Features       apijson.Field
	Reason         apijson.Field
	ResultType     apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StellarTransactionScanResponseValidationStellarValidationResultSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseValidationStellarValidationResultSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseValidationStellarValidationResultSchema) implementsStellarTransactionScanResponseValidation() {
}

type StellarTransactionScanResponseValidationStellarValidationResultSchemaFeature struct {
	// Address the feature refers to
	Address string `json:"address,required"`
	// Textual description
	Description string `json:"description,required"`
	FeatureID string `json:"feature_id,required"`
	// Feature Classification
	Type StellarTransactionScanResponseValidationStellarValidationResultSchemaFeaturesType `json:"type,required"`
	JSON stellarTransactionScanResponseValidationStellarValidationResultSchemaFeatureJSON  `json:"-"`
}

// stellarTransactionScanResponseValidationStellarValidationResultSchemaFeatureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseValidationStellarValidationResultSchemaFeature]
type stellarTransactionScanResponseValidationStellarValidationResultSchemaFeatureJSON struct {
	Address     apijson.Field
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseValidationStellarValidationResultSchemaFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseValidationStellarValidationResultSchemaFeatureJSON) RawJSON() string {
	return r.raw
}

// Feature Classification
type StellarTransactionScanResponseValidationStellarValidationResultSchemaFeaturesType string

const (
	StellarTransactionScanResponseValidationStellarValidationResultSchemaFeaturesTypeBenign    StellarTransactionScanResponseValidationStellarValidationResultSchemaFeaturesType = "Benign"
	StellarTransactionScanResponseValidationStellarValidationResultSchemaFeaturesTypeWarning   StellarTransactionScanResponseValidationStellarValidationResultSchemaFeaturesType = "Warning"
	StellarTransactionScanResponseValidationStellarValidationResultSchemaFeaturesTypeMalicious StellarTransactionScanResponseValidationStellarValidationResultSchemaFeaturesType = "Malicious"
	StellarTransactionScanResponseValidationStellarValidationResultSchemaFeaturesTypeInfo      StellarTransactionScanResponseValidationStellarValidationResultSchemaFeaturesType = "Info"
)

func (r StellarTransactionScanResponseValidationStellarValidationResultSchemaFeaturesType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseValidationStellarValidationResultSchemaFeaturesTypeBenign, StellarTransactionScanResponseValidationStellarValidationResultSchemaFeaturesTypeWarning, StellarTransactionScanResponseValidationStellarValidationResultSchemaFeaturesTypeMalicious, StellarTransactionScanResponseValidationStellarValidationResultSchemaFeaturesTypeInfo:
		return true
	}
	return false
}

// Verdict of the validation
type StellarTransactionScanResponseValidationStellarValidationResultSchemaResultType string

const (
	StellarTransactionScanResponseValidationStellarValidationResultSchemaResultTypeBenign    StellarTransactionScanResponseValidationStellarValidationResultSchemaResultType = "Benign"
	StellarTransactionScanResponseValidationStellarValidationResultSchemaResultTypeWarning   StellarTransactionScanResponseValidationStellarValidationResultSchemaResultType = "Warning"
	StellarTransactionScanResponseValidationStellarValidationResultSchemaResultTypeMalicious StellarTransactionScanResponseValidationStellarValidationResultSchemaResultType = "Malicious"
)

func (r StellarTransactionScanResponseValidationStellarValidationResultSchemaResultType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseValidationStellarValidationResultSchemaResultTypeBenign, StellarTransactionScanResponseValidationStellarValidationResultSchemaResultTypeWarning, StellarTransactionScanResponseValidationStellarValidationResultSchemaResultTypeMalicious:
		return true
	}
	return false
}

type StellarTransactionScanResponseValidationStellarValidationResultSchemaStatus string

const (
	StellarTransactionScanResponseValidationStellarValidationResultSchemaStatusSuccess StellarTransactionScanResponseValidationStellarValidationResultSchemaStatus = "Success"
)

func (r StellarTransactionScanResponseValidationStellarValidationResultSchemaStatus) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseValidationStellarValidationResultSchemaStatusSuccess:
		return true
	}
	return false
}

type StellarTransactionScanResponseValidationStellarValidationErrorSchema struct {
	// Error message
	Error  string                                                                     `json:"error,required"`
	Status StellarTransactionScanResponseValidationStellarValidationErrorSchemaStatus `json:"status,required"`
	JSON   stellarTransactionScanResponseValidationStellarValidationErrorSchemaJSON   `json:"-"`
}

// stellarTransactionScanResponseValidationStellarValidationErrorSchemaJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseValidationStellarValidationErrorSchema]
type stellarTransactionScanResponseValidationStellarValidationErrorSchemaJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseValidationStellarValidationErrorSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseValidationStellarValidationErrorSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseValidationStellarValidationErrorSchema) implementsStellarTransactionScanResponseValidation() {
}

type StellarTransactionScanResponseValidationStellarValidationErrorSchemaStatus string

const (
	StellarTransactionScanResponseValidationStellarValidationErrorSchemaStatusError StellarTransactionScanResponseValidationStellarValidationErrorSchemaStatus = "Error"
)

func (r StellarTransactionScanResponseValidationStellarValidationErrorSchemaStatus) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseValidationStellarValidationErrorSchemaStatusError:
		return true
	}
	return false
}

type StellarTransactionScanResponseValidationStatus string

const (
	StellarTransactionScanResponseValidationStatusSuccess StellarTransactionScanResponseValidationStatus = "Success"
	StellarTransactionScanResponseValidationStatusError   StellarTransactionScanResponseValidationStatus = "Error"
)

func (r StellarTransactionScanResponseValidationStatus) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseValidationStatusSuccess, StellarTransactionScanResponseValidationStatusError:
		return true
	}
	return false
}

// Verdict of the validation
type StellarTransactionScanResponseValidationResultType string

const (
	StellarTransactionScanResponseValidationResultTypeBenign    StellarTransactionScanResponseValidationResultType = "Benign"
	StellarTransactionScanResponseValidationResultTypeWarning   StellarTransactionScanResponseValidationResultType = "Warning"
	StellarTransactionScanResponseValidationResultTypeMalicious StellarTransactionScanResponseValidationResultType = "Malicious"
)

func (r StellarTransactionScanResponseValidationResultType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseValidationResultTypeBenign, StellarTransactionScanResponseValidationResultTypeWarning, StellarTransactionScanResponseValidationResultTypeMalicious:
		return true
	}
	return false
}
