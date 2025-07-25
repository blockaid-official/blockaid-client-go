// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo

import (
	"context"
	"net/http"
	"time"

	"github.com/blockaid-official/blockaid-client-go/internal/apijson"
	"github.com/blockaid-official/blockaid-client-go/internal/param"
	"github.com/blockaid-official/blockaid-client-go/internal/requestconfig"
	"github.com/blockaid-official/blockaid-client-go/option"
)

// ExchangeProtectionWithdrawalService contains methods and other services that
// help with interacting with the blockaid API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExchangeProtectionWithdrawalService] method instead.
type ExchangeProtectionWithdrawalService struct {
	Options []option.RequestOption
}

// NewExchangeProtectionWithdrawalService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewExchangeProtectionWithdrawalService(opts ...option.RequestOption) (r *ExchangeProtectionWithdrawalService) {
	r = &ExchangeProtectionWithdrawalService{}
	r.Options = opts
	return
}

// Gets a withdrawal and returns a full security assessment indicating weather or
// not this withdrawal is malicious as well as textual reasons of why the
// withdrawal was flagged that way.
func (r *ExchangeProtectionWithdrawalService) Scan(ctx context.Context, body ExchangeProtectionWithdrawalScanParams, opts ...option.RequestOption) (res *ExchangeProtectionWithdrawalScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/exchange/withdrawal/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ExchangeProtectionWithdrawalScanResponse struct {
	// An enumeration.
	RiskLevel   ExchangeProtectionWithdrawalScanResponseRiskLevel `json:"risk_level,required"`
	Description string                                            `json:"description"`
	Labels      []ExchangeProtectionWithdrawalScanResponseLabel   `json:"labels"`
	Reason      string                                            `json:"reason"`
	JSON        exchangeProtectionWithdrawalScanResponseJSON      `json:"-"`
}

// exchangeProtectionWithdrawalScanResponseJSON contains the JSON metadata for the
// struct [ExchangeProtectionWithdrawalScanResponse]
type exchangeProtectionWithdrawalScanResponseJSON struct {
	RiskLevel   apijson.Field
	Description apijson.Field
	Labels      apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExchangeProtectionWithdrawalScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exchangeProtectionWithdrawalScanResponseJSON) RawJSON() string {
	return r.raw
}

// An enumeration.
type ExchangeProtectionWithdrawalScanResponseRiskLevel string

const (
	ExchangeProtectionWithdrawalScanResponseRiskLevelLow    ExchangeProtectionWithdrawalScanResponseRiskLevel = "low"
	ExchangeProtectionWithdrawalScanResponseRiskLevelMedium ExchangeProtectionWithdrawalScanResponseRiskLevel = "medium"
	ExchangeProtectionWithdrawalScanResponseRiskLevelHigh   ExchangeProtectionWithdrawalScanResponseRiskLevel = "high"
	ExchangeProtectionWithdrawalScanResponseRiskLevelSevere ExchangeProtectionWithdrawalScanResponseRiskLevel = "severe"
)

func (r ExchangeProtectionWithdrawalScanResponseRiskLevel) IsKnown() bool {
	switch r {
	case ExchangeProtectionWithdrawalScanResponseRiskLevelLow, ExchangeProtectionWithdrawalScanResponseRiskLevelMedium, ExchangeProtectionWithdrawalScanResponseRiskLevelHigh, ExchangeProtectionWithdrawalScanResponseRiskLevelSevere:
		return true
	}
	return false
}

type ExchangeProtectionWithdrawalScanResponseLabel struct {
	Details ExchangeProtectionWithdrawalScanResponseLabelsDetails `json:"details,required"`
	Label   string                                                `json:"label,required"`
	// An enumeration.
	RiskLevel ExchangeProtectionWithdrawalScanResponseLabelsRiskLevel  `json:"risk_level,required"`
	Evidences []ExchangeProtectionWithdrawalScanResponseLabelsEvidence `json:"evidences"`
	JSON      exchangeProtectionWithdrawalScanResponseLabelJSON        `json:"-"`
}

// exchangeProtectionWithdrawalScanResponseLabelJSON contains the JSON metadata for
// the struct [ExchangeProtectionWithdrawalScanResponseLabel]
type exchangeProtectionWithdrawalScanResponseLabelJSON struct {
	Details     apijson.Field
	Label       apijson.Field
	RiskLevel   apijson.Field
	Evidences   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExchangeProtectionWithdrawalScanResponseLabel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exchangeProtectionWithdrawalScanResponseLabelJSON) RawJSON() string {
	return r.raw
}

type ExchangeProtectionWithdrawalScanResponseLabelsDetails struct {
	Description string                                                    `json:"description,required"`
	Title       string                                                    `json:"title,required"`
	JSON        exchangeProtectionWithdrawalScanResponseLabelsDetailsJSON `json:"-"`
}

// exchangeProtectionWithdrawalScanResponseLabelsDetailsJSON contains the JSON
// metadata for the struct [ExchangeProtectionWithdrawalScanResponseLabelsDetails]
type exchangeProtectionWithdrawalScanResponseLabelsDetailsJSON struct {
	Description apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExchangeProtectionWithdrawalScanResponseLabelsDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exchangeProtectionWithdrawalScanResponseLabelsDetailsJSON) RawJSON() string {
	return r.raw
}

// An enumeration.
type ExchangeProtectionWithdrawalScanResponseLabelsRiskLevel string

const (
	ExchangeProtectionWithdrawalScanResponseLabelsRiskLevelLow    ExchangeProtectionWithdrawalScanResponseLabelsRiskLevel = "low"
	ExchangeProtectionWithdrawalScanResponseLabelsRiskLevelMedium ExchangeProtectionWithdrawalScanResponseLabelsRiskLevel = "medium"
	ExchangeProtectionWithdrawalScanResponseLabelsRiskLevelHigh   ExchangeProtectionWithdrawalScanResponseLabelsRiskLevel = "high"
	ExchangeProtectionWithdrawalScanResponseLabelsRiskLevelSevere ExchangeProtectionWithdrawalScanResponseLabelsRiskLevel = "severe"
)

func (r ExchangeProtectionWithdrawalScanResponseLabelsRiskLevel) IsKnown() bool {
	switch r {
	case ExchangeProtectionWithdrawalScanResponseLabelsRiskLevelLow, ExchangeProtectionWithdrawalScanResponseLabelsRiskLevelMedium, ExchangeProtectionWithdrawalScanResponseLabelsRiskLevelHigh, ExchangeProtectionWithdrawalScanResponseLabelsRiskLevelSevere:
		return true
	}
	return false
}

type ExchangeProtectionWithdrawalScanResponseLabelsEvidence struct {
	Addresses []ExchangeProtectionWithdrawalScanResponseLabelsEvidencesAddress `json:"addresses"`
	URLs      []ExchangeProtectionWithdrawalScanResponseLabelsEvidencesURL     `json:"urls"`
	JSON      exchangeProtectionWithdrawalScanResponseLabelsEvidenceJSON       `json:"-"`
}

// exchangeProtectionWithdrawalScanResponseLabelsEvidenceJSON contains the JSON
// metadata for the struct [ExchangeProtectionWithdrawalScanResponseLabelsEvidence]
type exchangeProtectionWithdrawalScanResponseLabelsEvidenceJSON struct {
	Addresses   apijson.Field
	URLs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExchangeProtectionWithdrawalScanResponseLabelsEvidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exchangeProtectionWithdrawalScanResponseLabelsEvidenceJSON) RawJSON() string {
	return r.raw
}

type ExchangeProtectionWithdrawalScanResponseLabelsEvidencesAddress struct {
	Address string                                                                  `json:"address,required"`
	Labels  []ExchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabel `json:"labels,required"`
	JSON    exchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressJSON      `json:"-"`
}

// exchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressJSON contains the
// JSON metadata for the struct
// [ExchangeProtectionWithdrawalScanResponseLabelsEvidencesAddress]
type exchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressJSON struct {
	Address     apijson.Field
	Labels      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExchangeProtectionWithdrawalScanResponseLabelsEvidencesAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressJSON) RawJSON() string {
	return r.raw
}

type ExchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabel struct {
	Label  string `json:"label,required"`
	Origin string `json:"origin,required"`
	// An enumeration.
	Severity ExchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabelsSeverity `json:"severity,required"`
	JSON     exchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabelJSON      `json:"-"`
}

// exchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabelJSON
// contains the JSON metadata for the struct
// [ExchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabel]
type exchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabelJSON struct {
	Label       apijson.Field
	Origin      apijson.Field
	Severity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabelJSON) RawJSON() string {
	return r.raw
}

// An enumeration.
type ExchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabelsSeverity string

const (
	ExchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabelsSeverityLow    ExchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabelsSeverity = "low"
	ExchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabelsSeverityMedium ExchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabelsSeverity = "medium"
	ExchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabelsSeverityHigh   ExchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabelsSeverity = "high"
	ExchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabelsSeveritySevere ExchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabelsSeverity = "severe"
)

func (r ExchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabelsSeverity) IsKnown() bool {
	switch r {
	case ExchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabelsSeverityLow, ExchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabelsSeverityMedium, ExchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabelsSeverityHigh, ExchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabelsSeveritySevere:
		return true
	}
	return false
}

type ExchangeProtectionWithdrawalScanResponseLabelsEvidencesURL struct {
	Labels []ExchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabel `json:"labels,required"`
	URL    string                                                             `json:"url,required"`
	JSON   exchangeProtectionWithdrawalScanResponseLabelsEvidencesURLJSON     `json:"-"`
}

// exchangeProtectionWithdrawalScanResponseLabelsEvidencesURLJSON contains the JSON
// metadata for the struct
// [ExchangeProtectionWithdrawalScanResponseLabelsEvidencesURL]
type exchangeProtectionWithdrawalScanResponseLabelsEvidencesURLJSON struct {
	Labels      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExchangeProtectionWithdrawalScanResponseLabelsEvidencesURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exchangeProtectionWithdrawalScanResponseLabelsEvidencesURLJSON) RawJSON() string {
	return r.raw
}

type ExchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabel struct {
	Label  string `json:"label,required"`
	Origin string `json:"origin,required"`
	// An enumeration.
	Severity ExchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabelsSeverity `json:"severity,required"`
	JSON     exchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabelJSON      `json:"-"`
}

// exchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabelJSON contains
// the JSON metadata for the struct
// [ExchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabel]
type exchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabelJSON struct {
	Label       apijson.Field
	Origin      apijson.Field
	Severity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabelJSON) RawJSON() string {
	return r.raw
}

// An enumeration.
type ExchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabelsSeverity string

const (
	ExchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabelsSeverityLow    ExchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabelsSeverity = "low"
	ExchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabelsSeverityMedium ExchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabelsSeverity = "medium"
	ExchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabelsSeverityHigh   ExchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabelsSeverity = "high"
	ExchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabelsSeveritySevere ExchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabelsSeverity = "severe"
)

func (r ExchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabelsSeverity) IsKnown() bool {
	switch r {
	case ExchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabelsSeverityLow, ExchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabelsSeverityMedium, ExchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabelsSeverityHigh, ExchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabelsSeveritySevere:
		return true
	}
	return false
}

type ExchangeProtectionWithdrawalScanParams struct {
	Account            param.Field[ExchangeProtectionWithdrawalScanParamsAccount]            `json:"account,required"`
	EventTime          param.Field[time.Time]                                                `json:"event_time,required" format:"date-time"`
	OnchainTransaction param.Field[ExchangeProtectionWithdrawalScanParamsOnchainTransaction] `json:"onchain_transaction,required"`
	WithdrawalID       param.Field[string]                                                   `json:"withdrawal_id,required"`
	ConnectionMetadata param.Field[ExchangeProtectionWithdrawalScanParamsConnectionMetadata] `json:"connection_metadata"`
}

func (r ExchangeProtectionWithdrawalScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExchangeProtectionWithdrawalScanParamsAccount struct {
	AccountID                param.Field[string]    `json:"account_id,required"`
	AccountCreationTimestamp param.Field[time.Time] `json:"account_creation_timestamp" format:"date-time"`
	UserAge                  param.Field[int64]     `json:"user_age"`
	UserCountryCode          param.Field[string]    `json:"user_country_code"`
}

func (r ExchangeProtectionWithdrawalScanParamsAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExchangeProtectionWithdrawalScanParamsOnchainTransaction struct {
	Amount param.Field[float64] `json:"amount,required"`
	Asset  param.Field[string]  `json:"asset,required"`
	// An enumeration.
	Chain     param.Field[ExchangeProtectionWithdrawalScanParamsOnchainTransactionChain] `json:"chain,required"`
	ToAddress param.Field[string]                                                        `json:"to_address,required"`
}

func (r ExchangeProtectionWithdrawalScanParamsOnchainTransaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An enumeration.
type ExchangeProtectionWithdrawalScanParamsOnchainTransactionChain string

const (
	ExchangeProtectionWithdrawalScanParamsOnchainTransactionChainEthereum ExchangeProtectionWithdrawalScanParamsOnchainTransactionChain = "ethereum"
	ExchangeProtectionWithdrawalScanParamsOnchainTransactionChainBase     ExchangeProtectionWithdrawalScanParamsOnchainTransactionChain = "base"
	ExchangeProtectionWithdrawalScanParamsOnchainTransactionChainArbitrum ExchangeProtectionWithdrawalScanParamsOnchainTransactionChain = "arbitrum"
	ExchangeProtectionWithdrawalScanParamsOnchainTransactionChainOptimism ExchangeProtectionWithdrawalScanParamsOnchainTransactionChain = "optimism"
	ExchangeProtectionWithdrawalScanParamsOnchainTransactionChainPolygon  ExchangeProtectionWithdrawalScanParamsOnchainTransactionChain = "polygon"
	ExchangeProtectionWithdrawalScanParamsOnchainTransactionChainBitcoin  ExchangeProtectionWithdrawalScanParamsOnchainTransactionChain = "bitcoin"
	ExchangeProtectionWithdrawalScanParamsOnchainTransactionChainSolana   ExchangeProtectionWithdrawalScanParamsOnchainTransactionChain = "solana"
)

func (r ExchangeProtectionWithdrawalScanParamsOnchainTransactionChain) IsKnown() bool {
	switch r {
	case ExchangeProtectionWithdrawalScanParamsOnchainTransactionChainEthereum, ExchangeProtectionWithdrawalScanParamsOnchainTransactionChainBase, ExchangeProtectionWithdrawalScanParamsOnchainTransactionChainArbitrum, ExchangeProtectionWithdrawalScanParamsOnchainTransactionChainOptimism, ExchangeProtectionWithdrawalScanParamsOnchainTransactionChainPolygon, ExchangeProtectionWithdrawalScanParamsOnchainTransactionChainBitcoin, ExchangeProtectionWithdrawalScanParamsOnchainTransactionChainSolana:
		return true
	}
	return false
}

type ExchangeProtectionWithdrawalScanParamsConnectionMetadata struct {
	CustomerIP param.Field[string] `json:"customer_ip,required"`
	UserAgent  param.Field[string] `json:"user_agent,required"`
}

func (r ExchangeProtectionWithdrawalScanParamsConnectionMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
