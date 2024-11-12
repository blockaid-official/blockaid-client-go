// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/blockaid-official/blockaid-client-go"
	"github.com/blockaid-official/blockaid-client-go/internal/testutil"
	"github.com/blockaid-official/blockaid-client-go/option"
)

func TestStellarTransactionReportWithOptionalParams(t *testing.T) {
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
	_, err := client.Stellar.Transaction.Report(context.TODO(), blockaidclientgo.StellarTransactionReportParams{
		Details: blockaidclientgo.F("details"),
		Event:   blockaidclientgo.F(blockaidclientgo.StellarTransactionReportParamsEventShouldBeMalicious),
		Report: blockaidclientgo.F[blockaidclientgo.StellarTransactionReportParamsReportUnion](blockaidclientgo.StellarTransactionReportParamsReportStellarAppealRequestID{
			ID:   blockaidclientgo.F("id"),
			Type: blockaidclientgo.F(blockaidclientgo.StellarTransactionReportParamsReportStellarAppealRequestIDTypeRequestID),
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

func TestStellarTransactionScanWithOptionalParams(t *testing.T) {
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
	_, err := client.Stellar.Transaction.Scan(context.TODO(), blockaidclientgo.StellarTransactionScanParams{
		StellarTransactionScanRequest: blockaidclientgo.StellarTransactionScanRequestParam{
			AccountAddress: blockaidclientgo.F("GDPMFLKUGASUTWBN2XGYYKD27QGHCYH4BUFUTER4L23INYQ4JHDWFOIE"),
			Chain:          blockaidclientgo.F(blockaidclientgo.StellarTransactionScanRequestChainPubnet),
			Metadata: blockaidclientgo.F[blockaidclientgo.StellarTransactionScanRequestMetadataUnionParam](blockaidclientgo.StellarTransactionScanRequestMetadataStellarWalletRequestMetadataParam{
				Type: blockaidclientgo.F(blockaidclientgo.StellarTransactionScanRequestMetadataStellarWalletRequestMetadataTypeWallet),
				URL:  blockaidclientgo.F("localhost"),
			}),
			Transaction: blockaidclientgo.F("AAAAAgAAAADewq1UMCVJ2C3VzYwoevwMcWD8DQtJkjxetobiHEnHYgAAAAEAAAAAAAAAAgAAAAAAAAAAAAAAAQAAAAEAAAAA3sKtVDAlSdgt1c2MKHr8DHFg/A0LSZI8XraG4hxJx2IAAAABAAAAACI40RTBOFEE7uT5mZkoq30mbvxLPJpMUm9cIFHgK9SRAAAAAAAAAAAAmJaAAAAAAAAAAAA="),
			Options:     blockaidclientgo.F([]blockaidclientgo.StellarTransactionScanRequestOption{blockaidclientgo.StellarTransactionScanRequestOptionValidation, blockaidclientgo.StellarTransactionScanRequestOptionSimulation}),
		},
	})
	if err != nil {
		var apierr *blockaidclientgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
