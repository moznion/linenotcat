package linenotcat

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
	values.Set("message", string(msg))

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
