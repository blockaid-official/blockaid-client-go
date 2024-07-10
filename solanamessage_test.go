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
		Transactions: blockaidclientgo.F([]string{"vxBNpvao9QJmLKXUThbbjRnxm3ufu4Wku97kHd5a67FDjSqeHwcPrBKTjAHp4ECr61eWwoxvUEVTuuWX65P9bCNDJrTJpX64vjdtpHA8cogA4C92Ubj813wUUA8Ey4Bvcrdj5c1bSTCv27zKyx1AHWDepVVoS5ZV2Sb3Nuw8RGrmjsZgU3hvPzE9hRBosY25Xpbyqo4b3Vr1BLfrVRBqsz7PvB74APZ7dHxfH49Xb2edrFS2DZ84SwtsZYLyTGF5wtZ6WHWiZN3ixjKGMAh5NLNmT9imKMBgtxuTMAw"}),
		Chain:        blockaidclientgo.F("mainnet"),
		Encoding:     blockaidclientgo.F("base58"),
		Method:       blockaidclientgo.F("signAndSendTransaction"),
	})
	if err != nil {
		var apierr *blockaidclientgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
