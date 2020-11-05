package main

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	BASE_URL     = "https://api-frontend.kemdikbud.go.id/"
	abc          = "ABCDEFGHIJKLMNOPQRSTUVWZ"
	MODE_BASIC   = "basic"
	MODE_ADVANCE = "advance"
)

func removeDir(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func addcol(fname string, column []string) error {

	f, err := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.WriteString("\n" + strings.Join(column, ",")); err != nil {
		return err
	}

	return nil
}
