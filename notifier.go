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
	token string
}

const (
	lineNotifyEndpoint = "https://notify-api.line.me/api/notify"
)

func (l *lineNotifier) notifyMessage(msg string, tee bool) error {
	values := url.Values{}
	values.Set("message", msg)

	req, err := http.NewRequest(
		"POST",
		lineNotifyEndpoint,
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+l.token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if tee {
		fmt.Print(msg)
	}

	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err != nil {
		return err
	}

	return nil
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

func (l *lineNotifier) notifyImage(imageFilePath string) error {
	body := &bytes.Buffer{}
	mw := multipart.NewWriter(body)

	mw.WriteField("message", "Image file")

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

	req, err := http.NewRequest(
		"POST",
		lineNotifyEndpoint,
		body,
	)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+l.token)
	req.Header.Set("Content-Type", mw.FormDataContentType())

	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err != nil {
		return err
	}

	return nil
}
