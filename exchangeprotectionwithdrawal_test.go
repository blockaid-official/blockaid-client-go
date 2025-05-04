// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/blockaid-official/blockaid-client-go"
	"github.com/blockaid-official/blockaid-client-go/internal/testutil"
	"github.com/blockaid-official/blockaid-client-go/option"
)

func TestExchangeProtectionWithdrawalScanWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blockaidclientgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ExchangeProtection.Withdrawal.Scan(context.TODO(), blockaidclientgo.ExchangeProtectionWithdrawalScanParams{
		Account: blockaidclientgo.F(blockaidclientgo.ExchangeProtectionWithdrawalScanParamsAccount{
			AccountID:       blockaidclientgo.F("account_id"),
			UserCountryCode: blockaidclientgo.F("user_country_code"),
			AgeInYears:      blockaidclientgo.F(int64(1)),
			Created:         blockaidclientgo.F(time.Now()),
		}),
		EventTime: blockaidclientgo.F(time.Now()),
		OnchainTransaction: blockaidclientgo.F(blockaidclientgo.ExchangeProtectionWithdrawalScanParamsOnchainTransaction{
			Amount:    blockaidclientgo.F(1.000000),
			Asset:     blockaidclientgo.F("asset"),
			Chain:     blockaidclientgo.F(blockaidclientgo.ExchangeProtectionWithdrawalScanParamsOnchainTransactionChainEthereum),
			ToAddress: blockaidclientgo.F("to_address"),
		}),
		WithdrawalID: blockaidclientgo.F("withdrawal_id"),
		ConnectionMetadata: blockaidclientgo.F(blockaidclientgo.ExchangeProtectionWithdrawalScanParamsConnectionMetadata{
			CustomerIP: blockaidclientgo.F("customer_ip"),
			UserAgent:  blockaidclientgo.F("user_agent"),
		}),
	})
	if err != nil {
		var apierr *blockaidclientgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
