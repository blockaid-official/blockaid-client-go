// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/param"
	"github.com/blockaid-official/blockaid-client-go/internal/requestconfig"
	"github.com/blockaid-official/blockaid-client-go/option"
)

// SuiTransactionService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSuiTransactionService] method instead.
type SuiTransactionService struct {
	Options []option.RequestOption
}

// NewSuiTransactionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSuiTransactionService(opts ...option.RequestOption) (r *SuiTransactionService) {
	r = &SuiTransactionService{}
	r.Options = opts
	return
}

// Get a risk recommendation with plain-language reasons for a Sui transaction.
func (r *SuiTransactionService) Scan(ctx context.Context, body SuiTransactionScanParams, opts ...option.RequestOption) (res *SuiTransactionScanResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/sui/transaction/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type SuiTransactionScanParams struct {
	AccountAddress param.Field[string]                        `json:"account_address" api:"required"`
	Chain          param.Field[SuiTransactionScanParamsChain] `json:"chain" api:"required"`
	// Metadata
	Metadata    param.Field[SuiTransactionScanParamsMetadataUnion] `json:"metadata" api:"required"`
	Transaction param.Field[string]                                `json:"transaction" api:"required"`
	// List of options to include in the response
	//
	// - `Options.validation`: Include Options.validation output in the response
	//
	// - `Options.simulation`: Include Options.simulation output in the response
	Options param.Field[[]SuiTransactionScanParamsOption] `json:"options"`
	// Optional customer-supplied hints about transaction intent that cannot be
	// inferred from on-chain simulation.
	TransactionHints param.Field[[]SuiTransactionScanParamsTransactionHintUnion] `json:"transaction_hints"`
}

func (r SuiTransactionScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuiTransactionScanParamsChain string

const (
	SuiTransactionScanParamsChainMainnet SuiTransactionScanParamsChain = "mainnet"
)

func (r SuiTransactionScanParamsChain) IsKnown() bool {
	switch r {
	case SuiTransactionScanParamsChainMainnet:
		return true
	}
	return false
}

// Metadata
type SuiTransactionScanParamsMetadata struct {
	Account    param.Field[interface{}] `json:"account"`
	Connection param.Field[interface{}] `json:"connection"`
	// Metadata for wallet requests
	Type param.Field[SuiTransactionScanParamsMetadataType] `json:"type"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url"`
}

func (r SuiTransactionScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuiTransactionScanParamsMetadata) implementsSuiTransactionScanParamsMetadataUnion() {}

// Metadata
//
// Satisfied by [SuiTransactionScanParamsMetadataSuiWalletRequestMetadata],
// [SuiTransactionScanParamsMetadataSuiInAppRequestMetadata],
// [SuiTransactionScanParamsMetadata].
type SuiTransactionScanParamsMetadataUnion interface {
	implementsSuiTransactionScanParamsMetadataUnion()
}

type SuiTransactionScanParamsMetadataSuiWalletRequestMetadata struct {
	// Metadata for wallet requests
	Type param.Field[SuiTransactionScanParamsMetadataSuiWalletRequestMetadataType] `json:"type" api:"required"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url" api:"required"`
	// Account information associated with the request
	Account param.Field[SuiTransactionScanParamsMetadataSuiWalletRequestMetadataAccount] `json:"account"`
	// Connection metadata including user agent and IP information
	Connection param.Field[SuiTransactionScanParamsMetadataSuiWalletRequestMetadataConnection] `json:"connection"`
}

func (r SuiTransactionScanParamsMetadataSuiWalletRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuiTransactionScanParamsMetadataSuiWalletRequestMetadata) implementsSuiTransactionScanParamsMetadataUnion() {
}

// Metadata for wallet requests
type SuiTransactionScanParamsMetadataSuiWalletRequestMetadataType string

const (
	SuiTransactionScanParamsMetadataSuiWalletRequestMetadataTypeWallet SuiTransactionScanParamsMetadataSuiWalletRequestMetadataType = "wallet"
)

func (r SuiTransactionScanParamsMetadataSuiWalletRequestMetadataType) IsKnown() bool {
	switch r {
	case SuiTransactionScanParamsMetadataSuiWalletRequestMetadataTypeWallet:
		return true
	}
	return false
}

// Account information associated with the request
type SuiTransactionScanParamsMetadataSuiWalletRequestMetadataAccount struct {
	// Unique identifier for the account.
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// List of all account addresses in different chains based on the CAIPs standard
	// (https://github.com/ChainAgnostic/CAIPs). Ethereum mainnet example:
	// eip155:1:0xab16a96d359ec26a11e2c2b3d8f8b8942d5bfcdb
	AccountAddresses param.Field[[]string] `json:"account_addresses"`
	// Timestamp when the account was created.
	AccountCreationTimestamp param.Field[time.Time] `json:"account_creation_timestamp" format:"date-time"`
	// Age of the user in years
	UserAge param.Field[int64] `json:"user_age"`
	// ISO country code of the user's location.
	UserCountryCode param.Field[string] `json:"user_country_code"`
}

func (r SuiTransactionScanParamsMetadataSuiWalletRequestMetadataAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Connection metadata including user agent and IP information
type SuiTransactionScanParamsMetadataSuiWalletRequestMetadataConnection struct {
	// IP address of the customer making the request. Both IPv4 and IPv6 addresses are
	// supported.
	IPAddress param.Field[string] `json:"ip_address" api:"required" format:"ipvanyaddress"`
	// The full URL of the website that the request was directed to.
	Origin param.Field[string] `json:"origin" format:"uri"`
	// User agent string from the client's browser or application.
	UserAgent param.Field[string] `json:"user_agent"`
	// WalletConnect session description, when the request originates from a
	// WalletConnect session.
	WalletconnectDescription param.Field[string] `json:"walletconnect_description"`
	// WalletConnect session name, when the request originates from a WalletConnect
	// session.
	WalletconnectName param.Field[string] `json:"walletconnect_name"`
}

func (r SuiTransactionScanParamsMetadataSuiWalletRequestMetadataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuiTransactionScanParamsMetadataSuiInAppRequestMetadata struct {
	// Account information associated with the request
	Account param.Field[SuiTransactionScanParamsMetadataSuiInAppRequestMetadataAccount] `json:"account"`
	// Connection metadata including user agent and IP information
	Connection param.Field[SuiTransactionScanParamsMetadataSuiInAppRequestMetadataConnection] `json:"connection"`
	// Metadata for in-app requests
	Type param.Field[SuiTransactionScanParamsMetadataSuiInAppRequestMetadataType] `json:"type"`
}

func (r SuiTransactionScanParamsMetadataSuiInAppRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuiTransactionScanParamsMetadataSuiInAppRequestMetadata) implementsSuiTransactionScanParamsMetadataUnion() {
}

// Account information associated with the request
type SuiTransactionScanParamsMetadataSuiInAppRequestMetadataAccount struct {
	// Unique identifier for the account.
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// List of all account addresses in different chains based on the CAIPs standard
	// (https://github.com/ChainAgnostic/CAIPs). Ethereum mainnet example:
	// eip155:1:0xab16a96d359ec26a11e2c2b3d8f8b8942d5bfcdb
	AccountAddresses param.Field[[]string] `json:"account_addresses"`
	// Timestamp when the account was created.
	AccountCreationTimestamp param.Field[time.Time] `json:"account_creation_timestamp" format:"date-time"`
	// Age of the user in years
	UserAge param.Field[int64] `json:"user_age"`
	// ISO country code of the user's location.
	UserCountryCode param.Field[string] `json:"user_country_code"`
}

func (r SuiTransactionScanParamsMetadataSuiInAppRequestMetadataAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Connection metadata including user agent and IP information
type SuiTransactionScanParamsMetadataSuiInAppRequestMetadataConnection struct {
	// IP address of the customer making the request. Both IPv4 and IPv6 addresses are
	// supported.
	IPAddress param.Field[string] `json:"ip_address" api:"required" format:"ipvanyaddress"`
	// The full URL of the website that the request was directed to.
	Origin param.Field[string] `json:"origin" format:"uri"`
	// User agent string from the client's browser or application.
	UserAgent param.Field[string] `json:"user_agent"`
	// WalletConnect session description, when the request originates from a
	// WalletConnect session.
	WalletconnectDescription param.Field[string] `json:"walletconnect_description"`
	// WalletConnect session name, when the request originates from a WalletConnect
	// session.
	WalletconnectName param.Field[string] `json:"walletconnect_name"`
}

func (r SuiTransactionScanParamsMetadataSuiInAppRequestMetadataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Metadata for in-app requests
type SuiTransactionScanParamsMetadataSuiInAppRequestMetadataType string

const (
	SuiTransactionScanParamsMetadataSuiInAppRequestMetadataTypeInApp SuiTransactionScanParamsMetadataSuiInAppRequestMetadataType = "in_app"
)

func (r SuiTransactionScanParamsMetadataSuiInAppRequestMetadataType) IsKnown() bool {
	switch r {
	case SuiTransactionScanParamsMetadataSuiInAppRequestMetadataTypeInApp:
		return true
	}
	return false
}

// Metadata for wallet requests
type SuiTransactionScanParamsMetadataType string

const (
	SuiTransactionScanParamsMetadataTypeWallet SuiTransactionScanParamsMetadataType = "wallet"
	SuiTransactionScanParamsMetadataTypeInApp  SuiTransactionScanParamsMetadataType = "in_app"
)

func (r SuiTransactionScanParamsMetadataType) IsKnown() bool {
	switch r {
	case SuiTransactionScanParamsMetadataTypeWallet, SuiTransactionScanParamsMetadataTypeInApp:
		return true
	}
	return false
}

type SuiTransactionScanParamsOption string

const (
	SuiTransactionScanParamsOptionValidation SuiTransactionScanParamsOption = "validation"
	SuiTransactionScanParamsOptionSimulation SuiTransactionScanParamsOption = "simulation"
)

func (r SuiTransactionScanParamsOption) IsKnown() bool {
	switch r {
	case SuiTransactionScanParamsOptionValidation, SuiTransactionScanParamsOptionSimulation:
		return true
	}
	return false
}

// A single customer-supplied hint about transaction intent. The `type` field
// identifies which hint variant this is.
type SuiTransactionScanParamsTransactionHint struct {
	// Hint type discriminator (`cross_chain_bridge`).
	Type param.Field[string] `json:"type" api:"required"`
	// The intended recipient address on the destination chain. Required when the
	// bridge protocol does not emit this on-chain (e.g. Relay, some Across deposit
	// routes).
	DestinationAddress param.Field[string]      `json:"destination_address"`
	DestinationAsset   param.Field[interface{}] `json:"destination_asset"`
	DestinationChain   param.Field[interface{}] `json:"destination_chain"`
}

func (r SuiTransactionScanParamsTransactionHint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuiTransactionScanParamsTransactionHint) implementsSuiTransactionScanParamsTransactionHintUnion() {
}

// A single customer-supplied hint about transaction intent. The `type` field
// identifies which hint variant this is.
//
// Satisfied by [SuiTransactionScanParamsTransactionHintsCrossChainBridgeHint],
// [SuiTransactionScanParamsTransactionHintsGenericTransactionHint],
// [SuiTransactionScanParamsTransactionHint].
type SuiTransactionScanParamsTransactionHintUnion interface {
	implementsSuiTransactionScanParamsTransactionHintUnion()
}

// Customer-supplied context for a cross-chain bridge deposit where the protocol
// does not emit the destination on-chain.
type SuiTransactionScanParamsTransactionHintsCrossChainBridgeHint struct {
	// Hint type discriminator (`cross_chain_bridge`).
	Type param.Field[SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintType] `json:"type" api:"required"`
	// The intended recipient address on the destination chain. Required when the
	// bridge protocol does not emit this on-chain (e.g. Relay, some Across deposit
	// routes).
	DestinationAddress param.Field[string] `json:"destination_address"`
	// Details of the asset the recipient will receive on the destination chain. May
	// differ from the source asset (e.g. wrapped vs. native, canonical vs. bridged
	// token).
	DestinationAsset param.Field[SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetUnion] `json:"destination_asset"`
	// The destination chain for the bridged assets.
	DestinationChain param.Field[TransactionScanSupportedChain] `json:"destination_chain"`
}

func (r SuiTransactionScanParamsTransactionHintsCrossChainBridgeHint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuiTransactionScanParamsTransactionHintsCrossChainBridgeHint) implementsSuiTransactionScanParamsTransactionHintUnion() {
}

// Hint type discriminator (`cross_chain_bridge`).
type SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintType string

const (
	SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintTypeCrossChainBridge SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintType = "cross_chain_bridge"
)

func (r SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintType) IsKnown() bool {
	switch r {
	case SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintTypeCrossChainBridge:
		return true
	}
	return false
}

// Details of the asset the recipient will receive on the destination chain. May
// differ from the source asset (e.g. wrapped vs. native, canonical vs. bridged
// token).
type SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAsset struct {
	// Type of the asset (`NATIVE`)
	Type param.Field[SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetType] `json:"type" api:"required"`
	// Token contract address on the destination chain.
	Address param.Field[string] `json:"address"`
	// Amount to be received in the asset's smallest unit (before decimal division),
	// e.g. wei for ETH.
	RawValue param.Field[string] `json:"raw_value"`
	// Token ID of the specific NFT being bridged.
	TokenID param.Field[string] `json:"token_id"`
}

func (r SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAsset) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAsset) implementsSuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetUnion() {
}

// Details of the asset the recipient will receive on the destination chain. May
// differ from the source asset (e.g. wrapped vs. native, canonical vs. bridged
// token).
//
// Satisfied by
// [SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeNativeAsset],
// [SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeFungibleAsset],
// [SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeNonFungibleAsset],
// [SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAsset].
type SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetUnion interface {
	implementsSuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetUnion()
}

type SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeNativeAsset struct {
	// Type of the asset (`NATIVE`)
	Type param.Field[SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeNativeAssetType] `json:"type" api:"required"`
	// Amount to be received in the asset's smallest unit (before decimal division),
	// e.g. wei for ETH.
	RawValue param.Field[string] `json:"raw_value"`
}

func (r SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeNativeAsset) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeNativeAsset) implementsSuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetUnion() {
}

// Type of the asset (`NATIVE`)
type SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeNativeAssetType string

const (
	SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeNativeAssetTypeNative SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeNativeAssetType = "NATIVE"
)

func (r SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeNativeAssetType) IsKnown() bool {
	switch r {
	case SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeNativeAssetTypeNative:
		return true
	}
	return false
}

type SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeFungibleAsset struct {
	// Token contract address on the destination chain.
	Address param.Field[string] `json:"address" api:"required"`
	// Type of the asset (`FUNGIBLE`)
	Type param.Field[SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeFungibleAssetType] `json:"type" api:"required"`
	// Amount to be received in the asset's smallest unit (before decimal division),
	// e.g. base units for ERC-20 tokens.
	RawValue param.Field[string] `json:"raw_value"`
}

func (r SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeFungibleAsset) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeFungibleAsset) implementsSuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetUnion() {
}

// Type of the asset (`FUNGIBLE`)
type SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeFungibleAssetType string

const (
	SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeFungibleAssetTypeFungible SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeFungibleAssetType = "FUNGIBLE"
)

func (r SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeFungibleAssetType) IsKnown() bool {
	switch r {
	case SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeFungibleAssetTypeFungible:
		return true
	}
	return false
}

type SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeNonFungibleAsset struct {
	// NFT collection contract address on the destination chain.
	Address param.Field[string] `json:"address" api:"required"`
	// Token ID of the specific NFT being bridged.
	TokenID param.Field[string] `json:"token_id" api:"required"`
	// Type of the asset (`NON_FUNGIBLE`)
	Type param.Field[SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeNonFungibleAssetType] `json:"type" api:"required"`
}

func (r SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeNonFungibleAsset) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeNonFungibleAsset) implementsSuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetUnion() {
}

// Type of the asset (`NON_FUNGIBLE`)
type SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeNonFungibleAssetType string

const (
	SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeNonFungibleAssetTypeNonFungible SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeNonFungibleAssetType = "NON_FUNGIBLE"
)

func (r SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeNonFungibleAssetType) IsKnown() bool {
	switch r {
	case SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeNonFungibleAssetTypeNonFungible:
		return true
	}
	return false
}

// Type of the asset (`NATIVE`)
type SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetType string

const (
	SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetTypeNative      SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetType = "NATIVE"
	SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetTypeFungible    SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetType = "FUNGIBLE"
	SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetTypeNonFungible SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetType = "NON_FUNGIBLE"
)

func (r SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetType) IsKnown() bool {
	switch r {
	case SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetTypeNative, SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetTypeFungible, SuiTransactionScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetTypeNonFungible:
		return true
	}
	return false
}

// Fallback for unrecognized or future hint types. Accepts any hint with a `type`
// field.
type SuiTransactionScanParamsTransactionHintsGenericTransactionHint struct {
	// Hint type identifier for unrecognized or future hint types.
	Type param.Field[string] `json:"type" api:"required"`
}

func (r SuiTransactionScanParamsTransactionHintsGenericTransactionHint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuiTransactionScanParamsTransactionHintsGenericTransactionHint) implementsSuiTransactionScanParamsTransactionHintUnion() {
}
