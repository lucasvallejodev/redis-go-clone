package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func getBulkSize(reader *bufio.Reader) (int, error) {
	bulkSize, err := getReaderString(reader)
	if err != nil {
		return -1, err
	}
	size, err := strconv.Atoi(bulkSize)
	if err != nil {
		return -1, err
	}
	return size, nil
}

func getArgs(reader *bufio.Reader, bulkSize int) ([]string, error) {
	var args []string
	for i := 0; i < bulkSize; i++ {
		size, err := getReaderString(reader)
		if err != nil {
			return []string{}, err
		}
		argSize, err := strconv.Atoi(size)
		if err != nil {
			return []string{}, fmt.Errorf("error: strconv for argsize %s", err)
		}
		arg, err := getReaderString(reader)
		if err != nil {
			return []string{}, err
		}
		/* Debug */
		if len(arg) != argSize {
			return []string{}, fmt.Errorf("error: size is different when parsing arguments")
		}
		args = append(args, arg)
	}
	return args, nil
}

func getReaderString(reader *bufio.Reader) (string, error) {
	chunk, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			return "", nil
		}
		return "", err
	}
	trimmedChunk := strings.ToLower(strings.TrimFunc(chunk, func(r rune) bool {
		return r == '$' || r == '\r' || r == '\n' || r == '*'
	}))
	return trimmedChunk, nil
}
