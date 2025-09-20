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

// StellarTransactionService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStellarTransactionService] method instead.
type StellarTransactionService struct {
	Options []option.RequestOption
}

// NewStellarTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStellarTransactionService(opts ...option.RequestOption) (r *StellarTransactionService) {
	r = &StellarTransactionService{}
	r.Options = opts
	return
}

// Report Transaction
func (r *StellarTransactionService) Report(ctx context.Context, body StellarTransactionReportParams, opts ...option.RequestOption) (res *int64, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/stellar/transaction/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets a transaction and returns a full simulation indicating what will happen in
// the transaction together with a recommended action and some textual reasons of
// why the transaction was flagged that way.
func (r *StellarTransactionService) Scan(ctx context.Context, body StellarTransactionScanParams, opts ...option.RequestOption) (res *StellarTransactionScanResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/stellar/transaction/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type StellarTransactionReportParams struct {
	Details param.Field[string]                                    `json:"details,required"`
	Event   param.Field[StellarTransactionReportParamsEvent]       `json:"event,required"`
	Report  param.Field[StellarTransactionReportParamsReportUnion] `json:"report,required"`
}

func (r StellarTransactionReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StellarTransactionReportParamsEvent string

const (
	StellarTransactionReportParamsEventShouldBeMalicious     StellarTransactionReportParamsEvent = "should_be_malicious"
	StellarTransactionReportParamsEventShouldBeBenign        StellarTransactionReportParamsEvent = "should_be_benign"
	StellarTransactionReportParamsEventWrongSimulationResult StellarTransactionReportParamsEvent = "wrong_simulation_result"
)

func (r StellarTransactionReportParamsEvent) IsKnown() bool {
	switch r {
	case StellarTransactionReportParamsEventShouldBeMalicious, StellarTransactionReportParamsEventShouldBeBenign, StellarTransactionReportParamsEventWrongSimulationResult:
		return true
	}
	return false
}

type StellarTransactionReportParamsReport struct {
	ID     param.Field[string]                                   `json:"id"`
	Params param.Field[StellarTransactionScanRequestParam]       `json:"params"`
	Type   param.Field[StellarTransactionReportParamsReportType] `json:"type"`
}

func (r StellarTransactionReportParamsReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionReportParamsReport) implementsStellarTransactionReportParamsReportUnion() {}

// Satisfied by [StellarTransactionReportParamsReportStellarAppealRequestID],
// [StellarTransactionReportParamsReportStellarAppealTransactionDataReport],
// [StellarTransactionReportParamsReport].
type StellarTransactionReportParamsReportUnion interface {
	implementsStellarTransactionReportParamsReportUnion()
}

type StellarTransactionReportParamsReportStellarAppealRequestID struct {
	ID   param.Field[string]                                                         `json:"id,required"`
	Type param.Field[StellarTransactionReportParamsReportStellarAppealRequestIDType] `json:"type"`
}

func (r StellarTransactionReportParamsReportStellarAppealRequestID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionReportParamsReportStellarAppealRequestID) implementsStellarTransactionReportParamsReportUnion() {
}

type StellarTransactionReportParamsReportStellarAppealRequestIDType string

const (
	StellarTransactionReportParamsReportStellarAppealRequestIDTypeRequestID StellarTransactionReportParamsReportStellarAppealRequestIDType = "request_id"
)

func (r StellarTransactionReportParamsReportStellarAppealRequestIDType) IsKnown() bool {
	switch r {
	case StellarTransactionReportParamsReportStellarAppealRequestIDTypeRequestID:
		return true
	}
	return false
}

type StellarTransactionReportParamsReportStellarAppealTransactionDataReport struct {
	Params param.Field[StellarTransactionScanRequestParam]                                         `json:"params,required"`
	Type   param.Field[StellarTransactionReportParamsReportStellarAppealTransactionDataReportType] `json:"type"`
}

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReport) implementsStellarTransactionReportParamsReportUnion() {
}

type StellarTransactionReportParamsReportStellarAppealTransactionDataReportType string

const (
	StellarTransactionReportParamsReportStellarAppealTransactionDataReportTypeParams StellarTransactionReportParamsReportStellarAppealTransactionDataReportType = "params"
)

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportType) IsKnown() bool {
	switch r {
	case StellarTransactionReportParamsReportStellarAppealTransactionDataReportTypeParams:
		return true
	}
	return false
}

type StellarTransactionReportParamsReportType string

const (
	StellarTransactionReportParamsReportTypeRequestID StellarTransactionReportParamsReportType = "request_id"
	StellarTransactionReportParamsReportTypeParams    StellarTransactionReportParamsReportType = "params"
)

func (r StellarTransactionReportParamsReportType) IsKnown() bool {
	switch r {
	case StellarTransactionReportParamsReportTypeRequestID, StellarTransactionReportParamsReportTypeParams:
		return true
	}
	return false
}

type StellarTransactionScanParams struct {
	StellarTransactionScanRequest StellarTransactionScanRequestParam `json:"StellarTransactionScanRequest,required"`
}

func (r StellarTransactionScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.StellarTransactionScanRequest)
}
