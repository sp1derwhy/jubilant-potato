package types

type GoStruct struct {
	// Record struct name
	Name string
	// Record struct attributes and responding json name to map
	Members map[string]string
}

var TransDict  = map[string]string {
	"int":"i64",
	"int8":"i8",
	"int16":"i16",
	"int32":"i32",
	"int64":"i64",
	"uint":"u64",
	"uint8":"u8",
	"uint16":"u16",
	"uint32":"u32",
	"uint64":"u64",
	"string":"String",
	"float32":"f32",
	"float64":"f64",
	"bool":"bool",
}