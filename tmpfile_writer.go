package linenotcat

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writeTemp(lines chan string) (string, error) {
	tmp, err := ioutil.TempFile(os.TempDir(), "linenotcat-")
	if err != nil {
		return "", err
	}
	defer tmp.Close()

	w := bufio.NewWriter(tmp)
	for line := range lines {
		fmt.Fprintln(w, line)
	}
	w.Flush()

	return tmp.Name(), nil
}
