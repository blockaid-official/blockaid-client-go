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

// SiteService contains methods and other services that help with interacting with
// the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSiteService] method instead.
type SiteService struct {
	Options []option.RequestOption
}

// NewSiteService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSiteService(opts ...option.RequestOption) (r *SiteService) {
	r = &SiteService{}
	r.Options = opts
	return
}

// Report for misclassification of a site.
func (r *SiteService) Report(ctx context.Context, body SiteReportParams, opts ...option.RequestOption) (res *SiteReportResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/site/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Scan Site
func (r *SiteService) Scan(ctx context.Context, body SiteScanParams, opts ...option.RequestOption) (res *SiteScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/site/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SiteScanHitResponse struct {
	AttackTypes       map[string]SiteScanHitResponseAttackType `json:"attack_types,required"`
	ContractRead      SiteScanHitResponseContractRead          `json:"contract_read,required"`
	ContractWrite     SiteScanHitResponseContractWrite         `json:"contract_write,required"`
	IsMalicious       bool                                     `json:"is_malicious,required"`
	IsReachable       bool                                     `json:"is_reachable,required"`
	IsWeb3Site        bool                                     `json:"is_web3_site,required"`
	JsonRpcOperations []string                                 `json:"json_rpc_operations,required"`
	MaliciousScore    float64                                  `json:"malicious_score,required"`
	NetworkOperations []string                                 `json:"network_operations,required"`
	ScanEndTime       time.Time                                `json:"scan_end_time,required" format:"date-time"`
	ScanStartTime     time.Time                                `json:"scan_start_time,required" format:"date-time"`
	Status            SiteScanHitResponseStatus                `json:"status,required"`
	URL               string                                   `json:"url,required"`
	JSON              siteScanHitResponseJSON                  `json:"-"`
}

// siteScanHitResponseJSON contains the JSON metadata for the struct
// [SiteScanHitResponse]
type siteScanHitResponseJSON struct {
	AttackTypes       apijson.Field
	ContractRead      apijson.Field
	ContractWrite     apijson.Field
	IsMalicious       apijson.Field
	IsReachable       apijson.Field
	IsWeb3Site        apijson.Field
	JsonRpcOperations apijson.Field
	MaliciousScore    apijson.Field
	NetworkOperations apijson.Field
	ScanEndTime       apijson.Field
	ScanStartTime     apijson.Field
	Status            apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SiteScanHitResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteScanHitResponseJSON) RawJSON() string {
	return r.raw
}

func (r SiteScanHitResponse) implementsSiteScanResponse() {}

type SiteScanHitResponseAttackType struct {
	Score     float64                           `json:"score,required"`
	Threshold float64                           `json:"threshold,required"`
	JSON      siteScanHitResponseAttackTypeJSON `json:"-"`
}

// siteScanHitResponseAttackTypeJSON contains the JSON metadata for the struct
// [SiteScanHitResponseAttackType]
type siteScanHitResponseAttackTypeJSON struct {
	Score       apijson.Field
	Threshold   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteScanHitResponseAttackType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteScanHitResponseAttackTypeJSON) RawJSON() string {
	return r.raw
}

type SiteScanHitResponseContractRead struct {
	ContractAddresses []string                            `json:"contract_addresses,required"`
	Functions         map[string][]string                 `json:"functions,required"`
	JSON              siteScanHitResponseContractReadJSON `json:"-"`
}

// siteScanHitResponseContractReadJSON contains the JSON metadata for the struct
// [SiteScanHitResponseContractRead]
type siteScanHitResponseContractReadJSON struct {
	ContractAddresses apijson.Field
	Functions         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SiteScanHitResponseContractRead) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteScanHitResponseContractReadJSON) RawJSON() string {
	return r.raw
}

type SiteScanHitResponseContractWrite struct {
	ContractAddresses []string                             `json:"contract_addresses,required"`
	Functions         map[string][]string                  `json:"functions,required"`
	JSON              siteScanHitResponseContractWriteJSON `json:"-"`
}

// siteScanHitResponseContractWriteJSON contains the JSON metadata for the struct
// [SiteScanHitResponseContractWrite]
type siteScanHitResponseContractWriteJSON struct {
	ContractAddresses apijson.Field
	Functions         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SiteScanHitResponseContractWrite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteScanHitResponseContractWriteJSON) RawJSON() string {
	return r.raw
}

type SiteScanHitResponseStatus string

const (
	SiteScanHitResponseStatusHit SiteScanHitResponseStatus = "hit"
)

func (r SiteScanHitResponseStatus) IsKnown() bool {
	switch r {
	case SiteScanHitResponseStatusHit:
		return true
	}
	return false
}

type SiteScanMissResponse struct {
	Status SiteScanMissResponseStatus `json:"status,required"`
	JSON   siteScanMissResponseJSON   `json:"-"`
}

// siteScanMissResponseJSON contains the JSON metadata for the struct
// [SiteScanMissResponse]
type siteScanMissResponseJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteScanMissResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteScanMissResponseJSON) RawJSON() string {
	return r.raw
}

func (r SiteScanMissResponse) implementsSiteScanResponse() {}

type SiteScanMissResponseStatus string

const (
	SiteScanMissResponseStatusMiss SiteScanMissResponseStatus = "miss"
)

func (r SiteScanMissResponseStatus) IsKnown() bool {
	switch r {
	case SiteScanMissResponseStatusMiss:
		return true
	}
	return false
}

type SiteReportResponse = interface{}

type SiteScanResponse struct {
	Status         SiteScanResponseStatus `json:"status,required"`
	URL            string                 `json:"url"`
	ScanStartTime  time.Time              `json:"scan_start_time" format:"date-time"`
	ScanEndTime    time.Time              `json:"scan_end_time" format:"date-time"`
	MaliciousScore float64                `json:"malicious_score"`
	IsReachable    bool                   `json:"is_reachable"`
	IsWeb3Site     bool                   `json:"is_web3_site"`
	IsMalicious    bool                   `json:"is_malicious"`
	// This field can have the runtime type of
	// [map[string]SiteScanHitResponseAttackType].
	AttackTypes interface{} `json:"attack_types,required"`
	// This field can have the runtime type of [[]string].
	NetworkOperations interface{} `json:"network_operations,required"`
	// This field can have the runtime type of [[]string].
	JsonRpcOperations interface{} `json:"json_rpc_operations,required"`
	// This field can have the runtime type of [SiteScanHitResponseContractWrite].
	ContractWrite interface{} `json:"contract_write,required"`
	// This field can have the runtime type of [SiteScanHitResponseContractRead].
	ContractRead interface{}          `json:"contract_read,required"`
	JSON         siteScanResponseJSON `json:"-"`
	union        SiteScanResponseUnion
}

// siteScanResponseJSON contains the JSON metadata for the struct
// [SiteScanResponse]
type siteScanResponseJSON struct {
	Status            apijson.Field
	URL               apijson.Field
	ScanStartTime     apijson.Field
	ScanEndTime       apijson.Field
	MaliciousScore    apijson.Field
	IsReachable       apijson.Field
	IsWeb3Site        apijson.Field
	IsMalicious       apijson.Field
	AttackTypes       apijson.Field
	NetworkOperations apijson.Field
	JsonRpcOperations apijson.Field
	ContractWrite     apijson.Field
	ContractRead      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r siteScanResponseJSON) RawJSON() string {
	return r.raw
}

func (r *SiteScanResponse) UnmarshalJSON(data []byte) (err error) {
	*r = SiteScanResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SiteScanResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [SiteScanHitResponse],
// [SiteScanMissResponse].
func (r SiteScanResponse) AsUnion() SiteScanResponseUnion {
	return r.union
}

// Union satisfied by [SiteScanHitResponse] or [SiteScanMissResponse].
type SiteScanResponseUnion interface {
	implementsSiteScanResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SiteScanResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SiteScanHitResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SiteScanMissResponse{}),
		},
	)
}

type SiteScanResponseStatus string

const (
	SiteScanResponseStatusHit  SiteScanResponseStatus = "hit"
	SiteScanResponseStatusMiss SiteScanResponseStatus = "miss"
)

func (r SiteScanResponseStatus) IsKnown() bool {
	switch r {
	case SiteScanResponseStatusHit, SiteScanResponseStatusMiss:
		return true
	}
	return false
}

type SiteReportParams struct {
	// Details about the report.
	Details param.Field[string] `json:"details,required"`
	// An enumeration.
	Event param.Field[SiteReportParamsEvent] `json:"event,required"`
	// The report parameters.
	Report param.Field[SiteReportParamsReportUnion] `json:"report,required"`
}

func (r SiteReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An enumeration.
type SiteReportParamsEvent string

const (
	SiteReportParamsEventFalsePositive SiteReportParamsEvent = "FALSE_POSITIVE"
	SiteReportParamsEventFalseNegative SiteReportParamsEvent = "FALSE_NEGATIVE"
)

func (r SiteReportParamsEvent) IsKnown() bool {
	switch r {
	case SiteReportParamsEventFalsePositive, SiteReportParamsEventFalseNegative:
		return true
	}
	return false
}

// The report parameters.
type SiteReportParamsReport struct {
	Type      param.Field[SiteReportParamsReportType] `json:"type,required"`
	Params    param.Field[interface{}]                `json:"params,required"`
	RequestID param.Field[string]                     `json:"request_id"`
}

func (r SiteReportParamsReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SiteReportParamsReport) implementsSiteReportParamsReportUnion() {}

// The report parameters.
//
// Satisfied by [SiteReportParamsReportParamReportSiteReportParams],
// [SiteReportParamsReportRequestIDReport], [SiteReportParamsReport].
type SiteReportParamsReportUnion interface {
	implementsSiteReportParamsReportUnion()
}

type SiteReportParamsReportParamReportSiteReportParams struct {
	Params param.Field[SiteReportParamsReportParamReportSiteReportParamsParams] `json:"params,required"`
	Type   param.Field[SiteReportParamsReportParamReportSiteReportParamsType]   `json:"type,required"`
}

func (r SiteReportParamsReportParamReportSiteReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SiteReportParamsReportParamReportSiteReportParams) implementsSiteReportParamsReportUnion() {}

type SiteReportParamsReportParamReportSiteReportParamsParams struct {
	// The url of the site to report on.
	URL param.Field[string] `json:"url,required"`
}

func (r SiteReportParamsReportParamReportSiteReportParamsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteReportParamsReportParamReportSiteReportParamsType string

const (
	SiteReportParamsReportParamReportSiteReportParamsTypeParams SiteReportParamsReportParamReportSiteReportParamsType = "params"
)

func (r SiteReportParamsReportParamReportSiteReportParamsType) IsKnown() bool {
	switch r {
	case SiteReportParamsReportParamReportSiteReportParamsTypeParams:
		return true
	}
	return false
}

type SiteReportParamsReportRequestIDReport struct {
	RequestID param.Field[string]                                    `json:"request_id,required"`
	Type      param.Field[SiteReportParamsReportRequestIDReportType] `json:"type,required"`
}

func (r SiteReportParamsReportRequestIDReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SiteReportParamsReportRequestIDReport) implementsSiteReportParamsReportUnion() {}

type SiteReportParamsReportRequestIDReportType string

const (
	SiteReportParamsReportRequestIDReportTypeRequestID SiteReportParamsReportRequestIDReportType = "request_id"
)

func (r SiteReportParamsReportRequestIDReportType) IsKnown() bool {
	switch r {
	case SiteReportParamsReportRequestIDReportTypeRequestID:
		return true
	}
	return false
}

type SiteReportParamsReportType string

const (
	SiteReportParamsReportTypeParams    SiteReportParamsReportType = "params"
	SiteReportParamsReportTypeRequestID SiteReportParamsReportType = "request_id"
)

func (r SiteReportParamsReportType) IsKnown() bool {
	switch r {
	case SiteReportParamsReportTypeParams, SiteReportParamsReportTypeRequestID:
		return true
	}
	return false
}

type SiteScanParams struct {
	URL      param.Field[string]                      `json:"url,required"`
	Metadata param.Field[SiteScanParamsMetadataUnion] `json:"metadata"`
}

func (r SiteScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteScanParamsMetadata struct {
	Type                     param.Field[SiteScanParamsMetadataType] `json:"type,required"`
	AccountAddress           param.Field[string]                     `json:"account_address"`
	AccountAddresses         param.Field[interface{}]                `json:"account_addresses,required"`
	WalletconnectName        param.Field[string]                     `json:"walletconnect_name"`
	WalletconnectDescription param.Field[string]                     `json:"walletconnect_description"`
}

func (r SiteScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SiteScanParamsMetadata) implementsSiteScanParamsMetadataUnion() {}

// Satisfied by [SiteScanParamsMetadataCatalogRequestMetadata],
// [SiteScanParamsMetadataWalletRequestMetadata], [SiteScanParamsMetadata].
type SiteScanParamsMetadataUnion interface {
	implementsSiteScanParamsMetadataUnion()
}

type SiteScanParamsMetadataCatalogRequestMetadata struct {
	Type param.Field[SiteScanParamsMetadataCatalogRequestMetadataType] `json:"type,required"`
}

func (r SiteScanParamsMetadataCatalogRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SiteScanParamsMetadataCatalogRequestMetadata) implementsSiteScanParamsMetadataUnion() {}

type SiteScanParamsMetadataCatalogRequestMetadataType string

const (
	SiteScanParamsMetadataCatalogRequestMetadataTypeCatalog SiteScanParamsMetadataCatalogRequestMetadataType = "catalog"
)

func (r SiteScanParamsMetadataCatalogRequestMetadataType) IsKnown() bool {
	switch r {
	case SiteScanParamsMetadataCatalogRequestMetadataTypeCatalog:
		return true
	}
	return false
}

type SiteScanParamsMetadataWalletRequestMetadata struct {
	AccountAddress param.Field[string]                                          `json:"account_address,required"`
	Type           param.Field[SiteScanParamsMetadataWalletRequestMetadataType] `json:"type,required"`
	// List of all account addresses in different chains based on the CAIPs standard
	// (https://github.com/ChainAgnostic/CAIPs). Ethereum mainnet example:
	// eip155:1:0xab16a96d359ec26a11e2c2b3d8f8b8942d5bfcdb
	AccountAddresses         param.Field[[]string] `json:"account_addresses"`
	WalletconnectDescription param.Field[string]   `json:"walletconnect_description"`
	WalletconnectName        param.Field[string]   `json:"walletconnect_name"`
}

func (r SiteScanParamsMetadataWalletRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SiteScanParamsMetadataWalletRequestMetadata) implementsSiteScanParamsMetadataUnion() {}

type SiteScanParamsMetadataWalletRequestMetadataType string

const (
	SiteScanParamsMetadataWalletRequestMetadataTypeWallet SiteScanParamsMetadataWalletRequestMetadataType = "wallet"
)

func (r SiteScanParamsMetadataWalletRequestMetadataType) IsKnown() bool {
	switch r {
	case SiteScanParamsMetadataWalletRequestMetadataTypeWallet:
		return true
	}
	return false
}

type SiteScanParamsMetadataType string

const (
	SiteScanParamsMetadataTypeCatalog SiteScanParamsMetadataType = "catalog"
	SiteScanParamsMetadataTypeWallet  SiteScanParamsMetadataType = "wallet"
)

func (r SiteScanParamsMetadataType) IsKnown() bool {
	switch r {
	case SiteScanParamsMetadataTypeCatalog, SiteScanParamsMetadataTypeWallet:
		return true
	}
	return false
}
