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
	opts = append(r.Options[:], opts...)
	path := "v0/token-bulk/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type TokenBulkScanResponse struct {
	Results map[string]TokenBulkScanResponseResult `json:"results,required"`
	JSON    tokenBulkScanResponseJSON              `json:"-"`
}

// tokenBulkScanResponseJSON contains the JSON metadata for the struct
// [TokenBulkScanResponse]
type tokenBulkScanResponseJSON struct {
	Results     apijson.Field
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
	Address string `json:"address,required"`
	// Dictionary of detected attacks found during the scan
	AttackTypes map[string]TokenBulkScanResponseResultsAttackType `json:"attack_types,required"`
	// The chain name
	Chain TokenScanSupportedChain `json:"chain,required"`
	// Fees associated with the token
	Fees TokenBulkScanResponseResultsFees `json:"fees,required"`
	// financial stats of the token
	FinancialStats TokenBulkScanResponseResultsFinancialStats `json:"financial_stats,required"`
	// Score between 0 to 1 (double)
	MaliciousScore string `json:"malicious_score,required"`
	// Metadata of the token
	Metadata TokenBulkScanResponseResultsMetadata `json:"metadata,required"`
	// An enumeration.
	ResultType TokenBulkScanResponseResultsResultType `json:"result_type,required"`
	// Trading limits of the token
	TradingLimits TokenBulkScanResponseResultsTradingLimits `json:"trading_limits,required"`
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
	Score string `json:"score,required"`
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
	Buy float64 `json:"buy"`
	// Sell fee of the token
	Sell float64 `json:"sell"`
	// Transfer fee of the token
	Transfer float64                              `json:"transfer"`
	JSON     tokenBulkScanResponseResultsFeesJSON `json:"-"`
}

// tokenBulkScanResponseResultsFeesJSON contains the JSON metadata for the struct
// [TokenBulkScanResponseResultsFees]
type tokenBulkScanResponseResultsFeesJSON struct {
	Buy         apijson.Field
	Sell        apijson.Field
	Transfer    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenBulkScanResponseResultsFees) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsFeesJSON) RawJSON() string {
	return r.raw
}

// financial stats of the token
type TokenBulkScanResponseResultsFinancialStats struct {
	BurnedLiquidityPercentage float64                                               `json:"burned_liquidity_percentage"`
	HoldersCount              int64                                                 `json:"holders_count"`
	LockedLiquidityPercentage float64                                               `json:"locked_liquidity_percentage"`
	TopHolders                []TokenBulkScanResponseResultsFinancialStatsTopHolder `json:"top_holders"`
	UsdPricePerUnit           float64                                               `json:"usd_price_per_unit"`
	JSON                      tokenBulkScanResponseResultsFinancialStatsJSON        `json:"-"`
}

// tokenBulkScanResponseResultsFinancialStatsJSON contains the JSON metadata for
// the struct [TokenBulkScanResponseResultsFinancialStats]
type tokenBulkScanResponseResultsFinancialStatsJSON struct {
	BurnedLiquidityPercentage apijson.Field
	HoldersCount              apijson.Field
	LockedLiquidityPercentage apijson.Field
	TopHolders                apijson.Field
	UsdPricePerUnit           apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *TokenBulkScanResponseResultsFinancialStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsFinancialStatsJSON) RawJSON() string {
	return r.raw
}

type TokenBulkScanResponseResultsFinancialStatsTopHolder struct {
	Address           string                                                  `json:"address"`
	HoldingPercentage float64                                                 `json:"holding_percentage"`
	JSON              tokenBulkScanResponseResultsFinancialStatsTopHolderJSON `json:"-"`
}

// tokenBulkScanResponseResultsFinancialStatsTopHolderJSON contains the JSON
// metadata for the struct [TokenBulkScanResponseResultsFinancialStatsTopHolder]
type tokenBulkScanResponseResultsFinancialStatsTopHolderJSON struct {
	Address           apijson.Field
	HoldingPercentage apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TokenBulkScanResponseResultsFinancialStatsTopHolder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsFinancialStatsTopHolderJSON) RawJSON() string {
	return r.raw
}

// Metadata of the token
type TokenBulkScanResponseResultsMetadata struct {
	// The unique ID for the Rune
	ID string `json:"id"`
	// This field can have the runtime type of
	// [TokenBulkScanResponseResultsMetadataSolanaMetadataContractBalance],
	// [TokenBulkScanResponseResultsMetadataEvmMetadataTokenContractBalance].
	ContractBalance interface{} `json:"contract_balance"`
	// Contract deploy date
	CreationTimestamp time.Time `json:"creation_timestamp" format:"date-time"`
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer"`
	// Description of the token
	Description string `json:"description"`
	// This field can have the runtime type of
	// [TokenBulkScanResponseResultsMetadataSolanaMetadataExternalLinks],
	// [TokenBulkScanResponseResultsMetadataEvmMetadataTokenExternalLinks].
	ExternalLinks interface{} `json:"external_links"`
	// The formatted name of the rune, with spacers
	FormattedName string `json:"formatted_name"`
	// Solana token freeze authority account
	FreezeAuthority string `json:"freeze_authority"`
	// URL of the token image
	ImageURL string `json:"image_url"`
	// Solana token mint authority account
	MintAuthority string `json:"mint_authority"`
	// Name of the token
	Name string `json:"name"`
	// The rune's unique sequential number.
	Number int64 `json:"number"`
	// Contract owner address
	Owner string `json:"owner"`
	// This field can have the runtime type of
	// [TokenBulkScanResponseResultsMetadataSolanaMetadataOwnerBalance],
	// [TokenBulkScanResponseResultsMetadataEvmMetadataTokenOwnerBalance].
	OwnerBalance interface{} `json:"owner_balance"`
	// Symbol of the token
	Symbol string `json:"symbol"`
	// Type of the token
	Type string `json:"type"`
	// Solana token update authority account
	UpdateAuthority string                                   `json:"update_authority"`
	JSON            tokenBulkScanResponseResultsMetadataJSON `json:"-"`
	union           TokenBulkScanResponseResultsMetadataUnion
}

// tokenBulkScanResponseResultsMetadataJSON contains the JSON metadata for the
// struct [TokenBulkScanResponseResultsMetadata]
type tokenBulkScanResponseResultsMetadataJSON struct {
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
	ContractBalance TokenBulkScanResponseResultsMetadataSolanaMetadataContractBalance `json:"contract_balance"`
	// Contract deploy date
	CreationTimestamp time.Time `json:"creation_timestamp" format:"date-time"`
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer"`
	// Description of the token
	Description string `json:"description"`
	// social links of the token
	ExternalLinks TokenBulkScanResponseResultsMetadataSolanaMetadataExternalLinks `json:"external_links"`
	// Solana token freeze authority account
	FreezeAuthority string `json:"freeze_authority"`
	// URL of the token image
	ImageURL string `json:"image_url"`
	// Solana token mint authority account
	MintAuthority string `json:"mint_authority"`
	// Name of the token
	Name string `json:"name"`
	// Contract owner address
	Owner string `json:"owner"`
	// Contract owner balance
	OwnerBalance TokenBulkScanResponseResultsMetadataSolanaMetadataOwnerBalance `json:"owner_balance"`
	// Symbol of the token
	Symbol string `json:"symbol"`
	// Type of the token
	Type string `json:"type"`
	// Solana token update authority account
	UpdateAuthority string                                                 `json:"update_authority"`
	JSON            tokenBulkScanResponseResultsMetadataSolanaMetadataJSON `json:"-"`
}

// tokenBulkScanResponseResultsMetadataSolanaMetadataJSON contains the JSON
// metadata for the struct [TokenBulkScanResponseResultsMetadataSolanaMetadata]
type tokenBulkScanResponseResultsMetadataSolanaMetadataJSON struct {
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
	Amount    float64                                                               `json:"amount"`
	AmountWei string                                                                `json:"amount_wei"`
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

// social links of the token
type TokenBulkScanResponseResultsMetadataSolanaMetadataExternalLinks struct {
	Homepage          string                                                              `json:"homepage"`
	TelegramChannelID string                                                              `json:"telegram_channel_id"`
	TwitterPage       string                                                              `json:"twitter_page"`
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

// Contract owner balance
type TokenBulkScanResponseResultsMetadataSolanaMetadataOwnerBalance struct {
	Amount    float64                                                            `json:"amount"`
	AmountWei string                                                             `json:"amount_wei"`
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
	ID string `json:"id"`
	// The formatted name of the rune, with spacers
	FormattedName string `json:"formatted_name"`
	// Name of the token
	Name string `json:"name"`
	// The rune's unique sequential number.
	Number int64 `json:"number"`
	// Symbol of the token
	Symbol string `json:"symbol"`
	// Type of the token
	Type string                                                       `json:"type"`
	JSON tokenBulkScanResponseResultsMetadataBitcoinMetadataTokenJSON `json:"-"`
}

// tokenBulkScanResponseResultsMetadataBitcoinMetadataTokenJSON contains the JSON
// metadata for the struct
// [TokenBulkScanResponseResultsMetadataBitcoinMetadataToken]
type tokenBulkScanResponseResultsMetadataBitcoinMetadataTokenJSON struct {
	ID            apijson.Field
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
	ContractBalance TokenBulkScanResponseResultsMetadataEvmMetadataTokenContractBalance `json:"contract_balance"`
	// Contract deploy date
	CreationTimestamp time.Time `json:"creation_timestamp" format:"date-time"`
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer"`
	// Description of the token
	Description string `json:"description"`
	// social links of the token
	ExternalLinks TokenBulkScanResponseResultsMetadataEvmMetadataTokenExternalLinks `json:"external_links"`
	// URL of the token image
	ImageURL string `json:"image_url"`
	// Name of the token
	Name string `json:"name"`
	// Contract owner address
	Owner string `json:"owner"`
	// Contract owner balance
	OwnerBalance TokenBulkScanResponseResultsMetadataEvmMetadataTokenOwnerBalance `json:"owner_balance"`
	// Symbol of the token
	Symbol string `json:"symbol"`
	// Type of the token
	Type string                                                   `json:"type"`
	JSON tokenBulkScanResponseResultsMetadataEvmMetadataTokenJSON `json:"-"`
}

// tokenBulkScanResponseResultsMetadataEvmMetadataTokenJSON contains the JSON
// metadata for the struct [TokenBulkScanResponseResultsMetadataEvmMetadataToken]
type tokenBulkScanResponseResultsMetadataEvmMetadataTokenJSON struct {
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
	Amount    float64                                                                 `json:"amount"`
	AmountWei string                                                                  `json:"amount_wei"`
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

// social links of the token
type TokenBulkScanResponseResultsMetadataEvmMetadataTokenExternalLinks struct {
	Homepage          string                                                                `json:"homepage"`
	TelegramChannelID string                                                                `json:"telegram_channel_id"`
	TwitterPage       string                                                                `json:"twitter_page"`
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

// Contract owner balance
type TokenBulkScanResponseResultsMetadataEvmMetadataTokenOwnerBalance struct {
	Amount    float64                                                              `json:"amount"`
	AmountWei string                                                               `json:"amount_wei"`
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

// An enumeration.
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
	MaxBuy            TokenBulkScanResponseResultsTradingLimitsMaxBuy            `json:"max_buy"`
	MaxHolding        TokenBulkScanResponseResultsTradingLimitsMaxHolding        `json:"max_holding"`
	MaxSell           TokenBulkScanResponseResultsTradingLimitsMaxSell           `json:"max_sell"`
	SellLimitPerBlock TokenBulkScanResponseResultsTradingLimitsSellLimitPerBlock `json:"sell_limit_per_block"`
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

type TokenBulkScanResponseResultsTradingLimitsMaxBuy struct {
	Amount    float64                                             `json:"amount"`
	AmountWei string                                              `json:"amount_wei"`
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

type TokenBulkScanResponseResultsTradingLimitsMaxHolding struct {
	Amount    float64                                                 `json:"amount"`
	AmountWei string                                                  `json:"amount_wei"`
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

type TokenBulkScanResponseResultsTradingLimitsMaxSell struct {
	Amount    float64                                              `json:"amount"`
	AmountWei string                                               `json:"amount_wei"`
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

type TokenBulkScanResponseResultsTradingLimitsSellLimitPerBlock struct {
	Amount    float64                                                        `json:"amount"`
	AmountWei string                                                         `json:"amount_wei"`
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
	Description string `json:"description,required"`
	// An enumeration.
	FeatureID TokenBulkScanResponseResultsFeaturesFeatureID `json:"feature_id,required"`
	// An enumeration.
	Type TokenBulkScanResponseResultsFeaturesType `json:"type,required"`
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

// An enumeration.
type TokenBulkScanResponseResultsFeaturesFeatureID string

const (
	TokenBulkScanResponseResultsFeaturesFeatureIDVerifiedContract         TokenBulkScanResponseResultsFeaturesFeatureID = "VERIFIED_CONTRACT"
	TokenBulkScanResponseResultsFeaturesFeatureIDHighTradeVolume          TokenBulkScanResponseResultsFeaturesFeatureID = "HIGH_TRADE_VOLUME"
	TokenBulkScanResponseResultsFeaturesFeatureIDMarketPlaceSalesHistory  TokenBulkScanResponseResultsFeaturesFeatureID = "MARKET_PLACE_SALES_HISTORY"
	TokenBulkScanResponseResultsFeaturesFeatureIDHighReputationToken      TokenBulkScanResponseResultsFeaturesFeatureID = "HIGH_REPUTATION_TOKEN"
	TokenBulkScanResponseResultsFeaturesFeatureIDOnchainActivityValidator TokenBulkScanResponseResultsFeaturesFeatureID = "ONCHAIN_ACTIVITY_VALIDATOR"
	TokenBulkScanResponseResultsFeaturesFeatureIDStaticCodeSignature      TokenBulkScanResponseResultsFeaturesFeatureID = "STATIC_CODE_SIGNATURE"
	TokenBulkScanResponseResultsFeaturesFeatureIDKnownMalicious           TokenBulkScanResponseResultsFeaturesFeatureID = "KNOWN_MALICIOUS"
	TokenBulkScanResponseResultsFeaturesFeatureIDMetadata                 TokenBulkScanResponseResultsFeaturesFeatureID = "METADATA"
	TokenBulkScanResponseResultsFeaturesFeatureIDAirdropPattern           TokenBulkScanResponseResultsFeaturesFeatureID = "AIRDROP_PATTERN"
	TokenBulkScanResponseResultsFeaturesFeatureIDImpersonator             TokenBulkScanResponseResultsFeaturesFeatureID = "IMPERSONATOR"
	TokenBulkScanResponseResultsFeaturesFeatureIDInorganicVolume          TokenBulkScanResponseResultsFeaturesFeatureID = "INORGANIC_VOLUME"
	TokenBulkScanResponseResultsFeaturesFeatureIDDynamicAnalysis          TokenBulkScanResponseResultsFeaturesFeatureID = "DYNAMIC_ANALYSIS"
	TokenBulkScanResponseResultsFeaturesFeatureIDUnstableTokenPrice       TokenBulkScanResponseResultsFeaturesFeatureID = "UNSTABLE_TOKEN_PRICE"
	TokenBulkScanResponseResultsFeaturesFeatureIDRugpull                  TokenBulkScanResponseResultsFeaturesFeatureID = "RUGPULL"
	TokenBulkScanResponseResultsFeaturesFeatureIDConsumerOverride         TokenBulkScanResponseResultsFeaturesFeatureID = "CONSUMER_OVERRIDE"
	TokenBulkScanResponseResultsFeaturesFeatureIDInappropriateContent     TokenBulkScanResponseResultsFeaturesFeatureID = "INAPPROPRIATE_CONTENT"
	TokenBulkScanResponseResultsFeaturesFeatureIDHighTransferFee          TokenBulkScanResponseResultsFeaturesFeatureID = "HIGH_TRANSFER_FEE"
	TokenBulkScanResponseResultsFeaturesFeatureIDHighBuyFee               TokenBulkScanResponseResultsFeaturesFeatureID = "HIGH_BUY_FEE"
	TokenBulkScanResponseResultsFeaturesFeatureIDHighSellFee              TokenBulkScanResponseResultsFeaturesFeatureID = "HIGH_SELL_FEE"
	TokenBulkScanResponseResultsFeaturesFeatureIDIsMintable               TokenBulkScanResponseResultsFeaturesFeatureID = "IS_MINTABLE"
	TokenBulkScanResponseResultsFeaturesFeatureIDModifiableTaxes          TokenBulkScanResponseResultsFeaturesFeatureID = "MODIFIABLE_TAXES"
	TokenBulkScanResponseResultsFeaturesFeatureIDCanBlacklist             TokenBulkScanResponseResultsFeaturesFeatureID = "CAN_BLACKLIST"
	TokenBulkScanResponseResultsFeaturesFeatureIDCanWhitelist             TokenBulkScanResponseResultsFeaturesFeatureID = "CAN_WHITELIST"
	TokenBulkScanResponseResultsFeaturesFeatureIDHasTradingCooldown       TokenBulkScanResponseResultsFeaturesFeatureID = "HAS_TRADING_COOLDOWN"
	TokenBulkScanResponseResultsFeaturesFeatureIDExternalFunctions        TokenBulkScanResponseResultsFeaturesFeatureID = "EXTERNAL_FUNCTIONS"
	TokenBulkScanResponseResultsFeaturesFeatureIDHiddenOwner              TokenBulkScanResponseResultsFeaturesFeatureID = "HIDDEN_OWNER"
	TokenBulkScanResponseResultsFeaturesFeatureIDTransferPauseable        TokenBulkScanResponseResultsFeaturesFeatureID = "TRANSFER_PAUSEABLE"
	TokenBulkScanResponseResultsFeaturesFeatureIDOwnershipRenounced       TokenBulkScanResponseResultsFeaturesFeatureID = "OWNERSHIP_RENOUNCED"
	TokenBulkScanResponseResultsFeaturesFeatureIDProxyContract            TokenBulkScanResponseResultsFeaturesFeatureID = "PROXY_CONTRACT"
)

func (r TokenBulkScanResponseResultsFeaturesFeatureID) IsKnown() bool {
	switch r {
	case TokenBulkScanResponseResultsFeaturesFeatureIDVerifiedContract, TokenBulkScanResponseResultsFeaturesFeatureIDHighTradeVolume, TokenBulkScanResponseResultsFeaturesFeatureIDMarketPlaceSalesHistory, TokenBulkScanResponseResultsFeaturesFeatureIDHighReputationToken, TokenBulkScanResponseResultsFeaturesFeatureIDOnchainActivityValidator, TokenBulkScanResponseResultsFeaturesFeatureIDStaticCodeSignature, TokenBulkScanResponseResultsFeaturesFeatureIDKnownMalicious, TokenBulkScanResponseResultsFeaturesFeatureIDMetadata, TokenBulkScanResponseResultsFeaturesFeatureIDAirdropPattern, TokenBulkScanResponseResultsFeaturesFeatureIDImpersonator, TokenBulkScanResponseResultsFeaturesFeatureIDInorganicVolume, TokenBulkScanResponseResultsFeaturesFeatureIDDynamicAnalysis, TokenBulkScanResponseResultsFeaturesFeatureIDUnstableTokenPrice, TokenBulkScanResponseResultsFeaturesFeatureIDRugpull, TokenBulkScanResponseResultsFeaturesFeatureIDConsumerOverride, TokenBulkScanResponseResultsFeaturesFeatureIDInappropriateContent, TokenBulkScanResponseResultsFeaturesFeatureIDHighTransferFee, TokenBulkScanResponseResultsFeaturesFeatureIDHighBuyFee, TokenBulkScanResponseResultsFeaturesFeatureIDHighSellFee, TokenBulkScanResponseResultsFeaturesFeatureIDIsMintable, TokenBulkScanResponseResultsFeaturesFeatureIDModifiableTaxes, TokenBulkScanResponseResultsFeaturesFeatureIDCanBlacklist, TokenBulkScanResponseResultsFeaturesFeatureIDCanWhitelist, TokenBulkScanResponseResultsFeaturesFeatureIDHasTradingCooldown, TokenBulkScanResponseResultsFeaturesFeatureIDExternalFunctions, TokenBulkScanResponseResultsFeaturesFeatureIDHiddenOwner, TokenBulkScanResponseResultsFeaturesFeatureIDTransferPauseable, TokenBulkScanResponseResultsFeaturesFeatureIDOwnershipRenounced, TokenBulkScanResponseResultsFeaturesFeatureIDProxyContract:
		return true
	}
	return false
}

// An enumeration.
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

type TokenBulkScanParams struct {
	// The chain name
	Chain param.Field[TokenScanSupportedChain] `json:"chain,required"`
	// A list of token addresses to scan
	Tokens param.Field[[]string] `json:"tokens,required"`
	// Object of additional information to validate against.
	Metadata param.Field[TokenBulkScanParamsMetadata] `json:"metadata"`
}

func (r TokenBulkScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Object of additional information to validate against.
type TokenBulkScanParamsMetadata struct {
	// cross reference transaction against the domain.
	Domain param.Field[string] `json:"domain"`
}

func (r TokenBulkScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
