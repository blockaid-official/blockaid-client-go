// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"reflect"
	"time"

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

type StellarAssetContractDetails struct {
	// Address of the asset's contract
	Address string `json:"address,required"`
	// Asset code
	Name string `json:"name,required"`
	// Asset symbol
	Symbol string `json:"symbol,required"`
	// Type of the asset (`CONTRACT`)
	Type StellarAssetContractDetailsType `json:"type"`
	JSON stellarAssetContractDetailsJSON `json:"-"`
}

// stellarAssetContractDetailsJSON contains the JSON metadata for the struct
// [StellarAssetContractDetails]
type stellarAssetContractDetailsJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarAssetContractDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarAssetContractDetailsJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`CONTRACT`)
type StellarAssetContractDetailsType string

const (
	StellarAssetContractDetailsTypeContract StellarAssetContractDetailsType = "CONTRACT"
)

func (r StellarAssetContractDetailsType) IsKnown() bool {
	switch r {
	case StellarAssetContractDetailsTypeContract:
		return true
	}
	return false
}

type StellarAssetTransferDetails struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                         `json:"usd_price,nullable"`
	JSON     stellarAssetTransferDetailsJSON `json:"-"`
}

// stellarAssetTransferDetailsJSON contains the JSON metadata for the struct
// [StellarAssetTransferDetails]
type stellarAssetTransferDetailsJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarAssetTransferDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarAssetTransferDetailsJSON) RawJSON() string {
	return r.raw
}

type StellarLegacyAssetDetails struct {
	// Asset code
	Code string `json:"code,required"`
	// Asset issuer address
	Issuer string `json:"issuer,required"`
	// Organization name
	OrgName string `json:"org_name,required"`
	// Organization URL
	OrgURL string `json:"org_url,required"`
	// Type of the asset (`ASSET`)
	Type StellarLegacyAssetDetailsType `json:"type"`
	JSON stellarLegacyAssetDetailsJSON `json:"-"`
}

// stellarLegacyAssetDetailsJSON contains the JSON metadata for the struct
// [StellarLegacyAssetDetails]
type stellarLegacyAssetDetailsJSON struct {
	Code        apijson.Field
	Issuer      apijson.Field
	OrgName     apijson.Field
	OrgURL      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarLegacyAssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarLegacyAssetDetailsJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ASSET`)
type StellarLegacyAssetDetailsType string

const (
	StellarLegacyAssetDetailsTypeAsset StellarLegacyAssetDetailsType = "ASSET"
)

func (r StellarLegacyAssetDetailsType) IsKnown() bool {
	switch r {
	case StellarLegacyAssetDetailsTypeAsset:
		return true
	}
	return false
}

type StellarNativeAssetDetails struct {
	// Asset code
	Code StellarNativeAssetDetailsCode `json:"code"`
	// Type of the asset (`NATIVE`)
	Type StellarNativeAssetDetailsType `json:"type"`
	JSON stellarNativeAssetDetailsJSON `json:"-"`
}

// stellarNativeAssetDetailsJSON contains the JSON metadata for the struct
// [StellarNativeAssetDetails]
type stellarNativeAssetDetailsJSON struct {
	Code        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarNativeAssetDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarNativeAssetDetailsJSON) RawJSON() string {
	return r.raw
}

// Asset code
type StellarNativeAssetDetailsCode string

const (
	StellarNativeAssetDetailsCodeXlm StellarNativeAssetDetailsCode = "XLM"
)

func (r StellarNativeAssetDetailsCode) IsKnown() bool {
	switch r {
	case StellarNativeAssetDetailsCodeXlm:
		return true
	}
	return false
}

// Type of the asset (`NATIVE`)
type StellarNativeAssetDetailsType string

const (
	StellarNativeAssetDetailsTypeNative StellarNativeAssetDetailsType = "NATIVE"
)

func (r StellarNativeAssetDetailsType) IsKnown() bool {
	switch r {
	case StellarNativeAssetDetailsTypeNative:
		return true
	}
	return false
}

type StellarSingleAssetExposure struct {
	// Approval value of the ERC20 token
	Approval float64                              `json:"approval,required"`
	Exposure []StellarSingleAssetExposureExposure `json:"exposure,required"`
	// Expiration date of the approval
	Expiration time.Time `json:"expiration,nullable" format:"date-time"`
	// Summarized description of the exposure
	Summary string                         `json:"summary,nullable"`
	JSON    stellarSingleAssetExposureJSON `json:"-"`
}

// stellarSingleAssetExposureJSON contains the JSON metadata for the struct
// [StellarSingleAssetExposure]
type stellarSingleAssetExposureJSON struct {
	Approval    apijson.Field
	Exposure    apijson.Field
	Expiration  apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarSingleAssetExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarSingleAssetExposureJSON) RawJSON() string {
	return r.raw
}

type StellarSingleAssetExposureExposure struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                `json:"usd_price,nullable"`
	JSON     stellarSingleAssetExposureExposureJSON `json:"-"`
}

// stellarSingleAssetExposureExposureJSON contains the JSON metadata for the struct
// [StellarSingleAssetExposureExposure]
type stellarSingleAssetExposureExposureJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarSingleAssetExposureExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarSingleAssetExposureExposureJSON) RawJSON() string {
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
	// - `Options.validation`: Include Options.validation output in the response
	//
	// - `Options.simulation`: Include Options.simulation output in the response
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
	Type param.Field[StellarTransactionScanRequestMetadataType] `json:"type"`
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
	Type param.Field[StellarTransactionScanRequestMetadataStellarInAppRequestMetadataType] `json:"type"`
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
	// [StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummary].
	AccountSummary interface{} `json:"account_summary"`
	// This field can have the runtime type of
	// [[]StellarTransactionScanResponseSimulationStellarSimulationResultAddressDetail].
	AddressDetails interface{} `json:"address_details"`
	// This field can have the runtime type of
	// [map[string][]StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiff].
	AssetsDiffs interface{} `json:"assets_diffs"`
	// This field can have the runtime type of
	// [map[string][]StellarTransactionScanResponseSimulationStellarSimulationResultAssetsOwnershipDiff].
	AssetsOwnershipDiff interface{} `json:"assets_ownership_diff"`
	// Error message
	Error string `json:"error"`
	// This field can have the runtime type of
	// [map[string][]StellarTransactionScanResponseSimulationStellarSimulationResultExposure].
	Exposures interface{}                                  `json:"exposures"`
	JSON      stellarTransactionScanResponseSimulationJSON `json:"-"`
	union     StellarTransactionScanResponseSimulationUnion
}

// stellarTransactionScanResponseSimulationJSON contains the JSON metadata for the
// struct [StellarTransactionScanResponseSimulation]
type stellarTransactionScanResponseSimulationJSON struct {
	Status              apijson.Field
	AccountSummary      apijson.Field
	AddressDetails      apijson.Field
	AssetsDiffs         apijson.Field
	AssetsOwnershipDiff apijson.Field
	Error               apijson.Field
	Exposures           apijson.Field
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
// [StellarTransactionScanResponseSimulationStellarSimulationResult],
// [StellarTransactionScanResponseSimulationStellarSimulationErrorSchema].
func (r StellarTransactionScanResponseSimulation) AsUnion() StellarTransactionScanResponseSimulationUnion {
	return r.union
}

// Simulation result; Only present if simulation option is included in the request
//
// Union satisfied by
// [StellarTransactionScanResponseSimulationStellarSimulationResult] or
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
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResult{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationErrorSchema{}),
		},
	)
}

type StellarTransactionScanResponseSimulationStellarSimulationResult struct {
	// Summary of the actions and asset transfers that were made by the requested
	// account address
	AccountSummary StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummary `json:"account_summary,required"`
	// Ownership diffs of the account addresses
	AssetsOwnershipDiff map[string][]StellarTransactionScanResponseSimulationStellarSimulationResultAssetsOwnershipDiff `json:"assets_ownership_diff,required"`
	Status              StellarTransactionScanResponseSimulationStellarSimulationResultStatus                           `json:"status,required"`
	// Details of addresses involved in the transaction
	AddressDetails []StellarTransactionScanResponseSimulationStellarSimulationResultAddressDetail `json:"address_details"`
	// Mapping between the address of an account to the assets diff during the
	// transaction
	AssetsDiffs map[string][]StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiff `json:"assets_diffs"`
	// Mapping between the address of an account to the exposure of the assets during
	// the transaction
	Exposures map[string][]StellarTransactionScanResponseSimulationStellarSimulationResultExposure `json:"exposures"`
	JSON      stellarTransactionScanResponseSimulationStellarSimulationResultJSON                  `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultJSON contains the
// JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResult]
type stellarTransactionScanResponseSimulationStellarSimulationResultJSON struct {
	AccountSummary      apijson.Field
	AssetsOwnershipDiff apijson.Field
	Status              apijson.Field
	AddressDetails      apijson.Field
	AssetsDiffs         apijson.Field
	Exposures           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResult) implementsStellarTransactionScanResponseSimulation() {
}

// Summary of the actions and asset transfers that were made by the requested
// account address
type StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummary struct {
	// Exposures made by the requested account address
	AccountExposures []StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposure `json:"account_exposures,required"`
	// Ownership diffs of the requested account address
	AccountOwnershipsDiff []StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountOwnershipsDiff `json:"account_ownerships_diff,required"`
	// Total USD diff for the requested account address
	TotalUsdDiff StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryTotalUsdDiff `json:"total_usd_diff,required"`
	// Assets diffs of the requested account address
	AccountAssetsDiffs []StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiff `json:"account_assets_diffs"`
	// Total USD exposure for each of the spender addresses during the transaction
	TotalUsdExposure map[string]float64                                                                `json:"total_usd_exposure"`
	JSON             stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummary]
type stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryJSON struct {
	AccountExposures      apijson.Field
	AccountOwnershipsDiff apijson.Field
	TotalUsdDiff          apijson.Field
	AccountAssetsDiffs    apijson.Field
	TotalUsdExposure      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposure struct {
	// This field can have the runtime type of [StellarLegacyAssetDetails],
	// [StellarNativeAssetDetails].
	Asset interface{} `json:"asset,required"`
	// This field can have the runtime type of [map[string]StellarSingleAssetExposure].
	Spenders interface{}                                                                                      `json:"spenders"`
	JSON     stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposureJSON `json:"-"`
	union    StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresUnion
}

// stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposureJSON) RawJSON() string {
	return r.raw
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposure) UnmarshalJSON(data []byte) (err error) {
	*r = StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresStellarLegacyAssetExposure],
// [StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresStellarNativeAssetExposure].
func (r StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposure) AsUnion() StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresUnion {
	return r.union
}

// Union satisfied by
// [StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresStellarLegacyAssetExposure]
// or
// [StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresStellarNativeAssetExposure].
type StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresUnion interface {
	implementsStellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresStellarLegacyAssetExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresStellarNativeAssetExposure{}),
		},
	)
}

type StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresStellarLegacyAssetExposure struct {
	Asset StellarLegacyAssetDetails `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StellarSingleAssetExposure                                                                                       `json:"spenders"`
	JSON     stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresStellarLegacyAssetExposureJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresStellarLegacyAssetExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresStellarLegacyAssetExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresStellarLegacyAssetExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresStellarLegacyAssetExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresStellarLegacyAssetExposureJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresStellarLegacyAssetExposure) implementsStellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposure() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresStellarNativeAssetExposure struct {
	Asset StellarNativeAssetDetails `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StellarSingleAssetExposure                                                                                       `json:"spenders"`
	JSON     stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresStellarNativeAssetExposureJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresStellarNativeAssetExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresStellarNativeAssetExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresStellarNativeAssetExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresStellarNativeAssetExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresStellarNativeAssetExposureJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposuresStellarNativeAssetExposure) implementsStellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountExposure() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountOwnershipsDiff struct {
	// List of public keys that can sign on behalf of the account post-transaction
	PostSigners []string `json:"post_signers,required"`
	// List of public keys that can sign on behalf of the account pre-transaction
	PreSigners []string                                                                                               `json:"pre_signers,required"`
	Type       StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountOwnershipsDiffType `json:"type"`
	JSON       stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountOwnershipsDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountOwnershipsDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountOwnershipsDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountOwnershipsDiffJSON struct {
	PostSigners apijson.Field
	PreSigners  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountOwnershipsDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountOwnershipsDiffJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountOwnershipsDiffType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountOwnershipsDiffTypeSetOptions StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountOwnershipsDiffType = "SET_OPTIONS"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountOwnershipsDiffType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountOwnershipsDiffTypeSetOptions:
		return true
	}
	return false
}

// Total USD diff for the requested account address
type StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryTotalUsdDiff struct {
	// Total incoming USD transfers
	In float64 `json:"in,required"`
	// Total outgoing USD transfers
	Out float64 `json:"out,required"`
	// Total USD transfers
	Total float64                                                                                       `json:"total"`
	JSON  stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryTotalUsdDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryTotalUsdDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryTotalUsdDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiff struct {
	// This field can have the runtime type of [StellarLegacyAssetDetails],
	// [StellarNativeAssetDetails], [StellarAssetContractDetails].
	Asset interface{} `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StellarAssetTransferDetails `json:"in,nullable"`
	// Details of the outgoing transfer
	Out   StellarAssetTransferDetails                                                                        `json:"out,nullable"`
	JSON  stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffJSON `json:"-"`
	union StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsUnion
}

// stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiff],
// [StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarNativeAssetDiff],
// [StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarContractAssetDiff].
func (r StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiff) AsUnion() StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiff],
// [StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarNativeAssetDiff]
// or
// [StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarContractAssetDiff].
type StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsUnion interface {
	implementsStellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarNativeAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarContractAssetDiff{}),
		},
	)
}

type StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiff struct {
	Asset StellarLegacyAssetDetails `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StellarAssetTransferDetails `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StellarAssetTransferDetails                                                                                               `json:"out,nullable"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiff) implementsStellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiff() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarNativeAssetDiff struct {
	Asset StellarNativeAssetDetails `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StellarAssetTransferDetails `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StellarAssetTransferDetails                                                                                               `json:"out,nullable"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarNativeAssetDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarNativeAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarNativeAssetDiff) implementsStellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiff() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarContractAssetDiff struct {
	Asset StellarAssetContractDetails `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StellarAssetTransferDetails `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StellarAssetTransferDetails                                                                                                 `json:"out,nullable"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarContractAssetDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarContractAssetDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarContractAssetDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarContractAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarContractAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarContractAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiffsStellarContractAssetDiff) implementsStellarTransactionScanResponseSimulationStellarSimulationResultAccountSummaryAccountAssetsDiff() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResultAssetsOwnershipDiff struct {
	// List of public keys that can sign on behalf of the account post-transaction
	PostSigners []string `json:"post_signers,required"`
	// List of public keys that can sign on behalf of the account pre-transaction
	PreSigners []string                                                                               `json:"pre_signers,required"`
	Type       StellarTransactionScanResponseSimulationStellarSimulationResultAssetsOwnershipDiffType `json:"type"`
	JSON       stellarTransactionScanResponseSimulationStellarSimulationResultAssetsOwnershipDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultAssetsOwnershipDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultAssetsOwnershipDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultAssetsOwnershipDiffJSON struct {
	PostSigners apijson.Field
	PreSigners  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultAssetsOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultAssetsOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResultAssetsOwnershipDiffType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultAssetsOwnershipDiffTypeSetOptions StellarTransactionScanResponseSimulationStellarSimulationResultAssetsOwnershipDiffType = "SET_OPTIONS"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultAssetsOwnershipDiffType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultAssetsOwnershipDiffTypeSetOptions:
		return true
	}
	return false
}

type StellarTransactionScanResponseSimulationStellarSimulationResultStatus string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultStatusSuccess StellarTransactionScanResponseSimulationStellarSimulationResultStatus = "Success"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultStatus) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultStatusSuccess:
		return true
	}
	return false
}

type StellarTransactionScanResponseSimulationStellarSimulationResultAddressDetail struct {
	// Encoded public key of the account
	AccountAddress interface{} `json:"account_address,required"`
	// Description of the account
	Description string                                                                           `json:"description,nullable"`
	JSON        stellarTransactionScanResponseSimulationStellarSimulationResultAddressDetailJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultAddressDetailJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultAddressDetail]
type stellarTransactionScanResponseSimulationStellarSimulationResultAddressDetailJSON struct {
	AccountAddress apijson.Field
	Description    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultAddressDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultAddressDetailJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiff struct {
	// This field can have the runtime type of [StellarLegacyAssetDetails],
	// [StellarNativeAssetDetails], [StellarAssetContractDetails].
	Asset interface{} `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StellarAssetTransferDetails `json:"in,nullable"`
	// Details of the outgoing transfer
	Out   StellarAssetTransferDetails                                                   `json:"out,nullable"`
	JSON  stellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffJSON `json:"-"`
	union StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsUnion
}

// stellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarLegacyAssetDiff],
// [StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarNativeAssetDiff],
// [StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarContractAssetDiff].
func (r StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiff) AsUnion() StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarLegacyAssetDiff],
// [StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarNativeAssetDiff]
// or
// [StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarContractAssetDiff].
type StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsUnion interface {
	implementsStellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarLegacyAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarNativeAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarContractAssetDiff{}),
		},
	)
}

type StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarLegacyAssetDiff struct {
	Asset StellarLegacyAssetDetails `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StellarAssetTransferDetails `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StellarAssetTransferDetails                                                                          `json:"out,nullable"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarLegacyAssetDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarLegacyAssetDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarLegacyAssetDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarLegacyAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarLegacyAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarLegacyAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarLegacyAssetDiff) implementsStellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiff() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarNativeAssetDiff struct {
	Asset StellarNativeAssetDetails `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StellarAssetTransferDetails `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StellarAssetTransferDetails                                                                          `json:"out,nullable"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarNativeAssetDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarNativeAssetDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarNativeAssetDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarNativeAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarNativeAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarNativeAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarNativeAssetDiff) implementsStellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiff() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarContractAssetDiff struct {
	Asset StellarAssetContractDetails `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StellarAssetTransferDetails `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StellarAssetTransferDetails                                                                            `json:"out,nullable"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarContractAssetDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarContractAssetDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarContractAssetDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarContractAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarContractAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarContractAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiffsStellarContractAssetDiff) implementsStellarTransactionScanResponseSimulationStellarSimulationResultAssetsDiff() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResultExposure struct {
	// This field can have the runtime type of [StellarLegacyAssetDetails],
	// [StellarNativeAssetDetails].
	Asset interface{} `json:"asset,required"`
	// This field can have the runtime type of [map[string]StellarSingleAssetExposure].
	Spenders interface{}                                                                 `json:"spenders"`
	JSON     stellarTransactionScanResponseSimulationStellarSimulationResultExposureJSON `json:"-"`
	union    StellarTransactionScanResponseSimulationStellarSimulationResultExposuresUnion
}

// stellarTransactionScanResponseSimulationStellarSimulationResultExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResultExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultExposureJSON) RawJSON() string {
	return r.raw
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultExposure) UnmarshalJSON(data []byte) (err error) {
	*r = StellarTransactionScanResponseSimulationStellarSimulationResultExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StellarTransactionScanResponseSimulationStellarSimulationResultExposuresUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StellarTransactionScanResponseSimulationStellarSimulationResultExposuresStellarLegacyAssetExposure],
// [StellarTransactionScanResponseSimulationStellarSimulationResultExposuresStellarNativeAssetExposure].
func (r StellarTransactionScanResponseSimulationStellarSimulationResultExposure) AsUnion() StellarTransactionScanResponseSimulationStellarSimulationResultExposuresUnion {
	return r.union
}

// Union satisfied by
// [StellarTransactionScanResponseSimulationStellarSimulationResultExposuresStellarLegacyAssetExposure]
// or
// [StellarTransactionScanResponseSimulationStellarSimulationResultExposuresStellarNativeAssetExposure].
type StellarTransactionScanResponseSimulationStellarSimulationResultExposuresUnion interface {
	implementsStellarTransactionScanResponseSimulationStellarSimulationResultExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StellarTransactionScanResponseSimulationStellarSimulationResultExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultExposuresStellarLegacyAssetExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultExposuresStellarNativeAssetExposure{}),
		},
	)
}

type StellarTransactionScanResponseSimulationStellarSimulationResultExposuresStellarLegacyAssetExposure struct {
	Asset StellarLegacyAssetDetails `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StellarSingleAssetExposure                                                                  `json:"spenders"`
	JSON     stellarTransactionScanResponseSimulationStellarSimulationResultExposuresStellarLegacyAssetExposureJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultExposuresStellarLegacyAssetExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultExposuresStellarLegacyAssetExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResultExposuresStellarLegacyAssetExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultExposuresStellarLegacyAssetExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultExposuresStellarLegacyAssetExposureJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultExposuresStellarLegacyAssetExposure) implementsStellarTransactionScanResponseSimulationStellarSimulationResultExposure() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResultExposuresStellarNativeAssetExposure struct {
	Asset StellarNativeAssetDetails `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StellarSingleAssetExposure                                                                  `json:"spenders"`
	JSON     stellarTransactionScanResponseSimulationStellarSimulationResultExposuresStellarNativeAssetExposureJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultExposuresStellarNativeAssetExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultExposuresStellarNativeAssetExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResultExposuresStellarNativeAssetExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultExposuresStellarNativeAssetExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultExposuresStellarNativeAssetExposureJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultExposuresStellarNativeAssetExposure) implementsStellarTransactionScanResponseSimulationStellarSimulationResultExposure() {
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
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification"`
	// A textual description about the validation result
	Description string `json:"description"`
	// Error message
	Error string `json:"error"`
	// This field can have the runtime type of
	// [[]StellarTransactionScanResponseValidationStellarValidationResultFeature].
	Features interface{} `json:"features"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason string `json:"reason"`
	// Verdict of the validation
	ResultType StellarTransactionScanResponseValidationResultType `json:"result_type"`
	JSON       stellarTransactionScanResponseValidationJSON       `json:"-"`
	union      StellarTransactionScanResponseValidationUnion
}

// stellarTransactionScanResponseValidationJSON contains the JSON metadata for the
// struct [StellarTransactionScanResponseValidation]
type stellarTransactionScanResponseValidationJSON struct {
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
// [StellarTransactionScanResponseValidationStellarValidationResult],
// [StellarTransactionScanResponseValidationStellarValidationErrorSchema].
func (r StellarTransactionScanResponseValidation) AsUnion() StellarTransactionScanResponseValidationUnion {
	return r.union
}

// Validation result; Only present if validation option is included in the request
//
// Union satisfied by
// [StellarTransactionScanResponseValidationStellarValidationResult] or
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
			Type:       reflect.TypeOf(StellarTransactionScanResponseValidationStellarValidationResult{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseValidationStellarValidationErrorSchema{}),
		},
	)
}

type StellarTransactionScanResponseValidationStellarValidationResult struct {
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification,required"`
	// A textual description about the validation result
	Description string                                                                   `json:"description,required"`
	Features    []StellarTransactionScanResponseValidationStellarValidationResultFeature `json:"features,required"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason string `json:"reason,required"`
	// Verdict of the validation
	ResultType StellarTransactionScanResponseValidationStellarValidationResultResultType `json:"result_type,required"`
	Status     StellarTransactionScanResponseValidationStellarValidationResultStatus     `json:"status,required"`
	JSON       stellarTransactionScanResponseValidationStellarValidationResultJSON       `json:"-"`
}

// stellarTransactionScanResponseValidationStellarValidationResultJSON contains the
// JSON metadata for the struct
// [StellarTransactionScanResponseValidationStellarValidationResult]
type stellarTransactionScanResponseValidationStellarValidationResultJSON struct {
	Classification apijson.Field
	Description    apijson.Field
	Features       apijson.Field
	Reason         apijson.Field
	ResultType     apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StellarTransactionScanResponseValidationStellarValidationResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseValidationStellarValidationResultJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseValidationStellarValidationResult) implementsStellarTransactionScanResponseValidation() {
}

type StellarTransactionScanResponseValidationStellarValidationResultFeature struct {
	// Address the feature refers to
	Address string `json:"address,required"`
	// Textual description
	Description string `json:"description,required"`
	FeatureID   string `json:"feature_id,required"`
	// Feature Classification
	Type StellarTransactionScanResponseValidationStellarValidationResultFeaturesType `json:"type,required"`
	JSON stellarTransactionScanResponseValidationStellarValidationResultFeatureJSON  `json:"-"`
}

// stellarTransactionScanResponseValidationStellarValidationResultFeatureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseValidationStellarValidationResultFeature]
type stellarTransactionScanResponseValidationStellarValidationResultFeatureJSON struct {
	Address     apijson.Field
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseValidationStellarValidationResultFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseValidationStellarValidationResultFeatureJSON) RawJSON() string {
	return r.raw
}

// Feature Classification
type StellarTransactionScanResponseValidationStellarValidationResultFeaturesType string

const (
	StellarTransactionScanResponseValidationStellarValidationResultFeaturesTypeBenign    StellarTransactionScanResponseValidationStellarValidationResultFeaturesType = "Benign"
	StellarTransactionScanResponseValidationStellarValidationResultFeaturesTypeWarning   StellarTransactionScanResponseValidationStellarValidationResultFeaturesType = "Warning"
	StellarTransactionScanResponseValidationStellarValidationResultFeaturesTypeMalicious StellarTransactionScanResponseValidationStellarValidationResultFeaturesType = "Malicious"
	StellarTransactionScanResponseValidationStellarValidationResultFeaturesTypeInfo      StellarTransactionScanResponseValidationStellarValidationResultFeaturesType = "Info"
)

func (r StellarTransactionScanResponseValidationStellarValidationResultFeaturesType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseValidationStellarValidationResultFeaturesTypeBenign, StellarTransactionScanResponseValidationStellarValidationResultFeaturesTypeWarning, StellarTransactionScanResponseValidationStellarValidationResultFeaturesTypeMalicious, StellarTransactionScanResponseValidationStellarValidationResultFeaturesTypeInfo:
		return true
	}
	return false
}

// Verdict of the validation
type StellarTransactionScanResponseValidationStellarValidationResultResultType string

const (
	StellarTransactionScanResponseValidationStellarValidationResultResultTypeBenign    StellarTransactionScanResponseValidationStellarValidationResultResultType = "Benign"
	StellarTransactionScanResponseValidationStellarValidationResultResultTypeWarning   StellarTransactionScanResponseValidationStellarValidationResultResultType = "Warning"
	StellarTransactionScanResponseValidationStellarValidationResultResultTypeMalicious StellarTransactionScanResponseValidationStellarValidationResultResultType = "Malicious"
)

func (r StellarTransactionScanResponseValidationStellarValidationResultResultType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseValidationStellarValidationResultResultTypeBenign, StellarTransactionScanResponseValidationStellarValidationResultResultTypeWarning, StellarTransactionScanResponseValidationStellarValidationResultResultTypeMalicious:
		return true
	}
	return false
}

type StellarTransactionScanResponseValidationStellarValidationResultStatus string

const (
	StellarTransactionScanResponseValidationStellarValidationResultStatusSuccess StellarTransactionScanResponseValidationStellarValidationResultStatus = "Success"
)

func (r StellarTransactionScanResponseValidationStellarValidationResultStatus) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseValidationStellarValidationResultStatusSuccess:
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
