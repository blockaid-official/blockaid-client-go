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

// Report whether the end-user accepted the Blockaid classification on the entity
// being scanned.
func (r *ScanService) Status(ctx context.Context, body ScanStatusParams, opts ...option.RequestOption) (res *ScanStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/scan/status/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ScanStatusResponse = interface{}

type ScanStatusParams struct {
	// The x-request-id header returned from the previous Blockaid api request
	RequestID param.Field[string] `json:"request_id,required"`
	// An enumeration.
	Status param.Field[ScanStatusParamsStatus] `json:"status,required"`
}

func (r ScanStatusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An enumeration.
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
