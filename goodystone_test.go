package goodystone

import (
	"strings"
	"testing"
)

// TYPE : 00 UUID : ED D1 EB EA C0 4E 5D EF A0 17  INSTANCE : EB B1 CB 1C 7C C1  POWER : -21 RSSI : -48
const UidPacket1 = `04 3E 29 02 01 00 01 C1 7C 1C CB B1 EB 1D 02 01 06 03 03 AA FE 15 16 AA FE 00 EB ED D1 EB EA C0 4E 5D EF A0 17 EB B1 CB 1C 7C C1 D0`

// TYPE : 00 UUID : ED D1 EB EA C0 4E 5D EF A0 17  INSTANCE : EB B1 CB 1C 7C C1  POWER : -21 RSSI : -49
const UidPacket2 = `04 3E 29 02 01 00 01 C1 7C 1C CB B1 EB 1D 02 01 06 03 03 AA FE 15 16 AA FE 00 EB ED D1 EB EA C0 4E 5D EF A0 17 EB B1 CB 1C 7C C1 CF`

const TlmPacket = `04 3E 25 02 01 03 01 C1 7C 1C CB B1 EB 19 02 01 06 03 03 AA FE 11 16 AA FE 20 00 0B 6D 20 C0 00 06 1B 17 00 43 7C B8 D5`

// TYPE : 10 SCHEME : 02 URL : go.esti.be POWER : -27 RSSI : -54
const UrlPacket1 = `04 3E 24 02 01 03 01 C1 7C 1C CB B1 EB 18 02 01 06 03 03 AA FE 10 16 AA FE 10 E5 02 67 6F 2E 65 73 74 69 2E 62 65 CA`

// TYPE : 10 SCHEME : 02 URL : go.esti.ben POWER : -27 RSSI : -59
const UrlPacket2 = `04 3E 25 02 01 00 01 C1 7C 1C CB B1 EB 19 02 01 06 03 03 AA FE 11 16 AA FE 10 E5 02 67 6F 2E 65 73 74 69 2E 62 65 6E C5`

// TYPE : 10 SCHEME : 02 URL : go.esti.ben POWER : -21 RSSI : -50
const UrlPacket3 = `04 3E 25 02 01 03 01 C1 7C 1C CB B1 EB 19 02 01 06 03 03 AA FE 11 16 AA FE 10 EB 02 67 6F 2E 65 73 74 69 2E 62 65 6E CE`

// TYPE : 10 SCHEME : 02 URL : go.esti.be POWER : -21 RSSI : -45
const UrlPacket4 = `04 3E 24 02 01 03 01 C1 7C 1C CB B1 EB 18 02 01 06 03 03 AA FE 10 16 AA FE 10 EB 02 67 6F 2E 65 73 74 69 2E 62 65 D3`

func TestIsValid(t *testing.T) {
	if !IsValid(UidPacket1) {
		t.Errorf("Validation failed for Packet1")
	}

	if !IsValid(UidPacket2) {
		t.Errorf("Validation failed for Packet2")
	}

	if !IsValid(TlmPacket) {
		t.Errorf("Validation failed for TlmPacket")
	}

	if !IsValid(UrlPacket1) {
		t.Errorf("Validation failed for UrlPacket1")
	}

	if !IsValid(UrlPacket2) {
		t.Errorf("Validation failed for UrlPacket2")
	}

	if !IsValid(UrlPacket3) {
		t.Errorf("Validation failed for UrlPacket3")
	}

	if !IsValid(UrlPacket4) {
		t.Errorf("Validation failed for UrlPacket4")
	}
}

func TestParseUid(t *testing.T) {
	uid := parseUid(strings.Split(UidPacket1, " "))
	expectedUid := "EDD1EBEAC04E5DEFA017"
	if uid != expectedUid {
		t.Errorf("Parsing uid {%v} failed for packet1: %v", expectedUid, uid)
	}

	uid = parseUid(strings.Split(UidPacket2, " "))
	if uid != expectedUid {
		t.Errorf("Parsing uid {%v} failed for packet2: %v", expectedUid, uid)
	}
}

func TestParseInstance(t *testing.T) {
	instance := parseInstance(strings.Split(UidPacket1, " "))
	expectedInstance := "EBB1CB1C7CC1"
	if instance != expectedInstance {
		t.Errorf("Parsing instance {%v} failed for packet1: %v", expectedInstance, instance)
	}

	instance = parseInstance(strings.Split(UidPacket2, " "))
	if instance != expectedInstance {
		t.Errorf("Parsing instance {%v} failed for packet2: %v", expectedInstance, instance)
	}
}

func TestParseType(t *testing.T) {
	Type := parseType(strings.Split(UidPacket1, " "))
	expectedType := "00"
	if Type != expectedType {
		t.Errorf("Parsing type {%v} failed for packet1: %v", expectedType, Type)
	}

	Type = parseType(strings.Split(UidPacket2, " "))
	if Type != expectedType {
		t.Errorf("Parsing type {%v} failed for packet2: %v", expectedType, Type)
	}

	Type = parseType(strings.Split(TlmPacket, " "))
	expectedType = "20"
	if Type != expectedType {
		t.Errorf("Parsing type {%v} failed for tlm packet: %v", expectedType, Type)
	}

	Type = parseType(strings.Split(UrlPacket1, " "))
	expectedType = "10"
	if Type != expectedType {
		t.Errorf("Parsing type {%v} failed for url packet: %v", expectedType, Type)
	}
}

func TestParsePower(t *testing.T) {
	power := parsePower(strings.Split(UidPacket1, " "))
	expectedPower := -21
	if power != expectedPower {
		t.Errorf("Parsing power {%v} failed for packet1: %v", expectedPower, power)
	}

	power = parsePower(strings.Split(UidPacket2, " "))
	if power != expectedPower {
		t.Errorf("Parsing power {%v} failed for packet2: %v", expectedPower, power)
	}

	power = parsePower(strings.Split(UrlPacket2, " "))
	expectedPower = -27
	if power != expectedPower {
		t.Errorf("Parsing power {%v} failed for url packet2: %v", expectedPower, power)
	}

	power = parsePower(strings.Split(UrlPacket3, " "))
	expectedPower = -21
	if power != expectedPower {
		t.Errorf("Parsing power {%v} failed for url packet3: %v", expectedPower, power)
	}
}

func TestParseRssi(t *testing.T) {
	rssi := parseRssi(strings.Split(UidPacket1, " "))
	expectedRssi := -48
	if rssi != expectedRssi {
		t.Errorf("Parsing rssi {%v} failed for packet1: %v", expectedRssi, rssi)
	}

	rssi = parseRssi(strings.Split(UidPacket2, " "))
	expectedRssi = -49
	if rssi != expectedRssi {
		t.Errorf("Parsing rssi {%v} failed for packet2: %v", expectedRssi, rssi)
	}

	rssi = parseRssi(strings.Split(TlmPacket, " "))
	expectedRssi = -43
	if rssi != expectedRssi {
		t.Errorf("Parsing rssi {%v} failed for tlm packet: %v", expectedRssi, rssi)
	}

	rssi = parseRssi(strings.Split(UrlPacket1, " "))
	expectedRssi = -54
	if rssi != expectedRssi {
		t.Errorf("Parsing rssi {%v} failed for url packet1: %v", expectedRssi, rssi)
	}

	rssi = parseRssi(strings.Split(UrlPacket2, " "))
	expectedRssi = -59
	if rssi != expectedRssi {
		t.Errorf("Parsing rssi {%v} failed for tlm url packet2: %v", expectedRssi, rssi)
	}

}

func TestParseMac(t *testing.T) {
	mac := parseMac(strings.Split(UidPacket1, " "))
	expectedMac := "EBB1CB1C7CC1"
	if mac != expectedMac {
		t.Errorf("Parsing mac {%v} failed for packet1: %v", expectedMac, mac)
	}

	mac = parseMac(strings.Split(UidPacket2, " "))
	if mac != expectedMac {
		t.Errorf("Parsing mac {%v} failed for packet2: %v", expectedMac, mac)
	}

	mac = parseMac(strings.Split(TlmPacket, " "))
	if mac != expectedMac {
		t.Errorf("Parsing mac {%v} failed for tlm packet: %v", expectedMac, mac)
	}

	mac = parseMac(strings.Split(UrlPacket1, " "))
	if mac != expectedMac {
		t.Errorf("Parsing mac {%v} failed for url packet: %v", expectedMac, mac)
	}
}

func TestToString(t *testing.T) {
	uidP := NewUIDPacket(UidPacket1)
	expectedStr := "INT  UID EDD1EBEAC04E5DEFA017 INSTANCE EBB1CB1C7CC1 RSSI -48"
	if uidP.ToString() != expectedStr {
		t.Errorf("Parsing packet {%v} failed for packet1", uidP.ToString())
	}

	uidP = NewUIDPacket(UidPacket2)
	expectedStr = "INT  UID EDD1EBEAC04E5DEFA017 INSTANCE EBB1CB1C7CC1 RSSI -49"
	if uidP.ToString() != expectedStr {
		t.Errorf("Parsing packet {%v} failed for packet2", uidP.ToString())
	}
}

func TestMapKey(t *testing.T) {
	uidP := NewUIDPacket(UidPacket1)
	expectedStr := "EDD1EBEAC04E5DEFA017|EBB1CB1C7CC1"
	if uidP.MapKey() != expectedStr {
		t.Errorf("Parsing packet {%v} failed for packet1", uidP.MapKey())
	}

	uidP = NewUIDPacket(UidPacket2)
	if uidP.MapKey() != expectedStr {
		t.Errorf("Parsing packet {%v} failed for packet2", uidP.MapKey())
	}

	urlP := NewURLPacket(UrlPacket2)
	expectedStr = "EBB1CB1C7CC1"
	if urlP.MapKey() != expectedStr {
		t.Errorf("Parsing packet {%v} failed for url packet2", urlP.MapKey())
	}
}

func TestParseVersion(t *testing.T) {
	version := parseVersion(strings.Split(TlmPacket, " "))
	expectedVersion := "00"
	if version != expectedVersion {
		t.Errorf("Parsing version {%v} failed for packet: %v", expectedVersion, version)
	}
}

func TestParseBattery(t *testing.T) {
	battery := parseBattery(strings.Split(TlmPacket, " "))
	expectedBattery := 2925
	if battery != expectedBattery {
		t.Errorf("Parsing battery {%v} failed for packet: %v", expectedBattery, battery)
	}
}

func TestParseTemperature(t *testing.T) {
	temperature := parseTemperature(strings.Split(TlmPacket, " "))
	expectedTemperature := 32.75
	if temperature != expectedTemperature {
		t.Errorf("Parsing temperature {%v} failed for packet: %v", expectedTemperature, temperature)
	}
}

func TestParsePacketCount(t *testing.T) {
	packetCount := parsePacketCount(strings.Split(TlmPacket, " "))
	expectedPacketCount := 400151
	if packetCount != expectedPacketCount {
		t.Errorf("Parsing packetCount {%v} failed for packet: %v", expectedPacketCount, packetCount)
	}
}

func TestParseTimeCount(t *testing.T) {
	timeCount := parseTimeCount(strings.Split(TlmPacket, " "))
	expectedTimeCount := 4422840
	if timeCount != expectedTimeCount {
		t.Errorf("Parsing timeCount {%v} failed for packet: %v", expectedTimeCount, timeCount)
	}
}

func TestParseScheme(t *testing.T) {
	scheme := parseScheme(strings.Split(UrlPacket1, " "))
	expectedScheme := "http://"
	if scheme != expectedScheme {
		t.Errorf("Parsing scheme {%v} failed for packet: %v", expectedScheme, scheme)
	}
}

func TestParseUrl(t *testing.T) {
	url := parseUrl(strings.Split(UrlPacket1, " "))
	expectedUrl := "go.esti.be"
	if url != expectedUrl {
		t.Errorf("Parsing url {%v} failed for url packet1: %v", expectedUrl, url)
	}

	url = parseUrl(strings.Split(UrlPacket2, " "))
	expectedUrl = "go.esti.ben"
	if url != expectedUrl {
		t.Errorf("Parsing url {%v} failed for url packet2: %v", expectedUrl, url)
	}

	url = parseUrl(strings.Split(UrlPacket3, " "))
	expectedUrl = "go.esti.ben"
	if url != expectedUrl {
		t.Errorf("Parsing url {%v} failed for url packet3: %v", expectedUrl, url)
	}

	url = parseUrl(strings.Split(UrlPacket4, " "))
	expectedUrl = "go.esti.be"
	if url != expectedUrl {
		t.Errorf("Parsing url {%v} failed for url packet4: %v", expectedUrl, url)
	}
}
