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

// SolanaAddressService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSolanaAddressService] method instead.
type SolanaAddressService struct {
	Options []option.RequestOption
}

// NewSolanaAddressService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSolanaAddressService(opts ...option.RequestOption) (r *SolanaAddressService) {
	r = &SolanaAddressService{}
	r.Options = opts
	return
}

// Gets an address and returns a full security assessment indicating weather or not
// this address is malicious as well as textual reasons of why the address was
// flagged that way.
func (r *SolanaAddressService) Scan(ctx context.Context, body SolanaAddressScanParams, opts ...option.RequestOption) (res *SolanaAddressScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/solana/address/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SolanaAddressScanResponse struct {
	// Features about the result
	Features []SolanaAddressScanResponseFeature `json:"features,required"`
	// An enumeration.
	ResultType SolanaAddressScanResponseResultType `json:"result_type,required"`
	JSON       solanaAddressScanResponseJSON       `json:"-"`
}

// solanaAddressScanResponseJSON contains the JSON metadata for the struct
// [SolanaAddressScanResponse]
type solanaAddressScanResponseJSON struct {
	Features    apijson.Field
	ResultType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaAddressScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaAddressScanResponseJSON) RawJSON() string {
	return r.raw
}

type SolanaAddressScanResponseFeature struct {
	// Description of the feature
	Description string `json:"description,required"`
	// ID of the feature
	FeatureID string `json:"feature_id,required"`
	// An enumeration.
	Type SolanaAddressScanResponseFeaturesType `json:"type,required"`
	JSON solanaAddressScanResponseFeatureJSON  `json:"-"`
}

// solanaAddressScanResponseFeatureJSON contains the JSON metadata for the struct
// [SolanaAddressScanResponseFeature]
type solanaAddressScanResponseFeatureJSON struct {
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SolanaAddressScanResponseFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaAddressScanResponseFeatureJSON) RawJSON() string {
	return r.raw
}

// An enumeration.
type SolanaAddressScanResponseFeaturesType string

const (
	SolanaAddressScanResponseFeaturesTypeMalicious SolanaAddressScanResponseFeaturesType = "Malicious"
	SolanaAddressScanResponseFeaturesTypeWarning   SolanaAddressScanResponseFeaturesType = "Warning"
	SolanaAddressScanResponseFeaturesTypeBenign    SolanaAddressScanResponseFeaturesType = "Benign"
	SolanaAddressScanResponseFeaturesTypeInfo      SolanaAddressScanResponseFeaturesType = "Info"
)

func (r SolanaAddressScanResponseFeaturesType) IsKnown() bool {
	switch r {
	case SolanaAddressScanResponseFeaturesTypeMalicious, SolanaAddressScanResponseFeaturesTypeWarning, SolanaAddressScanResponseFeaturesTypeBenign, SolanaAddressScanResponseFeaturesTypeInfo:
		return true
	}
	return false
}

// An enumeration.
type SolanaAddressScanResponseResultType string

const (
	SolanaAddressScanResponseResultTypeMalicious SolanaAddressScanResponseResultType = "Malicious"
	SolanaAddressScanResponseResultTypeWarning   SolanaAddressScanResponseResultType = "Warning"
	SolanaAddressScanResponseResultTypeBenign    SolanaAddressScanResponseResultType = "Benign"
)

func (r SolanaAddressScanResponseResultType) IsKnown() bool {
	switch r {
	case SolanaAddressScanResponseResultTypeMalicious, SolanaAddressScanResponseResultTypeWarning, SolanaAddressScanResponseResultTypeBenign:
		return true
	}
	return false
}

type SolanaAddressScanParams struct {
	// Encoded public key
	Address  param.Field[string]                          `json:"address,required"`
	Metadata param.Field[SolanaAddressScanParamsMetadata] `json:"metadata,required"`
	// Chain to scan the transaction on
	Chain param.Field[string] `json:"chain"`
}

func (r SolanaAddressScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SolanaAddressScanParamsMetadata struct {
	// URL of the dApp related to the address
	URL param.Field[string] `json:"url,required"`
}

func (r SolanaAddressScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
