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

// BitcoinTransactionService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBitcoinTransactionService] method instead.
type BitcoinTransactionService struct {
	Options []option.RequestOption
}

// NewBitcoinTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBitcoinTransactionService(opts ...option.RequestOption) (r *BitcoinTransactionService) {
	r = &BitcoinTransactionService{}
	r.Options = opts
	return
}

// Report Transaction
func (r *BitcoinTransactionService) Report(ctx context.Context, body BitcoinTransactionReportParams, opts ...option.RequestOption) (res *int64, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/bitcoin/transaction/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Scan Transaction
func (r *BitcoinTransactionService) Scan(ctx context.Context, body BitcoinTransactionScanParams, opts ...option.RequestOption) (res *BitcoinTransactionScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/bitcoin/transaction/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BitcoinTransactionScanResponse struct {
	// Simulation result; Only present if simulation option is included in the request
	Simulation BitcoinTransactionScanResponseSimulation `json:"simulation,nullable"`
	// Validation result; Only present if validation option is included in the request
	Validation BitcoinTransactionScanResponseValidation `json:"validation,nullable"`
	JSON       bitcoinTransactionScanResponseJSON       `json:"-"`
}

// bitcoinTransactionScanResponseJSON contains the JSON metadata for the struct
// [BitcoinTransactionScanResponse]
type bitcoinTransactionScanResponseJSON struct {
	Simulation  apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseJSON) RawJSON() string {
	return r.raw
}

// Simulation result; Only present if simulation option is included in the request
type BitcoinTransactionScanResponseSimulation struct {
	// This field can have the runtime type of
	// [[]BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAddressDetail].
	AddressDetails interface{} `json:"address_details,required"`
	// This field can have the runtime type of
	// [map[string][]BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiff].
	AssetsDiffs    interface{}                                    `json:"assets_diffs,required"`
	Status         BitcoinTransactionScanResponseSimulationStatus `json:"status,required"`
	AccountSummary interface{}                                    `json:"account_summary"`
	// Error message
	Error string                                       `json:"error"`
	JSON  bitcoinTransactionScanResponseSimulationJSON `json:"-"`
	union BitcoinTransactionScanResponseSimulationUnion
}

// bitcoinTransactionScanResponseSimulationJSON contains the JSON metadata for the
// struct [BitcoinTransactionScanResponseSimulation]
type bitcoinTransactionScanResponseSimulationJSON struct {
	AddressDetails apijson.Field
	AssetsDiffs    apijson.Field
	Status         apijson.Field
	AccountSummary apijson.Field
	Error          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r bitcoinTransactionScanResponseSimulationJSON) RawJSON() string {
	return r.raw
}

func (r *BitcoinTransactionScanResponseSimulation) UnmarshalJSON(data []byte) (err error) {
	*r = BitcoinTransactionScanResponseSimulation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BitcoinTransactionScanResponseSimulationUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfiguration],
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchema].
func (r BitcoinTransactionScanResponseSimulation) AsUnion() BitcoinTransactionScanResponseSimulationUnion {
	return r.union
}

// Simulation result; Only present if simulation option is included in the request
//
// Union satisfied by
// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfiguration]
// or [BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchema].
type BitcoinTransactionScanResponseSimulationUnion interface {
	implementsBitcoinTransactionScanResponseSimulation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BitcoinTransactionScanResponseSimulationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchema{}),
		},
	)
}

type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfiguration struct {
	Status         BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationStatus `json:"status,required"`
	AccountSummary interface{}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   `json:"account_summary"`
	// Details of addresses involved in the transaction
	AddressDetails []BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAddressDetail `json:"address_details"`
	// Mapping between the address of an account to the assets diff during the
	// transaction
	AssetsDiffs map[string][]BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiff `json:"assets_diffs"`
	JSON        bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationJSON                    `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfiguration]
type bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationJSON struct {
	Status         apijson.Field
	AccountSummary apijson.Field
	AddressDetails apijson.Field
	AssetsDiffs    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfiguration) implementsBitcoinTransactionScanResponseSimulation() {
}

type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationStatus string

const (
	BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationStatusSuccess BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationStatus = "Success"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationStatus) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationStatusSuccess:
		return true
	}
	return false
}

type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAddressDetail struct {
	// Encoded public key of the account
	AccountAddress string `json:"account_address,required"`
	// Description of the account
	Description string                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   `json:"description,nullable"`
	JSON        bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAddressDetailJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAddressDetailJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAddressDetail]
type bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAddressDetailJSON struct {
	AccountAddress apijson.Field
	Description    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAddressDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAddressDetailJSON) RawJSON() string {
	return r.raw
}

type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiff struct {
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// This field can have the runtime type of
	// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaIn],
	// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaIn],
	// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaIn].
	In interface{} `json:"in,required"`
	// This field can have the runtime type of
	// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaOut],
	// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaOut],
	// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaOut].
	Out interface{} `json:"out,required"`
	// This field can have the runtime type of
	// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAsset],
	// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaAsset],
	// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaAsset].
	Asset interface{}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           `json:"asset"`
	JSON  bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffJSON `json:"-"`
	union BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsUnion
}

// bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiff]
type bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffJSON struct {
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	Asset       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffJSON) RawJSON() string {
	return r.raw
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiff) UnmarshalJSON(data []byte) (err error) {
	*r = BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiff{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchema],
// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchema],
// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchema].
func (r BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiff) AsUnion() BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsUnion {
	return r.union
}

// Union satisfied by
// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchema],
// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchema]
// or
// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchema].
type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsUnion interface {
	implementsBitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiff()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchema{}),
		},
	)
}

type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchema struct {
	Asset BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaOut  `json:"out,nullable"`
	JSON bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchema]
type bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchema) implementsBitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiff() {
}

type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAsset struct {
	// URL of the asset's logo
	LogoURL string `json:"logo_url,required,nullable"`
	// Decimals of the asset
	Decimals BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetDecimals `json:"decimals"`
	// Name of the asset
	Name BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetName `json:"name"`
	// Symbol of the asset
	Symbol BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetSymbol `json:"symbol"`
	// Type of the asset (`NATIVE`)
	Type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetType `json:"type"`
	JSON bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAsset]
type bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetJSON struct {
	LogoURL     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Decimals of the asset
type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetDecimals int64

const (
	BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetDecimals8 BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetDecimals = 8
)

func (r BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetDecimals) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetDecimals8:
		return true
	}
	return false
}

// Name of the asset
type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetName string

const (
	BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetNameBitcoin BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetName = "Bitcoin"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetName) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetNameBitcoin:
		return true
	}
	return false
}

// Symbol of the asset
type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetSymbol string

const (
	BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetSymbolBtc BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetSymbol = "BTC"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetSymbol) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetSymbolBtc:
		return true
	}
	return false
}

// Type of the asset (`NATIVE`)
type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetType string

const (
	BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetTypeNative BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetType = "NATIVE"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaAssetTypeNative:
		return true
	}
	return false
}

// Details of the incoming transfer
type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue string `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  `json:"summary,nullable"`
	JSON    bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaInJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaInJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaIn]
type bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaInJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue string `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   `json:"summary,nullable"`
	JSON    bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaOutJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaOutJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaOut]
type bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchema struct {
	Asset BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaOut  `json:"out,nullable"`
	JSON bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchema]
type bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchema) implementsBitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiff() {
}

type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaAsset struct {
	// token's name
	Name string `json:"name,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ORDINAL`)
	Type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaAssetType `json:"type"`
	JSON bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaAssetJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaAsset]
type bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaAssetJSON struct {
	Name        apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ORDINAL`)
type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaAssetType string

const (
	BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaAssetTypeOrdinal BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaAssetType = "ORDINAL"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaAssetType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaAssetTypeOrdinal:
		return true
	}
	return false
}

// Details of the incoming transfer
type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaIn struct {
	// Id of the ordinal
	ID string `json:"id,required"`
	// Raw value of the transfer
	RawValue string `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               `json:"summary,nullable"`
	JSON    bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaInJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaInJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaIn]
type bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaInJSON struct {
	ID          apijson.Field
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaOut struct {
	// Id of the ordinal
	ID string `json:"id,required"`
	// Raw value of the transfer
	RawValue string `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                `json:"summary,nullable"`
	JSON    bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaOutJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaOutJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaOut]
type bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaOutJSON struct {
	ID          apijson.Field
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchema struct {
	Asset BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaAsset `json:"asset,required"`
	// The type of the assets in this diff
	AssetType string `json:"asset_type,required"`
	// Details of the incoming transfer
	In BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaIn `json:"in,nullable"`
	// Details of the outgoing transfer
	Out  BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaOut  `json:"out,nullable"`
	JSON bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchema]
type bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaJSON struct {
	Asset       apijson.Field
	AssetType   apijson.Field
	In          apijson.Field
	Out         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaJSON) RawJSON() string {
	return r.raw
}

func (r BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchema) implementsBitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiff() {
}

type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaAsset struct {
	// The Rune ID
	ID string `json:"id,required"`
	// Decimals of the asset
	Decimals int64 `json:"decimals,required"`
	// The Rune name
	Name string `json:"name,required"`
	// The Rune spaced name
	SpacedName string `json:"spaced_name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`RUNE`)
	Type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaAssetType `json:"type"`
	JSON bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaAssetJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaAssetJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaAsset]
type bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaAssetJSON struct {
	ID          apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	SpacedName  apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaAssetJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`RUNE`)
type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaAssetType string

const (
	BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaAssetTypeRune BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaAssetType = "RUNE"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaAssetType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaAssetTypeRune:
		return true
	}
	return false
}

// Details of the incoming transfer
type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaIn struct {
	// Raw value of the transfer
	RawValue string `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         `json:"summary,nullable"`
	JSON    bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaInJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaInJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaIn]
type bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaInJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaInJSON) RawJSON() string {
	return r.raw
}

// Details of the outgoing transfer
type BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaOut struct {
	// Raw value of the transfer
	RawValue string `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          `json:"summary,nullable"`
	JSON    bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaOutJSON `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaOutJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaOut]
type bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaOutJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinBitcoinSimulationSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaUnionAnnotatedAccountSingleAssetDiffSchemaTypeNativeAssetDetailsSchemaNativeDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleNativeAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeOrdinalDetailsSchemaOrdinalDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleOrdinalAssetDiffAnnotatedAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaFieldInfoAnnotationNoneTypeRequiredTrueTitleRunesAssetDiffEmptyModelEmptyModelAddressDetailsBaseSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedEmptyModelSimulationSchemaConfigurationAssetsDiffsBitcoinAccountSingleAssetDiffSchemaTypeRuneDetailsSchemaRuneDiffSchemaOutJSON) RawJSON() string {
	return r.raw
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchema struct {
	// Error message
	Error  string                                                                     `json:"error,required"`
	Status BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchemaStatus `json:"status,required"`
	JSON   bitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchemaJSON   `json:"-"`
}

// bitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchemaJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchema]
type bitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchemaJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchemaJSON) RawJSON() string {
	return r.raw
}

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchema) implementsBitcoinTransactionScanResponseSimulation() {
}

type BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchemaStatus string

const (
	BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchemaStatusError BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchemaStatus = "Error"
)

func (r BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchemaStatus) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationBitcoinSimulationErrorSchemaStatusError:
		return true
	}
	return false
}

type BitcoinTransactionScanResponseSimulationStatus string

const (
	BitcoinTransactionScanResponseSimulationStatusSuccess BitcoinTransactionScanResponseSimulationStatus = "Success"
	BitcoinTransactionScanResponseSimulationStatusError   BitcoinTransactionScanResponseSimulationStatus = "Error"
)

func (r BitcoinTransactionScanResponseSimulationStatus) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseSimulationStatusSuccess, BitcoinTransactionScanResponseSimulationStatusError:
		return true
	}
	return false
}

// Validation result; Only present if validation option is included in the request
type BitcoinTransactionScanResponseValidation struct {
	// This field can have the runtime type of
	// [[]BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaFeature].
	Features interface{}                                    `json:"features,required"`
	Status   BitcoinTransactionScanResponseValidationStatus `json:"status,required"`
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification"`
	// A textual description about the validation result
	Description string `json:"description"`
	// Error message
	Error string `json:"error"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason string `json:"reason"`
	// Verdict of the validation
	ResultType BitcoinTransactionScanResponseValidationResultType `json:"result_type"`
	JSON       bitcoinTransactionScanResponseValidationJSON       `json:"-"`
	union      BitcoinTransactionScanResponseValidationUnion
}

// bitcoinTransactionScanResponseValidationJSON contains the JSON metadata for the
// struct [BitcoinTransactionScanResponseValidation]
type bitcoinTransactionScanResponseValidationJSON struct {
	Features       apijson.Field
	Status         apijson.Field
	Classification apijson.Field
	Description    apijson.Field
	Error          apijson.Field
	Reason         apijson.Field
	ResultType     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r bitcoinTransactionScanResponseValidationJSON) RawJSON() string {
	return r.raw
}

func (r *BitcoinTransactionScanResponseValidation) UnmarshalJSON(data []byte) (err error) {
	*r = BitcoinTransactionScanResponseValidation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BitcoinTransactionScanResponseValidationUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchema],
// [BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchema].
func (r BitcoinTransactionScanResponseValidation) AsUnion() BitcoinTransactionScanResponseValidationUnion {
	return r.union
}

// Validation result; Only present if validation option is included in the request
//
// Union satisfied by
// [BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchema]
// or [BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchema].
type BitcoinTransactionScanResponseValidationUnion interface {
	implementsBitcoinTransactionScanResponseValidation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BitcoinTransactionScanResponseValidationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchema{}),
		},
	)
}

type BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchema struct {
	// A textual classification that can be presented to the user explaining the
	// reason.
	Classification string `json:"classification,required"`
	// A textual description about the validation result
	Description string `json:"description,required"`
	// A list of features about this transaction explaining the validation
	Features []BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaFeature `json:"features,required"`
	// A textual description about the reasons the transaction was flagged with
	// result_type
	Reason string `json:"reason,required"`
	// Verdict of the validation
	ResultType BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaResultType `json:"result_type,required"`
	Status     BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaStatus     `json:"status,required"`
	JSON       bitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaJSON       `json:"-"`
}

// bitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchema]
type bitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaJSON struct {
	Classification apijson.Field
	Description    apijson.Field
	Features       apijson.Field
	Reason         apijson.Field
	ResultType     apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaJSON) RawJSON() string {
	return r.raw
}

func (r BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchema) implementsBitcoinTransactionScanResponseValidation() {
}

type BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaFeature struct {
	// Address the feature refers to
	Address string `json:"address,required"`
	// Textual description
	Description string `json:"description,required"`
	FeatureID string `json:"feature_id,required"`
	// Feature Classification
	Type BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaFeaturesType `json:"type,required"`
	JSON bitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaFeatureJSON  `json:"-"`
}

// bitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaFeatureJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaFeature]
type bitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaFeatureJSON struct {
	Address     apijson.Field
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaFeatureJSON) RawJSON() string {
	return r.raw
}

// Feature Classification
type BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaFeaturesType string

const (
	BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaFeaturesTypeBenign    BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaFeaturesType = "Benign"
	BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaFeaturesTypeWarning   BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaFeaturesType = "Warning"
	BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaFeaturesTypeMalicious BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaFeaturesType = "Malicious"
	BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaFeaturesTypeInfo      BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaFeaturesType = "Info"
)

func (r BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaFeaturesType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaFeaturesTypeBenign, BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaFeaturesTypeWarning, BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaFeaturesTypeMalicious, BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaFeaturesTypeInfo:
		return true
	}
	return false
}

// Verdict of the validation
type BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaResultType string

const (
	BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaResultTypeBenign    BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaResultType = "Benign"
	BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaResultTypeWarning   BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaResultType = "Warning"
	BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaResultTypeMalicious BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaResultType = "Malicious"
)

func (r BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaResultType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaResultTypeBenign, BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaResultTypeWarning, BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaResultTypeMalicious:
		return true
	}
	return false
}

type BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaStatus string

const (
	BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaStatusSuccess BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaStatus = "Success"
)

func (r BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaStatus) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseValidationBitcoinValidationResultSchemaTypeAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaStatusSuccess:
		return true
	}
	return false
}

type BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchema struct {
	// Error message
	Error  string                                                                     `json:"error,required"`
	Status BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchemaStatus `json:"status,required"`
	JSON   bitcoinTransactionScanResponseValidationBitcoinValidationErrorSchemaJSON   `json:"-"`
}

// bitcoinTransactionScanResponseValidationBitcoinValidationErrorSchemaJSON
// contains the JSON metadata for the struct
// [BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchema]
type bitcoinTransactionScanResponseValidationBitcoinValidationErrorSchemaJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bitcoinTransactionScanResponseValidationBitcoinValidationErrorSchemaJSON) RawJSON() string {
	return r.raw
}

func (r BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchema) implementsBitcoinTransactionScanResponseValidation() {
}

type BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchemaStatus string

const (
	BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchemaStatusError BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchemaStatus = "Error"
)

func (r BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchemaStatus) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseValidationBitcoinValidationErrorSchemaStatusError:
		return true
	}
	return false
}

type BitcoinTransactionScanResponseValidationStatus string

const (
	BitcoinTransactionScanResponseValidationStatusSuccess BitcoinTransactionScanResponseValidationStatus = "Success"
	BitcoinTransactionScanResponseValidationStatusError   BitcoinTransactionScanResponseValidationStatus = "Error"
)

func (r BitcoinTransactionScanResponseValidationStatus) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseValidationStatusSuccess, BitcoinTransactionScanResponseValidationStatusError:
		return true
	}
	return false
}

// Verdict of the validation
type BitcoinTransactionScanResponseValidationResultType string

const (
	BitcoinTransactionScanResponseValidationResultTypeBenign    BitcoinTransactionScanResponseValidationResultType = "Benign"
	BitcoinTransactionScanResponseValidationResultTypeWarning   BitcoinTransactionScanResponseValidationResultType = "Warning"
	BitcoinTransactionScanResponseValidationResultTypeMalicious BitcoinTransactionScanResponseValidationResultType = "Malicious"
)

func (r BitcoinTransactionScanResponseValidationResultType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanResponseValidationResultTypeBenign, BitcoinTransactionScanResponseValidationResultTypeWarning, BitcoinTransactionScanResponseValidationResultTypeMalicious:
		return true
	}
	return false
}

type BitcoinTransactionReportParams struct {
	Details param.Field[string]                                    `json:"details,required"`
	Event   param.Field[BitcoinTransactionReportParamsEvent]       `json:"event,required"`
	Report  param.Field[BitcoinTransactionReportParamsReportUnion] `json:"report,required"`
}

func (r BitcoinTransactionReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BitcoinTransactionReportParamsEvent string

const (
	BitcoinTransactionReportParamsEventShouldBeMalicious BitcoinTransactionReportParamsEvent = "should_be_malicious"
	BitcoinTransactionReportParamsEventShouldBeBenign    BitcoinTransactionReportParamsEvent = "should_be_benign"
)

func (r BitcoinTransactionReportParamsEvent) IsKnown() bool {
	switch r {
	case BitcoinTransactionReportParamsEventShouldBeMalicious, BitcoinTransactionReportParamsEventShouldBeBenign:
		return true
	}
	return false
}

type BitcoinTransactionReportParamsReport struct {
	Params param.Field[interface{}]                              `json:"params,required"`
	ID     param.Field[string]                                   `json:"id"`
	Type   param.Field[BitcoinTransactionReportParamsReportType] `json:"type"`
}

func (r BitcoinTransactionReportParamsReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionReportParamsReport) implementsBitcoinTransactionReportParamsReportUnion() {}

// Satisfied by [BitcoinTransactionReportParamsReportBitcoinAppealRequestID],
// [BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1],
// [BitcoinTransactionReportParamsReport].
type BitcoinTransactionReportParamsReportUnion interface {
	implementsBitcoinTransactionReportParamsReportUnion()
}

type BitcoinTransactionReportParamsReportBitcoinAppealRequestID struct {
	ID   param.Field[string]                                                         `json:"id,required"`
	Type param.Field[BitcoinTransactionReportParamsReportBitcoinAppealRequestIDType] `json:"type"`
}

func (r BitcoinTransactionReportParamsReportBitcoinAppealRequestID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionReportParamsReportBitcoinAppealRequestID) implementsBitcoinTransactionReportParamsReportUnion() {
}

type BitcoinTransactionReportParamsReportBitcoinAppealRequestIDType string

const (
	BitcoinTransactionReportParamsReportBitcoinAppealRequestIDTypeRequestID BitcoinTransactionReportParamsReportBitcoinAppealRequestIDType = "request_id"
)

func (r BitcoinTransactionReportParamsReportBitcoinAppealRequestIDType) IsKnown() bool {
	switch r {
	case BitcoinTransactionReportParamsReportBitcoinAppealRequestIDTypeRequestID:
		return true
	}
	return false
}

type BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1 struct {
	Params param.Field[BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1Params] `json:"params,required"`
	Type   param.Field[BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1Type]   `json:"type"`
}

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1) implementsBitcoinTransactionReportParamsReportUnion() {
}

type BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1Params struct {
	AccountAddress param.Field[string]                                                                                                                                                                                                                                                                                                                                                         `json:"account_address,required"`
	Chain          param.Field[BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsChain] `json:"chain,required"`
	// Metadata
	Metadata    param.Field[BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataUnion] `json:"metadata,required"`
	Transaction param.Field[string]                                                                                                                                                                                                                                                                                                                                                                 `json:"transaction,required"`
	// List of options to include in the response
	//
	// - `Options.validation`: Include Options.validation output in the response
	//
	// - `Options.simulation`: Include Options.simulation output in the response
	Options param.Field[[]BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsOption] `json:"options"`
}

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsChain string

const (
	BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsChainBitcoin BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsChain = "bitcoin"
)

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsChain) IsKnown() bool {
	switch r {
	case BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsChainBitcoin:
		return true
	}
	return false
}

// Metadata
type BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadata struct {
	// Metadata for wallet requests
	Type param.Field[BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataType] `json:"type"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url"`
}

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadata) implementsBitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataUnion() {
}

// Metadata
//
// Satisfied by
// [BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIWalletRequestMetadata],
// [BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIInAppRequestMetadata],
// [BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadata].
type BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataUnion interface {
	implementsBitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataUnion()
}

type BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIWalletRequestMetadata struct {
	// Metadata for wallet requests
	Type param.Field[BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIWalletRequestMetadataType] `json:"type,required"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url,required"`
}

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIWalletRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIWalletRequestMetadata) implementsBitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataUnion() {
}

// Metadata for wallet requests
type BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIWalletRequestMetadataType string

const (
	BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIWalletRequestMetadataTypeWallet BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIWalletRequestMetadataType = "wallet"
)

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIWalletRequestMetadataType) IsKnown() bool {
	switch r {
	case BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIWalletRequestMetadataTypeWallet:
		return true
	}
	return false
}

type BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIInAppRequestMetadata struct {
	// Metadata for in-app requests
	Type param.Field[BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIInAppRequestMetadataType] `json:"type"`
}

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIInAppRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIInAppRequestMetadata) implementsBitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataUnion() {
}

// Metadata for in-app requests
type BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIInAppRequestMetadataType string

const (
	BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIInAppRequestMetadataTypeInApp BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIInAppRequestMetadataType = "in_app"
)

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIInAppRequestMetadataType) IsKnown() bool {
	switch r {
	case BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIInAppRequestMetadataTypeInApp:
		return true
	}
	return false
}

// Metadata for wallet requests
type BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataType string

const (
	BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataTypeWallet BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataType = "wallet"
	BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataTypeInApp  BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataType = "in_app"
)

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataType) IsKnown() bool {
	switch r {
	case BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataTypeWallet, BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsMetadataTypeInApp:
		return true
	}
	return false
}

type BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsOption string

const (
	BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsOptionValidation BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsOption = "validation"
	BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsOptionSimulation BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsOption = "simulation"
)

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsOption) IsKnown() bool {
	switch r {
	case BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsOptionValidation, BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1ParamsOptionSimulation:
		return true
	}
	return false
}

type BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1Type string

const (
	BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1TypeParams BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1Type = "params"
)

func (r BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1Type) IsKnown() bool {
	switch r {
	case BitcoinTransactionReportParamsReportBitcoinAppealTransactionDataReportRequestSchemaTypeChainSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedStrSkipValidationPlainSerializerGetPydanticSchemaAnnotatedSetOptionsFieldInfoAnnotationNoneTypeRequiredFalseValidationSimulationTitleOptionsSimulationMetadataMinLenMinLength1TypeParams:
		return true
	}
	return false
}

type BitcoinTransactionReportParamsReportType string

const (
	BitcoinTransactionReportParamsReportTypeRequestID BitcoinTransactionReportParamsReportType = "request_id"
	BitcoinTransactionReportParamsReportTypeParams    BitcoinTransactionReportParamsReportType = "params"
)

func (r BitcoinTransactionReportParamsReportType) IsKnown() bool {
	switch r {
	case BitcoinTransactionReportParamsReportTypeRequestID, BitcoinTransactionReportParamsReportTypeParams:
		return true
	}
	return false
}

type BitcoinTransactionScanParams struct {
	AccountAddress param.Field[string]                            `json:"account_address,required"`
	Chain          param.Field[BitcoinTransactionScanParamsChain] `json:"chain,required"`
	// Metadata
	Metadata    param.Field[BitcoinTransactionScanParamsMetadataUnion] `json:"metadata,required"`
	Transaction param.Field[string]                                    `json:"transaction,required"`
	// List of options to include in the response
	//
	// - `Options.validation`: Include Options.validation output in the response
	//
	// - `Options.simulation`: Include Options.simulation output in the response
	Options param.Field[[]BitcoinTransactionScanParamsOption] `json:"options"`
}

func (r BitcoinTransactionScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BitcoinTransactionScanParamsChain string

const (
	BitcoinTransactionScanParamsChainBitcoin BitcoinTransactionScanParamsChain = "bitcoin"
)

func (r BitcoinTransactionScanParamsChain) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanParamsChainBitcoin:
		return true
	}
	return false
}

// Metadata
type BitcoinTransactionScanParamsMetadata struct {
	// Metadata for wallet requests
	Type param.Field[BitcoinTransactionScanParamsMetadataType] `json:"type"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url"`
}

func (r BitcoinTransactionScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionScanParamsMetadata) implementsBitcoinTransactionScanParamsMetadataUnion() {}

// Metadata
//
// Satisfied by
// [BitcoinTransactionScanParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIWalletRequestMetadata],
// [BitcoinTransactionScanParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIInAppRequestMetadata],
// [BitcoinTransactionScanParamsMetadata].
type BitcoinTransactionScanParamsMetadataUnion interface {
	implementsBitcoinTransactionScanParamsMetadataUnion()
}

type BitcoinTransactionScanParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIWalletRequestMetadata struct {
	// Metadata for wallet requests
	Type param.Field[BitcoinTransactionScanParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIWalletRequestMetadataType] `json:"type,required"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url,required"`
}

func (r BitcoinTransactionScanParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIWalletRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionScanParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIWalletRequestMetadata) implementsBitcoinTransactionScanParamsMetadataUnion() {
}

// Metadata for wallet requests
type BitcoinTransactionScanParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIWalletRequestMetadataType string

const (
	BitcoinTransactionScanParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIWalletRequestMetadataTypeWallet BitcoinTransactionScanParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIWalletRequestMetadataType = "wallet"
)

func (r BitcoinTransactionScanParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIWalletRequestMetadataType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIWalletRequestMetadataTypeWallet:
		return true
	}
	return false
}

type BitcoinTransactionScanParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIInAppRequestMetadata struct {
	// Metadata for in-app requests
	Type param.Field[BitcoinTransactionScanParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIInAppRequestMetadataType] `json:"type"`
}

func (r BitcoinTransactionScanParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIInAppRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionScanParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIInAppRequestMetadata) implementsBitcoinTransactionScanParamsMetadataUnion() {
}

// Metadata for in-app requests
type BitcoinTransactionScanParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIInAppRequestMetadataType string

const (
	BitcoinTransactionScanParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIInAppRequestMetadataTypeInApp BitcoinTransactionScanParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIInAppRequestMetadataType = "in_app"
)

func (r BitcoinTransactionScanParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIInAppRequestMetadataType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIInAppRequestMetadataTypeInApp:
		return true
	}
	return false
}

// Metadata for wallet requests
type BitcoinTransactionScanParamsMetadataType string

const (
	BitcoinTransactionScanParamsMetadataTypeWallet BitcoinTransactionScanParamsMetadataType = "wallet"
	BitcoinTransactionScanParamsMetadataTypeInApp  BitcoinTransactionScanParamsMetadataType = "in_app"
)

func (r BitcoinTransactionScanParamsMetadataType) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanParamsMetadataTypeWallet, BitcoinTransactionScanParamsMetadataTypeInApp:
		return true
	}
	return false
}

type BitcoinTransactionScanParamsOption string

const (
	BitcoinTransactionScanParamsOptionValidation BitcoinTransactionScanParamsOption = "validation"
	BitcoinTransactionScanParamsOptionSimulation BitcoinTransactionScanParamsOption = "simulation"
)

func (r BitcoinTransactionScanParamsOption) IsKnown() bool {
	switch r {
	case BitcoinTransactionScanParamsOptionValidation, BitcoinTransactionScanParamsOptionSimulation:
		return true
	}
	return false
}
