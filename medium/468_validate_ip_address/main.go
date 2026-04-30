package main

import (
	"strconv"
	"strings"
	"unicode"
)

func validIPv4(parts []string) bool {
	for _, v := range parts {
		if len(v) == 0 || len(v) != 1 && v[0] == '0' {
			return false
		}

		number, err := strconv.Atoi(v)
		if err != nil {
			return false
		}

		if number < 0 || number > 255 {
			return false
		}
	}

	return true
}

func validIPv6(parts []string) bool {
	for _, v := range parts {
		if len(v) == 0 {
			return false
		}

		if len(v) > 4 {
			return false
		}

		for _, rn := range []rune(v) {
			if !unicode.Is(unicode.ASCII_Hex_Digit, rn) {
				return false
			}
		}
	}

	return true
}

func validIPAddress(queryIP string) string {
	splitStringIPv4 := strings.Split(queryIP, ".")
	splitStringIPv6 := strings.Split(queryIP, ":")

	if len(splitStringIPv4) == 4 && validIPv4(splitStringIPv4) {
		return "IPv4"
	}
	if len(splitStringIPv6) == 8 && validIPv6(splitStringIPv6) {
		return "IPv6"
	}

	return "Neither"
}

func main() {

}
