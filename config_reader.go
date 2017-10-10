package linenotcat

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"runtime"
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
	var homedir string
	if runtime.GOOS == "windows" {
		if homedir = os.Getenv("USERPROFILE"); homedir == "" {
			return "", errors.New(`%USERPROFILE% not set`)
		}
	} else {
		if homedir = os.Getenv("HOME"); homedir == "" {
			return "", errors.New(`$HOME not set`)
		}
	}
	return filepath.Join(homedir, ".linenotcat"), nil
}
