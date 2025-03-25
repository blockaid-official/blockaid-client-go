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
// behavior
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
	Chain TokenScanSupportedChain `json:"chain,required"`
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

// Fees associated with the token
type TokenScanResponseFees struct {
	// Buy fee of the token
	Buy float64 `json:"buy,nullable"`
	// Sell fee of the token
	Sell float64 `json:"sell,nullable"`
	// Transfer fee of the token
	Transfer float64 `json:"transfer,nullable"`
	// The maximum value that a transfer fee will cost
	TransferFeeMaxAmount int64                     `json:"transfer_fee_max_amount,nullable"`
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

// financial stats of the token
type TokenScanResponseFinancialStats struct {
	// Token liquidity burned percentage
	BurnedLiquidityPercentage float64 `json:"burned_liquidity_percentage,nullable"`
	// Amount of token holders
	HoldersCount int64 `json:"holders_count,nullable"`
	// Token liquidity locked percentage
	LockedLiquidityPercentage float64 `json:"locked_liquidity_percentage,nullable"`
	// token supply
	Supply int64 `json:"supply,nullable"`
	// Top token holders
	TopHolders []TokenScanResponseFinancialStatsTopHolder `json:"top_holders"`
	// Total reserve in USD
	TotalReserveInUsd float64 `json:"total_reserve_in_usd,nullable"`
	// token price in USD
	UsdPricePerUnit float64                             `json:"usd_price_per_unit,nullable"`
	JSON            tokenScanResponseFinancialStatsJSON `json:"-"`
}

// tokenScanResponseFinancialStatsJSON contains the JSON metadata for the struct
// [TokenScanResponseFinancialStats]
type tokenScanResponseFinancialStatsJSON struct {
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

func (r *TokenScanResponseFinancialStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseFinancialStatsJSON) RawJSON() string {
	return r.raw
}

type TokenScanResponseFinancialStatsTopHolder struct {
	// Address
	Address string `json:"address,nullable"`
	// Holding position out of total token liquidity
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
	// [TokenScanResponseMetadataSolanaMetadataContractBalance],
	// [TokenScanResponseMetadataEvmMetadataTokenContractBalance].
	ContractBalance interface{} `json:"contract_balance"`
	// Contract deploy date
	CreationTimestamp string `json:"creation_timestamp,nullable"`
	// Decimals of the token
	Decimals int64 `json:"decimals,nullable"`
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer,nullable"`
	// This field can have the runtime type of
	// [TokenScanResponseMetadataSolanaMetadataDeployerBalance],
	// [TokenScanResponseMetadataEvmMetadataTokenDeployerBalance].
	DeployerBalance interface{} `json:"deployer_balance"`
	// Description of the token
	Description string `json:"description,nullable"`
	// This field can have the runtime type of
	// [TokenScanResponseMetadataSolanaMetadataExternalLinks],
	// [TokenScanResponseMetadataEvmMetadataTokenExternalLinks].
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
	// [TokenScanResponseMetadataSolanaMetadataOwnerBalance],
	// [TokenScanResponseMetadataEvmMetadataTokenOwnerBalance].
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
	ContractBalance TokenScanResponseMetadataSolanaMetadataContractBalance `json:"contract_balance,nullable"`
	// Contract deploy date
	CreationTimestamp string `json:"creation_timestamp,nullable"`
	// Decimals of the token
	Decimals int64 `json:"decimals,nullable"`
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer,nullable"`
	// Contract creator balance
	DeployerBalance TokenScanResponseMetadataSolanaMetadataDeployerBalance `json:"deployer_balance,nullable"`
	// Description of the token
	Description string `json:"description,nullable"`
	// social links of the token
	ExternalLinks TokenScanResponseMetadataSolanaMetadataExternalLinks `json:"external_links"`
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
	OwnerBalance TokenScanResponseMetadataSolanaMetadataOwnerBalance `json:"owner_balance,nullable"`
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
	URLs []string                                    `json:"urls,nullable"`
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

func (r *TokenScanResponseMetadataSolanaMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseMetadataSolanaMetadataJSON) RawJSON() string {
	return r.raw
}

func (r TokenScanResponseMetadataSolanaMetadata) implementsTokenScanResponseMetadata() {}

// Contract balance
type TokenScanResponseMetadataSolanaMetadataContractBalance struct {
	Amount    float64                                                    `json:"amount,nullable"`
	AmountWei string                                                     `json:"amount_wei,nullable"`
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
	Amount    float64                                                    `json:"amount,nullable"`
	AmountWei string                                                     `json:"amount_wei,nullable"`
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
	Homepage          string                                                   `json:"homepage,nullable"`
	TelegramChannelID string                                                   `json:"telegram_channel_id,nullable"`
	TwitterPage       string                                                   `json:"twitter_page,nullable"`
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

// Contract owner balance
type TokenScanResponseMetadataSolanaMetadataOwnerBalance struct {
	Amount    float64                                                 `json:"amount,nullable"`
	AmountWei string                                                  `json:"amount_wei,nullable"`
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
	Type string                                            `json:"type,nullable"`
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
	ContractBalance TokenScanResponseMetadataEvmMetadataTokenContractBalance `json:"contract_balance,nullable"`
	// Contract deploy date
	CreationTimestamp string `json:"creation_timestamp,nullable"`
	// Decimals of the token
	Decimals int64 `json:"decimals,nullable"`
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer,nullable"`
	// Contract creator balance
	DeployerBalance TokenScanResponseMetadataEvmMetadataTokenDeployerBalance `json:"deployer_balance,nullable"`
	// Description of the token
	Description string `json:"description,nullable"`
	// social links of the token
	ExternalLinks TokenScanResponseMetadataEvmMetadataTokenExternalLinks `json:"external_links"`
	// URL of the token image
	ImageURL string `json:"image_url,nullable"`
	// Malicious urls associated with the token
	MaliciousURLs []string `json:"malicious_urls,nullable"`
	// Name of the token
	Name string `json:"name,nullable"`
	// Contract owner address
	Owner string `json:"owner,nullable"`
	// Contract owner balance
	OwnerBalance TokenScanResponseMetadataEvmMetadataTokenOwnerBalance `json:"owner_balance,nullable"`
	// Symbol of the token
	Symbol string `json:"symbol,nullable"`
	// Address of the token creation initiator, only set if the tokens was created by a
	// well known token launch platform
	TokenCreationInitiator string `json:"token_creation_initiator,nullable"`
	// Type of the token
	Type string `json:"type,nullable"`
	// Urls associated with the token
	URLs []string                                      `json:"urls,nullable"`
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
	Amount    float64                                                      `json:"amount,nullable"`
	AmountWei string                                                       `json:"amount_wei,nullable"`
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
	Amount    float64                                                      `json:"amount,nullable"`
	AmountWei string                                                       `json:"amount_wei,nullable"`
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
	Homepage          string                                                     `json:"homepage,nullable"`
	TelegramChannelID string                                                     `json:"telegram_channel_id,nullable"`
	TwitterPage       string                                                     `json:"twitter_page,nullable"`
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

// Contract owner balance
type TokenScanResponseMetadataEvmMetadataTokenOwnerBalance struct {
	Amount    float64                                                   `json:"amount,nullable"`
	AmountWei string                                                    `json:"amount_wei,nullable"`
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
	MaxBuy TokenScanResponseTradingLimitsMaxBuy `json:"max_buy,nullable"`
	// Max amount that can be held by a single address
	MaxHolding TokenScanResponseTradingLimitsMaxHolding `json:"max_holding,nullable"`
	// Max amount that can be sold at once
	MaxSell TokenScanResponseTradingLimitsMaxSell `json:"max_sell,nullable"`
	// Maximum amount of the token that can be sold in a block
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

// Max amount that can be bought at once
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

// Max amount that can be held by a single address
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

// Max amount that can be sold at once
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

// Maximum amount of the token that can be sold in a block
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
	TokenScanResponseFeaturesFeatureIDUnverifiedContract             TokenScanResponseFeaturesFeatureID = "UNVERIFIED_CONTRACT"
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
	case TokenScanResponseFeaturesFeatureIDVerifiedContract, TokenScanResponseFeaturesFeatureIDUnverifiedContract, TokenScanResponseFeaturesFeatureIDHighTradeVolume, TokenScanResponseFeaturesFeatureIDMarketPlaceSalesHistory, TokenScanResponseFeaturesFeatureIDHighReputationToken, TokenScanResponseFeaturesFeatureIDOnchainActivityValidator, TokenScanResponseFeaturesFeatureIDStaticCodeSignature, TokenScanResponseFeaturesFeatureIDKnownMalicious, TokenScanResponseFeaturesFeatureIDMetadata, TokenScanResponseFeaturesFeatureIDAirdropPattern, TokenScanResponseFeaturesFeatureIDImpersonator, TokenScanResponseFeaturesFeatureIDInorganicVolume, TokenScanResponseFeaturesFeatureIDDynamicAnalysis, TokenScanResponseFeaturesFeatureIDConcentratedSupplyDistribution, TokenScanResponseFeaturesFeatureIDHoneypot, TokenScanResponseFeaturesFeatureIDInsufficientLockedLiquidity, TokenScanResponseFeaturesFeatureIDUnstableTokenPrice, TokenScanResponseFeaturesFeatureIDRugpull, TokenScanResponseFeaturesFeatureIDWashTrading, TokenScanResponseFeaturesFeatureIDConsumerOverride, TokenScanResponseFeaturesFeatureIDInappropriateContent, TokenScanResponseFeaturesFeatureIDHighTransferFee, TokenScanResponseFeaturesFeatureIDHighBuyFee, TokenScanResponseFeaturesFeatureIDHighSellFee, TokenScanResponseFeaturesFeatureIDUnsellableToken, TokenScanResponseFeaturesFeatureIDIsMintable, TokenScanResponseFeaturesFeatureIDRebaseToken, TokenScanResponseFeaturesFeatureIDLiquidStakingToken, TokenScanResponseFeaturesFeatureIDModifiableTaxes, TokenScanResponseFeaturesFeatureIDCanBlacklist, TokenScanResponseFeaturesFeatureIDCanWhitelist, TokenScanResponseFeaturesFeatureIDHasTradingCooldown, TokenScanResponseFeaturesFeatureIDExternalFunctions, TokenScanResponseFeaturesFeatureIDHiddenOwner, TokenScanResponseFeaturesFeatureIDTransferPauseable, TokenScanResponseFeaturesFeatureIDOwnershipRenounced, TokenScanResponseFeaturesFeatureIDOwnerCanChangeBalance, TokenScanResponseFeaturesFeatureIDProxyContract, TokenScanResponseFeaturesFeatureIDSimilarMaliciousContract, TokenScanResponseFeaturesFeatureIDFakeVolume, TokenScanResponseFeaturesFeatureIDHiddenSupplyByKeyHolder, TokenScanResponseFeaturesFeatureIDFakeTradeMakerCount:
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
// Satisfied by [TokenReportParamsReportParamReportTokenReportParams],
// [TokenReportParamsReportRequestIDReport], [TokenReportParamsReport].
type TokenReportParamsReportUnion interface {
	implementsTokenReportParamsReportUnion()
}

type TokenReportParamsReportParamReportTokenReportParams struct {
	Params param.Field[TokenReportParamsReportParamReportTokenReportParamsParams] `json:"params,required"`
	Type   param.Field[TokenReportParamsReportParamReportTokenReportParamsType]   `json:"type,required"`
}

func (r TokenReportParamsReportParamReportTokenReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TokenReportParamsReportParamReportTokenReportParams) implementsTokenReportParamsReportUnion() {
}

type TokenReportParamsReportParamReportTokenReportParamsParams struct {
	// The address of the token to report on.
	Address param.Field[string] `json:"address,required"`
	// The chain name
	Chain param.Field[TokenScanSupportedChain] `json:"chain,required"`
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
	RequestID param.Field[string]                                     `json:"request_id,required"`
	Type      param.Field[TokenReportParamsReportRequestIDReportType] `json:"type,required"`
}

func (r TokenReportParamsReportRequestIDReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TokenReportParamsReportRequestIDReport) implementsTokenReportParamsReportUnion() {}

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
	Address param.Field[string] `json:"address,required"`
	// The chain name
	Chain param.Field[TokenScanSupportedChain] `json:"chain,required"`
	// Object of additional information to validate against.
	Metadata param.Field[TokenScanParamsMetadata] `json:"metadata"`
}

func (r TokenScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Object of additional information to validate against.
type TokenScanParamsMetadata struct {
	// cross reference transaction against the domain.
	Domain param.Field[string] `json:"domain"`
}

func (r TokenScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
