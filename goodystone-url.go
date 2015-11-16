package goodystone

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type URLPacket struct {
	Mac    string
	Interf string
	Url    string
	Scheme string
	Type   string
	Power  int
	Rssi   int
}

func NewURLPacket(line string) *URLPacket {
	vals := strings.Split(line, " ")
	return &URLPacket{
		Mac:    parseMac(vals),
		Url:    parseUrl(vals),
		Scheme: parseScheme(vals),
		Type:   parseType(vals),
		Power:  parsePower(vals),
		Rssi:   parseRssi(vals)}
}

func parseUrl(vals []string) string {
	//	return strings.Join(vals[28:len(vals)-1], "")
	url, _ := hex.DecodeString(strings.Join(vals[28:len(vals)-1], ""))
	return string(url)
}

func parseScheme(vals []string) string {
	switch vals[27] {
	case "00":
		return "http://www."
	case "01":
		return "https://www."
	case "02":
		return "http://"
	case "03":
		return "https://"
	}
	return ""
}

func (edp *URLPacket) MapKey() string {
	return MapKey(edp.Mac)
}

func (edp *URLPacket) ToString() string {
	return fmt.Sprintf("INT %v URL %v RSSI %d", edp.Interf, edp.Url, edp.Rssi)
}
