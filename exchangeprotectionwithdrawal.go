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
	DateAdded string `json:"date_added,required"`
	IsPublic  bool   `json:"is_public,required"`
	Label     string `json:"label,required"`
	Origin    string `json:"origin,required"`
	// An enumeration.
	Severity              ExchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabelsSeverity `json:"severity,required"`
	Chains                []string                                                                       `json:"chains"`
	ReportingOrganization string                                                                         `json:"reporting_organization"`
	JSON                  exchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabelJSON      `json:"-"`
}

// exchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabelJSON
// contains the JSON metadata for the struct
// [ExchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabel]
type exchangeProtectionWithdrawalScanResponseLabelsEvidencesAddressesLabelJSON struct {
	DateAdded             apijson.Field
	IsPublic              apijson.Field
	Label                 apijson.Field
	Origin                apijson.Field
	Severity              apijson.Field
	Chains                apijson.Field
	ReportingOrganization apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
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
	DateAdded string `json:"date_added,required"`
	IsPublic  bool   `json:"is_public,required"`
	Label     string `json:"label,required"`
	Origin    string `json:"origin,required"`
	// An enumeration.
	Severity              ExchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabelsSeverity `json:"severity,required"`
	Chains                []string                                                                  `json:"chains"`
	ReportingOrganization string                                                                    `json:"reporting_organization"`
	JSON                  exchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabelJSON      `json:"-"`
}

// exchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabelJSON contains
// the JSON metadata for the struct
// [ExchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabel]
type exchangeProtectionWithdrawalScanResponseLabelsEvidencesURLsLabelJSON struct {
	DateAdded             apijson.Field
	IsPublic              apijson.Field
	Label                 apijson.Field
	Origin                apijson.Field
	Severity              apijson.Field
	Chains                apijson.Field
	ReportingOrganization apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
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
	Account            param.Field[ExchangeProtectionWithdrawalScanParamsAccount]                 `json:"account,required"`
	EventTime          param.Field[time.Time]                                                     `json:"event_time,required" format:"date-time"`
	OnchainTransaction param.Field[ExchangeProtectionWithdrawalScanParamsOnchainTransactionUnion] `json:"onchain_transaction,required"`
	WithdrawalID       param.Field[string]                                                        `json:"withdrawal_id,required"`
	ConnectionMetadata param.Field[ExchangeProtectionWithdrawalScanParamsConnectionMetadata]      `json:"connection_metadata"`
}

func (r ExchangeProtectionWithdrawalScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExchangeProtectionWithdrawalScanParamsAccount struct {
	AccountID       param.Field[string]    `json:"account_id,required"`
	UserCountryCode param.Field[string]    `json:"user_country_code,required"`
	AgeInYears      param.Field[int64]     `json:"age_in_years"`
	Created         param.Field[time.Time] `json:"created" format:"date-time"`
}

func (r ExchangeProtectionWithdrawalScanParamsAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExchangeProtectionWithdrawalScanParamsOnchainTransaction struct {
	Amount param.Field[float64]                                                       `json:"amount,required"`
	Asset  param.Field[ExchangeProtectionWithdrawalScanParamsOnchainTransactionAsset] `json:"asset,required"`
	// An enumeration.
	Chain     param.Field[ExchangeProtectionWithdrawalScanParamsOnchainTransactionChain] `json:"chain,required"`
	ToAddress param.Field[string]                                                        `json:"to_address,required"`
}

func (r ExchangeProtectionWithdrawalScanParamsOnchainTransaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExchangeProtectionWithdrawalScanParamsOnchainTransaction) implementsExchangeProtectionWithdrawalScanParamsOnchainTransactionUnion() {
}

// Satisfied by
// [ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransaction],
// [ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestBitcoinOnchainTransaction],
// [ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestStellarOnchainTransaction],
// [ExchangeProtectionWithdrawalScanParamsOnchainTransaction].
type ExchangeProtectionWithdrawalScanParamsOnchainTransactionUnion interface {
	implementsExchangeProtectionWithdrawalScanParamsOnchainTransactionUnion()
}

type ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransaction struct {
	Amount param.Field[float64]                                                                                                                  `json:"amount,required"`
	Asset  param.Field[ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionAsset] `json:"asset,required"`
	// An enumeration.
	Chain     param.Field[ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionChain] `json:"chain,required"`
	ToAddress param.Field[string]                                                                                                                   `json:"to_address,required"`
}

func (r ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransaction) implementsExchangeProtectionWithdrawalScanParamsOnchainTransactionUnion() {
}

type ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionAsset string

const (
	ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionAssetEth ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionAsset = "ETH"
	ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionAssetPol ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionAsset = "POL"
)

func (r ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionAsset) IsKnown() bool {
	switch r {
	case ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionAssetEth, ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionAssetPol:
		return true
	}
	return false
}

// An enumeration.
type ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionChain string

const (
	ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionChainEthereum ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionChain = "ethereum"
	ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionChainBase     ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionChain = "base"
	ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionChainArbitrum ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionChain = "arbitrum"
	ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionChainOptimism ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionChain = "optimism"
	ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionChainPolygon  ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionChain = "polygon"
)

func (r ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionChain) IsKnown() bool {
	switch r {
	case ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionChainEthereum, ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionChainBase, ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionChainArbitrum, ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionChainOptimism, ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestEvmOnchainTransactionChainPolygon:
		return true
	}
	return false
}

type ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestBitcoinOnchainTransaction struct {
	Amount    param.Field[float64]                                                                                                                      `json:"amount,required"`
	Asset     param.Field[ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestBitcoinOnchainTransactionAsset] `json:"asset,required"`
	Chain     param.Field[ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestBitcoinOnchainTransactionChain] `json:"chain,required"`
	ToAddress param.Field[string]                                                                                                                       `json:"to_address,required"`
}

func (r ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestBitcoinOnchainTransaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestBitcoinOnchainTransaction) implementsExchangeProtectionWithdrawalScanParamsOnchainTransactionUnion() {
}

type ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestBitcoinOnchainTransactionAsset string

const (
	ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestBitcoinOnchainTransactionAssetBtc ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestBitcoinOnchainTransactionAsset = "BTC"
)

func (r ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestBitcoinOnchainTransactionAsset) IsKnown() bool {
	switch r {
	case ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestBitcoinOnchainTransactionAssetBtc:
		return true
	}
	return false
}

type ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestBitcoinOnchainTransactionChain string

const (
	ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestBitcoinOnchainTransactionChainBitcoin ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestBitcoinOnchainTransactionChain = "bitcoin"
)

func (r ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestBitcoinOnchainTransactionChain) IsKnown() bool {
	switch r {
	case ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestBitcoinOnchainTransactionChainBitcoin:
		return true
	}
	return false
}

type ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestStellarOnchainTransaction struct {
	Amount    param.Field[float64]                                                                                                                      `json:"amount,required"`
	Asset     param.Field[ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestStellarOnchainTransactionAsset] `json:"asset,required"`
	Chain     param.Field[ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestStellarOnchainTransactionChain] `json:"chain,required"`
	ToAddress param.Field[string]                                                                                                                       `json:"to_address,required"`
}

func (r ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestStellarOnchainTransaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestStellarOnchainTransaction) implementsExchangeProtectionWithdrawalScanParamsOnchainTransactionUnion() {
}

type ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestStellarOnchainTransactionAsset string

const (
	ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestStellarOnchainTransactionAssetXlm ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestStellarOnchainTransactionAsset = "XLM"
)

func (r ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestStellarOnchainTransactionAsset) IsKnown() bool {
	switch r {
	case ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestStellarOnchainTransactionAssetXlm:
		return true
	}
	return false
}

type ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestStellarOnchainTransactionChain string

const (
	ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestStellarOnchainTransactionChainStellar ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestStellarOnchainTransactionChain = "stellar"
)

func (r ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestStellarOnchainTransactionChain) IsKnown() bool {
	switch r {
	case ExchangeProtectionWithdrawalScanParamsOnchainTransactionRoutersExchangeProtectionModelsRequestStellarOnchainTransactionChainStellar:
		return true
	}
	return false
}

type ExchangeProtectionWithdrawalScanParamsOnchainTransactionAsset string

const (
	ExchangeProtectionWithdrawalScanParamsOnchainTransactionAssetEth ExchangeProtectionWithdrawalScanParamsOnchainTransactionAsset = "ETH"
	ExchangeProtectionWithdrawalScanParamsOnchainTransactionAssetPol ExchangeProtectionWithdrawalScanParamsOnchainTransactionAsset = "POL"
	ExchangeProtectionWithdrawalScanParamsOnchainTransactionAssetBtc ExchangeProtectionWithdrawalScanParamsOnchainTransactionAsset = "BTC"
	ExchangeProtectionWithdrawalScanParamsOnchainTransactionAssetXlm ExchangeProtectionWithdrawalScanParamsOnchainTransactionAsset = "XLM"
)

func (r ExchangeProtectionWithdrawalScanParamsOnchainTransactionAsset) IsKnown() bool {
	switch r {
	case ExchangeProtectionWithdrawalScanParamsOnchainTransactionAssetEth, ExchangeProtectionWithdrawalScanParamsOnchainTransactionAssetPol, ExchangeProtectionWithdrawalScanParamsOnchainTransactionAssetBtc, ExchangeProtectionWithdrawalScanParamsOnchainTransactionAssetXlm:
		return true
	}
	return false
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
	ExchangeProtectionWithdrawalScanParamsOnchainTransactionChainStellar  ExchangeProtectionWithdrawalScanParamsOnchainTransactionChain = "stellar"
)

func (r ExchangeProtectionWithdrawalScanParamsOnchainTransactionChain) IsKnown() bool {
	switch r {
	case ExchangeProtectionWithdrawalScanParamsOnchainTransactionChainEthereum, ExchangeProtectionWithdrawalScanParamsOnchainTransactionChainBase, ExchangeProtectionWithdrawalScanParamsOnchainTransactionChainArbitrum, ExchangeProtectionWithdrawalScanParamsOnchainTransactionChainOptimism, ExchangeProtectionWithdrawalScanParamsOnchainTransactionChainPolygon, ExchangeProtectionWithdrawalScanParamsOnchainTransactionChainBitcoin, ExchangeProtectionWithdrawalScanParamsOnchainTransactionChainStellar:
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
