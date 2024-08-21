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

// TokenBulkService contains methods and other services that help with interacting
// with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTokenBulkService] method instead.
type TokenBulkService struct {
	Options []option.RequestOption
}

// NewTokenBulkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTokenBulkService(opts ...option.RequestOption) (r *TokenBulkService) {
	r = &TokenBulkService{}
	r.Options = opts
	return
}

// Gets a list of token addresses and scan the tokens to identify any indication of
// malicious behaviour
func (r *TokenBulkService) Scan(ctx context.Context, body TokenBulkScanParams, opts ...option.RequestOption) (res *TokenBulkScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/token-bulk/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type TokenBulkScanResponse struct {
	Results map[string]TokenBulkScanResponseResult `json:"results,required"`
	JSON    tokenBulkScanResponseJSON              `json:"-"`
}

// tokenBulkScanResponseJSON contains the JSON metadata for the struct
// [TokenBulkScanResponse]
type tokenBulkScanResponseJSON struct {
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenBulkScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseJSON) RawJSON() string {
	return r.raw
}

type TokenBulkScanResponseResult struct {
	// Dictionary of detected attacks found during the scan
	AttackTypes map[string]TokenBulkScanResponseResultsAttackType `json:"attack_types,required"`
	// Score between 0 to 1 (double)
	MaliciousScore string `json:"malicious_score,required"`
	// An enumeration.
	ResultType TokenBulkScanResponseResultsResultType `json:"result_type,required"`
	JSON       tokenBulkScanResponseResultJSON        `json:"-"`
}

// tokenBulkScanResponseResultJSON contains the JSON metadata for the struct
// [TokenBulkScanResponseResult]
type tokenBulkScanResponseResultJSON struct {
	AttackTypes    apijson.Field
	MaliciousScore apijson.Field
	ResultType     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TokenBulkScanResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultJSON) RawJSON() string {
	return r.raw
}

type TokenBulkScanResponseResultsAttackType struct {
	// Score between 0 to 1 (double) that indicates the assurance this attack happened
	Score string `json:"score,required"`
	// Object contains an extra information related to the attack
	Features interface{} `json:"features"`
	// If score is higher or equal to this field, the token is using this attack type
	Threshold string                                     `json:"threshold"`
	JSON      tokenBulkScanResponseResultsAttackTypeJSON `json:"-"`
}

// tokenBulkScanResponseResultsAttackTypeJSON contains the JSON metadata for the
// struct [TokenBulkScanResponseResultsAttackType]
type tokenBulkScanResponseResultsAttackTypeJSON struct {
	Score       apijson.Field
	Features    apijson.Field
	Threshold   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenBulkScanResponseResultsAttackType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenBulkScanResponseResultsAttackTypeJSON) RawJSON() string {
	return r.raw
}

// An enumeration.
type TokenBulkScanResponseResultsResultType string

const (
	TokenBulkScanResponseResultsResultTypeBenign    TokenBulkScanResponseResultsResultType = "Benign"
	TokenBulkScanResponseResultsResultTypeWarning   TokenBulkScanResponseResultsResultType = "Warning"
	TokenBulkScanResponseResultsResultTypeMalicious TokenBulkScanResponseResultsResultType = "Malicious"
	TokenBulkScanResponseResultsResultTypeSpam      TokenBulkScanResponseResultsResultType = "Spam"
)

func (r TokenBulkScanResponseResultsResultType) IsKnown() bool {
	switch r {
	case TokenBulkScanResponseResultsResultTypeBenign, TokenBulkScanResponseResultsResultTypeWarning, TokenBulkScanResponseResultsResultTypeMalicious, TokenBulkScanResponseResultsResultTypeSpam:
		return true
	}
	return false
}

type TokenBulkScanParams struct {
	// The chain name
	Chain param.Field[TokenScanSupportedChain] `json:"chain,required"`
	// A list of token addresses to scan
	Tokens param.Field[[]string] `json:"tokens,required"`
	// Object of additional information to validate against.
	Metadata param.Field[TokenBulkScanParamsMetadata] `json:"metadata"`
}

func (r TokenBulkScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Object of additional information to validate against.
type TokenBulkScanParamsMetadata struct {
	// cross reference transaction against the domain.
	Domain param.Field[string] `json:"domain"`
}

func (r TokenBulkScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
