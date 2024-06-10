package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

func CreateContext(method, uri string, form interface{}) (echo.Context, *httptest.ResponseRecorder) {
	var req *http.Request
	if method == echo.GET || method == echo.DELETE {
		req, _ = http.NewRequest(method, uri, nil)
	} else {
		f, _ := json.Marshal(form)
		req, _ = http.NewRequest(method, uri, bytes.NewReader(f))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	}
	rec := httptest.NewRecorder()

	return echo.New().NewContext(req, rec), rec
}
