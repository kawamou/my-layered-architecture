package infrastructure

import (
	"bufio"
	"os"
	"strings"
)

func Load() (err error) {
	filePath := ".env"

	f, err := os.Open(filePath)
	defer f.Close()

	lines := make([]string, 0, 100)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, l := range lines {
		pair := strings.Split(l, "=")
		os.Setenv(pair[0], pair[1])
	}
	return
}