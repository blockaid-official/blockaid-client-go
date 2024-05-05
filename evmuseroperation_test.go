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

func TestEvmUserOperationScanWithOptionalParams(t *testing.T) {
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
	_, err := client.Evm.UserOperation.Scan(context.TODO(), blockaidclientgo.EvmUserOperationScanParams{
		Chain: blockaidclientgo.F(blockaidclientgo.SupportedChainBaseSepolia),
		Data: blockaidclientgo.F(blockaidclientgo.EvmUserOperationScanParamsData{
			Operation: blockaidclientgo.F(blockaidclientgo.EvmUserOperationScanParamsDataOperation{
				Sender:               blockaidclientgo.F("0x77bA5AC3ca4864be26CA3112baDf07286CcC3324"),
				Nonce:                blockaidclientgo.F("0x22"),
				InitCode:             blockaidclientgo.F("0x"),
				CallData:             blockaidclientgo.F("0x51945447000000000000000000000000aeed57a826a998f9388ce2fd6cdb0b6aa75e3d190000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000044095ea7b300000000000000000000000050a9266605ba303b659ff105919205570f2af971000000000000000000000000000000000000000000000000000000000000006400000000000000000000000000000000000000000000000000000000"),
				CallGasLimit:         blockaidclientgo.F("0x3c38"),
				VerificationGasLimit: blockaidclientgo.F("0x1659f"),
				PreVerificationGas:   blockaidclientgo.F("0x2496ebc"),
				MaxFeePerGas:         blockaidclientgo.F("0x218fe7"),
				MaxPriorityFeePerGas: blockaidclientgo.F("0xf4240"),
				PaymasterAndData:     blockaidclientgo.F("0x9d6ac51b972544251fcc0f2902e633e3f9bd3f290000000000000000000000000000000000000000000000000000000065cc4c990000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001eb8343d03ec9fc28a877c2fcad21d9923c56e6ad156ea6647282a35ce215c931f9fbdf3bec37168f9c9b49e33a0818731c5892ff626852f9465e619538540221c"),
				Signature:            blockaidclientgo.F("0x"),
			}),
			Entrypoint: blockaidclientgo.F("0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789"),
		}),
		Metadata: blockaidclientgo.F(blockaidclientgo.MetadataParam{
			Domain: blockaidclientgo.F("example.com"),
		}),
		AccountAddress: blockaidclientgo.F("0x77bA5AC3ca4864be26CA3112baDf07286CcC3324"),
		Options:        blockaidclientgo.F([]blockaidclientgo.EvmUserOperationScanParamsOption{blockaidclientgo.EvmUserOperationScanParamsOptionSimulation, blockaidclientgo.EvmUserOperationScanParamsOptionValidation}),
	})
	if err != nil {
		var apierr *blockaidclientgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
