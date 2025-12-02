package input

import (
	"bufio"
	"bytes"
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

func ReadSeperator(name string, seperator byte, cb func(string)) error {
	f, err := files.Open("files/" + name + ".txt")
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReaderSize(f, 1<<20) // 1MB

	for {
		part, err := r.ReadSlice(seperator)

		if err == bufio.ErrBufferFull {
			acc := append([]byte{}, part...)
			for {
				chunk, err2 := r.ReadSlice(seperator)
				acc = append(acc, chunk...)
				if err2 != bufio.ErrBufferFull {
					err = err2
					part = acc
					break
				}
			}
		}

		token := bytes.TrimRight(part, ",\r\n \t")

		if err == io.EOF {
			if len(token) > 0 {
				cb(string(token))
			}
			return nil
		}

		if err != nil && err != bufio.ErrBufferFull {
			return err
		}

		if len(token) > 0 {
			cb(string(token))
		}
	}
}
