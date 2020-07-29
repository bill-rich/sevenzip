package sevenzip

import (
	"path/filepath"
	"testing"
)

func TestOpenReader(t *testing.T) {
	tables := map[string]struct {
		file string
	}{
		"no header compression": {
			file: "t0.7z",
		},
		"with header compression": {
			file: "t1.7z",
		},
	}

	for name, table := range tables {
		t.Run(name, func(t *testing.T) {
			if _, err := OpenReader(filepath.Join("testdata", table.file)); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestOpenReaderWithPassword(t *testing.T) {
	tables := map[string]struct {
		file     string
		password string
	}{
		"no header compression": {
			file:     "t2.7z",
			password: "password",
		},
		"with header compression": {
			file:     "t3.7z",
			password: "password",
		},
	}

	for name, table := range tables {
		t.Run(name, func(t *testing.T) {
			if _, err := OpenReaderWithPassword(filepath.Join("testdata", table.file), table.password); err != nil {
				t.Fatal(err)
			}
		})
	}
}
