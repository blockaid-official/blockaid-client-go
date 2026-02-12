// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/param"
	"github.com/blockaid-official/blockaid-client-go/internal/requestconfig"
	"github.com/blockaid-official/blockaid-client-go/option"
)

// TokenWebhookService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTokenWebhookService] method instead.
type TokenWebhookService struct {
	Options []option.RequestOption
}

// NewTokenWebhookService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTokenWebhookService(opts ...option.RequestOption) (r *TokenWebhookService) {
	r = &TokenWebhookService{}
	r.Options = opts
	return
}

// Creates a webhook subscription for a chain to receive real-time token scan
// updates.
func (r *TokenWebhookService) New(ctx context.Context, chain TokenScanSupportedChain, body TokenWebhookNewParams, opts ...option.RequestOption) (res *TokenWebhookNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v0/token/hooks/%v", chain)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Deletes the webhook subscription for a chain and stops further token scan
// updates.
func (r *TokenWebhookService) Delete(ctx context.Context, chain TokenScanSupportedChain, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("v0/token/hooks/%v", chain)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get information about an existing webhook for a given chain
func (r *TokenWebhookService) Get(ctx context.Context, chain TokenScanSupportedChain, opts ...option.RequestOption) (res *TokenWebhookGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v0/token/hooks/%v", chain)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all active webhook subscriptions across all chains
func (r *TokenWebhookService) GetAll(ctx context.Context, opts ...option.RequestOption) (res *[]TokenWebhookGetAllResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/token/hooks/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TokenWebhookNewResponse struct {
	Active bool `json:"active,required"`
	// The chain name
	Chain     TokenScanSupportedChain `json:"chain,required"`
	CreatedAt time.Time               `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time               `json:"updated_at,required" format:"date-time"`
	URL       string                  `json:"url,required" format:"uri"`
	// Optional shared secret key (32 characters), used to calculate the HMAC signature
	SharedSecretKey string                      `json:"shared_secret_key,nullable"`
	JSON            tokenWebhookNewResponseJSON `json:"-"`
}

// tokenWebhookNewResponseJSON contains the JSON metadata for the struct
// [TokenWebhookNewResponse]
type tokenWebhookNewResponseJSON struct {
	Active          apijson.Field
	Chain           apijson.Field
	CreatedAt       apijson.Field
	UpdatedAt       apijson.Field
	URL             apijson.Field
	SharedSecretKey apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TokenWebhookNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenWebhookNewResponseJSON) RawJSON() string {
	return r.raw
}

type TokenWebhookGetResponse struct {
	Active bool `json:"active,required"`
	// The chain name
	Chain     TokenScanSupportedChain `json:"chain,required"`
	CreatedAt time.Time               `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time               `json:"updated_at,required" format:"date-time"`
	URL       string                  `json:"url,required" format:"uri"`
	// Optional shared secret key (32 characters), used to calculate the HMAC signature
	SharedSecretKey string                      `json:"shared_secret_key,nullable"`
	JSON            tokenWebhookGetResponseJSON `json:"-"`
}

// tokenWebhookGetResponseJSON contains the JSON metadata for the struct
// [TokenWebhookGetResponse]
type tokenWebhookGetResponseJSON struct {
	Active          apijson.Field
	Chain           apijson.Field
	CreatedAt       apijson.Field
	UpdatedAt       apijson.Field
	URL             apijson.Field
	SharedSecretKey apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TokenWebhookGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenWebhookGetResponseJSON) RawJSON() string {
	return r.raw
}

type TokenWebhookGetAllResponse struct {
	Active bool `json:"active,required"`
	// The chain name
	Chain     TokenScanSupportedChain `json:"chain,required"`
	CreatedAt time.Time               `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time               `json:"updated_at,required" format:"date-time"`
	URL       string                  `json:"url,required" format:"uri"`
	// Optional shared secret key (32 characters), used to calculate the HMAC signature
	SharedSecretKey string                         `json:"shared_secret_key,nullable"`
	JSON            tokenWebhookGetAllResponseJSON `json:"-"`
}

// tokenWebhookGetAllResponseJSON contains the JSON metadata for the struct
// [TokenWebhookGetAllResponse]
type tokenWebhookGetAllResponseJSON struct {
	Active          apijson.Field
	Chain           apijson.Field
	CreatedAt       apijson.Field
	UpdatedAt       apijson.Field
	URL             apijson.Field
	SharedSecretKey apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TokenWebhookGetAllResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenWebhookGetAllResponseJSON) RawJSON() string {
	return r.raw
}

type TokenWebhookNewParams struct {
	URL param.Field[string] `json:"url,required" format:"uri"`
	// Filter for webhook updates
	Filter param.Field[TokenWebhookNewParamsFilterUnion] `json:"filter"`
	// Optional shared secret key (32 characters), used to calculate the HMAC signature
	SharedSecretKey param.Field[string] `json:"shared_secret_key"`
}

func (r TokenWebhookNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter for webhook updates
type TokenWebhookNewParamsFilter struct {
	// Type of filter applied to the webhook updates
	FilterType     param.Field[TokenWebhookNewParamsFilterFilterType] `json:"filter_type"`
	TokenAddresses param.Field[interface{}]                           `json:"token_addresses"`
	// Type of token to filter
	TokenType param.Field[TokenWebhookNewParamsFilterTokenType] `json:"token_type"`
}

func (r TokenWebhookNewParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TokenWebhookNewParamsFilter) implementsTokenWebhookNewParamsFilterUnion() {}

// Filter for webhook updates
//
// Satisfied by [TokenWebhookNewParamsFilterTokenAddressFilter],
// [TokenWebhookNewParamsFilterTokenTypeFilter], [TokenWebhookNewParamsFilter].
type TokenWebhookNewParamsFilterUnion interface {
	implementsTokenWebhookNewParamsFilterUnion()
}

type TokenWebhookNewParamsFilterTokenAddressFilter struct {
	// List of up to 100000 token addresses to filter webhook updates
	TokenAddresses param.Field[[]string] `json:"token_addresses,required"`
	// Type of filter applied to the webhook updates
	FilterType param.Field[TokenWebhookNewParamsFilterTokenAddressFilterFilterType] `json:"filter_type"`
}

func (r TokenWebhookNewParamsFilterTokenAddressFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TokenWebhookNewParamsFilterTokenAddressFilter) implementsTokenWebhookNewParamsFilterUnion() {}

// Type of filter applied to the webhook updates
type TokenWebhookNewParamsFilterTokenAddressFilterFilterType string

const (
	TokenWebhookNewParamsFilterTokenAddressFilterFilterTypeTokenAddress TokenWebhookNewParamsFilterTokenAddressFilterFilterType = "token_address"
)

func (r TokenWebhookNewParamsFilterTokenAddressFilterFilterType) IsKnown() bool {
	switch r {
	case TokenWebhookNewParamsFilterTokenAddressFilterFilterTypeTokenAddress:
		return true
	}
	return false
}

type TokenWebhookNewParamsFilterTokenTypeFilter struct {
	// Type of token to filter
	TokenType param.Field[TokenWebhookNewParamsFilterTokenTypeFilterTokenType] `json:"token_type,required"`
	// Type of filter applied to the webhook updates
	FilterType param.Field[TokenWebhookNewParamsFilterTokenTypeFilterFilterType] `json:"filter_type"`
}

func (r TokenWebhookNewParamsFilterTokenTypeFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TokenWebhookNewParamsFilterTokenTypeFilter) implementsTokenWebhookNewParamsFilterUnion() {}

// Type of token to filter
type TokenWebhookNewParamsFilterTokenTypeFilterTokenType string

const (
	TokenWebhookNewParamsFilterTokenTypeFilterTokenTypeFungible    TokenWebhookNewParamsFilterTokenTypeFilterTokenType = "Fungible"
	TokenWebhookNewParamsFilterTokenTypeFilterTokenTypeNonFungible TokenWebhookNewParamsFilterTokenTypeFilterTokenType = "NonFungible"
)

func (r TokenWebhookNewParamsFilterTokenTypeFilterTokenType) IsKnown() bool {
	switch r {
	case TokenWebhookNewParamsFilterTokenTypeFilterTokenTypeFungible, TokenWebhookNewParamsFilterTokenTypeFilterTokenTypeNonFungible:
		return true
	}
	return false
}

// Type of filter applied to the webhook updates
type TokenWebhookNewParamsFilterTokenTypeFilterFilterType string

const (
	TokenWebhookNewParamsFilterTokenTypeFilterFilterTypeTokenType TokenWebhookNewParamsFilterTokenTypeFilterFilterType = "token_type"
)

func (r TokenWebhookNewParamsFilterTokenTypeFilterFilterType) IsKnown() bool {
	switch r {
	case TokenWebhookNewParamsFilterTokenTypeFilterFilterTypeTokenType:
		return true
	}
	return false
}

// Type of filter applied to the webhook updates
type TokenWebhookNewParamsFilterFilterType string

const (
	TokenWebhookNewParamsFilterFilterTypeTokenAddress TokenWebhookNewParamsFilterFilterType = "token_address"
	TokenWebhookNewParamsFilterFilterTypeTokenType    TokenWebhookNewParamsFilterFilterType = "token_type"
)

func (r TokenWebhookNewParamsFilterFilterType) IsKnown() bool {
	switch r {
	case TokenWebhookNewParamsFilterFilterTypeTokenAddress, TokenWebhookNewParamsFilterFilterTypeTokenType:
		return true
	}
	return false
}

// Type of token to filter
type TokenWebhookNewParamsFilterTokenType string

const (
	TokenWebhookNewParamsFilterTokenTypeFungible    TokenWebhookNewParamsFilterTokenType = "Fungible"
	TokenWebhookNewParamsFilterTokenTypeNonFungible TokenWebhookNewParamsFilterTokenType = "NonFungible"
)

func (r TokenWebhookNewParamsFilterTokenType) IsKnown() bool {
	switch r {
	case TokenWebhookNewParamsFilterTokenTypeFungible, TokenWebhookNewParamsFilterTokenTypeNonFungible:
		return true
	}
	return false
}
