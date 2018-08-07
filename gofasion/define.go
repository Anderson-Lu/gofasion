package gofasion

type IFasion interface {
	Get(key string) *IFasion

	ValueStr() string

	ValueInt() int
	// ValueInt16() int16
	ValueInt32() int32
	ValueInt64() int64

	// ValueUInt() uint
	// ValueUInt8() uint8
	// ValueUInt16() uint16
	// ValueUInt32() uint32
	// ValueUInt64() uint64

	ValueBool() bool

	Array() []*Fasion

	Value(interface{}) error
}
