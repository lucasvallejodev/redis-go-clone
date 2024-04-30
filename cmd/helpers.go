package main

import (
	"fmt"
	"strconv"
)

const (
	SimpleStringR = '+'
	SimpleErrorR  = '-'
	BulkStringR   = '$'
	ArrayR        = '*'
	CLRFLength    = 2
)

type Parser struct {
	parsedInput []string
}

func (p *Parser) parseEverythingUntilCRLF(text string) (string, string) {
	result := ""
	for _, r := range text {
		if r == '\r' {
			continue
		} else if r == '\n' {
			break
		}
		result = result + string(r)
	}

	return result, text[len(result)+CLRFLength:]
}

func (p *Parser) parseArrayR(text string) ([]string, string, error) {
	strNumber, rest := p.parseEverythingUntilCRLF(text)

	numberOfItems, err := strconv.Atoi(strNumber)
	if err != nil {
		return nil, "", err
	}

	var result []string
	for i := 0; i < numberOfItems; i++ {
		elemParsed, restNotParsed, err := p.parse(rest)
		if err != nil {
			return nil, "", err
		}
		result = append(result, elemParsed...)
		rest = restNotParsed
	}

	return result, "", nil
}

func (p *Parser) parseBulkStringR(text string) ([]string, string, error) {
	strNumber, rest := p.parseEverythingUntilCRLF(text)
	numberBytes, err := strconv.Atoi(strNumber)
	if err != nil {
		return nil, "", err
	}

	var result []byte
	for i := 0; i < numberBytes; i++ {
		result = append(result, rest[i])
	}

	return []string{string(result)}, rest[numberBytes+CLRFLength:], nil
}

func (p *Parser) parse(input string) ([]string, string, error) {
	var result []string
	var partialResult []string
	var restNotParsed string
	var err error
	if len(input) == 0 {
		return result, "", nil
	}
	switch input[0] {
	case SimpleStringR:
		fmt.Println("SimpleStringR")
	case SimpleErrorR:
		fmt.Println("SimpleErrorR")
	case ArrayR:
		partialResult, restNotParsed, err = p.parseArrayR(input[1:])
	case BulkStringR:
		partialResult, restNotParsed, err = p.parseBulkStringR(input[1:])
	}
	result = append(result, partialResult...)

	return result, restNotParsed, err
}

func (p *Parser) parseInput(input string) error {
	parsed, _, err := p.parse(input)
	p.parsedInput = parsed

	return err
}
