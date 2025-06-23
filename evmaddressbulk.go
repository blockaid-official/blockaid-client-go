// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"net/http"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/requestconfig"
	"github.com/blockaid-official/blockaid-client-go/option"
)

// EvmAddressBulkService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvmAddressBulkService] method instead.
type EvmAddressBulkService struct {
	Options []option.RequestOption
}

// NewEvmAddressBulkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEvmAddressBulkService(opts ...option.RequestOption) (r *EvmAddressBulkService) {
	r = &EvmAddressBulkService{}
	r.Options = opts
	return
}

// Gets a list of addresses and returns a security assessment of each address.
func (r *EvmAddressBulkService) Scan(ctx context.Context, body EvmAddressBulkScanParams, opts ...option.RequestOption) (res *EvmAddressBulkScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/evm/address-bulk/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets a list of addresses and returns a security assessment with the feature
// information of each address.
func (r *EvmAddressBulkService) ScanExtended(ctx context.Context, body EvmAddressBulkScanExtendedParams, opts ...option.RequestOption) (res *ValidateBulkExtendedAddressesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/evm/address-bulk/scan-extended"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EvmAddressBulkScanResponse map[string]EvmAddressBulkScanResponseItem

// An enumeration.
type EvmAddressBulkScanResponseItem string

const (
	EvmAddressBulkScanResponseItemMalicious EvmAddressBulkScanResponseItem = "Malicious"
	EvmAddressBulkScanResponseItemWarning   EvmAddressBulkScanResponseItem = "Warning"
	EvmAddressBulkScanResponseItemBenign    EvmAddressBulkScanResponseItem = "Benign"
	EvmAddressBulkScanResponseItemError     EvmAddressBulkScanResponseItem = "Error"
)

func (r EvmAddressBulkScanResponseItem) IsKnown() bool {
	switch r {
	case EvmAddressBulkScanResponseItemMalicious, EvmAddressBulkScanResponseItemWarning, EvmAddressBulkScanResponseItemBenign, EvmAddressBulkScanResponseItemError:
		return true
	}
	return false
}

type EvmAddressBulkScanParams struct {
	ValidateBulkAddresses ValidateBulkAddressesParam `json:"ValidateBulkAddresses,required"`
}

func (r EvmAddressBulkScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ValidateBulkAddresses)
}

type EvmAddressBulkScanExtendedParams struct {
	ValidateBulkExtendedAddressesRequest ValidateBulkExtendedAddressesRequestParam `json:"ValidateBulkExtendedAddressesRequest,required"`
}

func (r EvmAddressBulkScanExtendedParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ValidateBulkExtendedAddressesRequest)
}
