# Evm

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#MetadataParam">MetadataParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#AddressAssetDiff">AddressAssetDiff</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#AddressAssetExposure">AddressAssetExposure</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#Erc1155Diff">Erc1155Diff</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#Erc1155Exposure">Erc1155Exposure</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#Erc1155TokenDetails">Erc1155TokenDetails</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#Erc20Diff">Erc20Diff</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#Erc20Exposure">Erc20Exposure</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#Erc20TokenDetails">Erc20TokenDetails</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#Erc721Diff">Erc721Diff</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#Erc721Exposure">Erc721Exposure</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#Erc721TokenDetails">Erc721TokenDetails</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#NativeDiff">NativeDiff</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#NonercTokenDetails">NonercTokenDetails</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#TransactionBulkResponse">TransactionBulkResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#TransactionScanFeature">TransactionScanFeature</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#TransactionScanResponse">TransactionScanResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#TransactionSimulation">TransactionSimulation</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#TransactionValidation">TransactionValidation</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#UsdDiff">UsdDiff</a>

## JsonRpc

Methods:

- <code title="post /evm/json-rpc/scan">client.Evm.JsonRpc.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#EvmJsonRpcService.Scan">Scan</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#EvmJsonRpcScanParams">EvmJsonRpcScanParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#TransactionScanResponse">TransactionScanResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Transaction

Methods:

- <code title="post /evm/transaction/scan">client.Evm.Transaction.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#EvmTransactionService.Scan">Scan</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#EvmTransactionScanParams">EvmTransactionScanParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#TransactionScanResponse">TransactionScanResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## TransactionBulk

Methods:

- <code title="post /evm/transaction-bulk/scan">client.Evm.TransactionBulk.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#EvmTransactionBulkService.Scan">Scan</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#EvmTransactionBulkScanParams">EvmTransactionBulkScanParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#TransactionBulkResponse">TransactionBulkResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## TransactionRaw

Methods:

- <code title="post /evm/transaction-raw/scan">client.Evm.TransactionRaw.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#EvmTransactionRawService.Scan">Scan</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#EvmTransactionRawScanParams">EvmTransactionRawScanParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#TransactionScanResponse">TransactionScanResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## UserOperation

# Site

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#SiteScanHitResponse">SiteScanHitResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#SiteScanMissResponse">SiteScanMissResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#SiteScanResponse">SiteScanResponse</a>

Methods:

- <code title="post /site/scan">client.Site.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#SiteService.Scan">Scan</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#SiteScanParams">SiteScanParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go">blockaid</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blockaid-go#SiteScanResponse">SiteScanResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
