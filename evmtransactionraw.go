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

// EvmTransactionRawService contains methods and other services that help with
// interacting with the blockaid API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewEvmTransactionRawService] method
// instead.
type EvmTransactionRawService struct {
	Options []option.RequestOption
}

// NewEvmTransactionRawService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEvmTransactionRawService(opts ...option.RequestOption) (r *EvmTransactionRawService) {
	r = &EvmTransactionRawService{}
	r.Options = opts
	return
}

// Gets a raw transaction and returns a full simulation indicating what will happen
// in the transaction together with a recommended action and some textual reasons
// of why the transaction was flagged that way.
func (r *EvmTransactionRawService) Scan(ctx context.Context, body EvmTransactionRawScanParams, opts ...option.RequestOption) (res *TransactionScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "evm/transaction-raw/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EvmTransactionRawScanParams struct {
	// The address to relate the transaction to. Account address determines in which
	// perspective the transaction is simulated and validated.
	AccountAddress param.Field[string] `json:"account_address,required"`
	// Hex string of the raw transaction data
	Data param.Field[string] `json:"data,required"`
	// Object of additional information to validate against.
	Metadata param.Field[MetadataParam] `json:"metadata,required"`
	// An enumeration.
	Chain param.Field[EvmTransactionRawScanParamsChain] `json:"chain"`
	// List of one or both of options for the desired output. "simulation" - include
	// simulation output in your response. "validation" - include security validation
	// of the transaction in your response. Default is ["validation"]
	Options param.Field[[]EvmTransactionRawScanParamsOption] `json:"options"`
}

func (r EvmTransactionRawScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An enumeration.
type EvmTransactionRawScanParamsChain string

const (
	EvmTransactionRawScanParamsChainArbitrum    EvmTransactionRawScanParamsChain = "arbitrum"
	EvmTransactionRawScanParamsChainAvalanche   EvmTransactionRawScanParamsChain = "avalanche"
	EvmTransactionRawScanParamsChainBase        EvmTransactionRawScanParamsChain = "base"
	EvmTransactionRawScanParamsChainBaseSepolia EvmTransactionRawScanParamsChain = "base-sepolia"
	EvmTransactionRawScanParamsChainBsc         EvmTransactionRawScanParamsChain = "bsc"
	EvmTransactionRawScanParamsChainEthereum    EvmTransactionRawScanParamsChain = "ethereum"
	EvmTransactionRawScanParamsChainOptimism    EvmTransactionRawScanParamsChain = "optimism"
	EvmTransactionRawScanParamsChainPolygon     EvmTransactionRawScanParamsChain = "polygon"
	EvmTransactionRawScanParamsChainZksync      EvmTransactionRawScanParamsChain = "zksync"
	EvmTransactionRawScanParamsChainZora        EvmTransactionRawScanParamsChain = "zora"
	EvmTransactionRawScanParamsChainLinea       EvmTransactionRawScanParamsChain = "linea"
	EvmTransactionRawScanParamsChainBlast       EvmTransactionRawScanParamsChain = "blast"
	EvmTransactionRawScanParamsChainUnknown     EvmTransactionRawScanParamsChain = "unknown"
)

func (r EvmTransactionRawScanParamsChain) IsKnown() bool {
	switch r {
	case EvmTransactionRawScanParamsChainArbitrum, EvmTransactionRawScanParamsChainAvalanche, EvmTransactionRawScanParamsChainBase, EvmTransactionRawScanParamsChainBaseSepolia, EvmTransactionRawScanParamsChainBsc, EvmTransactionRawScanParamsChainEthereum, EvmTransactionRawScanParamsChainOptimism, EvmTransactionRawScanParamsChainPolygon, EvmTransactionRawScanParamsChainZksync, EvmTransactionRawScanParamsChainZora, EvmTransactionRawScanParamsChainLinea, EvmTransactionRawScanParamsChainBlast, EvmTransactionRawScanParamsChainUnknown:
		return true
	}
	return false
}

// An enumeration.
type EvmTransactionRawScanParamsOption string

const (
	EvmTransactionRawScanParamsOptionValidation EvmTransactionRawScanParamsOption = "validation"
	EvmTransactionRawScanParamsOptionSimulation EvmTransactionRawScanParamsOption = "simulation"
)

func (r EvmTransactionRawScanParamsOption) IsKnown() bool {
	switch r {
	case EvmTransactionRawScanParamsOptionValidation, EvmTransactionRawScanParamsOptionSimulation:
		return true
	}
	return false
}