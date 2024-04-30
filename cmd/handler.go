package main

import (
	"fmt"
	"strings"
)

type Redis struct {
	data map[string]string
}

func initRedis() *Redis {
	redis := &Redis{}
	redis.data = make(map[string]string)
	return redis
}

func (db *Redis) executeCommand(input []string) ([]byte, error) {
	command := strings.ToLower(input[0])

	switch command {
	case "echo":
		return []byte(fmt.Sprintf("$%d\r\n%s\r\n", len(input[1]), input[1])), nil
	case "ping":
		return []byte("+PONG\r\n"), nil
	case "set":
		return db.set(input[1:])
	case "get":
		return db.get(input[1:])
	default:
		return nil, fmt.Errorf("unknown command %s", command)
	}
}
func (db *Redis) set(input []string) ([]byte, error) {
	if len(input) != 2 {
		return nil, fmt.Errorf("invalid number of arguments in 'set' command: %s", input)
	}

	db.data[input[0]] = input[1]
	fmt.Println(db.data)

	return []byte("+OK\r\n"), nil
}
func (db *Redis) get(input []string) ([]byte, error) {
	fmt.Println(db.data)
	if len(input) != 1 {
		return nil, fmt.Errorf("invalid number of arguments in 'get' command: %s", input)
	}

	if result, ok := db.data[input[0]]; ok {
		return []byte(fmt.Sprintf("$%d\r\n%s\r\n", len(result), result)), nil
	}

	return []byte("$-1\r\n"), nil
}
