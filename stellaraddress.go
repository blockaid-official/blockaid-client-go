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

// StellarAddressService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStellarAddressService] method instead.
type StellarAddressService struct {
	Options []option.RequestOption
}

// NewStellarAddressService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStellarAddressService(opts ...option.RequestOption) (r *StellarAddressService) {
	r = &StellarAddressService{}
	r.Options = opts
	return
}

// Gets an address and returns a full security assessment indicating weather or not
// this address is malicious as well as textual reasons of why the address was
// flagged that way.
func (r *StellarAddressService) Scan(ctx context.Context, body StellarAddressScanParams, opts ...option.RequestOption) (res *StellarAddressScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/stellar/address/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type StellarAddressScanResponse struct {
	Address    string                              `json:"address,required"`
	Chain      string                              `json:"chain,required"`
	Features   []StellarAddressScanResponseFeature `json:"features,required"`
	ResultType string                              `json:"result_type,required"`
	JSON       stellarAddressScanResponseJSON      `json:"-"`
}

// stellarAddressScanResponseJSON contains the JSON metadata for the struct
// [StellarAddressScanResponse]
type stellarAddressScanResponseJSON struct {
	Address     apijson.Field
	Chain       apijson.Field
	Features    apijson.Field
	ResultType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarAddressScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarAddressScanResponseJSON) RawJSON() string {
	return r.raw
}

type StellarAddressScanResponseFeature struct {
	Address     string                                 `json:"address,required"`
	Description string                                 `json:"description,required"`
	FeatureID   string                                 `json:"feature_id,required"`
	Type        StellarAddressScanResponseFeaturesType `json:"type,required"`
	JSON        stellarAddressScanResponseFeatureJSON  `json:"-"`
}

// stellarAddressScanResponseFeatureJSON contains the JSON metadata for the struct
// [StellarAddressScanResponseFeature]
type stellarAddressScanResponseFeatureJSON struct {
	Address     apijson.Field
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarAddressScanResponseFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarAddressScanResponseFeatureJSON) RawJSON() string {
	return r.raw
}

type StellarAddressScanResponseFeaturesType string

const (
	StellarAddressScanResponseFeaturesTypeBenign    StellarAddressScanResponseFeaturesType = "Benign"
	StellarAddressScanResponseFeaturesTypeWarning   StellarAddressScanResponseFeaturesType = "Warning"
	StellarAddressScanResponseFeaturesTypeMalicious StellarAddressScanResponseFeaturesType = "Malicious"
	StellarAddressScanResponseFeaturesTypeInfo      StellarAddressScanResponseFeaturesType = "Info"
)

func (r StellarAddressScanResponseFeaturesType) IsKnown() bool {
	switch r {
	case StellarAddressScanResponseFeaturesTypeBenign, StellarAddressScanResponseFeaturesTypeWarning, StellarAddressScanResponseFeaturesTypeMalicious, StellarAddressScanResponseFeaturesTypeInfo:
		return true
	}
	return false
}

type StellarAddressScanParams struct {
	Address param.Field[string] `json:"address,required"`
}

func (r StellarAddressScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
