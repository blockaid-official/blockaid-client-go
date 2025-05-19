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

func TestSuiTransactionScanWithOptionalParams(t *testing.T) {
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
	_, err := client.Sui.Transaction.Scan(context.TODO(), blockaidclientgo.SuiTransactionScanParams{
		AccountAddress: blockaidclientgo.F("0x45e90b3ea2e1920c43d92d224630d6a865c1b58a7b4e770c2ac156eab30eb491"),
		Chain:          blockaidclientgo.F(blockaidclientgo.SuiTransactionScanParamsChainMainnet),
		Metadata: blockaidclientgo.F[blockaidclientgo.SuiTransactionScanParamsMetadataUnion](blockaidclientgo.SuiTransactionScanParamsMetadataSuiWalletRequestMetadata{
			Type: blockaidclientgo.F(blockaidclientgo.SuiTransactionScanParamsMetadataSuiWalletRequestMetadataTypeWallet),
			URL:  blockaidclientgo.F("localhost"),
		}),
		Transaction: blockaidclientgo.F("AAACAAgA4fUFAAAAAAAgHvls2mKzo/48s/fPdWP8xKtE4BhIjR2O8gMaZ6bI1+sCAgABAQAAAQECAAABAQBF6Qs+ouGSDEPZLSJGMNaoZcG1intOdwwqwVbqsw60kQFySkLceU6uis9QxxK4CDYqttqK3ilc9/yEcCgxdaeA0cl/xhwAAAAAIEuXU9TpAtIJmbPVFpxdc70+RWUqlSrfyIUKT9q1Au0ERekLPqLhkgxD2S0iRjDWqGXBtYp7TncMKsFW6rMOtJHuAgAAAAAAACAKNQAAAAAAAA=="),
		Options:     blockaidclientgo.F([]blockaidclientgo.SuiTransactionScanParamsOption{blockaidclientgo.SuiTransactionScanParamsOptionSimulation}),
	})
	if err != nil {
		var apierr *blockaidclientgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
