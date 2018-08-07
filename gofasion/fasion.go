package gofasion

type Fasion struct {
	rawJson string
	errInfo error
	current interface{}
}

func NewFasion(rawJson string) *Fasion {
	return &Fasion{
		rawJson: rawJson,
	}
}

func NewFasionFromBytes(rawJson []byte) *Fasion {
	return &Fasion{
		rawJson: string(rawJson),
	}
}

func (self *Fasion) Get(key string) *Fasion {
	curMap, err := self.parseJson()
	if err != nil {
		self.errInfo = err
		self.current = nil
	}
	if v, ok := curMap[key]; ok {
		rawJson, err := self.toJson(v)
		if err == nil {
			return NewFasion(rawJson)
		}
	}
	return self
}
