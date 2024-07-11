// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsTransactionScanResponseEventsParamsValueUnion() {}
func (UnionString) ImplementsEvmJsonRpcScanParamsBlockUnion()                {}
func (UnionString) ImplementsEvmTransactionScanParamsBlockUnion()            {}
func (UnionString) ImplementsEvmTransactionBulkScanParamsBlockUnion()        {}
func (UnionString) ImplementsEvmTransactionRawScanParamsBlockUnion()         {}
func (UnionString) ImplementsEvmUserOperationScanParamsBlockUnion()          {}
func (UnionString) ImplementsEvmPostTransactionScanParamsBlockUnion()        {}
func (UnionString) ImplementsEvmPostTransactionBulkScanParamsBlockUnion()    {}

type UnionInt int64

func (UnionInt) ImplementsEvmJsonRpcScanParamsBlockUnion()             {}
func (UnionInt) ImplementsEvmTransactionScanParamsBlockUnion()         {}
func (UnionInt) ImplementsEvmTransactionBulkScanParamsBlockUnion()     {}
func (UnionInt) ImplementsEvmTransactionRawScanParamsBlockUnion()      {}
func (UnionInt) ImplementsEvmUserOperationScanParamsBlockUnion()       {}
func (UnionInt) ImplementsEvmPostTransactionScanParamsBlockUnion()     {}
func (UnionInt) ImplementsEvmPostTransactionBulkScanParamsBlockUnion() {}
