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

// Report Transaction
func (r *BitcoinTransactionRawService) Report(ctx context.Context, body BitcoinTransactionRawReportParams, opts ...option.RequestOption) (res *int64, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/bitcoin/transaction/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Scan Transaction
func (r *BitcoinTransactionRawService) Scan(ctx context.Context, body BitcoinTransactionRawScanParams, opts ...option.RequestOption) (res *BitcoinTransactionScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/bitcoin/transaction-raw/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BitcoinTransactionRawReportParams struct {
	Details param.Field[string]                                       `json:"details,required"`
	Event   param.Field[BitcoinTransactionRawReportParamsEvent]       `json:"event,required"`
	Report  param.Field[BitcoinTransactionRawReportParamsReportUnion] `json:"report,required"`
}

func (r BitcoinTransactionRawReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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

type BitcoinTransactionRawReportParamsReport struct {
	ID     param.Field[string]                                      `json:"id"`
	Params param.Field[BitcoinTransactionScanRequestParam]          `json:"params"`
	Type   param.Field[BitcoinTransactionRawReportParamsReportType] `json:"type"`
}

func (r BitcoinTransactionRawReportParamsReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionRawReportParamsReport) implementsBitcoinTransactionRawReportParamsReportUnion() {
}

// Satisfied by [BitcoinTransactionRawReportParamsReportBitcoinAppealRequestID],
// [BitcoinTransactionRawReportParamsReportBitcoinAppealTransactionDataReport],
// [BitcoinTransactionRawReportParamsReport].
type BitcoinTransactionRawReportParamsReportUnion interface {
	implementsBitcoinTransactionRawReportParamsReportUnion()
}

type BitcoinTransactionRawReportParamsReportBitcoinAppealRequestID struct {
	ID   param.Field[string]                                                            `json:"id,required"`
	Type param.Field[BitcoinTransactionRawReportParamsReportBitcoinAppealRequestIDType] `json:"type"`
}

func (r BitcoinTransactionRawReportParamsReportBitcoinAppealRequestID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionRawReportParamsReportBitcoinAppealRequestID) implementsBitcoinTransactionRawReportParamsReportUnion() {
}

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
	Params param.Field[BitcoinTransactionScanRequestParam]                                            `json:"params,required"`
	Type   param.Field[BitcoinTransactionRawReportParamsReportBitcoinAppealTransactionDataReportType] `json:"type"`
}

func (r BitcoinTransactionRawReportParamsReportBitcoinAppealTransactionDataReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BitcoinTransactionRawReportParamsReportBitcoinAppealTransactionDataReport) implementsBitcoinTransactionRawReportParamsReportUnion() {
}

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
	BitcoinTransactionScanRequest BitcoinTransactionScanRequestParam `json:"BitcoinTransactionScanRequest,required"`
}

func (r BitcoinTransactionRawScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.BitcoinTransactionScanRequest)
}
