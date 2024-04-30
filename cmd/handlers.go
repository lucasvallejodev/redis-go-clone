package main

import (
	"bufio"
	"bytes"
)

type Command struct {
	Args     []string
	BulkSize int
}

func NewCommand(data []byte) (Command, error) {
	reader := bufio.NewReader(bytes.NewReader(data))
	bulkSize, err := getBulkSize(reader)
	if err != nil {
		return Command{}, err
	}
	args, err := getArgs(reader, bulkSize)
	if err != nil {
		return Command{}, err
	}
	return Command{
		Args:     args,
		BulkSize: bulkSize,
	}, nil
}
