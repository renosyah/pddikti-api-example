package main

import (
	"bufio"
	"os"
	"strings"
)

const (
	BASE_URL     = "https://api-frontend.kemdikbud.go.id/"
	abc          = "ABCDEFGHIJKLMNOPQRSTUVWZ"
	MODE_BASIC   = "basic"
	MODE_ADVANCE = "advance"
)

func checkcol(fname string, column []string) (bool, error) {

	file, err := os.Open(fname)
	if err != nil {
		return false, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if scanner.Text() == strings.Join(column, ",") {
			return true, nil
		}
	}

	return false, nil
}

func addcol(fname string, column []string) error {

	f, err := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.WriteString(strings.Join(column, ",") + "\n"); err != nil {
		return err
	}

	return nil
}
