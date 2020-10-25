package requestcorrelation

import (
	"fmt"
	"strconv"
	"strings"
)

// TraceParent defines the model for the W3C Trace Parent
type TraceParent struct {
	Version string
	TraceID string
	SpanID  string
	Flags   string
}

// GenerateNewTraceParent returns a new trace parent which defines a new execution context
func GenerateNewTraceParent() TraceParent {
	tp := TraceParent{
		Version: "00",
		TraceID: GenerateRandomHexString(TraceIDLength),
		SpanID:  GenerateRandomHexString(SpanIDLength),
		Flags:   "00",
	}

	return tp
}

// NewSpanFromParent returns a new TraceParent which shares the same parent trace context and a new span which
// defines the new "child" context
func NewSpanFromParent(tp TraceParent) TraceParent {
	newTp := TraceParent{
		Version: tp.Version,
		TraceID: tp.TraceID,
		SpanID:  GenerateRandomHexString(SpanIDLength),
		Flags:   tp.Flags,
	}

	return newTp
}

// NewTraceParent parses and valdiates the input parameters and if valid returns a TraceParent record.
// If the parameters are invalid, it returns an error string
func NewTraceParent(version int, traceID string, spanID string, flags int) (*TraceParent, error) {

	err := validateTraceData(version, traceID, spanID, flags)
	if err != nil {
		return nil, err
	}

	return &TraceParent{
		Version: fmt.Sprintf("%02x", version),
		TraceID: traceID,
		SpanID:  spanID,
		Flags:   fmt.Sprintf("%02x", flags),
	}, nil
}

// NewTraceParentWithStr uses strings for version and flags for a simpler api definition
func NewTraceParentWithStr(version string, traceID string, spanID string, flags string) (*TraceParent, error) {

	nVersion, _ := strconv.ParseUint(version, 16, 32)
	nFlags, _ := strconv.ParseUint(flags, 16, 32)

	return NewTraceParent(
		int(nVersion),
		traceID,
		spanID,
		int(nFlags),
	)
}

func validateTraceData(version int, traceID string, spanID string, flags int) error {
	// https://www.w3.org/TR/trace-context/#version
	// TraceParent assumes version 00 by default. Version ff is forbidden
	if version < 0 || version > 254 {
		return &errorString{"Invalid version number. Must be [0, 255)"}
	}

	if err := IsValidTraceIDString(traceID); err != nil {
		return err
	}

	if err := IsValidSpanIDString(spanID); err != nil {
		return err
	}

	// https://www.w3.org/TR/trace-context/#trace-flags
	if version < 0 || version > 255 {
		return &errorString{"Invalid version number. Must be [0, 255]"}
	}

	return nil
}

// ToString returns a stringified version of the TraceParent
func (tp TraceParent) ToString() string {
	return fmt.Sprintf("%v-%v-%v-%v", tp.Version, tp.TraceID, tp.SpanID, tp.Flags)
}

// ToRequestID converts the W3C Trace Parent to a legacy Request-Id
func (tp TraceParent) ToRequestID() (*RequestID, error) {
	return NewRequestID(
		tp.TraceID,
		tp.SpanID,
	)
}

// TraceParentFromString parses a W3C trace parent string and returns a strongly types model
func TraceParentFromString(s string) (*TraceParent, error) {

	if s == "" {
		return nil, &errorString{"Invalid trace string"}
	}

	ss := strings.Split(s, "-")
	if len(ss) != 4 {
		return nil, &errorString{"Invalid trace string"}
	}

	version := ss[0]
	traceID := ss[1]
	spanID := ss[2]
	flags := ss[3]

	if !IsValidHexString(version, 2) || !IsValidHexString(flags, 2) {
		return nil, &errorString{"Invalid trace string"}
	}

	return NewTraceParentWithStr(version, traceID, spanID, flags)
}
