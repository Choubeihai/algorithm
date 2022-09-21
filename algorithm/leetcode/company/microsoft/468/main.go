package main

import (
	"strconv"
	"strings"
)

func validIPAddress(queryIP string) string {
	list := strings.Split(queryIP, ".")
	if len(list) == 4 {
		for i := 0; i < 4; i++ {
			if len(list[i]) > 1 && list[i][0] == '0' {
				return "Neither"
			}
			v, err := strconv.Atoi(list[i])
			if err != nil {
				return "Neither"
			}
			if v > 255 {
				return "Neither"
			}
		}
		return "IPv4"
	}
	list = strings.Split(queryIP, ":")
	if len(list) == 8 {
		for i := 0; i < 8; i++ {
			if len(list[i]) > 4 {
				return "Neither"
			}
			v, err := strconv.ParseInt(list[i], 16, 0)
			if err != nil {
				return "Neither"
			}
			if v > 65535 {
				return "Neither"
			}
		}
		return "IPv6"
	}
	return "Neither"
}
