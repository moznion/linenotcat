package linenotcat

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type status struct {
	apiRequestBuilder *apiRequestBuilder
}

func (s *status) getStatus() error {
	req, err := s.apiRequestBuilder.buildStatusRequest()
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	read, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(read))

	return checkHttpStatus(res)
}
