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

func TestSolanaMessageScanWithOptionalParams(t *testing.T) {
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
	_, err := client.Solana.Message.Scan(context.TODO(), blockaidclientgo.SolanaMessageScanParams{
		AccountAddress: blockaidclientgo.F("86xCnPeV69n6t3DnyGvkKobf9FdN2H9oiVDdaMpo2MMY"),
		Metadata: blockaidclientgo.F(blockaidclientgo.SolanaMessageScanParamsMetadata{
			URL: blockaidclientgo.F("https://example.com"),
		}),
		Transactions: blockaidclientgo.F([]string{"vxBNpvao9QJmLKXUThbbjRnxm3ufu4Wku97kHd5a67FDjSqeHwcPrBKTjAHp4ECr61eWwoxvUEVTuuWX65P9bCNDJrTJpX64vjdtpHA8cogA4C92Ubj813wUUA8Ey4Bvcrdj5c1bSTrGZVzb8QmCKyzMu9kMiSWpFtaFrNN8zb9grr81N3R3njrFgxCxNSjboFtomLyZ3iUQBaBkRF1DyzGyc1r1kd8FnptaDWteNCXJHUYFeH8wBDwZJzNZfz71CiugXhxBTJSAqSNC8JEWm7kmCqwjUqLd23L2x2s"}),
		Chain:        blockaidclientgo.F("mainnet"),
		Encoding:     blockaidclientgo.F(blockaidclientgo.SolanaMessageScanParamsEncodingBase58),
		Method:       blockaidclientgo.F("signAndSendTransaction"),
		Options:      blockaidclientgo.F([]blockaidclientgo.SolanaMessageScanParamsOption{blockaidclientgo.SolanaMessageScanParamsOptionSimulation, blockaidclientgo.SolanaMessageScanParamsOptionValidation}),
	})
	if err != nil {
		var apierr *blockaidclientgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
