// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"reflect"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/param"
	"github.com/blockaid-official/blockaid-client-go/option"
	"github.com/tidwall/gjson"
)

// StellarService contains methods and other services that help with interacting
// with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStellarService] method instead.
type StellarService struct {
	Options     []option.RequestOption
	Transaction *StellarTransactionService
}

// NewStellarService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStellarService(opts ...option.RequestOption) (r *StellarService) {
	r = &StellarService{}
	r.Options = opts
	r.Transaction = NewStellarTransactionService(opts...)
	return
}

type StellarAssetContractDetailsSchema struct {
	// Address of the asset's contract
	Address string `json:"address,required"`
	// Asset code
	Name string `json:"name,required"`
	// Asset symbol
	Symbol string `json:"symbol,required"`
	// Type of the asset (`CONTRACT`)
	Type StellarAssetContractDetailsSchemaType `json:"type"`
	JSON stellarAssetContractDetailsSchemaJSON `json:"-"`
}

// stellarAssetContractDetailsSchemaJSON contains the JSON metadata for the struct
// [StellarAssetContractDetailsSchema]
type stellarAssetContractDetailsSchemaJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarAssetContractDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarAssetContractDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StellarAssetContractDetailsSchema) implementsTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAsset() {
}

func (r StellarAssetContractDetailsSchema) implementsTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAsset() {
}

// Type of the asset (`CONTRACT`)
type StellarAssetContractDetailsSchemaType string

const (
	StellarAssetContractDetailsSchemaTypeContract StellarAssetContractDetailsSchemaType = "CONTRACT"
)

func (r StellarAssetContractDetailsSchemaType) IsKnown() bool {
	switch r {
	case StellarAssetContractDetailsSchemaTypeContract:
		return true
	}
	return false
}

type StellarAssetTransferDetailsSchema struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                               `json:"usd_price"`
	JSON     stellarAssetTransferDetailsSchemaJSON `json:"-"`
}

// stellarAssetTransferDetailsSchemaJSON contains the JSON metadata for the struct
// [StellarAssetTransferDetailsSchema]
type stellarAssetTransferDetailsSchemaJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarAssetTransferDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarAssetTransferDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

type StellarTransactionScanRequestParam struct {
	AccountAddress param.Field[string] `json:"account_address,required"`
	// A CAIP-2 chain ID or a Stellar network name
	Chain param.Field[StellarTransactionScanRequestChain] `json:"chain,required"`
	// Metadata
	Metadata param.Field[StellarTransactionScanRequestMetadataUnionParam] `json:"metadata,required"`
	// List of XDR-encoded transactions to be scanned
	Transactions param.Field[[]string] `json:"transactions,required"`
	// List of options to include in the response
	//
	// - `simulation`: Include simulation output in the response
	// - `validation`: Include security validation of the transaction in the response
	Options param.Field[[]StellarTransactionScanRequestOption] `json:"options"`
}

func (r StellarTransactionScanRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A CAIP-2 chain ID or a Stellar network name
type StellarTransactionScanRequestChain string

const (
	StellarTransactionScanRequestChainPubnet    StellarTransactionScanRequestChain = "pubnet"
	StellarTransactionScanRequestChainFuturenet StellarTransactionScanRequestChain = "futurenet"
)

func (r StellarTransactionScanRequestChain) IsKnown() bool {
	switch r {
	case StellarTransactionScanRequestChainPubnet, StellarTransactionScanRequestChainFuturenet:
		return true
	}
	return false
}

// Metadata
type StellarTransactionScanRequestMetadataParam struct {
	// Metadata for wallet requests
	Type param.Field[StellarTransactionScanRequestMetadataType] `json:"type"`
	// URL of the dApp that originated the transaction
	URL param.Field[string] `json:"url"`
}

func (r StellarTransactionScanRequestMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionScanRequestMetadataParam) implementsStellarTransactionScanRequestMetadataUnionParam() {
}

// Metadata
//
// Satisfied by
// [StellarTransactionScanRequestMetadataStellarWalletRequestMetadataParam],
// [StellarTransactionScanRequestMetadataStellarInAppRequestMetadataParam],
// [StellarTransactionScanRequestMetadataParam].
type StellarTransactionScanRequestMetadataUnionParam interface {
	implementsStellarTransactionScanRequestMetadataUnionParam()
}

type StellarTransactionScanRequestMetadataStellarWalletRequestMetadataParam struct {
	// URL of the dApp that originated the transaction
	URL param.Field[string] `json:"url,required"`
	// Metadata for wallet requests
	Type param.Field[StellarTransactionScanRequestMetadataStellarWalletRequestMetadataType] `json:"type"`
}

func (r StellarTransactionScanRequestMetadataStellarWalletRequestMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionScanRequestMetadataStellarWalletRequestMetadataParam) implementsStellarTransactionScanRequestMetadataUnionParam() {
}

// Metadata for wallet requests
type StellarTransactionScanRequestMetadataStellarWalletRequestMetadataType string

const (
	StellarTransactionScanRequestMetadataStellarWalletRequestMetadataTypeWallet StellarTransactionScanRequestMetadataStellarWalletRequestMetadataType = "wallet"
)

func (r StellarTransactionScanRequestMetadataStellarWalletRequestMetadataType) IsKnown() bool {
	switch r {
	case StellarTransactionScanRequestMetadataStellarWalletRequestMetadataTypeWallet:
		return true
	}
	return false
}

type StellarTransactionScanRequestMetadataStellarInAppRequestMetadataParam struct {
	// Metadata for in-app requests
	Type param.Field[StellarTransactionScanRequestMetadataStellarInAppRequestMetadataType] `json:"type"`
}

func (r StellarTransactionScanRequestMetadataStellarInAppRequestMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionScanRequestMetadataStellarInAppRequestMetadataParam) implementsStellarTransactionScanRequestMetadataUnionParam() {
}

// Metadata for in-app requests
type StellarTransactionScanRequestMetadataStellarInAppRequestMetadataType string

const (
	StellarTransactionScanRequestMetadataStellarInAppRequestMetadataTypeInApp StellarTransactionScanRequestMetadataStellarInAppRequestMetadataType = "in_app"
)

func (r StellarTransactionScanRequestMetadataStellarInAppRequestMetadataType) IsKnown() bool {
	switch r {
	case StellarTransactionScanRequestMetadataStellarInAppRequestMetadataTypeInApp:
		return true
	}
	return false
}

// Metadata for wallet requests
type StellarTransactionScanRequestMetadataType string

const (
	StellarTransactionScanRequestMetadataTypeWallet StellarTransactionScanRequestMetadataType = "wallet"
	StellarTransactionScanRequestMetadataTypeInApp  StellarTransactionScanRequestMetadataType = "in_app"
)

func (r StellarTransactionScanRequestMetadataType) IsKnown() bool {
	switch r {
	case StellarTransactionScanRequestMetadataTypeWallet, StellarTransactionScanRequestMetadataTypeInApp:
		return true
	}
	return false
}

type StellarTransactionScanRequestOption string

const (
	StellarTransactionScanRequestOptionValidation StellarTransactionScanRequestOption = "validation"
	StellarTransactionScanRequestOptionSimulation StellarTransactionScanRequestOption = "simulation"
)

func (r StellarTransactionScanRequestOption) IsKnown() bool {
	switch r {
	case StellarTransactionScanRequestOptionValidation, StellarTransactionScanRequestOptionSimulation:
		return true
	}
	return false
}

type TransactionScanResponse struct {
	// Simulation result; Only present if simulation option is included in the request
	Simulation TransactionScanResponseSimulation `json:"simulation,nullable"`
	// Validation result; Only present if validation option is included in the request
	Validation TransactionScanResponseValidation `json:"validation,nullable"`
	JSON       transactionScanResponseJSON       `json:"-"`
}

// transactionScanResponseJSON contains the JSON metadata for the struct
// [TransactionScanResponse]
type transactionScanResponseJSON struct {
	Simulation  apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseJSON) RawJSON() string {
	return r.raw
}

// Simulation result; Only present if simulation option is included in the request
type TransactionScanResponseSimulation struct {
	Status TransactionScanResponseSimulationStatus `json:"status,required"`
	// This field can have the runtime type of
	// [map[string][]TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiff].
	AssetsDiffs interface{} `json:"assets_diffs,required"`
	// This field can have the runtime type of
	// [map[string]TransactionScanResponseSimulationStellarSimulationResultSchemaExposure].
	Exposures interface{} `json:"exposures,required"`
	// This field can have the runtime type of
	// [map[string]TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsOwnershipDiff].
	AssetsOwnershipDiff interface{} `json:"assets_ownership_diff,required"`
	// This field can have the runtime type of
	// [[]TransactionScanResponseSimulationStellarSimulationResultSchemaAddressDetail].
	AddressDetails interface{} `json:"address_details,required"`
	// This field can have the runtime type of
	// [TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummary].
	AccountSummary interface{} `json:"account_summary,required"`
	// Error message
	Error string                                `json:"error"`
	JSON  transactionScanResponseSimulationJSON `json:"-"`
	union TransactionScanResponseSimulationUnion
}

// transactionScanResponseSimulationJSON contains the JSON metadata for the struct
// [TransactionScanResponseSimulation]
type transactionScanResponseSimulationJSON struct {
	Status              apijson.Field
	AssetsDiffs         apijson.Field
	Exposures           apijson.Field
	AssetsOwnershipDiff apijson.Field
	AddressDetails      apijson.Field
	AccountSummary      apijson.Field
	Error               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r transactionScanResponseSimulationJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionScanResponseSimulation) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TransactionScanResponseSimulationUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TransactionScanResponseSimulationStellarSimulationResultSchema],
// [TransactionScanResponseSimulationStellarSimulationErrorSchema].
func (r TransactionScanResponseSimulation) AsUnion() TransactionScanResponseSimulationUnion {
	return r.union
}

// Simulation result; Only present if simulation option is included in the request
//
// Union satisfied by
// [TransactionScanResponseSimulationStellarSimulationResultSchema] or
// [TransactionScanResponseSimulationStellarSimulationErrorSchema].
type TransactionScanResponseSimulationUnion interface {
	implementsTransactionScanResponseSimulation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionScanResponseSimulationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionScanResponseSimulationStellarSimulationResultSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionScanResponseSimulationStellarSimulationErrorSchema{}),
		},
	)
}

type TransactionScanResponseSimulationStellarSimulationResultSchema struct {
	// Summary of the actions and asset transfers that were made by the requested
	// account address
	AccountSummary TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummary `json:"account_summary,required"`
	Status         TransactionScanResponseSimulationStellarSimulationResultSchemaStatus         `json:"status,required"`
	// Details of addresses involved in the transaction
	AddressDetails []TransactionScanResponseSimulationStellarSimulationResultSchemaAddressDetail `json:"address_details"`
	// Mapping between the address of an account to the assets diff during the
	// transaction
	AssetsDiffs map[string][]TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiff `json:"assets_diffs"`
	// Mapping between the address of an account to the ownership diff of the account
	// during the transaction
	AssetsOwnershipDiff map[string]TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsOwnershipDiff `json:"assets_ownership_diff"`
	// Mapping between the address of an account to the exposure of the assets during
	// the transaction
	Exposures map[string]TransactionScanResponseSimulationStellarSimulationResultSchemaExposure `json:"exposures"`
	JSON      transactionScanResponseSimulationStellarSimulationResultSchemaJSON                `json:"-"`
}

// transactionScanResponseSimulationStellarSimulationResultSchemaJSON contains the
// JSON metadata for the struct
// [TransactionScanResponseSimulationStellarSimulationResultSchema]
type transactionScanResponseSimulationStellarSimulationResultSchemaJSON struct {
	AccountSummary      apijson.Field
	Status              apijson.Field
	AddressDetails      apijson.Field
	AssetsDiffs         apijson.Field
	AssetsOwnershipDiff apijson.Field
	Exposures           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *TransactionScanResponseSimulationStellarSimulationResultSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseSimulationStellarSimulationResultSchemaJSON) RawJSON() string {
	return r.raw
}

func (r TransactionScanResponseSimulationStellarSimulationResultSchema) implementsTransactionScanResponseSimulation() {
}

// Summary of the actions and asset transfers that were made by the requested
// account address
type TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummary struct {
	// Exposures made by the requested account address
	AccountExposures TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposures `json:"account_exposures,required"`
	// Account ownerships diff of the requested account address
	AccountOwnershipsDiff TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountOwnershipsDiff `json:"account_ownerships_diff,required"`
	// Total USD diff for the requested account address
	TotalUsdDiff TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryTotalUsdDiff `json:"total_usd_diff,required"`
	// Assets diffs of the requested account address
	AssetsDiffs []TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiff `json:"assets_diffs"`
	// Total USD exposure for each of the spender addresses during the transaction
	TotalUsdExposure map[string]float64                                                               `json:"total_usd_exposure"`
	JSON             transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryJSON `json:"-"`
}

// transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryJSON
// contains the JSON metadata for the struct
// [TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummary]
type transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryJSON struct {
	AccountExposures      apijson.Field
	AccountOwnershipsDiff apijson.Field
	TotalUsdDiff          apijson.Field
	AssetsDiffs           apijson.Field
	TotalUsdExposure      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryJSON) RawJSON() string {
	return r.raw
}

// Exposures made by the requested account address
type TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposures struct {
	Asset StellarAssetContractDetailsSchema `json:"asset,required"`
	// Mapping between the address of a Spender to the exposure of the asset during the
	// transaction
	Spenders map[string]TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposuresSpender `json:"spenders"`
	JSON     transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposuresJSON               `json:"-"`
}

// transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposuresJSON
// contains the JSON metadata for the struct
// [TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposures]
type transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposuresJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposures) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposuresJSON) RawJSON() string {
	return r.raw
}

type TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposuresSpender struct {
	// Raw value of the exposure
	RawValue int64 `json:"raw_value,required"`
	// Value of the exposure
	Value float64 `json:"value,required"`
	// Summarized description of the exposure
	Summary string                                                                                                  `json:"summary,nullable"`
	JSON    transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposuresSpenderJSON `json:"-"`
}

// transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposuresSpenderJSON
// contains the JSON metadata for the struct
// [TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposuresSpender]
type transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposuresSpenderJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposuresSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountExposuresSpenderJSON) RawJSON() string {
	return r.raw
}

// Account ownerships diff of the requested account address
type TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountOwnershipsDiff struct {
	// List of public keys that can sign on behalf of the account post-transaction
	PostSigners []string `json:"post_signers,required"`
	// List of public keys that can sign on behalf of the account pre-transaction
	PreSigners []string                                                                                              `json:"pre_signers,required"`
	JSON       transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountOwnershipsDiffJSON `json:"-"`
}

// transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountOwnershipsDiffJSON
// contains the JSON metadata for the struct
// [TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountOwnershipsDiff]
type transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountOwnershipsDiffJSON struct {
	PostSigners apijson.Field
	PreSigners  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountOwnershipsDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAccountOwnershipsDiffJSON) RawJSON() string {
	return r.raw
}

// Total USD diff for the requested account address
type TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryTotalUsdDiff struct {
	// Total incoming USD transfers
	In float64 `json:"in,required"`
	// Total outgoing USD transfers
	Out float64 `json:"out,required"`
	// Total USD transfers
	Total float64                                                                                      `json:"total"`
	JSON  transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryTotalUsdDiffJSON `json:"-"`
}

// transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryTotalUsdDiffJSON
// contains the JSON metadata for the struct
// [TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryTotalUsdDiff]
type transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

type TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiff struct {
	// Asset involved in the transfer
	Asset TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAsset `json:"asset,required"`
	// Incoming transfers of the asset
	In StellarAssetTransferDetailsSchema `json:"in,nullable"`
	// Outgoing transfers of the asset
	Out  StellarAssetTransferDetailsSchema                                                          `json:"out,nullable"`
	JSON transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffJSON `json:"-"`
}

// transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffJSON
// contains the JSON metadata for the struct
// [TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiff]
type transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffJSON) RawJSON() string {
	return r.raw
}

// Asset involved in the transfer
type TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAsset struct {
	// Type of the asset (`ASSET`)
	Type TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetType `json:"type"`
	// Asset code
	Code string `json:"code"`
	// Asset issuer address
	Issuer string `json:"issuer"`
	// Organization name
	OrgName string `json:"org_name"`
	// Organization URL
	OrgURL string `json:"org_url"`
	// Address of the asset's contract
	Address string `json:"address"`
	// Asset code
	Name string `json:"name"`
	// Asset symbol
	Symbol string                                                                                           `json:"symbol"`
	JSON   transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetJSON `json:"-"`
	union  TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetUnion
}

// transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetJSON
// contains the JSON metadata for the struct
// [TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAsset]
type transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetJSON struct {
	Type        apijson.Field
	Code        apijson.Field
	Issuer      apijson.Field
	OrgName     apijson.Field
	OrgURL      apijson.Field
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAsset) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchema],
// [TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchema],
// [StellarAssetContractDetailsSchema].
func (r TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAsset) AsUnion() TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetUnion {
	return r.union
}

// Asset involved in the transfer
//
// Union satisfied by
// [TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchema],
// [TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchema]
// or [StellarAssetContractDetailsSchema].
type TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetUnion interface {
	implementsTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarAssetContractDetailsSchema{}),
		},
	)
}

type TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchema struct {
	// Asset code
	Code string `json:"code,required"`
	// Asset issuer address
	Issuer string `json:"issuer,required"`
	// Organization name
	OrgName string `json:"org_name,required"`
	// Organization URL
	OrgURL string `json:"org_url,required"`
	// Type of the asset (`ASSET`)
	Type TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchemaType `json:"type"`
	JSON transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchemaJSON `json:"-"`
}

// transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchemaJSON
// contains the JSON metadata for the struct
// [TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchema]
type transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchemaJSON struct {
	Code        apijson.Field
	Issuer      apijson.Field
	OrgName     apijson.Field
	OrgURL      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchema) implementsTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAsset() {
}

// Type of the asset (`ASSET`)
type TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchemaType string

const (
	TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchemaTypeAsset TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchemaType = "ASSET"
)

func (r TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchemaType) IsKnown() bool {
	switch r {
	case TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarLegacyAssetDetailsSchemaTypeAsset:
		return true
	}
	return false
}

type TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchema struct {
	// Asset code
	Code TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaCode `json:"code"`
	// Type of the asset (`NATIVE`)
	Type TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaType `json:"type"`
	JSON transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaJSON `json:"-"`
}

// transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaJSON
// contains the JSON metadata for the struct
// [TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchema]
type transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaJSON struct {
	Code        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchema) implementsTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAsset() {
}

// Asset code
type TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaCode string

const (
	TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaCodeXlm TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaCode = "XLM"
)

func (r TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaCode) IsKnown() bool {
	switch r {
	case TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaCodeXlm:
		return true
	}
	return false
}

// Type of the asset (`NATIVE`)
type TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaType string

const (
	TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaTypeNative TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaType = "NATIVE"
)

func (r TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaType) IsKnown() bool {
	switch r {
	case TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetStellarNativeAssetDetailsSchemaTypeNative:
		return true
	}
	return false
}

// Type of the asset (`ASSET`)
type TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetType string

const (
	TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetTypeAsset    TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetType = "ASSET"
	TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetTypeNative   TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetType = "NATIVE"
	TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetTypeContract TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetType = "CONTRACT"
)

func (r TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetType) IsKnown() bool {
	switch r {
	case TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetTypeAsset, TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetTypeNative, TransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAssetTypeContract:
		return true
	}
	return false
}

type TransactionScanResponseSimulationStellarSimulationResultSchemaStatus string

const (
	TransactionScanResponseSimulationStellarSimulationResultSchemaStatusSuccess TransactionScanResponseSimulationStellarSimulationResultSchemaStatus = "Success"
)

func (r TransactionScanResponseSimulationStellarSimulationResultSchemaStatus) IsKnown() bool {
	switch r {
	case TransactionScanResponseSimulationStellarSimulationResultSchemaStatusSuccess:
		return true
	}
	return false
}

type TransactionScanResponseSimulationStellarSimulationResultSchemaAddressDetail struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// Description of the account
	Description string                                                                          `json:"description,nullable"`
	JSON        transactionScanResponseSimulationStellarSimulationResultSchemaAddressDetailJSON `json:"-"`
}

// transactionScanResponseSimulationStellarSimulationResultSchemaAddressDetailJSON
// contains the JSON metadata for the struct
// [TransactionScanResponseSimulationStellarSimulationResultSchemaAddressDetail]
type transactionScanResponseSimulationStellarSimulationResultSchemaAddressDetailJSON struct {
	AccountAddress apijson.Field
	Description    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionScanResponseSimulationStellarSimulationResultSchemaAddressDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseSimulationStellarSimulationResultSchemaAddressDetailJSON) RawJSON() string {
	return r.raw
}

type TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiff struct {
	// Asset involved in the transfer
	Asset TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAsset `json:"asset,required"`
	// Incoming transfers of the asset
	In StellarAssetTransferDetailsSchema `json:"in,nullable"`
	// Outgoing transfers of the asset
	Out  StellarAssetTransferDetailsSchema                                            `json:"out,nullable"`
	JSON transactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffJSON `json:"-"`
}

// transactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffJSON
// contains the JSON metadata for the struct
// [TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiff]
type transactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffJSON struct {
	Asset       apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffJSON) RawJSON() string {
	return r.raw
}

// Asset involved in the transfer
type TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAsset struct {
	// Type of the asset (`ASSET`)
	Type TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetType `json:"type"`
	// Asset code
	Code string `json:"code"`
	// Asset issuer address
	Issuer string `json:"issuer"`
	// Organization name
	OrgName string `json:"org_name"`
	// Organization URL
	OrgURL string `json:"org_url"`
	// Address of the asset's contract
	Address string `json:"address"`
	// Asset code
	Name string `json:"name"`
	// Asset symbol
	Symbol string                                                                             `json:"symbol"`
	JSON   transactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetJSON `json:"-"`
	union  TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetUnion
}

// transactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetJSON
// contains the JSON metadata for the struct
// [TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAsset]
type transactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetJSON struct {
	Type        apijson.Field
	Code        apijson.Field
	Issuer      apijson.Field
	OrgName     apijson.Field
	OrgURL      apijson.Field
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r transactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAsset) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchema],
// [TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchema],
// [StellarAssetContractDetailsSchema].
func (r TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAsset) AsUnion() TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetUnion {
	return r.union
}

// Asset involved in the transfer
//
// Union satisfied by
// [TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchema],
// [TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchema]
// or [StellarAssetContractDetailsSchema].
type TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetUnion interface {
	implementsTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAsset()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StellarAssetContractDetailsSchema{}),
		},
	)
}

type TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchema struct {
	// Asset code
	Code string `json:"code,required"`
	// Asset issuer address
	Issuer string `json:"issuer,required"`
	// Organization name
	OrgName string `json:"org_name,required"`
	// Organization URL
	OrgURL string `json:"org_url,required"`
	// Type of the asset (`ASSET`)
	Type TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchemaType `json:"type"`
	JSON transactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchemaJSON `json:"-"`
}

// transactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchemaJSON
// contains the JSON metadata for the struct
// [TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchema]
type transactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchemaJSON struct {
	Code        apijson.Field
	Issuer      apijson.Field
	OrgName     apijson.Field
	OrgURL      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchema) implementsTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAsset() {
}

// Type of the asset (`ASSET`)
type TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchemaType string

const (
	TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchemaTypeAsset TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchemaType = "ASSET"
)

func (r TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchemaType) IsKnown() bool {
	switch r {
	case TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarLegacyAssetDetailsSchemaTypeAsset:
		return true
	}
	return false
}

type TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchema struct {
	// Asset code
	Code TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaCode `json:"code"`
	// Type of the asset (`NATIVE`)
	Type TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaType `json:"type"`
	JSON transactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaJSON `json:"-"`
}

// transactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaJSON
// contains the JSON metadata for the struct
// [TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchema]
type transactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaJSON struct {
	Code        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchema) implementsTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAsset() {
}

// Asset code
type TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaCode string

const (
	TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaCodeXlm TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaCode = "XLM"
)

func (r TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaCode) IsKnown() bool {
	switch r {
	case TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaCodeXlm:
		return true
	}
	return false
}

// Type of the asset (`NATIVE`)
type TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaType string

const (
	TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaTypeNative TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaType = "NATIVE"
)

func (r TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaType) IsKnown() bool {
	switch r {
	case TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetStellarNativeAssetDetailsSchemaTypeNative:
		return true
	}
	return false
}

// Type of the asset (`ASSET`)
type TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetType string

const (
	TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetTypeAsset    TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetType = "ASSET"
	TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetTypeNative   TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetType = "NATIVE"
	TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetTypeContract TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetType = "CONTRACT"
)

func (r TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetType) IsKnown() bool {
	switch r {
	case TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetTypeAsset, TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetTypeNative, TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAssetTypeContract:
		return true
	}
	return false
}

type TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsOwnershipDiff struct {
	// List of public keys that can sign on behalf of the account post-transaction
	PostSigners []string `json:"post_signers,required"`
	// List of public keys that can sign on behalf of the account pre-transaction
	PreSigners []string                                                                              `json:"pre_signers,required"`
	JSON       transactionScanResponseSimulationStellarSimulationResultSchemaAssetsOwnershipDiffJSON `json:"-"`
}

// transactionScanResponseSimulationStellarSimulationResultSchemaAssetsOwnershipDiffJSON
// contains the JSON metadata for the struct
// [TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsOwnershipDiff]
type transactionScanResponseSimulationStellarSimulationResultSchemaAssetsOwnershipDiffJSON struct {
	PostSigners apijson.Field
	PreSigners  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionScanResponseSimulationStellarSimulationResultSchemaAssetsOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseSimulationStellarSimulationResultSchemaAssetsOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

type TransactionScanResponseSimulationStellarSimulationResultSchemaExposure struct {
	Asset StellarAssetContractDetailsSchema `json:"asset,required"`
	// Mapping between the address of a Spender to the exposure of the asset during the
	// transaction
	Spenders map[string]TransactionScanResponseSimulationStellarSimulationResultSchemaExposuresSpender `json:"spenders"`
	JSON     transactionScanResponseSimulationStellarSimulationResultSchemaExposureJSON                `json:"-"`
}

// transactionScanResponseSimulationStellarSimulationResultSchemaExposureJSON
// contains the JSON metadata for the struct
// [TransactionScanResponseSimulationStellarSimulationResultSchemaExposure]
type transactionScanResponseSimulationStellarSimulationResultSchemaExposureJSON struct {
	Asset       apijson.Field
	Spenders    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionScanResponseSimulationStellarSimulationResultSchemaExposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseSimulationStellarSimulationResultSchemaExposureJSON) RawJSON() string {
	return r.raw
}

type TransactionScanResponseSimulationStellarSimulationResultSchemaExposuresSpender struct {
	// Raw value of the exposure
	RawValue int64 `json:"raw_value,required"`
	// Value of the exposure
	Value float64 `json:"value,required"`
	// Summarized description of the exposure
	Summary string                                                                             `json:"summary,nullable"`
	JSON    transactionScanResponseSimulationStellarSimulationResultSchemaExposuresSpenderJSON `json:"-"`
}

// transactionScanResponseSimulationStellarSimulationResultSchemaExposuresSpenderJSON
// contains the JSON metadata for the struct
// [TransactionScanResponseSimulationStellarSimulationResultSchemaExposuresSpender]
type transactionScanResponseSimulationStellarSimulationResultSchemaExposuresSpenderJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionScanResponseSimulationStellarSimulationResultSchemaExposuresSpender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseSimulationStellarSimulationResultSchemaExposuresSpenderJSON) RawJSON() string {
	return r.raw
}

type TransactionScanResponseSimulationStellarSimulationErrorSchema struct {
	// Error message
	Error  string                                                              `json:"error,required"`
	Status TransactionScanResponseSimulationStellarSimulationErrorSchemaStatus `json:"status,required"`
	JSON   transactionScanResponseSimulationStellarSimulationErrorSchemaJSON   `json:"-"`
}

// transactionScanResponseSimulationStellarSimulationErrorSchemaJSON contains the
// JSON metadata for the struct
// [TransactionScanResponseSimulationStellarSimulationErrorSchema]
type transactionScanResponseSimulationStellarSimulationErrorSchemaJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionScanResponseSimulationStellarSimulationErrorSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseSimulationStellarSimulationErrorSchemaJSON) RawJSON() string {
	return r.raw
}

func (r TransactionScanResponseSimulationStellarSimulationErrorSchema) implementsTransactionScanResponseSimulation() {
}

type TransactionScanResponseSimulationStellarSimulationErrorSchemaStatus string

const (
	TransactionScanResponseSimulationStellarSimulationErrorSchemaStatusError TransactionScanResponseSimulationStellarSimulationErrorSchemaStatus = "Error"
)

func (r TransactionScanResponseSimulationStellarSimulationErrorSchemaStatus) IsKnown() bool {
	switch r {
	case TransactionScanResponseSimulationStellarSimulationErrorSchemaStatusError:
		return true
	}
	return false
}

type TransactionScanResponseSimulationStatus string

const (
	TransactionScanResponseSimulationStatusSuccess TransactionScanResponseSimulationStatus = "Success"
	TransactionScanResponseSimulationStatusError   TransactionScanResponseSimulationStatus = "Error"
)

func (r TransactionScanResponseSimulationStatus) IsKnown() bool {
	switch r {
	case TransactionScanResponseSimulationStatusSuccess, TransactionScanResponseSimulationStatusError:
		return true
	}
	return false
}

// Validation result; Only present if validation option is included in the request
type TransactionScanResponseValidation struct {
	Status TransactionScanResponseValidationStatus `json:"status,required"`
	// Verdict of the validation
	ResultType TransactionScanResponseValidationResultType `json:"result_type"`
	// A textual description about the validation result
	Description string `json:"description"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason TransactionScanResponseValidationReason `json:"reason"`
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification"`
	// This field can have the runtime type of
	// [[]TransactionScanResponseValidationStellarValidationResultSchemaFeature].
	Features interface{} `json:"features,required"`
	// Error message
	Error string                                `json:"error"`
	JSON  transactionScanResponseValidationJSON `json:"-"`
	union TransactionScanResponseValidationUnion
}

// transactionScanResponseValidationJSON contains the JSON metadata for the struct
// [TransactionScanResponseValidation]
type transactionScanResponseValidationJSON struct {
	Status         apijson.Field
	ResultType     apijson.Field
	Description    apijson.Field
	Reason         apijson.Field
	Classification apijson.Field
	Features       apijson.Field
	Error          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r transactionScanResponseValidationJSON) RawJSON() string {
	return r.raw
}

func (r *TransactionScanResponseValidation) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TransactionScanResponseValidationUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TransactionScanResponseValidationStellarValidationResultSchema],
// [TransactionScanResponseValidationStellarValidationErrorSchema].
func (r TransactionScanResponseValidation) AsUnion() TransactionScanResponseValidationUnion {
	return r.union
}

// Validation result; Only present if validation option is included in the request
//
// Union satisfied by
// [TransactionScanResponseValidationStellarValidationResultSchema] or
// [TransactionScanResponseValidationStellarValidationErrorSchema].
type TransactionScanResponseValidationUnion interface {
	implementsTransactionScanResponseValidation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TransactionScanResponseValidationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionScanResponseValidationStellarValidationResultSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionScanResponseValidationStellarValidationErrorSchema{}),
		},
	)
}

type TransactionScanResponseValidationStellarValidationResultSchema struct {
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification,required"`
	// A textual description about the validation result
	Description string `json:"description,required"`
	// A list of features about this transaction explaining the validation
	Features []TransactionScanResponseValidationStellarValidationResultSchemaFeature `json:"features,required"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason TransactionScanResponseValidationStellarValidationResultSchemaReason `json:"reason,required"`
	// Verdict of the validation
	ResultType TransactionScanResponseValidationStellarValidationResultSchemaResultType `json:"result_type,required"`
	Status     TransactionScanResponseValidationStellarValidationResultSchemaStatus     `json:"status,required"`
	JSON       transactionScanResponseValidationStellarValidationResultSchemaJSON       `json:"-"`
}

// transactionScanResponseValidationStellarValidationResultSchemaJSON contains the
// JSON metadata for the struct
// [TransactionScanResponseValidationStellarValidationResultSchema]
type transactionScanResponseValidationStellarValidationResultSchemaJSON struct {
	Classification apijson.Field
	Description    apijson.Field
	Features       apijson.Field
	Reason         apijson.Field
	ResultType     apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionScanResponseValidationStellarValidationResultSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseValidationStellarValidationResultSchemaJSON) RawJSON() string {
	return r.raw
}

func (r TransactionScanResponseValidationStellarValidationResultSchema) implementsTransactionScanResponseValidation() {
}

type TransactionScanResponseValidationStellarValidationResultSchemaFeature struct {
	// Address the feature refers to
	Address string `json:"address,required"`
	// Textual description
	Description string `json:"description,required"`
	FeatureID string `json:"feature_id,required"`
	// Feature Classification
	Type TransactionScanResponseValidationStellarValidationResultSchemaFeaturesType `json:"type,required"`
	JSON transactionScanResponseValidationStellarValidationResultSchemaFeatureJSON  `json:"-"`
}

// transactionScanResponseValidationStellarValidationResultSchemaFeatureJSON
// contains the JSON metadata for the struct
// [TransactionScanResponseValidationStellarValidationResultSchemaFeature]
type transactionScanResponseValidationStellarValidationResultSchemaFeatureJSON struct {
	Address     apijson.Field
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionScanResponseValidationStellarValidationResultSchemaFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseValidationStellarValidationResultSchemaFeatureJSON) RawJSON() string {
	return r.raw
}

// Feature Classification
type TransactionScanResponseValidationStellarValidationResultSchemaFeaturesType string

const (
	TransactionScanResponseValidationStellarValidationResultSchemaFeaturesTypeBenign    TransactionScanResponseValidationStellarValidationResultSchemaFeaturesType = "Benign"
	TransactionScanResponseValidationStellarValidationResultSchemaFeaturesTypeWarning   TransactionScanResponseValidationStellarValidationResultSchemaFeaturesType = "Warning"
	TransactionScanResponseValidationStellarValidationResultSchemaFeaturesTypeMalicious TransactionScanResponseValidationStellarValidationResultSchemaFeaturesType = "Malicious"
	TransactionScanResponseValidationStellarValidationResultSchemaFeaturesTypeInfo      TransactionScanResponseValidationStellarValidationResultSchemaFeaturesType = "Info"
)

func (r TransactionScanResponseValidationStellarValidationResultSchemaFeaturesType) IsKnown() bool {
	switch r {
	case TransactionScanResponseValidationStellarValidationResultSchemaFeaturesTypeBenign, TransactionScanResponseValidationStellarValidationResultSchemaFeaturesTypeWarning, TransactionScanResponseValidationStellarValidationResultSchemaFeaturesTypeMalicious, TransactionScanResponseValidationStellarValidationResultSchemaFeaturesTypeInfo:
		return true
	}
	return false
}

// A textual description about the reasons the transaction was flagged with
// result_type
type TransactionScanResponseValidationStellarValidationResultSchemaReason string

const (
	TransactionScanResponseValidationStellarValidationResultSchemaReasonEmpty                 TransactionScanResponseValidationStellarValidationResultSchemaReason = ""
	TransactionScanResponseValidationStellarValidationResultSchemaReasonKnownAttacker         TransactionScanResponseValidationStellarValidationResultSchemaReason = "known_attacker"
	TransactionScanResponseValidationStellarValidationResultSchemaReasonKnownFraudulentAsset  TransactionScanResponseValidationStellarValidationResultSchemaReason = "known_fraudulent_asset"
	TransactionScanResponseValidationStellarValidationResultSchemaReasonMaliciousMemo         TransactionScanResponseValidationStellarValidationResultSchemaReason = "malicious_memo"
	TransactionScanResponseValidationStellarValidationResultSchemaReasonUnfairTrade           TransactionScanResponseValidationStellarValidationResultSchemaReason = "unfair_trade"
	TransactionScanResponseValidationStellarValidationResultSchemaReasonTransferFarming       TransactionScanResponseValidationStellarValidationResultSchemaReason = "transfer_farming"
	TransactionScanResponseValidationStellarValidationResultSchemaReasonNativeOwnershipChange TransactionScanResponseValidationStellarValidationResultSchemaReason = "native_ownership_change"
	TransactionScanResponseValidationStellarValidationResultSchemaReasonOther                 TransactionScanResponseValidationStellarValidationResultSchemaReason = "other"
)

func (r TransactionScanResponseValidationStellarValidationResultSchemaReason) IsKnown() bool {
	switch r {
	case TransactionScanResponseValidationStellarValidationResultSchemaReasonEmpty, TransactionScanResponseValidationStellarValidationResultSchemaReasonKnownAttacker, TransactionScanResponseValidationStellarValidationResultSchemaReasonKnownFraudulentAsset, TransactionScanResponseValidationStellarValidationResultSchemaReasonMaliciousMemo, TransactionScanResponseValidationStellarValidationResultSchemaReasonUnfairTrade, TransactionScanResponseValidationStellarValidationResultSchemaReasonTransferFarming, TransactionScanResponseValidationStellarValidationResultSchemaReasonNativeOwnershipChange, TransactionScanResponseValidationStellarValidationResultSchemaReasonOther:
		return true
	}
	return false
}

// Verdict of the validation
type TransactionScanResponseValidationStellarValidationResultSchemaResultType string

const (
	TransactionScanResponseValidationStellarValidationResultSchemaResultTypeBenign    TransactionScanResponseValidationStellarValidationResultSchemaResultType = "Benign"
	TransactionScanResponseValidationStellarValidationResultSchemaResultTypeWarning   TransactionScanResponseValidationStellarValidationResultSchemaResultType = "Warning"
	TransactionScanResponseValidationStellarValidationResultSchemaResultTypeMalicious TransactionScanResponseValidationStellarValidationResultSchemaResultType = "Malicious"
)

func (r TransactionScanResponseValidationStellarValidationResultSchemaResultType) IsKnown() bool {
	switch r {
	case TransactionScanResponseValidationStellarValidationResultSchemaResultTypeBenign, TransactionScanResponseValidationStellarValidationResultSchemaResultTypeWarning, TransactionScanResponseValidationStellarValidationResultSchemaResultTypeMalicious:
		return true
	}
	return false
}

type TransactionScanResponseValidationStellarValidationResultSchemaStatus string

const (
	TransactionScanResponseValidationStellarValidationResultSchemaStatusSuccess TransactionScanResponseValidationStellarValidationResultSchemaStatus = "Success"
)

func (r TransactionScanResponseValidationStellarValidationResultSchemaStatus) IsKnown() bool {
	switch r {
	case TransactionScanResponseValidationStellarValidationResultSchemaStatusSuccess:
		return true
	}
	return false
}

type TransactionScanResponseValidationStellarValidationErrorSchema struct {
	// Error message
	Error  string                                                              `json:"error,required"`
	Status TransactionScanResponseValidationStellarValidationErrorSchemaStatus `json:"status,required"`
	JSON   transactionScanResponseValidationStellarValidationErrorSchemaJSON   `json:"-"`
}

// transactionScanResponseValidationStellarValidationErrorSchemaJSON contains the
// JSON metadata for the struct
// [TransactionScanResponseValidationStellarValidationErrorSchema]
type transactionScanResponseValidationStellarValidationErrorSchemaJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionScanResponseValidationStellarValidationErrorSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionScanResponseValidationStellarValidationErrorSchemaJSON) RawJSON() string {
	return r.raw
}

func (r TransactionScanResponseValidationStellarValidationErrorSchema) implementsTransactionScanResponseValidation() {
}

type TransactionScanResponseValidationStellarValidationErrorSchemaStatus string

const (
	TransactionScanResponseValidationStellarValidationErrorSchemaStatusError TransactionScanResponseValidationStellarValidationErrorSchemaStatus = "Error"
)

func (r TransactionScanResponseValidationStellarValidationErrorSchemaStatus) IsKnown() bool {
	switch r {
	case TransactionScanResponseValidationStellarValidationErrorSchemaStatusError:
		return true
	}
	return false
}

type TransactionScanResponseValidationStatus string

const (
	TransactionScanResponseValidationStatusSuccess TransactionScanResponseValidationStatus = "Success"
	TransactionScanResponseValidationStatusError   TransactionScanResponseValidationStatus = "Error"
)

func (r TransactionScanResponseValidationStatus) IsKnown() bool {
	switch r {
	case TransactionScanResponseValidationStatusSuccess, TransactionScanResponseValidationStatusError:
		return true
	}
	return false
}

// Verdict of the validation
type TransactionScanResponseValidationResultType string

const (
	TransactionScanResponseValidationResultTypeBenign    TransactionScanResponseValidationResultType = "Benign"
	TransactionScanResponseValidationResultTypeWarning   TransactionScanResponseValidationResultType = "Warning"
	TransactionScanResponseValidationResultTypeMalicious TransactionScanResponseValidationResultType = "Malicious"
)

func (r TransactionScanResponseValidationResultType) IsKnown() bool {
	switch r {
	case TransactionScanResponseValidationResultTypeBenign, TransactionScanResponseValidationResultTypeWarning, TransactionScanResponseValidationResultTypeMalicious:
		return true
	}
	return false
}

// A textual description about the reasons the transaction was flagged with
// result_type
type TransactionScanResponseValidationReason string

const (
	TransactionScanResponseValidationReasonEmpty                 TransactionScanResponseValidationReason = ""
	TransactionScanResponseValidationReasonKnownAttacker         TransactionScanResponseValidationReason = "known_attacker"
	TransactionScanResponseValidationReasonKnownFraudulentAsset  TransactionScanResponseValidationReason = "known_fraudulent_asset"
	TransactionScanResponseValidationReasonMaliciousMemo         TransactionScanResponseValidationReason = "malicious_memo"
	TransactionScanResponseValidationReasonUnfairTrade           TransactionScanResponseValidationReason = "unfair_trade"
	TransactionScanResponseValidationReasonTransferFarming       TransactionScanResponseValidationReason = "transfer_farming"
	TransactionScanResponseValidationReasonNativeOwnershipChange TransactionScanResponseValidationReason = "native_ownership_change"
	TransactionScanResponseValidationReasonOther                 TransactionScanResponseValidationReason = "other"
)

func (r TransactionScanResponseValidationReason) IsKnown() bool {
	switch r {
	case TransactionScanResponseValidationReasonEmpty, TransactionScanResponseValidationReasonKnownAttacker, TransactionScanResponseValidationReasonKnownFraudulentAsset, TransactionScanResponseValidationReasonMaliciousMemo, TransactionScanResponseValidationReasonUnfairTrade, TransactionScanResponseValidationReasonTransferFarming, TransactionScanResponseValidationReasonNativeOwnershipChange, TransactionScanResponseValidationReasonOther:
		return true
	}
	return false
}
