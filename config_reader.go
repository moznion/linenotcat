package linenotcat

import (
	"bufio"
	"fmt"
	"os"
)

func readToken(configFilePath string) (string, error) {
	fp, err := os.Open(configFilePath)
	if err != nil {
		return "", err
	}

	var token string

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		token = scanner.Text()
		if token != "" {
			break
		}
	}
	err = scanner.Err()
	if err != nil {
		return "", err
	}

	return token, nil
}

func readDefaultToken() (string, error) {
	defaultConfigFilePath, err := getDefaultConfigFilePath()
	if err != nil {
		return "", err
	}
	return readToken(defaultConfigFilePath)
}

func getDefaultConfigFilePath() (string, error) {
	homedir := os.Getenv("HOME")
	if homedir == "" {
		return "", fmt.Errorf("$HOME not set")
	}
	return homedir + "/.linenotcat", nil
}
