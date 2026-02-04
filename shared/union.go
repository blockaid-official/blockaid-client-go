// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsUserOperationRequestBlockUnionParam()                      {}
func (UnionString) ImplementsEvmJsonRpcScanResponseEventsParamsValueUnion()             {}
func (UnionString) ImplementsEvmJsonRpcScanParamsBlockUnion()                           {}
func (UnionString) ImplementsEvmTransactionScanResponseEventsParamsValueUnion()         {}
func (UnionString) ImplementsEvmTransactionScanParamsBlockUnion()                       {}
func (UnionString) ImplementsEvmTransactionBulkScanResponseEventsParamsValueUnion()     {}
func (UnionString) ImplementsEvmTransactionBulkScanParamsBlockUnion()                   {}
func (UnionString) ImplementsEvmTransactionRawScanResponseEventsParamsValueUnion()      {}
func (UnionString) ImplementsEvmTransactionRawScanParamsBlockUnion()                    {}
func (UnionString) ImplementsEvmUserOperationScanResponseEventsParamsValueUnion()       {}
func (UnionString) ImplementsEvmPostTransactionScanResponseEventsParamsValueUnion()     {}
func (UnionString) ImplementsEvmPostTransactionScanParamsBlockUnion()                   {}
func (UnionString) ImplementsEvmPostTransactionBulkScanResponseEventsParamsValueUnion() {}
func (UnionString) ImplementsEvmPostTransactionBulkScanParamsBlockUnion()               {}

type UnionInt int64

func (UnionInt) ImplementsUserOperationRequestBlockUnionParam()        {}
func (UnionInt) ImplementsEvmJsonRpcScanParamsBlockUnion()             {}
func (UnionInt) ImplementsEvmTransactionScanParamsBlockUnion()         {}
func (UnionInt) ImplementsEvmTransactionBulkScanParamsBlockUnion()     {}
func (UnionInt) ImplementsEvmTransactionRawScanParamsBlockUnion()      {}
func (UnionInt) ImplementsEvmPostTransactionScanParamsBlockUnion()     {}
func (UnionInt) ImplementsEvmPostTransactionBulkScanParamsBlockUnion() {}
