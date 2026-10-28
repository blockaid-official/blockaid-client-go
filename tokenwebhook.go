// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"slices"
	"time"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/param"
	"github.com/blockaid-official/blockaid-client-go/internal/requestconfig"
	"github.com/blockaid-official/blockaid-client-go/option"
	"github.com/tidwall/gjson"
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
	return res, err
}

// Updates an existing webhook subscription in place. Only the provided fields are
// changed; omitted fields keep their current value. The filter's token_addresses
// can be replaced in full via `filter`, or adjusted incrementally via
// `add_token_addresses`/`remove_token_addresses` (mutually exclusive with
// `filter`).
func (r *TokenWebhookService) Update(ctx context.Context, chain TokenScanSupportedChain, body TokenWebhookUpdateParams, opts ...option.RequestOption) (res *TokenWebhookUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v0/token/hooks/%v", chain)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Deletes the webhook subscription for a chain and stops further token scan
// updates.
func (r *TokenWebhookService) Delete(ctx context.Context, chain TokenScanSupportedChain, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("v0/token/hooks/%v", chain)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get information about an existing webhook for a given chain
func (r *TokenWebhookService) Get(ctx context.Context, chain TokenScanSupportedChain, opts ...option.RequestOption) (res *TokenWebhookGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v0/token/hooks/%v", chain)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List all active webhook subscriptions across all chains
func (r *TokenWebhookService) GetAll(ctx context.Context, opts ...option.RequestOption) (res *[]TokenWebhookGetAllResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/token/hooks/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type TokenWebhookNewResponse struct {
	Active bool `json:"active" api:"required"`
	// The chain name
	Chain     TokenScanSupportedChain `json:"chain" api:"required"`
	CreatedAt time.Time               `json:"created_at" api:"required" format:"date-time"`
	UpdatedAt time.Time               `json:"updated_at" api:"required" format:"date-time"`
	URL       string                  `json:"url" api:"required" format:"uri"`
	// Filter for webhook updates
	Filter TokenWebhookNewResponseFilter `json:"filter" api:"nullable"`
	// Optional shared secret key (32 characters), used to calculate the HMAC signature
	SharedSecretKey string                      `json:"shared_secret_key" api:"nullable"`
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
	Filter          apijson.Field
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

// Filter for webhook updates
type TokenWebhookNewResponseFilter struct {
	// Type of filter applied to the webhook updates
	FilterType TokenWebhookNewResponseFilterFilterType `json:"filter_type"`
	// This field can have the runtime type of [[]string].
	TokenAddresses interface{} `json:"token_addresses"`
	// Type of token to filter
	TokenType TokenWebhookNewResponseFilterTokenType `json:"token_type"`
	JSON      tokenWebhookNewResponseFilterJSON      `json:"-"`
	union     TokenWebhookNewResponseFilterUnion
}

// tokenWebhookNewResponseFilterJSON contains the JSON metadata for the struct
// [TokenWebhookNewResponseFilter]
type tokenWebhookNewResponseFilterJSON struct {
	FilterType     apijson.Field
	TokenAddresses apijson.Field
	TokenType      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r tokenWebhookNewResponseFilterJSON) RawJSON() string {
	return r.raw
}

func (r *TokenWebhookNewResponseFilter) UnmarshalJSON(data []byte) (err error) {
	*r = TokenWebhookNewResponseFilter{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TokenWebhookNewResponseFilterUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TokenWebhookNewResponseFilterTokenAddressFilter],
// [TokenWebhookNewResponseFilterTokenTypeFilter].
func (r TokenWebhookNewResponseFilter) AsUnion() TokenWebhookNewResponseFilterUnion {
	return r.union
}

// Filter for webhook updates
//
// Union satisfied by [TokenWebhookNewResponseFilterTokenAddressFilter] or
// [TokenWebhookNewResponseFilterTokenTypeFilter].
type TokenWebhookNewResponseFilterUnion interface {
	implementsTokenWebhookNewResponseFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TokenWebhookNewResponseFilterUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenWebhookNewResponseFilterTokenAddressFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenWebhookNewResponseFilterTokenTypeFilter{}),
		},
	)
}

type TokenWebhookNewResponseFilterTokenAddressFilter struct {
	// List of up to 100000 token addresses to filter webhook updates
	TokenAddresses []string `json:"token_addresses" api:"required"`
	// Type of filter applied to the webhook updates
	FilterType TokenWebhookNewResponseFilterTokenAddressFilterFilterType `json:"filter_type"`
	JSON       tokenWebhookNewResponseFilterTokenAddressFilterJSON       `json:"-"`
}

// tokenWebhookNewResponseFilterTokenAddressFilterJSON contains the JSON metadata
// for the struct [TokenWebhookNewResponseFilterTokenAddressFilter]
type tokenWebhookNewResponseFilterTokenAddressFilterJSON struct {
	TokenAddresses apijson.Field
	FilterType     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TokenWebhookNewResponseFilterTokenAddressFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenWebhookNewResponseFilterTokenAddressFilterJSON) RawJSON() string {
	return r.raw
}

func (r TokenWebhookNewResponseFilterTokenAddressFilter) implementsTokenWebhookNewResponseFilter() {}

// Type of filter applied to the webhook updates
type TokenWebhookNewResponseFilterTokenAddressFilterFilterType string

const (
	TokenWebhookNewResponseFilterTokenAddressFilterFilterTypeTokenAddress TokenWebhookNewResponseFilterTokenAddressFilterFilterType = "token_address"
)

func (r TokenWebhookNewResponseFilterTokenAddressFilterFilterType) IsKnown() bool {
	switch r {
	case TokenWebhookNewResponseFilterTokenAddressFilterFilterTypeTokenAddress:
		return true
	}
	return false
}

type TokenWebhookNewResponseFilterTokenTypeFilter struct {
	// Type of token to filter
	TokenType TokenWebhookNewResponseFilterTokenTypeFilterTokenType `json:"token_type" api:"required"`
	// Type of filter applied to the webhook updates
	FilterType TokenWebhookNewResponseFilterTokenTypeFilterFilterType `json:"filter_type"`
	JSON       tokenWebhookNewResponseFilterTokenTypeFilterJSON       `json:"-"`
}

// tokenWebhookNewResponseFilterTokenTypeFilterJSON contains the JSON metadata for
// the struct [TokenWebhookNewResponseFilterTokenTypeFilter]
type tokenWebhookNewResponseFilterTokenTypeFilterJSON struct {
	TokenType   apijson.Field
	FilterType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenWebhookNewResponseFilterTokenTypeFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenWebhookNewResponseFilterTokenTypeFilterJSON) RawJSON() string {
	return r.raw
}

func (r TokenWebhookNewResponseFilterTokenTypeFilter) implementsTokenWebhookNewResponseFilter() {}

// Type of token to filter
type TokenWebhookNewResponseFilterTokenTypeFilterTokenType string

const (
	TokenWebhookNewResponseFilterTokenTypeFilterTokenTypeFungible    TokenWebhookNewResponseFilterTokenTypeFilterTokenType = "Fungible"
	TokenWebhookNewResponseFilterTokenTypeFilterTokenTypeNonFungible TokenWebhookNewResponseFilterTokenTypeFilterTokenType = "NonFungible"
)

func (r TokenWebhookNewResponseFilterTokenTypeFilterTokenType) IsKnown() bool {
	switch r {
	case TokenWebhookNewResponseFilterTokenTypeFilterTokenTypeFungible, TokenWebhookNewResponseFilterTokenTypeFilterTokenTypeNonFungible:
		return true
	}
	return false
}

// Type of filter applied to the webhook updates
type TokenWebhookNewResponseFilterTokenTypeFilterFilterType string

const (
	TokenWebhookNewResponseFilterTokenTypeFilterFilterTypeTokenType TokenWebhookNewResponseFilterTokenTypeFilterFilterType = "token_type"
)

func (r TokenWebhookNewResponseFilterTokenTypeFilterFilterType) IsKnown() bool {
	switch r {
	case TokenWebhookNewResponseFilterTokenTypeFilterFilterTypeTokenType:
		return true
	}
	return false
}

// Type of filter applied to the webhook updates
type TokenWebhookNewResponseFilterFilterType string

const (
	TokenWebhookNewResponseFilterFilterTypeTokenAddress TokenWebhookNewResponseFilterFilterType = "token_address"
	TokenWebhookNewResponseFilterFilterTypeTokenType    TokenWebhookNewResponseFilterFilterType = "token_type"
)

func (r TokenWebhookNewResponseFilterFilterType) IsKnown() bool {
	switch r {
	case TokenWebhookNewResponseFilterFilterTypeTokenAddress, TokenWebhookNewResponseFilterFilterTypeTokenType:
		return true
	}
	return false
}

// Type of token to filter
type TokenWebhookNewResponseFilterTokenType string

const (
	TokenWebhookNewResponseFilterTokenTypeFungible    TokenWebhookNewResponseFilterTokenType = "Fungible"
	TokenWebhookNewResponseFilterTokenTypeNonFungible TokenWebhookNewResponseFilterTokenType = "NonFungible"
)

func (r TokenWebhookNewResponseFilterTokenType) IsKnown() bool {
	switch r {
	case TokenWebhookNewResponseFilterTokenTypeFungible, TokenWebhookNewResponseFilterTokenTypeNonFungible:
		return true
	}
	return false
}

type TokenWebhookUpdateResponse struct {
	Active bool `json:"active" api:"required"`
	// The chain name
	Chain     TokenScanSupportedChain `json:"chain" api:"required"`
	CreatedAt time.Time               `json:"created_at" api:"required" format:"date-time"`
	UpdatedAt time.Time               `json:"updated_at" api:"required" format:"date-time"`
	URL       string                  `json:"url" api:"required" format:"uri"`
	// Filter for webhook updates
	Filter TokenWebhookUpdateResponseFilter `json:"filter" api:"nullable"`
	// Optional shared secret key (32 characters), used to calculate the HMAC signature
	SharedSecretKey string                         `json:"shared_secret_key" api:"nullable"`
	JSON            tokenWebhookUpdateResponseJSON `json:"-"`
}

// tokenWebhookUpdateResponseJSON contains the JSON metadata for the struct
// [TokenWebhookUpdateResponse]
type tokenWebhookUpdateResponseJSON struct {
	Active          apijson.Field
	Chain           apijson.Field
	CreatedAt       apijson.Field
	UpdatedAt       apijson.Field
	URL             apijson.Field
	Filter          apijson.Field
	SharedSecretKey apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TokenWebhookUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenWebhookUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Filter for webhook updates
type TokenWebhookUpdateResponseFilter struct {
	// Type of filter applied to the webhook updates
	FilterType TokenWebhookUpdateResponseFilterFilterType `json:"filter_type"`
	// This field can have the runtime type of [[]string].
	TokenAddresses interface{} `json:"token_addresses"`
	// Type of token to filter
	TokenType TokenWebhookUpdateResponseFilterTokenType `json:"token_type"`
	JSON      tokenWebhookUpdateResponseFilterJSON      `json:"-"`
	union     TokenWebhookUpdateResponseFilterUnion
}

// tokenWebhookUpdateResponseFilterJSON contains the JSON metadata for the struct
// [TokenWebhookUpdateResponseFilter]
type tokenWebhookUpdateResponseFilterJSON struct {
	FilterType     apijson.Field
	TokenAddresses apijson.Field
	TokenType      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r tokenWebhookUpdateResponseFilterJSON) RawJSON() string {
	return r.raw
}

func (r *TokenWebhookUpdateResponseFilter) UnmarshalJSON(data []byte) (err error) {
	*r = TokenWebhookUpdateResponseFilter{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TokenWebhookUpdateResponseFilterUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TokenWebhookUpdateResponseFilterTokenAddressFilter],
// [TokenWebhookUpdateResponseFilterTokenTypeFilter].
func (r TokenWebhookUpdateResponseFilter) AsUnion() TokenWebhookUpdateResponseFilterUnion {
	return r.union
}

// Filter for webhook updates
//
// Union satisfied by [TokenWebhookUpdateResponseFilterTokenAddressFilter] or
// [TokenWebhookUpdateResponseFilterTokenTypeFilter].
type TokenWebhookUpdateResponseFilterUnion interface {
	implementsTokenWebhookUpdateResponseFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TokenWebhookUpdateResponseFilterUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenWebhookUpdateResponseFilterTokenAddressFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenWebhookUpdateResponseFilterTokenTypeFilter{}),
		},
	)
}

type TokenWebhookUpdateResponseFilterTokenAddressFilter struct {
	// List of up to 100000 token addresses to filter webhook updates
	TokenAddresses []string `json:"token_addresses" api:"required"`
	// Type of filter applied to the webhook updates
	FilterType TokenWebhookUpdateResponseFilterTokenAddressFilterFilterType `json:"filter_type"`
	JSON       tokenWebhookUpdateResponseFilterTokenAddressFilterJSON       `json:"-"`
}

// tokenWebhookUpdateResponseFilterTokenAddressFilterJSON contains the JSON
// metadata for the struct [TokenWebhookUpdateResponseFilterTokenAddressFilter]
type tokenWebhookUpdateResponseFilterTokenAddressFilterJSON struct {
	TokenAddresses apijson.Field
	FilterType     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TokenWebhookUpdateResponseFilterTokenAddressFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenWebhookUpdateResponseFilterTokenAddressFilterJSON) RawJSON() string {
	return r.raw
}

func (r TokenWebhookUpdateResponseFilterTokenAddressFilter) implementsTokenWebhookUpdateResponseFilter() {
}

// Type of filter applied to the webhook updates
type TokenWebhookUpdateResponseFilterTokenAddressFilterFilterType string

const (
	TokenWebhookUpdateResponseFilterTokenAddressFilterFilterTypeTokenAddress TokenWebhookUpdateResponseFilterTokenAddressFilterFilterType = "token_address"
)

func (r TokenWebhookUpdateResponseFilterTokenAddressFilterFilterType) IsKnown() bool {
	switch r {
	case TokenWebhookUpdateResponseFilterTokenAddressFilterFilterTypeTokenAddress:
		return true
	}
	return false
}

type TokenWebhookUpdateResponseFilterTokenTypeFilter struct {
	// Type of token to filter
	TokenType TokenWebhookUpdateResponseFilterTokenTypeFilterTokenType `json:"token_type" api:"required"`
	// Type of filter applied to the webhook updates
	FilterType TokenWebhookUpdateResponseFilterTokenTypeFilterFilterType `json:"filter_type"`
	JSON       tokenWebhookUpdateResponseFilterTokenTypeFilterJSON       `json:"-"`
}

// tokenWebhookUpdateResponseFilterTokenTypeFilterJSON contains the JSON metadata
// for the struct [TokenWebhookUpdateResponseFilterTokenTypeFilter]
type tokenWebhookUpdateResponseFilterTokenTypeFilterJSON struct {
	TokenType   apijson.Field
	FilterType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenWebhookUpdateResponseFilterTokenTypeFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenWebhookUpdateResponseFilterTokenTypeFilterJSON) RawJSON() string {
	return r.raw
}

func (r TokenWebhookUpdateResponseFilterTokenTypeFilter) implementsTokenWebhookUpdateResponseFilter() {
}

// Type of token to filter
type TokenWebhookUpdateResponseFilterTokenTypeFilterTokenType string

const (
	TokenWebhookUpdateResponseFilterTokenTypeFilterTokenTypeFungible    TokenWebhookUpdateResponseFilterTokenTypeFilterTokenType = "Fungible"
	TokenWebhookUpdateResponseFilterTokenTypeFilterTokenTypeNonFungible TokenWebhookUpdateResponseFilterTokenTypeFilterTokenType = "NonFungible"
)

func (r TokenWebhookUpdateResponseFilterTokenTypeFilterTokenType) IsKnown() bool {
	switch r {
	case TokenWebhookUpdateResponseFilterTokenTypeFilterTokenTypeFungible, TokenWebhookUpdateResponseFilterTokenTypeFilterTokenTypeNonFungible:
		return true
	}
	return false
}

// Type of filter applied to the webhook updates
type TokenWebhookUpdateResponseFilterTokenTypeFilterFilterType string

const (
	TokenWebhookUpdateResponseFilterTokenTypeFilterFilterTypeTokenType TokenWebhookUpdateResponseFilterTokenTypeFilterFilterType = "token_type"
)

func (r TokenWebhookUpdateResponseFilterTokenTypeFilterFilterType) IsKnown() bool {
	switch r {
	case TokenWebhookUpdateResponseFilterTokenTypeFilterFilterTypeTokenType:
		return true
	}
	return false
}

// Type of filter applied to the webhook updates
type TokenWebhookUpdateResponseFilterFilterType string

const (
	TokenWebhookUpdateResponseFilterFilterTypeTokenAddress TokenWebhookUpdateResponseFilterFilterType = "token_address"
	TokenWebhookUpdateResponseFilterFilterTypeTokenType    TokenWebhookUpdateResponseFilterFilterType = "token_type"
)

func (r TokenWebhookUpdateResponseFilterFilterType) IsKnown() bool {
	switch r {
	case TokenWebhookUpdateResponseFilterFilterTypeTokenAddress, TokenWebhookUpdateResponseFilterFilterTypeTokenType:
		return true
	}
	return false
}

// Type of token to filter
type TokenWebhookUpdateResponseFilterTokenType string

const (
	TokenWebhookUpdateResponseFilterTokenTypeFungible    TokenWebhookUpdateResponseFilterTokenType = "Fungible"
	TokenWebhookUpdateResponseFilterTokenTypeNonFungible TokenWebhookUpdateResponseFilterTokenType = "NonFungible"
)

func (r TokenWebhookUpdateResponseFilterTokenType) IsKnown() bool {
	switch r {
	case TokenWebhookUpdateResponseFilterTokenTypeFungible, TokenWebhookUpdateResponseFilterTokenTypeNonFungible:
		return true
	}
	return false
}

type TokenWebhookGetResponse struct {
	Active bool `json:"active" api:"required"`
	// The chain name
	Chain     TokenScanSupportedChain `json:"chain" api:"required"`
	CreatedAt time.Time               `json:"created_at" api:"required" format:"date-time"`
	UpdatedAt time.Time               `json:"updated_at" api:"required" format:"date-time"`
	URL       string                  `json:"url" api:"required" format:"uri"`
	// Filter for webhook updates
	Filter TokenWebhookGetResponseFilter `json:"filter" api:"nullable"`
	// Optional shared secret key (32 characters), used to calculate the HMAC signature
	SharedSecretKey string                      `json:"shared_secret_key" api:"nullable"`
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
	Filter          apijson.Field
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

// Filter for webhook updates
type TokenWebhookGetResponseFilter struct {
	// Type of filter applied to the webhook updates
	FilterType TokenWebhookGetResponseFilterFilterType `json:"filter_type"`
	// This field can have the runtime type of [[]string].
	TokenAddresses interface{} `json:"token_addresses"`
	// Type of token to filter
	TokenType TokenWebhookGetResponseFilterTokenType `json:"token_type"`
	JSON      tokenWebhookGetResponseFilterJSON      `json:"-"`
	union     TokenWebhookGetResponseFilterUnion
}

// tokenWebhookGetResponseFilterJSON contains the JSON metadata for the struct
// [TokenWebhookGetResponseFilter]
type tokenWebhookGetResponseFilterJSON struct {
	FilterType     apijson.Field
	TokenAddresses apijson.Field
	TokenType      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r tokenWebhookGetResponseFilterJSON) RawJSON() string {
	return r.raw
}

func (r *TokenWebhookGetResponseFilter) UnmarshalJSON(data []byte) (err error) {
	*r = TokenWebhookGetResponseFilter{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TokenWebhookGetResponseFilterUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TokenWebhookGetResponseFilterTokenAddressFilter],
// [TokenWebhookGetResponseFilterTokenTypeFilter].
func (r TokenWebhookGetResponseFilter) AsUnion() TokenWebhookGetResponseFilterUnion {
	return r.union
}

// Filter for webhook updates
//
// Union satisfied by [TokenWebhookGetResponseFilterTokenAddressFilter] or
// [TokenWebhookGetResponseFilterTokenTypeFilter].
type TokenWebhookGetResponseFilterUnion interface {
	implementsTokenWebhookGetResponseFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TokenWebhookGetResponseFilterUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenWebhookGetResponseFilterTokenAddressFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenWebhookGetResponseFilterTokenTypeFilter{}),
		},
	)
}

type TokenWebhookGetResponseFilterTokenAddressFilter struct {
	// List of up to 100000 token addresses to filter webhook updates
	TokenAddresses []string `json:"token_addresses" api:"required"`
	// Type of filter applied to the webhook updates
	FilterType TokenWebhookGetResponseFilterTokenAddressFilterFilterType `json:"filter_type"`
	JSON       tokenWebhookGetResponseFilterTokenAddressFilterJSON       `json:"-"`
}

// tokenWebhookGetResponseFilterTokenAddressFilterJSON contains the JSON metadata
// for the struct [TokenWebhookGetResponseFilterTokenAddressFilter]
type tokenWebhookGetResponseFilterTokenAddressFilterJSON struct {
	TokenAddresses apijson.Field
	FilterType     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TokenWebhookGetResponseFilterTokenAddressFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenWebhookGetResponseFilterTokenAddressFilterJSON) RawJSON() string {
	return r.raw
}

func (r TokenWebhookGetResponseFilterTokenAddressFilter) implementsTokenWebhookGetResponseFilter() {}

// Type of filter applied to the webhook updates
type TokenWebhookGetResponseFilterTokenAddressFilterFilterType string

const (
	TokenWebhookGetResponseFilterTokenAddressFilterFilterTypeTokenAddress TokenWebhookGetResponseFilterTokenAddressFilterFilterType = "token_address"
)

func (r TokenWebhookGetResponseFilterTokenAddressFilterFilterType) IsKnown() bool {
	switch r {
	case TokenWebhookGetResponseFilterTokenAddressFilterFilterTypeTokenAddress:
		return true
	}
	return false
}

type TokenWebhookGetResponseFilterTokenTypeFilter struct {
	// Type of token to filter
	TokenType TokenWebhookGetResponseFilterTokenTypeFilterTokenType `json:"token_type" api:"required"`
	// Type of filter applied to the webhook updates
	FilterType TokenWebhookGetResponseFilterTokenTypeFilterFilterType `json:"filter_type"`
	JSON       tokenWebhookGetResponseFilterTokenTypeFilterJSON       `json:"-"`
}

// tokenWebhookGetResponseFilterTokenTypeFilterJSON contains the JSON metadata for
// the struct [TokenWebhookGetResponseFilterTokenTypeFilter]
type tokenWebhookGetResponseFilterTokenTypeFilterJSON struct {
	TokenType   apijson.Field
	FilterType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenWebhookGetResponseFilterTokenTypeFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenWebhookGetResponseFilterTokenTypeFilterJSON) RawJSON() string {
	return r.raw
}

func (r TokenWebhookGetResponseFilterTokenTypeFilter) implementsTokenWebhookGetResponseFilter() {}

// Type of token to filter
type TokenWebhookGetResponseFilterTokenTypeFilterTokenType string

const (
	TokenWebhookGetResponseFilterTokenTypeFilterTokenTypeFungible    TokenWebhookGetResponseFilterTokenTypeFilterTokenType = "Fungible"
	TokenWebhookGetResponseFilterTokenTypeFilterTokenTypeNonFungible TokenWebhookGetResponseFilterTokenTypeFilterTokenType = "NonFungible"
)

func (r TokenWebhookGetResponseFilterTokenTypeFilterTokenType) IsKnown() bool {
	switch r {
	case TokenWebhookGetResponseFilterTokenTypeFilterTokenTypeFungible, TokenWebhookGetResponseFilterTokenTypeFilterTokenTypeNonFungible:
		return true
	}
	return false
}

// Type of filter applied to the webhook updates
type TokenWebhookGetResponseFilterTokenTypeFilterFilterType string

const (
	TokenWebhookGetResponseFilterTokenTypeFilterFilterTypeTokenType TokenWebhookGetResponseFilterTokenTypeFilterFilterType = "token_type"
)

func (r TokenWebhookGetResponseFilterTokenTypeFilterFilterType) IsKnown() bool {
	switch r {
	case TokenWebhookGetResponseFilterTokenTypeFilterFilterTypeTokenType:
		return true
	}
	return false
}

// Type of filter applied to the webhook updates
type TokenWebhookGetResponseFilterFilterType string

const (
	TokenWebhookGetResponseFilterFilterTypeTokenAddress TokenWebhookGetResponseFilterFilterType = "token_address"
	TokenWebhookGetResponseFilterFilterTypeTokenType    TokenWebhookGetResponseFilterFilterType = "token_type"
)

func (r TokenWebhookGetResponseFilterFilterType) IsKnown() bool {
	switch r {
	case TokenWebhookGetResponseFilterFilterTypeTokenAddress, TokenWebhookGetResponseFilterFilterTypeTokenType:
		return true
	}
	return false
}

// Type of token to filter
type TokenWebhookGetResponseFilterTokenType string

const (
	TokenWebhookGetResponseFilterTokenTypeFungible    TokenWebhookGetResponseFilterTokenType = "Fungible"
	TokenWebhookGetResponseFilterTokenTypeNonFungible TokenWebhookGetResponseFilterTokenType = "NonFungible"
)

func (r TokenWebhookGetResponseFilterTokenType) IsKnown() bool {
	switch r {
	case TokenWebhookGetResponseFilterTokenTypeFungible, TokenWebhookGetResponseFilterTokenTypeNonFungible:
		return true
	}
	return false
}

type TokenWebhookGetAllResponse struct {
	Active bool `json:"active" api:"required"`
	// The chain name
	Chain     TokenScanSupportedChain `json:"chain" api:"required"`
	CreatedAt time.Time               `json:"created_at" api:"required" format:"date-time"`
	UpdatedAt time.Time               `json:"updated_at" api:"required" format:"date-time"`
	URL       string                  `json:"url" api:"required" format:"uri"`
	// Filter for webhook updates
	Filter TokenWebhookGetAllResponseFilter `json:"filter" api:"nullable"`
	// Optional shared secret key (32 characters), used to calculate the HMAC signature
	SharedSecretKey string                         `json:"shared_secret_key" api:"nullable"`
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
	Filter          apijson.Field
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

// Filter for webhook updates
type TokenWebhookGetAllResponseFilter struct {
	// Type of filter applied to the webhook updates
	FilterType TokenWebhookGetAllResponseFilterFilterType `json:"filter_type"`
	// This field can have the runtime type of [[]string].
	TokenAddresses interface{} `json:"token_addresses"`
	// Type of token to filter
	TokenType TokenWebhookGetAllResponseFilterTokenType `json:"token_type"`
	JSON      tokenWebhookGetAllResponseFilterJSON      `json:"-"`
	union     TokenWebhookGetAllResponseFilterUnion
}

// tokenWebhookGetAllResponseFilterJSON contains the JSON metadata for the struct
// [TokenWebhookGetAllResponseFilter]
type tokenWebhookGetAllResponseFilterJSON struct {
	FilterType     apijson.Field
	TokenAddresses apijson.Field
	TokenType      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r tokenWebhookGetAllResponseFilterJSON) RawJSON() string {
	return r.raw
}

func (r *TokenWebhookGetAllResponseFilter) UnmarshalJSON(data []byte) (err error) {
	*r = TokenWebhookGetAllResponseFilter{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TokenWebhookGetAllResponseFilterUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TokenWebhookGetAllResponseFilterTokenAddressFilter],
// [TokenWebhookGetAllResponseFilterTokenTypeFilter].
func (r TokenWebhookGetAllResponseFilter) AsUnion() TokenWebhookGetAllResponseFilterUnion {
	return r.union
}

// Filter for webhook updates
//
// Union satisfied by [TokenWebhookGetAllResponseFilterTokenAddressFilter] or
// [TokenWebhookGetAllResponseFilterTokenTypeFilter].
type TokenWebhookGetAllResponseFilterUnion interface {
	implementsTokenWebhookGetAllResponseFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TokenWebhookGetAllResponseFilterUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenWebhookGetAllResponseFilterTokenAddressFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenWebhookGetAllResponseFilterTokenTypeFilter{}),
		},
	)
}

type TokenWebhookGetAllResponseFilterTokenAddressFilter struct {
	// List of up to 100000 token addresses to filter webhook updates
	TokenAddresses []string `json:"token_addresses" api:"required"`
	// Type of filter applied to the webhook updates
	FilterType TokenWebhookGetAllResponseFilterTokenAddressFilterFilterType `json:"filter_type"`
	JSON       tokenWebhookGetAllResponseFilterTokenAddressFilterJSON       `json:"-"`
}

// tokenWebhookGetAllResponseFilterTokenAddressFilterJSON contains the JSON
// metadata for the struct [TokenWebhookGetAllResponseFilterTokenAddressFilter]
type tokenWebhookGetAllResponseFilterTokenAddressFilterJSON struct {
	TokenAddresses apijson.Field
	FilterType     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TokenWebhookGetAllResponseFilterTokenAddressFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenWebhookGetAllResponseFilterTokenAddressFilterJSON) RawJSON() string {
	return r.raw
}

func (r TokenWebhookGetAllResponseFilterTokenAddressFilter) implementsTokenWebhookGetAllResponseFilter() {
}

// Type of filter applied to the webhook updates
type TokenWebhookGetAllResponseFilterTokenAddressFilterFilterType string

const (
	TokenWebhookGetAllResponseFilterTokenAddressFilterFilterTypeTokenAddress TokenWebhookGetAllResponseFilterTokenAddressFilterFilterType = "token_address"
)

func (r TokenWebhookGetAllResponseFilterTokenAddressFilterFilterType) IsKnown() bool {
	switch r {
	case TokenWebhookGetAllResponseFilterTokenAddressFilterFilterTypeTokenAddress:
		return true
	}
	return false
}

type TokenWebhookGetAllResponseFilterTokenTypeFilter struct {
	// Type of token to filter
	TokenType TokenWebhookGetAllResponseFilterTokenTypeFilterTokenType `json:"token_type" api:"required"`
	// Type of filter applied to the webhook updates
	FilterType TokenWebhookGetAllResponseFilterTokenTypeFilterFilterType `json:"filter_type"`
	JSON       tokenWebhookGetAllResponseFilterTokenTypeFilterJSON       `json:"-"`
}

// tokenWebhookGetAllResponseFilterTokenTypeFilterJSON contains the JSON metadata
// for the struct [TokenWebhookGetAllResponseFilterTokenTypeFilter]
type tokenWebhookGetAllResponseFilterTokenTypeFilterJSON struct {
	TokenType   apijson.Field
	FilterType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenWebhookGetAllResponseFilterTokenTypeFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenWebhookGetAllResponseFilterTokenTypeFilterJSON) RawJSON() string {
	return r.raw
}

func (r TokenWebhookGetAllResponseFilterTokenTypeFilter) implementsTokenWebhookGetAllResponseFilter() {
}

// Type of token to filter
type TokenWebhookGetAllResponseFilterTokenTypeFilterTokenType string

const (
	TokenWebhookGetAllResponseFilterTokenTypeFilterTokenTypeFungible    TokenWebhookGetAllResponseFilterTokenTypeFilterTokenType = "Fungible"
	TokenWebhookGetAllResponseFilterTokenTypeFilterTokenTypeNonFungible TokenWebhookGetAllResponseFilterTokenTypeFilterTokenType = "NonFungible"
)

func (r TokenWebhookGetAllResponseFilterTokenTypeFilterTokenType) IsKnown() bool {
	switch r {
	case TokenWebhookGetAllResponseFilterTokenTypeFilterTokenTypeFungible, TokenWebhookGetAllResponseFilterTokenTypeFilterTokenTypeNonFungible:
		return true
	}
	return false
}

// Type of filter applied to the webhook updates
type TokenWebhookGetAllResponseFilterTokenTypeFilterFilterType string

const (
	TokenWebhookGetAllResponseFilterTokenTypeFilterFilterTypeTokenType TokenWebhookGetAllResponseFilterTokenTypeFilterFilterType = "token_type"
)

func (r TokenWebhookGetAllResponseFilterTokenTypeFilterFilterType) IsKnown() bool {
	switch r {
	case TokenWebhookGetAllResponseFilterTokenTypeFilterFilterTypeTokenType:
		return true
	}
	return false
}

// Type of filter applied to the webhook updates
type TokenWebhookGetAllResponseFilterFilterType string

const (
	TokenWebhookGetAllResponseFilterFilterTypeTokenAddress TokenWebhookGetAllResponseFilterFilterType = "token_address"
	TokenWebhookGetAllResponseFilterFilterTypeTokenType    TokenWebhookGetAllResponseFilterFilterType = "token_type"
)

func (r TokenWebhookGetAllResponseFilterFilterType) IsKnown() bool {
	switch r {
	case TokenWebhookGetAllResponseFilterFilterTypeTokenAddress, TokenWebhookGetAllResponseFilterFilterTypeTokenType:
		return true
	}
	return false
}

// Type of token to filter
type TokenWebhookGetAllResponseFilterTokenType string

const (
	TokenWebhookGetAllResponseFilterTokenTypeFungible    TokenWebhookGetAllResponseFilterTokenType = "Fungible"
	TokenWebhookGetAllResponseFilterTokenTypeNonFungible TokenWebhookGetAllResponseFilterTokenType = "NonFungible"
)

func (r TokenWebhookGetAllResponseFilterTokenType) IsKnown() bool {
	switch r {
	case TokenWebhookGetAllResponseFilterTokenTypeFungible, TokenWebhookGetAllResponseFilterTokenTypeNonFungible:
		return true
	}
	return false
}

type TokenWebhookNewParams struct {
	URL param.Field[string] `json:"url" api:"required" format:"uri"`
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
	TokenAddresses param.Field[[]string] `json:"token_addresses" api:"required"`
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
	TokenType param.Field[TokenWebhookNewParamsFilterTokenTypeFilterTokenType] `json:"token_type" api:"required"`
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

type TokenWebhookUpdateParams struct {
	// Token addresses to add to the hook's existing token_address filter, without
	// having to resend the full list. Mutually exclusive with filter.
	AddTokenAddresses param.Field[[]string] `json:"add_token_addresses"`
	// If provided, replaces the hook's filter in full. Mutually exclusive with
	// add_token_addresses/remove_token_addresses.
	Filter param.Field[TokenWebhookUpdateParamsFilterUnion] `json:"filter"`
	// Token addresses to remove from the hook's existing token_address filter.
	// Mutually exclusive with filter.
	RemoveTokenAddresses param.Field[[]string] `json:"remove_token_addresses"`
	// If provided, replaces the hook's shared secret key
	SharedSecretKey param.Field[string] `json:"shared_secret_key"`
	// If provided, replaces the hook's url
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r TokenWebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If provided, replaces the hook's filter in full. Mutually exclusive with
// add_token_addresses/remove_token_addresses.
type TokenWebhookUpdateParamsFilter struct {
	// Type of filter applied to the webhook updates
	FilterType     param.Field[TokenWebhookUpdateParamsFilterFilterType] `json:"filter_type"`
	TokenAddresses param.Field[interface{}]                              `json:"token_addresses"`
	// Type of token to filter
	TokenType param.Field[TokenWebhookUpdateParamsFilterTokenType] `json:"token_type"`
}

func (r TokenWebhookUpdateParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TokenWebhookUpdateParamsFilter) implementsTokenWebhookUpdateParamsFilterUnion() {}

// If provided, replaces the hook's filter in full. Mutually exclusive with
// add_token_addresses/remove_token_addresses.
//
// Satisfied by [TokenWebhookUpdateParamsFilterTokenAddressFilter],
// [TokenWebhookUpdateParamsFilterTokenTypeFilter],
// [TokenWebhookUpdateParamsFilter].
type TokenWebhookUpdateParamsFilterUnion interface {
	implementsTokenWebhookUpdateParamsFilterUnion()
}

type TokenWebhookUpdateParamsFilterTokenAddressFilter struct {
	// List of up to 100000 token addresses to filter webhook updates
	TokenAddresses param.Field[[]string] `json:"token_addresses" api:"required"`
	// Type of filter applied to the webhook updates
	FilterType param.Field[TokenWebhookUpdateParamsFilterTokenAddressFilterFilterType] `json:"filter_type"`
}

func (r TokenWebhookUpdateParamsFilterTokenAddressFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TokenWebhookUpdateParamsFilterTokenAddressFilter) implementsTokenWebhookUpdateParamsFilterUnion() {
}

// Type of filter applied to the webhook updates
type TokenWebhookUpdateParamsFilterTokenAddressFilterFilterType string

const (
	TokenWebhookUpdateParamsFilterTokenAddressFilterFilterTypeTokenAddress TokenWebhookUpdateParamsFilterTokenAddressFilterFilterType = "token_address"
)

func (r TokenWebhookUpdateParamsFilterTokenAddressFilterFilterType) IsKnown() bool {
	switch r {
	case TokenWebhookUpdateParamsFilterTokenAddressFilterFilterTypeTokenAddress:
		return true
	}
	return false
}

type TokenWebhookUpdateParamsFilterTokenTypeFilter struct {
	// Type of token to filter
	TokenType param.Field[TokenWebhookUpdateParamsFilterTokenTypeFilterTokenType] `json:"token_type" api:"required"`
	// Type of filter applied to the webhook updates
	FilterType param.Field[TokenWebhookUpdateParamsFilterTokenTypeFilterFilterType] `json:"filter_type"`
}

func (r TokenWebhookUpdateParamsFilterTokenTypeFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TokenWebhookUpdateParamsFilterTokenTypeFilter) implementsTokenWebhookUpdateParamsFilterUnion() {
}

// Type of token to filter
type TokenWebhookUpdateParamsFilterTokenTypeFilterTokenType string

const (
	TokenWebhookUpdateParamsFilterTokenTypeFilterTokenTypeFungible    TokenWebhookUpdateParamsFilterTokenTypeFilterTokenType = "Fungible"
	TokenWebhookUpdateParamsFilterTokenTypeFilterTokenTypeNonFungible TokenWebhookUpdateParamsFilterTokenTypeFilterTokenType = "NonFungible"
)

func (r TokenWebhookUpdateParamsFilterTokenTypeFilterTokenType) IsKnown() bool {
	switch r {
	case TokenWebhookUpdateParamsFilterTokenTypeFilterTokenTypeFungible, TokenWebhookUpdateParamsFilterTokenTypeFilterTokenTypeNonFungible:
		return true
	}
	return false
}

// Type of filter applied to the webhook updates
type TokenWebhookUpdateParamsFilterTokenTypeFilterFilterType string

const (
	TokenWebhookUpdateParamsFilterTokenTypeFilterFilterTypeTokenType TokenWebhookUpdateParamsFilterTokenTypeFilterFilterType = "token_type"
)

func (r TokenWebhookUpdateParamsFilterTokenTypeFilterFilterType) IsKnown() bool {
	switch r {
	case TokenWebhookUpdateParamsFilterTokenTypeFilterFilterTypeTokenType:
		return true
	}
	return false
}

// Type of filter applied to the webhook updates
type TokenWebhookUpdateParamsFilterFilterType string

const (
	TokenWebhookUpdateParamsFilterFilterTypeTokenAddress TokenWebhookUpdateParamsFilterFilterType = "token_address"
	TokenWebhookUpdateParamsFilterFilterTypeTokenType    TokenWebhookUpdateParamsFilterFilterType = "token_type"
)

func (r TokenWebhookUpdateParamsFilterFilterType) IsKnown() bool {
	switch r {
	case TokenWebhookUpdateParamsFilterFilterTypeTokenAddress, TokenWebhookUpdateParamsFilterFilterTypeTokenType:
		return true
	}
	return false
}

// Type of token to filter
type TokenWebhookUpdateParamsFilterTokenType string

const (
	TokenWebhookUpdateParamsFilterTokenTypeFungible    TokenWebhookUpdateParamsFilterTokenType = "Fungible"
	TokenWebhookUpdateParamsFilterTokenTypeNonFungible TokenWebhookUpdateParamsFilterTokenType = "NonFungible"
)

func (r TokenWebhookUpdateParamsFilterTokenType) IsKnown() bool {
	switch r {
	case TokenWebhookUpdateParamsFilterTokenTypeFungible, TokenWebhookUpdateParamsFilterTokenTypeNonFungible:
		return true
	}
	return false
}
