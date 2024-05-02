// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaid_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stainless-sdks/blockaid-client-go"
	"github.com/stainless-sdks/blockaid-client-go/internal"
	"github.com/stainless-sdks/blockaid-client-go/option"
)

type closureTransport struct {
	fn func(req *http.Request) (*http.Response, error)
}

func (t *closureTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.fn(req)
}

func TestUserAgentHeader(t *testing.T) {
	var userAgent string
	client := blockaid.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					userAgent = req.Header.Get("User-Agent")
					return &http.Response{
						StatusCode: http.StatusOK,
					}, nil
				},
			},
		}),
	)
	client.Evm.JsonRpc.Scan(context.Background(), blockaid.EvmJsonRpcScanParams{
		Chain: blockaid.F("ethereum"),
		Data: blockaid.F(blockaid.EvmJsonRpcScanParamsData{
			Method: blockaid.F("eth_signTypedData_v4"),
			Params: blockaid.F([]interface{}{"0x49c73c9d361c04769a452E85D343b41aC38e0EE4", "{\"domain\":{\"chainId\":1,\"name\":\"Aave interest bearing WETH\",\"version\":\"1\",\"verifyingContract\":\"0x030ba81f1c18d280636f32af80b9aad02cf0854e\"},\"message\":{\"owner\":\"0x49c73c9d361c04769a452E85D343b41aC38e0EE4\",\"spender\":\"0xa74cbd5b80f73b5950768c8dc467f1c6307c00fd\",\"value\":\"115792089237316195423570985008687907853269984665640564039457584007913129639935\",\"nonce\":\"0\",\"deadline\":\"1988064000\",\"holder\":\"0x49c73c9d361c04769a452E85D343b41aC38e0EE4\"},\"primaryType\":\"Permit\",\"types\":{\"EIP712Domain\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"version\",\"type\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\"}],\"Permit\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\"}]}}"}),
		}),
		Metadata: blockaid.F(blockaid.MetadataParam{
			Domain: blockaid.F("https://boredapeyartclub.com"),
		}),
	})
	if userAgent != fmt.Sprintf("Blockaid/Go %s", internal.PackageVersion) {
		t.Errorf("Expected User-Agent to be correct, but got: %#v", userAgent)
	}
}

func TestRetryAfter(t *testing.T) {
	attempts := 0
	client := blockaid.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					attempts++
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After"): []string{"0.1"},
						},
					}, nil
				},
			},
		}),
	)
	res, err := client.Evm.JsonRpc.Scan(context.Background(), blockaid.EvmJsonRpcScanParams{
		Chain: blockaid.F("ethereum"),
		Data: blockaid.F(blockaid.EvmJsonRpcScanParamsData{
			Method: blockaid.F("eth_signTypedData_v4"),
			Params: blockaid.F([]interface{}{"0x49c73c9d361c04769a452E85D343b41aC38e0EE4", "{\"domain\":{\"chainId\":1,\"name\":\"Aave interest bearing WETH\",\"version\":\"1\",\"verifyingContract\":\"0x030ba81f1c18d280636f32af80b9aad02cf0854e\"},\"message\":{\"owner\":\"0x49c73c9d361c04769a452E85D343b41aC38e0EE4\",\"spender\":\"0xa74cbd5b80f73b5950768c8dc467f1c6307c00fd\",\"value\":\"115792089237316195423570985008687907853269984665640564039457584007913129639935\",\"nonce\":\"0\",\"deadline\":\"1988064000\",\"holder\":\"0x49c73c9d361c04769a452E85D343b41aC38e0EE4\"},\"primaryType\":\"Permit\",\"types\":{\"EIP712Domain\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"version\",\"type\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\"}],\"Permit\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\"}]}}"}),
		}),
		Metadata: blockaid.F(blockaid.MetadataParam{
			Domain: blockaid.F("https://boredapeyartclub.com"),
		}),
	})
	if err == nil || res != nil {
		t.Error("Expected there to be a cancel error and for the response to be nil")
	}
	if want := 3; attempts != want {
		t.Errorf("Expected %d attempts, got %d", want, attempts)
	}
}

func TestRetryAfterMs(t *testing.T) {
	attempts := 0
	client := blockaid.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					attempts++
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After-Ms"): []string{"100"},
						},
					}, nil
				},
			},
		}),
	)
	res, err := client.Evm.JsonRpc.Scan(context.Background(), blockaid.EvmJsonRpcScanParams{
		Chain: blockaid.F("ethereum"),
		Data: blockaid.F(blockaid.EvmJsonRpcScanParamsData{
			Method: blockaid.F("eth_signTypedData_v4"),
			Params: blockaid.F([]interface{}{"0x49c73c9d361c04769a452E85D343b41aC38e0EE4", "{\"domain\":{\"chainId\":1,\"name\":\"Aave interest bearing WETH\",\"version\":\"1\",\"verifyingContract\":\"0x030ba81f1c18d280636f32af80b9aad02cf0854e\"},\"message\":{\"owner\":\"0x49c73c9d361c04769a452E85D343b41aC38e0EE4\",\"spender\":\"0xa74cbd5b80f73b5950768c8dc467f1c6307c00fd\",\"value\":\"115792089237316195423570985008687907853269984665640564039457584007913129639935\",\"nonce\":\"0\",\"deadline\":\"1988064000\",\"holder\":\"0x49c73c9d361c04769a452E85D343b41aC38e0EE4\"},\"primaryType\":\"Permit\",\"types\":{\"EIP712Domain\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"version\",\"type\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\"}],\"Permit\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\"}]}}"}),
		}),
		Metadata: blockaid.F(blockaid.MetadataParam{
			Domain: blockaid.F("https://boredapeyartclub.com"),
		}),
	})
	if err == nil || res != nil {
		t.Error("Expected there to be a cancel error and for the response to be nil")
	}
	if want := 3; attempts != want {
		t.Errorf("Expected %d attempts, got %d", want, attempts)
	}
}

func TestContextCancel(t *testing.T) {
	client := blockaid.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					<-req.Context().Done()
					return nil, req.Context().Err()
				},
			},
		}),
	)
	cancelCtx, cancel := context.WithCancel(context.Background())
	cancel()
	res, err := client.Evm.JsonRpc.Scan(cancelCtx, blockaid.EvmJsonRpcScanParams{
		Chain: blockaid.F("ethereum"),
		Data: blockaid.F(blockaid.EvmJsonRpcScanParamsData{
			Method: blockaid.F("eth_signTypedData_v4"),
			Params: blockaid.F([]interface{}{"0x49c73c9d361c04769a452E85D343b41aC38e0EE4", "{\"domain\":{\"chainId\":1,\"name\":\"Aave interest bearing WETH\",\"version\":\"1\",\"verifyingContract\":\"0x030ba81f1c18d280636f32af80b9aad02cf0854e\"},\"message\":{\"owner\":\"0x49c73c9d361c04769a452E85D343b41aC38e0EE4\",\"spender\":\"0xa74cbd5b80f73b5950768c8dc467f1c6307c00fd\",\"value\":\"115792089237316195423570985008687907853269984665640564039457584007913129639935\",\"nonce\":\"0\",\"deadline\":\"1988064000\",\"holder\":\"0x49c73c9d361c04769a452E85D343b41aC38e0EE4\"},\"primaryType\":\"Permit\",\"types\":{\"EIP712Domain\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"version\",\"type\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\"}],\"Permit\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\"}]}}"}),
		}),
		Metadata: blockaid.F(blockaid.MetadataParam{
			Domain: blockaid.F("https://boredapeyartclub.com"),
		}),
	})
	if err == nil || res != nil {
		t.Error("Expected there to be a cancel error and for the response to be nil")
	}
}

func TestContextCancelDelay(t *testing.T) {
	client := blockaid.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					<-req.Context().Done()
					return nil, req.Context().Err()
				},
			},
		}),
	)
	cancelCtx, cancel := context.WithTimeout(context.Background(), 2*time.Millisecond)
	defer cancel()
	res, err := client.Evm.JsonRpc.Scan(cancelCtx, blockaid.EvmJsonRpcScanParams{
		Chain: blockaid.F("ethereum"),
		Data: blockaid.F(blockaid.EvmJsonRpcScanParamsData{
			Method: blockaid.F("eth_signTypedData_v4"),
			Params: blockaid.F([]interface{}{"0x49c73c9d361c04769a452E85D343b41aC38e0EE4", "{\"domain\":{\"chainId\":1,\"name\":\"Aave interest bearing WETH\",\"version\":\"1\",\"verifyingContract\":\"0x030ba81f1c18d280636f32af80b9aad02cf0854e\"},\"message\":{\"owner\":\"0x49c73c9d361c04769a452E85D343b41aC38e0EE4\",\"spender\":\"0xa74cbd5b80f73b5950768c8dc467f1c6307c00fd\",\"value\":\"115792089237316195423570985008687907853269984665640564039457584007913129639935\",\"nonce\":\"0\",\"deadline\":\"1988064000\",\"holder\":\"0x49c73c9d361c04769a452E85D343b41aC38e0EE4\"},\"primaryType\":\"Permit\",\"types\":{\"EIP712Domain\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"version\",\"type\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\"}],\"Permit\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\"}]}}"}),
		}),
		Metadata: blockaid.F(blockaid.MetadataParam{
			Domain: blockaid.F("https://boredapeyartclub.com"),
		}),
	})
	if err == nil || res != nil {
		t.Error("expected there to be a cancel error and for the response to be nil")
	}
}

func TestContextDeadline(t *testing.T) {
	testTimeout := time.After(3 * time.Second)
	testDone := make(chan struct{})

	deadline := time.Now().Add(100 * time.Millisecond)
	deadlineCtx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go func() {
		client := blockaid.NewClient(
			option.WithHTTPClient(&http.Client{
				Transport: &closureTransport{
					fn: func(req *http.Request) (*http.Response, error) {
						<-req.Context().Done()
						return nil, req.Context().Err()
					},
				},
			}),
		)
		res, err := client.Evm.JsonRpc.Scan(deadlineCtx, blockaid.EvmJsonRpcScanParams{
			Chain: blockaid.F("ethereum"),
			Data: blockaid.F(blockaid.EvmJsonRpcScanParamsData{
				Method: blockaid.F("eth_signTypedData_v4"),
				Params: blockaid.F([]interface{}{"0x49c73c9d361c04769a452E85D343b41aC38e0EE4", "{\"domain\":{\"chainId\":1,\"name\":\"Aave interest bearing WETH\",\"version\":\"1\",\"verifyingContract\":\"0x030ba81f1c18d280636f32af80b9aad02cf0854e\"},\"message\":{\"owner\":\"0x49c73c9d361c04769a452E85D343b41aC38e0EE4\",\"spender\":\"0xa74cbd5b80f73b5950768c8dc467f1c6307c00fd\",\"value\":\"115792089237316195423570985008687907853269984665640564039457584007913129639935\",\"nonce\":\"0\",\"deadline\":\"1988064000\",\"holder\":\"0x49c73c9d361c04769a452E85D343b41aC38e0EE4\"},\"primaryType\":\"Permit\",\"types\":{\"EIP712Domain\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"version\",\"type\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\"}],\"Permit\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\"}]}}"}),
			}),
			Metadata: blockaid.F(blockaid.MetadataParam{
				Domain: blockaid.F("https://boredapeyartclub.com"),
			}),
		})
		if err == nil || res != nil {
			t.Error("expected there to be a deadline error and for the response to be nil")
		}
		close(testDone)
	}()

	select {
	case <-testTimeout:
		t.Fatal("client didn't finish in time")
	case <-testDone:
		if diff := time.Since(deadline); diff < -30*time.Millisecond || 30*time.Millisecond < diff {
			t.Fatalf("client did not return within 30ms of context deadline, got %s", diff)
		}
	}
}
