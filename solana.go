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
	Simulation CombinedValidationResultSimulation `json:"simulation,nullable"`
	JSON       combinedValidationResultJSON       `json:"-"`
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

// Transaction simulation result
type CombinedValidationResultSimulation struct {
	// Summary of the requested account address
	AccountSummary CombinedValidationResultSimulationAccountSummary `json:"account_summary,required"`
	// Summary of the accounts involved in the transaction simulation
	AccountsDetails []CombinedValidationResultSimulationAccountsDetail `json:"accounts_details,required"`
	// Summary of the assets involved in the transaction simulation
	AssetsDiff map[string][]CombinedValidationResultSimulationAssetsDiff `json:"assets_diff,required"`
	// Summary of ownership changes; By account address
	AssetsOwnershipDiff map[string][]CombinedValidationResultSimulationAssetsOwnershipDiff `json:"assets_ownership_diff,required"`
	// Summary of the delegations, by account address
	Delegations map[string][]CombinedValidationResultSimulationDelegation `json:"delegations,required"`
	JSON        combinedValidationResultSimulationJSON                    `json:"-"`
}

// combinedValidationResultSimulationJSON contains the JSON metadata for the struct
// [CombinedValidationResultSimulation]
type combinedValidationResultSimulationJSON struct {
	AccountSummary      apijson.Field
	AccountsDetails     apijson.Field
	AssetsDiff          apijson.Field
	AssetsOwnershipDiff apijson.Field
	Delegations         apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CombinedValidationResultSimulation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationJSON) RawJSON() string {
	return r.raw
}

// Summary of the requested account address
type CombinedValidationResultSimulationAccountSummary struct {
	// Total USD diff for the requested account address
	TotalUsdDiff CombinedValidationResultSimulationAccountSummaryTotalUsdDiff `json:"total_usd_diff,required"`
	// Assets diff of the requested account address
	AccountAssetsDiff []CombinedValidationResultSimulationAccountSummaryAccountAssetsDiff `json:"account_assets_diff"`
	// Delegated assets of the requested account address
	AccountDelegations []CombinedValidationResultSimulationAccountSummaryAccountDelegation `json:"account_delegations"`
	// Account ownerships diff of the requested account address
	AccountOwnershipsDiff []CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiff `json:"account_ownerships_diff"`
	JSON                  combinedValidationResultSimulationAccountSummaryJSON                    `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryJSON contains the JSON metadata
// for the struct [CombinedValidationResultSimulationAccountSummary]
type combinedValidationResultSimulationAccountSummaryJSON struct {
	TotalUsdDiff          apijson.Field
	AccountAssetsDiff     apijson.Field
	AccountDelegations    apijson.Field
	AccountOwnershipsDiff apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryJSON) RawJSON() string {
	return r.raw
}

// Total USD diff for the requested account address
type CombinedValidationResultSimulationAccountSummaryTotalUsdDiff struct {
	// Total incoming USD transfers
	In float64 `json:"in,required"`
	// Total outgoing USD transfers
	Out float64 `json:"out,required"`
	// Total USD transfers
	Total float64                                                          `json:"total,required"`
	JSON  combinedValidationResultSimulationAccountSummaryTotalUsdDiffJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryTotalUsdDiffJSON contains the
// JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryTotalUsdDiff]
type combinedValidationResultSimulationAccountSummaryTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationAccountSummaryAccountAssetsDiff struct {
	// This field can have the runtime type of
	// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaAsset],
	// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaAsset],
	// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaAsset],
	// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaAsset].
	Asset interface{} `json:"asset"`
	// This field can have the runtime type of
	// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaIn],
	// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaIn],
	// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaIn],
	// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaIn].
	In interface{} `json:"in,required"`
	// This field can have the runtime type of
	// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaOut],
	// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaOut],
	// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaOut],
	// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaOut].
	Out   interface{}                                                           `json:"out,required"`
	JSON  combinedValidationResultSimulationAccountSummaryAccountAssetsDiffJSON `json:"-"`
	union CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffUnion
}

// combinedValidationResultSimulationAccountSummaryAccountAssetsDiffJSON contains
// the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiff]
type combinedValidationResultSimulationAccountSummaryAccountAssetsDiffJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r combinedValidationResultSimulationAccountSummaryAccountAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchema],
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchema],
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchema],
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchema].
func (r CombinedValidationResultSimulationAccountSummaryAccountAssetsDiff) AsUnion() CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffUnion {
	return r.union
}

// Union satisfied by
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchema],
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchema],
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchema]
// or
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchema].
type CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffUnion interface {
	implementsCombinedValidationResultSimulationAccountSummaryAccountAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchema{}),
		},
	)
}

type CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchema struct {
	Asset CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaAsset `json:"asset,required"`
	// Incoming transfers of the asset
	In   CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaIn   `json:"in,nullable"`
	Out  CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaOut  `json:"out,nullable"`
	JSON combinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchema]
type combinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchema) implementsCombinedValidationResultSimulationAccountSummaryAccountAssetsDiff() {
}

type CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaAsset struct {
	Decimals int64 `json:"decimals"`
	// Type of the asset (`"SOL"`)
	Type string                                                                                        `json:"type"`
	JSON combinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaAssetJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaAsset]
type combinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Incoming transfers of the asset
type CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                    `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaInJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaInJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaIn]
type combinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                     `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaOutJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaOutJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaOut]
type combinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchema struct {
	Asset CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaAsset `json:"asset,required"`
	// Incoming transfers of the asset
	In   CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaIn   `json:"in,nullable"`
	Out  CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaOut  `json:"out,nullable"`
	JSON combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchema]
type combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchema) implementsCombinedValidationResultSimulationAccountSummaryAccountAssetsDiff() {
}

type CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaAsset struct {
	Address  string `json:"address,required"`
	Decimals int64  `json:"decimals,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Logo     string `json:"logo"`
	// Type of the asset (`"TOKEN"`)
	Type string                                                                                               `json:"type"`
	JSON combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaAssetJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaAsset]
type combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Incoming transfers of the asset
type CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                           `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaInJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaInJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaIn]
type combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                            `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaOutJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaOutJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaOut]
type combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchema struct {
	Asset CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaAsset `json:"asset,required"`
	// Incoming transfers of the asset
	In   CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaIn   `json:"in,nullable"`
	Out  CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaOut  `json:"out,nullable"`
	JSON combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchema]
type combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchema) implementsCombinedValidationResultSimulationAccountSummaryAccountAssetsDiff() {
}

type CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaAsset struct {
	Address  string `json:"address,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	Logo     string `json:"logo"`
	// Type of the asset (`"NFT"`)
	Type string                                                                                                  `json:"type"`
	JSON combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaAssetJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaAsset]
type combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Incoming transfers of the asset
type CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                              `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaInJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaInJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaIn]
type combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                               `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaOutJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaOutJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaOut]
type combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchema struct {
	Asset CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaAsset `json:"asset,required"`
	// Incoming transfers of the asset
	In   CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaIn   `json:"in,nullable"`
	Out  CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaOut  `json:"out,nullable"`
	JSON combinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchema]
type combinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchema) implementsCombinedValidationResultSimulationAccountSummaryAccountAssetsDiff() {
}

type CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaAsset struct {
	Address  string `json:"address,required"`
	Decimals int64  `json:"decimals,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Logo     string `json:"logo"`
	// Type of the asset (`"CNFT"`)
	Type string                                                                                   `json:"type"`
	JSON combinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaAssetJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaAsset]
type combinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Incoming transfers of the asset
type CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                               `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaInJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaInJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaIn]
type combinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaOutJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaOutJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaOut]
type combinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationAccountSummaryAccountDelegation struct {
	Asset CombinedValidationResultSimulationAccountSummaryAccountDelegationsAsset `json:"asset,required"`
	// The delegate's address
	Delegate string `json:"delegate,required"`
	// Details of the delegation
	Delegation CombinedValidationResultSimulationAccountSummaryAccountDelegationsDelegation `json:"delegation,required"`
	JSON       combinedValidationResultSimulationAccountSummaryAccountDelegationJSON        `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountDelegationJSON contains
// the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountDelegation]
type combinedValidationResultSimulationAccountSummaryAccountDelegationJSON struct {
	Asset       apijson.Field
	Delegate    apijson.Field
	Delegation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountDelegationJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationAccountSummaryAccountDelegationsAsset struct {
	Address string `json:"address,required"`
	Symbol  string `json:"symbol,required"`
	Name    string `json:"name,required"`
	Logo    string `json:"logo"`
	// Type of the asset (`"TOKEN"`)
	Type     string                                                                      `json:"type"`
	Decimals int64                                                                       `json:"decimals"`
	JSON     combinedValidationResultSimulationAccountSummaryAccountDelegationsAssetJSON `json:"-"`
	union    CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetUnion
}

// combinedValidationResultSimulationAccountSummaryAccountDelegationsAssetJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountDelegationsAsset]
type combinedValidationResultSimulationAccountSummaryAccountDelegationsAssetJSON struct {
	Address     apijson.Field
	Symbol      apijson.Field
	Name        apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r combinedValidationResultSimulationAccountSummaryAccountDelegationsAssetJSON) RawJSON() string {
	return r.raw
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountDelegationsAsset) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetSplFungibleTokenDetailsSchema],
// [CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetSplNonFungibleTokenDetailsSchema],
// [CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetCnftDetailsSchema].
func (r CombinedValidationResultSimulationAccountSummaryAccountDelegationsAsset) AsUnion() CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetUnion {
	return r.union
}

// Union satisfied by
// [CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetSplFungibleTokenDetailsSchema],
// [CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetSplNonFungibleTokenDetailsSchema]
// or
// [CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetCnftDetailsSchema].
type CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetUnion interface {
	implementsCombinedValidationResultSimulationAccountSummaryAccountDelegationsAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetSplFungibleTokenDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetSplNonFungibleTokenDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetCnftDetailsSchema{}),
		},
	)
}

type CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetSplFungibleTokenDetailsSchema struct {
	Address  string `json:"address,required"`
	Decimals int64  `json:"decimals,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Logo     string `json:"logo"`
	// Type of the asset (`"TOKEN"`)
	Type string                                                                                                   `json:"type"`
	JSON combinedValidationResultSimulationAccountSummaryAccountDelegationsAssetSplFungibleTokenDetailsSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountDelegationsAssetSplFungibleTokenDetailsSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetSplFungibleTokenDetailsSchema]
type combinedValidationResultSimulationAccountSummaryAccountDelegationsAssetSplFungibleTokenDetailsSchemaJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetSplFungibleTokenDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountDelegationsAssetSplFungibleTokenDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetSplFungibleTokenDetailsSchema) implementsCombinedValidationResultSimulationAccountSummaryAccountDelegationsAsset() {
}

type CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetSplNonFungibleTokenDetailsSchema struct {
	Address  string `json:"address,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	Logo     string `json:"logo"`
	// Type of the asset (`"NFT"`)
	Type string                                                                                                      `json:"type"`
	JSON combinedValidationResultSimulationAccountSummaryAccountDelegationsAssetSplNonFungibleTokenDetailsSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountDelegationsAssetSplNonFungibleTokenDetailsSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetSplNonFungibleTokenDetailsSchema]
type combinedValidationResultSimulationAccountSummaryAccountDelegationsAssetSplNonFungibleTokenDetailsSchemaJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetSplNonFungibleTokenDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountDelegationsAssetSplNonFungibleTokenDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetSplNonFungibleTokenDetailsSchema) implementsCombinedValidationResultSimulationAccountSummaryAccountDelegationsAsset() {
}

type CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetCnftDetailsSchema struct {
	Address  string `json:"address,required"`
	Decimals int64  `json:"decimals,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Logo     string `json:"logo"`
	// Type of the asset (`"CNFT"`)
	Type string                                                                                       `json:"type"`
	JSON combinedValidationResultSimulationAccountSummaryAccountDelegationsAssetCnftDetailsSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountDelegationsAssetCnftDetailsSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetCnftDetailsSchema]
type combinedValidationResultSimulationAccountSummaryAccountDelegationsAssetCnftDetailsSchemaJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetCnftDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountDelegationsAssetCnftDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAccountSummaryAccountDelegationsAssetCnftDetailsSchema) implementsCombinedValidationResultSimulationAccountSummaryAccountDelegationsAsset() {
}

// Details of the delegation
type CombinedValidationResultSimulationAccountSummaryAccountDelegationsDelegation struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                          `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAccountSummaryAccountDelegationsDelegationJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountDelegationsDelegationJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountDelegationsDelegation]
type combinedValidationResultSimulationAccountSummaryAccountDelegationsDelegationJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountDelegationsDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountDelegationsDelegationJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiff struct {
	// This field can have the runtime type of
	// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaAsset],
	// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset],
	// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAsset].
	Asset interface{} `json:"asset,required"`
	// This field can have the runtime type of
	// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaIn],
	// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaIn],
	// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaIn].
	In interface{} `json:"in_,required"`
	// This field can have the runtime type of
	// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaOut],
	// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaOut],
	// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaOut].
	Out interface{} `json:"out,required"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,nullable"`
	// The owner post the transaction
	PostOwner string                                                                    `json:"post_owner,required"`
	JSON      combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffJSON `json:"-"`
	union     CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffUnion
}

// combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiff]
type combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	PostOwner   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiff) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchema],
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchema],
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchema].
func (r CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiff) AsUnion() CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffUnion {
	return r.union
}

// Union satisfied by
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchema],
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchema]
// or
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchema].
type CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffUnion interface {
	implementsCombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchema{}),
		},
	)
}

type CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchema struct {
	Asset CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaAsset `json:"asset,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required"`
	// Incoming transfers of the asset
	In CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaIn `json:"in_,nullable"`
	// Details of the moved value
	Out CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaOut `json:"out,nullable"`
	// The owner prior to the transaction
	PreOwner string                                                                                                `json:"pre_owner,nullable"`
	JSON     combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchema]
type combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaJSON struct {
	Asset       apijson.Field
	PostOwner   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchema) implementsCombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiff() {
}

type CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaAsset struct {
	Decimals int64 `json:"decimals"`
	// Type of the asset (`"SOL"`)
	Type string                                                                                                     `json:"type"`
	JSON combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaAssetJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaAsset]
type combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Incoming transfers of the asset
type CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                 `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaInJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaInJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaIn]
type combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

// Details of the moved value
type CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                  `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaOutJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaOutJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaOut]
type combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchema struct {
	Asset CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset `json:"asset,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required"`
	// Incoming transfers of the asset
	In CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaIn `json:"in_,nullable"`
	// Details of the moved value
	Out CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaOut `json:"out,nullable"`
	// The owner prior to the transaction
	PreOwner string                                                                                               `json:"pre_owner,nullable"`
	JSON     combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchema]
type combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaJSON struct {
	Asset       apijson.Field
	PostOwner   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchema) implementsCombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiff() {
}

type CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset struct {
	Address string `json:"address,required"`
	Symbol  string `json:"symbol,required"`
	Name    string `json:"name,required"`
	Logo    string `json:"logo"`
	// Type of the asset (`"TOKEN"`)
	Type     string                                                                                                    `json:"type"`
	Decimals int64                                                                                                     `json:"decimals"`
	JSON     combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetJSON `json:"-"`
	union    CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetUnion
}

// combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset]
type combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetJSON struct {
	Address     apijson.Field
	Symbol      apijson.Field
	Name        apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema],
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema].
func (r CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset) AsUnion() CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetUnion {
	return r.union
}

// Union satisfied by
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema]
// or
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema].
type CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetUnion interface {
	implementsCombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema{}),
		},
	)
}

type CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema struct {
	Address  string `json:"address,required"`
	Decimals int64  `json:"decimals,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Logo     string `json:"logo"`
	// Type of the asset (`"TOKEN"`)
	Type string                                                                                                                                 `json:"type"`
	JSON combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema]
type combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchemaJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema) implementsCombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset() {
}

type CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema struct {
	Address  string `json:"address,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	Logo     string `json:"logo"`
	// Type of the asset (`"NFT"`)
	Type string                                                                                                                                    `json:"type"`
	JSON combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema]
type combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchemaJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema) implementsCombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset() {
}

// Incoming transfers of the asset
type CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaInJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaInJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaIn]
type combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

// Details of the moved value
type CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                 `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaOutJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaOutJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaOut]
type combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchema struct {
	// The owner post the transaction
	PostOwner string                                                                                                         `json:"post_owner,required"`
	Asset     CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAsset `json:"asset"`
	// Incoming transfers of the asset
	In CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaIn `json:"in_,nullable"`
	// Details of the moved value
	Out CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaOut `json:"out,nullable"`
	// The owner prior to the transaction
	PreOwner string                                                                                                        `json:"pre_owner,nullable"`
	JSON     combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchema]
type combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaJSON struct {
	PostOwner   apijson.Field
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchema) implementsCombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiff() {
}

type CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAsset struct {
	Decimals int64 `json:"decimals"`
	// Type of the asset (`"STAKED_SOL"`)
	Type string                                                                                                             `json:"type"`
	JSON combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAssetJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAsset]
type combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Incoming transfers of the asset
type CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                         `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaInJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaInJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaIn]
type combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

// Details of the moved value
type CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                          `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaOutJSON `json:"-"`
}

// combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaOutJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaOut]
type combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationAccountsDetail struct {
	Type CombinedValidationResultSimulationAccountsDetailsType `json:"type"`
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
	Owner string                                               `json:"owner"`
	JSON  combinedValidationResultSimulationAccountsDetailJSON `json:"-"`
	union CombinedValidationResultSimulationAccountsDetailsUnion
}

// combinedValidationResultSimulationAccountsDetailJSON contains the JSON metadata
// for the struct [CombinedValidationResultSimulationAccountsDetail]
type combinedValidationResultSimulationAccountsDetailJSON struct {
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

func (r combinedValidationResultSimulationAccountsDetailJSON) RawJSON() string {
	return r.raw
}

func (r *CombinedValidationResultSimulationAccountsDetail) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CombinedValidationResultSimulationAccountsDetailsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CombinedValidationResultSimulationAccountsDetailsSystemAccountDetailsSchema],
// [CombinedValidationResultSimulationAccountsDetailsTokenAccountDetailsSchema],
// [CombinedValidationResultSimulationAccountsDetailsFungibleMintAccountDetailsSchema],
// [CombinedValidationResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchema],
// [CombinedValidationResultSimulationAccountsDetailsProgramAccountDetailsSchema],
// [CombinedValidationResultSimulationAccountsDetailsPdaAccountSchema],
// [CombinedValidationResultSimulationAccountsDetailsCnftMintAccountDetailsSchema].
func (r CombinedValidationResultSimulationAccountsDetail) AsUnion() CombinedValidationResultSimulationAccountsDetailsUnion {
	return r.union
}

// Union satisfied by
// [CombinedValidationResultSimulationAccountsDetailsSystemAccountDetailsSchema],
// [CombinedValidationResultSimulationAccountsDetailsTokenAccountDetailsSchema],
// [CombinedValidationResultSimulationAccountsDetailsFungibleMintAccountDetailsSchema],
// [CombinedValidationResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchema],
// [CombinedValidationResultSimulationAccountsDetailsProgramAccountDetailsSchema],
// [CombinedValidationResultSimulationAccountsDetailsPdaAccountSchema] or
// [CombinedValidationResultSimulationAccountsDetailsCnftMintAccountDetailsSchema].
type CombinedValidationResultSimulationAccountsDetailsUnion interface {
	implementsCombinedValidationResultSimulationAccountsDetail()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CombinedValidationResultSimulationAccountsDetailsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAccountsDetailsSystemAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAccountsDetailsTokenAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAccountsDetailsFungibleMintAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAccountsDetailsProgramAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAccountsDetailsPdaAccountSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAccountsDetailsCnftMintAccountDetailsSchema{}),
		},
	)
}

type CombinedValidationResultSimulationAccountsDetailsSystemAccountDetailsSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// Whether the account had been written to during the simulation
	WasWrittenTo bool `json:"was_written_to,required"`
	// Description of the account
	Description string                                                                          `json:"description,nullable"`
	Type        CombinedValidationResultSimulationAccountsDetailsSystemAccountDetailsSchemaType `json:"type"`
	JSON        combinedValidationResultSimulationAccountsDetailsSystemAccountDetailsSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAccountsDetailsSystemAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountsDetailsSystemAccountDetailsSchema]
type combinedValidationResultSimulationAccountsDetailsSystemAccountDetailsSchemaJSON struct {
	AccountAddress apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountsDetailsSystemAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountsDetailsSystemAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAccountsDetailsSystemAccountDetailsSchema) implementsCombinedValidationResultSimulationAccountsDetail() {
}

type CombinedValidationResultSimulationAccountsDetailsSystemAccountDetailsSchemaType string

const (
	CombinedValidationResultSimulationAccountsDetailsSystemAccountDetailsSchemaTypeSystemAccount CombinedValidationResultSimulationAccountsDetailsSystemAccountDetailsSchemaType = "SYSTEM_ACCOUNT"
)

func (r CombinedValidationResultSimulationAccountsDetailsSystemAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case CombinedValidationResultSimulationAccountsDetailsSystemAccountDetailsSchemaTypeSystemAccount:
		return true
	}
	return false
}

type CombinedValidationResultSimulationAccountsDetailsTokenAccountDetailsSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// Encoded public key of the mint
	MintAddress string `json:"mint_address,required"`
	// Encoded public key of the owner
	OwnerAddress string `json:"owner_address,required"`
	// Whether the account had been written to during the simulation
	WasWrittenTo bool `json:"was_written_to,required"`
	// Description of the account
	Description string                                                                         `json:"description,nullable"`
	Type        CombinedValidationResultSimulationAccountsDetailsTokenAccountDetailsSchemaType `json:"type"`
	JSON        combinedValidationResultSimulationAccountsDetailsTokenAccountDetailsSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAccountsDetailsTokenAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountsDetailsTokenAccountDetailsSchema]
type combinedValidationResultSimulationAccountsDetailsTokenAccountDetailsSchemaJSON struct {
	AccountAddress apijson.Field
	MintAddress    apijson.Field
	OwnerAddress   apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountsDetailsTokenAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountsDetailsTokenAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAccountsDetailsTokenAccountDetailsSchema) implementsCombinedValidationResultSimulationAccountsDetail() {
}

type CombinedValidationResultSimulationAccountsDetailsTokenAccountDetailsSchemaType string

const (
	CombinedValidationResultSimulationAccountsDetailsTokenAccountDetailsSchemaTypeTokenAccount CombinedValidationResultSimulationAccountsDetailsTokenAccountDetailsSchemaType = "TOKEN_ACCOUNT"
)

func (r CombinedValidationResultSimulationAccountsDetailsTokenAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case CombinedValidationResultSimulationAccountsDetailsTokenAccountDetailsSchemaTypeTokenAccount:
		return true
	}
	return false
}

type CombinedValidationResultSimulationAccountsDetailsFungibleMintAccountDetailsSchema struct {
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
	Logo string                                                                                `json:"logo"`
	Type CombinedValidationResultSimulationAccountsDetailsFungibleMintAccountDetailsSchemaType `json:"type"`
	JSON combinedValidationResultSimulationAccountsDetailsFungibleMintAccountDetailsSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAccountsDetailsFungibleMintAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountsDetailsFungibleMintAccountDetailsSchema]
type combinedValidationResultSimulationAccountsDetailsFungibleMintAccountDetailsSchemaJSON struct {
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

func (r *CombinedValidationResultSimulationAccountsDetailsFungibleMintAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountsDetailsFungibleMintAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAccountsDetailsFungibleMintAccountDetailsSchema) implementsCombinedValidationResultSimulationAccountsDetail() {
}

type CombinedValidationResultSimulationAccountsDetailsFungibleMintAccountDetailsSchemaType string

const (
	CombinedValidationResultSimulationAccountsDetailsFungibleMintAccountDetailsSchemaTypeFungibleMintAccount CombinedValidationResultSimulationAccountsDetailsFungibleMintAccountDetailsSchemaType = "FUNGIBLE_MINT_ACCOUNT"
)

func (r CombinedValidationResultSimulationAccountsDetailsFungibleMintAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case CombinedValidationResultSimulationAccountsDetailsFungibleMintAccountDetailsSchemaTypeFungibleMintAccount:
		return true
	}
	return false
}

type CombinedValidationResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchema struct {
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
	Logo string                                                                                   `json:"logo"`
	Type CombinedValidationResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchemaType `json:"type"`
	JSON combinedValidationResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchema]
type combinedValidationResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchemaJSON struct {
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

func (r *CombinedValidationResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchema) implementsCombinedValidationResultSimulationAccountsDetail() {
}

type CombinedValidationResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchemaType string

const (
	CombinedValidationResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchemaTypeNonFungibleMintAccount CombinedValidationResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchemaType = "NON_FUNGIBLE_MINT_ACCOUNT"
)

func (r CombinedValidationResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case CombinedValidationResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchemaTypeNonFungibleMintAccount:
		return true
	}
	return false
}

type CombinedValidationResultSimulationAccountsDetailsProgramAccountDetailsSchema struct {
	// Encoded public key of the account
	AccountAddress string                                                                           `json:"account_address,required"`
	Type           CombinedValidationResultSimulationAccountsDetailsProgramAccountDetailsSchemaType `json:"type,required"`
	// Whether the account had been written to during the simulation
	WasWrittenTo bool `json:"was_written_to,required"`
	// Description of the account
	Description string                                                                           `json:"description,nullable"`
	JSON        combinedValidationResultSimulationAccountsDetailsProgramAccountDetailsSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAccountsDetailsProgramAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountsDetailsProgramAccountDetailsSchema]
type combinedValidationResultSimulationAccountsDetailsProgramAccountDetailsSchemaJSON struct {
	AccountAddress apijson.Field
	Type           apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountsDetailsProgramAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountsDetailsProgramAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAccountsDetailsProgramAccountDetailsSchema) implementsCombinedValidationResultSimulationAccountsDetail() {
}

type CombinedValidationResultSimulationAccountsDetailsProgramAccountDetailsSchemaType string

const (
	CombinedValidationResultSimulationAccountsDetailsProgramAccountDetailsSchemaTypeProgram       CombinedValidationResultSimulationAccountsDetailsProgramAccountDetailsSchemaType = "PROGRAM"
	CombinedValidationResultSimulationAccountsDetailsProgramAccountDetailsSchemaTypeNativeProgram CombinedValidationResultSimulationAccountsDetailsProgramAccountDetailsSchemaType = "NATIVE_PROGRAM"
)

func (r CombinedValidationResultSimulationAccountsDetailsProgramAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case CombinedValidationResultSimulationAccountsDetailsProgramAccountDetailsSchemaTypeProgram, CombinedValidationResultSimulationAccountsDetailsProgramAccountDetailsSchemaTypeNativeProgram:
		return true
	}
	return false
}

type CombinedValidationResultSimulationAccountsDetailsPdaAccountSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// The address of the owning program
	Owner string `json:"owner,required"`
	// Whether the account had been written to during the simulation
	WasWrittenTo bool `json:"was_written_to,required"`
	// Description of the account
	Description string                                                                `json:"description,nullable"`
	Type        CombinedValidationResultSimulationAccountsDetailsPdaAccountSchemaType `json:"type"`
	JSON        combinedValidationResultSimulationAccountsDetailsPdaAccountSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAccountsDetailsPdaAccountSchemaJSON contains
// the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountsDetailsPdaAccountSchema]
type combinedValidationResultSimulationAccountsDetailsPdaAccountSchemaJSON struct {
	AccountAddress apijson.Field
	Owner          apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAccountsDetailsPdaAccountSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountsDetailsPdaAccountSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAccountsDetailsPdaAccountSchema) implementsCombinedValidationResultSimulationAccountsDetail() {
}

type CombinedValidationResultSimulationAccountsDetailsPdaAccountSchemaType string

const (
	CombinedValidationResultSimulationAccountsDetailsPdaAccountSchemaTypePda CombinedValidationResultSimulationAccountsDetailsPdaAccountSchemaType = "PDA"
)

func (r CombinedValidationResultSimulationAccountsDetailsPdaAccountSchemaType) IsKnown() bool {
	switch r {
	case CombinedValidationResultSimulationAccountsDetailsPdaAccountSchemaTypePda:
		return true
	}
	return false
}

type CombinedValidationResultSimulationAccountsDetailsCnftMintAccountDetailsSchema struct {
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
	Logo string                                                                            `json:"logo"`
	Type CombinedValidationResultSimulationAccountsDetailsCnftMintAccountDetailsSchemaType `json:"type"`
	JSON combinedValidationResultSimulationAccountsDetailsCnftMintAccountDetailsSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAccountsDetailsCnftMintAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAccountsDetailsCnftMintAccountDetailsSchema]
type combinedValidationResultSimulationAccountsDetailsCnftMintAccountDetailsSchemaJSON struct {
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

func (r *CombinedValidationResultSimulationAccountsDetailsCnftMintAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAccountsDetailsCnftMintAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAccountsDetailsCnftMintAccountDetailsSchema) implementsCombinedValidationResultSimulationAccountsDetail() {
}

type CombinedValidationResultSimulationAccountsDetailsCnftMintAccountDetailsSchemaType string

const (
	CombinedValidationResultSimulationAccountsDetailsCnftMintAccountDetailsSchemaTypeCnftMintAccount CombinedValidationResultSimulationAccountsDetailsCnftMintAccountDetailsSchemaType = "CNFT_MINT_ACCOUNT"
)

func (r CombinedValidationResultSimulationAccountsDetailsCnftMintAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case CombinedValidationResultSimulationAccountsDetailsCnftMintAccountDetailsSchemaTypeCnftMintAccount:
		return true
	}
	return false
}

type CombinedValidationResultSimulationAccountsDetailsType string

const (
	CombinedValidationResultSimulationAccountsDetailsTypeSystemAccount          CombinedValidationResultSimulationAccountsDetailsType = "SYSTEM_ACCOUNT"
	CombinedValidationResultSimulationAccountsDetailsTypeTokenAccount           CombinedValidationResultSimulationAccountsDetailsType = "TOKEN_ACCOUNT"
	CombinedValidationResultSimulationAccountsDetailsTypeFungibleMintAccount    CombinedValidationResultSimulationAccountsDetailsType = "FUNGIBLE_MINT_ACCOUNT"
	CombinedValidationResultSimulationAccountsDetailsTypeNonFungibleMintAccount CombinedValidationResultSimulationAccountsDetailsType = "NON_FUNGIBLE_MINT_ACCOUNT"
	CombinedValidationResultSimulationAccountsDetailsTypeProgram                CombinedValidationResultSimulationAccountsDetailsType = "PROGRAM"
	CombinedValidationResultSimulationAccountsDetailsTypeNativeProgram          CombinedValidationResultSimulationAccountsDetailsType = "NATIVE_PROGRAM"
	CombinedValidationResultSimulationAccountsDetailsTypePda                    CombinedValidationResultSimulationAccountsDetailsType = "PDA"
	CombinedValidationResultSimulationAccountsDetailsTypeCnftMintAccount        CombinedValidationResultSimulationAccountsDetailsType = "CNFT_MINT_ACCOUNT"
)

func (r CombinedValidationResultSimulationAccountsDetailsType) IsKnown() bool {
	switch r {
	case CombinedValidationResultSimulationAccountsDetailsTypeSystemAccount, CombinedValidationResultSimulationAccountsDetailsTypeTokenAccount, CombinedValidationResultSimulationAccountsDetailsTypeFungibleMintAccount, CombinedValidationResultSimulationAccountsDetailsTypeNonFungibleMintAccount, CombinedValidationResultSimulationAccountsDetailsTypeProgram, CombinedValidationResultSimulationAccountsDetailsTypeNativeProgram, CombinedValidationResultSimulationAccountsDetailsTypePda, CombinedValidationResultSimulationAccountsDetailsTypeCnftMintAccount:
		return true
	}
	return false
}

type CombinedValidationResultSimulationAssetsDiff struct {
	// This field can have the runtime type of
	// [CombinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaAsset],
	// [CombinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaAsset],
	// [CombinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaAsset],
	// [CombinedValidationResultSimulationAssetsDiffCnftDiffSchemaAsset].
	Asset interface{} `json:"asset"`
	// This field can have the runtime type of
	// [CombinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaIn],
	// [CombinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaIn],
	// [CombinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaIn],
	// [CombinedValidationResultSimulationAssetsDiffCnftDiffSchemaIn].
	In interface{} `json:"in,required"`
	// This field can have the runtime type of
	// [CombinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaOut],
	// [CombinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaOut],
	// [CombinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaOut],
	// [CombinedValidationResultSimulationAssetsDiffCnftDiffSchemaOut].
	Out   interface{}                                      `json:"out,required"`
	JSON  combinedValidationResultSimulationAssetsDiffJSON `json:"-"`
	union CombinedValidationResultSimulationAssetsDiffUnion
}

// combinedValidationResultSimulationAssetsDiffJSON contains the JSON metadata for
// the struct [CombinedValidationResultSimulationAssetsDiff]
type combinedValidationResultSimulationAssetsDiffJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r combinedValidationResultSimulationAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *CombinedValidationResultSimulationAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CombinedValidationResultSimulationAssetsDiffUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CombinedValidationResultSimulationAssetsDiffNativeSolDiffSchema],
// [CombinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchema],
// [CombinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchema],
// [CombinedValidationResultSimulationAssetsDiffCnftDiffSchema].
func (r CombinedValidationResultSimulationAssetsDiff) AsUnion() CombinedValidationResultSimulationAssetsDiffUnion {
	return r.union
}

// Union satisfied by
// [CombinedValidationResultSimulationAssetsDiffNativeSolDiffSchema],
// [CombinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchema],
// [CombinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchema] or
// [CombinedValidationResultSimulationAssetsDiffCnftDiffSchema].
type CombinedValidationResultSimulationAssetsDiffUnion interface {
	implementsCombinedValidationResultSimulationAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CombinedValidationResultSimulationAssetsDiffUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAssetsDiffNativeSolDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAssetsDiffCnftDiffSchema{}),
		},
	)
}

type CombinedValidationResultSimulationAssetsDiffNativeSolDiffSchema struct {
	Asset CombinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaAsset `json:"asset,required"`
	// Incoming transfers of the asset
	In   CombinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaIn   `json:"in,nullable"`
	Out  CombinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaOut  `json:"out,nullable"`
	JSON combinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaJSON contains the
// JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsDiffNativeSolDiffSchema]
type combinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsDiffNativeSolDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAssetsDiffNativeSolDiffSchema) implementsCombinedValidationResultSimulationAssetsDiff() {
}

type CombinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaAsset struct {
	Decimals int64 `json:"decimals"`
	// Type of the asset (`"SOL"`)
	Type string                                                                   `json:"type"`
	JSON combinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaAssetJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaAsset]
type combinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Incoming transfers of the asset
type CombinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                               `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaInJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaInJSON contains
// the JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaIn]
type combinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaOutJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaOutJSON contains
// the JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaOut]
type combinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsDiffNativeSolDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchema struct {
	Asset CombinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaAsset `json:"asset,required"`
	// Incoming transfers of the asset
	In   CombinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaIn   `json:"in,nullable"`
	Out  CombinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaOut  `json:"out,nullable"`
	JSON combinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchema]
type combinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchema) implementsCombinedValidationResultSimulationAssetsDiff() {
}

type CombinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaAsset struct {
	Address  string `json:"address,required"`
	Decimals int64  `json:"decimals,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Logo     string `json:"logo"`
	// Type of the asset (`"TOKEN"`)
	Type string                                                                          `json:"type"`
	JSON combinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaAssetJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaAsset]
type combinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Incoming transfers of the asset
type CombinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                      `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaInJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaInJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaIn]
type combinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                       `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaOutJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaOutJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaOut]
type combinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsDiffSplFungibleTokenDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchema struct {
	Asset CombinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaAsset `json:"asset,required"`
	// Incoming transfers of the asset
	In   CombinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaIn   `json:"in,nullable"`
	Out  CombinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaOut  `json:"out,nullable"`
	JSON combinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchema]
type combinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchema) implementsCombinedValidationResultSimulationAssetsDiff() {
}

type CombinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaAsset struct {
	Address  string `json:"address,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	Logo     string `json:"logo"`
	// Type of the asset (`"NFT"`)
	Type string                                                                             `json:"type"`
	JSON combinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaAssetJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaAsset]
type combinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Incoming transfers of the asset
type CombinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                         `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaInJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaInJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaIn]
type combinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                          `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaOutJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaOutJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaOut]
type combinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationAssetsDiffCnftDiffSchema struct {
	Asset CombinedValidationResultSimulationAssetsDiffCnftDiffSchemaAsset `json:"asset,required"`
	// Incoming transfers of the asset
	In   CombinedValidationResultSimulationAssetsDiffCnftDiffSchemaIn   `json:"in,nullable"`
	Out  CombinedValidationResultSimulationAssetsDiffCnftDiffSchemaOut  `json:"out,nullable"`
	JSON combinedValidationResultSimulationAssetsDiffCnftDiffSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsDiffCnftDiffSchemaJSON contains the JSON
// metadata for the struct
// [CombinedValidationResultSimulationAssetsDiffCnftDiffSchema]
type combinedValidationResultSimulationAssetsDiffCnftDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsDiffCnftDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsDiffCnftDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAssetsDiffCnftDiffSchema) implementsCombinedValidationResultSimulationAssetsDiff() {
}

type CombinedValidationResultSimulationAssetsDiffCnftDiffSchemaAsset struct {
	Address  string `json:"address,required"`
	Decimals int64  `json:"decimals,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Logo     string `json:"logo"`
	// Type of the asset (`"CNFT"`)
	Type string                                                              `json:"type"`
	JSON combinedValidationResultSimulationAssetsDiffCnftDiffSchemaAssetJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsDiffCnftDiffSchemaAssetJSON contains the
// JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsDiffCnftDiffSchemaAsset]
type combinedValidationResultSimulationAssetsDiffCnftDiffSchemaAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsDiffCnftDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsDiffCnftDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Incoming transfers of the asset
type CombinedValidationResultSimulationAssetsDiffCnftDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                          `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAssetsDiffCnftDiffSchemaInJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsDiffCnftDiffSchemaInJSON contains the
// JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsDiffCnftDiffSchemaIn]
type combinedValidationResultSimulationAssetsDiffCnftDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsDiffCnftDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsDiffCnftDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationAssetsDiffCnftDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                           `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAssetsDiffCnftDiffSchemaOutJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsDiffCnftDiffSchemaOutJSON contains the
// JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsDiffCnftDiffSchemaOut]
type combinedValidationResultSimulationAssetsDiffCnftDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsDiffCnftDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsDiffCnftDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationAssetsOwnershipDiff struct {
	// This field can have the runtime type of
	// [CombinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaAsset],
	// [CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset],
	// [CombinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAsset].
	Asset interface{} `json:"asset,required"`
	// This field can have the runtime type of
	// [CombinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaIn],
	// [CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaIn],
	// [CombinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaIn].
	In interface{} `json:"in_,required"`
	// This field can have the runtime type of
	// [CombinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaOut],
	// [CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaOut],
	// [CombinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaOut].
	Out interface{} `json:"out,required"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,nullable"`
	// The owner post the transaction
	PostOwner string                                                    `json:"post_owner,required"`
	JSON      combinedValidationResultSimulationAssetsOwnershipDiffJSON `json:"-"`
	union     CombinedValidationResultSimulationAssetsOwnershipDiffUnion
}

// combinedValidationResultSimulationAssetsOwnershipDiffJSON contains the JSON
// metadata for the struct [CombinedValidationResultSimulationAssetsOwnershipDiff]
type combinedValidationResultSimulationAssetsOwnershipDiffJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	PostOwner   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r combinedValidationResultSimulationAssetsOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

func (r *CombinedValidationResultSimulationAssetsOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CombinedValidationResultSimulationAssetsOwnershipDiffUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CombinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchema],
// [CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchema],
// [CombinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchema].
func (r CombinedValidationResultSimulationAssetsOwnershipDiff) AsUnion() CombinedValidationResultSimulationAssetsOwnershipDiffUnion {
	return r.union
}

// Union satisfied by
// [CombinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchema],
// [CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchema]
// or
// [CombinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchema].
type CombinedValidationResultSimulationAssetsOwnershipDiffUnion interface {
	implementsCombinedValidationResultSimulationAssetsOwnershipDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CombinedValidationResultSimulationAssetsOwnershipDiffUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchema{}),
		},
	)
}

type CombinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchema struct {
	Asset CombinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaAsset `json:"asset,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required"`
	// Incoming transfers of the asset
	In CombinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaIn `json:"in_,nullable"`
	// Details of the moved value
	Out CombinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaOut `json:"out,nullable"`
	// The owner prior to the transaction
	PreOwner string                                                                                `json:"pre_owner,nullable"`
	JSON     combinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchema]
type combinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaJSON struct {
	Asset       apijson.Field
	PostOwner   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchema) implementsCombinedValidationResultSimulationAssetsOwnershipDiff() {
}

type CombinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaAsset struct {
	Decimals int64 `json:"decimals"`
	// Type of the asset (`"SOL"`)
	Type string                                                                                     `json:"type"`
	JSON combinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaAssetJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaAsset]
type combinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Incoming transfers of the asset
type CombinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                 `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaInJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaInJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaIn]
type combinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

// Details of the moved value
type CombinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                  `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaOutJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaOutJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaOut]
type combinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchema struct {
	Asset CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset `json:"asset,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required"`
	// Incoming transfers of the asset
	In CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaIn `json:"in_,nullable"`
	// Details of the moved value
	Out CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaOut `json:"out,nullable"`
	// The owner prior to the transaction
	PreOwner string                                                                               `json:"pre_owner,nullable"`
	JSON     combinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchema]
type combinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaJSON struct {
	Asset       apijson.Field
	PostOwner   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchema) implementsCombinedValidationResultSimulationAssetsOwnershipDiff() {
}

type CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset struct {
	Address string `json:"address,required"`
	Symbol  string `json:"symbol,required"`
	Name    string `json:"name,required"`
	Logo    string `json:"logo"`
	// Type of the asset (`"TOKEN"`)
	Type     string                                                                                    `json:"type"`
	Decimals int64                                                                                     `json:"decimals"`
	JSON     combinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetJSON `json:"-"`
	union    CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetUnion
}

// combinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset]
type combinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetJSON struct {
	Address     apijson.Field
	Symbol      apijson.Field
	Name        apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r combinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

func (r *CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema],
// [CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema].
func (r CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset) AsUnion() CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetUnion {
	return r.union
}

// Union satisfied by
// [CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema]
// or
// [CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema].
type CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetUnion interface {
	implementsCombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema{}),
		},
	)
}

type CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema struct {
	Address  string `json:"address,required"`
	Decimals int64  `json:"decimals,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Logo     string `json:"logo"`
	// Type of the asset (`"TOKEN"`)
	Type string                                                                                                                 `json:"type"`
	JSON combinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema]
type combinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchemaJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplFungibleTokenDetailsSchema) implementsCombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset() {
}

type CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema struct {
	Address  string `json:"address,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	Logo     string `json:"logo"`
	// Type of the asset (`"NFT"`)
	Type string                                                                                                                    `json:"type"`
	JSON combinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema]
type combinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchemaJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetSplNonFungibleTokenDetailsSchema) implementsCombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset() {
}

// Incoming transfers of the asset
type CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaInJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaInJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaIn]
type combinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

// Details of the moved value
type CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                 `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaOutJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaOutJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaOut]
type combinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchema struct {
	// The owner post the transaction
	PostOwner string                                                                                         `json:"post_owner,required"`
	Asset     CombinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAsset `json:"asset"`
	// Incoming transfers of the asset
	In CombinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaIn `json:"in_,nullable"`
	// Details of the moved value
	Out CombinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaOut `json:"out,nullable"`
	// The owner prior to the transaction
	PreOwner string                                                                                        `json:"pre_owner,nullable"`
	JSON     combinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchema]
type combinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaJSON struct {
	PostOwner   apijson.Field
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchema) implementsCombinedValidationResultSimulationAssetsOwnershipDiff() {
}

type CombinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAsset struct {
	Decimals int64 `json:"decimals"`
	// Type of the asset (`"STAKED_SOL"`)
	Type string                                                                                             `json:"type"`
	JSON combinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAssetJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAsset]
type combinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Incoming transfers of the asset
type CombinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                         `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaInJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaInJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaIn]
type combinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

// Details of the moved value
type CombinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                          `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaOutJSON `json:"-"`
}

// combinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaOutJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaOut]
type combinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationDelegation struct {
	Asset CombinedValidationResultSimulationDelegationsAsset `json:"asset,required"`
	// The delegate's address
	Delegate string `json:"delegate,required"`
	// Details of the delegation
	Delegation CombinedValidationResultSimulationDelegationsDelegation `json:"delegation,required"`
	JSON       combinedValidationResultSimulationDelegationJSON        `json:"-"`
}

// combinedValidationResultSimulationDelegationJSON contains the JSON metadata for
// the struct [CombinedValidationResultSimulationDelegation]
type combinedValidationResultSimulationDelegationJSON struct {
	Asset       apijson.Field
	Delegate    apijson.Field
	Delegation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationDelegationJSON) RawJSON() string {
	return r.raw
}

type CombinedValidationResultSimulationDelegationsAsset struct {
	Address string `json:"address,required"`
	Symbol  string `json:"symbol,required"`
	Name    string `json:"name,required"`
	Logo    string `json:"logo"`
	// Type of the asset (`"TOKEN"`)
	Type     string                                                 `json:"type"`
	Decimals int64                                                  `json:"decimals"`
	JSON     combinedValidationResultSimulationDelegationsAssetJSON `json:"-"`
	union    CombinedValidationResultSimulationDelegationsAssetUnion
}

// combinedValidationResultSimulationDelegationsAssetJSON contains the JSON
// metadata for the struct [CombinedValidationResultSimulationDelegationsAsset]
type combinedValidationResultSimulationDelegationsAssetJSON struct {
	Address     apijson.Field
	Symbol      apijson.Field
	Name        apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r combinedValidationResultSimulationDelegationsAssetJSON) RawJSON() string {
	return r.raw
}

func (r *CombinedValidationResultSimulationDelegationsAsset) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CombinedValidationResultSimulationDelegationsAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CombinedValidationResultSimulationDelegationsAssetSplFungibleTokenDetailsSchema],
// [CombinedValidationResultSimulationDelegationsAssetSplNonFungibleTokenDetailsSchema],
// [CombinedValidationResultSimulationDelegationsAssetCnftDetailsSchema].
func (r CombinedValidationResultSimulationDelegationsAsset) AsUnion() CombinedValidationResultSimulationDelegationsAssetUnion {
	return r.union
}

// Union satisfied by
// [CombinedValidationResultSimulationDelegationsAssetSplFungibleTokenDetailsSchema],
// [CombinedValidationResultSimulationDelegationsAssetSplNonFungibleTokenDetailsSchema]
// or [CombinedValidationResultSimulationDelegationsAssetCnftDetailsSchema].
type CombinedValidationResultSimulationDelegationsAssetUnion interface {
	implementsCombinedValidationResultSimulationDelegationsAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CombinedValidationResultSimulationDelegationsAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationDelegationsAssetSplFungibleTokenDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationDelegationsAssetSplNonFungibleTokenDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CombinedValidationResultSimulationDelegationsAssetCnftDetailsSchema{}),
		},
	)
}

type CombinedValidationResultSimulationDelegationsAssetSplFungibleTokenDetailsSchema struct {
	Address  string `json:"address,required"`
	Decimals int64  `json:"decimals,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Logo     string `json:"logo"`
	// Type of the asset (`"TOKEN"`)
	Type string                                                                              `json:"type"`
	JSON combinedValidationResultSimulationDelegationsAssetSplFungibleTokenDetailsSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationDelegationsAssetSplFungibleTokenDetailsSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationDelegationsAssetSplFungibleTokenDetailsSchema]
type combinedValidationResultSimulationDelegationsAssetSplFungibleTokenDetailsSchemaJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationDelegationsAssetSplFungibleTokenDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationDelegationsAssetSplFungibleTokenDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationDelegationsAssetSplFungibleTokenDetailsSchema) implementsCombinedValidationResultSimulationDelegationsAsset() {
}

type CombinedValidationResultSimulationDelegationsAssetSplNonFungibleTokenDetailsSchema struct {
	Address  string `json:"address,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	Logo     string `json:"logo"`
	// Type of the asset (`"NFT"`)
	Type string                                                                                 `json:"type"`
	JSON combinedValidationResultSimulationDelegationsAssetSplNonFungibleTokenDetailsSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationDelegationsAssetSplNonFungibleTokenDetailsSchemaJSON
// contains the JSON metadata for the struct
// [CombinedValidationResultSimulationDelegationsAssetSplNonFungibleTokenDetailsSchema]
type combinedValidationResultSimulationDelegationsAssetSplNonFungibleTokenDetailsSchemaJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationDelegationsAssetSplNonFungibleTokenDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationDelegationsAssetSplNonFungibleTokenDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationDelegationsAssetSplNonFungibleTokenDetailsSchema) implementsCombinedValidationResultSimulationDelegationsAsset() {
}

type CombinedValidationResultSimulationDelegationsAssetCnftDetailsSchema struct {
	Address  string `json:"address,required"`
	Decimals int64  `json:"decimals,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Logo     string `json:"logo"`
	// Type of the asset (`"CNFT"`)
	Type string                                                                  `json:"type"`
	JSON combinedValidationResultSimulationDelegationsAssetCnftDetailsSchemaJSON `json:"-"`
}

// combinedValidationResultSimulationDelegationsAssetCnftDetailsSchemaJSON contains
// the JSON metadata for the struct
// [CombinedValidationResultSimulationDelegationsAssetCnftDetailsSchema]
type combinedValidationResultSimulationDelegationsAssetCnftDetailsSchemaJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationDelegationsAssetCnftDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationDelegationsAssetCnftDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CombinedValidationResultSimulationDelegationsAssetCnftDetailsSchema) implementsCombinedValidationResultSimulationDelegationsAsset() {
}

// Details of the delegation
type CombinedValidationResultSimulationDelegationsDelegation struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                     `json:"usd_price,nullable"`
	JSON     combinedValidationResultSimulationDelegationsDelegationJSON `json:"-"`
}

// combinedValidationResultSimulationDelegationsDelegationJSON contains the JSON
// metadata for the struct
// [CombinedValidationResultSimulationDelegationsDelegation]
type combinedValidationResultSimulationDelegationsDelegationJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CombinedValidationResultSimulationDelegationsDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r combinedValidationResultSimulationDelegationsDelegationJSON) RawJSON() string {
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
