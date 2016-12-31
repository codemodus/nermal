package nermal

import (
	"fmt"
	"os"
	"path/filepath"
	"unicode"
)

// Port normalizes and validates a string port number with preceding colon.
func Port(p string) (string, error) {
	if p[0] != ':' {
		p = ":" + p
	}

	for _, v := range p[1:] {
		if !unicode.IsDigit(v) {
			return "", fmt.Errorf("val %q not digit in %q", v, p)
		}
	}

	return p, nil
}

// File normalizes and validates a string filename by checking the writability
// of the containing directory.
func File(f string) error {
	d := filepath.Dir(f)
	if _, err := os.Stat(d); err != nil {
		return err
	}

	fp := filepath.Join(d, "touch.zilch")
	if err := TouchFile(fp); err != nil {
		return err
	}

	return os.Remove(fp)
}

// TouchFile ...
func TouchFile(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("cannot touch file in directory %q", filepath.Dir(name))
	}

	return f.Close()
}
