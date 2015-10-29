package goodystone

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const UID = "00"
const URL = "10"
const TLM = "20"

func IsValid(str string) bool {
	r, err := regexp.Compile(`^04\ 3E\ .{2}\ 02\ 01\ .{26}\ 02\ 01\ .{8}\ AA\ FE`)
	if err != nil {
		fmt.Printf("There is a problem with your regexp.\n")
		return false
	}

	return r.MatchString(str)
}

func GetType(str string) string {
	vals := strings.Split(str, " ")
	return parseType(vals)
}

func parseType(vals []string) string {
	return vals[25]
}

func parseRssi(vals []string) int {
	rssi, _ := strconv.ParseInt(vals[len(vals)-1], 16, 0)
	return int(rssi) - 256
}

func parsePower(vals []string) int {
	power, _ := strconv.ParseInt(vals[26], 16, 0)
	return int(power) - 256
}

func parseMac(vals []string) string {
	return strings.Join(
		[]string{
			vals[12],
			vals[11],
			vals[10],
			vals[9],
			vals[8],
			vals[7]},
		"")
}
