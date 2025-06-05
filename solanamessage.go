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

// SolanaMessageService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSolanaMessageService] method instead.
type SolanaMessageService struct {
	Options []option.RequestOption
}

// NewSolanaMessageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSolanaMessageService(opts ...option.RequestOption) (r *SolanaMessageService) {
	r = &SolanaMessageService{}
	r.Options = opts
	return
}

// Scan Message
func (r *SolanaMessageService) Scan(ctx context.Context, body SolanaMessageScanParams, opts ...option.RequestOption) (res *SolanaMessageScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/solana/message/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SolanaMessageScanResponse struct {
	Encoding SolanaMessageScanResponseEncoding `json:"encoding,required"`
	// Unique identifier of the request
	RequestID string                          `json:"request_id,required"`
	Status    SolanaMessageScanResponseStatus `json:"status,required"`
	// Error message if the simulation failed
	Error string `json:"error,nullable"`
	// Error details
	ErrorDetails SolanaMessageScanResponseErrorDetails `json:"error_details,nullable"`
	// Result of the request
	Result SolanaMessageScanResponseResult `json:"result,nullable"`
	JSON   solanaMessageScanResponseJSON   `json:"-"`
}

// solanaMessageScanResponseJSON contains the JSON metadata for the struct
// [SolanaMessageScanResponse]
type solanaMessageScanResponseJSON struct {
	Encoding     apijson.Field
	RequestID    apijson.Field
	Status       apijson.Field
	Error        apijson.Field
	ErrorDetails apijson.Field
	Result       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SolanaMessageScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseEncoding string

const (
	SolanaMessageScanResponseEncodingBase58 SolanaMessageScanResponseEncoding = "base58"
	SolanaMessageScanResponseEncodingBase64 SolanaMessageScanResponseEncoding = "base64"
)

func (r SolanaMessageScanResponseEncoding) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseEncodingBase58, SolanaMessageScanResponseEncodingBase64:
		return true
	}
	return false
}

type SolanaMessageScanResponseStatus string

const (
	SolanaMessageScanResponseStatusSuccess SolanaMessageScanResponseStatus = "SUCCESS"
	SolanaMessageScanResponseStatusError   SolanaMessageScanResponseStatus = "ERROR"
)

func (r SolanaMessageScanResponseStatus) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseStatusSuccess, SolanaMessageScanResponseStatusError:
		return true
	}
	return false
}

// Error details
type SolanaMessageScanResponseErrorDetails struct {
	// Advanced message of the error
	Message string `json:"message,required"`
	// Machine readable error code
	Code string `json:"code,nullable"`
	// Index of the instruction in the transaction
	InstructionIndex int64 `json:"instruction_index"`
	// Error number
	Number int64 `json:"number,nullable"`
	// The program account that caused the error
	ProgramAccount string `json:"program_account,nullable"`
	// Index of the transaction in the bulk
	TransactionIndex int64                                     `json:"transaction_index"`
	Type             SolanaMessageScanResponseErrorDetailsType `json:"type"`
	JSON             solanaMessageScanResponseErrorDetailsJSON `json:"-"`
	union            SolanaMessageScanResponseErrorDetailsUnion
}

// solanaMessageScanResponseErrorDetailsJSON contains the JSON metadata for the
// struct [SolanaMessageScanResponseErrorDetails]
type solanaMessageScanResponseErrorDetailsJSON struct {
	Message          apijson.Field
	Code             apijson.Field
	InstructionIndex apijson.Field
	Number           apijson.Field
	ProgramAccount   apijson.Field
	TransactionIndex apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r solanaMessageScanResponseErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseErrorDetails) UnmarshalJSON(data []byte) (err error) {
	*r = SolanaMessageScanResponseErrorDetails{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SolanaMessageScanResponseErrorDetailsUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SolanaMessageScanResponseErrorDetailsSolanaAPIErrorDetails],
// [SolanaMessageScanResponseErrorDetailsSolanaTransactionErrorDetails],
// [SolanaMessageScanResponseErrorDetailsSolanaInstructionErrorDetails].
func (r SolanaMessageScanResponseErrorDetails) AsUnion() SolanaMessageScanResponseErrorDetailsUnion {
	return r.union
}

// Error details
//
// Union satisfied by [SolanaMessageScanResponseErrorDetailsSolanaAPIErrorDetails],
// [SolanaMessageScanResponseErrorDetailsSolanaTransactionErrorDetails] or
// [SolanaMessageScanResponseErrorDetailsSolanaInstructionErrorDetails].
type SolanaMessageScanResponseErrorDetailsUnion interface {
	implementsSolanaMessageScanResponseErrorDetails()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseErrorDetailsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SolanaMessageScanResponseErrorDetailsSolanaAPIErrorDetails{}),
			DiscriminatorValue: "ApiError",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SolanaMessageScanResponseErrorDetailsSolanaTransactionErrorDetails{}),
			DiscriminatorValue: "TransactionError",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SolanaMessageScanResponseErrorDetailsSolanaInstructionErrorDetails{}),
			DiscriminatorValue: "InstructionError",
		},
	)
}

type SolanaMessageScanResponseErrorDetailsSolanaAPIErrorDetails struct {
	// Advanced message of the error
	Message string                                                         `json:"message,required"`
	Type    SolanaMessageScanResponseErrorDetailsSolanaAPIErrorDetailsType `json:"type"`
	JSON    solanaMessageScanResponseErrorDetailsSolanaAPIErrorDetailsJSON `json:"-"`
}

// solanaMessageScanResponseErrorDetailsSolanaAPIErrorDetailsJSON contains the JSON
// metadata for the struct
// [SolanaMessageScanResponseErrorDetailsSolanaAPIErrorDetails]
type solanaMessageScanResponseErrorDetailsSolanaAPIErrorDetailsJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseErrorDetailsSolanaAPIErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseErrorDetailsSolanaAPIErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseErrorDetailsSolanaAPIErrorDetails) implementsSolanaMessageScanResponseErrorDetails() {
}

type SolanaMessageScanResponseErrorDetailsSolanaAPIErrorDetailsType string

const (
	SolanaMessageScanResponseErrorDetailsSolanaAPIErrorDetailsTypeAPIError SolanaMessageScanResponseErrorDetailsSolanaAPIErrorDetailsType = "ApiError"
)

func (r SolanaMessageScanResponseErrorDetailsSolanaAPIErrorDetailsType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseErrorDetailsSolanaAPIErrorDetailsTypeAPIError:
		return true
	}
	return false
}

type SolanaMessageScanResponseErrorDetailsSolanaTransactionErrorDetails struct {
	// Advanced message of the error
	Message string `json:"message,required"`
	// Index of the transaction in the bulk
	TransactionIndex int64                                                                  `json:"transaction_index,required"`
	Type             SolanaMessageScanResponseErrorDetailsSolanaTransactionErrorDetailsType `json:"type"`
	JSON             solanaMessageScanResponseErrorDetailsSolanaTransactionErrorDetailsJSON `json:"-"`
}

// solanaMessageScanResponseErrorDetailsSolanaTransactionErrorDetailsJSON contains
// the JSON metadata for the struct
// [SolanaMessageScanResponseErrorDetailsSolanaTransactionErrorDetails]
type solanaMessageScanResponseErrorDetailsSolanaTransactionErrorDetailsJSON struct {
	Message          apijson.Field
	TransactionIndex apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SolanaMessageScanResponseErrorDetailsSolanaTransactionErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseErrorDetailsSolanaTransactionErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseErrorDetailsSolanaTransactionErrorDetails) implementsSolanaMessageScanResponseErrorDetails() {
}

type SolanaMessageScanResponseErrorDetailsSolanaTransactionErrorDetailsType string

const (
	SolanaMessageScanResponseErrorDetailsSolanaTransactionErrorDetailsTypeTransactionError SolanaMessageScanResponseErrorDetailsSolanaTransactionErrorDetailsType = "TransactionError"
)

func (r SolanaMessageScanResponseErrorDetailsSolanaTransactionErrorDetailsType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseErrorDetailsSolanaTransactionErrorDetailsTypeTransactionError:
		return true
	}
	return false
}

type SolanaMessageScanResponseErrorDetailsSolanaInstructionErrorDetails struct {
	// Index of the instruction in the transaction
	InstructionIndex int64 `json:"instruction_index,required"`
	// Human readable error
	Message string `json:"message,required"`
	// Index of the transaction in the bulk
	TransactionIndex int64 `json:"transaction_index,required"`
	// Machine readable error code
	Code string `json:"code,nullable"`
	// Error number
	Number int64 `json:"number,nullable"`
	// The program account that caused the error
	ProgramAccount string                                                                 `json:"program_account,nullable"`
	Type           SolanaMessageScanResponseErrorDetailsSolanaInstructionErrorDetailsType `json:"type"`
	JSON           solanaMessageScanResponseErrorDetailsSolanaInstructionErrorDetailsJSON `json:"-"`
}

// solanaMessageScanResponseErrorDetailsSolanaInstructionErrorDetailsJSON contains
// the JSON metadata for the struct
// [SolanaMessageScanResponseErrorDetailsSolanaInstructionErrorDetails]
type solanaMessageScanResponseErrorDetailsSolanaInstructionErrorDetailsJSON struct {
	InstructionIndex apijson.Field
	Message          apijson.Field
	TransactionIndex apijson.Field
	Code             apijson.Field
	Number           apijson.Field
	ProgramAccount   apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SolanaMessageScanResponseErrorDetailsSolanaInstructionErrorDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseErrorDetailsSolanaInstructionErrorDetailsJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseErrorDetailsSolanaInstructionErrorDetails) implementsSolanaMessageScanResponseErrorDetails() {
}

type SolanaMessageScanResponseErrorDetailsSolanaInstructionErrorDetailsType string

const (
	SolanaMessageScanResponseErrorDetailsSolanaInstructionErrorDetailsTypeInstructionError SolanaMessageScanResponseErrorDetailsSolanaInstructionErrorDetailsType = "InstructionError"
)

func (r SolanaMessageScanResponseErrorDetailsSolanaInstructionErrorDetailsType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseErrorDetailsSolanaInstructionErrorDetailsTypeInstructionError:
		return true
	}
	return false
}

type SolanaMessageScanResponseErrorDetailsType string

const (
	SolanaMessageScanResponseErrorDetailsTypeAPIError         SolanaMessageScanResponseErrorDetailsType = "ApiError"
	SolanaMessageScanResponseErrorDetailsTypeTransactionError SolanaMessageScanResponseErrorDetailsType = "TransactionError"
	SolanaMessageScanResponseErrorDetailsTypeInstructionError SolanaMessageScanResponseErrorDetailsType = "InstructionError"
)

func (r SolanaMessageScanResponseErrorDetailsType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseErrorDetailsTypeAPIError, SolanaMessageScanResponseErrorDetailsTypeTransactionError, SolanaMessageScanResponseErrorDetailsTypeInstructionError:
		return true
	}
	return false
}

// Result of the request
type SolanaMessageScanResponseResult struct {
	// Transaction Simulation Result
	Simulation SolanaMessageScanResponseResultSimulation `json:"simulation,required,nullable"`
	// Transaction Validation Result
	Validation SolanaMessageScanResponseResultValidation `json:"validation,required,nullable"`
	JSON       solanaMessageScanResponseResultJSON       `json:"-"`
}

// solanaMessageScanResponseResultJSON contains the JSON metadata for the struct
// [SolanaMessageScanResponseResult]
type solanaMessageScanResponseResultJSON struct {
	Simulation  apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultJSON) RawJSON() string {
	return r.raw
}

// Transaction Simulation Result
type SolanaMessageScanResponseResultSimulation struct {
	// Summary of the actions and asset transfers that were made by the requested
	// account address
	AccountSummary SolanaMessageScanResponseResultSimulationAccountSummary `json:"account_summary,required"`
	// Ownership diffs of the account addresses
	AssetsOwnershipDiff map[string][]SolanaMessageScanResponseResultSimulationAssetsOwnershipDiff `json:"assets_ownership_diff,required"`
	// Details of addresses involved in the transaction
	AccountsDetails []SolanaMessageScanResponseResultSimulationAccountsDetail `json:"accounts_details"`
	// Mapping between the address of an account to the assets diff during the
	// transaction
	AssetsDiff map[string][]SolanaMessageScanResponseResultSimulationAssetsDiff `json:"assets_diff"`
	// Mapping between the address of an account to the exposure of the assets during
	// the transaction
	Delegations map[string][]SolanaMessageScanResponseResultSimulationDelegation `json:"delegations"`
	JSON        solanaMessageScanResponseResultSimulationJSON                    `json:"-"`
}

// solanaMessageScanResponseResultSimulationJSON contains the JSON metadata for the
// struct [SolanaMessageScanResponseResultSimulation]
type solanaMessageScanResponseResultSimulationJSON struct {
	AccountSummary      apijson.Field
	AssetsOwnershipDiff apijson.Field
	AccountsDetails     apijson.Field
	AssetsDiff          apijson.Field
	Delegations         apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationJSON) RawJSON() string {
	return r.raw
}

// Summary of the actions and asset transfers that were made by the requested
// account address
type SolanaMessageScanResponseResultSimulationAccountSummary struct {
	// Exposures made by the requested account address
	AccountDelegations []SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegation `json:"account_delegations,required"`
	// Ownership diffs of the requested account address
	AccountOwnershipsDiff []SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiff `json:"account_ownerships_diff,required"`
	// Total USD diff for the requested account address
	TotalUsdDiff SolanaMessageScanResponseResultSimulationAccountSummaryTotalUsdDiff `json:"total_usd_diff,required"`
	// Assets diffs of the requested account address
	AccountAssetsDiff []SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiff `json:"account_assets_diff"`
	// Total USD exposure for each of the spender addresses during the transaction
	TotalUsdExposure map[string]float64                                          `json:"total_usd_exposure"`
	JSON             solanaMessageScanResponseResultSimulationAccountSummaryJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryJSON contains the JSON
// metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummary]
type solanaMessageScanResponseResultSimulationAccountSummaryJSON struct {
	AccountDelegations    apijson.Field
	AccountOwnershipsDiff apijson.Field
	TotalUsdDiff          apijson.Field
	AccountAssetsDiff     apijson.Field
	TotalUsdExposure      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegation struct {
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationAsset],
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAsset],
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAsset].
	Asset     interface{} `json:"asset,required"`
	AssetType string      `json:"asset_type,required"`
	Delegate  string      `json:"delegate,required"`
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationDelegation],
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationDelegation],
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationDelegation].
	Delegation interface{}                                                                  `json:"delegation,required"`
	JSON       solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationJSON `json:"-"`
	union      SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsUnion
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegation]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Delegate    apijson.Field
	Delegation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegation) UnmarshalJSON(data []byte) (err error) {
	*r = SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegation],
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegation],
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegation].
func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegation) AsUnion() SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsUnion {
	return r.union
}

// Union satisfied by
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegation],
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegation]
// or
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegation].
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsUnion interface {
	implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegation{}),
		},
	)
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegation struct {
	Asset      SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationAsset      `json:"asset,required"`
	AssetType  string                                                                                                  `json:"asset_type,required"`
	Delegate   string                                                                                                  `json:"delegate,required"`
	Delegation SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationDelegation `json:"delegation,required"`
	JSON       solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationJSON       `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegation]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Delegate    apijson.Field
	Delegation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegation) implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegation() {
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	// URL of the asset's logo
	Logo string                                                                                                 `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationAsset]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationAssetType string

const (
	SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationAssetTypeCnft SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationAssetType = "CNFT"
)

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationAssetTypeCnft:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationDelegation struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                     `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationDelegationJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationDelegation]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationDelegationJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaCnftDelegationDelegationJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegation struct {
	Asset      SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAsset      `json:"asset,required"`
	AssetType  string                                                                                                              `json:"asset_type,required"`
	Delegate   string                                                                                                              `json:"delegate,required"`
	Delegation SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationDelegation `json:"delegation,required"`
	JSON       solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationJSON       `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegation]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Delegate    apijson.Field
	Delegation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegation) implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegation() {
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's decimals
	Decimals int64 `json:"decimals,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	Logo string                                                                                                             `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAsset]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAssetType string

const (
	SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAssetTypeToken SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAssetType = "TOKEN"
)

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAssetTypeToken:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationDelegation struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                 `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationDelegationJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationDelegation]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationDelegationJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationDelegationJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegation struct {
	Asset      SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAsset      `json:"asset,required"`
	AssetType  string                                                                                                                 `json:"asset_type,required"`
	Delegate   string                                                                                                                 `json:"delegate,required"`
	Delegation SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationDelegation `json:"delegation,required"`
	JSON       solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationJSON       `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegation]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Delegate    apijson.Field
	Delegation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegation) implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegation() {
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	// URL of the asset's logo
	Logo string                                                                                                                `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAsset]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAssetType string

const (
	SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAssetTypeNFT SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAssetType = "NFT"
)

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAssetTypeNFT:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationDelegation struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                    `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationDelegationJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationDelegation]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationDelegationJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationDelegationJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiff struct {
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAsset],
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAsset],
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAsset],
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAsset].
	Asset interface{} `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required,nullable"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,required,nullable"`
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffIn],
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffIn],
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffIn],
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffIn].
	In interface{} `json:"in"`
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffOut],
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffOut],
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffOut],
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffOut].
	Out   interface{}                                                                      `json:"out"`
	JSON  solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffJSON `json:"-"`
	union SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffUnion
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiff]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	PostOwner   apijson.Field
	PreOwner    apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiff],
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiff],
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiff],
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiff].
func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiff) AsUnion() SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffUnion {
	return r.union
}

// Union satisfied by
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiff],
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiff],
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiff]
// or
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiff].
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffUnion interface {
	implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiff{}),
		},
	)
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiff struct {
	Asset SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required,nullable"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,required,nullable"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiff]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	PostOwner   apijson.Field
	PreOwner    apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiff) implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiff() {
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAsset struct {
	Decimals int64 `json:"decimals,required"`
	// Type of the asset (`"NativeToken"`)
	Type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetType `json:"type,required"`
	// Logo of SOL
	Logo string                                                                                                            `json:"logo,nullable"`
	JSON solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAsset]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	Logo        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`"NativeToken"`)
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetTypeSol SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetType = "SOL"
	SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetTypeEth SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetType = "ETH"
)

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetTypeSol, SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetTypeEth:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                        `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffIn]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                         `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffOut]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiff struct {
	Asset SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required,nullable"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,required,nullable"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiff]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	PostOwner   apijson.Field
	PreOwner    apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiff) implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiff() {
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAsset struct {
	Decimals int64 `json:"decimals,required"`
	// Type of the asset (`"StakedNative"`)
	Type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetType `json:"type,required"`
	// Logo of SOL
	Logo string                                                                                                            `json:"logo,nullable"`
	JSON solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAsset]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	Logo        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`"StakedNative"`)
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetTypeStakedSol SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetType = "STAKED_SOL"
	SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetTypeStakedEth SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetType = "STAKED_ETH"
)

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetTypeStakedSol, SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetTypeStakedEth:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                        `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffIn]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                         `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffOut]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiff struct {
	Asset SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required,nullable"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,required,nullable"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiff]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	PostOwner   apijson.Field
	PreOwner    apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiff) implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiff() {
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's decimals
	Decimals int64 `json:"decimals,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	Logo string                                                                                                              `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAsset]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAssetTypeToken SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAssetType = "TOKEN"
)

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAssetTypeToken:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                          `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffIn]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                           `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffOut]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiff struct {
	Asset SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required,nullable"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,required,nullable"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiff]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	PostOwner   apijson.Field
	PreOwner    apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiff) implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiff() {
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	// URL of the asset's logo
	Logo string                                                                                                                 `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAsset]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAssetTypeNFT SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAssetType = "NFT"
)

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAssetTypeNFT:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                             `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffIn]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                              `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffOut]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffOutJSON) RawJSON() string {
	return r.raw
}

// Total USD diff for the requested account address
type SolanaMessageScanResponseResultSimulationAccountSummaryTotalUsdDiff struct {
	// Total incoming USD transfers
	In float64 `json:"in,required"`
	// Total outgoing USD transfers
	Out float64 `json:"out,required"`
	// Total USD transfers
	Total float64                                                                 `json:"total"`
	JSON  solanaMessageScanResponseResultSimulationAccountSummaryTotalUsdDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryTotalUsdDiffJSON contains
// the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryTotalUsdDiff]
type solanaMessageScanResponseResultSimulationAccountSummaryTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiff struct {
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAsset],
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAsset],
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAsset],
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAsset].
	Asset interface{} `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffIn],
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffIn],
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffIn],
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffIn].
	In interface{} `json:"in"`
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffOut],
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffOut],
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffOut],
	// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffOut].
	Out   interface{}                                                                  `json:"out"`
	JSON  solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffJSON `json:"-"`
	union SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffUnion
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiff]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiff],
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiff],
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiff],
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiff].
func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiff) AsUnion() SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffUnion {
	return r.union
}

// Union satisfied by
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiff],
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiff],
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiff]
// or
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiff].
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffUnion interface {
	implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiff{}),
		},
	)
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiff struct {
	Asset SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiff]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiff) implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiff() {
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAsset struct {
	Decimals int64 `json:"decimals,required"`
	// Type of the asset (`"NativeToken"`)
	Type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetType `json:"type,required"`
	// Logo of SOL
	Logo string                                                                                                 `json:"logo,nullable"`
	JSON solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAsset]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	Logo        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`"NativeToken"`)
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetTypeSol SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetType = "SOL"
	SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetTypeEth SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetType = "ETH"
)

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetTypeSol, SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetTypeEth:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                             `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffIn]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                              `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffOut]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiff struct {
	Asset SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiff]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiff) implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiff() {
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's decimals
	Decimals int64 `json:"decimals,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	Logo string                                                                                                      `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAsset]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAssetTypeToken SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAssetType = "TOKEN"
)

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAssetTypeToken:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                  `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffIn]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                   `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffOut]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiff struct {
	Asset SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiff]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiff) implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiff() {
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	// URL of the asset's logo
	Logo string                                                                                                         `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAsset]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAssetTypeNFT SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAssetType = "NFT"
)

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAssetTypeNFT:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                     `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffIn]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                      `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffOut]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiff struct {
	Asset SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiff]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiff) implementsSolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiff() {
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	// URL of the asset's logo
	Logo string                                                                                               `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAsset]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAssetTypeCnft SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAssetType = "CNFT"
)

func (r SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAssetTypeCnft:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                           `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffIn]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                            `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffOut]
type solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiff struct {
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAsset],
	// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAsset],
	// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAsset],
	// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAsset].
	Asset interface{} `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required,nullable"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,required,nullable"`
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffIn],
	// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffIn],
	// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffIn],
	// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffIn].
	In interface{} `json:"in"`
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffOut],
	// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffOut],
	// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffOut],
	// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffOut].
	Out   interface{}                                                      `json:"out"`
	JSON  solanaMessageScanResponseResultSimulationAssetsOwnershipDiffJSON `json:"-"`
	union SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffUnion
}

// solanaMessageScanResponseResultSimulationAssetsOwnershipDiffJSON contains the
// JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiff]
type solanaMessageScanResponseResultSimulationAssetsOwnershipDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	PostOwner   apijson.Field
	PreOwner    apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r solanaMessageScanResponseResultSimulationAssetsOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseResultSimulationAssetsOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	*r = SolanaMessageScanResponseResultSimulationAssetsOwnershipDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiff],
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiff],
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiff],
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiff].
func (r SolanaMessageScanResponseResultSimulationAssetsOwnershipDiff) AsUnion() SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffUnion {
	return r.union
}

// Union satisfied by
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiff],
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiff],
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiff]
// or
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiff].
type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffUnion interface {
	implementsSolanaMessageScanResponseResultSimulationAssetsOwnershipDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiff{}),
		},
	)
}

type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiff struct {
	Asset SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required,nullable"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,required,nullable"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiff]
type solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	PostOwner   apijson.Field
	PreOwner    apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiff) implementsSolanaMessageScanResponseResultSimulationAssetsOwnershipDiff() {
}

type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAsset struct {
	Decimals int64 `json:"decimals,required"`
	// Type of the asset (`"NativeToken"`)
	Type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetType `json:"type,required"`
	// Logo of SOL
	Logo string                                                                                            `json:"logo,nullable"`
	JSON solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAsset]
type solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	Logo        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`"NativeToken"`)
type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetTypeSol SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetType = "SOL"
	SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetTypeEth SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetType = "ETH"
)

func (r SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetTypeSol, SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetTypeEth:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                        `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffIn]
type solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                         `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffOut]
type solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNativeSolOwnershipDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiff struct {
	Asset SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required,nullable"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,required,nullable"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiff]
type solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	PostOwner   apijson.Field
	PreOwner    apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiff) implementsSolanaMessageScanResponseResultSimulationAssetsOwnershipDiff() {
}

type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAsset struct {
	Decimals int64 `json:"decimals,required"`
	// Type of the asset (`"StakedNative"`)
	Type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetType `json:"type,required"`
	// Logo of SOL
	Logo string                                                                                            `json:"logo,nullable"`
	JSON solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAsset]
type solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	Logo        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`"StakedNative"`)
type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetTypeStakedSol SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetType = "STAKED_SOL"
	SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetTypeStakedEth SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetType = "STAKED_ETH"
)

func (r SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetTypeStakedSol, SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetTypeStakedEth:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                        `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffIn]
type solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                         `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffOut]
type solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaStakedSolOwnershipDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiff struct {
	Asset SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required,nullable"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,required,nullable"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiff]
type solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	PostOwner   apijson.Field
	PreOwner    apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiff) implementsSolanaMessageScanResponseResultSimulationAssetsOwnershipDiff() {
}

type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's decimals
	Decimals int64 `json:"decimals,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	Logo string                                                                                              `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAsset]
type solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAssetTypeToken SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAssetType = "TOKEN"
)

func (r SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAssetTypeToken:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                          `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffIn]
type solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                           `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffOut]
type solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiff struct {
	Asset SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required,nullable"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,required,nullable"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiff]
type solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	PostOwner   apijson.Field
	PreOwner    apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiff) implementsSolanaMessageScanResponseResultSimulationAssetsOwnershipDiff() {
}

type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	// URL of the asset's logo
	Logo string                                                                                                 `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAsset]
type solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAssetTypeNFT SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAssetType = "NFT"
)

func (r SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAssetTypeNFT:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                             `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffIn]
type solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                              `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffOut]
type solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAccountsDetail struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	WasWrittenTo   bool   `json:"was_written_to,required"`
	// Description of the account
	Description string `json:"description,nullable"`
	// Logo of the mint
	Logo string `json:"logo,nullable"`
	// Encoded public key of the mint
	MintAddress string `json:"mint_address"`
	// Name of the mint
	Name string `json:"name"`
	// The address of the owning program
	Owner string `json:"owner"`
	// Encoded public key of the owner
	OwnerAddress string `json:"owner_address"`
	// Symbol of the mint
	Symbol string                                                       `json:"symbol"`
	Type   SolanaMessageScanResponseResultSimulationAccountsDetailsType `json:"type"`
	// URI of the mint
	Uri   string                                                      `json:"uri"`
	JSON  solanaMessageScanResponseResultSimulationAccountsDetailJSON `json:"-"`
	union SolanaMessageScanResponseResultSimulationAccountsDetailsUnion
}

// solanaMessageScanResponseResultSimulationAccountsDetailJSON contains the JSON
// metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountsDetail]
type solanaMessageScanResponseResultSimulationAccountsDetailJSON struct {
	AccountAddress apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Logo           apijson.Field
	MintAddress    apijson.Field
	Name           apijson.Field
	Owner          apijson.Field
	OwnerAddress   apijson.Field
	Symbol         apijson.Field
	Type           apijson.Field
	Uri            apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r solanaMessageScanResponseResultSimulationAccountsDetailJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseResultSimulationAccountsDetail) UnmarshalJSON(data []byte) (err error) {
	*r = SolanaMessageScanResponseResultSimulationAccountsDetail{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SolanaMessageScanResponseResultSimulationAccountsDetailsUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaPdaAccountSchema],
// [SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaSystemAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaProgramAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaTokenAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaFungibleMintAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaNonFungibleMintAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaCnftMintAccountDetailsSchema].
func (r SolanaMessageScanResponseResultSimulationAccountsDetail) AsUnion() SolanaMessageScanResponseResultSimulationAccountsDetailsUnion {
	return r.union
}

// Union satisfied by
// [SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaPdaAccountSchema],
// [SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaSystemAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaProgramAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaTokenAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaFungibleMintAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaNonFungibleMintAccountDetailsSchema]
// or
// [SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaCnftMintAccountDetailsSchema].
type SolanaMessageScanResponseResultSimulationAccountsDetailsUnion interface {
	implementsSolanaMessageScanResponseResultSimulationAccountsDetail()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseResultSimulationAccountsDetailsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaPdaAccountSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaSystemAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaProgramAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaTokenAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaFungibleMintAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaNonFungibleMintAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaCnftMintAccountDetailsSchema{}),
		},
	)
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaPdaAccountSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// The address of the owning program
	Owner        string `json:"owner,required"`
	WasWrittenTo bool   `json:"was_written_to,required"`
	// Description of the account
	Description string                                                                             `json:"description,nullable"`
	Type        SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaPdaAccountSchemaType `json:"type"`
	JSON        solanaMessageScanResponseResultSimulationAccountsDetailsSolanaPdaAccountSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountsDetailsSolanaPdaAccountSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaPdaAccountSchema]
type solanaMessageScanResponseResultSimulationAccountsDetailsSolanaPdaAccountSchemaJSON struct {
	AccountAddress apijson.Field
	Owner          apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaPdaAccountSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountsDetailsSolanaPdaAccountSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaPdaAccountSchema) implementsSolanaMessageScanResponseResultSimulationAccountsDetail() {
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaPdaAccountSchemaType string

const (
	SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaPdaAccountSchemaTypePda SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaPdaAccountSchemaType = "PDA"
)

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaPdaAccountSchemaType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaPdaAccountSchemaTypePda:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaSystemAccountDetailsSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	WasWrittenTo   bool   `json:"was_written_to,required"`
	// Description of the account
	Description string                                                                                       `json:"description,nullable"`
	Type        SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaSystemAccountDetailsSchemaType `json:"type"`
	JSON        solanaMessageScanResponseResultSimulationAccountsDetailsSolanaSystemAccountDetailsSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountsDetailsSolanaSystemAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaSystemAccountDetailsSchema]
type solanaMessageScanResponseResultSimulationAccountsDetailsSolanaSystemAccountDetailsSchemaJSON struct {
	AccountAddress apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaSystemAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountsDetailsSolanaSystemAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaSystemAccountDetailsSchema) implementsSolanaMessageScanResponseResultSimulationAccountsDetail() {
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaSystemAccountDetailsSchemaType string

const (
	SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaSystemAccountDetailsSchemaTypeSystemAccount SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaSystemAccountDetailsSchemaType = "SYSTEM_ACCOUNT"
)

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaSystemAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaSystemAccountDetailsSchemaTypeSystemAccount:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaProgramAccountDetailsSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	WasWrittenTo   bool   `json:"was_written_to,required"`
	// Description of the account
	Description string                                                                                        `json:"description,nullable"`
	Type        SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaProgramAccountDetailsSchemaType `json:"type"`
	JSON        solanaMessageScanResponseResultSimulationAccountsDetailsSolanaProgramAccountDetailsSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountsDetailsSolanaProgramAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaProgramAccountDetailsSchema]
type solanaMessageScanResponseResultSimulationAccountsDetailsSolanaProgramAccountDetailsSchemaJSON struct {
	AccountAddress apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaProgramAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountsDetailsSolanaProgramAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaProgramAccountDetailsSchema) implementsSolanaMessageScanResponseResultSimulationAccountsDetail() {
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaProgramAccountDetailsSchemaType string

const (
	SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaProgramAccountDetailsSchemaTypeProgram       SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaProgramAccountDetailsSchemaType = "PROGRAM"
	SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaProgramAccountDetailsSchemaTypeNativeProgram SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaProgramAccountDetailsSchemaType = "NATIVE_PROGRAM"
)

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaProgramAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaProgramAccountDetailsSchemaTypeProgram, SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaProgramAccountDetailsSchemaTypeNativeProgram:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaTokenAccountDetailsSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// Encoded public key of the mint
	MintAddress string `json:"mint_address,required"`
	// Encoded public key of the owner
	OwnerAddress string `json:"owner_address,required"`
	WasWrittenTo bool   `json:"was_written_to,required"`
	// Description of the account
	Description string                                                                                      `json:"description,nullable"`
	Type        SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaTokenAccountDetailsSchemaType `json:"type"`
	JSON        solanaMessageScanResponseResultSimulationAccountsDetailsSolanaTokenAccountDetailsSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountsDetailsSolanaTokenAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaTokenAccountDetailsSchema]
type solanaMessageScanResponseResultSimulationAccountsDetailsSolanaTokenAccountDetailsSchemaJSON struct {
	AccountAddress apijson.Field
	MintAddress    apijson.Field
	OwnerAddress   apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaTokenAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountsDetailsSolanaTokenAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaTokenAccountDetailsSchema) implementsSolanaMessageScanResponseResultSimulationAccountsDetail() {
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaTokenAccountDetailsSchemaType string

const (
	SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaTokenAccountDetailsSchemaTypeTokenAccount SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaTokenAccountDetailsSchemaType = "TOKEN_ACCOUNT"
)

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaTokenAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaTokenAccountDetailsSchemaTypeTokenAccount:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaFungibleMintAccountDetailsSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// Logo of the mint
	Logo string `json:"logo,required,nullable"`
	// Name of the mint
	Name string `json:"name,required"`
	// Symbol of the mint
	Symbol       string `json:"symbol,required"`
	WasWrittenTo bool   `json:"was_written_to,required"`
	// Description of the account
	Description string                                                                                             `json:"description,nullable"`
	Type        SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaFungibleMintAccountDetailsSchemaType `json:"type"`
	JSON        solanaMessageScanResponseResultSimulationAccountsDetailsSolanaFungibleMintAccountDetailsSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountsDetailsSolanaFungibleMintAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaFungibleMintAccountDetailsSchema]
type solanaMessageScanResponseResultSimulationAccountsDetailsSolanaFungibleMintAccountDetailsSchemaJSON struct {
	AccountAddress apijson.Field
	Logo           apijson.Field
	Name           apijson.Field
	Symbol         apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaFungibleMintAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountsDetailsSolanaFungibleMintAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaFungibleMintAccountDetailsSchema) implementsSolanaMessageScanResponseResultSimulationAccountsDetail() {
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaFungibleMintAccountDetailsSchemaType string

const (
	SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaFungibleMintAccountDetailsSchemaTypeFungibleMintAccount SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaFungibleMintAccountDetailsSchemaType = "FUNGIBLE_MINT_ACCOUNT"
)

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaFungibleMintAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaFungibleMintAccountDetailsSchemaTypeFungibleMintAccount:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaNonFungibleMintAccountDetailsSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// Logo of the mint
	Logo string `json:"logo,required,nullable"`
	// Name of the mint
	Name string `json:"name,required"`
	// Symbol of the mint
	Symbol string `json:"symbol,required"`
	// URI of the mint
	Uri          string `json:"uri,required"`
	WasWrittenTo bool   `json:"was_written_to,required"`
	// Description of the account
	Description string                                                                                                `json:"description,nullable"`
	Type        SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaNonFungibleMintAccountDetailsSchemaType `json:"type"`
	JSON        solanaMessageScanResponseResultSimulationAccountsDetailsSolanaNonFungibleMintAccountDetailsSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountsDetailsSolanaNonFungibleMintAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaNonFungibleMintAccountDetailsSchema]
type solanaMessageScanResponseResultSimulationAccountsDetailsSolanaNonFungibleMintAccountDetailsSchemaJSON struct {
	AccountAddress apijson.Field
	Logo           apijson.Field
	Name           apijson.Field
	Symbol         apijson.Field
	Uri            apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaNonFungibleMintAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountsDetailsSolanaNonFungibleMintAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaNonFungibleMintAccountDetailsSchema) implementsSolanaMessageScanResponseResultSimulationAccountsDetail() {
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaNonFungibleMintAccountDetailsSchemaType string

const (
	SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaNonFungibleMintAccountDetailsSchemaTypeNonFungibleMintAccount SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaNonFungibleMintAccountDetailsSchemaType = "NON_FUNGIBLE_MINT_ACCOUNT"
)

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaNonFungibleMintAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaNonFungibleMintAccountDetailsSchemaTypeNonFungibleMintAccount:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaCnftMintAccountDetailsSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// Logo of the mint
	Logo string `json:"logo,required,nullable"`
	// Name of the mint
	Name string `json:"name,required"`
	// Symbol of the mint
	Symbol string `json:"symbol,required"`
	// URI of the mint
	Uri          string `json:"uri,required"`
	WasWrittenTo bool   `json:"was_written_to,required"`
	// Description of the account
	Description string                                                                                         `json:"description,nullable"`
	Type        SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaCnftMintAccountDetailsSchemaType `json:"type"`
	JSON        solanaMessageScanResponseResultSimulationAccountsDetailsSolanaCnftMintAccountDetailsSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAccountsDetailsSolanaCnftMintAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaCnftMintAccountDetailsSchema]
type solanaMessageScanResponseResultSimulationAccountsDetailsSolanaCnftMintAccountDetailsSchemaJSON struct {
	AccountAddress apijson.Field
	Logo           apijson.Field
	Name           apijson.Field
	Symbol         apijson.Field
	Uri            apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaCnftMintAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAccountsDetailsSolanaCnftMintAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaCnftMintAccountDetailsSchema) implementsSolanaMessageScanResponseResultSimulationAccountsDetail() {
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaCnftMintAccountDetailsSchemaType string

const (
	SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaCnftMintAccountDetailsSchemaTypeCnftMintAccount SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaCnftMintAccountDetailsSchemaType = "CNFT_MINT_ACCOUNT"
)

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaCnftMintAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountsDetailsSolanaCnftMintAccountDetailsSchemaTypeCnftMintAccount:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationAccountsDetailsType string

const (
	SolanaMessageScanResponseResultSimulationAccountsDetailsTypePda                    SolanaMessageScanResponseResultSimulationAccountsDetailsType = "PDA"
	SolanaMessageScanResponseResultSimulationAccountsDetailsTypeSystemAccount          SolanaMessageScanResponseResultSimulationAccountsDetailsType = "SYSTEM_ACCOUNT"
	SolanaMessageScanResponseResultSimulationAccountsDetailsTypeProgram                SolanaMessageScanResponseResultSimulationAccountsDetailsType = "PROGRAM"
	SolanaMessageScanResponseResultSimulationAccountsDetailsTypeNativeProgram          SolanaMessageScanResponseResultSimulationAccountsDetailsType = "NATIVE_PROGRAM"
	SolanaMessageScanResponseResultSimulationAccountsDetailsTypeTokenAccount           SolanaMessageScanResponseResultSimulationAccountsDetailsType = "TOKEN_ACCOUNT"
	SolanaMessageScanResponseResultSimulationAccountsDetailsTypeFungibleMintAccount    SolanaMessageScanResponseResultSimulationAccountsDetailsType = "FUNGIBLE_MINT_ACCOUNT"
	SolanaMessageScanResponseResultSimulationAccountsDetailsTypeNonFungibleMintAccount SolanaMessageScanResponseResultSimulationAccountsDetailsType = "NON_FUNGIBLE_MINT_ACCOUNT"
	SolanaMessageScanResponseResultSimulationAccountsDetailsTypeCnftMintAccount        SolanaMessageScanResponseResultSimulationAccountsDetailsType = "CNFT_MINT_ACCOUNT"
)

func (r SolanaMessageScanResponseResultSimulationAccountsDetailsType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAccountsDetailsTypePda, SolanaMessageScanResponseResultSimulationAccountsDetailsTypeSystemAccount, SolanaMessageScanResponseResultSimulationAccountsDetailsTypeProgram, SolanaMessageScanResponseResultSimulationAccountsDetailsTypeNativeProgram, SolanaMessageScanResponseResultSimulationAccountsDetailsTypeTokenAccount, SolanaMessageScanResponseResultSimulationAccountsDetailsTypeFungibleMintAccount, SolanaMessageScanResponseResultSimulationAccountsDetailsTypeNonFungibleMintAccount, SolanaMessageScanResponseResultSimulationAccountsDetailsTypeCnftMintAccount:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationAssetsDiff struct {
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffAsset],
	// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffAsset],
	// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffAsset],
	// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffAsset].
	Asset interface{} `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffIn],
	// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffIn],
	// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffIn],
	// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffIn].
	In interface{} `json:"in"`
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffOut],
	// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffOut],
	// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffOut],
	// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffOut].
	Out   interface{}                                             `json:"out"`
	JSON  solanaMessageScanResponseResultSimulationAssetsDiffJSON `json:"-"`
	union SolanaMessageScanResponseResultSimulationAssetsDiffUnion
}

// solanaMessageScanResponseResultSimulationAssetsDiffJSON contains the JSON
// metadata for the struct [SolanaMessageScanResponseResultSimulationAssetsDiff]
type solanaMessageScanResponseResultSimulationAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r solanaMessageScanResponseResultSimulationAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseResultSimulationAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = SolanaMessageScanResponseResultSimulationAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SolanaMessageScanResponseResultSimulationAssetsDiffUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiff],
// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiff],
// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiff],
// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiff].
func (r SolanaMessageScanResponseResultSimulationAssetsDiff) AsUnion() SolanaMessageScanResponseResultSimulationAssetsDiffUnion {
	return r.union
}

// Union satisfied by
// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiff],
// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiff],
// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiff]
// or [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiff].
type SolanaMessageScanResponseResultSimulationAssetsDiffUnion interface {
	implementsSolanaMessageScanResponseResultSimulationAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseResultSimulationAssetsDiffUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiff{}),
		},
	)
}

type SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiff struct {
	Asset SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiff]
type solanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiff) implementsSolanaMessageScanResponseResultSimulationAssetsDiff() {
}

type SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffAsset struct {
	Decimals int64 `json:"decimals,required"`
	// Type of the asset (`"NativeToken"`)
	Type SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffAssetType `json:"type,required"`
	// Logo of SOL
	Logo string                                                                            `json:"logo,nullable"`
	JSON solanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffAsset]
type solanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	Logo        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`"NativeToken"`)
type SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffAssetTypeSol SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffAssetType = "SOL"
	SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffAssetTypeEth SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffAssetType = "ETH"
)

func (r SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffAssetTypeSol, SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffAssetTypeEth:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                        `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffIn]
type solanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                         `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffOut]
type solanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsDiffSolanaNativeAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiff struct {
	Asset SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiff]
type solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiff) implementsSolanaMessageScanResponseResultSimulationAssetsDiff() {
}

type SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's decimals
	Decimals int64 `json:"decimals,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	Logo string                                                                                 `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffAsset]
type solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffAssetTypeToken SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffAssetType = "TOKEN"
)

func (r SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffAssetTypeToken:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                             `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffIn]
type solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                              `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffOut]
type solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplFungibleAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiff struct {
	Asset SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiff]
type solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiff) implementsSolanaMessageScanResponseResultSimulationAssetsDiff() {
}

type SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	// URL of the asset's logo
	Logo string                                                                                    `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffAsset]
type solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffAssetTypeNFT SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffAssetType = "NFT"
)

func (r SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffAssetTypeNFT:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffIn]
type solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                 `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffOut]
type solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsDiffSolanaSplNonFungibleAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiff struct {
	Asset SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiff]
type solanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiff) implementsSolanaMessageScanResponseResultSimulationAssetsDiff() {
}

type SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	// URL of the asset's logo
	Logo string                                                                          `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffAsset]
type solanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffAssetTypeCnft SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffAssetType = "CNFT"
)

func (r SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffAssetTypeCnft:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                      `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffIn]
type solanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                       `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffOut]
type solanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationAssetsDiffSolanaCnftAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationDelegation struct {
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationAsset],
	// [SolanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationAsset],
	// [SolanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationAsset].
	Asset     interface{} `json:"asset,required"`
	AssetType string      `json:"asset_type,required"`
	Delegate  string      `json:"delegate,required"`
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationDelegation],
	// [SolanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationDelegation],
	// [SolanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationDelegation].
	Delegation interface{}                                             `json:"delegation,required"`
	JSON       solanaMessageScanResponseResultSimulationDelegationJSON `json:"-"`
	union      SolanaMessageScanResponseResultSimulationDelegationsUnion
}

// solanaMessageScanResponseResultSimulationDelegationJSON contains the JSON
// metadata for the struct [SolanaMessageScanResponseResultSimulationDelegation]
type solanaMessageScanResponseResultSimulationDelegationJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Delegate    apijson.Field
	Delegation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r solanaMessageScanResponseResultSimulationDelegationJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseResultSimulationDelegation) UnmarshalJSON(data []byte) (err error) {
	*r = SolanaMessageScanResponseResultSimulationDelegation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SolanaMessageScanResponseResultSimulationDelegationsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SolanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegation],
// [SolanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegation],
// [SolanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegation].
func (r SolanaMessageScanResponseResultSimulationDelegation) AsUnion() SolanaMessageScanResponseResultSimulationDelegationsUnion {
	return r.union
}

// Union satisfied by
// [SolanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegation],
// [SolanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegation]
// or
// [SolanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegation].
type SolanaMessageScanResponseResultSimulationDelegationsUnion interface {
	implementsSolanaMessageScanResponseResultSimulationDelegation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseResultSimulationDelegationsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegation{}),
		},
	)
}

type SolanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegation struct {
	Asset      SolanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationAsset      `json:"asset,required"`
	AssetType  string                                                                             `json:"asset_type,required"`
	Delegate   string                                                                             `json:"delegate,required"`
	Delegation SolanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationDelegation `json:"delegation,required"`
	JSON       solanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationJSON       `json:"-"`
}

// solanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegation]
type solanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Delegate    apijson.Field
	Delegation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegation) implementsSolanaMessageScanResponseResultSimulationDelegation() {
}

type SolanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	// URL of the asset's logo
	Logo string                                                                            `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationAsset]
type solanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationAssetType string

const (
	SolanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationAssetTypeCnft SolanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationAssetType = "CNFT"
)

func (r SolanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationAssetTypeCnft:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationDelegation struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationDelegationJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationDelegation]
type solanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationDelegationJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationDelegationsSolanaCnftDelegationDelegationJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegation struct {
	Asset      SolanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationAsset      `json:"asset,required"`
	AssetType  string                                                                                         `json:"asset_type,required"`
	Delegate   string                                                                                         `json:"delegate,required"`
	Delegation SolanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationDelegation `json:"delegation,required"`
	JSON       solanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationJSON       `json:"-"`
}

// solanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegation]
type solanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Delegate    apijson.Field
	Delegation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegation) implementsSolanaMessageScanResponseResultSimulationDelegation() {
}

type SolanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's decimals
	Decimals int64 `json:"decimals,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	Logo string                                                                                        `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationAsset]
type solanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationAssetType string

const (
	SolanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationAssetTypeToken SolanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationAssetType = "TOKEN"
)

func (r SolanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationAssetTypeToken:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationDelegation struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                            `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationDelegationJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationDelegation]
type solanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationDelegationJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationDelegationsSolanaFungibleSplTokenDelegationDelegationJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegation struct {
	Asset      SolanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationAsset      `json:"asset,required"`
	AssetType  string                                                                                            `json:"asset_type,required"`
	Delegate   string                                                                                            `json:"delegate,required"`
	Delegation SolanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationDelegation `json:"delegation,required"`
	JSON       solanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationJSON       `json:"-"`
}

// solanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegation]
type solanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Delegate    apijson.Field
	Delegation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegation) implementsSolanaMessageScanResponseResultSimulationDelegation() {
}

type SolanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	// URL of the asset's logo
	Logo string                                                                                           `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationAsset]
type solanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationAssetType string

const (
	SolanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationAssetTypeNFT SolanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationAssetType = "NFT"
)

func (r SolanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationAssetTypeNFT:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationDelegation struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                               `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationDelegationJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationDelegation]
type solanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationDelegationJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationDelegationsSolanaNonFungibleSplTokenDelegationDelegationJSON) RawJSON() string {
	return r.raw
}

// Transaction Validation Result
type SolanaMessageScanResponseResultValidation struct {
	// A list of features explaining what is happening in the transaction in different
	// levels of severity
	ExtendedFeatures []SolanaMessageScanResponseResultValidationExtendedFeature `json:"extended_features,required"`
	// A list of features about this transaction explaining the validation
	Features []string `json:"features,required"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason string `json:"reason,required"`
	// Verdict of the validation
	ResultType SolanaMessageScanResponseResultValidationResultType `json:"result_type,required"`
	JSON       solanaMessageScanResponseResultValidationJSON       `json:"-"`
}

// solanaMessageScanResponseResultValidationJSON contains the JSON metadata for the
// struct [SolanaMessageScanResponseResultValidation]
type solanaMessageScanResponseResultValidationJSON struct {
	ExtendedFeatures apijson.Field
	Features         apijson.Field
	Reason           apijson.Field
	ResultType       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultValidationJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultValidationExtendedFeature struct {
	// Address the feature refers to
	Address string `json:"address,required,nullable"`
	// Textual description
	Description string `json:"description,required"`
	FeatureID   string `json:"feature_id,required"`
	// Feature Classification
	Type SolanaMessageScanResponseResultValidationExtendedFeaturesType `json:"type,required"`
	JSON solanaMessageScanResponseResultValidationExtendedFeatureJSON  `json:"-"`
}

// solanaMessageScanResponseResultValidationExtendedFeatureJSON contains the JSON
// metadata for the struct
// [SolanaMessageScanResponseResultValidationExtendedFeature]
type solanaMessageScanResponseResultValidationExtendedFeatureJSON struct {
	Address     apijson.Field
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultValidationExtendedFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultValidationExtendedFeatureJSON) RawJSON() string {
	return r.raw
}

// Feature Classification
type SolanaMessageScanResponseResultValidationExtendedFeaturesType string

const (
	SolanaMessageScanResponseResultValidationExtendedFeaturesTypeBenign    SolanaMessageScanResponseResultValidationExtendedFeaturesType = "Benign"
	SolanaMessageScanResponseResultValidationExtendedFeaturesTypeWarning   SolanaMessageScanResponseResultValidationExtendedFeaturesType = "Warning"
	SolanaMessageScanResponseResultValidationExtendedFeaturesTypeMalicious SolanaMessageScanResponseResultValidationExtendedFeaturesType = "Malicious"
	SolanaMessageScanResponseResultValidationExtendedFeaturesTypeInfo      SolanaMessageScanResponseResultValidationExtendedFeaturesType = "Info"
)

func (r SolanaMessageScanResponseResultValidationExtendedFeaturesType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultValidationExtendedFeaturesTypeBenign, SolanaMessageScanResponseResultValidationExtendedFeaturesTypeWarning, SolanaMessageScanResponseResultValidationExtendedFeaturesTypeMalicious, SolanaMessageScanResponseResultValidationExtendedFeaturesTypeInfo:
		return true
	}
	return false
}

// Verdict of the validation
type SolanaMessageScanResponseResultValidationResultType string

const (
	SolanaMessageScanResponseResultValidationResultTypeBenign    SolanaMessageScanResponseResultValidationResultType = "Benign"
	SolanaMessageScanResponseResultValidationResultTypeWarning   SolanaMessageScanResponseResultValidationResultType = "Warning"
	SolanaMessageScanResponseResultValidationResultTypeMalicious SolanaMessageScanResponseResultValidationResultType = "Malicious"
)

func (r SolanaMessageScanResponseResultValidationResultType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultValidationResultTypeBenign, SolanaMessageScanResponseResultValidationResultTypeWarning, SolanaMessageScanResponseResultValidationResultTypeMalicious:
		return true
	}
	return false
}

type SolanaMessageScanParams struct {
	AccountAddress param.Field[string]                          `json:"account_address,required"`
	Metadata       param.Field[SolanaMessageScanParamsMetadata] `json:"metadata,required"`
	// Transactions to scan
	Transactions param.Field[[]string]                        `json:"transactions,required"`
	Chain        param.Field[string]                          `json:"chain"`
	Encoding     param.Field[SolanaMessageScanParamsEncoding] `json:"encoding"`
	// The RPC method used by the dApp to propose the transaction
	Method param.Field[string] `json:"method"`
	// List of options to include in the response
	//
	// - `Options.validation`: Include Options.validation output in the response
	//
	// - `Options.simulation`: Include Options.simulation output in the response
	Options param.Field[[]SolanaMessageScanParamsOption] `json:"options"`
}

func (r SolanaMessageScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SolanaMessageScanParamsMetadata struct {
	// URL of the dApp that originated the transaction
	URL param.Field[string] `json:"url"`
}

func (r SolanaMessageScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SolanaMessageScanParamsEncoding string

const (
	SolanaMessageScanParamsEncodingBase58 SolanaMessageScanParamsEncoding = "base58"
	SolanaMessageScanParamsEncodingBase64 SolanaMessageScanParamsEncoding = "base64"
)

func (r SolanaMessageScanParamsEncoding) IsKnown() bool {
	switch r {
	case SolanaMessageScanParamsEncodingBase58, SolanaMessageScanParamsEncodingBase64:
		return true
	}
	return false
}

type SolanaMessageScanParamsOption string

const (
	SolanaMessageScanParamsOptionValidation SolanaMessageScanParamsOption = "validation"
	SolanaMessageScanParamsOptionSimulation SolanaMessageScanParamsOption = "simulation"
)

func (r SolanaMessageScanParamsOption) IsKnown() bool {
	switch r {
	case SolanaMessageScanParamsOptionValidation, SolanaMessageScanParamsOptionSimulation:
		return true
	}
	return false
}
