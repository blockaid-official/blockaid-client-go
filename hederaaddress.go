// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/param"
	"github.com/blockaid-official/blockaid-client-go/internal/requestconfig"
	"github.com/blockaid-official/blockaid-client-go/option"
)

// HederaAddressService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHederaAddressService] method instead.
type HederaAddressService struct {
	Options []option.RequestOption
}

// NewHederaAddressService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHederaAddressService(opts ...option.RequestOption) (r *HederaAddressService) {
	r = &HederaAddressService{}
	r.Options = opts
	return
}

// Scan Address
func (r *HederaAddressService) Scan(ctx context.Context, body HederaAddressScanParams, opts ...option.RequestOption) (res *HederaAddressScanResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/hedera/address/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type HederaAddressScanResponse struct {
	// The outcome of the address scan, indicating the assessed risk level of the
	// scanned address. <br /> Possible values: <br /> - `Malicious` — The address is
	// confirmed to be associated with malicious activity. <br /> - `High-Risk` — The
	// address shows strong indicators of high-risk or potentially malicious behavior.
	// <br /> - `Warning` — The address exhibits suspicious signals that may require
	// caution. <br /> - `Benign` — No known malicious or risky activity was detected
	// for the address. <br /> - `Error` — The scan could not be completed
	// successfully.
	ResultType HederaAddressScanResponseResultType `json:"result_type" api:"required"`
	// A list of human-readable textual features describing the scanned address,
	// intended for display to the end user.
	Features []HederaAddressScanResponseFeature `json:"features"`
	JSON     hederaAddressScanResponseJSON      `json:"-"`
}

// hederaAddressScanResponseJSON contains the JSON metadata for the struct
// [HederaAddressScanResponse]
type hederaAddressScanResponseJSON struct {
	ResultType  apijson.Field
	Features    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaAddressScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaAddressScanResponseJSON) RawJSON() string {
	return r.raw
}

// The outcome of the address scan, indicating the assessed risk level of the
// scanned address. <br /> Possible values: <br /> - `Malicious` — The address is
// confirmed to be associated with malicious activity. <br /> - `High-Risk` — The
// address shows strong indicators of high-risk or potentially malicious behavior.
// <br /> - `Warning` — The address exhibits suspicious signals that may require
// caution. <br /> - `Benign` — No known malicious or risky activity was detected
// for the address. <br /> - `Error` — The scan could not be completed
// successfully.
type HederaAddressScanResponseResultType string

const (
	HederaAddressScanResponseResultTypeMalicious HederaAddressScanResponseResultType = "Malicious"
	HederaAddressScanResponseResultTypeHighRisk  HederaAddressScanResponseResultType = "High-Risk"
	HederaAddressScanResponseResultTypeWarning   HederaAddressScanResponseResultType = "Warning"
	HederaAddressScanResponseResultTypeBenign    HederaAddressScanResponseResultType = "Benign"
	HederaAddressScanResponseResultTypeError     HederaAddressScanResponseResultType = "Error"
)

func (r HederaAddressScanResponseResultType) IsKnown() bool {
	switch r {
	case HederaAddressScanResponseResultTypeMalicious, HederaAddressScanResponseResultTypeHighRisk, HederaAddressScanResponseResultTypeWarning, HederaAddressScanResponseResultTypeBenign, HederaAddressScanResponseResultTypeError:
		return true
	}
	return false
}

type HederaAddressScanResponseFeature struct {
	// A human-readable description explaining the feature.
	Description string `json:"description" api:"required"`
	// The unique identifier for the feature.
	FeatureID string `json:"feature_id" api:"required"`
	// The risk classification indicated by the feature. Possible values are
	// `Malicious`, `High-Risk`, `Warning`, `Benign`, or `Info`.
	Type HederaAddressScanResponseFeaturesType `json:"type" api:"required"`
	JSON hederaAddressScanResponseFeatureJSON  `json:"-"`
}

// hederaAddressScanResponseFeatureJSON contains the JSON metadata for the struct
// [HederaAddressScanResponseFeature]
type hederaAddressScanResponseFeatureJSON struct {
	Description apijson.Field
	FeatureID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HederaAddressScanResponseFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hederaAddressScanResponseFeatureJSON) RawJSON() string {
	return r.raw
}

// The risk classification indicated by the feature. Possible values are
// `Malicious`, `High-Risk`, `Warning`, `Benign`, or `Info`.
type HederaAddressScanResponseFeaturesType string

const (
	HederaAddressScanResponseFeaturesTypeMalicious HederaAddressScanResponseFeaturesType = "Malicious"
	HederaAddressScanResponseFeaturesTypeWarning   HederaAddressScanResponseFeaturesType = "Warning"
	HederaAddressScanResponseFeaturesTypeBenign    HederaAddressScanResponseFeaturesType = "Benign"
	HederaAddressScanResponseFeaturesTypeHighRisk  HederaAddressScanResponseFeaturesType = "High-Risk"
	HederaAddressScanResponseFeaturesTypeInfo      HederaAddressScanResponseFeaturesType = "Info"
)

func (r HederaAddressScanResponseFeaturesType) IsKnown() bool {
	switch r {
	case HederaAddressScanResponseFeaturesTypeMalicious, HederaAddressScanResponseFeaturesTypeWarning, HederaAddressScanResponseFeaturesTypeBenign, HederaAddressScanResponseFeaturesTypeHighRisk, HederaAddressScanResponseFeaturesTypeInfo:
		return true
	}
	return false
}

type HederaAddressScanParams struct {
	// The address to validate.
	Address param.Field[string] `json:"address" api:"required"`
	// The chain the transaction runs on.
	Chain param.Field[HederaAddressScanParamsChain] `json:"chain" api:"required"`
}

func (r HederaAddressScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The chain the transaction runs on.
type HederaAddressScanParamsChain string

const (
	HederaAddressScanParamsChainMainnet HederaAddressScanParamsChain = "mainnet"
)

func (r HederaAddressScanParamsChain) IsKnown() bool {
	switch r {
	case HederaAddressScanParamsChainMainnet:
		return true
	}
	return false
}
