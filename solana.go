// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/option"
)

// SolanaService contains methods and other services that help with interacting
// with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSolanaService] method instead.
type SolanaService struct {
	Options []option.RequestOption
	Message *SolanaMessageService
	Address *SolanaAddressService
}

// NewSolanaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSolanaService(opts ...option.RequestOption) (r *SolanaService) {
	r = &SolanaService{}
	r.Options = opts
	r.Message = NewSolanaMessageService(opts...)
	r.Address = NewSolanaAddressService(opts...)
	return
}

type SolanaAccountRentFee struct {
	// Encoded public key of the account that rent was deposited for.
	AccountAddress string `json:"account_address" api:"required"`
	// Type of the newly created account, e.g. TOKEN_ACCOUNT or SYSTEM_ACCOUNT.
	AccountType string `json:"account_type" api:"required"`
	// Rent deposit amount in lamports, as a string.
	Lamports string                   `json:"lamports" api:"required"`
	JSON     solanaAccountRentFeeJSON `json:"-"`
}

// solanaAccountRentFeeJSON contains the JSON metadata for the struct
// [SolanaAccountRentFee]
type solanaAccountRentFeeJSON struct {
	AccountAddress apijson.Field
	AccountType    apijson.Field
	Lamports       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SolanaAccountRentFee) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaAccountRentFeeJSON) RawJSON() string {
	return r.raw
}

type SolanaGasEstimation struct {
	// Base transaction fee in lamports
	NetworkFee string `json:"network_fee" api:"required"`
	// Prioritization fee in lamports
	PriorityFee string `json:"priority_fee" api:"required"`
	// Total fee in lamports: network fee plus priority fee plus account rent fees.
	Total string `json:"total" api:"required"`
	// Total fee in lamports (equal to total; added for cross-chain consistency)
	Used string `json:"used" api:"required"`
	// Rent deposit fees for accounts created by the transaction; empty when no new
	// accounts are created.
	AccountRentFees []SolanaAccountRentFee  `json:"account_rent_fees"`
	JSON            solanaGasEstimationJSON `json:"-"`
}

// solanaGasEstimationJSON contains the JSON metadata for the struct
// [SolanaGasEstimation]
type solanaGasEstimationJSON struct {
	NetworkFee      apijson.Field
	PriorityFee     apijson.Field
	Total           apijson.Field
	Used            apijson.Field
	AccountRentFees apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SolanaGasEstimation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r solanaGasEstimationJSON) RawJSON() string {
	return r.raw
}
