package linenotcat

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type lineNotifier struct {
	apiRequestBuilder *apiRequestBuilder
}

func (l *lineNotifier) notifyMessage(msg string, tee bool) error {
	values := url.Values{}
	values.Set("message", msg)

	req, err := l.apiRequestBuilder.buildFormNotifyRequest(strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}

	if tee {
		fmt.Print(msg)
	}

	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err != nil {
		return err
	}

	return checkHTTPStatus(res)
}

func (l *lineNotifier) notifyMessages(msgs []string, tee bool) error {
	return l.notifyMessage(strings.Join(msgs, "\n"), tee)
}

func (l *lineNotifier) notifyFile(tmpFilePath string, tee bool) error {
	msg, err := ioutil.ReadFile(tmpFilePath)
	if err != nil {
		return err
	}

	return l.notifyMessage(string(msg), tee)
}

func (l *lineNotifier) notifyImage(imageFilePath, message string, tee bool) error {
	body := &bytes.Buffer{}
	mw := multipart.NewWriter(body)

	mw.WriteField("message", message)
	if tee {
		fmt.Print(message)
	}

	pw, err := mw.CreateFormFile("imageFile", imageFilePath)
	if err != nil {
		return err
	}

	file, err := os.Open(imageFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(pw, file)
	if err != nil {
		return err
	}

	err = mw.Close()
	if err != nil {
		return err
	}

	req, err := l.apiRequestBuilder.buildMultipartNotifyRequest(body, mw)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err != nil {
		return err
	}

	return checkHTTPStatus(res)
}
