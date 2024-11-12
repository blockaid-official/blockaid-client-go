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

// StarknetTransactionService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStarknetTransactionService] method instead.
type StarknetTransactionService struct {
	Options []option.RequestOption
}

// NewStarknetTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStarknetTransactionService(opts ...option.RequestOption) (r *StarknetTransactionService) {
	r = &StarknetTransactionService{}
	r.Options = opts
	return
}

// Scan Transactions
func (r *StarknetTransactionService) Scan(ctx context.Context, body StarknetTransactionScanParams, opts ...option.RequestOption) (res *StarknetTransactionScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/starknet/transaction/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type StarknetTransactionScanParams struct {
	AccountAddress param.Field[string] `json:"account_address,required"`
	// The chain name or chain ID
	Chain param.Field[StarknetTransactionScanParamsChain] `json:"chain,required"`
	// Metadata
	Metadata    param.Field[StarknetTransactionScanParamsMetadataUnion]    `json:"metadata,required"`
	Transaction param.Field[StarknetTransactionScanParamsTransactionUnion] `json:"transaction,required"`
	// Optional block number or tag context for the simulation
	BlockNumber param.Field[string] `json:"block_number"`
	// List of options to include in the response
	//
	// - `Options.validation`: Include Options.validation output in the response
	//
	// - `Options.simulation`: Include Options.simulation output in the response
	Options param.Field[[]StarknetTransactionScanParamsOption] `json:"options"`
}

func (r StarknetTransactionScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A CAIP-2 or a Starknet network name or a Starknet network name
type StarknetTransactionScanParamsChain string

const (
	StarknetTransactionScanParamsChainMainnet            StarknetTransactionScanParamsChain = "mainnet"
	StarknetTransactionScanParamsChainSepolia            StarknetTransactionScanParamsChain = "sepolia"
	StarknetTransactionScanParamsChainSepoliaIntegration StarknetTransactionScanParamsChain = "sepolia_integration"
)

func (r StarknetTransactionScanParamsChain) IsKnown() bool {
	switch r {
	case StarknetTransactionScanParamsChainMainnet, StarknetTransactionScanParamsChainSepolia, StarknetTransactionScanParamsChainSepoliaIntegration:
		return true
	}
	return false
}

// Metadata
type StarknetTransactionScanParamsMetadata struct {
	// Metadata for wallet requests
	Type param.Field[StarknetTransactionScanParamsMetadataType] `json:"type"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url"`
}

func (r StarknetTransactionScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionScanParamsMetadata) implementsStarknetTransactionScanParamsMetadataUnion() {
}

// Metadata
//
// Satisfied by
// [StarknetTransactionScanParamsMetadataStarknetWalletRequestMetadata],
// [StarknetTransactionScanParamsMetadataStarknetInAppRequestMetadata],
// [StarknetTransactionScanParamsMetadata].
type StarknetTransactionScanParamsMetadataUnion interface {
	implementsStarknetTransactionScanParamsMetadataUnion()
}

type StarknetTransactionScanParamsMetadataStarknetWalletRequestMetadata struct {
	// Metadata for wallet requests
	Type param.Field[StarknetTransactionScanParamsMetadataStarknetWalletRequestMetadataType] `json:"type,required"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url,required"`
}

func (r StarknetTransactionScanParamsMetadataStarknetWalletRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionScanParamsMetadataStarknetWalletRequestMetadata) implementsStarknetTransactionScanParamsMetadataUnion() {
}

// Metadata for wallet requests
type StarknetTransactionScanParamsMetadataStarknetWalletRequestMetadataType string

const (
	StarknetTransactionScanParamsMetadataStarknetWalletRequestMetadataTypeWallet StarknetTransactionScanParamsMetadataStarknetWalletRequestMetadataType = "wallet"
)

func (r StarknetTransactionScanParamsMetadataStarknetWalletRequestMetadataType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanParamsMetadataStarknetWalletRequestMetadataTypeWallet:
		return true
	}
	return false
}

type StarknetTransactionScanParamsMetadataStarknetInAppRequestMetadata struct {
	// Metadata for in-app requests
	Type param.Field[StarknetTransactionScanParamsMetadataStarknetInAppRequestMetadataType] `json:"type"`
}

func (r StarknetTransactionScanParamsMetadataStarknetInAppRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionScanParamsMetadataStarknetInAppRequestMetadata) implementsStarknetTransactionScanParamsMetadataUnion() {
}

// Metadata for in-app requests
type StarknetTransactionScanParamsMetadataStarknetInAppRequestMetadataType string

const (
	StarknetTransactionScanParamsMetadataStarknetInAppRequestMetadataTypeInApp StarknetTransactionScanParamsMetadataStarknetInAppRequestMetadataType = "in_app"
)

func (r StarknetTransactionScanParamsMetadataStarknetInAppRequestMetadataType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanParamsMetadataStarknetInAppRequestMetadataTypeInApp:
		return true
	}
	return false
}

// Metadata for wallet requests
type StarknetTransactionScanParamsMetadataType string

const (
	StarknetTransactionScanParamsMetadataTypeWallet StarknetTransactionScanParamsMetadataType = "wallet"
	StarknetTransactionScanParamsMetadataTypeInApp  StarknetTransactionScanParamsMetadataType = "in_app"
)

func (r StarknetTransactionScanParamsMetadataType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanParamsMetadataTypeWallet, StarknetTransactionScanParamsMetadataTypeInApp:
		return true
	}
	return false
}

type StarknetTransactionScanParamsTransaction struct {
	// The nonce of the transaction.
	Nonce param.Field[string] `json:"nonce,required"`
	// The version of the transaction.
	Version               param.Field[StarknetTransactionScanParamsTransactionVersion] `json:"version,required"`
	AccountDeploymentData param.Field[interface{}]                                     `json:"account_deployment_data"`
	Calldata              param.Field[interface{}]                                     `json:"calldata"`
	// The id of the chain to which the transaction is sent.
	ChainID param.Field[string] `json:"chain_id"`
	// The hash of the contract class.
	ClassHash           param.Field[string]      `json:"class_hash"`
	ConstructorCalldata param.Field[interface{}] `json:"constructor_calldata"`
	// The salt of the contract address.
	ContractAddressSalt param.Field[string] `json:"contract_address_salt"`
	// The maximum fee that the sender is willing to pay.
	MaxFee param.Field[string] `json:"max_fee"`
	// The nonce data availability mode.
	NonceDataAvailabilityMode param.Field[StarknetTransactionScanParamsTransactionNonceDataAvailabilityMode] `json:"nonce_data_availability_mode"`
	PaymasterData             param.Field[interface{}]                                                       `json:"paymaster_data"`
	// The address of the sender.
	SenderAddress param.Field[string] `json:"sender_address"`
}

func (r StarknetTransactionScanParamsTransaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionScanParamsTransaction) implementsStarknetTransactionScanParamsTransactionUnion() {
}

// Satisfied by
// [StarknetTransactionScanParamsTransactionStarknetInvokeV1TransactionSchema],
// [StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchema],
// [StarknetTransactionScanParamsTransactionStarknetDeployAccountV1TransactionSchema],
// [StarknetTransactionScanParamsTransactionStarknetDeployAccountV3TransactionSchema],
// [StarknetTransactionScanParamsTransaction].
type StarknetTransactionScanParamsTransactionUnion interface {
	implementsStarknetTransactionScanParamsTransactionUnion()
}

type StarknetTransactionScanParamsTransactionStarknetInvokeV1TransactionSchema struct {
	// The maximum fee that the sender is willing to pay.
	MaxFee param.Field[string] `json:"max_fee,required"`
	// The nonce of the transaction.
	Nonce param.Field[string] `json:"nonce,required"`
	// The address of the sender.
	SenderAddress param.Field[string] `json:"sender_address,required"`
	// The version of the transaction.
	Version param.Field[StarknetTransactionScanParamsTransactionStarknetInvokeV1TransactionSchemaVersion] `json:"version,required"`
	// The arguments that are passed to the validate and execute functions.
	Calldata param.Field[[]string] `json:"calldata"`
}

func (r StarknetTransactionScanParamsTransactionStarknetInvokeV1TransactionSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionScanParamsTransactionStarknetInvokeV1TransactionSchema) implementsStarknetTransactionScanParamsTransactionUnion() {
}

// The version of the transaction.
type StarknetTransactionScanParamsTransactionStarknetInvokeV1TransactionSchemaVersion int64

const (
	StarknetTransactionScanParamsTransactionStarknetInvokeV1TransactionSchemaVersion1 StarknetTransactionScanParamsTransactionStarknetInvokeV1TransactionSchemaVersion = 1
)

func (r StarknetTransactionScanParamsTransactionStarknetInvokeV1TransactionSchemaVersion) IsKnown() bool {
	switch r {
	case StarknetTransactionScanParamsTransactionStarknetInvokeV1TransactionSchemaVersion1:
		return true
	}
	return false
}

type StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchema struct {
	// The arguments that are passed to the validate and execute functions.
	Calldata param.Field[[]string] `json:"calldata,required"`
	// The id of the chain to which the transaction is sent.
	ChainID param.Field[string] `json:"chain_id,required"`
	// The nonce of the transaction.
	Nonce param.Field[string] `json:"nonce,required"`
	// The address of the sender.
	SenderAddress param.Field[string] `json:"sender_address,required"`
	// The version of the transaction.
	Version param.Field[StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaVersion] `json:"version,required"`
	// For future use. Currently this value is always empty.
	AccountDeploymentData param.Field[[]string] `json:"account_deployment_data"`
	// The nonce data availability mode.
	NonceDataAvailabilityMode param.Field[StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaNonceDataAvailabilityMode] `json:"nonce_data_availability_mode"`
	// For future use. Currently this value is always empty.
	PaymasterData param.Field[[]string] `json:"paymaster_data"`
}

func (r StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchema) implementsStarknetTransactionScanParamsTransactionUnion() {
}

// The version of the transaction.
type StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaVersion int64

const (
	StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaVersion3 StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaVersion = 3
)

func (r StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaVersion) IsKnown() bool {
	switch r {
	case StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaVersion3:
		return true
	}
	return false
}

// The nonce data availability mode.
type StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaNonceDataAvailabilityMode int64

const (
	StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaNonceDataAvailabilityMode0 StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaNonceDataAvailabilityMode = 0
)

func (r StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaNonceDataAvailabilityMode) IsKnown() bool {
	switch r {
	case StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaNonceDataAvailabilityMode0:
		return true
	}
	return false
}

type StarknetTransactionScanParamsTransactionStarknetDeployAccountV1TransactionSchema struct {
	// The hash of the contract class.
	ClassHash param.Field[string] `json:"class_hash,required"`
	// The arguments that are passed to the constructor function.
	ConstructorCalldata param.Field[[]string] `json:"constructor_calldata,required"`
	// The salt of the contract address.
	ContractAddressSalt param.Field[string] `json:"contract_address_salt,required"`
	// The maximum fee that the sender is willing to pay.
	MaxFee param.Field[string] `json:"max_fee,required"`
	// The nonce of the transaction.
	Nonce param.Field[string] `json:"nonce,required"`
	// The version of the transaction.
	Version param.Field[StarknetTransactionScanParamsTransactionStarknetDeployAccountV1TransactionSchemaVersion] `json:"version,required"`
}

func (r StarknetTransactionScanParamsTransactionStarknetDeployAccountV1TransactionSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionScanParamsTransactionStarknetDeployAccountV1TransactionSchema) implementsStarknetTransactionScanParamsTransactionUnion() {
}

// The version of the transaction.
type StarknetTransactionScanParamsTransactionStarknetDeployAccountV1TransactionSchemaVersion int64

const (
	StarknetTransactionScanParamsTransactionStarknetDeployAccountV1TransactionSchemaVersion1 StarknetTransactionScanParamsTransactionStarknetDeployAccountV1TransactionSchemaVersion = 1
)

func (r StarknetTransactionScanParamsTransactionStarknetDeployAccountV1TransactionSchemaVersion) IsKnown() bool {
	switch r {
	case StarknetTransactionScanParamsTransactionStarknetDeployAccountV1TransactionSchemaVersion1:
		return true
	}
	return false
}

type StarknetTransactionScanParamsTransactionStarknetDeployAccountV3TransactionSchema struct {
	// The hash of the contract class.
	ClassHash param.Field[string] `json:"class_hash,required"`
	// The arguments that are passed to the constructor function.
	ConstructorCalldata param.Field[[]string] `json:"constructor_calldata,required"`
	// The salt of the contract address.
	ContractAddressSalt param.Field[string] `json:"contract_address_salt,required"`
	// The maximum fee that the sender is willing to pay.
	MaxFee param.Field[string] `json:"max_fee,required"`
	// The nonce of the transaction.
	Nonce param.Field[string] `json:"nonce,required"`
	// The version of the transaction.
	Version param.Field[StarknetTransactionScanParamsTransactionStarknetDeployAccountV3TransactionSchemaVersion] `json:"version,required"`
}

func (r StarknetTransactionScanParamsTransactionStarknetDeployAccountV3TransactionSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionScanParamsTransactionStarknetDeployAccountV3TransactionSchema) implementsStarknetTransactionScanParamsTransactionUnion() {
}

// The version of the transaction.
type StarknetTransactionScanParamsTransactionStarknetDeployAccountV3TransactionSchemaVersion int64

const (
	StarknetTransactionScanParamsTransactionStarknetDeployAccountV3TransactionSchemaVersion3 StarknetTransactionScanParamsTransactionStarknetDeployAccountV3TransactionSchemaVersion = 3
)

func (r StarknetTransactionScanParamsTransactionStarknetDeployAccountV3TransactionSchemaVersion) IsKnown() bool {
	switch r {
	case StarknetTransactionScanParamsTransactionStarknetDeployAccountV3TransactionSchemaVersion3:
		return true
	}
	return false
}

// The version of the transaction.
type StarknetTransactionScanParamsTransactionVersion int64

const (
	StarknetTransactionScanParamsTransactionVersion1 StarknetTransactionScanParamsTransactionVersion = 1
	StarknetTransactionScanParamsTransactionVersion3 StarknetTransactionScanParamsTransactionVersion = 3
)

func (r StarknetTransactionScanParamsTransactionVersion) IsKnown() bool {
	switch r {
	case StarknetTransactionScanParamsTransactionVersion1, StarknetTransactionScanParamsTransactionVersion3:
		return true
	}
	return false
}

// The nonce data availability mode.
type StarknetTransactionScanParamsTransactionNonceDataAvailabilityMode int64

const (
	StarknetTransactionScanParamsTransactionNonceDataAvailabilityMode0 StarknetTransactionScanParamsTransactionNonceDataAvailabilityMode = 0
)

func (r StarknetTransactionScanParamsTransactionNonceDataAvailabilityMode) IsKnown() bool {
	switch r {
	case StarknetTransactionScanParamsTransactionNonceDataAvailabilityMode0:
		return true
	}
	return false
}

type StarknetTransactionScanParamsOption string

const (
	StarknetTransactionScanParamsOptionValidation StarknetTransactionScanParamsOption = "validation"
	StarknetTransactionScanParamsOptionSimulation StarknetTransactionScanParamsOption = "simulation"
)

func (r StarknetTransactionScanParamsOption) IsKnown() bool {
	switch r {
	case StarknetTransactionScanParamsOptionValidation, StarknetTransactionScanParamsOptionSimulation:
		return true
	}
	return false
}
