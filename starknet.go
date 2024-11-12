// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/option"
)

// StarknetService contains methods and other services that help with interacting
// with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStarknetService] method instead.
type StarknetService struct {
	Options     []option.RequestOption
	Transaction *StarknetTransactionService
}

// NewStarknetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStarknetService(opts ...option.RequestOption) (r *StarknetService) {
	r = &StarknetService{}
	r.Options = opts
	r.Transaction = NewStarknetTransactionService(opts...)
	return
}

type StarknetErc1155Diff struct {
	// Token ID of the transfer
	TokenID string `json:"token_id,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value int64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string                  `json:"summary,nullable"`
	JSON    starknetErc1155DiffJSON `json:"-"`
}

// starknetErc1155DiffJSON contains the JSON metadata for the struct
// [StarknetErc1155Diff]
type starknetErc1155DiffJSON struct {
	TokenID     apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetErc1155Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetErc1155DiffJSON) RawJSON() string {
	return r.raw
}

type StarknetErc20Diff struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Value of the transfer
	Value string `json:"value,required"`
	// Summarized description of the transfer
	Summary string                `json:"summary,nullable"`
	JSON    starknetErc20DiffJSON `json:"-"`
}

// starknetErc20DiffJSON contains the JSON metadata for the struct
// [StarknetErc20Diff]
type starknetErc20DiffJSON struct {
	RawValue    apijson.Field
	UsdPrice    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetErc20Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetErc20DiffJSON) RawJSON() string {
	return r.raw
}

type StarknetErc721Diff struct {
	// Token ID of the transfer
	TokenID string `json:"token_id,required"`
	// USD price of the asset
	UsdPrice string `json:"usd_price,required"`
	// Summarized description of the transfer
	Summary string                 `json:"summary,nullable"`
	JSON    starknetErc721DiffJSON `json:"-"`
}

// starknetErc721DiffJSON contains the JSON metadata for the struct
// [StarknetErc721Diff]
type starknetErc721DiffJSON struct {
	TokenID     apijson.Field
	UsdPrice    apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetErc721Diff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetErc721DiffJSON) RawJSON() string {
	return r.raw
}
