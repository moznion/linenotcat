package linenotcat

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func checkHttpStatus(res *http.Response) error {
	statusCode := res.StatusCode
	if statusCode < 200 || statusCode > 299 {
		read, _ := ioutil.ReadAll(res.Body)
		return fmt.Errorf(string(read))
	}

	return nil
}
