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

type StarknetAccountErc1155Exposure struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC1155`)
	Type StarknetAccountErc1155ExposureType `json:"type"`
	JSON starknetAccountErc1155ExposureJSON `json:"-"`
}

// starknetAccountErc1155ExposureJSON contains the JSON metadata for the struct
// [StarknetAccountErc1155Exposure]
type starknetAccountErc1155ExposureJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetAccountErc1155Exposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetAccountErc1155ExposureJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC1155`)
type StarknetAccountErc1155ExposureType string

const (
	StarknetAccountErc1155ExposureTypeErc1155 StarknetAccountErc1155ExposureType = "ERC1155"
)

func (r StarknetAccountErc1155ExposureType) IsKnown() bool {
	switch r {
	case StarknetAccountErc1155ExposureTypeErc1155:
		return true
	}
	return false
}

type StarknetAccountErc20Exposure struct {
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
	Type StarknetAccountErc20ExposureType `json:"type"`
	JSON starknetAccountErc20ExposureJSON `json:"-"`
}

// starknetAccountErc20ExposureJSON contains the JSON metadata for the struct
// [StarknetAccountErc20Exposure]
type starknetAccountErc20ExposureJSON struct {
	Address     apijson.Field
	Decimals    apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetAccountErc20Exposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetAccountErc20ExposureJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC20`)
type StarknetAccountErc20ExposureType string

const (
	StarknetAccountErc20ExposureTypeErc20 StarknetAccountErc20ExposureType = "ERC20"
)

func (r StarknetAccountErc20ExposureType) IsKnown() bool {
	switch r {
	case StarknetAccountErc20ExposureTypeErc20:
		return true
	}
	return false
}

type StarknetAccountErc721Exposure struct {
	// Address of the token's contract
	Address string `json:"address,required"`
	// token's name
	Name string `json:"name,required"`
	// token's symbol
	Symbol string `json:"symbol,required"`
	// URL of the asset's logo
	LogoURL string `json:"logo_url,nullable"`
	// Type of the asset (`ERC721`)
	Type StarknetAccountErc721ExposureType `json:"type"`
	JSON starknetAccountErc721ExposureJSON `json:"-"`
}

// starknetAccountErc721ExposureJSON contains the JSON metadata for the struct
// [StarknetAccountErc721Exposure]
type starknetAccountErc721ExposureJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	LogoURL     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StarknetAccountErc721Exposure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r starknetAccountErc721ExposureJSON) RawJSON() string {
	return r.raw
}

// Type of the asset (`ERC721`)
type StarknetAccountErc721ExposureType string

const (
	StarknetAccountErc721ExposureTypeErc721 StarknetAccountErc721ExposureType = "ERC721"
)

func (r StarknetAccountErc721ExposureType) IsKnown() bool {
	switch r {
	case StarknetAccountErc721ExposureTypeErc721:
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
