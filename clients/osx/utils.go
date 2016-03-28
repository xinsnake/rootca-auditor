package main

import (
	"bytes"
	"encoding/binary"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func AsciiToText(input string) string {

	inputBytes := []byte(input)

	re := regexp.MustCompile(`\\[0-7]{3}`)

	inputBytes = re.ReplaceAllFunc(inputBytes, func(m []byte) []byte {
		buf := new(bytes.Buffer)
		byteInt, _ := strconv.ParseInt(string(m[1:]), 8, 64)
		_ = binary.Write(buf, binary.LittleEndian, int8(byteInt))
		return buf.Bytes()
	})

	return string(inputBytes)
}

func StripSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
