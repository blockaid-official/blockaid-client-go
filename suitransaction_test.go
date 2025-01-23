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
		Transaction: blockaidclientgo.F("AAAMAAhAQ6YNAAAAAAEBd57GsbOhR6dNcrvF/qZ/PPV7Ea+Z6IEcLls/V6rNSQ9gukkWAAAAAAEAAQEAAQEACEBDpg0AAAAAABCGCXcJRGn2EAAAAAAAAAAAAQEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgEAAAAAAAAAAAEB9RRaesNFyoc2z4x2BH0A1tN48w6Bvm9utVcYTZ3pPHjvASkSAAAAAAAAEIYJdwlEafYQAAAAAAAAAAAAAQEAIEXpCz6i4ZIMQ9ktIkYw1qhlwbWKe053DCrBVuqzDrSRACBF6Qs+ouGSDEPZLSJGMNaoZcG1intOdwwqwVbqsw60kQ4CAAEBAAAAvY1EiXggQsb6+tTeS8al4LhKQ8bABkf/1wYtHiu3VJ4FdHJhZGUKZmxhc2hfc3dhcAIHAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIDc3VpA1NVSQAH26NGcuMMsGWx+T46tVMYdo/W/vZsFZQsn3y4RuL5AOcEdXNkYwRVU0RDAAcBAQABAgABAwABBAABBQABBgABBwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIHYmFsYW5jZQxkZXN0cm95X3plcm8BBwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACA3N1aQNTVUkAAQMBAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACBGNvaW4EemVybwEH26NGcuMMsGWx+T46tVMYdo/W/vZsFZQsn3y4RuL5AOcEdXNkYwRVU0RDAAAAvY1EiXggQsb6+tTeS8al4LhKQ8bABkf/1wYtHiu3VJ4FdHJhZGUSc3dhcF9yZWNlaXB0X2RlYnRzAAEDAQACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgRjb2luBXNwbGl0AQcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgNzdWkDU1VJAAIDAAAAAAMEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACBGNvaW4MaW50b19iYWxhbmNlAQcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgNzdWkDU1VJAAECBQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIEY29pbgxpbnRvX2JhbGFuY2UBB9ujRnLjDLBlsfk+OrVTGHaP1v72bBWULJ98uEbi+QDnBHVzZGMEVVNEQwABAwMAAAAAvY1EiXggQsb6+tTeS8al4LhKQ8bABkf/1wYtHiu3VJ4FdHJhZGUQcmVwYXlfZmxhc2hfc3dhcAIHAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIDc3VpA1NVSQAH26NGcuMMsGWx+T46tVMYdo/W/vZsFZQsn3y4RuL5AOcEdXNkYwRVU0RDAAUBAQADAQACAAIGAAIHAAEHAAB9ooXCIzqUefJ/USmzxuUpZCEZhBuZn5KMcZ85x9RTQg5zbGlwcGFnZV9jaGVjaw9hc3NlcnRfc2xpcHBhZ2UCBwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACA3N1aQNTVUkAB9ujRnLjDLBlsfk+OrVTGHaP1v72bBWULJ98uEbi+QDnBHVzZGMEVVNEQwADAQEAAQgAAQkAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACBGNvaW4MZnJvbV9iYWxhbmNlAQfbo0Zy4wywZbH5Pjq1Uxh2j9b+9mwVlCyffLhG4vkA5wR1c2RjBFVTREMAAQMBAAEAAQEDAAAAAAEKAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgRjb2luBXZhbHVlAQfbo0Zy4wywZbH5Pjq1Uxh2j9b+9mwVlCyffLhG4vkA5wR1c2RjBFVTREMAAQMKAAAAAQEDCgAAAAELAEXpCz6i4ZIMQ9ktIkYw1qhlwbWKe053DCrBVuqzDrSRAXJKQtx5Tq6Kz1DHErgINiq22oreKVz3/IRwKDF1p4DRyX/GHAAAAAAgS5dT1OkC0gmZs9UWnF1zvT5FZSqVKt/IhQpP2rUC7QRF6Qs+ouGSDEPZLSJGMNaoZcG1intOdwwqwVbqsw60ke4CAAAAAAAA8M9VAAAAAAAA"),
		Options:     blockaidclientgo.F([]blockaidclientgo.SuiTransactionScanParamsOption{blockaidclientgo.SuiTransactionScanParamsOptionValidation}),
	})
	if err != nil {
		var apierr *blockaidclientgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
