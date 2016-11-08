package linenotcat

import (
	"io"
	"mime/multipart"
	"net/http"
)

const (
	baseURL = "https://notify-api.line.me/api"
)

type apiRequestBuilder struct {
	token string
}

func (arb *apiRequestBuilder) buildFormNotifyRequest(body io.Reader) (*http.Request, error) {
	return arb.buildNotifyRequest(body, "application/x-www-form-urlencoded")
}

func (arb *apiRequestBuilder) buildMultipartNotifyRequest(body io.Reader, w *multipart.Writer) (*http.Request, error) {
	return arb.buildNotifyRequest(body, w.FormDataContentType())
}

func (arb *apiRequestBuilder) buildNotifyRequest(body io.Reader, contentType string) (*http.Request, error) {
	req, err := http.NewRequest(
		"POST",
		baseURL+"/notify",
		body,
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+arb.token)
	req.Header.Set("Content-Type", contentType)

	return req, nil
}
