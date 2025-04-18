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
	// Blockchain network
	Chain TokenScanSupportedChain `json:"chain,required"`
	// Fees associated with the token
	Fees TokenBulkScanResponseResultsFees `json:"fees,required"`
	// financial stats of the token
	FinancialStats TokenBulkScanResponseResultsFinancialStats `json:"financial_stats,required"`
	// Score between 0 to 1 (double)
	MaliciousScore string `json:"malicious_score,required"`
	// Metadata of the token
	Metadata TokenBulkScanResponseResultsMetadata `json:"metadata,required"`
	// General indication
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
	Buy float64 `json:"buy,nullable"`
	// Sell fee of the token
	Sell float64 `json:"sell,nullable"`
	// Transfer fee of the token
	Transfer float64 `json:"transfer,nullable"`
	// The maximum value that a transfer fee will cost
	TransferFeeMaxAmount int64                                `json:"transfer_fee_max_amount,nullable"`
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

// financial stats of the token
type TokenBulkScanResponseResultsFinancialStats struct {
	// Token liquidity burned percentage
	BurnedLiquidityPercentage float64 `json:"burned_liquidity_percentage,nullable"`
	// Amount of token holders
	HoldersCount int64 `json:"holders_count,nullable"`
	// Token liquidity locked percentage
	LockedLiquidityPercentage float64 `json:"locked_liquidity_percentage,nullable"`
	// token supply
	Supply int64 `json:"supply,nullable"`
	// Top token holders
	TopHolders []TokenBulkScanResponseResultsFinancialStatsTopHolder `json:"top_holders"`
	// Total reserve in USD
	TotalReserveInUsd float64 `json:"total_reserve_in_usd,nullable"`
	// token price in USD
	UsdPricePerUnit float64                                        `json:"usd_price_per_unit,nullable"`
	JSON            tokenBulkScanResponseResultsFinancialStatsJSON `json:"-"`
}

// tokenBulkScanResponseResultsFinancialStatsJSON contains the JSON metadata for
// the struct [TokenBulkScanResponseResultsFinancialStats]
type tokenBulkScanResponseResultsFinancialStatsJSON struct {
	BurnedLiquidityPercentage apijson.Field
	HoldersCount              apijson.Field
	LockedLiquidityPercentage apijson.Field
	Supply                    apijson.Field
	TopHolders                apijson.Field
	TotalReserveInUsd         apijson.Field
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
	// Address
	Address string `json:"address,nullable"`
	// Holding position out of total token liquidity
	HoldingPercentage float64                                                 `json:"holding_percentage,nullable"`
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
	ID string `json:"id,nullable"`
	// This field can have the runtime type of
	// [TokenBulkScanResponseResultsMetadataSolanaMetadataContractBalance],
	// [TokenBulkScanResponseResultsMetadataEvmMetadataTokenContractBalance].
	ContractBalance interface{} `json:"contract_balance"`
	// Contract deploy date
	CreationTimestamp string `json:"creation_timestamp,nullable"`
	// Decimals of the token
	Decimals int64 `json:"decimals,nullable"`
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer,nullable"`
	// This field can have the runtime type of
	// [TokenBulkScanResponseResultsMetadataSolanaMetadataDeployerBalance],
	// [TokenBulkScanResponseResultsMetadataEvmMetadataTokenDeployerBalance].
	DeployerBalance interface{} `json:"deployer_balance"`
	// Description of the token
	Description string `json:"description,nullable"`
	// This field can have the runtime type of
	// [TokenBulkScanResponseResultsMetadataSolanaMetadataExternalLinks],
	// [TokenBulkScanResponseResultsMetadataEvmMetadataTokenExternalLinks].
	ExternalLinks interface{} `json:"external_links"`
	// The formatted name of the rune, with spacers
	FormattedName string `json:"formatted_name,nullable"`
	// Solana token freeze authority account
	FreezeAuthority string `json:"freeze_authority,nullable"`
	// URL of the token image
	ImageURL string `json:"image_url,nullable"`
	// This field can have the runtime type of [[]string].
	MaliciousURLs interface{} `json:"malicious_urls"`
	// Solana token mint authority account
	MintAuthority string `json:"mint_authority,nullable"`
	// Name of the token
	Name string `json:"name,nullable"`
	// The rune's unique sequential number.
	Number int64 `json:"number,nullable"`
	// Contract owner address
	Owner string `json:"owner,nullable"`
	// This field can have the runtime type of
	// [TokenBulkScanResponseResultsMetadataSolanaMetadataOwnerBalance],
	// [TokenBulkScanResponseResultsMetadataEvmMetadataTokenOwnerBalance].
	OwnerBalance interface{} `json:"owner_balance"`
	// Symbol of the token
	Symbol string `json:"symbol,nullable"`
	// Address of the token creation initiator, only set if the tokens was created by a
	// well known token launch platform
	TokenCreationInitiator string `json:"token_creation_initiator,nullable"`
	// Type of the token
	Type string `json:"type,nullable"`
	// Solana token update authority account
	UpdateAuthority string `json:"update_authority,nullable"`
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
	MaliciousURLs          apijson.Field
	MintAuthority          apijson.Field
	Name                   apijson.Field
	Number                 apijson.Field
	Owner                  apijson.Field
	OwnerBalance           apijson.Field
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
	ContractBalance TokenBulkScanResponseResultsMetadataSolanaMetadataContractBalance `json:"contract_balance,nullable"`
	// Contract deploy date
	CreationTimestamp string `json:"creation_timestamp,nullable"`
	// Decimals of the token
	Decimals int64 `json:"decimals,nullable"`
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer,nullable"`
	// Contract creator balance
	DeployerBalance TokenBulkScanResponseResultsMetadataSolanaMetadataDeployerBalance `json:"deployer_balance,nullable"`
	// Description of the token
	Description string `json:"description,nullable"`
	// social links of the token
	ExternalLinks TokenBulkScanResponseResultsMetadataSolanaMetadataExternalLinks `json:"external_links"`
	// Solana token freeze authority account
	FreezeAuthority string `json:"freeze_authority,nullable"`
	// URL of the token image
	ImageURL string `json:"image_url,nullable"`
	// Malicious urls associated with the token
	MaliciousURLs []string `json:"malicious_urls,nullable"`
	// Solana token mint authority account
	MintAuthority string `json:"mint_authority,nullable"`
	// Name of the token
	Name string `json:"name,nullable"`
	// Contract owner address
	Owner string `json:"owner,nullable"`
	// Contract owner balance
	OwnerBalance TokenBulkScanResponseResultsMetadataSolanaMetadataOwnerBalance `json:"owner_balance,nullable"`
	// Symbol of the token
	Symbol string `json:"symbol,nullable"`
	// Address of the token creation initiator, only set if the tokens was created by a
	// well known token launch platform
	TokenCreationInitiator string `json:"token_creation_initiator,nullable"`
	// Type of the token
	Type string `json:"type,nullable"`
	// Solana token update authority account
	UpdateAuthority string `json:"update_authority,nullable"`
	// Urls associated with the token
	URLs []string                                               `json:"urls,nullable"`
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
	MaliciousURLs          apijson.Field
	MintAuthority          apijson.Field
	Name                   apijson.Field
	Owner                  apijson.Field
	OwnerBalance           apijson.Field
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
	Amount    float64                                                               `json:"amount,nullable"`
	AmountWei string                                                                `json:"amount_wei,nullable"`
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
	Amount    float64                                                               `json:"amount,nullable"`
	AmountWei string                                                                `json:"amount_wei,nullable"`
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
	Homepage          string                                                              `json:"homepage,nullable"`
	TelegramChannelID string                                                              `json:"telegram_channel_id,nullable"`
	TwitterPage       string                                                              `json:"twitter_page,nullable"`
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
	Amount    float64                                                            `json:"amount,nullable"`
	AmountWei string                                                             `json:"amount_wei,nullable"`
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
	ID string `json:"id,nullable"`
	// Decimals of the token
	Decimals int64 `json:"decimals,nullable"`
	// The formatted name of the rune, with spacers
	FormattedName string `json:"formatted_name,nullable"`
	// Name of the token
	Name string `json:"name,nullable"`
	// The rune's unique sequential number.
	Number int64 `json:"number,nullable"`
	// Symbol of the token
	Symbol string `json:"symbol,nullable"`
	// Type of the token
	Type string                                                       `json:"type,nullable"`
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
	ContractBalance TokenBulkScanResponseResultsMetadataEvmMetadataTokenContractBalance `json:"contract_balance,nullable"`
	// Contract deploy date
	CreationTimestamp string `json:"creation_timestamp,nullable"`
	// Decimals of the token
	Decimals int64 `json:"decimals,nullable"`
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer,nullable"`
	// Contract creator balance
	DeployerBalance TokenBulkScanResponseResultsMetadataEvmMetadataTokenDeployerBalance `json:"deployer_balance,nullable"`
	// Description of the token
	Description string `json:"description,nullable"`
	// social links of the token
	ExternalLinks TokenBulkScanResponseResultsMetadataEvmMetadataTokenExternalLinks `json:"external_links"`
	// URL of the token image
	ImageURL string `json:"image_url,nullable"`
	// Malicious urls associated with the token
	MaliciousURLs []string `json:"malicious_urls,nullable"`
	// Name of the token
	Name string `json:"name,nullable"`
	// Contract owner address
	Owner string `json:"owner,nullable"`
	// Contract owner balance
	OwnerBalance TokenBulkScanResponseResultsMetadataEvmMetadataTokenOwnerBalance `json:"owner_balance,nullable"`
	// Symbol of the token
	Symbol string `json:"symbol,nullable"`
	// Address of the token creation initiator, only set if the tokens was created by a
	// well known token launch platform
	TokenCreationInitiator string `json:"token_creation_initiator,nullable"`
	// Type of the token
	Type string `json:"type,nullable"`
	// Urls associated with the token
	URLs []string                                                 `json:"urls,nullable"`
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
	Amount    float64                                                                 `json:"amount,nullable"`
	AmountWei string                                                                  `json:"amount_wei,nullable"`
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
	Amount    float64                                                                 `json:"amount,nullable"`
	AmountWei string                                                                  `json:"amount_wei,nullable"`
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
	Homepage          string                                                                `json:"homepage,nullable"`
	TelegramChannelID string                                                                `json:"telegram_channel_id,nullable"`
	TwitterPage       string                                                                `json:"twitter_page,nullable"`
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
	Amount    float64                                                              `json:"amount,nullable"`
	AmountWei string                                                               `json:"amount_wei,nullable"`
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
	MaxBuy TokenBulkScanResponseResultsTradingLimitsMaxBuy `json:"max_buy,nullable"`
	// Max amount that can be held by a single address
	MaxHolding TokenBulkScanResponseResultsTradingLimitsMaxHolding `json:"max_holding,nullable"`
	// Max amount that can be sold at once
	MaxSell TokenBulkScanResponseResultsTradingLimitsMaxSell `json:"max_sell,nullable"`
	// Maximum amount of the token that can be sold in a block
	SellLimitPerBlock TokenBulkScanResponseResultsTradingLimitsSellLimitPerBlock `json:"sell_limit_per_block,nullable"`
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
	Amount    float64                                             `json:"amount,nullable"`
	AmountWei string                                              `json:"amount_wei,nullable"`
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
	Amount    float64                                                 `json:"amount,nullable"`
	AmountWei string                                                  `json:"amount_wei,nullable"`
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
	Amount    float64                                              `json:"amount,nullable"`
	AmountWei string                                               `json:"amount_wei,nullable"`
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
	Amount    float64                                                        `json:"amount,nullable"`
	AmountWei string                                                         `json:"amount_wei,nullable"`
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
	// Feature identifier
	FeatureID TokenBulkScanResponseResultsFeaturesFeatureID `json:"feature_id,required"`
	// Type of the feature
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
	TokenBulkScanResponseResultsFeaturesFeatureIDFakeVolume                     TokenBulkScanResponseResultsFeaturesFeatureID = "FAKE_VOLUME"
	TokenBulkScanResponseResultsFeaturesFeatureIDHiddenSupplyByKeyHolder        TokenBulkScanResponseResultsFeaturesFeatureID = "HIDDEN_SUPPLY_BY_KEY_HOLDER"
	TokenBulkScanResponseResultsFeaturesFeatureIDFakeTradeMakerCount            TokenBulkScanResponseResultsFeaturesFeatureID = "FAKE_TRADE_MAKER_COUNT"
)

func (r TokenBulkScanResponseResultsFeaturesFeatureID) IsKnown() bool {
	switch r {
	case TokenBulkScanResponseResultsFeaturesFeatureIDVerifiedContract, TokenBulkScanResponseResultsFeaturesFeatureIDUnverifiedContract, TokenBulkScanResponseResultsFeaturesFeatureIDHighTradeVolume, TokenBulkScanResponseResultsFeaturesFeatureIDMarketPlaceSalesHistory, TokenBulkScanResponseResultsFeaturesFeatureIDHighReputationToken, TokenBulkScanResponseResultsFeaturesFeatureIDOnchainActivityValidator, TokenBulkScanResponseResultsFeaturesFeatureIDStaticCodeSignature, TokenBulkScanResponseResultsFeaturesFeatureIDKnownMalicious, TokenBulkScanResponseResultsFeaturesFeatureIDMetadata, TokenBulkScanResponseResultsFeaturesFeatureIDAirdropPattern, TokenBulkScanResponseResultsFeaturesFeatureIDImpersonator, TokenBulkScanResponseResultsFeaturesFeatureIDInorganicVolume, TokenBulkScanResponseResultsFeaturesFeatureIDDynamicAnalysis, TokenBulkScanResponseResultsFeaturesFeatureIDConcentratedSupplyDistribution, TokenBulkScanResponseResultsFeaturesFeatureIDHoneypot, TokenBulkScanResponseResultsFeaturesFeatureIDInsufficientLockedLiquidity, TokenBulkScanResponseResultsFeaturesFeatureIDUnstableTokenPrice, TokenBulkScanResponseResultsFeaturesFeatureIDRugpull, TokenBulkScanResponseResultsFeaturesFeatureIDWashTrading, TokenBulkScanResponseResultsFeaturesFeatureIDConsumerOverride, TokenBulkScanResponseResultsFeaturesFeatureIDInappropriateContent, TokenBulkScanResponseResultsFeaturesFeatureIDHighTransferFee, TokenBulkScanResponseResultsFeaturesFeatureIDHighBuyFee, TokenBulkScanResponseResultsFeaturesFeatureIDHighSellFee, TokenBulkScanResponseResultsFeaturesFeatureIDUnsellableToken, TokenBulkScanResponseResultsFeaturesFeatureIDIsMintable, TokenBulkScanResponseResultsFeaturesFeatureIDRebaseToken, TokenBulkScanResponseResultsFeaturesFeatureIDLiquidStakingToken, TokenBulkScanResponseResultsFeaturesFeatureIDModifiableTaxes, TokenBulkScanResponseResultsFeaturesFeatureIDCanBlacklist, TokenBulkScanResponseResultsFeaturesFeatureIDCanWhitelist, TokenBulkScanResponseResultsFeaturesFeatureIDHasTradingCooldown, TokenBulkScanResponseResultsFeaturesFeatureIDExternalFunctions, TokenBulkScanResponseResultsFeaturesFeatureIDHiddenOwner, TokenBulkScanResponseResultsFeaturesFeatureIDTransferPauseable, TokenBulkScanResponseResultsFeaturesFeatureIDOwnershipRenounced, TokenBulkScanResponseResultsFeaturesFeatureIDOwnerCanChangeBalance, TokenBulkScanResponseResultsFeaturesFeatureIDProxyContract, TokenBulkScanResponseResultsFeaturesFeatureIDSimilarMaliciousContract, TokenBulkScanResponseResultsFeaturesFeatureIDFakeVolume, TokenBulkScanResponseResultsFeaturesFeatureIDHiddenSupplyByKeyHolder, TokenBulkScanResponseResultsFeaturesFeatureIDFakeTradeMakerCount:
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
