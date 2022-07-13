package data

import (
	"encoding/json"
	"errors"
)

type UUID struct {
	Valid bool
	Bytes [16]byte
}

func (id UUID) MarshalJSON() ([]byte, error) {
	if id.Valid {
		return json.Marshal(id.String())
	} else {
		return json.Marshal(nil)
	}
}

func (id *UUID) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	*id = ParseUUID(s)

	if !id.Valid {
		return errors.New("Invalid UUID")
	}

	return nil
}

func (id *UUID) String() string {
	src := id.Bytes
	out := [36]byte{}

	encodeHexBytes(out[0:8], src[0:4])
	out[8] = '-'
	encodeHexBytes(out[9:13], src[4:6])
	out[13] = '-'
	encodeHexBytes(out[14:18], src[6:8])
	out[18] = '-'
	encodeHexBytes(out[19:23], src[8:10])
	out[23] = '-'
	encodeHexBytes(out[24:36], src[10:16])

	return string(out[:])
}

/// Encodes `len(src)` hex bytes into `dst` from `src`. `src` is treated as bytes.
func encodeHexBytes(dst, src []byte) {
	// ensures that len(dst) has enough space
	if len(dst) < len(src)*2 {
		panic("Expected len(dst) < len(src)*2")
	}

	for i := 0; i < len(src); i++ {
		hi := encodeHexOctet(src[i] >> 4)
		lo := encodeHexOctet(src[i] & 0xF)
		if hi == 255 || lo == 255 {
			panic("Unexpected failure of encodeHexOctet")
		}

		dst[i*2] = hi
		dst[i*2+1] = lo
	}
}

/// Encodes a single hex octet. Returns 255 on error.
func encodeHexOctet(src byte) byte {
	if src < 10 {
		return src + '0'
	} else if src < 16 {
		return src - 10 + 'a'
	} else {
		return 255
	}
}

// Reads a UUID from a string. Returns a null UUID on failure.
func ParseUUID(src string) UUID {
	srcBytes := []byte(src)

	out := [16]byte{}

	if !decodeHexBytes(out[0:4], srcBytes[0:8]) ||
		src[8] != '-' ||
		!decodeHexBytes(out[4:6], srcBytes[9:13]) ||
		src[13] != '-' ||
		!decodeHexBytes(out[6:8], srcBytes[14:18]) ||
		src[18] != '-' ||
		!decodeHexBytes(out[8:10], srcBytes[19:23]) ||
		src[23] != '-' ||
		!decodeHexBytes(out[10:16], srcBytes[24:36]) {
		return UUID{}
	}

	return UUID{Valid: true, Bytes: out}
}

// Decodes `len(dst)` hex bytes into `dst` from `src`. `src` is treated as a string.
//
// Returns false on error.
func decodeHexBytes(dst, src []byte) bool {
	// ensures that len(src) has enough characters
	if len(src) < len(dst)*2 {
		return false
	}

	for i := 0; i < len(dst); i++ {
		hi := decodeHexOctet(src[i*2])
		lo := decodeHexOctet(src[i*2+1])
		if hi == 255 || lo == 255 {
			return false
		}

		dst[i] = hi<<4 | lo
	}

	return true
}

// Decodes a single hex octet. Returns 255 on error.
func decodeHexOctet(src byte) byte {
	if src >= '0' && src <= '9' {
		return src - '0'
	} else if src >= 'a' && src <= 'f' {
		return src - 'a' + 10
	} else {
		return 255
	}
}
