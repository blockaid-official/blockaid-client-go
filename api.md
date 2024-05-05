# Evm

Params Types:

- <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#Chain">Chain</a>
- <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#MetadataParam">MetadataParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#AddressAssetDiff">AddressAssetDiff</a>
- <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#AddressAssetExposure">AddressAssetExposure</a>
- <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#Erc1155Diff">Erc1155Diff</a>
- <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#Erc1155Exposure">Erc1155Exposure</a>
- <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#Erc1155TokenDetails">Erc1155TokenDetails</a>
- <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#Erc20Diff">Erc20Diff</a>
- <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#Erc20Exposure">Erc20Exposure</a>
- <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#Erc20TokenDetails">Erc20TokenDetails</a>
- <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#Erc721Diff">Erc721Diff</a>
- <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#Erc721Exposure">Erc721Exposure</a>
- <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#Erc721TokenDetails">Erc721TokenDetails</a>
- <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#NativeDiff">NativeDiff</a>
- <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#NonercTokenDetails">NonercTokenDetails</a>
- <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#TransactionBulkResponse">TransactionBulkResponse</a>
- <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#TransactionScanFeature">TransactionScanFeature</a>
- <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#TransactionScanResponse">TransactionScanResponse</a>
- <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#TransactionSimulation">TransactionSimulation</a>
- <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#TransactionValidation">TransactionValidation</a>
- <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#UsdDiff">UsdDiff</a>

## JsonRpc

Methods:

- <code title="post /evm/json-rpc/scan">client.Evm.JsonRpc.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#EvmJsonRpcService.Scan">Scan</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#EvmJsonRpcScanParams">EvmJsonRpcScanParams</a>) (<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#TransactionScanResponse">TransactionScanResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Transaction

Methods:

- <code title="post /evm/transaction/scan">client.Evm.Transaction.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#EvmTransactionService.Scan">Scan</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#EvmTransactionScanParams">EvmTransactionScanParams</a>) (<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#TransactionScanResponse">TransactionScanResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## TransactionBulk

Methods:

- <code title="post /evm/transaction-bulk/scan">client.Evm.TransactionBulk.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#EvmTransactionBulkService.Scan">Scan</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#EvmTransactionBulkScanParams">EvmTransactionBulkScanParams</a>) ([]<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#TransactionBulkResponse">TransactionBulkResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## TransactionRaw

Methods:

- <code title="post /evm/transaction-raw/scan">client.Evm.TransactionRaw.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#EvmTransactionRawService.Scan">Scan</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#EvmTransactionRawScanParams">EvmTransactionRawScanParams</a>) (<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#TransactionScanResponse">TransactionScanResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## UserOperation

Methods:

- <code title="post /evm/user-operation/scan">client.Evm.UserOperation.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#EvmUserOperationService.Scan">Scan</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#EvmUserOperationScanParams">EvmUserOperationScanParams</a>) (<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#TransactionScanResponse">TransactionScanResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Site

Response Types:

- <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#SiteScanHitResponse">SiteScanHitResponse</a>
- <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#SiteScanMissResponse">SiteScanMissResponse</a>
- <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#SiteScanResponse">SiteScanResponse</a>

Methods:

- <code title="post /site/scan">client.Site.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#SiteService.Scan">Scan</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#SiteScanParams">SiteScanParams</a>) (<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go">blockaidclientgo</a>.<a href="https://pkg.go.dev/github.com/blockaid-official/blockaid-client-go#SiteScanResponse">SiteScanResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
