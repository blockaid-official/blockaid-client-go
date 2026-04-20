// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"net/http"
	"reflect"
	"slices"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/param"
	"github.com/blockaid-official/blockaid-client-go/internal/requestconfig"
	"github.com/blockaid-official/blockaid-client-go/option"
	"github.com/tidwall/gjson"
)

// TokenBulkService contains methods and other services that help with interacting
// with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTokenBulkService] method instead.
type TokenBulkService struct {
	Options []option.RequestOption
}

// NewTokenBulkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTokenBulkService(opts ...option.RequestOption) (r *TokenBulkService) {
	r = &TokenBulkService{}
	r.Options = opts
	return
}

// Gets a list of token addresses and scan the tokens to identify any indication of
// malicious behaviour
func (r *TokenBulkService) Scan(ctx context.Context, body TokenBulkScanParams, opts ...option.RequestOption) (res *TokenBulkScanResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/token-bulk/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type TokenBulkScanResponse struct {
	Results map[string]TokenBulkScanResponseResult `json:"results" api:"required"`
	// Errors encountered during bulk scanning, keyed by token address
	Errors map[string]TokenBulkScanResponseError `json:"errors"`
	JSON   tokenBulkScanResponseJSON             `json:"-"`
}

// tokenBulkScanResponseJSON contains the JSON metadata for the struct
// [TokenBulkScanResponse]
type tokenBulkScanResponseJSON struct {
	Results     apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenBulkScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseJSON) RawJSON() string {
	return r.raw
}

type TokenBulkScanResponseResult struct {
	// Token address to validate (EVM / Solana)
	Address string `json:"address" api:"required"`
	// Dictionary of detected attacks found during the scan
	AttackTypes map[string]TokenBulkScanResponseResultsAttackType `json:"attack_types" api:"required"`
	// Blockchain network
	Chain TokenScanSupportedChain `json:"chain" api:"required"`
	// Fees associated with the token
	Fees TokenBulkScanResponseResultsFees `json:"fees" api:"required"`
	// financial stats of the token
	FinancialStats FinancialStats `json:"financial_stats" api:"required"`
	// Score between 0 to 1 (double)
	MaliciousScore string `json:"malicious_score" api:"required"`
	// Metadata of the token
	Metadata TokenBulkScanResponseResultsMetadata `json:"metadata" api:"required"`
	// General indication
	ResultType TokenBulkScanResponseResultsResultType `json:"result_type" api:"required"`
	// Trading limits of the token
	TradingLimits TokenBulkScanResponseResultsTradingLimits `json:"trading_limits" api:"required"`
	// List of features associated with the token
	Features []TokenBulkScanResponseResultsFeature `json:"features"`
	JSON     tokenBulkScanResponseResultJSON       `json:"-"`
}

// tokenBulkScanResponseResultJSON contains the JSON metadata for the struct
// [TokenBulkScanResponseResult]
type tokenBulkScanResponseResultJSON struct {
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

func (r *TokenBulkScanResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultJSON) RawJSON() string {
	return r.raw
}

type TokenBulkScanResponseResultsAttackType struct {
	// Score between 0 to 1 (double) that indicates the assurance this attack happened
	Score string `json:"score" api:"required"`
	// Object contains an extra information related to the attack
	Features interface{} `json:"features"`
	// If score is higher or equal to this field, the token is using this attack type
	Threshold string                                     `json:"threshold"`
	JSON      tokenBulkScanResponseResultsAttackTypeJSON `json:"-"`
}

// tokenBulkScanResponseResultsAttackTypeJSON contains the JSON metadata for the
// struct [TokenBulkScanResponseResultsAttackType]
type tokenBulkScanResponseResultsAttackTypeJSON struct {
	Score       apijson.Field
	Features    apijson.Field
	Threshold   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenBulkScanResponseResultsAttackType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsAttackTypeJSON) RawJSON() string {
	return r.raw
}

// Fees associated with the token
type TokenBulkScanResponseResultsFees struct {
	// Buy fee of the token
	Buy float64 `json:"buy" api:"nullable"`
	// Sell fee of the token
	Sell float64 `json:"sell" api:"nullable"`
	// Transfer fee of the token
	Transfer float64 `json:"transfer" api:"nullable"`
	// The maximum value that a transfer fee will cost
	TransferFeeMaxAmount int64                                `json:"transfer_fee_max_amount" api:"nullable"`
	JSON                 tokenBulkScanResponseResultsFeesJSON `json:"-"`
}

// tokenBulkScanResponseResultsFeesJSON contains the JSON metadata for the struct
// [TokenBulkScanResponseResultsFees]
type tokenBulkScanResponseResultsFeesJSON struct {
	Buy                  apijson.Field
	Sell                 apijson.Field
	Transfer             apijson.Field
	TransferFeeMaxAmount apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TokenBulkScanResponseResultsFees) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsFeesJSON) RawJSON() string {
	return r.raw
}

// Metadata of the token
type TokenBulkScanResponseResultsMetadata struct {
	// The unique ID for the Rune
	ID string `json:"id" api:"nullable"`
	// This field can have the runtime type of
	// [TokenBulkScanResponseResultsMetadataSolanaMetadataContractBalance],
	// [TokenBulkScanResponseResultsMetadataEvmMetadataTokenContractBalance].
	ContractBalance interface{} `json:"contract_balance"`
	// Contract deploy date
	CreationTimestamp string `json:"creation_timestamp" api:"nullable"`
	// Decimals of the token
	Decimals int64 `json:"decimals" api:"nullable"`
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer" api:"nullable"`
	// This field can have the runtime type of
	// [TokenBulkScanResponseResultsMetadataSolanaMetadataDeployerBalance],
	// [TokenBulkScanResponseResultsMetadataEvmMetadataTokenDeployerBalance].
	DeployerBalance interface{} `json:"deployer_balance"`
	// Description of the token
	Description string `json:"description" api:"nullable"`
	// This field can have the runtime type of
	// [TokenBulkScanResponseResultsMetadataSolanaMetadataExternalLinks],
	// [TokenBulkScanResponseResultsMetadataEvmMetadataTokenExternalLinks].
	ExternalLinks interface{} `json:"external_links"`
	// The formatted name of the rune, with spacers
	FormattedName string `json:"formatted_name" api:"nullable"`
	// Solana token freeze authority account
	FreezeAuthority string `json:"freeze_authority" api:"nullable"`
	// URL of the token image
	ImageURL string `json:"image_url" api:"nullable"`
	// This field can have the runtime type of
	// [[]TokenBulkScanResponseResultsMetadataSolanaMetadataImpersonationTarget],
	// [[]TokenBulkScanResponseResultsMetadataEvmMetadataTokenImpersonationTarget].
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
	// [TokenBulkScanResponseResultsMetadataSolanaMetadataOwnerBalance],
	// [TokenBulkScanResponseResultsMetadataEvmMetadataTokenOwnerBalance].
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
	URLs  interface{}                              `json:"urls"`
	JSON  tokenBulkScanResponseResultsMetadataJSON `json:"-"`
	union TokenBulkScanResponseResultsMetadataUnion
}

// tokenBulkScanResponseResultsMetadataJSON contains the JSON metadata for the
// struct [TokenBulkScanResponseResultsMetadata]
type tokenBulkScanResponseResultsMetadataJSON struct {
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

func (r tokenBulkScanResponseResultsMetadataJSON) RawJSON() string {
	return r.raw
}

func (r *TokenBulkScanResponseResultsMetadata) UnmarshalJSON(data []byte) (err error) {
	*r = TokenBulkScanResponseResultsMetadata{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TokenBulkScanResponseResultsMetadataUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TokenBulkScanResponseResultsMetadataSolanaMetadata],
// [TokenBulkScanResponseResultsMetadataBitcoinMetadataToken],
// [TokenBulkScanResponseResultsMetadataEvmMetadataToken].
func (r TokenBulkScanResponseResultsMetadata) AsUnion() TokenBulkScanResponseResultsMetadataUnion {
	return r.union
}

// Metadata of the token
//
// Union satisfied by [TokenBulkScanResponseResultsMetadataSolanaMetadata],
// [TokenBulkScanResponseResultsMetadataBitcoinMetadataToken] or
// [TokenBulkScanResponseResultsMetadataEvmMetadataToken].
type TokenBulkScanResponseResultsMetadataUnion interface {
	implementsTokenBulkScanResponseResultsMetadata()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TokenBulkScanResponseResultsMetadataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenBulkScanResponseResultsMetadataSolanaMetadata{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenBulkScanResponseResultsMetadataBitcoinMetadataToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenBulkScanResponseResultsMetadataEvmMetadataToken{}),
		},
	)
}

type TokenBulkScanResponseResultsMetadataSolanaMetadata struct {
	// Contract balance
	ContractBalance TokenBulkScanResponseResultsMetadataSolanaMetadataContractBalance `json:"contract_balance" api:"nullable"`
	// Contract deploy date
	CreationTimestamp string `json:"creation_timestamp" api:"nullable"`
	// Decimals of the token
	Decimals int64 `json:"decimals" api:"nullable"`
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer" api:"nullable"`
	// Contract creator balance
	DeployerBalance TokenBulkScanResponseResultsMetadataSolanaMetadataDeployerBalance `json:"deployer_balance" api:"nullable"`
	// Description of the token
	Description string `json:"description" api:"nullable"`
	// social links of the token
	ExternalLinks TokenBulkScanResponseResultsMetadataSolanaMetadataExternalLinks `json:"external_links"`
	// Solana token freeze authority account
	FreezeAuthority string `json:"freeze_authority" api:"nullable"`
	// URL of the token image
	ImageURL string `json:"image_url" api:"nullable"`
	// List of tokens that this token is impersonating, if detected as an impersonator
	ImpersonationTargets []TokenBulkScanResponseResultsMetadataSolanaMetadataImpersonationTarget `json:"impersonation_targets" api:"nullable"`
	// Malicious urls associated with the token
	MaliciousURLs []string `json:"malicious_urls" api:"nullable"`
	// Solana token mint authority account
	MintAuthority string `json:"mint_authority" api:"nullable"`
	// Name of the token
	Name string `json:"name" api:"nullable"`
	// Contract owner address
	Owner string `json:"owner" api:"nullable"`
	// Contract owner balance
	OwnerBalance TokenBulkScanResponseResultsMetadataSolanaMetadataOwnerBalance `json:"owner_balance" api:"nullable"`
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
	URLs []string                                               `json:"urls" api:"nullable"`
	JSON tokenBulkScanResponseResultsMetadataSolanaMetadataJSON `json:"-"`
}

// tokenBulkScanResponseResultsMetadataSolanaMetadataJSON contains the JSON
// metadata for the struct [TokenBulkScanResponseResultsMetadataSolanaMetadata]
type tokenBulkScanResponseResultsMetadataSolanaMetadataJSON struct {
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

func (r *TokenBulkScanResponseResultsMetadataSolanaMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsMetadataSolanaMetadataJSON) RawJSON() string {
	return r.raw
}

func (r TokenBulkScanResponseResultsMetadataSolanaMetadata) implementsTokenBulkScanResponseResultsMetadata() {
}

// Contract balance
type TokenBulkScanResponseResultsMetadataSolanaMetadataContractBalance struct {
	Amount    float64                                                               `json:"amount" api:"nullable"`
	AmountWei string                                                                `json:"amount_wei" api:"nullable"`
	JSON      tokenBulkScanResponseResultsMetadataSolanaMetadataContractBalanceJSON `json:"-"`
}

// tokenBulkScanResponseResultsMetadataSolanaMetadataContractBalanceJSON contains
// the JSON metadata for the struct
// [TokenBulkScanResponseResultsMetadataSolanaMetadataContractBalance]
type tokenBulkScanResponseResultsMetadataSolanaMetadataContractBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenBulkScanResponseResultsMetadataSolanaMetadataContractBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsMetadataSolanaMetadataContractBalanceJSON) RawJSON() string {
	return r.raw
}

// Contract creator balance
type TokenBulkScanResponseResultsMetadataSolanaMetadataDeployerBalance struct {
	Amount    float64                                                               `json:"amount" api:"nullable"`
	AmountWei string                                                                `json:"amount_wei" api:"nullable"`
	JSON      tokenBulkScanResponseResultsMetadataSolanaMetadataDeployerBalanceJSON `json:"-"`
}

// tokenBulkScanResponseResultsMetadataSolanaMetadataDeployerBalanceJSON contains
// the JSON metadata for the struct
// [TokenBulkScanResponseResultsMetadataSolanaMetadataDeployerBalance]
type tokenBulkScanResponseResultsMetadataSolanaMetadataDeployerBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenBulkScanResponseResultsMetadataSolanaMetadataDeployerBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsMetadataSolanaMetadataDeployerBalanceJSON) RawJSON() string {
	return r.raw
}

// social links of the token
type TokenBulkScanResponseResultsMetadataSolanaMetadataExternalLinks struct {
	Homepage          string                                                              `json:"homepage" api:"nullable"`
	TelegramChannelID string                                                              `json:"telegram_channel_id" api:"nullable"`
	TwitterPage       string                                                              `json:"twitter_page" api:"nullable"`
	JSON              tokenBulkScanResponseResultsMetadataSolanaMetadataExternalLinksJSON `json:"-"`
}

// tokenBulkScanResponseResultsMetadataSolanaMetadataExternalLinksJSON contains the
// JSON metadata for the struct
// [TokenBulkScanResponseResultsMetadataSolanaMetadataExternalLinks]
type tokenBulkScanResponseResultsMetadataSolanaMetadataExternalLinksJSON struct {
	Homepage          apijson.Field
	TelegramChannelID apijson.Field
	TwitterPage       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TokenBulkScanResponseResultsMetadataSolanaMetadataExternalLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsMetadataSolanaMetadataExternalLinksJSON) RawJSON() string {
	return r.raw
}

type TokenBulkScanResponseResultsMetadataSolanaMetadataImpersonationTarget struct {
	// Address of the token being impersonated
	Address string `json:"address" api:"required"`
	// Blockchain network of the target token
	Chain string `json:"chain" api:"required"`
	// Name of the token being impersonated
	Name string `json:"name" api:"nullable"`
	// Source of the impersonation match
	Source TokenBulkScanResponseResultsMetadataSolanaMetadataImpersonationTargetsSource `json:"source" api:"nullable"`
	// Symbol of the token being impersonated
	Symbol string                                                                    `json:"symbol" api:"nullable"`
	JSON   tokenBulkScanResponseResultsMetadataSolanaMetadataImpersonationTargetJSON `json:"-"`
}

// tokenBulkScanResponseResultsMetadataSolanaMetadataImpersonationTargetJSON
// contains the JSON metadata for the struct
// [TokenBulkScanResponseResultsMetadataSolanaMetadataImpersonationTarget]
type tokenBulkScanResponseResultsMetadataSolanaMetadataImpersonationTargetJSON struct {
	Address     apijson.Field
	Chain       apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenBulkScanResponseResultsMetadataSolanaMetadataImpersonationTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsMetadataSolanaMetadataImpersonationTargetJSON) RawJSON() string {
	return r.raw
}

// Source of the impersonation match
type TokenBulkScanResponseResultsMetadataSolanaMetadataImpersonationTargetsSource string

const (
	TokenBulkScanResponseResultsMetadataSolanaMetadataImpersonationTargetsSourceTopToken    TokenBulkScanResponseResultsMetadataSolanaMetadataImpersonationTargetsSource = "TOP_TOKEN"
	TokenBulkScanResponseResultsMetadataSolanaMetadataImpersonationTargetsSourceUserDefined TokenBulkScanResponseResultsMetadataSolanaMetadataImpersonationTargetsSource = "USER_DEFINED"
)

func (r TokenBulkScanResponseResultsMetadataSolanaMetadataImpersonationTargetsSource) IsKnown() bool {
	switch r {
	case TokenBulkScanResponseResultsMetadataSolanaMetadataImpersonationTargetsSourceTopToken, TokenBulkScanResponseResultsMetadataSolanaMetadataImpersonationTargetsSourceUserDefined:
		return true
	}
	return false
}

// Contract owner balance
type TokenBulkScanResponseResultsMetadataSolanaMetadataOwnerBalance struct {
	Amount    float64                                                            `json:"amount" api:"nullable"`
	AmountWei string                                                             `json:"amount_wei" api:"nullable"`
	JSON      tokenBulkScanResponseResultsMetadataSolanaMetadataOwnerBalanceJSON `json:"-"`
}

// tokenBulkScanResponseResultsMetadataSolanaMetadataOwnerBalanceJSON contains the
// JSON metadata for the struct
// [TokenBulkScanResponseResultsMetadataSolanaMetadataOwnerBalance]
type tokenBulkScanResponseResultsMetadataSolanaMetadataOwnerBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenBulkScanResponseResultsMetadataSolanaMetadataOwnerBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsMetadataSolanaMetadataOwnerBalanceJSON) RawJSON() string {
	return r.raw
}

type TokenBulkScanResponseResultsMetadataBitcoinMetadataToken struct {
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
	Type string                                                       `json:"type" api:"nullable"`
	JSON tokenBulkScanResponseResultsMetadataBitcoinMetadataTokenJSON `json:"-"`
}

// tokenBulkScanResponseResultsMetadataBitcoinMetadataTokenJSON contains the JSON
// metadata for the struct
// [TokenBulkScanResponseResultsMetadataBitcoinMetadataToken]
type tokenBulkScanResponseResultsMetadataBitcoinMetadataTokenJSON struct {
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

func (r *TokenBulkScanResponseResultsMetadataBitcoinMetadataToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsMetadataBitcoinMetadataTokenJSON) RawJSON() string {
	return r.raw
}

func (r TokenBulkScanResponseResultsMetadataBitcoinMetadataToken) implementsTokenBulkScanResponseResultsMetadata() {
}

type TokenBulkScanResponseResultsMetadataEvmMetadataToken struct {
	// Contract balance
	ContractBalance TokenBulkScanResponseResultsMetadataEvmMetadataTokenContractBalance `json:"contract_balance" api:"nullable"`
	// Contract deploy date
	CreationTimestamp string `json:"creation_timestamp" api:"nullable"`
	// Decimals of the token
	Decimals int64 `json:"decimals" api:"nullable"`
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer" api:"nullable"`
	// Contract creator balance
	DeployerBalance TokenBulkScanResponseResultsMetadataEvmMetadataTokenDeployerBalance `json:"deployer_balance" api:"nullable"`
	// Description of the token
	Description string `json:"description" api:"nullable"`
	// social links of the token
	ExternalLinks TokenBulkScanResponseResultsMetadataEvmMetadataTokenExternalLinks `json:"external_links"`
	// URL of the token image
	ImageURL string `json:"image_url" api:"nullable"`
	// List of tokens that this token is impersonating, if detected as an impersonator
	ImpersonationTargets []TokenBulkScanResponseResultsMetadataEvmMetadataTokenImpersonationTarget `json:"impersonation_targets" api:"nullable"`
	// Malicious urls associated with the token
	MaliciousURLs []string `json:"malicious_urls" api:"nullable"`
	// Name of the token
	Name string `json:"name" api:"nullable"`
	// Contract owner address
	Owner string `json:"owner" api:"nullable"`
	// Contract owner balance
	OwnerBalance TokenBulkScanResponseResultsMetadataEvmMetadataTokenOwnerBalance `json:"owner_balance" api:"nullable"`
	// Symbol of the token
	Symbol string `json:"symbol" api:"nullable"`
	// Address of the token creation initiator, only set if the tokens was created by a
	// well known token launch platform
	TokenCreationInitiator string `json:"token_creation_initiator" api:"nullable"`
	// Type of the token
	Type string `json:"type" api:"nullable"`
	// Urls associated with the token
	URLs []string                                                 `json:"urls" api:"nullable"`
	JSON tokenBulkScanResponseResultsMetadataEvmMetadataTokenJSON `json:"-"`
}

// tokenBulkScanResponseResultsMetadataEvmMetadataTokenJSON contains the JSON
// metadata for the struct [TokenBulkScanResponseResultsMetadataEvmMetadataToken]
type tokenBulkScanResponseResultsMetadataEvmMetadataTokenJSON struct {
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

func (r *TokenBulkScanResponseResultsMetadataEvmMetadataToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsMetadataEvmMetadataTokenJSON) RawJSON() string {
	return r.raw
}

func (r TokenBulkScanResponseResultsMetadataEvmMetadataToken) implementsTokenBulkScanResponseResultsMetadata() {
}

// Contract balance
type TokenBulkScanResponseResultsMetadataEvmMetadataTokenContractBalance struct {
	Amount    float64                                                                 `json:"amount" api:"nullable"`
	AmountWei string                                                                  `json:"amount_wei" api:"nullable"`
	JSON      tokenBulkScanResponseResultsMetadataEvmMetadataTokenContractBalanceJSON `json:"-"`
}

// tokenBulkScanResponseResultsMetadataEvmMetadataTokenContractBalanceJSON contains
// the JSON metadata for the struct
// [TokenBulkScanResponseResultsMetadataEvmMetadataTokenContractBalance]
type tokenBulkScanResponseResultsMetadataEvmMetadataTokenContractBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenBulkScanResponseResultsMetadataEvmMetadataTokenContractBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsMetadataEvmMetadataTokenContractBalanceJSON) RawJSON() string {
	return r.raw
}

// Contract creator balance
type TokenBulkScanResponseResultsMetadataEvmMetadataTokenDeployerBalance struct {
	Amount    float64                                                                 `json:"amount" api:"nullable"`
	AmountWei string                                                                  `json:"amount_wei" api:"nullable"`
	JSON      tokenBulkScanResponseResultsMetadataEvmMetadataTokenDeployerBalanceJSON `json:"-"`
}

// tokenBulkScanResponseResultsMetadataEvmMetadataTokenDeployerBalanceJSON contains
// the JSON metadata for the struct
// [TokenBulkScanResponseResultsMetadataEvmMetadataTokenDeployerBalance]
type tokenBulkScanResponseResultsMetadataEvmMetadataTokenDeployerBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenBulkScanResponseResultsMetadataEvmMetadataTokenDeployerBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsMetadataEvmMetadataTokenDeployerBalanceJSON) RawJSON() string {
	return r.raw
}

// social links of the token
type TokenBulkScanResponseResultsMetadataEvmMetadataTokenExternalLinks struct {
	Homepage          string                                                                `json:"homepage" api:"nullable"`
	TelegramChannelID string                                                                `json:"telegram_channel_id" api:"nullable"`
	TwitterPage       string                                                                `json:"twitter_page" api:"nullable"`
	JSON              tokenBulkScanResponseResultsMetadataEvmMetadataTokenExternalLinksJSON `json:"-"`
}

// tokenBulkScanResponseResultsMetadataEvmMetadataTokenExternalLinksJSON contains
// the JSON metadata for the struct
// [TokenBulkScanResponseResultsMetadataEvmMetadataTokenExternalLinks]
type tokenBulkScanResponseResultsMetadataEvmMetadataTokenExternalLinksJSON struct {
	Homepage          apijson.Field
	TelegramChannelID apijson.Field
	TwitterPage       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TokenBulkScanResponseResultsMetadataEvmMetadataTokenExternalLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsMetadataEvmMetadataTokenExternalLinksJSON) RawJSON() string {
	return r.raw
}

type TokenBulkScanResponseResultsMetadataEvmMetadataTokenImpersonationTarget struct {
	// Address of the token being impersonated
	Address string `json:"address" api:"required"`
	// Blockchain network of the target token
	Chain string `json:"chain" api:"required"`
	// Name of the token being impersonated
	Name string `json:"name" api:"nullable"`
	// Source of the impersonation match
	Source TokenBulkScanResponseResultsMetadataEvmMetadataTokenImpersonationTargetsSource `json:"source" api:"nullable"`
	// Symbol of the token being impersonated
	Symbol string                                                                      `json:"symbol" api:"nullable"`
	JSON   tokenBulkScanResponseResultsMetadataEvmMetadataTokenImpersonationTargetJSON `json:"-"`
}

// tokenBulkScanResponseResultsMetadataEvmMetadataTokenImpersonationTargetJSON
// contains the JSON metadata for the struct
// [TokenBulkScanResponseResultsMetadataEvmMetadataTokenImpersonationTarget]
type tokenBulkScanResponseResultsMetadataEvmMetadataTokenImpersonationTargetJSON struct {
	Address     apijson.Field
	Chain       apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenBulkScanResponseResultsMetadataEvmMetadataTokenImpersonationTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsMetadataEvmMetadataTokenImpersonationTargetJSON) RawJSON() string {
	return r.raw
}

// Source of the impersonation match
type TokenBulkScanResponseResultsMetadataEvmMetadataTokenImpersonationTargetsSource string

const (
	TokenBulkScanResponseResultsMetadataEvmMetadataTokenImpersonationTargetsSourceTopToken    TokenBulkScanResponseResultsMetadataEvmMetadataTokenImpersonationTargetsSource = "TOP_TOKEN"
	TokenBulkScanResponseResultsMetadataEvmMetadataTokenImpersonationTargetsSourceUserDefined TokenBulkScanResponseResultsMetadataEvmMetadataTokenImpersonationTargetsSource = "USER_DEFINED"
)

func (r TokenBulkScanResponseResultsMetadataEvmMetadataTokenImpersonationTargetsSource) IsKnown() bool {
	switch r {
	case TokenBulkScanResponseResultsMetadataEvmMetadataTokenImpersonationTargetsSourceTopToken, TokenBulkScanResponseResultsMetadataEvmMetadataTokenImpersonationTargetsSourceUserDefined:
		return true
	}
	return false
}

// Contract owner balance
type TokenBulkScanResponseResultsMetadataEvmMetadataTokenOwnerBalance struct {
	Amount    float64                                                              `json:"amount" api:"nullable"`
	AmountWei string                                                               `json:"amount_wei" api:"nullable"`
	JSON      tokenBulkScanResponseResultsMetadataEvmMetadataTokenOwnerBalanceJSON `json:"-"`
}

// tokenBulkScanResponseResultsMetadataEvmMetadataTokenOwnerBalanceJSON contains
// the JSON metadata for the struct
// [TokenBulkScanResponseResultsMetadataEvmMetadataTokenOwnerBalance]
type tokenBulkScanResponseResultsMetadataEvmMetadataTokenOwnerBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenBulkScanResponseResultsMetadataEvmMetadataTokenOwnerBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsMetadataEvmMetadataTokenOwnerBalanceJSON) RawJSON() string {
	return r.raw
}

// General indication
type TokenBulkScanResponseResultsResultType string

const (
	TokenBulkScanResponseResultsResultTypeBenign    TokenBulkScanResponseResultsResultType = "Benign"
	TokenBulkScanResponseResultsResultTypeWarning   TokenBulkScanResponseResultsResultType = "Warning"
	TokenBulkScanResponseResultsResultTypeMalicious TokenBulkScanResponseResultsResultType = "Malicious"
	TokenBulkScanResponseResultsResultTypeSpam      TokenBulkScanResponseResultsResultType = "Spam"
)

func (r TokenBulkScanResponseResultsResultType) IsKnown() bool {
	switch r {
	case TokenBulkScanResponseResultsResultTypeBenign, TokenBulkScanResponseResultsResultTypeWarning, TokenBulkScanResponseResultsResultTypeMalicious, TokenBulkScanResponseResultsResultTypeSpam:
		return true
	}
	return false
}

// Trading limits of the token
type TokenBulkScanResponseResultsTradingLimits struct {
	// Max amount that can be bought at once
	MaxBuy TokenBulkScanResponseResultsTradingLimitsMaxBuy `json:"max_buy" api:"nullable"`
	// Max amount that can be held by a single address
	MaxHolding TokenBulkScanResponseResultsTradingLimitsMaxHolding `json:"max_holding" api:"nullable"`
	// Max amount that can be sold at once
	MaxSell TokenBulkScanResponseResultsTradingLimitsMaxSell `json:"max_sell" api:"nullable"`
	// Maximum amount of the token that can be sold in a block
	SellLimitPerBlock TokenBulkScanResponseResultsTradingLimitsSellLimitPerBlock `json:"sell_limit_per_block" api:"nullable"`
	JSON              tokenBulkScanResponseResultsTradingLimitsJSON              `json:"-"`
}

// tokenBulkScanResponseResultsTradingLimitsJSON contains the JSON metadata for the
// struct [TokenBulkScanResponseResultsTradingLimits]
type tokenBulkScanResponseResultsTradingLimitsJSON struct {
	MaxBuy            apijson.Field
	MaxHolding        apijson.Field
	MaxSell           apijson.Field
	SellLimitPerBlock apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TokenBulkScanResponseResultsTradingLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsTradingLimitsJSON) RawJSON() string {
	return r.raw
}

// Max amount that can be bought at once
type TokenBulkScanResponseResultsTradingLimitsMaxBuy struct {
	Amount    float64                                             `json:"amount" api:"nullable"`
	AmountWei string                                              `json:"amount_wei" api:"nullable"`
	JSON      tokenBulkScanResponseResultsTradingLimitsMaxBuyJSON `json:"-"`
}

// tokenBulkScanResponseResultsTradingLimitsMaxBuyJSON contains the JSON metadata
// for the struct [TokenBulkScanResponseResultsTradingLimitsMaxBuy]
type tokenBulkScanResponseResultsTradingLimitsMaxBuyJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenBulkScanResponseResultsTradingLimitsMaxBuy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsTradingLimitsMaxBuyJSON) RawJSON() string {
	return r.raw
}

// Max amount that can be held by a single address
type TokenBulkScanResponseResultsTradingLimitsMaxHolding struct {
	Amount    float64                                                 `json:"amount" api:"nullable"`
	AmountWei string                                                  `json:"amount_wei" api:"nullable"`
	JSON      tokenBulkScanResponseResultsTradingLimitsMaxHoldingJSON `json:"-"`
}

// tokenBulkScanResponseResultsTradingLimitsMaxHoldingJSON contains the JSON
// metadata for the struct [TokenBulkScanResponseResultsTradingLimitsMaxHolding]
type tokenBulkScanResponseResultsTradingLimitsMaxHoldingJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenBulkScanResponseResultsTradingLimitsMaxHolding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsTradingLimitsMaxHoldingJSON) RawJSON() string {
	return r.raw
}

// Max amount that can be sold at once
type TokenBulkScanResponseResultsTradingLimitsMaxSell struct {
	Amount    float64                                              `json:"amount" api:"nullable"`
	AmountWei string                                               `json:"amount_wei" api:"nullable"`
	JSON      tokenBulkScanResponseResultsTradingLimitsMaxSellJSON `json:"-"`
}

// tokenBulkScanResponseResultsTradingLimitsMaxSellJSON contains the JSON metadata
// for the struct [TokenBulkScanResponseResultsTradingLimitsMaxSell]
type tokenBulkScanResponseResultsTradingLimitsMaxSellJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenBulkScanResponseResultsTradingLimitsMaxSell) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsTradingLimitsMaxSellJSON) RawJSON() string {
	return r.raw
}

// Maximum amount of the token that can be sold in a block
type TokenBulkScanResponseResultsTradingLimitsSellLimitPerBlock struct {
	Amount    float64                                                        `json:"amount" api:"nullable"`
	AmountWei string                                                         `json:"amount_wei" api:"nullable"`
	JSON      tokenBulkScanResponseResultsTradingLimitsSellLimitPerBlockJSON `json:"-"`
}

// tokenBulkScanResponseResultsTradingLimitsSellLimitPerBlockJSON contains the JSON
// metadata for the struct
// [TokenBulkScanResponseResultsTradingLimitsSellLimitPerBlock]
type tokenBulkScanResponseResultsTradingLimitsSellLimitPerBlockJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenBulkScanResponseResultsTradingLimitsSellLimitPerBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsTradingLimitsSellLimitPerBlockJSON) RawJSON() string {
	return r.raw
}

type TokenBulkScanResponseResultsFeature struct {
	// Description of the feature
	Description string `json:"description" api:"required"`
	// Feature identifier
	FeatureID TokenBulkScanResponseResultsFeaturesFeatureID `json:"feature_id" api:"required"`
	// Type of the feature
	Type TokenBulkScanResponseResultsFeaturesType `json:"type" api:"required"`
	JSON tokenBulkScanResponseResultsFeatureJSON  `json:"-"`
}

// tokenBulkScanResponseResultsFeatureJSON contains the JSON metadata for the
// struct [TokenBulkScanResponseResultsFeature]
type tokenBulkScanResponseResultsFeatureJSON struct {
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenBulkScanResponseResultsFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsFeatureJSON) RawJSON() string {
	return r.raw
}

// Feature identifier
type TokenBulkScanResponseResultsFeaturesFeatureID string

const (
	TokenBulkScanResponseResultsFeaturesFeatureIDVerifiedContract               TokenBulkScanResponseResultsFeaturesFeatureID = "VERIFIED_CONTRACT"
	TokenBulkScanResponseResultsFeaturesFeatureIDUnverifiedContract             TokenBulkScanResponseResultsFeaturesFeatureID = "UNVERIFIED_CONTRACT"
	TokenBulkScanResponseResultsFeaturesFeatureIDHighTradeVolume                TokenBulkScanResponseResultsFeaturesFeatureID = "HIGH_TRADE_VOLUME"
	TokenBulkScanResponseResultsFeaturesFeatureIDMarketPlaceSalesHistory        TokenBulkScanResponseResultsFeaturesFeatureID = "MARKET_PLACE_SALES_HISTORY"
	TokenBulkScanResponseResultsFeaturesFeatureIDHighReputationToken            TokenBulkScanResponseResultsFeaturesFeatureID = "HIGH_REPUTATION_TOKEN"
	TokenBulkScanResponseResultsFeaturesFeatureIDOnchainActivityValidator       TokenBulkScanResponseResultsFeaturesFeatureID = "ONCHAIN_ACTIVITY_VALIDATOR"
	TokenBulkScanResponseResultsFeaturesFeatureIDStaticCodeSignature            TokenBulkScanResponseResultsFeaturesFeatureID = "STATIC_CODE_SIGNATURE"
	TokenBulkScanResponseResultsFeaturesFeatureIDKnownMalicious                 TokenBulkScanResponseResultsFeaturesFeatureID = "KNOWN_MALICIOUS"
	TokenBulkScanResponseResultsFeaturesFeatureIDIsEoa                          TokenBulkScanResponseResultsFeaturesFeatureID = "IS_EOA"
	TokenBulkScanResponseResultsFeaturesFeatureIDIsContract                     TokenBulkScanResponseResultsFeaturesFeatureID = "IS_CONTRACT"
	TokenBulkScanResponseResultsFeaturesFeatureIDErc20Contract                  TokenBulkScanResponseResultsFeaturesFeatureID = "ERC20_CONTRACT"
	TokenBulkScanResponseResultsFeaturesFeatureIDTrustedContract                TokenBulkScanResponseResultsFeaturesFeatureID = "TRUSTED_CONTRACT"
	TokenBulkScanResponseResultsFeaturesFeatureIDBenignCreator                  TokenBulkScanResponseResultsFeaturesFeatureID = "BENIGN_CREATOR"
	TokenBulkScanResponseResultsFeaturesFeatureIDMetadata                       TokenBulkScanResponseResultsFeaturesFeatureID = "METADATA"
	TokenBulkScanResponseResultsFeaturesFeatureIDAirdropPattern                 TokenBulkScanResponseResultsFeaturesFeatureID = "AIRDROP_PATTERN"
	TokenBulkScanResponseResultsFeaturesFeatureIDImpersonator                   TokenBulkScanResponseResultsFeaturesFeatureID = "IMPERSONATOR"
	TokenBulkScanResponseResultsFeaturesFeatureIDInorganicVolume                TokenBulkScanResponseResultsFeaturesFeatureID = "INORGANIC_VOLUME"
	TokenBulkScanResponseResultsFeaturesFeatureIDDynamicAnalysis                TokenBulkScanResponseResultsFeaturesFeatureID = "DYNAMIC_ANALYSIS"
	TokenBulkScanResponseResultsFeaturesFeatureIDConcentratedSupplyDistribution TokenBulkScanResponseResultsFeaturesFeatureID = "CONCENTRATED_SUPPLY_DISTRIBUTION"
	TokenBulkScanResponseResultsFeaturesFeatureIDHoneypot                       TokenBulkScanResponseResultsFeaturesFeatureID = "HONEYPOT"
	TokenBulkScanResponseResultsFeaturesFeatureIDInsufficientLockedLiquidity    TokenBulkScanResponseResultsFeaturesFeatureID = "INSUFFICIENT_LOCKED_LIQUIDITY"
	TokenBulkScanResponseResultsFeaturesFeatureIDUnstableTokenPrice             TokenBulkScanResponseResultsFeaturesFeatureID = "UNSTABLE_TOKEN_PRICE"
	TokenBulkScanResponseResultsFeaturesFeatureIDRugpull                        TokenBulkScanResponseResultsFeaturesFeatureID = "RUGPULL"
	TokenBulkScanResponseResultsFeaturesFeatureIDWashTrading                    TokenBulkScanResponseResultsFeaturesFeatureID = "WASH_TRADING"
	TokenBulkScanResponseResultsFeaturesFeatureIDConsumerOverride               TokenBulkScanResponseResultsFeaturesFeatureID = "CONSUMER_OVERRIDE"
	TokenBulkScanResponseResultsFeaturesFeatureIDInappropriateContent           TokenBulkScanResponseResultsFeaturesFeatureID = "INAPPROPRIATE_CONTENT"
	TokenBulkScanResponseResultsFeaturesFeatureIDHighTransferFee                TokenBulkScanResponseResultsFeaturesFeatureID = "HIGH_TRANSFER_FEE"
	TokenBulkScanResponseResultsFeaturesFeatureIDHighBuyFee                     TokenBulkScanResponseResultsFeaturesFeatureID = "HIGH_BUY_FEE"
	TokenBulkScanResponseResultsFeaturesFeatureIDHighSellFee                    TokenBulkScanResponseResultsFeaturesFeatureID = "HIGH_SELL_FEE"
	TokenBulkScanResponseResultsFeaturesFeatureIDUnsellableToken                TokenBulkScanResponseResultsFeaturesFeatureID = "UNSELLABLE_TOKEN"
	TokenBulkScanResponseResultsFeaturesFeatureIDIsMintable                     TokenBulkScanResponseResultsFeaturesFeatureID = "IS_MINTABLE"
	TokenBulkScanResponseResultsFeaturesFeatureIDRebaseToken                    TokenBulkScanResponseResultsFeaturesFeatureID = "REBASE_TOKEN"
	TokenBulkScanResponseResultsFeaturesFeatureIDLiquidStakingToken             TokenBulkScanResponseResultsFeaturesFeatureID = "LIQUID_STAKING_TOKEN"
	TokenBulkScanResponseResultsFeaturesFeatureIDModifiableTaxes                TokenBulkScanResponseResultsFeaturesFeatureID = "MODIFIABLE_TAXES"
	TokenBulkScanResponseResultsFeaturesFeatureIDCanBlacklist                   TokenBulkScanResponseResultsFeaturesFeatureID = "CAN_BLACKLIST"
	TokenBulkScanResponseResultsFeaturesFeatureIDCanWhitelist                   TokenBulkScanResponseResultsFeaturesFeatureID = "CAN_WHITELIST"
	TokenBulkScanResponseResultsFeaturesFeatureIDHasTradingCooldown             TokenBulkScanResponseResultsFeaturesFeatureID = "HAS_TRADING_COOLDOWN"
	TokenBulkScanResponseResultsFeaturesFeatureIDExternalFunctions              TokenBulkScanResponseResultsFeaturesFeatureID = "EXTERNAL_FUNCTIONS"
	TokenBulkScanResponseResultsFeaturesFeatureIDHiddenOwner                    TokenBulkScanResponseResultsFeaturesFeatureID = "HIDDEN_OWNER"
	TokenBulkScanResponseResultsFeaturesFeatureIDTransferPauseable              TokenBulkScanResponseResultsFeaturesFeatureID = "TRANSFER_PAUSEABLE"
	TokenBulkScanResponseResultsFeaturesFeatureIDOwnershipRenounced             TokenBulkScanResponseResultsFeaturesFeatureID = "OWNERSHIP_RENOUNCED"
	TokenBulkScanResponseResultsFeaturesFeatureIDOwnerCanChangeBalance          TokenBulkScanResponseResultsFeaturesFeatureID = "OWNER_CAN_CHANGE_BALANCE"
	TokenBulkScanResponseResultsFeaturesFeatureIDProxyContract                  TokenBulkScanResponseResultsFeaturesFeatureID = "PROXY_CONTRACT"
	TokenBulkScanResponseResultsFeaturesFeatureIDSimilarMaliciousContract       TokenBulkScanResponseResultsFeaturesFeatureID = "SIMILAR_MALICIOUS_CONTRACT"
	TokenBulkScanResponseResultsFeaturesFeatureIDImpersonatorSensitiveAsset     TokenBulkScanResponseResultsFeaturesFeatureID = "IMPERSONATOR_SENSITIVE_ASSET"
	TokenBulkScanResponseResultsFeaturesFeatureIDImpersonatorHighConfidence     TokenBulkScanResponseResultsFeaturesFeatureID = "IMPERSONATOR_HIGH_CONFIDENCE"
	TokenBulkScanResponseResultsFeaturesFeatureIDImpersonatorMediumConfidence   TokenBulkScanResponseResultsFeaturesFeatureID = "IMPERSONATOR_MEDIUM_CONFIDENCE"
	TokenBulkScanResponseResultsFeaturesFeatureIDImpersonatorLowConfidence      TokenBulkScanResponseResultsFeaturesFeatureID = "IMPERSONATOR_LOW_CONFIDENCE"
	TokenBulkScanResponseResultsFeaturesFeatureIDImpersonationProtected         TokenBulkScanResponseResultsFeaturesFeatureID = "IMPERSONATION_PROTECTED"
	TokenBulkScanResponseResultsFeaturesFeatureIDFakeVolume                     TokenBulkScanResponseResultsFeaturesFeatureID = "FAKE_VOLUME"
	TokenBulkScanResponseResultsFeaturesFeatureIDHiddenSupplyByKeyHolder        TokenBulkScanResponseResultsFeaturesFeatureID = "HIDDEN_SUPPLY_BY_KEY_HOLDER"
	TokenBulkScanResponseResultsFeaturesFeatureIDFakeTradeMakerCount            TokenBulkScanResponseResultsFeaturesFeatureID = "FAKE_TRADE_MAKER_COUNT"
	TokenBulkScanResponseResultsFeaturesFeatureIDTransferFromReverts            TokenBulkScanResponseResultsFeaturesFeatureID = "TRANSFER_FROM_REVERTS"
	TokenBulkScanResponseResultsFeaturesFeatureIDOffensiveTokenMetadata         TokenBulkScanResponseResultsFeaturesFeatureID = "OFFENSIVE_TOKEN_METADATA"
	TokenBulkScanResponseResultsFeaturesFeatureIDListedOnCentralizedExchange    TokenBulkScanResponseResultsFeaturesFeatureID = "LISTED_ON_CENTRALIZED_EXCHANGE"
	TokenBulkScanResponseResultsFeaturesFeatureIDSanctionedCreator              TokenBulkScanResponseResultsFeaturesFeatureID = "SANCTIONED_CREATOR"
	TokenBulkScanResponseResultsFeaturesFeatureIDSpamText                       TokenBulkScanResponseResultsFeaturesFeatureID = "SPAM_TEXT"
	TokenBulkScanResponseResultsFeaturesFeatureIDBondingCurveToken              TokenBulkScanResponseResultsFeaturesFeatureID = "BONDING_CURVE_TOKEN"
	TokenBulkScanResponseResultsFeaturesFeatureIDHeavilySniped                  TokenBulkScanResponseResultsFeaturesFeatureID = "HEAVILY_SNIPED"
	TokenBulkScanResponseResultsFeaturesFeatureIDSolanaToken2022                TokenBulkScanResponseResultsFeaturesFeatureID = "SOLANA_TOKEN_2022"
	TokenBulkScanResponseResultsFeaturesFeatureIDPostDump                       TokenBulkScanResponseResultsFeaturesFeatureID = "POST_DUMP"
	TokenBulkScanResponseResultsFeaturesFeatureIDDexPaid                        TokenBulkScanResponseResultsFeaturesFeatureID = "DEX_PAID"
	TokenBulkScanResponseResultsFeaturesFeatureIDLowReputationCreator           TokenBulkScanResponseResultsFeaturesFeatureID = "LOW_REPUTATION_CREATOR"
	TokenBulkScanResponseResultsFeaturesFeatureIDSnipeAtMint                    TokenBulkScanResponseResultsFeaturesFeatureID = "SNIPE_AT_MINT"
	TokenBulkScanResponseResultsFeaturesFeatureIDTransferHookEnabled            TokenBulkScanResponseResultsFeaturesFeatureID = "TRANSFER_HOOK_ENABLED"
	TokenBulkScanResponseResultsFeaturesFeatureIDConfidentialTransfersEnabled   TokenBulkScanResponseResultsFeaturesFeatureID = "CONFIDENTIAL_TRANSFERS_ENABLED"
	TokenBulkScanResponseResultsFeaturesFeatureIDNonTranserable                 TokenBulkScanResponseResultsFeaturesFeatureID = "NON_TRANSERABLE"
	TokenBulkScanResponseResultsFeaturesFeatureIDTokenBackdoor                  TokenBulkScanResponseResultsFeaturesFeatureID = "TOKEN_BACKDOOR"
	TokenBulkScanResponseResultsFeaturesFeatureIDCreatedViaLaunchpad            TokenBulkScanResponseResultsFeaturesFeatureID = "CREATED_VIA_LAUNCHPAD"
	TokenBulkScanResponseResultsFeaturesFeatureIDCompromisedToken               TokenBulkScanResponseResultsFeaturesFeatureID = "COMPROMISED_TOKEN"
	TokenBulkScanResponseResultsFeaturesFeatureIDLongFundTrail                  TokenBulkScanResponseResultsFeaturesFeatureID = "LONG_FUND_TRAIL"
)

func (r TokenBulkScanResponseResultsFeaturesFeatureID) IsKnown() bool {
	switch r {
	case TokenBulkScanResponseResultsFeaturesFeatureIDVerifiedContract, TokenBulkScanResponseResultsFeaturesFeatureIDUnverifiedContract, TokenBulkScanResponseResultsFeaturesFeatureIDHighTradeVolume, TokenBulkScanResponseResultsFeaturesFeatureIDMarketPlaceSalesHistory, TokenBulkScanResponseResultsFeaturesFeatureIDHighReputationToken, TokenBulkScanResponseResultsFeaturesFeatureIDOnchainActivityValidator, TokenBulkScanResponseResultsFeaturesFeatureIDStaticCodeSignature, TokenBulkScanResponseResultsFeaturesFeatureIDKnownMalicious, TokenBulkScanResponseResultsFeaturesFeatureIDIsEoa, TokenBulkScanResponseResultsFeaturesFeatureIDIsContract, TokenBulkScanResponseResultsFeaturesFeatureIDErc20Contract, TokenBulkScanResponseResultsFeaturesFeatureIDTrustedContract, TokenBulkScanResponseResultsFeaturesFeatureIDBenignCreator, TokenBulkScanResponseResultsFeaturesFeatureIDMetadata, TokenBulkScanResponseResultsFeaturesFeatureIDAirdropPattern, TokenBulkScanResponseResultsFeaturesFeatureIDImpersonator, TokenBulkScanResponseResultsFeaturesFeatureIDInorganicVolume, TokenBulkScanResponseResultsFeaturesFeatureIDDynamicAnalysis, TokenBulkScanResponseResultsFeaturesFeatureIDConcentratedSupplyDistribution, TokenBulkScanResponseResultsFeaturesFeatureIDHoneypot, TokenBulkScanResponseResultsFeaturesFeatureIDInsufficientLockedLiquidity, TokenBulkScanResponseResultsFeaturesFeatureIDUnstableTokenPrice, TokenBulkScanResponseResultsFeaturesFeatureIDRugpull, TokenBulkScanResponseResultsFeaturesFeatureIDWashTrading, TokenBulkScanResponseResultsFeaturesFeatureIDConsumerOverride, TokenBulkScanResponseResultsFeaturesFeatureIDInappropriateContent, TokenBulkScanResponseResultsFeaturesFeatureIDHighTransferFee, TokenBulkScanResponseResultsFeaturesFeatureIDHighBuyFee, TokenBulkScanResponseResultsFeaturesFeatureIDHighSellFee, TokenBulkScanResponseResultsFeaturesFeatureIDUnsellableToken, TokenBulkScanResponseResultsFeaturesFeatureIDIsMintable, TokenBulkScanResponseResultsFeaturesFeatureIDRebaseToken, TokenBulkScanResponseResultsFeaturesFeatureIDLiquidStakingToken, TokenBulkScanResponseResultsFeaturesFeatureIDModifiableTaxes, TokenBulkScanResponseResultsFeaturesFeatureIDCanBlacklist, TokenBulkScanResponseResultsFeaturesFeatureIDCanWhitelist, TokenBulkScanResponseResultsFeaturesFeatureIDHasTradingCooldown, TokenBulkScanResponseResultsFeaturesFeatureIDExternalFunctions, TokenBulkScanResponseResultsFeaturesFeatureIDHiddenOwner, TokenBulkScanResponseResultsFeaturesFeatureIDTransferPauseable, TokenBulkScanResponseResultsFeaturesFeatureIDOwnershipRenounced, TokenBulkScanResponseResultsFeaturesFeatureIDOwnerCanChangeBalance, TokenBulkScanResponseResultsFeaturesFeatureIDProxyContract, TokenBulkScanResponseResultsFeaturesFeatureIDSimilarMaliciousContract, TokenBulkScanResponseResultsFeaturesFeatureIDImpersonatorSensitiveAsset, TokenBulkScanResponseResultsFeaturesFeatureIDImpersonatorHighConfidence, TokenBulkScanResponseResultsFeaturesFeatureIDImpersonatorMediumConfidence, TokenBulkScanResponseResultsFeaturesFeatureIDImpersonatorLowConfidence, TokenBulkScanResponseResultsFeaturesFeatureIDImpersonationProtected, TokenBulkScanResponseResultsFeaturesFeatureIDFakeVolume, TokenBulkScanResponseResultsFeaturesFeatureIDHiddenSupplyByKeyHolder, TokenBulkScanResponseResultsFeaturesFeatureIDFakeTradeMakerCount, TokenBulkScanResponseResultsFeaturesFeatureIDTransferFromReverts, TokenBulkScanResponseResultsFeaturesFeatureIDOffensiveTokenMetadata, TokenBulkScanResponseResultsFeaturesFeatureIDListedOnCentralizedExchange, TokenBulkScanResponseResultsFeaturesFeatureIDSanctionedCreator, TokenBulkScanResponseResultsFeaturesFeatureIDSpamText, TokenBulkScanResponseResultsFeaturesFeatureIDBondingCurveToken, TokenBulkScanResponseResultsFeaturesFeatureIDHeavilySniped, TokenBulkScanResponseResultsFeaturesFeatureIDSolanaToken2022, TokenBulkScanResponseResultsFeaturesFeatureIDPostDump, TokenBulkScanResponseResultsFeaturesFeatureIDDexPaid, TokenBulkScanResponseResultsFeaturesFeatureIDLowReputationCreator, TokenBulkScanResponseResultsFeaturesFeatureIDSnipeAtMint, TokenBulkScanResponseResultsFeaturesFeatureIDTransferHookEnabled, TokenBulkScanResponseResultsFeaturesFeatureIDConfidentialTransfersEnabled, TokenBulkScanResponseResultsFeaturesFeatureIDNonTranserable, TokenBulkScanResponseResultsFeaturesFeatureIDTokenBackdoor, TokenBulkScanResponseResultsFeaturesFeatureIDCreatedViaLaunchpad, TokenBulkScanResponseResultsFeaturesFeatureIDCompromisedToken, TokenBulkScanResponseResultsFeaturesFeatureIDLongFundTrail:
		return true
	}
	return false
}

// Type of the feature
type TokenBulkScanResponseResultsFeaturesType string

const (
	TokenBulkScanResponseResultsFeaturesTypeBenign    TokenBulkScanResponseResultsFeaturesType = "Benign"
	TokenBulkScanResponseResultsFeaturesTypeInfo      TokenBulkScanResponseResultsFeaturesType = "Info"
	TokenBulkScanResponseResultsFeaturesTypeWarning   TokenBulkScanResponseResultsFeaturesType = "Warning"
	TokenBulkScanResponseResultsFeaturesTypeMalicious TokenBulkScanResponseResultsFeaturesType = "Malicious"
)

func (r TokenBulkScanResponseResultsFeaturesType) IsKnown() bool {
	switch r {
	case TokenBulkScanResponseResultsFeaturesTypeBenign, TokenBulkScanResponseResultsFeaturesTypeInfo, TokenBulkScanResponseResultsFeaturesTypeWarning, TokenBulkScanResponseResultsFeaturesTypeMalicious:
		return true
	}
	return false
}

type TokenBulkScanResponseError struct {
	// Error message describing why the scan failed for this token
	Message TokenBulkScanResponseErrorsMessage `json:"message" api:"required"`
	JSON    tokenBulkScanResponseErrorJSON     `json:"-"`
}

// tokenBulkScanResponseErrorJSON contains the JSON metadata for the struct
// [TokenBulkScanResponseError]
type tokenBulkScanResponseErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenBulkScanResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseErrorJSON) RawJSON() string {
	return r.raw
}

// Error message describing why the scan failed for this token
type TokenBulkScanResponseErrorsMessage string

const (
	TokenBulkScanResponseErrorsMessagePendingScanResult TokenBulkScanResponseErrorsMessage = "PENDING_SCAN_RESULT"
	TokenBulkScanResponseErrorsMessageNotAToken         TokenBulkScanResponseErrorsMessage = "NOT_A_TOKEN"
)

func (r TokenBulkScanResponseErrorsMessage) IsKnown() bool {
	switch r {
	case TokenBulkScanResponseErrorsMessagePendingScanResult, TokenBulkScanResponseErrorsMessageNotAToken:
		return true
	}
	return false
}

type TokenBulkScanParams struct {
	// The chain name
	Chain param.Field[TokenScanSupportedChain] `json:"chain" api:"required"`
	// A list of token addresses to scan
	Tokens param.Field[[]string] `json:"tokens" api:"required"`
	// Optional token metadata context (e.g., source/integration hints) used to enrich
	// results.
	Metadata param.Field[TokenBulkScanParamsMetadata] `json:"metadata"`
}

func (r TokenBulkScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional token metadata context (e.g., source/integration hints) used to enrich
// results.
type TokenBulkScanParamsMetadata struct {
	// cross reference transaction against the domain.
	Domain param.Field[string] `json:"domain"`
}

func (r TokenBulkScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
