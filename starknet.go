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

// StarknetService contains methods and other services that help with interacting
// with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStarknetService] method instead.
type StarknetService struct {
	Options     []option.RequestOption
	Transaction *StarknetTransactionService
}

// NewStarknetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStarknetService(opts ...option.RequestOption) (r *StarknetService) {
	r = &StarknetService{}
	r.Options = opts
	r.Transaction = NewStarknetTransactionService(opts...)
	return
}

type StarknetErc1155Details struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC1155`)
	Type StarknetErc1155DetailsType `json:"type"`
	JSON starknetErc1155DetailsJSON `json:"-"`
}

// starknetErc1155DetailsJSON contains the JSON metadata for the struct
// [StarknetErc1155Details]
type starknetErc1155DetailsJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetErc1155Details) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetErc1155DetailsJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC1155`)
type StarknetErc1155DetailsType string

const (
	StarknetErc1155DetailsTypeErc1155 StarknetErc1155DetailsType = "ERC1155"
)

func (r StarknetErc1155DetailsType) IsKnown() bool {
	switch r {
	case StarknetErc1155DetailsTypeErc1155:
		return true
	}
	return false
}

type StarknetErc1155Diff struct {
	// Token ID of the transfer
	TokenID string `json:"token_id,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value int64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string                  `json:"summary,nullable"`
	JSON    starknetErc1155DiffJSON `json:"-"`
}

// starknetErc1155DiffJSON contains the JSON metadata for the struct
// [StarknetErc1155Diff]
type starknetErc1155DiffJSON struct {
	TokenID     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetErc1155DiffJSON) RawJSON() string {
	return r.raw
}

type StarknetErc20Details struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's decimals
	Decimals int64 `json:"decimals,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC20`)
	Type StarknetErc20DetailsType `json:"type"`
	JSON starknetErc20DetailsJSON `json:"-"`
}

// starknetErc20DetailsJSON contains the JSON metadata for the struct
// [StarknetErc20Details]
type starknetErc20DetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetErc20Details) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetErc20DetailsJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC20`)
type StarknetErc20DetailsType string

const (
	StarknetErc20DetailsTypeErc20 StarknetErc20DetailsType = "ERC20"
)

func (r StarknetErc20DetailsType) IsKnown() bool {
	switch r {
	case StarknetErc20DetailsTypeErc20:
		return true
	}
	return false
}

type StarknetErc20Diff struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                `json:"summary,nullable"`
	JSON    starknetErc20DiffJSON `json:"-"`
}

// starknetErc20DiffJSON contains the JSON metadata for the struct
// [StarknetErc20Diff]
type starknetErc20DiffJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetErc20DiffJSON) RawJSON() string {
	return r.raw
}

type StarknetErc721Details struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC721`)
	Type StarknetErc721DetailsType `json:"type"`
	JSON starknetErc721DetailsJSON `json:"-"`
}

// starknetErc721DetailsJSON contains the JSON metadata for the struct
// [StarknetErc721Details]
type starknetErc721DetailsJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetErc721Details) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetErc721DetailsJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC721`)
type StarknetErc721DetailsType string

const (
	StarknetErc721DetailsTypeErc721 StarknetErc721DetailsType = "ERC721"
)

func (r StarknetErc721DetailsType) IsKnown() bool {
	switch r {
	case StarknetErc721DetailsTypeErc721:
		return true
	}
	return false
}

type StarknetErc721Diff struct {
	// Token ID of the transfer
	TokenID string `json:"token_id,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Summarized description of the transfer
	Summary string                 `json:"summary,nullable"`
	JSON    starknetErc721DiffJSON `json:"-"`
}

// starknetErc721DiffJSON contains the JSON metadata for the struct
// [StarknetErc721Diff]
type starknetErc721DiffJSON struct {
	TokenID     apijson.Field
	UsdPrice    apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetErc721DiffJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanRequestParam struct {
	AccountAddress param.Field[string] `json:"account_address,required"`
	// The chain name or chain ID
	Chain param.Field[StarknetTransactionScanRequestChain] `json:"chain,required"`
	// Metadata
	Metadata    param.Field[StarknetTransactionScanRequestMetadataUnionParam]    `json:"metadata,required"`
	Transaction param.Field[StarknetTransactionScanRequestTransactionUnionParam] `json:"transaction,required"`
	// Optional block number or tag context for the simulation
	BlockNumber param.Field[string] `json:"block_number"`
	// List of options to include in the response
	//
	// - `Options.validation`: Include Options.validation output in the response
	//
	// - `Options.simulation`: Include Options.simulation output in the response
	Options param.Field[[]StarknetTransactionScanRequestOption] `json:"options"`
}

func (r StarknetTransactionScanRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A CAIP-2 or a Starknet network name or a Starknet network name
type StarknetTransactionScanRequestChain string

const (
	StarknetTransactionScanRequestChainMainnet StarknetTransactionScanRequestChain = "mainnet"
	StarknetTransactionScanRequestChainSepolia StarknetTransactionScanRequestChain = "sepolia"
)

func (r StarknetTransactionScanRequestChain) IsKnown() bool {
	switch r {
	case StarknetTransactionScanRequestChainMainnet, StarknetTransactionScanRequestChainSepolia:
		return true
	}
	return false
}

// Metadata
type StarknetTransactionScanRequestMetadataParam struct {
	// Metadata for wallet requests
	Type param.Field[StarknetTransactionScanRequestMetadataType] `json:"type"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url"`
}

func (r StarknetTransactionScanRequestMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionScanRequestMetadataParam) implementsStarknetTransactionScanRequestMetadataUnionParam() {
}

// Metadata
//
// Satisfied by
// [StarknetTransactionScanRequestMetadataStarknetWalletRequestMetadataParam],
// [StarknetTransactionScanRequestMetadataStarknetInAppRequestMetadataParam],
// [StarknetTransactionScanRequestMetadataParam].
type StarknetTransactionScanRequestMetadataUnionParam interface {
	implementsStarknetTransactionScanRequestMetadataUnionParam()
}

type StarknetTransactionScanRequestMetadataStarknetWalletRequestMetadataParam struct {
	// Metadata for wallet requests
	Type param.Field[StarknetTransactionScanRequestMetadataStarknetWalletRequestMetadataType] `json:"type,required"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url,required"`
}

func (r StarknetTransactionScanRequestMetadataStarknetWalletRequestMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionScanRequestMetadataStarknetWalletRequestMetadataParam) implementsStarknetTransactionScanRequestMetadataUnionParam() {
}

// Metadata for wallet requests
type StarknetTransactionScanRequestMetadataStarknetWalletRequestMetadataType string

const (
	StarknetTransactionScanRequestMetadataStarknetWalletRequestMetadataTypeWallet StarknetTransactionScanRequestMetadataStarknetWalletRequestMetadataType = "wallet"
)

func (r StarknetTransactionScanRequestMetadataStarknetWalletRequestMetadataType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanRequestMetadataStarknetWalletRequestMetadataTypeWallet:
		return true
	}
	return false
}

type StarknetTransactionScanRequestMetadataStarknetInAppRequestMetadataParam struct {
	// Metadata for in-app requests
	Type param.Field[StarknetTransactionScanRequestMetadataStarknetInAppRequestMetadataType] `json:"type"`
}

func (r StarknetTransactionScanRequestMetadataStarknetInAppRequestMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionScanRequestMetadataStarknetInAppRequestMetadataParam) implementsStarknetTransactionScanRequestMetadataUnionParam() {
}

// Metadata for in-app requests
type StarknetTransactionScanRequestMetadataStarknetInAppRequestMetadataType string

const (
	StarknetTransactionScanRequestMetadataStarknetInAppRequestMetadataTypeInApp StarknetTransactionScanRequestMetadataStarknetInAppRequestMetadataType = "in_app"
)

func (r StarknetTransactionScanRequestMetadataStarknetInAppRequestMetadataType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanRequestMetadataStarknetInAppRequestMetadataTypeInApp:
		return true
	}
	return false
}

// Metadata for wallet requests
type StarknetTransactionScanRequestMetadataType string

const (
	StarknetTransactionScanRequestMetadataTypeWallet StarknetTransactionScanRequestMetadataType = "wallet"
	StarknetTransactionScanRequestMetadataTypeInApp  StarknetTransactionScanRequestMetadataType = "in_app"
)

func (r StarknetTransactionScanRequestMetadataType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanRequestMetadataTypeWallet, StarknetTransactionScanRequestMetadataTypeInApp:
		return true
	}
	return false
}

type StarknetTransactionScanRequestTransactionParam struct {
	// The nonce of the transaction.
	Nonce param.Field[string] `json:"nonce,required"`
	// The version of the transaction.
	Version               param.Field[StarknetTransactionScanRequestTransactionVersion] `json:"version,required"`
	AccountDeploymentData param.Field[interface{}]                                      `json:"account_deployment_data"`
	Calldata              param.Field[interface{}]                                      `json:"calldata"`
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
	NonceDataAvailabilityMode param.Field[int64]       `json:"nonce_data_availability_mode"`
	PaymasterData             param.Field[interface{}] `json:"paymaster_data"`
	// The address of the sender.
	SenderAddress param.Field[string] `json:"sender_address"`
}

func (r StarknetTransactionScanRequestTransactionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionScanRequestTransactionParam) implementsStarknetTransactionScanRequestTransactionUnionParam() {
}

// Satisfied by
// [StarknetTransactionScanRequestTransactionStarknetInvokeV1TransactionSchemaParam],
// [StarknetTransactionScanRequestTransactionStarknetInvokeV3TransactionSchemaParam],
// [StarknetTransactionScanRequestTransactionStarknetDeployAccountV1TransactionSchemaParam],
// [StarknetTransactionScanRequestTransactionStarknetDeployAccountV3TransactionSchemaParam],
// [StarknetTransactionScanRequestTransactionParam].
type StarknetTransactionScanRequestTransactionUnionParam interface {
	implementsStarknetTransactionScanRequestTransactionUnionParam()
}

type StarknetTransactionScanRequestTransactionStarknetInvokeV1TransactionSchemaParam struct {
	// The maximum fee that the sender is willing to pay.
	MaxFee param.Field[string] `json:"max_fee,required"`
	// The nonce of the transaction.
	Nonce param.Field[string] `json:"nonce,required"`
	// The address of the sender.
	SenderAddress param.Field[string] `json:"sender_address,required"`
	// The version of the transaction.
	Version param.Field[StarknetTransactionScanRequestTransactionStarknetInvokeV1TransactionSchemaVersion] `json:"version,required"`
	// The arguments that are passed to the validate and execute functions.
	Calldata param.Field[[]string] `json:"calldata"`
}

func (r StarknetTransactionScanRequestTransactionStarknetInvokeV1TransactionSchemaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionScanRequestTransactionStarknetInvokeV1TransactionSchemaParam) implementsStarknetTransactionScanRequestTransactionUnionParam() {
}

// The version of the transaction.
type StarknetTransactionScanRequestTransactionStarknetInvokeV1TransactionSchemaVersion int64

const (
	StarknetTransactionScanRequestTransactionStarknetInvokeV1TransactionSchemaVersion1 StarknetTransactionScanRequestTransactionStarknetInvokeV1TransactionSchemaVersion = 1
)

func (r StarknetTransactionScanRequestTransactionStarknetInvokeV1TransactionSchemaVersion) IsKnown() bool {
	switch r {
	case StarknetTransactionScanRequestTransactionStarknetInvokeV1TransactionSchemaVersion1:
		return true
	}
	return false
}

type StarknetTransactionScanRequestTransactionStarknetInvokeV3TransactionSchemaParam struct {
	// The arguments that are passed to the validate and execute functions.
	Calldata param.Field[[]string] `json:"calldata,required"`
	// The id of the chain to which the transaction is sent.
	ChainID param.Field[string] `json:"chain_id,required"`
	// The nonce of the transaction.
	Nonce param.Field[string] `json:"nonce,required"`
	// The address of the sender.
	SenderAddress param.Field[string] `json:"sender_address,required"`
	// The version of the transaction.
	Version param.Field[StarknetTransactionScanRequestTransactionStarknetInvokeV3TransactionSchemaVersion] `json:"version,required"`
	// For future use. Currently this value is always empty.
	AccountDeploymentData param.Field[[]string] `json:"account_deployment_data"`
	// The nonce data availability mode.
	NonceDataAvailabilityMode param.Field[int64] `json:"nonce_data_availability_mode"`
	// For future use. Currently this value is always empty.
	PaymasterData param.Field[[]string] `json:"paymaster_data"`
}

func (r StarknetTransactionScanRequestTransactionStarknetInvokeV3TransactionSchemaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionScanRequestTransactionStarknetInvokeV3TransactionSchemaParam) implementsStarknetTransactionScanRequestTransactionUnionParam() {
}

// The version of the transaction.
type StarknetTransactionScanRequestTransactionStarknetInvokeV3TransactionSchemaVersion int64

const (
	StarknetTransactionScanRequestTransactionStarknetInvokeV3TransactionSchemaVersion3 StarknetTransactionScanRequestTransactionStarknetInvokeV3TransactionSchemaVersion = 3
)

func (r StarknetTransactionScanRequestTransactionStarknetInvokeV3TransactionSchemaVersion) IsKnown() bool {
	switch r {
	case StarknetTransactionScanRequestTransactionStarknetInvokeV3TransactionSchemaVersion3:
		return true
	}
	return false
}

type StarknetTransactionScanRequestTransactionStarknetDeployAccountV1TransactionSchemaParam struct {
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
	Version param.Field[StarknetTransactionScanRequestTransactionStarknetDeployAccountV1TransactionSchemaVersion] `json:"version,required"`
}

func (r StarknetTransactionScanRequestTransactionStarknetDeployAccountV1TransactionSchemaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionScanRequestTransactionStarknetDeployAccountV1TransactionSchemaParam) implementsStarknetTransactionScanRequestTransactionUnionParam() {
}

// The version of the transaction.
type StarknetTransactionScanRequestTransactionStarknetDeployAccountV1TransactionSchemaVersion int64

const (
	StarknetTransactionScanRequestTransactionStarknetDeployAccountV1TransactionSchemaVersion1 StarknetTransactionScanRequestTransactionStarknetDeployAccountV1TransactionSchemaVersion = 1
)

func (r StarknetTransactionScanRequestTransactionStarknetDeployAccountV1TransactionSchemaVersion) IsKnown() bool {
	switch r {
	case StarknetTransactionScanRequestTransactionStarknetDeployAccountV1TransactionSchemaVersion1:
		return true
	}
	return false
}

type StarknetTransactionScanRequestTransactionStarknetDeployAccountV3TransactionSchemaParam struct {
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
	Version param.Field[StarknetTransactionScanRequestTransactionStarknetDeployAccountV3TransactionSchemaVersion] `json:"version,required"`
}

func (r StarknetTransactionScanRequestTransactionStarknetDeployAccountV3TransactionSchemaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionScanRequestTransactionStarknetDeployAccountV3TransactionSchemaParam) implementsStarknetTransactionScanRequestTransactionUnionParam() {
}

// The version of the transaction.
type StarknetTransactionScanRequestTransactionStarknetDeployAccountV3TransactionSchemaVersion int64

const (
	StarknetTransactionScanRequestTransactionStarknetDeployAccountV3TransactionSchemaVersion3 StarknetTransactionScanRequestTransactionStarknetDeployAccountV3TransactionSchemaVersion = 3
)

func (r StarknetTransactionScanRequestTransactionStarknetDeployAccountV3TransactionSchemaVersion) IsKnown() bool {
	switch r {
	case StarknetTransactionScanRequestTransactionStarknetDeployAccountV3TransactionSchemaVersion3:
		return true
	}
	return false
}

// The version of the transaction.
type StarknetTransactionScanRequestTransactionVersion int64

const (
	StarknetTransactionScanRequestTransactionVersion1 StarknetTransactionScanRequestTransactionVersion = 1
	StarknetTransactionScanRequestTransactionVersion3 StarknetTransactionScanRequestTransactionVersion = 3
)

func (r StarknetTransactionScanRequestTransactionVersion) IsKnown() bool {
	switch r {
	case StarknetTransactionScanRequestTransactionVersion1, StarknetTransactionScanRequestTransactionVersion3:
		return true
	}
	return false
}

type StarknetTransactionScanRequestOption string

const (
	StarknetTransactionScanRequestOptionValidation StarknetTransactionScanRequestOption = "validation"
	StarknetTransactionScanRequestOptionSimulation StarknetTransactionScanRequestOption = "simulation"
)

func (r StarknetTransactionScanRequestOption) IsKnown() bool {
	switch r {
	case StarknetTransactionScanRequestOptionValidation, StarknetTransactionScanRequestOptionSimulation:
		return true
	}
	return false
}

type StarknetTransactionScanResponse struct {
	// Simulation result; Only present if simulation option is included in the request
	Simulation StarknetTransactionScanResponseSimulation `json:"simulation,nullable"`
	// Validation result; Only present if validation option is included in the request
	Validation StarknetTransactionScanResponseValidation `json:"validation,nullable"`
	JSON       starknetTransactionScanResponseJSON       `json:"-"`
}

// starknetTransactionScanResponseJSON contains the JSON metadata for the struct
// [StarknetTransactionScanResponse]
type starknetTransactionScanResponseJSON struct {
	Simulation  apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseJSON) RawJSON() string {
	return r.raw
}

// Simulation result; Only present if simulation option is included in the request
type StarknetTransactionScanResponseSimulation struct {
	Status StarknetTransactionScanResponseSimulationStatus `json:"status,required"`
	// This field can have the runtime type of
	// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummary].
	AccountSummary interface{} `json:"account_summary"`
	// This field can have the runtime type of
	// [[]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAddressDetail].
	AddressDetails interface{} `json:"address_details"`
	// This field can have the runtime type of
	// [map[string][]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff].
	AssetsDiffs interface{} `json:"assets_diffs"`
	// Optional block number or tag context for the simulation
	BlockNumber string `json:"block_number,nullable"`
	// Error message
	Error string `json:"error"`
	// This field can have the runtime type of
	// [map[string][]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure].
	Exposures interface{}                                   `json:"exposures"`
	JSON      starknetTransactionScanResponseSimulationJSON `json:"-"`
	union     StarknetTransactionScanResponseSimulationUnion
}

// starknetTransactionScanResponseSimulationJSON contains the JSON metadata for the
// struct [StarknetTransactionScanResponseSimulation]
type starknetTransactionScanResponseSimulationJSON struct {
	Status         apijson.Field
	AccountSummary apijson.Field
	AddressDetails apijson.Field
	AssetsDiffs    apijson.Field
	BlockNumber    apijson.Field
	Error          apijson.Field
	Exposures      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r starknetTransactionScanResponseSimulationJSON) RawJSON() string {
	return r.raw
}

func (r *StarknetTransactionScanResponseSimulation) UnmarshalJSON(data []byte) (err error) {
	*r = StarknetTransactionScanResponseSimulation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [StarknetTransactionScanResponseSimulationUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchema],
// [StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchema].
func (r StarknetTransactionScanResponseSimulation) AsUnion() StarknetTransactionScanResponseSimulationUnion {
	return r.union
}

// Simulation result; Only present if simulation option is included in the request
//
// Union satisfied by
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchema]
// or [StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchema].
type StarknetTransactionScanResponseSimulationUnion interface {
	implementsStarknetTransactionScanResponseSimulation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StarknetTransactionScanResponseSimulationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchema{}),
		},
	)
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchema struct {
	// Summary of the actions and asset transfers that were made by the requested
	// account address
	AccountSummary StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummary `json:"account_summary,required"`
	Status         StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaStatus         `json:"status,required"`
	// Details of addresses involved in the transaction
	AddressDetails []StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAddressDetail `json:"address_details"`
	// Mapping between the address of an account to the assets diff during the
	// transaction
	AssetsDiffs map[string][]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff `json:"assets_diffs"`
	// Optional block number or tag context for the simulation
	BlockNumber string `json:"block_number,nullable"`
	// Mapping between the address of an account to the exposure of the assets during
	// the transaction
	Exposures map[string][]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure `json:"exposures"`
	JSON      starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaJSON                  `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchema]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaJSON struct {
	AccountSummary apijson.Field
	Status         apijson.Field
	AddressDetails apijson.Field
	AssetsDiffs    apijson.Field
	BlockNumber    apijson.Field
	Exposures      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchema) implementsStarknetTransactionScanResponseSimulation() {
}

// Summary of the actions and asset transfers that were made by the requested
// account address
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummary struct {
	// Exposures made by the requested account address
	AccountExposures []StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure `json:"account_exposures,required"`
	// Total USD diff for the requested account address
	TotalUsdDiff StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryTotalUsdDiff `json:"total_usd_diff,required"`
	// Assets diffs of the requested account address
	AccountAssetsDiffs []StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiff `json:"account_assets_diffs"`
	// Total USD exposure for each of the spender addresses during the transaction
	TotalUsdExposure map[string]float64                                                                                `json:"total_usd_exposure"`
	JSON             starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummary]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryJSON struct {
	AccountExposures   apijson.Field
	TotalUsdDiff       apijson.Field
	AccountAssetsDiffs apijson.Field
	TotalUsdExposure   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure struct {
	// This field can have the runtime type of [StarknetErc20Details],
	// [StarknetErc721Details], [StarknetErc1155Details].
	Asset interface{} `json:"asset,required"`
	// This field can have the runtime type of
	// [map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpender],
	// [map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpender],
	// [map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpender].
	Spenders interface{}                                                                                                      `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposureJSON `json:"-"`
	union    StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresUnion
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposureJSON) RawJSON() string {
	return r.raw
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure) UnmarshalJSON(data []byte) (err error) {
	*r = StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20Exposure],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721Exposure],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155Exposure].
func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure) AsUnion() StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresUnion {
	return r.union
}

// Union satisfied by
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20Exposure],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721Exposure]
// or
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155Exposure].
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresUnion interface {
	implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20Exposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721Exposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155Exposure{}),
		},
	)
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20Exposure struct {
	Asset StarknetErc20Details `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpender `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureJSON               `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20Exposure]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20Exposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20Exposure) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpender struct {
	// Approval value of the ERC20 token
	Approval string              `json:"approval,required"`
	Exposure []StarknetErc20Diff `json:"exposure,required"`
	// Expiration date of the approval
	Expiration time.Time `json:"expiration,nullable" format:"date-time"`
	// Summarized description of the exposure
	Summary string                                                                                                                                        `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpenderJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpenderJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpender]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpenderJSON struct {
	Approval    apijson.Field
	Exposure    apijson.Field
	Expiration  apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc20ExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721Exposure struct {
	Asset StarknetErc721Details `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpender `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureJSON               `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721Exposure]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721Exposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721Exposure) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpender struct {
	Exposure []StarknetErc721Diff `json:"exposure,required"`
	// Whether `setApprovalForAll` was invoked
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// Summarized description of the exposure
	Summary string                                                                                                                                         `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpenderJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpenderJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpender]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc721ExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155Exposure struct {
	Asset StarknetErc1155Details `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpender `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureJSON               `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155Exposure]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155Exposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155Exposure) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposure() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpender struct {
	Exposure []StarknetErc1155Diff `json:"exposure,required"`
	// Whether `setApprovalForAll` was invoked
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// Summarized description of the exposure
	Summary string                                                                                                                                          `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpenderJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpenderJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpender]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountExposuresStarknetErc1155ExposureSpenderJSON) RawJSON() string {
	return r.raw
}

// Total USD diff for the requested account address
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryTotalUsdDiff struct {
	// Total incoming USD transfers
	In float64 `json:"in,required"`
	// Total outgoing USD transfers
	Out float64 `json:"out,required"`
	// Total USD transfers
	Total float64                                                                                                       `json:"total"`
	JSON  starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryTotalUsdDiffJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryTotalUsdDiffJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryTotalUsdDiff]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiff struct {
	// This field can have the runtime type of [StarknetErc20Details],
	// [StarknetErc721Details], [StarknetErc1155Details].
	Asset interface{} `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of [StarknetErc20Diff],
	// [StarknetErc721Diff], [StarknetErc1155Diff].
	In interface{} `json:"in"`
	// This field can have the runtime type of [StarknetErc20Diff],
	// [StarknetErc721Diff], [StarknetErc1155Diff].
	Out   interface{}                                                                                                        `json:"out"`
	JSON  starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffJSON `json:"-"`
	union StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsUnion
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiff]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiff],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiff],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiff].
func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiff) AsUnion() StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiff],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiff]
// or
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiff].
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsUnion interface {
	implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiff{}),
		},
	)
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiff struct {
	Asset StarknetErc20Details `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StarknetErc20Diff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetErc20Diff                                                                                                                         `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiff]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc20AssetDiff) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiff struct {
	Asset StarknetErc721Details `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StarknetErc721Diff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetErc721Diff                                                                                                                         `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiff]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc721AssetDiff) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiff struct {
	Asset StarknetErc1155Details `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StarknetErc1155Diff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetErc1155Diff                                                                                                                         `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiff]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiffsStarknetErc1155AssetDiff) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAccountSummaryAccountAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaStatus string

const (
	StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaStatusSuccess StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaStatus = "Success"
)

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaStatus) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaStatusSuccess:
		return true
	}
	return false
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAddressDetail struct {
	// Encoded public key of the account
	AccountAddress interface{} `json:"account_address,required"`
	// Description of the account
	Description string                                                                                           `json:"description,nullable"`
	JSON        starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAddressDetailJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAddressDetailJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAddressDetail]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAddressDetailJSON struct {
	AccountAddress apijson.Field
	Description    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAddressDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAddressDetailJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff struct {
	// This field can have the runtime type of [StarknetErc20Details],
	// [StarknetErc721Details], [StarknetErc1155Details].
	Asset interface{} `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of [StarknetErc20Diff],
	// [StarknetErc721Diff], [StarknetErc1155Diff].
	In interface{} `json:"in"`
	// This field can have the runtime type of [StarknetErc20Diff],
	// [StarknetErc721Diff], [StarknetErc1155Diff].
	Out   interface{}                                                                                   `json:"out"`
	JSON  starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffJSON `json:"-"`
	union StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsUnion
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiff],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiff],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiff].
func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff) AsUnion() StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiff],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiff]
// or
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiff].
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsUnion interface {
	implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiff{}),
		},
	)
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiff struct {
	Asset StarknetErc20Details `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StarknetErc20Diff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetErc20Diff                                                                                                    `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiff]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc20AssetDiff) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiff struct {
	Asset StarknetErc721Details `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StarknetErc721Diff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetErc721Diff                                                                                                    `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiff]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc721AssetDiff) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiff struct {
	Asset StarknetErc1155Details `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In StarknetErc1155Diff `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  StarknetErc1155Diff                                                                                                    `json:"out,nullable"`
	JSON starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiff]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiffsStarknetErc1155AssetDiff) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaAssetsDiff() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure struct {
	// This field can have the runtime type of [StarknetErc20Details],
	// [StarknetErc721Details], [StarknetErc1155Details].
	Asset interface{} `json:"asset,required"`
	// This field can have the runtime type of
	// [map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpender],
	// [map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpender],
	// [map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpender].
	Spenders interface{}                                                                                 `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposureJSON `json:"-"`
	union    StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresUnion
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposureJSON) RawJSON() string {
	return r.raw
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure) UnmarshalJSON(data []byte) (err error) {
	*r = StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20Exposure],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721Exposure],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155Exposure].
func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure) AsUnion() StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresUnion {
	return r.union
}

// Union satisfied by
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20Exposure],
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721Exposure]
// or
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155Exposure].
type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresUnion interface {
	implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20Exposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721Exposure{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155Exposure{}),
		},
	)
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20Exposure struct {
	Asset StarknetErc20Details `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpender `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureJSON               `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20Exposure]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20Exposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20Exposure) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpender struct {
	// Approval value of the ERC20 token
	Approval string              `json:"approval,required"`
	Exposure []StarknetErc20Diff `json:"exposure,required"`
	// Expiration date of the approval
	Expiration time.Time `json:"expiration,nullable" format:"date-time"`
	// Summarized description of the exposure
	Summary string                                                                                                                   `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpenderJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpenderJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpender]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpenderJSON struct {
	Approval    apijson.Field
	Exposure    apijson.Field
	Expiration  apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc20ExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721Exposure struct {
	Asset StarknetErc721Details `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpender `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureJSON               `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721Exposure]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721Exposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721Exposure) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpender struct {
	Exposure []StarknetErc721Diff `json:"exposure,required"`
	// Whether `setApprovalForAll` was invoked
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// Summarized description of the exposure
	Summary string                                                                                                                    `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpenderJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpenderJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpender]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc721ExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155Exposure struct {
	Asset StarknetErc1155Details `json:"asset,required"`
	// Mapping between the spender address and the exposure of the asset
	Spenders map[string]StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpender `json:"spenders"`
	JSON     starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureJSON               `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155Exposure]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155Exposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155Exposure) implementsStarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposure() {
}

type StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpender struct {
	Exposure []StarknetErc1155Diff `json:"exposure,required"`
	// Whether `setApprovalForAll` was invoked
	IsApprovedForAll bool `json:"is_approved_for_all,required"`
	// Summarized description of the exposure
	Summary string                                                                                                                     `json:"summary,nullable"`
	JSON    starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpenderJSON `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpenderJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpender]
type starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpenderJSON struct {
	Exposure         apijson.Field
	IsApprovedForAll apijson.Field
	Summary          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetStarknetSimulationResultSchemaExposuresStarknetErc1155ExposureSpenderJSON) RawJSON() string {
	return r.raw
}

type StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchema struct {
	// Error message
	Error  string                                                                       `json:"error,required"`
	Status StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchemaStatus `json:"status,required"`
	JSON   starknetTransactionScanResponseSimulationStarknetSimulationErrorSchemaJSON   `json:"-"`
}

// starknetTransactionScanResponseSimulationStarknetSimulationErrorSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchema]
type starknetTransactionScanResponseSimulationStarknetSimulationErrorSchemaJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseSimulationStarknetSimulationErrorSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchema) implementsStarknetTransactionScanResponseSimulation() {
}

type StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchemaStatus string

const (
	StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchemaStatusError StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchemaStatus = "Error"
)

func (r StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchemaStatus) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStarknetSimulationErrorSchemaStatusError:
		return true
	}
	return false
}

type StarknetTransactionScanResponseSimulationStatus string

const (
	StarknetTransactionScanResponseSimulationStatusSuccess StarknetTransactionScanResponseSimulationStatus = "Success"
	StarknetTransactionScanResponseSimulationStatusError   StarknetTransactionScanResponseSimulationStatus = "Error"
)

func (r StarknetTransactionScanResponseSimulationStatus) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseSimulationStatusSuccess, StarknetTransactionScanResponseSimulationStatusError:
		return true
	}
	return false
}

// Validation result; Only present if validation option is included in the request
type StarknetTransactionScanResponseValidation struct {
	Status StarknetTransactionScanResponseValidationStatus `json:"status,required"`
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification"`
	// A textual description about the validation result
	Description string `json:"description"`
	// Error message
	Error string `json:"error"`
	// This field can have the runtime type of
	// [[]StarknetTransactionScanResponseValidationStarknetValidationResultFeature].
	Features interface{} `json:"features"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason string `json:"reason"`
	// Verdict of the validation
	ResultType StarknetTransactionScanResponseValidationResultType `json:"result_type"`
	JSON       starknetTransactionScanResponseValidationJSON       `json:"-"`
	union      StarknetTransactionScanResponseValidationUnion
}

// starknetTransactionScanResponseValidationJSON contains the JSON metadata for the
// struct [StarknetTransactionScanResponseValidation]
type starknetTransactionScanResponseValidationJSON struct {
	Status         apijson.Field
	Classification apijson.Field
	Description    apijson.Field
	Error          apijson.Field
	Features       apijson.Field
	Reason         apijson.Field
	ResultType     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r starknetTransactionScanResponseValidationJSON) RawJSON() string {
	return r.raw
}

func (r *StarknetTransactionScanResponseValidation) UnmarshalJSON(data []byte) (err error) {
	*r = StarknetTransactionScanResponseValidation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [StarknetTransactionScanResponseValidationUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StarknetTransactionScanResponseValidationStarknetValidationResult],
// [StarknetTransactionScanResponseValidationStarknetValidationErrorSchema].
func (r StarknetTransactionScanResponseValidation) AsUnion() StarknetTransactionScanResponseValidationUnion {
	return r.union
}

// Validation result; Only present if validation option is included in the request
//
// Union satisfied by
// [StarknetTransactionScanResponseValidationStarknetValidationResult] or
// [StarknetTransactionScanResponseValidationStarknetValidationErrorSchema].
type StarknetTransactionScanResponseValidationUnion interface {
	implementsStarknetTransactionScanResponseValidation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StarknetTransactionScanResponseValidationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseValidationStarknetValidationResult{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StarknetTransactionScanResponseValidationStarknetValidationErrorSchema{}),
		},
	)
}

type StarknetTransactionScanResponseValidationStarknetValidationResult struct {
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification,required"`
	// A textual description about the validation result
	Description string                                                                     `json:"description,required"`
	Features    []StarknetTransactionScanResponseValidationStarknetValidationResultFeature `json:"features,required"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason string `json:"reason,required"`
	// Verdict of the validation
	ResultType StarknetTransactionScanResponseValidationStarknetValidationResultResultType `json:"result_type,required"`
	Status     StarknetTransactionScanResponseValidationStarknetValidationResultStatus     `json:"status,required"`
	JSON       starknetTransactionScanResponseValidationStarknetValidationResultJSON       `json:"-"`
}

// starknetTransactionScanResponseValidationStarknetValidationResultJSON contains
// the JSON metadata for the struct
// [StarknetTransactionScanResponseValidationStarknetValidationResult]
type starknetTransactionScanResponseValidationStarknetValidationResultJSON struct {
	Classification apijson.Field
	Description    apijson.Field
	Features       apijson.Field
	Reason         apijson.Field
	ResultType     apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseValidationStarknetValidationResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseValidationStarknetValidationResultJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseValidationStarknetValidationResult) implementsStarknetTransactionScanResponseValidation() {
}

type StarknetTransactionScanResponseValidationStarknetValidationResultFeature struct {
	// Address the feature refers to
	Address string `json:"address,required"`
	// Textual description
	Description string `json:"description,required"`
	FeatureID   string `json:"feature_id,required"`
	// Feature Classification
	Type StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesType `json:"type,required"`
	JSON starknetTransactionScanResponseValidationStarknetValidationResultFeatureJSON  `json:"-"`
}

// starknetTransactionScanResponseValidationStarknetValidationResultFeatureJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseValidationStarknetValidationResultFeature]
type starknetTransactionScanResponseValidationStarknetValidationResultFeatureJSON struct {
	Address     apijson.Field
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseValidationStarknetValidationResultFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseValidationStarknetValidationResultFeatureJSON) RawJSON() string {
	return r.raw
}

// Feature Classification
type StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesType string

const (
	StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesTypeBenign    StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesType = "Benign"
	StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesTypeWarning   StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesType = "Warning"
	StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesTypeMalicious StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesType = "Malicious"
	StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesTypeInfo      StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesType = "Info"
)

func (r StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesTypeBenign, StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesTypeWarning, StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesTypeMalicious, StarknetTransactionScanResponseValidationStarknetValidationResultFeaturesTypeInfo:
		return true
	}
	return false
}

// Verdict of the validation
type StarknetTransactionScanResponseValidationStarknetValidationResultResultType string

const (
	StarknetTransactionScanResponseValidationStarknetValidationResultResultTypeBenign    StarknetTransactionScanResponseValidationStarknetValidationResultResultType = "Benign"
	StarknetTransactionScanResponseValidationStarknetValidationResultResultTypeWarning   StarknetTransactionScanResponseValidationStarknetValidationResultResultType = "Warning"
	StarknetTransactionScanResponseValidationStarknetValidationResultResultTypeMalicious StarknetTransactionScanResponseValidationStarknetValidationResultResultType = "Malicious"
)

func (r StarknetTransactionScanResponseValidationStarknetValidationResultResultType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseValidationStarknetValidationResultResultTypeBenign, StarknetTransactionScanResponseValidationStarknetValidationResultResultTypeWarning, StarknetTransactionScanResponseValidationStarknetValidationResultResultTypeMalicious:
		return true
	}
	return false
}

type StarknetTransactionScanResponseValidationStarknetValidationResultStatus string

const (
	StarknetTransactionScanResponseValidationStarknetValidationResultStatusSuccess StarknetTransactionScanResponseValidationStarknetValidationResultStatus = "Success"
)

func (r StarknetTransactionScanResponseValidationStarknetValidationResultStatus) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseValidationStarknetValidationResultStatusSuccess:
		return true
	}
	return false
}

type StarknetTransactionScanResponseValidationStarknetValidationErrorSchema struct {
	// Error message
	Error  string                                                                       `json:"error,required"`
	Status StarknetTransactionScanResponseValidationStarknetValidationErrorSchemaStatus `json:"status,required"`
	JSON   starknetTransactionScanResponseValidationStarknetValidationErrorSchemaJSON   `json:"-"`
}

// starknetTransactionScanResponseValidationStarknetValidationErrorSchemaJSON
// contains the JSON metadata for the struct
// [StarknetTransactionScanResponseValidationStarknetValidationErrorSchema]
type starknetTransactionScanResponseValidationStarknetValidationErrorSchemaJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetTransactionScanResponseValidationStarknetValidationErrorSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetTransactionScanResponseValidationStarknetValidationErrorSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StarknetTransactionScanResponseValidationStarknetValidationErrorSchema) implementsStarknetTransactionScanResponseValidation() {
}

type StarknetTransactionScanResponseValidationStarknetValidationErrorSchemaStatus string

const (
	StarknetTransactionScanResponseValidationStarknetValidationErrorSchemaStatusError StarknetTransactionScanResponseValidationStarknetValidationErrorSchemaStatus = "Error"
)

func (r StarknetTransactionScanResponseValidationStarknetValidationErrorSchemaStatus) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseValidationStarknetValidationErrorSchemaStatusError:
		return true
	}
	return false
}

type StarknetTransactionScanResponseValidationStatus string

const (
	StarknetTransactionScanResponseValidationStatusSuccess StarknetTransactionScanResponseValidationStatus = "Success"
	StarknetTransactionScanResponseValidationStatusError   StarknetTransactionScanResponseValidationStatus = "Error"
)

func (r StarknetTransactionScanResponseValidationStatus) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseValidationStatusSuccess, StarknetTransactionScanResponseValidationStatusError:
		return true
	}
	return false
}

// Verdict of the validation
type StarknetTransactionScanResponseValidationResultType string

const (
	StarknetTransactionScanResponseValidationResultTypeBenign    StarknetTransactionScanResponseValidationResultType = "Benign"
	StarknetTransactionScanResponseValidationResultTypeWarning   StarknetTransactionScanResponseValidationResultType = "Warning"
	StarknetTransactionScanResponseValidationResultTypeMalicious StarknetTransactionScanResponseValidationResultType = "Malicious"
)

func (r StarknetTransactionScanResponseValidationResultType) IsKnown() bool {
	switch r {
	case StarknetTransactionScanResponseValidationResultTypeBenign, StarknetTransactionScanResponseValidationResultTypeWarning, StarknetTransactionScanResponseValidationResultTypeMalicious:
		return true
	}
	return false
}
