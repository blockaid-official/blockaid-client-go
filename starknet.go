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

type StarknetErc1155Details struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC1155`)
	Type StarknetErc1155DetailsType `json:"type"`
	JSON starknetErc1155DetailsJSON `json:"-"`
}

// starknetErc1155DetailsJSON contains the JSON metadata for the struct
// [StarknetErc1155Details]
type starknetErc1155DetailsJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetErc1155Details) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetErc1155DetailsJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC1155`)
type StarknetErc1155DetailsType string

const (
	StarknetErc1155DetailsTypeErc1155 StarknetErc1155DetailsType = "ERC1155"
)

func (r StarknetErc1155DetailsType) IsKnown() bool {
	switch r {
	case StarknetErc1155DetailsTypeErc1155:
		return true
	}
	return false
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

type StarknetErc20Details struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's decimals
	Decimals int64 `json:"decimals,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC20`)
	Type StarknetErc20DetailsType `json:"type"`
	JSON starknetErc20DetailsJSON `json:"-"`
}

// starknetErc20DetailsJSON contains the JSON metadata for the struct
// [StarknetErc20Details]
type starknetErc20DetailsJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetErc20Details) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetErc20DetailsJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC20`)
type StarknetErc20DetailsType string

const (
	StarknetErc20DetailsTypeErc20 StarknetErc20DetailsType = "ERC20"
)

func (r StarknetErc20DetailsType) IsKnown() bool {
	switch r {
	case StarknetErc20DetailsTypeErc20:
		return true
	}
	return false
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

type StarknetErc721Details struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC721`)
	Type StarknetErc721DetailsType `json:"type"`
	JSON starknetErc721DetailsJSON `json:"-"`
}

// starknetErc721DetailsJSON contains the JSON metadata for the struct
// [StarknetErc721Details]
type starknetErc721DetailsJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetErc721Details) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetErc721DetailsJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC721`)
type StarknetErc721DetailsType string

const (
	StarknetErc721DetailsTypeErc721 StarknetErc721DetailsType = "ERC721"
)

func (r StarknetErc721DetailsType) IsKnown() bool {
	switch r {
	case StarknetErc721DetailsTypeErc721:
		return true
	}
	return false
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
