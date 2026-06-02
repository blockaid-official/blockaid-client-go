// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/param"
	"github.com/blockaid-official/blockaid-client-go/internal/requestconfig"
	"github.com/blockaid-official/blockaid-client-go/option"
)

// ScanService contains methods and other services that help with interacting with
// the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScanService] method instead.
type ScanService struct {
	Options []option.RequestOption
}

// NewScanService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewScanService(opts ...option.RequestOption) (r *ScanService) {
	r = &ScanService{}
	r.Options = opts
	return
}

// Report a misclassification of any entity (transaction, address, site, token,
// etc.) using a request ID from a previous scan.
func (r *ScanService) Report(ctx context.Context, body ScanReportParams, opts ...option.RequestOption) (res *ScanReportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/scan/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Report whether the end-user accepted the Blockaid classification on the entity
// being scanned.
func (r *ScanService) Status(ctx context.Context, body ScanStatusParams, opts ...option.RequestOption) (res *ScanStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/scan/status/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ScanReportResponse = interface{}

type ScanStatusResponse = interface{}

type ScanReportParams struct {
	// Details about the report.
	Details param.Field[string] `json:"details" api:"required"`
	// The event type of the report. Could be `FALSE_POSITIVE` or `FALSE_NEGATIVE`.
	Event param.Field[ScanReportParamsEvent] `json:"event" api:"required"`
	// Client-side context attached to a report, identifying the originating dApp and
	// end-user.
	Metadata param.Field[ScanReportParamsMetadata] `json:"metadata" api:"required"`
	// The request ID of a previous scan. This can be found in the value of the
	// `x-request-id` field in the headers of the response of the previous request. For
	// instance: `6c3cf6c1-a80d-4927-91b9-03d841ea61fe`.
	RequestID param.Field[string] `json:"request_id" api:"required"`
}

func (r ScanReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The event type of the report. Could be `FALSE_POSITIVE` or `FALSE_NEGATIVE`.
type ScanReportParamsEvent string

const (
	ScanReportParamsEventFalsePositive ScanReportParamsEvent = "FALSE_POSITIVE"
	ScanReportParamsEventFalseNegative ScanReportParamsEvent = "FALSE_NEGATIVE"
)

func (r ScanReportParamsEvent) IsKnown() bool {
	switch r {
	case ScanReportParamsEventFalsePositive, ScanReportParamsEventFalseNegative:
		return true
	}
	return false
}

// Client-side context attached to a report, identifying the originating dApp and
// end-user.
type ScanReportParamsMetadata struct {
	// End-user account information (ID, age, country, creation time).
	Account param.Field[ScanReportParamsMetadataAccount] `json:"account"`
	// Network connection details (IP address, user agent).
	Connection param.Field[ScanReportParamsMetadataConnection] `json:"connection"`
	// The dApp or origin URL that triggered the interaction being reported.
	Domain param.Field[string] `json:"domain"`
}

func (r ScanReportParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// End-user account information (ID, age, country, creation time).
type ScanReportParamsMetadataAccount struct {
	// Unique identifier for the account.
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// Timestamp when the account was created.
	AccountCreationTimestamp param.Field[time.Time] `json:"account_creation_timestamp" format:"date-time"`
	// Age of the user in years
	UserAge param.Field[int64] `json:"user_age"`
	// ISO country code of the user's location.
	UserCountryCode param.Field[string] `json:"user_country_code"`
}

func (r ScanReportParamsMetadataAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Network connection details (IP address, user agent).
type ScanReportParamsMetadataConnection struct {
	// IP address of the customer making the request.
	IPAddress param.Field[string] `json:"ip_address" api:"required" format:"ipvanyaddress"`
	// User agent string from the client's browser or application.
	UserAgent param.Field[string] `json:"user_agent"`
}

func (r ScanReportParamsMetadataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScanStatusParams struct {
	// The x-request-id header returned from the previous Blockaid api request
	RequestID param.Field[string] `json:"request_id" api:"required"`
	// Status of a scan-status lookup request.
	Status param.Field[ScanStatusParamsStatus] `json:"status" api:"required"`
}

func (r ScanStatusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Status of a scan-status lookup request.
type ScanStatusParamsStatus string

const (
	ScanStatusParamsStatusAccepted ScanStatusParamsStatus = "accepted"
	ScanStatusParamsStatusRejected ScanStatusParamsStatus = "rejected"
)

func (r ScanStatusParamsStatus) IsKnown() bool {
	switch r {
	case ScanStatusParamsStatusAccepted, ScanStatusParamsStatusRejected:
		return true
	}
	return false
}
