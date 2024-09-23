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
	TotalUsdDiff TotalUsdDiffSchema `json:"total_usd_diff,required"`
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

type AccountSummarySchemaAccountAssetsDiff struct {
	// Type of the asset involved in the transfer
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of [NativeSolDetailsSchema],
	// [SplFungibleTokenDetailsSchema], [SplNonFungibleTokenDetailsSchema],
	// [CnftDetailsSchema].
	Asset interface{} `json:"asset"`
	// Incoming transfers of the asset
	In    AssetTransferDetailsSchema                `json:"in,nullable"`
	Out   AssetTransferDetailsSchema                `json:"out,nullable"`
	JSON  accountSummarySchemaAccountAssetsDiffJSON `json:"-"`
	union AccountSummarySchemaAccountAssetsDiffUnion
}

// accountSummarySchemaAccountAssetsDiffJSON contains the JSON metadata for the
// struct [AccountSummarySchemaAccountAssetsDiff]
type accountSummarySchemaAccountAssetsDiffJSON struct {
	AssetType   apijson.Field
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
	*r = AccountSummarySchemaAccountAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccountSummarySchemaAccountAssetsDiffUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [NativeSolDiffSchema],
// [SplFungibleTokenDiffSchema], [SplNonFungibleTokenDiffSchema], [CnftDiffSchema].
func (r AccountSummarySchemaAccountAssetsDiff) AsUnion() AccountSummarySchemaAccountAssetsDiffUnion {
	return r.union
}

// Union satisfied by [NativeSolDiffSchema], [SplFungibleTokenDiffSchema],
// [SplNonFungibleTokenDiffSchema] or [CnftDiffSchema].
type AccountSummarySchemaAccountAssetsDiffUnion interface {
	implementsAccountSummarySchemaAccountAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountSummarySchemaAccountAssetsDiffUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NativeSolDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SplFungibleTokenDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SplNonFungibleTokenDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CnftDiffSchema{}),
		},
	)
}

type AccountSummarySchemaAccountOwnershipsDiff struct {
	// Type of the asset involved in the transfer
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of [NativeSolDetailsSchema],
	// [SplTokenOwnershipDiffSchemaAsset], [StakedSolAssetDetailsSchema].
	Asset interface{} `json:"asset,required"`
	// Incoming transfers of the asset
	In AssetTransferDetailsSchema `json:"in_,nullable"`
	// Details of the moved value
	Out AssetTransferDetailsSchema `json:"out,nullable"`
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
	AssetType   apijson.Field
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
	*r = AccountSummarySchemaAccountOwnershipsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccountSummarySchemaAccountOwnershipsDiffUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [NativeSolOwnershipDiffSchema],
// [SplTokenOwnershipDiffSchema], [StakedSolWithdrawAuthorityDiffSchema].
func (r AccountSummarySchemaAccountOwnershipsDiff) AsUnion() AccountSummarySchemaAccountOwnershipsDiffUnion {
	return r.union
}

// Union satisfied by [NativeSolOwnershipDiffSchema], [SplTokenOwnershipDiffSchema]
// or [StakedSolWithdrawAuthorityDiffSchema].
type AccountSummarySchemaAccountOwnershipsDiffUnion interface {
	implementsAccountSummarySchemaAccountOwnershipsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountSummarySchemaAccountOwnershipsDiffUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NativeSolOwnershipDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SplTokenOwnershipDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StakedSolWithdrawAuthorityDiffSchema{}),
		},
	)
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

type APIErrorDetails struct {
	// Advanced message of the error
	Message string              `json:"message,required"`
	Type    string              `json:"type"`
	JSON    apiErrorDetailsJSON `json:"-"`
}

// apiErrorDetailsJSON contains the JSON metadata for the struct [APIErrorDetails]
type apiErrorDetailsJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r APIErrorDetails) implementsResponseSchemaErrorDetails() {}

type AssetTransferDetailsSchema struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                        `json:"usd_price,nullable"`
	JSON     assetTransferDetailsSchemaJSON `json:"-"`
}

// assetTransferDetailsSchemaJSON contains the JSON metadata for the struct
// [AssetTransferDetailsSchema]
type assetTransferDetailsSchemaJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetTransferDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetTransferDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

type CnftDetailsSchema struct {
	Address  string `json:"address,required"`
	Decimals int64  `json:"decimals,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Logo     string `json:"logo"`
	// Type of the asset (`"CNFT"`)
	Type string                `json:"type"`
	JSON cnftDetailsSchemaJSON `json:"-"`
}

// cnftDetailsSchemaJSON contains the JSON metadata for the struct
// [CnftDetailsSchema]
type cnftDetailsSchemaJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CnftDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cnftDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CnftDetailsSchema) implementsDelegatedAssetDetailsSchemaAsset() {}

type CnftDiffSchema struct {
	Asset CnftDetailsSchema `json:"asset,required"`
	// Type of the asset involved in the transfer
	AssetType string `json:"asset_type,required"`
	// Incoming transfers of the asset
	In   AssetTransferDetailsSchema `json:"in,nullable"`
	Out  AssetTransferDetailsSchema `json:"out,nullable"`
	JSON cnftDiffSchemaJSON         `json:"-"`
}

// cnftDiffSchemaJSON contains the JSON metadata for the struct [CnftDiffSchema]
type cnftDiffSchemaJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CnftDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cnftDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CnftDiffSchema) implementsAccountSummarySchemaAccountAssetsDiff() {}

func (r CnftDiffSchema) implementsSuccessfulSimulationResultSchemaAssetsDiff() {}

type CnftMintAccountDetailsSchema struct {
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
	Logo string                           `json:"logo"`
	Type CnftMintAccountDetailsSchemaType `json:"type"`
	JSON cnftMintAccountDetailsSchemaJSON `json:"-"`
}

// cnftMintAccountDetailsSchemaJSON contains the JSON metadata for the struct
// [CnftMintAccountDetailsSchema]
type cnftMintAccountDetailsSchemaJSON struct {
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

func (r *CnftMintAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cnftMintAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r CnftMintAccountDetailsSchema) implementsSuccessfulSimulationResultSchemaAccountsDetail() {}

type CnftMintAccountDetailsSchemaType string

const (
	CnftMintAccountDetailsSchemaTypeCnftMintAccount CnftMintAccountDetailsSchemaType = "CNFT_MINT_ACCOUNT"
)

func (r CnftMintAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case CnftMintAccountDetailsSchemaTypeCnftMintAccount:
		return true
	}
	return false
}

type CombinedValidationResult struct {
	// Transaction simulation result
	Simulation SuccessfulSimulationResultSchema `json:"simulation,nullable"`
	// Transaction validation result
	Validation CombinedValidationResultValidation `json:"validation,nullable"`
	JSON       combinedValidationResultJSON       `json:"-"`
}

// combinedValidationResultJSON contains the JSON metadata for the struct
// [CombinedValidationResult]
type combinedValidationResultJSON struct {
	Simulation  apijson.Field
	Validation  apijson.Field
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
	// A list of features explaining what is happening in the transaction in different
	// levels of severity
	ExtendedFeatures []ValidationFeature `json:"extended_features,required"`
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
	ExtendedFeatures apijson.Field
	Features         apijson.Field
	Reason           apijson.Field
	Verdict          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
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
	// Type of the asset involved in the delegation
	AssetType string `json:"asset_type,required"`
	// The delegate's address
	Delegate string `json:"delegate,required"`
	// Details of the delegation
	Delegation AssetTransferDetailsSchema      `json:"delegation,required"`
	JSON       delegatedAssetDetailsSchemaJSON `json:"-"`
}

// delegatedAssetDetailsSchemaJSON contains the JSON metadata for the struct
// [DelegatedAssetDetailsSchema]
type delegatedAssetDetailsSchemaJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
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
	*r = DelegatedAssetDetailsSchemaAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DelegatedAssetDetailsSchemaAssetUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [SplFungibleTokenDetailsSchema],
// [SplNonFungibleTokenDetailsSchema], [CnftDetailsSchema].
func (r DelegatedAssetDetailsSchemaAsset) AsUnion() DelegatedAssetDetailsSchemaAssetUnion {
	return r.union
}

// Union satisfied by [SplFungibleTokenDetailsSchema],
// [SplNonFungibleTokenDetailsSchema] or [CnftDetailsSchema].
type DelegatedAssetDetailsSchemaAssetUnion interface {
	implementsDelegatedAssetDetailsSchemaAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DelegatedAssetDetailsSchemaAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SplFungibleTokenDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SplNonFungibleTokenDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CnftDetailsSchema{}),
		},
	)
}

type FungibleMintAccountDetailsSchema struct {
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
	Logo string                               `json:"logo"`
	Type FungibleMintAccountDetailsSchemaType `json:"type"`
	JSON fungibleMintAccountDetailsSchemaJSON `json:"-"`
}

// fungibleMintAccountDetailsSchemaJSON contains the JSON metadata for the struct
// [FungibleMintAccountDetailsSchema]
type fungibleMintAccountDetailsSchemaJSON struct {
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

func (r *FungibleMintAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fungibleMintAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r FungibleMintAccountDetailsSchema) implementsSuccessfulSimulationResultSchemaAccountsDetail() {
}

type FungibleMintAccountDetailsSchemaType string

const (
	FungibleMintAccountDetailsSchemaTypeFungibleMintAccount FungibleMintAccountDetailsSchemaType = "FUNGIBLE_MINT_ACCOUNT"
)

func (r FungibleMintAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case FungibleMintAccountDetailsSchemaTypeFungibleMintAccount:
		return true
	}
	return false
}

type InstructionErrorDetails struct {
	// Index of the instruction in the transaction
	InstructionIndex int64 `json:"instruction_index,required"`
	// Advanced message of the error
	Message string `json:"message,required"`
	// Index of the transaction in the bulk
	TransactionIndex int64 `json:"transaction_index,required"`
	// The program account that caused the error
	ProgramAccount string                      `json:"program_account,nullable"`
	Type           string                      `json:"type"`
	JSON           instructionErrorDetailsJSON `json:"-"`
}

// instructionErrorDetailsJSON contains the JSON metadata for the struct
// [InstructionErrorDetails]
type instructionErrorDetailsJSON struct {
	InstructionIndex apijson.Field
	Message          apijson.Field
	TransactionIndex apijson.Field
	ProgramAccount   apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InstructionErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instructionErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r InstructionErrorDetails) implementsResponseSchemaErrorDetails() {}

type NativeSolDetailsSchema struct {
	Decimals int64 `json:"decimals"`
	// Logo of Sol
	Logo string `json:"logo,nullable"`
	// Type of the asset (`"SOL"`)
	Type string                     `json:"type"`
	JSON nativeSolDetailsSchemaJSON `json:"-"`
}

// nativeSolDetailsSchemaJSON contains the JSON metadata for the struct
// [NativeSolDetailsSchema]
type nativeSolDetailsSchemaJSON struct {
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NativeSolDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nativeSolDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

type NativeSolDiffSchema struct {
	Asset NativeSolDetailsSchema `json:"asset,required"`
	// Type of the asset involved in the transfer
	AssetType string `json:"asset_type,required"`
	// Incoming transfers of the asset
	In   AssetTransferDetailsSchema `json:"in,nullable"`
	Out  AssetTransferDetailsSchema `json:"out,nullable"`
	JSON nativeSolDiffSchemaJSON    `json:"-"`
}

// nativeSolDiffSchemaJSON contains the JSON metadata for the struct
// [NativeSolDiffSchema]
type nativeSolDiffSchemaJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NativeSolDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nativeSolDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r NativeSolDiffSchema) implementsAccountSummarySchemaAccountAssetsDiff() {}

func (r NativeSolDiffSchema) implementsSuccessfulSimulationResultSchemaAssetsDiff() {}

type NativeSolOwnershipDiffSchema struct {
	Asset NativeSolDetailsSchema `json:"asset,required"`
	// Type of the asset involved in the transfer
	AssetType string `json:"asset_type,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required"`
	// Incoming transfers of the asset
	In AssetTransferDetailsSchema `json:"in_,nullable"`
	// Details of the moved value
	Out AssetTransferDetailsSchema `json:"out,nullable"`
	// The owner prior to the transaction
	PreOwner string                           `json:"pre_owner,nullable"`
	JSON     nativeSolOwnershipDiffSchemaJSON `json:"-"`
}

// nativeSolOwnershipDiffSchemaJSON contains the JSON metadata for the struct
// [NativeSolOwnershipDiffSchema]
type nativeSolOwnershipDiffSchemaJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	PostOwner   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NativeSolOwnershipDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nativeSolOwnershipDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r NativeSolOwnershipDiffSchema) implementsAccountSummarySchemaAccountOwnershipsDiff() {}

func (r NativeSolOwnershipDiffSchema) implementsSuccessfulSimulationResultSchemaAssetsOwnershipDiff() {
}

type NonFungibleMintAccountDetailsSchema struct {
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
	Logo string                                  `json:"logo"`
	Type NonFungibleMintAccountDetailsSchemaType `json:"type"`
	JSON nonFungibleMintAccountDetailsSchemaJSON `json:"-"`
}

// nonFungibleMintAccountDetailsSchemaJSON contains the JSON metadata for the
// struct [NonFungibleMintAccountDetailsSchema]
type nonFungibleMintAccountDetailsSchemaJSON struct {
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

func (r *NonFungibleMintAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nonFungibleMintAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r NonFungibleMintAccountDetailsSchema) implementsSuccessfulSimulationResultSchemaAccountsDetail() {
}

type NonFungibleMintAccountDetailsSchemaType string

const (
	NonFungibleMintAccountDetailsSchemaTypeNonFungibleMintAccount NonFungibleMintAccountDetailsSchemaType = "NON_FUNGIBLE_MINT_ACCOUNT"
)

func (r NonFungibleMintAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case NonFungibleMintAccountDetailsSchemaTypeNonFungibleMintAccount:
		return true
	}
	return false
}

type PdaAccountSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// The address of the owning program
	Owner string `json:"owner,required"`
	// Whether the account had been written to during the simulation
	WasWrittenTo bool `json:"was_written_to,required"`
	// Description of the account
	Description string               `json:"description,nullable"`
	Type        PdaAccountSchemaType `json:"type"`
	JSON        pdaAccountSchemaJSON `json:"-"`
}

// pdaAccountSchemaJSON contains the JSON metadata for the struct
// [PdaAccountSchema]
type pdaAccountSchemaJSON struct {
	AccountAddress apijson.Field
	Owner          apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PdaAccountSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pdaAccountSchemaJSON) RawJSON() string {
	return r.raw
}

func (r PdaAccountSchema) implementsSuccessfulSimulationResultSchemaAccountsDetail() {}

type PdaAccountSchemaType string

const (
	PdaAccountSchemaTypePda PdaAccountSchemaType = "PDA"
)

func (r PdaAccountSchemaType) IsKnown() bool {
	switch r {
	case PdaAccountSchemaTypePda:
		return true
	}
	return false
}

type ProgramAccountDetailsSchema struct {
	// Encoded public key of the account
	AccountAddress string                          `json:"account_address,required"`
	Type           ProgramAccountDetailsSchemaType `json:"type,required"`
	// Whether the account had been written to during the simulation
	WasWrittenTo bool `json:"was_written_to,required"`
	// Description of the account
	Description string                          `json:"description,nullable"`
	JSON        programAccountDetailsSchemaJSON `json:"-"`
}

// programAccountDetailsSchemaJSON contains the JSON metadata for the struct
// [ProgramAccountDetailsSchema]
type programAccountDetailsSchemaJSON struct {
	AccountAddress apijson.Field
	Type           apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProgramAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r programAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r ProgramAccountDetailsSchema) implementsSuccessfulSimulationResultSchemaAccountsDetail() {}

type ProgramAccountDetailsSchemaType string

const (
	ProgramAccountDetailsSchemaTypeProgram       ProgramAccountDetailsSchemaType = "PROGRAM"
	ProgramAccountDetailsSchemaTypeNativeProgram ProgramAccountDetailsSchemaType = "NATIVE_PROGRAM"
)

func (r ProgramAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case ProgramAccountDetailsSchemaTypeProgram, ProgramAccountDetailsSchemaTypeNativeProgram:
		return true
	}
	return false
}

type ResponseSchema struct {
	// Unique identifier of the request
	RequestID string `json:"request_id,required"`
	// An enumeration.
	Status ResponseSchemaStatus `json:"status,required"`
	// Encoding of the public keys
	Encoding string `json:"encoding"`
	// Error message if the simulation failed
	Error string `json:"error,nullable"`
	// Advanced error details
	ErrorDetails ResponseSchemaErrorDetails `json:"error_details,nullable"`
	// Summary of the result
	Result CombinedValidationResult `json:"result,nullable"`
	JSON   responseSchemaJSON       `json:"-"`
}

// responseSchemaJSON contains the JSON metadata for the struct [ResponseSchema]
type responseSchemaJSON struct {
	RequestID    apijson.Field
	Status       apijson.Field
	Encoding     apijson.Field
	Error        apijson.Field
	ErrorDetails apijson.Field
	Result       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResponseSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseSchemaJSON) RawJSON() string {
	return r.raw
}

// An enumeration.
type ResponseSchemaStatus string

const (
	ResponseSchemaStatusSuccess ResponseSchemaStatus = "SUCCESS"
	ResponseSchemaStatusError   ResponseSchemaStatus = "ERROR"
)

func (r ResponseSchemaStatus) IsKnown() bool {
	switch r {
	case ResponseSchemaStatusSuccess, ResponseSchemaStatusError:
		return true
	}
	return false
}

// Advanced error details
type ResponseSchemaErrorDetails struct {
	Type string `json:"type"`
	// Advanced message of the error
	Message string `json:"message,required"`
	// Index of the transaction in the bulk
	TransactionIndex int64 `json:"transaction_index"`
	// Index of the instruction in the transaction
	InstructionIndex int64 `json:"instruction_index"`
	// The program account that caused the error
	ProgramAccount string                         `json:"program_account,nullable"`
	JSON           responseSchemaErrorDetailsJSON `json:"-"`
	union          ResponseSchemaErrorDetailsUnion
}

// responseSchemaErrorDetailsJSON contains the JSON metadata for the struct
// [ResponseSchemaErrorDetails]
type responseSchemaErrorDetailsJSON struct {
	Type             apijson.Field
	Message          apijson.Field
	TransactionIndex apijson.Field
	InstructionIndex apijson.Field
	ProgramAccount   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r responseSchemaErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r *ResponseSchemaErrorDetails) UnmarshalJSON(data []byte) (err error) {
	*r = ResponseSchemaErrorDetails{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ResponseSchemaErrorDetailsUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [APIErrorDetails],
// [TransactionErrorDetails], [InstructionErrorDetails].
func (r ResponseSchemaErrorDetails) AsUnion() ResponseSchemaErrorDetailsUnion {
	return r.union
}

// Advanced error details
//
// Union satisfied by [APIErrorDetails], [TransactionErrorDetails] or
// [InstructionErrorDetails].
type ResponseSchemaErrorDetailsUnion interface {
	implementsResponseSchemaErrorDetails()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ResponseSchemaErrorDetailsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(APIErrorDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionErrorDetails{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InstructionErrorDetails{}),
		},
	)
}

type SplFungibleTokenDetailsSchema struct {
	Address  string `json:"address,required"`
	Decimals int64  `json:"decimals,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Logo     string `json:"logo"`
	// Type of the asset (`"TOKEN"`)
	Type string                            `json:"type"`
	JSON splFungibleTokenDetailsSchemaJSON `json:"-"`
}

// splFungibleTokenDetailsSchemaJSON contains the JSON metadata for the struct
// [SplFungibleTokenDetailsSchema]
type splFungibleTokenDetailsSchemaJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplFungibleTokenDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splFungibleTokenDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SplFungibleTokenDetailsSchema) implementsDelegatedAssetDetailsSchemaAsset() {}

func (r SplFungibleTokenDetailsSchema) implementsSplTokenOwnershipDiffSchemaAsset() {}

type SplFungibleTokenDiffSchema struct {
	Asset SplFungibleTokenDetailsSchema `json:"asset,required"`
	// Type of the asset involved in the transfer
	AssetType string `json:"asset_type,required"`
	// Incoming transfers of the asset
	In   AssetTransferDetailsSchema     `json:"in,nullable"`
	Out  AssetTransferDetailsSchema     `json:"out,nullable"`
	JSON splFungibleTokenDiffSchemaJSON `json:"-"`
}

// splFungibleTokenDiffSchemaJSON contains the JSON metadata for the struct
// [SplFungibleTokenDiffSchema]
type splFungibleTokenDiffSchemaJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplFungibleTokenDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splFungibleTokenDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SplFungibleTokenDiffSchema) implementsAccountSummarySchemaAccountAssetsDiff() {}

func (r SplFungibleTokenDiffSchema) implementsSuccessfulSimulationResultSchemaAssetsDiff() {}

type SplNonFungibleTokenDetailsSchema struct {
	Address  string `json:"address,required"`
	Name     string `json:"name,required"`
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	Logo     string `json:"logo"`
	// Type of the asset (`"NFT"`)
	Type string                               `json:"type"`
	JSON splNonFungibleTokenDetailsSchemaJSON `json:"-"`
}

// splNonFungibleTokenDetailsSchemaJSON contains the JSON metadata for the struct
// [SplNonFungibleTokenDetailsSchema]
type splNonFungibleTokenDetailsSchemaJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplNonFungibleTokenDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splNonFungibleTokenDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SplNonFungibleTokenDetailsSchema) implementsDelegatedAssetDetailsSchemaAsset() {}

func (r SplNonFungibleTokenDetailsSchema) implementsSplTokenOwnershipDiffSchemaAsset() {}

type SplNonFungibleTokenDiffSchema struct {
	Asset SplNonFungibleTokenDetailsSchema `json:"asset,required"`
	// Type of the asset involved in the transfer
	AssetType string `json:"asset_type,required"`
	// Incoming transfers of the asset
	In   AssetTransferDetailsSchema        `json:"in,nullable"`
	Out  AssetTransferDetailsSchema        `json:"out,nullable"`
	JSON splNonFungibleTokenDiffSchemaJSON `json:"-"`
}

// splNonFungibleTokenDiffSchemaJSON contains the JSON metadata for the struct
// [SplNonFungibleTokenDiffSchema]
type splNonFungibleTokenDiffSchemaJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplNonFungibleTokenDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splNonFungibleTokenDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SplNonFungibleTokenDiffSchema) implementsAccountSummarySchemaAccountAssetsDiff() {}

func (r SplNonFungibleTokenDiffSchema) implementsSuccessfulSimulationResultSchemaAssetsDiff() {}

type SplTokenOwnershipDiffSchema struct {
	Asset SplTokenOwnershipDiffSchemaAsset `json:"asset,required"`
	// Type of the asset involved in the transfer
	AssetType string `json:"asset_type,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required"`
	// Incoming transfers of the asset
	In AssetTransferDetailsSchema `json:"in_,nullable"`
	// Details of the moved value
	Out AssetTransferDetailsSchema `json:"out,nullable"`
	// The owner prior to the transaction
	PreOwner string                          `json:"pre_owner,nullable"`
	JSON     splTokenOwnershipDiffSchemaJSON `json:"-"`
}

// splTokenOwnershipDiffSchemaJSON contains the JSON metadata for the struct
// [SplTokenOwnershipDiffSchema]
type splTokenOwnershipDiffSchemaJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	PostOwner   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplTokenOwnershipDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splTokenOwnershipDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SplTokenOwnershipDiffSchema) implementsAccountSummarySchemaAccountOwnershipsDiff() {}

func (r SplTokenOwnershipDiffSchema) implementsSuccessfulSimulationResultSchemaAssetsOwnershipDiff() {
}

type SplTokenOwnershipDiffSchemaAsset struct {
	Address string `json:"address,required"`
	Symbol  string `json:"symbol,required"`
	Name    string `json:"name,required"`
	Logo    string `json:"logo"`
	// Type of the asset (`"TOKEN"`)
	Type     string                               `json:"type"`
	Decimals int64                                `json:"decimals"`
	JSON     splTokenOwnershipDiffSchemaAssetJSON `json:"-"`
	union    SplTokenOwnershipDiffSchemaAssetUnion
}

// splTokenOwnershipDiffSchemaAssetJSON contains the JSON metadata for the struct
// [SplTokenOwnershipDiffSchemaAsset]
type splTokenOwnershipDiffSchemaAssetJSON struct {
	Address     apijson.Field
	Symbol      apijson.Field
	Name        apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r splTokenOwnershipDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

func (r *SplTokenOwnershipDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	*r = SplTokenOwnershipDiffSchemaAsset{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SplTokenOwnershipDiffSchemaAssetUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [SplFungibleTokenDetailsSchema],
// [SplNonFungibleTokenDetailsSchema].
func (r SplTokenOwnershipDiffSchemaAsset) AsUnion() SplTokenOwnershipDiffSchemaAssetUnion {
	return r.union
}

// Union satisfied by [SplFungibleTokenDetailsSchema] or
// [SplNonFungibleTokenDetailsSchema].
type SplTokenOwnershipDiffSchemaAssetUnion interface {
	implementsSplTokenOwnershipDiffSchemaAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SplTokenOwnershipDiffSchemaAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SplFungibleTokenDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SplNonFungibleTokenDetailsSchema{}),
		},
	)
}

type StakedSolAssetDetailsSchema struct {
	Decimals int64 `json:"decimals"`
	// Logo of Sol
	Logo string `json:"logo,nullable"`
	// Type of the asset (`"STAKED_SOL"`)
	Type string                          `json:"type"`
	JSON stakedSolAssetDetailsSchemaJSON `json:"-"`
}

// stakedSolAssetDetailsSchemaJSON contains the JSON metadata for the struct
// [StakedSolAssetDetailsSchema]
type stakedSolAssetDetailsSchemaJSON struct {
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StakedSolAssetDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stakedSolAssetDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

type StakedSolWithdrawAuthorityDiffSchema struct {
	// Type of the asset involved in the transfer
	AssetType string `json:"asset_type,required"`
	// The owner post the transaction
	PostOwner string                      `json:"post_owner,required"`
	Asset     StakedSolAssetDetailsSchema `json:"asset"`
	// Incoming transfers of the asset
	In AssetTransferDetailsSchema `json:"in_,nullable"`
	// Details of the moved value
	Out AssetTransferDetailsSchema `json:"out,nullable"`
	// The owner prior to the transaction
	PreOwner string                                   `json:"pre_owner,nullable"`
	JSON     stakedSolWithdrawAuthorityDiffSchemaJSON `json:"-"`
}

// stakedSolWithdrawAuthorityDiffSchemaJSON contains the JSON metadata for the
// struct [StakedSolWithdrawAuthorityDiffSchema]
type stakedSolWithdrawAuthorityDiffSchemaJSON struct {
	AssetType   apijson.Field
	PostOwner   apijson.Field
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StakedSolWithdrawAuthorityDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stakedSolWithdrawAuthorityDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StakedSolWithdrawAuthorityDiffSchema) implementsAccountSummarySchemaAccountOwnershipsDiff() {}

func (r StakedSolWithdrawAuthorityDiffSchema) implementsSuccessfulSimulationResultSchemaAssetsOwnershipDiff() {
}

type SuccessfulSimulationResultSchema struct {
	// Summary of the requested account address
	AccountSummary AccountSummarySchema `json:"account_summary,required"`
	// Summary of the accounts involved in the transaction simulation
	AccountsDetails []SuccessfulSimulationResultSchemaAccountsDetail `json:"accounts_details,required"`
	// Summary of the assets involved in the transaction simulation
	AssetsDiff map[string][]SuccessfulSimulationResultSchemaAssetsDiff `json:"assets_diff,required"`
	// Summary of ownership changes, By account address
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
	*r = SuccessfulSimulationResultSchemaAccountsDetail{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SuccessfulSimulationResultSchemaAccountsDetailsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [SystemAccountDetailsSchema],
// [TokenAccountDetailsSchema], [FungibleMintAccountDetailsSchema],
// [NonFungibleMintAccountDetailsSchema], [ProgramAccountDetailsSchema],
// [PdaAccountSchema], [CnftMintAccountDetailsSchema].
func (r SuccessfulSimulationResultSchemaAccountsDetail) AsUnion() SuccessfulSimulationResultSchemaAccountsDetailsUnion {
	return r.union
}

// Union satisfied by [SystemAccountDetailsSchema], [TokenAccountDetailsSchema],
// [FungibleMintAccountDetailsSchema], [NonFungibleMintAccountDetailsSchema],
// [ProgramAccountDetailsSchema], [PdaAccountSchema] or
// [CnftMintAccountDetailsSchema].
type SuccessfulSimulationResultSchemaAccountsDetailsUnion interface {
	implementsSuccessfulSimulationResultSchemaAccountsDetail()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SuccessfulSimulationResultSchemaAccountsDetailsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SystemAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FungibleMintAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NonFungibleMintAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProgramAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PdaAccountSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CnftMintAccountDetailsSchema{}),
		},
	)
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
	// Type of the asset involved in the transfer
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of [NativeSolDetailsSchema],
	// [SplFungibleTokenDetailsSchema], [SplNonFungibleTokenDetailsSchema],
	// [CnftDetailsSchema].
	Asset interface{} `json:"asset"`
	// Incoming transfers of the asset
	In    AssetTransferDetailsSchema                     `json:"in,nullable"`
	Out   AssetTransferDetailsSchema                     `json:"out,nullable"`
	JSON  successfulSimulationResultSchemaAssetsDiffJSON `json:"-"`
	union SuccessfulSimulationResultSchemaAssetsDiffUnion
}

// successfulSimulationResultSchemaAssetsDiffJSON contains the JSON metadata for
// the struct [SuccessfulSimulationResultSchemaAssetsDiff]
type successfulSimulationResultSchemaAssetsDiffJSON struct {
	AssetType   apijson.Field
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
	*r = SuccessfulSimulationResultSchemaAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SuccessfulSimulationResultSchemaAssetsDiffUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [NativeSolDiffSchema],
// [SplFungibleTokenDiffSchema], [SplNonFungibleTokenDiffSchema], [CnftDiffSchema].
func (r SuccessfulSimulationResultSchemaAssetsDiff) AsUnion() SuccessfulSimulationResultSchemaAssetsDiffUnion {
	return r.union
}

// Union satisfied by [NativeSolDiffSchema], [SplFungibleTokenDiffSchema],
// [SplNonFungibleTokenDiffSchema] or [CnftDiffSchema].
type SuccessfulSimulationResultSchemaAssetsDiffUnion interface {
	implementsSuccessfulSimulationResultSchemaAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SuccessfulSimulationResultSchemaAssetsDiffUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NativeSolDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SplFungibleTokenDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SplNonFungibleTokenDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CnftDiffSchema{}),
		},
	)
}

type SuccessfulSimulationResultSchemaAssetsOwnershipDiff struct {
	// Type of the asset involved in the transfer
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of [NativeSolDetailsSchema],
	// [SplTokenOwnershipDiffSchemaAsset], [StakedSolAssetDetailsSchema].
	Asset interface{} `json:"asset,required"`
	// Incoming transfers of the asset
	In AssetTransferDetailsSchema `json:"in_,nullable"`
	// Details of the moved value
	Out AssetTransferDetailsSchema `json:"out,nullable"`
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
	AssetType   apijson.Field
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
	*r = SuccessfulSimulationResultSchemaAssetsOwnershipDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SuccessfulSimulationResultSchemaAssetsOwnershipDiffUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [NativeSolOwnershipDiffSchema],
// [SplTokenOwnershipDiffSchema], [StakedSolWithdrawAuthorityDiffSchema].
func (r SuccessfulSimulationResultSchemaAssetsOwnershipDiff) AsUnion() SuccessfulSimulationResultSchemaAssetsOwnershipDiffUnion {
	return r.union
}

// Union satisfied by [NativeSolOwnershipDiffSchema], [SplTokenOwnershipDiffSchema]
// or [StakedSolWithdrawAuthorityDiffSchema].
type SuccessfulSimulationResultSchemaAssetsOwnershipDiffUnion interface {
	implementsSuccessfulSimulationResultSchemaAssetsOwnershipDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SuccessfulSimulationResultSchemaAssetsOwnershipDiffUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NativeSolOwnershipDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SplTokenOwnershipDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StakedSolWithdrawAuthorityDiffSchema{}),
		},
	)
}

type SystemAccountDetailsSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// Whether the account had been written to during the simulation
	WasWrittenTo bool `json:"was_written_to,required"`
	// Description of the account
	Description string                         `json:"description,nullable"`
	Type        SystemAccountDetailsSchemaType `json:"type"`
	JSON        systemAccountDetailsSchemaJSON `json:"-"`
}

// systemAccountDetailsSchemaJSON contains the JSON metadata for the struct
// [SystemAccountDetailsSchema]
type systemAccountDetailsSchemaJSON struct {
	AccountAddress apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SystemAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r systemAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SystemAccountDetailsSchema) implementsSuccessfulSimulationResultSchemaAccountsDetail() {}

type SystemAccountDetailsSchemaType string

const (
	SystemAccountDetailsSchemaTypeSystemAccount SystemAccountDetailsSchemaType = "SYSTEM_ACCOUNT"
)

func (r SystemAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case SystemAccountDetailsSchemaTypeSystemAccount:
		return true
	}
	return false
}

type TokenAccountDetailsSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// Encoded public key of the mint
	MintAddress string `json:"mint_address,required"`
	// Encoded public key of the owner
	OwnerAddress string `json:"owner_address,required"`
	// Whether the account had been written to during the simulation
	WasWrittenTo bool `json:"was_written_to,required"`
	// Description of the account
	Description string                        `json:"description,nullable"`
	Type        TokenAccountDetailsSchemaType `json:"type"`
	JSON        tokenAccountDetailsSchemaJSON `json:"-"`
}

// tokenAccountDetailsSchemaJSON contains the JSON metadata for the struct
// [TokenAccountDetailsSchema]
type tokenAccountDetailsSchemaJSON struct {
	AccountAddress apijson.Field
	MintAddress    apijson.Field
	OwnerAddress   apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TokenAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r TokenAccountDetailsSchema) implementsSuccessfulSimulationResultSchemaAccountsDetail() {}

type TokenAccountDetailsSchemaType string

const (
	TokenAccountDetailsSchemaTypeTokenAccount TokenAccountDetailsSchemaType = "TOKEN_ACCOUNT"
)

func (r TokenAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case TokenAccountDetailsSchemaTypeTokenAccount:
		return true
	}
	return false
}

type TotalUsdDiffSchema struct {
	// Total incoming USD transfers
	In float64 `json:"in,required"`
	// Total outgoing USD transfers
	Out float64 `json:"out,required"`
	// Total USD transfers
	Total float64                `json:"total,required"`
	JSON  totalUsdDiffSchemaJSON `json:"-"`
}

// totalUsdDiffSchemaJSON contains the JSON metadata for the struct
// [TotalUsdDiffSchema]
type totalUsdDiffSchemaJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TotalUsdDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r totalUsdDiffSchemaJSON) RawJSON() string {
	return r.raw
}

type TransactionErrorDetails struct {
	// Advanced message of the error
	Message string `json:"message,required"`
	// Index of the transaction in the bulk
	TransactionIndex int64                       `json:"transaction_index,required"`
	Type             string                      `json:"type"`
	JSON             transactionErrorDetailsJSON `json:"-"`
}

// transactionErrorDetailsJSON contains the JSON metadata for the struct
// [TransactionErrorDetails]
type transactionErrorDetailsJSON struct {
	Message          apijson.Field
	TransactionIndex apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransactionErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r TransactionErrorDetails) implementsResponseSchemaErrorDetails() {}

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
	Method  param.Field[string]   `json:"method"`
	Options param.Field[[]string] `json:"options"`
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

type ValidationFeature struct {
	// Textual description
	Description string `json:"description,required"`
	// Feature name
	FeatureID string `json:"feature_id,required"`
	// An enumeration.
	Type ValidationFeatureType `json:"type,required"`
	// Address the feature refers to
	Address string                `json:"address,nullable"`
	JSON    validationFeatureJSON `json:"-"`
}

// validationFeatureJSON contains the JSON metadata for the struct
// [ValidationFeature]
type validationFeatureJSON struct {
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidationFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r validationFeatureJSON) RawJSON() string {
	return r.raw
}

// An enumeration.
type ValidationFeatureType string

const (
	ValidationFeatureTypeInfo      ValidationFeatureType = "Info"
	ValidationFeatureTypeBenign    ValidationFeatureType = "Benign"
	ValidationFeatureTypeWarning   ValidationFeatureType = "Warning"
	ValidationFeatureTypeMalicious ValidationFeatureType = "Malicious"
)

func (r ValidationFeatureType) IsKnown() bool {
	switch r {
	case ValidationFeatureTypeInfo, ValidationFeatureTypeBenign, ValidationFeatureTypeWarning, ValidationFeatureTypeMalicious:
		return true
	}
	return false
}
