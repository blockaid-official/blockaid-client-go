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

// Scan an address
func (r *SolanaAddressService) Scan(ctx context.Context, body SolanaAddressScanParams, opts ...option.RequestOption) (res *SolanaAddressScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/solana/address/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SolanaAddressScanResponse struct {
	// Features about the result
	Features []SolanaAddressScanResponseFeature `json:"features,required"`
	// Verdict of Result
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
	// Address the feature refers to
	Address string `json:"address,required,nullable"`
	// Textual description
	Description string `json:"description,required"`
	FeatureID   string `json:"feature_id,required"`
	// Feature Classification
	Type SolanaAddressScanResponseFeaturesType `json:"type,required"`
	JSON solanaAddressScanResponseFeatureJSON  `json:"-"`
}

// solanaAddressScanResponseFeatureJSON contains the JSON metadata for the struct
// [SolanaAddressScanResponseFeature]
type solanaAddressScanResponseFeatureJSON struct {
	Address     apijson.Field
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

// Feature Classification
type SolanaAddressScanResponseFeaturesType string

const (
	SolanaAddressScanResponseFeaturesTypeBenign    SolanaAddressScanResponseFeaturesType = "Benign"
	SolanaAddressScanResponseFeaturesTypeWarning   SolanaAddressScanResponseFeaturesType = "Warning"
	SolanaAddressScanResponseFeaturesTypeMalicious SolanaAddressScanResponseFeaturesType = "Malicious"
	SolanaAddressScanResponseFeaturesTypeInfo      SolanaAddressScanResponseFeaturesType = "Info"
)

func (r SolanaAddressScanResponseFeaturesType) IsKnown() bool {
	switch r {
	case SolanaAddressScanResponseFeaturesTypeBenign, SolanaAddressScanResponseFeaturesTypeWarning, SolanaAddressScanResponseFeaturesTypeMalicious, SolanaAddressScanResponseFeaturesTypeInfo:
		return true
	}
	return false
}

// Verdict of Result
type SolanaAddressScanResponseResultType string

const (
	SolanaAddressScanResponseResultTypeBenign    SolanaAddressScanResponseResultType = "Benign"
	SolanaAddressScanResponseResultTypeWarning   SolanaAddressScanResponseResultType = "Warning"
	SolanaAddressScanResponseResultTypeMalicious SolanaAddressScanResponseResultType = "Malicious"
)

func (r SolanaAddressScanResponseResultType) IsKnown() bool {
	switch r {
	case SolanaAddressScanResponseResultTypeBenign, SolanaAddressScanResponseResultTypeWarning, SolanaAddressScanResponseResultTypeMalicious:
		return true
	}
	return false
}

type SolanaAddressScanParams struct {
	Address  param.Field[string]                          `json:"address,required"`
	Metadata param.Field[SolanaAddressScanParamsMetadata] `json:"metadata,required"`
	Chain    param.Field[string]                          `json:"chain"`
}

func (r SolanaAddressScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SolanaAddressScanParamsMetadata struct {
	// URL of the dApp that originated the transaction
	URL param.Field[string] `json:"url"`
}

func (r SolanaAddressScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
