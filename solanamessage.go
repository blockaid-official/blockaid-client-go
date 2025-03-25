// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"net/http"
	"reflect"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
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

// Scan a message
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
	// Simulation result; Only present if simulation option is included in the request
	Simulation SolanaMessageScanResponseResultSimulation `json:"simulation,nullable"`
	// Validation result; Only present if validation option is included in the request
	Validation SolanaMessageScanResponseResultValidation `json:"validation,nullable"`
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

// Simulation result; Only present if simulation option is included in the request
type SolanaMessageScanResponseResultSimulation struct {
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummary].
	AccountSummary interface{} `json:"account_summary"`
	// This field can have the runtime type of
	// [[]SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetail].
	AccountsDetails interface{} `json:"accounts_details"`
	// This field can have the runtime type of
	// [map[string][]SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiff].
	AssetsDiff interface{} `json:"assets_diff"`
	// This field can have the runtime type of
	// [map[string][]SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiff].
	AssetsOwnershipDiff interface{} `json:"assets_ownership_diff"`
	// This field can have the runtime type of
	// [map[string][]SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegation].
	Delegations interface{} `json:"delegations"`
	// Error message
	Error  string                                          `json:"error"`
	Status SolanaMessageScanResponseResultSimulationStatus `json:"status"`
	JSON   solanaMessageScanResponseResultSimulationJSON   `json:"-"`
	union  SolanaMessageScanResponseResultSimulationUnion
}

// solanaMessageScanResponseResultSimulationJSON contains the JSON metadata for the
// struct [SolanaMessageScanResponseResultSimulation]
type solanaMessageScanResponseResultSimulationJSON struct {
	AccountSummary      apijson.Field
	AccountsDetails     apijson.Field
	AssetsDiff          apijson.Field
	AssetsOwnershipDiff apijson.Field
	Delegations         apijson.Field
	Error               apijson.Field
	Status              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r solanaMessageScanResponseResultSimulationJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseResultSimulation) UnmarshalJSON(data []byte) (err error) {
	*r = SolanaMessageScanResponseResultSimulation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SolanaMessageScanResponseResultSimulationUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchema],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationErrorSchema].
func (r SolanaMessageScanResponseResultSimulation) AsUnion() SolanaMessageScanResponseResultSimulationUnion {
	return r.union
}

// Simulation result; Only present if simulation option is included in the request
//
// Union satisfied by
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchema] or
// [SolanaMessageScanResponseResultSimulationSolanaSimulationErrorSchema].
type SolanaMessageScanResponseResultSimulationUnion interface {
	implementsSolanaMessageScanResponseResultSimulation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseResultSimulationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationErrorSchema{}),
		},
	)
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchema struct {
	// Summary of the actions and asset transfers that were made by the requested
	// account address
	AccountSummary SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummary `json:"account_summary,required"`
	// Ownership diffs of the account addresses
	AssetsOwnershipDiff map[string][]SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiff `json:"assets_ownership_diff,required"`
	// Details of addresses involved in the transaction
	AccountsDetails []SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetail `json:"accounts_details"`
	// Mapping between the address of an account to the assets diff during the
	// transaction
	AssetsDiff map[string][]SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiff `json:"assets_diff"`
	// Mapping between the address of an account to the exposure of the assets during
	// the transaction
	Delegations map[string][]SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegation `json:"delegations"`
	JSON        solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaJSON                    `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchema]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaJSON struct {
	AccountSummary      apijson.Field
	AssetsOwnershipDiff apijson.Field
	AccountsDetails     apijson.Field
	AssetsDiff          apijson.Field
	Delegations         apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchema) implementsSolanaMessageScanResponseResultSimulation() {
}

// Summary of the actions and asset transfers that were made by the requested
// account address
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummary struct {
	// Exposures made by the requested account address
	AccountDelegations []SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegation `json:"account_delegations,required"`
	// Ownership diffs of the requested account address
	AccountOwnershipsDiff []SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiff `json:"account_ownerships_diff,required"`
	// Total USD diff for the requested account address
	TotalUsdDiff SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryTotalUsdDiff `json:"total_usd_diff,required"`
	// Assets diffs of the requested account address
	AccountAssetsDiff []SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiff `json:"account_assets_diff"`
	// Total USD exposure for each of the spender addresses during the transaction
	TotalUsdExposure map[string]float64                                                                      `json:"total_usd_exposure"`
	JSON             solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummary]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryJSON struct {
	AccountDelegations    apijson.Field
	AccountOwnershipsDiff apijson.Field
	TotalUsdDiff          apijson.Field
	AccountAssetsDiff     apijson.Field
	TotalUsdExposure      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegation struct {
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationAsset],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAsset],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAsset].
	Asset     interface{} `json:"asset,required"`
	AssetType string      `json:"asset_type,required"`
	Delegate  string      `json:"delegate,required"`
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationDelegation],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationDelegation],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationDelegation].
	Delegation interface{}                                                                                              `json:"delegation,required"`
	JSON       solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationJSON `json:"-"`
	union      SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsUnion
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegation]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Delegate    apijson.Field
	Delegation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegation) UnmarshalJSON(data []byte) (err error) {
	*r = SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegation],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegation],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegation].
func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegation) AsUnion() SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsUnion {
	return r.union
}

// Union satisfied by
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegation],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegation]
// or
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegation].
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsUnion interface {
	implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegation{}),
		},
	)
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegation struct {
	Asset      SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationAsset      `json:"asset,required"`
	AssetType  string                                                                                                                              `json:"asset_type,required"`
	Delegate   string                                                                                                                              `json:"delegate,required"`
	Delegation SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationDelegation `json:"delegation,required"`
	JSON       solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationJSON       `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegation]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Delegate    apijson.Field
	Delegation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegation) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegation() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	// URL of the asset's logo
	Logo string                                                                                                                             `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationAsset]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationAssetType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationAssetTypeCnft SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationAssetType = "CNFT"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationAssetTypeCnft:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationDelegation struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                                 `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationDelegationJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationDelegation]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationDelegationJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaCnftDelegationDelegationJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegation struct {
	Asset      SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAsset      `json:"asset,required"`
	AssetType  string                                                                                                                                          `json:"asset_type,required"`
	Delegate   string                                                                                                                                          `json:"delegate,required"`
	Delegation SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationDelegation `json:"delegation,required"`
	JSON       solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationJSON       `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegation]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Delegate    apijson.Field
	Delegation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegation) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegation() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's decimals
	Decimals int64 `json:"decimals,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	Logo string                                                                                                                                         `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAsset]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAssetType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAssetTypeToken SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAssetType = "TOKEN"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationAssetTypeToken:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationDelegation struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                                             `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationDelegationJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationDelegation]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationDelegationJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaFungibleSplTokenDelegationDelegationJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegation struct {
	Asset      SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAsset      `json:"asset,required"`
	AssetType  string                                                                                                                                             `json:"asset_type,required"`
	Delegate   string                                                                                                                                             `json:"delegate,required"`
	Delegation SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationDelegation `json:"delegation,required"`
	JSON       solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationJSON       `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegation]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Delegate    apijson.Field
	Delegation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegation) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegation() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	// URL of the asset's logo
	Logo string                                                                                                                                            `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAsset]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAssetType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAssetTypeNFT SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAssetType = "NFT"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationAssetTypeNFT:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationDelegation struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                                                `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationDelegationJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationDelegation]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationDelegationJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountDelegationsSolanaNonFungibleSplTokenDelegationDelegationJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiff struct {
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAsset],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAsset],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAsset],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAsset].
	Asset interface{} `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required,nullable"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,required,nullable"`
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffIn],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffIn],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffIn],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffIn].
	In interface{} `json:"in"`
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffOut],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffOut],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffOut],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffOut].
	Out   interface{}                                                                                                  `json:"out"`
	JSON  solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffJSON `json:"-"`
	union SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffUnion
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiff]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	PostOwner   apijson.Field
	PreOwner    apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiff],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiff],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiff],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiff].
func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiff) AsUnion() SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffUnion {
	return r.union
}

// Union satisfied by
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiff],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiff],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiff]
// or
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiff].
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffUnion interface {
	implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiff{}),
		},
	)
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiff struct {
	Asset SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required,nullable"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,required,nullable"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiff]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	PostOwner   apijson.Field
	PreOwner    apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiff) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiff() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAsset struct {
	Decimals int64 `json:"decimals,required"`
	// Type of the asset (`"NativeToken"`)
	Type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetType `json:"type,required"`
	// Logo of SOL
	Logo string                                                                                                                                        `json:"logo,nullable"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAsset]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	Logo        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`"NativeToken"`)
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetTypeSol SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetType = "SOL"
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetTypeEth SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetType = "ETH"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetTypeSol, SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffAssetTypeEth:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                                    `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffIn]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                                     `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffOut]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNativeSolOwnershipDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiff struct {
	Asset SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required,nullable"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,required,nullable"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiff]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	PostOwner   apijson.Field
	PreOwner    apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiff) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiff() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAsset struct {
	Decimals int64 `json:"decimals,required"`
	// Type of the asset (`"StakedNative"`)
	Type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetType `json:"type,required"`
	// Logo of SOL
	Logo string                                                                                                                                        `json:"logo,nullable"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAsset]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	Logo        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`"StakedNative"`)
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetTypeStakedSol SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetType = "STAKED_SOL"
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetTypeStakedEth SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetType = "STAKED_ETH"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetTypeStakedSol, SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffAssetTypeStakedEth:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                                    `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffIn]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                                     `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffOut]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaStakedSolOwnershipDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiff struct {
	Asset SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required,nullable"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,required,nullable"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiff]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	PostOwner   apijson.Field
	PreOwner    apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiff) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiff() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's decimals
	Decimals int64 `json:"decimals,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	Logo string                                                                                                                                          `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAsset]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAssetTypeToken SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAssetType = "TOKEN"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffAssetTypeToken:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                                      `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffIn]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                                       `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffOut]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaFungibleSolOwnershipDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiff struct {
	Asset SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required,nullable"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,required,nullable"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiff]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	PostOwner   apijson.Field
	PreOwner    apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiff) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiff() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	// URL of the asset's logo
	Logo string                                                                                                                                             `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAsset]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAssetTypeNFT SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAssetType = "NFT"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffAssetTypeNFT:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                                         `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffIn]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                                          `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffOut]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountOwnershipsDiffSolanaNonFungibleSolOwnershipDiffOutJSON) RawJSON() string {
	return r.raw
}

// Total USD diff for the requested account address
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryTotalUsdDiff struct {
	// Total incoming USD transfers
	In float64 `json:"in,required"`
	// Total outgoing USD transfers
	Out float64 `json:"out,required"`
	// Total USD transfers
	Total float64                                                                                             `json:"total"`
	JSON  solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryTotalUsdDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryTotalUsdDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryTotalUsdDiff]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryTotalUsdDiffJSON struct {
	In          apijson.Field
	Out         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryTotalUsdDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryTotalUsdDiffJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiff struct {
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAsset],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAsset],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAsset],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAsset].
	Asset interface{} `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffIn],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffIn],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffIn],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffIn].
	In interface{} `json:"in"`
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffOut],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffOut],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffOut],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffOut].
	Out   interface{}                                                                                              `json:"out"`
	JSON  solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffJSON `json:"-"`
	union SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffUnion
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiff]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiff],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiff],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiff],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiff].
func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiff) AsUnion() SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffUnion {
	return r.union
}

// Union satisfied by
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiff],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiff],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiff]
// or
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiff].
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffUnion interface {
	implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiff{}),
		},
	)
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiff struct {
	Asset SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiff]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiff) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiff() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAsset struct {
	Decimals int64 `json:"decimals,required"`
	// Type of the asset (`"NativeToken"`)
	Type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetType `json:"type,required"`
	// Logo of SOL
	Logo string                                                                                                                             `json:"logo,nullable"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAsset]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	Logo        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`"NativeToken"`)
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetTypeSol SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetType = "SOL"
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetTypeEth SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetType = "ETH"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetTypeSol, SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffAssetTypeEth:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                         `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffIn]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                          `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffOut]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaNativeAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiff struct {
	Asset SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiff]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiff) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiff() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's decimals
	Decimals int64 `json:"decimals,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	Logo string                                                                                                                                  `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAsset]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAssetTypeToken SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAssetType = "TOKEN"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffAssetTypeToken:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                              `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffIn]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                               `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffOut]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplFungibleAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiff struct {
	Asset SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiff]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiff) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiff() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	// URL of the asset's logo
	Logo string                                                                                                                                     `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAsset]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAssetTypeNFT SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAssetType = "NFT"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffAssetTypeNFT:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                                 `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffIn]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                                  `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffOut]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaSplNonFungibleAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiff struct {
	Asset SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiff]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiff) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiff() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	// URL of the asset's logo
	Logo string                                                                                                                           `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAsset]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAssetTypeCnft SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAssetType = "CNFT"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffAssetTypeCnft:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                       `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffIn]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                        `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffOut]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountSummaryAccountAssetsDiffSolanaCnftAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiff struct {
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAsset],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAsset],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAsset],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAsset].
	Asset interface{} `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required,nullable"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,required,nullable"`
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffIn],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffIn],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffIn],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffIn].
	In interface{} `json:"in"`
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffOut],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffOut],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffOut],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffOut].
	Out   interface{}                                                                                  `json:"out"`
	JSON  solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffJSON `json:"-"`
	union SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffUnion
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiff]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	PostOwner   apijson.Field
	PreOwner    apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	*r = SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiff],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiff],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiff],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiff].
func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiff) AsUnion() SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffUnion {
	return r.union
}

// Union satisfied by
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiff],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiff],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiff]
// or
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiff].
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffUnion interface {
	implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiff{}),
		},
	)
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiff struct {
	Asset SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required,nullable"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,required,nullable"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiff]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	PostOwner   apijson.Field
	PreOwner    apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiff) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiff() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAsset struct {
	Decimals int64 `json:"decimals,required"`
	// Type of the asset (`"NativeToken"`)
	Type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetType `json:"type,required"`
	// Logo of SOL
	Logo string                                                                                                                        `json:"logo,nullable"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAsset]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	Logo        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`"NativeToken"`)
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetTypeSol SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetType = "SOL"
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetTypeEth SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetType = "ETH"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetTypeSol, SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffAssetTypeEth:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                    `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffIn]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                     `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffOut]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNativeSolOwnershipDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiff struct {
	Asset SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required,nullable"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,required,nullable"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiff]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	PostOwner   apijson.Field
	PreOwner    apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiff) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiff() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAsset struct {
	Decimals int64 `json:"decimals,required"`
	// Type of the asset (`"StakedNative"`)
	Type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetType `json:"type,required"`
	// Logo of SOL
	Logo string                                                                                                                        `json:"logo,nullable"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAsset]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	Logo        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`"StakedNative"`)
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetTypeStakedSol SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetType = "STAKED_SOL"
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetTypeStakedEth SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetType = "STAKED_ETH"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetTypeStakedSol, SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffAssetTypeStakedEth:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                    `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffIn]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                     `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffOut]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaStakedSolOwnershipDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiff struct {
	Asset SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required,nullable"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,required,nullable"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiff]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	PostOwner   apijson.Field
	PreOwner    apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiff) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiff() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's decimals
	Decimals int64 `json:"decimals,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	Logo string                                                                                                                          `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAsset]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAssetTypeToken SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAssetType = "TOKEN"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffAssetTypeToken:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                      `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffIn]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                       `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffOut]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaFungibleSolOwnershipDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiff struct {
	Asset SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// The owner post the transaction
	PostOwner string `json:"post_owner,required,nullable"`
	// The owner prior to the transaction
	PreOwner string `json:"pre_owner,required,nullable"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiff]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	PostOwner   apijson.Field
	PreOwner    apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiff) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiff() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	// URL of the asset's logo
	Logo string                                                                                                                             `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAsset]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAssetTypeNFT SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAssetType = "NFT"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffAssetTypeNFT:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                         `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffIn]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                          `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffOut]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsOwnershipDiffSolanaNonFungibleSolOwnershipDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetail struct {
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
	Symbol string                                                                                   `json:"symbol"`
	Type   SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsType `json:"type"`
	// URI of the mint
	Uri   string                                                                                  `json:"uri"`
	JSON  solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailJSON `json:"-"`
	union SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsUnion
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetail]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailJSON struct {
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

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetail) UnmarshalJSON(data []byte) (err error) {
	*r = SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetail{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaPdaAccountSchema],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaSystemAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaProgramAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaTokenAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaFungibleMintAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaNonFungibleMintAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaCnftMintAccountDetailsSchema].
func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetail) AsUnion() SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsUnion {
	return r.union
}

// Union satisfied by
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaPdaAccountSchema],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaSystemAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaProgramAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaTokenAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaFungibleMintAccountDetailsSchema],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaNonFungibleMintAccountDetailsSchema]
// or
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaCnftMintAccountDetailsSchema].
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsUnion interface {
	implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetail()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaPdaAccountSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaSystemAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaProgramAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaTokenAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaFungibleMintAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaNonFungibleMintAccountDetailsSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaCnftMintAccountDetailsSchema{}),
		},
	)
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaPdaAccountSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// The address of the owning program
	Owner        string `json:"owner,required"`
	WasWrittenTo bool   `json:"was_written_to,required"`
	// Description of the account
	Description string                                                                                                         `json:"description,nullable"`
	Type        SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaPdaAccountSchemaType `json:"type"`
	JSON        solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaPdaAccountSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaPdaAccountSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaPdaAccountSchema]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaPdaAccountSchemaJSON struct {
	AccountAddress apijson.Field
	Owner          apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaPdaAccountSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaPdaAccountSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaPdaAccountSchema) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetail() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaPdaAccountSchemaType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaPdaAccountSchemaTypePda SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaPdaAccountSchemaType = "PDA"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaPdaAccountSchemaType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaPdaAccountSchemaTypePda:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaSystemAccountDetailsSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	WasWrittenTo   bool   `json:"was_written_to,required"`
	// Description of the account
	Description string                                                                                                                   `json:"description,nullable"`
	Type        SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaSystemAccountDetailsSchemaType `json:"type"`
	JSON        solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaSystemAccountDetailsSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaSystemAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaSystemAccountDetailsSchema]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaSystemAccountDetailsSchemaJSON struct {
	AccountAddress apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaSystemAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaSystemAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaSystemAccountDetailsSchema) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetail() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaSystemAccountDetailsSchemaType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaSystemAccountDetailsSchemaTypeSystemAccount SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaSystemAccountDetailsSchemaType = "SYSTEM_ACCOUNT"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaSystemAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaSystemAccountDetailsSchemaTypeSystemAccount:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaProgramAccountDetailsSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	WasWrittenTo   bool   `json:"was_written_to,required"`
	// Description of the account
	Description string                                                                                                                    `json:"description,nullable"`
	Type        SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaProgramAccountDetailsSchemaType `json:"type"`
	JSON        solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaProgramAccountDetailsSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaProgramAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaProgramAccountDetailsSchema]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaProgramAccountDetailsSchemaJSON struct {
	AccountAddress apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaProgramAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaProgramAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaProgramAccountDetailsSchema) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetail() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaProgramAccountDetailsSchemaType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaProgramAccountDetailsSchemaTypeProgram       SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaProgramAccountDetailsSchemaType = "PROGRAM"
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaProgramAccountDetailsSchemaTypeNativeProgram SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaProgramAccountDetailsSchemaType = "NATIVE_PROGRAM"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaProgramAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaProgramAccountDetailsSchemaTypeProgram, SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaProgramAccountDetailsSchemaTypeNativeProgram:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaTokenAccountDetailsSchema struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// Encoded public key of the mint
	MintAddress string `json:"mint_address,required"`
	// Encoded public key of the owner
	OwnerAddress string `json:"owner_address,required"`
	WasWrittenTo bool   `json:"was_written_to,required"`
	// Description of the account
	Description string                                                                                                                  `json:"description,nullable"`
	Type        SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaTokenAccountDetailsSchemaType `json:"type"`
	JSON        solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaTokenAccountDetailsSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaTokenAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaTokenAccountDetailsSchema]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaTokenAccountDetailsSchemaJSON struct {
	AccountAddress apijson.Field
	MintAddress    apijson.Field
	OwnerAddress   apijson.Field
	WasWrittenTo   apijson.Field
	Description    apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaTokenAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaTokenAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaTokenAccountDetailsSchema) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetail() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaTokenAccountDetailsSchemaType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaTokenAccountDetailsSchemaTypeTokenAccount     SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaTokenAccountDetailsSchemaType = "TOKEN_ACCOUNT"
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaTokenAccountDetailsSchemaTypeToken2022Account SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaTokenAccountDetailsSchemaType = "TOKEN_2022_ACCOUNT"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaTokenAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaTokenAccountDetailsSchemaTypeTokenAccount, SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaTokenAccountDetailsSchemaTypeToken2022Account:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaFungibleMintAccountDetailsSchema struct {
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
	Description string                                                                                                                         `json:"description,nullable"`
	Type        SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaFungibleMintAccountDetailsSchemaType `json:"type"`
	JSON        solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaFungibleMintAccountDetailsSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaFungibleMintAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaFungibleMintAccountDetailsSchema]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaFungibleMintAccountDetailsSchemaJSON struct {
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

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaFungibleMintAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaFungibleMintAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaFungibleMintAccountDetailsSchema) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetail() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaFungibleMintAccountDetailsSchemaType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaFungibleMintAccountDetailsSchemaTypeFungibleMintAccount     SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaFungibleMintAccountDetailsSchemaType = "FUNGIBLE_MINT_ACCOUNT"
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaFungibleMintAccountDetailsSchemaTypeFungibleMintAccount2022 SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaFungibleMintAccountDetailsSchemaType = "FUNGIBLE_MINT_ACCOUNT_2022"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaFungibleMintAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaFungibleMintAccountDetailsSchemaTypeFungibleMintAccount, SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaFungibleMintAccountDetailsSchemaTypeFungibleMintAccount2022:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaNonFungibleMintAccountDetailsSchema struct {
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
	Description string                                                                                                                            `json:"description,nullable"`
	Type        SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaNonFungibleMintAccountDetailsSchemaType `json:"type"`
	JSON        solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaNonFungibleMintAccountDetailsSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaNonFungibleMintAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaNonFungibleMintAccountDetailsSchema]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaNonFungibleMintAccountDetailsSchemaJSON struct {
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

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaNonFungibleMintAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaNonFungibleMintAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaNonFungibleMintAccountDetailsSchema) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetail() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaNonFungibleMintAccountDetailsSchemaType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaNonFungibleMintAccountDetailsSchemaTypeNonFungibleMintAccount     SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaNonFungibleMintAccountDetailsSchemaType = "NON_FUNGIBLE_MINT_ACCOUNT"
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaNonFungibleMintAccountDetailsSchemaTypeNonFungibleMintAccount2022 SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaNonFungibleMintAccountDetailsSchemaType = "NON_FUNGIBLE_MINT_ACCOUNT_2022"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaNonFungibleMintAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaNonFungibleMintAccountDetailsSchemaTypeNonFungibleMintAccount, SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaNonFungibleMintAccountDetailsSchemaTypeNonFungibleMintAccount2022:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaCnftMintAccountDetailsSchema struct {
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
	Description string                                                                                                                     `json:"description,nullable"`
	Type        SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaCnftMintAccountDetailsSchemaType `json:"type"`
	JSON        solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaCnftMintAccountDetailsSchemaJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaCnftMintAccountDetailsSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaCnftMintAccountDetailsSchema]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaCnftMintAccountDetailsSchemaJSON struct {
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

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaCnftMintAccountDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaCnftMintAccountDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaCnftMintAccountDetailsSchema) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetail() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaCnftMintAccountDetailsSchemaType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaCnftMintAccountDetailsSchemaTypeCnftMintAccount SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaCnftMintAccountDetailsSchemaType = "CNFT_MINT_ACCOUNT"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaCnftMintAccountDetailsSchemaType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsSolanaCnftMintAccountDetailsSchemaTypeCnftMintAccount:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsTypePda                        SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsType = "PDA"
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsTypeSystemAccount              SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsType = "SYSTEM_ACCOUNT"
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsTypeProgram                    SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsType = "PROGRAM"
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsTypeNativeProgram              SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsType = "NATIVE_PROGRAM"
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsTypeTokenAccount               SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsType = "TOKEN_ACCOUNT"
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsTypeToken2022Account           SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsType = "TOKEN_2022_ACCOUNT"
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsTypeFungibleMintAccount        SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsType = "FUNGIBLE_MINT_ACCOUNT"
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsTypeFungibleMintAccount2022    SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsType = "FUNGIBLE_MINT_ACCOUNT_2022"
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsTypeNonFungibleMintAccount     SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsType = "NON_FUNGIBLE_MINT_ACCOUNT"
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsTypeNonFungibleMintAccount2022 SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsType = "NON_FUNGIBLE_MINT_ACCOUNT_2022"
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsTypeCnftMintAccount            SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsType = "CNFT_MINT_ACCOUNT"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsTypePda, SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsTypeSystemAccount, SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsTypeProgram, SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsTypeNativeProgram, SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsTypeTokenAccount, SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsTypeToken2022Account, SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsTypeFungibleMintAccount, SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsTypeFungibleMintAccount2022, SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsTypeNonFungibleMintAccount, SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsTypeNonFungibleMintAccount2022, SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAccountsDetailsTypeCnftMintAccount:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiff struct {
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffAsset],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffAsset],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffAsset],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffAsset].
	Asset interface{} `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffIn],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffIn],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffIn],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffIn].
	In interface{} `json:"in"`
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffOut],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffOut],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffOut],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffOut].
	Out   interface{}                                                                         `json:"out"`
	JSON  solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffJSON `json:"-"`
	union SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffUnion
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiff]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiff],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiff],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiff],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiff].
func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiff) AsUnion() SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffUnion {
	return r.union
}

// Union satisfied by
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiff],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiff],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiff]
// or
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiff].
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffUnion interface {
	implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiff{}),
		},
	)
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiff struct {
	Asset SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiff]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiff) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiff() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffAsset struct {
	Decimals int64 `json:"decimals,required"`
	// Type of the asset (`"NativeToken"`)
	Type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffAssetType `json:"type,required"`
	// Logo of SOL
	Logo string                                                                                                        `json:"logo,nullable"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffAsset]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffAssetJSON struct {
	Decimals    apijson.Field
	Type        apijson.Field
	Logo        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`"NativeToken"`)
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffAssetTypeSol SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffAssetType = "SOL"
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffAssetTypeEth SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffAssetType = "ETH"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffAssetTypeSol, SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffAssetTypeEth:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                    `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffIn]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                     `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffOut]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaNativeAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiff struct {
	Asset SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiff]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiff) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiff() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffAsset struct {
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
	Type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffAsset]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffAssetTypeToken SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffAssetType = "TOKEN"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffAssetTypeToken:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                         `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffIn]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                          `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffOut]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplFungibleAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiff struct {
	Asset SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiff]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiff) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiff() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	// URL of the asset's logo
	Logo string                                                                                                                `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffAsset]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffAssetTypeNFT SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffAssetType = "NFT"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffAssetTypeNFT:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                            `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffIn]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                             `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffOut]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaSplNonFungibleAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiff struct {
	Asset SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffOut  `json:"out,nullable"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiff]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiff) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiff() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	// URL of the asset's logo
	Logo string                                                                                                      `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffAsset]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffAssetType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffAssetTypeCnft SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffAssetType = "CNFT"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffAssetTypeCnft:
		return true
	}
	return false
}

// Details of the incoming transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffIn struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                  `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffInJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffInJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffIn]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffInJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffOut struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                   `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffOutJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffOutJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffOut]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffOutJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaAssetsDiffSolanaCnftAssetDiffOutJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegation struct {
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationAsset],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationAsset],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationAsset].
	Asset     interface{} `json:"asset,required"`
	AssetType string      `json:"asset_type,required"`
	Delegate  string      `json:"delegate,required"`
	// This field can have the runtime type of
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationDelegation],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationDelegation],
	// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationDelegation].
	Delegation interface{}                                                                         `json:"delegation,required"`
	JSON       solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationJSON `json:"-"`
	union      SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsUnion
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegation]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Delegate    apijson.Field
	Delegation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegation) UnmarshalJSON(data []byte) (err error) {
	*r = SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegation],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegation],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegation].
func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegation) AsUnion() SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsUnion {
	return r.union
}

// Union satisfied by
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegation],
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegation]
// or
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegation].
type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsUnion interface {
	implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegation{}),
		},
	)
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegation struct {
	Asset      SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationAsset      `json:"asset,required"`
	AssetType  string                                                                                                         `json:"asset_type,required"`
	Delegate   string                                                                                                         `json:"delegate,required"`
	Delegation SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationDelegation `json:"delegation,required"`
	JSON       solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationJSON       `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegation]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Delegate    apijson.Field
	Delegation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegation) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegation() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	// URL of the asset's logo
	Logo string                                                                                                        `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationAsset]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationAssetType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationAssetTypeCnft SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationAssetType = "CNFT"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationAssetTypeCnft:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationDelegation struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                            `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationDelegationJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationDelegation]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationDelegationJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaCnftDelegationDelegationJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegation struct {
	Asset      SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationAsset      `json:"asset,required"`
	AssetType  string                                                                                                                     `json:"asset_type,required"`
	Delegate   string                                                                                                                     `json:"delegate,required"`
	Delegation SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationDelegation `json:"delegation,required"`
	JSON       solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationJSON       `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegation]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Delegate    apijson.Field
	Delegation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegation) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegation() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's decimals
	Decimals int64 `json:"decimals,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	Logo string                                                                                                                    `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationAsset]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationAssetJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationAssetType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationAssetTypeToken SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationAssetType = "TOKEN"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationAssetTypeToken:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationDelegation struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                        `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationDelegationJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationDelegation]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationDelegationJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaFungibleSplTokenDelegationDelegationJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegation struct {
	Asset      SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationAsset      `json:"asset,required"`
	AssetType  string                                                                                                                        `json:"asset_type,required"`
	Delegate   string                                                                                                                        `json:"delegate,required"`
	Delegation SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationDelegation `json:"delegation,required"`
	JSON       solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationJSON       `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegation]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	Delegate    apijson.Field
	Delegation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegation) implementsSolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegation() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationAsset struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol   string `json:"symbol,required"`
	Decimals int64  `json:"decimals"`
	// URL of the asset's logo
	Logo string                                                                                                                       `json:"logo,nullable"`
	Type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationAssetType `json:"type"`
	JSON solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationAssetJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationAssetJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationAsset]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationAssetJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Decimals    apijson.Field
	Logo        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationAssetJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationAssetType string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationAssetTypeNFT SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationAssetType = "NFT"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationAssetType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationAssetTypeNFT:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationDelegation struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                                                                                                                           `json:"usd_price,nullable"`
	JSON     solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationDelegationJSON `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationDelegationJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationDelegation]
type solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationDelegationJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationDelegation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationResultSchemaDelegationsSolanaNonFungibleSplTokenDelegationDelegationJSON) RawJSON() string {
	return r.raw
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationErrorSchema struct {
	// Error message
	Error  string                                                                     `json:"error,required"`
	Status SolanaMessageScanResponseResultSimulationSolanaSimulationErrorSchemaStatus `json:"status,required"`
	JSON   solanaMessageScanResponseResultSimulationSolanaSimulationErrorSchemaJSON   `json:"-"`
}

// solanaMessageScanResponseResultSimulationSolanaSimulationErrorSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultSimulationSolanaSimulationErrorSchema]
type solanaMessageScanResponseResultSimulationSolanaSimulationErrorSchemaJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultSimulationSolanaSimulationErrorSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultSimulationSolanaSimulationErrorSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationErrorSchema) implementsSolanaMessageScanResponseResultSimulation() {
}

type SolanaMessageScanResponseResultSimulationSolanaSimulationErrorSchemaStatus string

const (
	SolanaMessageScanResponseResultSimulationSolanaSimulationErrorSchemaStatusError SolanaMessageScanResponseResultSimulationSolanaSimulationErrorSchemaStatus = "Error"
)

func (r SolanaMessageScanResponseResultSimulationSolanaSimulationErrorSchemaStatus) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationSolanaSimulationErrorSchemaStatusError:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultSimulationStatus string

const (
	SolanaMessageScanResponseResultSimulationStatusError SolanaMessageScanResponseResultSimulationStatus = "Error"
)

func (r SolanaMessageScanResponseResultSimulationStatus) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultSimulationStatusError:
		return true
	}
	return false
}

// Validation result; Only present if validation option is included in the request
type SolanaMessageScanResponseResultValidation struct {
	Status SolanaMessageScanResponseResultValidationStatus `json:"status,required"`
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification"`
	// A textual description about the validation result
	Description string `json:"description"`
	// Error message
	Error string `json:"error"`
	// This field can have the runtime type of
	// [[]SolanaMessageScanResponseResultValidationSolanaValidationResultFeature].
	Features interface{} `json:"features"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason string `json:"reason"`
	// Verdict of the validation
	ResultType SolanaMessageScanResponseResultValidationResultType `json:"result_type"`
	JSON       solanaMessageScanResponseResultValidationJSON       `json:"-"`
	union      SolanaMessageScanResponseResultValidationUnion
}

// solanaMessageScanResponseResultValidationJSON contains the JSON metadata for the
// struct [SolanaMessageScanResponseResultValidation]
type solanaMessageScanResponseResultValidationJSON struct {
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

func (r solanaMessageScanResponseResultValidationJSON) RawJSON() string {
	return r.raw
}

func (r *SolanaMessageScanResponseResultValidation) UnmarshalJSON(data []byte) (err error) {
	*r = SolanaMessageScanResponseResultValidation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SolanaMessageScanResponseResultValidationUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SolanaMessageScanResponseResultValidationSolanaValidationResult],
// [SolanaMessageScanResponseResultValidationSolanaValidationErrorSchema].
func (r SolanaMessageScanResponseResultValidation) AsUnion() SolanaMessageScanResponseResultValidationUnion {
	return r.union
}

// Validation result; Only present if validation option is included in the request
//
// Union satisfied by
// [SolanaMessageScanResponseResultValidationSolanaValidationResult] or
// [SolanaMessageScanResponseResultValidationSolanaValidationErrorSchema].
type SolanaMessageScanResponseResultValidationUnion interface {
	implementsSolanaMessageScanResponseResultValidation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SolanaMessageScanResponseResultValidationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultValidationSolanaValidationResult{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SolanaMessageScanResponseResultValidationSolanaValidationErrorSchema{}),
		},
	)
}

type SolanaMessageScanResponseResultValidationSolanaValidationResult struct {
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification,required"`
	// A textual description about the validation result
	Description string                                                                   `json:"description,required"`
	Features    []SolanaMessageScanResponseResultValidationSolanaValidationResultFeature `json:"features,required"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason string `json:"reason,required"`
	// Verdict of the validation
	ResultType SolanaMessageScanResponseResultValidationSolanaValidationResultResultType `json:"result_type,required"`
	Status     SolanaMessageScanResponseResultValidationSolanaValidationResultStatus     `json:"status,required"`
	JSON       solanaMessageScanResponseResultValidationSolanaValidationResultJSON       `json:"-"`
}

// solanaMessageScanResponseResultValidationSolanaValidationResultJSON contains the
// JSON metadata for the struct
// [SolanaMessageScanResponseResultValidationSolanaValidationResult]
type solanaMessageScanResponseResultValidationSolanaValidationResultJSON struct {
	Classification apijson.Field
	Description    apijson.Field
	Features       apijson.Field
	Reason         apijson.Field
	ResultType     apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultValidationSolanaValidationResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultValidationSolanaValidationResultJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultValidationSolanaValidationResult) implementsSolanaMessageScanResponseResultValidation() {
}

type SolanaMessageScanResponseResultValidationSolanaValidationResultFeature struct {
	// Address the feature refers to
	Address string `json:"address,required"`
	// Textual description
	Description string `json:"description,required"`
	FeatureID   string `json:"feature_id,required"`
	// Feature Classification
	Type SolanaMessageScanResponseResultValidationSolanaValidationResultFeaturesType `json:"type,required"`
	JSON solanaMessageScanResponseResultValidationSolanaValidationResultFeatureJSON  `json:"-"`
}

// solanaMessageScanResponseResultValidationSolanaValidationResultFeatureJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultValidationSolanaValidationResultFeature]
type solanaMessageScanResponseResultValidationSolanaValidationResultFeatureJSON struct {
	Address     apijson.Field
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultValidationSolanaValidationResultFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultValidationSolanaValidationResultFeatureJSON) RawJSON() string {
	return r.raw
}

// Feature Classification
type SolanaMessageScanResponseResultValidationSolanaValidationResultFeaturesType string

const (
	SolanaMessageScanResponseResultValidationSolanaValidationResultFeaturesTypeBenign    SolanaMessageScanResponseResultValidationSolanaValidationResultFeaturesType = "Benign"
	SolanaMessageScanResponseResultValidationSolanaValidationResultFeaturesTypeWarning   SolanaMessageScanResponseResultValidationSolanaValidationResultFeaturesType = "Warning"
	SolanaMessageScanResponseResultValidationSolanaValidationResultFeaturesTypeMalicious SolanaMessageScanResponseResultValidationSolanaValidationResultFeaturesType = "Malicious"
	SolanaMessageScanResponseResultValidationSolanaValidationResultFeaturesTypeInfo      SolanaMessageScanResponseResultValidationSolanaValidationResultFeaturesType = "Info"
)

func (r SolanaMessageScanResponseResultValidationSolanaValidationResultFeaturesType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultValidationSolanaValidationResultFeaturesTypeBenign, SolanaMessageScanResponseResultValidationSolanaValidationResultFeaturesTypeWarning, SolanaMessageScanResponseResultValidationSolanaValidationResultFeaturesTypeMalicious, SolanaMessageScanResponseResultValidationSolanaValidationResultFeaturesTypeInfo:
		return true
	}
	return false
}

// Verdict of the validation
type SolanaMessageScanResponseResultValidationSolanaValidationResultResultType string

const (
	SolanaMessageScanResponseResultValidationSolanaValidationResultResultTypeBenign    SolanaMessageScanResponseResultValidationSolanaValidationResultResultType = "Benign"
	SolanaMessageScanResponseResultValidationSolanaValidationResultResultTypeWarning   SolanaMessageScanResponseResultValidationSolanaValidationResultResultType = "Warning"
	SolanaMessageScanResponseResultValidationSolanaValidationResultResultTypeMalicious SolanaMessageScanResponseResultValidationSolanaValidationResultResultType = "Malicious"
)

func (r SolanaMessageScanResponseResultValidationSolanaValidationResultResultType) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultValidationSolanaValidationResultResultTypeBenign, SolanaMessageScanResponseResultValidationSolanaValidationResultResultTypeWarning, SolanaMessageScanResponseResultValidationSolanaValidationResultResultTypeMalicious:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultValidationSolanaValidationResultStatus string

const (
	SolanaMessageScanResponseResultValidationSolanaValidationResultStatusSuccess SolanaMessageScanResponseResultValidationSolanaValidationResultStatus = "Success"
)

func (r SolanaMessageScanResponseResultValidationSolanaValidationResultStatus) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultValidationSolanaValidationResultStatusSuccess:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultValidationSolanaValidationErrorSchema struct {
	// Error message
	Error  string                                                                     `json:"error,required"`
	Status SolanaMessageScanResponseResultValidationSolanaValidationErrorSchemaStatus `json:"status,required"`
	JSON   solanaMessageScanResponseResultValidationSolanaValidationErrorSchemaJSON   `json:"-"`
}

// solanaMessageScanResponseResultValidationSolanaValidationErrorSchemaJSON
// contains the JSON metadata for the struct
// [SolanaMessageScanResponseResultValidationSolanaValidationErrorSchema]
type solanaMessageScanResponseResultValidationSolanaValidationErrorSchemaJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaMessageScanResponseResultValidationSolanaValidationErrorSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaMessageScanResponseResultValidationSolanaValidationErrorSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SolanaMessageScanResponseResultValidationSolanaValidationErrorSchema) implementsSolanaMessageScanResponseResultValidation() {
}

type SolanaMessageScanResponseResultValidationSolanaValidationErrorSchemaStatus string

const (
	SolanaMessageScanResponseResultValidationSolanaValidationErrorSchemaStatusError SolanaMessageScanResponseResultValidationSolanaValidationErrorSchemaStatus = "Error"
)

func (r SolanaMessageScanResponseResultValidationSolanaValidationErrorSchemaStatus) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultValidationSolanaValidationErrorSchemaStatusError:
		return true
	}
	return false
}

type SolanaMessageScanResponseResultValidationStatus string

const (
	SolanaMessageScanResponseResultValidationStatusSuccess SolanaMessageScanResponseResultValidationStatus = "Success"
	SolanaMessageScanResponseResultValidationStatusError   SolanaMessageScanResponseResultValidationStatus = "Error"
)

func (r SolanaMessageScanResponseResultValidationStatus) IsKnown() bool {
	switch r {
	case SolanaMessageScanResponseResultValidationStatusSuccess, SolanaMessageScanResponseResultValidationStatusError:
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
	Body interface{} `json:"body,required"`
}

func (r SolanaMessageScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
