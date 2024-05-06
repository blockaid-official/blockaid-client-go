// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsTransactionScanResponseEventsParamsValueUnion() {}
func (UnionString) ImplementsTransactionValidationErrorLocUnion()            {}

type UnionInt int64

func (UnionInt) ImplementsTransactionValidationErrorLocUnion() {}
