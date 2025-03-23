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
	Results interface{}               `json:"results,required"`
	JSON    tokenBulkScanResponseJSON `json:"-"`
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

type TokenBulkScanParams struct {
	// The chain name
	Chain param.Field[TokenBulkScanParamsChain] `json:"chain,required"`
	// A list of token addresses to scan
	Tokens param.Field[[]string] `json:"tokens,required"`
	// Object of additional information to validate against.
	Metadata param.Field[TokenBulkScanParamsMetadata] `json:"metadata"`
}

func (r TokenBulkScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The chain name
type TokenBulkScanParamsChain string

const (
	TokenBulkScanParamsChainArbitrum    TokenBulkScanParamsChain = "arbitrum"
	TokenBulkScanParamsChainAvalanche   TokenBulkScanParamsChain = "avalanche"
	TokenBulkScanParamsChainBase        TokenBulkScanParamsChain = "base"
	TokenBulkScanParamsChainBsc         TokenBulkScanParamsChain = "bsc"
	TokenBulkScanParamsChainEthereum    TokenBulkScanParamsChain = "ethereum"
	TokenBulkScanParamsChainOptimism    TokenBulkScanParamsChain = "optimism"
	TokenBulkScanParamsChainPolygon     TokenBulkScanParamsChain = "polygon"
	TokenBulkScanParamsChainZora        TokenBulkScanParamsChain = "zora"
	TokenBulkScanParamsChainSolana      TokenBulkScanParamsChain = "solana"
	TokenBulkScanParamsChainStarknet    TokenBulkScanParamsChain = "starknet"
	TokenBulkScanParamsChainStellar     TokenBulkScanParamsChain = "stellar"
	TokenBulkScanParamsChainLinea       TokenBulkScanParamsChain = "linea"
	TokenBulkScanParamsChainBlast       TokenBulkScanParamsChain = "blast"
	TokenBulkScanParamsChainZksync      TokenBulkScanParamsChain = "zksync"
	TokenBulkScanParamsChainScroll      TokenBulkScanParamsChain = "scroll"
	TokenBulkScanParamsChainDegen       TokenBulkScanParamsChain = "degen"
	TokenBulkScanParamsChainAbstract    TokenBulkScanParamsChain = "abstract"
	TokenBulkScanParamsChainSoneium     TokenBulkScanParamsChain = "soneium"
	TokenBulkScanParamsChainInk         TokenBulkScanParamsChain = "ink"
	TokenBulkScanParamsChainZeroNetwork TokenBulkScanParamsChain = "zero-network"
	TokenBulkScanParamsChainBerachain   TokenBulkScanParamsChain = "berachain"
	TokenBulkScanParamsChainUnichain    TokenBulkScanParamsChain = "unichain"
)

func (r TokenBulkScanParamsChain) IsKnown() bool {
	switch r {
	case TokenBulkScanParamsChainArbitrum, TokenBulkScanParamsChainAvalanche, TokenBulkScanParamsChainBase, TokenBulkScanParamsChainBsc, TokenBulkScanParamsChainEthereum, TokenBulkScanParamsChainOptimism, TokenBulkScanParamsChainPolygon, TokenBulkScanParamsChainZora, TokenBulkScanParamsChainSolana, TokenBulkScanParamsChainStarknet, TokenBulkScanParamsChainStellar, TokenBulkScanParamsChainLinea, TokenBulkScanParamsChainBlast, TokenBulkScanParamsChainZksync, TokenBulkScanParamsChainScroll, TokenBulkScanParamsChainDegen, TokenBulkScanParamsChainAbstract, TokenBulkScanParamsChainSoneium, TokenBulkScanParamsChainInk, TokenBulkScanParamsChainZeroNetwork, TokenBulkScanParamsChainBerachain, TokenBulkScanParamsChainUnichain:
		return true
	}
	return false
}

// Object of additional information to validate against.
type TokenBulkScanParamsMetadata struct {
	// cross reference transaction against the domain.
	Domain param.Field[string] `json:"domain"`
}

func (r TokenBulkScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
