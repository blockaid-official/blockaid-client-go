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

// SolanaMessageService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSolanaMessageService] method instead.
type SolanaMessageService struct {
	Options []option.RequestOption
}

// NewSolanaMessageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSolanaMessageService(opts ...option.RequestOption) (r *SolanaMessageService) {
	r = &SolanaMessageService{}
	r.Options = opts
	return
}

// Scan Messages
func (r *SolanaMessageService) Scan(ctx context.Context, body SolanaMessageScanParams, opts ...option.RequestOption) (res *SolanaMessageScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/solana/message/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SolanaMessageScanResponse struct {
	// Encoding of the public keys
	Encoding string `json:"encoding"`
	// Error message if the simulation failed
	Error string `json:"error,nullable"`
	// Summary of the result
	Result SolanaMessageScanResponseResult `json:"result,nullable"`
	JSON   solanaMessageScanResponseJSON   `json:"-"`
}

// solanaMessageScanResponseJSON contains the JSON metadata for the struct
// [SolanaMessageScanResponse]
type solanaMessageScanResponseJSON struct {
	Encoding    apijson.Field
	Error       apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseJSON) RawJSON() string {
	return r.raw
}

// Summary of the result
type SolanaMessageScanResponseResult struct {
	// Transaction validation result
	Validation SolanaMessageScanResponseResultValidation `json:"validation,required"`
	// Transaction simulation result
	Simulation SolanaMessageScanResponseResultSimulation `json:"simulation,nullable"`
	JSON       solanaMessageScanResponseResultJSON       `json:"-"`
}

// solanaMessageScanResponseResultJSON contains the JSON metadata for the struct
// [SolanaMessageScanResponseResult]
type solanaMessageScanResponseResultJSON struct {
	Validation  apijson.Field
	Simulation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultJSON) RawJSON() string {
	return r.raw
}

// Transaction validation result
type SolanaMessageScanResponseResultValidation struct {
	// A list of features about this transaction explaining the validation
	Features []string `json:"features,required"`
	// An enumeration.
	Reason SolanaMessageScanResponseResultValidationReason `json:"reason,required"`
	// An enumeration.
	Verdict SolanaMessageScanResponseResultValidationVerdict `json:"verdict,required"`
	JSON    solanaMessageScanResponseResultValidationJSON    `json:"-"`
}

// solanaMessageScanResponseResultValidationJSON contains the JSON metadata for the
// struct [SolanaMessageScanResponseResultValidation]
type solanaMessageScanResponseResultValidationJSON struct {
	Features    apijson.Field
	Reason      apijson.Field
	Verdict     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultValidationJSON) RawJSON() string {
	return r.raw
}

// An enumeration.
type SolanaMessageScanResponseResultValidationReason string

const (
	SolanaMessageScanResponseResultValidationReasonEmpty                   SolanaMessageScanResponseResultValidationReason = ""
	SolanaMessageScanResponseResultValidationReasonSharedStateInBulk       SolanaMessageScanResponseResultValidationReason = "shared_state_in_bulk"
	SolanaMessageScanResponseResultValidationReasonUnknownProfiter         SolanaMessageScanResponseResultValidationReason = "unknown_profiter"
	SolanaMessageScanResponseResultValidationReasonUnfairTrade             SolanaMessageScanResponseResultValidationReason = "unfair_trade"
	SolanaMessageScanResponseResultValidationReasonTransferFarming         SolanaMessageScanResponseResultValidationReason = "transfer_farming"
	SolanaMessageScanResponseResultValidationReasonWritableAccountsFarming SolanaMessageScanResponseResultValidationReason = "writable_accounts_farming"
	SolanaMessageScanResponseResultValidationReasonNativeOwnershipChange   SolanaMessageScanResponseResultValidationReason = "native_ownership_change"
	SolanaMessageScanResponseResultValidationReasonSplTokenOwnershipChange SolanaMessageScanResponseResultValidationReason = "spl_token_ownership_change"
	SolanaMessageScanResponseResultValidationReasonExposureFarming         SolanaMessageScanResponseResultValidationReason = "exposure_farming"
	SolanaMessageScanResponseResultValidationReasonKnownAttacker           SolanaMessageScanResponseResultValidationReason = "known_attacker"
	SolanaMessageScanResponseResultValidationReasonInvalidSignature        SolanaMessageScanResponseResultValidationReason = "invalid_signature"
	SolanaMessageScanResponseResultValidationReasonOther                   SolanaMessageScanResponseResultValidationReason = "other"
)

func (r SolanaMessageScanResponseResultValidationReason) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultValidationReasonEmpty, SolanaMessageScanResponseResultValidationReasonSharedStateInBulk, SolanaMessageScanResponseResultValidationReasonUnknownProfiter, SolanaMessageScanResponseResultValidationReasonUnfairTrade, SolanaMessageScanResponseResultValidationReasonTransferFarming, SolanaMessageScanResponseResultValidationReasonWritableAccountsFarming, SolanaMessageScanResponseResultValidationReasonNativeOwnershipChange, SolanaMessageScanResponseResultValidationReasonSplTokenOwnershipChange, SolanaMessageScanResponseResultValidationReasonExposureFarming, SolanaMessageScanResponseResultValidationReasonKnownAttacker, SolanaMessageScanResponseResultValidationReasonInvalidSignature, SolanaMessageScanResponseResultValidationReasonOther:
		return true
	}
	return false
}

// An enumeration.
type SolanaMessageScanResponseResultValidationVerdict string

const (
	SolanaMessageScanResponseResultValidationVerdictBenign    SolanaMessageScanResponseResultValidationVerdict = "Benign"
	SolanaMessageScanResponseResultValidationVerdictWarning   SolanaMessageScanResponseResultValidationVerdict = "Warning"
	SolanaMessageScanResponseResultValidationVerdictMalicious SolanaMessageScanResponseResultValidationVerdict = "Malicious"
)

func (r SolanaMessageScanResponseResultValidationVerdict) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultValidationVerdictBenign, SolanaMessageScanResponseResultValidationVerdictWarning, SolanaMessageScanResponseResultValidationVerdictMalicious:
		return true
	}
	return false
}

// Transaction simulation result
type SolanaMessageScanResponseResultSimulation struct {
	// Summary of the requested account address
	AccountSummary SolanaMessageScanResponseResultSimulationAccountSummary `json:"account_summary,required"`
	// Summary of the accounts involved in the transaction simulation
	AccountsDetails []SolanaMessageScanResponseResultSimulationAccountsDetail `json:"accounts_details,required"`
	// Summary of the assets involved in the transaction simulation
	AssetsDiff map[string][]SolanaMessageScanResponseResultSimulationAssetsDiff `json:"assets_diff,required"`
	// Summary of ownership changes; By account address
	AssetsOwnershipDiff map[string][]SolanaMessageScanResponseResultSimulationAssetsOwnershipDiff `json:"assets_ownership_diff,required"`
	// Summary of the delegations, by account address
	Delegations map[string][]SolanaMessageScanResponseResultSimulationDelegation `json:"delegations,required"`
	JSON        solanaMessageScanResponseResultSimulationJSON                    `json:"-"`
}

// solanaMessageScanResponseResultSimulationJSON contains the JSON metadata for the
// struct [SolanaMessageScanResponseResultSimulation]
type solanaMessageScanResponseResultSimulationJSON struct {
	AccountSummary      apijson.Field
	AccountsDetails     apijson.Field
	AssetsDiff          apijson.Field
	AssetsOwnershipDiff apijson.Field
	Delegations         apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationJSON) RawJSON() string {
	return r.raw
}

// Summary of the requested account address
type SolanaMessageScanResponseResultSimulationAccountSummary struct {
	// Total USD diff for the requested account address
	TotalUsdDiff SolanaMessageScanResponseResultSimulationAccountSummaryTotalUsdDiff `json:"total_usd_diff,required"`
	// Assets diff of the requested account address
	AccountAssetsDiff []SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiff `json:"account_assets_diff"`
	// Delegated assets of the requested account address
	AccountDelegations []SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegation `json:"account_delegations"`
	// Account ownerships diff of the requested account address
	AccountOwnershipsDiff []SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiff `json:"account_ownerships_diff"`
	JSON                  solanaMessageScanResponseResultSimulationAccountSummaryJSON                    `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryJSON contains the JSON
// metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummary]
type solanaMessageScanResponseResultSimulationAccountSummaryJSON struct {
	TotalUsdDiff          apijson.Field
	AccountAssetsDiff     apijson.Field
	AccountDelegations    apijson.Field
	AccountOwnershipsDiff apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryJSON) RawJSON() string {
	return r.raw
}

// Total USD diff for the requested account address
type SolanaMessageScanResponseResultSimulationAccountSummaryTotalUsdDiff struct {
	// Total incoming USD transfers
	In float64 `json:"in,required"`
	// Total outgoing USD transfers
	Out float64 `json:"out,required"`
	// Total USD transfers
	Total float64                                                                 `json:"total,required"`
	JSON  solanaMessageScanResponseResultSimulationAccountSummaryTotalUsdDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryTotalUsdDiffJSON contains
// the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryTotalUsdDiff]
type solanaMessageScanResponseResultSimulationAccountSummaryTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiff struct {
	// This field can have the runtime type of [NativeSolDetailsSchema],
	// [SplFungibleTokenDetailsSchema], [SplNonFungibleTokenDetailsSchema],
	// [CnftDetailsSchema].
	Asset interface{} `json:"asset"`
	// Incoming transfers of the asset
	In    AssetTransferDetailsSchema                                                   `json:"in,nullable"`
	Out   AssetTransferDetailsSchema                                                   `json:"out,nullable"`
	JSON  solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffJSON `json:"-"`
	union SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffUnion
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiff]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchema],
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchema],
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchema],
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchema].
func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiff) AsUnion() SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffUnion {
	return r.union
}

// Union satisfied by
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchema],
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchema],
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchema]
// or
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchema].
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffUnion interface {
	implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchema{}),
		},
	)
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchema struct {
	Asset NativeSolDetailsSchema `json:"asset,required"`
	// Incoming transfers of the asset
	In   AssetTransferDetailsSchema                                                                      `json:"in,nullable"`
	Out  AssetTransferDetailsSchema                                                                      `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchema]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffNativeSolDiffSchema) implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiff() {
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchema struct {
	Asset SplFungibleTokenDetailsSchema `json:"asset,required"`
	// Incoming transfers of the asset
	In   AssetTransferDetailsSchema                                                                             `json:"in,nullable"`
	Out  AssetTransferDetailsSchema                                                                             `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchema]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSplFungibleTokenDiffSchema) implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiff() {
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchema struct {
	Asset SplNonFungibleTokenDetailsSchema `json:"asset,required"`
	// Incoming transfers of the asset
	In   AssetTransferDetailsSchema                                                                                `json:"in,nullable"`
	Out  AssetTransferDetailsSchema                                                                                `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchema]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSplNonFungibleTokenDiffSchema) implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiff() {
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchema struct {
	Asset CnftDetailsSchema `json:"asset,required"`
	// Incoming transfers of the asset
	In   AssetTransferDetailsSchema                                                                 `json:"in,nullable"`
	Out  AssetTransferDetailsSchema                                                                 `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchema]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffCnftDiffSchema) implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiff() {
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegation struct {
	Asset SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsAsset `json:"asset,required"`
	// The delegate's address
	Delegate string `json:"delegate,required"`
	// Details of the delegation
	Delegation AssetTransferDetailsSchema                                                   `json:"delegation,required"`
	JSON       solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegation]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationJSON struct {
	Asset       apijson.Field
	Delegate    apijson.Field
	Delegation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsAsset struct {
	Address string `json:"address,required"`
	Symbol  string `json:"symbol,required"`
	Name    string `json:"name,required"`
	Logo    string `json:"logo"`
	// Type of the asset (`"TOKEN"`)
	Type     string                                                                             `json:"type"`
	Decimals int64                                                                              `json:"decimals"`
	JSON     solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsAssetJSON `json:"-"`
	union    SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsAssetUnion
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsAsset]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsAssetJSON struct {
	Address     apijson.Field
	Symbol      apijson.Field
	Name        apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsAssetJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsAsset) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [SplFungibleTokenDetailsSchema],
// [SplNonFungibleTokenDetailsSchema], [CnftDetailsSchema].
func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsAsset) AsUnion() SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsAssetUnion {
	return r.union
}

// Union satisfied by [SplFungibleTokenDetailsSchema],
// [SplNonFungibleTokenDetailsSchema] or [CnftDetailsSchema].
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsAssetUnion interface {
	implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsAssetUnion)(nil)).Elem(),
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

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiff struct {
	// This field can have the runtime type of [NativeSolDetailsSchema],
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset],
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAsset].
	Asset interface{} `json:"asset,required"`
	// Incoming transfers of the asset
	In AssetTransferDetailsSchema `json:"in_,nullable"`
	// Details of the moved value
	Out AssetTransferDetailsSchema `json:"out,nullable"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,nullable"`
	// The owner post the transaction
	PostOwner string                                                                           `json:"post_owner,required"`
	JSON      solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffJSON `json:"-"`
	union     SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffUnion
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiff]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	PostOwner   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiff) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchema],
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchema],
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchema].
func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiff) AsUnion() SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffUnion {
	return r.union
}

// Union satisfied by
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchema],
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchema]
// or
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchema].
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffUnion interface {
	implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchema{}),
		},
	)
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchema struct {
	Asset NativeSolDetailsSchema `json:"asset,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required"`
	// Incoming transfers of the asset
	In AssetTransferDetailsSchema `json:"in_,nullable"`
	// Details of the moved value
	Out AssetTransferDetailsSchema `json:"out,nullable"`
	// The owner prior to the transaction
	PreOwner string                                                                                                       `json:"pre_owner,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchema]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaJSON struct {
	Asset       apijson.Field
	PostOwner   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffNativeSolOwnershipDiffSchema) implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiff() {
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchema struct {
	Asset SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset `json:"asset,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required"`
	// Incoming transfers of the asset
	In AssetTransferDetailsSchema `json:"in_,nullable"`
	// Details of the moved value
	Out AssetTransferDetailsSchema `json:"out,nullable"`
	// The owner prior to the transaction
	PreOwner string                                                                                                      `json:"pre_owner,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchema]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaJSON struct {
	Asset       apijson.Field
	PostOwner   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchema) implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiff() {
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset struct {
	Address string `json:"address,required"`
	Symbol  string `json:"symbol,required"`
	Name    string `json:"name,required"`
	Logo    string `json:"logo"`
	// Type of the asset (`"TOKEN"`)
	Type     string                                                                                                           `json:"type"`
	Decimals int64                                                                                                            `json:"decimals"`
	JSON     solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetJSON `json:"-"`
	union    SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetUnion
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetJSON struct {
	Address     apijson.Field
	Symbol      apijson.Field
	Name        apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [SplFungibleTokenDetailsSchema],
// [SplNonFungibleTokenDetailsSchema].
func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset) AsUnion() SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetUnion {
	return r.union
}

// Union satisfied by [SplFungibleTokenDetailsSchema] or
// [SplNonFungibleTokenDetailsSchema].
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetUnion interface {
	implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSplTokenOwnershipDiffSchemaAssetUnion)(nil)).Elem(),
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

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchema struct {
	// The owner post the transaction
	PostOwner string                                                                                                                `json:"post_owner,required"`
	Asset     SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAsset `json:"asset"`
	// Incoming transfers of the asset
	In AssetTransferDetailsSchema `json:"in_,nullable"`
	// Details of the moved value
	Out AssetTransferDetailsSchema `json:"out,nullable"`
	// The owner prior to the transaction
	PreOwner string                                                                                                               `json:"pre_owner,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchema]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaJSON struct {
	PostOwner   apijson.Field
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchema) implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiff() {
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAsset struct {
	Decimals int64 `json:"decimals"`
	// Type of the asset (`"STAKED_SOL"`)
	Type string                                                                                                                    `json:"type"`
	JSON solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAsset]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffStakedSolWithdrawAuthorityDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAccountsDetail struct {
	Type SolanaMessageScanResponseResultSimulationAccountsDetailsType `json:"type"`
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
	Owner string                                                      `json:"owner"`
	JSON  solanaMessageScanResponseResultSimulationAccountsDetailJSON `json:"-"`
	union SolanaMessageScanResponseResultSimulationAccountsDetailsUnion
}

// solanaMessageScanResponseResultSimulationAccountsDetailJSON contains the JSON
// metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountsDetail]
type solanaMessageScanResponseResultSimulationAccountsDetailJSON struct {
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

func (r solanaMessageScanResponseResultSimulationAccountsDetailJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseResultSimulationAccountsDetail) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SolanaMessageScanResponseResultSimulationAccountsDetailsUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SolanaMessageScanResponseResultSimulationAccountsDetailsSystemAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationAccountsDetailsTokenAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationAccountsDetailsFungibleMintAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationAccountsDetailsProgramAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationAccountsDetailsPdaAccountSchema],
// [SolanaMessageScanResponseResultSimulationAccountsDetailsCnftMintAccountDetailsSchema].
func (r SolanaMessageScanResponseResultSimulationAccountsDetail) AsUnion() SolanaMessageScanResponseResultSimulationAccountsDetailsUnion {
	return r.union
}

// Union satisfied by
// [SolanaMessageScanResponseResultSimulationAccountsDetailsSystemAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationAccountsDetailsTokenAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationAccountsDetailsFungibleMintAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationAccountsDetailsProgramAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationAccountsDetailsPdaAccountSchema] or
// [SolanaMessageScanResponseResultSimulationAccountsDetailsCnftMintAccountDetailsSchema].
type SolanaMessageScanResponseResultSimulationAccountsDetailsUnion interface {
	implementsSolanaMessageScanResponseResultSimulationAccountsDetail()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseResultSimulationAccountsDetailsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountsDetailsSystemAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountsDetailsTokenAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountsDetailsFungibleMintAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountsDetailsProgramAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountsDetailsPdaAccountSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountsDetailsCnftMintAccountDetailsSchema{}),
		},
	)
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsSystemAccountDetailsSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// Whether the account had been written to during the simulation
	WasWrittenTo bool `json:"was_written_to,required"`
	// Description of the account
	Description string                                                                                 `json:"description,nullable"`
	Type        SolanaMessageScanResponseResultSimulationAccountsDetailsSystemAccountDetailsSchemaType `json:"type"`
	JSON        solanaMessageScanResponseResultSimulationAccountsDetailsSystemAccountDetailsSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountsDetailsSystemAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountsDetailsSystemAccountDetailsSchema]
type solanaMessageScanResponseResultSimulationAccountsDetailsSystemAccountDetailsSchemaJSON struct {
	AccountAddress apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountsDetailsSystemAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountsDetailsSystemAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsSystemAccountDetailsSchema) implementsSolanaMessageScanResponseResultSimulationAccountsDetail() {
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsSystemAccountDetailsSchemaType string

const (
	SolanaMessageScanResponseResultSimulationAccountsDetailsSystemAccountDetailsSchemaTypeSystemAccount SolanaMessageScanResponseResultSimulationAccountsDetailsSystemAccountDetailsSchemaType = "SYSTEM_ACCOUNT"
)

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsSystemAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountsDetailsSystemAccountDetailsSchemaTypeSystemAccount:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsTokenAccountDetailsSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// Encoded public key of the mint
	MintAddress string `json:"mint_address,required"`
	// Encoded public key of the owner
	OwnerAddress string `json:"owner_address,required"`
	// Whether the account had been written to during the simulation
	WasWrittenTo bool `json:"was_written_to,required"`
	// Description of the account
	Description string                                                                                `json:"description,nullable"`
	Type        SolanaMessageScanResponseResultSimulationAccountsDetailsTokenAccountDetailsSchemaType `json:"type"`
	JSON        solanaMessageScanResponseResultSimulationAccountsDetailsTokenAccountDetailsSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountsDetailsTokenAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountsDetailsTokenAccountDetailsSchema]
type solanaMessageScanResponseResultSimulationAccountsDetailsTokenAccountDetailsSchemaJSON struct {
	AccountAddress apijson.Field
	MintAddress    apijson.Field
	OwnerAddress   apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountsDetailsTokenAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountsDetailsTokenAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsTokenAccountDetailsSchema) implementsSolanaMessageScanResponseResultSimulationAccountsDetail() {
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsTokenAccountDetailsSchemaType string

const (
	SolanaMessageScanResponseResultSimulationAccountsDetailsTokenAccountDetailsSchemaTypeTokenAccount SolanaMessageScanResponseResultSimulationAccountsDetailsTokenAccountDetailsSchemaType = "TOKEN_ACCOUNT"
)

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsTokenAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountsDetailsTokenAccountDetailsSchemaTypeTokenAccount:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsFungibleMintAccountDetailsSchema struct {
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
	Logo string                                                                                       `json:"logo"`
	Type SolanaMessageScanResponseResultSimulationAccountsDetailsFungibleMintAccountDetailsSchemaType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationAccountsDetailsFungibleMintAccountDetailsSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountsDetailsFungibleMintAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountsDetailsFungibleMintAccountDetailsSchema]
type solanaMessageScanResponseResultSimulationAccountsDetailsFungibleMintAccountDetailsSchemaJSON struct {
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

func (r *SolanaMessageScanResponseResultSimulationAccountsDetailsFungibleMintAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountsDetailsFungibleMintAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsFungibleMintAccountDetailsSchema) implementsSolanaMessageScanResponseResultSimulationAccountsDetail() {
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsFungibleMintAccountDetailsSchemaType string

const (
	SolanaMessageScanResponseResultSimulationAccountsDetailsFungibleMintAccountDetailsSchemaTypeFungibleMintAccount SolanaMessageScanResponseResultSimulationAccountsDetailsFungibleMintAccountDetailsSchemaType = "FUNGIBLE_MINT_ACCOUNT"
)

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsFungibleMintAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountsDetailsFungibleMintAccountDetailsSchemaTypeFungibleMintAccount:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchema struct {
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
	Logo string                                                                                          `json:"logo"`
	Type SolanaMessageScanResponseResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchemaType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchema]
type solanaMessageScanResponseResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchemaJSON struct {
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

func (r *SolanaMessageScanResponseResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchema) implementsSolanaMessageScanResponseResultSimulationAccountsDetail() {
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchemaType string

const (
	SolanaMessageScanResponseResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchemaTypeNonFungibleMintAccount SolanaMessageScanResponseResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchemaType = "NON_FUNGIBLE_MINT_ACCOUNT"
)

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountsDetailsNonFungibleMintAccountDetailsSchemaTypeNonFungibleMintAccount:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsProgramAccountDetailsSchema struct {
	// Encoded public key of the account
	AccountAddress string                                                                                  `json:"account_address,required"`
	Type           SolanaMessageScanResponseResultSimulationAccountsDetailsProgramAccountDetailsSchemaType `json:"type,required"`
	// Whether the account had been written to during the simulation
	WasWrittenTo bool `json:"was_written_to,required"`
	// Description of the account
	Description string                                                                                  `json:"description,nullable"`
	JSON        solanaMessageScanResponseResultSimulationAccountsDetailsProgramAccountDetailsSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountsDetailsProgramAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountsDetailsProgramAccountDetailsSchema]
type solanaMessageScanResponseResultSimulationAccountsDetailsProgramAccountDetailsSchemaJSON struct {
	AccountAddress apijson.Field
	Type           apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountsDetailsProgramAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountsDetailsProgramAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsProgramAccountDetailsSchema) implementsSolanaMessageScanResponseResultSimulationAccountsDetail() {
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsProgramAccountDetailsSchemaType string

const (
	SolanaMessageScanResponseResultSimulationAccountsDetailsProgramAccountDetailsSchemaTypeProgram       SolanaMessageScanResponseResultSimulationAccountsDetailsProgramAccountDetailsSchemaType = "PROGRAM"
	SolanaMessageScanResponseResultSimulationAccountsDetailsProgramAccountDetailsSchemaTypeNativeProgram SolanaMessageScanResponseResultSimulationAccountsDetailsProgramAccountDetailsSchemaType = "NATIVE_PROGRAM"
)

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsProgramAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountsDetailsProgramAccountDetailsSchemaTypeProgram, SolanaMessageScanResponseResultSimulationAccountsDetailsProgramAccountDetailsSchemaTypeNativeProgram:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsPdaAccountSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// The address of the owning program
	Owner string `json:"owner,required"`
	// Whether the account had been written to during the simulation
	WasWrittenTo bool `json:"was_written_to,required"`
	// Description of the account
	Description string                                                                       `json:"description,nullable"`
	Type        SolanaMessageScanResponseResultSimulationAccountsDetailsPdaAccountSchemaType `json:"type"`
	JSON        solanaMessageScanResponseResultSimulationAccountsDetailsPdaAccountSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountsDetailsPdaAccountSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountsDetailsPdaAccountSchema]
type solanaMessageScanResponseResultSimulationAccountsDetailsPdaAccountSchemaJSON struct {
	AccountAddress apijson.Field
	Owner          apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountsDetailsPdaAccountSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountsDetailsPdaAccountSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsPdaAccountSchema) implementsSolanaMessageScanResponseResultSimulationAccountsDetail() {
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsPdaAccountSchemaType string

const (
	SolanaMessageScanResponseResultSimulationAccountsDetailsPdaAccountSchemaTypePda SolanaMessageScanResponseResultSimulationAccountsDetailsPdaAccountSchemaType = "PDA"
)

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsPdaAccountSchemaType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountsDetailsPdaAccountSchemaTypePda:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsCnftMintAccountDetailsSchema struct {
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
	Type SolanaMessageScanResponseResultSimulationAccountsDetailsCnftMintAccountDetailsSchemaType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationAccountsDetailsCnftMintAccountDetailsSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountsDetailsCnftMintAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountsDetailsCnftMintAccountDetailsSchema]
type solanaMessageScanResponseResultSimulationAccountsDetailsCnftMintAccountDetailsSchemaJSON struct {
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

func (r *SolanaMessageScanResponseResultSimulationAccountsDetailsCnftMintAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountsDetailsCnftMintAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsCnftMintAccountDetailsSchema) implementsSolanaMessageScanResponseResultSimulationAccountsDetail() {
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsCnftMintAccountDetailsSchemaType string

const (
	SolanaMessageScanResponseResultSimulationAccountsDetailsCnftMintAccountDetailsSchemaTypeCnftMintAccount SolanaMessageScanResponseResultSimulationAccountsDetailsCnftMintAccountDetailsSchemaType = "CNFT_MINT_ACCOUNT"
)

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsCnftMintAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountsDetailsCnftMintAccountDetailsSchemaTypeCnftMintAccount:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsType string

const (
	SolanaMessageScanResponseResultSimulationAccountsDetailsTypeSystemAccount          SolanaMessageScanResponseResultSimulationAccountsDetailsType = "SYSTEM_ACCOUNT"
	SolanaMessageScanResponseResultSimulationAccountsDetailsTypeTokenAccount           SolanaMessageScanResponseResultSimulationAccountsDetailsType = "TOKEN_ACCOUNT"
	SolanaMessageScanResponseResultSimulationAccountsDetailsTypeFungibleMintAccount    SolanaMessageScanResponseResultSimulationAccountsDetailsType = "FUNGIBLE_MINT_ACCOUNT"
	SolanaMessageScanResponseResultSimulationAccountsDetailsTypeNonFungibleMintAccount SolanaMessageScanResponseResultSimulationAccountsDetailsType = "NON_FUNGIBLE_MINT_ACCOUNT"
	SolanaMessageScanResponseResultSimulationAccountsDetailsTypeProgram                SolanaMessageScanResponseResultSimulationAccountsDetailsType = "PROGRAM"
	SolanaMessageScanResponseResultSimulationAccountsDetailsTypeNativeProgram          SolanaMessageScanResponseResultSimulationAccountsDetailsType = "NATIVE_PROGRAM"
	SolanaMessageScanResponseResultSimulationAccountsDetailsTypePda                    SolanaMessageScanResponseResultSimulationAccountsDetailsType = "PDA"
	SolanaMessageScanResponseResultSimulationAccountsDetailsTypeCnftMintAccount        SolanaMessageScanResponseResultSimulationAccountsDetailsType = "CNFT_MINT_ACCOUNT"
)

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountsDetailsTypeSystemAccount, SolanaMessageScanResponseResultSimulationAccountsDetailsTypeTokenAccount, SolanaMessageScanResponseResultSimulationAccountsDetailsTypeFungibleMintAccount, SolanaMessageScanResponseResultSimulationAccountsDetailsTypeNonFungibleMintAccount, SolanaMessageScanResponseResultSimulationAccountsDetailsTypeProgram, SolanaMessageScanResponseResultSimulationAccountsDetailsTypeNativeProgram, SolanaMessageScanResponseResultSimulationAccountsDetailsTypePda, SolanaMessageScanResponseResultSimulationAccountsDetailsTypeCnftMintAccount:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationAssetsDiff struct {
	// This field can have the runtime type of [NativeSolDetailsSchema],
	// [SplFungibleTokenDetailsSchema], [SplNonFungibleTokenDetailsSchema],
	// [CnftDetailsSchema].
	Asset interface{} `json:"asset"`
	// Incoming transfers of the asset
	In    AssetTransferDetailsSchema                              `json:"in,nullable"`
	Out   AssetTransferDetailsSchema                              `json:"out,nullable"`
	JSON  solanaMessageScanResponseResultSimulationAssetsDiffJSON `json:"-"`
	union SolanaMessageScanResponseResultSimulationAssetsDiffUnion
}

// solanaMessageScanResponseResultSimulationAssetsDiffJSON contains the JSON
// metadata for the struct [SolanaMessageScanResponseResultSimulationAssetsDiff]
type solanaMessageScanResponseResultSimulationAssetsDiffJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r solanaMessageScanResponseResultSimulationAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseResultSimulationAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SolanaMessageScanResponseResultSimulationAssetsDiffUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SolanaMessageScanResponseResultSimulationAssetsDiffNativeSolDiffSchema],
// [SolanaMessageScanResponseResultSimulationAssetsDiffSplFungibleTokenDiffSchema],
// [SolanaMessageScanResponseResultSimulationAssetsDiffSplNonFungibleTokenDiffSchema],
// [SolanaMessageScanResponseResultSimulationAssetsDiffCnftDiffSchema].
func (r SolanaMessageScanResponseResultSimulationAssetsDiff) AsUnion() SolanaMessageScanResponseResultSimulationAssetsDiffUnion {
	return r.union
}

// Union satisfied by
// [SolanaMessageScanResponseResultSimulationAssetsDiffNativeSolDiffSchema],
// [SolanaMessageScanResponseResultSimulationAssetsDiffSplFungibleTokenDiffSchema],
// [SolanaMessageScanResponseResultSimulationAssetsDiffSplNonFungibleTokenDiffSchema]
// or [SolanaMessageScanResponseResultSimulationAssetsDiffCnftDiffSchema].
type SolanaMessageScanResponseResultSimulationAssetsDiffUnion interface {
	implementsSolanaMessageScanResponseResultSimulationAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseResultSimulationAssetsDiffUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAssetsDiffNativeSolDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAssetsDiffSplFungibleTokenDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAssetsDiffSplNonFungibleTokenDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAssetsDiffCnftDiffSchema{}),
		},
	)
}

type SolanaMessageScanResponseResultSimulationAssetsDiffNativeSolDiffSchema struct {
	Asset NativeSolDetailsSchema `json:"asset,required"`
	// Incoming transfers of the asset
	In   AssetTransferDetailsSchema                                                 `json:"in,nullable"`
	Out  AssetTransferDetailsSchema                                                 `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationAssetsDiffNativeSolDiffSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsDiffNativeSolDiffSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsDiffNativeSolDiffSchema]
type solanaMessageScanResponseResultSimulationAssetsDiffNativeSolDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsDiffNativeSolDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsDiffNativeSolDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAssetsDiffNativeSolDiffSchema) implementsSolanaMessageScanResponseResultSimulationAssetsDiff() {
}

type SolanaMessageScanResponseResultSimulationAssetsDiffSplFungibleTokenDiffSchema struct {
	Asset SplFungibleTokenDetailsSchema `json:"asset,required"`
	// Incoming transfers of the asset
	In   AssetTransferDetailsSchema                                                        `json:"in,nullable"`
	Out  AssetTransferDetailsSchema                                                        `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationAssetsDiffSplFungibleTokenDiffSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsDiffSplFungibleTokenDiffSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsDiffSplFungibleTokenDiffSchema]
type solanaMessageScanResponseResultSimulationAssetsDiffSplFungibleTokenDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsDiffSplFungibleTokenDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsDiffSplFungibleTokenDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAssetsDiffSplFungibleTokenDiffSchema) implementsSolanaMessageScanResponseResultSimulationAssetsDiff() {
}

type SolanaMessageScanResponseResultSimulationAssetsDiffSplNonFungibleTokenDiffSchema struct {
	Asset SplNonFungibleTokenDetailsSchema `json:"asset,required"`
	// Incoming transfers of the asset
	In   AssetTransferDetailsSchema                                                           `json:"in,nullable"`
	Out  AssetTransferDetailsSchema                                                           `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsDiffSplNonFungibleTokenDiffSchema]
type solanaMessageScanResponseResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsDiffSplNonFungibleTokenDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsDiffSplNonFungibleTokenDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAssetsDiffSplNonFungibleTokenDiffSchema) implementsSolanaMessageScanResponseResultSimulationAssetsDiff() {
}

type SolanaMessageScanResponseResultSimulationAssetsDiffCnftDiffSchema struct {
	Asset CnftDetailsSchema `json:"asset,required"`
	// Incoming transfers of the asset
	In   AssetTransferDetailsSchema                                            `json:"in,nullable"`
	Out  AssetTransferDetailsSchema                                            `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationAssetsDiffCnftDiffSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsDiffCnftDiffSchemaJSON contains
// the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsDiffCnftDiffSchema]
type solanaMessageScanResponseResultSimulationAssetsDiffCnftDiffSchemaJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsDiffCnftDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsDiffCnftDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAssetsDiffCnftDiffSchema) implementsSolanaMessageScanResponseResultSimulationAssetsDiff() {
}

type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiff struct {
	// This field can have the runtime type of [NativeSolDetailsSchema],
	// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset],
	// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAsset].
	Asset interface{} `json:"asset,required"`
	// Incoming transfers of the asset
	In AssetTransferDetailsSchema `json:"in_,nullable"`
	// Details of the moved value
	Out AssetTransferDetailsSchema `json:"out,nullable"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,nullable"`
	// The owner post the transaction
	PostOwner string                                                           `json:"post_owner,required"`
	JSON      solanaMessageScanResponseResultSimulationAssetsOwnershipDiffJSON `json:"-"`
	union     SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffUnion
}

// solanaMessageScanResponseResultSimulationAssetsOwnershipDiffJSON contains the
// JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiff]
type solanaMessageScanResponseResultSimulationAssetsOwnershipDiffJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	PostOwner   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r solanaMessageScanResponseResultSimulationAssetsOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseResultSimulationAssetsOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchema],
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchema],
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchema].
func (r SolanaMessageScanResponseResultSimulationAssetsOwnershipDiff) AsUnion() SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffUnion {
	return r.union
}

// Union satisfied by
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchema],
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchema]
// or
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchema].
type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffUnion interface {
	implementsSolanaMessageScanResponseResultSimulationAssetsOwnershipDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchema{}),
		},
	)
}

type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchema struct {
	Asset NativeSolDetailsSchema `json:"asset,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required"`
	// Incoming transfers of the asset
	In AssetTransferDetailsSchema `json:"in_,nullable"`
	// Details of the moved value
	Out AssetTransferDetailsSchema `json:"out,nullable"`
	// The owner prior to the transaction
	PreOwner string                                                                                       `json:"pre_owner,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchema]
type solanaMessageScanResponseResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaJSON struct {
	Asset       apijson.Field
	PostOwner   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffNativeSolOwnershipDiffSchema) implementsSolanaMessageScanResponseResultSimulationAssetsOwnershipDiff() {
}

type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchema struct {
	Asset SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset `json:"asset,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required"`
	// Incoming transfers of the asset
	In AssetTransferDetailsSchema `json:"in_,nullable"`
	// Details of the moved value
	Out AssetTransferDetailsSchema `json:"out,nullable"`
	// The owner prior to the transaction
	PreOwner string                                                                                      `json:"pre_owner,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchema]
type solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaJSON struct {
	Asset       apijson.Field
	PostOwner   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchema) implementsSolanaMessageScanResponseResultSimulationAssetsOwnershipDiff() {
}

type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset struct {
	Address string `json:"address,required"`
	Symbol  string `json:"symbol,required"`
	Name    string `json:"name,required"`
	Logo    string `json:"logo"`
	// Type of the asset (`"TOKEN"`)
	Type     string                                                                                           `json:"type"`
	Decimals int64                                                                                            `json:"decimals"`
	JSON     solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetJSON `json:"-"`
	union    SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetUnion
}

// solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset]
type solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetJSON struct {
	Address     apijson.Field
	Symbol      apijson.Field
	Name        apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [SplFungibleTokenDetailsSchema],
// [SplNonFungibleTokenDetailsSchema].
func (r SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset) AsUnion() SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetUnion {
	return r.union
}

// Union satisfied by [SplFungibleTokenDetailsSchema] or
// [SplNonFungibleTokenDetailsSchema].
type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetUnion interface {
	implementsSolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSplTokenOwnershipDiffSchemaAssetUnion)(nil)).Elem(),
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

type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchema struct {
	// The owner post the transaction
	PostOwner string                                                                                                `json:"post_owner,required"`
	Asset     SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAsset `json:"asset"`
	// Incoming transfers of the asset
	In AssetTransferDetailsSchema `json:"in_,nullable"`
	// Details of the moved value
	Out AssetTransferDetailsSchema `json:"out,nullable"`
	// The owner prior to the transaction
	PreOwner string                                                                                               `json:"pre_owner,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchema]
type solanaMessageScanResponseResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaJSON struct {
	PostOwner   apijson.Field
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	PreOwner    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchema) implementsSolanaMessageScanResponseResultSimulationAssetsOwnershipDiff() {
}

type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAsset struct {
	Decimals int64 `json:"decimals"`
	// Type of the asset (`"STAKED_SOL"`)
	Type string                                                                                                    `json:"type"`
	JSON solanaMessageScanResponseResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAsset]
type solanaMessageScanResponseResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsOwnershipDiffStakedSolWithdrawAuthorityDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationDelegation struct {
	Asset SolanaMessageScanResponseResultSimulationDelegationsAsset `json:"asset,required"`
	// The delegate's address
	Delegate string `json:"delegate,required"`
	// Details of the delegation
	Delegation AssetTransferDetailsSchema                              `json:"delegation,required"`
	JSON       solanaMessageScanResponseResultSimulationDelegationJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationDelegationJSON contains the JSON
// metadata for the struct [SolanaMessageScanResponseResultSimulationDelegation]
type solanaMessageScanResponseResultSimulationDelegationJSON struct {
	Asset       apijson.Field
	Delegate    apijson.Field
	Delegation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationDelegationJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationDelegationsAsset struct {
	Address string `json:"address,required"`
	Symbol  string `json:"symbol,required"`
	Name    string `json:"name,required"`
	Logo    string `json:"logo"`
	// Type of the asset (`"TOKEN"`)
	Type     string                                                        `json:"type"`
	Decimals int64                                                         `json:"decimals"`
	JSON     solanaMessageScanResponseResultSimulationDelegationsAssetJSON `json:"-"`
	union    SolanaMessageScanResponseResultSimulationDelegationsAssetUnion
}

// solanaMessageScanResponseResultSimulationDelegationsAssetJSON contains the JSON
// metadata for the struct
// [SolanaMessageScanResponseResultSimulationDelegationsAsset]
type solanaMessageScanResponseResultSimulationDelegationsAssetJSON struct {
	Address     apijson.Field
	Symbol      apijson.Field
	Name        apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	Decimals    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r solanaMessageScanResponseResultSimulationDelegationsAssetJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseResultSimulationDelegationsAsset) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SolanaMessageScanResponseResultSimulationDelegationsAssetUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [SplFungibleTokenDetailsSchema],
// [SplNonFungibleTokenDetailsSchema], [CnftDetailsSchema].
func (r SolanaMessageScanResponseResultSimulationDelegationsAsset) AsUnion() SolanaMessageScanResponseResultSimulationDelegationsAssetUnion {
	return r.union
}

// Union satisfied by [SplFungibleTokenDetailsSchema],
// [SplNonFungibleTokenDetailsSchema] or [CnftDetailsSchema].
type SolanaMessageScanResponseResultSimulationDelegationsAssetUnion interface {
	implementsSolanaMessageScanResponseResultSimulationDelegationsAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseResultSimulationDelegationsAssetUnion)(nil)).Elem(),
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

type SolanaMessageScanParams struct {
	// Encoded public key of the account to simulate the transaction on
	AccountAddress param.Field[string]                          `json:"account_address,required"`
	Metadata       param.Field[SolanaMessageScanParamsMetadata] `json:"metadata,required"`
	// Transactions to scan
	Transactions param.Field[[]string] `json:"transactions,required"`
	// Chain to scan the transaction on
	Chain param.Field[string] `json:"chain"`
	// Encoding of the transaction and public keys
	Encoding param.Field[string] `json:"encoding"`
	// The RPC method used by dApp to propose the transaction
	Method param.Field[string] `json:"method"`
}

func (r SolanaMessageScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SolanaMessageScanParamsMetadata struct {
	// URL of the dApp that originated the transaction
	URL param.Field[string] `json:"url"`
}

func (r SolanaMessageScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
