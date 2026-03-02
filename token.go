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

// TokenService contains methods and other services that help with interacting with
// the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTokenService] method instead.
type TokenService struct {
	Options []option.RequestOption
}

// NewTokenService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTokenService(opts ...option.RequestOption) (r *TokenService) {
	r = &TokenService{}
	r.Options = opts
	return
}

// Report for misclassification of a token.
func (r *TokenService) Report(ctx context.Context, body TokenReportParams, opts ...option.RequestOption) (res *TokenReportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/token/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets a token address and scan the token to identify any indication of malicious
// behavior
func (r *TokenService) Scan(ctx context.Context, body TokenScanParams, opts ...option.RequestOption) (res *TokenScanResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/token/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type FinancialStats struct {
	// Percentage of token currently held by bundlers - wallets that bought in the
	// exact same Solana slot, at any point in the token's life-cycle. Currently
	// available for Solana only.
	BundlersHoldingPercentage float64 `json:"bundlers_holding_percentage" api:"nullable"`
	// Token liquidity burned percentage
	BurnedLiquidityPercentage float64 `json:"burned_liquidity_percentage" api:"nullable"`
	// Percentage of token's supply held in known developer wallets (0.0 to 100.0)
	DevHoldingPercentage float64 `json:"dev_holding_percentage" api:"nullable"`
	// Amount of token holders
	HoldersCount int64 `json:"holders_count" api:"nullable"`
	// Percentage of token's supply _currently_ held by sniper bots (0.0 to 100.0).
	// Currently available for Solana only.
	InitialSnipersHoldingPercentage float64 `json:"initial_snipers_holding_percentage" api:"nullable"`
	// Percentage of supply that is currently held by insiders - defined as wallets
	// exhibiting early acquisition behaviors typically associated with insider
	// activity.
	InsidersHoldingPercentage float64 `json:"insiders_holding_percentage" api:"nullable"`
	// Token liquidity locked percentage
	LockedLiquidityPercentage float64 `json:"locked_liquidity_percentage" api:"nullable"`
	// Token markets/pools
	Markets []TokenMarket `json:"markets"`
	// Percentage of token's supply _initially_ held by sniper bots (0.0 to 100.0).
	// Currently available for Solana only.
	SnipersHoldingPercentage float64 `json:"snipers_holding_percentage" api:"nullable"`
	// token supply
	Supply int64 `json:"supply" api:"nullable"`
	// Top token holders
	TopHolders []TopHolder `json:"top_holders"`
	// Total reserve in USD
	TotalReserveInUsd float64 `json:"total_reserve_in_usd" api:"nullable"`
	// token price in USD
	UsdPricePerUnit float64            `json:"usd_price_per_unit" api:"nullable"`
	JSON            financialStatsJSON `json:"-"`
}

// financialStatsJSON contains the JSON metadata for the struct [FinancialStats]
type financialStatsJSON struct {
	BundlersHoldingPercentage       apijson.Field
	BurnedLiquidityPercentage       apijson.Field
	DevHoldingPercentage            apijson.Field
	HoldersCount                    apijson.Field
	InitialSnipersHoldingPercentage apijson.Field
	InsidersHoldingPercentage       apijson.Field
	LockedLiquidityPercentage       apijson.Field
	Markets                         apijson.Field
	SnipersHoldingPercentage        apijson.Field
	Supply                          apijson.Field
	TopHolders                      apijson.Field
	TotalReserveInUsd               apijson.Field
	UsdPricePerUnit                 apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *FinancialStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r financialStatsJSON) RawJSON() string {
	return r.raw
}

type MarketType string

const (
	MarketTypeBondingCurve MarketType = "BONDING_CURVE"
	MarketTypeAmm          MarketType = "AMM"
	MarketTypeUnknown      MarketType = "UNKNOWN"
)

func (r MarketType) IsKnown() bool {
	switch r {
	case MarketTypeBondingCurve, MarketTypeAmm, MarketTypeUnknown:
		return true
	}
	return false
}

type TokenMarket struct {
	Address      string          `json:"address" api:"required"`
	BaseToken    string          `json:"base_token" api:"required"`
	MarketName   string          `json:"market_name" api:"required"`
	MarketType   MarketType      `json:"market_type" api:"required"`
	PairName     string          `json:"pair_name" api:"required"`
	QuoteToken   string          `json:"quote_token" api:"required"`
	ReserveInUsd float64         `json:"reserve_in_usd" api:"required"`
	JSON         tokenMarketJSON `json:"-"`
}

// tokenMarketJSON contains the JSON metadata for the struct [TokenMarket]
type tokenMarketJSON struct {
	Address      apijson.Field
	BaseToken    apijson.Field
	MarketName   apijson.Field
	MarketType   apijson.Field
	PairName     apijson.Field
	QuoteToken   apijson.Field
	ReserveInUsd apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TokenMarket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenMarketJSON) RawJSON() string {
	return r.raw
}

type TopHolder struct {
	// Address
	Address string `json:"address" api:"nullable"`
	// Holding position out of total token liquidity
	HoldingPercentage float64        `json:"holding_percentage" api:"nullable"`
	Label             TopHolderLabel `json:"label" api:"nullable"`
	Name              string         `json:"name" api:"nullable"`
	JSON              topHolderJSON  `json:"-"`
}

// topHolderJSON contains the JSON metadata for the struct [TopHolder]
type topHolderJSON struct {
	Address           apijson.Field
	HoldingPercentage apijson.Field
	Label             apijson.Field
	Name              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TopHolder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r topHolderJSON) RawJSON() string {
	return r.raw
}

type TopHolderLabel string

const (
	TopHolderLabelMarket   TopHolderLabel = "market"
	TopHolderLabelLocker   TopHolderLabel = "locker"
	TopHolderLabelWallet   TopHolderLabel = "wallet"
	TopHolderLabelContract TopHolderLabel = "contract"
	TopHolderLabelProgram  TopHolderLabel = "program"
)

func (r TopHolderLabel) IsKnown() bool {
	switch r {
	case TopHolderLabelMarket, TopHolderLabelLocker, TopHolderLabelWallet, TopHolderLabelContract, TopHolderLabelProgram:
		return true
	}
	return false
}

type TokenReportResponse = interface{}

type TokenScanResponse struct {
	// Token address to validate (EVM / Solana)
	Address string `json:"address" api:"required"`
	// Dictionary of detected attacks found during the scan
	AttackTypes map[string]TokenScanResponseAttackType `json:"attack_types" api:"required"`
	// Blockchain network
	Chain TokenScanSupportedChain `json:"chain" api:"required"`
	// Fees associated with the token
	Fees TokenScanResponseFees `json:"fees" api:"required"`
	// financial stats of the token
	FinancialStats FinancialStats `json:"financial_stats" api:"required"`
	// Score between 0 to 1 (double)
	MaliciousScore string `json:"malicious_score" api:"required"`
	// Metadata of the token
	Metadata TokenScanResponseMetadata `json:"metadata" api:"required"`
	// General indication
	ResultType TokenScanResponseResultType `json:"result_type" api:"required"`
	// Trading limits of the token
	TradingLimits TokenScanResponseTradingLimits `json:"trading_limits" api:"required"`
	// List of features associated with the token
	Features []TokenScanResponseFeature `json:"features"`
	JSON     tokenScanResponseJSON      `json:"-"`
}

// tokenScanResponseJSON contains the JSON metadata for the struct
// [TokenScanResponse]
type tokenScanResponseJSON struct {
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

func (r *TokenScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseJSON) RawJSON() string {
	return r.raw
}

type TokenScanResponseAttackType struct {
	// Score between 0 to 1 (double) that indicates the assurance this attack happened
	Score string `json:"score" api:"required"`
	// Object contains an extra information related to the attack
	Features interface{} `json:"features"`
	// If score is higher or equal to this field, the token is using this attack type
	Threshold string                          `json:"threshold"`
	JSON      tokenScanResponseAttackTypeJSON `json:"-"`
}

// tokenScanResponseAttackTypeJSON contains the JSON metadata for the struct
// [TokenScanResponseAttackType]
type tokenScanResponseAttackTypeJSON struct {
	Score       apijson.Field
	Features    apijson.Field
	Threshold   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenScanResponseAttackType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseAttackTypeJSON) RawJSON() string {
	return r.raw
}

// Fees associated with the token
type TokenScanResponseFees struct {
	// Buy fee of the token
	Buy float64 `json:"buy" api:"nullable"`
	// Sell fee of the token
	Sell float64 `json:"sell" api:"nullable"`
	// Transfer fee of the token
	Transfer float64 `json:"transfer" api:"nullable"`
	// The maximum value that a transfer fee will cost
	TransferFeeMaxAmount int64                     `json:"transfer_fee_max_amount" api:"nullable"`
	JSON                 tokenScanResponseFeesJSON `json:"-"`
}

// tokenScanResponseFeesJSON contains the JSON metadata for the struct
// [TokenScanResponseFees]
type tokenScanResponseFeesJSON struct {
	Buy                  apijson.Field
	Sell                 apijson.Field
	Transfer             apijson.Field
	TransferFeeMaxAmount apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TokenScanResponseFees) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseFeesJSON) RawJSON() string {
	return r.raw
}

// Metadata of the token
type TokenScanResponseMetadata struct {
	// The unique ID for the Rune
	ID string `json:"id" api:"nullable"`
	// This field can have the runtime type of
	// [TokenScanResponseMetadataSolanaMetadataContractBalance],
	// [TokenScanResponseMetadataEvmMetadataTokenContractBalance].
	ContractBalance interface{} `json:"contract_balance"`
	// Contract deploy date
	CreationTimestamp string `json:"creation_timestamp" api:"nullable"`
	// Decimals of the token
	Decimals int64 `json:"decimals" api:"nullable"`
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer" api:"nullable"`
	// This field can have the runtime type of
	// [TokenScanResponseMetadataSolanaMetadataDeployerBalance],
	// [TokenScanResponseMetadataEvmMetadataTokenDeployerBalance].
	DeployerBalance interface{} `json:"deployer_balance"`
	// Description of the token
	Description string `json:"description" api:"nullable"`
	// This field can have the runtime type of
	// [TokenScanResponseMetadataSolanaMetadataExternalLinks],
	// [TokenScanResponseMetadataEvmMetadataTokenExternalLinks].
	ExternalLinks interface{} `json:"external_links"`
	// The formatted name of the rune, with spacers
	FormattedName string `json:"formatted_name" api:"nullable"`
	// Solana token freeze authority account
	FreezeAuthority string `json:"freeze_authority" api:"nullable"`
	// URL of the token image
	ImageURL string `json:"image_url" api:"nullable"`
	// This field can have the runtime type of
	// [[]TokenScanResponseMetadataSolanaMetadataImpersonationTarget],
	// [[]TokenScanResponseMetadataEvmMetadataTokenImpersonationTarget].
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
	// [TokenScanResponseMetadataSolanaMetadataOwnerBalance],
	// [TokenScanResponseMetadataEvmMetadataTokenOwnerBalance].
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
	URLs  interface{}                   `json:"urls"`
	JSON  tokenScanResponseMetadataJSON `json:"-"`
	union TokenScanResponseMetadataUnion
}

// tokenScanResponseMetadataJSON contains the JSON metadata for the struct
// [TokenScanResponseMetadata]
type tokenScanResponseMetadataJSON struct {
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

func (r tokenScanResponseMetadataJSON) RawJSON() string {
	return r.raw
}

func (r *TokenScanResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	*r = TokenScanResponseMetadata{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TokenScanResponseMetadataUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TokenScanResponseMetadataSolanaMetadata],
// [TokenScanResponseMetadataBitcoinMetadataToken],
// [TokenScanResponseMetadataEvmMetadataToken].
func (r TokenScanResponseMetadata) AsUnion() TokenScanResponseMetadataUnion {
	return r.union
}

// Metadata of the token
//
// Union satisfied by [TokenScanResponseMetadataSolanaMetadata],
// [TokenScanResponseMetadataBitcoinMetadataToken] or
// [TokenScanResponseMetadataEvmMetadataToken].
type TokenScanResponseMetadataUnion interface {
	implementsTokenScanResponseMetadata()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TokenScanResponseMetadataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenScanResponseMetadataSolanaMetadata{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenScanResponseMetadataBitcoinMetadataToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenScanResponseMetadataEvmMetadataToken{}),
		},
	)
}

type TokenScanResponseMetadataSolanaMetadata struct {
	// Contract balance
	ContractBalance TokenScanResponseMetadataSolanaMetadataContractBalance `json:"contract_balance" api:"nullable"`
	// Contract deploy date
	CreationTimestamp string `json:"creation_timestamp" api:"nullable"`
	// Decimals of the token
	Decimals int64 `json:"decimals" api:"nullable"`
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer" api:"nullable"`
	// Contract creator balance
	DeployerBalance TokenScanResponseMetadataSolanaMetadataDeployerBalance `json:"deployer_balance" api:"nullable"`
	// Description of the token
	Description string `json:"description" api:"nullable"`
	// social links of the token
	ExternalLinks TokenScanResponseMetadataSolanaMetadataExternalLinks `json:"external_links"`
	// Solana token freeze authority account
	FreezeAuthority string `json:"freeze_authority" api:"nullable"`
	// URL of the token image
	ImageURL string `json:"image_url" api:"nullable"`
	// List of tokens that this token is impersonating, if detected as an impersonator
	ImpersonationTargets []TokenScanResponseMetadataSolanaMetadataImpersonationTarget `json:"impersonation_targets" api:"nullable"`
	// Malicious urls associated with the token
	MaliciousURLs []string `json:"malicious_urls" api:"nullable"`
	// Solana token mint authority account
	MintAuthority string `json:"mint_authority" api:"nullable"`
	// Name of the token
	Name string `json:"name" api:"nullable"`
	// Contract owner address
	Owner string `json:"owner" api:"nullable"`
	// Contract owner balance
	OwnerBalance TokenScanResponseMetadataSolanaMetadataOwnerBalance `json:"owner_balance" api:"nullable"`
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
	URLs []string                                    `json:"urls" api:"nullable"`
	JSON tokenScanResponseMetadataSolanaMetadataJSON `json:"-"`
}

// tokenScanResponseMetadataSolanaMetadataJSON contains the JSON metadata for the
// struct [TokenScanResponseMetadataSolanaMetadata]
type tokenScanResponseMetadataSolanaMetadataJSON struct {
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

func (r *TokenScanResponseMetadataSolanaMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseMetadataSolanaMetadataJSON) RawJSON() string {
	return r.raw
}

func (r TokenScanResponseMetadataSolanaMetadata) implementsTokenScanResponseMetadata() {}

// Contract balance
type TokenScanResponseMetadataSolanaMetadataContractBalance struct {
	Amount    float64                                                    `json:"amount" api:"nullable"`
	AmountWei string                                                     `json:"amount_wei" api:"nullable"`
	JSON      tokenScanResponseMetadataSolanaMetadataContractBalanceJSON `json:"-"`
}

// tokenScanResponseMetadataSolanaMetadataContractBalanceJSON contains the JSON
// metadata for the struct [TokenScanResponseMetadataSolanaMetadataContractBalance]
type tokenScanResponseMetadataSolanaMetadataContractBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenScanResponseMetadataSolanaMetadataContractBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseMetadataSolanaMetadataContractBalanceJSON) RawJSON() string {
	return r.raw
}

// Contract creator balance
type TokenScanResponseMetadataSolanaMetadataDeployerBalance struct {
	Amount    float64                                                    `json:"amount" api:"nullable"`
	AmountWei string                                                     `json:"amount_wei" api:"nullable"`
	JSON      tokenScanResponseMetadataSolanaMetadataDeployerBalanceJSON `json:"-"`
}

// tokenScanResponseMetadataSolanaMetadataDeployerBalanceJSON contains the JSON
// metadata for the struct [TokenScanResponseMetadataSolanaMetadataDeployerBalance]
type tokenScanResponseMetadataSolanaMetadataDeployerBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenScanResponseMetadataSolanaMetadataDeployerBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseMetadataSolanaMetadataDeployerBalanceJSON) RawJSON() string {
	return r.raw
}

// social links of the token
type TokenScanResponseMetadataSolanaMetadataExternalLinks struct {
	Homepage          string                                                   `json:"homepage" api:"nullable"`
	TelegramChannelID string                                                   `json:"telegram_channel_id" api:"nullable"`
	TwitterPage       string                                                   `json:"twitter_page" api:"nullable"`
	JSON              tokenScanResponseMetadataSolanaMetadataExternalLinksJSON `json:"-"`
}

// tokenScanResponseMetadataSolanaMetadataExternalLinksJSON contains the JSON
// metadata for the struct [TokenScanResponseMetadataSolanaMetadataExternalLinks]
type tokenScanResponseMetadataSolanaMetadataExternalLinksJSON struct {
	Homepage          apijson.Field
	TelegramChannelID apijson.Field
	TwitterPage       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TokenScanResponseMetadataSolanaMetadataExternalLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseMetadataSolanaMetadataExternalLinksJSON) RawJSON() string {
	return r.raw
}

type TokenScanResponseMetadataSolanaMetadataImpersonationTarget struct {
	// Address of the token being impersonated
	Address string `json:"address" api:"required"`
	// Blockchain network of the target token
	Chain string `json:"chain" api:"required"`
	// Name of the token being impersonated
	Name string `json:"name" api:"nullable"`
	// Source of the impersonation match
	Source TokenScanResponseMetadataSolanaMetadataImpersonationTargetsSource `json:"source" api:"nullable"`
	// Symbol of the token being impersonated
	Symbol string                                                         `json:"symbol" api:"nullable"`
	JSON   tokenScanResponseMetadataSolanaMetadataImpersonationTargetJSON `json:"-"`
}

// tokenScanResponseMetadataSolanaMetadataImpersonationTargetJSON contains the JSON
// metadata for the struct
// [TokenScanResponseMetadataSolanaMetadataImpersonationTarget]
type tokenScanResponseMetadataSolanaMetadataImpersonationTargetJSON struct {
	Address     apijson.Field
	Chain       apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenScanResponseMetadataSolanaMetadataImpersonationTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseMetadataSolanaMetadataImpersonationTargetJSON) RawJSON() string {
	return r.raw
}

// Source of the impersonation match
type TokenScanResponseMetadataSolanaMetadataImpersonationTargetsSource string

const (
	TokenScanResponseMetadataSolanaMetadataImpersonationTargetsSourceTopToken    TokenScanResponseMetadataSolanaMetadataImpersonationTargetsSource = "TOP_TOKEN"
	TokenScanResponseMetadataSolanaMetadataImpersonationTargetsSourceUserDefined TokenScanResponseMetadataSolanaMetadataImpersonationTargetsSource = "USER_DEFINED"
)

func (r TokenScanResponseMetadataSolanaMetadataImpersonationTargetsSource) IsKnown() bool {
	switch r {
	case TokenScanResponseMetadataSolanaMetadataImpersonationTargetsSourceTopToken, TokenScanResponseMetadataSolanaMetadataImpersonationTargetsSourceUserDefined:
		return true
	}
	return false
}

// Contract owner balance
type TokenScanResponseMetadataSolanaMetadataOwnerBalance struct {
	Amount    float64                                                 `json:"amount" api:"nullable"`
	AmountWei string                                                  `json:"amount_wei" api:"nullable"`
	JSON      tokenScanResponseMetadataSolanaMetadataOwnerBalanceJSON `json:"-"`
}

// tokenScanResponseMetadataSolanaMetadataOwnerBalanceJSON contains the JSON
// metadata for the struct [TokenScanResponseMetadataSolanaMetadataOwnerBalance]
type tokenScanResponseMetadataSolanaMetadataOwnerBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenScanResponseMetadataSolanaMetadataOwnerBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseMetadataSolanaMetadataOwnerBalanceJSON) RawJSON() string {
	return r.raw
}

type TokenScanResponseMetadataBitcoinMetadataToken struct {
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
	Type string                                            `json:"type" api:"nullable"`
	JSON tokenScanResponseMetadataBitcoinMetadataTokenJSON `json:"-"`
}

// tokenScanResponseMetadataBitcoinMetadataTokenJSON contains the JSON metadata for
// the struct [TokenScanResponseMetadataBitcoinMetadataToken]
type tokenScanResponseMetadataBitcoinMetadataTokenJSON struct {
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

func (r *TokenScanResponseMetadataBitcoinMetadataToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseMetadataBitcoinMetadataTokenJSON) RawJSON() string {
	return r.raw
}

func (r TokenScanResponseMetadataBitcoinMetadataToken) implementsTokenScanResponseMetadata() {}

type TokenScanResponseMetadataEvmMetadataToken struct {
	// Contract balance
	ContractBalance TokenScanResponseMetadataEvmMetadataTokenContractBalance `json:"contract_balance" api:"nullable"`
	// Contract deploy date
	CreationTimestamp string `json:"creation_timestamp" api:"nullable"`
	// Decimals of the token
	Decimals int64 `json:"decimals" api:"nullable"`
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer" api:"nullable"`
	// Contract creator balance
	DeployerBalance TokenScanResponseMetadataEvmMetadataTokenDeployerBalance `json:"deployer_balance" api:"nullable"`
	// Description of the token
	Description string `json:"description" api:"nullable"`
	// social links of the token
	ExternalLinks TokenScanResponseMetadataEvmMetadataTokenExternalLinks `json:"external_links"`
	// URL of the token image
	ImageURL string `json:"image_url" api:"nullable"`
	// List of tokens that this token is impersonating, if detected as an impersonator
	ImpersonationTargets []TokenScanResponseMetadataEvmMetadataTokenImpersonationTarget `json:"impersonation_targets" api:"nullable"`
	// Malicious urls associated with the token
	MaliciousURLs []string `json:"malicious_urls" api:"nullable"`
	// Name of the token
	Name string `json:"name" api:"nullable"`
	// Contract owner address
	Owner string `json:"owner" api:"nullable"`
	// Contract owner balance
	OwnerBalance TokenScanResponseMetadataEvmMetadataTokenOwnerBalance `json:"owner_balance" api:"nullable"`
	// Symbol of the token
	Symbol string `json:"symbol" api:"nullable"`
	// Address of the token creation initiator, only set if the tokens was created by a
	// well known token launch platform
	TokenCreationInitiator string `json:"token_creation_initiator" api:"nullable"`
	// Type of the token
	Type string `json:"type" api:"nullable"`
	// Urls associated with the token
	URLs []string                                      `json:"urls" api:"nullable"`
	JSON tokenScanResponseMetadataEvmMetadataTokenJSON `json:"-"`
}

// tokenScanResponseMetadataEvmMetadataTokenJSON contains the JSON metadata for the
// struct [TokenScanResponseMetadataEvmMetadataToken]
type tokenScanResponseMetadataEvmMetadataTokenJSON struct {
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

func (r *TokenScanResponseMetadataEvmMetadataToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseMetadataEvmMetadataTokenJSON) RawJSON() string {
	return r.raw
}

func (r TokenScanResponseMetadataEvmMetadataToken) implementsTokenScanResponseMetadata() {}

// Contract balance
type TokenScanResponseMetadataEvmMetadataTokenContractBalance struct {
	Amount    float64                                                      `json:"amount" api:"nullable"`
	AmountWei string                                                       `json:"amount_wei" api:"nullable"`
	JSON      tokenScanResponseMetadataEvmMetadataTokenContractBalanceJSON `json:"-"`
}

// tokenScanResponseMetadataEvmMetadataTokenContractBalanceJSON contains the JSON
// metadata for the struct
// [TokenScanResponseMetadataEvmMetadataTokenContractBalance]
type tokenScanResponseMetadataEvmMetadataTokenContractBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenScanResponseMetadataEvmMetadataTokenContractBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseMetadataEvmMetadataTokenContractBalanceJSON) RawJSON() string {
	return r.raw
}

// Contract creator balance
type TokenScanResponseMetadataEvmMetadataTokenDeployerBalance struct {
	Amount    float64                                                      `json:"amount" api:"nullable"`
	AmountWei string                                                       `json:"amount_wei" api:"nullable"`
	JSON      tokenScanResponseMetadataEvmMetadataTokenDeployerBalanceJSON `json:"-"`
}

// tokenScanResponseMetadataEvmMetadataTokenDeployerBalanceJSON contains the JSON
// metadata for the struct
// [TokenScanResponseMetadataEvmMetadataTokenDeployerBalance]
type tokenScanResponseMetadataEvmMetadataTokenDeployerBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenScanResponseMetadataEvmMetadataTokenDeployerBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseMetadataEvmMetadataTokenDeployerBalanceJSON) RawJSON() string {
	return r.raw
}

// social links of the token
type TokenScanResponseMetadataEvmMetadataTokenExternalLinks struct {
	Homepage          string                                                     `json:"homepage" api:"nullable"`
	TelegramChannelID string                                                     `json:"telegram_channel_id" api:"nullable"`
	TwitterPage       string                                                     `json:"twitter_page" api:"nullable"`
	JSON              tokenScanResponseMetadataEvmMetadataTokenExternalLinksJSON `json:"-"`
}

// tokenScanResponseMetadataEvmMetadataTokenExternalLinksJSON contains the JSON
// metadata for the struct [TokenScanResponseMetadataEvmMetadataTokenExternalLinks]
type tokenScanResponseMetadataEvmMetadataTokenExternalLinksJSON struct {
	Homepage          apijson.Field
	TelegramChannelID apijson.Field
	TwitterPage       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TokenScanResponseMetadataEvmMetadataTokenExternalLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseMetadataEvmMetadataTokenExternalLinksJSON) RawJSON() string {
	return r.raw
}

type TokenScanResponseMetadataEvmMetadataTokenImpersonationTarget struct {
	// Address of the token being impersonated
	Address string `json:"address" api:"required"`
	// Blockchain network of the target token
	Chain string `json:"chain" api:"required"`
	// Name of the token being impersonated
	Name string `json:"name" api:"nullable"`
	// Source of the impersonation match
	Source TokenScanResponseMetadataEvmMetadataTokenImpersonationTargetsSource `json:"source" api:"nullable"`
	// Symbol of the token being impersonated
	Symbol string                                                           `json:"symbol" api:"nullable"`
	JSON   tokenScanResponseMetadataEvmMetadataTokenImpersonationTargetJSON `json:"-"`
}

// tokenScanResponseMetadataEvmMetadataTokenImpersonationTargetJSON contains the
// JSON metadata for the struct
// [TokenScanResponseMetadataEvmMetadataTokenImpersonationTarget]
type tokenScanResponseMetadataEvmMetadataTokenImpersonationTargetJSON struct {
	Address     apijson.Field
	Chain       apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenScanResponseMetadataEvmMetadataTokenImpersonationTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseMetadataEvmMetadataTokenImpersonationTargetJSON) RawJSON() string {
	return r.raw
}

// Source of the impersonation match
type TokenScanResponseMetadataEvmMetadataTokenImpersonationTargetsSource string

const (
	TokenScanResponseMetadataEvmMetadataTokenImpersonationTargetsSourceTopToken    TokenScanResponseMetadataEvmMetadataTokenImpersonationTargetsSource = "TOP_TOKEN"
	TokenScanResponseMetadataEvmMetadataTokenImpersonationTargetsSourceUserDefined TokenScanResponseMetadataEvmMetadataTokenImpersonationTargetsSource = "USER_DEFINED"
)

func (r TokenScanResponseMetadataEvmMetadataTokenImpersonationTargetsSource) IsKnown() bool {
	switch r {
	case TokenScanResponseMetadataEvmMetadataTokenImpersonationTargetsSourceTopToken, TokenScanResponseMetadataEvmMetadataTokenImpersonationTargetsSourceUserDefined:
		return true
	}
	return false
}

// Contract owner balance
type TokenScanResponseMetadataEvmMetadataTokenOwnerBalance struct {
	Amount    float64                                                   `json:"amount" api:"nullable"`
	AmountWei string                                                    `json:"amount_wei" api:"nullable"`
	JSON      tokenScanResponseMetadataEvmMetadataTokenOwnerBalanceJSON `json:"-"`
}

// tokenScanResponseMetadataEvmMetadataTokenOwnerBalanceJSON contains the JSON
// metadata for the struct [TokenScanResponseMetadataEvmMetadataTokenOwnerBalance]
type tokenScanResponseMetadataEvmMetadataTokenOwnerBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenScanResponseMetadataEvmMetadataTokenOwnerBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseMetadataEvmMetadataTokenOwnerBalanceJSON) RawJSON() string {
	return r.raw
}

// General indication
type TokenScanResponseResultType string

const (
	TokenScanResponseResultTypeBenign    TokenScanResponseResultType = "Benign"
	TokenScanResponseResultTypeWarning   TokenScanResponseResultType = "Warning"
	TokenScanResponseResultTypeMalicious TokenScanResponseResultType = "Malicious"
	TokenScanResponseResultTypeSpam      TokenScanResponseResultType = "Spam"
)

func (r TokenScanResponseResultType) IsKnown() bool {
	switch r {
	case TokenScanResponseResultTypeBenign, TokenScanResponseResultTypeWarning, TokenScanResponseResultTypeMalicious, TokenScanResponseResultTypeSpam:
		return true
	}
	return false
}

// Trading limits of the token
type TokenScanResponseTradingLimits struct {
	// Max amount that can be bought at once
	MaxBuy TokenScanResponseTradingLimitsMaxBuy `json:"max_buy" api:"nullable"`
	// Max amount that can be held by a single address
	MaxHolding TokenScanResponseTradingLimitsMaxHolding `json:"max_holding" api:"nullable"`
	// Max amount that can be sold at once
	MaxSell TokenScanResponseTradingLimitsMaxSell `json:"max_sell" api:"nullable"`
	// Maximum amount of the token that can be sold in a block
	SellLimitPerBlock TokenScanResponseTradingLimitsSellLimitPerBlock `json:"sell_limit_per_block" api:"nullable"`
	JSON              tokenScanResponseTradingLimitsJSON              `json:"-"`
}

// tokenScanResponseTradingLimitsJSON contains the JSON metadata for the struct
// [TokenScanResponseTradingLimits]
type tokenScanResponseTradingLimitsJSON struct {
	MaxBuy            apijson.Field
	MaxHolding        apijson.Field
	MaxSell           apijson.Field
	SellLimitPerBlock apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TokenScanResponseTradingLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseTradingLimitsJSON) RawJSON() string {
	return r.raw
}

// Max amount that can be bought at once
type TokenScanResponseTradingLimitsMaxBuy struct {
	Amount    float64                                  `json:"amount" api:"nullable"`
	AmountWei string                                   `json:"amount_wei" api:"nullable"`
	JSON      tokenScanResponseTradingLimitsMaxBuyJSON `json:"-"`
}

// tokenScanResponseTradingLimitsMaxBuyJSON contains the JSON metadata for the
// struct [TokenScanResponseTradingLimitsMaxBuy]
type tokenScanResponseTradingLimitsMaxBuyJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenScanResponseTradingLimitsMaxBuy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseTradingLimitsMaxBuyJSON) RawJSON() string {
	return r.raw
}

// Max amount that can be held by a single address
type TokenScanResponseTradingLimitsMaxHolding struct {
	Amount    float64                                      `json:"amount" api:"nullable"`
	AmountWei string                                       `json:"amount_wei" api:"nullable"`
	JSON      tokenScanResponseTradingLimitsMaxHoldingJSON `json:"-"`
}

// tokenScanResponseTradingLimitsMaxHoldingJSON contains the JSON metadata for the
// struct [TokenScanResponseTradingLimitsMaxHolding]
type tokenScanResponseTradingLimitsMaxHoldingJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenScanResponseTradingLimitsMaxHolding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseTradingLimitsMaxHoldingJSON) RawJSON() string {
	return r.raw
}

// Max amount that can be sold at once
type TokenScanResponseTradingLimitsMaxSell struct {
	Amount    float64                                   `json:"amount" api:"nullable"`
	AmountWei string                                    `json:"amount_wei" api:"nullable"`
	JSON      tokenScanResponseTradingLimitsMaxSellJSON `json:"-"`
}

// tokenScanResponseTradingLimitsMaxSellJSON contains the JSON metadata for the
// struct [TokenScanResponseTradingLimitsMaxSell]
type tokenScanResponseTradingLimitsMaxSellJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenScanResponseTradingLimitsMaxSell) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseTradingLimitsMaxSellJSON) RawJSON() string {
	return r.raw
}

// Maximum amount of the token that can be sold in a block
type TokenScanResponseTradingLimitsSellLimitPerBlock struct {
	Amount    float64                                             `json:"amount" api:"nullable"`
	AmountWei string                                              `json:"amount_wei" api:"nullable"`
	JSON      tokenScanResponseTradingLimitsSellLimitPerBlockJSON `json:"-"`
}

// tokenScanResponseTradingLimitsSellLimitPerBlockJSON contains the JSON metadata
// for the struct [TokenScanResponseTradingLimitsSellLimitPerBlock]
type tokenScanResponseTradingLimitsSellLimitPerBlockJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenScanResponseTradingLimitsSellLimitPerBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseTradingLimitsSellLimitPerBlockJSON) RawJSON() string {
	return r.raw
}

type TokenScanResponseFeature struct {
	// Description of the feature
	Description string `json:"description" api:"required"`
	// Feature identifier
	FeatureID TokenScanResponseFeaturesFeatureID `json:"feature_id" api:"required"`
	// Type of the feature
	Type TokenScanResponseFeaturesType `json:"type" api:"required"`
	JSON tokenScanResponseFeatureJSON  `json:"-"`
}

// tokenScanResponseFeatureJSON contains the JSON metadata for the struct
// [TokenScanResponseFeature]
type tokenScanResponseFeatureJSON struct {
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenScanResponseFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseFeatureJSON) RawJSON() string {
	return r.raw
}

// Feature identifier
type TokenScanResponseFeaturesFeatureID string

const (
	TokenScanResponseFeaturesFeatureIDVerifiedContract               TokenScanResponseFeaturesFeatureID = "VERIFIED_CONTRACT"
	TokenScanResponseFeaturesFeatureIDUnverifiedContract             TokenScanResponseFeaturesFeatureID = "UNVERIFIED_CONTRACT"
	TokenScanResponseFeaturesFeatureIDHighTradeVolume                TokenScanResponseFeaturesFeatureID = "HIGH_TRADE_VOLUME"
	TokenScanResponseFeaturesFeatureIDMarketPlaceSalesHistory        TokenScanResponseFeaturesFeatureID = "MARKET_PLACE_SALES_HISTORY"
	TokenScanResponseFeaturesFeatureIDHighReputationToken            TokenScanResponseFeaturesFeatureID = "HIGH_REPUTATION_TOKEN"
	TokenScanResponseFeaturesFeatureIDOnchainActivityValidator       TokenScanResponseFeaturesFeatureID = "ONCHAIN_ACTIVITY_VALIDATOR"
	TokenScanResponseFeaturesFeatureIDStaticCodeSignature            TokenScanResponseFeaturesFeatureID = "STATIC_CODE_SIGNATURE"
	TokenScanResponseFeaturesFeatureIDKnownMalicious                 TokenScanResponseFeaturesFeatureID = "KNOWN_MALICIOUS"
	TokenScanResponseFeaturesFeatureIDIsEoa                          TokenScanResponseFeaturesFeatureID = "IS_EOA"
	TokenScanResponseFeaturesFeatureIDIsContract                     TokenScanResponseFeaturesFeatureID = "IS_CONTRACT"
	TokenScanResponseFeaturesFeatureIDErc20Contract                  TokenScanResponseFeaturesFeatureID = "ERC20_CONTRACT"
	TokenScanResponseFeaturesFeatureIDTrustedContract                TokenScanResponseFeaturesFeatureID = "TRUSTED_CONTRACT"
	TokenScanResponseFeaturesFeatureIDBenignCreator                  TokenScanResponseFeaturesFeatureID = "BENIGN_CREATOR"
	TokenScanResponseFeaturesFeatureIDMetadata                       TokenScanResponseFeaturesFeatureID = "METADATA"
	TokenScanResponseFeaturesFeatureIDAirdropPattern                 TokenScanResponseFeaturesFeatureID = "AIRDROP_PATTERN"
	TokenScanResponseFeaturesFeatureIDImpersonator                   TokenScanResponseFeaturesFeatureID = "IMPERSONATOR"
	TokenScanResponseFeaturesFeatureIDInorganicVolume                TokenScanResponseFeaturesFeatureID = "INORGANIC_VOLUME"
	TokenScanResponseFeaturesFeatureIDDynamicAnalysis                TokenScanResponseFeaturesFeatureID = "DYNAMIC_ANALYSIS"
	TokenScanResponseFeaturesFeatureIDConcentratedSupplyDistribution TokenScanResponseFeaturesFeatureID = "CONCENTRATED_SUPPLY_DISTRIBUTION"
	TokenScanResponseFeaturesFeatureIDHoneypot                       TokenScanResponseFeaturesFeatureID = "HONEYPOT"
	TokenScanResponseFeaturesFeatureIDInsufficientLockedLiquidity    TokenScanResponseFeaturesFeatureID = "INSUFFICIENT_LOCKED_LIQUIDITY"
	TokenScanResponseFeaturesFeatureIDUnstableTokenPrice             TokenScanResponseFeaturesFeatureID = "UNSTABLE_TOKEN_PRICE"
	TokenScanResponseFeaturesFeatureIDRugpull                        TokenScanResponseFeaturesFeatureID = "RUGPULL"
	TokenScanResponseFeaturesFeatureIDWashTrading                    TokenScanResponseFeaturesFeatureID = "WASH_TRADING"
	TokenScanResponseFeaturesFeatureIDConsumerOverride               TokenScanResponseFeaturesFeatureID = "CONSUMER_OVERRIDE"
	TokenScanResponseFeaturesFeatureIDInappropriateContent           TokenScanResponseFeaturesFeatureID = "INAPPROPRIATE_CONTENT"
	TokenScanResponseFeaturesFeatureIDHighTransferFee                TokenScanResponseFeaturesFeatureID = "HIGH_TRANSFER_FEE"
	TokenScanResponseFeaturesFeatureIDHighBuyFee                     TokenScanResponseFeaturesFeatureID = "HIGH_BUY_FEE"
	TokenScanResponseFeaturesFeatureIDHighSellFee                    TokenScanResponseFeaturesFeatureID = "HIGH_SELL_FEE"
	TokenScanResponseFeaturesFeatureIDUnsellableToken                TokenScanResponseFeaturesFeatureID = "UNSELLABLE_TOKEN"
	TokenScanResponseFeaturesFeatureIDIsMintable                     TokenScanResponseFeaturesFeatureID = "IS_MINTABLE"
	TokenScanResponseFeaturesFeatureIDRebaseToken                    TokenScanResponseFeaturesFeatureID = "REBASE_TOKEN"
	TokenScanResponseFeaturesFeatureIDLiquidStakingToken             TokenScanResponseFeaturesFeatureID = "LIQUID_STAKING_TOKEN"
	TokenScanResponseFeaturesFeatureIDModifiableTaxes                TokenScanResponseFeaturesFeatureID = "MODIFIABLE_TAXES"
	TokenScanResponseFeaturesFeatureIDCanBlacklist                   TokenScanResponseFeaturesFeatureID = "CAN_BLACKLIST"
	TokenScanResponseFeaturesFeatureIDCanWhitelist                   TokenScanResponseFeaturesFeatureID = "CAN_WHITELIST"
	TokenScanResponseFeaturesFeatureIDHasTradingCooldown             TokenScanResponseFeaturesFeatureID = "HAS_TRADING_COOLDOWN"
	TokenScanResponseFeaturesFeatureIDExternalFunctions              TokenScanResponseFeaturesFeatureID = "EXTERNAL_FUNCTIONS"
	TokenScanResponseFeaturesFeatureIDHiddenOwner                    TokenScanResponseFeaturesFeatureID = "HIDDEN_OWNER"
	TokenScanResponseFeaturesFeatureIDTransferPauseable              TokenScanResponseFeaturesFeatureID = "TRANSFER_PAUSEABLE"
	TokenScanResponseFeaturesFeatureIDOwnershipRenounced             TokenScanResponseFeaturesFeatureID = "OWNERSHIP_RENOUNCED"
	TokenScanResponseFeaturesFeatureIDOwnerCanChangeBalance          TokenScanResponseFeaturesFeatureID = "OWNER_CAN_CHANGE_BALANCE"
	TokenScanResponseFeaturesFeatureIDProxyContract                  TokenScanResponseFeaturesFeatureID = "PROXY_CONTRACT"
	TokenScanResponseFeaturesFeatureIDSimilarMaliciousContract       TokenScanResponseFeaturesFeatureID = "SIMILAR_MALICIOUS_CONTRACT"
	TokenScanResponseFeaturesFeatureIDFakeVolume                     TokenScanResponseFeaturesFeatureID = "FAKE_VOLUME"
	TokenScanResponseFeaturesFeatureIDHiddenSupplyByKeyHolder        TokenScanResponseFeaturesFeatureID = "HIDDEN_SUPPLY_BY_KEY_HOLDER"
	TokenScanResponseFeaturesFeatureIDFakeTradeMakerCount            TokenScanResponseFeaturesFeatureID = "FAKE_TRADE_MAKER_COUNT"
)

func (r TokenScanResponseFeaturesFeatureID) IsKnown() bool {
	switch r {
	case TokenScanResponseFeaturesFeatureIDVerifiedContract, TokenScanResponseFeaturesFeatureIDUnverifiedContract, TokenScanResponseFeaturesFeatureIDHighTradeVolume, TokenScanResponseFeaturesFeatureIDMarketPlaceSalesHistory, TokenScanResponseFeaturesFeatureIDHighReputationToken, TokenScanResponseFeaturesFeatureIDOnchainActivityValidator, TokenScanResponseFeaturesFeatureIDStaticCodeSignature, TokenScanResponseFeaturesFeatureIDKnownMalicious, TokenScanResponseFeaturesFeatureIDIsEoa, TokenScanResponseFeaturesFeatureIDIsContract, TokenScanResponseFeaturesFeatureIDErc20Contract, TokenScanResponseFeaturesFeatureIDTrustedContract, TokenScanResponseFeaturesFeatureIDBenignCreator, TokenScanResponseFeaturesFeatureIDMetadata, TokenScanResponseFeaturesFeatureIDAirdropPattern, TokenScanResponseFeaturesFeatureIDImpersonator, TokenScanResponseFeaturesFeatureIDInorganicVolume, TokenScanResponseFeaturesFeatureIDDynamicAnalysis, TokenScanResponseFeaturesFeatureIDConcentratedSupplyDistribution, TokenScanResponseFeaturesFeatureIDHoneypot, TokenScanResponseFeaturesFeatureIDInsufficientLockedLiquidity, TokenScanResponseFeaturesFeatureIDUnstableTokenPrice, TokenScanResponseFeaturesFeatureIDRugpull, TokenScanResponseFeaturesFeatureIDWashTrading, TokenScanResponseFeaturesFeatureIDConsumerOverride, TokenScanResponseFeaturesFeatureIDInappropriateContent, TokenScanResponseFeaturesFeatureIDHighTransferFee, TokenScanResponseFeaturesFeatureIDHighBuyFee, TokenScanResponseFeaturesFeatureIDHighSellFee, TokenScanResponseFeaturesFeatureIDUnsellableToken, TokenScanResponseFeaturesFeatureIDIsMintable, TokenScanResponseFeaturesFeatureIDRebaseToken, TokenScanResponseFeaturesFeatureIDLiquidStakingToken, TokenScanResponseFeaturesFeatureIDModifiableTaxes, TokenScanResponseFeaturesFeatureIDCanBlacklist, TokenScanResponseFeaturesFeatureIDCanWhitelist, TokenScanResponseFeaturesFeatureIDHasTradingCooldown, TokenScanResponseFeaturesFeatureIDExternalFunctions, TokenScanResponseFeaturesFeatureIDHiddenOwner, TokenScanResponseFeaturesFeatureIDTransferPauseable, TokenScanResponseFeaturesFeatureIDOwnershipRenounced, TokenScanResponseFeaturesFeatureIDOwnerCanChangeBalance, TokenScanResponseFeaturesFeatureIDProxyContract, TokenScanResponseFeaturesFeatureIDSimilarMaliciousContract, TokenScanResponseFeaturesFeatureIDFakeVolume, TokenScanResponseFeaturesFeatureIDHiddenSupplyByKeyHolder, TokenScanResponseFeaturesFeatureIDFakeTradeMakerCount:
		return true
	}
	return false
}

// Type of the feature
type TokenScanResponseFeaturesType string

const (
	TokenScanResponseFeaturesTypeBenign    TokenScanResponseFeaturesType = "Benign"
	TokenScanResponseFeaturesTypeInfo      TokenScanResponseFeaturesType = "Info"
	TokenScanResponseFeaturesTypeWarning   TokenScanResponseFeaturesType = "Warning"
	TokenScanResponseFeaturesTypeMalicious TokenScanResponseFeaturesType = "Malicious"
)

func (r TokenScanResponseFeaturesType) IsKnown() bool {
	switch r {
	case TokenScanResponseFeaturesTypeBenign, TokenScanResponseFeaturesTypeInfo, TokenScanResponseFeaturesTypeWarning, TokenScanResponseFeaturesTypeMalicious:
		return true
	}
	return false
}

type TokenReportParams struct {
	// Details about the report.
	Details param.Field[string] `json:"details" api:"required"`
	// The event type of the report. Could be `FALSE_POSITIVE` or `FALSE_NEGATIVE`.
	Event param.Field[TokenReportParamsEvent] `json:"event" api:"required"`
	// Parameters identifying the token to report, provided either as token details
	// (address and chain) or as a request ID from a previous scan.
	Report param.Field[TokenReportParamsReportUnion] `json:"report" api:"required"`
}

func (r TokenReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The event type of the report. Could be `FALSE_POSITIVE` or `FALSE_NEGATIVE`.
type TokenReportParamsEvent string

const (
	TokenReportParamsEventFalsePositive TokenReportParamsEvent = "FALSE_POSITIVE"
	TokenReportParamsEventFalseNegative TokenReportParamsEvent = "FALSE_NEGATIVE"
)

func (r TokenReportParamsEvent) IsKnown() bool {
	switch r {
	case TokenReportParamsEventFalsePositive, TokenReportParamsEventFalseNegative:
		return true
	}
	return false
}

// Parameters identifying the token to report, provided either as token details
// (address and chain) or as a request ID from a previous scan.
type TokenReportParamsReport struct {
	Type   param.Field[TokenReportParamsReportType] `json:"type" api:"required"`
	Params param.Field[interface{}]                 `json:"params"`
	// The request ID of a previous request. This can be found in the value of the
	// `x-request-id` field in the headers of the response of the previous request. For
	// instance: `6c3cf6c1-a80d-4927-91b9-03d841ea61fe`.
	RequestID param.Field[string] `json:"request_id"`
}

func (r TokenReportParamsReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TokenReportParamsReport) implementsTokenReportParamsReportUnion() {}

// Parameters identifying the token to report, provided either as token details
// (address and chain) or as a request ID from a previous scan.
//
// Satisfied by [TokenReportParamsReportParamReportTokenReportParams],
// [TokenReportParamsReportRequestIDReport], [TokenReportParamsReport].
type TokenReportParamsReportUnion interface {
	implementsTokenReportParamsReportUnion()
}

type TokenReportParamsReportParamReportTokenReportParams struct {
	Params param.Field[TokenReportParamsReportParamReportTokenReportParamsParams] `json:"params" api:"required"`
	Type   param.Field[TokenReportParamsReportParamReportTokenReportParamsType]   `json:"type" api:"required"`
}

func (r TokenReportParamsReportParamReportTokenReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TokenReportParamsReportParamReportTokenReportParams) implementsTokenReportParamsReportUnion() {
}

type TokenReportParamsReportParamReportTokenReportParamsParams struct {
	// The address of the token to report on.
	Address param.Field[string] `json:"address" api:"required"`
	// The chain name
	Chain param.Field[TokenScanSupportedChain] `json:"chain" api:"required"`
}

func (r TokenReportParamsReportParamReportTokenReportParamsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TokenReportParamsReportParamReportTokenReportParamsType string

const (
	TokenReportParamsReportParamReportTokenReportParamsTypeParams TokenReportParamsReportParamReportTokenReportParamsType = "params"
)

func (r TokenReportParamsReportParamReportTokenReportParamsType) IsKnown() bool {
	switch r {
	case TokenReportParamsReportParamReportTokenReportParamsTypeParams:
		return true
	}
	return false
}

type TokenReportParamsReportRequestIDReport struct {
	// The request ID of a previous request. This can be found in the value of the
	// `x-request-id` field in the headers of the response of the previous request. For
	// instance: `6c3cf6c1-a80d-4927-91b9-03d841ea61fe`.
	RequestID param.Field[string] `json:"request_id" api:"required"`
	// The type identifier indicating that a request ID from a previous scan is being
	// used.
	Type param.Field[TokenReportParamsReportRequestIDReportType] `json:"type" api:"required"`
}

func (r TokenReportParamsReportRequestIDReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TokenReportParamsReportRequestIDReport) implementsTokenReportParamsReportUnion() {}

// The type identifier indicating that a request ID from a previous scan is being
// used.
type TokenReportParamsReportRequestIDReportType string

const (
	TokenReportParamsReportRequestIDReportTypeRequestID TokenReportParamsReportRequestIDReportType = "request_id"
)

func (r TokenReportParamsReportRequestIDReportType) IsKnown() bool {
	switch r {
	case TokenReportParamsReportRequestIDReportTypeRequestID:
		return true
	}
	return false
}

type TokenReportParamsReportType string

const (
	TokenReportParamsReportTypeParams    TokenReportParamsReportType = "params"
	TokenReportParamsReportTypeRequestID TokenReportParamsReportType = "request_id"
)

func (r TokenReportParamsReportType) IsKnown() bool {
	switch r {
	case TokenReportParamsReportTypeParams, TokenReportParamsReportTypeRequestID:
		return true
	}
	return false
}

type TokenScanParams struct {
	// Token address to validate (EVM / Solana / Stellar / Starknet)
	Address param.Field[string] `json:"address" api:"required"`
	// The chain name
	Chain param.Field[TokenScanSupportedChain] `json:"chain" api:"required"`
	// Optional token metadata context (e.g., source/integration hints) used to enrich
	// results.
	Metadata param.Field[TokenScanParamsMetadata] `json:"metadata"`
	// The ID of the specific NFT within an ERC-721 or ERC-1155 collection.
	TokenID param.Field[int64] `json:"token_id"`
}

func (r TokenScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional token metadata context (e.g., source/integration hints) used to enrich
// results.
type TokenScanParamsMetadata struct {
	// cross reference transaction against the domain.
	Domain param.Field[string] `json:"domain"`
}

func (r TokenScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
