// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/param"
	"github.com/blockaid-official/blockaid-client-go/internal/requestconfig"
	"github.com/blockaid-official/blockaid-client-go/option"
)

// TokenBulkExportService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTokenBulkExportService] method instead.
type TokenBulkExportService struct {
	Options []option.RequestOption
}

// NewTokenBulkExportService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTokenBulkExportService(opts ...option.RequestOption) (r *TokenBulkExportService) {
	r = &TokenBulkExportService{}
	r.Options = opts
	return
}

// Creates an asynchronous job to export token scan data for a chain in the
// requested format.
func (r *TokenBulkExportService) New(ctx context.Context, body TokenBulkExportNewParams, opts ...option.RequestOption) (res *TokenBulkExportNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/token/bulk-export/jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves the current status of the job and download URLs when status is
// `succeeded`.
func (r *TokenBulkExportService) Status(ctx context.Context, jobID string, opts ...option.RequestOption) (res *TokenBulkExportStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/token/bulk-export/jobs/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type TokenBulkExportNewResponse struct {
	JobID string `json:"job_id" api:"required"`
	// Current state of the job: `queued` (accepted, waiting to start), `running` (in
	// progress), `succeeded` (completed successfully), `failed` (completed with an
	// error).
	Status TokenBulkExportNewResponseStatus `json:"status" api:"required"`
	JSON   tokenBulkExportNewResponseJSON   `json:"-"`
}

// tokenBulkExportNewResponseJSON contains the JSON metadata for the struct
// [TokenBulkExportNewResponse]
type tokenBulkExportNewResponseJSON struct {
	JobID       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenBulkExportNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkExportNewResponseJSON) RawJSON() string {
	return r.raw
}

// Current state of the job: `queued` (accepted, waiting to start), `running` (in
// progress), `succeeded` (completed successfully), `failed` (completed with an
// error).
type TokenBulkExportNewResponseStatus string

const (
	TokenBulkExportNewResponseStatusQueued    TokenBulkExportNewResponseStatus = "queued"
	TokenBulkExportNewResponseStatusRunning   TokenBulkExportNewResponseStatus = "running"
	TokenBulkExportNewResponseStatusSucceeded TokenBulkExportNewResponseStatus = "succeeded"
	TokenBulkExportNewResponseStatusFailed    TokenBulkExportNewResponseStatus = "failed"
)

func (r TokenBulkExportNewResponseStatus) IsKnown() bool {
	switch r {
	case TokenBulkExportNewResponseStatusQueued, TokenBulkExportNewResponseStatusRunning, TokenBulkExportNewResponseStatusSucceeded, TokenBulkExportNewResponseStatusFailed:
		return true
	}
	return false
}

type TokenBulkExportStatusResponse struct {
	// Current state of the job: `queued` (accepted, waiting to start), `running` (in
	// progress), `succeeded` (completed successfully), `failed` (completed with an
	// error).
	Status TokenBulkExportStatusResponseStatus `json:"status" api:"required"`
	// The chain name
	Chain TokenScanSupportedChain `json:"chain" api:"nullable"`
	// Seconds until URLs expire
	ExpiresIn int64    `json:"expires_in" api:"nullable"`
	Files     []string `json:"files" api:"nullable"`
	// Number of tokens exported
	TokensCount int64                             `json:"tokens_count" api:"nullable"`
	JSON        tokenBulkExportStatusResponseJSON `json:"-"`
}

// tokenBulkExportStatusResponseJSON contains the JSON metadata for the struct
// [TokenBulkExportStatusResponse]
type tokenBulkExportStatusResponseJSON struct {
	Status      apijson.Field
	Chain       apijson.Field
	ExpiresIn   apijson.Field
	Files       apijson.Field
	TokensCount apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenBulkExportStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkExportStatusResponseJSON) RawJSON() string {
	return r.raw
}

// Current state of the job: `queued` (accepted, waiting to start), `running` (in
// progress), `succeeded` (completed successfully), `failed` (completed with an
// error).
type TokenBulkExportStatusResponseStatus string

const (
	TokenBulkExportStatusResponseStatusQueued    TokenBulkExportStatusResponseStatus = "queued"
	TokenBulkExportStatusResponseStatusRunning   TokenBulkExportStatusResponseStatus = "running"
	TokenBulkExportStatusResponseStatusSucceeded TokenBulkExportStatusResponseStatus = "succeeded"
	TokenBulkExportStatusResponseStatusFailed    TokenBulkExportStatusResponseStatus = "failed"
)

func (r TokenBulkExportStatusResponseStatus) IsKnown() bool {
	switch r {
	case TokenBulkExportStatusResponseStatusQueued, TokenBulkExportStatusResponseStatusRunning, TokenBulkExportStatusResponseStatusSucceeded, TokenBulkExportStatusResponseStatusFailed:
		return true
	}
	return false
}

type TokenBulkExportNewParams struct {
	// The chain name
	Chain param.Field[TokenScanSupportedChain] `json:"chain" api:"required"`
	// Output file format for the export: `jsonl_gzip` (JSON Lines compressed with
	// gzip) or `parquet_snappy` (Apache Parquet compressed with Snappy).
	Format param.Field[TokenBulkExportNewParamsFormat] `json:"format" api:"required"`
	// Export only records with timestamps at or after this Unix timestamp. If empty or
	// 0, all records will be exported.
	SyncStartTime param.Field[int64] `json:"sync_start_time"`
	// Type of token standard: `fungible` for interchangeable tokens (e.g., ERC-20), or
	// `non_fungible` for unique tokens (e.g., NFTs such as ERC-721/ERC-1155).
	TokenType param.Field[TokenBulkExportNewParamsTokenType] `json:"token_type"`
}

func (r TokenBulkExportNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Output file format for the export: `jsonl_gzip` (JSON Lines compressed with
// gzip) or `parquet_snappy` (Apache Parquet compressed with Snappy).
type TokenBulkExportNewParamsFormat string

const (
	TokenBulkExportNewParamsFormatJSONLGzip     TokenBulkExportNewParamsFormat = "jsonl_gzip"
	TokenBulkExportNewParamsFormatParquetSnappy TokenBulkExportNewParamsFormat = "parquet_snappy"
)

func (r TokenBulkExportNewParamsFormat) IsKnown() bool {
	switch r {
	case TokenBulkExportNewParamsFormatJSONLGzip, TokenBulkExportNewParamsFormatParquetSnappy:
		return true
	}
	return false
}

// Type of token standard: `fungible` for interchangeable tokens (e.g., ERC-20), or
// `non_fungible` for unique tokens (e.g., NFTs such as ERC-721/ERC-1155).
type TokenBulkExportNewParamsTokenType string

const (
	TokenBulkExportNewParamsTokenTypeFungible    TokenBulkExportNewParamsTokenType = "fungible"
	TokenBulkExportNewParamsTokenTypeNonFungible TokenBulkExportNewParamsTokenType = "non_fungible"
)

func (r TokenBulkExportNewParamsTokenType) IsKnown() bool {
	switch r {
	case TokenBulkExportNewParamsTokenTypeFungible, TokenBulkExportNewParamsTokenTypeNonFungible:
		return true
	}
	return false
}
