package util

import (
	"strconv"
	"strings"
)

func StrToInt(intstr string, prune []string) int64 {
	var intRes int64 = 0
	if prune == nil {
		prune = append(prune, `,`)
	}
	for _, ostr := range prune {
		intstr = strings.ReplaceAll(intstr, ostr, ``)
	}
	if intstr != "" {
		intRes, _ = strconv.ParseInt(intstr, 10, 64)
	}
	return intRes
}
