package skyroom

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var (
	ErrUnexpectedConnectionError = errors.New("unexpected connection error")
	ErrResponseDecodeError       = errors.New("couldn't decode response")
	ErrBadRequestError           = errors.New("bad request")
)

type baseRequest struct {
	Action string      `json:"action"`
	Params interface{} `json:"params"`
}

type BaseResponse struct {
	Ok           bool            `json:"ok"`
	ErrorCode    int             `json:"error_code"`
	ErrorMessage string          `json:"error_message"`
	Result       json.RawMessage `json:"result"`
}

type Result json.RawMessage

func (sky *Skyroom) Post(serviceName string, data interface{}) (json.RawMessage, error) {
	endpoint, _ := url.JoinPath(sky.BaseURL.String(), sky.APIKey)

	request := baseRequest{
		Action: serviceName,
		Params: data,
	}

	marshaledRequest, _ := json.Marshal(request)

	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(marshaledRequest))

	if err != nil {
		fmt.Println(err)
		return nil, ErrUnexpectedConnectionError
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	body, _ := io.ReadAll(resp.Body)

	if (resp.StatusCode != http.StatusOK) && (resp.StatusCode != http.StatusCreated) {
		fmt.Println(ErrUnexpectedConnectionError)
		return nil, ErrUnexpectedConnectionError
	}

	unmarshalledResp := &BaseResponse{}

	if err := json.Unmarshal(body, unmarshalledResp); err != nil {
		fmt.Println(err)
		return nil, ErrResponseDecodeError
	}

	if !unmarshalledResp.Ok {
		fmt.Println(unmarshalledResp.ErrorCode, unmarshalledResp.ErrorMessage)
		return nil, ErrBadRequestError
	}

	result := unmarshalledResp.Result

	return result, nil
}
