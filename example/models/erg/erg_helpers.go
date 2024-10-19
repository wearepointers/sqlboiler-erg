// Generated by sqlboiler-erg: DO NOT EDIT.
package erg

import (
	"strings"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
)

type json any
type decimal float64

func contains(a []string, s string) bool {
	for _, v := range a {
		if strings.EqualFold(v, s) {
			return true
		}
	}
	return false
}

func doesNotContain(a []string, s string) bool {
	return !contains(a, s)
}

/////////////////////////////////////////////////////////////

func nullDotUintToUintPtr(n null.Uint) *uint {
	return n.Ptr()
}

func nullDotBytesToByteSlicePtr(n null.Bytes) *[]byte {
	return n.Ptr()
}

func nullDotByteToBytePtr(n null.Byte) *byte {
	return n.Ptr()
}

func nullDotBoolToBoolPtr(n null.Bool) *bool {
	return n.Ptr()
}

func nullDotStringToStringPtr(n null.String) *string {
	return n.Ptr()
}

func nullDotJSONToJSONPtr(n null.JSON) *json {
	if n.IsZero() {
		return nil
	}

	j := json(n)
	return &j
}

func typesDotJSONToJSON(t types.JSON) json {
	return json(t)
}

func typesDotStringArrayToStringSlice(tsa types.StringArray) []string {
	return tsa
}

func typesDotBoolArrayToBoolSlice(tba types.BoolArray) []bool {
	return tba
}

func typesDotBytesArrayToBytesSlice(tba types.BytesArray) [][]byte {
	return tba
}

func typesDotByteToByte(t types.Byte) byte {
	return byte(t)
}

func typesDotInt64ArrayToInt64Slice(tsa types.Int64Array) []int64 {
	return tsa
}

func typesDotFloat64ArrayToFloat64Slice(tsa types.Float64Array) []float64 {
	return tsa
}

func nullDotTimeToTimeDotTimePtr(n null.Time) *time.Time {
	return n.Ptr()
}

func nullDotTimeToTimePtr(n null.Time) *time.Time {
	return n.Ptr()
}

func typesDotDecimalToDecimal(t types.Decimal) decimal {
	if t.Big == nil {
		return 0
	}
	f, _ := t.Float64()
	return decimal(f)
}

func nullDotIntToIntPtr(n null.Int) *int {
	return n.Ptr()
}

func nullDotBytesToBytesPtr(n null.Bytes) *[]byte {
	return n.Ptr()
}

func typesDotDecimalArrayToDecimalSlice(t types.DecimalArray) []decimal {
	m := make([]decimal, len(t))
	for i, td := range t {
		m[i] = typesDotDecimalToDecimal(td)
	}

	return m
}

func typesDotNullDecimalToNullDecimal(t types.NullDecimal) *decimal {
	if t.Big == nil {
		return nil
	}
	f, _ := t.Float64()
	d := decimal(f)
	return &d
}