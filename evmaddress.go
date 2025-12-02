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
func (r *EvmAddressService) Report(ctx context.Context, body EvmAddressReportParams, opts ...option.RequestOption) (res *EvmAddressReportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/evm/address/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets an address and returns a full security assessment indicating weather or not
// this address is malicious as well as textual reasons of why the address was
// flagged that way.
func (r *EvmAddressService) Scan(ctx context.Context, body EvmAddressScanParams, opts ...option.RequestOption) (res *AddressValidation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/evm/address/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EvmAddressReportResponse = interface{}

type EvmAddressReportParams struct {
	// Details about the report.
	Details param.Field[string] `json:"details,required"`
	// The event type of the report. Could be FALSE_POSITIVE or FALSE_NEGATIVE.
	Event param.Field[EvmAddressReportParamsEvent] `json:"event,required"`
	// The report parameters.
	Report param.Field[EvmAddressReportParamsReportUnion] `json:"report,required"`
}

func (r EvmAddressReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The event type of the report. Could be FALSE_POSITIVE or FALSE_NEGATIVE.
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

// The report parameters.
type EvmAddressReportParamsReport struct {
	Type      param.Field[EvmAddressReportParamsReportType] `json:"type,required"`
	Params    param.Field[interface{}]                      `json:"params"`
	RequestID param.Field[string]                           `json:"request_id"`
}

func (r EvmAddressReportParamsReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmAddressReportParamsReport) implementsEvmAddressReportParamsReportUnion() {}

// The report parameters.
//
// Satisfied by [EvmAddressReportParamsReportParamReportEvmAddressReportParams],
// [EvmAddressReportParamsReportRequestIDReport], [EvmAddressReportParamsReport].
type EvmAddressReportParamsReportUnion interface {
	implementsEvmAddressReportParamsReportUnion()
}

type EvmAddressReportParamsReportParamReportEvmAddressReportParams struct {
	Params param.Field[EvmAddressReportParamsReportParamReportEvmAddressReportParamsParams] `json:"params,required"`
	Type   param.Field[EvmAddressReportParamsReportParamReportEvmAddressReportParamsType]   `json:"type,required"`
}

func (r EvmAddressReportParamsReportParamReportEvmAddressReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmAddressReportParamsReportParamReportEvmAddressReportParams) implementsEvmAddressReportParamsReportUnion() {
}

type EvmAddressReportParamsReportParamReportEvmAddressReportParamsParams struct {
	// The address to report on.
	Address param.Field[string] `json:"address,required"`
	// The chain name
	Chain param.Field[EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain] `json:"chain,required"`
	// The domain related to this address.
	Domain param.Field[string] `json:"domain,required"`
}

func (r EvmAddressReportParamsReportParamReportEvmAddressReportParamsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The chain name
type EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain string

const (
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainArbitrum              EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "arbitrum"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainAvalanche             EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "avalanche"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainBase                  EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "base"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainBaseSepolia           EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "base-sepolia"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainLordchain             EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "lordchain"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainLordchainTestnet      EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "lordchain-testnet"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainMetacade              EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "metacade"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainMetacadeTestnet       EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "metacade-testnet"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainBsc                   EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "bsc"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainEthereum              EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "ethereum"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainOptimism              EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "optimism"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainPolygon               EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "polygon"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainZksync                EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "zksync"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainZksyncSepolia         EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "zksync-sepolia"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainZora                  EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "zora"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainLinea                 EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "linea"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainBlast                 EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "blast"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainScroll                EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "scroll"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainEthereumSepolia       EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "ethereum-sepolia"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainDegen                 EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "degen"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainAvalancheFuji         EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "avalanche-fuji"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainImmutableZkevm        EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "immutable-zkevm"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainImmutableZkevmTestnet EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "immutable-zkevm-testnet"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainGnosis                EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "gnosis"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainWorldchain            EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "worldchain"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainSoneiumMinato         EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "soneium-minato"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainRonin                 EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "ronin"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainApechain              EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "apechain"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainZeroNetwork           EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "zero-network"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainBerachain             EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "berachain"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainBerachainBartio       EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "berachain-bartio"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainInk                   EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "ink"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainInkSepolia            EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "ink-sepolia"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainAbstract              EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "abstract"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainAbstractTestnet       EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "abstract-testnet"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainSoneium               EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "soneium"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainUnichain              EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "unichain"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainSei                   EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "sei"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainFlowEvm               EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "flow-evm"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainHyperevm              EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "hyperevm"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainKatana                EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "katana"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainPlume                 EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "plume"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainXlayer                EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "xlayer"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainMonad                 EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "monad"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainMonadTestnet          EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "monad-testnet"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainHedera                EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "hedera"
	EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainTempoTestnet          EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain = "tempo-testnet"
)

func (r EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChain) IsKnown() bool {
	switch r {
	case EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainArbitrum, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainAvalanche, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainBase, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainBaseSepolia, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainLordchain, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainLordchainTestnet, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainMetacade, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainMetacadeTestnet, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainBsc, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainEthereum, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainOptimism, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainPolygon, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainZksync, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainZksyncSepolia, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainZora, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainLinea, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainBlast, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainScroll, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainEthereumSepolia, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainDegen, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainAvalancheFuji, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainImmutableZkevm, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainImmutableZkevmTestnet, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainGnosis, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainWorldchain, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainSoneiumMinato, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainRonin, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainApechain, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainZeroNetwork, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainBerachain, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainBerachainBartio, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainInk, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainInkSepolia, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainAbstract, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainAbstractTestnet, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainSoneium, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainUnichain, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainSei, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainFlowEvm, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainHyperevm, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainKatana, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainPlume, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainXlayer, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainMonad, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainMonadTestnet, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainHedera, EvmAddressReportParamsReportParamReportEvmAddressReportParamsParamsChainTempoTestnet:
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
	RequestID param.Field[string]                                          `json:"request_id,required"`
	Type      param.Field[EvmAddressReportParamsReportRequestIDReportType] `json:"type,required"`
}

func (r EvmAddressReportParamsReportRequestIDReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvmAddressReportParamsReportRequestIDReport) implementsEvmAddressReportParamsReportUnion() {}

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
	ValidateAddress ValidateAddressParam `json:"ValidateAddress,required"`
}

func (r EvmAddressScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ValidateAddress)
}
