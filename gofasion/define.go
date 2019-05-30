package gofasion

/*

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

	ValueDefaultStr(string) string
	ValueDefaultInt(int) int
	ValueDefaultInt16(int16) int16
	ValueDefaultInt32(int32) int32
	ValueDefaultInt64(int64) int64
	ValueDefaultFloat32(float32) float32
	ValueDefaultFloat64(float64) float64
	ValueDefaultBool(bool) bool

	Array() []*Fasion
	ArrayForEach(func(int, *Fasion))

	Keys() []string
	HasKey(key string) bool

	Value(interface{}) error

	//for v1.3 or later version, support check if key exists
	ValStr() (bool, string)
	ValInt64() (bool, int64)
	ValInt32() (bool, int32)
	ValInt16() (bool, int16)
	ValInt() (bool, int)
	ValFloat32() (bool,float32)
	ValFloat32N(int) (bool,float32)
	ValFloat64() (bool,float64)
	ValFloat64N(int) (bool,float64)
	ValBool() (bool,bool)
*/
