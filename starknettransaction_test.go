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

func TestStarknetTransactionScanWithOptionalParams(t *testing.T) {
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
	_, err := client.Starknet.Transaction.Scan(context.TODO(), blockaidclientgo.StarknetTransactionScanParams{
		AccountAddress: blockaidclientgo.F("account_address"),
		Chain:          blockaidclientgo.F(blockaidclientgo.StarknetTransactionScanParamsChainMainnet),
		Metadata: blockaidclientgo.F[blockaidclientgo.StarknetTransactionScanParamsMetadataUnion](blockaidclientgo.StarknetTransactionScanParamsMetadataStarknetWalletRequestMetadata{
			Type: blockaidclientgo.F(blockaidclientgo.StarknetTransactionScanParamsMetadataStarknetWalletRequestMetadataTypeWallet),
			URL:  blockaidclientgo.F("url"),
		}),
		Transaction: blockaidclientgo.F[blockaidclientgo.StarknetTransactionScanParamsTransactionUnion](blockaidclientgo.StarknetTransactionScanParamsTransactionStarknetInvokeV1TransactionSchema{
			MaxFee:        blockaidclientgo.F("max_fee"),
			Nonce:         blockaidclientgo.F("nonce"),
			SenderAddress: blockaidclientgo.F("sender_address"),
			Version:       blockaidclientgo.F(blockaidclientgo.StarknetTransactionScanParamsTransactionStarknetInvokeV1TransactionSchemaVersion1),
			Calldata:      blockaidclientgo.F([]string{"string", "string", "string"}),
		}),
		Options: blockaidclientgo.F([]blockaidclientgo.StarknetTransactionScanParamsOption{blockaidclientgo.StarknetTransactionScanParamsOptionValidation}),
	})
	if err != nil {
		var apierr *blockaidclientgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}