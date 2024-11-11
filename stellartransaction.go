// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"net/http"
	"reflect"
	"time"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/param"
	"github.com/blockaid-official/blockaid-client-go/internal/requestconfig"
	"github.com/blockaid-official/blockaid-client-go/option"
	"github.com/tidwall/gjson"
)

// StellarTransactionService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStellarTransactionService] method instead.
type StellarTransactionService struct {
	Options []option.RequestOption
}

// NewStellarTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStellarTransactionService(opts ...option.RequestOption) (r *StellarTransactionService) {
	r = &StellarTransactionService{}
	r.Options = opts
	return
}

// Report Transaction
func (r *StellarTransactionService) Report(ctx context.Context, body StellarTransactionReportParams, opts ...option.RequestOption) (res *int64, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/stellar/transaction/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Scan Transaction
func (r *StellarTransactionService) Scan(ctx context.Context, body StellarTransactionScanParams, opts ...option.RequestOption) (res *StellarTransactionScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/stellar/transaction/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
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
	// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummary].
	AccountSummary interface{} `json:"account_summary,required"`
	// This field can have the runtime type of
	// [[]StellarTransactionScanResponseSimulationStellarSimulationResponseAddressDetail].
	AddressDetails interface{} `json:"address_details,required"`
	// This field can have the runtime type of
	// [map[string][]StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiff].
	AssetsDiffs interface{} `json:"assets_diffs,required"`
	// This field can have the runtime type of
	// [map[string][]StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsOwnershipDiff].
	AssetsOwnershipDiff interface{} `json:"assets_ownership_diff,required"`
	// This field can have the runtime type of
	// [map[string][]StellarTransactionScanResponseSimulationStellarSimulationResponseExposure].
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
// [StellarTransactionScanResponseSimulationStellarSimulationResponse],
// [StellarTransactionScanResponseSimulationStellarSimulationErrorSchema].
func (r StellarTransactionScanResponseSimulation) AsUnion() StellarTransactionScanResponseSimulationUnion {
	return r.union
}

// Simulation result; Only present if simulation option is included in the request
//
// Union satisfied by
// [StellarTransactionScanResponseSimulationStellarSimulationResponse] or
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
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationErrorSchema{}),
		},
	)
}

type StellarTransactionScanResponseSimulationStellarSimulationResponse struct {
	// Summary of the actions and asset transfers that were made by the requested
	// account address
	AccountSummary StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummary `json:"account_summary,required"`
	// Ownership diffs of the account addresses
	AssetsOwnershipDiff map[string][]StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsOwnershipDiff `json:"assets_ownership_diff,required"`
	Status              StellarTransactionScanResponseSimulationStellarSimulationResponseStatus                           `json:"status,required"`
	// Details of addresses involved in the transaction
	AddressDetails []StellarTransactionScanResponseSimulationStellarSimulationResponseAddressDetail `json:"address_details"`
	// Mapping between the address of an account to the assets diff during the
	// transaction
	AssetsDiffs map[string][]StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiff `json:"assets_diffs"`
	// Mapping between the address of an account to the exposure of the assets during
	// the transaction
	Exposures map[string][]StellarTransactionScanResponseSimulationStellarSimulationResponseExposure `json:"exposures"`
	JSON      stellarTransactionScanResponseSimulationStellarSimulationResponseJSON                  `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseJSON contains
// the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponse]
type stellarTransactionScanResponseSimulationStellarSimulationResponseJSON struct {
	AccountSummary      apijson.Field
	AssetsOwnershipDiff apijson.Field
	Status              apijson.Field
	AddressDetails      apijson.Field
	AssetsDiffs         apijson.Field
	Exposures           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResponse) implementsStellarTransactionScanResponseSimulation() {
}

// Summary of the actions and asset transfers that were made by the requested
// account address
type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummary struct {
	// Exposures made by the requested account address
	AccountExposures []StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposure `json:"account_exposures,required"`
	// Ownership diffs of the requested account address
	AccountOwnershipsDiff []StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountOwnershipsDiff `json:"account_ownerships_diff,required"`
	// Total USD diff for the requested account address
	TotalUsdDiff StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryTotalUsdDiff `json:"total_usd_diff,required"`
	// Assets diffs of the requested account address
	AccountAssetsDiffs []StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiff `json:"account_assets_diffs"`
	// Total USD exposure for each of the spender addresses during the transaction
	TotalUsdExposure map[string]float64                                                                  `json:"total_usd_exposure"`
	JSON             stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummary]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryJSON struct {
	AccountExposures      apijson.Field
	AccountOwnershipsDiff apijson.Field
	TotalUsdDiff          apijson.Field
	AccountAssetsDiffs    apijson.Field
	TotalUsdExposure      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposure struct {
	// This field can have the runtime type of
	// [map[string]StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureSpender],
	// [map[string]StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureSpender].
	Spenders interface{} `json:"spenders,required"`
	// This field can have the runtime type of
	// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureAsset],
	// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureAsset].
	Asset interface{}                                                                                        `json:"asset"`
	JSON  stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposureJSON `json:"-"`
	union StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresUnion
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposureJSON struct {
	Spenders    apijson.Field
	Asset       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposureJSON) RawJSON() string {
	return r.raw
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposure) UnmarshalJSON(data []byte) (err error) {
	*r = StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposure],
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposure].
func (r StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposure) AsUnion() StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresUnion {
	return r.union
}

// Union satisfied by
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposure]
// or
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposure].
type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresUnion interface {
	implementsStellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposure{}),
		},
	)
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposure struct {
	Asset StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureAsset `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureSpender `json:"spenders"`
	JSON     stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureJSON               `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposure) implementsStellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposure() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureAsset struct {
	// Asset code
	Code string `json:"code,required"`
	// Asset issuer address
	Issuer string `json:"issuer,required"`
	// Organization name
	OrgName string `json:"org_name,required"`
	// Organization URL
	OrgURL string `json:"org_url,required"`
	// Type of the asset (`ASSET`)
	Type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureAssetType `json:"type"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureAssetJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureAssetJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureAsset]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureAssetJSON struct {
	Code        apijson.Field
	Issuer      apijson.Field
	OrgName     apijson.Field
	OrgURL      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ASSET`)
type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureAssetType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureAssetTypeAsset StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureAssetType = "ASSET"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureAssetType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureAssetTypeAsset:
		return true
	}
	return false
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureSpender struct {
	// Approval value of the ERC20 token
	Approval string                                                                                                                                      `json:"approval,required"`
	Exposure []StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureSpendersExposure `json:"exposure,required"`
	// Expiration date of the approval
	Expiration time.Time `json:"expiration,nullable" format:"date-time"`
	// Summarized description of the exposure
	Summary string                                                                                                                               `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureSpenderJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureSpenderJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureSpender]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureSpenderJSON struct {
	Approval    apijson.Field
	Exposure    apijson.Field
	Expiration  apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureSpendersExposure struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                                        `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureSpendersExposureJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureSpendersExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureSpendersExposureJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarLegacyAssetExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposure struct {
	Asset StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureAsset `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureSpender `json:"spenders"`
	JSON     stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureJSON               `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposure) implementsStellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposure() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureAsset struct {
	// Asset code
	Code StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureAssetCode `json:"code"`
	// Type of the asset (`NATIVE`)
	Type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureAssetType `json:"type"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureAssetJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureAssetJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureAsset]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureAssetJSON struct {
	Code        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureAssetJSON) RawJSON() string {
	return r.raw
}

// Asset code
type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureAssetCode string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureAssetCodeXlm StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureAssetCode = "XLM"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureAssetCode) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureAssetCodeXlm:
		return true
	}
	return false
}

// Type of the asset (`NATIVE`)
type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureAssetType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureAssetTypeNative StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureAssetType = "NATIVE"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureAssetType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureAssetTypeNative:
		return true
	}
	return false
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureSpender struct {
	// Approval value of the ERC20 token
	Approval string                                                                                                                                      `json:"approval,required"`
	Exposure []StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureSpendersExposure `json:"exposure,required"`
	// Expiration date of the approval
	Expiration time.Time `json:"expiration,nullable" format:"date-time"`
	// Summarized description of the exposure
	Summary string                                                                                                                               `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureSpenderJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureSpenderJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureSpender]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureSpenderJSON struct {
	Approval    apijson.Field
	Exposure    apijson.Field
	Expiration  apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureSpendersExposure struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                                        `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureSpendersExposureJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureSpendersExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureSpendersExposureJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountExposuresStellarNativeAssetExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountOwnershipsDiff struct {
	// List of public keys that can sign on behalf of the account post-transaction
	PostSigners []string `json:"post_signers,required"`
	// List of public keys that can sign on behalf of the account pre-transaction
	PreSigners []string                                                                                                 `json:"pre_signers,required"`
	Type       StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountOwnershipsDiffType `json:"type"`
	JSON       stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountOwnershipsDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountOwnershipsDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountOwnershipsDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountOwnershipsDiffJSON struct {
	PostSigners apijson.Field
	PreSigners  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountOwnershipsDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountOwnershipsDiffJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountOwnershipsDiffType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountOwnershipsDiffTypeSetOptions StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountOwnershipsDiffType = "SET_OPTIONS"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountOwnershipsDiffType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountOwnershipsDiffTypeSetOptions:
		return true
	}
	return false
}

// Total USD diff for the requested account address
type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryTotalUsdDiff struct {
	// Total incoming USD transfers
	In float64 `json:"in,required"`
	// Total outgoing USD transfers
	Out float64 `json:"out,required"`
	// Total USD transfers
	Total float64                                                                                         `json:"total"`
	JSON  stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryTotalUsdDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryTotalUsdDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryTotalUsdDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiff struct {
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of
	// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffIn],
	// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffIn],
	// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffIn].
	In interface{} `json:"in,required"`
	// This field can have the runtime type of
	// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffOut],
	// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffOut],
	// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffOut].
	Out interface{} `json:"out,required"`
	// This field can have the runtime type of
	// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAsset],
	// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAsset],
	// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffAsset].
	Asset interface{}                                                                                          `json:"asset"`
	JSON  stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffJSON `json:"-"`
	union StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsUnion
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffJSON struct {
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	Asset       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiff],
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiff],
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiff].
func (r StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiff) AsUnion() StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiff],
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiff]
// or
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiff].
type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsUnion interface {
	implementsStellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiff{}),
		},
	)
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiff struct {
	Asset StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffOut  `json:"out,nullable"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiff) implementsStellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiff() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAsset struct {
	// Asset code
	Code string `json:"code,required"`
	// Asset issuer address
	Issuer string `json:"issuer,required"`
	// Organization name
	OrgName string `json:"org_name,required"`
	// Organization URL
	OrgURL string `json:"org_url,required"`
	// Type of the asset (`ASSET`)
	Type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAssetType `json:"type"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAssetJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAsset]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAssetJSON struct {
	Code        apijson.Field
	Issuer      apijson.Field
	OrgName     apijson.Field
	OrgURL      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ASSET`)
type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAssetType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAssetTypeAsset StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAssetType = "ASSET"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAssetType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffAssetTypeAsset:
		return true
	}
	return false
}

// Details of the incoming transfer
type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                        `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffInJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffInJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffIn]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffInJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                         `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffOutJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffOutJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffOut]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffOutJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarLegacyAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiff struct {
	Asset StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffOut  `json:"out,nullable"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiff) implementsStellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiff() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAsset struct {
	// Asset code
	Code StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetCode `json:"code"`
	// Type of the asset (`NATIVE`)
	Type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetType `json:"type"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAsset]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetJSON struct {
	Code        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Asset code
type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetCode string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetCodeXlm StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetCode = "XLM"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetCode) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetCodeXlm:
		return true
	}
	return false
}

// Type of the asset (`NATIVE`)
type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetTypeNative StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetType = "NATIVE"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffAssetTypeNative:
		return true
	}
	return false
}

// Details of the incoming transfer
type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                        `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffInJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffInJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffIn]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffInJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                         `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffOutJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffOutJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffOut]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffOutJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarNativeAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiff struct {
	Asset StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffOut  `json:"out,nullable"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiff) implementsStellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiff() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffAsset struct {
	// Address of the asset's contract
	Address string `json:"address,required"`
	// Asset code
	Name string `json:"name,required"`
	// Asset symbol
	Symbol string `json:"symbol,required"`
	// Type of the asset (`CONTRACT`)
	Type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffAssetType `json:"type"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffAssetJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffAsset]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`CONTRACT`)
type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffAssetType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffAssetTypeContract StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffAssetType = "CONTRACT"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffAssetType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffAssetTypeContract:
		return true
	}
	return false
}

// Details of the incoming transfer
type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                          `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffInJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffInJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffIn]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffInJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                           `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffOutJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffOutJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffOut]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffOutJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAccountSummaryAccountAssetsDiffsStellarContractAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsOwnershipDiff struct {
	// List of public keys that can sign on behalf of the account post-transaction
	PostSigners []string `json:"post_signers,required"`
	// List of public keys that can sign on behalf of the account pre-transaction
	PreSigners []string                                                                                 `json:"pre_signers,required"`
	Type       StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsOwnershipDiffType `json:"type"`
	JSON       stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsOwnershipDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsOwnershipDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsOwnershipDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsOwnershipDiffJSON struct {
	PostSigners apijson.Field
	PreSigners  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsOwnershipDiffType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsOwnershipDiffTypeSetOptions StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsOwnershipDiffType = "SET_OPTIONS"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsOwnershipDiffType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsOwnershipDiffTypeSetOptions:
		return true
	}
	return false
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseStatus string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResponseStatusSuccess StellarTransactionScanResponseSimulationStellarSimulationResponseStatus = "Success"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseStatus) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResponseStatusSuccess:
		return true
	}
	return false
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAddressDetail struct {
	// Encoded public key of the account
	AccountAddress interface{} `json:"account_address,required"`
	// Description of the account
	Description string                                                                             `json:"description,nullable"`
	JSON        stellarTransactionScanResponseSimulationStellarSimulationResponseAddressDetailJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAddressDetailJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAddressDetail]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAddressDetailJSON struct {
	AccountAddress apijson.Field
	Description    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAddressDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAddressDetailJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiff struct {
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of
	// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffIn],
	// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffIn],
	// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffIn].
	In interface{} `json:"in,required"`
	// This field can have the runtime type of
	// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffOut],
	// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffOut],
	// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffOut].
	Out interface{} `json:"out,required"`
	// This field can have the runtime type of
	// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffAsset],
	// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffAsset],
	// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffAsset].
	Asset interface{}                                                                     `json:"asset"`
	JSON  stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffJSON `json:"-"`
	union StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsUnion
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffJSON struct {
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	Asset       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiff],
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiff],
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiff].
func (r StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiff) AsUnion() StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiff],
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiff]
// or
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiff].
type StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsUnion interface {
	implementsStellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiff{}),
		},
	)
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiff struct {
	Asset StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffOut  `json:"out,nullable"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiff) implementsStellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiff() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffAsset struct {
	// Asset code
	Code string `json:"code,required"`
	// Asset issuer address
	Issuer string `json:"issuer,required"`
	// Organization name
	OrgName string `json:"org_name,required"`
	// Organization URL
	OrgURL string `json:"org_url,required"`
	// Type of the asset (`ASSET`)
	Type StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffAssetType `json:"type"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffAssetJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffAsset]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffAssetJSON struct {
	Code        apijson.Field
	Issuer      apijson.Field
	OrgName     apijson.Field
	OrgURL      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ASSET`)
type StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffAssetType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffAssetTypeAsset StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffAssetType = "ASSET"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffAssetType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffAssetTypeAsset:
		return true
	}
	return false
}

// Details of the incoming transfer
type StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                   `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffInJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffInJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffIn]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffInJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                    `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffOutJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffOutJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffOut]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffOutJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarLegacyAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiff struct {
	Asset StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffOut  `json:"out,nullable"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiff) implementsStellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiff() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffAsset struct {
	// Asset code
	Code StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffAssetCode `json:"code"`
	// Type of the asset (`NATIVE`)
	Type StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffAssetType `json:"type"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffAssetJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffAsset]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffAssetJSON struct {
	Code        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Asset code
type StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffAssetCode string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffAssetCodeXlm StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffAssetCode = "XLM"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffAssetCode) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffAssetCodeXlm:
		return true
	}
	return false
}

// Type of the asset (`NATIVE`)
type StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffAssetType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffAssetTypeNative StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffAssetType = "NATIVE"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffAssetType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffAssetTypeNative:
		return true
	}
	return false
}

// Details of the incoming transfer
type StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                   `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffInJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffInJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffIn]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffInJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                    `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffOutJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffOutJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffOut]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffOutJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarNativeAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiff struct {
	Asset StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffOut  `json:"out,nullable"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiff]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiff) implementsStellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiff() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffAsset struct {
	// Address of the asset's contract
	Address string `json:"address,required"`
	// Asset code
	Name string `json:"name,required"`
	// Asset symbol
	Symbol string `json:"symbol,required"`
	// Type of the asset (`CONTRACT`)
	Type StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffAssetType `json:"type"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffAssetJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffAsset]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`CONTRACT`)
type StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffAssetType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffAssetTypeContract StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffAssetType = "CONTRACT"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffAssetType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffAssetTypeContract:
		return true
	}
	return false
}

// Details of the incoming transfer
type StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                     `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffInJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffInJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffIn]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffInJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                      `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffOutJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffOutJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffOut]
type stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffOutJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseAssetsDiffsStellarContractAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseExposure struct {
	// This field can have the runtime type of
	// [map[string]StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureSpender],
	// [map[string]StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureSpender].
	Spenders interface{} `json:"spenders,required"`
	// This field can have the runtime type of
	// [StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureAsset],
	// [StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureAsset].
	Asset interface{}                                                                   `json:"asset"`
	JSON  stellarTransactionScanResponseSimulationStellarSimulationResponseExposureJSON `json:"-"`
	union StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresUnion
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResponseExposureJSON struct {
	Spenders    apijson.Field
	Asset       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseExposureJSON) RawJSON() string {
	return r.raw
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseExposure) UnmarshalJSON(data []byte) (err error) {
	*r = StellarTransactionScanResponseSimulationStellarSimulationResponseExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposure],
// [StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposure].
func (r StellarTransactionScanResponseSimulationStellarSimulationResponseExposure) AsUnion() StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresUnion {
	return r.union
}

// Union satisfied by
// [StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposure]
// or
// [StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposure].
type StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresUnion interface {
	implementsStellarTransactionScanResponseSimulationStellarSimulationResponseExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposure{}),
		},
	)
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposure struct {
	Asset StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureAsset `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureSpender `json:"spenders"`
	JSON     stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureJSON               `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposure) implementsStellarTransactionScanResponseSimulationStellarSimulationResponseExposure() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureAsset struct {
	// Asset code
	Code string `json:"code,required"`
	// Asset issuer address
	Issuer string `json:"issuer,required"`
	// Organization name
	OrgName string `json:"org_name,required"`
	// Organization URL
	OrgURL string `json:"org_url,required"`
	// Type of the asset (`ASSET`)
	Type StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureAssetType `json:"type"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureAssetJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureAssetJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureAsset]
type stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureAssetJSON struct {
	Code        apijson.Field
	Issuer      apijson.Field
	OrgName     apijson.Field
	OrgURL      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ASSET`)
type StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureAssetType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureAssetTypeAsset StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureAssetType = "ASSET"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureAssetType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureAssetTypeAsset:
		return true
	}
	return false
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureSpender struct {
	// Approval value of the ERC20 token
	Approval string                                                                                                                 `json:"approval,required"`
	Exposure []StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureSpendersExposure `json:"exposure,required"`
	// Expiration date of the approval
	Expiration time.Time `json:"expiration,nullable" format:"date-time"`
	// Summarized description of the exposure
	Summary string                                                                                                          `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureSpenderJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureSpenderJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureSpender]
type stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureSpenderJSON struct {
	Approval    apijson.Field
	Exposure    apijson.Field
	Expiration  apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureSpendersExposure struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                   `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureSpendersExposureJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureSpendersExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureSpendersExposureJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarLegacyAssetExposureSpendersExposureJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposure struct {
	Asset StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureAsset `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureSpender `json:"spenders"`
	JSON     stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureJSON               `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureJSON) RawJSON() string {
	return r.raw
}

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposure) implementsStellarTransactionScanResponseSimulationStellarSimulationResponseExposure() {
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureAsset struct {
	// Asset code
	Code StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureAssetCode `json:"code"`
	// Type of the asset (`NATIVE`)
	Type StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureAssetType `json:"type"`
	JSON stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureAssetJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureAssetJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureAsset]
type stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureAssetJSON struct {
	Code        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureAssetJSON) RawJSON() string {
	return r.raw
}

// Asset code
type StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureAssetCode string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureAssetCodeXlm StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureAssetCode = "XLM"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureAssetCode) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureAssetCodeXlm:
		return true
	}
	return false
}

// Type of the asset (`NATIVE`)
type StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureAssetType string

const (
	StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureAssetTypeNative StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureAssetType = "NATIVE"
)

func (r StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureAssetType) IsKnown() bool {
	switch r {
	case StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureAssetTypeNative:
		return true
	}
	return false
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureSpender struct {
	// Approval value of the ERC20 token
	Approval string                                                                                                                 `json:"approval,required"`
	Exposure []StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureSpendersExposure `json:"exposure,required"`
	// Expiration date of the approval
	Expiration time.Time `json:"expiration,nullable" format:"date-time"`
	// Summarized description of the exposure
	Summary string                                                                                                          `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureSpenderJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureSpenderJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureSpender]
type stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureSpenderJSON struct {
	Approval    apijson.Field
	Exposure    apijson.Field
	Expiration  apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureSpendersExposure struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                   `json:"summary,nullable"`
	JSON    stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureSpendersExposureJSON `json:"-"`
}

// stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureSpendersExposureJSON
// contains the JSON metadata for the struct
// [StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureSpendersExposure]
type stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureSpendersExposureJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureSpendersExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarTransactionScanResponseSimulationStellarSimulationResponseExposuresStellarNativeAssetExposureSpendersExposureJSON) RawJSON() string {
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
	// [[]StellarTransactionScanResponseValidationStellarValidationResultFeature].
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
	FeatureID string `json:"feature_id,required"`
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

type StellarTransactionReportParams struct {
	Details param.Field[string]                                    `json:"details,required"`
	Event   param.Field[StellarTransactionReportParamsEvent]       `json:"event,required"`
	Report  param.Field[StellarTransactionReportParamsReportUnion] `json:"report,required"`
}

func (r StellarTransactionReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StellarTransactionReportParamsEvent string

const (
	StellarTransactionReportParamsEventShouldBeMalicious StellarTransactionReportParamsEvent = "should_be_malicious"
	StellarTransactionReportParamsEventShouldBeBenign    StellarTransactionReportParamsEvent = "should_be_benign"
)

func (r StellarTransactionReportParamsEvent) IsKnown() bool {
	switch r {
	case StellarTransactionReportParamsEventShouldBeMalicious, StellarTransactionReportParamsEventShouldBeBenign:
		return true
	}
	return false
}

type StellarTransactionReportParamsReport struct {
	Params param.Field[interface{}]                              `json:"params,required"`
	ID     param.Field[string]                                   `json:"id"`
	Type   param.Field[StellarTransactionReportParamsReportType] `json:"type"`
}

func (r StellarTransactionReportParamsReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionReportParamsReport) implementsStellarTransactionReportParamsReportUnion() {}

// Satisfied by [StellarTransactionReportParamsReportStellarAppealRequestID],
// [StellarTransactionReportParamsReportStellarAppealTransactionDataReport],
// [StellarTransactionReportParamsReport].
type StellarTransactionReportParamsReportUnion interface {
	implementsStellarTransactionReportParamsReportUnion()
}

type StellarTransactionReportParamsReportStellarAppealRequestID struct {
	ID   param.Field[string]                                                         `json:"id,required"`
	Type param.Field[StellarTransactionReportParamsReportStellarAppealRequestIDType] `json:"type"`
}

func (r StellarTransactionReportParamsReportStellarAppealRequestID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionReportParamsReportStellarAppealRequestID) implementsStellarTransactionReportParamsReportUnion() {
}

type StellarTransactionReportParamsReportStellarAppealRequestIDType string

const (
	StellarTransactionReportParamsReportStellarAppealRequestIDTypeRequestID StellarTransactionReportParamsReportStellarAppealRequestIDType = "request_id"
)

func (r StellarTransactionReportParamsReportStellarAppealRequestIDType) IsKnown() bool {
	switch r {
	case StellarTransactionReportParamsReportStellarAppealRequestIDTypeRequestID:
		return true
	}
	return false
}

type StellarTransactionReportParamsReportStellarAppealTransactionDataReport struct {
	Params param.Field[StellarTransactionReportParamsReportStellarAppealTransactionDataReportParams] `json:"params,required"`
	Type   param.Field[StellarTransactionReportParamsReportStellarAppealTransactionDataReportType]   `json:"type"`
}

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReport) implementsStellarTransactionReportParamsReportUnion() {
}

type StellarTransactionReportParamsReportStellarAppealTransactionDataReportParams struct {
	AccountAddress param.Field[string] `json:"account_address,required"`
	// A CAIP-2 chain ID or a Stellar network name
	Chain param.Field[StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsChain] `json:"chain,required"`
	// Metadata
	Metadata    param.Field[StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataUnion] `json:"metadata,required"`
	Transaction param.Field[string]                                                                                    `json:"transaction,required"`
	// List of options to include in the response
	//
	// - `Options.validation`: Include Options.validation output in the response
	//
	// - `Options.simulation`: Include Options.simulation output in the response
	Options param.Field[[]StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsOption] `json:"options"`
}

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A CAIP-2 chain ID or a Stellar network name
type StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsChain string

const (
	StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsChainPubnet    StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsChain = "pubnet"
	StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsChainFuturenet StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsChain = "futurenet"
	StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsChainTestnet   StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsChain = "testnet"
)

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsChain) IsKnown() bool {
	switch r {
	case StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsChainPubnet, StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsChainFuturenet, StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsChainTestnet:
		return true
	}
	return false
}

// Metadata
type StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadata struct {
	// Metadata for wallet requests
	Type param.Field[StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataType] `json:"type"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url"`
}

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadata) implementsStellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataUnion() {
}

// Metadata
//
// Satisfied by
// [StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarWalletRequestMetadata],
// [StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarInAppRequestMetadata],
// [StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadata].
type StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataUnion interface {
	implementsStellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataUnion()
}

type StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarWalletRequestMetadata struct {
	// Metadata for wallet requests
	Type param.Field[StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarWalletRequestMetadataType] `json:"type,required"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url,required"`
}

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarWalletRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarWalletRequestMetadata) implementsStellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataUnion() {
}

// Metadata for wallet requests
type StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarWalletRequestMetadataType string

const (
	StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarWalletRequestMetadataTypeWallet StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarWalletRequestMetadataType = "wallet"
)

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarWalletRequestMetadataType) IsKnown() bool {
	switch r {
	case StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarWalletRequestMetadataTypeWallet:
		return true
	}
	return false
}

type StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarInAppRequestMetadata struct {
	// Metadata for in-app requests
	Type param.Field[StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarInAppRequestMetadataType] `json:"type"`
}

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarInAppRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarInAppRequestMetadata) implementsStellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataUnion() {
}

// Metadata for in-app requests
type StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarInAppRequestMetadataType string

const (
	StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarInAppRequestMetadataTypeInApp StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarInAppRequestMetadataType = "in_app"
)

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarInAppRequestMetadataType) IsKnown() bool {
	switch r {
	case StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarInAppRequestMetadataTypeInApp:
		return true
	}
	return false
}

// Metadata for wallet requests
type StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataType string

const (
	StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataTypeWallet StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataType = "wallet"
	StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataTypeInApp  StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataType = "in_app"
)

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataType) IsKnown() bool {
	switch r {
	case StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataTypeWallet, StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataTypeInApp:
		return true
	}
	return false
}

type StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsOption string

const (
	StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsOptionValidation StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsOption = "validation"
	StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsOptionSimulation StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsOption = "simulation"
)

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsOption) IsKnown() bool {
	switch r {
	case StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsOptionValidation, StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsOptionSimulation:
		return true
	}
	return false
}

type StellarTransactionReportParamsReportStellarAppealTransactionDataReportType string

const (
	StellarTransactionReportParamsReportStellarAppealTransactionDataReportTypeParams StellarTransactionReportParamsReportStellarAppealTransactionDataReportType = "params"
)

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportType) IsKnown() bool {
	switch r {
	case StellarTransactionReportParamsReportStellarAppealTransactionDataReportTypeParams:
		return true
	}
	return false
}

type StellarTransactionReportParamsReportType string

const (
	StellarTransactionReportParamsReportTypeRequestID StellarTransactionReportParamsReportType = "request_id"
	StellarTransactionReportParamsReportTypeParams    StellarTransactionReportParamsReportType = "params"
)

func (r StellarTransactionReportParamsReportType) IsKnown() bool {
	switch r {
	case StellarTransactionReportParamsReportTypeRequestID, StellarTransactionReportParamsReportTypeParams:
		return true
	}
	return false
}

type StellarTransactionScanParams struct {
	AccountAddress param.Field[string] `json:"account_address,required"`
	// A CAIP-2 chain ID or a Stellar network name
	Chain param.Field[StellarTransactionScanParamsChain] `json:"chain,required"`
	// Metadata
	Metadata    param.Field[StellarTransactionScanParamsMetadataUnion] `json:"metadata,required"`
	Transaction param.Field[string]                                    `json:"transaction,required"`
	// List of options to include in the response
	//
	// - `Options.validation`: Include Options.validation output in the response
	//
	// - `Options.simulation`: Include Options.simulation output in the response
	Options param.Field[[]StellarTransactionScanParamsOption] `json:"options"`
}

func (r StellarTransactionScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A CAIP-2 chain ID or a Stellar network name
type StellarTransactionScanParamsChain string

const (
	StellarTransactionScanParamsChainPubnet    StellarTransactionScanParamsChain = "pubnet"
	StellarTransactionScanParamsChainFuturenet StellarTransactionScanParamsChain = "futurenet"
	StellarTransactionScanParamsChainTestnet   StellarTransactionScanParamsChain = "testnet"
)

func (r StellarTransactionScanParamsChain) IsKnown() bool {
	switch r {
	case StellarTransactionScanParamsChainPubnet, StellarTransactionScanParamsChainFuturenet, StellarTransactionScanParamsChainTestnet:
		return true
	}
	return false
}

// Metadata
type StellarTransactionScanParamsMetadata struct {
	// Metadata for wallet requests
	Type param.Field[StellarTransactionScanParamsMetadataType] `json:"type"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url"`
}

func (r StellarTransactionScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionScanParamsMetadata) implementsStellarTransactionScanParamsMetadataUnion() {}

// Metadata
//
// Satisfied by [StellarTransactionScanParamsMetadataStellarWalletRequestMetadata],
// [StellarTransactionScanParamsMetadataStellarInAppRequestMetadata],
// [StellarTransactionScanParamsMetadata].
type StellarTransactionScanParamsMetadataUnion interface {
	implementsStellarTransactionScanParamsMetadataUnion()
}

type StellarTransactionScanParamsMetadataStellarWalletRequestMetadata struct {
	// Metadata for wallet requests
	Type param.Field[StellarTransactionScanParamsMetadataStellarWalletRequestMetadataType] `json:"type,required"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url,required"`
}

func (r StellarTransactionScanParamsMetadataStellarWalletRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionScanParamsMetadataStellarWalletRequestMetadata) implementsStellarTransactionScanParamsMetadataUnion() {
}

// Metadata for wallet requests
type StellarTransactionScanParamsMetadataStellarWalletRequestMetadataType string

const (
	StellarTransactionScanParamsMetadataStellarWalletRequestMetadataTypeWallet StellarTransactionScanParamsMetadataStellarWalletRequestMetadataType = "wallet"
)

func (r StellarTransactionScanParamsMetadataStellarWalletRequestMetadataType) IsKnown() bool {
	switch r {
	case StellarTransactionScanParamsMetadataStellarWalletRequestMetadataTypeWallet:
		return true
	}
	return false
}

type StellarTransactionScanParamsMetadataStellarInAppRequestMetadata struct {
	// Metadata for in-app requests
	Type param.Field[StellarTransactionScanParamsMetadataStellarInAppRequestMetadataType] `json:"type"`
}

func (r StellarTransactionScanParamsMetadataStellarInAppRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionScanParamsMetadataStellarInAppRequestMetadata) implementsStellarTransactionScanParamsMetadataUnion() {
}

// Metadata for in-app requests
type StellarTransactionScanParamsMetadataStellarInAppRequestMetadataType string

const (
	StellarTransactionScanParamsMetadataStellarInAppRequestMetadataTypeInApp StellarTransactionScanParamsMetadataStellarInAppRequestMetadataType = "in_app"
)

func (r StellarTransactionScanParamsMetadataStellarInAppRequestMetadataType) IsKnown() bool {
	switch r {
	case StellarTransactionScanParamsMetadataStellarInAppRequestMetadataTypeInApp:
		return true
	}
	return false
}

// Metadata for wallet requests
type StellarTransactionScanParamsMetadataType string

const (
	StellarTransactionScanParamsMetadataTypeWallet StellarTransactionScanParamsMetadataType = "wallet"
	StellarTransactionScanParamsMetadataTypeInApp  StellarTransactionScanParamsMetadataType = "in_app"
)

func (r StellarTransactionScanParamsMetadataType) IsKnown() bool {
	switch r {
	case StellarTransactionScanParamsMetadataTypeWallet, StellarTransactionScanParamsMetadataTypeInApp:
		return true
	}
	return false
}

type StellarTransactionScanParamsOption string

const (
	StellarTransactionScanParamsOptionValidation StellarTransactionScanParamsOption = "validation"
	StellarTransactionScanParamsOptionSimulation StellarTransactionScanParamsOption = "simulation"
)

func (r StellarTransactionScanParamsOption) IsKnown() bool {
	switch r {
	case StellarTransactionScanParamsOptionValidation, StellarTransactionScanParamsOptionSimulation:
		return true
	}
	return false
}
