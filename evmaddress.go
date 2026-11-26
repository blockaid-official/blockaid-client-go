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

// EvmAddressService contains methods and other services that help with interacting
// with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvmAddressService] method instead.
type EvmAddressService struct {
	Options []option.RequestOption
}

// NewEvmAddressService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEvmAddressService(opts ...option.RequestOption) (r *EvmAddressService) {
	r = &EvmAddressService{}
	r.Options = opts
	return
}

// Report a misclassification of an EVM address.
//
// Deprecated: deprecated
func (r *EvmAddressService) Report(ctx context.Context, body EvmAddressReportParams, opts ...option.RequestOption) (res *EvmAddressReportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/evm/address/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a risk recommendation with plain-language reasons for an address.
func (r *EvmAddressService) Scan(ctx context.Context, body EvmAddressScanParams, opts ...option.RequestOption) (res *AddressValidation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/evm/address/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type EvmAddressReportResponse = interface{}

type EvmAddressReportParams struct {
	// Details about the report.
	Details param.Field[string] `json:"details" api:"required"`
	// The event type of the report. Could be `FALSE_POSITIVE` or `FALSE_NEGATIVE`.
	Event param.Field[EvmAddressReportParamsEvent] `json:"event" api:"required"`
	// Parameters identifying the address to report, provided either as address details
	// (address, domain, and chain) or as a request ID from a previous scan.
	Report param.Field[EvmAddressReportParamsReportUnion] `json:"report" api:"required"`
}

func (r EvmAddressReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The event type of the report. Could be `FALSE_POSITIVE` or `FALSE_NEGATIVE`.
type EvmAddressReportParamsEvent string

const (
	EvmAddressReportParamsEventFalsePositive EvmAddressReportParamsEvent = "FALSE_POSITIVE"
	EvmAddressReportParamsEventFalseNegative EvmAddressReportParamsEvent = "FALSE_NEGATIVE"
)

func (r EvmAddressReportParamsEvent) IsKnown() bool {
	switch r {
	case EvmAddressReportParamsEventFalsePositive, EvmAddressReportParamsEventFalseNegative:
		return true
	}
	return false
}

// Parameters identifying the address to report, provided either as address details
// (address, domain, and chain) or as a request ID from a previous scan.
type EvmAddressReportParamsReport struct {
	Type   param.Field[EvmAddressReportParamsReportType] `json:"type" api:"required"`
	Params param.Field[interface{}]                      `json:"params"`
	// The request ID of a previous request. This can be found in the value of the
	// `x-request-id` field in the headers of the response of the previous request. For
	// instance: `6c3cf6c1-a80d-4927-91b9-03d841ea61fe`.
	RequestID param.Field[string] `json:"request_id"`
}

func (r EvmAddressReportParamsReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmAddressReportParamsReport) implementsEvmAddressReportParamsReportUnion() {}

// Parameters identifying the address to report, provided either as address details
// (address, domain, and chain) or as a request ID from a previous scan.
//
// Satisfied by [EvmAddressReportParamsReportParamReportEvmAddressReportParams],
// [EvmAddressReportParamsReportRequestIDReport], [EvmAddressReportParamsReport].
type EvmAddressReportParamsReportUnion interface {
	implementsEvmAddressReportParamsReportUnion()
}

type EvmAddressReportParamsReportParamReportEvmAddressReportParams struct {
	Params param.Field[EvmAddressReportParamsReportParamReportEvmAddressReportParamsParams] `json:"params" api:"required"`
	Type   param.Field[EvmAddressReportParamsReportParamReportEvmAddressReportParamsType]   `json:"type" api:"required"`
}

func (r EvmAddressReportParamsReportParamReportEvmAddressReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmAddressReportParamsReportParamReportEvmAddressReportParams) implementsEvmAddressReportParamsReportUnion() {
}

type EvmAddressReportParamsReportParamReportEvmAddressReportParamsParams struct {
	// The address to report on.
	Address param.Field[string] `json:"address" api:"required"`
	// The chain name
	Chain param.Field[EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain] `json:"chain" api:"required"`
	// The domain related to this address.
	Domain param.Field[string] `json:"domain" api:"required"`
}

func (r EvmAddressReportParamsReportParamReportEvmAddressReportParamsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The chain name
type EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain string

const (
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainAbstract        EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "abstract"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainAbstractTestnet EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "abstract-testnet"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainApechain        EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "apechain"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainArbitrum        EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "arbitrum"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainArc             EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "arc"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainAvalanche       EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "avalanche"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainAvalancheFuji   EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "avalanche-fuji"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainBase            EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "base"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainBaseSepolia     EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "base-sepolia"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainBerachain       EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "berachain"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainBsc             EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "bsc"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainEthereum        EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "ethereum"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainEthereumSepolia EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "ethereum-sepolia"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainFlare           EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "flare"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainFlowEvm         EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "flow-evm"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainGiwaSepolia     EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "giwa-sepolia"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainGnosis          EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "gnosis"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainHedera          EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "hedera"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainHyperevm        EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "hyperevm"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainInk             EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "ink"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainInkSepolia      EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "ink-sepolia"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainKaia            EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "kaia"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainKatana          EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "katana"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainKiteAI          EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "kite-ai"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainLinea           EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "linea"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainMantle          EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "mantle"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainMegaeth         EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "megaeth"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainMonad           EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "monad"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainMonadTestnet    EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "monad-testnet"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainOptimism        EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "optimism"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainPlasma          EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "plasma"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainPlume           EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "plume"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainPolygon         EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "polygon"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainRobinhood       EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "robinhood"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainRonin           EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "ronin"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainScroll          EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "scroll"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainSei             EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "sei"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainSoneium         EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "soneium"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainSoneiumMinato   EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "soneium-minato"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainTempo           EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "tempo"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainTempoTestnet    EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "tempo-testnet"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainTron            EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "tron"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainUnichain        EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "unichain"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainWorldchain      EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "worldchain"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainXdc             EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "xdc"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainXlayer          EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "xlayer"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainZksync          EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "zksync"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainZksyncSepolia   EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "zksync-sepolia"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainZora            EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "zora"
)

func (r EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain) IsKnown() bool {
	switch r {
	case EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainAbstract, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainAbstractTestnet, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainApechain, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainArbitrum, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainArc, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainAvalanche, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainAvalancheFuji, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainBase, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainBaseSepolia, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainBerachain, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainBsc, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainEthereum, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainEthereumSepolia, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainFlare, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainFlowEvm, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainGiwaSepolia, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainGnosis, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainHedera, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainHyperevm, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainInk, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainInkSepolia, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainKaia, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainKatana, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainKiteAI, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainLinea, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainMantle, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainMegaeth, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainMonad, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainMonadTestnet, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainOptimism, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainPlasma, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainPlume, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainPolygon, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainRobinhood, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainRonin, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainScroll, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainSei, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainSoneium, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainSoneiumMinato, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainTempo, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainTempoTestnet, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainTron, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainUnichain, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainWorldchain, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainXdc, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainXlayer, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainZksync, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainZksyncSepolia, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainZora:
		return true
	}
	return false
}

type EvmAddressReportParamsReportParamReportEvmAddressReportParamsType string

const (
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsTypeParams EvmAddressReportParamsReportParamReportEvmAddressReportParamsType = "params"
)

func (r EvmAddressReportParamsReportParamReportEvmAddressReportParamsType) IsKnown() bool {
	switch r {
	case EvmAddressReportParamsReportParamReportEvmAddressReportParamsTypeParams:
		return true
	}
	return false
}

type EvmAddressReportParamsReportRequestIDReport struct {
	// The request ID of a previous request. This can be found in the value of the
	// `x-request-id` field in the headers of the response of the previous request. For
	// instance: `6c3cf6c1-a80d-4927-91b9-03d841ea61fe`.
	RequestID param.Field[string] `json:"request_id" api:"required"`
	// The type identifier indicating that a request ID from a previous scan is being
	// used.
	Type param.Field[EvmAddressReportParamsReportRequestIDReportType] `json:"type" api:"required"`
}

func (r EvmAddressReportParamsReportRequestIDReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmAddressReportParamsReportRequestIDReport) implementsEvmAddressReportParamsReportUnion() {}

// The type identifier indicating that a request ID from a previous scan is being
// used.
type EvmAddressReportParamsReportRequestIDReportType string

const (
	EvmAddressReportParamsReportRequestIDReportTypeRequestID EvmAddressReportParamsReportRequestIDReportType = "request_id"
)

func (r EvmAddressReportParamsReportRequestIDReportType) IsKnown() bool {
	switch r {
	case EvmAddressReportParamsReportRequestIDReportTypeRequestID:
		return true
	}
	return false
}

type EvmAddressReportParamsReportType string

const (
	EvmAddressReportParamsReportTypeParams    EvmAddressReportParamsReportType = "params"
	EvmAddressReportParamsReportTypeRequestID EvmAddressReportParamsReportType = "request_id"
)

func (r EvmAddressReportParamsReportType) IsKnown() bool {
	switch r {
	case EvmAddressReportParamsReportTypeParams, EvmAddressReportParamsReportTypeRequestID:
		return true
	}
	return false
}

type EvmAddressScanParams struct {
	ValidateAddress ValidateAddressParam `json:"ValidateAddress" api:"required"`
}

func (r EvmAddressScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ValidateAddress)
}
