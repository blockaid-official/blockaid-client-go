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

func TestEvmAddressBulkScan(t *testing.T) {
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
	_, err := client.Evm.AddressBulk.Scan(context.TODO(), blockaidclientgo.EvmAddressBulkScanParams{
		ValidateBulkAddresses: blockaidclientgo.ValidateBulkAddressesParam{
			Addresses: blockaidclientgo.F([]string{"0xb85492afC686d5CA405E3CD4f50b05D358c75Ede", "0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97", "0x1f9090aaE28b8a3dCeaDf281B0F12828e676c326", "0xD6E4aA932147A3FE5311dA1b67D9e73da06F9cEf"}),
			Chain:     blockaidclientgo.F(blockaidclientgo.TransactionScanSupportedChainEthereum),
			Metadata: blockaidclientgo.F(blockaidclientgo.ValidateBulkAddressesMetadataParam{
				Domain: blockaidclientgo.F("www.example.xyz"),
			}),
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

func TestEvmAddressBulkScanExtended(t *testing.T) {
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
	_, err := client.Evm.AddressBulk.ScanExtended(context.TODO(), blockaidclientgo.EvmAddressBulkScanExtendedParams{
		ValidateBulkExtendedAddressesRequest: blockaidclientgo.ValidateBulkExtendedAddressesRequestParam{
			Addresses: blockaidclientgo.F([]string{"0xb85492afC686d5CA405E3CD4f50b05D358c75Ede", "0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97", "0x1f9090aaE28b8a3dCeaDf281B0F12828e676c326", "0xD6E4aA932147A3FE5311dA1b67D9e73da06F9cEf"}),
			Chain:     blockaidclientgo.F(blockaidclientgo.TransactionScanSupportedChainEthereum),
			Metadata: blockaidclientgo.F(blockaidclientgo.ValidateBulkExtendedAddressesRequestMetadataParam{
				Account: blockaidclientgo.F(blockaidclientgo.ValidateBulkExtendedAddressesRequestMetadataAccountParam{
					AccountID:       blockaidclientgo.F("user123"),
					UserCountryCode: blockaidclientgo.F("US"),
					AgeInYears:      blockaidclientgo.F(int64(3)),
					Created:         blockaidclientgo.F(time.Now()),
				}),
				Connection: blockaidclientgo.F(blockaidclientgo.ValidateBulkExtendedAddressesRequestMetadataConnectionParam{
					IPAddress: blockaidclientgo.F("192.168.1.1"),
					UserAgent: blockaidclientgo.F("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36"),
				}),
			}),
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
