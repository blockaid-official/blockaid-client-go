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

// OrganizationRiskExposureConfigurationService contains methods and other services
// that help with interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationRiskExposureConfigurationService] method instead.
type OrganizationRiskExposureConfigurationService struct {
	Options []option.RequestOption
}

// NewOrganizationRiskExposureConfigurationService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewOrganizationRiskExposureConfigurationService(opts ...option.RequestOption) (r *OrganizationRiskExposureConfigurationService) {
	r = &OrganizationRiskExposureConfigurationService{}
	r.Options = opts
	return
}

// Returns your organization's current risk exposure settings, including category
// risk scores and verdict thresholds.
func (r *OrganizationRiskExposureConfigurationService) Get(ctx context.Context, opts ...option.RequestOption) (res *OrganizationRiskConfigView, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/platform/organization/configuration/risk-exposure"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update your organization's risk exposure settings. Pass a category name or
// verdict level with a new value to override the platform default, or pass null to
// revert to the default. Any key you omit stays unchanged.
func (r *OrganizationRiskExposureConfigurationService) Update(ctx context.Context, body OrganizationRiskExposureConfigurationUpdateParams, opts ...option.RequestOption) (res *OrganizationRiskConfigView, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/platform/organization/configuration/risk-exposure"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// A risk-exposure category with its platform default score and the organization's
// override (null when unset).
type OrganizationCategoryRiskView struct {
	// The exposure category name, such as sanctioned_entity, stolen_funds, or mixer.
	Category string `json:"category" api:"required"`
	// Scores at or above this value contribute to this category's risk. Platform
	// default, set by Blockaid.
	DefaultRiskScore float64 `json:"default_risk_score" api:"required"`
	// Your organization's custom risk score for this category. Overrides the platform
	// default when set, or null to use the default.
	OverrideRiskScore float64                          `json:"override_risk_score" api:"required,nullable"`
	JSON              organizationCategoryRiskViewJSON `json:"-"`
}

// organizationCategoryRiskViewJSON contains the JSON metadata for the struct
// [OrganizationCategoryRiskView]
type organizationCategoryRiskViewJSON struct {
	Category          apijson.Field
	DefaultRiskScore  apijson.Field
	OverrideRiskScore apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *OrganizationCategoryRiskView) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationCategoryRiskViewJSON) RawJSON() string {
	return r.raw
}

// Your organization's risk exposure settings, including category risk scores and
// verdict thresholds.
type OrganizationRiskConfigView struct {
	// The risk score configuration for each exposure category.
	Categories []OrganizationCategoryRiskView `json:"categories" api:"required"`
	// The score thresholds that determine each verdict level.
	Verdicts []VerdictThresholdView         `json:"verdicts" api:"required"`
	JSON     organizationRiskConfigViewJSON `json:"-"`
}

// organizationRiskConfigViewJSON contains the JSON metadata for the struct
// [OrganizationRiskConfigView]
type organizationRiskConfigViewJSON struct {
	Categories  apijson.Field
	Verdicts    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRiskConfigView) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationRiskConfigViewJSON) RawJSON() string {
	return r.raw
}

// A partial update to your organization's risk exposure settings. Only the fields
// you include are changed.
type PatchRiskConfigRequestParam struct {
	// Risk score overrides by category name. Pass a value in [0, 1] to set an
	// override, or null to revert to the platform default.
	CategoryOverrides param.Field[map[string]float64] `json:"category_overrides"`
	// Verdict threshold overrides by level. Pass a value in [0, 1] to set an override,
	// or null to revert to the platform default. The resulting thresholds must be
	// strictly increasing (warning < high_risk < malicious).
	VerdictOverrides param.Field[PatchRiskConfigRequestVerdictOverridesParam] `json:"verdict_overrides"`
}

func (r PatchRiskConfigRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Verdict threshold overrides by level. Pass a value in [0, 1] to set an override,
// or null to revert to the platform default. The resulting thresholds must be
// strictly increasing (warning < high_risk < malicious).
type PatchRiskConfigRequestVerdictOverridesParam struct {
	// Scores at or above this value are classified as high_risk. Set to null to revert
	// to the platform default.
	HighRisk param.Field[float64] `json:"high_risk"`
	// Scores at or above this value are classified as malicious. Set to null to revert
	// to the platform default.
	Malicious param.Field[float64] `json:"malicious"`
	// Scores at or above this value are classified as warning. Set to null to revert
	// to the platform default.
	Warning param.Field[float64] `json:"warning"`
}

func (r PatchRiskConfigRequestVerdictOverridesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A verdict level with its platform default threshold and the organization's
// override (null when unset).
type VerdictThresholdView struct {
	// The minimum aggregate score at which this verdict is applied, as set by
	// Blockaid.
	DefaultThreshold float64 `json:"default_threshold" api:"required"`
	// The verdict level this threshold applies to. benign is excluded as it is the
	// implicit default below warning.
	Level VerdictThresholdViewLevel `json:"level" api:"required"`
	// Your organization's custom minimum score for this verdict. Overrides the
	// platform default when set, or null to use the default.
	OverrideThreshold float64                  `json:"override_threshold" api:"required,nullable"`
	JSON              verdictThresholdViewJSON `json:"-"`
}

// verdictThresholdViewJSON contains the JSON metadata for the struct
// [VerdictThresholdView]
type verdictThresholdViewJSON struct {
	DefaultThreshold  apijson.Field
	Level             apijson.Field
	OverrideThreshold apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *VerdictThresholdView) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verdictThresholdViewJSON) RawJSON() string {
	return r.raw
}

// The verdict level this threshold applies to. benign is excluded as it is the
// implicit default below warning.
type VerdictThresholdViewLevel string

const (
	VerdictThresholdViewLevelWarning   VerdictThresholdViewLevel = "warning"
	VerdictThresholdViewLevelHighRisk  VerdictThresholdViewLevel = "high_risk"
	VerdictThresholdViewLevelMalicious VerdictThresholdViewLevel = "malicious"
)

func (r VerdictThresholdViewLevel) IsKnown() bool {
	switch r {
	case VerdictThresholdViewLevelWarning, VerdictThresholdViewLevelHighRisk, VerdictThresholdViewLevelMalicious:
		return true
	}
	return false
}

type OrganizationRiskExposureConfigurationUpdateParams struct {
	// A partial update to your organization's risk exposure settings. Only the fields
	// you include are changed.
	PatchRiskConfigRequest PatchRiskConfigRequestParam `json:"PatchRiskConfigRequest" api:"required"`
}

func (r OrganizationRiskExposureConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PatchRiskConfigRequest)
}
