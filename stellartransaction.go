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

// StellarTransactionService contains methods and other services that help with
// interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStellarTransactionService] method instead.
type StellarTransactionService struct {
	Options []option.RequestOption
}

// NewStellarTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStellarTransactionService(opts ...option.RequestOption) (r *StellarTransactionService) {
	r = &StellarTransactionService{}
	r.Options = opts
	return
}

// Report Transaction
func (r *StellarTransactionService) Report(ctx context.Context, body StellarTransactionReportParams, opts ...option.RequestOption) (res *int64, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/stellar/transaction/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Scan Transaction
func (r *StellarTransactionService) Scan(ctx context.Context, body StellarTransactionScanParams, opts ...option.RequestOption) (res *StellarTransactionScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/stellar/transaction/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type StellarTransactionReportParams struct {
	Details param.Field[string]                                    `json:"details,required"`
	Event   param.Field[StellarTransactionReportParamsEvent]       `json:"event,required"`
	Report  param.Field[StellarTransactionReportParamsReportUnion] `json:"report,required"`
}

func (r StellarTransactionReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StellarTransactionReportParamsEvent string

const (
	StellarTransactionReportParamsEventShouldBeMalicious StellarTransactionReportParamsEvent = "should_be_malicious"
	StellarTransactionReportParamsEventShouldBeBenign    StellarTransactionReportParamsEvent = "should_be_benign"
)

func (r StellarTransactionReportParamsEvent) IsKnown() bool {
	switch r {
	case StellarTransactionReportParamsEventShouldBeMalicious, StellarTransactionReportParamsEventShouldBeBenign:
		return true
	}
	return false
}

type StellarTransactionReportParamsReport struct {
	Params param.Field[interface{}]                              `json:"params,required"`
	ID     param.Field[string]                                   `json:"id"`
	Type   param.Field[StellarTransactionReportParamsReportType] `json:"type"`
}

func (r StellarTransactionReportParamsReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionReportParamsReport) implementsStellarTransactionReportParamsReportUnion() {}

// Satisfied by [StellarTransactionReportParamsReportStellarAppealRequestID],
// [StellarTransactionReportParamsReportStellarAppealTransactionDataReport],
// [StellarTransactionReportParamsReport].
type StellarTransactionReportParamsReportUnion interface {
	implementsStellarTransactionReportParamsReportUnion()
}

type StellarTransactionReportParamsReportStellarAppealRequestID struct {
	ID   param.Field[string]                                                         `json:"id,required"`
	Type param.Field[StellarTransactionReportParamsReportStellarAppealRequestIDType] `json:"type"`
}

func (r StellarTransactionReportParamsReportStellarAppealRequestID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionReportParamsReportStellarAppealRequestID) implementsStellarTransactionReportParamsReportUnion() {
}

type StellarTransactionReportParamsReportStellarAppealRequestIDType string

const (
	StellarTransactionReportParamsReportStellarAppealRequestIDTypeRequestID StellarTransactionReportParamsReportStellarAppealRequestIDType = "request_id"
)

func (r StellarTransactionReportParamsReportStellarAppealRequestIDType) IsKnown() bool {
	switch r {
	case StellarTransactionReportParamsReportStellarAppealRequestIDTypeRequestID:
		return true
	}
	return false
}

type StellarTransactionReportParamsReportStellarAppealTransactionDataReport struct {
	Params param.Field[StellarTransactionReportParamsReportStellarAppealTransactionDataReportParams] `json:"params,required"`
	Type   param.Field[StellarTransactionReportParamsReportStellarAppealTransactionDataReportType]   `json:"type"`
}

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReport) implementsStellarTransactionReportParamsReportUnion() {
}

type StellarTransactionReportParamsReportStellarAppealTransactionDataReportParams struct {
	AccountAddress param.Field[string] `json:"account_address,required"`
	// A CAIP-2 chain ID or a Stellar network name
	Chain param.Field[StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsChain] `json:"chain,required"`
	// Metadata
	Metadata    param.Field[StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataUnion] `json:"metadata,required"`
	Transaction param.Field[string]                                                                                    `json:"transaction,required"`
	// List of options to include in the response
	//
	// - `Options.validation`: Include Options.validation output in the response
	//
	// - `Options.simulation`: Include Options.simulation output in the response
	Options param.Field[[]StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsOption] `json:"options"`
}

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A CAIP-2 chain ID or a Stellar network name
type StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsChain string

const (
	StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsChainPubnet    StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsChain = "pubnet"
	StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsChainFuturenet StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsChain = "futurenet"
	StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsChainTestnet   StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsChain = "testnet"
)

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsChain) IsKnown() bool {
	switch r {
	case StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsChainPubnet, StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsChainFuturenet, StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsChainTestnet:
		return true
	}
	return false
}

// Metadata
type StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadata struct {
	// Metadata for wallet requests
	Type param.Field[StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataType] `json:"type"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url"`
}

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadata) implementsStellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataUnion() {
}

// Metadata
//
// Satisfied by
// [StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarWalletRequestMetadata],
// [StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarInAppRequestMetadata],
// [StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadata].
type StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataUnion interface {
	implementsStellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataUnion()
}

type StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarWalletRequestMetadata struct {
	// Metadata for wallet requests
	Type param.Field[StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarWalletRequestMetadataType] `json:"type,required"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url,required"`
}

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarWalletRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarWalletRequestMetadata) implementsStellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataUnion() {
}

// Metadata for wallet requests
type StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarWalletRequestMetadataType string

const (
	StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarWalletRequestMetadataTypeWallet StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarWalletRequestMetadataType = "wallet"
)

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarWalletRequestMetadataType) IsKnown() bool {
	switch r {
	case StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarWalletRequestMetadataTypeWallet:
		return true
	}
	return false
}

type StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarInAppRequestMetadata struct {
	// Metadata for in-app requests
	Type param.Field[StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarInAppRequestMetadataType] `json:"type"`
}

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarInAppRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarInAppRequestMetadata) implementsStellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataUnion() {
}

// Metadata for in-app requests
type StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarInAppRequestMetadataType string

const (
	StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarInAppRequestMetadataTypeInApp StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarInAppRequestMetadataType = "in_app"
)

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarInAppRequestMetadataType) IsKnown() bool {
	switch r {
	case StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataStellarInAppRequestMetadataTypeInApp:
		return true
	}
	return false
}

// Metadata for wallet requests
type StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataType string

const (
	StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataTypeWallet StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataType = "wallet"
	StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataTypeInApp  StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataType = "in_app"
)

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataType) IsKnown() bool {
	switch r {
	case StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataTypeWallet, StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsMetadataTypeInApp:
		return true
	}
	return false
}

type StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsOption string

const (
	StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsOptionValidation StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsOption = "validation"
	StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsOptionSimulation StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsOption = "simulation"
)

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsOption) IsKnown() bool {
	switch r {
	case StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsOptionValidation, StellarTransactionReportParamsReportStellarAppealTransactionDataReportParamsOptionSimulation:
		return true
	}
	return false
}

type StellarTransactionReportParamsReportStellarAppealTransactionDataReportType string

const (
	StellarTransactionReportParamsReportStellarAppealTransactionDataReportTypeParams StellarTransactionReportParamsReportStellarAppealTransactionDataReportType = "params"
)

func (r StellarTransactionReportParamsReportStellarAppealTransactionDataReportType) IsKnown() bool {
	switch r {
	case StellarTransactionReportParamsReportStellarAppealTransactionDataReportTypeParams:
		return true
	}
	return false
}

type StellarTransactionReportParamsReportType string

const (
	StellarTransactionReportParamsReportTypeRequestID StellarTransactionReportParamsReportType = "request_id"
	StellarTransactionReportParamsReportTypeParams    StellarTransactionReportParamsReportType = "params"
)

func (r StellarTransactionReportParamsReportType) IsKnown() bool {
	switch r {
	case StellarTransactionReportParamsReportTypeRequestID, StellarTransactionReportParamsReportTypeParams:
		return true
	}
	return false
}

type StellarTransactionScanParams struct {
	AccountAddress param.Field[string] `json:"account_address,required"`
	// A CAIP-2 chain ID or a Stellar network name
	Chain param.Field[StellarTransactionScanParamsChain] `json:"chain,required"`
	// Metadata
	Metadata    param.Field[StellarTransactionScanParamsMetadataUnion] `json:"metadata,required"`
	Transaction param.Field[string]                                    `json:"transaction,required"`
	// List of options to include in the response
	//
	// - `Options.validation`: Include Options.validation output in the response
	//
	// - `Options.simulation`: Include Options.simulation output in the response
	Options param.Field[[]StellarTransactionScanParamsOption] `json:"options"`
}

func (r StellarTransactionScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A CAIP-2 chain ID or a Stellar network name
type StellarTransactionScanParamsChain string

const (
	StellarTransactionScanParamsChainPubnet    StellarTransactionScanParamsChain = "pubnet"
	StellarTransactionScanParamsChainFuturenet StellarTransactionScanParamsChain = "futurenet"
	StellarTransactionScanParamsChainTestnet   StellarTransactionScanParamsChain = "testnet"
)

func (r StellarTransactionScanParamsChain) IsKnown() bool {
	switch r {
	case StellarTransactionScanParamsChainPubnet, StellarTransactionScanParamsChainFuturenet, StellarTransactionScanParamsChainTestnet:
		return true
	}
	return false
}

// Metadata
type StellarTransactionScanParamsMetadata struct {
	// Metadata for wallet requests
	Type param.Field[StellarTransactionScanParamsMetadataType] `json:"type"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url"`
}

func (r StellarTransactionScanParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionScanParamsMetadata) implementsStellarTransactionScanParamsMetadataUnion() {}

// Metadata
//
// Satisfied by [StellarTransactionScanParamsMetadataStellarWalletRequestMetadata],
// [StellarTransactionScanParamsMetadataStellarInAppRequestMetadata],
// [StellarTransactionScanParamsMetadata].
type StellarTransactionScanParamsMetadataUnion interface {
	implementsStellarTransactionScanParamsMetadataUnion()
}

type StellarTransactionScanParamsMetadataStellarWalletRequestMetadata struct {
	// Metadata for wallet requests
	Type param.Field[StellarTransactionScanParamsMetadataStellarWalletRequestMetadataType] `json:"type,required"`
	// URL of the dApp originating the transaction
	URL param.Field[string] `json:"url,required"`
}

func (r StellarTransactionScanParamsMetadataStellarWalletRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionScanParamsMetadataStellarWalletRequestMetadata) implementsStellarTransactionScanParamsMetadataUnion() {
}

// Metadata for wallet requests
type StellarTransactionScanParamsMetadataStellarWalletRequestMetadataType string

const (
	StellarTransactionScanParamsMetadataStellarWalletRequestMetadataTypeWallet StellarTransactionScanParamsMetadataStellarWalletRequestMetadataType = "wallet"
)

func (r StellarTransactionScanParamsMetadataStellarWalletRequestMetadataType) IsKnown() bool {
	switch r {
	case StellarTransactionScanParamsMetadataStellarWalletRequestMetadataTypeWallet:
		return true
	}
	return false
}

type StellarTransactionScanParamsMetadataStellarInAppRequestMetadata struct {
	// Metadata for in-app requests
	Type param.Field[StellarTransactionScanParamsMetadataStellarInAppRequestMetadataType] `json:"type"`
}

func (r StellarTransactionScanParamsMetadataStellarInAppRequestMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StellarTransactionScanParamsMetadataStellarInAppRequestMetadata) implementsStellarTransactionScanParamsMetadataUnion() {
}

// Metadata for in-app requests
type StellarTransactionScanParamsMetadataStellarInAppRequestMetadataType string

const (
	StellarTransactionScanParamsMetadataStellarInAppRequestMetadataTypeInApp StellarTransactionScanParamsMetadataStellarInAppRequestMetadataType = "in_app"
)

func (r StellarTransactionScanParamsMetadataStellarInAppRequestMetadataType) IsKnown() bool {
	switch r {
	case StellarTransactionScanParamsMetadataStellarInAppRequestMetadataTypeInApp:
		return true
	}
	return false
}

// Metadata for wallet requests
type StellarTransactionScanParamsMetadataType string

const (
	StellarTransactionScanParamsMetadataTypeWallet StellarTransactionScanParamsMetadataType = "wallet"
	StellarTransactionScanParamsMetadataTypeInApp  StellarTransactionScanParamsMetadataType = "in_app"
)

func (r StellarTransactionScanParamsMetadataType) IsKnown() bool {
	switch r {
	case StellarTransactionScanParamsMetadataTypeWallet, StellarTransactionScanParamsMetadataTypeInApp:
		return true
	}
	return false
}

type StellarTransactionScanParamsOption string

const (
	StellarTransactionScanParamsOptionValidation StellarTransactionScanParamsOption = "validation"
	StellarTransactionScanParamsOptionSimulation StellarTransactionScanParamsOption = "simulation"
)

func (r StellarTransactionScanParamsOption) IsKnown() bool {
	switch r {
	case StellarTransactionScanParamsOptionValidation, StellarTransactionScanParamsOptionSimulation:
		return true
	}
	return false
}
