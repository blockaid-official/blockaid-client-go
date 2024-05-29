// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsTransactionScanResponseEventsParamsValueUnion() {}
func (UnionString) ImplementsEvmJsonRpcScanParamsChainUnion()                {}
func (UnionString) ImplementsEvmTransactionScanParamsChainUnion()            {}
func (UnionString) ImplementsEvmTransactionBulkScanParamsChainUnion()        {}
func (UnionString) ImplementsEvmTransactionRawScanParamsChainUnion()         {}
func (UnionString) ImplementsEvmUserOperationScanParamsChainUnion()          {}
func (UnionString) ImplementsTokenScanParamsAddressUnion()                   {}
