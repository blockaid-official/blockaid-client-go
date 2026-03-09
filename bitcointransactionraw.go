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

// BitcoinTransactionRawService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBitcoinTransactionRawService] method instead.
type BitcoinTransactionRawService struct {
	Options []option.RequestOption
}

// NewBitcoinTransactionRawService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBitcoinTransactionRawService(opts ...option.RequestOption) (r *BitcoinTransactionRawService) {
	r = &BitcoinTransactionRawService{}
	r.Options = opts
	return
}

// Submit an appeal or false-positive report for a Bitcoin transaction scan. Use
// when you believe a scan result was incorrect (e.g. should_be_benign,
// should_be_malicious, wrong_simulation_result).
func (r *BitcoinTransactionRawService) Report(ctx context.Context, body BitcoinTransactionRawReportParams, opts ...option.RequestOption) (res *int64, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/bitcoin/transaction/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Scan a raw Bitcoin transaction for security risks before signing. Returns a
// validation verdict (Benign, Warning, or Malicious) and, when requested, a
// simulation of asset changes.
func (r *BitcoinTransactionRawService) Scan(ctx context.Context, body BitcoinTransactionRawScanParams, opts ...option.RequestOption) (res *BitcoinTransactionScanResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/bitcoin/transaction-raw/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BitcoinTransactionRawReportParams struct {
	// Free-text explanation or context for the report.
	Details param.Field[string] `json:"details" api:"required"`
	// Type of appeal: what you believe was wrong with the scan result.
	Event param.Field[BitcoinTransactionRawReportParamsEvent] `json:"event" api:"required"`
	// Either a prior scan request ID or the full transaction parameters being
	// reported.
	Report param.Field[BitcoinTransactionRawReportParamsReportUnion] `json:"report" api:"required"`
}

func (r BitcoinTransactionRawReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of appeal: what you believe was wrong with the scan result.
type BitcoinTransactionRawReportParamsEvent string

const (
	BitcoinTransactionRawReportParamsEventShouldBeMalicious     BitcoinTransactionRawReportParamsEvent = "should_be_malicious"
	BitcoinTransactionRawReportParamsEventShouldBeBenign        BitcoinTransactionRawReportParamsEvent = "should_be_benign"
	BitcoinTransactionRawReportParamsEventWrongSimulationResult BitcoinTransactionRawReportParamsEvent = "wrong_simulation_result"
)

func (r BitcoinTransactionRawReportParamsEvent) IsKnown() bool {
	switch r {
	case BitcoinTransactionRawReportParamsEventShouldBeMalicious, BitcoinTransactionRawReportParamsEventShouldBeBenign, BitcoinTransactionRawReportParamsEventWrongSimulationResult:
		return true
	}
	return false
}

// Either a prior scan request ID or the full transaction parameters being
// reported.
type BitcoinTransactionRawReportParamsReport struct {
	// Request ID from a previous transaction scan response.
	ID param.Field[string] `json:"id"`
	// Full transaction scan request (same shape as the scan endpoint).
	Params param.Field[BitcoinTransactionScanRequestParam] `json:"params"`
	// Discriminator; use "request_id" when referencing a prior scan.
	Type param.Field[BitcoinTransactionRawReportParamsReportType] `json:"type"`
}

func (r BitcoinTransactionRawReportParamsReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionRawReportParamsReport) implementsBitcoinTransactionRawReportParamsReportUnion() {
}

// Either a prior scan request ID or the full transaction parameters being
// reported.
//
// Satisfied by [BitcoinTransactionRawReportParamsReportBitcoinAppealRequestID],
// [BitcoinTransactionRawReportParamsReportBitcoinAppealTransactionDataReport],
// [BitcoinTransactionRawReportParamsReport].
type BitcoinTransactionRawReportParamsReportUnion interface {
	implementsBitcoinTransactionRawReportParamsReportUnion()
}

type BitcoinTransactionRawReportParamsReportBitcoinAppealRequestID struct {
	// Request ID from a previous transaction scan response.
	ID param.Field[string] `json:"id" api:"required"`
	// Discriminator; use "request_id" when referencing a prior scan.
	Type param.Field[BitcoinTransactionRawReportParamsReportBitcoinAppealRequestIDType] `json:"type"`
}

func (r BitcoinTransactionRawReportParamsReportBitcoinAppealRequestID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionRawReportParamsReportBitcoinAppealRequestID) implementsBitcoinTransactionRawReportParamsReportUnion() {
}

// Discriminator; use "request_id" when referencing a prior scan.
type BitcoinTransactionRawReportParamsReportBitcoinAppealRequestIDType string

const (
	BitcoinTransactionRawReportParamsReportBitcoinAppealRequestIDTypeRequestID BitcoinTransactionRawReportParamsReportBitcoinAppealRequestIDType = "request_id"
)

func (r BitcoinTransactionRawReportParamsReportBitcoinAppealRequestIDType) IsKnown() bool {
	switch r {
	case BitcoinTransactionRawReportParamsReportBitcoinAppealRequestIDTypeRequestID:
		return true
	}
	return false
}

type BitcoinTransactionRawReportParamsReportBitcoinAppealTransactionDataReport struct {
	// Full transaction scan request (same shape as the scan endpoint).
	Params param.Field[BitcoinTransactionScanRequestParam] `json:"params" api:"required"`
	// Discriminator; use "params" when supplying full transaction data.
	Type param.Field[BitcoinTransactionRawReportParamsReportBitcoinAppealTransactionDataReportType] `json:"type"`
}

func (r BitcoinTransactionRawReportParamsReportBitcoinAppealTransactionDataReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionRawReportParamsReportBitcoinAppealTransactionDataReport) implementsBitcoinTransactionRawReportParamsReportUnion() {
}

// Discriminator; use "params" when supplying full transaction data.
type BitcoinTransactionRawReportParamsReportBitcoinAppealTransactionDataReportType string

const (
	BitcoinTransactionRawReportParamsReportBitcoinAppealTransactionDataReportTypeParams BitcoinTransactionRawReportParamsReportBitcoinAppealTransactionDataReportType = "params"
)

func (r BitcoinTransactionRawReportParamsReportBitcoinAppealTransactionDataReportType) IsKnown() bool {
	switch r {
	case BitcoinTransactionRawReportParamsReportBitcoinAppealTransactionDataReportTypeParams:
		return true
	}
	return false
}

// Discriminator; use "request_id" when referencing a prior scan.
type BitcoinTransactionRawReportParamsReportType string

const (
	BitcoinTransactionRawReportParamsReportTypeRequestID BitcoinTransactionRawReportParamsReportType = "request_id"
	BitcoinTransactionRawReportParamsReportTypeParams    BitcoinTransactionRawReportParamsReportType = "params"
)

func (r BitcoinTransactionRawReportParamsReportType) IsKnown() bool {
	switch r {
	case BitcoinTransactionRawReportParamsReportTypeRequestID, BitcoinTransactionRawReportParamsReportTypeParams:
		return true
	}
	return false
}

type BitcoinTransactionRawScanParams struct {
	BitcoinTransactionScanRequest BitcoinTransactionScanRequestParam `json:"BitcoinTransactionScanRequest" api:"required"`
}

func (r BitcoinTransactionRawScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.BitcoinTransactionScanRequest)
}
