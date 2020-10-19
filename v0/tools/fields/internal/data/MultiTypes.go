package data

import (
	"fmt"
	"github.com/foxcapades/lib-go-discord/v0/tools/fields/internal/types"
)

const (
	soEqual    = "ShouldEqual"
	soResemble = "ShouldResemble"
	soAlmost   = "ShouldAlmostEqual"
)

func wrapString(i string) string {
	return fmt.Sprintf("string(%s)", i)
}

func wrapFloat32(i string) string {
	return fmt.Sprintf("func() float32 {v, e := strconv.ParseFloat(%s, 32); if e != nil {panic(e)}; return float32(v)}()", wrapString(i))
}

func wrapFloat64(i string) string {
	return fmt.Sprintf("func() float64 {v, e := strconv.ParseFloat(%s, 64); if e != nil {panic(e)}; return v}()", wrapString(i))
}

var Fields = []types.FieldDef{
	{
		Name: "I8",
		Type: "int8",
		Testing: types.Testing{
			JSONTest: types.Test{
				ActualTransform: wrapString,
				InputValue:      "120",
				ExpectValue:     `"120"`,
				Comparison:      soEqual,
			},
			ValueTest: types.Test{
				InputValue:  "120",
				ExpectValue: "120",
				Comparison:  soEqual,
			},
		},
	},
	{
		Name: "I16",
		Type: "int16",
		Testing: types.Testing{
			JSONTest: types.Test{
				ActualTransform: wrapString,
				InputValue:      "32_215",
				ExpectValue:     `"32215"`,
				Comparison:      soEqual,
			},
			ValueTest: types.Test{
				InputValue:  "32_215",
				ExpectValue: "32_215",
				Comparison:  soEqual,
			},
		},
	},
	{
		Name: "I32",
		Type: "int32",
		Testing: types.Testing{
			JSONTest: types.Test{
				ActualTransform: wrapString,
				InputValue:      "3_331_115",
				ExpectValue:     `"3331115"`,
				Comparison:      soEqual,
			},
			ValueTest: types.Test{
				InputValue:  "3_331_115",
				ExpectValue: "3_331_115",
				Comparison:  soEqual,
			},
		},
	},
	{
		Name: "I64",
		Type: "int64",
		Testing: types.Testing{
			JSONTest: types.Test{
				ActualTransform: wrapString,
				InputValue:      "123_456_789_123",
				ExpectValue:     `"123456789123"`,
				Comparison:      soEqual,
			},
			ValueTest: types.Test{
				InputValue:  "123_456_789_123",
				ExpectValue: "123_456_789_123",
				Comparison:  soEqual,
			},
		},
	},
	{
		Name: "U8",
		Type: "uint8",
		Testing: types.Testing{
			JSONTest: types.Test{
				ActualTransform: wrapString,
				InputValue:      "250",
				ExpectValue:     `"250"`,
				Comparison:      soEqual,
			},
			ValueTest: types.Test{
				InputValue:  "250",
				ExpectValue: "250",
				Comparison:  soEqual,
			},
		},
	},
	{
		Name: "U16",
		Type: "uint16",
		Testing: types.Testing{
			JSONTest: types.Test{
				ActualTransform: wrapString,
				InputValue:      "62_000",
				ExpectValue:     `"62000"`,
				Comparison:      soEqual,
			},
			ValueTest: types.Test{
				InputValue:  "62_000",
				ExpectValue: "62_000",
				Comparison:  soEqual,
			},
		},
	},
	{
		Name: "U32",
		Type: "uint32",
		Testing: types.Testing{
			JSONTest: types.Test{
				ActualTransform: wrapString,
				InputValue:      "10_880_000",
				ExpectValue:     `"10880000"`,
				Comparison:      soEqual,
			},
			ValueTest: types.Test{
				InputValue:  "10_880_000",
				ExpectValue: "10_880_000",
				Comparison:  soEqual,
			},
		},
	},
	{
		Name: "U64",
		Type: "uint64",
		Testing: types.Testing{
			JSONTest: types.Test{
				ActualTransform: wrapString,
				InputValue:      "999_999_999_999",
				ExpectValue:     `"999999999999"`,
				Comparison:      soEqual,
			},
			ValueTest: types.Test{
				InputValue:  "999_999_999_999",
				ExpectValue: "999_999_999_999",
				Comparison:  soEqual,
			},
		},
	},
	{
		Name: "F32",
		Type: "float32",
		Testing: types.Testing{
			JSONTest: types.Test{
				ActualTransform: wrapFloat32,
				InputValue:      "1.234512",
				ExpectValue:     "1.234512",
				Comparison:      soAlmost,
				CompareParams:   []string{"0.00001"},
			},
			ValueTest: types.Test{
				InputValue:    "1.234512",
				ExpectValue:   "1.234512",
				Comparison:    soAlmost,
				CompareParams: []string{"0.00001"},
			},
		},
	},
	{
		Name: "F64",
		Type: "float64",
		Testing: types.Testing{
			JSONTest: types.Test{
				ActualTransform: wrapFloat64,
				InputValue:      "12345.1234567891",
				ExpectValue:     "12345.1234567891",
				Comparison:      soAlmost,
			},
			ValueTest: types.Test{
				InputValue:  "12345.123456789",
				ExpectValue: "12345.123456789",
				Comparison:  soAlmost,
			},
		},
	},
	{
		Name: "Bool",
		Type: "bool",
		Testing: types.Testing{
			JSONTest: types.Test{
				ActualTransform: wrapString,
				InputValue:      "true",
				ExpectValue:     `"true"`,
				Comparison:      soEqual,
			},
			ValueTest: types.Test{
				InputValue:  "true",
				ExpectValue: "true",
				Comparison:  soEqual,
			},
		},
	},
	{
		Name: "Str",
		Type: "string",
		Testing: types.Testing{
			JSONTest: types.Test{
				ActualTransform: wrapString,
				InputValue:      `"hello world"`,
				ExpectValue:     `"\"hello world\""`,
				Comparison:      soEqual,
			},
			ValueTest: types.Test{
				InputValue:  `"hello world"`,
				ExpectValue: `"hello world"`,
				Comparison:  soEqual,
			},
		},
	},
	{
		Name: "Snowflake",
		Type: "Snowflake",
		Testing: types.Testing{
			JSONTest: types.Test{
				ActualTransform: wrapString,
				InputValue:      "func() Snowflake {out := Snowflake{}; out.SetRawValue(123_456_789_123); return out}()",
				ExpectValue:     `"123456789123"`,
				Comparison:      soEqual,
			},
			ValueTest: types.Test{
				ActualTransform: func(s string) string {
					return s + ".RawValue()"
				},
				InputValue:  "func() Snowflake {out := Snowflake{}; out.SetRawValue(123_456_789_123); return out}()",
				ExpectValue: "123456789123",
				Comparison:  soEqual,
			},
		},
	},
	{
		Name: "Time",
		Type: "time.Time",
		Testing: types.Testing{
			JSONTest:  types.Test{
				ActualTransform: wrapString,
				InputValue: "time.Unix(0, 0).UTC()",
				ExpectValue: `"\"1970-01-01T00:00:00Z\""`,
				Comparison: soEqual,
			},
			ValueTest: types.Test{
				InputValue: "time.Unix(0, 0).UTC()",
				ExpectValue: "time.Unix(0, 0).UTC()",
				Comparison: soResemble,
			},
		},
	},
}

