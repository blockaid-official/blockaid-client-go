// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"reflect"
	"time"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
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
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                `json:"summary,nullable"`
	JSON    stellarAssetTransferDetailsSchemaJSON `json:"-"`
}

// stellarAssetTransferDetailsSchemaJSON contains the JSON metadata for the struct
// [StellarAssetTransferDetailsSchema]
type stellarAssetTransferDetailsSchemaJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarAssetTransferDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarAssetTransferDetailsSchemaJSON) RawJSON() string {
	return r.raw
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
	// This field can have the runtime type of
	// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummary].
	AccountSummary interface{} `json:"account_summary,required"`
	// This field can have the runtime type of
	// [[]StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAddressDetail].
	AddressDetails interface{} `json:"address_details,required"`
	// This field can have the runtime type of
	// [map[string][]StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiff].
	AssetsDiffs interface{} `json:"assets_diffs,required"`
	// This field can have the runtime type of
	// [map[string][]StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsOwnershipDiff].
	AssetsOwnershipDiff interface{} `json:"assets_ownership_diff,required"`
	// This field can have the runtime type of
	// [map[string][]StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposure].
	Exposures interface{}                                    `json:"exposures,required"`
	Status    StellarTransactionScanResponseSimulationStatus `json:"status,required"`
	// Error message
	Error string                                       `json:"error"`
	JSON  stellarTransactionScanResponseSimulationJSON `json:"-"`
	union StellarTransactionScanResponseSimulationUnion
}

// stellarTransactionScanResponseSimulationJSON contains the JSON metadata for the
// struct [StellarTransactionScanResponseSimulation]
type stellarTransactionScanResponseSimulationJSON struct {
	AccountSummary      apijson.Field
	AddressDetails      apijson.Field
	AssetsDiffs         apijson.Field
	AssetsOwnershipDiff apijson.Field
	Exposures           apijson.Field
	Status              apijson.Field
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
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaType],
// [StellarTransactionScanResponseSimulationStellarSimulationErrorSchema].
func (r StellarTransactionScanResponseSimulation) AsUnion() StellarTransactionScanResponseSimulationUnion {
	return r.union
}

// Simulation result; Only present if simulation option is included in the request
//
// Union satisfied by
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaType] or
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
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultSchemaType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationErrorSchema{}),
		},
	)
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaType struct {
	// Summary of the actions and asset transfers that were made by the requested
	// account address
	AccountSummary StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummary `json:"account_summary,required"`
	// Ownership diffs of the account addresses
	AssetsOwnershipDiff map[string][]StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsOwnershipDiff `json:"assets_ownership_diff,required"`
	Status              StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeStatus                           `json:"status,required"`
	// Details of addresses involved in the transaction
	AddressDetails []StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAddressDetail `json:"address_details"`
	// Mapping between the address of an account to the assets diff during the
	// transaction
	AssetsDiffs map[string][]StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiff `json:"assets_diffs"`
	// Mapping between the address of an account to the exposure of the assets during
	// the transaction
	Exposures map[string][]StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposure `json:"exposures"`
	JSON      stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeJSON                  `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaType]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeJSON struct {
	AccountSummary      apijson.Field
	AssetsOwnershipDiff apijson.Field
	Status              apijson.Field
	AddressDetails      apijson.Field
	AssetsDiffs         apijson.Field
	Exposures           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaType) implementsStellarTransactionScanResponseSimulation() {
}

// Summary of the actions and asset transfers that were made by the requested
// account address
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummary struct {
	// Exposures made by the requested account address
	AccountExposures []StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposure `json:"account_exposures,required"`
	// Ownership diffs of the requested account address
	AccountOwnershipsDiff []StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountOwnershipsDiff `json:"account_ownerships_diff,required"`
	// Total USD diff for the requested account address
	TotalUsdDiff StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryTotalUsdDiff `json:"total_usd_diff,required"`
	// Assets diffs of the requested account address
	AccountAssetsDiffs []StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiff `json:"account_assets_diffs"`
	// Total USD exposure for each of the spender addresses during the transaction
	TotalUsdExposure map[string]float64                                                                          `json:"total_usd_exposure"`
	JSON             stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummary]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryJSON struct {
	AccountExposures      apijson.Field
	AccountOwnershipsDiff apijson.Field
	TotalUsdDiff          apijson.Field
	AccountAssetsDiffs    apijson.Field
	TotalUsdExposure      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposure struct {
	// This field can have the runtime type of
	// [map[string]StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureSpender],
	// [map[string]StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureSpender].
	Spenders interface{} `json:"spenders,required"`
	// This field can have the runtime type of
	// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureAsset],
	// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureAsset].
	Asset interface{}                                                                                                `json:"asset"`
	JSON  stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposureJSON `json:"-"`
	union StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresUnion
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposureJSON struct {
	Spenders    apijson.Field
	Asset       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposureJSON) RawJSON() string {
	return r.raw
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposure) UnmarshalJSON(data []byte) (err error) {
	*r = StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposure],
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposure].
func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposure) AsUnion() StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresUnion {
	return r.union
}

// Union satisfied by
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposure]
// or
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposure].
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresUnion interface {
	implementsStellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposure{}),
		},
	)
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposure struct {
	Asset StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureAsset `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureSpender `json:"spenders"`
	JSON     stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureJSON               `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposure) implementsStellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposure() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureAsset struct {
	// Asset code
	Code string `json:"code,required"`
	// Asset issuer address
	Issuer string `json:"issuer,required"`
	// Organization name
	OrgName string `json:"org_name,required"`
	// Organization URL
	OrgURL string `json:"org_url,required"`
	// Type of the asset (`ASSET`)
	Type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureAssetType `json:"type"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureAssetJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureAssetJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureAsset]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureAssetJSON struct {
	Code        apijson.Field
	Issuer      apijson.Field
	OrgName     apijson.Field
	OrgURL      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ASSET`)
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureAssetType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureAssetTypeAsset StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureAssetType = "ASSET"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureAssetType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureAssetTypeAsset:
		return true
	}
	return false
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureSpender struct {
	// Approval value of the ERC20 token
	Approval string                                                                                                                                              `json:"approval,required"`
	Exposure []StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureSpendersExposure `json:"exposure,required"`
	// Expiration date of the approval
	Expiration time.Time `json:"expiration,nullable" format:"date-time"`
	// Summarized description of the exposure
	Summary string                                                                                                                                       `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureSpenderJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureSpenderJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureSpender]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureSpenderJSON struct {
	Approval    apijson.Field
	Exposure    apijson.Field
	Expiration  apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureSpendersExposure struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                                                `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureSpendersExposureJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureSpendersExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureSpendersExposureJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarLegacyAssetExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposure struct {
	Asset StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureAsset `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureSpender `json:"spenders"`
	JSON     stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureJSON               `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposure) implementsStellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposure() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureAsset struct {
	// Asset code
	Code StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureAssetCode `json:"code"`
	// Type of the asset (`NATIVE`)
	Type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureAssetType `json:"type"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureAssetJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureAssetJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureAsset]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureAssetJSON struct {
	Code        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureAssetJSON) RawJSON() string {
	return r.raw
}

// Asset code
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureAssetCode string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureAssetCodeXlm StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureAssetCode = "XLM"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureAssetCode) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureAssetCodeXlm:
		return true
	}
	return false
}

// Type of the asset (`NATIVE`)
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureAssetType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureAssetTypeNative StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureAssetType = "NATIVE"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureAssetType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureAssetTypeNative:
		return true
	}
	return false
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureSpender struct {
	// Approval value of the ERC20 token
	Approval string                                                                                                                                              `json:"approval,required"`
	Exposure []StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureSpendersExposure `json:"exposure,required"`
	// Expiration date of the approval
	Expiration time.Time `json:"expiration,nullable" format:"date-time"`
	// Summarized description of the exposure
	Summary string                                                                                                                                       `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureSpenderJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureSpenderJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureSpender]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureSpenderJSON struct {
	Approval    apijson.Field
	Exposure    apijson.Field
	Expiration  apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureSpendersExposure struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                                                `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureSpendersExposureJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureSpendersExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureSpendersExposureJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountExposuresStellarNativeAssetExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountOwnershipsDiff struct {
	// List of public keys that can sign on behalf of the account post-transaction
	PostSigners []string `json:"post_signers,required"`
	// List of public keys that can sign on behalf of the account pre-transaction
	PreSigners []string                                                                                                         `json:"pre_signers,required"`
	Type       StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountOwnershipsDiffType `json:"type"`
	JSON       stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountOwnershipsDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountOwnershipsDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountOwnershipsDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountOwnershipsDiffJSON struct {
	PostSigners apijson.Field
	PreSigners  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountOwnershipsDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountOwnershipsDiffJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountOwnershipsDiffType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountOwnershipsDiffTypeSetOptions StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountOwnershipsDiffType = "SET_OPTIONS"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountOwnershipsDiffType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountOwnershipsDiffTypeSetOptions:
		return true
	}
	return false
}

// Total USD diff for the requested account address
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryTotalUsdDiff struct {
	// Total incoming USD transfers
	In float64 `json:"in,required"`
	// Total outgoing USD transfers
	Out float64 `json:"out,required"`
	// Total USD transfers
	Total float64                                                                                                 `json:"total"`
	JSON  stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryTotalUsdDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryTotalUsdDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryTotalUsdDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiff struct {
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of
	// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAsset],
	// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAsset],
	// [StellarAssetContractDetailsSchema].
	Asset interface{} `json:"asset"`
	// Details of the incoming transfer
	In StellarAssetTransferDetailsSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out   StellarAssetTransferDetailsSchema                                                                            `json:"out,nullable"`
	JSON  stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffJSON `json:"-"`
	union StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsUnion
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffJSON struct {
	AssetType   apijson.Field
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiff],
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiff],
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarContractAssetDiff].
func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiff) AsUnion() StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiff],
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiff]
// or
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarContractAssetDiff].
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsUnion interface {
	implementsStellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarContractAssetDiff{}),
		},
	)
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiff struct {
	Asset StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StellarAssetTransferDetailsSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StellarAssetTransferDetailsSchema                                                                                                   `json:"out,nullable"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiff) implementsStellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiff() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAsset struct {
	// Asset code
	Code string `json:"code,required"`
	// Asset issuer address
	Issuer string `json:"issuer,required"`
	// Organization name
	OrgName string `json:"org_name,required"`
	// Organization URL
	OrgURL string `json:"org_url,required"`
	// Type of the asset (`ASSET`)
	Type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAssetType `json:"type"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAssetJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAsset]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAssetJSON struct {
	Code        apijson.Field
	Issuer      apijson.Field
	OrgName     apijson.Field
	OrgURL      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ASSET`)
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAssetType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAssetTypeAsset StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAssetType = "ASSET"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAssetType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAssetTypeAsset:
		return true
	}
	return false
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiff struct {
	Asset StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StellarAssetTransferDetailsSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StellarAssetTransferDetailsSchema                                                                                                   `json:"out,nullable"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiff) implementsStellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiff() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAsset struct {
	// Asset code
	Code StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetCode `json:"code"`
	// Type of the asset (`NATIVE`)
	Type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetType `json:"type"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAsset]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetJSON struct {
	Code        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Asset code
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetCode string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetCodeXlm StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetCode = "XLM"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetCode) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetCodeXlm:
		return true
	}
	return false
}

// Type of the asset (`NATIVE`)
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetTypeNative StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetType = "NATIVE"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetTypeNative:
		return true
	}
	return false
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarContractAssetDiff struct {
	Asset StellarAssetContractDetailsSchema `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StellarAssetTransferDetailsSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StellarAssetTransferDetailsSchema                                                                                                     `json:"out,nullable"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarContractAssetDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarContractAssetDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarContractAssetDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarContractAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarContractAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarContractAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiffsStellarContractAssetDiff) implementsStellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAccountSummaryAccountAssetsDiff() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsOwnershipDiff struct {
	// List of public keys that can sign on behalf of the account post-transaction
	PostSigners []string `json:"post_signers,required"`
	// List of public keys that can sign on behalf of the account pre-transaction
	PreSigners []string                                                                                         `json:"pre_signers,required"`
	Type       StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsOwnershipDiffType `json:"type"`
	JSON       stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsOwnershipDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsOwnershipDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsOwnershipDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsOwnershipDiffJSON struct {
	PostSigners apijson.Field
	PreSigners  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsOwnershipDiffType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsOwnershipDiffTypeSetOptions StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsOwnershipDiffType = "SET_OPTIONS"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsOwnershipDiffType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsOwnershipDiffTypeSetOptions:
		return true
	}
	return false
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeStatus string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeStatusSuccess StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeStatus = "Success"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeStatus) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeStatusSuccess:
		return true
	}
	return false
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAddressDetail struct {
	// Encoded public key of the account
	AccountAddress interface{} `json:"account_address,required"`
	// Description of the account
	Description string                                                                                     `json:"description,nullable"`
	JSON        stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAddressDetailJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAddressDetailJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAddressDetail]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAddressDetailJSON struct {
	AccountAddress apijson.Field
	Description    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAddressDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAddressDetailJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiff struct {
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of
	// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiffAsset],
	// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiffAsset],
	// [StellarAssetContractDetailsSchema].
	Asset interface{} `json:"asset"`
	// Details of the incoming transfer
	In StellarAssetTransferDetailsSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out   StellarAssetTransferDetailsSchema                                                       `json:"out,nullable"`
	JSON  stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffJSON `json:"-"`
	union StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsUnion
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffJSON struct {
	AssetType   apijson.Field
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiff],
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiff],
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarContractAssetDiff].
func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiff) AsUnion() StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiff],
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiff]
// or
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarContractAssetDiff].
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsUnion interface {
	implementsStellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarContractAssetDiff{}),
		},
	)
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiff struct {
	Asset StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StellarAssetTransferDetailsSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StellarAssetTransferDetailsSchema                                                                              `json:"out,nullable"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiff) implementsStellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiff() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiffAsset struct {
	// Asset code
	Code string `json:"code,required"`
	// Asset issuer address
	Issuer string `json:"issuer,required"`
	// Organization name
	OrgName string `json:"org_name,required"`
	// Organization URL
	OrgURL string `json:"org_url,required"`
	// Type of the asset (`ASSET`)
	Type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiffAssetType `json:"type"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiffAssetJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiffAsset]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiffAssetJSON struct {
	Code        apijson.Field
	Issuer      apijson.Field
	OrgName     apijson.Field
	OrgURL      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ASSET`)
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiffAssetType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiffAssetTypeAsset StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiffAssetType = "ASSET"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiffAssetType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarLegacyAssetDiffAssetTypeAsset:
		return true
	}
	return false
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiff struct {
	Asset StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StellarAssetTransferDetailsSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StellarAssetTransferDetailsSchema                                                                              `json:"out,nullable"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiff) implementsStellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiff() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiffAsset struct {
	// Asset code
	Code StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiffAssetCode `json:"code"`
	// Type of the asset (`NATIVE`)
	Type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiffAssetType `json:"type"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiffAssetJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiffAsset]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiffAssetJSON struct {
	Code        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Asset code
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiffAssetCode string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiffAssetCodeXlm StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiffAssetCode = "XLM"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiffAssetCode) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiffAssetCodeXlm:
		return true
	}
	return false
}

// Type of the asset (`NATIVE`)
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiffAssetType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiffAssetTypeNative StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiffAssetType = "NATIVE"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiffAssetType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarNativeAssetDiffAssetTypeNative:
		return true
	}
	return false
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarContractAssetDiff struct {
	Asset StellarAssetContractDetailsSchema `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StellarAssetTransferDetailsSchema `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StellarAssetTransferDetailsSchema                                                                                `json:"out,nullable"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarContractAssetDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarContractAssetDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarContractAssetDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarContractAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarContractAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarContractAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiffsStellarContractAssetDiff) implementsStellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeAssetsDiff() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposure struct {
	// This field can have the runtime type of
	// [map[string]StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureSpender],
	// [map[string]StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureSpender].
	Spenders interface{} `json:"spenders,required"`
	// This field can have the runtime type of
	// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureAsset],
	// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureAsset].
	Asset interface{}                                                                           `json:"asset"`
	JSON  stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposureJSON `json:"-"`
	union StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresUnion
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposureJSON struct {
	Spenders    apijson.Field
	Asset       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposureJSON) RawJSON() string {
	return r.raw
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposure) UnmarshalJSON(data []byte) (err error) {
	*r = StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposure],
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposure].
func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposure) AsUnion() StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresUnion {
	return r.union
}

// Union satisfied by
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposure]
// or
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposure].
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresUnion interface {
	implementsStellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposure{}),
		},
	)
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposure struct {
	Asset StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureAsset `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureSpender `json:"spenders"`
	JSON     stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureJSON               `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposure) implementsStellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposure() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureAsset struct {
	// Asset code
	Code string `json:"code,required"`
	// Asset issuer address
	Issuer string `json:"issuer,required"`
	// Organization name
	OrgName string `json:"org_name,required"`
	// Organization URL
	OrgURL string `json:"org_url,required"`
	// Type of the asset (`ASSET`)
	Type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureAssetType `json:"type"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureAssetJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureAssetJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureAsset]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureAssetJSON struct {
	Code        apijson.Field
	Issuer      apijson.Field
	OrgName     apijson.Field
	OrgURL      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ASSET`)
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureAssetType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureAssetTypeAsset StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureAssetType = "ASSET"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureAssetType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureAssetTypeAsset:
		return true
	}
	return false
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureSpender struct {
	// Approval value of the ERC20 token
	Approval string                                                                                                                         `json:"approval,required"`
	Exposure []StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureSpendersExposure `json:"exposure,required"`
	// Expiration date of the approval
	Expiration time.Time `json:"expiration,nullable" format:"date-time"`
	// Summarized description of the exposure
	Summary string                                                                                                                  `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureSpenderJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureSpenderJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureSpender]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureSpenderJSON struct {
	Approval    apijson.Field
	Exposure    apijson.Field
	Expiration  apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureSpendersExposure struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                           `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureSpendersExposureJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureSpendersExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureSpendersExposureJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarLegacyAssetExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposure struct {
	Asset StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureAsset `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureSpender `json:"spenders"`
	JSON     stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureJSON               `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposure) implementsStellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposure() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureAsset struct {
	// Asset code
	Code StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureAssetCode `json:"code"`
	// Type of the asset (`NATIVE`)
	Type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureAssetType `json:"type"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureAssetJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureAssetJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureAsset]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureAssetJSON struct {
	Code        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureAssetJSON) RawJSON() string {
	return r.raw
}

// Asset code
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureAssetCode string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureAssetCodeXlm StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureAssetCode = "XLM"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureAssetCode) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureAssetCodeXlm:
		return true
	}
	return false
}

// Type of the asset (`NATIVE`)
type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureAssetType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureAssetTypeNative StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureAssetType = "NATIVE"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureAssetType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureAssetTypeNative:
		return true
	}
	return false
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureSpender struct {
	// Approval value of the ERC20 token
	Approval string                                                                                                                         `json:"approval,required"`
	Exposure []StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureSpendersExposure `json:"exposure,required"`
	// Expiration date of the approval
	Expiration time.Time `json:"expiration,nullable" format:"date-time"`
	// Summarized description of the exposure
	Summary string                                                                                                                  `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureSpenderJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureSpenderJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureSpender]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureSpenderJSON struct {
	Approval    apijson.Field
	Exposure    apijson.Field
	Expiration  apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureSpendersExposure struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                           `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureSpendersExposureJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureSpendersExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureSpendersExposureJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResultSchemaTypeExposuresStellarNativeAssetExposureSpendersExposureJSON) RawJSON() string {
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
	// This field can have the runtime type of
	// [[]StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaFeature].
	Features interface{}                                    `json:"features,required"`
	Status   StellarTransactionScanResponseValidationStatus `json:"status,required"`
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification"`
	// A textual description about the validation result
	Description string `json:"description"`
	// Error message
	Error string `json:"error"`
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
	Features       apijson.Field
	Status         apijson.Field
	Classification apijson.Field
	Description    apijson.Field
	Error          apijson.Field
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
// [StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchema],
// [StellarTransactionScanResponseValidationStellarValidationErrorSchema].
func (r StellarTransactionScanResponseValidation) AsUnion() StellarTransactionScanResponseValidationUnion {
	return r.union
}

// Validation result; Only present if validation option is included in the request
//
// Union satisfied by
// [StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchema]
// or [StellarTransactionScanResponseValidationStellarValidationErrorSchema].
type StellarTransactionScanResponseValidationUnion interface {
	implementsStellarTransactionScanResponseValidation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StellarTransactionScanResponseValidationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseValidationStellarValidationErrorSchema{}),
		},
	)
}

type StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchema struct {
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification,required"`
	// A textual description about the validation result
	Description string                                                                                   `json:"description,required"`
	Features    []StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaFeature `json:"features,required"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason string `json:"reason,required"`
	// Verdict of the validation
	ResultType StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaResultType `json:"result_type,required"`
	Status     StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaStatus     `json:"status,required"`
	JSON       stellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaJSON       `json:"-"`
}

// stellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchema]
type stellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaJSON struct {
	Classification apijson.Field
	Description    apijson.Field
	Features       apijson.Field
	Reason         apijson.Field
	ResultType     apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchema) implementsStellarTransactionScanResponseValidation() {
}

type StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaFeature struct {
	// Address the feature refers to
	Address string `json:"address,required"`
	// Textual description
	Description string `json:"description,required"`
	FeatureID string `json:"feature_id,required"`
	// Feature Classification
	Type StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaFeaturesType `json:"type,required"`
	JSON stellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaFeatureJSON  `json:"-"`
}

// stellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaFeatureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaFeature]
type stellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaFeatureJSON struct {
	Address     apijson.Field
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaFeatureJSON) RawJSON() string {
	return r.raw
}

// Feature Classification
type StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaFeaturesType string

const (
	StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaFeaturesTypeBenign    StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaFeaturesType = "Benign"
	StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaFeaturesTypeWarning   StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaFeaturesType = "Warning"
	StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaFeaturesTypeMalicious StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaFeaturesType = "Malicious"
	StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaFeaturesTypeInfo      StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaFeaturesType = "Info"
)

func (r StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaFeaturesType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaFeaturesTypeBenign, StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaFeaturesTypeWarning, StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaFeaturesTypeMalicious, StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaFeaturesTypeInfo:
		return true
	}
	return false
}

// Verdict of the validation
type StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaResultType string

const (
	StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaResultTypeBenign    StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaResultType = "Benign"
	StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaResultTypeWarning   StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaResultType = "Warning"
	StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaResultTypeMalicious StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaResultType = "Malicious"
)

func (r StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaResultType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaResultTypeBenign, StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaResultTypeWarning, StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaResultTypeMalicious:
		return true
	}
	return false
}

type StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaStatus string

const (
	StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaStatusSuccess StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaStatus = "Success"
)

func (r StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaStatus) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseValidationStellarValidationSuccessfulResultSchemaStatusSuccess:
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
