package linenotcat

import (
	"bufio"
	"os"
)

func readFromStdin(lines chan string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines <- scanner.Text()
	}
	close(lines)
}
