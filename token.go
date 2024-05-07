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

// TokenService contains methods and other services that help with interacting with
// the blockaid API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewTokenService] method instead.
type TokenService struct {
	Options []option.RequestOption
}

// NewTokenService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTokenService(opts ...option.RequestOption) (r *TokenService) {
	r = &TokenService{}
	r.Options = opts
	return
}

// Gets a token address and scan the token to identify any indication of malicious
// behaviour
func (r *TokenService) Scan(ctx context.Context, body TokenScanParams, opts ...option.RequestOption) (res *TokenScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/token/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type TokenScanResponse struct {
	// Dictionary of detected attacks found during the scan
	AttackTypes map[string]TokenScanResponseAttackType `json:"attack_types,required"`
	// Score between 0 to 1 (double)
	MaliciousScore string `json:"malicious_score,required"`
	// An enumeration.
	ResultType TokenScanResponseResultType `json:"result_type,required"`
	JSON       tokenScanResponseJSON       `json:"-"`
}

// tokenScanResponseJSON contains the JSON metadata for the struct
// [TokenScanResponse]
type tokenScanResponseJSON struct {
	AttackTypes    apijson.Field
	MaliciousScore apijson.Field
	ResultType     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TokenScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseJSON) RawJSON() string {
	return r.raw
}

type TokenScanResponseAttackType struct {
	// Score between 0 to 1 (double) that indicates the assurance this attack happened
	Score string `json:"score,required"`
	// Object contains an extra information related to the attack
	Features interface{} `json:"features"`
	// If score is higher or equal to this field, the token is using this attack type
	Threshold string                          `json:"threshold"`
	JSON      tokenScanResponseAttackTypeJSON `json:"-"`
}

// tokenScanResponseAttackTypeJSON contains the JSON metadata for the struct
// [TokenScanResponseAttackType]
type tokenScanResponseAttackTypeJSON struct {
	Score       apijson.Field
	Features    apijson.Field
	Threshold   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenScanResponseAttackType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenScanResponseAttackTypeJSON) RawJSON() string {
	return r.raw
}

// An enumeration.
type TokenScanResponseResultType string

const (
	TokenScanResponseResultTypeBenign    TokenScanResponseResultType = "Benign"
	TokenScanResponseResultTypeWarning   TokenScanResponseResultType = "Warning"
	TokenScanResponseResultTypeMalicious TokenScanResponseResultType = "Malicious"
	TokenScanResponseResultTypeSpam      TokenScanResponseResultType = "Spam"
)

func (r TokenScanResponseResultType) IsKnown() bool {
	switch r {
	case TokenScanResponseResultTypeBenign, TokenScanResponseResultTypeWarning, TokenScanResponseResultTypeMalicious, TokenScanResponseResultTypeSpam:
		return true
	}
	return false
}

type TokenScanParams struct {
	// Token address to validate (EVM / Solana)
	Address param.Field[TokenScanParamsAddressUnion] `json:"address,required"`
	// The chain name
	Chain param.Field[TokenScanSupportedChain] `json:"chain,required"`
	// Object of additional information to validate against.
	Metadata param.Field[TokenScanParamsMetadata] `json:"metadata"`
	// (optional) Token ID for ERC721 or ERC1155
	TokenID param.Field[int64] `json:"token_id"`
}

func (r TokenScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Token address to validate (EVM / Solana)
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type TokenScanParamsAddressUnion interface {
	ImplementsTokenScanParamsAddressUnion()
}

// Object of additional information to validate against.
type TokenScanParamsMetadata struct {
	// cross reference transaction against the domain.
	Domain param.Field[string] `json:"domain"`
}

func (r TokenScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
