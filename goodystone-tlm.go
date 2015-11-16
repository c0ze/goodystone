package goodystone

import (
	"strconv"
	"strings"
)

type TLMPacket struct {
	Interf      string
	Mac         string
	Type        string
	Version     string
	Battery     int
	Temperature float64
	PacketCount int
	TimeCount   int
	Rssi        int
}

func NewTLMPacket(line string) *TLMPacket {
	vals := strings.Split(line, " ")
	return &TLMPacket{
		Mac:         parseMac(vals),
		Type:        parseType(vals),
		Version:     parseVersion(vals),
		Battery:     parseBattery(vals),
		Temperature: parseTemperature(vals),
		PacketCount: parsePacketCount(vals),
		TimeCount:   parseTimeCount(vals),
		Rssi:        parseRssi(vals)}
}

func (tlp *TLMPacket) MapKey() string {
	return MapKey(tlp.Mac)
}

func parseVersion(vals []string) string {
	return vals[26]
}

func parseBattery(vals []string) int {
	battery, _ := strconv.ParseInt(strings.Join(vals[27:29], ""), 16, 0)
	return int(battery)
}

func parseTemperature(vals []string) float64 {
	temperature, _ := strconv.ParseInt(strings.Join(vals[29:31], ""), 16, 0)
	return float64(temperature) / 256
}

func parsePacketCount(vals []string) int {
	count, _ := strconv.ParseInt(strings.Join(vals[31:35], ""), 16, 0)
	return int(count)
}

func parseTimeCount(vals []string) int {
	count, _ := strconv.ParseInt(strings.Join(vals[35:39], ""), 16, 0)
	return int(count)
}
