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

// StarknetTransactionService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStarknetTransactionService] method instead.
type StarknetTransactionService struct {
	Options []option.RequestOption
}

// NewStarknetTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStarknetTransactionService(opts ...option.RequestOption) (r *StarknetTransactionService) {
	r = &StarknetTransactionService{}
	r.Options = opts
	return
}

// Report Transaction
func (r *StarknetTransactionService) Report(ctx context.Context, body StarknetTransactionReportParams, opts ...option.RequestOption) (res *int64, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/starknet/transaction/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Scan Transactions
func (r *StarknetTransactionService) Scan(ctx context.Context, body StarknetTransactionScanParams, opts ...option.RequestOption) (res *StarknetTransactionScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/starknet/transaction/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type StarknetTransactionReportParams struct {
	Details param.Field[string]                                     `json:"details,required"`
	Event   param.Field[StarknetTransactionReportParamsEvent]       `json:"event,required"`
	Report  param.Field[StarknetTransactionReportParamsReportUnion] `json:"report,required"`
}

func (r StarknetTransactionReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StarknetTransactionReportParamsEvent string

const (
	StarknetTransactionReportParamsEventShouldBeMalicious     StarknetTransactionReportParamsEvent = "should_be_malicious"
	StarknetTransactionReportParamsEventShouldBeBenign        StarknetTransactionReportParamsEvent = "should_be_benign"
	StarknetTransactionReportParamsEventWrongSimulationResult StarknetTransactionReportParamsEvent = "wrong_simulation_result"
)

func (r StarknetTransactionReportParamsEvent) IsKnown() bool {
	switch r {
	case StarknetTransactionReportParamsEventShouldBeMalicious, StarknetTransactionReportParamsEventShouldBeBenign, StarknetTransactionReportParamsEventWrongSimulationResult:
		return true
	}
	return false
}

type StarknetTransactionReportParamsReport struct {
	ID     param.Field[string]                                    `json:"id"`
	Params param.Field[StarknetTransactionScanRequestParam]       `json:"params"`
	Type   param.Field[StarknetTransactionReportParamsReportType] `json:"type"`
}

func (r StarknetTransactionReportParamsReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionReportParamsReport) implementsStarknetTransactionReportParamsReportUnion() {
}

// Satisfied by [StarknetTransactionReportParamsReportStarknetAppealRequestID],
// [StarknetTransactionReportParamsReportStarknetAppealTransactionDataReport],
// [StarknetTransactionReportParamsReport].
type StarknetTransactionReportParamsReportUnion interface {
	implementsStarknetTransactionReportParamsReportUnion()
}

type StarknetTransactionReportParamsReportStarknetAppealRequestID struct {
	ID   param.Field[string]                                                           `json:"id,required"`
	Type param.Field[StarknetTransactionReportParamsReportStarknetAppealRequestIDType] `json:"type"`
}

func (r StarknetTransactionReportParamsReportStarknetAppealRequestID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionReportParamsReportStarknetAppealRequestID) implementsStarknetTransactionReportParamsReportUnion() {
}

type StarknetTransactionReportParamsReportStarknetAppealRequestIDType string

const (
	StarknetTransactionReportParamsReportStarknetAppealRequestIDTypeRequestID StarknetTransactionReportParamsReportStarknetAppealRequestIDType = "request_id"
)

func (r StarknetTransactionReportParamsReportStarknetAppealRequestIDType) IsKnown() bool {
	switch r {
	case StarknetTransactionReportParamsReportStarknetAppealRequestIDTypeRequestID:
		return true
	}
	return false
}

type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReport struct {
	Params param.Field[StarknetTransactionScanRequestParam]                                          `json:"params,required"`
	Type   param.Field[StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportType] `json:"type"`
}

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReport) implementsStarknetTransactionReportParamsReportUnion() {
}

type StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportType string

const (
	StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportTypeParams StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportType = "params"
)

func (r StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportType) IsKnown() bool {
	switch r {
	case StarknetTransactionReportParamsReportStarknetAppealTransactionDataReportTypeParams:
		return true
	}
	return false
}

type StarknetTransactionReportParamsReportType string

const (
	StarknetTransactionReportParamsReportTypeRequestID StarknetTransactionReportParamsReportType = "request_id"
	StarknetTransactionReportParamsReportTypeParams    StarknetTransactionReportParamsReportType = "params"
)

func (r StarknetTransactionReportParamsReportType) IsKnown() bool {
	switch r {
	case StarknetTransactionReportParamsReportTypeRequestID, StarknetTransactionReportParamsReportTypeParams:
		return true
	}
	return false
}

type StarknetTransactionScanParams struct {
	StarknetTransactionScanRequest StarknetTransactionScanRequestParam `json:"StarknetTransactionScanRequest,required"`
}

func (r StarknetTransactionScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.StarknetTransactionScanRequest)
}
