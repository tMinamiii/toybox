package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type (
	Data struct {
		Content string `json:"content"`
	}

	Response struct {
		Data []Data `json:"data"`
	}
)

const MaxLimit = 3

const TimeOutDuration = 3 * time.Second

type (
	SearchImpl struct {
		url     string
		client  *http.Client
		timeout time.Duration
	}
)

func (r *SearchImpl) Search(ctx context.Context, q string) (*Response, error) {
	var (
		body  []byte
		req   *http.Request
		err   error
		tries = 0
	)

	for tries < MaxLimit {
		req, err = r.buildRequest(q)
		if err != nil {
			return nil, err
		}

		body, err = r.doRequest(ctx, req)
		if err != nil {
			tries++
			continue
		}

		resp := &Response{}
		err = json.Unmarshal(body, resp)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return resp, nil
	}

	return nil, errors.Wrap(err, "retry limit exceeded")
}

func (r *SearchImpl) buildRequest(q string) (*http.Request, error) {
	vs := make(url.Values, 6)
	vs.Add("keyword", q)

	req, err := http.NewRequest(http.MethodPost, r.url, strings.NewReader(vs.Encode()))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("accept", "application/json")

	return req, nil
}

func (r *SearchImpl) doRequest(ctx context.Context, req *http.Request) ([]byte, error) {
	reqCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	reqWithCtx := req.WithContext(reqCtx)
	res, err := r.client.Do(reqWithCtx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return bodyBytes, nil
	}

	switch {
	case res.StatusCode/100 == 4:
		return nil, errors.Errorf("error occurred. status code = %d", res.StatusCode)
	case res.StatusCode/100 == 5:
		return nil, errors.Errorf("error occurred. status code = %d", res.StatusCode)
	default:
		return nil, errors.Errorf("error occurred. status code = %d", res.StatusCode)
	}
}
