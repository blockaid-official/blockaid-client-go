// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"net/http"
	"reflect"
	"time"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/param"
	"github.com/blockaid-official/blockaid-client-go/internal/requestconfig"
	"github.com/blockaid-official/blockaid-client-go/option"
	"github.com/tidwall/gjson"
)

// ChainAgnosticTransactionService contains methods and other services that help
// with interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChainAgnosticTransactionService] method instead.
type ChainAgnosticTransactionService struct {
	Options []option.RequestOption
}

// NewChainAgnosticTransactionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewChainAgnosticTransactionService(opts ...option.RequestOption) (r *ChainAgnosticTransactionService) {
	r = &ChainAgnosticTransactionService{}
	r.Options = opts
	return
}

// Gets a transaction and returns a full security assessment indicating whether or
// not the transaction is malicious, along with textual reasons explaining why it
// was flagged as such.
func (r *ChainAgnosticTransactionService) Scan(ctx context.Context, body ChainAgnosticTransactionScanParams, opts ...option.RequestOption) (res *ChainAgnosticTransactionScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/transaction/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ChainAgnosticTransactionScanResponse struct {
	// Complete validation result containing all scan details and findings
	Validation ChainAgnosticTransactionScanResponseValidation `json:"validation,required"`
	JSON       chainAgnosticTransactionScanResponseJSON       `json:"-"`
}

// chainAgnosticTransactionScanResponseJSON contains the JSON metadata for the
// struct [ChainAgnosticTransactionScanResponse]
type chainAgnosticTransactionScanResponseJSON struct {
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChainAgnosticTransactionScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chainAgnosticTransactionScanResponseJSON) RawJSON() string {
	return r.raw
}

// Complete validation result containing all scan details and findings
type ChainAgnosticTransactionScanResponseValidation struct {
	// Classification of the scan result based on the detected features
	Classification string `json:"classification,required"`
	// Detailed description of the validation result
	Description string `json:"description,required"`
	// The type of validation result
	ResultType ChainAgnosticTransactionScanResponseValidationResultType `json:"result_type,required"`
	// The status of the transaction scan
	Status ChainAgnosticTransactionScanResponseValidationStatus `json:"status,required"`
	// List of features detected during the scan
	Features []ChainAgnosticTransactionScanResponseValidationFeature `json:"features"`
	JSON     chainAgnosticTransactionScanResponseValidationJSON      `json:"-"`
}

// chainAgnosticTransactionScanResponseValidationJSON contains the JSON metadata
// for the struct [ChainAgnosticTransactionScanResponseValidation]
type chainAgnosticTransactionScanResponseValidationJSON struct {
	Classification apijson.Field
	Description    apijson.Field
	ResultType     apijson.Field
	Status         apijson.Field
	Features       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ChainAgnosticTransactionScanResponseValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chainAgnosticTransactionScanResponseValidationJSON) RawJSON() string {
	return r.raw
}

// The type of validation result
type ChainAgnosticTransactionScanResponseValidationResultType string

const (
	ChainAgnosticTransactionScanResponseValidationResultTypeMalicious ChainAgnosticTransactionScanResponseValidationResultType = "Malicious"
	ChainAgnosticTransactionScanResponseValidationResultTypeWarning   ChainAgnosticTransactionScanResponseValidationResultType = "Warning"
	ChainAgnosticTransactionScanResponseValidationResultTypeBenign    ChainAgnosticTransactionScanResponseValidationResultType = "Benign"
	ChainAgnosticTransactionScanResponseValidationResultTypeHighRisk  ChainAgnosticTransactionScanResponseValidationResultType = "High-Risk"
)

func (r ChainAgnosticTransactionScanResponseValidationResultType) IsKnown() bool {
	switch r {
	case ChainAgnosticTransactionScanResponseValidationResultTypeMalicious, ChainAgnosticTransactionScanResponseValidationResultTypeWarning, ChainAgnosticTransactionScanResponseValidationResultTypeBenign, ChainAgnosticTransactionScanResponseValidationResultTypeHighRisk:
		return true
	}
	return false
}

// The status of the transaction scan
type ChainAgnosticTransactionScanResponseValidationStatus string

const (
	ChainAgnosticTransactionScanResponseValidationStatusSuccess ChainAgnosticTransactionScanResponseValidationStatus = "Success"
	ChainAgnosticTransactionScanResponseValidationStatusError   ChainAgnosticTransactionScanResponseValidationStatus = "Error"
)

func (r ChainAgnosticTransactionScanResponseValidationStatus) IsKnown() bool {
	switch r {
	case ChainAgnosticTransactionScanResponseValidationStatusSuccess, ChainAgnosticTransactionScanResponseValidationStatusError:
		return true
	}
	return false
}

type ChainAgnosticTransactionScanResponseValidationFeature struct {
	// Detailed description of the detected feature
	Description string `json:"description,required"`
	// Unique identifier for the detected feature
	FeatureID string `json:"feature_id,required"`
	// The type of feature detected in the scan
	Type ChainAgnosticTransactionScanResponseValidationFeaturesType `json:"type,required"`
	// The blockchain address where the feature was detected
	Address string `json:"address"`
	// This field can have the runtime type of [[]string].
	AssociatedURLs interface{} `json:"associated_urls"`
	// Type of entity associated with the feature
	Entity ChainAgnosticTransactionScanResponseValidationFeaturesEntity `json:"entity"`
	// The URL where the feature was detected
	URL   string                                                    `json:"url"`
	JSON  chainAgnosticTransactionScanResponseValidationFeatureJSON `json:"-"`
	union ChainAgnosticTransactionScanResponseValidationFeaturesUnion
}

// chainAgnosticTransactionScanResponseValidationFeatureJSON contains the JSON
// metadata for the struct [ChainAgnosticTransactionScanResponseValidationFeature]
type chainAgnosticTransactionScanResponseValidationFeatureJSON struct {
	Description    apijson.Field
	FeatureID      apijson.Field
	Type           apijson.Field
	Address        apijson.Field
	AssociatedURLs apijson.Field
	Entity         apijson.Field
	URL            apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r chainAgnosticTransactionScanResponseValidationFeatureJSON) RawJSON() string {
	return r.raw
}

func (r *ChainAgnosticTransactionScanResponseValidationFeature) UnmarshalJSON(data []byte) (err error) {
	*r = ChainAgnosticTransactionScanResponseValidationFeature{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ChainAgnosticTransactionScanResponseValidationFeaturesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeature],
// [ChainAgnosticTransactionScanResponseValidationFeaturesURLFeature],
// [ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeature].
func (r ChainAgnosticTransactionScanResponseValidationFeature) AsUnion() ChainAgnosticTransactionScanResponseValidationFeaturesUnion {
	return r.union
}

// Union satisfied by
// [ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeature],
// [ChainAgnosticTransactionScanResponseValidationFeaturesURLFeature] or
// [ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeature].
type ChainAgnosticTransactionScanResponseValidationFeaturesUnion interface {
	implementsChainAgnosticTransactionScanResponseValidationFeature()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ChainAgnosticTransactionScanResponseValidationFeaturesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeature{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ChainAgnosticTransactionScanResponseValidationFeaturesURLFeature{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeature{}),
		},
	)
}

type ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeature struct {
	// The blockchain address where the feature was detected
	Address string `json:"address,required"`
	// List of URLs associated with this address
	AssociatedURLs []string `json:"associated_urls,required"`
	// Detailed description of the detected feature
	Description string `json:"description,required"`
	// Unique identifier for the detected feature
	FeatureID string `json:"feature_id,required"`
	// The type of feature detected in the scan
	Type ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureType `json:"type,required"`
	// Type of entity associated with the feature
	Entity ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureEntity `json:"entity"`
	JSON   chainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureJSON   `json:"-"`
}

// chainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureJSON
// contains the JSON metadata for the struct
// [ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeature]
type chainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureJSON struct {
	Address        apijson.Field
	AssociatedURLs apijson.Field
	Description    apijson.Field
	FeatureID      apijson.Field
	Type           apijson.Field
	Entity         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureJSON) RawJSON() string {
	return r.raw
}

func (r ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeature) implementsChainAgnosticTransactionScanResponseValidationFeature() {
}

// The type of feature detected in the scan
type ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureType string

const (
	ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureTypeMalicious ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureType = "Malicious"
	ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureTypeWarning   ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureType = "Warning"
	ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureTypeBenign    ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureType = "Benign"
	ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureTypeHighRisk  ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureType = "High-Risk"
	ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureTypeInfo      ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureType = "Info"
)

func (r ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureType) IsKnown() bool {
	switch r {
	case ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureTypeMalicious, ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureTypeWarning, ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureTypeBenign, ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureTypeHighRisk, ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureTypeInfo:
		return true
	}
	return false
}

// Type of entity associated with the feature
type ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureEntity string

const (
	ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureEntityAddress ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureEntity = "address"
)

func (r ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureEntity) IsKnown() bool {
	switch r {
	case ChainAgnosticTransactionScanResponseValidationFeaturesAddressFeatureEntityAddress:
		return true
	}
	return false
}

type ChainAgnosticTransactionScanResponseValidationFeaturesURLFeature struct {
	// Detailed description of the detected feature
	Description string `json:"description,required"`
	// Unique identifier for the detected feature
	FeatureID string `json:"feature_id,required"`
	// The type of feature detected in the scan
	Type ChainAgnosticTransactionScanResponseValidationFeaturesURLFeatureType `json:"type,required"`
	// The URL where the feature was detected
	URL string `json:"url,required"`
	// Type of entity associated with the feature
	Entity ChainAgnosticTransactionScanResponseValidationFeaturesURLFeatureEntity `json:"entity"`
	JSON   chainAgnosticTransactionScanResponseValidationFeaturesURLFeatureJSON   `json:"-"`
}

// chainAgnosticTransactionScanResponseValidationFeaturesURLFeatureJSON contains
// the JSON metadata for the struct
// [ChainAgnosticTransactionScanResponseValidationFeaturesURLFeature]
type chainAgnosticTransactionScanResponseValidationFeaturesURLFeatureJSON struct {
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	Entity      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChainAgnosticTransactionScanResponseValidationFeaturesURLFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chainAgnosticTransactionScanResponseValidationFeaturesURLFeatureJSON) RawJSON() string {
	return r.raw
}

func (r ChainAgnosticTransactionScanResponseValidationFeaturesURLFeature) implementsChainAgnosticTransactionScanResponseValidationFeature() {
}

// The type of feature detected in the scan
type ChainAgnosticTransactionScanResponseValidationFeaturesURLFeatureType string

const (
	ChainAgnosticTransactionScanResponseValidationFeaturesURLFeatureTypeMalicious ChainAgnosticTransactionScanResponseValidationFeaturesURLFeatureType = "Malicious"
	ChainAgnosticTransactionScanResponseValidationFeaturesURLFeatureTypeWarning   ChainAgnosticTransactionScanResponseValidationFeaturesURLFeatureType = "Warning"
	ChainAgnosticTransactionScanResponseValidationFeaturesURLFeatureTypeBenign    ChainAgnosticTransactionScanResponseValidationFeaturesURLFeatureType = "Benign"
	ChainAgnosticTransactionScanResponseValidationFeaturesURLFeatureTypeHighRisk  ChainAgnosticTransactionScanResponseValidationFeaturesURLFeatureType = "High-Risk"
	ChainAgnosticTransactionScanResponseValidationFeaturesURLFeatureTypeInfo      ChainAgnosticTransactionScanResponseValidationFeaturesURLFeatureType = "Info"
)

func (r ChainAgnosticTransactionScanResponseValidationFeaturesURLFeatureType) IsKnown() bool {
	switch r {
	case ChainAgnosticTransactionScanResponseValidationFeaturesURLFeatureTypeMalicious, ChainAgnosticTransactionScanResponseValidationFeaturesURLFeatureTypeWarning, ChainAgnosticTransactionScanResponseValidationFeaturesURLFeatureTypeBenign, ChainAgnosticTransactionScanResponseValidationFeaturesURLFeatureTypeHighRisk, ChainAgnosticTransactionScanResponseValidationFeaturesURLFeatureTypeInfo:
		return true
	}
	return false
}

// Type of entity associated with the feature
type ChainAgnosticTransactionScanResponseValidationFeaturesURLFeatureEntity string

const (
	ChainAgnosticTransactionScanResponseValidationFeaturesURLFeatureEntityURL ChainAgnosticTransactionScanResponseValidationFeaturesURLFeatureEntity = "url"
)

func (r ChainAgnosticTransactionScanResponseValidationFeaturesURLFeatureEntity) IsKnown() bool {
	switch r {
	case ChainAgnosticTransactionScanResponseValidationFeaturesURLFeatureEntityURL:
		return true
	}
	return false
}

type ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeature struct {
	// Detailed description of the detected feature
	Description string `json:"description,required"`
	// Unique identifier for the detected feature
	FeatureID string `json:"feature_id,required"`
	// The type of feature detected in the scan
	Type ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureType `json:"type,required"`
	// Type of entity associated with the feature
	Entity ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureEntity `json:"entity"`
	JSON   chainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureJSON   `json:"-"`
}

// chainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureJSON
// contains the JSON metadata for the struct
// [ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeature]
type chainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureJSON struct {
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	Entity      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureJSON) RawJSON() string {
	return r.raw
}

func (r ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeature) implementsChainAgnosticTransactionScanResponseValidationFeature() {
}

// The type of feature detected in the scan
type ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureType string

const (
	ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureTypeMalicious ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureType = "Malicious"
	ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureTypeWarning   ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureType = "Warning"
	ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureTypeBenign    ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureType = "Benign"
	ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureTypeHighRisk  ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureType = "High-Risk"
	ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureTypeInfo      ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureType = "Info"
)

func (r ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureType) IsKnown() bool {
	switch r {
	case ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureTypeMalicious, ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureTypeWarning, ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureTypeBenign, ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureTypeHighRisk, ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureTypeInfo:
		return true
	}
	return false
}

// Type of entity associated with the feature
type ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureEntity string

const (
	ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureEntityTransaction ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureEntity = "transaction"
)

func (r ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureEntity) IsKnown() bool {
	switch r {
	case ChainAgnosticTransactionScanResponseValidationFeaturesTransactionFeatureEntityTransaction:
		return true
	}
	return false
}

// The type of feature detected in the scan
type ChainAgnosticTransactionScanResponseValidationFeaturesType string

const (
	ChainAgnosticTransactionScanResponseValidationFeaturesTypeMalicious ChainAgnosticTransactionScanResponseValidationFeaturesType = "Malicious"
	ChainAgnosticTransactionScanResponseValidationFeaturesTypeWarning   ChainAgnosticTransactionScanResponseValidationFeaturesType = "Warning"
	ChainAgnosticTransactionScanResponseValidationFeaturesTypeBenign    ChainAgnosticTransactionScanResponseValidationFeaturesType = "Benign"
	ChainAgnosticTransactionScanResponseValidationFeaturesTypeHighRisk  ChainAgnosticTransactionScanResponseValidationFeaturesType = "High-Risk"
	ChainAgnosticTransactionScanResponseValidationFeaturesTypeInfo      ChainAgnosticTransactionScanResponseValidationFeaturesType = "Info"
)

func (r ChainAgnosticTransactionScanResponseValidationFeaturesType) IsKnown() bool {
	switch r {
	case ChainAgnosticTransactionScanResponseValidationFeaturesTypeMalicious, ChainAgnosticTransactionScanResponseValidationFeaturesTypeWarning, ChainAgnosticTransactionScanResponseValidationFeaturesTypeBenign, ChainAgnosticTransactionScanResponseValidationFeaturesTypeHighRisk, ChainAgnosticTransactionScanResponseValidationFeaturesTypeInfo:
		return true
	}
	return false
}

// Type of entity associated with the feature
type ChainAgnosticTransactionScanResponseValidationFeaturesEntity string

const (
	ChainAgnosticTransactionScanResponseValidationFeaturesEntityAddress     ChainAgnosticTransactionScanResponseValidationFeaturesEntity = "address"
	ChainAgnosticTransactionScanResponseValidationFeaturesEntityURL         ChainAgnosticTransactionScanResponseValidationFeaturesEntity = "url"
	ChainAgnosticTransactionScanResponseValidationFeaturesEntityTransaction ChainAgnosticTransactionScanResponseValidationFeaturesEntity = "transaction"
)

func (r ChainAgnosticTransactionScanResponseValidationFeaturesEntity) IsKnown() bool {
	switch r {
	case ChainAgnosticTransactionScanResponseValidationFeaturesEntityAddress, ChainAgnosticTransactionScanResponseValidationFeaturesEntityURL, ChainAgnosticTransactionScanResponseValidationFeaturesEntityTransaction:
		return true
	}
	return false
}

type ChainAgnosticTransactionScanParams struct {
	// Transaction data
	Data param.Field[ChainAgnosticTransactionScanParamsData] `json:"data,required"`
	// Additional metadata about the request including account and connection
	// information
	Metadata param.Field[ChainAgnosticTransactionScanParamsMetadata] `json:"metadata,required"`
	// List of options to apply during the transaction scan
	Options param.Field[[]ChainAgnosticTransactionScanParamsOption] `json:"options,required"`
}

func (r ChainAgnosticTransactionScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Transaction data
type ChainAgnosticTransactionScanParamsData struct {
	// Amount of the transaction in the specified asset
	Amount param.Field[float64] `json:"amount,required"`
	// Asset information
	Asset param.Field[ChainAgnosticTransactionScanParamsDataAssetUnion] `json:"asset,required"`
	// The chain name
	Chain param.Field[ChainAgnosticTransactionScanParamsDataChain] `json:"chain,required"`
	// Destination address or identifier for the transaction
	To param.Field[string] `json:"to,required"`
	// Source address or identifier for the transaction
	From param.Field[string] `json:"from"`
}

func (r ChainAgnosticTransactionScanParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Asset information
type ChainAgnosticTransactionScanParamsDataAsset struct {
	// The address of the asset
	Address param.Field[string] `json:"address"`
	// The symbol of the asset
	Symbol param.Field[string] `json:"symbol"`
}

func (r ChainAgnosticTransactionScanParamsDataAsset) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChainAgnosticTransactionScanParamsDataAsset) implementsChainAgnosticTransactionScanParamsDataAssetUnion() {
}

// Asset information
//
// Satisfied by [ChainAgnosticTransactionScanParamsDataAssetAssetAddress],
// [ChainAgnosticTransactionScanParamsDataAssetAssetSymbol],
// [ChainAgnosticTransactionScanParamsDataAsset].
type ChainAgnosticTransactionScanParamsDataAssetUnion interface {
	implementsChainAgnosticTransactionScanParamsDataAssetUnion()
}

type ChainAgnosticTransactionScanParamsDataAssetAssetAddress struct {
	// The address of the asset
	Address param.Field[string] `json:"address,required"`
}

func (r ChainAgnosticTransactionScanParamsDataAssetAssetAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChainAgnosticTransactionScanParamsDataAssetAssetAddress) implementsChainAgnosticTransactionScanParamsDataAssetUnion() {
}

type ChainAgnosticTransactionScanParamsDataAssetAssetSymbol struct {
	// The symbol of the asset
	Symbol param.Field[string] `json:"symbol,required"`
}

func (r ChainAgnosticTransactionScanParamsDataAssetAssetSymbol) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChainAgnosticTransactionScanParamsDataAssetAssetSymbol) implementsChainAgnosticTransactionScanParamsDataAssetUnion() {
}

// The chain name
type ChainAgnosticTransactionScanParamsDataChain string

const (
	ChainAgnosticTransactionScanParamsDataChainEthereum ChainAgnosticTransactionScanParamsDataChain = "ethereum"
	ChainAgnosticTransactionScanParamsDataChainBase     ChainAgnosticTransactionScanParamsDataChain = "base"
	ChainAgnosticTransactionScanParamsDataChainArbitrum ChainAgnosticTransactionScanParamsDataChain = "arbitrum"
	ChainAgnosticTransactionScanParamsDataChainOptimism ChainAgnosticTransactionScanParamsDataChain = "optimism"
	ChainAgnosticTransactionScanParamsDataChainPolygon  ChainAgnosticTransactionScanParamsDataChain = "polygon"
	ChainAgnosticTransactionScanParamsDataChainSolana   ChainAgnosticTransactionScanParamsDataChain = "solana"
	ChainAgnosticTransactionScanParamsDataChainStellar  ChainAgnosticTransactionScanParamsDataChain = "stellar"
	ChainAgnosticTransactionScanParamsDataChainBitcoin  ChainAgnosticTransactionScanParamsDataChain = "bitcoin"
)

func (r ChainAgnosticTransactionScanParamsDataChain) IsKnown() bool {
	switch r {
	case ChainAgnosticTransactionScanParamsDataChainEthereum, ChainAgnosticTransactionScanParamsDataChainBase, ChainAgnosticTransactionScanParamsDataChainArbitrum, ChainAgnosticTransactionScanParamsDataChainOptimism, ChainAgnosticTransactionScanParamsDataChainPolygon, ChainAgnosticTransactionScanParamsDataChainSolana, ChainAgnosticTransactionScanParamsDataChainStellar, ChainAgnosticTransactionScanParamsDataChainBitcoin:
		return true
	}
	return false
}

// Additional metadata about the request including account and connection
// information
type ChainAgnosticTransactionScanParamsMetadata struct {
	// Account information associated with the request
	Account param.Field[ChainAgnosticTransactionScanParamsMetadataAccount] `json:"account"`
	// Connection metadata including user agent and IP information
	Connection param.Field[ChainAgnosticTransactionScanParamsMetadataConnection] `json:"connection"`
}

func (r ChainAgnosticTransactionScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Account information associated with the request
type ChainAgnosticTransactionScanParamsMetadataAccount struct {
	// Unique identifier for the account
	AccountID param.Field[string] `json:"account_id,required"`
	// Timestamp when the account was created
	AccountCreationTimestamp param.Field[time.Time] `json:"account_creation_timestamp" format:"date-time"`
	// Age of the user in years
	UserAge param.Field[int64] `json:"user_age"`
	// ISO country code of the user's location
	UserCountryCode param.Field[string] `json:"user_country_code"`
}

func (r ChainAgnosticTransactionScanParamsMetadataAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Connection metadata including user agent and IP information
type ChainAgnosticTransactionScanParamsMetadataConnection struct {
	// IP address of the customer making the request
	IPAddress param.Field[string] `json:"ip_address,required" format:"ipvanyaddress"`
	// User agent string from the client's browser or application
	UserAgent param.Field[string] `json:"user_agent"`
}

func (r ChainAgnosticTransactionScanParamsMetadataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An enumeration.
type ChainAgnosticTransactionScanParamsOption string

const (
	ChainAgnosticTransactionScanParamsOptionValidation ChainAgnosticTransactionScanParamsOption = "validation"
)

func (r ChainAgnosticTransactionScanParamsOption) IsKnown() bool {
	switch r {
	case ChainAgnosticTransactionScanParamsOptionValidation:
		return true
	}
	return false
}
