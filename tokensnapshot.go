// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/apiquery"
	"github.com/blockaid-official/blockaid-client-go/internal/param"
	"github.com/blockaid-official/blockaid-client-go/internal/requestconfig"
	"github.com/blockaid-official/blockaid-client-go/option"
	"github.com/tidwall/gjson"
)

// TokenSnapshotService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTokenSnapshotService] method instead.
type TokenSnapshotService struct {
	Options []option.RequestOption
}

// NewTokenSnapshotService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTokenSnapshotService(opts ...option.RequestOption) (r *TokenSnapshotService) {
	r = &TokenSnapshotService{}
	r.Options = opts
	return
}

// Retrieve tokens that experienced a state change within the specified timeframe
//
// Specify your preferred page size to manage response size. Specify the time frame
// in minutes. it is recommended to use timeframes shorter than 30 minutes
//
// Use the provided cursor to navigate through the pages. The cursor value will be
// null on the final page, indicating there are no more results to fetch. To
// retrieve the complete data set, iterate through all pages using the cursors
// provided in the response
func (r *TokenSnapshotService) Diff(ctx context.Context, query TokenSnapshotDiffParams, opts ...option.RequestOption) (res *TokenSnapshotDiffResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/token/snapshot/diff"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// This endpoint provides a state snapshot of all token scans, allowing you to
// synchronize with the latest state
//
// Specify your preferred page size to manage response size. Use the provided
// cursor to navigate through the pages. The cursor value will be null on the final
// page, indicating there are no more results to fetch.
//
// To retrieve the complete data set, iterate through all pages using the cursors
// provided in the response
func (r *TokenSnapshotService) Full(ctx context.Context, query TokenSnapshotFullParams, opts ...option.RequestOption) (res *TokenSnapshotFullResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/token/snapshot/full"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type TokenSnapshotDiffResponse struct {
	// Cursor to refetch the current page
	CurrentPage string                          `json:"current_page" api:"required"`
	Items       []TokenSnapshotDiffResponseItem `json:"items" api:"required"`
	// Cursor for the next page
	NextPage string `json:"next_page" api:"nullable"`
	// Cursor for the previous page
	PreviousPage string                        `json:"previous_page" api:"nullable"`
	JSON         tokenSnapshotDiffResponseJSON `json:"-"`
}

// tokenSnapshotDiffResponseJSON contains the JSON metadata for the struct
// [TokenSnapshotDiffResponse]
type tokenSnapshotDiffResponseJSON struct {
	CurrentPage  apijson.Field
	Items        apijson.Field
	NextPage     apijson.Field
	PreviousPage apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TokenSnapshotDiffResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotDiffResponseJSON) RawJSON() string {
	return r.raw
}

type TokenSnapshotDiffResponseItem struct {
	// Token address to validate (EVM / Solana)
	Address string `json:"address" api:"required"`
	// Dictionary of detected attacks found during the scan
	AttackTypes map[string]TokenSnapshotDiffResponseItemsAttackType `json:"attack_types" api:"required"`
	// Blockchain network
	Chain TokenScanSupportedChain `json:"chain" api:"required"`
	// Fees associated with the token
	Fees TokenSnapshotDiffResponseItemsFees `json:"fees" api:"required"`
	// financial stats of the token
	FinancialStats FinancialStats `json:"financial_stats" api:"required"`
	// Score between 0 to 1 (double)
	MaliciousScore string `json:"malicious_score" api:"required"`
	// Metadata of the token
	Metadata TokenSnapshotDiffResponseItemsMetadata `json:"metadata" api:"required"`
	// General indication
	ResultType TokenSnapshotDiffResponseItemsResultType `json:"result_type" api:"required"`
	// Trading limits of the token
	TradingLimits TokenSnapshotDiffResponseItemsTradingLimits `json:"trading_limits" api:"required"`
	// List of features associated with the token
	Features []TokenSnapshotDiffResponseItemsFeature `json:"features"`
	JSON     tokenSnapshotDiffResponseItemJSON       `json:"-"`
}

// tokenSnapshotDiffResponseItemJSON contains the JSON metadata for the struct
// [TokenSnapshotDiffResponseItem]
type tokenSnapshotDiffResponseItemJSON struct {
	Address        apijson.Field
	AttackTypes    apijson.Field
	Chain          apijson.Field
	Fees           apijson.Field
	FinancialStats apijson.Field
	MaliciousScore apijson.Field
	Metadata       apijson.Field
	ResultType     apijson.Field
	TradingLimits  apijson.Field
	Features       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TokenSnapshotDiffResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotDiffResponseItemJSON) RawJSON() string {
	return r.raw
}

type TokenSnapshotDiffResponseItemsAttackType struct {
	// Score between 0 to 1 (double) that indicates the assurance this attack happened
	Score string `json:"score" api:"required"`
	// Object contains an extra information related to the attack
	Features interface{} `json:"features"`
	// If score is higher or equal to this field, the token is using this attack type
	Threshold string                                       `json:"threshold"`
	JSON      tokenSnapshotDiffResponseItemsAttackTypeJSON `json:"-"`
}

// tokenSnapshotDiffResponseItemsAttackTypeJSON contains the JSON metadata for the
// struct [TokenSnapshotDiffResponseItemsAttackType]
type tokenSnapshotDiffResponseItemsAttackTypeJSON struct {
	Score       apijson.Field
	Features    apijson.Field
	Threshold   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotDiffResponseItemsAttackType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotDiffResponseItemsAttackTypeJSON) RawJSON() string {
	return r.raw
}

// Fees associated with the token
type TokenSnapshotDiffResponseItemsFees struct {
	// Buy fee of the token
	Buy float64 `json:"buy" api:"nullable"`
	// Sell fee of the token
	Sell float64 `json:"sell" api:"nullable"`
	// Transfer fee of the token
	Transfer float64 `json:"transfer" api:"nullable"`
	// The maximum value that a transfer fee will cost
	TransferFeeMaxAmount int64                                  `json:"transfer_fee_max_amount" api:"nullable"`
	JSON                 tokenSnapshotDiffResponseItemsFeesJSON `json:"-"`
}

// tokenSnapshotDiffResponseItemsFeesJSON contains the JSON metadata for the struct
// [TokenSnapshotDiffResponseItemsFees]
type tokenSnapshotDiffResponseItemsFeesJSON struct {
	Buy                  apijson.Field
	Sell                 apijson.Field
	Transfer             apijson.Field
	TransferFeeMaxAmount apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TokenSnapshotDiffResponseItemsFees) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotDiffResponseItemsFeesJSON) RawJSON() string {
	return r.raw
}

// Metadata of the token
type TokenSnapshotDiffResponseItemsMetadata struct {
	// The unique ID for the Rune
	ID string `json:"id" api:"nullable"`
	// This field can have the runtime type of
	// [TokenSnapshotDiffResponseItemsMetadataSolanaMetadataContractBalance],
	// [TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenContractBalance].
	ContractBalance interface{} `json:"contract_balance"`
	// Contract deploy date
	CreationTimestamp string `json:"creation_timestamp" api:"nullable"`
	// Decimals of the token
	Decimals int64 `json:"decimals" api:"nullable"`
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer" api:"nullable"`
	// This field can have the runtime type of
	// [TokenSnapshotDiffResponseItemsMetadataSolanaMetadataDeployerBalance],
	// [TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenDeployerBalance].
	DeployerBalance interface{} `json:"deployer_balance"`
	// Description of the token
	Description string `json:"description" api:"nullable"`
	// This field can have the runtime type of
	// [TokenSnapshotDiffResponseItemsMetadataSolanaMetadataExternalLinks],
	// [TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenExternalLinks].
	ExternalLinks interface{} `json:"external_links"`
	// The formatted name of the rune, with spacers
	FormattedName string `json:"formatted_name" api:"nullable"`
	// Solana token freeze authority account
	FreezeAuthority string `json:"freeze_authority" api:"nullable"`
	// URL of the token image
	ImageURL string `json:"image_url" api:"nullable"`
	// This field can have the runtime type of
	// [[]TokenSnapshotDiffResponseItemsMetadataSolanaMetadataImpersonationTarget],
	// [[]TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenImpersonationTarget].
	ImpersonationTargets interface{} `json:"impersonation_targets"`
	// This field can have the runtime type of [[]string].
	MaliciousURLs interface{} `json:"malicious_urls"`
	// Solana token mint authority account
	MintAuthority string `json:"mint_authority" api:"nullable"`
	// Name of the token
	Name string `json:"name" api:"nullable"`
	// The rune's unique sequential number.
	Number int64 `json:"number" api:"nullable"`
	// Contract owner address
	Owner string `json:"owner" api:"nullable"`
	// This field can have the runtime type of
	// [TokenSnapshotDiffResponseItemsMetadataSolanaMetadataOwnerBalance],
	// [TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenOwnerBalance].
	OwnerBalance interface{} `json:"owner_balance"`
	// Solana token permanent delegate account
	PermanentDelegate string `json:"permanent_delegate" api:"nullable"`
	// Symbol of the token
	Symbol string `json:"symbol" api:"nullable"`
	// Address of the token creation initiator, only set if the tokens was created by a
	// well known token launch platform
	TokenCreationInitiator string `json:"token_creation_initiator" api:"nullable"`
	// Type of the token
	Type string `json:"type" api:"nullable"`
	// Solana token update authority account
	UpdateAuthority string `json:"update_authority" api:"nullable"`
	// This field can have the runtime type of [[]string].
	URLs  interface{}                                `json:"urls"`
	JSON  tokenSnapshotDiffResponseItemsMetadataJSON `json:"-"`
	union TokenSnapshotDiffResponseItemsMetadataUnion
}

// tokenSnapshotDiffResponseItemsMetadataJSON contains the JSON metadata for the
// struct [TokenSnapshotDiffResponseItemsMetadata]
type tokenSnapshotDiffResponseItemsMetadataJSON struct {
	ID                     apijson.Field
	ContractBalance        apijson.Field
	CreationTimestamp      apijson.Field
	Decimals               apijson.Field
	Deployer               apijson.Field
	DeployerBalance        apijson.Field
	Description            apijson.Field
	ExternalLinks          apijson.Field
	FormattedName          apijson.Field
	FreezeAuthority        apijson.Field
	ImageURL               apijson.Field
	ImpersonationTargets   apijson.Field
	MaliciousURLs          apijson.Field
	MintAuthority          apijson.Field
	Name                   apijson.Field
	Number                 apijson.Field
	Owner                  apijson.Field
	OwnerBalance           apijson.Field
	PermanentDelegate      apijson.Field
	Symbol                 apijson.Field
	TokenCreationInitiator apijson.Field
	Type                   apijson.Field
	UpdateAuthority        apijson.Field
	URLs                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r tokenSnapshotDiffResponseItemsMetadataJSON) RawJSON() string {
	return r.raw
}

func (r *TokenSnapshotDiffResponseItemsMetadata) UnmarshalJSON(data []byte) (err error) {
	*r = TokenSnapshotDiffResponseItemsMetadata{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TokenSnapshotDiffResponseItemsMetadataUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TokenSnapshotDiffResponseItemsMetadataSolanaMetadata],
// [TokenSnapshotDiffResponseItemsMetadataBitcoinMetadataToken],
// [TokenSnapshotDiffResponseItemsMetadataEvmMetadataToken].
func (r TokenSnapshotDiffResponseItemsMetadata) AsUnion() TokenSnapshotDiffResponseItemsMetadataUnion {
	return r.union
}

// Metadata of the token
//
// Union satisfied by [TokenSnapshotDiffResponseItemsMetadataSolanaMetadata],
// [TokenSnapshotDiffResponseItemsMetadataBitcoinMetadataToken] or
// [TokenSnapshotDiffResponseItemsMetadataEvmMetadataToken].
type TokenSnapshotDiffResponseItemsMetadataUnion interface {
	implementsTokenSnapshotDiffResponseItemsMetadata()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TokenSnapshotDiffResponseItemsMetadataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenSnapshotDiffResponseItemsMetadataSolanaMetadata{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenSnapshotDiffResponseItemsMetadataBitcoinMetadataToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenSnapshotDiffResponseItemsMetadataEvmMetadataToken{}),
		},
	)
}

type TokenSnapshotDiffResponseItemsMetadataSolanaMetadata struct {
	// Contract balance
	ContractBalance TokenSnapshotDiffResponseItemsMetadataSolanaMetadataContractBalance `json:"contract_balance" api:"nullable"`
	// Contract deploy date
	CreationTimestamp string `json:"creation_timestamp" api:"nullable"`
	// Decimals of the token
	Decimals int64 `json:"decimals" api:"nullable"`
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer" api:"nullable"`
	// Contract creator balance
	DeployerBalance TokenSnapshotDiffResponseItemsMetadataSolanaMetadataDeployerBalance `json:"deployer_balance" api:"nullable"`
	// Description of the token
	Description string `json:"description" api:"nullable"`
	// social links of the token
	ExternalLinks TokenSnapshotDiffResponseItemsMetadataSolanaMetadataExternalLinks `json:"external_links"`
	// Solana token freeze authority account
	FreezeAuthority string `json:"freeze_authority" api:"nullable"`
	// URL of the token image
	ImageURL string `json:"image_url" api:"nullable"`
	// List of tokens that this token is impersonating, if detected as an impersonator
	ImpersonationTargets []TokenSnapshotDiffResponseItemsMetadataSolanaMetadataImpersonationTarget `json:"impersonation_targets" api:"nullable"`
	// Malicious urls associated with the token
	MaliciousURLs []string `json:"malicious_urls" api:"nullable"`
	// Solana token mint authority account
	MintAuthority string `json:"mint_authority" api:"nullable"`
	// Name of the token
	Name string `json:"name" api:"nullable"`
	// Contract owner address
	Owner string `json:"owner" api:"nullable"`
	// Contract owner balance
	OwnerBalance TokenSnapshotDiffResponseItemsMetadataSolanaMetadataOwnerBalance `json:"owner_balance" api:"nullable"`
	// Solana token permanent delegate account
	PermanentDelegate string `json:"permanent_delegate" api:"nullable"`
	// Symbol of the token
	Symbol string `json:"symbol" api:"nullable"`
	// Address of the token creation initiator, only set if the tokens was created by a
	// well known token launch platform
	TokenCreationInitiator string `json:"token_creation_initiator" api:"nullable"`
	// Type of the token
	Type string `json:"type" api:"nullable"`
	// Solana token update authority account
	UpdateAuthority string `json:"update_authority" api:"nullable"`
	// Urls associated with the token
	URLs []string                                                 `json:"urls" api:"nullable"`
	JSON tokenSnapshotDiffResponseItemsMetadataSolanaMetadataJSON `json:"-"`
}

// tokenSnapshotDiffResponseItemsMetadataSolanaMetadataJSON contains the JSON
// metadata for the struct [TokenSnapshotDiffResponseItemsMetadataSolanaMetadata]
type tokenSnapshotDiffResponseItemsMetadataSolanaMetadataJSON struct {
	ContractBalance        apijson.Field
	CreationTimestamp      apijson.Field
	Decimals               apijson.Field
	Deployer               apijson.Field
	DeployerBalance        apijson.Field
	Description            apijson.Field
	ExternalLinks          apijson.Field
	FreezeAuthority        apijson.Field
	ImageURL               apijson.Field
	ImpersonationTargets   apijson.Field
	MaliciousURLs          apijson.Field
	MintAuthority          apijson.Field
	Name                   apijson.Field
	Owner                  apijson.Field
	OwnerBalance           apijson.Field
	PermanentDelegate      apijson.Field
	Symbol                 apijson.Field
	TokenCreationInitiator apijson.Field
	Type                   apijson.Field
	UpdateAuthority        apijson.Field
	URLs                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *TokenSnapshotDiffResponseItemsMetadataSolanaMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotDiffResponseItemsMetadataSolanaMetadataJSON) RawJSON() string {
	return r.raw
}

func (r TokenSnapshotDiffResponseItemsMetadataSolanaMetadata) implementsTokenSnapshotDiffResponseItemsMetadata() {
}

// Contract balance
type TokenSnapshotDiffResponseItemsMetadataSolanaMetadataContractBalance struct {
	Amount    float64                                                                 `json:"amount" api:"nullable"`
	AmountWei string                                                                  `json:"amount_wei" api:"nullable"`
	JSON      tokenSnapshotDiffResponseItemsMetadataSolanaMetadataContractBalanceJSON `json:"-"`
}

// tokenSnapshotDiffResponseItemsMetadataSolanaMetadataContractBalanceJSON contains
// the JSON metadata for the struct
// [TokenSnapshotDiffResponseItemsMetadataSolanaMetadataContractBalance]
type tokenSnapshotDiffResponseItemsMetadataSolanaMetadataContractBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotDiffResponseItemsMetadataSolanaMetadataContractBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotDiffResponseItemsMetadataSolanaMetadataContractBalanceJSON) RawJSON() string {
	return r.raw
}

// Contract creator balance
type TokenSnapshotDiffResponseItemsMetadataSolanaMetadataDeployerBalance struct {
	Amount    float64                                                                 `json:"amount" api:"nullable"`
	AmountWei string                                                                  `json:"amount_wei" api:"nullable"`
	JSON      tokenSnapshotDiffResponseItemsMetadataSolanaMetadataDeployerBalanceJSON `json:"-"`
}

// tokenSnapshotDiffResponseItemsMetadataSolanaMetadataDeployerBalanceJSON contains
// the JSON metadata for the struct
// [TokenSnapshotDiffResponseItemsMetadataSolanaMetadataDeployerBalance]
type tokenSnapshotDiffResponseItemsMetadataSolanaMetadataDeployerBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotDiffResponseItemsMetadataSolanaMetadataDeployerBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotDiffResponseItemsMetadataSolanaMetadataDeployerBalanceJSON) RawJSON() string {
	return r.raw
}

// social links of the token
type TokenSnapshotDiffResponseItemsMetadataSolanaMetadataExternalLinks struct {
	Homepage          string                                                                `json:"homepage" api:"nullable"`
	TelegramChannelID string                                                                `json:"telegram_channel_id" api:"nullable"`
	TwitterPage       string                                                                `json:"twitter_page" api:"nullable"`
	JSON              tokenSnapshotDiffResponseItemsMetadataSolanaMetadataExternalLinksJSON `json:"-"`
}

// tokenSnapshotDiffResponseItemsMetadataSolanaMetadataExternalLinksJSON contains
// the JSON metadata for the struct
// [TokenSnapshotDiffResponseItemsMetadataSolanaMetadataExternalLinks]
type tokenSnapshotDiffResponseItemsMetadataSolanaMetadataExternalLinksJSON struct {
	Homepage          apijson.Field
	TelegramChannelID apijson.Field
	TwitterPage       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TokenSnapshotDiffResponseItemsMetadataSolanaMetadataExternalLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotDiffResponseItemsMetadataSolanaMetadataExternalLinksJSON) RawJSON() string {
	return r.raw
}

type TokenSnapshotDiffResponseItemsMetadataSolanaMetadataImpersonationTarget struct {
	// Address of the token being impersonated
	Address string `json:"address" api:"required"`
	// Blockchain network of the target token
	Chain string `json:"chain" api:"required"`
	// Name of the token being impersonated
	Name string `json:"name" api:"nullable"`
	// Source of the impersonation match
	Source TokenSnapshotDiffResponseItemsMetadataSolanaMetadataImpersonationTargetsSource `json:"source" api:"nullable"`
	// Symbol of the token being impersonated
	Symbol string                                                                      `json:"symbol" api:"nullable"`
	JSON   tokenSnapshotDiffResponseItemsMetadataSolanaMetadataImpersonationTargetJSON `json:"-"`
}

// tokenSnapshotDiffResponseItemsMetadataSolanaMetadataImpersonationTargetJSON
// contains the JSON metadata for the struct
// [TokenSnapshotDiffResponseItemsMetadataSolanaMetadataImpersonationTarget]
type tokenSnapshotDiffResponseItemsMetadataSolanaMetadataImpersonationTargetJSON struct {
	Address     apijson.Field
	Chain       apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotDiffResponseItemsMetadataSolanaMetadataImpersonationTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotDiffResponseItemsMetadataSolanaMetadataImpersonationTargetJSON) RawJSON() string {
	return r.raw
}

// Source of the impersonation match
type TokenSnapshotDiffResponseItemsMetadataSolanaMetadataImpersonationTargetsSource string

const (
	TokenSnapshotDiffResponseItemsMetadataSolanaMetadataImpersonationTargetsSourceTopToken    TokenSnapshotDiffResponseItemsMetadataSolanaMetadataImpersonationTargetsSource = "TOP_TOKEN"
	TokenSnapshotDiffResponseItemsMetadataSolanaMetadataImpersonationTargetsSourceUserDefined TokenSnapshotDiffResponseItemsMetadataSolanaMetadataImpersonationTargetsSource = "USER_DEFINED"
)

func (r TokenSnapshotDiffResponseItemsMetadataSolanaMetadataImpersonationTargetsSource) IsKnown() bool {
	switch r {
	case TokenSnapshotDiffResponseItemsMetadataSolanaMetadataImpersonationTargetsSourceTopToken, TokenSnapshotDiffResponseItemsMetadataSolanaMetadataImpersonationTargetsSourceUserDefined:
		return true
	}
	return false
}

// Contract owner balance
type TokenSnapshotDiffResponseItemsMetadataSolanaMetadataOwnerBalance struct {
	Amount    float64                                                              `json:"amount" api:"nullable"`
	AmountWei string                                                               `json:"amount_wei" api:"nullable"`
	JSON      tokenSnapshotDiffResponseItemsMetadataSolanaMetadataOwnerBalanceJSON `json:"-"`
}

// tokenSnapshotDiffResponseItemsMetadataSolanaMetadataOwnerBalanceJSON contains
// the JSON metadata for the struct
// [TokenSnapshotDiffResponseItemsMetadataSolanaMetadataOwnerBalance]
type tokenSnapshotDiffResponseItemsMetadataSolanaMetadataOwnerBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotDiffResponseItemsMetadataSolanaMetadataOwnerBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotDiffResponseItemsMetadataSolanaMetadataOwnerBalanceJSON) RawJSON() string {
	return r.raw
}

type TokenSnapshotDiffResponseItemsMetadataBitcoinMetadataToken struct {
	// The unique ID for the Rune
	ID string `json:"id" api:"nullable"`
	// Decimals of the token
	Decimals int64 `json:"decimals" api:"nullable"`
	// The formatted name of the rune, with spacers
	FormattedName string `json:"formatted_name" api:"nullable"`
	// Name of the token
	Name string `json:"name" api:"nullable"`
	// The rune's unique sequential number.
	Number int64 `json:"number" api:"nullable"`
	// Symbol of the token
	Symbol string `json:"symbol" api:"nullable"`
	// Type of the token
	Type string                                                         `json:"type" api:"nullable"`
	JSON tokenSnapshotDiffResponseItemsMetadataBitcoinMetadataTokenJSON `json:"-"`
}

// tokenSnapshotDiffResponseItemsMetadataBitcoinMetadataTokenJSON contains the JSON
// metadata for the struct
// [TokenSnapshotDiffResponseItemsMetadataBitcoinMetadataToken]
type tokenSnapshotDiffResponseItemsMetadataBitcoinMetadataTokenJSON struct {
	ID            apijson.Field
	Decimals      apijson.Field
	FormattedName apijson.Field
	Name          apijson.Field
	Number        apijson.Field
	Symbol        apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TokenSnapshotDiffResponseItemsMetadataBitcoinMetadataToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotDiffResponseItemsMetadataBitcoinMetadataTokenJSON) RawJSON() string {
	return r.raw
}

func (r TokenSnapshotDiffResponseItemsMetadataBitcoinMetadataToken) implementsTokenSnapshotDiffResponseItemsMetadata() {
}

type TokenSnapshotDiffResponseItemsMetadataEvmMetadataToken struct {
	// Contract balance
	ContractBalance TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenContractBalance `json:"contract_balance" api:"nullable"`
	// Contract deploy date
	CreationTimestamp string `json:"creation_timestamp" api:"nullable"`
	// Decimals of the token
	Decimals int64 `json:"decimals" api:"nullable"`
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer" api:"nullable"`
	// Contract creator balance
	DeployerBalance TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenDeployerBalance `json:"deployer_balance" api:"nullable"`
	// Description of the token
	Description string `json:"description" api:"nullable"`
	// social links of the token
	ExternalLinks TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenExternalLinks `json:"external_links"`
	// URL of the token image
	ImageURL string `json:"image_url" api:"nullable"`
	// List of tokens that this token is impersonating, if detected as an impersonator
	ImpersonationTargets []TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenImpersonationTarget `json:"impersonation_targets" api:"nullable"`
	// Malicious urls associated with the token
	MaliciousURLs []string `json:"malicious_urls" api:"nullable"`
	// Name of the token
	Name string `json:"name" api:"nullable"`
	// Contract owner address
	Owner string `json:"owner" api:"nullable"`
	// Contract owner balance
	OwnerBalance TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenOwnerBalance `json:"owner_balance" api:"nullable"`
	// Symbol of the token
	Symbol string `json:"symbol" api:"nullable"`
	// Address of the token creation initiator, only set if the tokens was created by a
	// well known token launch platform
	TokenCreationInitiator string `json:"token_creation_initiator" api:"nullable"`
	// Type of the token
	Type string `json:"type" api:"nullable"`
	// Urls associated with the token
	URLs []string                                                   `json:"urls" api:"nullable"`
	JSON tokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenJSON `json:"-"`
}

// tokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenJSON contains the JSON
// metadata for the struct [TokenSnapshotDiffResponseItemsMetadataEvmMetadataToken]
type tokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenJSON struct {
	ContractBalance        apijson.Field
	CreationTimestamp      apijson.Field
	Decimals               apijson.Field
	Deployer               apijson.Field
	DeployerBalance        apijson.Field
	Description            apijson.Field
	ExternalLinks          apijson.Field
	ImageURL               apijson.Field
	ImpersonationTargets   apijson.Field
	MaliciousURLs          apijson.Field
	Name                   apijson.Field
	Owner                  apijson.Field
	OwnerBalance           apijson.Field
	Symbol                 apijson.Field
	TokenCreationInitiator apijson.Field
	Type                   apijson.Field
	URLs                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *TokenSnapshotDiffResponseItemsMetadataEvmMetadataToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenJSON) RawJSON() string {
	return r.raw
}

func (r TokenSnapshotDiffResponseItemsMetadataEvmMetadataToken) implementsTokenSnapshotDiffResponseItemsMetadata() {
}

// Contract balance
type TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenContractBalance struct {
	Amount    float64                                                                   `json:"amount" api:"nullable"`
	AmountWei string                                                                    `json:"amount_wei" api:"nullable"`
	JSON      tokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenContractBalanceJSON `json:"-"`
}

// tokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenContractBalanceJSON
// contains the JSON metadata for the struct
// [TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenContractBalance]
type tokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenContractBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenContractBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenContractBalanceJSON) RawJSON() string {
	return r.raw
}

// Contract creator balance
type TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenDeployerBalance struct {
	Amount    float64                                                                   `json:"amount" api:"nullable"`
	AmountWei string                                                                    `json:"amount_wei" api:"nullable"`
	JSON      tokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenDeployerBalanceJSON `json:"-"`
}

// tokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenDeployerBalanceJSON
// contains the JSON metadata for the struct
// [TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenDeployerBalance]
type tokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenDeployerBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenDeployerBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenDeployerBalanceJSON) RawJSON() string {
	return r.raw
}

// social links of the token
type TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenExternalLinks struct {
	Homepage          string                                                                  `json:"homepage" api:"nullable"`
	TelegramChannelID string                                                                  `json:"telegram_channel_id" api:"nullable"`
	TwitterPage       string                                                                  `json:"twitter_page" api:"nullable"`
	JSON              tokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenExternalLinksJSON `json:"-"`
}

// tokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenExternalLinksJSON contains
// the JSON metadata for the struct
// [TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenExternalLinks]
type tokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenExternalLinksJSON struct {
	Homepage          apijson.Field
	TelegramChannelID apijson.Field
	TwitterPage       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenExternalLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenExternalLinksJSON) RawJSON() string {
	return r.raw
}

type TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenImpersonationTarget struct {
	// Address of the token being impersonated
	Address string `json:"address" api:"required"`
	// Blockchain network of the target token
	Chain string `json:"chain" api:"required"`
	// Name of the token being impersonated
	Name string `json:"name" api:"nullable"`
	// Source of the impersonation match
	Source TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenImpersonationTargetsSource `json:"source" api:"nullable"`
	// Symbol of the token being impersonated
	Symbol string                                                                        `json:"symbol" api:"nullable"`
	JSON   tokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenImpersonationTargetJSON `json:"-"`
}

// tokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenImpersonationTargetJSON
// contains the JSON metadata for the struct
// [TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenImpersonationTarget]
type tokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenImpersonationTargetJSON struct {
	Address     apijson.Field
	Chain       apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenImpersonationTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenImpersonationTargetJSON) RawJSON() string {
	return r.raw
}

// Source of the impersonation match
type TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenImpersonationTargetsSource string

const (
	TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenImpersonationTargetsSourceTopToken    TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenImpersonationTargetsSource = "TOP_TOKEN"
	TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenImpersonationTargetsSourceUserDefined TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenImpersonationTargetsSource = "USER_DEFINED"
)

func (r TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenImpersonationTargetsSource) IsKnown() bool {
	switch r {
	case TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenImpersonationTargetsSourceTopToken, TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenImpersonationTargetsSourceUserDefined:
		return true
	}
	return false
}

// Contract owner balance
type TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenOwnerBalance struct {
	Amount    float64                                                                `json:"amount" api:"nullable"`
	AmountWei string                                                                 `json:"amount_wei" api:"nullable"`
	JSON      tokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenOwnerBalanceJSON `json:"-"`
}

// tokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenOwnerBalanceJSON contains
// the JSON metadata for the struct
// [TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenOwnerBalance]
type tokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenOwnerBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenOwnerBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotDiffResponseItemsMetadataEvmMetadataTokenOwnerBalanceJSON) RawJSON() string {
	return r.raw
}

// General indication
type TokenSnapshotDiffResponseItemsResultType string

const (
	TokenSnapshotDiffResponseItemsResultTypeBenign    TokenSnapshotDiffResponseItemsResultType = "Benign"
	TokenSnapshotDiffResponseItemsResultTypeWarning   TokenSnapshotDiffResponseItemsResultType = "Warning"
	TokenSnapshotDiffResponseItemsResultTypeMalicious TokenSnapshotDiffResponseItemsResultType = "Malicious"
	TokenSnapshotDiffResponseItemsResultTypeSpam      TokenSnapshotDiffResponseItemsResultType = "Spam"
)

func (r TokenSnapshotDiffResponseItemsResultType) IsKnown() bool {
	switch r {
	case TokenSnapshotDiffResponseItemsResultTypeBenign, TokenSnapshotDiffResponseItemsResultTypeWarning, TokenSnapshotDiffResponseItemsResultTypeMalicious, TokenSnapshotDiffResponseItemsResultTypeSpam:
		return true
	}
	return false
}

// Trading limits of the token
type TokenSnapshotDiffResponseItemsTradingLimits struct {
	// Max amount that can be bought at once
	MaxBuy TokenSnapshotDiffResponseItemsTradingLimitsMaxBuy `json:"max_buy" api:"nullable"`
	// Max amount that can be held by a single address
	MaxHolding TokenSnapshotDiffResponseItemsTradingLimitsMaxHolding `json:"max_holding" api:"nullable"`
	// Max amount that can be sold at once
	MaxSell TokenSnapshotDiffResponseItemsTradingLimitsMaxSell `json:"max_sell" api:"nullable"`
	// Maximum amount of the token that can be sold in a block
	SellLimitPerBlock TokenSnapshotDiffResponseItemsTradingLimitsSellLimitPerBlock `json:"sell_limit_per_block" api:"nullable"`
	JSON              tokenSnapshotDiffResponseItemsTradingLimitsJSON              `json:"-"`
}

// tokenSnapshotDiffResponseItemsTradingLimitsJSON contains the JSON metadata for
// the struct [TokenSnapshotDiffResponseItemsTradingLimits]
type tokenSnapshotDiffResponseItemsTradingLimitsJSON struct {
	MaxBuy            apijson.Field
	MaxHolding        apijson.Field
	MaxSell           apijson.Field
	SellLimitPerBlock apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TokenSnapshotDiffResponseItemsTradingLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotDiffResponseItemsTradingLimitsJSON) RawJSON() string {
	return r.raw
}

// Max amount that can be bought at once
type TokenSnapshotDiffResponseItemsTradingLimitsMaxBuy struct {
	Amount    float64                                               `json:"amount" api:"nullable"`
	AmountWei string                                                `json:"amount_wei" api:"nullable"`
	JSON      tokenSnapshotDiffResponseItemsTradingLimitsMaxBuyJSON `json:"-"`
}

// tokenSnapshotDiffResponseItemsTradingLimitsMaxBuyJSON contains the JSON metadata
// for the struct [TokenSnapshotDiffResponseItemsTradingLimitsMaxBuy]
type tokenSnapshotDiffResponseItemsTradingLimitsMaxBuyJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotDiffResponseItemsTradingLimitsMaxBuy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotDiffResponseItemsTradingLimitsMaxBuyJSON) RawJSON() string {
	return r.raw
}

// Max amount that can be held by a single address
type TokenSnapshotDiffResponseItemsTradingLimitsMaxHolding struct {
	Amount    float64                                                   `json:"amount" api:"nullable"`
	AmountWei string                                                    `json:"amount_wei" api:"nullable"`
	JSON      tokenSnapshotDiffResponseItemsTradingLimitsMaxHoldingJSON `json:"-"`
}

// tokenSnapshotDiffResponseItemsTradingLimitsMaxHoldingJSON contains the JSON
// metadata for the struct [TokenSnapshotDiffResponseItemsTradingLimitsMaxHolding]
type tokenSnapshotDiffResponseItemsTradingLimitsMaxHoldingJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotDiffResponseItemsTradingLimitsMaxHolding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotDiffResponseItemsTradingLimitsMaxHoldingJSON) RawJSON() string {
	return r.raw
}

// Max amount that can be sold at once
type TokenSnapshotDiffResponseItemsTradingLimitsMaxSell struct {
	Amount    float64                                                `json:"amount" api:"nullable"`
	AmountWei string                                                 `json:"amount_wei" api:"nullable"`
	JSON      tokenSnapshotDiffResponseItemsTradingLimitsMaxSellJSON `json:"-"`
}

// tokenSnapshotDiffResponseItemsTradingLimitsMaxSellJSON contains the JSON
// metadata for the struct [TokenSnapshotDiffResponseItemsTradingLimitsMaxSell]
type tokenSnapshotDiffResponseItemsTradingLimitsMaxSellJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotDiffResponseItemsTradingLimitsMaxSell) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotDiffResponseItemsTradingLimitsMaxSellJSON) RawJSON() string {
	return r.raw
}

// Maximum amount of the token that can be sold in a block
type TokenSnapshotDiffResponseItemsTradingLimitsSellLimitPerBlock struct {
	Amount    float64                                                          `json:"amount" api:"nullable"`
	AmountWei string                                                           `json:"amount_wei" api:"nullable"`
	JSON      tokenSnapshotDiffResponseItemsTradingLimitsSellLimitPerBlockJSON `json:"-"`
}

// tokenSnapshotDiffResponseItemsTradingLimitsSellLimitPerBlockJSON contains the
// JSON metadata for the struct
// [TokenSnapshotDiffResponseItemsTradingLimitsSellLimitPerBlock]
type tokenSnapshotDiffResponseItemsTradingLimitsSellLimitPerBlockJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotDiffResponseItemsTradingLimitsSellLimitPerBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotDiffResponseItemsTradingLimitsSellLimitPerBlockJSON) RawJSON() string {
	return r.raw
}

type TokenSnapshotDiffResponseItemsFeature struct {
	// Description of the feature
	Description string `json:"description" api:"required"`
	// Feature identifier
	FeatureID TokenSnapshotDiffResponseItemsFeaturesFeatureID `json:"feature_id" api:"required"`
	// Type of the feature
	Type TokenSnapshotDiffResponseItemsFeaturesType `json:"type" api:"required"`
	JSON tokenSnapshotDiffResponseItemsFeatureJSON  `json:"-"`
}

// tokenSnapshotDiffResponseItemsFeatureJSON contains the JSON metadata for the
// struct [TokenSnapshotDiffResponseItemsFeature]
type tokenSnapshotDiffResponseItemsFeatureJSON struct {
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotDiffResponseItemsFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotDiffResponseItemsFeatureJSON) RawJSON() string {
	return r.raw
}

// Feature identifier
type TokenSnapshotDiffResponseItemsFeaturesFeatureID string

const (
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDVerifiedContract               TokenSnapshotDiffResponseItemsFeaturesFeatureID = "VERIFIED_CONTRACT"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDUnverifiedContract             TokenSnapshotDiffResponseItemsFeaturesFeatureID = "UNVERIFIED_CONTRACT"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDHighTradeVolume                TokenSnapshotDiffResponseItemsFeaturesFeatureID = "HIGH_TRADE_VOLUME"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDMarketPlaceSalesHistory        TokenSnapshotDiffResponseItemsFeaturesFeatureID = "MARKET_PLACE_SALES_HISTORY"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDHighReputationToken            TokenSnapshotDiffResponseItemsFeaturesFeatureID = "HIGH_REPUTATION_TOKEN"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDOnchainActivityValidator       TokenSnapshotDiffResponseItemsFeaturesFeatureID = "ONCHAIN_ACTIVITY_VALIDATOR"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDStaticCodeSignature            TokenSnapshotDiffResponseItemsFeaturesFeatureID = "STATIC_CODE_SIGNATURE"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDKnownMalicious                 TokenSnapshotDiffResponseItemsFeaturesFeatureID = "KNOWN_MALICIOUS"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDIsEoa                          TokenSnapshotDiffResponseItemsFeaturesFeatureID = "IS_EOA"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDIsContract                     TokenSnapshotDiffResponseItemsFeaturesFeatureID = "IS_CONTRACT"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDErc20Contract                  TokenSnapshotDiffResponseItemsFeaturesFeatureID = "ERC20_CONTRACT"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDTrustedContract                TokenSnapshotDiffResponseItemsFeaturesFeatureID = "TRUSTED_CONTRACT"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDBenignCreator                  TokenSnapshotDiffResponseItemsFeaturesFeatureID = "BENIGN_CREATOR"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDMetadata                       TokenSnapshotDiffResponseItemsFeaturesFeatureID = "METADATA"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDAirdropPattern                 TokenSnapshotDiffResponseItemsFeaturesFeatureID = "AIRDROP_PATTERN"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDImpersonator                   TokenSnapshotDiffResponseItemsFeaturesFeatureID = "IMPERSONATOR"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDInorganicVolume                TokenSnapshotDiffResponseItemsFeaturesFeatureID = "INORGANIC_VOLUME"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDDynamicAnalysis                TokenSnapshotDiffResponseItemsFeaturesFeatureID = "DYNAMIC_ANALYSIS"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDConcentratedSupplyDistribution TokenSnapshotDiffResponseItemsFeaturesFeatureID = "CONCENTRATED_SUPPLY_DISTRIBUTION"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDHoneypot                       TokenSnapshotDiffResponseItemsFeaturesFeatureID = "HONEYPOT"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDInsufficientLockedLiquidity    TokenSnapshotDiffResponseItemsFeaturesFeatureID = "INSUFFICIENT_LOCKED_LIQUIDITY"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDUnstableTokenPrice             TokenSnapshotDiffResponseItemsFeaturesFeatureID = "UNSTABLE_TOKEN_PRICE"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDRugpull                        TokenSnapshotDiffResponseItemsFeaturesFeatureID = "RUGPULL"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDWashTrading                    TokenSnapshotDiffResponseItemsFeaturesFeatureID = "WASH_TRADING"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDConsumerOverride               TokenSnapshotDiffResponseItemsFeaturesFeatureID = "CONSUMER_OVERRIDE"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDInappropriateContent           TokenSnapshotDiffResponseItemsFeaturesFeatureID = "INAPPROPRIATE_CONTENT"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDHighTransferFee                TokenSnapshotDiffResponseItemsFeaturesFeatureID = "HIGH_TRANSFER_FEE"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDHighBuyFee                     TokenSnapshotDiffResponseItemsFeaturesFeatureID = "HIGH_BUY_FEE"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDHighSellFee                    TokenSnapshotDiffResponseItemsFeaturesFeatureID = "HIGH_SELL_FEE"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDUnsellableToken                TokenSnapshotDiffResponseItemsFeaturesFeatureID = "UNSELLABLE_TOKEN"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDIsMintable                     TokenSnapshotDiffResponseItemsFeaturesFeatureID = "IS_MINTABLE"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDRebaseToken                    TokenSnapshotDiffResponseItemsFeaturesFeatureID = "REBASE_TOKEN"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDLiquidStakingToken             TokenSnapshotDiffResponseItemsFeaturesFeatureID = "LIQUID_STAKING_TOKEN"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDModifiableTaxes                TokenSnapshotDiffResponseItemsFeaturesFeatureID = "MODIFIABLE_TAXES"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDCanBlacklist                   TokenSnapshotDiffResponseItemsFeaturesFeatureID = "CAN_BLACKLIST"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDCanWhitelist                   TokenSnapshotDiffResponseItemsFeaturesFeatureID = "CAN_WHITELIST"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDHasTradingCooldown             TokenSnapshotDiffResponseItemsFeaturesFeatureID = "HAS_TRADING_COOLDOWN"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDExternalFunctions              TokenSnapshotDiffResponseItemsFeaturesFeatureID = "EXTERNAL_FUNCTIONS"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDHiddenOwner                    TokenSnapshotDiffResponseItemsFeaturesFeatureID = "HIDDEN_OWNER"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDTransferPauseable              TokenSnapshotDiffResponseItemsFeaturesFeatureID = "TRANSFER_PAUSEABLE"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDOwnershipRenounced             TokenSnapshotDiffResponseItemsFeaturesFeatureID = "OWNERSHIP_RENOUNCED"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDOwnerCanChangeBalance          TokenSnapshotDiffResponseItemsFeaturesFeatureID = "OWNER_CAN_CHANGE_BALANCE"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDProxyContract                  TokenSnapshotDiffResponseItemsFeaturesFeatureID = "PROXY_CONTRACT"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDSimilarMaliciousContract       TokenSnapshotDiffResponseItemsFeaturesFeatureID = "SIMILAR_MALICIOUS_CONTRACT"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDImpersonatorSensitiveAsset     TokenSnapshotDiffResponseItemsFeaturesFeatureID = "IMPERSONATOR_SENSITIVE_ASSET"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDImpersonatorHighConfidence     TokenSnapshotDiffResponseItemsFeaturesFeatureID = "IMPERSONATOR_HIGH_CONFIDENCE"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDImpersonatorMediumConfidence   TokenSnapshotDiffResponseItemsFeaturesFeatureID = "IMPERSONATOR_MEDIUM_CONFIDENCE"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDImpersonatorLowConfidence      TokenSnapshotDiffResponseItemsFeaturesFeatureID = "IMPERSONATOR_LOW_CONFIDENCE"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDImpersonationProtected         TokenSnapshotDiffResponseItemsFeaturesFeatureID = "IMPERSONATION_PROTECTED"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDFakeVolume                     TokenSnapshotDiffResponseItemsFeaturesFeatureID = "FAKE_VOLUME"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDHiddenSupplyByKeyHolder        TokenSnapshotDiffResponseItemsFeaturesFeatureID = "HIDDEN_SUPPLY_BY_KEY_HOLDER"
	TokenSnapshotDiffResponseItemsFeaturesFeatureIDFakeTradeMakerCount            TokenSnapshotDiffResponseItemsFeaturesFeatureID = "FAKE_TRADE_MAKER_COUNT"
)

func (r TokenSnapshotDiffResponseItemsFeaturesFeatureID) IsKnown() bool {
	switch r {
	case TokenSnapshotDiffResponseItemsFeaturesFeatureIDVerifiedContract, TokenSnapshotDiffResponseItemsFeaturesFeatureIDUnverifiedContract, TokenSnapshotDiffResponseItemsFeaturesFeatureIDHighTradeVolume, TokenSnapshotDiffResponseItemsFeaturesFeatureIDMarketPlaceSalesHistory, TokenSnapshotDiffResponseItemsFeaturesFeatureIDHighReputationToken, TokenSnapshotDiffResponseItemsFeaturesFeatureIDOnchainActivityValidator, TokenSnapshotDiffResponseItemsFeaturesFeatureIDStaticCodeSignature, TokenSnapshotDiffResponseItemsFeaturesFeatureIDKnownMalicious, TokenSnapshotDiffResponseItemsFeaturesFeatureIDIsEoa, TokenSnapshotDiffResponseItemsFeaturesFeatureIDIsContract, TokenSnapshotDiffResponseItemsFeaturesFeatureIDErc20Contract, TokenSnapshotDiffResponseItemsFeaturesFeatureIDTrustedContract, TokenSnapshotDiffResponseItemsFeaturesFeatureIDBenignCreator, TokenSnapshotDiffResponseItemsFeaturesFeatureIDMetadata, TokenSnapshotDiffResponseItemsFeaturesFeatureIDAirdropPattern, TokenSnapshotDiffResponseItemsFeaturesFeatureIDImpersonator, TokenSnapshotDiffResponseItemsFeaturesFeatureIDInorganicVolume, TokenSnapshotDiffResponseItemsFeaturesFeatureIDDynamicAnalysis, TokenSnapshotDiffResponseItemsFeaturesFeatureIDConcentratedSupplyDistribution, TokenSnapshotDiffResponseItemsFeaturesFeatureIDHoneypot, TokenSnapshotDiffResponseItemsFeaturesFeatureIDInsufficientLockedLiquidity, TokenSnapshotDiffResponseItemsFeaturesFeatureIDUnstableTokenPrice, TokenSnapshotDiffResponseItemsFeaturesFeatureIDRugpull, TokenSnapshotDiffResponseItemsFeaturesFeatureIDWashTrading, TokenSnapshotDiffResponseItemsFeaturesFeatureIDConsumerOverride, TokenSnapshotDiffResponseItemsFeaturesFeatureIDInappropriateContent, TokenSnapshotDiffResponseItemsFeaturesFeatureIDHighTransferFee, TokenSnapshotDiffResponseItemsFeaturesFeatureIDHighBuyFee, TokenSnapshotDiffResponseItemsFeaturesFeatureIDHighSellFee, TokenSnapshotDiffResponseItemsFeaturesFeatureIDUnsellableToken, TokenSnapshotDiffResponseItemsFeaturesFeatureIDIsMintable, TokenSnapshotDiffResponseItemsFeaturesFeatureIDRebaseToken, TokenSnapshotDiffResponseItemsFeaturesFeatureIDLiquidStakingToken, TokenSnapshotDiffResponseItemsFeaturesFeatureIDModifiableTaxes, TokenSnapshotDiffResponseItemsFeaturesFeatureIDCanBlacklist, TokenSnapshotDiffResponseItemsFeaturesFeatureIDCanWhitelist, TokenSnapshotDiffResponseItemsFeaturesFeatureIDHasTradingCooldown, TokenSnapshotDiffResponseItemsFeaturesFeatureIDExternalFunctions, TokenSnapshotDiffResponseItemsFeaturesFeatureIDHiddenOwner, TokenSnapshotDiffResponseItemsFeaturesFeatureIDTransferPauseable, TokenSnapshotDiffResponseItemsFeaturesFeatureIDOwnershipRenounced, TokenSnapshotDiffResponseItemsFeaturesFeatureIDOwnerCanChangeBalance, TokenSnapshotDiffResponseItemsFeaturesFeatureIDProxyContract, TokenSnapshotDiffResponseItemsFeaturesFeatureIDSimilarMaliciousContract, TokenSnapshotDiffResponseItemsFeaturesFeatureIDImpersonatorSensitiveAsset, TokenSnapshotDiffResponseItemsFeaturesFeatureIDImpersonatorHighConfidence, TokenSnapshotDiffResponseItemsFeaturesFeatureIDImpersonatorMediumConfidence, TokenSnapshotDiffResponseItemsFeaturesFeatureIDImpersonatorLowConfidence, TokenSnapshotDiffResponseItemsFeaturesFeatureIDImpersonationProtected, TokenSnapshotDiffResponseItemsFeaturesFeatureIDFakeVolume, TokenSnapshotDiffResponseItemsFeaturesFeatureIDHiddenSupplyByKeyHolder, TokenSnapshotDiffResponseItemsFeaturesFeatureIDFakeTradeMakerCount:
		return true
	}
	return false
}

// Type of the feature
type TokenSnapshotDiffResponseItemsFeaturesType string

const (
	TokenSnapshotDiffResponseItemsFeaturesTypeBenign    TokenSnapshotDiffResponseItemsFeaturesType = "Benign"
	TokenSnapshotDiffResponseItemsFeaturesTypeInfo      TokenSnapshotDiffResponseItemsFeaturesType = "Info"
	TokenSnapshotDiffResponseItemsFeaturesTypeWarning   TokenSnapshotDiffResponseItemsFeaturesType = "Warning"
	TokenSnapshotDiffResponseItemsFeaturesTypeMalicious TokenSnapshotDiffResponseItemsFeaturesType = "Malicious"
)

func (r TokenSnapshotDiffResponseItemsFeaturesType) IsKnown() bool {
	switch r {
	case TokenSnapshotDiffResponseItemsFeaturesTypeBenign, TokenSnapshotDiffResponseItemsFeaturesTypeInfo, TokenSnapshotDiffResponseItemsFeaturesTypeWarning, TokenSnapshotDiffResponseItemsFeaturesTypeMalicious:
		return true
	}
	return false
}

type TokenSnapshotFullResponse struct {
	// Cursor to refetch the current page
	CurrentPage string                          `json:"current_page" api:"required"`
	Items       []TokenSnapshotFullResponseItem `json:"items" api:"required"`
	// Cursor for the next page
	NextPage string `json:"next_page" api:"nullable"`
	// Cursor for the previous page
	PreviousPage string                        `json:"previous_page" api:"nullable"`
	JSON         tokenSnapshotFullResponseJSON `json:"-"`
}

// tokenSnapshotFullResponseJSON contains the JSON metadata for the struct
// [TokenSnapshotFullResponse]
type tokenSnapshotFullResponseJSON struct {
	CurrentPage  apijson.Field
	Items        apijson.Field
	NextPage     apijson.Field
	PreviousPage apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TokenSnapshotFullResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotFullResponseJSON) RawJSON() string {
	return r.raw
}

type TokenSnapshotFullResponseItem struct {
	// Token address to validate (EVM / Solana)
	Address string `json:"address" api:"required"`
	// Dictionary of detected attacks found during the scan
	AttackTypes map[string]TokenSnapshotFullResponseItemsAttackType `json:"attack_types" api:"required"`
	// Blockchain network
	Chain TokenScanSupportedChain `json:"chain" api:"required"`
	// Fees associated with the token
	Fees TokenSnapshotFullResponseItemsFees `json:"fees" api:"required"`
	// financial stats of the token
	FinancialStats FinancialStats `json:"financial_stats" api:"required"`
	// Score between 0 to 1 (double)
	MaliciousScore string `json:"malicious_score" api:"required"`
	// Metadata of the token
	Metadata TokenSnapshotFullResponseItemsMetadata `json:"metadata" api:"required"`
	// General indication
	ResultType TokenSnapshotFullResponseItemsResultType `json:"result_type" api:"required"`
	// Trading limits of the token
	TradingLimits TokenSnapshotFullResponseItemsTradingLimits `json:"trading_limits" api:"required"`
	// List of features associated with the token
	Features []TokenSnapshotFullResponseItemsFeature `json:"features"`
	JSON     tokenSnapshotFullResponseItemJSON       `json:"-"`
}

// tokenSnapshotFullResponseItemJSON contains the JSON metadata for the struct
// [TokenSnapshotFullResponseItem]
type tokenSnapshotFullResponseItemJSON struct {
	Address        apijson.Field
	AttackTypes    apijson.Field
	Chain          apijson.Field
	Fees           apijson.Field
	FinancialStats apijson.Field
	MaliciousScore apijson.Field
	Metadata       apijson.Field
	ResultType     apijson.Field
	TradingLimits  apijson.Field
	Features       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TokenSnapshotFullResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotFullResponseItemJSON) RawJSON() string {
	return r.raw
}

type TokenSnapshotFullResponseItemsAttackType struct {
	// Score between 0 to 1 (double) that indicates the assurance this attack happened
	Score string `json:"score" api:"required"`
	// Object contains an extra information related to the attack
	Features interface{} `json:"features"`
	// If score is higher or equal to this field, the token is using this attack type
	Threshold string                                       `json:"threshold"`
	JSON      tokenSnapshotFullResponseItemsAttackTypeJSON `json:"-"`
}

// tokenSnapshotFullResponseItemsAttackTypeJSON contains the JSON metadata for the
// struct [TokenSnapshotFullResponseItemsAttackType]
type tokenSnapshotFullResponseItemsAttackTypeJSON struct {
	Score       apijson.Field
	Features    apijson.Field
	Threshold   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotFullResponseItemsAttackType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotFullResponseItemsAttackTypeJSON) RawJSON() string {
	return r.raw
}

// Fees associated with the token
type TokenSnapshotFullResponseItemsFees struct {
	// Buy fee of the token
	Buy float64 `json:"buy" api:"nullable"`
	// Sell fee of the token
	Sell float64 `json:"sell" api:"nullable"`
	// Transfer fee of the token
	Transfer float64 `json:"transfer" api:"nullable"`
	// The maximum value that a transfer fee will cost
	TransferFeeMaxAmount int64                                  `json:"transfer_fee_max_amount" api:"nullable"`
	JSON                 tokenSnapshotFullResponseItemsFeesJSON `json:"-"`
}

// tokenSnapshotFullResponseItemsFeesJSON contains the JSON metadata for the struct
// [TokenSnapshotFullResponseItemsFees]
type tokenSnapshotFullResponseItemsFeesJSON struct {
	Buy                  apijson.Field
	Sell                 apijson.Field
	Transfer             apijson.Field
	TransferFeeMaxAmount apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TokenSnapshotFullResponseItemsFees) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotFullResponseItemsFeesJSON) RawJSON() string {
	return r.raw
}

// Metadata of the token
type TokenSnapshotFullResponseItemsMetadata struct {
	// The unique ID for the Rune
	ID string `json:"id" api:"nullable"`
	// This field can have the runtime type of
	// [TokenSnapshotFullResponseItemsMetadataSolanaMetadataContractBalance],
	// [TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenContractBalance].
	ContractBalance interface{} `json:"contract_balance"`
	// Contract deploy date
	CreationTimestamp string `json:"creation_timestamp" api:"nullable"`
	// Decimals of the token
	Decimals int64 `json:"decimals" api:"nullable"`
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer" api:"nullable"`
	// This field can have the runtime type of
	// [TokenSnapshotFullResponseItemsMetadataSolanaMetadataDeployerBalance],
	// [TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenDeployerBalance].
	DeployerBalance interface{} `json:"deployer_balance"`
	// Description of the token
	Description string `json:"description" api:"nullable"`
	// This field can have the runtime type of
	// [TokenSnapshotFullResponseItemsMetadataSolanaMetadataExternalLinks],
	// [TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenExternalLinks].
	ExternalLinks interface{} `json:"external_links"`
	// The formatted name of the rune, with spacers
	FormattedName string `json:"formatted_name" api:"nullable"`
	// Solana token freeze authority account
	FreezeAuthority string `json:"freeze_authority" api:"nullable"`
	// URL of the token image
	ImageURL string `json:"image_url" api:"nullable"`
	// This field can have the runtime type of
	// [[]TokenSnapshotFullResponseItemsMetadataSolanaMetadataImpersonationTarget],
	// [[]TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenImpersonationTarget].
	ImpersonationTargets interface{} `json:"impersonation_targets"`
	// This field can have the runtime type of [[]string].
	MaliciousURLs interface{} `json:"malicious_urls"`
	// Solana token mint authority account
	MintAuthority string `json:"mint_authority" api:"nullable"`
	// Name of the token
	Name string `json:"name" api:"nullable"`
	// The rune's unique sequential number.
	Number int64 `json:"number" api:"nullable"`
	// Contract owner address
	Owner string `json:"owner" api:"nullable"`
	// This field can have the runtime type of
	// [TokenSnapshotFullResponseItemsMetadataSolanaMetadataOwnerBalance],
	// [TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenOwnerBalance].
	OwnerBalance interface{} `json:"owner_balance"`
	// Solana token permanent delegate account
	PermanentDelegate string `json:"permanent_delegate" api:"nullable"`
	// Symbol of the token
	Symbol string `json:"symbol" api:"nullable"`
	// Address of the token creation initiator, only set if the tokens was created by a
	// well known token launch platform
	TokenCreationInitiator string `json:"token_creation_initiator" api:"nullable"`
	// Type of the token
	Type string `json:"type" api:"nullable"`
	// Solana token update authority account
	UpdateAuthority string `json:"update_authority" api:"nullable"`
	// This field can have the runtime type of [[]string].
	URLs  interface{}                                `json:"urls"`
	JSON  tokenSnapshotFullResponseItemsMetadataJSON `json:"-"`
	union TokenSnapshotFullResponseItemsMetadataUnion
}

// tokenSnapshotFullResponseItemsMetadataJSON contains the JSON metadata for the
// struct [TokenSnapshotFullResponseItemsMetadata]
type tokenSnapshotFullResponseItemsMetadataJSON struct {
	ID                     apijson.Field
	ContractBalance        apijson.Field
	CreationTimestamp      apijson.Field
	Decimals               apijson.Field
	Deployer               apijson.Field
	DeployerBalance        apijson.Field
	Description            apijson.Field
	ExternalLinks          apijson.Field
	FormattedName          apijson.Field
	FreezeAuthority        apijson.Field
	ImageURL               apijson.Field
	ImpersonationTargets   apijson.Field
	MaliciousURLs          apijson.Field
	MintAuthority          apijson.Field
	Name                   apijson.Field
	Number                 apijson.Field
	Owner                  apijson.Field
	OwnerBalance           apijson.Field
	PermanentDelegate      apijson.Field
	Symbol                 apijson.Field
	TokenCreationInitiator apijson.Field
	Type                   apijson.Field
	UpdateAuthority        apijson.Field
	URLs                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r tokenSnapshotFullResponseItemsMetadataJSON) RawJSON() string {
	return r.raw
}

func (r *TokenSnapshotFullResponseItemsMetadata) UnmarshalJSON(data []byte) (err error) {
	*r = TokenSnapshotFullResponseItemsMetadata{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TokenSnapshotFullResponseItemsMetadataUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TokenSnapshotFullResponseItemsMetadataSolanaMetadata],
// [TokenSnapshotFullResponseItemsMetadataBitcoinMetadataToken],
// [TokenSnapshotFullResponseItemsMetadataEvmMetadataToken].
func (r TokenSnapshotFullResponseItemsMetadata) AsUnion() TokenSnapshotFullResponseItemsMetadataUnion {
	return r.union
}

// Metadata of the token
//
// Union satisfied by [TokenSnapshotFullResponseItemsMetadataSolanaMetadata],
// [TokenSnapshotFullResponseItemsMetadataBitcoinMetadataToken] or
// [TokenSnapshotFullResponseItemsMetadataEvmMetadataToken].
type TokenSnapshotFullResponseItemsMetadataUnion interface {
	implementsTokenSnapshotFullResponseItemsMetadata()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TokenSnapshotFullResponseItemsMetadataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenSnapshotFullResponseItemsMetadataSolanaMetadata{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenSnapshotFullResponseItemsMetadataBitcoinMetadataToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenSnapshotFullResponseItemsMetadataEvmMetadataToken{}),
		},
	)
}

type TokenSnapshotFullResponseItemsMetadataSolanaMetadata struct {
	// Contract balance
	ContractBalance TokenSnapshotFullResponseItemsMetadataSolanaMetadataContractBalance `json:"contract_balance" api:"nullable"`
	// Contract deploy date
	CreationTimestamp string `json:"creation_timestamp" api:"nullable"`
	// Decimals of the token
	Decimals int64 `json:"decimals" api:"nullable"`
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer" api:"nullable"`
	// Contract creator balance
	DeployerBalance TokenSnapshotFullResponseItemsMetadataSolanaMetadataDeployerBalance `json:"deployer_balance" api:"nullable"`
	// Description of the token
	Description string `json:"description" api:"nullable"`
	// social links of the token
	ExternalLinks TokenSnapshotFullResponseItemsMetadataSolanaMetadataExternalLinks `json:"external_links"`
	// Solana token freeze authority account
	FreezeAuthority string `json:"freeze_authority" api:"nullable"`
	// URL of the token image
	ImageURL string `json:"image_url" api:"nullable"`
	// List of tokens that this token is impersonating, if detected as an impersonator
	ImpersonationTargets []TokenSnapshotFullResponseItemsMetadataSolanaMetadataImpersonationTarget `json:"impersonation_targets" api:"nullable"`
	// Malicious urls associated with the token
	MaliciousURLs []string `json:"malicious_urls" api:"nullable"`
	// Solana token mint authority account
	MintAuthority string `json:"mint_authority" api:"nullable"`
	// Name of the token
	Name string `json:"name" api:"nullable"`
	// Contract owner address
	Owner string `json:"owner" api:"nullable"`
	// Contract owner balance
	OwnerBalance TokenSnapshotFullResponseItemsMetadataSolanaMetadataOwnerBalance `json:"owner_balance" api:"nullable"`
	// Solana token permanent delegate account
	PermanentDelegate string `json:"permanent_delegate" api:"nullable"`
	// Symbol of the token
	Symbol string `json:"symbol" api:"nullable"`
	// Address of the token creation initiator, only set if the tokens was created by a
	// well known token launch platform
	TokenCreationInitiator string `json:"token_creation_initiator" api:"nullable"`
	// Type of the token
	Type string `json:"type" api:"nullable"`
	// Solana token update authority account
	UpdateAuthority string `json:"update_authority" api:"nullable"`
	// Urls associated with the token
	URLs []string                                                 `json:"urls" api:"nullable"`
	JSON tokenSnapshotFullResponseItemsMetadataSolanaMetadataJSON `json:"-"`
}

// tokenSnapshotFullResponseItemsMetadataSolanaMetadataJSON contains the JSON
// metadata for the struct [TokenSnapshotFullResponseItemsMetadataSolanaMetadata]
type tokenSnapshotFullResponseItemsMetadataSolanaMetadataJSON struct {
	ContractBalance        apijson.Field
	CreationTimestamp      apijson.Field
	Decimals               apijson.Field
	Deployer               apijson.Field
	DeployerBalance        apijson.Field
	Description            apijson.Field
	ExternalLinks          apijson.Field
	FreezeAuthority        apijson.Field
	ImageURL               apijson.Field
	ImpersonationTargets   apijson.Field
	MaliciousURLs          apijson.Field
	MintAuthority          apijson.Field
	Name                   apijson.Field
	Owner                  apijson.Field
	OwnerBalance           apijson.Field
	PermanentDelegate      apijson.Field
	Symbol                 apijson.Field
	TokenCreationInitiator apijson.Field
	Type                   apijson.Field
	UpdateAuthority        apijson.Field
	URLs                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *TokenSnapshotFullResponseItemsMetadataSolanaMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotFullResponseItemsMetadataSolanaMetadataJSON) RawJSON() string {
	return r.raw
}

func (r TokenSnapshotFullResponseItemsMetadataSolanaMetadata) implementsTokenSnapshotFullResponseItemsMetadata() {
}

// Contract balance
type TokenSnapshotFullResponseItemsMetadataSolanaMetadataContractBalance struct {
	Amount    float64                                                                 `json:"amount" api:"nullable"`
	AmountWei string                                                                  `json:"amount_wei" api:"nullable"`
	JSON      tokenSnapshotFullResponseItemsMetadataSolanaMetadataContractBalanceJSON `json:"-"`
}

// tokenSnapshotFullResponseItemsMetadataSolanaMetadataContractBalanceJSON contains
// the JSON metadata for the struct
// [TokenSnapshotFullResponseItemsMetadataSolanaMetadataContractBalance]
type tokenSnapshotFullResponseItemsMetadataSolanaMetadataContractBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotFullResponseItemsMetadataSolanaMetadataContractBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotFullResponseItemsMetadataSolanaMetadataContractBalanceJSON) RawJSON() string {
	return r.raw
}

// Contract creator balance
type TokenSnapshotFullResponseItemsMetadataSolanaMetadataDeployerBalance struct {
	Amount    float64                                                                 `json:"amount" api:"nullable"`
	AmountWei string                                                                  `json:"amount_wei" api:"nullable"`
	JSON      tokenSnapshotFullResponseItemsMetadataSolanaMetadataDeployerBalanceJSON `json:"-"`
}

// tokenSnapshotFullResponseItemsMetadataSolanaMetadataDeployerBalanceJSON contains
// the JSON metadata for the struct
// [TokenSnapshotFullResponseItemsMetadataSolanaMetadataDeployerBalance]
type tokenSnapshotFullResponseItemsMetadataSolanaMetadataDeployerBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotFullResponseItemsMetadataSolanaMetadataDeployerBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotFullResponseItemsMetadataSolanaMetadataDeployerBalanceJSON) RawJSON() string {
	return r.raw
}

// social links of the token
type TokenSnapshotFullResponseItemsMetadataSolanaMetadataExternalLinks struct {
	Homepage          string                                                                `json:"homepage" api:"nullable"`
	TelegramChannelID string                                                                `json:"telegram_channel_id" api:"nullable"`
	TwitterPage       string                                                                `json:"twitter_page" api:"nullable"`
	JSON              tokenSnapshotFullResponseItemsMetadataSolanaMetadataExternalLinksJSON `json:"-"`
}

// tokenSnapshotFullResponseItemsMetadataSolanaMetadataExternalLinksJSON contains
// the JSON metadata for the struct
// [TokenSnapshotFullResponseItemsMetadataSolanaMetadataExternalLinks]
type tokenSnapshotFullResponseItemsMetadataSolanaMetadataExternalLinksJSON struct {
	Homepage          apijson.Field
	TelegramChannelID apijson.Field
	TwitterPage       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TokenSnapshotFullResponseItemsMetadataSolanaMetadataExternalLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotFullResponseItemsMetadataSolanaMetadataExternalLinksJSON) RawJSON() string {
	return r.raw
}

type TokenSnapshotFullResponseItemsMetadataSolanaMetadataImpersonationTarget struct {
	// Address of the token being impersonated
	Address string `json:"address" api:"required"`
	// Blockchain network of the target token
	Chain string `json:"chain" api:"required"`
	// Name of the token being impersonated
	Name string `json:"name" api:"nullable"`
	// Source of the impersonation match
	Source TokenSnapshotFullResponseItemsMetadataSolanaMetadataImpersonationTargetsSource `json:"source" api:"nullable"`
	// Symbol of the token being impersonated
	Symbol string                                                                      `json:"symbol" api:"nullable"`
	JSON   tokenSnapshotFullResponseItemsMetadataSolanaMetadataImpersonationTargetJSON `json:"-"`
}

// tokenSnapshotFullResponseItemsMetadataSolanaMetadataImpersonationTargetJSON
// contains the JSON metadata for the struct
// [TokenSnapshotFullResponseItemsMetadataSolanaMetadataImpersonationTarget]
type tokenSnapshotFullResponseItemsMetadataSolanaMetadataImpersonationTargetJSON struct {
	Address     apijson.Field
	Chain       apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotFullResponseItemsMetadataSolanaMetadataImpersonationTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotFullResponseItemsMetadataSolanaMetadataImpersonationTargetJSON) RawJSON() string {
	return r.raw
}

// Source of the impersonation match
type TokenSnapshotFullResponseItemsMetadataSolanaMetadataImpersonationTargetsSource string

const (
	TokenSnapshotFullResponseItemsMetadataSolanaMetadataImpersonationTargetsSourceTopToken    TokenSnapshotFullResponseItemsMetadataSolanaMetadataImpersonationTargetsSource = "TOP_TOKEN"
	TokenSnapshotFullResponseItemsMetadataSolanaMetadataImpersonationTargetsSourceUserDefined TokenSnapshotFullResponseItemsMetadataSolanaMetadataImpersonationTargetsSource = "USER_DEFINED"
)

func (r TokenSnapshotFullResponseItemsMetadataSolanaMetadataImpersonationTargetsSource) IsKnown() bool {
	switch r {
	case TokenSnapshotFullResponseItemsMetadataSolanaMetadataImpersonationTargetsSourceTopToken, TokenSnapshotFullResponseItemsMetadataSolanaMetadataImpersonationTargetsSourceUserDefined:
		return true
	}
	return false
}

// Contract owner balance
type TokenSnapshotFullResponseItemsMetadataSolanaMetadataOwnerBalance struct {
	Amount    float64                                                              `json:"amount" api:"nullable"`
	AmountWei string                                                               `json:"amount_wei" api:"nullable"`
	JSON      tokenSnapshotFullResponseItemsMetadataSolanaMetadataOwnerBalanceJSON `json:"-"`
}

// tokenSnapshotFullResponseItemsMetadataSolanaMetadataOwnerBalanceJSON contains
// the JSON metadata for the struct
// [TokenSnapshotFullResponseItemsMetadataSolanaMetadataOwnerBalance]
type tokenSnapshotFullResponseItemsMetadataSolanaMetadataOwnerBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotFullResponseItemsMetadataSolanaMetadataOwnerBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotFullResponseItemsMetadataSolanaMetadataOwnerBalanceJSON) RawJSON() string {
	return r.raw
}

type TokenSnapshotFullResponseItemsMetadataBitcoinMetadataToken struct {
	// The unique ID for the Rune
	ID string `json:"id" api:"nullable"`
	// Decimals of the token
	Decimals int64 `json:"decimals" api:"nullable"`
	// The formatted name of the rune, with spacers
	FormattedName string `json:"formatted_name" api:"nullable"`
	// Name of the token
	Name string `json:"name" api:"nullable"`
	// The rune's unique sequential number.
	Number int64 `json:"number" api:"nullable"`
	// Symbol of the token
	Symbol string `json:"symbol" api:"nullable"`
	// Type of the token
	Type string                                                         `json:"type" api:"nullable"`
	JSON tokenSnapshotFullResponseItemsMetadataBitcoinMetadataTokenJSON `json:"-"`
}

// tokenSnapshotFullResponseItemsMetadataBitcoinMetadataTokenJSON contains the JSON
// metadata for the struct
// [TokenSnapshotFullResponseItemsMetadataBitcoinMetadataToken]
type tokenSnapshotFullResponseItemsMetadataBitcoinMetadataTokenJSON struct {
	ID            apijson.Field
	Decimals      apijson.Field
	FormattedName apijson.Field
	Name          apijson.Field
	Number        apijson.Field
	Symbol        apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TokenSnapshotFullResponseItemsMetadataBitcoinMetadataToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotFullResponseItemsMetadataBitcoinMetadataTokenJSON) RawJSON() string {
	return r.raw
}

func (r TokenSnapshotFullResponseItemsMetadataBitcoinMetadataToken) implementsTokenSnapshotFullResponseItemsMetadata() {
}

type TokenSnapshotFullResponseItemsMetadataEvmMetadataToken struct {
	// Contract balance
	ContractBalance TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenContractBalance `json:"contract_balance" api:"nullable"`
	// Contract deploy date
	CreationTimestamp string `json:"creation_timestamp" api:"nullable"`
	// Decimals of the token
	Decimals int64 `json:"decimals" api:"nullable"`
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer" api:"nullable"`
	// Contract creator balance
	DeployerBalance TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenDeployerBalance `json:"deployer_balance" api:"nullable"`
	// Description of the token
	Description string `json:"description" api:"nullable"`
	// social links of the token
	ExternalLinks TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenExternalLinks `json:"external_links"`
	// URL of the token image
	ImageURL string `json:"image_url" api:"nullable"`
	// List of tokens that this token is impersonating, if detected as an impersonator
	ImpersonationTargets []TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenImpersonationTarget `json:"impersonation_targets" api:"nullable"`
	// Malicious urls associated with the token
	MaliciousURLs []string `json:"malicious_urls" api:"nullable"`
	// Name of the token
	Name string `json:"name" api:"nullable"`
	// Contract owner address
	Owner string `json:"owner" api:"nullable"`
	// Contract owner balance
	OwnerBalance TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenOwnerBalance `json:"owner_balance" api:"nullable"`
	// Symbol of the token
	Symbol string `json:"symbol" api:"nullable"`
	// Address of the token creation initiator, only set if the tokens was created by a
	// well known token launch platform
	TokenCreationInitiator string `json:"token_creation_initiator" api:"nullable"`
	// Type of the token
	Type string `json:"type" api:"nullable"`
	// Urls associated with the token
	URLs []string                                                   `json:"urls" api:"nullable"`
	JSON tokenSnapshotFullResponseItemsMetadataEvmMetadataTokenJSON `json:"-"`
}

// tokenSnapshotFullResponseItemsMetadataEvmMetadataTokenJSON contains the JSON
// metadata for the struct [TokenSnapshotFullResponseItemsMetadataEvmMetadataToken]
type tokenSnapshotFullResponseItemsMetadataEvmMetadataTokenJSON struct {
	ContractBalance        apijson.Field
	CreationTimestamp      apijson.Field
	Decimals               apijson.Field
	Deployer               apijson.Field
	DeployerBalance        apijson.Field
	Description            apijson.Field
	ExternalLinks          apijson.Field
	ImageURL               apijson.Field
	ImpersonationTargets   apijson.Field
	MaliciousURLs          apijson.Field
	Name                   apijson.Field
	Owner                  apijson.Field
	OwnerBalance           apijson.Field
	Symbol                 apijson.Field
	TokenCreationInitiator apijson.Field
	Type                   apijson.Field
	URLs                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *TokenSnapshotFullResponseItemsMetadataEvmMetadataToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotFullResponseItemsMetadataEvmMetadataTokenJSON) RawJSON() string {
	return r.raw
}

func (r TokenSnapshotFullResponseItemsMetadataEvmMetadataToken) implementsTokenSnapshotFullResponseItemsMetadata() {
}

// Contract balance
type TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenContractBalance struct {
	Amount    float64                                                                   `json:"amount" api:"nullable"`
	AmountWei string                                                                    `json:"amount_wei" api:"nullable"`
	JSON      tokenSnapshotFullResponseItemsMetadataEvmMetadataTokenContractBalanceJSON `json:"-"`
}

// tokenSnapshotFullResponseItemsMetadataEvmMetadataTokenContractBalanceJSON
// contains the JSON metadata for the struct
// [TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenContractBalance]
type tokenSnapshotFullResponseItemsMetadataEvmMetadataTokenContractBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenContractBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotFullResponseItemsMetadataEvmMetadataTokenContractBalanceJSON) RawJSON() string {
	return r.raw
}

// Contract creator balance
type TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenDeployerBalance struct {
	Amount    float64                                                                   `json:"amount" api:"nullable"`
	AmountWei string                                                                    `json:"amount_wei" api:"nullable"`
	JSON      tokenSnapshotFullResponseItemsMetadataEvmMetadataTokenDeployerBalanceJSON `json:"-"`
}

// tokenSnapshotFullResponseItemsMetadataEvmMetadataTokenDeployerBalanceJSON
// contains the JSON metadata for the struct
// [TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenDeployerBalance]
type tokenSnapshotFullResponseItemsMetadataEvmMetadataTokenDeployerBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenDeployerBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotFullResponseItemsMetadataEvmMetadataTokenDeployerBalanceJSON) RawJSON() string {
	return r.raw
}

// social links of the token
type TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenExternalLinks struct {
	Homepage          string                                                                  `json:"homepage" api:"nullable"`
	TelegramChannelID string                                                                  `json:"telegram_channel_id" api:"nullable"`
	TwitterPage       string                                                                  `json:"twitter_page" api:"nullable"`
	JSON              tokenSnapshotFullResponseItemsMetadataEvmMetadataTokenExternalLinksJSON `json:"-"`
}

// tokenSnapshotFullResponseItemsMetadataEvmMetadataTokenExternalLinksJSON contains
// the JSON metadata for the struct
// [TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenExternalLinks]
type tokenSnapshotFullResponseItemsMetadataEvmMetadataTokenExternalLinksJSON struct {
	Homepage          apijson.Field
	TelegramChannelID apijson.Field
	TwitterPage       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenExternalLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotFullResponseItemsMetadataEvmMetadataTokenExternalLinksJSON) RawJSON() string {
	return r.raw
}

type TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenImpersonationTarget struct {
	// Address of the token being impersonated
	Address string `json:"address" api:"required"`
	// Blockchain network of the target token
	Chain string `json:"chain" api:"required"`
	// Name of the token being impersonated
	Name string `json:"name" api:"nullable"`
	// Source of the impersonation match
	Source TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenImpersonationTargetsSource `json:"source" api:"nullable"`
	// Symbol of the token being impersonated
	Symbol string                                                                        `json:"symbol" api:"nullable"`
	JSON   tokenSnapshotFullResponseItemsMetadataEvmMetadataTokenImpersonationTargetJSON `json:"-"`
}

// tokenSnapshotFullResponseItemsMetadataEvmMetadataTokenImpersonationTargetJSON
// contains the JSON metadata for the struct
// [TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenImpersonationTarget]
type tokenSnapshotFullResponseItemsMetadataEvmMetadataTokenImpersonationTargetJSON struct {
	Address     apijson.Field
	Chain       apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenImpersonationTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotFullResponseItemsMetadataEvmMetadataTokenImpersonationTargetJSON) RawJSON() string {
	return r.raw
}

// Source of the impersonation match
type TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenImpersonationTargetsSource string

const (
	TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenImpersonationTargetsSourceTopToken    TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenImpersonationTargetsSource = "TOP_TOKEN"
	TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenImpersonationTargetsSourceUserDefined TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenImpersonationTargetsSource = "USER_DEFINED"
)

func (r TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenImpersonationTargetsSource) IsKnown() bool {
	switch r {
	case TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenImpersonationTargetsSourceTopToken, TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenImpersonationTargetsSourceUserDefined:
		return true
	}
	return false
}

// Contract owner balance
type TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenOwnerBalance struct {
	Amount    float64                                                                `json:"amount" api:"nullable"`
	AmountWei string                                                                 `json:"amount_wei" api:"nullable"`
	JSON      tokenSnapshotFullResponseItemsMetadataEvmMetadataTokenOwnerBalanceJSON `json:"-"`
}

// tokenSnapshotFullResponseItemsMetadataEvmMetadataTokenOwnerBalanceJSON contains
// the JSON metadata for the struct
// [TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenOwnerBalance]
type tokenSnapshotFullResponseItemsMetadataEvmMetadataTokenOwnerBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotFullResponseItemsMetadataEvmMetadataTokenOwnerBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotFullResponseItemsMetadataEvmMetadataTokenOwnerBalanceJSON) RawJSON() string {
	return r.raw
}

// General indication
type TokenSnapshotFullResponseItemsResultType string

const (
	TokenSnapshotFullResponseItemsResultTypeBenign    TokenSnapshotFullResponseItemsResultType = "Benign"
	TokenSnapshotFullResponseItemsResultTypeWarning   TokenSnapshotFullResponseItemsResultType = "Warning"
	TokenSnapshotFullResponseItemsResultTypeMalicious TokenSnapshotFullResponseItemsResultType = "Malicious"
	TokenSnapshotFullResponseItemsResultTypeSpam      TokenSnapshotFullResponseItemsResultType = "Spam"
)

func (r TokenSnapshotFullResponseItemsResultType) IsKnown() bool {
	switch r {
	case TokenSnapshotFullResponseItemsResultTypeBenign, TokenSnapshotFullResponseItemsResultTypeWarning, TokenSnapshotFullResponseItemsResultTypeMalicious, TokenSnapshotFullResponseItemsResultTypeSpam:
		return true
	}
	return false
}

// Trading limits of the token
type TokenSnapshotFullResponseItemsTradingLimits struct {
	// Max amount that can be bought at once
	MaxBuy TokenSnapshotFullResponseItemsTradingLimitsMaxBuy `json:"max_buy" api:"nullable"`
	// Max amount that can be held by a single address
	MaxHolding TokenSnapshotFullResponseItemsTradingLimitsMaxHolding `json:"max_holding" api:"nullable"`
	// Max amount that can be sold at once
	MaxSell TokenSnapshotFullResponseItemsTradingLimitsMaxSell `json:"max_sell" api:"nullable"`
	// Maximum amount of the token that can be sold in a block
	SellLimitPerBlock TokenSnapshotFullResponseItemsTradingLimitsSellLimitPerBlock `json:"sell_limit_per_block" api:"nullable"`
	JSON              tokenSnapshotFullResponseItemsTradingLimitsJSON              `json:"-"`
}

// tokenSnapshotFullResponseItemsTradingLimitsJSON contains the JSON metadata for
// the struct [TokenSnapshotFullResponseItemsTradingLimits]
type tokenSnapshotFullResponseItemsTradingLimitsJSON struct {
	MaxBuy            apijson.Field
	MaxHolding        apijson.Field
	MaxSell           apijson.Field
	SellLimitPerBlock apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TokenSnapshotFullResponseItemsTradingLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotFullResponseItemsTradingLimitsJSON) RawJSON() string {
	return r.raw
}

// Max amount that can be bought at once
type TokenSnapshotFullResponseItemsTradingLimitsMaxBuy struct {
	Amount    float64                                               `json:"amount" api:"nullable"`
	AmountWei string                                                `json:"amount_wei" api:"nullable"`
	JSON      tokenSnapshotFullResponseItemsTradingLimitsMaxBuyJSON `json:"-"`
}

// tokenSnapshotFullResponseItemsTradingLimitsMaxBuyJSON contains the JSON metadata
// for the struct [TokenSnapshotFullResponseItemsTradingLimitsMaxBuy]
type tokenSnapshotFullResponseItemsTradingLimitsMaxBuyJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotFullResponseItemsTradingLimitsMaxBuy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotFullResponseItemsTradingLimitsMaxBuyJSON) RawJSON() string {
	return r.raw
}

// Max amount that can be held by a single address
type TokenSnapshotFullResponseItemsTradingLimitsMaxHolding struct {
	Amount    float64                                                   `json:"amount" api:"nullable"`
	AmountWei string                                                    `json:"amount_wei" api:"nullable"`
	JSON      tokenSnapshotFullResponseItemsTradingLimitsMaxHoldingJSON `json:"-"`
}

// tokenSnapshotFullResponseItemsTradingLimitsMaxHoldingJSON contains the JSON
// metadata for the struct [TokenSnapshotFullResponseItemsTradingLimitsMaxHolding]
type tokenSnapshotFullResponseItemsTradingLimitsMaxHoldingJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotFullResponseItemsTradingLimitsMaxHolding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotFullResponseItemsTradingLimitsMaxHoldingJSON) RawJSON() string {
	return r.raw
}

// Max amount that can be sold at once
type TokenSnapshotFullResponseItemsTradingLimitsMaxSell struct {
	Amount    float64                                                `json:"amount" api:"nullable"`
	AmountWei string                                                 `json:"amount_wei" api:"nullable"`
	JSON      tokenSnapshotFullResponseItemsTradingLimitsMaxSellJSON `json:"-"`
}

// tokenSnapshotFullResponseItemsTradingLimitsMaxSellJSON contains the JSON
// metadata for the struct [TokenSnapshotFullResponseItemsTradingLimitsMaxSell]
type tokenSnapshotFullResponseItemsTradingLimitsMaxSellJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotFullResponseItemsTradingLimitsMaxSell) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotFullResponseItemsTradingLimitsMaxSellJSON) RawJSON() string {
	return r.raw
}

// Maximum amount of the token that can be sold in a block
type TokenSnapshotFullResponseItemsTradingLimitsSellLimitPerBlock struct {
	Amount    float64                                                          `json:"amount" api:"nullable"`
	AmountWei string                                                           `json:"amount_wei" api:"nullable"`
	JSON      tokenSnapshotFullResponseItemsTradingLimitsSellLimitPerBlockJSON `json:"-"`
}

// tokenSnapshotFullResponseItemsTradingLimitsSellLimitPerBlockJSON contains the
// JSON metadata for the struct
// [TokenSnapshotFullResponseItemsTradingLimitsSellLimitPerBlock]
type tokenSnapshotFullResponseItemsTradingLimitsSellLimitPerBlockJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotFullResponseItemsTradingLimitsSellLimitPerBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotFullResponseItemsTradingLimitsSellLimitPerBlockJSON) RawJSON() string {
	return r.raw
}

type TokenSnapshotFullResponseItemsFeature struct {
	// Description of the feature
	Description string `json:"description" api:"required"`
	// Feature identifier
	FeatureID TokenSnapshotFullResponseItemsFeaturesFeatureID `json:"feature_id" api:"required"`
	// Type of the feature
	Type TokenSnapshotFullResponseItemsFeaturesType `json:"type" api:"required"`
	JSON tokenSnapshotFullResponseItemsFeatureJSON  `json:"-"`
}

// tokenSnapshotFullResponseItemsFeatureJSON contains the JSON metadata for the
// struct [TokenSnapshotFullResponseItemsFeature]
type tokenSnapshotFullResponseItemsFeatureJSON struct {
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSnapshotFullResponseItemsFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSnapshotFullResponseItemsFeatureJSON) RawJSON() string {
	return r.raw
}

// Feature identifier
type TokenSnapshotFullResponseItemsFeaturesFeatureID string

const (
	TokenSnapshotFullResponseItemsFeaturesFeatureIDVerifiedContract               TokenSnapshotFullResponseItemsFeaturesFeatureID = "VERIFIED_CONTRACT"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDUnverifiedContract             TokenSnapshotFullResponseItemsFeaturesFeatureID = "UNVERIFIED_CONTRACT"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDHighTradeVolume                TokenSnapshotFullResponseItemsFeaturesFeatureID = "HIGH_TRADE_VOLUME"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDMarketPlaceSalesHistory        TokenSnapshotFullResponseItemsFeaturesFeatureID = "MARKET_PLACE_SALES_HISTORY"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDHighReputationToken            TokenSnapshotFullResponseItemsFeaturesFeatureID = "HIGH_REPUTATION_TOKEN"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDOnchainActivityValidator       TokenSnapshotFullResponseItemsFeaturesFeatureID = "ONCHAIN_ACTIVITY_VALIDATOR"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDStaticCodeSignature            TokenSnapshotFullResponseItemsFeaturesFeatureID = "STATIC_CODE_SIGNATURE"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDKnownMalicious                 TokenSnapshotFullResponseItemsFeaturesFeatureID = "KNOWN_MALICIOUS"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDIsEoa                          TokenSnapshotFullResponseItemsFeaturesFeatureID = "IS_EOA"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDIsContract                     TokenSnapshotFullResponseItemsFeaturesFeatureID = "IS_CONTRACT"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDErc20Contract                  TokenSnapshotFullResponseItemsFeaturesFeatureID = "ERC20_CONTRACT"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDTrustedContract                TokenSnapshotFullResponseItemsFeaturesFeatureID = "TRUSTED_CONTRACT"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDBenignCreator                  TokenSnapshotFullResponseItemsFeaturesFeatureID = "BENIGN_CREATOR"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDMetadata                       TokenSnapshotFullResponseItemsFeaturesFeatureID = "METADATA"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDAirdropPattern                 TokenSnapshotFullResponseItemsFeaturesFeatureID = "AIRDROP_PATTERN"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDImpersonator                   TokenSnapshotFullResponseItemsFeaturesFeatureID = "IMPERSONATOR"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDInorganicVolume                TokenSnapshotFullResponseItemsFeaturesFeatureID = "INORGANIC_VOLUME"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDDynamicAnalysis                TokenSnapshotFullResponseItemsFeaturesFeatureID = "DYNAMIC_ANALYSIS"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDConcentratedSupplyDistribution TokenSnapshotFullResponseItemsFeaturesFeatureID = "CONCENTRATED_SUPPLY_DISTRIBUTION"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDHoneypot                       TokenSnapshotFullResponseItemsFeaturesFeatureID = "HONEYPOT"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDInsufficientLockedLiquidity    TokenSnapshotFullResponseItemsFeaturesFeatureID = "INSUFFICIENT_LOCKED_LIQUIDITY"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDUnstableTokenPrice             TokenSnapshotFullResponseItemsFeaturesFeatureID = "UNSTABLE_TOKEN_PRICE"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDRugpull                        TokenSnapshotFullResponseItemsFeaturesFeatureID = "RUGPULL"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDWashTrading                    TokenSnapshotFullResponseItemsFeaturesFeatureID = "WASH_TRADING"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDConsumerOverride               TokenSnapshotFullResponseItemsFeaturesFeatureID = "CONSUMER_OVERRIDE"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDInappropriateContent           TokenSnapshotFullResponseItemsFeaturesFeatureID = "INAPPROPRIATE_CONTENT"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDHighTransferFee                TokenSnapshotFullResponseItemsFeaturesFeatureID = "HIGH_TRANSFER_FEE"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDHighBuyFee                     TokenSnapshotFullResponseItemsFeaturesFeatureID = "HIGH_BUY_FEE"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDHighSellFee                    TokenSnapshotFullResponseItemsFeaturesFeatureID = "HIGH_SELL_FEE"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDUnsellableToken                TokenSnapshotFullResponseItemsFeaturesFeatureID = "UNSELLABLE_TOKEN"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDIsMintable                     TokenSnapshotFullResponseItemsFeaturesFeatureID = "IS_MINTABLE"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDRebaseToken                    TokenSnapshotFullResponseItemsFeaturesFeatureID = "REBASE_TOKEN"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDLiquidStakingToken             TokenSnapshotFullResponseItemsFeaturesFeatureID = "LIQUID_STAKING_TOKEN"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDModifiableTaxes                TokenSnapshotFullResponseItemsFeaturesFeatureID = "MODIFIABLE_TAXES"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDCanBlacklist                   TokenSnapshotFullResponseItemsFeaturesFeatureID = "CAN_BLACKLIST"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDCanWhitelist                   TokenSnapshotFullResponseItemsFeaturesFeatureID = "CAN_WHITELIST"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDHasTradingCooldown             TokenSnapshotFullResponseItemsFeaturesFeatureID = "HAS_TRADING_COOLDOWN"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDExternalFunctions              TokenSnapshotFullResponseItemsFeaturesFeatureID = "EXTERNAL_FUNCTIONS"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDHiddenOwner                    TokenSnapshotFullResponseItemsFeaturesFeatureID = "HIDDEN_OWNER"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDTransferPauseable              TokenSnapshotFullResponseItemsFeaturesFeatureID = "TRANSFER_PAUSEABLE"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDOwnershipRenounced             TokenSnapshotFullResponseItemsFeaturesFeatureID = "OWNERSHIP_RENOUNCED"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDOwnerCanChangeBalance          TokenSnapshotFullResponseItemsFeaturesFeatureID = "OWNER_CAN_CHANGE_BALANCE"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDProxyContract                  TokenSnapshotFullResponseItemsFeaturesFeatureID = "PROXY_CONTRACT"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDSimilarMaliciousContract       TokenSnapshotFullResponseItemsFeaturesFeatureID = "SIMILAR_MALICIOUS_CONTRACT"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDImpersonatorSensitiveAsset     TokenSnapshotFullResponseItemsFeaturesFeatureID = "IMPERSONATOR_SENSITIVE_ASSET"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDImpersonatorHighConfidence     TokenSnapshotFullResponseItemsFeaturesFeatureID = "IMPERSONATOR_HIGH_CONFIDENCE"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDImpersonatorMediumConfidence   TokenSnapshotFullResponseItemsFeaturesFeatureID = "IMPERSONATOR_MEDIUM_CONFIDENCE"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDImpersonatorLowConfidence      TokenSnapshotFullResponseItemsFeaturesFeatureID = "IMPERSONATOR_LOW_CONFIDENCE"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDImpersonationProtected         TokenSnapshotFullResponseItemsFeaturesFeatureID = "IMPERSONATION_PROTECTED"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDFakeVolume                     TokenSnapshotFullResponseItemsFeaturesFeatureID = "FAKE_VOLUME"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDHiddenSupplyByKeyHolder        TokenSnapshotFullResponseItemsFeaturesFeatureID = "HIDDEN_SUPPLY_BY_KEY_HOLDER"
	TokenSnapshotFullResponseItemsFeaturesFeatureIDFakeTradeMakerCount            TokenSnapshotFullResponseItemsFeaturesFeatureID = "FAKE_TRADE_MAKER_COUNT"
)

func (r TokenSnapshotFullResponseItemsFeaturesFeatureID) IsKnown() bool {
	switch r {
	case TokenSnapshotFullResponseItemsFeaturesFeatureIDVerifiedContract, TokenSnapshotFullResponseItemsFeaturesFeatureIDUnverifiedContract, TokenSnapshotFullResponseItemsFeaturesFeatureIDHighTradeVolume, TokenSnapshotFullResponseItemsFeaturesFeatureIDMarketPlaceSalesHistory, TokenSnapshotFullResponseItemsFeaturesFeatureIDHighReputationToken, TokenSnapshotFullResponseItemsFeaturesFeatureIDOnchainActivityValidator, TokenSnapshotFullResponseItemsFeaturesFeatureIDStaticCodeSignature, TokenSnapshotFullResponseItemsFeaturesFeatureIDKnownMalicious, TokenSnapshotFullResponseItemsFeaturesFeatureIDIsEoa, TokenSnapshotFullResponseItemsFeaturesFeatureIDIsContract, TokenSnapshotFullResponseItemsFeaturesFeatureIDErc20Contract, TokenSnapshotFullResponseItemsFeaturesFeatureIDTrustedContract, TokenSnapshotFullResponseItemsFeaturesFeatureIDBenignCreator, TokenSnapshotFullResponseItemsFeaturesFeatureIDMetadata, TokenSnapshotFullResponseItemsFeaturesFeatureIDAirdropPattern, TokenSnapshotFullResponseItemsFeaturesFeatureIDImpersonator, TokenSnapshotFullResponseItemsFeaturesFeatureIDInorganicVolume, TokenSnapshotFullResponseItemsFeaturesFeatureIDDynamicAnalysis, TokenSnapshotFullResponseItemsFeaturesFeatureIDConcentratedSupplyDistribution, TokenSnapshotFullResponseItemsFeaturesFeatureIDHoneypot, TokenSnapshotFullResponseItemsFeaturesFeatureIDInsufficientLockedLiquidity, TokenSnapshotFullResponseItemsFeaturesFeatureIDUnstableTokenPrice, TokenSnapshotFullResponseItemsFeaturesFeatureIDRugpull, TokenSnapshotFullResponseItemsFeaturesFeatureIDWashTrading, TokenSnapshotFullResponseItemsFeaturesFeatureIDConsumerOverride, TokenSnapshotFullResponseItemsFeaturesFeatureIDInappropriateContent, TokenSnapshotFullResponseItemsFeaturesFeatureIDHighTransferFee, TokenSnapshotFullResponseItemsFeaturesFeatureIDHighBuyFee, TokenSnapshotFullResponseItemsFeaturesFeatureIDHighSellFee, TokenSnapshotFullResponseItemsFeaturesFeatureIDUnsellableToken, TokenSnapshotFullResponseItemsFeaturesFeatureIDIsMintable, TokenSnapshotFullResponseItemsFeaturesFeatureIDRebaseToken, TokenSnapshotFullResponseItemsFeaturesFeatureIDLiquidStakingToken, TokenSnapshotFullResponseItemsFeaturesFeatureIDModifiableTaxes, TokenSnapshotFullResponseItemsFeaturesFeatureIDCanBlacklist, TokenSnapshotFullResponseItemsFeaturesFeatureIDCanWhitelist, TokenSnapshotFullResponseItemsFeaturesFeatureIDHasTradingCooldown, TokenSnapshotFullResponseItemsFeaturesFeatureIDExternalFunctions, TokenSnapshotFullResponseItemsFeaturesFeatureIDHiddenOwner, TokenSnapshotFullResponseItemsFeaturesFeatureIDTransferPauseable, TokenSnapshotFullResponseItemsFeaturesFeatureIDOwnershipRenounced, TokenSnapshotFullResponseItemsFeaturesFeatureIDOwnerCanChangeBalance, TokenSnapshotFullResponseItemsFeaturesFeatureIDProxyContract, TokenSnapshotFullResponseItemsFeaturesFeatureIDSimilarMaliciousContract, TokenSnapshotFullResponseItemsFeaturesFeatureIDImpersonatorSensitiveAsset, TokenSnapshotFullResponseItemsFeaturesFeatureIDImpersonatorHighConfidence, TokenSnapshotFullResponseItemsFeaturesFeatureIDImpersonatorMediumConfidence, TokenSnapshotFullResponseItemsFeaturesFeatureIDImpersonatorLowConfidence, TokenSnapshotFullResponseItemsFeaturesFeatureIDImpersonationProtected, TokenSnapshotFullResponseItemsFeaturesFeatureIDFakeVolume, TokenSnapshotFullResponseItemsFeaturesFeatureIDHiddenSupplyByKeyHolder, TokenSnapshotFullResponseItemsFeaturesFeatureIDFakeTradeMakerCount:
		return true
	}
	return false
}

// Type of the feature
type TokenSnapshotFullResponseItemsFeaturesType string

const (
	TokenSnapshotFullResponseItemsFeaturesTypeBenign    TokenSnapshotFullResponseItemsFeaturesType = "Benign"
	TokenSnapshotFullResponseItemsFeaturesTypeInfo      TokenSnapshotFullResponseItemsFeaturesType = "Info"
	TokenSnapshotFullResponseItemsFeaturesTypeWarning   TokenSnapshotFullResponseItemsFeaturesType = "Warning"
	TokenSnapshotFullResponseItemsFeaturesTypeMalicious TokenSnapshotFullResponseItemsFeaturesType = "Malicious"
)

func (r TokenSnapshotFullResponseItemsFeaturesType) IsKnown() bool {
	switch r {
	case TokenSnapshotFullResponseItemsFeaturesTypeBenign, TokenSnapshotFullResponseItemsFeaturesTypeInfo, TokenSnapshotFullResponseItemsFeaturesTypeWarning, TokenSnapshotFullResponseItemsFeaturesTypeMalicious:
		return true
	}
	return false
}

type TokenSnapshotDiffParams struct {
	// The chain name
	Chain param.Field[TokenScanSupportedChain] `query:"chain" api:"required"`
	// Cursor to start from, if not provided, the first page will be returned
	Cursor param.Field[string] `query:"cursor"`
	// Number of tokens to return in a page
	Size param.Field[int64] `query:"size"`
	// Timeframe in minutes
	Timeframe param.Field[int64] `query:"timeframe"`
}

// URLQuery serializes [TokenSnapshotDiffParams]'s query parameters as
// `url.Values`.
func (r TokenSnapshotDiffParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TokenSnapshotFullParams struct {
	// The chain name
	Chain param.Field[TokenScanSupportedChain] `query:"chain" api:"required"`
	// Cursor to start from, if not provided, the first page will be returned
	Cursor param.Field[string] `query:"cursor"`
	// Number of tokens to return in a page
	Size param.Field[int64] `query:"size"`
	// Filter tokens by type. Allowed values are "Fungible" or "NonFungible". If not
	// provided, all types will be included.
	TokenType param.Field[TokenSnapshotFullParamsTokenType] `query:"token_type"`
}

// URLQuery serializes [TokenSnapshotFullParams]'s query parameters as
// `url.Values`.
func (r TokenSnapshotFullParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter tokens by type. Allowed values are "Fungible" or "NonFungible". If not
// provided, all types will be included.
type TokenSnapshotFullParamsTokenType string

const (
	TokenSnapshotFullParamsTokenTypeFungible    TokenSnapshotFullParamsTokenType = "Fungible"
	TokenSnapshotFullParamsTokenTypeNonFungible TokenSnapshotFullParamsTokenType = "NonFungible"
)

func (r TokenSnapshotFullParamsTokenType) IsKnown() bool {
	switch r {
	case TokenSnapshotFullParamsTokenTypeFungible, TokenSnapshotFullParamsTokenTypeNonFungible:
		return true
	}
	return false
}
