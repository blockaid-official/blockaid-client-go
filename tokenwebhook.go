// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"fmt"
	"net/http"
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

// Create a new webhook for a given chain. The webhook will receive real-time token
// scan updates.
//
// ## Webhook Endpoint Requirements
//
// To ensure proper functionality, your webhook endpoint must meet the following
// requirements:
//
// ### ✅ Response Criteria
//
//   - Must return an **HTTP 200 OK** status code upon successful receipt.
//   - Must respond within **1 second**.
//   - If the endpoint is unreachable for more than **1 hour**, the webhook will be
//     **automatically disabled**.
//
// ### 🔄 Synchronization Process
//
// 1. **Subscribe to the webhook** to start receiving updates.
// 2. **Receive real-time updates** for new token scans
//
// ### Webhook Request Format
//
// The webhook will send `POST` requests containing an array of token scan results:
//
// ```json
// List[TokenValidationResponse]
// ```
//
// For detailed response structure, see the
// [Token Scan Response Reference](https://docs.blockaid.io/reference/token-scan-response-reference).
//
// ### Note
//
//   - The same token will have multiple scan results over time
//   - Ensure that your system properly handles state overrides to reflect the most
//     up-to-date token status.
func (r *TokenWebhookService) New(ctx context.Context, chain TokenScanSupportedChain, body TokenWebhookNewParams, opts ...option.RequestOption) (res *TokenWebhookNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v0/token/hooks/%v", chain)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete an existing webhook subscription for a given chain.
//
// This will immediately stop sending token scan updates to the webhook URL.
// Returns a 204 status code on successful deletion.
func (r *TokenWebhookService) Delete(ctx context.Context, chain TokenScanSupportedChain, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("v0/token/hooks/%v", chain)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get information about an existing webhook for a given chain
func (r *TokenWebhookService) Get(ctx context.Context, chain TokenScanSupportedChain, opts ...option.RequestOption) (res *TokenWebhookGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v0/token/hooks/%v", chain)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all active webhook subscriptions across all chains
func (r *TokenWebhookService) GetAll(ctx context.Context, opts ...option.RequestOption) (res *[]TokenWebhookGetAllResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/token/hooks/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TokenWebhookNewResponse struct {
	Active bool `json:"active,required"`
	// The chain name
	Chain     TokenScanSupportedChain     `json:"chain,required"`
	CreatedAt time.Time                   `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time                   `json:"updated_at,required" format:"date-time"`
	URL       string                      `json:"url,required" format:"uri"`
	JSON      tokenWebhookNewResponseJSON `json:"-"`
}

// tokenWebhookNewResponseJSON contains the JSON metadata for the struct
// [TokenWebhookNewResponse]
type tokenWebhookNewResponseJSON struct {
	Active      apijson.Field
	Chain       apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
	Chain     TokenScanSupportedChain     `json:"chain,required"`
	CreatedAt time.Time                   `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time                   `json:"updated_at,required" format:"date-time"`
	URL       string                      `json:"url,required" format:"uri"`
	JSON      tokenWebhookGetResponseJSON `json:"-"`
}

// tokenWebhookGetResponseJSON contains the JSON metadata for the struct
// [TokenWebhookGetResponse]
type tokenWebhookGetResponseJSON struct {
	Active      apijson.Field
	Chain       apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
	Chain     TokenScanSupportedChain        `json:"chain,required"`
	CreatedAt time.Time                      `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time                      `json:"updated_at,required" format:"date-time"`
	URL       string                         `json:"url,required" format:"uri"`
	JSON      tokenWebhookGetAllResponseJSON `json:"-"`
}

// tokenWebhookGetAllResponseJSON contains the JSON metadata for the struct
// [TokenWebhookGetAllResponse]
type tokenWebhookGetAllResponseJSON struct {
	Active      apijson.Field
	Chain       apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenWebhookGetAllResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenWebhookGetAllResponseJSON) RawJSON() string {
	return r.raw
}

type TokenWebhookNewParams struct {
	URL param.Field[string] `json:"url,required" format:"uri"`
}

func (r TokenWebhookNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
