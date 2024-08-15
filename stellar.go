// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/option"
)

// StellarService contains methods and other services that help with interacting
// with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStellarService] method instead.
type StellarService struct {
	Options     []option.RequestOption
	Transaction *StellarTransactionService
}

// NewStellarService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStellarService(opts ...option.RequestOption) (r *StellarService) {
	r = &StellarService{}
	r.Options = opts
	r.Transaction = NewStellarTransactionService(opts...)
	return
}

type StellarAssetContractDetailsSchema struct {
	// Address of the asset's contract
	Address string `json:"address,required"`
	// Asset code
	Name string `json:"name,required"`
	// Asset symbol
	Symbol string `json:"symbol,required"`
	// Type of the asset (`CONTRACT`)
	Type StellarAssetContractDetailsSchemaType `json:"type"`
	JSON stellarAssetContractDetailsSchemaJSON `json:"-"`
}

// stellarAssetContractDetailsSchemaJSON contains the JSON metadata for the struct
// [StellarAssetContractDetailsSchema]
type stellarAssetContractDetailsSchemaJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarAssetContractDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarAssetContractDetailsSchemaJSON) RawJSON() string {
	return r.raw
}

func (r StellarAssetContractDetailsSchema) implementsStellarTransactionScanResponseSimulationStellarSimulationResultSchemaAccountSummaryAssetsDiffsAsset() {
}

func (r StellarAssetContractDetailsSchema) implementsStellarTransactionScanResponseSimulationStellarSimulationResultSchemaAssetsDiffsAsset() {
}

// Type of the asset (`CONTRACT`)
type StellarAssetContractDetailsSchemaType string

const (
	StellarAssetContractDetailsSchemaTypeContract StellarAssetContractDetailsSchemaType = "CONTRACT"
)

func (r StellarAssetContractDetailsSchemaType) IsKnown() bool {
	switch r {
	case StellarAssetContractDetailsSchemaTypeContract:
		return true
	}
	return false
}

type StellarAssetTransferDetailsSchema struct {
	// Raw value of the transfer
	RawValue int64 `json:"raw_value,required"`
	// Value of the transfer
	Value float64 `json:"value,required"`
	// Summarized description of the transfer
	Summary string `json:"summary,nullable"`
	// USD price of the asset
	UsdPrice float64                               `json:"usd_price"`
	JSON     stellarAssetTransferDetailsSchemaJSON `json:"-"`
}

// stellarAssetTransferDetailsSchemaJSON contains the JSON metadata for the struct
// [StellarAssetTransferDetailsSchema]
type stellarAssetTransferDetailsSchemaJSON struct {
	RawValue    apijson.Field
	Value       apijson.Field
	Summary     apijson.Field
	UsdPrice    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StellarAssetTransferDetailsSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stellarAssetTransferDetailsSchemaJSON) RawJSON() string {
	return r.raw
}
