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
	opts = append(r.Options[:], opts...)
	path := "v0/token/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets a token address and scan the token to identify any indication of malicious
// behaviour
func (r *TokenService) Scan(ctx context.Context, body TokenScanParams, opts ...option.RequestOption) (res *TokenScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/token/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type TokenReportResponse = interface{}

type TokenScanResponse struct {
	// Token address to validate (EVM / Solana)
	Address string `json:"address,required"`
	// Dictionary of detected attacks found during the scan
	AttackTypes map[string]TokenScanResponseAttackType `json:"attack_types,required"`
	// Blockchain network
	Chain TokenScanResponseChain `json:"chain,required"`
	// Fees associated with the token
	Fees TokenScanResponseFees `json:"fees,required"`
	// financial stats of the token
	FinancialStats TokenScanResponseFinancialStats `json:"financial_stats,required"`
	// Score between 0 to 1 (double)
	MaliciousScore string `json:"malicious_score,required"`
	// Metadata of the token
	Metadata TokenScanResponseMetadata `json:"metadata,required"`
	// General indication
	ResultType TokenScanResponseResultType `json:"result_type,required"`
	// Trading limits of the token
	TradingLimits TokenScanResponseTradingLimits `json:"trading_limits,required"`
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
	Score string `json:"score,required"`
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

// Blockchain network
type TokenScanResponseChain string

const (
	TokenScanResponseChainArbitrum    TokenScanResponseChain = "arbitrum"
	TokenScanResponseChainAvalanche   TokenScanResponseChain = "avalanche"
	TokenScanResponseChainBase        TokenScanResponseChain = "base"
	TokenScanResponseChainBsc         TokenScanResponseChain = "bsc"
	TokenScanResponseChainEthereum    TokenScanResponseChain = "ethereum"
	TokenScanResponseChainOptimism    TokenScanResponseChain = "optimism"
	TokenScanResponseChainPolygon     TokenScanResponseChain = "polygon"
	TokenScanResponseChainZora        TokenScanResponseChain = "zora"
	TokenScanResponseChainSolana      TokenScanResponseChain = "solana"
	TokenScanResponseChainStarknet    TokenScanResponseChain = "starknet"
	TokenScanResponseChainStellar     TokenScanResponseChain = "stellar"
	TokenScanResponseChainLinea       TokenScanResponseChain = "linea"
	TokenScanResponseChainBlast       TokenScanResponseChain = "blast"
	TokenScanResponseChainZksync      TokenScanResponseChain = "zksync"
	TokenScanResponseChainScroll      TokenScanResponseChain = "scroll"
	TokenScanResponseChainDegen       TokenScanResponseChain = "degen"
	TokenScanResponseChainAbstract    TokenScanResponseChain = "abstract"
	TokenScanResponseChainSoneium     TokenScanResponseChain = "soneium"
	TokenScanResponseChainInk         TokenScanResponseChain = "ink"
	TokenScanResponseChainZeroNetwork TokenScanResponseChain = "zero-network"
	TokenScanResponseChainBerachain   TokenScanResponseChain = "berachain"
	TokenScanResponseChainUnichain    TokenScanResponseChain = "unichain"
)

func (r TokenScanResponseChain) IsKnown() bool {
	switch r {
	case TokenScanResponseChainArbitrum, TokenScanResponseChainAvalanche, TokenScanResponseChainBase, TokenScanResponseChainBsc, TokenScanResponseChainEthereum, TokenScanResponseChainOptimism, TokenScanResponseChainPolygon, TokenScanResponseChainZora, TokenScanResponseChainSolana, TokenScanResponseChainStarknet, TokenScanResponseChainStellar, TokenScanResponseChainLinea, TokenScanResponseChainBlast, TokenScanResponseChainZksync, TokenScanResponseChainScroll, TokenScanResponseChainDegen, TokenScanResponseChainAbstract, TokenScanResponseChainSoneium, TokenScanResponseChainInk, TokenScanResponseChainZeroNetwork, TokenScanResponseChainBerachain, TokenScanResponseChainUnichain:
		return true
	}
	return false
}

// Fees associated with the token
type TokenScanResponseFees struct {
	// Buy fee of the token
	Buy float64 `json:"buy,nullable"`
	// Sell fee of the token
	Sell float64 `json:"sell,nullable"`
	// Transfer fee of the token
	Transfer float64                   `json:"transfer,nullable"`
	JSON     tokenScanResponseFeesJSON `json:"-"`
}

// tokenScanResponseFeesJSON contains the JSON metadata for the struct
// [TokenScanResponseFees]
type tokenScanResponseFeesJSON struct {
	Buy         apijson.Field
	Sell        apijson.Field
	Transfer    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenScanResponseFees) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseFeesJSON) RawJSON() string {
	return r.raw
}

// financial stats of the token
type TokenScanResponseFinancialStats struct {
	BurnedLiquidityPercentage float64                                    `json:"burned_liquidity_percentage,nullable"`
	HoldersCount              int64                                      `json:"holders_count,nullable"`
	LockedLiquidityPercentage float64                                    `json:"locked_liquidity_percentage,nullable"`
	TopHolders                []TokenScanResponseFinancialStatsTopHolder `json:"top_holders"`
	// Total reserve in USD
	TotalReserveInUsd float64                             `json:"total_reserve_in_usd,nullable"`
	UsdPricePerUnit   float64                             `json:"usd_price_per_unit,nullable"`
	JSON              tokenScanResponseFinancialStatsJSON `json:"-"`
}

// tokenScanResponseFinancialStatsJSON contains the JSON metadata for the struct
// [TokenScanResponseFinancialStats]
type tokenScanResponseFinancialStatsJSON struct {
	BurnedLiquidityPercentage apijson.Field
	HoldersCount              apijson.Field
	LockedLiquidityPercentage apijson.Field
	TopHolders                apijson.Field
	TotalReserveInUsd         apijson.Field
	UsdPricePerUnit           apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *TokenScanResponseFinancialStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseFinancialStatsJSON) RawJSON() string {
	return r.raw
}

type TokenScanResponseFinancialStatsTopHolder struct {
	Address           string                                       `json:"address,nullable"`
	HoldingPercentage float64                                      `json:"holding_percentage,nullable"`
	JSON              tokenScanResponseFinancialStatsTopHolderJSON `json:"-"`
}

// tokenScanResponseFinancialStatsTopHolderJSON contains the JSON metadata for the
// struct [TokenScanResponseFinancialStatsTopHolder]
type tokenScanResponseFinancialStatsTopHolderJSON struct {
	Address           apijson.Field
	HoldingPercentage apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TokenScanResponseFinancialStatsTopHolder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseFinancialStatsTopHolderJSON) RawJSON() string {
	return r.raw
}

// Metadata of the token
type TokenScanResponseMetadata struct {
	// The unique ID for the Rune
	ID string `json:"id,nullable"`
	// This field can have the runtime type of
	// [TokenScanResponseMetadataTokenSolanaMetadataContractBalance],
	// [TokenScanResponseMetadataTokenEvmMetadataTokenContractBalance].
	ContractBalance interface{} `json:"contract_balance"`
	// Contract deploy date
	CreationTimestamp time.Time `json:"creation_timestamp,nullable" format:"date-time"`
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer,nullable"`
	// Description of the token
	Description string `json:"description,nullable"`
	// This field can have the runtime type of
	// [TokenScanResponseMetadataTokenSolanaMetadataExternalLinks],
	// [TokenScanResponseMetadataTokenEvmMetadataTokenExternalLinks].
	ExternalLinks interface{} `json:"external_links"`
	// The formatted name of the rune, with spacers
	FormattedName string `json:"formatted_name,nullable"`
	// Solana token freeze authority account
	FreezeAuthority string `json:"freeze_authority,nullable"`
	// URL of the token image
	ImageURL string `json:"image_url,nullable"`
	// Solana token mint authority account
	MintAuthority string `json:"mint_authority,nullable"`
	// Name of the token
	Name string `json:"name,nullable"`
	// The rune's unique sequential number.
	Number int64 `json:"number,nullable"`
	// Contract owner address
	Owner string `json:"owner,nullable"`
	// This field can have the runtime type of
	// [TokenScanResponseMetadataTokenSolanaMetadataOwnerBalance],
	// [TokenScanResponseMetadataTokenEvmMetadataTokenOwnerBalance].
	OwnerBalance interface{} `json:"owner_balance"`
	// Symbol of the token
	Symbol string `json:"symbol,nullable"`
	// Type of the token
	Type string `json:"type,nullable"`
	// Solana token update authority account
	UpdateAuthority string                        `json:"update_authority,nullable"`
	JSON            tokenScanResponseMetadataJSON `json:"-"`
	union           TokenScanResponseMetadataUnion
}

// tokenScanResponseMetadataJSON contains the JSON metadata for the struct
// [TokenScanResponseMetadata]
type tokenScanResponseMetadataJSON struct {
	ID                apijson.Field
	ContractBalance   apijson.Field
	CreationTimestamp apijson.Field
	Deployer          apijson.Field
	Description       apijson.Field
	ExternalLinks     apijson.Field
	FormattedName     apijson.Field
	FreezeAuthority   apijson.Field
	ImageURL          apijson.Field
	MintAuthority     apijson.Field
	Name              apijson.Field
	Number            apijson.Field
	Owner             apijson.Field
	OwnerBalance      apijson.Field
	Symbol            apijson.Field
	Type              apijson.Field
	UpdateAuthority   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
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
// [TokenScanResponseMetadataTokenSolanaMetadata],
// [TokenScanResponseMetadataTokenBitcoinMetadataToken],
// [TokenScanResponseMetadataTokenEvmMetadataToken].
func (r TokenScanResponseMetadata) AsUnion() TokenScanResponseMetadataUnion {
	return r.union
}

// Metadata of the token
//
// Union satisfied by [TokenScanResponseMetadataTokenSolanaMetadata],
// [TokenScanResponseMetadataTokenBitcoinMetadataToken] or
// [TokenScanResponseMetadataTokenEvmMetadataToken].
type TokenScanResponseMetadataUnion interface {
	implementsTokenScanResponseMetadata()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TokenScanResponseMetadataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenScanResponseMetadataTokenSolanaMetadata{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenScanResponseMetadataTokenBitcoinMetadataToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenScanResponseMetadataTokenEvmMetadataToken{}),
		},
	)
}

type TokenScanResponseMetadataTokenSolanaMetadata struct {
	// Contract balance
	ContractBalance TokenScanResponseMetadataTokenSolanaMetadataContractBalance `json:"contract_balance,nullable"`
	// Contract deploy date
	CreationTimestamp time.Time `json:"creation_timestamp,nullable" format:"date-time"`
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer,nullable"`
	// Description of the token
	Description string `json:"description,nullable"`
	// social links of the token
	ExternalLinks TokenScanResponseMetadataTokenSolanaMetadataExternalLinks `json:"external_links"`
	// Solana token freeze authority account
	FreezeAuthority string `json:"freeze_authority,nullable"`
	// URL of the token image
	ImageURL string `json:"image_url,nullable"`
	// Solana token mint authority account
	MintAuthority string `json:"mint_authority,nullable"`
	// Name of the token
	Name string `json:"name,nullable"`
	// Contract owner address
	Owner string `json:"owner,nullable"`
	// Contract owner balance
	OwnerBalance TokenScanResponseMetadataTokenSolanaMetadataOwnerBalance `json:"owner_balance,nullable"`
	// Symbol of the token
	Symbol string `json:"symbol,nullable"`
	// Type of the token
	Type string `json:"type,nullable"`
	// Solana token update authority account
	UpdateAuthority string                                           `json:"update_authority,nullable"`
	JSON            tokenScanResponseMetadataTokenSolanaMetadataJSON `json:"-"`
}

// tokenScanResponseMetadataTokenSolanaMetadataJSON contains the JSON metadata for
// the struct [TokenScanResponseMetadataTokenSolanaMetadata]
type tokenScanResponseMetadataTokenSolanaMetadataJSON struct {
	ContractBalance   apijson.Field
	CreationTimestamp apijson.Field
	Deployer          apijson.Field
	Description       apijson.Field
	ExternalLinks     apijson.Field
	FreezeAuthority   apijson.Field
	ImageURL          apijson.Field
	MintAuthority     apijson.Field
	Name              apijson.Field
	Owner             apijson.Field
	OwnerBalance      apijson.Field
	Symbol            apijson.Field
	Type              apijson.Field
	UpdateAuthority   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TokenScanResponseMetadataTokenSolanaMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseMetadataTokenSolanaMetadataJSON) RawJSON() string {
	return r.raw
}

func (r TokenScanResponseMetadataTokenSolanaMetadata) implementsTokenScanResponseMetadata() {}

// Contract balance
type TokenScanResponseMetadataTokenSolanaMetadataContractBalance struct {
	Amount    float64                                                         `json:"amount,nullable"`
	AmountWei string                                                          `json:"amount_wei,nullable"`
	JSON      tokenScanResponseMetadataTokenSolanaMetadataContractBalanceJSON `json:"-"`
}

// tokenScanResponseMetadataTokenSolanaMetadataContractBalanceJSON contains the
// JSON metadata for the struct
// [TokenScanResponseMetadataTokenSolanaMetadataContractBalance]
type tokenScanResponseMetadataTokenSolanaMetadataContractBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenScanResponseMetadataTokenSolanaMetadataContractBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseMetadataTokenSolanaMetadataContractBalanceJSON) RawJSON() string {
	return r.raw
}

// social links of the token
type TokenScanResponseMetadataTokenSolanaMetadataExternalLinks struct {
	Homepage          string                                                        `json:"homepage,nullable"`
	TelegramChannelID string                                                        `json:"telegram_channel_id,nullable"`
	TwitterPage       string                                                        `json:"twitter_page,nullable"`
	JSON              tokenScanResponseMetadataTokenSolanaMetadataExternalLinksJSON `json:"-"`
}

// tokenScanResponseMetadataTokenSolanaMetadataExternalLinksJSON contains the JSON
// metadata for the struct
// [TokenScanResponseMetadataTokenSolanaMetadataExternalLinks]
type tokenScanResponseMetadataTokenSolanaMetadataExternalLinksJSON struct {
	Homepage          apijson.Field
	TelegramChannelID apijson.Field
	TwitterPage       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TokenScanResponseMetadataTokenSolanaMetadataExternalLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseMetadataTokenSolanaMetadataExternalLinksJSON) RawJSON() string {
	return r.raw
}

// Contract owner balance
type TokenScanResponseMetadataTokenSolanaMetadataOwnerBalance struct {
	Amount    float64                                                      `json:"amount,nullable"`
	AmountWei string                                                       `json:"amount_wei,nullable"`
	JSON      tokenScanResponseMetadataTokenSolanaMetadataOwnerBalanceJSON `json:"-"`
}

// tokenScanResponseMetadataTokenSolanaMetadataOwnerBalanceJSON contains the JSON
// metadata for the struct
// [TokenScanResponseMetadataTokenSolanaMetadataOwnerBalance]
type tokenScanResponseMetadataTokenSolanaMetadataOwnerBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenScanResponseMetadataTokenSolanaMetadataOwnerBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseMetadataTokenSolanaMetadataOwnerBalanceJSON) RawJSON() string {
	return r.raw
}

type TokenScanResponseMetadataTokenBitcoinMetadataToken struct {
	// The unique ID for the Rune
	ID string `json:"id,nullable"`
	// The formatted name of the rune, with spacers
	FormattedName string `json:"formatted_name,nullable"`
	// Name of the token
	Name string `json:"name,nullable"`
	// The rune's unique sequential number.
	Number int64 `json:"number,nullable"`
	// Symbol of the token
	Symbol string `json:"symbol,nullable"`
	// Type of the token
	Type string                                                 `json:"type,nullable"`
	JSON tokenScanResponseMetadataTokenBitcoinMetadataTokenJSON `json:"-"`
}

// tokenScanResponseMetadataTokenBitcoinMetadataTokenJSON contains the JSON
// metadata for the struct [TokenScanResponseMetadataTokenBitcoinMetadataToken]
type tokenScanResponseMetadataTokenBitcoinMetadataTokenJSON struct {
	ID            apijson.Field
	FormattedName apijson.Field
	Name          apijson.Field
	Number        apijson.Field
	Symbol        apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TokenScanResponseMetadataTokenBitcoinMetadataToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseMetadataTokenBitcoinMetadataTokenJSON) RawJSON() string {
	return r.raw
}

func (r TokenScanResponseMetadataTokenBitcoinMetadataToken) implementsTokenScanResponseMetadata() {}

type TokenScanResponseMetadataTokenEvmMetadataToken struct {
	// Contract balance
	ContractBalance TokenScanResponseMetadataTokenEvmMetadataTokenContractBalance `json:"contract_balance,nullable"`
	// Contract deploy date
	CreationTimestamp time.Time `json:"creation_timestamp,nullable" format:"date-time"`
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer,nullable"`
	// Description of the token
	Description string `json:"description,nullable"`
	// social links of the token
	ExternalLinks TokenScanResponseMetadataTokenEvmMetadataTokenExternalLinks `json:"external_links"`
	// URL of the token image
	ImageURL string `json:"image_url,nullable"`
	// Name of the token
	Name string `json:"name,nullable"`
	// Contract owner address
	Owner string `json:"owner,nullable"`
	// Contract owner balance
	OwnerBalance TokenScanResponseMetadataTokenEvmMetadataTokenOwnerBalance `json:"owner_balance,nullable"`
	// Symbol of the token
	Symbol string `json:"symbol,nullable"`
	// Type of the token
	Type string                                             `json:"type,nullable"`
	JSON tokenScanResponseMetadataTokenEvmMetadataTokenJSON `json:"-"`
}

// tokenScanResponseMetadataTokenEvmMetadataTokenJSON contains the JSON metadata
// for the struct [TokenScanResponseMetadataTokenEvmMetadataToken]
type tokenScanResponseMetadataTokenEvmMetadataTokenJSON struct {
	ContractBalance   apijson.Field
	CreationTimestamp apijson.Field
	Deployer          apijson.Field
	Description       apijson.Field
	ExternalLinks     apijson.Field
	ImageURL          apijson.Field
	Name              apijson.Field
	Owner             apijson.Field
	OwnerBalance      apijson.Field
	Symbol            apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TokenScanResponseMetadataTokenEvmMetadataToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseMetadataTokenEvmMetadataTokenJSON) RawJSON() string {
	return r.raw
}

func (r TokenScanResponseMetadataTokenEvmMetadataToken) implementsTokenScanResponseMetadata() {}

// Contract balance
type TokenScanResponseMetadataTokenEvmMetadataTokenContractBalance struct {
	Amount    float64                                                           `json:"amount,nullable"`
	AmountWei string                                                            `json:"amount_wei,nullable"`
	JSON      tokenScanResponseMetadataTokenEvmMetadataTokenContractBalanceJSON `json:"-"`
}

// tokenScanResponseMetadataTokenEvmMetadataTokenContractBalanceJSON contains the
// JSON metadata for the struct
// [TokenScanResponseMetadataTokenEvmMetadataTokenContractBalance]
type tokenScanResponseMetadataTokenEvmMetadataTokenContractBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenScanResponseMetadataTokenEvmMetadataTokenContractBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseMetadataTokenEvmMetadataTokenContractBalanceJSON) RawJSON() string {
	return r.raw
}

// social links of the token
type TokenScanResponseMetadataTokenEvmMetadataTokenExternalLinks struct {
	Homepage          string                                                          `json:"homepage,nullable"`
	TelegramChannelID string                                                          `json:"telegram_channel_id,nullable"`
	TwitterPage       string                                                          `json:"twitter_page,nullable"`
	JSON              tokenScanResponseMetadataTokenEvmMetadataTokenExternalLinksJSON `json:"-"`
}

// tokenScanResponseMetadataTokenEvmMetadataTokenExternalLinksJSON contains the
// JSON metadata for the struct
// [TokenScanResponseMetadataTokenEvmMetadataTokenExternalLinks]
type tokenScanResponseMetadataTokenEvmMetadataTokenExternalLinksJSON struct {
	Homepage          apijson.Field
	TelegramChannelID apijson.Field
	TwitterPage       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TokenScanResponseMetadataTokenEvmMetadataTokenExternalLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseMetadataTokenEvmMetadataTokenExternalLinksJSON) RawJSON() string {
	return r.raw
}

// Contract owner balance
type TokenScanResponseMetadataTokenEvmMetadataTokenOwnerBalance struct {
	Amount    float64                                                        `json:"amount,nullable"`
	AmountWei string                                                         `json:"amount_wei,nullable"`
	JSON      tokenScanResponseMetadataTokenEvmMetadataTokenOwnerBalanceJSON `json:"-"`
}

// tokenScanResponseMetadataTokenEvmMetadataTokenOwnerBalanceJSON contains the JSON
// metadata for the struct
// [TokenScanResponseMetadataTokenEvmMetadataTokenOwnerBalance]
type tokenScanResponseMetadataTokenEvmMetadataTokenOwnerBalanceJSON struct {
	Amount      apijson.Field
	AmountWei   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenScanResponseMetadataTokenEvmMetadataTokenOwnerBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseMetadataTokenEvmMetadataTokenOwnerBalanceJSON) RawJSON() string {
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
	MaxBuy            TokenScanResponseTradingLimitsMaxBuy            `json:"max_buy,nullable"`
	MaxHolding        TokenScanResponseTradingLimitsMaxHolding        `json:"max_holding,nullable"`
	MaxSell           TokenScanResponseTradingLimitsMaxSell           `json:"max_sell,nullable"`
	SellLimitPerBlock TokenScanResponseTradingLimitsSellLimitPerBlock `json:"sell_limit_per_block,nullable"`
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

type TokenScanResponseTradingLimitsMaxBuy struct {
	Amount    float64                                  `json:"amount,nullable"`
	AmountWei string                                   `json:"amount_wei,nullable"`
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

type TokenScanResponseTradingLimitsMaxHolding struct {
	Amount    float64                                      `json:"amount,nullable"`
	AmountWei string                                       `json:"amount_wei,nullable"`
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

type TokenScanResponseTradingLimitsMaxSell struct {
	Amount    float64                                   `json:"amount,nullable"`
	AmountWei string                                    `json:"amount_wei,nullable"`
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

type TokenScanResponseTradingLimitsSellLimitPerBlock struct {
	Amount    float64                                             `json:"amount,nullable"`
	AmountWei string                                              `json:"amount_wei,nullable"`
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
	Description string `json:"description,required"`
	// Feature identifier
	FeatureID TokenScanResponseFeaturesFeatureID `json:"feature_id,required"`
	// Type of the feature
	Type TokenScanResponseFeaturesType `json:"type,required"`
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
	TokenScanResponseFeaturesFeatureIDHighTradeVolume                TokenScanResponseFeaturesFeatureID = "HIGH_TRADE_VOLUME"
	TokenScanResponseFeaturesFeatureIDMarketPlaceSalesHistory        TokenScanResponseFeaturesFeatureID = "MARKET_PLACE_SALES_HISTORY"
	TokenScanResponseFeaturesFeatureIDHighReputationToken            TokenScanResponseFeaturesFeatureID = "HIGH_REPUTATION_TOKEN"
	TokenScanResponseFeaturesFeatureIDOnchainActivityValidator       TokenScanResponseFeaturesFeatureID = "ONCHAIN_ACTIVITY_VALIDATOR"
	TokenScanResponseFeaturesFeatureIDStaticCodeSignature            TokenScanResponseFeaturesFeatureID = "STATIC_CODE_SIGNATURE"
	TokenScanResponseFeaturesFeatureIDKnownMalicious                 TokenScanResponseFeaturesFeatureID = "KNOWN_MALICIOUS"
	TokenScanResponseFeaturesFeatureIDMetadata                       TokenScanResponseFeaturesFeatureID = "METADATA"
	TokenScanResponseFeaturesFeatureIDAirdropPattern                 TokenScanResponseFeaturesFeatureID = "AIRDROP_PATTERN"
	TokenScanResponseFeaturesFeatureIDImpersonator                   TokenScanResponseFeaturesFeatureID = "IMPERSONATOR"
	TokenScanResponseFeaturesFeatureIDInorganicVolume                TokenScanResponseFeaturesFeatureID = "INORGANIC_VOLUME"
	TokenScanResponseFeaturesFeatureIDDynamicAnalysis                TokenScanResponseFeaturesFeatureID = "DYNAMIC_ANALYSIS"
	TokenScanResponseFeaturesFeatureIDUnstableTokenPrice             TokenScanResponseFeaturesFeatureID = "UNSTABLE_TOKEN_PRICE"
	TokenScanResponseFeaturesFeatureIDRugpull                        TokenScanResponseFeaturesFeatureID = "RUGPULL"
	TokenScanResponseFeaturesFeatureIDConsumerOverride               TokenScanResponseFeaturesFeatureID = "CONSUMER_OVERRIDE"
	TokenScanResponseFeaturesFeatureIDInappropriateContent           TokenScanResponseFeaturesFeatureID = "INAPPROPRIATE_CONTENT"
	TokenScanResponseFeaturesFeatureIDHighTransferFee                TokenScanResponseFeaturesFeatureID = "HIGH_TRANSFER_FEE"
	TokenScanResponseFeaturesFeatureIDHighBuyFee                     TokenScanResponseFeaturesFeatureID = "HIGH_BUY_FEE"
	TokenScanResponseFeaturesFeatureIDHighSellFee                    TokenScanResponseFeaturesFeatureID = "HIGH_SELL_FEE"
	TokenScanResponseFeaturesFeatureIDIsMintable                     TokenScanResponseFeaturesFeatureID = "IS_MINTABLE"
	TokenScanResponseFeaturesFeatureIDModifiableTaxes                TokenScanResponseFeaturesFeatureID = "MODIFIABLE_TAXES"
	TokenScanResponseFeaturesFeatureIDCanBlacklist                   TokenScanResponseFeaturesFeatureID = "CAN_BLACKLIST"
	TokenScanResponseFeaturesFeatureIDCanWhitelist                   TokenScanResponseFeaturesFeatureID = "CAN_WHITELIST"
	TokenScanResponseFeaturesFeatureIDHasTradingCooldown             TokenScanResponseFeaturesFeatureID = "HAS_TRADING_COOLDOWN"
	TokenScanResponseFeaturesFeatureIDExternalFunctions              TokenScanResponseFeaturesFeatureID = "EXTERNAL_FUNCTIONS"
	TokenScanResponseFeaturesFeatureIDHiddenOwner                    TokenScanResponseFeaturesFeatureID = "HIDDEN_OWNER"
	TokenScanResponseFeaturesFeatureIDTransferPauseable              TokenScanResponseFeaturesFeatureID = "TRANSFER_PAUSEABLE"
	TokenScanResponseFeaturesFeatureIDOwnershipRenounced             TokenScanResponseFeaturesFeatureID = "OWNERSHIP_RENOUNCED"
	TokenScanResponseFeaturesFeatureIDProxyContract                  TokenScanResponseFeaturesFeatureID = "PROXY_CONTRACT"
	TokenScanResponseFeaturesFeatureIDLiquidStake                    TokenScanResponseFeaturesFeatureID = "LIQUID_STAKE"
	TokenScanResponseFeaturesFeatureIDRebaseToken                    TokenScanResponseFeaturesFeatureID = "REBASE_TOKEN"
	TokenScanResponseFeaturesFeatureIDUnsellableToken                TokenScanResponseFeaturesFeatureID = "UNSELLABLE_TOKEN"
	TokenScanResponseFeaturesFeatureIDConcentratedSupplyDistribution TokenScanResponseFeaturesFeatureID = "CONCENTRATED_SUPPLY_DISTRIBUTION"
	TokenScanResponseFeaturesFeatureIDInsufficientLockedLiquidity    TokenScanResponseFeaturesFeatureID = "INSUFFICIENT_LOCKED_LIQUIDITY"
	TokenScanResponseFeaturesFeatureIDHoneypot                       TokenScanResponseFeaturesFeatureID = "HONEYPOT"
)

func (r TokenScanResponseFeaturesFeatureID) IsKnown() bool {
	switch r {
	case TokenScanResponseFeaturesFeatureIDVerifiedContract, TokenScanResponseFeaturesFeatureIDHighTradeVolume, TokenScanResponseFeaturesFeatureIDMarketPlaceSalesHistory, TokenScanResponseFeaturesFeatureIDHighReputationToken, TokenScanResponseFeaturesFeatureIDOnchainActivityValidator, TokenScanResponseFeaturesFeatureIDStaticCodeSignature, TokenScanResponseFeaturesFeatureIDKnownMalicious, TokenScanResponseFeaturesFeatureIDMetadata, TokenScanResponseFeaturesFeatureIDAirdropPattern, TokenScanResponseFeaturesFeatureIDImpersonator, TokenScanResponseFeaturesFeatureIDInorganicVolume, TokenScanResponseFeaturesFeatureIDDynamicAnalysis, TokenScanResponseFeaturesFeatureIDUnstableTokenPrice, TokenScanResponseFeaturesFeatureIDRugpull, TokenScanResponseFeaturesFeatureIDConsumerOverride, TokenScanResponseFeaturesFeatureIDInappropriateContent, TokenScanResponseFeaturesFeatureIDHighTransferFee, TokenScanResponseFeaturesFeatureIDHighBuyFee, TokenScanResponseFeaturesFeatureIDHighSellFee, TokenScanResponseFeaturesFeatureIDIsMintable, TokenScanResponseFeaturesFeatureIDModifiableTaxes, TokenScanResponseFeaturesFeatureIDCanBlacklist, TokenScanResponseFeaturesFeatureIDCanWhitelist, TokenScanResponseFeaturesFeatureIDHasTradingCooldown, TokenScanResponseFeaturesFeatureIDExternalFunctions, TokenScanResponseFeaturesFeatureIDHiddenOwner, TokenScanResponseFeaturesFeatureIDTransferPauseable, TokenScanResponseFeaturesFeatureIDOwnershipRenounced, TokenScanResponseFeaturesFeatureIDProxyContract, TokenScanResponseFeaturesFeatureIDLiquidStake, TokenScanResponseFeaturesFeatureIDRebaseToken, TokenScanResponseFeaturesFeatureIDUnsellableToken, TokenScanResponseFeaturesFeatureIDConcentratedSupplyDistribution, TokenScanResponseFeaturesFeatureIDInsufficientLockedLiquidity, TokenScanResponseFeaturesFeatureIDHoneypot:
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
	Details param.Field[string] `json:"details,required"`
	// The event type of the report. Could be FALSE_POSITIVE or FALSE_NEGATIVE.
	Event param.Field[TokenReportParamsEvent] `json:"event,required"`
	// The report parameters.
	Report param.Field[TokenReportParamsReportUnion] `json:"report,required"`
}

func (r TokenReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The event type of the report. Could be FALSE_POSITIVE or FALSE_NEGATIVE.
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

// The report parameters.
type TokenReportParamsReport struct {
	Type      param.Field[TokenReportParamsReportType] `json:"type,required"`
	Params    param.Field[interface{}]                 `json:"params"`
	RequestID param.Field[string]                      `json:"request_id"`
}

func (r TokenReportParamsReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TokenReportParamsReport) implementsTokenReportParamsReportUnion() {}

// The report parameters.
//
// Satisfied by [TokenReportParamsReportTokenParamReportTokenReportParams],
// [TokenReportParamsReportTokenRequestIDReport], [TokenReportParamsReport].
type TokenReportParamsReportUnion interface {
	implementsTokenReportParamsReportUnion()
}

type TokenReportParamsReportTokenParamReportTokenReportParams struct {
	Params param.Field[TokenReportParamsReportTokenParamReportTokenReportParamsParams] `json:"params,required"`
	Type   param.Field[TokenReportParamsReportTokenParamReportTokenReportParamsType]   `json:"type,required"`
}

func (r TokenReportParamsReportTokenParamReportTokenReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TokenReportParamsReportTokenParamReportTokenReportParams) implementsTokenReportParamsReportUnion() {
}

type TokenReportParamsReportTokenParamReportTokenReportParamsParams struct {
	// The address of the token to report on.
	Address param.Field[string] `json:"address,required"`
	// The chain name
	Chain param.Field[TokenReportParamsReportTokenParamReportTokenReportParamsParamsChain] `json:"chain,required"`
}

func (r TokenReportParamsReportTokenParamReportTokenReportParamsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The chain name
type TokenReportParamsReportTokenParamReportTokenReportParamsParamsChain string

const (
	TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainArbitrum    TokenReportParamsReportTokenParamReportTokenReportParamsParamsChain = "arbitrum"
	TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainAvalanche   TokenReportParamsReportTokenParamReportTokenReportParamsParamsChain = "avalanche"
	TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainBase        TokenReportParamsReportTokenParamReportTokenReportParamsParamsChain = "base"
	TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainBsc         TokenReportParamsReportTokenParamReportTokenReportParamsParamsChain = "bsc"
	TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainEthereum    TokenReportParamsReportTokenParamReportTokenReportParamsParamsChain = "ethereum"
	TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainOptimism    TokenReportParamsReportTokenParamReportTokenReportParamsParamsChain = "optimism"
	TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainPolygon     TokenReportParamsReportTokenParamReportTokenReportParamsParamsChain = "polygon"
	TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainZora        TokenReportParamsReportTokenParamReportTokenReportParamsParamsChain = "zora"
	TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainSolana      TokenReportParamsReportTokenParamReportTokenReportParamsParamsChain = "solana"
	TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainStarknet    TokenReportParamsReportTokenParamReportTokenReportParamsParamsChain = "starknet"
	TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainStellar     TokenReportParamsReportTokenParamReportTokenReportParamsParamsChain = "stellar"
	TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainLinea       TokenReportParamsReportTokenParamReportTokenReportParamsParamsChain = "linea"
	TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainBlast       TokenReportParamsReportTokenParamReportTokenReportParamsParamsChain = "blast"
	TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainZksync      TokenReportParamsReportTokenParamReportTokenReportParamsParamsChain = "zksync"
	TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainScroll      TokenReportParamsReportTokenParamReportTokenReportParamsParamsChain = "scroll"
	TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainDegen       TokenReportParamsReportTokenParamReportTokenReportParamsParamsChain = "degen"
	TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainAbstract    TokenReportParamsReportTokenParamReportTokenReportParamsParamsChain = "abstract"
	TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainSoneium     TokenReportParamsReportTokenParamReportTokenReportParamsParamsChain = "soneium"
	TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainInk         TokenReportParamsReportTokenParamReportTokenReportParamsParamsChain = "ink"
	TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainZeroNetwork TokenReportParamsReportTokenParamReportTokenReportParamsParamsChain = "zero-network"
	TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainBerachain   TokenReportParamsReportTokenParamReportTokenReportParamsParamsChain = "berachain"
	TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainUnichain    TokenReportParamsReportTokenParamReportTokenReportParamsParamsChain = "unichain"
)

func (r TokenReportParamsReportTokenParamReportTokenReportParamsParamsChain) IsKnown() bool {
	switch r {
	case TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainArbitrum, TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainAvalanche, TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainBase, TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainBsc, TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainEthereum, TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainOptimism, TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainPolygon, TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainZora, TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainSolana, TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainStarknet, TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainStellar, TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainLinea, TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainBlast, TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainZksync, TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainScroll, TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainDegen, TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainAbstract, TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainSoneium, TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainInk, TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainZeroNetwork, TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainBerachain, TokenReportParamsReportTokenParamReportTokenReportParamsParamsChainUnichain:
		return true
	}
	return false
}

type TokenReportParamsReportTokenParamReportTokenReportParamsType string

const (
	TokenReportParamsReportTokenParamReportTokenReportParamsTypeParams TokenReportParamsReportTokenParamReportTokenReportParamsType = "params"
)

func (r TokenReportParamsReportTokenParamReportTokenReportParamsType) IsKnown() bool {
	switch r {
	case TokenReportParamsReportTokenParamReportTokenReportParamsTypeParams:
		return true
	}
	return false
}

type TokenReportParamsReportTokenRequestIDReport struct {
	RequestID param.Field[string]                                          `json:"request_id,required"`
	Type      param.Field[TokenReportParamsReportTokenRequestIDReportType] `json:"type,required"`
}

func (r TokenReportParamsReportTokenRequestIDReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TokenReportParamsReportTokenRequestIDReport) implementsTokenReportParamsReportUnion() {}

type TokenReportParamsReportTokenRequestIDReportType string

const (
	TokenReportParamsReportTokenRequestIDReportTypeRequestID TokenReportParamsReportTokenRequestIDReportType = "request_id"
)

func (r TokenReportParamsReportTokenRequestIDReportType) IsKnown() bool {
	switch r {
	case TokenReportParamsReportTokenRequestIDReportTypeRequestID:
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
	Address param.Field[string] `json:"address,required"`
	// The chain name
	Chain param.Field[TokenScanParamsChain] `json:"chain,required"`
	// Object of additional information to validate against.
	Metadata param.Field[TokenScanParamsMetadata] `json:"metadata"`
}

func (r TokenScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The chain name
type TokenScanParamsChain string

const (
	TokenScanParamsChainArbitrum    TokenScanParamsChain = "arbitrum"
	TokenScanParamsChainAvalanche   TokenScanParamsChain = "avalanche"
	TokenScanParamsChainBase        TokenScanParamsChain = "base"
	TokenScanParamsChainBsc         TokenScanParamsChain = "bsc"
	TokenScanParamsChainEthereum    TokenScanParamsChain = "ethereum"
	TokenScanParamsChainOptimism    TokenScanParamsChain = "optimism"
	TokenScanParamsChainPolygon     TokenScanParamsChain = "polygon"
	TokenScanParamsChainZora        TokenScanParamsChain = "zora"
	TokenScanParamsChainSolana      TokenScanParamsChain = "solana"
	TokenScanParamsChainStarknet    TokenScanParamsChain = "starknet"
	TokenScanParamsChainStellar     TokenScanParamsChain = "stellar"
	TokenScanParamsChainLinea       TokenScanParamsChain = "linea"
	TokenScanParamsChainBlast       TokenScanParamsChain = "blast"
	TokenScanParamsChainZksync      TokenScanParamsChain = "zksync"
	TokenScanParamsChainScroll      TokenScanParamsChain = "scroll"
	TokenScanParamsChainDegen       TokenScanParamsChain = "degen"
	TokenScanParamsChainAbstract    TokenScanParamsChain = "abstract"
	TokenScanParamsChainSoneium     TokenScanParamsChain = "soneium"
	TokenScanParamsChainInk         TokenScanParamsChain = "ink"
	TokenScanParamsChainZeroNetwork TokenScanParamsChain = "zero-network"
	TokenScanParamsChainBerachain   TokenScanParamsChain = "berachain"
	TokenScanParamsChainUnichain    TokenScanParamsChain = "unichain"
)

func (r TokenScanParamsChain) IsKnown() bool {
	switch r {
	case TokenScanParamsChainArbitrum, TokenScanParamsChainAvalanche, TokenScanParamsChainBase, TokenScanParamsChainBsc, TokenScanParamsChainEthereum, TokenScanParamsChainOptimism, TokenScanParamsChainPolygon, TokenScanParamsChainZora, TokenScanParamsChainSolana, TokenScanParamsChainStarknet, TokenScanParamsChainStellar, TokenScanParamsChainLinea, TokenScanParamsChainBlast, TokenScanParamsChainZksync, TokenScanParamsChainScroll, TokenScanParamsChainDegen, TokenScanParamsChainAbstract, TokenScanParamsChainSoneium, TokenScanParamsChainInk, TokenScanParamsChainZeroNetwork, TokenScanParamsChainBerachain, TokenScanParamsChainUnichain:
		return true
	}
	return false
}

// Object of additional information to validate against.
type TokenScanParamsMetadata struct {
	// cross reference transaction against the domain.
	Domain param.Field[string] `json:"domain"`
}

func (r TokenScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
