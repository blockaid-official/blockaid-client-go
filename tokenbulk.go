// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"net/http"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/param"
	"github.com/blockaid-official/blockaid-client-go/internal/requestconfig"
	"github.com/blockaid-official/blockaid-client-go/option"
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
	FreezeAuthority string `json:"freeze_authority,required"`
	MintAuthority   string `json:"mint_authority,required"`
	// internal metadata
	TokenMetadata   TokenBulkScanResponseResultsMetadataTokenMetadata `json:"token_metadata,required"`
	UpdateAuthority string                                            `json:"update_authority,required"`
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer"`
	// Description of the token
	Description string `json:"description"`
	// URL of the token image
	ImageURL string `json:"image_url"`
	// Name of the token
	Name string `json:"name"`
	// Price per unit of the token. For NFT, it's the price of the NFT. For ERC20, it's
	// the price of a single token. Can be updated daily.
	PricePerUnit float64 `json:"price_per_unit"`
	// Symbol of the token
	Symbol string `json:"symbol"`
	// Type of the token
	Type string                                   `json:"type"`
	JSON tokenBulkScanResponseResultsMetadataJSON `json:"-"`
}

// tokenBulkScanResponseResultsMetadataJSON contains the JSON metadata for the
// struct [TokenBulkScanResponseResultsMetadata]
type tokenBulkScanResponseResultsMetadataJSON struct {
	FreezeAuthority apijson.Field
	MintAuthority   apijson.Field
	TokenMetadata   apijson.Field
	UpdateAuthority apijson.Field
	Deployer        apijson.Field
	Description     apijson.Field
	ImageURL        apijson.Field
	Name            apijson.Field
	PricePerUnit    apijson.Field
	Symbol          apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TokenBulkScanResponseResultsMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsMetadataJSON) RawJSON() string {
	return r.raw
}

// internal metadata
type TokenBulkScanResponseResultsMetadataTokenMetadata struct {
	// Address of the deployer of the fungible token
	Deployer string `json:"deployer"`
	// Description of the token
	Description string `json:"description"`
	// Number of owners of the fungible token, updated daily
	HoldersCount int64 `json:"holders_count"`
	// URL of the token image
	ImageURL string `json:"image_url"`
	// List of malicious_urls
	MaliciousURLs []string `json:"malicious_urls"`
	// Name of the token
	Name string `json:"name"`
	// Symbol of the token
	Symbol string `json:"symbol"`
	// An enumeration.
	Type TokenBulkScanResponseResultsMetadataTokenMetadataType `json:"type"`
	// Uri of the token
	Uri string `json:"uri"`
	// List of urls
	URLs []string                                              `json:"urls"`
	JSON tokenBulkScanResponseResultsMetadataTokenMetadataJSON `json:"-"`
}

// tokenBulkScanResponseResultsMetadataTokenMetadataJSON contains the JSON metadata
// for the struct [TokenBulkScanResponseResultsMetadataTokenMetadata]
type tokenBulkScanResponseResultsMetadataTokenMetadataJSON struct {
	Deployer      apijson.Field
	Description   apijson.Field
	HoldersCount  apijson.Field
	ImageURL      apijson.Field
	MaliciousURLs apijson.Field
	Name          apijson.Field
	Symbol        apijson.Field
	Type          apijson.Field
	Uri           apijson.Field
	URLs          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TokenBulkScanResponseResultsMetadataTokenMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsMetadataTokenMetadataJSON) RawJSON() string {
	return r.raw
}

// An enumeration.
type TokenBulkScanResponseResultsMetadataTokenMetadataType string

const (
	TokenBulkScanResponseResultsMetadataTokenMetadataTypeErc20       TokenBulkScanResponseResultsMetadataTokenMetadataType = "erc20"
	TokenBulkScanResponseResultsMetadataTokenMetadataTypeErc721      TokenBulkScanResponseResultsMetadataTokenMetadataType = "erc721"
	TokenBulkScanResponseResultsMetadataTokenMetadataTypeErc1155     TokenBulkScanResponseResultsMetadataTokenMetadataType = "erc1155"
	TokenBulkScanResponseResultsMetadataTokenMetadataTypeFungible    TokenBulkScanResponseResultsMetadataTokenMetadataType = "FUNGIBLE"
	TokenBulkScanResponseResultsMetadataTokenMetadataTypeNonFungible TokenBulkScanResponseResultsMetadataTokenMetadataType = "NonFungible"
)

func (r TokenBulkScanResponseResultsMetadataTokenMetadataType) IsKnown() bool {
	switch r {
	case TokenBulkScanResponseResultsMetadataTokenMetadataTypeErc20, TokenBulkScanResponseResultsMetadataTokenMetadataTypeErc721, TokenBulkScanResponseResultsMetadataTokenMetadataTypeErc1155, TokenBulkScanResponseResultsMetadataTokenMetadataTypeFungible, TokenBulkScanResponseResultsMetadataTokenMetadataTypeNonFungible:
		return true
	}
	return false
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
	TokenBulkScanResponseResultsFeaturesFeatureIDVerifiedContract        TokenBulkScanResponseResultsFeaturesFeatureID = "VERIFIED_CONTRACT"
	TokenBulkScanResponseResultsFeaturesFeatureIDHighTradeVolume         TokenBulkScanResponseResultsFeaturesFeatureID = "HIGH_TRADE_VOLUME"
	TokenBulkScanResponseResultsFeaturesFeatureIDMarketPlaceSalesHistory TokenBulkScanResponseResultsFeaturesFeatureID = "MARKET_PLACE_SALES_HISTORY"
	TokenBulkScanResponseResultsFeaturesFeatureIDHighReputationToken     TokenBulkScanResponseResultsFeaturesFeatureID = "HIGH_REPUTATION_TOKEN"
	TokenBulkScanResponseResultsFeaturesFeatureIDStaticCodeSignature     TokenBulkScanResponseResultsFeaturesFeatureID = "STATIC_CODE_SIGNATURE"
	TokenBulkScanResponseResultsFeaturesFeatureIDKnownMalicious          TokenBulkScanResponseResultsFeaturesFeatureID = "KNOWN_MALICIOUS"
	TokenBulkScanResponseResultsFeaturesFeatureIDMetadata                TokenBulkScanResponseResultsFeaturesFeatureID = "METADATA"
	TokenBulkScanResponseResultsFeaturesFeatureIDAirdropPattern          TokenBulkScanResponseResultsFeaturesFeatureID = "AIRDROP_PATTERN"
	TokenBulkScanResponseResultsFeaturesFeatureIDImpersonator            TokenBulkScanResponseResultsFeaturesFeatureID = "IMPERSONATOR"
	TokenBulkScanResponseResultsFeaturesFeatureIDInorganicVolume         TokenBulkScanResponseResultsFeaturesFeatureID = "INORGANIC_VOLUME"
	TokenBulkScanResponseResultsFeaturesFeatureIDDynamicAnalysis         TokenBulkScanResponseResultsFeaturesFeatureID = "DYNAMIC_ANALYSIS"
	TokenBulkScanResponseResultsFeaturesFeatureIDUnstableTokenPrice      TokenBulkScanResponseResultsFeaturesFeatureID = "UNSTABLE_TOKEN_PRICE"
	TokenBulkScanResponseResultsFeaturesFeatureIDRugpull                 TokenBulkScanResponseResultsFeaturesFeatureID = "RUGPULL"
	TokenBulkScanResponseResultsFeaturesFeatureIDConsumerOverride        TokenBulkScanResponseResultsFeaturesFeatureID = "CONSUMER_OVERRIDE"
	TokenBulkScanResponseResultsFeaturesFeatureIDInappropriateContent    TokenBulkScanResponseResultsFeaturesFeatureID = "INAPPROPRIATE_CONTENT"
	TokenBulkScanResponseResultsFeaturesFeatureIDHighTransferFee         TokenBulkScanResponseResultsFeaturesFeatureID = "HIGH_TRANSFER_FEE"
	TokenBulkScanResponseResultsFeaturesFeatureIDHighBuyFee              TokenBulkScanResponseResultsFeaturesFeatureID = "HIGH_BUY_FEE"
	TokenBulkScanResponseResultsFeaturesFeatureIDHighSellFee             TokenBulkScanResponseResultsFeaturesFeatureID = "HIGH_SELL_FEE"
)

func (r TokenBulkScanResponseResultsFeaturesFeatureID) IsKnown() bool {
	switch r {
	case TokenBulkScanResponseResultsFeaturesFeatureIDVerifiedContract, TokenBulkScanResponseResultsFeaturesFeatureIDHighTradeVolume, TokenBulkScanResponseResultsFeaturesFeatureIDMarketPlaceSalesHistory, TokenBulkScanResponseResultsFeaturesFeatureIDHighReputationToken, TokenBulkScanResponseResultsFeaturesFeatureIDStaticCodeSignature, TokenBulkScanResponseResultsFeaturesFeatureIDKnownMalicious, TokenBulkScanResponseResultsFeaturesFeatureIDMetadata, TokenBulkScanResponseResultsFeaturesFeatureIDAirdropPattern, TokenBulkScanResponseResultsFeaturesFeatureIDImpersonator, TokenBulkScanResponseResultsFeaturesFeatureIDInorganicVolume, TokenBulkScanResponseResultsFeaturesFeatureIDDynamicAnalysis, TokenBulkScanResponseResultsFeaturesFeatureIDUnstableTokenPrice, TokenBulkScanResponseResultsFeaturesFeatureIDRugpull, TokenBulkScanResponseResultsFeaturesFeatureIDConsumerOverride, TokenBulkScanResponseResultsFeaturesFeatureIDInappropriateContent, TokenBulkScanResponseResultsFeaturesFeatureIDHighTransferFee, TokenBulkScanResponseResultsFeaturesFeatureIDHighBuyFee, TokenBulkScanResponseResultsFeaturesFeatureIDHighSellFee:
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
