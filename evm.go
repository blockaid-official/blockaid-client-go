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
	// An enumeration.
	ResultType AddressValidationResultType `json:"result_type,required"`
	// An error message if the validation failed.
	Error string `json:"error"`
	// A list of textual features about this transaction that can be presented to the
	// user.
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

// An enumeration.
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

// A list of textual features about this transaction that can be presented to the
// user.
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

type AddressValidationFeaturesArray []string

func (r AddressValidationFeaturesArray) implementsAddressValidationFeaturesUnion() {}

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
	// Unique identifier for the account
	AccountID param.Field[string] `json:"account_id,required"`
	// Timestamp when the account was created
	AccountCreationTimestamp param.Field[time.Time] `json:"account_creation_timestamp" format:"date-time"`
	// Age of the user in years
	UserAge param.Field[int64] `json:"user_age"`
	// ISO country code of the user's location
	UserCountryCode param.Field[string] `json:"user_country_code"`
}

func (r MetadataParamAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Connection metadata including user agent and IP information
type MetadataParamConnection struct {
	// IP address of the customer making the request
	IPAddress param.Field[string] `json:"ip_address,required" format:"ipvanyaddress"`
	// User agent string from the client's browser or application
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
)

func (r TokenScanSupportedChain) IsKnown() bool {
	switch r {
	case TokenScanSupportedChainArbitrum, TokenScanSupportedChainAvalanche, TokenScanSupportedChainBase, TokenScanSupportedChainBsc, TokenScanSupportedChainEthereum, TokenScanSupportedChainOptimism, TokenScanSupportedChainPolygon, TokenScanSupportedChainZora, TokenScanSupportedChainSolana, TokenScanSupportedChainStarknet, TokenScanSupportedChainStarknetSepolia, TokenScanSupportedChainStellar, TokenScanSupportedChainLinea, TokenScanSupportedChainDegen, TokenScanSupportedChainZksync, TokenScanSupportedChainScroll, TokenScanSupportedChainBlast, TokenScanSupportedChainSoneiumMinato, TokenScanSupportedChainBaseSepolia, TokenScanSupportedChainBitcoin, TokenScanSupportedChainAbstract, TokenScanSupportedChainSoneium, TokenScanSupportedChainInk, TokenScanSupportedChainZeroNetwork, TokenScanSupportedChainBerachain, TokenScanSupportedChainUnichain, TokenScanSupportedChainRonin, TokenScanSupportedChainSui, TokenScanSupportedChainHedera, TokenScanSupportedChainHyperevm, TokenScanSupportedChainXlayer, TokenScanSupportedChainMonad:
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
	TransactionScanSupportedChainKatana                TransactionScanSupportedChain = "katana"
	TransactionScanSupportedChainPlume                 TransactionScanSupportedChain = "plume"
	TransactionScanSupportedChainXlayer                TransactionScanSupportedChain = "xlayer"
	TransactionScanSupportedChainMonad                 TransactionScanSupportedChain = "monad"
	TransactionScanSupportedChainMonadTestnet          TransactionScanSupportedChain = "monad-testnet"
)

func (r TransactionScanSupportedChain) IsKnown() bool {
	switch r {
	case TransactionScanSupportedChainArbitrum, TransactionScanSupportedChainAvalanche, TransactionScanSupportedChainBase, TransactionScanSupportedChainBaseSepolia, TransactionScanSupportedChainLordchain, TransactionScanSupportedChainLordchainTestnet, TransactionScanSupportedChainMetacade, TransactionScanSupportedChainMetacadeTestnet, TransactionScanSupportedChainBsc, TransactionScanSupportedChainEthereum, TransactionScanSupportedChainOptimism, TransactionScanSupportedChainPolygon, TransactionScanSupportedChainZksync, TransactionScanSupportedChainZksyncSepolia, TransactionScanSupportedChainZora, TransactionScanSupportedChainLinea, TransactionScanSupportedChainBlast, TransactionScanSupportedChainScroll, TransactionScanSupportedChainEthereumSepolia, TransactionScanSupportedChainDegen, TransactionScanSupportedChainAvalancheFuji, TransactionScanSupportedChainImmutableZkevm, TransactionScanSupportedChainImmutableZkevmTestnet, TransactionScanSupportedChainGnosis, TransactionScanSupportedChainWorldchain, TransactionScanSupportedChainSoneiumMinato, TransactionScanSupportedChainRonin, TransactionScanSupportedChainApechain, TransactionScanSupportedChainZeroNetwork, TransactionScanSupportedChainBerachain, TransactionScanSupportedChainBerachainBartio, TransactionScanSupportedChainInk, TransactionScanSupportedChainInkSepolia, TransactionScanSupportedChainAbstract, TransactionScanSupportedChainAbstractTestnet, TransactionScanSupportedChainSoneium, TransactionScanSupportedChainUnichain, TransactionScanSupportedChainSei, TransactionScanSupportedChainFlowEvm, TransactionScanSupportedChainHyperevm, TransactionScanSupportedChainKatana, TransactionScanSupportedChainPlume, TransactionScanSupportedChainXlayer, TransactionScanSupportedChainMonad, TransactionScanSupportedChainMonadTestnet:
		return true
	}
	return false
}

type ValidateAddressParam struct {
	// The address to validate.
	Address param.Field[string] `json:"address,required"`
	// The chain name
	Chain param.Field[TransactionScanSupportedChain] `json:"chain,required"`
	// Object of additional information to validate against.
	Metadata param.Field[ValidateAddressMetadataUnionParam] `json:"metadata,required"`
}

func (r ValidateAddressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Object of additional information to validate against.
type ValidateAddressMetadataParam struct {
	// The full URL of the DApp or website that initiated the transaction, for
	// cross-reference. Must use the https or http scheme and contain a valid hostname.
	// Cannot contain JSON, braces, or other embedded data structures.
	Domain param.Field[string] `json:"domain"`
	// Indicates that the transaction was not initiated by a dapp.
	NonDapp param.Field[ValidateAddressMetadataNonDapp] `json:"non_dapp"`
}

func (r ValidateAddressMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ValidateAddressMetadataParam) implementsValidateAddressMetadataUnionParam() {}

// Object of additional information to validate against.
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
	Domain param.Field[string] `json:"domain,required"`
}

func (r ValidateAddressMetadataRoutersEvmModelsMetadataDappParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ValidateAddressMetadataRoutersEvmModelsMetadataDappParam) implementsValidateAddressMetadataUnionParam() {
}

// Indicates that the transaction was not initiated by a dapp.
type ValidateAddressMetadataNonDapp bool

const (
	ValidateAddressMetadataNonDappTrue ValidateAddressMetadataNonDapp = true
)

func (r ValidateAddressMetadataNonDapp) IsKnown() bool {
	switch r {
	case ValidateAddressMetadataNonDappTrue:
		return true
	}
	return false
}

type ValidateBulkAddressesParam struct {
	// List of addresses to validate.
	Addresses param.Field[[]string] `json:"addresses,required"`
	// The chain name
	Chain param.Field[TransactionScanSupportedChain] `json:"chain,required"`
	// Object of additional information to validate against.
	Metadata param.Field[ValidateBulkAddressesMetadataUnionParam] `json:"metadata,required"`
}

func (r ValidateBulkAddressesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Object of additional information to validate against.
type ValidateBulkAddressesMetadataParam struct {
	// The full URL of the DApp or website that initiated the transaction, for
	// cross-reference. Must use the https or http scheme and contain a valid hostname.
	// Cannot contain JSON, braces, or other embedded data structures.
	Domain param.Field[string] `json:"domain"`
	// Indicates that the transaction was not initiated by a dapp.
	NonDapp param.Field[ValidateBulkAddressesMetadataNonDapp] `json:"non_dapp"`
}

func (r ValidateBulkAddressesMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ValidateBulkAddressesMetadataParam) implementsValidateBulkAddressesMetadataUnionParam() {}

// Object of additional information to validate against.
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
	Domain param.Field[string] `json:"domain,required"`
}

func (r ValidateBulkAddressesMetadataRoutersEvmModelsMetadataDappParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ValidateBulkAddressesMetadataRoutersEvmModelsMetadataDappParam) implementsValidateBulkAddressesMetadataUnionParam() {
}

// Indicates that the transaction was not initiated by a dapp.
type ValidateBulkAddressesMetadataNonDapp bool

const (
	ValidateBulkAddressesMetadataNonDappTrue ValidateBulkAddressesMetadataNonDapp = true
)

func (r ValidateBulkAddressesMetadataNonDapp) IsKnown() bool {
	switch r {
	case ValidateBulkAddressesMetadataNonDappTrue:
		return true
	}
	return false
}

type ValidateBulkExtendedAddressesRequestParam struct {
	// List of addresses to validate.
	Addresses param.Field[[]string] `json:"addresses,required"`
	// The chain name
	Chain param.Field[TransactionScanSupportedChain] `json:"chain,required"`
	// Object of additional information to validate against.
	Metadata param.Field[ValidateBulkExtendedAddressesRequestMetadataParam] `json:"metadata,required"`
}

func (r ValidateBulkExtendedAddressesRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Object of additional information to validate against.
type ValidateBulkExtendedAddressesRequestMetadataParam struct {
	Account    param.Field[ValidateBulkExtendedAddressesRequestMetadataAccountParam]    `json:"account,required"`
	Connection param.Field[ValidateBulkExtendedAddressesRequestMetadataConnectionParam] `json:"connection,required"`
}

func (r ValidateBulkExtendedAddressesRequestMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ValidateBulkExtendedAddressesRequestMetadataAccountParam struct {
	AccountID       param.Field[string]    `json:"account_id,required"`
	UserCountryCode param.Field[string]    `json:"user_country_code,required"`
	AgeInYears      param.Field[int64]     `json:"age_in_years"`
	Created         param.Field[time.Time] `json:"created" format:"date-time"`
}

func (r ValidateBulkExtendedAddressesRequestMetadataAccountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ValidateBulkExtendedAddressesRequestMetadataConnectionParam struct {
	IPAddress param.Field[string] `json:"ip_address,required"`
	UserAgent param.Field[string] `json:"user_agent,required"`
}

func (r ValidateBulkExtendedAddressesRequestMetadataConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ValidateBulkExtendedAddressesResponse struct {
	Results []ValidateBulkExtendedAddressesResponseResult `json:"results,required"`
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
	Address string `json:"address,required"`
	// Label of the address.
	Label string `json:"label,required"`
	// Validation result.
	Validation ValidateBulkExtendedAddressesResponseResultsValidation `json:"validation,required"`
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
	// An enumeration.
	ResultType ValidateBulkExtendedAddressesResponseResultsValidationResultType `json:"result_type,required"`
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

// An enumeration.
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
	Description string `json:"description,required"`
	// Feature identifier
	FeatureID ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID `json:"feature_id,required"`
	// Type of the feature
	Type ValidateBulkExtendedAddressesResponseResultsValidationFeaturesType `json:"type,required"`
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
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDFakeVolume                     ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "FAKE_VOLUME"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHiddenSupplyByKeyHolder        ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "HIDDEN_SUPPLY_BY_KEY_HOLDER"
	ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDFakeTradeMakerCount            ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID = "FAKE_TRADE_MAKER_COUNT"
)

func (r ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureID) IsKnown() bool {
	switch r {
	case ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDVerifiedContract, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDUnverifiedContract, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHighTradeVolume, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDMarketPlaceSalesHistory, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHighReputationToken, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDOnchainActivityValidator, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDStaticCodeSignature, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDKnownMalicious, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDMetadata, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDAirdropPattern, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDImpersonator, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDInorganicVolume, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDDynamicAnalysis, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDConcentratedSupplyDistribution, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHoneypot, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDInsufficientLockedLiquidity, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDUnstableTokenPrice, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDRugpull, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDWashTrading, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDConsumerOverride, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDInappropriateContent, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHighTransferFee, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHighBuyFee, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHighSellFee, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDUnsellableToken, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDIsMintable, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDRebaseToken, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDLiquidStakingToken, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDModifiableTaxes, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDCanBlacklist, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDCanWhitelist, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHasTradingCooldown, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDExternalFunctions, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHiddenOwner, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDTransferPauseable, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDOwnershipRenounced, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDOwnerCanChangeBalance, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDProxyContract, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDSimilarMaliciousContract, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDFakeVolume, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDHiddenSupplyByKeyHolder, ValidateBulkExtendedAddressesResponseResultsValidationFeaturesFeatureIDFakeTradeMakerCount:
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
