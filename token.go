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
	// The chain name
	Chain TokenScanSupportedChain `json:"chain,required"`
	// Fees associated with the token
	Fees TokenScanResponseFees `json:"fees,required"`
	// financial stats of the token
	FinancialStats TokenScanResponseFinancialStats `json:"financial_stats,required"`
	// Score between 0 to 1 (double)
	MaliciousScore string `json:"malicious_score,required"`
	// Metadata of the token
	Metadata TokenScanResponseMetadata `json:"metadata,required"`
	// An enumeration.
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
	Buy float64 `json:"buy"`
	// Sell fee of the token
	Sell float64 `json:"sell"`
	// Transfer fee of the token
	Transfer float64                   `json:"transfer"`
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
	BurnedLiquidityPercentage float64                                    `json:"burned_liquidity_percentage"`
	HoldersCount              int64                                      `json:"holders_count"`
	LockedLiquidityPercentage float64                                    `json:"locked_liquidity_percentage"`
	TopHolders                []TokenScanResponseFinancialStatsTopHolder `json:"top_holders"`
	UsdPricePerUnit           float64                                    `json:"usd_price_per_unit"`
	JSON                      tokenScanResponseFinancialStatsJSON        `json:"-"`
}

// tokenScanResponseFinancialStatsJSON contains the JSON metadata for the struct
// [TokenScanResponseFinancialStats]
type tokenScanResponseFinancialStatsJSON struct {
	BurnedLiquidityPercentage apijson.Field
	HoldersCount              apijson.Field
	LockedLiquidityPercentage apijson.Field
	TopHolders                apijson.Field
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
	Address           string                                       `json:"address"`
	HoldingPercentage float64                                      `json:"holding_percentage"`
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
	FreezeAuthority string `json:"freeze_authority,required"`
	MintAuthority   string `json:"mint_authority,required"`
	// internal metadata
	TokenMetadata   TokenScanResponseMetadataTokenMetadata `json:"token_metadata,required"`
	UpdateAuthority string                                 `json:"update_authority,required"`
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
	Type string                        `json:"type"`
	JSON tokenScanResponseMetadataJSON `json:"-"`
}

// tokenScanResponseMetadataJSON contains the JSON metadata for the struct
// [TokenScanResponseMetadata]
type tokenScanResponseMetadataJSON struct {
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

func (r *TokenScanResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseMetadataJSON) RawJSON() string {
	return r.raw
}

// internal metadata
type TokenScanResponseMetadataTokenMetadata struct {
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
	Type TokenScanResponseMetadataTokenMetadataType `json:"type"`
	// Uri of the token
	Uri string `json:"uri"`
	// List of urls
	URLs []string                                   `json:"urls"`
	JSON tokenScanResponseMetadataTokenMetadataJSON `json:"-"`
}

// tokenScanResponseMetadataTokenMetadataJSON contains the JSON metadata for the
// struct [TokenScanResponseMetadataTokenMetadata]
type tokenScanResponseMetadataTokenMetadataJSON struct {
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

func (r *TokenScanResponseMetadataTokenMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseMetadataTokenMetadataJSON) RawJSON() string {
	return r.raw
}

// An enumeration.
type TokenScanResponseMetadataTokenMetadataType string

const (
	TokenScanResponseMetadataTokenMetadataTypeErc20       TokenScanResponseMetadataTokenMetadataType = "erc20"
	TokenScanResponseMetadataTokenMetadataTypeErc721      TokenScanResponseMetadataTokenMetadataType = "erc721"
	TokenScanResponseMetadataTokenMetadataTypeErc1155     TokenScanResponseMetadataTokenMetadataType = "erc1155"
	TokenScanResponseMetadataTokenMetadataTypeFungible    TokenScanResponseMetadataTokenMetadataType = "FUNGIBLE"
	TokenScanResponseMetadataTokenMetadataTypeNonFungible TokenScanResponseMetadataTokenMetadataType = "NonFungible"
)

func (r TokenScanResponseMetadataTokenMetadataType) IsKnown() bool {
	switch r {
	case TokenScanResponseMetadataTokenMetadataTypeErc20, TokenScanResponseMetadataTokenMetadataTypeErc721, TokenScanResponseMetadataTokenMetadataTypeErc1155, TokenScanResponseMetadataTokenMetadataTypeFungible, TokenScanResponseMetadataTokenMetadataTypeNonFungible:
		return true
	}
	return false
}

// An enumeration.
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
	MaxBuy            TokenScanResponseTradingLimitsMaxBuy            `json:"max_buy"`
	MaxHolding        TokenScanResponseTradingLimitsMaxHolding        `json:"max_holding"`
	MaxSell           TokenScanResponseTradingLimitsMaxSell           `json:"max_sell"`
	SellLimitPerBlock TokenScanResponseTradingLimitsSellLimitPerBlock `json:"sell_limit_per_block"`
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
	Amount    float64                                  `json:"amount"`
	AmountWei string                                   `json:"amount_wei"`
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
	Amount    float64                                      `json:"amount"`
	AmountWei string                                       `json:"amount_wei"`
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
	Amount    float64                                   `json:"amount"`
	AmountWei string                                    `json:"amount_wei"`
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
	Amount    float64                                             `json:"amount"`
	AmountWei string                                              `json:"amount_wei"`
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
	// An enumeration.
	FeatureID TokenScanResponseFeaturesFeatureID `json:"feature_id,required"`
	// An enumeration.
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

// An enumeration.
type TokenScanResponseFeaturesFeatureID string

const (
	TokenScanResponseFeaturesFeatureIDVerifiedContract        TokenScanResponseFeaturesFeatureID = "VERIFIED_CONTRACT"
	TokenScanResponseFeaturesFeatureIDHighTradeVolume         TokenScanResponseFeaturesFeatureID = "HIGH_TRADE_VOLUME"
	TokenScanResponseFeaturesFeatureIDMarketPlaceSalesHistory TokenScanResponseFeaturesFeatureID = "MARKET_PLACE_SALES_HISTORY"
	TokenScanResponseFeaturesFeatureIDHighReputationToken     TokenScanResponseFeaturesFeatureID = "HIGH_REPUTATION_TOKEN"
	TokenScanResponseFeaturesFeatureIDStaticCodeSignature     TokenScanResponseFeaturesFeatureID = "STATIC_CODE_SIGNATURE"
	TokenScanResponseFeaturesFeatureIDKnownMalicious          TokenScanResponseFeaturesFeatureID = "KNOWN_MALICIOUS"
	TokenScanResponseFeaturesFeatureIDMetadata                TokenScanResponseFeaturesFeatureID = "METADATA"
	TokenScanResponseFeaturesFeatureIDAirdropPattern          TokenScanResponseFeaturesFeatureID = "AIRDROP_PATTERN"
	TokenScanResponseFeaturesFeatureIDImpersonator            TokenScanResponseFeaturesFeatureID = "IMPERSONATOR"
	TokenScanResponseFeaturesFeatureIDInorganicVolume         TokenScanResponseFeaturesFeatureID = "INORGANIC_VOLUME"
	TokenScanResponseFeaturesFeatureIDDynamicAnalysis         TokenScanResponseFeaturesFeatureID = "DYNAMIC_ANALYSIS"
	TokenScanResponseFeaturesFeatureIDUnstableTokenPrice      TokenScanResponseFeaturesFeatureID = "UNSTABLE_TOKEN_PRICE"
	TokenScanResponseFeaturesFeatureIDRugpull                 TokenScanResponseFeaturesFeatureID = "RUGPULL"
	TokenScanResponseFeaturesFeatureIDConsumerOverride        TokenScanResponseFeaturesFeatureID = "CONSUMER_OVERRIDE"
	TokenScanResponseFeaturesFeatureIDInappropriateContent    TokenScanResponseFeaturesFeatureID = "INAPPROPRIATE_CONTENT"
	TokenScanResponseFeaturesFeatureIDHighTransferFee         TokenScanResponseFeaturesFeatureID = "HIGH_TRANSFER_FEE"
	TokenScanResponseFeaturesFeatureIDHighBuyFee              TokenScanResponseFeaturesFeatureID = "HIGH_BUY_FEE"
	TokenScanResponseFeaturesFeatureIDHighSellFee             TokenScanResponseFeaturesFeatureID = "HIGH_SELL_FEE"
)

func (r TokenScanResponseFeaturesFeatureID) IsKnown() bool {
	switch r {
	case TokenScanResponseFeaturesFeatureIDVerifiedContract, TokenScanResponseFeaturesFeatureIDHighTradeVolume, TokenScanResponseFeaturesFeatureIDMarketPlaceSalesHistory, TokenScanResponseFeaturesFeatureIDHighReputationToken, TokenScanResponseFeaturesFeatureIDStaticCodeSignature, TokenScanResponseFeaturesFeatureIDKnownMalicious, TokenScanResponseFeaturesFeatureIDMetadata, TokenScanResponseFeaturesFeatureIDAirdropPattern, TokenScanResponseFeaturesFeatureIDImpersonator, TokenScanResponseFeaturesFeatureIDInorganicVolume, TokenScanResponseFeaturesFeatureIDDynamicAnalysis, TokenScanResponseFeaturesFeatureIDUnstableTokenPrice, TokenScanResponseFeaturesFeatureIDRugpull, TokenScanResponseFeaturesFeatureIDConsumerOverride, TokenScanResponseFeaturesFeatureIDInappropriateContent, TokenScanResponseFeaturesFeatureIDHighTransferFee, TokenScanResponseFeaturesFeatureIDHighBuyFee, TokenScanResponseFeaturesFeatureIDHighSellFee:
		return true
	}
	return false
}

// An enumeration.
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
	// An enumeration.
	Event param.Field[TokenReportParamsEvent] `json:"event,required"`
	// The report parameters.
	Report param.Field[TokenReportParamsReportUnion] `json:"report,required"`
}

func (r TokenReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An enumeration.
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
	Params    param.Field[interface{}]                 `json:"params,required"`
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
