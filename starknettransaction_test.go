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
		AccountAddress: blockaidclientgo.F("0x62a2959fa6502b30cbfb51199fbbe72e72ee4f5a86ec754b4172c7d7beb6ff4"),
		Chain:          blockaidclientgo.F(blockaidclientgo.StarknetTransactionScanParamsChainMainnet),
		Metadata: blockaidclientgo.F[blockaidclientgo.StarknetTransactionScanParamsMetadataUnion](blockaidclientgo.StarknetTransactionScanParamsMetadataStarknetWalletRequestMetadata{
			Type: blockaidclientgo.F(blockaidclientgo.StarknetTransactionScanParamsMetadataStarknetWalletRequestMetadataTypeWallet),
			URL:  blockaidclientgo.F("giftnostra.com"),
		}),
		Transaction: blockaidclientgo.F[blockaidclientgo.StarknetTransactionScanParamsTransactionUnion](blockaidclientgo.StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchema{
			Calldata:      blockaidclientgo.F([]string{"0x1", "0x4b33a775b537a39c8960120806e815764f173e4fa76546e6706c31aa1b0ac4a", "0x994f23fef04108984d50a4f870723cd46f95d592ed3de9a13f97d5c55846fb", "0x9", "0x1", "0x53c91253bc9682c04929ca02ed00b3e423f6710d2ee7e0d5ebb06f3ecf368a8", "0x1", "0x62a2959fa6502b30cbfb51199fbbe72e72ee4f5a86ec754b4172c7d7beb6ff4", "0x1", "0x5f612ce", "0x0", "0x0", "0x0"}),
			ChainID:       blockaidclientgo.F("0x534e5f4d41494e"),
			Nonce:         blockaidclientgo.F("0xc"),
			SenderAddress: blockaidclientgo.F("0x1840b3c89a446c74a3962647a2a7fb449d83905c4511027dfa9e099c6886691"),
			Version:       blockaidclientgo.F(blockaidclientgo.StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaVersion3),
			AccountDeploymentData: blockaidclientgo.F([]string{}),
			NonceDataAvailabilityMode: blockaidclientgo.F(blockaidclientgo.StarknetTransactionScanParamsTransactionStarknetInvokeV3TransactionSchemaNonceDataAvailabilityMode0),
			PaymasterData: blockaidclientgo.F([]string{}),
		}),
		BlockNumber: blockaidclientgo.F("0xa12e3"),
		Options:     blockaidclientgo.F([]blockaidclientgo.StarknetTransactionScanParamsOption{blockaidclientgo.StarknetTransactionScanParamsOptionValidation, blockaidclientgo.StarknetTransactionScanParamsOptionSimulation}),
	})
	if err != nil {
		var apierr *blockaidclientgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
