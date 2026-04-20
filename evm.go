// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"reflect"
	"time"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/param"
	"github.com/blockaid-official/blockaid-client-go/option"
	"github.com/tidwall/gjson"
)

// EvmService contains methods and other services that help with interacting with
// the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvmService] method instead.
type EvmService struct {
	Options             []option.RequestOption
	JsonRpc             *EvmJsonRpcService
	Transaction         *EvmTransactionService
	TransactionBulk     *EvmTransactionBulkService
	TransactionRaw      *EvmTransactionRawService
	UserOperation       *EvmUserOperationService
	PostTransaction     *EvmPostTransactionService
	PostTransactionBulk *EvmPostTransactionBulkService
	Address             *EvmAddressService
	AddressBulk         *EvmAddressBulkService
}

// NewEvmService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEvmService(opts ...option.RequestOption) (r *EvmService) {
	r = &EvmService{}
	r.Options = opts
	r.JsonRpc = NewEvmJsonRpcService(opts...)
	r.Transaction = NewEvmTransactionService(opts...)
	r.TransactionBulk = NewEvmTransactionBulkService(opts...)
	r.TransactionRaw = NewEvmTransactionRawService(opts...)
	r.UserOperation = NewEvmUserOperationService(opts...)
	r.PostTransaction = NewEvmPostTransactionService(opts...)
	r.PostTransactionBulk = NewEvmPostTransactionBulkService(opts...)
	r.Address = NewEvmAddressService(opts...)
	r.AddressBulk = NewEvmAddressBulkService(opts...)
	return
}

type AddressValidation struct {
	// Overall validation outcome for the scan.
	ResultType AddressValidationResultType `json:"result_type" api:"required"`
	// An error message returned when `result_type` is `Error`.
	Error string `json:"error"`
	// A list of features explaining the scan result (each feature includes a type,
	// feature_id, and description).
	Features AddressValidationFeaturesUnion `json:"features"`
	JSON     addressValidationJSON          `json:"-"`
}

// addressValidationJSON contains the JSON metadata for the struct
// [AddressValidation]
type addressValidationJSON struct {
	ResultType  apijson.Field
	Error       apijson.Field
	Features    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressValidationJSON) RawJSON() string {
	return r.raw
}

// Overall validation outcome for the scan.
type AddressValidationResultType string

const (
	AddressValidationResultTypeMalicious AddressValidationResultType = "Malicious"
	AddressValidationResultTypeWarning   AddressValidationResultType = "Warning"
	AddressValidationResultTypeBenign    AddressValidationResultType = "Benign"
	AddressValidationResultTypeError     AddressValidationResultType = "Error"
)

func (r AddressValidationResultType) IsKnown() bool {
	switch r {
	case AddressValidationResultTypeMalicious, AddressValidationResultTypeWarning, AddressValidationResultTypeBenign, AddressValidationResultTypeError:
		return true
	}
	return false
}

// A list of features explaining the scan result (each feature includes a type,
// feature_id, and description).
//
// Union satisfied by [AddressValidationFeaturesArray] or
// [AddressValidationFeaturesArray].
type AddressValidationFeaturesUnion interface {
	implementsAddressValidationFeaturesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressValidationFeaturesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AddressValidationFeaturesArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AddressValidationFeaturesArray{}),
		},
	)
}

type AddressValidationFeaturesArray []AddressValidationFeaturesArrayItem

func (r AddressValidationFeaturesArray) implementsAddressValidationFeaturesUnion() {}

type AddressValidationFeaturesArrayItem struct {
	// Description of the feature
	Description string `json:"description" api:"required"`
	// Feature identifier
	FeatureID AddressValidationFeaturesArrayFeatureID `json:"feature_id" api:"required"`
	// Type of the feature
	Type AddressValidationFeaturesArrayType     `json:"type" api:"required"`
	JSON addressValidationFeaturesArrayItemJSON `json:"-"`
}

// addressValidationFeaturesArrayItemJSON contains the JSON metadata for the struct
// [AddressValidationFeaturesArrayItem]
type addressValidationFeaturesArrayItemJSON struct {
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressValidationFeaturesArrayItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressValidationFeaturesArrayItemJSON) RawJSON() string {
	return r.raw
}

// Feature identifier
type AddressValidationFeaturesArrayFeatureID string

const (
	AddressValidationFeaturesArrayFeatureIDVerifiedContract               AddressValidationFeaturesArrayFeatureID = "VERIFIED_CONTRACT"
	AddressValidationFeaturesArrayFeatureIDUnverifiedContract             AddressValidationFeaturesArrayFeatureID = "UNVERIFIED_CONTRACT"
	AddressValidationFeaturesArrayFeatureIDHighTradeVolume                AddressValidationFeaturesArrayFeatureID = "HIGH_TRADE_VOLUME"
	AddressValidationFeaturesArrayFeatureIDMarketPlaceSalesHistory        AddressValidationFeaturesArrayFeatureID = "MARKET_PLACE_SALES_HISTORY"
	AddressValidationFeaturesArrayFeatureIDHighReputationToken            AddressValidationFeaturesArrayFeatureID = "HIGH_REPUTATION_TOKEN"
	AddressValidationFeaturesArrayFeatureIDOnchainActivityValidator       AddressValidationFeaturesArrayFeatureID = "ONCHAIN_ACTIVITY_VALIDATOR"
	AddressValidationFeaturesArrayFeatureIDStaticCodeSignature            AddressValidationFeaturesArrayFeatureID = "STATIC_CODE_SIGNATURE"
	AddressValidationFeaturesArrayFeatureIDKnownMalicious                 AddressValidationFeaturesArrayFeatureID = "KNOWN_MALICIOUS"
	AddressValidationFeaturesArrayFeatureIDIsEoa                          AddressValidationFeaturesArrayFeatureID = "IS_EOA"
	AddressValidationFeaturesArrayFeatureIDIsContract                     AddressValidationFeaturesArrayFeatureID = "IS_CONTRACT"
	AddressValidationFeaturesArrayFeatureIDErc20Contract                  AddressValidationFeaturesArrayFeatureID = "ERC20_CONTRACT"
	AddressValidationFeaturesArrayFeatureIDTrustedContract                AddressValidationFeaturesArrayFeatureID = "TRUSTED_CONTRACT"
	AddressValidationFeaturesArrayFeatureIDBenignCreator                  AddressValidationFeaturesArrayFeatureID = "BENIGN_CREATOR"
	AddressValidationFeaturesArrayFeatureIDMetadata                       AddressValidationFeaturesArrayFeatureID = "METADATA"
	AddressValidationFeaturesArrayFeatureIDAirdropPattern                 AddressValidationFeaturesArrayFeatureID = "AIRDROP_PATTERN"
	AddressValidationFeaturesArrayFeatureIDImpersonator                   AddressValidationFeaturesArrayFeatureID = "IMPERSONATOR"
	AddressValidationFeaturesArrayFeatureIDInorganicVolume                AddressValidationFeaturesArrayFeatureID = "INORGANIC_VOLUME"
	AddressValidationFeaturesArrayFeatureIDDynamicAnalysis                AddressValidationFeaturesArrayFeatureID = "DYNAMIC_ANALYSIS"
	AddressValidationFeaturesArrayFeatureIDConcentratedSupplyDistribution AddressValidationFeaturesArrayFeatureID = "CONCENTRATED_SUPPLY_DISTRIBUTION"
	AddressValidationFeaturesArrayFeatureIDHoneypot                       AddressValidationFeaturesArrayFeatureID = "HONEYPOT"
	AddressValidationFeaturesArrayFeatureIDInsufficientLockedLiquidity    AddressValidationFeaturesArrayFeatureID = "INSUFFICIENT_LOCKED_LIQUIDITY"
	AddressValidationFeaturesArrayFeatureIDUnstableTokenPrice             AddressValidationFeaturesArrayFeatureID = "UNSTABLE_TOKEN_PRICE"
	AddressValidationFeaturesArrayFeatureIDRugpull                        AddressValidationFeaturesArrayFeatureID = "RUGPULL"
	AddressValidationFeaturesArrayFeatureIDWashTrading                    AddressValidationFeaturesArrayFeatureID = "WASH_TRADING"
	AddressValidationFeaturesArrayFeatureIDConsumerOverride               AddressValidationFeaturesArrayFeatureID = "CONSUMER_OVERRIDE"
	AddressValidationFeaturesArrayFeatureIDInappropriateContent           AddressValidationFeaturesArrayFeatureID = "INAPPROPRIATE_CONTENT"
	AddressValidationFeaturesArrayFeatureIDHighTransferFee                AddressValidationFeaturesArrayFeatureID = "HIGH_TRANSFER_FEE"
	AddressValidationFeaturesArrayFeatureIDHighBuyFee                     AddressValidationFeaturesArrayFeatureID = "HIGH_BUY_FEE"
	AddressValidationFeaturesArrayFeatureIDHighSellFee                    AddressValidationFeaturesArrayFeatureID = "HIGH_SELL_FEE"
	AddressValidationFeaturesArrayFeatureIDUnsellableToken                AddressValidationFeaturesArrayFeatureID = "UNSELLABLE_TOKEN"
	AddressValidationFeaturesArrayFeatureIDIsMintable                     AddressValidationFeaturesArrayFeatureID = "IS_MINTABLE"
	AddressValidationFeaturesArrayFeatureIDRebaseToken                    AddressValidationFeaturesArrayFeatureID = "REBASE_TOKEN"
	AddressValidationFeaturesArrayFeatureIDLiquidStakingToken             AddressValidationFeaturesArrayFeatureID = "LIQUID_STAKING_TOKEN"
	AddressValidationFeaturesArrayFeatureIDModifiableTaxes                AddressValidationFeaturesArrayFeatureID = "MODIFIABLE_TAXES"
	AddressValidationFeaturesArrayFeatureIDCanBlacklist                   AddressValidationFeaturesArrayFeatureID = "CAN_BLACKLIST"
	AddressValidationFeaturesArrayFeatureIDCanWhitelist                   AddressValidationFeaturesArrayFeatureID = "CAN_WHITELIST"
	AddressValidationFeaturesArrayFeatureIDHasTradingCooldown             AddressValidationFeaturesArrayFeatureID = "HAS_TRADING_COOLDOWN"
	AddressValidationFeaturesArrayFeatureIDExternalFunctions              AddressValidationFeaturesArrayFeatureID = "EXTERNAL_FUNCTIONS"
	AddressValidationFeaturesArrayFeatureIDHiddenOwner                    AddressValidationFeaturesArrayFeatureID = "HIDDEN_OWNER"
	AddressValidationFeaturesArrayFeatureIDTransferPauseable              AddressValidationFeaturesArrayFeatureID = "TRANSFER_PAUSEABLE"
	AddressValidationFeaturesArrayFeatureIDOwnershipRenounced             AddressValidationFeaturesArrayFeatureID = "OWNERSHIP_RENOUNCED"
	AddressValidationFeaturesArrayFeatureIDOwnerCanChangeBalance          AddressValidationFeaturesArrayFeatureID = "OWNER_CAN_CHANGE_BALANCE"
	AddressValidationFeaturesArrayFeatureIDProxyContract                  AddressValidationFeaturesArrayFeatureID = "PROXY_CONTRACT"
	AddressValidationFeaturesArrayFeatureIDSimilarMaliciousContract       AddressValidationFeaturesArrayFeatureID = "SIMILAR_MALICIOUS_CONTRACT"
	AddressValidationFeaturesArrayFeatureIDImpersonatorSensitiveAsset     AddressValidationFeaturesArrayFeatureID = "IMPERSONATOR_SENSITIVE_ASSET"
	AddressValidationFeaturesArrayFeatureIDImpersonatorHighConfidence     AddressValidationFeaturesArrayFeatureID = "IMPERSONATOR_HIGH_CONFIDENCE"
	AddressValidationFeaturesArrayFeatureIDImpersonatorMediumConfidence   AddressValidationFeaturesArrayFeatureID = "IMPERSONATOR_MEDIUM_CONFIDENCE"
	AddressValidationFeaturesArrayFeatureIDImpersonatorLowConfidence      AddressValidationFeaturesArrayFeatureID = "IMPERSONATOR_LOW_CONFIDENCE"
	AddressValidationFeaturesArrayFeatureIDImpersonationProtected         AddressValidationFeaturesArrayFeatureID = "IMPERSONATION_PROTECTED"
	AddressValidationFeaturesArrayFeatureIDFakeVolume                     AddressValidationFeaturesArrayFeatureID = "FAKE_VOLUME"
	AddressValidationFeaturesArrayFeatureIDHiddenSupplyByKeyHolder        AddressValidationFeaturesArrayFeatureID = "HIDDEN_SUPPLY_BY_KEY_HOLDER"
	AddressValidationFeaturesArrayFeatureIDFakeTradeMakerCount            AddressValidationFeaturesArrayFeatureID = "FAKE_TRADE_MAKER_COUNT"
	AddressValidationFeaturesArrayFeatureIDTransferFromReverts            AddressValidationFeaturesArrayFeatureID = "TRANSFER_FROM_REVERTS"
	AddressValidationFeaturesArrayFeatureIDOffensiveTokenMetadata         AddressValidationFeaturesArrayFeatureID = "OFFENSIVE_TOKEN_METADATA"
	AddressValidationFeaturesArrayFeatureIDListedOnCentralizedExchange    AddressValidationFeaturesArrayFeatureID = "LISTED_ON_CENTRALIZED_EXCHANGE"
	AddressValidationFeaturesArrayFeatureIDSanctionedCreator              AddressValidationFeaturesArrayFeatureID = "SANCTIONED_CREATOR"
	AddressValidationFeaturesArrayFeatureIDSpamText                       AddressValidationFeaturesArrayFeatureID = "SPAM_TEXT"
	AddressValidationFeaturesArrayFeatureIDBondingCurveToken              AddressValidationFeaturesArrayFeatureID = "BONDING_CURVE_TOKEN"
	AddressValidationFeaturesArrayFeatureIDHeavilySniped                  AddressValidationFeaturesArrayFeatureID = "HEAVILY_SNIPED"
	AddressValidationFeaturesArrayFeatureIDSolanaToken2022                AddressValidationFeaturesArrayFeatureID = "SOLANA_TOKEN_2022"
	AddressValidationFeaturesArrayFeatureIDPostDump                       AddressValidationFeaturesArrayFeatureID = "POST_DUMP"
	AddressValidationFeaturesArrayFeatureIDDexPaid                        AddressValidationFeaturesArrayFeatureID = "DEX_PAID"
	AddressValidationFeaturesArrayFeatureIDLowReputationCreator           AddressValidationFeaturesArrayFeatureID = "LOW_REPUTATION_CREATOR"
	AddressValidationFeaturesArrayFeatureIDSnipeAtMint                    AddressValidationFeaturesArrayFeatureID = "SNIPE_AT_MINT"
	AddressValidationFeaturesArrayFeatureIDTransferHookEnabled            AddressValidationFeaturesArrayFeatureID = "TRANSFER_HOOK_ENABLED"
	AddressValidationFeaturesArrayFeatureIDConfidentialTransfersEnabled   AddressValidationFeaturesArrayFeatureID = "CONFIDENTIAL_TRANSFERS_ENABLED"
	AddressValidationFeaturesArrayFeatureIDNonTranserable                 AddressValidationFeaturesArrayFeatureID = "NON_TRANSERABLE"
	AddressValidationFeaturesArrayFeatureIDTokenBackdoor                  AddressValidationFeaturesArrayFeatureID = "TOKEN_BACKDOOR"
	AddressValidationFeaturesArrayFeatureIDCreatedViaLaunchpad            AddressValidationFeaturesArrayFeatureID = "CREATED_VIA_LAUNCHPAD"
	AddressValidationFeaturesArrayFeatureIDCompromisedToken               AddressValidationFeaturesArrayFeatureID = "COMPROMISED_TOKEN"
	AddressValidationFeaturesArrayFeatureIDLongFundTrail                  AddressValidationFeaturesArrayFeatureID = "LONG_FUND_TRAIL"
)

func (r AddressValidationFeaturesArrayFeatureID) IsKnown() bool {
	switch r {
	case AddressValidationFeaturesArrayFeatureIDVerifiedContract, AddressValidationFeaturesArrayFeatureIDUnverifiedContract, AddressValidationFeaturesArrayFeatureIDHighTradeVolume, AddressValidationFeaturesArrayFeatureIDMarketPlaceSalesHistory, AddressValidationFeaturesArrayFeatureIDHighReputationToken, AddressValidationFeaturesArrayFeatureIDOnchainActivityValidator, AddressValidationFeaturesArrayFeatureIDStaticCodeSignature, AddressValidationFeaturesArrayFeatureIDKnownMalicious, AddressValidationFeaturesArrayFeatureIDIsEoa, AddressValidationFeaturesArrayFeatureIDIsContract, AddressValidationFeaturesArrayFeatureIDErc20Contract, AddressValidationFeaturesArrayFeatureIDTrustedContract, AddressValidationFeaturesArrayFeatureIDBenignCreator, AddressValidationFeaturesArrayFeatureIDMetadata, AddressValidationFeaturesArrayFeatureIDAirdropPattern, AddressValidationFeaturesArrayFeatureIDImpersonator, AddressValidationFeaturesArrayFeatureIDInorganicVolume, AddressValidationFeaturesArrayFeatureIDDynamicAnalysis, AddressValidationFeaturesArrayFeatureIDConcentratedSupplyDistribution, AddressValidationFeaturesArrayFeatureIDHoneypot, AddressValidationFeaturesArrayFeatureIDInsufficientLockedLiquidity, AddressValidationFeaturesArrayFeatureIDUnstableTokenPrice, AddressValidationFeaturesArrayFeatureIDRugpull, AddressValidationFeaturesArrayFeatureIDWashTrading, AddressValidationFeaturesArrayFeatureIDConsumerOverride, AddressValidationFeaturesArrayFeatureIDInappropriateContent, AddressValidationFeaturesArrayFeatureIDHighTransferFee, AddressValidationFeaturesArrayFeatureIDHighBuyFee, AddressValidationFeaturesArrayFeatureIDHighSellFee, AddressValidationFeaturesArrayFeatureIDUnsellableToken, AddressValidationFeaturesArrayFeatureIDIsMintable, AddressValidationFeaturesArrayFeatureIDRebaseToken, AddressValidationFeaturesArrayFeatureIDLiquidStakingToken, AddressValidationFeaturesArrayFeatureIDModifiableTaxes, AddressValidationFeaturesArrayFeatureIDCanBlacklist, AddressValidationFeaturesArrayFeatureIDCanWhitelist, AddressValidationFeaturesArrayFeatureIDHasTradingCooldown, AddressValidationFeaturesArrayFeatureIDExternalFunctions, AddressValidationFeaturesArrayFeatureIDHiddenOwner, AddressValidationFeaturesArrayFeatureIDTransferPauseable, AddressValidationFeaturesArrayFeatureIDOwnershipRenounced, AddressValidationFeaturesArrayFeatureIDOwnerCanChangeBalance, AddressValidationFeaturesArrayFeatureIDProxyContract, AddressValidationFeaturesArrayFeatureIDSimilarMaliciousContract, AddressValidationFeaturesArrayFeatureIDImpersonatorSensitiveAsset, AddressValidationFeaturesArrayFeatureIDImpersonatorHighConfidence, AddressValidationFeaturesArrayFeatureIDImpersonatorMediumConfidence, AddressValidationFeaturesArrayFeatureIDImpersonatorLowConfidence, AddressValidationFeaturesArrayFeatureIDImpersonationProtected, AddressValidationFeaturesArrayFeatureIDFakeVolume, AddressValidationFeaturesArrayFeatureIDHiddenSupplyByKeyHolder, AddressValidationFeaturesArrayFeatureIDFakeTradeMakerCount, AddressValidationFeaturesArrayFeatureIDTransferFromReverts, AddressValidationFeaturesArrayFeatureIDOffensiveTokenMetadata, AddressValidationFeaturesArrayFeatureIDListedOnCentralizedExchange, AddressValidationFeaturesArrayFeatureIDSanctionedCreator, AddressValidationFeaturesArrayFeatureIDSpamText, AddressValidationFeaturesArrayFeatureIDBondingCurveToken, AddressValidationFeaturesArrayFeatureIDHeavilySniped, AddressValidationFeaturesArrayFeatureIDSolanaToken2022, AddressValidationFeaturesArrayFeatureIDPostDump, AddressValidationFeaturesArrayFeatureIDDexPaid, AddressValidationFeaturesArrayFeatureIDLowReputationCreator, AddressValidationFeaturesArrayFeatureIDSnipeAtMint, AddressValidationFeaturesArrayFeatureIDTransferHookEnabled, AddressValidationFeaturesArrayFeatureIDConfidentialTransfersEnabled, AddressValidationFeaturesArrayFeatureIDNonTranserable, AddressValidationFeaturesArrayFeatureIDTokenBackdoor, AddressValidationFeaturesArrayFeatureIDCreatedViaLaunchpad, AddressValidationFeaturesArrayFeatureIDCompromisedToken, AddressValidationFeaturesArrayFeatureIDLongFundTrail:
		return true
	}
	return false
}

// Type of the feature
type AddressValidationFeaturesArrayType string

const (
	AddressValidationFeaturesArrayTypeBenign    AddressValidationFeaturesArrayType = "Benign"
	AddressValidationFeaturesArrayTypeInfo      AddressValidationFeaturesArrayType = "Info"
	AddressValidationFeaturesArrayTypeWarning   AddressValidationFeaturesArrayType = "Warning"
	AddressValidationFeaturesArrayTypeMalicious AddressValidationFeaturesArrayType = "Malicious"
)

func (r AddressValidationFeaturesArrayType) IsKnown() bool {
	switch r {
	case AddressValidationFeaturesArrayTypeBenign, AddressValidationFeaturesArrayTypeInfo, AddressValidationFeaturesArrayTypeWarning, AddressValidationFeaturesArrayTypeMalicious:
		return true
	}
	return false
}

type AuthorizationParam struct {
	// The delegation designation address
	Address param.Field[string] `json:"address" api:"required"`
	// The chain ID as hex string
	ChainID param.Field[string] `json:"chainId"`
	// The authority address of the delegation, should be provided when the signature
	// (r,s,yParity) is not provided in order to simulate the transaction with the
	// correct delegation
	Eoa param.Field[string] `json:"eoa"`
	// The nonce value as hex string
	Nonce param.Field[string] `json:"nonce"`
	// The r value as hex string
	R param.Field[string] `json:"r"`
	// The s value as hex string
	S param.Field[string] `json:"s"`
	// The yParity value as hex string
	YParity param.Field[string] `json:"yParity"`
}

func (r AuthorizationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MetadataParam struct {
	// Account information associated with the request
	Account param.Field[MetadataParamAccount] `json:"account"`
	// Connection metadata including user agent and IP information
	Connection param.Field[MetadataParamConnection] `json:"connection"`
}

func (r MetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Account information associated with the request
type MetadataParamAccount struct {
	// Unique identifier for the account.
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// Timestamp when the account was created.
	AccountCreationTimestamp param.Field[time.Time] `json:"account_creation_timestamp" format:"date-time"`
	// Age of the user in years
	UserAge param.Field[int64] `json:"user_age"`
	// ISO country code of the user's location.
	UserCountryCode param.Field[string] `json:"user_country_code"`
}

func (r MetadataParamAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Connection metadata including user agent and IP information
type MetadataParamConnection struct {
	// IP address of the customer making the request.
	IPAddress param.Field[string] `json:"ip_address" api:"required" format:"ipvanyaddress"`
	// User agent string from the client's browser or application.
	UserAgent param.Field[string] `json:"user_agent"`
}

func (r MetadataParamConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The chain name
type TokenScanSupportedChain string

const (
	TokenScanSupportedChainArbitrum        TokenScanSupportedChain = "arbitrum"
	TokenScanSupportedChainAvalanche       TokenScanSupportedChain = "avalanche"
	TokenScanSupportedChainBase            TokenScanSupportedChain = "base"
	TokenScanSupportedChainBsc             TokenScanSupportedChain = "bsc"
	TokenScanSupportedChainEthereum        TokenScanSupportedChain = "ethereum"
	TokenScanSupportedChainOptimism        TokenScanSupportedChain = "optimism"
	TokenScanSupportedChainPolygon         TokenScanSupportedChain = "polygon"
	TokenScanSupportedChainZora            TokenScanSupportedChain = "zora"
	TokenScanSupportedChainSolana          TokenScanSupportedChain = "solana"
	TokenScanSupportedChainStarknet        TokenScanSupportedChain = "starknet"
	TokenScanSupportedChainStarknetSepolia TokenScanSupportedChain = "starknet-sepolia"
	TokenScanSupportedChainStellar         TokenScanSupportedChain = "stellar"
	TokenScanSupportedChainLinea           TokenScanSupportedChain = "linea"
	TokenScanSupportedChainDegen           TokenScanSupportedChain = "degen"
	TokenScanSupportedChainZksync          TokenScanSupportedChain = "zksync"
	TokenScanSupportedChainScroll          TokenScanSupportedChain = "scroll"
	TokenScanSupportedChainBlast           TokenScanSupportedChain = "blast"
	TokenScanSupportedChainSoneiumMinato   TokenScanSupportedChain = "soneium-minato"
	TokenScanSupportedChainBaseSepolia     TokenScanSupportedChain = "base-sepolia"
	TokenScanSupportedChainBitcoin         TokenScanSupportedChain = "bitcoin"
	TokenScanSupportedChainAbstract        TokenScanSupportedChain = "abstract"
	TokenScanSupportedChainSoneium         TokenScanSupportedChain = "soneium"
	TokenScanSupportedChainInk             TokenScanSupportedChain = "ink"
	TokenScanSupportedChainZeroNetwork     TokenScanSupportedChain = "zero-network"
	TokenScanSupportedChainBerachain       TokenScanSupportedChain = "berachain"
	TokenScanSupportedChainUnichain        TokenScanSupportedChain = "unichain"
	TokenScanSupportedChainRonin           TokenScanSupportedChain = "ronin"
	TokenScanSupportedChainSui             TokenScanSupportedChain = "sui"
	TokenScanSupportedChainHedera          TokenScanSupportedChain = "hedera"
	TokenScanSupportedChainHyperevm        TokenScanSupportedChain = "hyperevm"
	TokenScanSupportedChainXlayer          TokenScanSupportedChain = "xlayer"
	TokenScanSupportedChainMonad           TokenScanSupportedChain = "monad"
	TokenScanSupportedChainMegaeth         TokenScanSupportedChain = "megaeth"
	TokenScanSupportedChainTempo           TokenScanSupportedChain = "tempo"
	TokenScanSupportedChainSei             TokenScanSupportedChain = "sei"
	TokenScanSupportedChainTron            TokenScanSupportedChain = "tron"
)

func (r TokenScanSupportedChain) IsKnown() bool {
	switch r {
	case TokenScanSupportedChainArbitrum, TokenScanSupportedChainAvalanche, TokenScanSupportedChainBase, TokenScanSupportedChainBsc, TokenScanSupportedChainEthereum, TokenScanSupportedChainOptimism, TokenScanSupportedChainPolygon, TokenScanSupportedChainZora, TokenScanSupportedChainSolana, TokenScanSupportedChainStarknet, TokenScanSupportedChainStarknetSepolia, TokenScanSupportedChainStellar, TokenScanSupportedChainLinea, TokenScanSupportedChainDegen, TokenScanSupportedChainZksync, TokenScanSupportedChainScroll, TokenScanSupportedChainBlast, TokenScanSupportedChainSoneiumMinato, TokenScanSupportedChainBaseSepolia, TokenScanSupportedChainBitcoin, TokenScanSupportedChainAbstract, TokenScanSupportedChainSoneium, TokenScanSupportedChainInk, TokenScanSupportedChainZeroNetwork, TokenScanSupportedChainBerachain, TokenScanSupportedChainUnichain, TokenScanSupportedChainRonin, TokenScanSupportedChainSui, TokenScanSupportedChainHedera, TokenScanSupportedChainHyperevm, TokenScanSupportedChainXlayer, TokenScanSupportedChainMonad, TokenScanSupportedChainMegaeth, TokenScanSupportedChainTempo, TokenScanSupportedChainSei, TokenScanSupportedChainTron:
		return true
	}
	return false
}

// The chain name
type TransactionScanSupportedChain string

const (
	TransactionScanSupportedChainArbitrum              TransactionScanSupportedChain = "arbitrum"
	TransactionScanSupportedChainAvalanche             TransactionScanSupportedChain = "avalanche"
	TransactionScanSupportedChainBase                  TransactionScanSupportedChain = "base"
	TransactionScanSupportedChainBaseSepolia           TransactionScanSupportedChain = "base-sepolia"
	TransactionScanSupportedChainLordchain             TransactionScanSupportedChain = "lordchain"
	TransactionScanSupportedChainLordchainTestnet      TransactionScanSupportedChain = "lordchain-testnet"
	TransactionScanSupportedChainMetacade              TransactionScanSupportedChain = "metacade"
	TransactionScanSupportedChainMetacadeTestnet       TransactionScanSupportedChain = "metacade-testnet"
	TransactionScanSupportedChainBsc                   TransactionScanSupportedChain = "bsc"
	TransactionScanSupportedChainEthereum              TransactionScanSupportedChain = "ethereum"
	TransactionScanSupportedChainOptimism              TransactionScanSupportedChain = "optimism"
	TransactionScanSupportedChainPolygon               TransactionScanSupportedChain = "polygon"
	TransactionScanSupportedChainZksync                TransactionScanSupportedChain = "zksync"
	TransactionScanSupportedChainZksyncSepolia         TransactionScanSupportedChain = "zksync-sepolia"
	TransactionScanSupportedChainZora                  TransactionScanSupportedChain = "zora"
	TransactionScanSupportedChainLinea                 TransactionScanSupportedChain = "linea"
	TransactionScanSupportedChainBlast                 TransactionScanSupportedChain = "blast"
	TransactionScanSupportedChainScroll                TransactionScanSupportedChain = "scroll"
	TransactionScanSupportedChainEthereumSepolia       TransactionScanSupportedChain = "ethereum-sepolia"
	TransactionScanSupportedChainDegen                 TransactionScanSupportedChain = "degen"
	TransactionScanSupportedChainAvalancheFuji         TransactionScanSupportedChain = "avalanche-fuji"
	TransactionScanSupportedChainImmutableZkevm        TransactionScanSupportedChain = "immutable-zkevm"
	TransactionScanSupportedChainImmutableZkevmTestnet TransactionScanSupportedChain = "immutable-zkevm-testnet"
	TransactionScanSupportedChainGnosis                TransactionScanSupportedChain = "gnosis"
	TransactionScanSupportedChainWorldchain            TransactionScanSupportedChain = "worldchain"
	TransactionScanSupportedChainSoneiumMinato         TransactionScanSupportedChain = "soneium-minato"
	TransactionScanSupportedChainRonin                 TransactionScanSupportedChain = "ronin"
	TransactionScanSupportedChainApechain              TransactionScanSupportedChain = "apechain"
	TransactionScanSupportedChainZeroNetwork           TransactionScanSupportedChain = "zero-network"
	TransactionScanSupportedChainBerachain             TransactionScanSupportedChain = "berachain"
	TransactionScanSupportedChainBerachainBartio       TransactionScanSupportedChain = "berachain-bartio"
	TransactionScanSupportedChainInk                   TransactionScanSupportedChain = "ink"
	TransactionScanSupportedChainInkSepolia            TransactionScanSupportedChain = "ink-sepolia"
	TransactionScanSupportedChainAbstract              TransactionScanSupportedChain = "abstract"
	TransactionScanSupportedChainAbstractTestnet       TransactionScanSupportedChain = "abstract-testnet"
	TransactionScanSupportedChainSoneium               TransactionScanSupportedChain = "soneium"
	TransactionScanSupportedChainUnichain              TransactionScanSupportedChain = "unichain"
	TransactionScanSupportedChainSei                   TransactionScanSupportedChain = "sei"
	TransactionScanSupportedChainFlowEvm               TransactionScanSupportedChain = "flow-evm"
	TransactionScanSupportedChainHyperevm              TransactionScanSupportedChain = "hyperevm"
	TransactionScanSupportedChainMegaeth               TransactionScanSupportedChain = "megaeth"
	TransactionScanSupportedChainKatana                TransactionScanSupportedChain = "katana"
	TransactionScanSupportedChainPlume                 TransactionScanSupportedChain = "plume"
	TransactionScanSupportedChainXlayer                TransactionScanSupportedChain = "xlayer"
	TransactionScanSupportedChainMonad                 TransactionScanSupportedChain = "monad"
	TransactionScanSupportedChainMonadTestnet          TransactionScanSupportedChain = "monad-testnet"
	TransactionScanSupportedChainTempo                 TransactionScanSupportedChain = "tempo"
	TransactionScanSupportedChainTempoTestnet          TransactionScanSupportedChain = "tempo-testnet"
	TransactionScanSupportedChainKiteAI                TransactionScanSupportedChain = "kite-ai"
)

func (r TransactionScanSupportedChain) IsKnown() bool {
	switch r {
	case TransactionScanSupportedChainArbitrum, TransactionScanSupportedChainAvalanche, TransactionScanSupportedChainBase, TransactionScanSupportedChainBaseSepolia, TransactionScanSupportedChainLordchain, TransactionScanSupportedChainLordchainTestnet, TransactionScanSupportedChainMetacade, TransactionScanSupportedChainMetacadeTestnet, TransactionScanSupportedChainBsc, TransactionScanSupportedChainEthereum, TransactionScanSupportedChainOptimism, TransactionScanSupportedChainPolygon, TransactionScanSupportedChainZksync, TransactionScanSupportedChainZksyncSepolia, TransactionScanSupportedChainZora, TransactionScanSupportedChainLinea, TransactionScanSupportedChainBlast, TransactionScanSupportedChainScroll, TransactionScanSupportedChainEthereumSepolia, TransactionScanSupportedChainDegen, TransactionScanSupportedChainAvalancheFuji, TransactionScanSupportedChainImmutableZkevm, TransactionScanSupportedChainImmutableZkevmTestnet, TransactionScanSupportedChainGnosis, TransactionScanSupportedChainWorldchain, TransactionScanSupportedChainSoneiumMinato, TransactionScanSupportedChainRonin, TransactionScanSupportedChainApechain, TransactionScanSupportedChainZeroNetwork, TransactionScanSupportedChainBerachain, TransactionScanSupportedChainBerachainBartio, TransactionScanSupportedChainInk, TransactionScanSupportedChainInkSepolia, TransactionScanSupportedChainAbstract, TransactionScanSupportedChainAbstractTestnet, TransactionScanSupportedChainSoneium, TransactionScanSupportedChainUnichain, TransactionScanSupportedChainSei, TransactionScanSupportedChainFlowEvm, TransactionScanSupportedChainHyperevm, TransactionScanSupportedChainMegaeth, TransactionScanSupportedChainKatana, TransactionScanSupportedChainPlume, TransactionScanSupportedChainXlayer, TransactionScanSupportedChainMonad, TransactionScanSupportedChainMonadTestnet, TransactionScanSupportedChainTempo, TransactionScanSupportedChainTempoTestnet, TransactionScanSupportedChainKiteAI:
		return true
	}
	return false
}

type UserOperationDataParam struct {
	// The operation parameters of the user operation request
	Operation param.Field[UserOperationDataOperationUnionParam] `json:"operation" api:"required"`
	// The address of the entrypoint receiving the request in hex string format
	Entrypoint param.Field[string] `json:"entrypoint"`
}

func (r UserOperationDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The operation parameters of the user operation request
type UserOperationDataOperationParam struct {
	// The account gas limits value in hex string format.
	AccountGasLimits param.Field[string] `json:"account_gas_limits"`
	// The call data value in hex string format.
	CallData param.Field[string] `json:"call_data"`
	// The call gas limit value in hex string format.
	CallGasLimit param.Field[string] `json:"call_gas_limit"`
	// The EIP-7702 authorization tuple for the user operation (optional)
	Eip7702Auth param.Field[AuthorizationParam] `json:"eip7702_auth"`
	// The gas fees value in hex string format.
	GasFees param.Field[string] `json:"gas_fees"`
	// The init code value in hex string format.
	InitCode param.Field[string] `json:"init_code"`
	// The max fee per gas value in hex string format.
	MaxFeePerGas param.Field[string] `json:"max_fee_per_gas"`
	// The max priority fee per gas value in hex string format.
	MaxPriorityFeePerGas param.Field[string] `json:"max_priority_fee_per_gas"`
	// The nonce value in hex string format.
	Nonce param.Field[string] `json:"nonce"`
	// The paymaster and data value in hex string format.
	PaymasterAndData param.Field[string] `json:"paymaster_and_data"`
	// The pre verification gas value in hex string format.
	PreVerificationGas param.Field[string] `json:"pre_verification_gas"`
	// The sender address of the operation in hex string format
	Sender param.Field[string] `json:"sender"`
	// The signature value in hex string format.
	Signature param.Field[string] `json:"signature"`
	// The verification gas limit value in hex string format.
	VerificationGasLimit param.Field[string] `json:"verification_gas_limit"`
}

func (r UserOperationDataOperationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserOperationDataOperationParam) implementsUserOperationDataOperationUnionParam() {}

// The operation parameters of the user operation request
//
// Satisfied by [UserOperationV6Param], [UserOperationV7Param],
// [UserOperationDataOperationParam].
type UserOperationDataOperationUnionParam interface {
	implementsUserOperationDataOperationUnionParam()
}

type UserOperationRequestParam struct {
	// The chain name or chain ID
	Chain param.Field[TransactionScanSupportedChain] `json:"chain" api:"required"`
	// The user operation request that was received by the wallet
	Data param.Field[UserOperationDataParam] `json:"data" api:"required"`
	// Additional context for the scan (e.g., dapp URL/domain, integration source).
	// Used to enrich results and reduce false positives/negatives.
	Metadata param.Field[UserOperationRequestMetadataUnionParam] `json:"metadata" api:"required"`
	// The address of the account (wallet) sending the request in hex string format
	AccountAddress param.Field[string] `json:"account_address"`
	// The relative block for the block validation. Can be "latest" or a block number.
	Block param.Field[UserOperationRequestBlockUnionParam] `json:"block"`
	// List of one or more of options for the desired output. "simulation" - include
	// simulation output in your response. "validation" - include security validation
	// of the transaction in your response. "gas_estimation" - include gas estimation
	// result in your response. Default is ["validation"]
	Options param.Field[[]UserOperationRequestOption] `json:"options"`
	// For simulations, determine whether to calculate missing balances in user
	// operations.
	ShouldCalculateMissingBalance param.Field[bool] `json:"should_calculate_missing_balance"`
	// Simulate transactions using gas estimation result. This requires
	// "gas_estimation" option to be enabled.
	SimulateWithEstimatedGas param.Field[bool] `json:"simulate_with_estimated_gas"`
	// Override the state of the chain. This is useful for testing purposes.
	StateOverride param.Field[map[string]UserOperationRequestStateOverrideParam] `json:"state_override"`
}

func (r UserOperationRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Additional context for the scan (e.g., dapp URL/domain, integration source).
// Used to enrich results and reduce false positives/negatives.
type UserOperationRequestMetadataParam struct {
	// The full URL of the DApp or website that initiated the transaction, for
	// cross-reference. Must use the https or http scheme and contain a valid hostname.
	// Cannot contain JSON, braces, or other embedded data structures.
	Domain param.Field[string] `json:"domain"`
	// Indicates that the transaction was not initiated by a dapp.
	NonDapp param.Field[bool] `json:"non_dapp"`
}

func (r UserOperationRequestMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserOperationRequestMetadataParam) implementsUserOperationRequestMetadataUnionParam() {}

// Additional context for the scan (e.g., dapp URL/domain, integration source).
// Used to enrich results and reduce false positives/negatives.
//
// Satisfied by [UserOperationRequestMetadataRoutersEvmModelsMetadataNonDappParam],
// [UserOperationRequestMetadataRoutersEvmModelsMetadataDappParam],
// [UserOperationRequestMetadataParam].
type UserOperationRequestMetadataUnionParam interface {
	implementsUserOperationRequestMetadataUnionParam()
}

type UserOperationRequestMetadataRoutersEvmModelsMetadataNonDappParam struct {
	// Indicates that the transaction was not initiated by a dapp.
	NonDapp param.Field[UserOperationRequestMetadataRoutersEvmModelsMetadataNonDappNonDapp] `json:"non_dapp"`
}

func (r UserOperationRequestMetadataRoutersEvmModelsMetadataNonDappParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserOperationRequestMetadataRoutersEvmModelsMetadataNonDappParam) implementsUserOperationRequestMetadataUnionParam() {
}

// Indicates that the transaction was not initiated by a dapp.
type UserOperationRequestMetadataRoutersEvmModelsMetadataNonDappNonDapp bool

const (
	UserOperationRequestMetadataRoutersEvmModelsMetadataNonDappNonDappTrue UserOperationRequestMetadataRoutersEvmModelsMetadataNonDappNonDapp = true
)

func (r UserOperationRequestMetadataRoutersEvmModelsMetadataNonDappNonDapp) IsKnown() bool {
	switch r {
	case UserOperationRequestMetadataRoutersEvmModelsMetadataNonDappNonDappTrue:
		return true
	}
	return false
}

type UserOperationRequestMetadataRoutersEvmModelsMetadataDappParam struct {
	// The full URL of the DApp or website that initiated the transaction, for
	// cross-reference. Must use the https or http scheme and contain a valid hostname.
	// Cannot contain JSON, braces, or other embedded data structures.
	Domain param.Field[string] `json:"domain" api:"required"`
	// Indicates that the transaction was not initiated by a dapp. Use false when the
	// transaction is from a dapp.
	NonDapp param.Field[bool] `json:"non_dapp"`
}

func (r UserOperationRequestMetadataRoutersEvmModelsMetadataDappParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserOperationRequestMetadataRoutersEvmModelsMetadataDappParam) implementsUserOperationRequestMetadataUnionParam() {
}

// The relative block for the block validation. Can be "latest" or a block number.
//
// Satisfied by [shared.UnionInt], [shared.UnionString].
type UserOperationRequestBlockUnionParam interface {
	ImplementsUserOperationRequestBlockUnionParam()
}

// Response sections to include (e.g., validation, simulation, gas estimation,
// events).
type UserOperationRequestOption string

const (
	UserOperationRequestOptionValidation    UserOperationRequestOption = "validation"
	UserOperationRequestOptionSimulation    UserOperationRequestOption = "simulation"
	UserOperationRequestOptionGasEstimation UserOperationRequestOption = "gas_estimation"
	UserOperationRequestOptionEvents        UserOperationRequestOption = "events"
)

func (r UserOperationRequestOption) IsKnown() bool {
	switch r {
	case UserOperationRequestOptionValidation, UserOperationRequestOptionSimulation, UserOperationRequestOptionGasEstimation, UserOperationRequestOptionEvents:
		return true
	}
	return false
}

type UserOperationRequestStateOverrideParam struct {
	// Fake balance to set for the account before executing the call.
	Balance param.Field[string] `json:"balance"`
	// Fake EVM bytecode to inject into the account before executing the call.
	Code param.Field[string] `json:"code"`
	// Moves precompile to given address
	MovePrecompileToAddress param.Field[string] `json:"movePrecompileToAddress"`
	// Fake nonce to set for the account before executing the call.
	Nonce param.Field[string] `json:"nonce"`
	// Fake key-value mapping to override all slots in the account storage before
	// executing the call.
	State param.Field[map[string]string] `json:"state"`
	// Fake key-value mapping to override individual slots in the account storage
	// before executing the call.
	StateDiff param.Field[map[string]string] `json:"stateDiff"`
}

func (r UserOperationRequestStateOverrideParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserOperationV6Param struct {
	// The call data value in hex string format.
	CallData param.Field[string] `json:"call_data"`
	// The call gas limit value in hex string format.
	CallGasLimit param.Field[string] `json:"call_gas_limit"`
	// The EIP-7702 authorization tuple for the user operation (optional)
	Eip7702Auth param.Field[AuthorizationParam] `json:"eip7702_auth"`
	// The init code value in hex string format.
	InitCode param.Field[string] `json:"init_code"`
	// The max fee per gas value in hex string format.
	MaxFeePerGas param.Field[string] `json:"max_fee_per_gas"`
	// The max priority fee per gas value in hex string format.
	MaxPriorityFeePerGas param.Field[string] `json:"max_priority_fee_per_gas"`
	// The nonce value in hex string format.
	Nonce param.Field[string] `json:"nonce"`
	// The paymaster and data value in hex string format.
	PaymasterAndData param.Field[string] `json:"paymaster_and_data"`
	// The pre verification gas value in hex string format.
	PreVerificationGas param.Field[string] `json:"pre_verification_gas"`
	// The sender address of the operation in hex string format
	Sender param.Field[string] `json:"sender"`
	// The signature value in hex string format.
	Signature param.Field[string] `json:"signature"`
	// The verification gas limit value in hex string format.
	VerificationGasLimit param.Field[string] `json:"verification_gas_limit"`
}

func (r UserOperationV6Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserOperationV6Param) implementsUserOperationDataOperationUnionParam() {}

type UserOperationV6GasEstimation struct {
	CallGasEstimate            string                             `json:"call_gas_estimate" api:"required"`
	PreVerificationGasEstimate string                             `json:"pre_verification_gas_estimate" api:"required"`
	Status                     UserOperationV6GasEstimationStatus `json:"status" api:"required"`
	VerificationGasEstimate    string                             `json:"verification_gas_estimate" api:"required"`
	JSON                       userOperationV6GasEstimationJSON   `json:"-"`
}

// userOperationV6GasEstimationJSON contains the JSON metadata for the struct
// [UserOperationV6GasEstimation]
type userOperationV6GasEstimationJSON struct {
	CallGasEstimate            apijson.Field
	PreVerificationGasEstimate apijson.Field
	Status                     apijson.Field
	VerificationGasEstimate    apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *UserOperationV6GasEstimation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOperationV6GasEstimationJSON) RawJSON() string {
	return r.raw
}

func (r UserOperationV6GasEstimation) implementsEvmJsonRpcScanResponseUserOperationGasEstimation() {}

func (r UserOperationV6GasEstimation) implementsEvmTransactionScanResponseUserOperationGasEstimation() {
}

func (r UserOperationV6GasEstimation) implementsEvmTransactionBulkScanResponseUserOperationGasEstimation() {
}

func (r UserOperationV6GasEstimation) implementsEvmTransactionRawScanResponseUserOperationGasEstimation() {
}

func (r UserOperationV6GasEstimation) implementsEvmUserOperationScanResponseUserOperationGasEstimation() {
}

func (r UserOperationV6GasEstimation) implementsEvmPostTransactionScanResponseUserOperationGasEstimation() {
}

func (r UserOperationV6GasEstimation) implementsEvmPostTransactionBulkScanResponseUserOperationGasEstimation() {
}

type UserOperationV6GasEstimationStatus string

const (
	UserOperationV6GasEstimationStatusSuccess UserOperationV6GasEstimationStatus = "Success"
)

func (r UserOperationV6GasEstimationStatus) IsKnown() bool {
	switch r {
	case UserOperationV6GasEstimationStatusSuccess:
		return true
	}
	return false
}

type UserOperationV7Param struct {
	// The account gas limits value in hex string format.
	AccountGasLimits param.Field[string] `json:"account_gas_limits"`
	// The call data value in hex string format.
	CallData param.Field[string] `json:"call_data"`
	// The EIP-7702 authorization tuple for the user operation (optional)
	Eip7702Auth param.Field[AuthorizationParam] `json:"eip7702_auth"`
	// The gas fees value in hex string format.
	GasFees param.Field[string] `json:"gas_fees"`
	// The init code value in hex string format.
	InitCode param.Field[string] `json:"init_code"`
	// The nonce value in hex string format.
	Nonce param.Field[string] `json:"nonce"`
	// The paymaster and data value in hex string format.
	PaymasterAndData param.Field[string] `json:"paymaster_and_data"`
	// The pre verification gas value in hex string format.
	PreVerificationGas param.Field[string] `json:"pre_verification_gas"`
	// The sender address of the operation in hex string format
	Sender param.Field[string] `json:"sender"`
	// The signature value in hex string format.
	Signature param.Field[string] `json:"signature"`
}

func (r UserOperationV7Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserOperationV7Param) implementsUserOperationDataOperationUnionParam() {}

type UserOperationV7GasEstimation struct {
	CallGasEstimate                  string                             `json:"call_gas_estimate" api:"required"`
	PaymasterVerificationGasEstimate string                             `json:"paymaster_verification_gas_estimate" api:"required"`
	PreVerificationGasEstimate       string                             `json:"pre_verification_gas_estimate" api:"required"`
	Status                           UserOperationV7GasEstimationStatus `json:"status" api:"required"`
	VerificationGasEstimate          string                             `json:"verification_gas_estimate" api:"required"`
	JSON                             userOperationV7GasEstimationJSON   `json:"-"`
}

// userOperationV7GasEstimationJSON contains the JSON metadata for the struct
// [UserOperationV7GasEstimation]
type userOperationV7GasEstimationJSON struct {
	CallGasEstimate                  apijson.Field
	PaymasterVerificationGasEstimate apijson.Field
	PreVerificationGasEstimate       apijson.Field
	Status                           apijson.Field
	VerificationGasEstimate          apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r *UserOperationV7GasEstimation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOperationV7GasEstimationJSON) RawJSON() string {
	return r.raw
}

func (r UserOperationV7GasEstimation) implementsEvmJsonRpcScanResponseUserOperationGasEstimation() {}

func (r UserOperationV7GasEstimation) implementsEvmTransactionScanResponseUserOperationGasEstimation() {
}

func (r UserOperationV7GasEstimation) implementsEvmTransactionBulkScanResponseUserOperationGasEstimation() {
}

func (r UserOperationV7GasEstimation) implementsEvmTransactionRawScanResponseUserOperationGasEstimation() {
}

func (r UserOperationV7GasEstimation) implementsEvmUserOperationScanResponseUserOperationGasEstimation() {
}

func (r UserOperationV7GasEstimation) implementsEvmPostTransactionScanResponseUserOperationGasEstimation() {
}

func (r UserOperationV7GasEstimation) implementsEvmPostTransactionBulkScanResponseUserOperationGasEstimation() {
}

type UserOperationV7GasEstimationStatus string

const (
	UserOperationV7GasEstimationStatusSuccess UserOperationV7GasEstimationStatus = "Success"
)

func (r UserOperationV7GasEstimationStatus) IsKnown() bool {
	switch r {
	case UserOperationV7GasEstimationStatusSuccess:
		return true
	}
	return false
}

type ValidateAddressParam struct {
	// The address to validate.
	Address param.Field[string] `json:"address" api:"required"`
	// The chain name
	Chain param.Field[TransactionScanSupportedChain] `json:"chain" api:"required"`
	// Additional context for the scan (e.g., dapp URL/domain, integration source).
	// Used to enrich results and reduce false positives/negatives.
	Metadata param.Field[ValidateAddressMetadataUnionParam] `json:"metadata" api:"required"`
}

func (r ValidateAddressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Additional context for the scan (e.g., dapp URL/domain, integration source).
// Used to enrich results and reduce false positives/negatives.
type ValidateAddressMetadataParam struct {
	// The full URL of the DApp or website that initiated the transaction, for
	// cross-reference. Must use the https or http scheme and contain a valid hostname.
	// Cannot contain JSON, braces, or other embedded data structures.
	Domain param.Field[string] `json:"domain"`
	// Indicates that the transaction was not initiated by a dapp.
	NonDapp param.Field[bool] `json:"non_dapp"`
}

func (r ValidateAddressMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ValidateAddressMetadataParam) implementsValidateAddressMetadataUnionParam() {}

// Additional context for the scan (e.g., dapp URL/domain, integration source).
// Used to enrich results and reduce false positives/negatives.
//
// Satisfied by [ValidateAddressMetadataRoutersEvmModelsMetadataNonDappParam],
// [ValidateAddressMetadataRoutersEvmModelsMetadataDappParam],
// [ValidateAddressMetadataParam].
type ValidateAddressMetadataUnionParam interface {
	implementsValidateAddressMetadataUnionParam()
}

type ValidateAddressMetadataRoutersEvmModelsMetadataNonDappParam struct {
	// Indicates that the transaction was not initiated by a dapp.
	NonDapp param.Field[ValidateAddressMetadataRoutersEvmModelsMetadataNonDappNonDapp] `json:"non_dapp"`
}

func (r ValidateAddressMetadataRoutersEvmModelsMetadataNonDappParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ValidateAddressMetadataRoutersEvmModelsMetadataNonDappParam) implementsValidateAddressMetadataUnionParam() {
}

// Indicates that the transaction was not initiated by a dapp.
type ValidateAddressMetadataRoutersEvmModelsMetadataNonDappNonDapp bool

const (
	ValidateAddressMetadataRoutersEvmModelsMetadataNonDappNonDappTrue ValidateAddressMetadataRoutersEvmModelsMetadataNonDappNonDapp = true
)

func (r ValidateAddressMetadataRoutersEvmModelsMetadataNonDappNonDapp) IsKnown() bool {
	switch r {
	case ValidateAddressMetadataRoutersEvmModelsMetadataNonDappNonDappTrue:
		return true
	}
	return false
}

type ValidateAddressMetadataRoutersEvmModelsMetadataDappParam struct {
	// The full URL of the DApp or website that initiated the transaction, for
	// cross-reference. Must use the https or http scheme and contain a valid hostname.
	// Cannot contain JSON, braces, or other embedded data structures.
	Domain param.Field[string] `json:"domain" api:"required"`
	// Indicates that the transaction was not initiated by a dapp. Use false when the
	// transaction is from a dapp.
	NonDapp param.Field[bool] `json:"non_dapp"`
}

func (r ValidateAddressMetadataRoutersEvmModelsMetadataDappParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ValidateAddressMetadataRoutersEvmModelsMetadataDappParam) implementsValidateAddressMetadataUnionParam() {
}

type ValidateBulkAddressesParam struct {
	// List of addresses to validate.
	Addresses param.Field[[]string] `json:"addresses" api:"required"`
	// The chain name
	Chain param.Field[TransactionScanSupportedChain] `json:"chain" api:"required"`
	// Additional context for the scan (e.g., dapp URL/domain, integration source).
	// Used to enrich results and reduce false positives/negatives.
	Metadata param.Field[ValidateBulkAddressesMetadataUnionParam] `json:"metadata" api:"required"`
}

func (r ValidateBulkAddressesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Additional context for the scan (e.g., dapp URL/domain, integration source).
// Used to enrich results and reduce false positives/negatives.
type ValidateBulkAddressesMetadataParam struct {
	// The full URL of the DApp or website that initiated the transaction, for
	// cross-reference. Must use the https or http scheme and contain a valid hostname.
	// Cannot contain JSON, braces, or other embedded data structures.
	Domain param.Field[string] `json:"domain"`
	// Indicates that the transaction was not initiated by a dapp.
	NonDapp param.Field[bool] `json:"non_dapp"`
}

func (r ValidateBulkAddressesMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ValidateBulkAddressesMetadataParam) implementsValidateBulkAddressesMetadataUnionParam() {}

// Additional context for the scan (e.g., dapp URL/domain, integration source).
// Used to enrich results and reduce false positives/negatives.
//
// Satisfied by
// [ValidateBulkAddressesMetadataRoutersEvmModelsMetadataNonDappParam],
// [ValidateBulkAddressesMetadataRoutersEvmModelsMetadataDappParam],
// [ValidateBulkAddressesMetadataParam].
type ValidateBulkAddressesMetadataUnionParam interface {
	implementsValidateBulkAddressesMetadataUnionParam()
}

type ValidateBulkAddressesMetadataRoutersEvmModelsMetadataNonDappParam struct {
	// Indicates that the transaction was not initiated by a dapp.
	NonDapp param.Field[ValidateBulkAddressesMetadataRoutersEvmModelsMetadataNonDappNonDapp] `json:"non_dapp"`
}

func (r ValidateBulkAddressesMetadataRoutersEvmModelsMetadataNonDappParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ValidateBulkAddressesMetadataRoutersEvmModelsMetadataNonDappParam) implementsValidateBulkAddressesMetadataUnionParam() {
}

// Indicates that the transaction was not initiated by a dapp.
type ValidateBulkAddressesMetadataRoutersEvmModelsMetadataNonDappNonDapp bool

const (
	ValidateBulkAddressesMetadataRoutersEvmModelsMetadataNonDappNonDappTrue ValidateBulkAddressesMetadataRoutersEvmModelsMetadataNonDappNonDapp = true
)

func (r ValidateBulkAddressesMetadataRoutersEvmModelsMetadataNonDappNonDapp) IsKnown() bool {
	switch r {
	case ValidateBulkAddressesMetadataRoutersEvmModelsMetadataNonDappNonDappTrue:
		return true
	}
	return false
}

type ValidateBulkAddressesMetadataRoutersEvmModelsMetadataDappParam struct {
	// The full URL of the DApp or website that initiated the transaction, for
	// cross-reference. Must use the https or http scheme and contain a valid hostname.
	// Cannot contain JSON, braces, or other embedded data structures.
	Domain param.Field[string] `json:"domain" api:"required"`
	// Indicates that the transaction was not initiated by a dapp. Use false when the
	// transaction is from a dapp.
	NonDapp param.Field[bool] `json:"non_dapp"`
}

func (r ValidateBulkAddressesMetadataRoutersEvmModelsMetadataDappParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ValidateBulkAddressesMetadataRoutersEvmModelsMetadataDappParam) implementsValidateBulkAddressesMetadataUnionParam() {
}

type ValidateBulkExtendedAddressesRequestParam struct {
	// List of addresses to validate.
	Addresses param.Field[[]string] `json:"addresses" api:"required"`
	// The chain name
	Chain param.Field[TransactionScanSupportedChain] `json:"chain" api:"required"`
	// Additional context for the scan (e.g., integration source, user/account
	// details). Used to enrich results and reduce false positives/negatives.
	Metadata param.Field[ValidateBulkExtendedAddressesRequestMetadataParam] `json:"metadata" api:"required"`
}

func (r ValidateBulkExtendedAddressesRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Additional context for the scan (e.g., integration source, user/account
// details). Used to enrich results and reduce false positives/negatives.
type ValidateBulkExtendedAddressesRequestMetadataParam struct {
	Account    param.Field[ValidateBulkExtendedAddressesRequestMetadataAccountParam]    `json:"account" api:"required"`
	Connection param.Field[ValidateBulkExtendedAddressesRequestMetadataConnectionParam] `json:"connection" api:"required"`
}

func (r ValidateBulkExtendedAddressesRequestMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ValidateBulkExtendedAddressesRequestMetadataAccountParam struct {
	AccountID       param.Field[string]    `json:"account_id" api:"required"`
	UserCountryCode param.Field[string]    `json:"user_country_code" api:"required"`
	AgeInYears      param.Field[int64]     `json:"age_in_years"`
	Created         param.Field[time.Time] `json:"created" format:"date-time"`
}

func (r ValidateBulkExtendedAddressesRequestMetadataAccountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ValidateBulkExtendedAddressesRequestMetadataConnectionParam struct {
	IPAddress param.Field[string] `json:"ip_address" api:"required"`
	UserAgent param.Field[string] `json:"user_agent" api:"required"`
}

func (r ValidateBulkExtendedAddressesRequestMetadataConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ValidateBulkExtendedAddressesResponse struct {
	Results []ValidateBulkExtendedAddressesResponseResult `json:"results" api:"required"`
	JSON    validateBulkExtendedAddressesResponseJSON     `json:"-"`
}

// validateBulkExtendedAddressesResponseJSON contains the JSON metadata for the
// struct [ValidateBulkExtendedAddressesResponse]
type validateBulkExtendedAddressesResponseJSON struct {
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateBulkExtendedAddressesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r validateBulkExtendedAddressesResponseJSON) RawJSON() string {
	return r.raw
}

type ValidateBulkExtendedAddressesResponseResult struct {
	// Address to validate.
	Address string `json:"address" api:"required"`
	// Label of the address.
	Label string `json:"label" api:"required"`
	// Validation result.
	Validation ValidateBulkExtendedAddressesResponseResultsValidation `json:"validation" api:"required"`
	JSON       validateBulkExtendedAddressesResponseResultJSON        `json:"-"`
}

// validateBulkExtendedAddressesResponseResultJSON contains the JSON metadata for
// the struct [ValidateBulkExtendedAddressesResponseResult]
type validateBulkExtendedAddressesResponseResultJSON struct {
	Address     apijson.Field
	Label       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateBulkExtendedAddressesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r validateBulkExtendedAddressesResponseResultJSON) RawJSON() string {
	return r.raw
}

// Validation result.
type ValidateBulkExtendedAddressesResponseResultsValidation struct {
	// Overall validation outcome for the scan.
	ResultType ValidateBulkExtendedAddressesResponseResultsValidationResultType `json:"result_type" api:"required"`
	// A list of textual features about this address that can be presented to the user.
	Features []ValidateBulkExtendedAddressesResponseResultsValidationFeature `json:"features"`
	JSON     validateBulkExtendedAddressesResponseResultsValidationJSON      `json:"-"`
}

// validateBulkExtendedAddressesResponseResultsValidationJSON contains the JSON
// metadata for the struct [ValidateBulkExtendedAddressesResponseResultsValidation]
type validateBulkExtendedAddressesResponseResultsValidationJSON struct {
	ResultType  apijson.Field
	Features    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateBulkExtendedAddressesResponseResultsValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r validateBulkExtendedAddressesResponseResultsValidationJSON) RawJSON() string {
	return r.raw
}

// Overall validation outcome for the scan.
type ValidateBulkExtendedAddressesResponseResultsValidationResultType string

const (
	ValidateBulkExtendedAddressesResponseResultsValidationResultTypeMalicious ValidateBulkExtendedAddressesResponseResultsValidationResultType = "Malicious"
	ValidateBulkExtendedAddressesResponseResultsValidationResultTypeWarning   ValidateBulkExtendedAddressesResponseResultsValidationResultType = "Warning"
	ValidateBulkExtendedAddressesResponseResultsValidationResultTypeBenign    ValidateBulkExtendedAddressesResponseResultsValidationResultType = "Benign"
	ValidateBulkExtendedAddressesResponseResultsValidationResultTypeError     ValidateBulkExtendedAddressesResponseResultsValidationResultType = "Error"
)

func (r ValidateBulkExtendedAddressesResponseResultsValidationResultType) IsKnown() bool {
	switch r {
	case ValidateBulkExtendedAddressesResponseResultsValidationResultTypeMalicious, ValidateBulkExtendedAddressesResponseResultsValidationResultTypeWarning, ValidateBulkExtendedAddressesResponseResultsValidationResultTypeBenign, ValidateBulkExtendedAddressesResponseResultsValidationResultTypeError:
		return true
	}
	return false
}

type ValidateBulkExtendedAddressesResponseResultsValidationFeature struct {
	// Description of the feature
	Description string `json:"description" api:"required"`
	// Feature identifier
	FeatureID ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID `json:"feature_id" api:"required"`
	// Type of the feature
	Type ValidateBulkExtendedAddressesResponseResultsValidationFeaturesType `json:"type" api:"required"`
	JSON validateBulkExtendedAddressesResponseResultsValidationFeatureJSON  `json:"-"`
}

// validateBulkExtendedAddressesResponseResultsValidationFeatureJSON contains the
// JSON metadata for the struct
// [ValidateBulkExtendedAddressesResponseResultsValidationFeature]
type validateBulkExtendedAddressesResponseResultsValidationFeatureJSON struct {
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateBulkExtendedAddressesResponseResultsValidationFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r validateBulkExtendedAddressesResponseResultsValidationFeatureJSON) RawJSON() string {
	return r.raw
}

// Feature identifier
type ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID string

const (
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDVerifiedContract               ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "VERIFIED_CONTRACT"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDUnverifiedContract             ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "UNVERIFIED_CONTRACT"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHighTradeVolume                ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "HIGH_TRADE_VOLUME"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDMarketPlaceSalesHistory        ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "MARKET_PLACE_SALES_HISTORY"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHighReputationToken            ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "HIGH_REPUTATION_TOKEN"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDOnchainActivityValidator       ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "ONCHAIN_ACTIVITY_VALIDATOR"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDStaticCodeSignature            ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "STATIC_CODE_SIGNATURE"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDKnownMalicious                 ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "KNOWN_MALICIOUS"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDIsEoa                          ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "IS_EOA"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDIsContract                     ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "IS_CONTRACT"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDErc20Contract                  ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "ERC20_CONTRACT"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDTrustedContract                ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "TRUSTED_CONTRACT"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDBenignCreator                  ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "BENIGN_CREATOR"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDMetadata                       ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "METADATA"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDAirdropPattern                 ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "AIRDROP_PATTERN"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDImpersonator                   ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "IMPERSONATOR"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDInorganicVolume                ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "INORGANIC_VOLUME"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDDynamicAnalysis                ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "DYNAMIC_ANALYSIS"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDConcentratedSupplyDistribution ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "CONCENTRATED_SUPPLY_DISTRIBUTION"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHoneypot                       ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "HONEYPOT"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDInsufficientLockedLiquidity    ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "INSUFFICIENT_LOCKED_LIQUIDITY"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDUnstableTokenPrice             ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "UNSTABLE_TOKEN_PRICE"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDRugpull                        ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "RUGPULL"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDWashTrading                    ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "WASH_TRADING"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDConsumerOverride               ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "CONSUMER_OVERRIDE"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDInappropriateContent           ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "INAPPROPRIATE_CONTENT"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHighTransferFee                ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "HIGH_TRANSFER_FEE"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHighBuyFee                     ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "HIGH_BUY_FEE"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHighSellFee                    ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "HIGH_SELL_FEE"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDUnsellableToken                ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "UNSELLABLE_TOKEN"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDIsMintable                     ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "IS_MINTABLE"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDRebaseToken                    ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "REBASE_TOKEN"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDLiquidStakingToken             ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "LIQUID_STAKING_TOKEN"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDModifiableTaxes                ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "MODIFIABLE_TAXES"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDCanBlacklist                   ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "CAN_BLACKLIST"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDCanWhitelist                   ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "CAN_WHITELIST"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHasTradingCooldown             ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "HAS_TRADING_COOLDOWN"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDExternalFunctions              ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "EXTERNAL_FUNCTIONS"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHiddenOwner                    ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "HIDDEN_OWNER"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDTransferPauseable              ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "TRANSFER_PAUSEABLE"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDOwnershipRenounced             ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "OWNERSHIP_RENOUNCED"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDOwnerCanChangeBalance          ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "OWNER_CAN_CHANGE_BALANCE"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDProxyContract                  ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "PROXY_CONTRACT"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDSimilarMaliciousContract       ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "SIMILAR_MALICIOUS_CONTRACT"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDImpersonatorSensitiveAsset     ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "IMPERSONATOR_SENSITIVE_ASSET"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDImpersonatorHighConfidence     ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "IMPERSONATOR_HIGH_CONFIDENCE"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDImpersonatorMediumConfidence   ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "IMPERSONATOR_MEDIUM_CONFIDENCE"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDImpersonatorLowConfidence      ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "IMPERSONATOR_LOW_CONFIDENCE"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDImpersonationProtected         ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "IMPERSONATION_PROTECTED"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDFakeVolume                     ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "FAKE_VOLUME"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHiddenSupplyByKeyHolder        ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "HIDDEN_SUPPLY_BY_KEY_HOLDER"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDFakeTradeMakerCount            ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "FAKE_TRADE_MAKER_COUNT"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDTransferFromReverts            ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "TRANSFER_FROM_REVERTS"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDOffensiveTokenMetadata         ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "OFFENSIVE_TOKEN_METADATA"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDListedOnCentralizedExchange    ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "LISTED_ON_CENTRALIZED_EXCHANGE"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDSanctionedCreator              ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "SANCTIONED_CREATOR"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDSpamText                       ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "SPAM_TEXT"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDBondingCurveToken              ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "BONDING_CURVE_TOKEN"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHeavilySniped                  ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "HEAVILY_SNIPED"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDSolanaToken2022                ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "SOLANA_TOKEN_2022"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDPostDump                       ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "POST_DUMP"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDDexPaid                        ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "DEX_PAID"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDLowReputationCreator           ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "LOW_REPUTATION_CREATOR"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDSnipeAtMint                    ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "SNIPE_AT_MINT"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDTransferHookEnabled            ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "TRANSFER_HOOK_ENABLED"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDConfidentialTransfersEnabled   ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "CONFIDENTIAL_TRANSFERS_ENABLED"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDNonTranserable                 ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "NON_TRANSERABLE"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDTokenBackdoor                  ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "TOKEN_BACKDOOR"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDCreatedViaLaunchpad            ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "CREATED_VIA_LAUNCHPAD"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDCompromisedToken               ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "COMPROMISED_TOKEN"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDLongFundTrail                  ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "LONG_FUND_TRAIL"
)

func (r ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID) IsKnown() bool {
	switch r {
	case ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDVerifiedContract, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDUnverifiedContract, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHighTradeVolume, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDMarketPlaceSalesHistory, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHighReputationToken, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDOnchainActivityValidator, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDStaticCodeSignature, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDKnownMalicious, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDIsEoa, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDIsContract, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDErc20Contract, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDTrustedContract, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDBenignCreator, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDMetadata, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDAirdropPattern, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDImpersonator, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDInorganicVolume, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDDynamicAnalysis, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDConcentratedSupplyDistribution, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHoneypot, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDInsufficientLockedLiquidity, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDUnstableTokenPrice, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDRugpull, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDWashTrading, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDConsumerOverride, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDInappropriateContent, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHighTransferFee, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHighBuyFee, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHighSellFee, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDUnsellableToken, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDIsMintable, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDRebaseToken, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDLiquidStakingToken, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDModifiableTaxes, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDCanBlacklist, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDCanWhitelist, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHasTradingCooldown, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDExternalFunctions, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHiddenOwner, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDTransferPauseable, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDOwnershipRenounced, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDOwnerCanChangeBalance, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDProxyContract, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDSimilarMaliciousContract, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDImpersonatorSensitiveAsset, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDImpersonatorHighConfidence, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDImpersonatorMediumConfidence, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDImpersonatorLowConfidence, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDImpersonationProtected, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDFakeVolume, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHiddenSupplyByKeyHolder, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDFakeTradeMakerCount, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDTransferFromReverts, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDOffensiveTokenMetadata, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDListedOnCentralizedExchange, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDSanctionedCreator, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDSpamText, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDBondingCurveToken, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHeavilySniped, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDSolanaToken2022, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDPostDump, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDDexPaid, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDLowReputationCreator, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDSnipeAtMint, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDTransferHookEnabled, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDConfidentialTransfersEnabled, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDNonTranserable, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDTokenBackdoor, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDCreatedViaLaunchpad, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDCompromisedToken, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDLongFundTrail:
		return true
	}
	return false
}

// Type of the feature
type ValidateBulkExtendedAddressesResponseResultsValidationFeaturesType string

const (
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesTypeBenign    ValidateBulkExtendedAddressesResponseResultsValidationFeaturesType = "Benign"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesTypeInfo      ValidateBulkExtendedAddressesResponseResultsValidationFeaturesType = "Info"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesTypeWarning   ValidateBulkExtendedAddressesResponseResultsValidationFeaturesType = "Warning"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesTypeMalicious ValidateBulkExtendedAddressesResponseResultsValidationFeaturesType = "Malicious"
)

func (r ValidateBulkExtendedAddressesResponseResultsValidationFeaturesType) IsKnown() bool {
	switch r {
	case ValidateBulkExtendedAddressesResponseResultsValidationFeaturesTypeBenign, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesTypeInfo, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesTypeWarning, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesTypeMalicious:
		return true
	}
	return false
}
