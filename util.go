package main

import (
	"os"
	"strings"
)

const (
	BASE_URL     = "https://api-frontend.kemdikbud.go.id/"
	abc          = "ABCDEFGHIJKLMNOPQRSTUVWZ"
	MODE_BASIC   = "basic"
	MODE_ADVANCE = "advance"
)

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
