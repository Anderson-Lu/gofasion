package gofasion

type IFasion interface {
	Get(key string) *IFasion
	GetFromPath(dir string) *IFasion

	IsValidJson() bool

	ValueStr() string
	ValueInt() int
	ValueInt16() int16
	ValueInt32() int32
	ValueInt64() int64
	ValueFloat32() float32
	ValueFloat64() float64
	ValueBool() bool

	Array() []*Fasion
	ArrayForEach(func(int, *Fasion))

	Keys() []string
	HasKey(key string) bool

	Value(interface{}) error
}
