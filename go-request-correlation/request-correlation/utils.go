package requestcorrelation

import (
	"bytes"
	"crypto/rand"
	"log"
	"math/big"
	"strings"
)

// TraceIDLength is the length of a W3C Trace ID String
const TraceIDLength = 32

// SpanIDLength is the length of a W3C Span ID String
const SpanIDLength = 16

// HexAlphabet defines the allowed alphabet for a W3C Trace hex string
const HexAlphabet = "0123456789abcdef"

// IsValidHexString checks if the input string is a valid hexadecimal string
// as per the W3C specification for a Trace Hex string
func IsValidHexString(s string, length int) bool {
	if s == "" || len(s) != length {
		return false
	}

	for _, r := range s {
		if !strings.ContainsRune(HexAlphabet, r) {
			return false
		}
	}

	// all checks pass
	return true
}

// GenerateRandomHexString generates a random hexadecimal string using cryptographic.
// Values are compliant with the W3C specification for a trace parent hex string
func GenerateRandomHexString(length int) string {
	const max = int64(len(HexAlphabet))

	var b bytes.Buffer

	for i := 0; i < length; i++ {
		randInt := cryptoRandSecure(max)
		s := HexAlphabet[randInt : randInt+1]
		b.WriteString(s)
	}

	return b.String()

}

// IsValidTraceIDString checks if the input traceID is a valid W3C Trace ID hex string
func IsValidTraceIDString(traceID string) error {

	// https://www.w3.org/TR/trace-context/#trace-id
	// All bytes as zero (00000000000000000000000000000000) is considered an invalid value.
	if !IsValidHexString(traceID, TraceIDLength) || traceID == "00000000000000000000000000000000" {
		return &errorString{"Invalid TraceID String"}
	}

	return nil
}

// IsValidSpanIDString checks if the input spanID is a valid W3C Span ID hex string
func IsValidSpanIDString(spanID string) error {
	// https://www.w3.org/TR/trace-context/#parent-id
	// All bytes as zero (0000000000000000) is considered an invalid value.
	if !IsValidHexString(spanID, SpanIDLength) || spanID == "0000000000000000" {
		return &errorString{"Invalid SpanID String"}
	}

	return nil
}

func cryptoRandSecure(max int64) int64 {
	randInt, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		log.Println("Error generating crypto rand. Error", err)
	}

	return (*randInt).Int64()
}
