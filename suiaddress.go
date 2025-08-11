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

// SuiAddressService contains methods and other services that help with interacting
// with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSuiAddressService] method instead.
type SuiAddressService struct {
	Options []option.RequestOption
}

// NewSuiAddressService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSuiAddressService(opts ...option.RequestOption) (r *SuiAddressService) {
	r = &SuiAddressService{}
	r.Options = opts
	return
}

// Gets an address and returns a full security assessment indicating weather or not
// this address is malicious as well as textual reasons of why the address was
// flagged that way.
func (r *SuiAddressService) Scan(ctx context.Context, body SuiAddressScanParams, opts ...option.RequestOption) (res *SuiAddressScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/sui/address/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SuiAddressScanResponse struct {
	// Verdict of the validation
	ResultType SuiAddressScanResponseResultType `json:"result_type,required"`
	// A list of textual features about this transaction that can be presented to the
	// user.
	Features []SuiAddressScanResponseFeature `json:"features"`
	JSON     suiAddressScanResponseJSON      `json:"-"`
}

// suiAddressScanResponseJSON contains the JSON metadata for the struct
// [SuiAddressScanResponse]
type suiAddressScanResponseJSON struct {
	ResultType  apijson.Field
	Features    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiAddressScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiAddressScanResponseJSON) RawJSON() string {
	return r.raw
}

// Verdict of the validation
type SuiAddressScanResponseResultType string

const (
	SuiAddressScanResponseResultTypeBenign    SuiAddressScanResponseResultType = "Benign"
	SuiAddressScanResponseResultTypeSpam      SuiAddressScanResponseResultType = "Spam"
	SuiAddressScanResponseResultTypeWarning   SuiAddressScanResponseResultType = "Warning"
	SuiAddressScanResponseResultTypeMalicious SuiAddressScanResponseResultType = "Malicious"
	SuiAddressScanResponseResultTypeError     SuiAddressScanResponseResultType = "Error"
)

func (r SuiAddressScanResponseResultType) IsKnown() bool {
	switch r {
	case SuiAddressScanResponseResultTypeBenign, SuiAddressScanResponseResultTypeSpam, SuiAddressScanResponseResultTypeWarning, SuiAddressScanResponseResultTypeMalicious, SuiAddressScanResponseResultTypeError:
		return true
	}
	return false
}

type SuiAddressScanResponseFeature struct {
	Description string                             `json:"description,required"`
	FeatureID   string                             `json:"feature_id,required"`
	Type        SuiAddressScanResponseFeaturesType `json:"type,required"`
	JSON        suiAddressScanResponseFeatureJSON  `json:"-"`
}

// suiAddressScanResponseFeatureJSON contains the JSON metadata for the struct
// [SuiAddressScanResponseFeature]
type suiAddressScanResponseFeatureJSON struct {
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuiAddressScanResponseFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suiAddressScanResponseFeatureJSON) RawJSON() string {
	return r.raw
}

type SuiAddressScanResponseFeaturesType string

const (
	SuiAddressScanResponseFeaturesTypeBenign    SuiAddressScanResponseFeaturesType = "Benign"
	SuiAddressScanResponseFeaturesTypeWarning   SuiAddressScanResponseFeaturesType = "Warning"
	SuiAddressScanResponseFeaturesTypeMalicious SuiAddressScanResponseFeaturesType = "Malicious"
	SuiAddressScanResponseFeaturesTypeInfo      SuiAddressScanResponseFeaturesType = "Info"
)

func (r SuiAddressScanResponseFeaturesType) IsKnown() bool {
	switch r {
	case SuiAddressScanResponseFeaturesTypeBenign, SuiAddressScanResponseFeaturesTypeWarning, SuiAddressScanResponseFeaturesTypeMalicious, SuiAddressScanResponseFeaturesTypeInfo:
		return true
	}
	return false
}

type SuiAddressScanParams struct {
	Address param.Field[string]                    `json:"address,required"`
	Chain   param.Field[SuiAddressScanParamsChain] `json:"chain,required"`
}

func (r SuiAddressScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuiAddressScanParamsChain string

const (
	SuiAddressScanParamsChainMainnet SuiAddressScanParamsChain = "mainnet"
	SuiAddressScanParamsChainTestnet SuiAddressScanParamsChain = "testnet"
	SuiAddressScanParamsChainDevnet  SuiAddressScanParamsChain = "devnet"
)

func (r SuiAddressScanParamsChain) IsKnown() bool {
	switch r {
	case SuiAddressScanParamsChainMainnet, SuiAddressScanParamsChainTestnet, SuiAddressScanParamsChainDevnet:
		return true
	}
	return false
}
