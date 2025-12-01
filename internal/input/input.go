package input

import (
	"bufio"
	"embed"
	"fmt"
	"io"
)

//go:embed files/*.txt
var files embed.FS

func ReadAll(name string) (string, error) {
	data, err := files.ReadFile("files/" + name + ".txt")
	if err != nil {
		return "", fmt.Errorf("input: failed to read %s: %w", name, err)
	}
	return string(data), nil
}

func ReadLines(name string) ([]string, error) {
	f, err := files.Open("files/" + name + ".txt")
	if err != nil {
		return nil, fmt.Errorf("input: failed to open %s: %w", name, err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil && err != io.EOF {
		return nil, fmt.Errorf("input: scan error in %s: %w", name, err)
	}
	return lines, nil
}
