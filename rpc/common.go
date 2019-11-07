package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Focinfi/ckb-sdk-go/types/errtypes"
)

const (
	jsonrpcVersion = "2.0"
)

type ReqBody struct {
	ID      int64       `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	JSONRPC string      `json:"jsonrpc"`
}

type RespBody struct {
	ID      int64           `json:"id"`
	JSONRPC string          `json:"jsonrpc"`
	Error   interface{}     `json:"error,omitempty"`
	Result  json.RawMessage `json:"result"`
}

func RawHTTPPost(ctx context.Context, url string, method string, params interface{}) ([]byte, error) {
	reqBody := ReqBody{
		ID:      1,
		Method:  method,
		Params:  params,
		JSONRPC: jsonrpcVersion,
	}

	b, err := json.Marshal(reqBody)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrMarshalRequestBodyFail, err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	req = req.WithContext(ctx)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrNewRequestFail, err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrHTTPRequestFail, err)
	}
	defer resp.Body.Close()

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrReadResponseBodyFail, err)
	}

	respBody := &RespBody{}
	if err := json.Unmarshal(respBodyBytes, respBody); err != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrUnmarshalResponseBodyFail, err)
	}
	if respBody.Error != nil {
		return nil, errtypes.WrapErr(errtypes.RPCErrResultErrorFail, errors.New(fmt.Sprint(respBody.Error)))
	}
	return respBody.Result, nil
}
