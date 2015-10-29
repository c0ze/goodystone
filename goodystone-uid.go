package goodystone

import (
	"fmt"
	"strings"
)

type UIDPacket struct {
	Mac      string
	Interf   string
	Uid      string
	Instance string
	Type     string
	Power    int
	Rssi     int
}

func NewUIDPacket(line string) *UIDPacket {
	vals := strings.Split(line, " ")
	return &UIDPacket{
		Mac:      parseMac(vals),
		Uid:      parseUid(vals),
		Instance: parseInstance(vals),
		Type:     parseType(vals),
		Power:    parsePower(vals),
		Rssi:     parseRssi(vals)}
}

func parseUid(vals []string) string {
	return strings.Join(vals[27:37], "")
}

func parseInstance(vals []string) string {
	return strings.Join(vals[37:43], "")
}

func (edp *UIDPacket) MapKey() string {
	return fmt.Sprintf("%v|%v", edp.Uid, edp.Instance)
}

func (edp *UIDPacket) ToString() string {
	return fmt.Sprintf("INT %v UID %v INSTANCE %v RSSI %d", edp.Interf, edp.Uid, edp.Instance, edp.Rssi)
}
