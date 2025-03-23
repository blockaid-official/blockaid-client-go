// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/param"
	"github.com/blockaid-official/blockaid-client-go/option"
)

// SolanaService contains methods and other services that help with interacting
// with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSolanaService] method instead.
type SolanaService struct {
	Options []option.RequestOption
	Message *SolanaMessageService
	Address *SolanaAddressService
}

// NewSolanaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSolanaService(opts ...option.RequestOption) (r *SolanaService) {
	r = &SolanaService{}
	r.Options = opts
	r.Message = NewSolanaMessageService(opts...)
	r.Address = NewSolanaAddressService(opts...)
	return
}

type AddressScanRequestSchemaParam struct {
	// Encoded public key
	Address  param.Field[string]                                `json:"address,required"`
	Metadata param.Field[AddressScanRequestSchemaMetadataParam] `json:"metadata,required"`
	// Chain to scan the transaction on
	Chain param.Field[string] `json:"chain"`
}

func (r AddressScanRequestSchemaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressScanRequestSchemaMetadataParam struct {
	// URL of the dApp related to the address
	URL param.Field[string] `json:"url,required"`
}

func (r AddressScanRequestSchemaMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressScanResponseSchema struct {
	// Features about the result
	Features []AddressScanResponseSchemaFeature `json:"features,required"`
	// Verdict of Result
	ResultType AddressScanResponseSchemaResultType `json:"result_type,required"`
	JSON       addressScanResponseSchemaJSON       `json:"-"`
}

// addressScanResponseSchemaJSON contains the JSON metadata for the struct
// [AddressScanResponseSchema]
type addressScanResponseSchemaJSON struct {
	Features    apijson.Field
	ResultType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressScanResponseSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressScanResponseSchemaJSON) RawJSON() string {
	return r.raw
}

type AddressScanResponseSchemaFeature struct {
	// Description of the feature
	Description string `json:"description,required"`
	// ID of the feature
	FeatureID string `json:"feature_id,required"`
	// An enumeration.
	Type AddressScanResponseSchemaFeaturesType `json:"type,required"`
	JSON addressScanResponseSchemaFeatureJSON  `json:"-"`
}

// addressScanResponseSchemaFeatureJSON contains the JSON metadata for the struct
// [AddressScanResponseSchemaFeature]
type addressScanResponseSchemaFeatureJSON struct {
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressScanResponseSchemaFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressScanResponseSchemaFeatureJSON) RawJSON() string {
	return r.raw
}

// An enumeration.
type AddressScanResponseSchemaFeaturesType string

const (
	AddressScanResponseSchemaFeaturesTypeMalicious AddressScanResponseSchemaFeaturesType = "Malicious"
	AddressScanResponseSchemaFeaturesTypeWarning   AddressScanResponseSchemaFeaturesType = "Warning"
	AddressScanResponseSchemaFeaturesTypeBenign    AddressScanResponseSchemaFeaturesType = "Benign"
	AddressScanResponseSchemaFeaturesTypeInfo      AddressScanResponseSchemaFeaturesType = "Info"
)

func (r AddressScanResponseSchemaFeaturesType) IsKnown() bool {
	switch r {
	case AddressScanResponseSchemaFeaturesTypeMalicious, AddressScanResponseSchemaFeaturesTypeWarning, AddressScanResponseSchemaFeaturesTypeBenign, AddressScanResponseSchemaFeaturesTypeInfo:
		return true
	}
	return false
}

// Verdict of Result
type AddressScanResponseSchemaResultType string

const (
	AddressScanResponseSchemaResultTypeBenign    AddressScanResponseSchemaResultType = "Benign"
	AddressScanResponseSchemaResultTypeWarning   AddressScanResponseSchemaResultType = "Warning"
	AddressScanResponseSchemaResultTypeMalicious AddressScanResponseSchemaResultType = "Malicious"
	AddressScanResponseSchemaResultTypeSpam      AddressScanResponseSchemaResultType = "Spam"
)

func (r AddressScanResponseSchemaResultType) IsKnown() bool {
	switch r {
	case AddressScanResponseSchemaResultTypeBenign, AddressScanResponseSchemaResultTypeWarning, AddressScanResponseSchemaResultTypeMalicious, AddressScanResponseSchemaResultTypeSpam:
		return true
	}
	return false
}
